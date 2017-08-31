package vcs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/Masterminds/vcs"

	"github.com/kamilchm/go2nix"
	"github.com/kamilchm/go2nix/gopath"
)

// LocalSourceInferrer tries to find given package in GOPATH.
// It then use Masterminds vcs package to return pakage source repo.
type LocalSourceInferrer struct{}

func (s *LocalSourceInferrer) Infer(pkg go2nix.GoPackage) (go2nix.GoPackage, error) {
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
			Url:  repo.Remote(),
		}

		ref, err := repo.Current()
		if err == nil {
			version, err := semver.NewVersion(ref)
			if err == nil {
				pkg.Version = version.String()
			}
		}

		if pkg.Source != nil {
			src.Revision = pkg.Source.Revision
		}

		commit, err := repo.CommitInfo(ref)
		if err == nil {
			src.Revision = commit.Commit
		}

		pkg.Source = src

		return pkg, nil
	}
	return pkg, fmt.Errorf("Couldn't find package source in GOPATH=%v", gopath.GoPath())
}
