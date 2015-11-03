package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func get(pkg string) {
	goPath, err := prepareGoPath()
	if err != nil {
		log.Fatal(err)
	}

	if err := goGet(pkg, goPath); err != nil {
		log.Fatal(err)
	}

	save(pkg, goPath)
}

func prepareGoPath() (string, error) {
	return ioutil.TempDir("", "go2nix")
}

func goGet(pkg, goPath string) error {
	goGetCmd := exec.Command("go", "get", "-v", pkg)
	goGetCmd.Env = setGoPath(os.Environ(), goPath)
	goGetCmd.Stderr = os.Stderr
	goGetCmd.Stdout = os.Stdout
	return goGetCmd.Run()
}

func setGoPath(env []string, goPath string) []string {
	found := false
	for i, envVar := range env {
		if strings.HasPrefix(envVar, "GOPATH") {
			env[i] = "GOPATH=" + goPath
			found = true
		}
	}
	if !found {
		env = append(env, "GOPATH="+goPath)
	}

	return env
}
