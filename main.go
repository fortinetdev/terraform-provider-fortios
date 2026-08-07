package main

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-fortios/fortios"
	"github.com/terraform-providers/terraform-provider-fortios/framework"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func main() {
	ctx := context.Background()
	primary := fortios.Provider()

	providers := []func() tfprotov5.ProviderServer{
		func() tfprotov5.ProviderServer {
			return schema.NewGRPCProviderServer(primary)
		},
		providerserver.NewProtocol5(framework.New(primary)),
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf5server.ServeOpt

	err = tf5server.Serve("registry.terraform.io/fortinetdev/fortios", muxServer.ProviderServer, serveOpts...)
}
