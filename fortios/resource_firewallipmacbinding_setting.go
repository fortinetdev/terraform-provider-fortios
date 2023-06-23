// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IP to MAC binding settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIpmacbindingSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIpmacbindingSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIpmacbindingSetting(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIpmacbindingSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallIpmacbindingSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallIpmacbindingSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpmacbindingSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallIpmacbindingSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpmacbindingSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingSetting resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpmacbindingSettingBindthroughfw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpmacbindingSettingBindtofw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpmacbindingSettingUndefinedhost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallIpmacbindingSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("bindthroughfw", flattenFirewallIpmacbindingSettingBindthroughfw(o["bindthroughfw"], d, "bindthroughfw", sv)); err != nil {
		if !fortiAPIPatch(o["bindthroughfw"]) {
			return fmt.Errorf("Error reading bindthroughfw: %v", err)
		}
	}

	if err = d.Set("bindtofw", flattenFirewallIpmacbindingSettingBindtofw(o["bindtofw"], d, "bindtofw", sv)); err != nil {
		if !fortiAPIPatch(o["bindtofw"]) {
			return fmt.Errorf("Error reading bindtofw: %v", err)
		}
	}

	if err = d.Set("undefinedhost", flattenFirewallIpmacbindingSettingUndefinedhost(o["undefinedhost"], d, "undefinedhost", sv)); err != nil {
		if !fortiAPIPatch(o["undefinedhost"]) {
			return fmt.Errorf("Error reading undefinedhost: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpmacbindingSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallIpmacbindingSettingBindthroughfw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingSettingBindtofw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingSettingUndefinedhost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpmacbindingSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bindthroughfw"); ok {
		if setArgNil {
			obj["bindthroughfw"] = nil
		} else {
			t, err := expandFirewallIpmacbindingSettingBindthroughfw(d, v, "bindthroughfw", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bindthroughfw"] = t
			}
		}
	}

	if v, ok := d.GetOk("bindtofw"); ok {
		if setArgNil {
			obj["bindtofw"] = nil
		} else {
			t, err := expandFirewallIpmacbindingSettingBindtofw(d, v, "bindtofw", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bindtofw"] = t
			}
		}
	}

	if v, ok := d.GetOk("undefinedhost"); ok {
		if setArgNil {
			obj["undefinedhost"] = nil
		} else {
			t, err := expandFirewallIpmacbindingSettingUndefinedhost(d, v, "undefinedhost", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["undefinedhost"] = t
			}
		}
	}

	return &obj, nil
}
