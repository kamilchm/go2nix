gitDeps:
with import <nixpkgs>{};
with builtins;

let
  git2github = goDep: with goDep;
  let
    owner = elemAt (lib.splitString "/" fetch.url) 3;
    repo = elemAt (lib.splitString "/" fetch.url) 4;
  in ''
    {
      goPackagePath = "${goPackagePath}";
      fetch = {
        type = "FromGitHub";
        owner = "${owner}";
        repo = "${repo}";
        rev = "${fetch.rev}";
        sha256 = "${fetch.sha256}";
      };
    }
    '';

  rewrite = goDep: with goDep; ''
    {
      goPackagePath = "${goPackagePath}";
      fetch = {
        type = "${fetch.type}";
        url = "${fetch.url}";
        rev = "${fetch.rev}";
        sha256 = "${fetch.sha256}";
      };
    }
    '';

  oldDeps = lib.partition (d:
    d.fetch.type == "git" && lib.hasPrefix "https://github.com/" d.fetch.url
    ) (import gitDeps)
  ;
    
  toConvert = oldDeps.right;
  dontTouch = map rewrite oldDeps.wrong;

  convertedDeps = map git2github toConvert; 

  newDeps = convertedDeps ++ dontTouch;

  newDepsNix = toFile "deps.nix" ( foldl' (p: n: p + n) "[\n" newDeps );
in
stdenv.mkDerivation {
  name = "goDeps";
  src = gitDeps;
  phases = [ "installPhase" ];

  installPhase = ''
    cp ${newDepsNix} $out
  '';
}
