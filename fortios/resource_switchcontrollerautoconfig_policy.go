// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch Auto-Config QoS policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"qos_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"storm_control_policy": &schema.Schema{
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
			"igmp_flood_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_flood_traffic": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerAutoConfigPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAutoConfigPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerAutoConfigPolicy(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerAutoConfigPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAutoConfigPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerAutoConfigPolicy(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerAutoConfigPolicy(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerAutoConfigPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAutoConfigPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAutoConfigPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAutoConfigPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyStormControlPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyPoeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyIgmpFloodReport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAutoConfigPolicyIgmpFloodTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerAutoConfigPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerAutoConfigPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenSwitchControllerAutoConfigPolicyQosPolicy(o["qos-policy"], d, "qos_policy", sv)); err != nil {
		if !fortiAPIPatch(o["qos-policy"]) {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("storm_control_policy", flattenSwitchControllerAutoConfigPolicyStormControlPolicy(o["storm-control-policy"], d, "storm_control_policy", sv)); err != nil {
		if !fortiAPIPatch(o["storm-control-policy"]) {
			return fmt.Errorf("Error reading storm_control_policy: %v", err)
		}
	}

	if err = d.Set("poe_status", flattenSwitchControllerAutoConfigPolicyPoeStatus(o["poe-status"], d, "poe_status", sv)); err != nil {
		if !fortiAPIPatch(o["poe-status"]) {
			return fmt.Errorf("Error reading poe_status: %v", err)
		}
	}

	if err = d.Set("igmp_flood_report", flattenSwitchControllerAutoConfigPolicyIgmpFloodReport(o["igmp-flood-report"], d, "igmp_flood_report", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-flood-report"]) {
			return fmt.Errorf("Error reading igmp_flood_report: %v", err)
		}
	}

	if err = d.Set("igmp_flood_traffic", flattenSwitchControllerAutoConfigPolicyIgmpFloodTraffic(o["igmp-flood-traffic"], d, "igmp_flood_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-flood-traffic"]) {
			return fmt.Errorf("Error reading igmp_flood_traffic: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerAutoConfigPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerAutoConfigPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyStormControlPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyPoeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyIgmpFloodReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAutoConfigPolicyIgmpFloodTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAutoConfigPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerAutoConfigPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok {

		t, err := expandSwitchControllerAutoConfigPolicyQosPolicy(d, v, "qos_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("storm_control_policy"); ok {

		t, err := expandSwitchControllerAutoConfigPolicyStormControlPolicy(d, v, "storm_control_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control-policy"] = t
		}
	}

	if v, ok := d.GetOk("poe_status"); ok {

		t, err := expandSwitchControllerAutoConfigPolicyPoeStatus(d, v, "poe_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-status"] = t
		}
	}

	if v, ok := d.GetOk("igmp_flood_report"); ok {

		t, err := expandSwitchControllerAutoConfigPolicyIgmpFloodReport(d, v, "igmp_flood_report", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-flood-report"] = t
		}
	}

	if v, ok := d.GetOk("igmp_flood_traffic"); ok {

		t, err := expandSwitchControllerAutoConfigPolicyIgmpFloodTraffic(d, v, "igmp_flood_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-flood-traffic"] = t
		}
	}

	return &obj, nil
}
