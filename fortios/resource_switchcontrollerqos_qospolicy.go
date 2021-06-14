// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch QoS policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Required:     true,
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerQosQosPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerQosQosPolicy(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerQosQosPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQosPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQosQosPolicy(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerQosQosPolicy(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerQosQosPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosQosPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosQosPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyDefaultCos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyTrustDot1PMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyTrustIpDscpMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQosPolicyQueuePolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerQosQosPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerQosQosPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("default_cos", flattenSwitchControllerQosQosPolicyDefaultCos(o["default-cos"], d, "default_cos", sv)); err != nil {
		if !fortiAPIPatch(o["default-cos"]) {
			return fmt.Errorf("Error reading default_cos: %v", err)
		}
	}

	if err = d.Set("trust_dot1p_map", flattenSwitchControllerQosQosPolicyTrustDot1PMap(o["trust-dot1p-map"], d, "trust_dot1p_map", sv)); err != nil {
		if !fortiAPIPatch(o["trust-dot1p-map"]) {
			return fmt.Errorf("Error reading trust_dot1p_map: %v", err)
		}
	}

	if err = d.Set("trust_ip_dscp_map", flattenSwitchControllerQosQosPolicyTrustIpDscpMap(o["trust-ip-dscp-map"], d, "trust_ip_dscp_map", sv)); err != nil {
		if !fortiAPIPatch(o["trust-ip-dscp-map"]) {
			return fmt.Errorf("Error reading trust_ip_dscp_map: %v", err)
		}
	}

	if err = d.Set("queue_policy", flattenSwitchControllerQosQosPolicyQueuePolicy(o["queue-policy"], d, "queue_policy", sv)); err != nil {
		if !fortiAPIPatch(o["queue-policy"]) {
			return fmt.Errorf("Error reading queue_policy: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerQosQosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerQosQosPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyDefaultCos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyTrustDot1PMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyTrustIpDscpMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQosPolicyQueuePolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosQosPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerQosQosPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("default_cos"); ok {

		t, err := expandSwitchControllerQosQosPolicyDefaultCos(d, v, "default_cos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-cos"] = t
		}
	}

	if v, ok := d.GetOk("trust_dot1p_map"); ok {

		t, err := expandSwitchControllerQosQosPolicyTrustDot1PMap(d, v, "trust_dot1p_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-dot1p-map"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_dscp_map"); ok {

		t, err := expandSwitchControllerQosQosPolicyTrustIpDscpMap(d, v, "trust_ip_dscp_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-dscp-map"] = t
		}
	}

	if v, ok := d.GetOk("queue_policy"); ok {

		t, err := expandSwitchControllerQosQosPolicyQueuePolicy(d, v, "queue_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queue-policy"] = t
		}
	}

	return &obj, nil
}
