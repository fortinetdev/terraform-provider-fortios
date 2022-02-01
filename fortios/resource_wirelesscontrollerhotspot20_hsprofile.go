// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure hotspot profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20HsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20HsProfileCreate,
		Read:   resourceWirelessControllerHotspot20HsProfileRead,
		Update: resourceWirelessControllerHotspot20HsProfileUpdate,
		Delete: resourceWirelessControllerHotspot20HsProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"release": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3),
				Optional:     true,
				Computed:     true,
			},
			"access_network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_internet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_asra": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_esr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_network_uesa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hessid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2tif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pame_bi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anqp_domain_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"domain_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"osu_ssid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"gas_comeback_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 10000),
				Optional:     true,
				Computed:     true,
			},
			"gas_fragmentation_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(512, 4096),
				Optional:     true,
				Computed:     true,
			},
			"dgaf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deauth_request_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 120),
				Optional:     true,
				Computed:     true,
			},
			"wnm_sleep_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bss_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"venue_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"venue_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"roaming_consortium": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"nai_realm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"oper_friendly_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"oper_icon": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"advice_of_charge": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"osu_provider_nai": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"terms_and_conditions": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"osu_provider": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"wan_metrics": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"network_auth": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"n3gpp_plmn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"conn_cap": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"qos_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ip_addr_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerHotspot20HsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20HsProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20HsProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20HsProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20HsProfile")
	}

	return resourceWirelessControllerHotspot20HsProfileRead(d, m)
}

func resourceWirelessControllerHotspot20HsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20HsProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20HsProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20HsProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20HsProfile")
	}

	return resourceWirelessControllerHotspot20HsProfileRead(d, m)
}

func resourceWirelessControllerHotspot20HsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20HsProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20HsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerHotspot20HsProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20HsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20HsProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20HsProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20HsProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileRelease(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkInternet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkAsra(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkEsr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAccessNetworkUesa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileVenueGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileVenueType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileHessid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileProxyArp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileL2Tif(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfilePameBi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAnqpDomainId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileOsuSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileGasComebackDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileDgaf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileDeauthRequestTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWnmSleepMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileBssTransition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileVenueName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileVenueUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileRoamingConsortium(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileNaiRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileOperFriendlyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileOperIcon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileAdviceOfCharge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileOsuProviderNai(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileTermsAndConditions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileOsuProvider(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerHotspot20HsProfileOsuProviderName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerHotspot20HsProfileOsuProviderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileWanMetrics(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileNetworkAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfile3GppPlmn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileConnCap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileQosMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20HsProfileIpAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20HsProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("release", flattenWirelessControllerHotspot20HsProfileRelease(o["release"], d, "release", sv)); err != nil {
		if !fortiAPIPatch(o["release"]) {
			return fmt.Errorf("Error reading release: %v", err)
		}
	}

	if err = d.Set("access_network_type", flattenWirelessControllerHotspot20HsProfileAccessNetworkType(o["access-network-type"], d, "access_network_type", sv)); err != nil {
		if !fortiAPIPatch(o["access-network-type"]) {
			return fmt.Errorf("Error reading access_network_type: %v", err)
		}
	}

	if err = d.Set("access_network_internet", flattenWirelessControllerHotspot20HsProfileAccessNetworkInternet(o["access-network-internet"], d, "access_network_internet", sv)); err != nil {
		if !fortiAPIPatch(o["access-network-internet"]) {
			return fmt.Errorf("Error reading access_network_internet: %v", err)
		}
	}

	if err = d.Set("access_network_asra", flattenWirelessControllerHotspot20HsProfileAccessNetworkAsra(o["access-network-asra"], d, "access_network_asra", sv)); err != nil {
		if !fortiAPIPatch(o["access-network-asra"]) {
			return fmt.Errorf("Error reading access_network_asra: %v", err)
		}
	}

	if err = d.Set("access_network_esr", flattenWirelessControllerHotspot20HsProfileAccessNetworkEsr(o["access-network-esr"], d, "access_network_esr", sv)); err != nil {
		if !fortiAPIPatch(o["access-network-esr"]) {
			return fmt.Errorf("Error reading access_network_esr: %v", err)
		}
	}

	if err = d.Set("access_network_uesa", flattenWirelessControllerHotspot20HsProfileAccessNetworkUesa(o["access-network-uesa"], d, "access_network_uesa", sv)); err != nil {
		if !fortiAPIPatch(o["access-network-uesa"]) {
			return fmt.Errorf("Error reading access_network_uesa: %v", err)
		}
	}

	if err = d.Set("venue_group", flattenWirelessControllerHotspot20HsProfileVenueGroup(o["venue-group"], d, "venue_group", sv)); err != nil {
		if !fortiAPIPatch(o["venue-group"]) {
			return fmt.Errorf("Error reading venue_group: %v", err)
		}
	}

	if err = d.Set("venue_type", flattenWirelessControllerHotspot20HsProfileVenueType(o["venue-type"], d, "venue_type", sv)); err != nil {
		if !fortiAPIPatch(o["venue-type"]) {
			return fmt.Errorf("Error reading venue_type: %v", err)
		}
	}

	if err = d.Set("hessid", flattenWirelessControllerHotspot20HsProfileHessid(o["hessid"], d, "hessid", sv)); err != nil {
		if !fortiAPIPatch(o["hessid"]) {
			return fmt.Errorf("Error reading hessid: %v", err)
		}
	}

	if err = d.Set("proxy_arp", flattenWirelessControllerHotspot20HsProfileProxyArp(o["proxy-arp"], d, "proxy_arp", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-arp"]) {
			return fmt.Errorf("Error reading proxy_arp: %v", err)
		}
	}

	if err = d.Set("l2tif", flattenWirelessControllerHotspot20HsProfileL2Tif(o["l2tif"], d, "l2tif", sv)); err != nil {
		if !fortiAPIPatch(o["l2tif"]) {
			return fmt.Errorf("Error reading l2tif: %v", err)
		}
	}

	if err = d.Set("pame_bi", flattenWirelessControllerHotspot20HsProfilePameBi(o["pame-bi"], d, "pame_bi", sv)); err != nil {
		if !fortiAPIPatch(o["pame-bi"]) {
			return fmt.Errorf("Error reading pame_bi: %v", err)
		}
	}

	if err = d.Set("anqp_domain_id", flattenWirelessControllerHotspot20HsProfileAnqpDomainId(o["anqp-domain-id"], d, "anqp_domain_id", sv)); err != nil {
		if !fortiAPIPatch(o["anqp-domain-id"]) {
			return fmt.Errorf("Error reading anqp_domain_id: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenWirelessControllerHotspot20HsProfileDomainName(o["domain-name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("osu_ssid", flattenWirelessControllerHotspot20HsProfileOsuSsid(o["osu-ssid"], d, "osu_ssid", sv)); err != nil {
		if !fortiAPIPatch(o["osu-ssid"]) {
			return fmt.Errorf("Error reading osu_ssid: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenWirelessControllerHotspot20HsProfileGasComebackDelay(o["gas-comeback-delay"], d, "gas_comeback_delay", sv)); err != nil {
		if !fortiAPIPatch(o["gas-comeback-delay"]) {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenWirelessControllerHotspot20HsProfileGasFragmentationLimit(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit", sv)); err != nil {
		if !fortiAPIPatch(o["gas-fragmentation-limit"]) {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("dgaf", flattenWirelessControllerHotspot20HsProfileDgaf(o["dgaf"], d, "dgaf", sv)); err != nil {
		if !fortiAPIPatch(o["dgaf"]) {
			return fmt.Errorf("Error reading dgaf: %v", err)
		}
	}

	if err = d.Set("deauth_request_timeout", flattenWirelessControllerHotspot20HsProfileDeauthRequestTimeout(o["deauth-request-timeout"], d, "deauth_request_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["deauth-request-timeout"]) {
			return fmt.Errorf("Error reading deauth_request_timeout: %v", err)
		}
	}

	if err = d.Set("wnm_sleep_mode", flattenWirelessControllerHotspot20HsProfileWnmSleepMode(o["wnm-sleep-mode"], d, "wnm_sleep_mode", sv)); err != nil {
		if !fortiAPIPatch(o["wnm-sleep-mode"]) {
			return fmt.Errorf("Error reading wnm_sleep_mode: %v", err)
		}
	}

	if err = d.Set("bss_transition", flattenWirelessControllerHotspot20HsProfileBssTransition(o["bss-transition"], d, "bss_transition", sv)); err != nil {
		if !fortiAPIPatch(o["bss-transition"]) {
			return fmt.Errorf("Error reading bss_transition: %v", err)
		}
	}

	if err = d.Set("venue_name", flattenWirelessControllerHotspot20HsProfileVenueName(o["venue-name"], d, "venue_name", sv)); err != nil {
		if !fortiAPIPatch(o["venue-name"]) {
			return fmt.Errorf("Error reading venue_name: %v", err)
		}
	}

	if err = d.Set("venue_url", flattenWirelessControllerHotspot20HsProfileVenueUrl(o["venue-url"], d, "venue_url", sv)); err != nil {
		if !fortiAPIPatch(o["venue-url"]) {
			return fmt.Errorf("Error reading venue_url: %v", err)
		}
	}

	if err = d.Set("roaming_consortium", flattenWirelessControllerHotspot20HsProfileRoamingConsortium(o["roaming-consortium"], d, "roaming_consortium", sv)); err != nil {
		if !fortiAPIPatch(o["roaming-consortium"]) {
			return fmt.Errorf("Error reading roaming_consortium: %v", err)
		}
	}

	if err = d.Set("nai_realm", flattenWirelessControllerHotspot20HsProfileNaiRealm(o["nai-realm"], d, "nai_realm", sv)); err != nil {
		if !fortiAPIPatch(o["nai-realm"]) {
			return fmt.Errorf("Error reading nai_realm: %v", err)
		}
	}

	if err = d.Set("oper_friendly_name", flattenWirelessControllerHotspot20HsProfileOperFriendlyName(o["oper-friendly-name"], d, "oper_friendly_name", sv)); err != nil {
		if !fortiAPIPatch(o["oper-friendly-name"]) {
			return fmt.Errorf("Error reading oper_friendly_name: %v", err)
		}
	}

	if err = d.Set("oper_icon", flattenWirelessControllerHotspot20HsProfileOperIcon(o["oper-icon"], d, "oper_icon", sv)); err != nil {
		if !fortiAPIPatch(o["oper-icon"]) {
			return fmt.Errorf("Error reading oper_icon: %v", err)
		}
	}

	if err = d.Set("advice_of_charge", flattenWirelessControllerHotspot20HsProfileAdviceOfCharge(o["advice-of-charge"], d, "advice_of_charge", sv)); err != nil {
		if !fortiAPIPatch(o["advice-of-charge"]) {
			return fmt.Errorf("Error reading advice_of_charge: %v", err)
		}
	}

	if err = d.Set("osu_provider_nai", flattenWirelessControllerHotspot20HsProfileOsuProviderNai(o["osu-provider-nai"], d, "osu_provider_nai", sv)); err != nil {
		if !fortiAPIPatch(o["osu-provider-nai"]) {
			return fmt.Errorf("Error reading osu_provider_nai: %v", err)
		}
	}

	if err = d.Set("terms_and_conditions", flattenWirelessControllerHotspot20HsProfileTermsAndConditions(o["terms-and-conditions"], d, "terms_and_conditions", sv)); err != nil {
		if !fortiAPIPatch(o["terms-and-conditions"]) {
			return fmt.Errorf("Error reading terms_and_conditions: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("osu_provider", flattenWirelessControllerHotspot20HsProfileOsuProvider(o["osu-provider"], d, "osu_provider", sv)); err != nil {
			if !fortiAPIPatch(o["osu-provider"]) {
				return fmt.Errorf("Error reading osu_provider: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("osu_provider"); ok {
			if err = d.Set("osu_provider", flattenWirelessControllerHotspot20HsProfileOsuProvider(o["osu-provider"], d, "osu_provider", sv)); err != nil {
				if !fortiAPIPatch(o["osu-provider"]) {
					return fmt.Errorf("Error reading osu_provider: %v", err)
				}
			}
		}
	}

	if err = d.Set("wan_metrics", flattenWirelessControllerHotspot20HsProfileWanMetrics(o["wan-metrics"], d, "wan_metrics", sv)); err != nil {
		if !fortiAPIPatch(o["wan-metrics"]) {
			return fmt.Errorf("Error reading wan_metrics: %v", err)
		}
	}

	if err = d.Set("network_auth", flattenWirelessControllerHotspot20HsProfileNetworkAuth(o["network-auth"], d, "network_auth", sv)); err != nil {
		if !fortiAPIPatch(o["network-auth"]) {
			return fmt.Errorf("Error reading network_auth: %v", err)
		}
	}

	if err = d.Set("n3gpp_plmn", flattenWirelessControllerHotspot20HsProfile3GppPlmn(o["3gpp-plmn"], d, "n3gpp_plmn", sv)); err != nil {
		if !fortiAPIPatch(o["3gpp-plmn"]) {
			return fmt.Errorf("Error reading n3gpp_plmn: %v", err)
		}
	}

	if err = d.Set("conn_cap", flattenWirelessControllerHotspot20HsProfileConnCap(o["conn-cap"], d, "conn_cap", sv)); err != nil {
		if !fortiAPIPatch(o["conn-cap"]) {
			return fmt.Errorf("Error reading conn_cap: %v", err)
		}
	}

	if err = d.Set("qos_map", flattenWirelessControllerHotspot20HsProfileQosMap(o["qos-map"], d, "qos_map", sv)); err != nil {
		if !fortiAPIPatch(o["qos-map"]) {
			return fmt.Errorf("Error reading qos_map: %v", err)
		}
	}

	if err = d.Set("ip_addr_type", flattenWirelessControllerHotspot20HsProfileIpAddrType(o["ip-addr-type"], d, "ip_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["ip-addr-type"]) {
			return fmt.Errorf("Error reading ip_addr_type: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20HsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20HsProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileRelease(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkInternet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkAsra(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkEsr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAccessNetworkUesa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileVenueGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileVenueType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileHessid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileProxyArp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileL2Tif(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfilePameBi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAnqpDomainId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileOsuSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileGasComebackDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileDgaf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileDeauthRequestTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWnmSleepMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileBssTransition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileVenueName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileVenueUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileRoamingConsortium(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileNaiRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileOperFriendlyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileOperIcon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileAdviceOfCharge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileOsuProviderNai(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileTermsAndConditions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileOsuProvider(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerHotspot20HsProfileOsuProviderName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20HsProfileOsuProviderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileWanMetrics(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileNetworkAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfile3GppPlmn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileConnCap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileQosMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20HsProfileIpAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20HsProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("release"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileRelease(d, v, "release", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["release"] = t
		}
	}

	if v, ok := d.GetOk("access_network_type"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkType(d, v, "access_network_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-type"] = t
		}
	}

	if v, ok := d.GetOk("access_network_internet"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkInternet(d, v, "access_network_internet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-internet"] = t
		}
	}

	if v, ok := d.GetOk("access_network_asra"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkAsra(d, v, "access_network_asra", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-asra"] = t
		}
	}

	if v, ok := d.GetOk("access_network_esr"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkEsr(d, v, "access_network_esr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-esr"] = t
		}
	}

	if v, ok := d.GetOk("access_network_uesa"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAccessNetworkUesa(d, v, "access_network_uesa", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-network-uesa"] = t
		}
	}

	if v, ok := d.GetOk("venue_group"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileVenueGroup(d, v, "venue_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-group"] = t
		}
	}

	if v, ok := d.GetOk("venue_type"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileVenueType(d, v, "venue_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-type"] = t
		}
	}

	if v, ok := d.GetOk("hessid"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileHessid(d, v, "hessid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hessid"] = t
		}
	}

	if v, ok := d.GetOk("proxy_arp"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileProxyArp(d, v, "proxy_arp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-arp"] = t
		}
	}

	if v, ok := d.GetOk("l2tif"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileL2Tif(d, v, "l2tif", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2tif"] = t
		}
	}

	if v, ok := d.GetOk("pame_bi"); ok {

		t, err := expandWirelessControllerHotspot20HsProfilePameBi(d, v, "pame_bi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pame-bi"] = t
		}
	}

	if v, ok := d.GetOkExists("anqp_domain_id"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAnqpDomainId(d, v, "anqp_domain_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anqp-domain-id"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileDomainName(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("osu_ssid"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileOsuSsid(d, v, "osu_ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-ssid"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileGasComebackDelay(d, v, "gas_comeback_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileGasFragmentationLimit(d, v, "gas_fragmentation_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("dgaf"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileDgaf(d, v, "dgaf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dgaf"] = t
		}
	}

	if v, ok := d.GetOk("deauth_request_timeout"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileDeauthRequestTimeout(d, v, "deauth_request_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-request-timeout"] = t
		}
	}

	if v, ok := d.GetOk("wnm_sleep_mode"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileWnmSleepMode(d, v, "wnm_sleep_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wnm-sleep-mode"] = t
		}
	}

	if v, ok := d.GetOk("bss_transition"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileBssTransition(d, v, "bss_transition", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("venue_name"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileVenueName(d, v, "venue_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-name"] = t
		}
	}

	if v, ok := d.GetOk("venue_url"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileVenueUrl(d, v, "venue_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["venue-url"] = t
		}
	}

	if v, ok := d.GetOk("roaming_consortium"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileRoamingConsortium(d, v, "roaming_consortium", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-consortium"] = t
		}
	}

	if v, ok := d.GetOk("nai_realm"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileNaiRealm(d, v, "nai_realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-realm"] = t
		}
	}

	if v, ok := d.GetOk("oper_friendly_name"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileOperFriendlyName(d, v, "oper_friendly_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oper-friendly-name"] = t
		}
	}

	if v, ok := d.GetOk("oper_icon"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileOperIcon(d, v, "oper_icon", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oper-icon"] = t
		}
	}

	if v, ok := d.GetOk("advice_of_charge"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileAdviceOfCharge(d, v, "advice_of_charge", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advice-of-charge"] = t
		}
	}

	if v, ok := d.GetOk("osu_provider_nai"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileOsuProviderNai(d, v, "osu_provider_nai", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-provider-nai"] = t
		}
	}

	if v, ok := d.GetOk("terms_and_conditions"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileTermsAndConditions(d, v, "terms_and_conditions", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["terms-and-conditions"] = t
		}
	}

	if v, ok := d.GetOk("osu_provider"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileOsuProvider(d, v, "osu_provider", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osu-provider"] = t
		}
	}

	if v, ok := d.GetOk("wan_metrics"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileWanMetrics(d, v, "wan_metrics", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-metrics"] = t
		}
	}

	if v, ok := d.GetOk("network_auth"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileNetworkAuth(d, v, "network_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-auth"] = t
		}
	}

	if v, ok := d.GetOk("n3gpp_plmn"); ok {

		t, err := expandWirelessControllerHotspot20HsProfile3GppPlmn(d, v, "n3gpp_plmn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["3gpp-plmn"] = t
		}
	}

	if v, ok := d.GetOk("conn_cap"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileConnCap(d, v, "conn_cap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-cap"] = t
		}
	}

	if v, ok := d.GetOk("qos_map"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileQosMap(d, v, "qos_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-map"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_type"); ok {

		t, err := expandWirelessControllerHotspot20HsProfileIpAddrType(d, v, "ip_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-type"] = t
		}
	}

	return &obj, nil
}
