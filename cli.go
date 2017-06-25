// All go2nix CLI related stuff
package main

//go:generate go-bindata -o assets.go templates/

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jawher/mow.cli"
)

// go2nix entry-point
func main() {
	go2nix := cli.App("go2nix", "Nix derivations for Go packages")
	go2nix.Version("v version", "go2nix "+version)

	go2nix.Command("save", "Saves dependencies for cwd within GOPATH", func(cmd *cli.Cmd) {
		outputFile := cmd.StringOpt("o output", "default.nix",
			"Write the resulting nix file to the named output file")
		depsFile := cmd.StringOpt("d deps-file", "deps.nix",
			"Write the resulting dependencies file to the named output file")
		testImports := cmd.BoolOpt("t test-imports", false,
			"Include test imports")
		buildTags := cmd.StringOpt("tags", "",
			"The dependencies will be generated with the specified build tags")

		cmd.Action = func() {
			goPath := os.Getenv("GOPATH")
			if goPath == "" {
				log.Fatal("No GOPATH set, can't find dependencies")
			}
			currPkg, err := currentPackage(goPath)
			if err != nil {
				log.Fatal(err)
			}
			buildTagsList := strings.Split(*buildTags, ",")
			if err := save(currPkg, goPath, *outputFile, *depsFile,
				*testImports, buildTagsList); err != nil {
				log.Fatal(err)
			}
		}
	})

	go2nix.Run(os.Args)
}

func currentPackage(goPath string) (string, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if packagePath, found := trimGopath(goPath, currDir); found {
		return packagePath, nil
	}

	return "", fmt.Errorf("Current dir %v is outside of GOPATH(%v). "+
		"Can't get current package name", currDir, goPath)
}

func trimGopath(goPath string, packagePath string) (trimmedPath string, trimmed bool) {
	for _, goPathDir := range filepath.SplitList(goPath) {
		goPathSrc, err := filepath.Abs(filepath.Join(goPathDir, "/src"))
		if err != nil {
			continue
		}
		if strings.HasPrefix(packagePath, goPathSrc) {
			return strings.TrimPrefix(packagePath, goPathSrc+"/"), true
		}
	}

	return packagePath, false
}
