export GOPATH=$(dirname "$PWD")
make -C ../src/github.com/terraform-providers/terraform-provider-fortios/  fmt
golint ../src/github.com/terraform-providers/terraform-provider-fortios/fortios
make -C ../src/github.com/terraform-providers/terraform-provider-fortios/  build

