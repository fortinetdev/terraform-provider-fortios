export GOPATH=$(dirname "$PWD")
export FORTIOS_ACCESS_TOKEN=8N6m0mb03yq40whQyb7hkx3yHdgfNt #This should be updated according the actual FortiGate VM
export FORTIOS_ACCESS_HOSTNAME=54.226.179.231 # #This should be updated according the actual FortiGate VM
make -C ../src/github.com/terraform-providers/terraform-provider-fortios/  testacc

