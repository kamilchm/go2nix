package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kamilchm/go2nix"
	"github.com/kamilchm/go2nix/dep"
	"github.com/kamilchm/go2nix/gopath"
)

func main() {
	loader := go2nix.PackageLoader(&gopath.PackageLoader{})
	solver := go2nix.DepSolver(&dep.Solver{})

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Couldn't get current dir: %v", err)
	}

	goPkg, err := loader.Package(currDir)
	if err != nil {
		log.Fatalf("Couldn't load Go package: %v", err)
	}

	deps, err := solver.Dependencies(goPkg, gopath.GoPath())
	if err != nil {
		log.Fatalf("Couldn't solve package dependencies: %v", err)
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
