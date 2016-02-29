package main

import (
	"log"
	"os/exec"
	"strings"
)

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

	// nix-prefetch-git after https://github.com/NixOS/nixpkgs/pull/11671
	if pathType == "git" && prefetchLines[len(prefetchLines)-1][0] == '}' {
		return prefetchLines[len(prefetchLines)-2]
	}

	// regular nix-prefetch-* output
	return "  sha512 = \"" + prefetchLines[len(prefetchLines)-1] + "\""
}
