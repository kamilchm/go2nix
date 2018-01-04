package dep

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"

	"github.com/kamilchm/go2nix"
)

const (
	lockFile = "Gopkg.lock"
)

type depLock struct {
	Projects []depProject `toml:"projects"`
}

type depProject struct {
	Name     string   `toml:"name"`
	Source   string   `toml:"source"`
	Revision string   `toml:"revision"`
	Packages []string `toml:"packages"`
}

type Solver struct{}

func (s *Solver) Dependencies(pkg go2nix.GoPackage, goPath string) ([]go2nix.GoPackage, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Couldn't get working dir: %v", err)
	}
	lock, err := readLock(wd)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read Dep lock file: %v", err)
	}

	return convertDeps(lock), nil
}

func readLock(dir string) (*depLock, error) {
	lPath := filepath.Join(dir, lockFile)
	lTree, err := toml.LoadFile(lPath)
	if err != nil {
		return nil, fmt.Errorf("Couldn't load %s: %v", lPath, err)
	}

	lock := &depLock{}
	err = lTree.Unmarshal(lock)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse dep lock tree: %v", err)
	}

	return lock, nil
}

func convertDeps(lock *depLock) (deps []go2nix.GoPackage) {
	for _, p := range lock.Projects {
		pkg := go2nix.GoPackage{
			Name:        go2nix.ImportPath(p.Name),
			Source:      &go2nix.PkgSource{Url: p.Source, Revision: p.Revision},
			SubPackages: p.Packages,
		}

		deps = append(deps, pkg)
	}

	return deps
}
