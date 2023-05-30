package main

import (
	"github.com/fortinetdev/terraform-provider-fortios/fortios"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortios.Provider})
}
