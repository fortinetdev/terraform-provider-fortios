p=$(dirname "$PWD");
export GOPATH=${p%/*/*/*/*}"/"
make -C ../  fmt
golint ../fortios
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortios/auth
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortios/config
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortios/request
golint ../vendor/github.com/fortinetdev/forti-sdk-go/fortios/sdkcore
make -C ../  build

