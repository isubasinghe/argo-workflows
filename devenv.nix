{ pkgs, ... }: 



{ 
  packages = [
    pkgs.git
    pkgs.jq
    pkgs.kube3d
    pkgs.docker
  ];

  enterShell = ''
    echo hello
    jq --version
  ''; 

}
