// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure default auto-config QoS policy for FortiSwitch.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAutoConfigDefault() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigDefaultUpdate,
		Read:   resourceSwitchControllerAutoConfigDefaultRead,
		Update: resourceSwitchControllerAutoConfigDefaultUpdate,
		Delete: resourceSwitchControllerAutoConfigDefaultDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fgt_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"isl_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"icl_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigDefaultUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerAutoConfigDefault(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerAutoConfigDefault(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerAutoConfigDefault")
	}

	return resourceSwitchControllerAutoConfigDefaultRead(d, m)
}

func resourceSwitchControllerAutoConfigDefaultDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerAutoConfigDefault(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigDefault resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerAutoConfigDefault(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerAutoConfigDefault resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigDefaultRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerAutoConfigDefault(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigDefault resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigDefault(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigDefault resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigDefaultFgtPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigDefaultIslPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigDefaultIclPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fgt_policy", flattenSwitchControllerAutoConfigDefaultFgtPolicy(o["fgt-policy"], d, "fgt_policy", sv)); err != nil {
		if !fortiAPIPatch(o["fgt-policy"]) {
			return fmt.Errorf("Error reading fgt_policy: %v", err)
		}
	}

	if err = d.Set("isl_policy", flattenSwitchControllerAutoConfigDefaultIslPolicy(o["isl-policy"], d, "isl_policy", sv)); err != nil {
		if !fortiAPIPatch(o["isl-policy"]) {
			return fmt.Errorf("Error reading isl_policy: %v", err)
		}
	}

	if err = d.Set("icl_policy", flattenSwitchControllerAutoConfigDefaultIclPolicy(o["icl-policy"], d, "icl_policy", sv)); err != nil {
		if !fortiAPIPatch(o["icl-policy"]) {
			return fmt.Errorf("Error reading icl_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigDefaultFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerAutoConfigDefaultFgtPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigDefaultIslPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigDefaultIclPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigDefault(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fgt_policy"); ok {
		if setArgNil {
			obj["fgt-policy"] = nil
		} else {

			t, err := expandSwitchControllerAutoConfigDefaultFgtPolicy(d, v, "fgt_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fgt-policy"] = t
			}
		}
	}

	if v, ok := d.GetOk("isl_policy"); ok {
		if setArgNil {
			obj["isl-policy"] = nil
		} else {

			t, err := expandSwitchControllerAutoConfigDefaultIslPolicy(d, v, "isl_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["isl-policy"] = t
			}
		}
	}

	if v, ok := d.GetOk("icl_policy"); ok {
		if setArgNil {
			obj["icl-policy"] = nil
		} else {

			t, err := expandSwitchControllerAutoConfigDefaultIclPolicy(d, v, "icl_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["icl-policy"] = t
			}
		}
	}

	return &obj, nil
}
