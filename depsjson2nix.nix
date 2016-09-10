# helps to migrate deps.json -> deps.nix
# eg. cat $(nix-build -E "import ./depsjson2nix.nix ~/nixpkgs/pkgs/servers/caddy/deps.json") > ~/nixpkgs/pkgs/servers/caddy/deps.nix
depsJson:
with import <nixpkgs>{};
with builtins;

let
  dep2nix = goDep: with goDep; ''
    {
      goPackagePath = "${goPackagePath}";
      fetch = {
        type = ${fetch.type};
        url = "${fetch.url}";
        rev = "${fetch.rev}";
        sha256 = "${fetch.sha256}";
      };
    }
    '';

  importGodeps = { depsFile, filterPackages ? [] }:
  let
    deps = lib.importJSON depsFile;
    external = filter (d: d ? include) deps;
    direct = filter (d: d ? goPackagePath && (length filterPackages == 0 || elem d.goPackagePath filterPackages)) deps;
  in
    foldl' (p: n: p + n) "[\n"
    ((map importGodeps (map (d: { depsFile = <nixpkgs> + "/pkgs/development/go-modules/generic/" + d.include; filterPackages = d.packages; }) external)) ++ (map dep2nix direct)) + "]\n";

  depsNix = toFile "deps.nix" (importGodeps { depsFile = toPath depsJson; });
in
stdenv.mkDerivation {
  name = "goDeps";
  src = depsJson;
  phases = [ "installPhase" ];

  installPhase = ''
    cp ${depsNix} $out
  '';
}
