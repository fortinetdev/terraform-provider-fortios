// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure CAPWAP timers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerTimers() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerTimersUpdate,
		Read:   resourceWirelessControllerTimersRead,
		Update: resourceWirelessControllerTimersUpdate,
		Delete: resourceWirelessControllerTimersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"echo_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"discovery_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 180),
				Optional:     true,
				Computed:     true,
			},
			"client_idle_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 3600),
				Optional:     true,
				Computed:     true,
			},
			"rogue_ap_log": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1440),
				Optional:     true,
				Computed:     true,
			},
			"fake_ap_log": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"darrp_optimize": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"darrp_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"darrp_time": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 5),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"sta_stats_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"vap_stats_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"radio_stats_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"sta_capability_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"sta_locate_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_intf_cleanup": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 3600),
				Optional:     true,
				Computed:     true,
			},
			"ble_scan_report_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerTimersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerTimers(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerTimers resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerTimers(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerTimers resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerTimers")
	}

	return resourceWirelessControllerTimersRead(d, m)
}

func resourceWirelessControllerTimersDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerTimers(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerTimers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerTimersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerTimers(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerTimers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerTimers(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerTimers resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerTimersEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersDiscoveryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersClientIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersRogueApLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersFakeApLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersDarrpOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersDarrpDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersDarrpTime(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := i["time"]; ok {
			tmp["time"] = flattenWirelessControllerTimersDarrpTimeTime(i["time"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerTimersDarrpTimeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaStatsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersVapStatsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersRadioStatsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaCapabilityInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaLocateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersIpsecIntfCleanup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerTimersBleScanReportIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerTimers(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("echo_interval", flattenWirelessControllerTimersEchoInterval(o["echo-interval"], d, "echo_interval")); err != nil {
		if !fortiAPIPatch(o["echo-interval"]) {
			return fmt.Errorf("Error reading echo_interval: %v", err)
		}
	}

	if err = d.Set("discovery_interval", flattenWirelessControllerTimersDiscoveryInterval(o["discovery-interval"], d, "discovery_interval")); err != nil {
		if !fortiAPIPatch(o["discovery-interval"]) {
			return fmt.Errorf("Error reading discovery_interval: %v", err)
		}
	}

	if err = d.Set("client_idle_timeout", flattenWirelessControllerTimersClientIdleTimeout(o["client-idle-timeout"], d, "client_idle_timeout")); err != nil {
		if !fortiAPIPatch(o["client-idle-timeout"]) {
			return fmt.Errorf("Error reading client_idle_timeout: %v", err)
		}
	}

	if err = d.Set("rogue_ap_log", flattenWirelessControllerTimersRogueApLog(o["rogue-ap-log"], d, "rogue_ap_log")); err != nil {
		if !fortiAPIPatch(o["rogue-ap-log"]) {
			return fmt.Errorf("Error reading rogue_ap_log: %v", err)
		}
	}

	if err = d.Set("fake_ap_log", flattenWirelessControllerTimersFakeApLog(o["fake-ap-log"], d, "fake_ap_log")); err != nil {
		if !fortiAPIPatch(o["fake-ap-log"]) {
			return fmt.Errorf("Error reading fake_ap_log: %v", err)
		}
	}

	if err = d.Set("darrp_optimize", flattenWirelessControllerTimersDarrpOptimize(o["darrp-optimize"], d, "darrp_optimize")); err != nil {
		if !fortiAPIPatch(o["darrp-optimize"]) {
			return fmt.Errorf("Error reading darrp_optimize: %v", err)
		}
	}

	if err = d.Set("darrp_day", flattenWirelessControllerTimersDarrpDay(o["darrp-day"], d, "darrp_day")); err != nil {
		if !fortiAPIPatch(o["darrp-day"]) {
			return fmt.Errorf("Error reading darrp_day: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("darrp_time", flattenWirelessControllerTimersDarrpTime(o["darrp-time"], d, "darrp_time")); err != nil {
			if !fortiAPIPatch(o["darrp-time"]) {
				return fmt.Errorf("Error reading darrp_time: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("darrp_time"); ok {
			if err = d.Set("darrp_time", flattenWirelessControllerTimersDarrpTime(o["darrp-time"], d, "darrp_time")); err != nil {
				if !fortiAPIPatch(o["darrp-time"]) {
					return fmt.Errorf("Error reading darrp_time: %v", err)
				}
			}
		}
	}

	if err = d.Set("sta_stats_interval", flattenWirelessControllerTimersStaStatsInterval(o["sta-stats-interval"], d, "sta_stats_interval")); err != nil {
		if !fortiAPIPatch(o["sta-stats-interval"]) {
			return fmt.Errorf("Error reading sta_stats_interval: %v", err)
		}
	}

	if err = d.Set("vap_stats_interval", flattenWirelessControllerTimersVapStatsInterval(o["vap-stats-interval"], d, "vap_stats_interval")); err != nil {
		if !fortiAPIPatch(o["vap-stats-interval"]) {
			return fmt.Errorf("Error reading vap_stats_interval: %v", err)
		}
	}

	if err = d.Set("radio_stats_interval", flattenWirelessControllerTimersRadioStatsInterval(o["radio-stats-interval"], d, "radio_stats_interval")); err != nil {
		if !fortiAPIPatch(o["radio-stats-interval"]) {
			return fmt.Errorf("Error reading radio_stats_interval: %v", err)
		}
	}

	if err = d.Set("sta_capability_interval", flattenWirelessControllerTimersStaCapabilityInterval(o["sta-capability-interval"], d, "sta_capability_interval")); err != nil {
		if !fortiAPIPatch(o["sta-capability-interval"]) {
			return fmt.Errorf("Error reading sta_capability_interval: %v", err)
		}
	}

	if err = d.Set("sta_locate_timer", flattenWirelessControllerTimersStaLocateTimer(o["sta-locate-timer"], d, "sta_locate_timer")); err != nil {
		if !fortiAPIPatch(o["sta-locate-timer"]) {
			return fmt.Errorf("Error reading sta_locate_timer: %v", err)
		}
	}

	if err = d.Set("ipsec_intf_cleanup", flattenWirelessControllerTimersIpsecIntfCleanup(o["ipsec-intf-cleanup"], d, "ipsec_intf_cleanup")); err != nil {
		if !fortiAPIPatch(o["ipsec-intf-cleanup"]) {
			return fmt.Errorf("Error reading ipsec_intf_cleanup: %v", err)
		}
	}

	if err = d.Set("ble_scan_report_intv", flattenWirelessControllerTimersBleScanReportIntv(o["ble-scan-report-intv"], d, "ble_scan_report_intv")); err != nil {
		if !fortiAPIPatch(o["ble-scan-report-intv"]) {
			return fmt.Errorf("Error reading ble_scan_report_intv: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerTimersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerTimersEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDiscoveryInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersClientIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRogueApLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersFakeApLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDarrpOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDarrpDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDarrpTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["time"], _ = expandWirelessControllerTimersDarrpTimeTime(d, i["time"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerTimersDarrpTimeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaStatsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersVapStatsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRadioStatsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaCapabilityInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaLocateTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersIpsecIntfCleanup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersBleScanReportIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerTimers(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("echo_interval"); ok {
		t, err := expandWirelessControllerTimersEchoInterval(d, v, "echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("discovery_interval"); ok {
		t, err := expandWirelessControllerTimersDiscoveryInterval(d, v, "discovery_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discovery-interval"] = t
		}
	}

	if v, ok := d.GetOk("client_idle_timeout"); ok {
		t, err := expandWirelessControllerTimersClientIdleTimeout(d, v, "client_idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("rogue_ap_log"); ok {
		t, err := expandWirelessControllerTimersRogueApLog(d, v, "rogue_ap_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rogue-ap-log"] = t
		}
	}

	if v, ok := d.GetOk("fake_ap_log"); ok {
		t, err := expandWirelessControllerTimersFakeApLog(d, v, "fake_ap_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fake-ap-log"] = t
		}
	}

	if v, ok := d.GetOkExists("darrp_optimize"); ok {
		t, err := expandWirelessControllerTimersDarrpOptimize(d, v, "darrp_optimize")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-optimize"] = t
		}
	}

	if v, ok := d.GetOk("darrp_day"); ok {
		t, err := expandWirelessControllerTimersDarrpDay(d, v, "darrp_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-day"] = t
		}
	}

	if v, ok := d.GetOk("darrp_time"); ok {
		t, err := expandWirelessControllerTimersDarrpTime(d, v, "darrp_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["darrp-time"] = t
		}
	}

	if v, ok := d.GetOk("sta_stats_interval"); ok {
		t, err := expandWirelessControllerTimersStaStatsInterval(d, v, "sta_stats_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-stats-interval"] = t
		}
	}

	if v, ok := d.GetOk("vap_stats_interval"); ok {
		t, err := expandWirelessControllerTimersVapStatsInterval(d, v, "vap_stats_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vap-stats-interval"] = t
		}
	}

	if v, ok := d.GetOk("radio_stats_interval"); ok {
		t, err := expandWirelessControllerTimersRadioStatsInterval(d, v, "radio_stats_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-stats-interval"] = t
		}
	}

	if v, ok := d.GetOk("sta_capability_interval"); ok {
		t, err := expandWirelessControllerTimersStaCapabilityInterval(d, v, "sta_capability_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-capability-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("sta_locate_timer"); ok {
		t, err := expandWirelessControllerTimersStaLocateTimer(d, v, "sta_locate_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sta-locate-timer"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_intf_cleanup"); ok {
		t, err := expandWirelessControllerTimersIpsecIntfCleanup(d, v, "ipsec_intf_cleanup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-intf-cleanup"] = t
		}
	}

	if v, ok := d.GetOk("ble_scan_report_intv"); ok {
		t, err := expandWirelessControllerTimersBleScanReportIntv(d, v, "ble_scan_report_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-scan-report-intv"] = t
		}
	}

	return &obj, nil
}
