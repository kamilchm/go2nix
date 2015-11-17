package main

import (
	"fmt"
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

func save(pkgName, goPath string) {
	deps := discoverDeps(pkgName, goPath)

	pkgRoot, err := govcs.RepoRootForImportPath(pkgName, false)
	if err != nil {
		fmt.Errorf("Can't get repo for import path", err)
	}

	pkg, err := Init(*pkgRoot, goPath)
	if err != nil {
		fmt.Errorf("Can't initialize package data", err)
	}

	pkgDef := struct {
		Pkg  *GoPackage
		Deps []*GoPackage
	}{pkg, deps}

	if err = writeFromTemplate("default.nix", pkgDef); err != nil {
		log.Fatal(err)
	}

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

func discoverDeps(pkg, goPath string) (packages []*GoPackage) {
	roots := discoverRootPackages(goPath)
	for _, root := range roots {
		if root.ImportPath == pkg {
			continue
		}

		packages = append(packages, root)
	}

	return packages
}

func Init(rr govcs.RepoRoot, goPath string) (*GoPackage, error) {
	repo, err := mmvcs.NewRepo(rr.Repo, goPath+"/src/"+rr.Root)
	if err != nil {
		return nil, err
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

func discoverRootPackages(goPath string) map[string]*GoPackage {
	pkgsBase := goPath + "/src/"
	baseLen := len(pkgsBase)
	packages := make(map[string]struct{})
	roots := make(map[string]*GoPackage)
	filepath.Walk(pkgsBase,
		func(pth string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(pth, ".go") {
				return nil
			}

			pkgDir := path.Dir(pth)

			if _, added := packages[pkgDir]; added {
				return nil
			}

			packages[pkgDir] = inSet

			importPath := pkgDir[baseLen:]

			for _, root := range roots {
				if strings.HasPrefix(importPath, root.ImportPath) {
					if strings.Contains(importPath, "vendor/") {
						root.Vendored = []*VendoredPackage{&VendoredPackage{
							Name:       nixName(importPath),
							ImportPath: strings.Split(importPath, "vendor/")[1],
							PkgDir:     strings.Split(importPath, root.ImportPath+"/")[1],
						}}
					}
					return nil
				}
			}

			rr, err := govcs.RepoRootForImportPath(importPath, false)
			if err != nil {
				fmt.Errorf("Can't get repo for import path", err)
			}

			root, err := Init(*rr, goPath)
			if err != nil {
				log.Println("Can't create package metadata: ", err)
			} else {
				roots[importPath] = root
			}

			return nil
		},
	)
	return roots
}
