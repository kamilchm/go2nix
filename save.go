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
	govcs "golang.org/x/tools/go/vcs"
)

var (
	inSet = struct{}{}
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

	pkg, err := NewPacakge(pkgName, goPath)
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

func NewPacakge(importPath string, goPath string) (*GoPackage, error) {
	rr, err := govcs.RepoRootForImportPath(importPath, false)
	if err != nil {
		return nil, fmt.Errorf("Can't get repo for import path: %v", err)
	}

	repo, err := mmvcs.NewRepo("", goPath+"/src/"+rr.Root)
	if err != nil {
		return nil, fmt.Errorf("Error while creating repo for (%v, %v): %v",
			rr.Repo, goPath+"/src/"+rr.Root, err)
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
		Name:       nixName(rr.Root),
		ImportPath: rr.Root,
		VcsRepo:    rr.Repo,
		VcsCommand: rr.VCS.Cmd,
		Revision:   revision,
		Hash:       calculateHash("file://"+repo.LocalPath(), rr.VCS.Cmd),
		UpdateDate: updateDate,
	}, nil
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

	log.Println("Anaalyzing", pkg.ImportPath)

	pkgBase := goPath + "/src/" + pkg.ImportPath
	filepath.Walk(pkgBase,
		func(pth string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(pth, ".go") {
				return nil
			}

			subDir, err := build.ImportDir(path.Dir(pth), build.AllowBinary)
			if err != nil {
				return err
			}

			subPkg, err := NewPacakge(subDir.ImportPath, goPath)
			if err != nil {
				return err
			}
			depsMap[subPkg.ImportPath] = subPkg

			allImports := append(subDir.Imports, subDir.TestImports...)

			for _, imp := range allImports {
				goroot, err := isGoroot(imp)
				if err != nil {
					return err
				}
				if goroot {
					continue
				}

				depPkg, err := NewPacakge(imp, goPath)
				if err != nil {
					return err
				}

				depsMap, err = depsRecursive(depPkg, goPath, depsMap)
				if err != nil {
					return err
				}
			}

			return nil
		},
	)
	return depsMap, nil
}

func isGoroot(importPath string) (bool, error) {
	pkg, err := build.Import(importPath, "", build.AllowBinary)
	if err != nil {
		return false, err
	}
	if pkg.Goroot {
		return true, nil
	}

	return false, nil
}
