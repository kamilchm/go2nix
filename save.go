package main

// go:generate go-bindata .

import (
	"encoding/json"
	"fmt"
	"go/build"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/vcs"
)

var (
	excludePath = "example"
)

type context struct {
	soFar map[string]bool
	ctx   build.Context
	dir   string
}

type NixDepenency struct {
	GoPackagePath string      `json:"goPackagePath"`
	Fetch         interface{} `json:"fetch"`
}

type FetchGit struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

type GoPackage struct {
	Name       string
	ImportPath string
	VcsRepo    string
	VcsCommand string
	Revision   string
	Hash       string
	UpdateDate time.Time
	Vendored   []*VendoredPackage
}

type VendoredPackage struct {
	Name       string
	ImportPath string
	PkgDir     string
}

func save(pkgName, goPath, nixFile string, testImports bool, buildTags []string) error {

	pkg, err := NewPackage(pkgName, goPath)
	if err != nil {
		return err
	}

	deps, err := findDeps(pkgName, goPath, testImports, buildTags)
	if err != nil {
		return err
	}

	pkgsSoFar := make(map[string]bool)
	var depsPkgs []*NixDepenency
	for _, dep := range deps {
		p, err := NewPackage(dep, goPath)
		if err != nil {
			return fmt.Errorf("Can't create package for: %v", dep)
		}
		if !pkgsSoFar[p.ImportPath] {
			depsPkgs = append(depsPkgs, &NixDepenency{
				GoPackagePath: p.ImportPath,
				Fetch: FetchGit{
					Type:   "git",
					Url:    p.VcsRepo,
					Rev:    p.Revision,
					Sha256: p.Hash,
				},
			})
			pkgsSoFar[p.ImportPath] = true
		}
	}

	if err := saveDeps(depsPkgs); err != nil {
		return err
	}

	pkgDef := struct {
		Pkg       *GoPackage
		BuildTags string
	}{pkg, strings.Join(buildTags, ",")}

	if err = writeFromTemplate(nixFile, pkgDef); err != nil {
		return err
	}

	return nil
}

func saveDeps(deps []*NixDepenency) error {
	depsFile, err := os.Create("deps.json")
	if err != nil {
		return err
	}
	defer depsFile.Close()

	j, jerr := json.MarshalIndent(deps, "", "  ")
	if jerr != nil {
		fmt.Println("jerr:", jerr.Error())
	}

	_, werr := depsFile.Write(j)
	if werr != nil {
		fmt.Println("werr:", werr.Error())
	}

	return nil
}

func writeFromTemplate(filename string, data interface{}) error {
	templateData, err := Asset("templates/default.nix")
	if err != nil {
		return err
	}

	t, err := template.New(filename).Delims("[[", "]]").Parse(string(templateData))
	if err != nil {
		return err
	}

	target, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer target.Close()

	t.Execute(target, data)
	return nil
}

func NewPackage(importPath string, goPath string) (*GoPackage, error) {
	fullPath := goPath + "/src/" + importPath

	repoRoot, err := repoRoot(fullPath)
	if err != nil {
		return nil, fmt.Errorf("Cannot find repo root for %v: %v",
			fullPath, err)
	}

	pkgRoot := strings.TrimPrefix(repoRoot, goPath+"/src/")

	repo, err := vcs.NewRepo("", repoRoot)
	if err != nil {
		return nil, fmt.Errorf("Error while creating repo for %v: %v",
			repoRoot, err)
	}
	revision, err := repo.Version()
	if err != nil {
		return nil, err
	}
	updateDate, err := repo.Date()
	if err != nil {
		return nil, err
	}

	return &GoPackage{
		Name:       nixName(pkgRoot),
		ImportPath: pkgRoot,
		VcsRepo:    repo.Remote(),
		VcsCommand: string(repo.Vcs()),
		Revision:   revision,
		Hash:       calculateHash("file://"+repo.LocalPath(), string(repo.Vcs())),
		UpdateDate: updateDate,
	}, nil
}

func repoRoot(pth string) (string, error) {
	_, err := vcs.DetectVcsFromFS(pth)
	if err == vcs.ErrCannotDetectVCS {
		if pth == "/" {
			return pth, err
		}
		return repoRoot(path.Dir(pth))
	}
	if err != nil {
		return pth, fmt.Errorf("Error while detecting repo root for %v: %v",
			pth, err)
	}
	return pth, nil
}

func nixName(goImportPath string) string {
	parts := strings.Split(goImportPath, "/")
	return strings.Replace(parts[len(parts)-1], ".", "-", -1)
}

func findDeps(name, gopath string, testImports bool, buildTags []string) ([]string, error) {
	ctx := build.Default
	ctx.BuildTags = buildTags

	if gopath != "" {
		ctx.GOPATH = gopath
	}
	c := &context{
		soFar: make(map[string]bool),
		ctx:   ctx,
		dir:   gopath + "/src/" + name,
	}
	if err := c.find(name, testImports); err != nil {
		return nil, err
	}
	var deps []string
	for p := range c.soFar {
		if p != name {
			deps = append(deps, p)
		}
	}
	sort.Strings(deps)
	return deps, nil
}

func (c *context) find(name string, testImports bool) error {
	if name == "C" {
		return nil
	}
	pkg, err := c.ctx.Import(name, c.dir, 0)
	if err != nil {
		return err
	}
	if pkg.Goroot {
		return nil
	}

	if name != "." {
		c.soFar[pkg.ImportPath] = true
	}
	imports := pkg.Imports
	if testImports {
		imports = append(imports, pkg.TestImports...)
	}
	for _, imp := range imports {
		if !c.soFar[imp] {
			if err := c.find(imp, testImports); err != nil {
				return err
			}
		}
	}
	return nil
}
