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
	return hashFromNixPrefetch(prefetchOut)
}

func hashFromNixPrefetch(prefetchOut []byte) string {
	prefetchStr := strings.TrimSpace(string(prefetchOut))
	prefetchLines := strings.Split(prefetchStr, "\n")
	return prefetchLines[len(prefetchLines)-1]
}
