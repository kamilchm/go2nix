// All go2nix CLI related stuff
package main

//go:generate go-bindata -o assets.go templates/

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

// go2nix entry-point
func main() {
	app := kingpin.New("go2nix", "Nix derivations for Go packages")

	saveCmd := app.Command("save", "Saves dependencies for cwd and current GOPATH")
	testImports := saveCmd.Flag("test-imports", "Include test imports.").Short('t').Bool()

	buildTags := saveCmd.Flag("tags", "the dependencies will be generated with the specified build tags").String()

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case saveCmd.FullCommand():
		goPath := os.Getenv("GOPATH")
		if goPath == "" {
			log.Fatal("No GOPATH set, can't find dependencies")
		}
		currPkg, err := currentPackage(goPath)
		if err != nil {
			log.Fatal(err)
		}
		buildTagsList := strings.Split(*buildTags, ",")
		if err := save(currPkg, goPath, *testImports, buildTagsList); err != nil {
			log.Fatal(err)
		}
	}
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
