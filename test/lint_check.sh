p=$(dirname "$PWD");  
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../fortios
golint ../vendor/github.com/fortios/fortios-sdk/auth
golint ../vendor/github.com/fortios/fortios-sdk/config
golint ../vendor/github.com/fortios/fortios-sdk/request
golint ../vendor/github.com/fortios/fortios-sdk/sdkcore
make -C ../  build

