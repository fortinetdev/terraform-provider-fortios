// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VDOM settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingsUpdate,
		Read:   resourceSystemSettingsRead,
		Update: resourceSystemSettingsUpdate,
		Delete: resourceSystemSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"opmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ngfw_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"implicit_allow_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"consolidated_firewall_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_external_dest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manageip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"manageip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),
				Optional:     true,
				Computed:     true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),
				Optional:     true,
				Computed:     true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),
				Optional:     true,
				Computed:     true,
			},
			"bfd_dont_enforce_src_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utf8_spam_tagging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wccp_cache_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_stats_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_stats_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"v4_ecmp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"fw_session_hairpin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prp_trailer_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snat_hairpin_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy_interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_proxy_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"central_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_default_policy_columns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_down_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auxiliary_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ses_denied_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_src_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_linkdown_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asymroute6_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_session_without_init": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_expectation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_helper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_nat_trace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sip_tcp_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"sip_udp_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"sip_ssl_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"sccp_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"multicast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_ttl_notchange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_skip_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_subnet_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deny_tcp_with_icmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ecmp_max_paths": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"discovered_device_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 365),
				Optional:     true,
				Computed:     true,
			},
			"email_portal_check_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_voip_alg_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_icap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_nat46_64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_implicit_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dns_database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_load_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_multicast_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dos_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_object_colors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_replacement_message_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_voip_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ap_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_security_profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dynamic_profile_display": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_local_in_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_local_reports": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wanopt_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_explicit_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dynamic_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dlp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_sslvpn_personal_bookmarks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_sslvpn_realms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_policy_based_ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_multiple_utm_profiles": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_spamfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_file_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_application_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_ips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_endpoint_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_endpoint_control_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dhcp_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_vpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wireless_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_switch_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortiap_split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_webfilter_advanced": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_traffic_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_wan_load_balancing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_antivirus": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_webfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_dnsfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_fortiextender_controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_advanced_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_allow_unnamed_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_email_collection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_domain_ip_reputation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_multiple_interface_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_policy_disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_per_policy_disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_policy_learning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compliance_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_session_resume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_quick_crash_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_dn_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ike_natt_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 65535),
				Optional:     true,
				Computed:     true,
			},
			"block_land_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSettings")
	}

	return resourceSystemSettingsRead(d, m)
}

func resourceSystemSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSettings resource from API: %v", err)
	}
	return nil
}

func flattenSystemSettingsComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsOpmode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsInspectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsNgfwMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsImplicitAllowDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSslSshProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsConsolidatedFirewallMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsHttpExternalDest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsManageip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsManageip6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdDetectMult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBfdDontEnforceSrcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsUtf8SpamTagging(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsWccpCacheEngine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsVpnStatsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsVpnStatsPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsV4EcmpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMacTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsFwSessionHairpin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsPrpTrailerAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSnatHairpinTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDhcpProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDhcpProxyInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDhcpProxyInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDhcpServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDhcp6ServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsCentralNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDefaultPolicyColumns(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSettingsGuiDefaultPolicyColumnsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSettingsGuiDefaultPolicyColumnsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsLldpReception(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsLldpTransmission(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsLinkDownAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAuxiliarySession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAsymrouteIcmp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSesDeniedTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsStrictSrcCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAllowLinkdownPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAsymroute6Icmp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSctpSessionWithoutInit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipExpectation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipHelper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipNatTrace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipTcpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipUdpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSipSslPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsSccpPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMulticastForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMulticastTtlNotchange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsMulticastSkipPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsAllowSubnetOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDenyTcpWithIcmp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsEcmpMaxPaths(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDiscoveredDeviceTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsEmailPortalCheckDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsDefaultVoipAlgMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiIcap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiNat4664(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiImplicitPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDnsDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiLoadBalance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiMulticastPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDosPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiObjectColors(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiReplacementMessageGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiVoipProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiApProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiSecurityProfileGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDynamicProfileDisplay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiLocalInPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiLocalReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiWanoptCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiExplicitProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDynamicRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDlp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiSslvpnPersonalBookmarks(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiSslvpnRealms(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiPolicyBasedIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiThreatWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiMultipleUtmProfiles(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiSpamfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiFileFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiApplicationControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiIps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiEndpointControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiEndpointControlAdvanced(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDhcpAdvanced(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiVpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiWirelessController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiSwitchController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiFortiapSplitTunneling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiWebfilterAdvanced(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiTrafficShaping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiWanLoadBalancing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiAntivirus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiWebfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDnsfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiWafProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiFortiextenderController(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiAdvancedPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiAllowUnnamedPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiEmailCollection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiDomainIpReputation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiMultipleInterfacePolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiPerPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsGuiPolicyLearning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsComplianceCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIkeSessionResume(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIkeQuickCrashDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIkeDnFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIkePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsIkeNattPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSettingsBlockLandAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("comments", flattenSystemSettingsComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("opmode", flattenSystemSettingsOpmode(o["opmode"], d, "opmode", sv)); err != nil {
		if !fortiAPIPatch(o["opmode"]) {
			return fmt.Errorf("Error reading opmode: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenSystemSettingsInspectionMode(o["inspection-mode"], d, "inspection_mode", sv)); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("ngfw_mode", flattenSystemSettingsNgfwMode(o["ngfw-mode"], d, "ngfw_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ngfw-mode"]) {
			return fmt.Errorf("Error reading ngfw_mode: %v", err)
		}
	}

	if err = d.Set("implicit_allow_dns", flattenSystemSettingsImplicitAllowDns(o["implicit-allow-dns"], d, "implicit_allow_dns", sv)); err != nil {
		if !fortiAPIPatch(o["implicit-allow-dns"]) {
			return fmt.Errorf("Error reading implicit_allow_dns: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenSystemSettingsSslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("consolidated_firewall_mode", flattenSystemSettingsConsolidatedFirewallMode(o["consolidated-firewall-mode"], d, "consolidated_firewall_mode", sv)); err != nil {
		if !fortiAPIPatch(o["consolidated-firewall-mode"]) {
			return fmt.Errorf("Error reading consolidated_firewall_mode: %v", err)
		}
	}

	if err = d.Set("http_external_dest", flattenSystemSettingsHttpExternalDest(o["http-external-dest"], d, "http_external_dest", sv)); err != nil {
		if !fortiAPIPatch(o["http-external-dest"]) {
			return fmt.Errorf("Error reading http_external_dest: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenSystemSettingsFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-session-dirty"]) {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("manageip", flattenSystemSettingsManageip(o["manageip"], d, "manageip", sv)); err != nil {
		if !fortiAPIPatch(o["manageip"]) {
			return fmt.Errorf("Error reading manageip: %v", err)
		}
	}

	if err = d.Set("gateway", flattenSystemSettingsGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemSettingsIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("manageip6", flattenSystemSettingsManageip6(o["manageip6"], d, "manageip6", sv)); err != nil {
		if !fortiAPIPatch(o["manageip6"]) {
			return fmt.Errorf("Error reading manageip6: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenSystemSettingsGateway6(o["gateway6"], d, "gateway6", sv)); err != nil {
		if !fortiAPIPatch(o["gateway6"]) {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("ip6", flattenSystemSettingsIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemSettingsDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("bfd", flattenSystemSettingsBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", flattenSystemSettingsBfdDesiredMinTx(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-desired-min-tx"]) {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", flattenSystemSettingsBfdRequiredMinRx(o["bfd-required-min-rx"], d, "bfd_required_min_rx", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-required-min-rx"]) {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", flattenSystemSettingsBfdDetectMult(o["bfd-detect-mult"], d, "bfd_detect_mult", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-detect-mult"]) {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_dont_enforce_src_port", flattenSystemSettingsBfdDontEnforceSrcPort(o["bfd-dont-enforce-src-port"], d, "bfd_dont_enforce_src_port", sv)); err != nil {
		if !fortiAPIPatch(o["bfd-dont-enforce-src-port"]) {
			return fmt.Errorf("Error reading bfd_dont_enforce_src_port: %v", err)
		}
	}

	if err = d.Set("utf8_spam_tagging", flattenSystemSettingsUtf8SpamTagging(o["utf8-spam-tagging"], d, "utf8_spam_tagging", sv)); err != nil {
		if !fortiAPIPatch(o["utf8-spam-tagging"]) {
			return fmt.Errorf("Error reading utf8_spam_tagging: %v", err)
		}
	}

	if err = d.Set("wccp_cache_engine", flattenSystemSettingsWccpCacheEngine(o["wccp-cache-engine"], d, "wccp_cache_engine", sv)); err != nil {
		if !fortiAPIPatch(o["wccp-cache-engine"]) {
			return fmt.Errorf("Error reading wccp_cache_engine: %v", err)
		}
	}

	if err = d.Set("vpn_stats_log", flattenSystemSettingsVpnStatsLog(o["vpn-stats-log"], d, "vpn_stats_log", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-stats-log"]) {
			return fmt.Errorf("Error reading vpn_stats_log: %v", err)
		}
	}

	if err = d.Set("vpn_stats_period", flattenSystemSettingsVpnStatsPeriod(o["vpn-stats-period"], d, "vpn_stats_period", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-stats-period"]) {
			return fmt.Errorf("Error reading vpn_stats_period: %v", err)
		}
	}

	if err = d.Set("v4_ecmp_mode", flattenSystemSettingsV4EcmpMode(o["v4-ecmp-mode"], d, "v4_ecmp_mode", sv)); err != nil {
		if !fortiAPIPatch(o["v4-ecmp-mode"]) {
			return fmt.Errorf("Error reading v4_ecmp_mode: %v", err)
		}
	}

	if err = d.Set("mac_ttl", flattenSystemSettingsMacTtl(o["mac-ttl"], d, "mac_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["mac-ttl"]) {
			return fmt.Errorf("Error reading mac_ttl: %v", err)
		}
	}

	if err = d.Set("fw_session_hairpin", flattenSystemSettingsFwSessionHairpin(o["fw-session-hairpin"], d, "fw_session_hairpin", sv)); err != nil {
		if !fortiAPIPatch(o["fw-session-hairpin"]) {
			return fmt.Errorf("Error reading fw_session_hairpin: %v", err)
		}
	}

	if err = d.Set("prp_trailer_action", flattenSystemSettingsPrpTrailerAction(o["prp-trailer-action"], d, "prp_trailer_action", sv)); err != nil {
		if !fortiAPIPatch(o["prp-trailer-action"]) {
			return fmt.Errorf("Error reading prp_trailer_action: %v", err)
		}
	}

	if err = d.Set("snat_hairpin_traffic", flattenSystemSettingsSnatHairpinTraffic(o["snat-hairpin-traffic"], d, "snat_hairpin_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["snat-hairpin-traffic"]) {
			return fmt.Errorf("Error reading snat_hairpin_traffic: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy", flattenSystemSettingsDhcpProxy(o["dhcp-proxy"], d, "dhcp_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-proxy"]) {
			return fmt.Errorf("Error reading dhcp_proxy: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy_interface_select_method", flattenSystemSettingsDhcpProxyInterfaceSelectMethod(o["dhcp-proxy-interface-select-method"], d, "dhcp_proxy_interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-proxy-interface-select-method"]) {
			return fmt.Errorf("Error reading dhcp_proxy_interface_select_method: %v", err)
		}
	}

	if err = d.Set("dhcp_proxy_interface", flattenSystemSettingsDhcpProxyInterface(o["dhcp-proxy-interface"], d, "dhcp_proxy_interface", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-proxy-interface"]) {
			return fmt.Errorf("Error reading dhcp_proxy_interface: %v", err)
		}
	}

	if err = d.Set("dhcp_server_ip", flattenSystemSettingsDhcpServerIp(o["dhcp-server-ip"], d, "dhcp_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-server-ip"]) {
			return fmt.Errorf("Error reading dhcp_server_ip: %v", err)
		}
	}

	if err = d.Set("dhcp6_server_ip", flattenSystemSettingsDhcp6ServerIp(o["dhcp6-server-ip"], d, "dhcp6_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp6-server-ip"]) {
			return fmt.Errorf("Error reading dhcp6_server_ip: %v", err)
		}
	}

	if err = d.Set("central_nat", flattenSystemSettingsCentralNat(o["central-nat"], d, "central_nat", sv)); err != nil {
		if !fortiAPIPatch(o["central-nat"]) {
			return fmt.Errorf("Error reading central_nat: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("gui_default_policy_columns", flattenSystemSettingsGuiDefaultPolicyColumns(o["gui-default-policy-columns"], d, "gui_default_policy_columns", sv)); err != nil {
			if !fortiAPIPatch(o["gui-default-policy-columns"]) {
				return fmt.Errorf("Error reading gui_default_policy_columns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gui_default_policy_columns"); ok {
			if err = d.Set("gui_default_policy_columns", flattenSystemSettingsGuiDefaultPolicyColumns(o["gui-default-policy-columns"], d, "gui_default_policy_columns", sv)); err != nil {
				if !fortiAPIPatch(o["gui-default-policy-columns"]) {
					return fmt.Errorf("Error reading gui_default_policy_columns: %v", err)
				}
			}
		}
	}

	if err = d.Set("lldp_reception", flattenSystemSettingsLldpReception(o["lldp-reception"], d, "lldp_reception", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-reception"]) {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", flattenSystemSettingsLldpTransmission(o["lldp-transmission"], d, "lldp_transmission", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-transmission"]) {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("link_down_access", flattenSystemSettingsLinkDownAccess(o["link-down-access"], d, "link_down_access", sv)); err != nil {
		if !fortiAPIPatch(o["link-down-access"]) {
			return fmt.Errorf("Error reading link_down_access: %v", err)
		}
	}

	if err = d.Set("auxiliary_session", flattenSystemSettingsAuxiliarySession(o["auxiliary-session"], d, "auxiliary_session", sv)); err != nil {
		if !fortiAPIPatch(o["auxiliary-session"]) {
			return fmt.Errorf("Error reading auxiliary_session: %v", err)
		}
	}

	if err = d.Set("asymroute", flattenSystemSettingsAsymroute(o["asymroute"], d, "asymroute", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute"]) {
			return fmt.Errorf("Error reading asymroute: %v", err)
		}
	}

	if err = d.Set("asymroute_icmp", flattenSystemSettingsAsymrouteIcmp(o["asymroute-icmp"], d, "asymroute_icmp", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute-icmp"]) {
			return fmt.Errorf("Error reading asymroute_icmp: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenSystemSettingsTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-session-without-syn"]) {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("ses_denied_traffic", flattenSystemSettingsSesDeniedTraffic(o["ses-denied-traffic"], d, "ses_denied_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ses-denied-traffic"]) {
			return fmt.Errorf("Error reading ses_denied_traffic: %v", err)
		}
	}

	if err = d.Set("strict_src_check", flattenSystemSettingsStrictSrcCheck(o["strict-src-check"], d, "strict_src_check", sv)); err != nil {
		if !fortiAPIPatch(o["strict-src-check"]) {
			return fmt.Errorf("Error reading strict_src_check: %v", err)
		}
	}

	if err = d.Set("allow_linkdown_path", flattenSystemSettingsAllowLinkdownPath(o["allow-linkdown-path"], d, "allow_linkdown_path", sv)); err != nil {
		if !fortiAPIPatch(o["allow-linkdown-path"]) {
			return fmt.Errorf("Error reading allow_linkdown_path: %v", err)
		}
	}

	if err = d.Set("asymroute6", flattenSystemSettingsAsymroute6(o["asymroute6"], d, "asymroute6", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute6"]) {
			return fmt.Errorf("Error reading asymroute6: %v", err)
		}
	}

	if err = d.Set("asymroute6_icmp", flattenSystemSettingsAsymroute6Icmp(o["asymroute6-icmp"], d, "asymroute6_icmp", sv)); err != nil {
		if !fortiAPIPatch(o["asymroute6-icmp"]) {
			return fmt.Errorf("Error reading asymroute6_icmp: %v", err)
		}
	}

	if err = d.Set("sctp_session_without_init", flattenSystemSettingsSctpSessionWithoutInit(o["sctp-session-without-init"], d, "sctp_session_without_init", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-session-without-init"]) {
			return fmt.Errorf("Error reading sctp_session_without_init: %v", err)
		}
	}

	if err = d.Set("sip_expectation", flattenSystemSettingsSipExpectation(o["sip-expectation"], d, "sip_expectation", sv)); err != nil {
		if !fortiAPIPatch(o["sip-expectation"]) {
			return fmt.Errorf("Error reading sip_expectation: %v", err)
		}
	}

	if err = d.Set("sip_helper", flattenSystemSettingsSipHelper(o["sip-helper"], d, "sip_helper", sv)); err != nil {
		if !fortiAPIPatch(o["sip-helper"]) {
			return fmt.Errorf("Error reading sip_helper: %v", err)
		}
	}

	if err = d.Set("sip_nat_trace", flattenSystemSettingsSipNatTrace(o["sip-nat-trace"], d, "sip_nat_trace", sv)); err != nil {
		if !fortiAPIPatch(o["sip-nat-trace"]) {
			return fmt.Errorf("Error reading sip_nat_trace: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("sip_tcp_port", flattenSystemSettingsSipTcpPort(o["sip-tcp-port"], d, "sip_tcp_port", sv)); err != nil {
		if !fortiAPIPatch(o["sip-tcp-port"]) {
			return fmt.Errorf("Error reading sip_tcp_port: %v", err)
		}
	}

	if err = d.Set("sip_udp_port", flattenSystemSettingsSipUdpPort(o["sip-udp-port"], d, "sip_udp_port", sv)); err != nil {
		if !fortiAPIPatch(o["sip-udp-port"]) {
			return fmt.Errorf("Error reading sip_udp_port: %v", err)
		}
	}

	if err = d.Set("sip_ssl_port", flattenSystemSettingsSipSslPort(o["sip-ssl-port"], d, "sip_ssl_port", sv)); err != nil {
		if !fortiAPIPatch(o["sip-ssl-port"]) {
			return fmt.Errorf("Error reading sip_ssl_port: %v", err)
		}
	}

	if err = d.Set("sccp_port", flattenSystemSettingsSccpPort(o["sccp-port"], d, "sccp_port", sv)); err != nil {
		if !fortiAPIPatch(o["sccp-port"]) {
			return fmt.Errorf("Error reading sccp_port: %v", err)
		}
	}

	if err = d.Set("multicast_forward", flattenSystemSettingsMulticastForward(o["multicast-forward"], d, "multicast_forward", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-forward"]) {
			return fmt.Errorf("Error reading multicast_forward: %v", err)
		}
	}

	if err = d.Set("multicast_ttl_notchange", flattenSystemSettingsMulticastTtlNotchange(o["multicast-ttl-notchange"], d, "multicast_ttl_notchange", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-ttl-notchange"]) {
			return fmt.Errorf("Error reading multicast_ttl_notchange: %v", err)
		}
	}

	if err = d.Set("multicast_skip_policy", flattenSystemSettingsMulticastSkipPolicy(o["multicast-skip-policy"], d, "multicast_skip_policy", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-skip-policy"]) {
			return fmt.Errorf("Error reading multicast_skip_policy: %v", err)
		}
	}

	if err = d.Set("allow_subnet_overlap", flattenSystemSettingsAllowSubnetOverlap(o["allow-subnet-overlap"], d, "allow_subnet_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["allow-subnet-overlap"]) {
			return fmt.Errorf("Error reading allow_subnet_overlap: %v", err)
		}
	}

	if err = d.Set("deny_tcp_with_icmp", flattenSystemSettingsDenyTcpWithIcmp(o["deny-tcp-with-icmp"], d, "deny_tcp_with_icmp", sv)); err != nil {
		if !fortiAPIPatch(o["deny-tcp-with-icmp"]) {
			return fmt.Errorf("Error reading deny_tcp_with_icmp: %v", err)
		}
	}

	if err = d.Set("ecmp_max_paths", flattenSystemSettingsEcmpMaxPaths(o["ecmp-max-paths"], d, "ecmp_max_paths", sv)); err != nil {
		if !fortiAPIPatch(o["ecmp-max-paths"]) {
			return fmt.Errorf("Error reading ecmp_max_paths: %v", err)
		}
	}

	if err = d.Set("discovered_device_timeout", flattenSystemSettingsDiscoveredDeviceTimeout(o["discovered-device-timeout"], d, "discovered_device_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["discovered-device-timeout"]) {
			return fmt.Errorf("Error reading discovered_device_timeout: %v", err)
		}
	}

	if err = d.Set("email_portal_check_dns", flattenSystemSettingsEmailPortalCheckDns(o["email-portal-check-dns"], d, "email_portal_check_dns", sv)); err != nil {
		if !fortiAPIPatch(o["email-portal-check-dns"]) {
			return fmt.Errorf("Error reading email_portal_check_dns: %v", err)
		}
	}

	if err = d.Set("default_voip_alg_mode", flattenSystemSettingsDefaultVoipAlgMode(o["default-voip-alg-mode"], d, "default_voip_alg_mode", sv)); err != nil {
		if !fortiAPIPatch(o["default-voip-alg-mode"]) {
			return fmt.Errorf("Error reading default_voip_alg_mode: %v", err)
		}
	}

	if err = d.Set("gui_icap", flattenSystemSettingsGuiIcap(o["gui-icap"], d, "gui_icap", sv)); err != nil {
		if !fortiAPIPatch(o["gui-icap"]) {
			return fmt.Errorf("Error reading gui_icap: %v", err)
		}
	}

	if err = d.Set("gui_nat46_64", flattenSystemSettingsGuiNat4664(o["gui-nat46-64"], d, "gui_nat46_64", sv)); err != nil {
		if !fortiAPIPatch(o["gui-nat46-64"]) {
			return fmt.Errorf("Error reading gui_nat46_64: %v", err)
		}
	}

	if err = d.Set("gui_implicit_policy", flattenSystemSettingsGuiImplicitPolicy(o["gui-implicit-policy"], d, "gui_implicit_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-implicit-policy"]) {
			return fmt.Errorf("Error reading gui_implicit_policy: %v", err)
		}
	}

	if err = d.Set("gui_dns_database", flattenSystemSettingsGuiDnsDatabase(o["gui-dns-database"], d, "gui_dns_database", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dns-database"]) {
			return fmt.Errorf("Error reading gui_dns_database: %v", err)
		}
	}

	if err = d.Set("gui_load_balance", flattenSystemSettingsGuiLoadBalance(o["gui-load-balance"], d, "gui_load_balance", sv)); err != nil {
		if !fortiAPIPatch(o["gui-load-balance"]) {
			return fmt.Errorf("Error reading gui_load_balance: %v", err)
		}
	}

	if err = d.Set("gui_multicast_policy", flattenSystemSettingsGuiMulticastPolicy(o["gui-multicast-policy"], d, "gui_multicast_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-multicast-policy"]) {
			return fmt.Errorf("Error reading gui_multicast_policy: %v", err)
		}
	}

	if err = d.Set("gui_dos_policy", flattenSystemSettingsGuiDosPolicy(o["gui-dos-policy"], d, "gui_dos_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dos-policy"]) {
			return fmt.Errorf("Error reading gui_dos_policy: %v", err)
		}
	}

	if err = d.Set("gui_object_colors", flattenSystemSettingsGuiObjectColors(o["gui-object-colors"], d, "gui_object_colors", sv)); err != nil {
		if !fortiAPIPatch(o["gui-object-colors"]) {
			return fmt.Errorf("Error reading gui_object_colors: %v", err)
		}
	}

	if err = d.Set("gui_replacement_message_groups", flattenSystemSettingsGuiReplacementMessageGroups(o["gui-replacement-message-groups"], d, "gui_replacement_message_groups", sv)); err != nil {
		if !fortiAPIPatch(o["gui-replacement-message-groups"]) {
			return fmt.Errorf("Error reading gui_replacement_message_groups: %v", err)
		}
	}

	if err = d.Set("gui_voip_profile", flattenSystemSettingsGuiVoipProfile(o["gui-voip-profile"], d, "gui_voip_profile", sv)); err != nil {
		if !fortiAPIPatch(o["gui-voip-profile"]) {
			return fmt.Errorf("Error reading gui_voip_profile: %v", err)
		}
	}

	if err = d.Set("gui_ap_profile", flattenSystemSettingsGuiApProfile(o["gui-ap-profile"], d, "gui_ap_profile", sv)); err != nil {
		if !fortiAPIPatch(o["gui-ap-profile"]) {
			return fmt.Errorf("Error reading gui_ap_profile: %v", err)
		}
	}

	if err = d.Set("gui_security_profile_group", flattenSystemSettingsGuiSecurityProfileGroup(o["gui-security-profile-group"], d, "gui_security_profile_group", sv)); err != nil {
		if !fortiAPIPatch(o["gui-security-profile-group"]) {
			return fmt.Errorf("Error reading gui_security_profile_group: %v", err)
		}
	}

	if err = d.Set("gui_dynamic_profile_display", flattenSystemSettingsGuiDynamicProfileDisplay(o["gui-dynamic-profile-display"], d, "gui_dynamic_profile_display", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dynamic-profile-display"]) {
			return fmt.Errorf("Error reading gui_dynamic_profile_display: %v", err)
		}
	}

	if err = d.Set("gui_local_in_policy", flattenSystemSettingsGuiLocalInPolicy(o["gui-local-in-policy"], d, "gui_local_in_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-local-in-policy"]) {
			return fmt.Errorf("Error reading gui_local_in_policy: %v", err)
		}
	}

	if err = d.Set("gui_local_reports", flattenSystemSettingsGuiLocalReports(o["gui-local-reports"], d, "gui_local_reports", sv)); err != nil {
		if !fortiAPIPatch(o["gui-local-reports"]) {
			return fmt.Errorf("Error reading gui_local_reports: %v", err)
		}
	}

	if err = d.Set("gui_wanopt_cache", flattenSystemSettingsGuiWanoptCache(o["gui-wanopt-cache"], d, "gui_wanopt_cache", sv)); err != nil {
		if !fortiAPIPatch(o["gui-wanopt-cache"]) {
			return fmt.Errorf("Error reading gui_wanopt_cache: %v", err)
		}
	}

	if err = d.Set("gui_explicit_proxy", flattenSystemSettingsGuiExplicitProxy(o["gui-explicit-proxy"], d, "gui_explicit_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-explicit-proxy"]) {
			return fmt.Errorf("Error reading gui_explicit_proxy: %v", err)
		}
	}

	if err = d.Set("gui_dynamic_routing", flattenSystemSettingsGuiDynamicRouting(o["gui-dynamic-routing"], d, "gui_dynamic_routing", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dynamic-routing"]) {
			return fmt.Errorf("Error reading gui_dynamic_routing: %v", err)
		}
	}

	if err = d.Set("gui_dlp", flattenSystemSettingsGuiDlp(o["gui-dlp"], d, "gui_dlp", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dlp"]) {
			return fmt.Errorf("Error reading gui_dlp: %v", err)
		}
	}

	if err = d.Set("gui_sslvpn_personal_bookmarks", flattenSystemSettingsGuiSslvpnPersonalBookmarks(o["gui-sslvpn-personal-bookmarks"], d, "gui_sslvpn_personal_bookmarks", sv)); err != nil {
		if !fortiAPIPatch(o["gui-sslvpn-personal-bookmarks"]) {
			return fmt.Errorf("Error reading gui_sslvpn_personal_bookmarks: %v", err)
		}
	}

	if err = d.Set("gui_sslvpn_realms", flattenSystemSettingsGuiSslvpnRealms(o["gui-sslvpn-realms"], d, "gui_sslvpn_realms", sv)); err != nil {
		if !fortiAPIPatch(o["gui-sslvpn-realms"]) {
			return fmt.Errorf("Error reading gui_sslvpn_realms: %v", err)
		}
	}

	if err = d.Set("gui_policy_based_ipsec", flattenSystemSettingsGuiPolicyBasedIpsec(o["gui-policy-based-ipsec"], d, "gui_policy_based_ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["gui-policy-based-ipsec"]) {
			return fmt.Errorf("Error reading gui_policy_based_ipsec: %v", err)
		}
	}

	if err = d.Set("gui_threat_weight", flattenSystemSettingsGuiThreatWeight(o["gui-threat-weight"], d, "gui_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["gui-threat-weight"]) {
			return fmt.Errorf("Error reading gui_threat_weight: %v", err)
		}
	}

	if err = d.Set("gui_multiple_utm_profiles", flattenSystemSettingsGuiMultipleUtmProfiles(o["gui-multiple-utm-profiles"], d, "gui_multiple_utm_profiles", sv)); err != nil {
		if !fortiAPIPatch(o["gui-multiple-utm-profiles"]) {
			return fmt.Errorf("Error reading gui_multiple_utm_profiles: %v", err)
		}
	}

	if err = d.Set("gui_spamfilter", flattenSystemSettingsGuiSpamfilter(o["gui-spamfilter"], d, "gui_spamfilter", sv)); err != nil {
		if !fortiAPIPatch(o["gui-spamfilter"]) {
			return fmt.Errorf("Error reading gui_spamfilter: %v", err)
		}
	}

	if err = d.Set("gui_file_filter", flattenSystemSettingsGuiFileFilter(o["gui-file-filter"], d, "gui_file_filter", sv)); err != nil {
		if !fortiAPIPatch(o["gui-file-filter"]) {
			return fmt.Errorf("Error reading gui_file_filter: %v", err)
		}
	}

	if err = d.Set("gui_application_control", flattenSystemSettingsGuiApplicationControl(o["gui-application-control"], d, "gui_application_control", sv)); err != nil {
		if !fortiAPIPatch(o["gui-application-control"]) {
			return fmt.Errorf("Error reading gui_application_control: %v", err)
		}
	}

	if err = d.Set("gui_ips", flattenSystemSettingsGuiIps(o["gui-ips"], d, "gui_ips", sv)); err != nil {
		if !fortiAPIPatch(o["gui-ips"]) {
			return fmt.Errorf("Error reading gui_ips: %v", err)
		}
	}

	if err = d.Set("gui_endpoint_control", flattenSystemSettingsGuiEndpointControl(o["gui-endpoint-control"], d, "gui_endpoint_control", sv)); err != nil {
		if !fortiAPIPatch(o["gui-endpoint-control"]) {
			return fmt.Errorf("Error reading gui_endpoint_control: %v", err)
		}
	}

	if err = d.Set("gui_endpoint_control_advanced", flattenSystemSettingsGuiEndpointControlAdvanced(o["gui-endpoint-control-advanced"], d, "gui_endpoint_control_advanced", sv)); err != nil {
		if !fortiAPIPatch(o["gui-endpoint-control-advanced"]) {
			return fmt.Errorf("Error reading gui_endpoint_control_advanced: %v", err)
		}
	}

	if err = d.Set("gui_dhcp_advanced", flattenSystemSettingsGuiDhcpAdvanced(o["gui-dhcp-advanced"], d, "gui_dhcp_advanced", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dhcp-advanced"]) {
			return fmt.Errorf("Error reading gui_dhcp_advanced: %v", err)
		}
	}

	if err = d.Set("gui_vpn", flattenSystemSettingsGuiVpn(o["gui-vpn"], d, "gui_vpn", sv)); err != nil {
		if !fortiAPIPatch(o["gui-vpn"]) {
			return fmt.Errorf("Error reading gui_vpn: %v", err)
		}
	}

	if err = d.Set("gui_wireless_controller", flattenSystemSettingsGuiWirelessController(o["gui-wireless-controller"], d, "gui_wireless_controller", sv)); err != nil {
		if !fortiAPIPatch(o["gui-wireless-controller"]) {
			return fmt.Errorf("Error reading gui_wireless_controller: %v", err)
		}
	}

	if err = d.Set("gui_switch_controller", flattenSystemSettingsGuiSwitchController(o["gui-switch-controller"], d, "gui_switch_controller", sv)); err != nil {
		if !fortiAPIPatch(o["gui-switch-controller"]) {
			return fmt.Errorf("Error reading gui_switch_controller: %v", err)
		}
	}

	if err = d.Set("gui_fortiap_split_tunneling", flattenSystemSettingsGuiFortiapSplitTunneling(o["gui-fortiap-split-tunneling"], d, "gui_fortiap_split_tunneling", sv)); err != nil {
		if !fortiAPIPatch(o["gui-fortiap-split-tunneling"]) {
			return fmt.Errorf("Error reading gui_fortiap_split_tunneling: %v", err)
		}
	}

	if err = d.Set("gui_webfilter_advanced", flattenSystemSettingsGuiWebfilterAdvanced(o["gui-webfilter-advanced"], d, "gui_webfilter_advanced", sv)); err != nil {
		if !fortiAPIPatch(o["gui-webfilter-advanced"]) {
			return fmt.Errorf("Error reading gui_webfilter_advanced: %v", err)
		}
	}

	if err = d.Set("gui_traffic_shaping", flattenSystemSettingsGuiTrafficShaping(o["gui-traffic-shaping"], d, "gui_traffic_shaping", sv)); err != nil {
		if !fortiAPIPatch(o["gui-traffic-shaping"]) {
			return fmt.Errorf("Error reading gui_traffic_shaping: %v", err)
		}
	}

	if err = d.Set("gui_wan_load_balancing", flattenSystemSettingsGuiWanLoadBalancing(o["gui-wan-load-balancing"], d, "gui_wan_load_balancing", sv)); err != nil {
		if !fortiAPIPatch(o["gui-wan-load-balancing"]) {
			return fmt.Errorf("Error reading gui_wan_load_balancing: %v", err)
		}
	}

	if err = d.Set("gui_antivirus", flattenSystemSettingsGuiAntivirus(o["gui-antivirus"], d, "gui_antivirus", sv)); err != nil {
		if !fortiAPIPatch(o["gui-antivirus"]) {
			return fmt.Errorf("Error reading gui_antivirus: %v", err)
		}
	}

	if err = d.Set("gui_webfilter", flattenSystemSettingsGuiWebfilter(o["gui-webfilter"], d, "gui_webfilter", sv)); err != nil {
		if !fortiAPIPatch(o["gui-webfilter"]) {
			return fmt.Errorf("Error reading gui_webfilter: %v", err)
		}
	}

	if err = d.Set("gui_dnsfilter", flattenSystemSettingsGuiDnsfilter(o["gui-dnsfilter"], d, "gui_dnsfilter", sv)); err != nil {
		if !fortiAPIPatch(o["gui-dnsfilter"]) {
			return fmt.Errorf("Error reading gui_dnsfilter: %v", err)
		}
	}

	if err = d.Set("gui_waf_profile", flattenSystemSettingsGuiWafProfile(o["gui-waf-profile"], d, "gui_waf_profile", sv)); err != nil {
		if !fortiAPIPatch(o["gui-waf-profile"]) {
			return fmt.Errorf("Error reading gui_waf_profile: %v", err)
		}
	}

	if err = d.Set("gui_fortiextender_controller", flattenSystemSettingsGuiFortiextenderController(o["gui-fortiextender-controller"], d, "gui_fortiextender_controller", sv)); err != nil {
		if !fortiAPIPatch(o["gui-fortiextender-controller"]) {
			return fmt.Errorf("Error reading gui_fortiextender_controller: %v", err)
		}
	}

	if err = d.Set("gui_advanced_policy", flattenSystemSettingsGuiAdvancedPolicy(o["gui-advanced-policy"], d, "gui_advanced_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-advanced-policy"]) {
			return fmt.Errorf("Error reading gui_advanced_policy: %v", err)
		}
	}

	if err = d.Set("gui_allow_unnamed_policy", flattenSystemSettingsGuiAllowUnnamedPolicy(o["gui-allow-unnamed-policy"], d, "gui_allow_unnamed_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-allow-unnamed-policy"]) {
			return fmt.Errorf("Error reading gui_allow_unnamed_policy: %v", err)
		}
	}

	if err = d.Set("gui_email_collection", flattenSystemSettingsGuiEmailCollection(o["gui-email-collection"], d, "gui_email_collection", sv)); err != nil {
		if !fortiAPIPatch(o["gui-email-collection"]) {
			return fmt.Errorf("Error reading gui_email_collection: %v", err)
		}
	}

	if err = d.Set("gui_domain_ip_reputation", flattenSystemSettingsGuiDomainIpReputation(o["gui-domain-ip-reputation"], d, "gui_domain_ip_reputation", sv)); err != nil {
		if !fortiAPIPatch(o["gui-domain-ip-reputation"]) {
			return fmt.Errorf("Error reading gui_domain_ip_reputation: %v", err)
		}
	}

	if err = d.Set("gui_multiple_interface_policy", flattenSystemSettingsGuiMultipleInterfacePolicy(o["gui-multiple-interface-policy"], d, "gui_multiple_interface_policy", sv)); err != nil {
		if !fortiAPIPatch(o["gui-multiple-interface-policy"]) {
			return fmt.Errorf("Error reading gui_multiple_interface_policy: %v", err)
		}
	}

	if err = d.Set("gui_policy_disclaimer", flattenSystemSettingsGuiPolicyDisclaimer(o["gui-policy-disclaimer"], d, "gui_policy_disclaimer", sv)); err != nil {
		if !fortiAPIPatch(o["gui-policy-disclaimer"]) {
			return fmt.Errorf("Error reading gui_policy_disclaimer: %v", err)
		}
	}

	if err = d.Set("gui_per_policy_disclaimer", flattenSystemSettingsGuiPerPolicyDisclaimer(o["gui-per-policy-disclaimer"], d, "gui_per_policy_disclaimer", sv)); err != nil {
		if !fortiAPIPatch(o["gui-per-policy-disclaimer"]) {
			return fmt.Errorf("Error reading gui_per_policy_disclaimer: %v", err)
		}
	}

	if err = d.Set("gui_policy_learning", flattenSystemSettingsGuiPolicyLearning(o["gui-policy-learning"], d, "gui_policy_learning", sv)); err != nil {
		if !fortiAPIPatch(o["gui-policy-learning"]) {
			return fmt.Errorf("Error reading gui_policy_learning: %v", err)
		}
	}

	if err = d.Set("compliance_check", flattenSystemSettingsComplianceCheck(o["compliance-check"], d, "compliance_check", sv)); err != nil {
		if !fortiAPIPatch(o["compliance-check"]) {
			return fmt.Errorf("Error reading compliance_check: %v", err)
		}
	}

	if err = d.Set("ike_session_resume", flattenSystemSettingsIkeSessionResume(o["ike-session-resume"], d, "ike_session_resume", sv)); err != nil {
		if !fortiAPIPatch(o["ike-session-resume"]) {
			return fmt.Errorf("Error reading ike_session_resume: %v", err)
		}
	}

	if err = d.Set("ike_quick_crash_detect", flattenSystemSettingsIkeQuickCrashDetect(o["ike-quick-crash-detect"], d, "ike_quick_crash_detect", sv)); err != nil {
		if !fortiAPIPatch(o["ike-quick-crash-detect"]) {
			return fmt.Errorf("Error reading ike_quick_crash_detect: %v", err)
		}
	}

	if err = d.Set("ike_dn_format", flattenSystemSettingsIkeDnFormat(o["ike-dn-format"], d, "ike_dn_format", sv)); err != nil {
		if !fortiAPIPatch(o["ike-dn-format"]) {
			return fmt.Errorf("Error reading ike_dn_format: %v", err)
		}
	}

	if err = d.Set("ike_port", flattenSystemSettingsIkePort(o["ike-port"], d, "ike_port", sv)); err != nil {
		if !fortiAPIPatch(o["ike-port"]) {
			return fmt.Errorf("Error reading ike_port: %v", err)
		}
	}

	if err = d.Set("ike_natt_port", flattenSystemSettingsIkeNattPort(o["ike-natt-port"], d, "ike_natt_port", sv)); err != nil {
		if !fortiAPIPatch(o["ike-natt-port"]) {
			return fmt.Errorf("Error reading ike_natt_port: %v", err)
		}
	}

	if err = d.Set("block_land_attack", flattenSystemSettingsBlockLandAttack(o["block-land-attack"], d, "block_land_attack", sv)); err != nil {
		if !fortiAPIPatch(o["block-land-attack"]) {
			return fmt.Errorf("Error reading block_land_attack: %v", err)
		}
	}

	return nil
}

func flattenSystemSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSettingsComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsOpmode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsInspectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsNgfwMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsImplicitAllowDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSslSshProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsConsolidatedFirewallMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsHttpExternalDest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsManageip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsManageip6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDetectMult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBfdDontEnforceSrcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsUtf8SpamTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsWccpCacheEngine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsVpnStatsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsVpnStatsPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsV4EcmpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMacTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsFwSessionHairpin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsPrpTrailerAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSnatHairpinTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpProxyInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpProxyInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcpServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDhcp6ServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsCentralNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDefaultPolicyColumns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSettingsGuiDefaultPolicyColumnsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSettingsGuiDefaultPolicyColumnsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLldpReception(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLldpTransmission(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsLinkDownAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAuxiliarySession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymrouteIcmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsTcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSesDeniedTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsStrictSrcCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAllowLinkdownPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAsymroute6Icmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSctpSessionWithoutInit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipExpectation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipHelper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipNatTrace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipTcpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipUdpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSipSslPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsSccpPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastTtlNotchange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsMulticastSkipPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsAllowSubnetOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDenyTcpWithIcmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsEcmpMaxPaths(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDiscoveredDeviceTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsEmailPortalCheckDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsDefaultVoipAlgMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiIcap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiNat4664(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiImplicitPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDnsDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiLoadBalance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiMulticastPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDosPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiObjectColors(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiReplacementMessageGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiVoipProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiApProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSecurityProfileGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDynamicProfileDisplay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiLocalInPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiLocalReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWanoptCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiExplicitProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDynamicRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDlp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSslvpnPersonalBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSslvpnRealms(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPolicyBasedIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiThreatWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiMultipleUtmProfiles(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSpamfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiApplicationControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiIps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEndpointControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEndpointControlAdvanced(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDhcpAdvanced(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiVpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWirelessController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiSwitchController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiFortiapSplitTunneling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWebfilterAdvanced(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiTrafficShaping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWanLoadBalancing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAntivirus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWebfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDnsfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiWafProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiFortiextenderController(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAdvancedPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiAllowUnnamedPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiEmailCollection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiDomainIpReputation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiMultipleInterfacePolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPerPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsGuiPolicyLearning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsComplianceCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeSessionResume(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeQuickCrashDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeDnFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsIkeNattPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSettingsBlockLandAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandSystemSettingsComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("opmode"); ok {

		t, err := expandSystemSettingsOpmode(d, v, "opmode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opmode"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {

		t, err := expandSystemSettingsInspectionMode(d, v, "inspection_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("ngfw_mode"); ok {

		t, err := expandSystemSettingsNgfwMode(d, v, "ngfw_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ngfw-mode"] = t
		}
	}

	if v, ok := d.GetOk("implicit_allow_dns"); ok {

		t, err := expandSystemSettingsImplicitAllowDns(d, v, "implicit_allow_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["implicit-allow-dns"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {

		t, err := expandSystemSettingsSslSshProfile(d, v, "ssl_ssh_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("consolidated_firewall_mode"); ok {

		t, err := expandSystemSettingsConsolidatedFirewallMode(d, v, "consolidated_firewall_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["consolidated-firewall-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_external_dest"); ok {

		t, err := expandSystemSettingsHttpExternalDest(d, v, "http_external_dest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-external-dest"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok {

		t, err := expandSystemSettingsFirewallSessionDirty(d, v, "firewall_session_dirty", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("manageip"); ok {

		t, err := expandSystemSettingsManageip(d, v, "manageip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manageip"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {

		t, err := expandSystemSettingsGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandSystemSettingsIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("manageip6"); ok {

		t, err := expandSystemSettingsManageip6(d, v, "manageip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["manageip6"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok {

		t, err := expandSystemSettingsGateway6(d, v, "gateway6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {

		t, err := expandSystemSettingsIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {

		t, err := expandSystemSettingsDevice(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {

		t, err := expandSystemSettingsBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("bfd_desired_min_tx"); ok {

		t, err := expandSystemSettingsBfdDesiredMinTx(d, v, "bfd_desired_min_tx", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-desired-min-tx"] = t
		}
	}

	if v, ok := d.GetOk("bfd_required_min_rx"); ok {

		t, err := expandSystemSettingsBfdRequiredMinRx(d, v, "bfd_required_min_rx", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-required-min-rx"] = t
		}
	}

	if v, ok := d.GetOk("bfd_detect_mult"); ok {

		t, err := expandSystemSettingsBfdDetectMult(d, v, "bfd_detect_mult", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-detect-mult"] = t
		}
	}

	if v, ok := d.GetOk("bfd_dont_enforce_src_port"); ok {

		t, err := expandSystemSettingsBfdDontEnforceSrcPort(d, v, "bfd_dont_enforce_src_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-dont-enforce-src-port"] = t
		}
	}

	if v, ok := d.GetOk("utf8_spam_tagging"); ok {

		t, err := expandSystemSettingsUtf8SpamTagging(d, v, "utf8_spam_tagging", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utf8-spam-tagging"] = t
		}
	}

	if v, ok := d.GetOk("wccp_cache_engine"); ok {

		t, err := expandSystemSettingsWccpCacheEngine(d, v, "wccp_cache_engine", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp-cache-engine"] = t
		}
	}

	if v, ok := d.GetOk("vpn_stats_log"); ok {

		t, err := expandSystemSettingsVpnStatsLog(d, v, "vpn_stats_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-stats-log"] = t
		}
	}

	if v, ok := d.GetOkExists("vpn_stats_period"); ok {

		t, err := expandSystemSettingsVpnStatsPeriod(d, v, "vpn_stats_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-stats-period"] = t
		}
	}

	if v, ok := d.GetOk("v4_ecmp_mode"); ok {

		t, err := expandSystemSettingsV4EcmpMode(d, v, "v4_ecmp_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["v4-ecmp-mode"] = t
		}
	}

	if v, ok := d.GetOk("mac_ttl"); ok {

		t, err := expandSystemSettingsMacTtl(d, v, "mac_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-ttl"] = t
		}
	}

	if v, ok := d.GetOk("fw_session_hairpin"); ok {

		t, err := expandSystemSettingsFwSessionHairpin(d, v, "fw_session_hairpin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fw-session-hairpin"] = t
		}
	}

	if v, ok := d.GetOk("prp_trailer_action"); ok {

		t, err := expandSystemSettingsPrpTrailerAction(d, v, "prp_trailer_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prp-trailer-action"] = t
		}
	}

	if v, ok := d.GetOk("snat_hairpin_traffic"); ok {

		t, err := expandSystemSettingsSnatHairpinTraffic(d, v, "snat_hairpin_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snat-hairpin-traffic"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy"); ok {

		t, err := expandSystemSettingsDhcpProxy(d, v, "dhcp_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy_interface_select_method"); ok {

		t, err := expandSystemSettingsDhcpProxyInterfaceSelectMethod(d, v, "dhcp_proxy_interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy-interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_proxy_interface"); ok {

		t, err := expandSystemSettingsDhcpProxyInterface(d, v, "dhcp_proxy_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-proxy-interface"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_ip"); ok {

		t, err := expandSystemSettingsDhcpServerIp(d, v, "dhcp_server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_server_ip"); ok {

		t, err := expandSystemSettingsDhcp6ServerIp(d, v, "dhcp6_server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("central_nat"); ok {

		t, err := expandSystemSettingsCentralNat(d, v, "central_nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["central-nat"] = t
		}
	}

	if v, ok := d.GetOk("gui_default_policy_columns"); ok {

		t, err := expandSystemSettingsGuiDefaultPolicyColumns(d, v, "gui_default_policy_columns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-default-policy-columns"] = t
		}
	}

	if v, ok := d.GetOk("lldp_reception"); ok {

		t, err := expandSystemSettingsLldpReception(d, v, "lldp_reception", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-reception"] = t
		}
	}

	if v, ok := d.GetOk("lldp_transmission"); ok {

		t, err := expandSystemSettingsLldpTransmission(d, v, "lldp_transmission", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-transmission"] = t
		}
	}

	if v, ok := d.GetOk("link_down_access"); ok {

		t, err := expandSystemSettingsLinkDownAccess(d, v, "link_down_access", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-access"] = t
		}
	}

	if v, ok := d.GetOk("auxiliary_session"); ok {

		t, err := expandSystemSettingsAuxiliarySession(d, v, "auxiliary_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auxiliary-session"] = t
		}
	}

	if v, ok := d.GetOk("asymroute"); ok {

		t, err := expandSystemSettingsAsymroute(d, v, "asymroute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute"] = t
		}
	}

	if v, ok := d.GetOk("asymroute_icmp"); ok {

		t, err := expandSystemSettingsAsymrouteIcmp(d, v, "asymroute_icmp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute-icmp"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok {

		t, err := expandSystemSettingsTcpSessionWithoutSyn(d, v, "tcp_session_without_syn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("ses_denied_traffic"); ok {

		t, err := expandSystemSettingsSesDeniedTraffic(d, v, "ses_denied_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ses-denied-traffic"] = t
		}
	}

	if v, ok := d.GetOk("strict_src_check"); ok {

		t, err := expandSystemSettingsStrictSrcCheck(d, v, "strict_src_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-src-check"] = t
		}
	}

	if v, ok := d.GetOk("allow_linkdown_path"); ok {

		t, err := expandSystemSettingsAllowLinkdownPath(d, v, "allow_linkdown_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-linkdown-path"] = t
		}
	}

	if v, ok := d.GetOk("asymroute6"); ok {

		t, err := expandSystemSettingsAsymroute6(d, v, "asymroute6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute6"] = t
		}
	}

	if v, ok := d.GetOk("asymroute6_icmp"); ok {

		t, err := expandSystemSettingsAsymroute6Icmp(d, v, "asymroute6_icmp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asymroute6-icmp"] = t
		}
	}

	if v, ok := d.GetOk("sctp_session_without_init"); ok {

		t, err := expandSystemSettingsSctpSessionWithoutInit(d, v, "sctp_session_without_init", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-session-without-init"] = t
		}
	}

	if v, ok := d.GetOk("sip_expectation"); ok {

		t, err := expandSystemSettingsSipExpectation(d, v, "sip_expectation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-expectation"] = t
		}
	}

	if v, ok := d.GetOk("sip_helper"); ok {

		t, err := expandSystemSettingsSipHelper(d, v, "sip_helper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-helper"] = t
		}
	}

	if v, ok := d.GetOk("sip_nat_trace"); ok {

		t, err := expandSystemSettingsSipNatTrace(d, v, "sip_nat_trace", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-nat-trace"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSystemSettingsStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("sip_tcp_port"); ok {

		t, err := expandSystemSettingsSipTcpPort(d, v, "sip_tcp_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-tcp-port"] = t
		}
	}

	if v, ok := d.GetOk("sip_udp_port"); ok {

		t, err := expandSystemSettingsSipUdpPort(d, v, "sip_udp_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-udp-port"] = t
		}
	}

	if v, ok := d.GetOkExists("sip_ssl_port"); ok {

		t, err := expandSystemSettingsSipSslPort(d, v, "sip_ssl_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-ssl-port"] = t
		}
	}

	if v, ok := d.GetOkExists("sccp_port"); ok {

		t, err := expandSystemSettingsSccpPort(d, v, "sccp_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sccp-port"] = t
		}
	}

	if v, ok := d.GetOk("multicast_forward"); ok {

		t, err := expandSystemSettingsMulticastForward(d, v, "multicast_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-forward"] = t
		}
	}

	if v, ok := d.GetOk("multicast_ttl_notchange"); ok {

		t, err := expandSystemSettingsMulticastTtlNotchange(d, v, "multicast_ttl_notchange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-ttl-notchange"] = t
		}
	}

	if v, ok := d.GetOk("multicast_skip_policy"); ok {

		t, err := expandSystemSettingsMulticastSkipPolicy(d, v, "multicast_skip_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-skip-policy"] = t
		}
	}

	if v, ok := d.GetOk("allow_subnet_overlap"); ok {

		t, err := expandSystemSettingsAllowSubnetOverlap(d, v, "allow_subnet_overlap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-subnet-overlap"] = t
		}
	}

	if v, ok := d.GetOk("deny_tcp_with_icmp"); ok {

		t, err := expandSystemSettingsDenyTcpWithIcmp(d, v, "deny_tcp_with_icmp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny-tcp-with-icmp"] = t
		}
	}

	if v, ok := d.GetOk("ecmp_max_paths"); ok {

		t, err := expandSystemSettingsEcmpMaxPaths(d, v, "ecmp_max_paths", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ecmp-max-paths"] = t
		}
	}

	if v, ok := d.GetOk("discovered_device_timeout"); ok {

		t, err := expandSystemSettingsDiscoveredDeviceTimeout(d, v, "discovered_device_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovered-device-timeout"] = t
		}
	}

	if v, ok := d.GetOk("email_portal_check_dns"); ok {

		t, err := expandSystemSettingsEmailPortalCheckDns(d, v, "email_portal_check_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-portal-check-dns"] = t
		}
	}

	if v, ok := d.GetOk("default_voip_alg_mode"); ok {

		t, err := expandSystemSettingsDefaultVoipAlgMode(d, v, "default_voip_alg_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-voip-alg-mode"] = t
		}
	}

	if v, ok := d.GetOk("gui_icap"); ok {

		t, err := expandSystemSettingsGuiIcap(d, v, "gui_icap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-icap"] = t
		}
	}

	if v, ok := d.GetOk("gui_nat46_64"); ok {

		t, err := expandSystemSettingsGuiNat4664(d, v, "gui_nat46_64", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-nat46-64"] = t
		}
	}

	if v, ok := d.GetOk("gui_implicit_policy"); ok {

		t, err := expandSystemSettingsGuiImplicitPolicy(d, v, "gui_implicit_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-implicit-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_dns_database"); ok {

		t, err := expandSystemSettingsGuiDnsDatabase(d, v, "gui_dns_database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dns-database"] = t
		}
	}

	if v, ok := d.GetOk("gui_load_balance"); ok {

		t, err := expandSystemSettingsGuiLoadBalance(d, v, "gui_load_balance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-load-balance"] = t
		}
	}

	if v, ok := d.GetOk("gui_multicast_policy"); ok {

		t, err := expandSystemSettingsGuiMulticastPolicy(d, v, "gui_multicast_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-multicast-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_dos_policy"); ok {

		t, err := expandSystemSettingsGuiDosPolicy(d, v, "gui_dos_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dos-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_object_colors"); ok {

		t, err := expandSystemSettingsGuiObjectColors(d, v, "gui_object_colors", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-object-colors"] = t
		}
	}

	if v, ok := d.GetOk("gui_replacement_message_groups"); ok {

		t, err := expandSystemSettingsGuiReplacementMessageGroups(d, v, "gui_replacement_message_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-replacement-message-groups"] = t
		}
	}

	if v, ok := d.GetOk("gui_voip_profile"); ok {

		t, err := expandSystemSettingsGuiVoipProfile(d, v, "gui_voip_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_ap_profile"); ok {

		t, err := expandSystemSettingsGuiApProfile(d, v, "gui_ap_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ap-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_security_profile_group"); ok {

		t, err := expandSystemSettingsGuiSecurityProfileGroup(d, v, "gui_security_profile_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-security-profile-group"] = t
		}
	}

	if v, ok := d.GetOk("gui_dynamic_profile_display"); ok {

		t, err := expandSystemSettingsGuiDynamicProfileDisplay(d, v, "gui_dynamic_profile_display", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dynamic-profile-display"] = t
		}
	}

	if v, ok := d.GetOk("gui_local_in_policy"); ok {

		t, err := expandSystemSettingsGuiLocalInPolicy(d, v, "gui_local_in_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-local-in-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_local_reports"); ok {

		t, err := expandSystemSettingsGuiLocalReports(d, v, "gui_local_reports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-local-reports"] = t
		}
	}

	if v, ok := d.GetOk("gui_wanopt_cache"); ok {

		t, err := expandSystemSettingsGuiWanoptCache(d, v, "gui_wanopt_cache", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wanopt-cache"] = t
		}
	}

	if v, ok := d.GetOk("gui_explicit_proxy"); ok {

		t, err := expandSystemSettingsGuiExplicitProxy(d, v, "gui_explicit_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-explicit-proxy"] = t
		}
	}

	if v, ok := d.GetOk("gui_dynamic_routing"); ok {

		t, err := expandSystemSettingsGuiDynamicRouting(d, v, "gui_dynamic_routing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dynamic-routing"] = t
		}
	}

	if v, ok := d.GetOk("gui_dlp"); ok {

		t, err := expandSystemSettingsGuiDlp(d, v, "gui_dlp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dlp"] = t
		}
	}

	if v, ok := d.GetOk("gui_sslvpn_personal_bookmarks"); ok {

		t, err := expandSystemSettingsGuiSslvpnPersonalBookmarks(d, v, "gui_sslvpn_personal_bookmarks", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-sslvpn-personal-bookmarks"] = t
		}
	}

	if v, ok := d.GetOk("gui_sslvpn_realms"); ok {

		t, err := expandSystemSettingsGuiSslvpnRealms(d, v, "gui_sslvpn_realms", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-sslvpn-realms"] = t
		}
	}

	if v, ok := d.GetOk("gui_policy_based_ipsec"); ok {

		t, err := expandSystemSettingsGuiPolicyBasedIpsec(d, v, "gui_policy_based_ipsec", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-policy-based-ipsec"] = t
		}
	}

	if v, ok := d.GetOk("gui_threat_weight"); ok {

		t, err := expandSystemSettingsGuiThreatWeight(d, v, "gui_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-threat-weight"] = t
		}
	}

	if v, ok := d.GetOk("gui_multiple_utm_profiles"); ok {

		t, err := expandSystemSettingsGuiMultipleUtmProfiles(d, v, "gui_multiple_utm_profiles", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-multiple-utm-profiles"] = t
		}
	}

	if v, ok := d.GetOk("gui_spamfilter"); ok {

		t, err := expandSystemSettingsGuiSpamfilter(d, v, "gui_spamfilter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-spamfilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_file_filter"); ok {

		t, err := expandSystemSettingsGuiFileFilter(d, v, "gui_file_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-file-filter"] = t
		}
	}

	if v, ok := d.GetOk("gui_application_control"); ok {

		t, err := expandSystemSettingsGuiApplicationControl(d, v, "gui_application_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-application-control"] = t
		}
	}

	if v, ok := d.GetOk("gui_ips"); ok {

		t, err := expandSystemSettingsGuiIps(d, v, "gui_ips", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ips"] = t
		}
	}

	if v, ok := d.GetOk("gui_endpoint_control"); ok {

		t, err := expandSystemSettingsGuiEndpointControl(d, v, "gui_endpoint_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-endpoint-control"] = t
		}
	}

	if v, ok := d.GetOk("gui_endpoint_control_advanced"); ok {

		t, err := expandSystemSettingsGuiEndpointControlAdvanced(d, v, "gui_endpoint_control_advanced", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-endpoint-control-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_dhcp_advanced"); ok {

		t, err := expandSystemSettingsGuiDhcpAdvanced(d, v, "gui_dhcp_advanced", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dhcp-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_vpn"); ok {

		t, err := expandSystemSettingsGuiVpn(d, v, "gui_vpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-vpn"] = t
		}
	}

	if v, ok := d.GetOk("gui_wireless_controller"); ok {

		t, err := expandSystemSettingsGuiWirelessController(d, v, "gui_wireless_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wireless-controller"] = t
		}
	}

	if v, ok := d.GetOk("gui_switch_controller"); ok {

		t, err := expandSystemSettingsGuiSwitchController(d, v, "gui_switch_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-switch-controller"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortiap_split_tunneling"); ok {

		t, err := expandSystemSettingsGuiFortiapSplitTunneling(d, v, "gui_fortiap_split_tunneling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortiap-split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("gui_webfilter_advanced"); ok {

		t, err := expandSystemSettingsGuiWebfilterAdvanced(d, v, "gui_webfilter_advanced", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-webfilter-advanced"] = t
		}
	}

	if v, ok := d.GetOk("gui_traffic_shaping"); ok {

		t, err := expandSystemSettingsGuiTrafficShaping(d, v, "gui_traffic_shaping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-traffic-shaping"] = t
		}
	}

	if v, ok := d.GetOk("gui_wan_load_balancing"); ok {

		t, err := expandSystemSettingsGuiWanLoadBalancing(d, v, "gui_wan_load_balancing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-wan-load-balancing"] = t
		}
	}

	if v, ok := d.GetOk("gui_antivirus"); ok {

		t, err := expandSystemSettingsGuiAntivirus(d, v, "gui_antivirus", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-antivirus"] = t
		}
	}

	if v, ok := d.GetOk("gui_webfilter"); ok {

		t, err := expandSystemSettingsGuiWebfilter(d, v, "gui_webfilter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-webfilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_dnsfilter"); ok {

		t, err := expandSystemSettingsGuiDnsfilter(d, v, "gui_dnsfilter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-dnsfilter"] = t
		}
	}

	if v, ok := d.GetOk("gui_waf_profile"); ok {

		t, err := expandSystemSettingsGuiWafProfile(d, v, "gui_waf_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("gui_fortiextender_controller"); ok {

		t, err := expandSystemSettingsGuiFortiextenderController(d, v, "gui_fortiextender_controller", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-fortiextender-controller"] = t
		}
	}

	if v, ok := d.GetOk("gui_advanced_policy"); ok {

		t, err := expandSystemSettingsGuiAdvancedPolicy(d, v, "gui_advanced_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-advanced-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_allow_unnamed_policy"); ok {

		t, err := expandSystemSettingsGuiAllowUnnamedPolicy(d, v, "gui_allow_unnamed_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-allow-unnamed-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_email_collection"); ok {

		t, err := expandSystemSettingsGuiEmailCollection(d, v, "gui_email_collection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-email-collection"] = t
		}
	}

	if v, ok := d.GetOk("gui_domain_ip_reputation"); ok {

		t, err := expandSystemSettingsGuiDomainIpReputation(d, v, "gui_domain_ip_reputation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-domain-ip-reputation"] = t
		}
	}

	if v, ok := d.GetOk("gui_multiple_interface_policy"); ok {

		t, err := expandSystemSettingsGuiMultipleInterfacePolicy(d, v, "gui_multiple_interface_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-multiple-interface-policy"] = t
		}
	}

	if v, ok := d.GetOk("gui_policy_disclaimer"); ok {

		t, err := expandSystemSettingsGuiPolicyDisclaimer(d, v, "gui_policy_disclaimer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-policy-disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("gui_per_policy_disclaimer"); ok {

		t, err := expandSystemSettingsGuiPerPolicyDisclaimer(d, v, "gui_per_policy_disclaimer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-per-policy-disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("gui_policy_learning"); ok {

		t, err := expandSystemSettingsGuiPolicyLearning(d, v, "gui_policy_learning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-policy-learning"] = t
		}
	}

	if v, ok := d.GetOk("compliance_check"); ok {

		t, err := expandSystemSettingsComplianceCheck(d, v, "compliance_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compliance-check"] = t
		}
	}

	if v, ok := d.GetOk("ike_session_resume"); ok {

		t, err := expandSystemSettingsIkeSessionResume(d, v, "ike_session_resume", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-session-resume"] = t
		}
	}

	if v, ok := d.GetOk("ike_quick_crash_detect"); ok {

		t, err := expandSystemSettingsIkeQuickCrashDetect(d, v, "ike_quick_crash_detect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-quick-crash-detect"] = t
		}
	}

	if v, ok := d.GetOk("ike_dn_format"); ok {

		t, err := expandSystemSettingsIkeDnFormat(d, v, "ike_dn_format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-dn-format"] = t
		}
	}

	if v, ok := d.GetOk("ike_port"); ok {

		t, err := expandSystemSettingsIkePort(d, v, "ike_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-port"] = t
		}
	}

	if v, ok := d.GetOk("ike_natt_port"); ok {

		t, err := expandSystemSettingsIkeNattPort(d, v, "ike_natt_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-natt-port"] = t
		}
	}

	if v, ok := d.GetOk("block_land_attack"); ok {

		t, err := expandSystemSettingsBlockLandAttack(d, v, "block_land_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-land-attack"] = t
		}
	}

	return &obj, nil
}
