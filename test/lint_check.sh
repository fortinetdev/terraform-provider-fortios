p=$(dirname "$PWD");  
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../fortios
make -C ../  build

