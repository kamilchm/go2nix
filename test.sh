#! /usr/bin/env nix-shell
#! nix-shell -i bash -p ripgrep -I nixpkgs=https://github.com/NixOS/nixpkgs-channels/archive/nixos-unstable.tar.gz
set -e

export NIXPKGS=$(nix-instantiate --eval -E '<nixpkgs>')

printPackage () {
    #echo $1
    #nix-instantiate --eval -E '((import <nixpkgs> {}).callPackage '"$1"' {}).goPackagePath'
    NAME=$(nix-instantiate --eval -E '((import <nixpkgs> {}).callPackage '"$1"' {}).name')
    SRC=$(grep "name = \"" $1 | cut -d '"' -f2 | cut -d '-' -f1 | xargs -I {} nix-build -A "{}.src")

    if [ -f "$SRC/Gopkg.lock" ]; then
        echo "$NAME: got Gopkg.lock >> ready to CONVERT"
    else
        echo "$NAME: NO Gopkg.lock!"
    fi
}

pushd $NIXPKGS
rg -t nix -l buildGoPackage | rg -vF "top-level" | rg -vF "nixos" | \
    rg -vF \
    -e "ethereum" \
    -e "xhyve.nix" \
    -e "minikube/default.nix" \
    -e "terraform" \
    -e "gopherclient" \
    -e "rancher" \
    -e "mop" \
    -e "heroku" \
    -e "beats" \
    -e "gh-ost" \
    -e "sudolikeaboss" \
    -e "mqtt-bench" \
    -e "parrot" \
    -e "docker-machine" \
    -e "xmpp-client" \
    -e "docker-distribution" \
    -e "cloudfoundry-cli" \
    -e "compile-daemon" \
    -e "container-linux" \
    -e "buildkite-agent" \
    -e "drone/default.nix" \
    -e "gdm/default.nix" \
    -e "golint" \
    -e "kube-aws" \
    -e "fscrypt" \
    -e "cloud-print-connector" \
    -e "coredns" \
    -e "mailhog" \
    -e "heapster" \
    -e "alertmanager" \
    -e "prometheus/" \
    -e "cockroach" \
    -e "scaleway-cli" \
    -e "wal-g" \
    -e "azure-vhd-utils" \
    -e "mongodb-tools" \
    -e "phraseapp-client" \
    -e "mynewt-newt" \
    -e "systemd-journal2gelf" \
    -e "platinum-searcher" \
    | sort | uniq | \
    while read goPackage; do
        printPackage "$goPackage"
    done
popd

