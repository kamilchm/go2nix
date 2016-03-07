# This file was generated by go2nix.
#with import <nixpkgs> {};
{ stdenv, lib, go, goPackages, fetchgit, fetchhg, fetchbzr, fetchsvn }:

with goPackages;

buildGoPackage rec {
  name = "[[ .Pkg.Name ]]-${version}";
  version = "[[ .Pkg.UpdateDate.Format "20060102" ]]-${stdenv.lib.strings.substring 0 7 rev}";
  rev = "[[ .Pkg.Revision ]]";

  buildInputs = [ go ];
  [[ if ne .BuildTags "" ]]
  buildFlags = "--tags [[ .BuildTags ]]";
  [[ end ]]
  goPackagePath = "[[ .Pkg.ImportPath ]]";

  src = fetchgit {
    inherit rev;
    url = "[[ .Pkg.VcsRepo ]]";
    sha256 = "[[ .Pkg.Hash ]]";
  };

  extraSrcs = (builtins.attrValues rec {
    [[ range $dep := .Deps ]][[ $dep.Name ]] = {
      goPackagePath = "[[ $dep.ImportPath ]]";

      src = fetch[[ $dep.VcsCommand ]] {
        url = "[[ $dep.VcsRepo ]]";
        rev = "[[ $dep.Revision ]]";
        sha256 = "[[ $dep.Hash ]]";
      };
    };
    [[ range $vend := $dep.Vendored ]][[ $vend.Name ]] = {
      goPackagePath = "[[ $vend.ImportPath ]]";

      src = "${[[ $dep.Name ]].src}/[[ $vend.PkgDir ]]";
    };
    [[ end ]][[ end ]] # can be improved with Go 1.6 http://talks.golang.org/2016/state-of-go.slide#14
  });
}
