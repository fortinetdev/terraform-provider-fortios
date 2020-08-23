// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IP to MAC binding settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIpmacbindingSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpmacbindingSettingUpdate,
		Read:   resourceFirewallIpmacbindingSettingRead,
		Update: resourceFirewallIpmacbindingSettingUpdate,
		Delete: resourceFirewallIpmacbindingSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bindthroughfw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bindtofw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"undefinedhost": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallIpmacbindingSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallIpmacbindingSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIpmacbindingSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIpmacbindingSetting")
	}

	return resourceFirewallIpmacbindingSettingRead(d, m)
}

func resourceFirewallIpmacbindingSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallIpmacbindingSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIpmacbindingSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpmacbindingSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallIpmacbindingSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpmacbindingSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingSetting resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpmacbindingSettingBindthroughfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingSettingBindtofw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingSettingUndefinedhost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIpmacbindingSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("bindthroughfw", flattenFirewallIpmacbindingSettingBindthroughfw(o["bindthroughfw"], d, "bindthroughfw")); err != nil {
		if !fortiAPIPatch(o["bindthroughfw"]) {
			return fmt.Errorf("Error reading bindthroughfw: %v", err)
		}
	}

	if err = d.Set("bindtofw", flattenFirewallIpmacbindingSettingBindtofw(o["bindtofw"], d, "bindtofw")); err != nil {
		if !fortiAPIPatch(o["bindtofw"]) {
			return fmt.Errorf("Error reading bindtofw: %v", err)
		}
	}

	if err = d.Set("undefinedhost", flattenFirewallIpmacbindingSettingUndefinedhost(o["undefinedhost"], d, "undefinedhost")); err != nil {
		if !fortiAPIPatch(o["undefinedhost"]) {
			return fmt.Errorf("Error reading undefinedhost: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpmacbindingSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIpmacbindingSettingBindthroughfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingSettingBindtofw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingSettingUndefinedhost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpmacbindingSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bindthroughfw"); ok {
		t, err := expandFirewallIpmacbindingSettingBindthroughfw(d, v, "bindthroughfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bindthroughfw"] = t
		}
	}

	if v, ok := d.GetOk("bindtofw"); ok {
		t, err := expandFirewallIpmacbindingSettingBindtofw(d, v, "bindtofw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bindtofw"] = t
		}
	}

	if v, ok := d.GetOk("undefinedhost"); ok {
		t, err := expandFirewallIpmacbindingSettingUndefinedhost(d, v, "undefinedhost")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["undefinedhost"] = t
		}
	}

	return &obj, nil
}
