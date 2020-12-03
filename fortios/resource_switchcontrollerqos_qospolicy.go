// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerQosQosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosQosPolicyCreate,
		Read:   resourceSwitchControllerQosQosPolicyRead,
		Update: resourceSwitchControllerQosQosPolicyUpdate,
		Delete: resourceSwitchControllerQosQosPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"default_cos": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Required:     true,
			},
			"trust_dot1p_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"trust_ip_dscp_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"queue_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchControllerQosQosPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerQosQosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerQosQosPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQosPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosQosPolicy")
	}

	return resourceSwitchControllerQosQosPolicyRead(d, m)
}

func resourceSwitchControllerQosQosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerQosQosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQosQosPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosQosPolicy")
	}

	return resourceSwitchControllerQosQosPolicyRead(d, m)
}

func resourceSwitchControllerQosQosPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerQosQosPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosQosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosQosPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerQosQosPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosQosPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosQosPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyDefaultCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyTrustDot1PMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyTrustIpDscpMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyQueuePolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosQosPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerQosQosPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("default_cos", flattenSwitchControllerQosQosPolicyDefaultCos(o["default-cos"], d, "default_cos")); err != nil {
		if !fortiAPIPatch(o["default-cos"]) {
			return fmt.Errorf("Error reading default_cos: %v", err)
		}
	}

	if err = d.Set("trust_dot1p_map", flattenSwitchControllerQosQosPolicyTrustDot1PMap(o["trust-dot1p-map"], d, "trust_dot1p_map")); err != nil {
		if !fortiAPIPatch(o["trust-dot1p-map"]) {
			return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
		}
	}

	if err = d.Set("trust_ip_dscp_map", flattenSwitchControllerQosQosPolicyTrustIpDscpMap(o["trust-ip-dscp-map"], d, "trust_ip_dscp_map")); err != nil {
		if !fortiAPIPatch(o["trust-ip-dscp-map"]) {
			return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
		}
	}

	if err = d.Set("queue_policy", flattenSwitchControllerQosQosPolicyQueuePolicy(o["queue-policy"], d, "queue_policy")); err != nil {
		if !fortiAPIPatch(o["queue-policy"]) {
			return fmt.Errorf("Error reading queue_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosQosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerQosQosPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyDefaultCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyTrustDot1PMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyTrustIpDscpMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyQueuePolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosQosPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerQosQosPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("default_cos"); ok {
		t, err := expandSwitchControllerQosQosPolicyDefaultCos(d, v, "default_cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cos"] = t
		}
	}

	if v, ok := d.GetOk("trust_dot1p_map"); ok {
		t, err := expandSwitchControllerQosQosPolicyTrustDot1PMap(d, v, "trust_dot1p_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-dot1p-map"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_dscp_map"); ok {
		t, err := expandSwitchControllerQosQosPolicyTrustIpDscpMap(d, v, "trust_ip_dscp_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-dscp-map"] = t
		}
	}

	if v, ok := d.GetOk("queue_policy"); ok {
		t, err := expandSwitchControllerQosQosPolicyQueuePolicy(d, v, "queue_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue-policy"] = t
		}
	}

	return &obj, nil
}
