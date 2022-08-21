{
  description = "bookmarks allows users to set folder bookmarks in the cli";

  outputs = { self, nixpkgs }: {

    defaultPackage.aarch64-darwin = nixpkgs.legacyPackages.aarch64-darwin.buildGoModule {
      name = "bookmarks";
      src = ./.;
      vendorSha256 = "1fq2nyf7hddkxdv8y89jca1zr74pqf42l1xcrszwxfig9jikf71w";
    };

    devShells.aarch_86-darwin.default = nixpkgs.mkShell {
      name = "my shell";
      packages = [
        nixpkgs.figlet
      ];
    };
  };
}
