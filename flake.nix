{
  description = "WhatsMorse";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = {
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in rec {
      devShells.default = pkgs.mkShell {
        packages = [
          pkgs.bashInteractive
          pkgs.go
        ];
      };
      packages.whatsmorse = pkgs.buildGoModule {
        name = "whatsmorse";
        src = ./.;
        # Because vendor file exists, need to set to null
        vendorHash = null;
        meta = with pkgs.lib; {
          description = "A morse-code instant messenger";
          homepage = "https://github.com/humaidq/whatsmorse";
          license = licenses.mit;
        };
      };
      defaultPackage = packages.whatsmorse;
    });
}
