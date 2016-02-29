package main

import (
	"log"
	"os/exec"
	"strings"
)

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
		log.Fatal(err)
	}
	return hashFromNixPrefetch(prefetchOut)
}

func hashFromNixPrefetch(prefetchOut []byte) string {
	prefetchStr := strings.TrimSpace(string(prefetchOut))
	prefetchLines := strings.Split(prefetchStr, "\n")
	return prefetchLines[len(prefetchLines)-1]
}
