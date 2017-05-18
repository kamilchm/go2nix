package dep

import (
	"fmt"

	"github.com/golang/dep"
	"github.com/golang/dep/gps"

	"github.com/kamilchm/go2nix"
)

type Solver struct{}

func (s *Solver) Dependencies(pkg go2nix.GoPackage) ([]go2nix.GoPackage, error) {
	ctx, err := dep.NewContext()
	if err != nil {
		return nil, fmt.Errorf("Couldn't create Go dep context: %v", err)
	}

	project, err := ctx.LoadProject("")
	if err != nil {
		return nil, fmt.Errorf("Couldn't load Go project: %v", err)
	}

	return convertDeps(project.Lock.Projects()), nil
}

func convertDeps(gpsDeps []gps.LockedProject) (deps []go2nix.GoPackage) {
	for _, p := range gpsDeps {
		pkg := go2nix.GoPackage{
			Name:    go2nix.ImportPath(p.Ident().ProjectRoot),
			Version: p.Version().String(),
		}

		if v, ok := p.Version().(gps.PairedVersion); ok {
			pkg.Revision = v.Underlying().String()
		}

		deps = append(deps, pkg)
	}

	return deps
}
