// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure wireless controller event log filters.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerLogUpdate,
		Read:   resourceWirelessControllerLogRead,
		Update: resourceWirelessControllerLogUpdate,
		Delete: resourceWirelessControllerLogDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addrgrp_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ble_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clb_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_starv_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"led_sched_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rogue_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sta_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sta_locate_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wids_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wtp_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wtp_fips_event_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerLogUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerLog(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLog resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerLog(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerLog")
	}

	return resourceWirelessControllerLogRead(d, m)
}

func resourceWirelessControllerLogDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerLog(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerLog resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerLog(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WirelessControllerLog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerLogRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerLog(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerLog(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerLog resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerLogStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogAddrgrpLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogBleLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogClbLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogDhcpStarvLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogLedSchedLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogRadioEventLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogRogueEventLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogStaEventLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogStaLocateLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogWidsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogWtpEventLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerLogWtpFipsEventLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerLog(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenWirelessControllerLogStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("addrgrp_log", flattenWirelessControllerLogAddrgrpLog(o["addrgrp-log"], d, "addrgrp_log", sv)); err != nil {
		if !fortiAPIPatch(o["addrgrp-log"]) {
			return fmt.Errorf("Error reading addrgrp_log: %v", err)
		}
	}

	if err = d.Set("ble_log", flattenWirelessControllerLogBleLog(o["ble-log"], d, "ble_log", sv)); err != nil {
		if !fortiAPIPatch(o["ble-log"]) {
			return fmt.Errorf("Error reading ble_log: %v", err)
		}
	}

	if err = d.Set("clb_log", flattenWirelessControllerLogClbLog(o["clb-log"], d, "clb_log", sv)); err != nil {
		if !fortiAPIPatch(o["clb-log"]) {
			return fmt.Errorf("Error reading clb_log: %v", err)
		}
	}

	if err = d.Set("dhcp_starv_log", flattenWirelessControllerLogDhcpStarvLog(o["dhcp-starv-log"], d, "dhcp_starv_log", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-starv-log"]) {
			return fmt.Errorf("Error reading dhcp_starv_log: %v", err)
		}
	}

	if err = d.Set("led_sched_log", flattenWirelessControllerLogLedSchedLog(o["led-sched-log"], d, "led_sched_log", sv)); err != nil {
		if !fortiAPIPatch(o["led-sched-log"]) {
			return fmt.Errorf("Error reading led_sched_log: %v", err)
		}
	}

	if err = d.Set("radio_event_log", flattenWirelessControllerLogRadioEventLog(o["radio-event-log"], d, "radio_event_log", sv)); err != nil {
		if !fortiAPIPatch(o["radio-event-log"]) {
			return fmt.Errorf("Error reading radio_event_log: %v", err)
		}
	}

	if err = d.Set("rogue_event_log", flattenWirelessControllerLogRogueEventLog(o["rogue-event-log"], d, "rogue_event_log", sv)); err != nil {
		if !fortiAPIPatch(o["rogue-event-log"]) {
			return fmt.Errorf("Error reading rogue_event_log: %v", err)
		}
	}

	if err = d.Set("sta_event_log", flattenWirelessControllerLogStaEventLog(o["sta-event-log"], d, "sta_event_log", sv)); err != nil {
		if !fortiAPIPatch(o["sta-event-log"]) {
			return fmt.Errorf("Error reading sta_event_log: %v", err)
		}
	}

	if err = d.Set("sta_locate_log", flattenWirelessControllerLogStaLocateLog(o["sta-locate-log"], d, "sta_locate_log", sv)); err != nil {
		if !fortiAPIPatch(o["sta-locate-log"]) {
			return fmt.Errorf("Error reading sta_locate_log: %v", err)
		}
	}

	if err = d.Set("wids_log", flattenWirelessControllerLogWidsLog(o["wids-log"], d, "wids_log", sv)); err != nil {
		if !fortiAPIPatch(o["wids-log"]) {
			return fmt.Errorf("Error reading wids_log: %v", err)
		}
	}

	if err = d.Set("wtp_event_log", flattenWirelessControllerLogWtpEventLog(o["wtp-event-log"], d, "wtp_event_log", sv)); err != nil {
		if !fortiAPIPatch(o["wtp-event-log"]) {
			return fmt.Errorf("Error reading wtp_event_log: %v", err)
		}
	}

	if err = d.Set("wtp_fips_event_log", flattenWirelessControllerLogWtpFipsEventLog(o["wtp-fips-event-log"], d, "wtp_fips_event_log", sv)); err != nil {
		if !fortiAPIPatch(o["wtp-fips-event-log"]) {
			return fmt.Errorf("Error reading wtp_fips_event_log: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerLogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerLogStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogAddrgrpLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogBleLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogClbLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogDhcpStarvLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogLedSchedLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogRadioEventLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogRogueEventLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogStaEventLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogStaLocateLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogWidsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogWtpEventLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerLogWtpFipsEventLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerLog(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandWirelessControllerLogStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("addrgrp_log"); ok {
		if setArgNil {
			obj["addrgrp-log"] = nil
		} else {
			t, err := expandWirelessControllerLogAddrgrpLog(d, v, "addrgrp_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["addrgrp-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("ble_log"); ok {
		if setArgNil {
			obj["ble-log"] = nil
		} else {
			t, err := expandWirelessControllerLogBleLog(d, v, "ble_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ble-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("clb_log"); ok {
		if setArgNil {
			obj["clb-log"] = nil
		} else {
			t, err := expandWirelessControllerLogClbLog(d, v, "clb_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["clb-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_starv_log"); ok {
		if setArgNil {
			obj["dhcp-starv-log"] = nil
		} else {
			t, err := expandWirelessControllerLogDhcpStarvLog(d, v, "dhcp_starv_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-starv-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("led_sched_log"); ok {
		if setArgNil {
			obj["led-sched-log"] = nil
		} else {
			t, err := expandWirelessControllerLogLedSchedLog(d, v, "led_sched_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["led-sched-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("radio_event_log"); ok {
		if setArgNil {
			obj["radio-event-log"] = nil
		} else {
			t, err := expandWirelessControllerLogRadioEventLog(d, v, "radio_event_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radio-event-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("rogue_event_log"); ok {
		if setArgNil {
			obj["rogue-event-log"] = nil
		} else {
			t, err := expandWirelessControllerLogRogueEventLog(d, v, "rogue_event_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rogue-event-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("sta_event_log"); ok {
		if setArgNil {
			obj["sta-event-log"] = nil
		} else {
			t, err := expandWirelessControllerLogStaEventLog(d, v, "sta_event_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sta-event-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("sta_locate_log"); ok {
		if setArgNil {
			obj["sta-locate-log"] = nil
		} else {
			t, err := expandWirelessControllerLogStaLocateLog(d, v, "sta_locate_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sta-locate-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("wids_log"); ok {
		if setArgNil {
			obj["wids-log"] = nil
		} else {
			t, err := expandWirelessControllerLogWidsLog(d, v, "wids_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wids-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("wtp_event_log"); ok {
		if setArgNil {
			obj["wtp-event-log"] = nil
		} else {
			t, err := expandWirelessControllerLogWtpEventLog(d, v, "wtp_event_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wtp-event-log"] = t
			}
		}
	}

	if v, ok := d.GetOk("wtp_fips_event_log"); ok {
		if setArgNil {
			obj["wtp-fips-event-log"] = nil
		} else {
			t, err := expandWirelessControllerLogWtpFipsEventLog(d, v, "wtp_fips_event_log", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wtp-fips-event-log"] = t
			}
		}
	}

	return &obj, nil
}
