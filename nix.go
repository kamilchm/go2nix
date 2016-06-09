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

	var newSrcDeps []*NixDependency
	newSrcInclude := NixDependency{IncludeFile: dstFile}
	newDstDeps := dstDepsList

	for packagePath, srcDep := range srcDeps {
		if dstDep, exist := dstDeps[packagePath]; exist {
			if srcDep.Fetch.Rev == dstDep.Fetch.Rev {
				fmt.Printf("Same version of %v found in both files, removing from %v\n",
					packagePath, srcFile)
				newSrcInclude.Packages = append(newSrcInclude.Packages, packagePath)
			} else {
				fmt.Printf("Package %v found in both files but in they use different version. You need to agree on its version manually.\n", packagePath)
				newSrcDeps = append(newSrcDeps, srcDep)
			}
		} else {
			fmt.Printf("Moving %v from %v to %v\n", packagePath, srcFile, dstFile)
			dstDeps[packagePath] = srcDep
			newDstDeps = append(newDstDeps, srcDep)
			newSrcInclude.Packages = append(newSrcInclude.Packages, packagePath)
		}
	}

	if len(newSrcInclude.Packages) > 0 {
		newSrcDeps = append(newSrcDeps, &newSrcInclude)
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

func groupBySource(depsList []*NixDependency) map[string]*NixDependency {
	depsMap := make(map[string]*NixDependency)
	for _, dep := range depsList {
		depsMap[dep.GoPackagePath] = dep
	}
	return depsMap
}

func saveDeps(deps []*NixDependency, depsFilename string) error {
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

func loadDeps(depsFilename string) ([]*NixDependency, error) {
	depsFile, err := ioutil.ReadFile(depsFilename)
	if err != nil {
		return nil, err
	}
	var deps []*NixDependency
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
