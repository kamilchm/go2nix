package main

// go:generate go-bindata .

import (
	"os"
	"strings"
	"text/template"
)

type NixDependency struct {
	GoPackagePath string    `json:"goPackagePath,omitempty"`
	Fetch         *FetchGit `json:"fetch,omitempty"`
	IncludeFile   string    `json:"include,omitempty"`
	Packages      []string  `json:"packages,omitempty"`
}

type FetchGit struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

func saveDeps(deps []*NixDependency, depsFilename string) error {
	return writeFromTemplate(depsFilename, "deps.nix", struct {
		Deps    []*NixDependency
		Version string
	}{deps, version})
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
