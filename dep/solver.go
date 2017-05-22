package dep

import (
	"fmt"
	"os"

	"github.com/golang/dep"

	"github.com/kamilchm/go2nix"
)

type Solver struct{}

func (s *Solver) Dependencies(pkg go2nix.GoPackage, goPath string) ([]go2nix.GoPackage, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("Couldn't get working dir: %v", err)
	}
	ctx, err := dep.NewContext(wd, []string{"GOPATH=" + goPath}, nil)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create Go dep context: %v", err)
	}

	project, err := ctx.LoadProject("")
	if err != nil {
		return nil, fmt.Errorf("Couldn't load Go project: %v", err)
	}

	return convertDeps(project), nil
}

func convertDeps(depProject *dep.Project) (deps []go2nix.GoPackage) {
	for _, p := range depProject.Lock.Projects() {
		pkg := go2nix.GoPackage{
			Name:    go2nix.ImportPath(p.Ident().ProjectRoot),
			Version: p.Version().String(),
		}

		//if v, ok := p.Version().(gps.PairedVersion); ok {
		//	pkg.Revision = v.Underlying().String()
		//}

		deps = append(deps, pkg)
	}

	return deps
}
