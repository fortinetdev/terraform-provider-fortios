// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WiFi Automatic Radio Resource Provisioning (ARRP) profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerArrpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerArrpProfileCreate,
		Read:   resourceWirelessControllerArrpProfileRead,
		Update: resourceWirelessControllerArrpProfileUpdate,
		Delete: resourceWirelessControllerArrpProfileDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"selection_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"monitor_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"weight_managed_ap": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"weight_rogue_ap": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"weight_noise_floor": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"weight_channel_load": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"weight_spectral_rssi": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"weight_weather_channel": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"weight_dfs_channel": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"threshold_ap": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 500),
				Optional:     true,
				Computed:     true,
			},
			"threshold_noise_floor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"threshold_channel_load": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"threshold_spectral_rssi": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"threshold_tx_retries": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000),
				Optional:     true,
				Computed:     true,
			},
			"threshold_rx_errors": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"include_weather_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_dfs_channel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerArrpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerArrpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerArrpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerArrpProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerArrpProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerArrpProfile")
	}

	return resourceWirelessControllerArrpProfileRead(d, m)
}

func resourceWirelessControllerArrpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerArrpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerArrpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerArrpProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerArrpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerArrpProfile")
	}

	return resourceWirelessControllerArrpProfileRead(d, m)
}

func resourceWirelessControllerArrpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerArrpProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerArrpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerArrpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerArrpProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerArrpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerArrpProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerArrpProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerArrpProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileSelectionPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileMonitorPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightManagedAp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightRogueAp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightNoiseFloor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightChannelLoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightSpectralRssi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightWeatherChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileWeightDfsChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdAp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdNoiseFloor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdChannelLoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdSpectralRssi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdTxRetries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileThresholdRxErrors(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileIncludeWeatherChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerArrpProfileIncludeDfsChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerArrpProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerArrpProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerArrpProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("selection_period", flattenWirelessControllerArrpProfileSelectionPeriod(o["selection-period"], d, "selection_period", sv)); err != nil {
		if !fortiAPIPatch(o["selection-period"]) {
			return fmt.Errorf("Error reading selection_period: %v", err)
		}
	}

	if err = d.Set("monitor_period", flattenWirelessControllerArrpProfileMonitorPeriod(o["monitor-period"], d, "monitor_period", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-period"]) {
			return fmt.Errorf("Error reading monitor_period: %v", err)
		}
	}

	if err = d.Set("weight_managed_ap", flattenWirelessControllerArrpProfileWeightManagedAp(o["weight-managed-ap"], d, "weight_managed_ap", sv)); err != nil {
		if !fortiAPIPatch(o["weight-managed-ap"]) {
			return fmt.Errorf("Error reading weight_managed_ap: %v", err)
		}
	}

	if err = d.Set("weight_rogue_ap", flattenWirelessControllerArrpProfileWeightRogueAp(o["weight-rogue-ap"], d, "weight_rogue_ap", sv)); err != nil {
		if !fortiAPIPatch(o["weight-rogue-ap"]) {
			return fmt.Errorf("Error reading weight_rogue_ap: %v", err)
		}
	}

	if err = d.Set("weight_noise_floor", flattenWirelessControllerArrpProfileWeightNoiseFloor(o["weight-noise-floor"], d, "weight_noise_floor", sv)); err != nil {
		if !fortiAPIPatch(o["weight-noise-floor"]) {
			return fmt.Errorf("Error reading weight_noise_floor: %v", err)
		}
	}

	if err = d.Set("weight_channel_load", flattenWirelessControllerArrpProfileWeightChannelLoad(o["weight-channel-load"], d, "weight_channel_load", sv)); err != nil {
		if !fortiAPIPatch(o["weight-channel-load"]) {
			return fmt.Errorf("Error reading weight_channel_load: %v", err)
		}
	}

	if err = d.Set("weight_spectral_rssi", flattenWirelessControllerArrpProfileWeightSpectralRssi(o["weight-spectral-rssi"], d, "weight_spectral_rssi", sv)); err != nil {
		if !fortiAPIPatch(o["weight-spectral-rssi"]) {
			return fmt.Errorf("Error reading weight_spectral_rssi: %v", err)
		}
	}

	if err = d.Set("weight_weather_channel", flattenWirelessControllerArrpProfileWeightWeatherChannel(o["weight-weather-channel"], d, "weight_weather_channel", sv)); err != nil {
		if !fortiAPIPatch(o["weight-weather-channel"]) {
			return fmt.Errorf("Error reading weight_weather_channel: %v", err)
		}
	}

	if err = d.Set("weight_dfs_channel", flattenWirelessControllerArrpProfileWeightDfsChannel(o["weight-dfs-channel"], d, "weight_dfs_channel", sv)); err != nil {
		if !fortiAPIPatch(o["weight-dfs-channel"]) {
			return fmt.Errorf("Error reading weight_dfs_channel: %v", err)
		}
	}

	if err = d.Set("threshold_ap", flattenWirelessControllerArrpProfileThresholdAp(o["threshold-ap"], d, "threshold_ap", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-ap"]) {
			return fmt.Errorf("Error reading threshold_ap: %v", err)
		}
	}

	if err = d.Set("threshold_noise_floor", flattenWirelessControllerArrpProfileThresholdNoiseFloor(o["threshold-noise-floor"], d, "threshold_noise_floor", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-noise-floor"]) {
			return fmt.Errorf("Error reading threshold_noise_floor: %v", err)
		}
	}

	if err = d.Set("threshold_channel_load", flattenWirelessControllerArrpProfileThresholdChannelLoad(o["threshold-channel-load"], d, "threshold_channel_load", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-channel-load"]) {
			return fmt.Errorf("Error reading threshold_channel_load: %v", err)
		}
	}

	if err = d.Set("threshold_spectral_rssi", flattenWirelessControllerArrpProfileThresholdSpectralRssi(o["threshold-spectral-rssi"], d, "threshold_spectral_rssi", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-spectral-rssi"]) {
			return fmt.Errorf("Error reading threshold_spectral_rssi: %v", err)
		}
	}

	if err = d.Set("threshold_tx_retries", flattenWirelessControllerArrpProfileThresholdTxRetries(o["threshold-tx-retries"], d, "threshold_tx_retries", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-tx-retries"]) {
			return fmt.Errorf("Error reading threshold_tx_retries: %v", err)
		}
	}

	if err = d.Set("threshold_rx_errors", flattenWirelessControllerArrpProfileThresholdRxErrors(o["threshold-rx-errors"], d, "threshold_rx_errors", sv)); err != nil {
		if !fortiAPIPatch(o["threshold-rx-errors"]) {
			return fmt.Errorf("Error reading threshold_rx_errors: %v", err)
		}
	}

	if err = d.Set("include_weather_channel", flattenWirelessControllerArrpProfileIncludeWeatherChannel(o["include-weather-channel"], d, "include_weather_channel", sv)); err != nil {
		if !fortiAPIPatch(o["include-weather-channel"]) {
			return fmt.Errorf("Error reading include_weather_channel: %v", err)
		}
	}

	if err = d.Set("include_dfs_channel", flattenWirelessControllerArrpProfileIncludeDfsChannel(o["include-dfs-channel"], d, "include_dfs_channel", sv)); err != nil {
		if !fortiAPIPatch(o["include-dfs-channel"]) {
			return fmt.Errorf("Error reading include_dfs_channel: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerArrpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerArrpProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileSelectionPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileMonitorPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightManagedAp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightRogueAp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightNoiseFloor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightChannelLoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightSpectralRssi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightWeatherChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileWeightDfsChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdAp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdNoiseFloor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdChannelLoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdSpectralRssi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdTxRetries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileThresholdRxErrors(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileIncludeWeatherChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerArrpProfileIncludeDfsChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerArrpProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerArrpProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandWirelessControllerArrpProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOkExists("selection_period"); ok {

		t, err := expandWirelessControllerArrpProfileSelectionPeriod(d, v, "selection_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selection-period"] = t
		}
	}

	if v, ok := d.GetOkExists("monitor_period"); ok {

		t, err := expandWirelessControllerArrpProfileMonitorPeriod(d, v, "monitor_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-period"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_managed_ap"); ok {

		t, err := expandWirelessControllerArrpProfileWeightManagedAp(d, v, "weight_managed_ap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-managed-ap"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_rogue_ap"); ok {

		t, err := expandWirelessControllerArrpProfileWeightRogueAp(d, v, "weight_rogue_ap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-rogue-ap"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_noise_floor"); ok {

		t, err := expandWirelessControllerArrpProfileWeightNoiseFloor(d, v, "weight_noise_floor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-noise-floor"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_channel_load"); ok {

		t, err := expandWirelessControllerArrpProfileWeightChannelLoad(d, v, "weight_channel_load", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-channel-load"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_spectral_rssi"); ok {

		t, err := expandWirelessControllerArrpProfileWeightSpectralRssi(d, v, "weight_spectral_rssi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-spectral-rssi"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_weather_channel"); ok {

		t, err := expandWirelessControllerArrpProfileWeightWeatherChannel(d, v, "weight_weather_channel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-weather-channel"] = t
		}
	}

	if v, ok := d.GetOkExists("weight_dfs_channel"); ok {

		t, err := expandWirelessControllerArrpProfileWeightDfsChannel(d, v, "weight_dfs_channel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-dfs-channel"] = t
		}
	}

	if v, ok := d.GetOkExists("threshold_ap"); ok {

		t, err := expandWirelessControllerArrpProfileThresholdAp(d, v, "threshold_ap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-ap"] = t
		}
	}

	if v, ok := d.GetOk("threshold_noise_floor"); ok {

		t, err := expandWirelessControllerArrpProfileThresholdNoiseFloor(d, v, "threshold_noise_floor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-noise-floor"] = t
		}
	}

	if v, ok := d.GetOkExists("threshold_channel_load"); ok {

		t, err := expandWirelessControllerArrpProfileThresholdChannelLoad(d, v, "threshold_channel_load", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-channel-load"] = t
		}
	}

	if v, ok := d.GetOk("threshold_spectral_rssi"); ok {

		t, err := expandWirelessControllerArrpProfileThresholdSpectralRssi(d, v, "threshold_spectral_rssi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-spectral-rssi"] = t
		}
	}

	if v, ok := d.GetOkExists("threshold_tx_retries"); ok {

		t, err := expandWirelessControllerArrpProfileThresholdTxRetries(d, v, "threshold_tx_retries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-tx-retries"] = t
		}
	}

	if v, ok := d.GetOkExists("threshold_rx_errors"); ok {

		t, err := expandWirelessControllerArrpProfileThresholdRxErrors(d, v, "threshold_rx_errors", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold-rx-errors"] = t
		}
	}

	if v, ok := d.GetOk("include_weather_channel"); ok {

		t, err := expandWirelessControllerArrpProfileIncludeWeatherChannel(d, v, "include_weather_channel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-weather-channel"] = t
		}
	}

	if v, ok := d.GetOk("include_dfs_channel"); ok {

		t, err := expandWirelessControllerArrpProfileIncludeDfsChannel(d, v, "include_dfs_channel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-dfs-channel"] = t
		}
	}

	return &obj, nil
}
