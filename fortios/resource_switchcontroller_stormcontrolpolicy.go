// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch storm control policy to be applied on managed-switch ports.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerStormControlPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerStormControlPolicyCreate,
		Read:   resourceSwitchControllerStormControlPolicyRead,
		Update: resourceSwitchControllerStormControlPolicyUpdate,
		Delete: resourceSwitchControllerStormControlPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"storm_control_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),
				Optional:     true,
				Computed:     true,
			},
			"unknown_unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerStormControlPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerStormControlPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerStormControlPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerStormControlPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerStormControlPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerStormControlPolicy")
	}

	return resourceSwitchControllerStormControlPolicyRead(d, m)
}

func resourceSwitchControllerStormControlPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerStormControlPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControlPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerStormControlPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerStormControlPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerStormControlPolicy")
	}

	return resourceSwitchControllerStormControlPolicyRead(d, m)
}

func resourceSwitchControllerStormControlPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerStormControlPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerStormControlPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerStormControlPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerStormControlPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControlPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerStormControlPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerStormControlPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerStormControlPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyStormControlMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyUnknownUnicast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerStormControlPolicyBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerStormControlPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerStormControlPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerStormControlPolicyDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("storm_control_mode", flattenSwitchControllerStormControlPolicyStormControlMode(o["storm-control-mode"], d, "storm_control_mode", sv)); err != nil {
		if !fortiAPIPatch(o["storm-control-mode"]) {
			return fmt.Errorf("Error reading storm_control_mode: %v", err)
		}
	}

	if err = d.Set("rate", flattenSwitchControllerStormControlPolicyRate(o["rate"], d, "rate", sv)); err != nil {
		if !fortiAPIPatch(o["rate"]) {
			return fmt.Errorf("Error reading rate: %v", err)
		}
	}

	if err = d.Set("unknown_unicast", flattenSwitchControllerStormControlPolicyUnknownUnicast(o["unknown-unicast"], d, "unknown_unicast", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-unicast"]) {
			return fmt.Errorf("Error reading unknown_unicast: %v", err)
		}
	}

	if err = d.Set("unknown_multicast", flattenSwitchControllerStormControlPolicyUnknownMulticast(o["unknown-multicast"], d, "unknown_multicast", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-multicast"]) {
			return fmt.Errorf("Error reading unknown_multicast: %v", err)
		}
	}

	if err = d.Set("broadcast", flattenSwitchControllerStormControlPolicyBroadcast(o["broadcast"], d, "broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["broadcast"]) {
			return fmt.Errorf("Error reading broadcast: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerStormControlPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerStormControlPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyStormControlMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyUnknownUnicast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerStormControlPolicyBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerStormControlPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerStormControlPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchControllerStormControlPolicyDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("storm_control_mode"); ok {

		t, err := expandSwitchControllerStormControlPolicyStormControlMode(d, v, "storm_control_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control-mode"] = t
		}
	}

	if v, ok := d.GetOkExists("rate"); ok {

		t, err := expandSwitchControllerStormControlPolicyRate(d, v, "rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate"] = t
		}
	}

	if v, ok := d.GetOk("unknown_unicast"); ok {

		t, err := expandSwitchControllerStormControlPolicyUnknownUnicast(d, v, "unknown_unicast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-unicast"] = t
		}
	}

	if v, ok := d.GetOk("unknown_multicast"); ok {

		t, err := expandSwitchControllerStormControlPolicyUnknownMulticast(d, v, "unknown_multicast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-multicast"] = t
		}
	}

	if v, ok := d.GetOk("broadcast"); ok {

		t, err := expandSwitchControllerStormControlPolicyBroadcast(d, v, "broadcast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast"] = t
		}
	}

	return &obj, nil
}
