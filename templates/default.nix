# This file was generated by https://github.com/kamilchm/go2nix v[[ .Version ]]
{ stdenv, buildGoPackage, fetchgit, fetchhg, fetchbzr, fetchsvn }:

buildGoPackage rec {
  pname = "[[ .Pkg.Name ]]-unstable";
  version = "[[ .Pkg.UpdateDate.Format "2006-01-02" ]]";
  rev = "[[ .Pkg.Revision ]]";

  [[ if ne .BuildTags "" ]]
  buildFlags = "--tags [[ .BuildTags ]]";
  [[ end -]]
  goPackagePath = "[[ .Pkg.ImportPath ]]";

  src = fetchgit {
    inherit rev;
    url = "[[ .Pkg.VcsRepo ]]";
    sha256 = "[[ .Pkg.Hash ]]";
  };

  goDeps = ./deps.nix;

  # TODO: add metadata https://nixos.org/nixpkgs/manual/#sec-standard-meta-attributes
  meta = {
  };
}
