package main

import (
	"fmt"
	"go/build"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
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
	ImportPath string
	PkgDir     string
}

func save(pkgName, goPath, nixFile string, depsFile string, testImports bool, buildTags []string) error {

	pkg, err := NewPackage(pkgName, goPath)
	if err != nil {
		return err
	}

	deps, err := findDeps(pkgName, goPath, testImports, buildTags)
	if err != nil {
		return err
	}

	pkgsSoFar := make(map[string]bool)
	var depsPkgs []*NixDependency
	for _, dep := range deps {
		p, err := NewPackage(dep, goPath)
		if err != nil {
			return fmt.Errorf("Can't create package for: %v", dep)
		}
		if p == nil || p.VcsRepo == pkg.VcsRepo {
			continue
		}
		if !pkgsSoFar[p.ImportPath] {
			depsPkgs = append(depsPkgs, &NixDependency{
				GoPackagePath: p.ImportPath,
				Fetch: &FetchGit{
					Type:   "git",
					Url:    p.VcsRepo,
					Rev:    p.Revision,
					Sha256: p.Hash,
				},
			})
			pkgsSoFar[p.ImportPath] = true
		}
	}

	if err := saveDeps(depsPkgs, depsFile); err != nil {
		return fmt.Errorf("Error while saving %v: %v", depsFile, err)
	}

	pkgDef := struct {
		Pkg       *GoPackage
		BuildTags string
		Version   string
	}{pkg, strings.Join(buildTags, ","), version}

	if err = writeFromTemplate(nixFile, "default.nix", pkgDef); err != nil {
		return fmt.Errorf("Error while writing %v: %v", nixFile, err)
	}

	return nil
}

func NewPackage(importPath string, goPath string) (*GoPackage, error) {
	fullPath, err := goPackageDir(importPath, goPath)
	if err != nil {
		return nil, err
	}

	repoRoot, err := repoRoot(fullPath)
	if err != nil {
		return nil, fmt.Errorf("Cannot find repo root for %v: %v",
			fullPath, err)
	}

	pkgRoot, _ := trimGopath(goPath, repoRoot)

	if strings.Contains(pkgRoot, "/vendor/") {
		return nil, nil
	}

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

func goPackageDir(importPath, goPath string) (string, error) {
	for _, goPathDir := range filepath.SplitList(goPath) {
		expectedPath := filepath.Clean(goPathDir + "/src/" + importPath)
		if _, err := os.Stat(expectedPath); err == nil {
			return expectedPath, nil
		}
	}

	return "", fmt.Errorf("Cannot find package %v dir in GOPATH: %v",
		importPath, goPath)
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

	if strings.Contains(c.dir, "/vendor") {
		return nil
	}

	imports := pkg.Imports
	if testImports {
		imports = append(imports, pkg.TestImports...)
	}

	for _, imp := range imports {
		if !c.soFar[imp] {
			topDir := c.dir

			repoRoot, _ := repoRoot(c.ctx.GOPATH + "/src/" + imp)
			if f, err := os.Stat(repoRoot + "/vendor"); err == nil && f.IsDir() {
				c.dir = repoRoot + "/vendor"
			}

			if err := c.find(imp, testImports); err != nil {
				return err
			}
			c.dir = topDir
		}
	}

	return nil
}
