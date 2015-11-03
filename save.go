package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

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
}

func save(pkg, goPath string) {
	packages := discoverDeps(pkg, goPath)

	if err := writeFromTemplate("deps.nix", packages); err != nil {
		log.Fatal(err)
	}

	pkgRoot, err := govcs.RepoRootForImportPath(pkg, false)
	if err != nil {
		fmt.Errorf("Can't get repo for import path", err)
	}

	p, err := Init(*pkgRoot, goPath)
	if err != nil {
		fmt.Errorf("Can't initialize package data", err)
	}

	if err = writeFromTemplate("default.nix", p); err != nil {
		log.Fatal(err)
	}

}

func writeFromTemplate(filename string, data interface{}) error {
	templateData, err := Asset("templates/" + filename)
	if err != nil {
		return err
	}

	t, err := template.New(filename).Parse(string(templateData))
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
	for root, _ := range roots {
		if root.Root == pkg {
			continue
		}

		p, err := Init(root, goPath)
		if err != nil {
			log.Fatal(err)
		}

		packages = append(packages, p)
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

	return &GoPackage{
		Name:       nixName(rr.Root),
		ImportPath: rr.Root,
		VcsRepo:    rr.Repo,
		VcsCommand: rr.VCS.Cmd,
		Revision:   revision,
		Hash:       calculateHash("file://"+repo.LocalPath(), rr.VCS.Cmd),
	}, nil
}

func nixName(goImportPath string) string {
	parts := strings.Split(goImportPath, "/")
	return parts[len(parts)-1]
}

func discoverRootPackages(goPath string) map[govcs.RepoRoot]struct{} {
	pkgsBase := goPath + "/src/"
	baseLen := len(pkgsBase)
	packages := make(map[string]struct{})
	roots := make(map[govcs.RepoRoot]struct{})
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

			for repo, _ := range roots {
				if strings.HasPrefix(importPath, repo.Root) {
					return nil
				}
			}

			rr, err := govcs.RepoRootForImportPath(importPath, false)
			if err != nil {
				fmt.Errorf("Can't get repo for import path", err)
			} else {
				roots[*rr] = inSet
			}

			return nil
		},
	)
	return roots
}
