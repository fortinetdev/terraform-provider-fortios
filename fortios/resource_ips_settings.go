// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS VDOM parameter.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsSettingsUpdate,
		Read:   resourceIpsSettingsRead,
		Update: resourceIpsSettingsUpdate,
		Delete: resourceIpsSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"packet_log_history": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"packet_log_post_attack": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"packet_log_memory": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 8192),
				Optional:     true,
				Computed:     true,
			},
			"ips_packet_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIpsSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IpsSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IpsSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsSettings")
	}

	return resourceIpsSettingsRead(d, m)
}

func resourceIpsSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating IpsSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateIpsSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing IpsSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadIpsSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IpsSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IpsSettings resource from API: %v", err)
	}
	return nil
}

func flattenIpsSettingsPacketLogHistory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogPostAttack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogMemory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsSettingsIpsPacketQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIpsSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("packet_log_history", flattenIpsSettingsPacketLogHistory(o["packet-log-history"], d, "packet_log_history", sv)); err != nil {
		if !fortiAPIPatch(o["packet-log-history"]) {
			return fmt.Errorf("Error reading packet_log_history: %v", err)
		}
	}

	if err = d.Set("packet_log_post_attack", flattenIpsSettingsPacketLogPostAttack(o["packet-log-post-attack"], d, "packet_log_post_attack", sv)); err != nil {
		if !fortiAPIPatch(o["packet-log-post-attack"]) {
			return fmt.Errorf("Error reading packet_log_post_attack: %v", err)
		}
	}

	if err = d.Set("packet_log_memory", flattenIpsSettingsPacketLogMemory(o["packet-log-memory"], d, "packet_log_memory", sv)); err != nil {
		if !fortiAPIPatch(o["packet-log-memory"]) {
			return fmt.Errorf("Error reading packet_log_memory: %v", err)
		}
	}

	if err = d.Set("ips_packet_quota", flattenIpsSettingsIpsPacketQuota(o["ips-packet-quota"], d, "ips_packet_quota", sv)); err != nil {
		if !fortiAPIPatch(o["ips-packet-quota"]) {
			return fmt.Errorf("Error reading ips_packet_quota: %v", err)
		}
	}

	return nil
}

func flattenIpsSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIpsSettingsPacketLogHistory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogPostAttack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogMemory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsIpsPacketQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIpsSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("packet_log_history"); ok {
		if setArgNil {
			obj["packet-log-history"] = nil
		} else {

			t, err := expandIpsSettingsPacketLogHistory(d, v, "packet_log_history", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["packet-log-history"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("packet_log_post_attack"); ok {
		if setArgNil {
			obj["packet-log-post-attack"] = nil
		} else {

			t, err := expandIpsSettingsPacketLogPostAttack(d, v, "packet_log_post_attack", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["packet-log-post-attack"] = t
			}
		}
	}

	if v, ok := d.GetOk("packet_log_memory"); ok {
		if setArgNil {
			obj["packet-log-memory"] = nil
		} else {

			t, err := expandIpsSettingsPacketLogMemory(d, v, "packet_log_memory", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["packet-log-memory"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("ips_packet_quota"); ok {
		if setArgNil {
			obj["ips-packet-quota"] = nil
		} else {

			t, err := expandIpsSettingsIpsPacketQuota(d, v, "ips_packet_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ips-packet-quota"] = t
			}
		}
	}

	return &obj, nil
}
