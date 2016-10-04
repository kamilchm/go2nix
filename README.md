[![Circle CI](https://circleci.com/gh/kamilchm/go2nix.svg?style=shield)](https://circleci.com/gh/kamilchm/go2nix)

# go2nix - because Go and Nix are both amazing

## For Nixers - packaging Go applications

`go2nix` provides an autmatic way to create Nix derivations for Go applications.

1. Start with app sources than can be built on your machine with `go build`.
   It means that you need to get all dependencies into current `GOPATH`.
2. Run `go2nix save` in application source dir where `main` package lives.
   This will create 2 files `default.nix` and `deps.nix` that can be moved
   into its own directory under `nixpkgs`.

## For Gophers - reproducible development and build environments

```
stay tuned
```

# Installation

The preferred way of installing `go2nix` is to use `nix` like `nix-env -iA go2nix` or using it declaratively.

But you can also use `go get github.com/kamilchm/go2nix`.
