package govcs

import (
	"fmt"

	"golang.org/x/tools/go/vcs"

	"github.com/kamilchm/go2nix"
)

// RemoteSourceInferrer tries to find remote repo location using Go Tools VCS package.
type RemoteRepoInferrer struct{}

func (s *RemoteRepoInferrer) Infer(pkg go2nix.GoPackage) (go2nix.GoPackage, error) {
	repo := string(pkg.Name)
	// The repository url can be different from the import
	// path. See the source attribute in the Gopkg.toml constraint
	// section.
	if pkg.Source != nil && pkg.Source.Url != "" {
		repo = pkg.Source.Url
	}
	repoRoot, err := vcs.RepoRootForImportPath(repo, false)
	if err != nil {
		return pkg, fmt.Errorf("Unknown repo root for '%s': %v", repo, err)
	}

	src := &go2nix.PkgSource{
		Type: fetchType(repoRoot.VCS.Name),
		Url:  repoRoot.Repo,
	}
	if pkg.Source != nil {
		src.Revision = pkg.Source.Revision
	}

	pkg.Source = src

	return pkg, nil
}

// fetchType maps Go vcs name to go2nix FetchType.
// https://github.com/golang/tools/blob/63c6481f3be3d4c29183574fa76516c4e7f54c6e/go/vcs/vcs.go#L55
func fetchType(vcsName string) go2nix.FetchType {
	fetchTypes := map[string]go2nix.FetchType{
		"Mercurial":  go2nix.Mercurial,
		"Git":        go2nix.Git,
		"Subversion": go2nix.Subversion,
		"Bazaar":     go2nix.Bazaar,
	}

	return fetchTypes[vcsName]
}
