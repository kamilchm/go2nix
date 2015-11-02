// All go2nix CLI related stuff
package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// go2nix entry-point
func main() {
	app := kingpin.New("go2nix", "Nix derivations for Go packages")

	getCmd := app.Command("get", "") // TODO: prefetch alias?
	getPackage := getCmd.Arg("package", "").Required().String()

	//saveCmd := app.Command("save", "Saves dependencies for current GOPATH")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case getCmd.FullCommand():
		get(*getPackage)
		//case saveCmd.FullCommand():
		//	goPath := os.Getenv("GOPATH")
		//	if goPath == "" {
		//		log.Fatal("No GOPATH set, can't save dependencies")
		//	}
		//	cwd, _ := os.Getwd()
		//	save("", cwd)
	}
}
