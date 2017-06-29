package vcs

import (
	"path/filepath"
	"strings"

	"github.com/Masterminds/vcs"
	"github.com/kamilchm/go2nix"
	"github.com/kamilchm/go2nix/gopath"
)

type LocalSourceSolver struct{}

func (s *LocalSourceSolver) Source(pkg go2nix.GoPackage) (*go2nix.PkgSource, error) {
	for _, goPathDir := range filepath.SplitList(gopath.GoPath()) {
		pkgPath := []string{goPathDir, "src"}
		pkgPath = append(pkgPath, strings.Split(string(pkg.Name), "/")...)
		localSrc, err := filepath.Abs(filepath.Join(pkgPath...))
		if err != nil {
			continue
		}

		repo, err := vcs.NewRepo("", localSrc)
		if err != nil {
			continue
		}

		src := &go2nix.PkgSource{
			Type: pkg.Source.Type,
		}

		if repo.IsDirty() {
			src.Url = repo.LocalPath()
		} else {
			src.Url = repo.Remote()
		}

		ref, err := repo.Current()
		if err != nil && pkg.Source != nil {
			src.Revision = pkg.Source.Revision
		}

		commit, err := repo.CommitInfo(ref)
		src.Revision = commit.Commit

		return src, nil
	}
	return pkg.Source, nil
}
