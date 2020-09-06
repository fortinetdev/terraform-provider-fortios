// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch Auto-Config QoS policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerAutoConfigPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAutoConfigPolicyCreate,
		Read:   resourceSwitchControllerAutoConfigPolicyRead,
		Update: resourceSwitchControllerAutoConfigPolicyUpdate,
		Delete: resourceSwitchControllerAutoConfigPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"qos_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"poe_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerAutoConfigPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerAutoConfigPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerAutoConfigPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerAutoConfigPolicy")
	}

	return resourceSwitchControllerAutoConfigPolicyRead(d, m)
}

func resourceSwitchControllerAutoConfigPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerAutoConfigPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerAutoConfigPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerAutoConfigPolicy")
	}

	return resourceSwitchControllerAutoConfigPolicyRead(d, m)
}

func resourceSwitchControllerAutoConfigPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerAutoConfigPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAutoConfigPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerAutoConfigPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyPoeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerAutoConfigPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenSwitchControllerAutoConfigPolicyQosPolicy(o["qos-policy"], d, "qos_policy")); err != nil {
		if !fortiAPIPatch(o["qos-policy"]) {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("poe_status", flattenSwitchControllerAutoConfigPolicyPoeStatus(o["poe-status"], d, "poe_status")); err != nil {
		if !fortiAPIPatch(o["poe-status"]) {
			return fmt.Errorf("Error reading poe_status: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerAutoConfigPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyPoeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerAutoConfigPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok {
		t, err := expandSwitchControllerAutoConfigPolicyQosPolicy(d, v, "qos_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("poe_status"); ok {
		t, err := expandSwitchControllerAutoConfigPolicyPoeStatus(d, v, "poe_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-status"] = t
		}
	}

	return &obj, nil
}
