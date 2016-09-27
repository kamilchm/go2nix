#!/usr/bin/env bash

set -e

url="https://github.com/jeffail/leaps";
rev="5cf7328a8c498041d2a887e89f22f138498f4621"
pkg_path="github.com/jeffail/leaps";
submodule="cmd/leaps"

GOPATH=`mktemp -d`

echo "Running test in ${GOPATH}"

git clone ${url} ${GOPATH}/src/$pkg_path
pushd $GOPATH/src/$pkg_path
git checkout ${rev}
git submodule update --init --recursive

cd $submodule
go get ./...

echo "Running go2nix save"

go2nix save

sort > expected.txt << EOF
golang.org/x/net
EOF

cat deps.nix | grep "goPackagePath" | cut -d "\"" -f 2 | sort > actual.txt

git diff --no-index expected.txt actual.txt && echo "PASS" || echo "FAIL"

popd
#rm -rf ${gopath}
