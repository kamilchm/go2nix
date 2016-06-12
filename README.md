[![Circle CI](https://circleci.com/gh/kamilchm/go2nix.svg?style=shield)](https://circleci.com/gh/kamilchm/go2nix)

# go2nix - because Go and Nix are both amazing

## For Nixers - packaging Go applications

`go2nix` provides an autmatic way to create Nix derivations for Go applications.

1. Start with app sources than can be built on your machine with `go build`.
   It means that you need to get all dependencies into current `GOPATH`.
2. Run `go2nix save` in application source dir where `main` package lives.
   This will create 2 files `default.nix` and `deps.json` that can be moved
   into its own directory under `nixpkgs`.
3. Run `go2nix merge <new_app_deps.json> <nixpkgs/development/go-modules/libs.json>`
   to merge app dependecies into common go libs from `nixpkgs`.

## For Gophers - reproducible development and build environments

```
stay tuned
```
