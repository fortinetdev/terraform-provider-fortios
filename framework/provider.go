package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &fwprovider{}
var _ provider.ProviderWithEphemeralResources = &fwprovider{}

// New returns a new, initialized Terraform Plugin Framework-style provider instance.
// The provider instance is fully configured once the `Configure` method has been called.
func New(primary interface{ Meta() interface{} }) provider.Provider {
	return &fwprovider{
		Primary: primary,
	}
}

type fwprovider struct {
	Primary interface{ Meta() interface{} }
}

func (f *fwprovider) Metadata(ctx context.Context, request provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "fortios"
}

func (f *fwprovider) Schema(ctx context.Context, request provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				Optional:    true,
				Description: "The hostname/IP address of the FortiOS to be connected",
			},

			"token": schema.StringAttribute{
				Optional:    true,
				Description: "",
			},

			"username": schema.StringAttribute{
				Optional:    true,
				Description: "The username of the user.",
			},

			"password": schema.StringAttribute{
				Optional:    true,
				Description: "The password of the user.",
			},

			"insecure": schema.BoolAttribute{
				Optional:    true,
				Description: "",
			},

			"cabundlefile": schema.StringAttribute{
				Optional:    true,
				Description: "CA Bundle file",
			},

			"cabundlecontent": schema.StringAttribute{
				Optional:    true,
				Description: "CA Bundle file content",
			},

			"http_proxy": schema.StringAttribute{
				Optional:    true,
				Description: "HTTP proxy address",
			},

			"peerauth": schema.StringAttribute{
				Optional:    true,
				Description: "Enable/disable peer authentication, can be 'enable' or 'disable'",
			},

			"cacert": schema.StringAttribute{
				Optional:    true,
				Description: "CA certtificate(Optional)",
			},

			"clientcert": schema.StringAttribute{
				Optional:    true,
				Description: "User certificate",
			},

			"clientkey": schema.StringAttribute{
				Optional:    true,
				Description: "User private key",
			},

			"vdom": schema.StringAttribute{
				Optional:    true,
				Description: "Vdom name of FortiOS. It will apply to all resources. Specify variable `vdomparam` on each resource will override the vdom value on that resource.",
			},

			"update_if_exist": schema.BoolAttribute{
				Optional: true,
			},

			"fmg_hostname": schema.StringAttribute{
				Optional:    true,
				Description: "Hostname/IP address of the FortiManager to connect to",
			},

			"fmg_username": schema.StringAttribute{
				Optional:    true,
				Description: "",
			},

			"fmg_passwd": schema.StringAttribute{
				Optional:    true,
				Description: "",
			},

			"fmg_insecure": schema.BoolAttribute{
				Optional:    true,
				Description: "",
			},

			"fmg_cabundlefile": schema.StringAttribute{
				Optional:    true,
				Description: "CA Bundle file",
			},
		},
	}
}

func (f *fwprovider) Configure(ctx context.Context, request provider.ConfigureRequest, response *provider.ConfigureResponse) {
	// Provider's parsed configuration (its instance state) is available through the primary provider's Meta() method.
	v := f.Primary.Meta()
	response.DataSourceData = v
	response.ResourceData = v
	response.EphemeralResourceData = v
}

func (f *fwprovider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (f *fwprovider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewResourceSystemVirtualSwitchPortAssignment,
	}

}

func (f *fwprovider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return nil
}
