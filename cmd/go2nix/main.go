package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamilchm/go2nix"
	"github.com/kamilchm/go2nix/dep"
	"github.com/kamilchm/go2nix/gopath"
	"github.com/kamilchm/go2nix/govcs"
	"github.com/kamilchm/go2nix/nixhash"
	"github.com/kamilchm/go2nix/vcs"
)

func main() {
	loader := go2nix.PackageLoader(&gopath.PackageLoader{})
	depSolver := go2nix.DepSolver(&dep.Solver{})
	packageInferrers := []go2nix.PackageInferrer{
		&govcs.RemoteRepoInferrer{},
		&vcs.LocalSourceInferrer{},
		&nixhash.HashInferrer{},
	}

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Couldn't get current dir: %v", err)
	}

	goPkg, err := loader.Package(currDir)
	if err != nil {
		log.Fatalf("Couldn't load Go package: %v", err)
	}
	for _, inferrer := range packageInferrers {
		goPkg, err = inferrer.Infer(goPkg)
		if err != nil {
			log.Fatalf("Error while trying to infer Go package '%s': %v", goPkg.Name, err)
		}
	}

	deps, err := depSolver.Dependencies(goPkg, gopath.GoPath())
	if err != nil {
		log.Fatalf("Couldn't solve package dependencies: %v", err)
	}

	depsInferrers := []go2nix.PackageInferrer{
		&govcs.RemoteRepoInferrer{},
		&nixhash.HashInferrer{},
	}

	for i := range deps {
		// TODO: infer packages concurrently
		for _, inferrer := range depsInferrers {
			pkg, err := inferrer.Infer(deps[i])
			if err != nil {
				log.Fatalf("Error while trying to infer Go package '%s': %v", deps[i].Name, err)
			}
			deps[i] = pkg
		}
	}

	if err = go2nix.WriteDepsNix(deps); err != nil {
		log.Fatalf("Error while trying to write deps.nix: %v", err)
	}

	nixPkg := go2nix.NixPackage{GoPackage: goPkg}
	// TODO: filename, err :=
	if err = go2nix.WriteDefaultNix(nixPkg); err != nil {
		log.Fatalf("Error while trying to write default.nix: %v", err)
	}

	fmt.Printf("New Nix derivation written to default.nix")
	if len(deps) > 0 {
		fmt.Printf(" with deps in deps.nix\n")
	} else {
		fmt.Printf("\n")
	}
}
