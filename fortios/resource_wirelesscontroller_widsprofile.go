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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed:     true,
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
			"ap_bgscan_disable_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_bgscan_disable_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_bgscan_disable_end": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dynamic_sort_subtable": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

func flattenWirelessControllerWidsProfileApBgscanPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanIdle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileApBgscanDisableSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerWidsProfileApBgscanDisableSchedulesName(i["name"], d, pre_append, sv)
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
	return v
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
	return v
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
	return v
}

func flattenWirelessControllerWidsProfileAuthFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFrameFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFloodTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileAssocFloodThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
	return v
}

func flattenWirelessControllerWidsProfileEapolStartIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolLogoffIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolSuccIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolFailIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreSuccIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailFlood(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileEapolPreFailIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWidsProfileDeauthUnknownSrcThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerWidsProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

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

	if isImportTable() {
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

			tmp["name"], _ = expandWirelessControllerWidsProfileApBgscanDisableSchedulesName(d, i["name"], pre_append, sv)
		}

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

	if v, ok := d.GetOk("ap_bgscan_disable_schedules"); ok {

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
	}

	if v, ok := d.GetOk("ap_bgscan_disable_start"); ok {

		t, err := expandWirelessControllerWidsProfileApBgscanDisableStart(d, v, "ap_bgscan_disable_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-start"] = t
		}
	}

	if v, ok := d.GetOk("ap_bgscan_disable_end"); ok {

		t, err := expandWirelessControllerWidsProfileApBgscanDisableEnd(d, v, "ap_bgscan_disable_end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-bgscan-disable-end"] = t
		}
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

	return &obj, nil
}
