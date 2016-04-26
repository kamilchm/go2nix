[![Circle CI](https://circleci.com/gh/kamilchm/go2nix.svg?style=shield)](https://circleci.com/gh/kamilchm/go2nix)

# go2nix

It's just a PoC and I working on major rework on ``common-libs`` branch. So there's no docs here :/ and I plan to write about usage when I settle with implementation. But it should be usable right now generating derivations with complete dependency set.
First, you need to setup you Go environment with proper ``GOPATH`` https://github.com/golang/go/wiki/GOPATH.
Then try to do:
```
$ go get -u github.com/btcsuite/btcd/...
$ cd $GOPATH/src/github.com/btcsuite/btcd
$ go2nix save
```
It should produce complete ``default.nix`` that you can build or include somewhere in nixpkgs.

### Related Works

* https://github.com/golang/go/blob/master/src/cmd/go/list.go
* https://github.com/cespare/deplist
* https://github.com/golang/tools/blob/9d2ff756b797a862da0686e2e41e09cd87da017b/go/types/package.go#L57
* https://github.com/golang/tools/blob/ba766134cc38dea5ed957cf2b40266a8e0aa5660/go/buildutil/allpackages.go#L52
* https://github.com/NixOS/nixpkgs/pull/12010
* go2nix save --exclude example
