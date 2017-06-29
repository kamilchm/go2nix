package go2nix

//go:generate go-bindata -o assets.go -pkg go2nix templates/

import (
	"os"
	"strings"
	"text/template"
)

func WriteDepsNix(packages []GoPackage) error {
	return writeFromTemplate("deps.nix", "deps.nix", struct {
		Deps    []GoPackage
		Version string
	}{packages, version})
}

func WriteDefaultNix(pkg NixPackage) error {
	return writeFromTemplate("default.nix", "default.nix", struct {
		Pkg       NixPackage
		Version   string
		BuildTags string
	}{pkg, version, ""})
}

func writeFromTemplate(filename, templFile string, data interface{}) error {
	templateData, err := Asset("templates/" + templFile)
	if err != nil {
		return err
	}

	t, err := template.New(filename).Delims("[[", "]]").Parse(string(templateData))
	if err != nil {
		return err
	}

	target, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer target.Close()

	t.Execute(target, data)
	return nil
}

func nixName(goImportPath string) string {
	parts := strings.Split(goImportPath, "/")
	return strings.Replace(parts[len(parts)-1], ".", "-", -1)
}
