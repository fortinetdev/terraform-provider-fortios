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

	obj, err := getObjectIpsSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsSettings(obj, mkey)
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

	err := c.DeleteIpsSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting IpsSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadIpsSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IpsSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsSettings resource from API: %v", err)
	}
	return nil
}

func flattenIpsSettingsPacketLogHistory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogPostAttack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsPacketLogMemory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsSettingsIpsPacketQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("packet_log_history", flattenIpsSettingsPacketLogHistory(o["packet-log-history"], d, "packet_log_history")); err != nil {
		if !fortiAPIPatch(o["packet-log-history"]) {
			return fmt.Errorf("Error reading packet_log_history: %v", err)
		}
	}

	if err = d.Set("packet_log_post_attack", flattenIpsSettingsPacketLogPostAttack(o["packet-log-post-attack"], d, "packet_log_post_attack")); err != nil {
		if !fortiAPIPatch(o["packet-log-post-attack"]) {
			return fmt.Errorf("Error reading packet_log_post_attack: %v", err)
		}
	}

	if err = d.Set("packet_log_memory", flattenIpsSettingsPacketLogMemory(o["packet-log-memory"], d, "packet_log_memory")); err != nil {
		if !fortiAPIPatch(o["packet-log-memory"]) {
			return fmt.Errorf("Error reading packet_log_memory: %v", err)
		}
	}

	if err = d.Set("ips_packet_quota", flattenIpsSettingsIpsPacketQuota(o["ips-packet-quota"], d, "ips_packet_quota")); err != nil {
		if !fortiAPIPatch(o["ips-packet-quota"]) {
			return fmt.Errorf("Error reading ips_packet_quota: %v", err)
		}
	}

	return nil
}

func flattenIpsSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsSettingsPacketLogHistory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogPostAttack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsPacketLogMemory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsSettingsIpsPacketQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("packet_log_history"); ok {
		t, err := expandIpsSettingsPacketLogHistory(d, v, "packet_log_history")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-history"] = t
		}
	}

	if v, ok := d.GetOkExists("packet_log_post_attack"); ok {
		t, err := expandIpsSettingsPacketLogPostAttack(d, v, "packet_log_post_attack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-post-attack"] = t
		}
	}

	if v, ok := d.GetOk("packet_log_memory"); ok {
		t, err := expandIpsSettingsPacketLogMemory(d, v, "packet_log_memory")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-log-memory"] = t
		}
	}

	if v, ok := d.GetOkExists("ips_packet_quota"); ok {
		t, err := expandIpsSettingsIpsPacketQuota(d, v, "ips_packet_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-packet-quota"] = t
		}
	}

	return &obj, nil
}
