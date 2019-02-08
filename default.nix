let
pkgs = import <nixpkgs> {};
in
  with pkgs; [
    nodejs

    #required by reprocessing
    gcc
    gnumake
    file
  ]


