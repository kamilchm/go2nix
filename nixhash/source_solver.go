package nixhash

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/kamilchm/go2nix"
)

type SourceSolver struct{}

type Package struct {
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

func (s *SourceSolver) Source(pkg go2nix.GoPackage) (*go2nix.PkgSource, error) {
	cmd, args := prefetchCmd(pkg.Source.Type)
	args = append(args, pkg.Source.Url)

	prefetchOut, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return nil, fmt.Errorf("Command %v failed: %v", cmd, err)
	}

	hash, err := hashFromNixPrefetch(pkg.Source.Type, prefetchOut)
	if err != nil {
		return nil, fmt.Errorf("Unable to extract hash from '%v' output: %v", cmd, err)
	}

	src := pkg.Source
	src.Sha256 = hash
	return src, nil
}

func hashFromNixPrefetch(fetchType go2nix.FetchType, prefetchOut []byte) (string, error) {
	prefetchStr := strings.TrimSpace(string(prefetchOut))
	prefetchLines := strings.Split(prefetchStr, "\n")

	// nix-prefetch-git after https://github.com/NixOS/nixpkgs/pull/13584
	if fetchType == go2nix.Git && prefetchLines[len(prefetchLines)-1][0] == '}' {
		var p Package
		err := json.Unmarshal([]byte(prefetchStr), &p)
		if err != nil {
			return "", fmt.Errorf("Unable to parse prefetch-git output as JSON: %v", err)
		}
		return p.Sha256, nil
	}

	// regular nix-prefetch-* output
	return prefetchLines[len(prefetchLines)-1], nil
}

func prefetchCmd(fetchType go2nix.FetchType) (string, []string) {
	fetchCmds := map[go2nix.FetchType][]string{
		go2nix.Mercurial:  {"nix-prefetch-hg"},
		go2nix.Git:        {"nix-prefetch-git", "--fetch-submodules"},
		go2nix.Subversion: {"nix-prefetch-svn"},
		go2nix.Bazaar:     {"nix-prefetch-bzr"},
	}

	cmd := fetchCmds[fetchType]
	args := []string{}
	if len(cmd) > 1 {
		args = cmd[1 : len(cmd)-1]
	}
	return cmd[0], args
}
