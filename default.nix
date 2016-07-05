with import <nixpkgs>{};

buildGoPackage rec {
  name = "go2nix-dev";

  goPackagePath = "github.com/kamilchm/go2nix";

  src = ./.;

  buildInputs = [ go-bindata.bin ];

  preBuild = ''
    go generate ./...
  '';

  goDeps = import ./deps.nix { inherit (pkgs) fetchgit fetchhg fetchbzr fetchsvn; };
}
