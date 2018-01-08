# go2nix v2

This version is under development and is not stable.

## Usage

Go2nix v2 uses the `Gopkg.lock` file to generate Nix expressions.

```
$ dep ensure
$ go2nix
```

If the Go project doesn't have any `Gopkg.lock` file, the
[`dep`](https://github.com/golang/dep) tool is able to generate this
file from several other dependency managers (at least `Glide` and
`govendor`). We then first need to create the `Gopkg.lock` file by
using `dep init`.

```
$ dep init
$ go2nix
```

## TODO

* [ ] package version discovery
* [ ] optional unstable in package name
* [ ] works with all nixpkgs Go packages
* [ ] integrations tests
* [ ] test coverage
* [ ] usage documentation with [diagram](https://github.com/golang/dep/#usage)
* [ ] binary subPackages discovery
* [ ] fetchFromGitHub ?
