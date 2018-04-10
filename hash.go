package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

type Package struct {
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

func calculateHash(url, pathType string) (hash string) {
	args := []string{}

	if pathType == "git" {
		// `fetchgit` passes this argument by default
		args = append(args, "--fetch-submodules")
	}

	args = append(args, url)
	prefetchCmd := exec.Command("nix-prefetch-"+pathType, args...)
	prefetchOut, err := prefetchCmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			log.Print(string(exitErr.Stderr))
		}
		log.Fatalf("Command %v %v failed: %v", prefetchCmd.Path, prefetchCmd.Args, err)
	}
	return hashFromNixPrefetch(pathType, prefetchOut)
}

func hashFromNixPrefetch(pathType string, prefetchOut []byte) string {
	prefetchStr := strings.TrimSpace(string(prefetchOut))
	prefetchLines := strings.Split(prefetchStr, "\n")

	// nix-prefetch-git after https://github.com/NixOS/nixpkgs/pull/13584
	if pathType == "git" && prefetchLines[len(prefetchLines)-1][0] == '}' {
		var p Package
		err := json.Unmarshal([]byte(prefetchStr), &p)
		if err != nil {
			log.Fatal(err)
		}
		return p.Sha256
	}

	// regular nix-prefetch-* output
	return prefetchLines[len(prefetchLines)-1]
}
