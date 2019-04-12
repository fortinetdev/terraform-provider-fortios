p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../fortios
golint ../vendor/github.com/fgtdev/fortios-sdk-go/auth
golint ../vendor/github.com/fgtdev/fortios-sdk-go/config
golint ../vendor/github.com/fgtdev/fortios-sdk-go/request
golint ../vendor/github.com/fgtdev/fortios-sdk-go/sdkcore
make -C ../  build

