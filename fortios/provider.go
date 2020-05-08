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
				Description: "Hostname/IP address of the FortiOS to connect to",
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

			"fmg_hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Hostname/IP address of the FortiManager to connect to",
			},

			"fmg_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"fmg_passwd": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"fmg_insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},

			"fmg_cabundlefile": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA Bundle file",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortios_networking_route_static":                 resourceNetworkingRouteStatic(),
			"fortios_networking_interface_port":               resourceNetworkingInterfacePort(),
			"fortios_system_admin_profiles":                   resourceSystemAdminProfiles(),
			"fortios_system_admin_administrator":              resourceSystemAdminAdministrator(),
			"fortios_firewall_object_address":                 resourceFirewallObjectAddress(),
			"fortios_firewall_object_addressgroup":            resourceFirewallObjectAddressGroup(),
			"fortios_firewall_object_service":                 resourceFirewallObjectService(),
			"fortios_firewall_object_servicegroup":            resourceFirewallObjectServiceGroup(),
			"fortios_firewall_object_servicecategory":         resourceFirewallObjectServiceCategory(),
			"fortios_firewall_object_vip":                     resourceFirewallObjectVip(),
			"fortios_firewall_object_vipgroup":                resourceFirewallObjectVipGroup(),
			"fortios_firewall_object_ippool":                  resourceFirewallObjectIPPool(),
			"fortios_firewall_security_policy":                resourceFirewallSecurityPolicy(),
			"fortios_firewall_security_policyseq":             resourceFirewallSecurityPolicySeq(),
			"fortios_system_setting_global":                   resourceSystemSettingGlobal(),
			"fortios_system_setting_dns":                      resourceSystemSettingDNS(),
			"fortios_system_setting_ntp":                      resourceSystemSettingNTP(),
			"fortios_log_syslog_setting":                      resourceLogSyslogSetting(),
			"fortios_log_fortianalyzer_setting":               resourceLogFortiAnalyzerSetting(),
			"fortios_system_apiuser_setting":                  resourceSystemAPIUserSetting(),
			"fortios_system_vdom_setting":                     resourceSystemVdomSetting(),
			"fortios_system_license_forticare":                resourceSystemLicenseFortiCare(),
			"fortios_system_license_vdom":                     resourceSystemLicenseVDOM(),
			"fortios_system_license_vm":                       resourceSystemLicenseVM(),
			"fortios_vpn_ipsec_phase1interface":               resourceVPNIPsecPhase1Interface(),
			"fortios_vpn_ipsec_phase2interface":               resourceVPNIPsecPhase2Interface(),
			"fortios_json_generic_api":                        resourceJSONGenericAPI(),
			"fortios_fmg_system_admin_profiles":               resourceFortimanagerSystemAdminProfiles(),
			"fortios_fmg_system_admin_user":                   resourceFortimanagerSystemAdminUser(),
			"fortios_fmg_devicemanager_device":                resourceFortimanagerDVMDevice(),
			"fortios_fmg_devicemanager_script":                resourceFortimanagerDVMScript(),
			"fortios_fmg_devicemanager_script_execute":        resourceFortimanagerDVMScriptExecute(),
			"fortios_fmg_devicemanager_install_device":        resourceFortimanagerDVMInstallDev(),
			"fortios_fmg_devicemanager_install_policypackage": resourceFortimanagerDVMInstallPolicyPackage(),
			"fortios_fmg_firewall_security_policy":            resourceFortimanagerFirewallSecurityPolicy(),
			"fortios_fmg_firewall_security_policypackage":     resourceFortimanagerFirewallSecurityPolicyPackage(),
			"fortios_fmg_object_adom_revision":                resourceFortimanagerObjectAdomRevision(),
			"fortios_fmg_firewall_object_address":             resourceFortimanagerFirewallObjectAddress(),
			"fortios_fmg_firewall_object_service":             resourceFortimanagerFirewallObjectService(),
			"fortios_fmg_firewall_object_ippool":              resourceFortimanagerFirewallObjectIppool(),
			"fortios_fmg_firewall_object_vip":                 resourceFortimanagerFirewallObjectVip(),
			"fortios_fmg_system_syslogserver":                 resourceFortimanagerSystemSyslogServer(),
			"fortios_fmg_system_network_interface":            resourceFortimanagerSystemNetworkInterface(),
			"fortios_fmg_system_network_route":                resourceFortimanagerSystemNetworkRoute(),
			"fortios_fmg_system_adom":                         resourceFortimanagerSystemAdom(),
			"fortios_fmg_system_global":                       resourceFortimanagerSystemGlobal(),
			"fortios_fmg_system_admin":                        resourceFortimanagerSystemAdmin(),
			"fortios_fmg_system_dns":                          resourceFortimanagerSystemDNS(),
			"fortios_fmg_system_ntp":                          resourceFortimanagerSystemNTP(),
			"fortios_fmg_system_license_vm":                   resourceFortimanagerSystemLicenseVM(),
			"fortios_fmg_system_license_forticare":            resourceFortimanagerSystemLicenseFortiCare(),
			"fortios_fmg_jsonrpc_request":                     resourceFortimanagerJSONRPCRequest(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	// Init client config with the values from TF files
	config := Config{
		Hostname:     d.Get("hostname").(string),
		Token:        d.Get("token").(string),
		CABundle:     d.Get("cabundlefile").(string),
		Vdom:         d.Get("vdom").(string),
		FMG_Hostname: d.Get("fmg_hostname").(string),
		FMG_CABundle: d.Get("fmg_cabundlefile").(string),
		FMG_Username: d.Get("fmg_username").(string),
		FMG_Passwd:   d.Get("fmg_passwd").(string),
	}

	v1, ok1 := d.GetOkExists("insecure")
	if ok1 {
		insecure := v1.(bool)
		config.Insecure = &insecure
	}

	v2, ok2 := d.GetOkExists("fmg_insecure")
	if ok2 {
		insecure := v2.(bool)
		config.FMG_Insecure = &insecure
	}

	// Create Client for later connections
	return config.CreateClient()
}
