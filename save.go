package main

// go:generate go-bindata .

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	mmvcs "github.com/Masterminds/vcs"
)

var (
	excludePath = "example"
)

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

func save(pkgName, goPath string) error {

	pkg, err := NewPackage(pkgName, goPath)
	if err != nil {
		return err
	}

	deps, err := discoverDeps(pkg, goPath)
	if err != nil {
		return err
	}

	for _, dep := range deps {
		log.Println(dep.ImportPath)
	}

	pkgDef := struct {
		Pkg  *GoPackage
		Deps []*GoPackage
	}{pkg, deps}

	if err = writeFromTemplate("default.nix", pkgDef); err != nil {
		return err
	}

	return nil
}

func writeFromTemplate(filename string, data interface{}) error {
	templateData, err := Asset("templates/" + filename)
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

	repo, err := mmvcs.NewRepo("", repoRoot)
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
	_, err := mmvcs.DetectVcsFromFS(pth)
	if err == mmvcs.ErrCannotDetectVCS {
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

func discoverDeps(pkg *GoPackage, goPath string) ([]*GoPackage, error) {
	depsMap := make(map[string]*GoPackage)
	depsMap, err := depsRecursive(pkg, goPath, depsMap)
	if err != nil {
		return nil, err
	}

	rrSet := make(map[string]*GoPackage)
	for _, pkg := range depsMap {
		if _, exist := rrSet[pkg.ImportPath]; !exist {
			rrSet[pkg.ImportPath] = pkg
		}
	}

	packages := make([]*GoPackage, len(rrSet))

	i := 0
	for _, pkg := range rrSet {
		packages[i] = pkg
		i++
	}

	return packages, nil
}

func depsRecursive(pkg *GoPackage, goPath string, depsMap map[string]*GoPackage) (packages map[string]*GoPackage, err error) {
	if _, exist := depsMap[pkg.ImportPath]; exist {
		return depsMap, nil
	}

	log.Println("Analyzing", pkg.ImportPath)

	pkgBase := goPath + "/src/" + pkg.ImportPath

	err = filepath.Walk(pkgBase,
		func(pth string, info os.FileInfo, err error) error {
			if strings.Contains(pth, excludePath) {
				return nil
			}

			if !strings.HasSuffix(pth, ".go") {
				return nil
			}

			subDir, err := build.ImportDir(path.Dir(pth), build.AllowBinary)
			if err != nil {
				// no Go buildable package
				return nil
			}
			provided, err := isProvided(subDir.ImportPath)
			if err != nil {
				return err
			}
			if provided {
				return nil
			}

			if _, exist := depsMap[subDir.ImportPath]; exist {
				return nil
			}

			subPkg, err := NewPackage(subDir.ImportPath, goPath)
			if err != nil {
				return fmt.Errorf("Cannot create package %v: %v", subDir.ImportPath, err)
			}
			depsMap[subDir.ImportPath] = subPkg

			allImports := append(subDir.Imports, subDir.TestImports...)

			for _, imp := range allImports {
				provided, err := isProvided(imp)
				if err != nil {
					return err
				}
				if provided {
					continue
				}

				depPkg, err := NewPackage(imp, goPath)
				if err != nil {
					return err
				}

				//depsMap[imp] = depPkg
				depsMap, err = depsRecursive(depPkg, goPath, depsMap)
				if err != nil {
					return err
				}
			}

			return nil
		},
	)

	if err != nil {
		return depsMap, err
	}
	return depsMap, nil
}

func isProvided(importPath string) (bool, error) {
	if importPath == "C" {
		return true, nil
	}
	if importPath == "." {
		return true, nil
	}
	pkg, err := build.Import(importPath, "", build.AllowBinary)
	if err != nil {
		return false, fmt.Errorf("Cannot import %v: %v", importPath, err)
	}
	if pkg.Goroot {
		return true, nil
	}

	return false, nil
}
