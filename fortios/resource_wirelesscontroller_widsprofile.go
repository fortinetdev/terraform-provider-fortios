// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure wireless intrusion detection system (WIDS) profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWidsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWidsProfileCreate,
		Read:   resourceWirelessControllerWidsProfileRead,
		Update: resourceWirelessControllerWidsProfileUpdate,
		Delete: resourceWirelessControllerWidsProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"sensor_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_scan_channel_list_2g_5g": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional:     true,
						},
					},
				},
			},
			"ap_scan_channel_list_6g": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional:     true,
						},
					},
				},
			},
			"ap_bgscan_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"ap_bgscan_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 600),
				Optional:     true,
				Computed:     true,
			},
			"ap_bgscan_duration": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000),
				Optional:     true,
				Computed:     true,
			},
			"ap_bgscan_idle": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000),
				Optional:     true,
				Computed:     true,
			},
			"ap_bgscan_report_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 600),
				Optional:     true,
				Computed:     true,
			},
			"ap_bgscan_disable_schedules": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"ap_bgscan_disable_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_bgscan_disable_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_bgscan_disable_end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_fgscan_report_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(15, 600),
				Optional:     true,
				Computed:     true,
			},
			"ap_scan_passive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_scan_threshold": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"ap_auto_suppress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_bridge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deauth_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"null_ssid_probe_resp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_duration_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"long_duration_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1000, 32767),
				Optional:     true,
				Computed:     true,
			},
			"invalid_mac_oui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weak_wep_iv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 120),
				Optional:     true,
				Computed:     true,
			},
			"auth_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"assoc_frame_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 120),
				Optional:     true,
				Computed:     true,
			},
			"assoc_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"reassoc_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reassoc_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"reassoc_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"probe_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"probe_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"bcn_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bcn_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"bcn_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"rts_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rts_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"rts_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"cts_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cts_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"cts_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"client_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"client_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"block_ack_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_ack_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"block_ack_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"pspoll_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pspoll_flood_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"pspoll_flood_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"netstumbler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netstumbler_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"netstumbler_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"wellenreiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wellenreiter_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"wellenreiter_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65100),
				Optional:     true,
				Computed:     true,
			},
			"spoofed_deauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"asleap_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_start_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_start_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),
				Optional:     true,
				Computed:     true,
			},
			"eapol_start_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"eapol_logoff_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_logoff_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),
				Optional:     true,
				Computed:     true,
			},
			"eapol_logoff_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"eapol_succ_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_succ_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),
				Optional:     true,
				Computed:     true,
			},
			"eapol_succ_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"eapol_fail_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_fail_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),
				Optional:     true,
				Computed:     true,
			},
			"eapol_fail_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"eapol_pre_succ_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_succ_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),
				Optional:     true,
				Computed:     true,
			},
			"eapol_pre_succ_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"eapol_pre_fail_flood": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_pre_fail_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 100),
				Optional:     true,
				Computed:     true,
			},
			"eapol_pre_fail_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"deauth_unknown_src_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"windows_bridge": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disassoc_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_spoofing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"chan_based_mitm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adhoc_valid_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adhoc_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eapol_key_overflow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_impersonation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"invalid_addr_combination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"beacon_wrong_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ht_greenfield": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overflow_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_ht_ie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"malformed_association": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ht_40mhz_intolerance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_ssid_misuse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"valid_client_misassociation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hotspotter_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pwsave_dos_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"omerta_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disconnect_station": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unencrypted_valid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fata_jack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"risky_encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fuzzed_beacon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fuzzed_probe_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fuzzed_probe_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"air_jack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wpa_ft_attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerWidsProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWirelessControllerWidsProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWidsProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWidsProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWidsProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWidsProfile")
	}

	return resourceWirelessControllerWidsProfileRead(d, m)
}

func resourceWirelessControllerWidsProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWirelessControllerWidsProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWidsProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWidsProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWidsProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWidsProfile")
	}

	return resourceWirelessControllerWidsProfileRead(d, m)
}

func resourceWirelessControllerWidsProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerWidsProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWidsProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWidsProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadWirelessControllerWidsProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWidsProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWidsProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWidsProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWidsProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileSensorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScanChannelList2G5G(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "chan", "chan")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["chan"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
			}
			tmp["chan"] = flattenWirelessControllerWidsProfileApScanChannelList2G5GChan(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWidsProfileApScanChannelList2G5GChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScanChannelList6G(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "chan", "chan")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["chan"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
			}
			tmp["chan"] = flattenWirelessControllerWidsProfileApScanChannelList6GChan(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWidsProfileApScanChannelList6GChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileApBgscanIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileApBgscanDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileApBgscanIdle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileApBgscanReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileApBgscanDisableSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenWirelessControllerWidsProfileApBgscanDisableSchedulesName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWidsProfileApBgscanDisableSchedulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanDisableDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanDisableStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanDisableEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApFgscanReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileApScanPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApScanThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApAutoSuppress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWirelessBridge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDeauthBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNullSsidProbeResp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileLongDurationAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileLongDurationThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileInvalidMacOui(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWeakWepIv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAuthFrameFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAuthFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileAuthFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileAssocFrameFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileAssocFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileReassocFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileReassocFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileReassocFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileProbeFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileProbeFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileProbeFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileBcnFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBcnFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileBcnFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileRtsFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileRtsFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileRtsFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileCtsFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileCtsFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileCtsFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileClientFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileClientFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileClientFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileBlockAckFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBlockAckFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileBlockAckFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfilePspollFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfilePspollFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfilePspollFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileNetstumbler(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileNetstumblerTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileNetstumblerThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileWellenreiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWellenreiterTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileWellenreiterThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileSpoofedDeauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAsleapAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolStartFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolStartThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolStartIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolLogoffFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolLogoffIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolSuccFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolSuccIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolFailFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolFailIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolPreSuccFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolPreSuccIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolPreFailFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileEapolPreFailIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileDeauthUnknownSrcThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerWidsProfileWindowsBridge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDisassocBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApSpoofing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileChanBasedMitm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAdhocValidSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAdhocNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolKeyOverflow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApImpersonation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileInvalidAddrCombination(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileBeaconWrongChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileHtGreenfield(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileOverflowIe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileMalformedHtIe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileMalformedAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileMalformedAssociation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileHt40MhzIntolerance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileValidSsidMisuse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileValidClientMisassociation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileHotspotterAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfilePwsaveDosAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileOmertaAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDisconnectStation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileUnencryptedValid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFataJack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileRiskyEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFuzzedBeacon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFuzzedProbeRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileFuzzedProbeResponse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAirJack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileWpaFtAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerWidsProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerWidsProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerWidsProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("sensor_mode", flattenWirelessControllerWidsProfileSensorMode(o["sensor-mode"], d, "sensor_mode", sv)); err != nil {
		if !fortiAPIPatch(o["sensor-mode"]) {
			return fmt.Errorf("Error reading sensor_mode: %v", err)
		}
	}

	if err = d.Set("ap_scan", flattenWirelessControllerWidsProfileApScan(o["ap-scan"], d, "ap_scan", sv)); err != nil {
		if !fortiAPIPatch(o["ap-scan"]) {
			return fmt.Errorf("Error reading ap_scan: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ap_scan_channel_list_2g_5g", flattenWirelessControllerWidsProfileApScanChannelList2G5G(o["ap-scan-channel-list-2G-5G"], d, "ap_scan_channel_list_2g_5g", sv)); err != nil {
			if !fortiAPIPatch(o["ap-scan-channel-list-2G-5G"]) {
				return fmt.Errorf("Error reading ap_scan_channel_list_2g_5g: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ap_scan_channel_list_2g_5g"); ok {
			if err = d.Set("ap_scan_channel_list_2g_5g", flattenWirelessControllerWidsProfileApScanChannelList2G5G(o["ap-scan-channel-list-2G-5G"], d, "ap_scan_channel_list_2g_5g", sv)); err != nil {
				if !fortiAPIPatch(o["ap-scan-channel-list-2G-5G"]) {
					return fmt.Errorf("Error reading ap_scan_channel_list_2g_5g: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ap_scan_channel_list_6g", flattenWirelessControllerWidsProfileApScanChannelList6G(o["ap-scan-channel-list-6G"], d, "ap_scan_channel_list_6g", sv)); err != nil {
			if !fortiAPIPatch(o["ap-scan-channel-list-6G"]) {
				return fmt.Errorf("Error reading ap_scan_channel_list_6g: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ap_scan_channel_list_6g"); ok {
			if err = d.Set("ap_scan_channel_list_6g", flattenWirelessControllerWidsProfileApScanChannelList6G(o["ap-scan-channel-list-6G"], d, "ap_scan_channel_list_6g", sv)); err != nil {
				if !fortiAPIPatch(o["ap-scan-channel-list-6G"]) {
					return fmt.Errorf("Error reading ap_scan_channel_list_6g: %v", err)
				}
			}
		}
	}

	if err = d.Set("ap_bgscan_period", flattenWirelessControllerWidsProfileApBgscanPeriod(o["ap-bgscan-period"], d, "ap_bgscan_period", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-period"]) {
			return fmt.Errorf("Error reading ap_bgscan_period: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_intv", flattenWirelessControllerWidsProfileApBgscanIntv(o["ap-bgscan-intv"], d, "ap_bgscan_intv", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-intv"]) {
			return fmt.Errorf("Error reading ap_bgscan_intv: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_duration", flattenWirelessControllerWidsProfileApBgscanDuration(o["ap-bgscan-duration"], d, "ap_bgscan_duration", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-duration"]) {
			return fmt.Errorf("Error reading ap_bgscan_duration: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_idle", flattenWirelessControllerWidsProfileApBgscanIdle(o["ap-bgscan-idle"], d, "ap_bgscan_idle", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-idle"]) {
			return fmt.Errorf("Error reading ap_bgscan_idle: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_report_intv", flattenWirelessControllerWidsProfileApBgscanReportIntv(o["ap-bgscan-report-intv"], d, "ap_bgscan_report_intv", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-report-intv"]) {
			return fmt.Errorf("Error reading ap_bgscan_report_intv: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ap_bgscan_disable_schedules", flattenWirelessControllerWidsProfileApBgscanDisableSchedules(o["ap-bgscan-disable-schedules"], d, "ap_bgscan_disable_schedules", sv)); err != nil {
			if !fortiAPIPatch(o["ap-bgscan-disable-schedules"]) {
				return fmt.Errorf("Error reading ap_bgscan_disable_schedules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ap_bgscan_disable_schedules"); ok {
			if err = d.Set("ap_bgscan_disable_schedules", flattenWirelessControllerWidsProfileApBgscanDisableSchedules(o["ap-bgscan-disable-schedules"], d, "ap_bgscan_disable_schedules", sv)); err != nil {
				if !fortiAPIPatch(o["ap-bgscan-disable-schedules"]) {
					return fmt.Errorf("Error reading ap_bgscan_disable_schedules: %v", err)
				}
			}
		}
	}

	if err = d.Set("ap_bgscan_disable_day", flattenWirelessControllerWidsProfileApBgscanDisableDay(o["ap-bgscan-disable-day"], d, "ap_bgscan_disable_day", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-disable-day"]) {
			return fmt.Errorf("Error reading ap_bgscan_disable_day: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_start", flattenWirelessControllerWidsProfileApBgscanDisableStart(o["ap-bgscan-disable-start"], d, "ap_bgscan_disable_start", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-disable-start"]) {
			return fmt.Errorf("Error reading ap_bgscan_disable_start: %v", err)
		}
	}

	if err = d.Set("ap_bgscan_disable_end", flattenWirelessControllerWidsProfileApBgscanDisableEnd(o["ap-bgscan-disable-end"], d, "ap_bgscan_disable_end", sv)); err != nil {
		if !fortiAPIPatch(o["ap-bgscan-disable-end"]) {
			return fmt.Errorf("Error reading ap_bgscan_disable_end: %v", err)
		}
	}

	if err = d.Set("ap_fgscan_report_intv", flattenWirelessControllerWidsProfileApFgscanReportIntv(o["ap-fgscan-report-intv"], d, "ap_fgscan_report_intv", sv)); err != nil {
		if !fortiAPIPatch(o["ap-fgscan-report-intv"]) {
			return fmt.Errorf("Error reading ap_fgscan_report_intv: %v", err)
		}
	}

	if err = d.Set("ap_scan_passive", flattenWirelessControllerWidsProfileApScanPassive(o["ap-scan-passive"], d, "ap_scan_passive", sv)); err != nil {
		if !fortiAPIPatch(o["ap-scan-passive"]) {
			return fmt.Errorf("Error reading ap_scan_passive: %v", err)
		}
	}

	if err = d.Set("ap_scan_threshold", flattenWirelessControllerWidsProfileApScanThreshold(o["ap-scan-threshold"], d, "ap_scan_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["ap-scan-threshold"]) {
			return fmt.Errorf("Error reading ap_scan_threshold: %v", err)
		}
	}

	if err = d.Set("ap_auto_suppress", flattenWirelessControllerWidsProfileApAutoSuppress(o["ap-auto-suppress"], d, "ap_auto_suppress", sv)); err != nil {
		if !fortiAPIPatch(o["ap-auto-suppress"]) {
			return fmt.Errorf("Error reading ap_auto_suppress: %v", err)
		}
	}

	if err = d.Set("wireless_bridge", flattenWirelessControllerWidsProfileWirelessBridge(o["wireless-bridge"], d, "wireless_bridge", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-bridge"]) {
			return fmt.Errorf("Error reading wireless_bridge: %v", err)
		}
	}

	if err = d.Set("deauth_broadcast", flattenWirelessControllerWidsProfileDeauthBroadcast(o["deauth-broadcast"], d, "deauth_broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["deauth-broadcast"]) {
			return fmt.Errorf("Error reading deauth_broadcast: %v", err)
		}
	}

	if err = d.Set("null_ssid_probe_resp", flattenWirelessControllerWidsProfileNullSsidProbeResp(o["null-ssid-probe-resp"], d, "null_ssid_probe_resp", sv)); err != nil {
		if !fortiAPIPatch(o["null-ssid-probe-resp"]) {
			return fmt.Errorf("Error reading null_ssid_probe_resp: %v", err)
		}
	}

	if err = d.Set("long_duration_attack", flattenWirelessControllerWidsProfileLongDurationAttack(o["long-duration-attack"], d, "long_duration_attack", sv)); err != nil {
		if !fortiAPIPatch(o["long-duration-attack"]) {
			return fmt.Errorf("Error reading long_duration_attack: %v", err)
		}
	}

	if err = d.Set("long_duration_thresh", flattenWirelessControllerWidsProfileLongDurationThresh(o["long-duration-thresh"], d, "long_duration_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["long-duration-thresh"]) {
			return fmt.Errorf("Error reading long_duration_thresh: %v", err)
		}
	}

	if err = d.Set("invalid_mac_oui", flattenWirelessControllerWidsProfileInvalidMacOui(o["invalid-mac-oui"], d, "invalid_mac_oui", sv)); err != nil {
		if !fortiAPIPatch(o["invalid-mac-oui"]) {
			return fmt.Errorf("Error reading invalid_mac_oui: %v", err)
		}
	}

	if err = d.Set("weak_wep_iv", flattenWirelessControllerWidsProfileWeakWepIv(o["weak-wep-iv"], d, "weak_wep_iv", sv)); err != nil {
		if !fortiAPIPatch(o["weak-wep-iv"]) {
			return fmt.Errorf("Error reading weak_wep_iv: %v", err)
		}
	}

	if err = d.Set("auth_frame_flood", flattenWirelessControllerWidsProfileAuthFrameFlood(o["auth-frame-flood"], d, "auth_frame_flood", sv)); err != nil {
		if !fortiAPIPatch(o["auth-frame-flood"]) {
			return fmt.Errorf("Error reading auth_frame_flood: %v", err)
		}
	}

	if err = d.Set("auth_flood_time", flattenWirelessControllerWidsProfileAuthFloodTime(o["auth-flood-time"], d, "auth_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["auth-flood-time"]) {
			return fmt.Errorf("Error reading auth_flood_time: %v", err)
		}
	}

	if err = d.Set("auth_flood_thresh", flattenWirelessControllerWidsProfileAuthFloodThresh(o["auth-flood-thresh"], d, "auth_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["auth-flood-thresh"]) {
			return fmt.Errorf("Error reading auth_flood_thresh: %v", err)
		}
	}

	if err = d.Set("assoc_frame_flood", flattenWirelessControllerWidsProfileAssocFrameFlood(o["assoc-frame-flood"], d, "assoc_frame_flood", sv)); err != nil {
		if !fortiAPIPatch(o["assoc-frame-flood"]) {
			return fmt.Errorf("Error reading assoc_frame_flood: %v", err)
		}
	}

	if err = d.Set("assoc_flood_time", flattenWirelessControllerWidsProfileAssocFloodTime(o["assoc-flood-time"], d, "assoc_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["assoc-flood-time"]) {
			return fmt.Errorf("Error reading assoc_flood_time: %v", err)
		}
	}

	if err = d.Set("assoc_flood_thresh", flattenWirelessControllerWidsProfileAssocFloodThresh(o["assoc-flood-thresh"], d, "assoc_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["assoc-flood-thresh"]) {
			return fmt.Errorf("Error reading assoc_flood_thresh: %v", err)
		}
	}

	if err = d.Set("reassoc_flood", flattenWirelessControllerWidsProfileReassocFlood(o["reassoc-flood"], d, "reassoc_flood", sv)); err != nil {
		if !fortiAPIPatch(o["reassoc-flood"]) {
			return fmt.Errorf("Error reading reassoc_flood: %v", err)
		}
	}

	if err = d.Set("reassoc_flood_time", flattenWirelessControllerWidsProfileReassocFloodTime(o["reassoc-flood-time"], d, "reassoc_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["reassoc-flood-time"]) {
			return fmt.Errorf("Error reading reassoc_flood_time: %v", err)
		}
	}

	if err = d.Set("reassoc_flood_thresh", flattenWirelessControllerWidsProfileReassocFloodThresh(o["reassoc-flood-thresh"], d, "reassoc_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["reassoc-flood-thresh"]) {
			return fmt.Errorf("Error reading reassoc_flood_thresh: %v", err)
		}
	}

	if err = d.Set("probe_flood", flattenWirelessControllerWidsProfileProbeFlood(o["probe-flood"], d, "probe_flood", sv)); err != nil {
		if !fortiAPIPatch(o["probe-flood"]) {
			return fmt.Errorf("Error reading probe_flood: %v", err)
		}
	}

	if err = d.Set("probe_flood_time", flattenWirelessControllerWidsProfileProbeFloodTime(o["probe-flood-time"], d, "probe_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["probe-flood-time"]) {
			return fmt.Errorf("Error reading probe_flood_time: %v", err)
		}
	}

	if err = d.Set("probe_flood_thresh", flattenWirelessControllerWidsProfileProbeFloodThresh(o["probe-flood-thresh"], d, "probe_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["probe-flood-thresh"]) {
			return fmt.Errorf("Error reading probe_flood_thresh: %v", err)
		}
	}

	if err = d.Set("bcn_flood", flattenWirelessControllerWidsProfileBcnFlood(o["bcn-flood"], d, "bcn_flood", sv)); err != nil {
		if !fortiAPIPatch(o["bcn-flood"]) {
			return fmt.Errorf("Error reading bcn_flood: %v", err)
		}
	}

	if err = d.Set("bcn_flood_time", flattenWirelessControllerWidsProfileBcnFloodTime(o["bcn-flood-time"], d, "bcn_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["bcn-flood-time"]) {
			return fmt.Errorf("Error reading bcn_flood_time: %v", err)
		}
	}

	if err = d.Set("bcn_flood_thresh", flattenWirelessControllerWidsProfileBcnFloodThresh(o["bcn-flood-thresh"], d, "bcn_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["bcn-flood-thresh"]) {
			return fmt.Errorf("Error reading bcn_flood_thresh: %v", err)
		}
	}

	if err = d.Set("rts_flood", flattenWirelessControllerWidsProfileRtsFlood(o["rts-flood"], d, "rts_flood", sv)); err != nil {
		if !fortiAPIPatch(o["rts-flood"]) {
			return fmt.Errorf("Error reading rts_flood: %v", err)
		}
	}

	if err = d.Set("rts_flood_time", flattenWirelessControllerWidsProfileRtsFloodTime(o["rts-flood-time"], d, "rts_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["rts-flood-time"]) {
			return fmt.Errorf("Error reading rts_flood_time: %v", err)
		}
	}

	if err = d.Set("rts_flood_thresh", flattenWirelessControllerWidsProfileRtsFloodThresh(o["rts-flood-thresh"], d, "rts_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["rts-flood-thresh"]) {
			return fmt.Errorf("Error reading rts_flood_thresh: %v", err)
		}
	}

	if err = d.Set("cts_flood", flattenWirelessControllerWidsProfileCtsFlood(o["cts-flood"], d, "cts_flood", sv)); err != nil {
		if !fortiAPIPatch(o["cts-flood"]) {
			return fmt.Errorf("Error reading cts_flood: %v", err)
		}
	}

	if err = d.Set("cts_flood_time", flattenWirelessControllerWidsProfileCtsFloodTime(o["cts-flood-time"], d, "cts_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["cts-flood-time"]) {
			return fmt.Errorf("Error reading cts_flood_time: %v", err)
		}
	}

	if err = d.Set("cts_flood_thresh", flattenWirelessControllerWidsProfileCtsFloodThresh(o["cts-flood-thresh"], d, "cts_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["cts-flood-thresh"]) {
			return fmt.Errorf("Error reading cts_flood_thresh: %v", err)
		}
	}

	if err = d.Set("client_flood", flattenWirelessControllerWidsProfileClientFlood(o["client-flood"], d, "client_flood", sv)); err != nil {
		if !fortiAPIPatch(o["client-flood"]) {
			return fmt.Errorf("Error reading client_flood: %v", err)
		}
	}

	if err = d.Set("client_flood_time", flattenWirelessControllerWidsProfileClientFloodTime(o["client-flood-time"], d, "client_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["client-flood-time"]) {
			return fmt.Errorf("Error reading client_flood_time: %v", err)
		}
	}

	if err = d.Set("client_flood_thresh", flattenWirelessControllerWidsProfileClientFloodThresh(o["client-flood-thresh"], d, "client_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["client-flood-thresh"]) {
			return fmt.Errorf("Error reading client_flood_thresh: %v", err)
		}
	}

	if err = d.Set("block_ack_flood", flattenWirelessControllerWidsProfileBlockAckFlood(o["block_ack-flood"], d, "block_ack_flood", sv)); err != nil {
		if !fortiAPIPatch(o["block_ack-flood"]) {
			return fmt.Errorf("Error reading block_ack_flood: %v", err)
		}
	}

	if err = d.Set("block_ack_flood_time", flattenWirelessControllerWidsProfileBlockAckFloodTime(o["block_ack-flood-time"], d, "block_ack_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["block_ack-flood-time"]) {
			return fmt.Errorf("Error reading block_ack_flood_time: %v", err)
		}
	}

	if err = d.Set("block_ack_flood_thresh", flattenWirelessControllerWidsProfileBlockAckFloodThresh(o["block_ack-flood-thresh"], d, "block_ack_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["block_ack-flood-thresh"]) {
			return fmt.Errorf("Error reading block_ack_flood_thresh: %v", err)
		}
	}

	if err = d.Set("pspoll_flood", flattenWirelessControllerWidsProfilePspollFlood(o["pspoll-flood"], d, "pspoll_flood", sv)); err != nil {
		if !fortiAPIPatch(o["pspoll-flood"]) {
			return fmt.Errorf("Error reading pspoll_flood: %v", err)
		}
	}

	if err = d.Set("pspoll_flood_time", flattenWirelessControllerWidsProfilePspollFloodTime(o["pspoll-flood-time"], d, "pspoll_flood_time", sv)); err != nil {
		if !fortiAPIPatch(o["pspoll-flood-time"]) {
			return fmt.Errorf("Error reading pspoll_flood_time: %v", err)
		}
	}

	if err = d.Set("pspoll_flood_thresh", flattenWirelessControllerWidsProfilePspollFloodThresh(o["pspoll-flood-thresh"], d, "pspoll_flood_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["pspoll-flood-thresh"]) {
			return fmt.Errorf("Error reading pspoll_flood_thresh: %v", err)
		}
	}

	if err = d.Set("netstumbler", flattenWirelessControllerWidsProfileNetstumbler(o["netstumbler"], d, "netstumbler", sv)); err != nil {
		if !fortiAPIPatch(o["netstumbler"]) {
			return fmt.Errorf("Error reading netstumbler: %v", err)
		}
	}

	if err = d.Set("netstumbler_time", flattenWirelessControllerWidsProfileNetstumblerTime(o["netstumbler-time"], d, "netstumbler_time", sv)); err != nil {
		if !fortiAPIPatch(o["netstumbler-time"]) {
			return fmt.Errorf("Error reading netstumbler_time: %v", err)
		}
	}

	if err = d.Set("netstumbler_thresh", flattenWirelessControllerWidsProfileNetstumblerThresh(o["netstumbler-thresh"], d, "netstumbler_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["netstumbler-thresh"]) {
			return fmt.Errorf("Error reading netstumbler_thresh: %v", err)
		}
	}

	if err = d.Set("wellenreiter", flattenWirelessControllerWidsProfileWellenreiter(o["wellenreiter"], d, "wellenreiter", sv)); err != nil {
		if !fortiAPIPatch(o["wellenreiter"]) {
			return fmt.Errorf("Error reading wellenreiter: %v", err)
		}
	}

	if err = d.Set("wellenreiter_time", flattenWirelessControllerWidsProfileWellenreiterTime(o["wellenreiter-time"], d, "wellenreiter_time", sv)); err != nil {
		if !fortiAPIPatch(o["wellenreiter-time"]) {
			return fmt.Errorf("Error reading wellenreiter_time: %v", err)
		}
	}

	if err = d.Set("wellenreiter_thresh", flattenWirelessControllerWidsProfileWellenreiterThresh(o["wellenreiter-thresh"], d, "wellenreiter_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["wellenreiter-thresh"]) {
			return fmt.Errorf("Error reading wellenreiter_thresh: %v", err)
		}
	}

	if err = d.Set("spoofed_deauth", flattenWirelessControllerWidsProfileSpoofedDeauth(o["spoofed-deauth"], d, "spoofed_deauth", sv)); err != nil {
		if !fortiAPIPatch(o["spoofed-deauth"]) {
			return fmt.Errorf("Error reading spoofed_deauth: %v", err)
		}
	}

	if err = d.Set("asleap_attack", flattenWirelessControllerWidsProfileAsleapAttack(o["asleap-attack"], d, "asleap_attack", sv)); err != nil {
		if !fortiAPIPatch(o["asleap-attack"]) {
			return fmt.Errorf("Error reading asleap_attack: %v", err)
		}
	}

	if err = d.Set("eapol_start_flood", flattenWirelessControllerWidsProfileEapolStartFlood(o["eapol-start-flood"], d, "eapol_start_flood", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-start-flood"]) {
			return fmt.Errorf("Error reading eapol_start_flood: %v", err)
		}
	}

	if err = d.Set("eapol_start_thresh", flattenWirelessControllerWidsProfileEapolStartThresh(o["eapol-start-thresh"], d, "eapol_start_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-start-thresh"]) {
			return fmt.Errorf("Error reading eapol_start_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_start_intv", flattenWirelessControllerWidsProfileEapolStartIntv(o["eapol-start-intv"], d, "eapol_start_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-start-intv"]) {
			return fmt.Errorf("Error reading eapol_start_intv: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_flood", flattenWirelessControllerWidsProfileEapolLogoffFlood(o["eapol-logoff-flood"], d, "eapol_logoff_flood", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-logoff-flood"]) {
			return fmt.Errorf("Error reading eapol_logoff_flood: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_thresh", flattenWirelessControllerWidsProfileEapolLogoffThresh(o["eapol-logoff-thresh"], d, "eapol_logoff_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-logoff-thresh"]) {
			return fmt.Errorf("Error reading eapol_logoff_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_logoff_intv", flattenWirelessControllerWidsProfileEapolLogoffIntv(o["eapol-logoff-intv"], d, "eapol_logoff_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-logoff-intv"]) {
			return fmt.Errorf("Error reading eapol_logoff_intv: %v", err)
		}
	}

	if err = d.Set("eapol_succ_flood", flattenWirelessControllerWidsProfileEapolSuccFlood(o["eapol-succ-flood"], d, "eapol_succ_flood", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-succ-flood"]) {
			return fmt.Errorf("Error reading eapol_succ_flood: %v", err)
		}
	}

	if err = d.Set("eapol_succ_thresh", flattenWirelessControllerWidsProfileEapolSuccThresh(o["eapol-succ-thresh"], d, "eapol_succ_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-succ-thresh"]) {
			return fmt.Errorf("Error reading eapol_succ_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_succ_intv", flattenWirelessControllerWidsProfileEapolSuccIntv(o["eapol-succ-intv"], d, "eapol_succ_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-succ-intv"]) {
			return fmt.Errorf("Error reading eapol_succ_intv: %v", err)
		}
	}

	if err = d.Set("eapol_fail_flood", flattenWirelessControllerWidsProfileEapolFailFlood(o["eapol-fail-flood"], d, "eapol_fail_flood", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-fail-flood"]) {
			return fmt.Errorf("Error reading eapol_fail_flood: %v", err)
		}
	}

	if err = d.Set("eapol_fail_thresh", flattenWirelessControllerWidsProfileEapolFailThresh(o["eapol-fail-thresh"], d, "eapol_fail_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-fail-thresh"]) {
			return fmt.Errorf("Error reading eapol_fail_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_fail_intv", flattenWirelessControllerWidsProfileEapolFailIntv(o["eapol-fail-intv"], d, "eapol_fail_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-fail-intv"]) {
			return fmt.Errorf("Error reading eapol_fail_intv: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_flood", flattenWirelessControllerWidsProfileEapolPreSuccFlood(o["eapol-pre-succ-flood"], d, "eapol_pre_succ_flood", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-pre-succ-flood"]) {
			return fmt.Errorf("Error reading eapol_pre_succ_flood: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_thresh", flattenWirelessControllerWidsProfileEapolPreSuccThresh(o["eapol-pre-succ-thresh"], d, "eapol_pre_succ_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-pre-succ-thresh"]) {
			return fmt.Errorf("Error reading eapol_pre_succ_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_pre_succ_intv", flattenWirelessControllerWidsProfileEapolPreSuccIntv(o["eapol-pre-succ-intv"], d, "eapol_pre_succ_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-pre-succ-intv"]) {
			return fmt.Errorf("Error reading eapol_pre_succ_intv: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_flood", flattenWirelessControllerWidsProfileEapolPreFailFlood(o["eapol-pre-fail-flood"], d, "eapol_pre_fail_flood", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-pre-fail-flood"]) {
			return fmt.Errorf("Error reading eapol_pre_fail_flood: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_thresh", flattenWirelessControllerWidsProfileEapolPreFailThresh(o["eapol-pre-fail-thresh"], d, "eapol_pre_fail_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-pre-fail-thresh"]) {
			return fmt.Errorf("Error reading eapol_pre_fail_thresh: %v", err)
		}
	}

	if err = d.Set("eapol_pre_fail_intv", flattenWirelessControllerWidsProfileEapolPreFailIntv(o["eapol-pre-fail-intv"], d, "eapol_pre_fail_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-pre-fail-intv"]) {
			return fmt.Errorf("Error reading eapol_pre_fail_intv: %v", err)
		}
	}

	if err = d.Set("deauth_unknown_src_thresh", flattenWirelessControllerWidsProfileDeauthUnknownSrcThresh(o["deauth-unknown-src-thresh"], d, "deauth_unknown_src_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["deauth-unknown-src-thresh"]) {
			return fmt.Errorf("Error reading deauth_unknown_src_thresh: %v", err)
		}
	}

	if err = d.Set("windows_bridge", flattenWirelessControllerWidsProfileWindowsBridge(o["windows-bridge"], d, "windows_bridge", sv)); err != nil {
		if !fortiAPIPatch(o["windows-bridge"]) {
			return fmt.Errorf("Error reading windows_bridge: %v", err)
		}
	}

	if err = d.Set("disassoc_broadcast", flattenWirelessControllerWidsProfileDisassocBroadcast(o["disassoc-broadcast"], d, "disassoc_broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["disassoc-broadcast"]) {
			return fmt.Errorf("Error reading disassoc_broadcast: %v", err)
		}
	}

	if err = d.Set("ap_spoofing", flattenWirelessControllerWidsProfileApSpoofing(o["ap-spoofing"], d, "ap_spoofing", sv)); err != nil {
		if !fortiAPIPatch(o["ap-spoofing"]) {
			return fmt.Errorf("Error reading ap_spoofing: %v", err)
		}
	}

	if err = d.Set("chan_based_mitm", flattenWirelessControllerWidsProfileChanBasedMitm(o["chan-based-mitm"], d, "chan_based_mitm", sv)); err != nil {
		if !fortiAPIPatch(o["chan-based-mitm"]) {
			return fmt.Errorf("Error reading chan_based_mitm: %v", err)
		}
	}

	if err = d.Set("adhoc_valid_ssid", flattenWirelessControllerWidsProfileAdhocValidSsid(o["adhoc-valid-ssid"], d, "adhoc_valid_ssid", sv)); err != nil {
		if !fortiAPIPatch(o["adhoc-valid-ssid"]) {
			return fmt.Errorf("Error reading adhoc_valid_ssid: %v", err)
		}
	}

	if err = d.Set("adhoc_network", flattenWirelessControllerWidsProfileAdhocNetwork(o["adhoc-network"], d, "adhoc_network", sv)); err != nil {
		if !fortiAPIPatch(o["adhoc-network"]) {
			return fmt.Errorf("Error reading adhoc_network: %v", err)
		}
	}

	if err = d.Set("eapol_key_overflow", flattenWirelessControllerWidsProfileEapolKeyOverflow(o["eapol-key-overflow"], d, "eapol_key_overflow", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-key-overflow"]) {
			return fmt.Errorf("Error reading eapol_key_overflow: %v", err)
		}
	}

	if err = d.Set("ap_impersonation", flattenWirelessControllerWidsProfileApImpersonation(o["ap-impersonation"], d, "ap_impersonation", sv)); err != nil {
		if !fortiAPIPatch(o["ap-impersonation"]) {
			return fmt.Errorf("Error reading ap_impersonation: %v", err)
		}
	}

	if err = d.Set("invalid_addr_combination", flattenWirelessControllerWidsProfileInvalidAddrCombination(o["invalid-addr-combination"], d, "invalid_addr_combination", sv)); err != nil {
		if !fortiAPIPatch(o["invalid-addr-combination"]) {
			return fmt.Errorf("Error reading invalid_addr_combination: %v", err)
		}
	}

	if err = d.Set("beacon_wrong_channel", flattenWirelessControllerWidsProfileBeaconWrongChannel(o["beacon-wrong-channel"], d, "beacon_wrong_channel", sv)); err != nil {
		if !fortiAPIPatch(o["beacon-wrong-channel"]) {
			return fmt.Errorf("Error reading beacon_wrong_channel: %v", err)
		}
	}

	if err = d.Set("ht_greenfield", flattenWirelessControllerWidsProfileHtGreenfield(o["ht-greenfield"], d, "ht_greenfield", sv)); err != nil {
		if !fortiAPIPatch(o["ht-greenfield"]) {
			return fmt.Errorf("Error reading ht_greenfield: %v", err)
		}
	}

	if err = d.Set("overflow_ie", flattenWirelessControllerWidsProfileOverflowIe(o["overflow-ie"], d, "overflow_ie", sv)); err != nil {
		if !fortiAPIPatch(o["overflow-ie"]) {
			return fmt.Errorf("Error reading overflow_ie: %v", err)
		}
	}

	if err = d.Set("malformed_ht_ie", flattenWirelessControllerWidsProfileMalformedHtIe(o["malformed-ht-ie"], d, "malformed_ht_ie", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-ht-ie"]) {
			return fmt.Errorf("Error reading malformed_ht_ie: %v", err)
		}
	}

	if err = d.Set("malformed_auth", flattenWirelessControllerWidsProfileMalformedAuth(o["malformed-auth"], d, "malformed_auth", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-auth"]) {
			return fmt.Errorf("Error reading malformed_auth: %v", err)
		}
	}

	if err = d.Set("malformed_association", flattenWirelessControllerWidsProfileMalformedAssociation(o["malformed-association"], d, "malformed_association", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-association"]) {
			return fmt.Errorf("Error reading malformed_association: %v", err)
		}
	}

	if err = d.Set("ht_40mhz_intolerance", flattenWirelessControllerWidsProfileHt40MhzIntolerance(o["ht-40mhz-intolerance"], d, "ht_40mhz_intolerance", sv)); err != nil {
		if !fortiAPIPatch(o["ht-40mhz-intolerance"]) {
			return fmt.Errorf("Error reading ht_40mhz_intolerance: %v", err)
		}
	}

	if err = d.Set("valid_ssid_misuse", flattenWirelessControllerWidsProfileValidSsidMisuse(o["valid-ssid-misuse"], d, "valid_ssid_misuse", sv)); err != nil {
		if !fortiAPIPatch(o["valid-ssid-misuse"]) {
			return fmt.Errorf("Error reading valid_ssid_misuse: %v", err)
		}
	}

	if err = d.Set("valid_client_misassociation", flattenWirelessControllerWidsProfileValidClientMisassociation(o["valid-client-misassociation"], d, "valid_client_misassociation", sv)); err != nil {
		if !fortiAPIPatch(o["valid-client-misassociation"]) {
			return fmt.Errorf("Error reading valid_client_misassociation: %v", err)
		}
	}

	if err = d.Set("hotspotter_attack", flattenWirelessControllerWidsProfileHotspotterAttack(o["hotspotter-attack"], d, "hotspotter_attack", sv)); err != nil {
		if !fortiAPIPatch(o["hotspotter-attack"]) {
			return fmt.Errorf("Error reading hotspotter_attack: %v", err)
		}
	}

	if err = d.Set("pwsave_dos_attack", flattenWirelessControllerWidsProfilePwsaveDosAttack(o["pwsave-dos-attack"], d, "pwsave_dos_attack", sv)); err != nil {
		if !fortiAPIPatch(o["pwsave-dos-attack"]) {
			return fmt.Errorf("Error reading pwsave_dos_attack: %v", err)
		}
	}

	if err = d.Set("omerta_attack", flattenWirelessControllerWidsProfileOmertaAttack(o["omerta-attack"], d, "omerta_attack", sv)); err != nil {
		if !fortiAPIPatch(o["omerta-attack"]) {
			return fmt.Errorf("Error reading omerta_attack: %v", err)
		}
	}

	if err = d.Set("disconnect_station", flattenWirelessControllerWidsProfileDisconnectStation(o["disconnect-station"], d, "disconnect_station", sv)); err != nil {
		if !fortiAPIPatch(o["disconnect-station"]) {
			return fmt.Errorf("Error reading disconnect_station: %v", err)
		}
	}

	if err = d.Set("unencrypted_valid", flattenWirelessControllerWidsProfileUnencryptedValid(o["unencrypted-valid"], d, "unencrypted_valid", sv)); err != nil {
		if !fortiAPIPatch(o["unencrypted-valid"]) {
			return fmt.Errorf("Error reading unencrypted_valid: %v", err)
		}
	}

	if err = d.Set("fata_jack", flattenWirelessControllerWidsProfileFataJack(o["fata-jack"], d, "fata_jack", sv)); err != nil {
		if !fortiAPIPatch(o["fata-jack"]) {
			return fmt.Errorf("Error reading fata_jack: %v", err)
		}
	}

	if err = d.Set("risky_encryption", flattenWirelessControllerWidsProfileRiskyEncryption(o["risky-encryption"], d, "risky_encryption", sv)); err != nil {
		if !fortiAPIPatch(o["risky-encryption"]) {
			return fmt.Errorf("Error reading risky_encryption: %v", err)
		}
	}

	if err = d.Set("fuzzed_beacon", flattenWirelessControllerWidsProfileFuzzedBeacon(o["fuzzed-beacon"], d, "fuzzed_beacon", sv)); err != nil {
		if !fortiAPIPatch(o["fuzzed-beacon"]) {
			return fmt.Errorf("Error reading fuzzed_beacon: %v", err)
		}
	}

	if err = d.Set("fuzzed_probe_request", flattenWirelessControllerWidsProfileFuzzedProbeRequest(o["fuzzed-probe-request"], d, "fuzzed_probe_request", sv)); err != nil {
		if !fortiAPIPatch(o["fuzzed-probe-request"]) {
			return fmt.Errorf("Error reading fuzzed_probe_request: %v", err)
		}
	}

	if err = d.Set("fuzzed_probe_response", flattenWirelessControllerWidsProfileFuzzedProbeResponse(o["fuzzed-probe-response"], d, "fuzzed_probe_response", sv)); err != nil {
		if !fortiAPIPatch(o["fuzzed-probe-response"]) {
			return fmt.Errorf("Error reading fuzzed_probe_response: %v", err)
		}
	}

	if err = d.Set("air_jack", flattenWirelessControllerWidsProfileAirJack(o["air-jack"], d, "air_jack", sv)); err != nil {
		if !fortiAPIPatch(o["air-jack"]) {
			return fmt.Errorf("Error reading air_jack: %v", err)
		}
	}

	if err = d.Set("wpa_ft_attack", flattenWirelessControllerWidsProfileWpaFtAttack(o["wpa-ft-attack"], d, "wpa_ft_attack", sv)); err != nil {
		if !fortiAPIPatch(o["wpa-ft-attack"]) {
			return fmt.Errorf("Error reading wpa_ft_attack: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWidsProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerWidsProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileSensorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScanChannelList2G5G(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["chan"], _ = expandWirelessControllerWidsProfileApScanChannelList2G5GChan(d, i["chan"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWidsProfileApScanChannelList2G5GChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScanChannelList6G(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["chan"], _ = expandWirelessControllerWidsProfileApScanChannelList6GChan(d, i["chan"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWidsProfileApScanChannelList6GChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanIdle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanDisableSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandWirelessControllerWidsProfileApBgscanDisableSchedulesName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWidsProfileApBgscanDisableSchedulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanDisableDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanDisableStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApBgscanDisableEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApFgscanReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScanPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApScanThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApAutoSuppress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWirelessBridge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDeauthBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNullSsidProbeResp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileLongDurationAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileLongDurationThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileInvalidMacOui(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWeakWepIv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAuthFrameFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAuthFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAuthFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAssocFrameFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAssocFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAssocFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileReassocFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileReassocFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileReassocFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileProbeFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileProbeFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileProbeFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBcnFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBcnFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBcnFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRtsFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRtsFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRtsFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileCtsFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileCtsFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileCtsFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileClientFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileClientFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileClientFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBlockAckFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBlockAckFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBlockAckFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePspollFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePspollFloodTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePspollFloodThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNetstumbler(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNetstumblerTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileNetstumblerThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWellenreiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWellenreiterTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWellenreiterThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileSpoofedDeauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAsleapAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolStartFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolStartThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolStartIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolLogoffFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolLogoffThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolLogoffIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolSuccFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolSuccThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolSuccIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolFailFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolFailThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolFailIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreSuccFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreSuccThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreSuccIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreFailFlood(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreFailThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolPreFailIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDeauthUnknownSrcThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWindowsBridge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDisassocBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApSpoofing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileChanBasedMitm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAdhocValidSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAdhocNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileEapolKeyOverflow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileApImpersonation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileInvalidAddrCombination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileBeaconWrongChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileHtGreenfield(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileOverflowIe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileMalformedHtIe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileMalformedAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileMalformedAssociation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileHt40MhzIntolerance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileValidSsidMisuse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileValidClientMisassociation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileHotspotterAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfilePwsaveDosAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileOmertaAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileDisconnectStation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileUnencryptedValid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFataJack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileRiskyEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFuzzedBeacon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFuzzedProbeRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileFuzzedProbeResponse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileAirJack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWidsProfileWpaFtAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWidsProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWidsProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerWidsProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("sensor_mode"); ok {
		t, err := expandWirelessControllerWidsProfileSensorMode(d, v, "sensor_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensor-mode"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan"); ok {
		t, err := expandWirelessControllerWidsProfileApScan(d, v, "ap_scan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_channel_list_2g_5g"); ok || d.HasChange("ap_scan_channel_list_2g_5g") {
		t, err := expandWirelessControllerWidsProfileApScanChannelList2G5G(d, v, "ap_scan_channel_list_2g_5g", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-channel-list-2G-5G"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_channel_list_6g"); ok || d.HasChange("ap_scan_channel_list_6g") {
		t, err := expandWirelessControllerWidsProfileApScanChannelList6G(d, v, "ap_scan_channel_list_6g", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-channel-list-6G"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_period"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanPeriod(d, v, "ap_bgscan_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-period"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_intv"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanIntv(d, v, "ap_bgscan_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_duration"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanDuration(d, v, "ap_bgscan_duration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-duration"] = t
		}
	}

	if v, ok := d.GetOkExists("ap_bgscan_idle"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanIdle(d, v, "ap_bgscan_idle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-idle"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_report_intv"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanReportIntv(d, v, "ap_bgscan_report_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_schedules"); ok || d.HasChange("ap_bgscan_disable_schedules") {
		t, err := expandWirelessControllerWidsProfileApBgscanDisableSchedules(d, v, "ap_bgscan_disable_schedules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-schedules"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_day"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanDisableDay(d, v, "ap_bgscan_disable_day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-day"] = t
		}
	} else if d.HasChange("ap_bgscan_disable_day") {
		obj["ap-bgscan-disable-day"] = nil
	}

	if v, ok := d.GetOk("ap_bgscan_disable_start"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanDisableStart(d, v, "ap_bgscan_disable_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-start"] = t
		}
	} else if d.HasChange("ap_bgscan_disable_start") {
		obj["ap-bgscan-disable-start"] = nil
	}

	if v, ok := d.GetOk("ap_bgscan_disable_end"); ok {
		t, err := expandWirelessControllerWidsProfileApBgscanDisableEnd(d, v, "ap_bgscan_disable_end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-end"] = t
		}
	} else if d.HasChange("ap_bgscan_disable_end") {
		obj["ap-bgscan-disable-end"] = nil
	}

	if v, ok := d.GetOk("ap_fgscan_report_intv"); ok {
		t, err := expandWirelessControllerWidsProfileApFgscanReportIntv(d, v, "ap_fgscan_report_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-fgscan-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_passive"); ok {
		t, err := expandWirelessControllerWidsProfileApScanPassive(d, v, "ap_scan_passive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-passive"] = t
		}
	}

	if v, ok := d.GetOk("ap_scan_threshold"); ok {
		t, err := expandWirelessControllerWidsProfileApScanThreshold(d, v, "ap_scan_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-scan-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ap_auto_suppress"); ok {
		t, err := expandWirelessControllerWidsProfileApAutoSuppress(d, v, "ap_auto_suppress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-auto-suppress"] = t
		}
	}

	if v, ok := d.GetOk("wireless_bridge"); ok {
		t, err := expandWirelessControllerWidsProfileWirelessBridge(d, v, "wireless_bridge", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wireless-bridge"] = t
		}
	}

	if v, ok := d.GetOk("deauth_broadcast"); ok {
		t, err := expandWirelessControllerWidsProfileDeauthBroadcast(d, v, "deauth_broadcast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("null_ssid_probe_resp"); ok {
		t, err := expandWirelessControllerWidsProfileNullSsidProbeResp(d, v, "null_ssid_probe_resp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-ssid-probe-resp"] = t
		}
	}

	if v, ok := d.GetOk("long_duration_attack"); ok {
		t, err := expandWirelessControllerWidsProfileLongDurationAttack(d, v, "long_duration_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-duration-attack"] = t
		}
	}

	if v, ok := d.GetOk("long_duration_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileLongDurationThresh(d, v, "long_duration_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["long-duration-thresh"] = t
		}
	}

	if v, ok := d.GetOk("invalid_mac_oui"); ok {
		t, err := expandWirelessControllerWidsProfileInvalidMacOui(d, v, "invalid_mac_oui", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-mac-oui"] = t
		}
	}

	if v, ok := d.GetOk("weak_wep_iv"); ok {
		t, err := expandWirelessControllerWidsProfileWeakWepIv(d, v, "weak_wep_iv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weak-wep-iv"] = t
		}
	}

	if v, ok := d.GetOk("auth_frame_flood"); ok {
		t, err := expandWirelessControllerWidsProfileAuthFrameFlood(d, v, "auth_frame_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-frame-flood"] = t
		}
	}

	if v, ok := d.GetOk("auth_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileAuthFloodTime(d, v, "auth_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("auth_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileAuthFloodThresh(d, v, "auth_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("assoc_frame_flood"); ok {
		t, err := expandWirelessControllerWidsProfileAssocFrameFlood(d, v, "assoc_frame_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-frame-flood"] = t
		}
	}

	if v, ok := d.GetOk("assoc_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileAssocFloodTime(d, v, "assoc_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("assoc_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileAssocFloodThresh(d, v, "assoc_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assoc-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("reassoc_flood"); ok {
		t, err := expandWirelessControllerWidsProfileReassocFlood(d, v, "reassoc_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reassoc-flood"] = t
		}
	}

	if v, ok := d.GetOk("reassoc_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileReassocFloodTime(d, v, "reassoc_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reassoc-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("reassoc_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileReassocFloodThresh(d, v, "reassoc_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reassoc-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("probe_flood"); ok {
		t, err := expandWirelessControllerWidsProfileProbeFlood(d, v, "probe_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-flood"] = t
		}
	}

	if v, ok := d.GetOk("probe_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileProbeFloodTime(d, v, "probe_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("probe_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileProbeFloodThresh(d, v, "probe_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("bcn_flood"); ok {
		t, err := expandWirelessControllerWidsProfileBcnFlood(d, v, "bcn_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bcn-flood"] = t
		}
	}

	if v, ok := d.GetOk("bcn_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileBcnFloodTime(d, v, "bcn_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bcn-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("bcn_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileBcnFloodThresh(d, v, "bcn_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bcn-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("rts_flood"); ok {
		t, err := expandWirelessControllerWidsProfileRtsFlood(d, v, "rts_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-flood"] = t
		}
	}

	if v, ok := d.GetOk("rts_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileRtsFloodTime(d, v, "rts_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("rts_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileRtsFloodThresh(d, v, "rts_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("cts_flood"); ok {
		t, err := expandWirelessControllerWidsProfileCtsFlood(d, v, "cts_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cts-flood"] = t
		}
	}

	if v, ok := d.GetOk("cts_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileCtsFloodTime(d, v, "cts_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cts-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("cts_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileCtsFloodThresh(d, v, "cts_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cts-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("client_flood"); ok {
		t, err := expandWirelessControllerWidsProfileClientFlood(d, v, "client_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-flood"] = t
		}
	}

	if v, ok := d.GetOk("client_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileClientFloodTime(d, v, "client_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("client_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileClientFloodThresh(d, v, "client_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("block_ack_flood"); ok {
		t, err := expandWirelessControllerWidsProfileBlockAckFlood(d, v, "block_ack_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_ack-flood"] = t
		}
	}

	if v, ok := d.GetOk("block_ack_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfileBlockAckFloodTime(d, v, "block_ack_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_ack-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("block_ack_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileBlockAckFloodThresh(d, v, "block_ack_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block_ack-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("pspoll_flood"); ok {
		t, err := expandWirelessControllerWidsProfilePspollFlood(d, v, "pspoll_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pspoll-flood"] = t
		}
	}

	if v, ok := d.GetOk("pspoll_flood_time"); ok {
		t, err := expandWirelessControllerWidsProfilePspollFloodTime(d, v, "pspoll_flood_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pspoll-flood-time"] = t
		}
	}

	if v, ok := d.GetOk("pspoll_flood_thresh"); ok {
		t, err := expandWirelessControllerWidsProfilePspollFloodThresh(d, v, "pspoll_flood_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pspoll-flood-thresh"] = t
		}
	}

	if v, ok := d.GetOk("netstumbler"); ok {
		t, err := expandWirelessControllerWidsProfileNetstumbler(d, v, "netstumbler", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netstumbler"] = t
		}
	}

	if v, ok := d.GetOk("netstumbler_time"); ok {
		t, err := expandWirelessControllerWidsProfileNetstumblerTime(d, v, "netstumbler_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netstumbler-time"] = t
		}
	}

	if v, ok := d.GetOk("netstumbler_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileNetstumblerThresh(d, v, "netstumbler_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netstumbler-thresh"] = t
		}
	}

	if v, ok := d.GetOk("wellenreiter"); ok {
		t, err := expandWirelessControllerWidsProfileWellenreiter(d, v, "wellenreiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wellenreiter"] = t
		}
	}

	if v, ok := d.GetOk("wellenreiter_time"); ok {
		t, err := expandWirelessControllerWidsProfileWellenreiterTime(d, v, "wellenreiter_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wellenreiter-time"] = t
		}
	}

	if v, ok := d.GetOk("wellenreiter_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileWellenreiterThresh(d, v, "wellenreiter_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wellenreiter-thresh"] = t
		}
	}

	if v, ok := d.GetOk("spoofed_deauth"); ok {
		t, err := expandWirelessControllerWidsProfileSpoofedDeauth(d, v, "spoofed_deauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spoofed-deauth"] = t
		}
	}

	if v, ok := d.GetOk("asleap_attack"); ok {
		t, err := expandWirelessControllerWidsProfileAsleapAttack(d, v, "asleap_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["asleap-attack"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_flood"); ok {
		t, err := expandWirelessControllerWidsProfileEapolStartFlood(d, v, "eapol_start_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileEapolStartThresh(d, v, "eapol_start_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_start_intv"); ok {
		t, err := expandWirelessControllerWidsProfileEapolStartIntv(d, v, "eapol_start_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-start-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_flood"); ok {
		t, err := expandWirelessControllerWidsProfileEapolLogoffFlood(d, v, "eapol_logoff_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileEapolLogoffThresh(d, v, "eapol_logoff_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_logoff_intv"); ok {
		t, err := expandWirelessControllerWidsProfileEapolLogoffIntv(d, v, "eapol_logoff_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-logoff-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_flood"); ok {
		t, err := expandWirelessControllerWidsProfileEapolSuccFlood(d, v, "eapol_succ_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileEapolSuccThresh(d, v, "eapol_succ_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_succ_intv"); ok {
		t, err := expandWirelessControllerWidsProfileEapolSuccIntv(d, v, "eapol_succ_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-succ-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_flood"); ok {
		t, err := expandWirelessControllerWidsProfileEapolFailFlood(d, v, "eapol_fail_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileEapolFailThresh(d, v, "eapol_fail_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_fail_intv"); ok {
		t, err := expandWirelessControllerWidsProfileEapolFailIntv(d, v, "eapol_fail_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-fail-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_flood"); ok {
		t, err := expandWirelessControllerWidsProfileEapolPreSuccFlood(d, v, "eapol_pre_succ_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileEapolPreSuccThresh(d, v, "eapol_pre_succ_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_succ_intv"); ok {
		t, err := expandWirelessControllerWidsProfileEapolPreSuccIntv(d, v, "eapol_pre_succ_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-succ-intv"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_flood"); ok {
		t, err := expandWirelessControllerWidsProfileEapolPreFailFlood(d, v, "eapol_pre_fail_flood", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-flood"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileEapolPreFailThresh(d, v, "eapol_pre_fail_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-thresh"] = t
		}
	}

	if v, ok := d.GetOk("eapol_pre_fail_intv"); ok {
		t, err := expandWirelessControllerWidsProfileEapolPreFailIntv(d, v, "eapol_pre_fail_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-pre-fail-intv"] = t
		}
	}

	if v, ok := d.GetOkExists("deauth_unknown_src_thresh"); ok {
		t, err := expandWirelessControllerWidsProfileDeauthUnknownSrcThresh(d, v, "deauth_unknown_src_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deauth-unknown-src-thresh"] = t
		}
	}

	if v, ok := d.GetOk("windows_bridge"); ok {
		t, err := expandWirelessControllerWidsProfileWindowsBridge(d, v, "windows_bridge", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-bridge"] = t
		}
	}

	if v, ok := d.GetOk("disassoc_broadcast"); ok {
		t, err := expandWirelessControllerWidsProfileDisassocBroadcast(d, v, "disassoc_broadcast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disassoc-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("ap_spoofing"); ok {
		t, err := expandWirelessControllerWidsProfileApSpoofing(d, v, "ap_spoofing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-spoofing"] = t
		}
	}

	if v, ok := d.GetOk("chan_based_mitm"); ok {
		t, err := expandWirelessControllerWidsProfileChanBasedMitm(d, v, "chan_based_mitm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chan-based-mitm"] = t
		}
	}

	if v, ok := d.GetOk("adhoc_valid_ssid"); ok {
		t, err := expandWirelessControllerWidsProfileAdhocValidSsid(d, v, "adhoc_valid_ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adhoc-valid-ssid"] = t
		}
	}

	if v, ok := d.GetOk("adhoc_network"); ok {
		t, err := expandWirelessControllerWidsProfileAdhocNetwork(d, v, "adhoc_network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adhoc-network"] = t
		}
	}

	if v, ok := d.GetOk("eapol_key_overflow"); ok {
		t, err := expandWirelessControllerWidsProfileEapolKeyOverflow(d, v, "eapol_key_overflow", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-overflow"] = t
		}
	}

	if v, ok := d.GetOk("ap_impersonation"); ok {
		t, err := expandWirelessControllerWidsProfileApImpersonation(d, v, "ap_impersonation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-impersonation"] = t
		}
	}

	if v, ok := d.GetOk("invalid_addr_combination"); ok {
		t, err := expandWirelessControllerWidsProfileInvalidAddrCombination(d, v, "invalid_addr_combination", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["invalid-addr-combination"] = t
		}
	}

	if v, ok := d.GetOk("beacon_wrong_channel"); ok {
		t, err := expandWirelessControllerWidsProfileBeaconWrongChannel(d, v, "beacon_wrong_channel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-wrong-channel"] = t
		}
	}

	if v, ok := d.GetOk("ht_greenfield"); ok {
		t, err := expandWirelessControllerWidsProfileHtGreenfield(d, v, "ht_greenfield", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ht-greenfield"] = t
		}
	}

	if v, ok := d.GetOk("overflow_ie"); ok {
		t, err := expandWirelessControllerWidsProfileOverflowIe(d, v, "overflow_ie", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["overflow-ie"] = t
		}
	}

	if v, ok := d.GetOk("malformed_ht_ie"); ok {
		t, err := expandWirelessControllerWidsProfileMalformedHtIe(d, v, "malformed_ht_ie", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-ht-ie"] = t
		}
	}

	if v, ok := d.GetOk("malformed_auth"); ok {
		t, err := expandWirelessControllerWidsProfileMalformedAuth(d, v, "malformed_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-auth"] = t
		}
	}

	if v, ok := d.GetOk("malformed_association"); ok {
		t, err := expandWirelessControllerWidsProfileMalformedAssociation(d, v, "malformed_association", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-association"] = t
		}
	}

	if v, ok := d.GetOk("ht_40mhz_intolerance"); ok {
		t, err := expandWirelessControllerWidsProfileHt40MhzIntolerance(d, v, "ht_40mhz_intolerance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ht-40mhz-intolerance"] = t
		}
	}

	if v, ok := d.GetOk("valid_ssid_misuse"); ok {
		t, err := expandWirelessControllerWidsProfileValidSsidMisuse(d, v, "valid_ssid_misuse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["valid-ssid-misuse"] = t
		}
	}

	if v, ok := d.GetOk("valid_client_misassociation"); ok {
		t, err := expandWirelessControllerWidsProfileValidClientMisassociation(d, v, "valid_client_misassociation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["valid-client-misassociation"] = t
		}
	}

	if v, ok := d.GetOk("hotspotter_attack"); ok {
		t, err := expandWirelessControllerWidsProfileHotspotterAttack(d, v, "hotspotter_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspotter-attack"] = t
		}
	}

	if v, ok := d.GetOk("pwsave_dos_attack"); ok {
		t, err := expandWirelessControllerWidsProfilePwsaveDosAttack(d, v, "pwsave_dos_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pwsave-dos-attack"] = t
		}
	}

	if v, ok := d.GetOk("omerta_attack"); ok {
		t, err := expandWirelessControllerWidsProfileOmertaAttack(d, v, "omerta_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["omerta-attack"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_station"); ok {
		t, err := expandWirelessControllerWidsProfileDisconnectStation(d, v, "disconnect_station", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-station"] = t
		}
	}

	if v, ok := d.GetOk("unencrypted_valid"); ok {
		t, err := expandWirelessControllerWidsProfileUnencryptedValid(d, v, "unencrypted_valid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unencrypted-valid"] = t
		}
	}

	if v, ok := d.GetOk("fata_jack"); ok {
		t, err := expandWirelessControllerWidsProfileFataJack(d, v, "fata_jack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fata-jack"] = t
		}
	}

	if v, ok := d.GetOk("risky_encryption"); ok {
		t, err := expandWirelessControllerWidsProfileRiskyEncryption(d, v, "risky_encryption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risky-encryption"] = t
		}
	}

	if v, ok := d.GetOk("fuzzed_beacon"); ok {
		t, err := expandWirelessControllerWidsProfileFuzzedBeacon(d, v, "fuzzed_beacon", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fuzzed-beacon"] = t
		}
	}

	if v, ok := d.GetOk("fuzzed_probe_request"); ok {
		t, err := expandWirelessControllerWidsProfileFuzzedProbeRequest(d, v, "fuzzed_probe_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fuzzed-probe-request"] = t
		}
	}

	if v, ok := d.GetOk("fuzzed_probe_response"); ok {
		t, err := expandWirelessControllerWidsProfileFuzzedProbeResponse(d, v, "fuzzed_probe_response", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fuzzed-probe-response"] = t
		}
	}

	if v, ok := d.GetOk("air_jack"); ok {
		t, err := expandWirelessControllerWidsProfileAirJack(d, v, "air_jack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["air-jack"] = t
		}
	}

	if v, ok := d.GetOk("wpa_ft_attack"); ok {
		t, err := expandWirelessControllerWidsProfileWpaFtAttack(d, v, "wpa_ft_attack", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wpa-ft-attack"] = t
		}
	}

	return &obj, nil
}
