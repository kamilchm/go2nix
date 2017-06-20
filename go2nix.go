package go2nix

import (
	"fmt"
)

type ImportPath string

type FetchType int

const (
	Git FetchType = iota
	Mercurial
	Bazaar
	Subversion
)

type PkgSource struct {
	Type   FetchType
	Url    string
	Sha256 string
}

type GoPackage struct {
	Name        ImportPath
	Source      *PkgSource
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

type SourceSolver interface {
	Source(GoPackage) (*PkgSource, error)
}

func WriteDepsNix(packages []GoPackage) error {
	return fmt.Errorf("Not implemented")
}

func WriteDefaultNix(NixPackage) error {
	return fmt.Errorf("Not implemented")
}
