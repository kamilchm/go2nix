package gopath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kamilchm/go2nix"
)

type PackageLoader struct{}

func (l *PackageLoader) Package(dir string) (go2nix.GoPackage, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return go2nix.GoPackage{}, fmt.Errorf("No GOPATH set, couldn't discover Go package")
	}

	for _, goPathDir := range strings.Split(goPath, ":") {
		goPathSrc, err := filepath.Abs(filepath.Join(goPathDir, "/src"))
		if err != nil {
			continue
		}
		if strings.HasPrefix(dir, goPathSrc) {
			return go2nix.GoPackage{
				Name: go2nix.ImportPath(strings.TrimPrefix(dir, goPathSrc+"/")),
			}, nil
		}
	}

	return go2nix.GoPackage{}, fmt.Errorf("Current dir %v is outside of GOPATH(%v). "+
		"Couldn't get current package name", dir, goPath)
}
