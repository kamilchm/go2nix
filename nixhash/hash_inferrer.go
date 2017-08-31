package nixhash

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/kamilchm/go2nix"
)

type HashInferrer struct{}

type Package struct {
	Url    string `json:"url"`
	Rev    string `json:"rev"`
	Sha256 string `json:"sha256"`
}

func (s *HashInferrer) Infer(pkg go2nix.GoPackage) (go2nix.GoPackage, error) {
	cmd, args := prefetchCmd(pkg.Source.Type)
	args = append(args, pkg.Source.Url)
	args = append(args, pkg.Source.Revision)

	prefetchOut, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return pkg, fmt.Errorf("Command '%v %s' failed: %v", cmd, strings.Join(args, " "), err)
	}

	hash, err := hashFromNixPrefetch(pkg.Source.Type, prefetchOut)
	if err != nil {
		return pkg, fmt.Errorf("Unable to extract hash from '%v' output: %v", cmd, err)
	}

	pkg.Source.Sha256 = hash
	return pkg, nil
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
	cmd := "nix-prefetch-" + fetchType.String()

	if fetchType == go2nix.Git {
		return cmd, []string{"--fetch-submodules"}
	}

	return cmd, []string{}
}
