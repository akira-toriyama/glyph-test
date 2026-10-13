{
  # glyph-test — the fleet's live-fire harness. This flake pins a buildGoModule vendorHash
  # so a dependabot gomod bump here has exactly the stale-hash problem fleet-repin
  # (akira-toriyama/.github) fixes; the re-pin is scripts/repin-vendorhash.sh,
  # the canonical fleet-sync ships to every repo with a root flake.nix + go.sum.
  description = "glyph-test — live ammunition for fleet-repin";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        packages.default = pkgs.buildGoModule {
          pname = "glyph-test";
          version = "0.0.0";
          src = ./.;
          vendorHash = "sha256-fxxp7ECuMMmLw7L5/lPi7lfmJM2p+Te9AqXm5Xed3G0=";
        };
      });
}
