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
	GoPackagePath string   `json:"goPackagePath"`
	Fetch         FetchGit `json:"fetch"`
}

type FetchGit struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

func MergeDeps(srcFile string, dstFile string) error {
	srcDepsList, err := loadDeps(srcFile)
	if err != nil {
		return err
	}
	dstDepsList, err := loadDeps(dstFile)
	if err != nil {
		return err
	}

	srcDeps := groupBySource(srcDepsList)
	dstDeps := groupBySource(dstDepsList)

	var newSrcDeps []*NixDepenency
	var newDstDeps []*NixDepenency

	for packagePath, srcDep := range srcDeps {
		if dstDep, exist := dstDeps[packagePath]; exist {
			if srcDep.Fetch.Rev == dstDep.Fetch.Rev {
				fmt.Printf("Same version of %v found in both files, removing from %v\n",
					packagePath, srcFile)
				newDstDeps = append(newDstDeps, dstDep)
			} else {
				fmt.Printf("Package %v found in both files but in they use different version. You need to agree on its version manually.\n")
				newSrcDeps = append(newSrcDeps, srcDep)
				newDstDeps = append(newDstDeps, dstDep)
			}
		} else {
			fmt.Printf("Moving %v from %v to %v\n", packagePath, srcFile, dstFile)
			dstDeps[packagePath] = srcDep
			newDstDeps = append(newDstDeps, srcDep)
		}
	}

	if len(newSrcDeps) == 0 {
		if err := os.Remove(srcFile); err != nil {
			return err
		}
		fmt.Printf("%v removed after all dependencies moved to %v\n", srcFile, dstFile)
	} else {
		if err := saveDeps(newSrcDeps, srcFile); err != nil {
			return err
		}
		fmt.Printf("New %v saved\n", srcFile)
	}
	if err := saveDeps(newDstDeps, dstFile); err != nil {
		return err
	}
	fmt.Printf("New %v saved\n", dstFile)

	return nil
}

func groupBySource(depsList []*NixDepenency) map[string]*NixDepenency {
	depsMap := make(map[string]*NixDepenency)
	for _, dep := range depsList {
		depsMap[dep.GoPackagePath] = dep
	}
	return depsMap
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
