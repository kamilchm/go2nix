package main

// go:generate go-bindata .

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type NixDepenency struct {
	GoPackagePath string      `json:"goPackagePath"`
	Fetch         interface{} `json:"fetch"`
}

type FetchGit struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

func saveDeps(deps []*NixDepenency, depsFilename string) error {
	depsFile, err := os.Create(depsFilename)
	if err != nil {
		return err
	}
	defer depsFile.Close()

	j, jerr := json.MarshalIndent(deps, "", "  ")
	if jerr != nil {
		fmt.Println("jerr:", jerr.Error())
	}

	_, werr := depsFile.Write(j)
	if werr != nil {
		fmt.Println("werr:", werr.Error())
	}

	return nil
}

func loadDeps(depsFilename string) ([]*NixDepenency, error) {
	depsFile, err := ioutil.ReadFile(depsFilename)
	if err != nil {
		return nil, err
	}
	var deps []*NixDepenency
	err = json.Unmarshal(depsFile, &deps)
	return deps, err
}

func writeFromTemplate(filename string, data interface{}) error {
	templateData, err := Asset("templates/default.nix")
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
