with import <nixpkgs> {};

buildGoPackage rec {
  name = "go2nix-${version}";
  version = "20160603-${stdenv.lib.strings.substring 0 7 rev}";
  rev = "ca903924339f654fda93c07f02dc56d8df3c87c3";

  goPackagePath = "github.com/kamilchm/go2nix";

  src = ./.;

  goDeps = ./deps.json;
}
