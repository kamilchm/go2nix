// All go2nix CLI related stuff
package main

//go:generate go-bindata -o assets.go templates/

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jawher/mow.cli"
)

// go2nix entry-point
func main() {
	go2nix := cli.App("go2nix", "Nix derivations for Go packages")

	go2nix.Command("save", "Saves dependecies for cwd within GOPATH", func(cmd *cli.Cmd) {
		outputFile := cmd.StringOpt("o output", "default.nix",
			"Write the resulting nix file to the named output file")
		depsFile := cmd.StringOpt("d deps-file", "deps.json",
			"Write the resulting dependencies file to the named output file")
		testImports := cmd.BoolOpt("t test-imports", false,
			"Include test imports.")
		buildTags := cmd.StringOpt("tags", "",
			"the dependencies will be generated with the specified build tags")

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

	go2nix.Command("merge", "Takes deps from one file and tries to merge it into another one", func(cmd *cli.Cmd) {
		srcFile := cmd.StringArg("SRC", "",
			"File with dependencies to merge into DST")
		dstFile := cmd.StringArg("DST", "",
			"Where to merge dependencies?")

		cmd.Action = func() {
			if err := MergeDeps(*srcFile, *dstFile); err != nil {
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

	if !strings.HasPrefix(currDir, goPath+"/src/") {
		return "", fmt.Errorf("Current dir %v is outside of GOPATH(%v). "+
			"Can't get current package name", currDir, goPath)
	}

	return strings.TrimPrefix(currDir, goPath+"/src/"), nil
}
