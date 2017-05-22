package go2nix

import "fmt"

type ImportPath string

type GoPackage struct {
	Name        ImportPath
	Source      string
	Revision    string
	Version     string
	SubPackages []string
}

type NixPackage struct {
	GoPackage
	Deps []GoPackage
}

type DepSolver interface {
	Dependencies(GoPackage, string) ([]GoPackage, error)
}

type PackageLoader interface {
	Package(dir string) (GoPackage, error)
}

func WriteDepsNix([]GoPackage) error {
	return fmt.Errorf("Not implemented")
}

func WriteDefaultNix(NixPackage) error {
	return fmt.Errorf("Not implemented")
}
