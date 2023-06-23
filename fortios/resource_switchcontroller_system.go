// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure system-wide switch controller settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSystemUpdate,
		Read:   resourceSwitchControllerSystemRead,
		Update: resourceSwitchControllerSystemUpdate,
		Delete: resourceSwitchControllerSystemDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"parallel_process_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parallel_process": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"data_sync_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 1800),
				Optional:     true,
				Computed:     true,
			},
			"iot_weight_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_scan_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_holdoff": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"iot_mac_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nac_periodic_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 180),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_periodic_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 180),
				Optional:     true,
				Computed:     true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"caputp_echo_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(8, 600),
				Optional:     true,
				Computed:     true,
			},
			"caputp_max_retransmit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerSystemUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSystem(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSystem(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSystem")
	}

	return resourceSwitchControllerSystemRead(d, m)
}

func resourceSwitchControllerSystemDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSystem(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSystem resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSystem(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerSystem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSystemRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerSystem(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSystem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSystem(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSystem resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSystemParallelProcessOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemParallelProcess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemDataSyncInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotWeightThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotScanInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotHoldoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemIotMacIdle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemNacPeriodicInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemDynamicPeriodicInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemTunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemCaputpEchoInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSystemCaputpMaxRetransmit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSystem(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("parallel_process_override", flattenSwitchControllerSystemParallelProcessOverride(o["parallel-process-override"], d, "parallel_process_override", sv)); err != nil {
		if !fortiAPIPatch(o["parallel-process-override"]) {
			return fmt.Errorf("Error reading parallel_process_override: %v", err)
		}
	}

	if err = d.Set("parallel_process", flattenSwitchControllerSystemParallelProcess(o["parallel-process"], d, "parallel_process", sv)); err != nil {
		if !fortiAPIPatch(o["parallel-process"]) {
			return fmt.Errorf("Error reading parallel_process: %v", err)
		}
	}

	if err = d.Set("data_sync_interval", flattenSwitchControllerSystemDataSyncInterval(o["data-sync-interval"], d, "data_sync_interval", sv)); err != nil {
		if !fortiAPIPatch(o["data-sync-interval"]) {
			return fmt.Errorf("Error reading data_sync_interval: %v", err)
		}
	}

	if err = d.Set("iot_weight_threshold", flattenSwitchControllerSystemIotWeightThreshold(o["iot-weight-threshold"], d, "iot_weight_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["iot-weight-threshold"]) {
			return fmt.Errorf("Error reading iot_weight_threshold: %v", err)
		}
	}

	if err = d.Set("iot_scan_interval", flattenSwitchControllerSystemIotScanInterval(o["iot-scan-interval"], d, "iot_scan_interval", sv)); err != nil {
		if !fortiAPIPatch(o["iot-scan-interval"]) {
			return fmt.Errorf("Error reading iot_scan_interval: %v", err)
		}
	}

	if err = d.Set("iot_holdoff", flattenSwitchControllerSystemIotHoldoff(o["iot-holdoff"], d, "iot_holdoff", sv)); err != nil {
		if !fortiAPIPatch(o["iot-holdoff"]) {
			return fmt.Errorf("Error reading iot_holdoff: %v", err)
		}
	}

	if err = d.Set("iot_mac_idle", flattenSwitchControllerSystemIotMacIdle(o["iot-mac-idle"], d, "iot_mac_idle", sv)); err != nil {
		if !fortiAPIPatch(o["iot-mac-idle"]) {
			return fmt.Errorf("Error reading iot_mac_idle: %v", err)
		}
	}

	if err = d.Set("nac_periodic_interval", flattenSwitchControllerSystemNacPeriodicInterval(o["nac-periodic-interval"], d, "nac_periodic_interval", sv)); err != nil {
		if !fortiAPIPatch(o["nac-periodic-interval"]) {
			return fmt.Errorf("Error reading nac_periodic_interval: %v", err)
		}
	}

	if err = d.Set("dynamic_periodic_interval", flattenSwitchControllerSystemDynamicPeriodicInterval(o["dynamic-periodic-interval"], d, "dynamic_periodic_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-periodic-interval"]) {
			return fmt.Errorf("Error reading dynamic_periodic_interval: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenSwitchControllerSystemTunnelMode(o["tunnel-mode"], d, "tunnel_mode", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-mode"]) {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("caputp_echo_interval", flattenSwitchControllerSystemCaputpEchoInterval(o["caputp-echo-interval"], d, "caputp_echo_interval", sv)); err != nil {
		if !fortiAPIPatch(o["caputp-echo-interval"]) {
			return fmt.Errorf("Error reading caputp_echo_interval: %v", err)
		}
	}

	if err = d.Set("caputp_max_retransmit", flattenSwitchControllerSystemCaputpMaxRetransmit(o["caputp-max-retransmit"], d, "caputp_max_retransmit", sv)); err != nil {
		if !fortiAPIPatch(o["caputp-max-retransmit"]) {
			return fmt.Errorf("Error reading caputp_max_retransmit: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSystemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSystemParallelProcessOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemParallelProcess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemDataSyncInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotWeightThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotScanInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotHoldoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemIotMacIdle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemNacPeriodicInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemDynamicPeriodicInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemTunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemCaputpEchoInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSystemCaputpMaxRetransmit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSystem(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("parallel_process_override"); ok {
		if setArgNil {
			obj["parallel-process-override"] = nil
		} else {
			t, err := expandSwitchControllerSystemParallelProcessOverride(d, v, "parallel_process_override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["parallel-process-override"] = t
			}
		}
	}

	if v, ok := d.GetOk("parallel_process"); ok {
		if setArgNil {
			obj["parallel-process"] = nil
		} else {
			t, err := expandSwitchControllerSystemParallelProcess(d, v, "parallel_process", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["parallel-process"] = t
			}
		}
	}

	if v, ok := d.GetOk("data_sync_interval"); ok {
		if setArgNil {
			obj["data-sync-interval"] = nil
		} else {
			t, err := expandSwitchControllerSystemDataSyncInterval(d, v, "data_sync_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["data-sync-interval"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("iot_weight_threshold"); ok {
		if setArgNil {
			obj["iot-weight-threshold"] = nil
		} else {
			t, err := expandSwitchControllerSystemIotWeightThreshold(d, v, "iot_weight_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["iot-weight-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("iot_scan_interval"); ok {
		if setArgNil {
			obj["iot-scan-interval"] = nil
		} else {
			t, err := expandSwitchControllerSystemIotScanInterval(d, v, "iot_scan_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["iot-scan-interval"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("iot_holdoff"); ok {
		if setArgNil {
			obj["iot-holdoff"] = nil
		} else {
			t, err := expandSwitchControllerSystemIotHoldoff(d, v, "iot_holdoff", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["iot-holdoff"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("iot_mac_idle"); ok {
		if setArgNil {
			obj["iot-mac-idle"] = nil
		} else {
			t, err := expandSwitchControllerSystemIotMacIdle(d, v, "iot_mac_idle", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["iot-mac-idle"] = t
			}
		}
	}

	if v, ok := d.GetOk("nac_periodic_interval"); ok {
		if setArgNil {
			obj["nac-periodic-interval"] = nil
		} else {
			t, err := expandSwitchControllerSystemNacPeriodicInterval(d, v, "nac_periodic_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nac-periodic-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("dynamic_periodic_interval"); ok {
		if setArgNil {
			obj["dynamic-periodic-interval"] = nil
		} else {
			t, err := expandSwitchControllerSystemDynamicPeriodicInterval(d, v, "dynamic_periodic_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dynamic-periodic-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok {
		if setArgNil {
			obj["tunnel-mode"] = nil
		} else {
			t, err := expandSwitchControllerSystemTunnelMode(d, v, "tunnel_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("caputp_echo_interval"); ok {
		if setArgNil {
			obj["caputp-echo-interval"] = nil
		} else {
			t, err := expandSwitchControllerSystemCaputpEchoInterval(d, v, "caputp_echo_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["caputp-echo-interval"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("caputp_max_retransmit"); ok {
		if setArgNil {
			obj["caputp-max-retransmit"] = nil
		} else {
			t, err := expandSwitchControllerSystemCaputpMaxRetransmit(d, v, "caputp_max_retransmit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["caputp-max-retransmit"] = t
			}
		}
	}

	return &obj, nil
}
