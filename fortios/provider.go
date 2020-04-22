package fortios

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider creates and returns the FortiOS terraform.ResourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Hostname/IP address of the Fortinet Device to connect to",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},

			"cabundlefile": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA Bundle file",
			},

			"vdom": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortios_networking_route_static":         resourceNetworkingRouteStatic(),
			"fortios_networking_interface_dhcp":       resourceNetworkingInterfaceDHCP(),
			"fortios_networking_interface_dhcp_ipres": resourceNetworkingInterfaceDHCPIpRes(),
			"fortios_networking_interface_port":       resourceNetworkingInterfacePort(),
			"fortios_system_admin_profiles":           resourceSystemAdminProfiles(),
			"fortios_system_admin_administrator":      resourceSystemAdminAdministrator(),
			"fortios_firewall_object_address":         resourceFirewallObjectAddress(),
			"fortios_firewall_object_addressgroup":    resourceFirewallObjectAddressGroup(),
			"fortios_firewall_object_service":         resourceFirewallObjectService(),
			"fortios_firewall_object_servicegroup":    resourceFirewallObjectServiceGroup(),
			"fortios_firewall_object_servicecategory": resourceFirewallObjectServiceCategory(),
			"fortios_firewall_object_vip":             resourceFirewallObjectVip(),
			"fortios_firewall_object_vipgroup":        resourceFirewallObjectVipGroup(),
			"fortios_firewall_object_ippool":          resourceFirewallObjectIPPool(),
			"fortios_firewall_security_policy":        resourceFirewallSecurityPolicy(),
			"fortios_firewall_security_policyseq":     resourceFirewallSecurityPolicySeq(),
			"fortios_system_setting_global":           resourceSystemSettingGlobal(),
			"fortios_system_setting_dns":              resourceSystemSettingDNS(),
			"fortios_system_setting_ntp":              resourceSystemSettingNTP(),
			"fortios_log_syslog_setting":              resourceLogSyslogSetting(),
			"fortios_log_fortianalyzer_setting":       resourceLogFortiAnalyzerSetting(),
			"fortios_system_apiuser_setting":          resourceSystemAPIUserSetting(),
			"fortios_system_vdom_setting":             resourceSystemVdomSetting(),
			"fortios_system_license_forticare":        resourceSystemLicenseFortiCare(),
			"fortios_system_license_vdom":             resourceSystemLicenseVDOM(),
			"fortios_system_license_vm":               resourceSystemLicenseVM(),
			"fortios_vpn_ipsec_phase1interface":       resourceVPNIPsecPhase1Interface(),
			"fortios_vpn_ipsec_phase2interface":       resourceVPNIPsecPhase2Interface(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	// Init client config with the values from TF files
	config := Config{
		Hostname: d.Get("hostname").(string),
		Token:    d.Get("token").(string),
		CABundle: d.Get("cabundlefile").(string),
		Vdom:     d.Get("vdom").(string),
	}

	v, ok := d.GetOkExists("insecure")
	if ok {
		insecure := v.(bool)
		config.Insecure = &insecure
	}

	// Create Client for later connections
	return config.CreateClient()
}
