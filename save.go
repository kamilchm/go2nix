package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/sdboyer/gps"
)

type NaiveAnalyzer struct{}

// DeriveManifestAndLock is called when the solver needs manifest/lock data
// for a particular dependency project (identified by the gps.ProjectRoot
// parameter) at a particular version. That version will be checked out in a
// directory rooted at path.
func (a NaiveAnalyzer) DeriveManifestAndLock(path string, n gps.ProjectRoot) (gps.Manifest, gps.Lock, error) {
	return nil, nil, nil
}

// Reports the name and version of the analyzer. This is used internally as part
// of gps' hashing memoization scheme.
func (a NaiveAnalyzer) Info() (string, *semver.Version) {
	v, _ := semver.NewVersion(version)
	return "go2nix", v
}

func save(pkgName, goPath, nixFile string, depsFile string, testImports bool, buildTags []string) error {
	// Assume the current directory is correctly placed on a GOPATH, and that it's the
	// root of the project.
	root, _ := os.Getwd()
	srcprefix := filepath.Join(build.Default.GOPATH, "src") + string(filepath.Separator)
	importroot := filepath.ToSlash(strings.TrimPrefix(root, srcprefix))

	// Set up params, including tracing
	params := gps.SolveParameters{
		RootDir:     root,
		Trace:       true,
		TraceLogger: log.New(os.Stdout, "go2nix: ", 0),
	}
	// Perform static analysis on the current project to find all of its imports.
	params.RootPackageTree, _ = gps.ListPackages(root, importroot)

	// Set up a SourceManager. This manages interaction with sources (repositories).
	tempdir, _ := ioutil.TempDir("", "go2nix-cache")
	sourcemgr, _ := gps.NewSourceManager(NaiveAnalyzer{}, filepath.Join(tempdir))
	defer sourcemgr.Release()

	// Prep and run the solver
	solver, _ := gps.Prepare(params, sourcemgr)
	solution, err := solver.Solve()
	if err != nil {
		return fmt.Errorf("Couldn't solve package dependencies: %v", err)
	}

	nixDeps := make([]NixDependency, len(solution.Projects()))
	for i, p := range solution.Projects() {
		nixDeps[i] = NixDependency{
			GoPackagePath: string(p.Ident().ProjectRoot),
			Fetch: &FetchGit{
				Type:   "git",
				Url:    p.Ident().Source,
				Rev:    p.Version().String(),
				Sha256: "xxxxxxxxxx",
			},
		}
	}

	if err := saveDeps(nixDeps, depsFile); err != nil {
		return fmt.Errorf("Error while saving %v: %v", depsFile, err)
	}

	if err = writeFromTemplate(nixFile, "default.nix", nixDeps[0]); err != nil {
		return fmt.Errorf("Error while writing %v: %v", nixFile, err)
	}

	return nil
}
