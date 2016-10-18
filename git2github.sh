#!/usr/bin/env bash
set -e
cat $(nix-build -E "import ./git2github.nix ~/nixpkgs/${1}") > ~/nixpkgs/${1}.new
rm ~/nixpkgs/${1}
mv ~/nixpkgs/${1}.new ~/nixpkgs/${1}
