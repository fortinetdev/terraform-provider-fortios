// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure CAPWAP timers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"client_idle_rehome_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 3600),
				Optional:     true,
				Computed:     true,
			},
			"auth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 30),
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
			"rogue_ap_cleanup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"drma_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
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

func resourceWirelessControllerTimersUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerTimers(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerTimers resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerTimers(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerTimers(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerTimers resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerTimers(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WirelessControllerTimers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerTimersRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerTimers(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerTimers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerTimers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerTimers resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerTimersEchoInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersDiscoveryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersClientIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersClientIdleRehomeTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersAuthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersRogueApLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersFakeApLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersRogueApCleanup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersDarrpOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersDarrpDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersDarrpTime(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := i["time"]; ok {

			tmp["time"] = flattenWirelessControllerTimersDarrpTimeTime(i["time"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "time", d)
	return result
}

func flattenWirelessControllerTimersDarrpTimeTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaStatsInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersVapStatsInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersRadioStatsInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaCapabilityInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersStaLocateTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersIpsecIntfCleanup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersBleScanReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerTimersDrmaInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerTimers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("echo_interval", flattenWirelessControllerTimersEchoInterval(o["echo-interval"], d, "echo_interval", sv)); err != nil {
		if !fortiAPIPatch(o["echo-interval"]) {
			return fmt.Errorf("Error reading echo_interval: %v", err)
		}
	}

	if err = d.Set("discovery_interval", flattenWirelessControllerTimersDiscoveryInterval(o["discovery-interval"], d, "discovery_interval", sv)); err != nil {
		if !fortiAPIPatch(o["discovery-interval"]) {
			return fmt.Errorf("Error reading discovery_interval: %v", err)
		}
	}

	if err = d.Set("client_idle_timeout", flattenWirelessControllerTimersClientIdleTimeout(o["client-idle-timeout"], d, "client_idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["client-idle-timeout"]) {
			return fmt.Errorf("Error reading client_idle_timeout: %v", err)
		}
	}

	if err = d.Set("client_idle_rehome_timeout", flattenWirelessControllerTimersClientIdleRehomeTimeout(o["client-idle-rehome-timeout"], d, "client_idle_rehome_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["client-idle-rehome-timeout"]) {
			return fmt.Errorf("Error reading client_idle_rehome_timeout: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenWirelessControllerTimersAuthTimeout(o["auth-timeout"], d, "auth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auth-timeout"]) {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if err = d.Set("rogue_ap_log", flattenWirelessControllerTimersRogueApLog(o["rogue-ap-log"], d, "rogue_ap_log", sv)); err != nil {
		if !fortiAPIPatch(o["rogue-ap-log"]) {
			return fmt.Errorf("Error reading rogue_ap_log: %v", err)
		}
	}

	if err = d.Set("fake_ap_log", flattenWirelessControllerTimersFakeApLog(o["fake-ap-log"], d, "fake_ap_log", sv)); err != nil {
		if !fortiAPIPatch(o["fake-ap-log"]) {
			return fmt.Errorf("Error reading fake_ap_log: %v", err)
		}
	}

	if err = d.Set("rogue_ap_cleanup", flattenWirelessControllerTimersRogueApCleanup(o["rogue-ap-cleanup"], d, "rogue_ap_cleanup", sv)); err != nil {
		if !fortiAPIPatch(o["rogue-ap-cleanup"]) {
			return fmt.Errorf("Error reading rogue_ap_cleanup: %v", err)
		}
	}

	if err = d.Set("darrp_optimize", flattenWirelessControllerTimersDarrpOptimize(o["darrp-optimize"], d, "darrp_optimize", sv)); err != nil {
		if !fortiAPIPatch(o["darrp-optimize"]) {
			return fmt.Errorf("Error reading darrp_optimize: %v", err)
		}
	}

	if err = d.Set("darrp_day", flattenWirelessControllerTimersDarrpDay(o["darrp-day"], d, "darrp_day", sv)); err != nil {
		if !fortiAPIPatch(o["darrp-day"]) {
			return fmt.Errorf("Error reading darrp_day: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("darrp_time", flattenWirelessControllerTimersDarrpTime(o["darrp-time"], d, "darrp_time", sv)); err != nil {
			if !fortiAPIPatch(o["darrp-time"]) {
				return fmt.Errorf("Error reading darrp_time: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("darrp_time"); ok {
			if err = d.Set("darrp_time", flattenWirelessControllerTimersDarrpTime(o["darrp-time"], d, "darrp_time", sv)); err != nil {
				if !fortiAPIPatch(o["darrp-time"]) {
					return fmt.Errorf("Error reading darrp_time: %v", err)
				}
			}
		}
	}

	if err = d.Set("sta_stats_interval", flattenWirelessControllerTimersStaStatsInterval(o["sta-stats-interval"], d, "sta_stats_interval", sv)); err != nil {
		if !fortiAPIPatch(o["sta-stats-interval"]) {
			return fmt.Errorf("Error reading sta_stats_interval: %v", err)
		}
	}

	if err = d.Set("vap_stats_interval", flattenWirelessControllerTimersVapStatsInterval(o["vap-stats-interval"], d, "vap_stats_interval", sv)); err != nil {
		if !fortiAPIPatch(o["vap-stats-interval"]) {
			return fmt.Errorf("Error reading vap_stats_interval: %v", err)
		}
	}

	if err = d.Set("radio_stats_interval", flattenWirelessControllerTimersRadioStatsInterval(o["radio-stats-interval"], d, "radio_stats_interval", sv)); err != nil {
		if !fortiAPIPatch(o["radio-stats-interval"]) {
			return fmt.Errorf("Error reading radio_stats_interval: %v", err)
		}
	}

	if err = d.Set("sta_capability_interval", flattenWirelessControllerTimersStaCapabilityInterval(o["sta-capability-interval"], d, "sta_capability_interval", sv)); err != nil {
		if !fortiAPIPatch(o["sta-capability-interval"]) {
			return fmt.Errorf("Error reading sta_capability_interval: %v", err)
		}
	}

	if err = d.Set("sta_locate_timer", flattenWirelessControllerTimersStaLocateTimer(o["sta-locate-timer"], d, "sta_locate_timer", sv)); err != nil {
		if !fortiAPIPatch(o["sta-locate-timer"]) {
			return fmt.Errorf("Error reading sta_locate_timer: %v", err)
		}
	}

	if err = d.Set("ipsec_intf_cleanup", flattenWirelessControllerTimersIpsecIntfCleanup(o["ipsec-intf-cleanup"], d, "ipsec_intf_cleanup", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-intf-cleanup"]) {
			return fmt.Errorf("Error reading ipsec_intf_cleanup: %v", err)
		}
	}

	if err = d.Set("ble_scan_report_intv", flattenWirelessControllerTimersBleScanReportIntv(o["ble-scan-report-intv"], d, "ble_scan_report_intv", sv)); err != nil {
		if !fortiAPIPatch(o["ble-scan-report-intv"]) {
			return fmt.Errorf("Error reading ble_scan_report_intv: %v", err)
		}
	}

	if err = d.Set("drma_interval", flattenWirelessControllerTimersDrmaInterval(o["drma-interval"], d, "drma_interval", sv)); err != nil {
		if !fortiAPIPatch(o["drma-interval"]) {
			return fmt.Errorf("Error reading drma_interval: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerTimersFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerTimersEchoInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDiscoveryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersClientIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersClientIdleRehomeTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersAuthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRogueApLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersFakeApLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRogueApCleanup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDarrpOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDarrpDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDarrpTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "time"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["time"], _ = expandWirelessControllerTimersDarrpTimeTime(d, i["time"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerTimersDarrpTimeTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaStatsInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersVapStatsInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersRadioStatsInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaCapabilityInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersStaLocateTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersIpsecIntfCleanup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersBleScanReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerTimersDrmaInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerTimers(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("echo_interval"); ok {
		if setArgNil {
			obj["echo-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersEchoInterval(d, v, "echo_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["echo-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("discovery_interval"); ok {
		if setArgNil {
			obj["discovery-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersDiscoveryInterval(d, v, "discovery_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["discovery-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("client_idle_timeout"); ok {
		if setArgNil {
			obj["client-idle-timeout"] = nil
		} else {

			t, err := expandWirelessControllerTimersClientIdleTimeout(d, v, "client_idle_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["client-idle-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("client_idle_rehome_timeout"); ok {
		if setArgNil {
			obj["client-idle-rehome-timeout"] = nil
		} else {

			t, err := expandWirelessControllerTimersClientIdleRehomeTimeout(d, v, "client_idle_rehome_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["client-idle-rehome-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok {
		if setArgNil {
			obj["auth-timeout"] = nil
		} else {

			t, err := expandWirelessControllerTimersAuthTimeout(d, v, "auth_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("rogue_ap_log"); ok {
		if setArgNil {
			obj["rogue-ap-log"] = nil
		} else {

			t, err := expandWirelessControllerTimersRogueApLog(d, v, "rogue_ap_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rogue-ap-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("fake_ap_log"); ok {
		if setArgNil {
			obj["fake-ap-log"] = nil
		} else {

			t, err := expandWirelessControllerTimersFakeApLog(d, v, "fake_ap_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fake-ap-log"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("rogue_ap_cleanup"); ok {
		if setArgNil {
			obj["rogue-ap-cleanup"] = nil
		} else {

			t, err := expandWirelessControllerTimersRogueApCleanup(d, v, "rogue_ap_cleanup", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rogue-ap-cleanup"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("darrp_optimize"); ok {
		if setArgNil {
			obj["darrp-optimize"] = nil
		} else {

			t, err := expandWirelessControllerTimersDarrpOptimize(d, v, "darrp_optimize", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["darrp-optimize"] = t
			}
		}
	}

	if v, ok := d.GetOk("darrp_day"); ok {
		if setArgNil {
			obj["darrp-day"] = nil
		} else {

			t, err := expandWirelessControllerTimersDarrpDay(d, v, "darrp_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["darrp-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("darrp_time"); ok || d.HasChange("darrp_time") {
		if setArgNil {
			obj["darrp-time"] = make([]struct{}, 0)
		} else {

			t, err := expandWirelessControllerTimersDarrpTime(d, v, "darrp_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["darrp-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("sta_stats_interval"); ok {
		if setArgNil {
			obj["sta-stats-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersStaStatsInterval(d, v, "sta_stats_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sta-stats-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("vap_stats_interval"); ok {
		if setArgNil {
			obj["vap-stats-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersVapStatsInterval(d, v, "vap_stats_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vap-stats-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("radio_stats_interval"); ok {
		if setArgNil {
			obj["radio-stats-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersRadioStatsInterval(d, v, "radio_stats_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radio-stats-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("sta_capability_interval"); ok {
		if setArgNil {
			obj["sta-capability-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersStaCapabilityInterval(d, v, "sta_capability_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sta-capability-interval"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("sta_locate_timer"); ok {
		if setArgNil {
			obj["sta-locate-timer"] = nil
		} else {

			t, err := expandWirelessControllerTimersStaLocateTimer(d, v, "sta_locate_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sta-locate-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_intf_cleanup"); ok {
		if setArgNil {
			obj["ipsec-intf-cleanup"] = nil
		} else {

			t, err := expandWirelessControllerTimersIpsecIntfCleanup(d, v, "ipsec_intf_cleanup", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-intf-cleanup"] = t
			}
		}
	}

	if v, ok := d.GetOk("ble_scan_report_intv"); ok {
		if setArgNil {
			obj["ble-scan-report-intv"] = nil
		} else {

			t, err := expandWirelessControllerTimersBleScanReportIntv(d, v, "ble_scan_report_intv", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ble-scan-report-intv"] = t
			}
		}
	}

	if v, ok := d.GetOk("drma_interval"); ok {
		if setArgNil {
			obj["drma-interval"] = nil
		} else {

			t, err := expandWirelessControllerTimersDrmaInterval(d, v, "drma_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["drma-interval"] = t
			}
		}
	}

	return &obj, nil
}
