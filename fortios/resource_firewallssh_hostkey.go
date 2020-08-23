// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: SSH proxy host public keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallSshHostKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSshHostKeyCreate,
		Read:   resourceFirewallSshHostKeyRead,
		Update: resourceFirewallSshHostKeyUpdate,
		Delete: resourceFirewallSshHostKeyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"public_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),
				Optional:     true,
			},
		},
	}
}

func resourceFirewallSshHostKeyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSshHostKey(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSshHostKey resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallSshHostKey(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallSshHostKey resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSshHostKey")
	}

	return resourceFirewallSshHostKeyRead(d, m)
}

func resourceFirewallSshHostKeyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSshHostKey(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshHostKey resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSshHostKey(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshHostKey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSshHostKey")
	}

	return resourceFirewallSshHostKeyRead(d, m)
}

func resourceFirewallSshHostKeyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallSshHostKey(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSshHostKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshHostKeyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallSshHostKey(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshHostKey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSshHostKey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshHostKey resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSshHostKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyNid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSshHostKeyPublicKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSshHostKey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallSshHostKeyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSshHostKeyStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallSshHostKeyType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("nid", flattenFirewallSshHostKeyNid(o["nid"], d, "nid")); err != nil {
		if !fortiAPIPatch(o["nid"]) {
			return fmt.Errorf("Error reading nid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallSshHostKeyIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallSshHostKeyPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("hostname", flattenFirewallSshHostKeyHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("public_key", flattenFirewallSshHostKeyPublicKey(o["public-key"], d, "public_key")); err != nil {
		if !fortiAPIPatch(o["public-key"]) {
			return fmt.Errorf("Error reading public_key: %v", err)
		}
	}

	return nil
}

func flattenFirewallSshHostKeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSshHostKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyNid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshHostKeyPublicKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSshHostKey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallSshHostKeyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallSshHostKeyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandFirewallSshHostKeyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("nid"); ok {
		t, err := expandFirewallSshHostKeyNid(d, v, "nid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nid"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFirewallSshHostKeyIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandFirewallSshHostKeyPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandFirewallSshHostKeyHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("public_key"); ok {
		t, err := expandFirewallSshHostKeyPublicKey(d, v, "public_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["public-key"] = t
		}
	}

	return &obj, nil
}
