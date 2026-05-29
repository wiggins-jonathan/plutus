{
  description = "Plutus development environment";

  nixConfig.bash-prompt-prefix = "(✶ ) ";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = ["x86_64-linux" "x86_64-darwin" "aarch64-linux"];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system: f {
        pkgs = import nixpkgs { inherit system; };
      });
    in {
      devShells = forEachSystem ({ pkgs }: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
          ];

          shellHook = ''
            echo "$(go version)"
          '';
        };
      });
    };
}
