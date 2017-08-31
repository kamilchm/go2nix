package go2nix

type ImportPath string

type FetchType int

func (f FetchType) String() string {
	names := map[FetchType]string{
		Mercurial:  "hg",
		Git:        "git",
		Subversion: "svn",
		Bazaar:     "bzr",
	}
	return names[f]
}

const (
	Git FetchType = iota
	Mercurial
	Bazaar
	Subversion
)

type PkgSource struct {
	Type     FetchType
	Url      string
	Revision string
	Sha256   string
}

type GoPackage struct {
	Name        ImportPath
	Source      *PkgSource
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

type PackageInferrer interface {
	Infer(GoPackage) (GoPackage, error)
}
