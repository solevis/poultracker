{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
}:

pkgs.buildGoApplication {
  name = "poultracker";
  pwd = ./.;
  src = ./.;
  subPackages = [ "cmd/poultracker" ];
  modules = ./gomod2nix.toml;
}
