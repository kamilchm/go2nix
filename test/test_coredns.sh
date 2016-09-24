#!/usr/bin/env bash

set -e

url="https://github.com/miekg/coredns";
rev="d250742d9e8a1851db370c5218e3ae90d373dccb"
pkg_path="github.com/miekg/coredns";

GOPATH=`mktemp -d`

echo "Running test in ${GOPATH}"

git clone ${url} ${GOPATH}/src/$pkg_path
pushd $GOPATH/src/$pkg_path
git checkout ${rev}
git submodule update --init --recursive

go get -v ./...

echo "Running go2nix save"
go2nix save

sort > expected.txt << EOF
github.com/beorn7/perks
github.com/coreos/etcd
github.com/flynn/go-shlex
github.com/fsnotify/fsnotify
github.com/golang/protobuf
github.com/hashicorp/go-syslog
github.com/matttproud/golang_protobuf_extensions
github.com/mholt/caddy
github.com/miekg/dns
github.com/patrickmn/go-cache
github.com/prometheus/client_golang
github.com/prometheus/client_model
github.com/prometheus/common
github.com/prometheus/procfs
github.com/ugorji/go
golang.org/x/net
golang.org/x/sys
k8s.io/kubernetes
EOF

cat deps.nix | grep "goPackagePath" | cut -d "\"" -f 2 | sort > actual.txt

git diff --no-index expected.txt actual.txt && echo "PASS" || echo "FAIL"

popd
#rm -rf ${gopath}
