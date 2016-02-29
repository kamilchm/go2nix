package main

import (
	"log"
	"os/exec"
	"encoding/json"
	"strings"
)

type Package struct {
    Url string `json:"url"`
    Rev string `json:"rev"`
    Sha256 string `json:"sha256"`
}

func calculateHash(url, pathType string) (hash string) {
	prefetchCmd := exec.Command("nix-prefetch-"+pathType, url)
	prefetchOut, err := prefetchCmd.Output()
	if err != nil {
		log.Fatal(err)
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
	return "  sha512 = \"" + prefetchLines[len(prefetchLines)-1] + "\""
}
