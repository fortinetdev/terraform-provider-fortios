// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure shared traffic shaper.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallShaperTrafficShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallShaperTrafficShaperCreate,
		Read:   resourceFirewallShaperTrafficShaperRead,
		Update: resourceFirewallShaperTrafficShaperUpdate,
		Delete: resourceFirewallShaperTrafficShaperDelete,

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
			"guaranteed_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"maximum_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallShaperTrafficShaperCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallShaperTrafficShaper(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperTrafficShaper resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallShaperTrafficShaper(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallShaperTrafficShaper resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShaperTrafficShaper")
	}

	return resourceFirewallShaperTrafficShaperRead(d, m)
}

func resourceFirewallShaperTrafficShaperUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallShaperTrafficShaper(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperTrafficShaper resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallShaperTrafficShaper(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallShaperTrafficShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallShaperTrafficShaper")
	}

	return resourceFirewallShaperTrafficShaperRead(d, m)
}

func resourceFirewallShaperTrafficShaperDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallShaperTrafficShaper(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallShaperTrafficShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallShaperTrafficShaperRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallShaperTrafficShaper(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperTrafficShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallShaperTrafficShaper(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallShaperTrafficShaper resource from API: %v", err)
	}
	return nil
}

func flattenFirewallShaperTrafficShaperName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperMaximumBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperPerPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallShaperTrafficShaperDiffservcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallShaperTrafficShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallShaperTrafficShaperName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenFirewallShaperTrafficShaperGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if !fortiAPIPatch(o["guaranteed-bandwidth"]) {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("maximum_bandwidth", flattenFirewallShaperTrafficShaperMaximumBandwidth(o["maximum-bandwidth"], d, "maximum_bandwidth")); err != nil {
		if !fortiAPIPatch(o["maximum-bandwidth"]) {
			return fmt.Errorf("Error reading maximum_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_unit", flattenFirewallShaperTrafficShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit")); err != nil {
		if !fortiAPIPatch(o["bandwidth-unit"]) {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("priority", flattenFirewallShaperTrafficShaperPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("per_policy", flattenFirewallShaperTrafficShaperPerPolicy(o["per-policy"], d, "per_policy")); err != nil {
		if !fortiAPIPatch(o["per-policy"]) {
			return fmt.Errorf("Error reading per_policy: %v", err)
		}
	}

	if err = d.Set("diffserv", flattenFirewallShaperTrafficShaperDiffserv(o["diffserv"], d, "diffserv")); err != nil {
		if !fortiAPIPatch(o["diffserv"]) {
			return fmt.Errorf("Error reading diffserv: %v", err)
		}
	}

	if err = d.Set("diffservcode", flattenFirewallShaperTrafficShaperDiffservcode(o["diffservcode"], d, "diffservcode")); err != nil {
		if !fortiAPIPatch(o["diffservcode"]) {
			return fmt.Errorf("Error reading diffservcode: %v", err)
		}
	}

	return nil
}

func flattenFirewallShaperTrafficShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallShaperTrafficShaperName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperMaximumBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperBandwidthUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperPerPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallShaperTrafficShaperDiffservcode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallShaperTrafficShaper(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallShaperTrafficShaperName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("guaranteed_bandwidth"); ok {
		t, err := expandFirewallShaperTrafficShaperGuaranteedBandwidth(d, v, "guaranteed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOkExists("maximum_bandwidth"); ok {
		t, err := expandFirewallShaperTrafficShaperMaximumBandwidth(d, v, "maximum_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_unit"); ok {
		t, err := expandFirewallShaperTrafficShaperBandwidthUnit(d, v, "bandwidth_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-unit"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandFirewallShaperTrafficShaperPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("per_policy"); ok {
		t, err := expandFirewallShaperTrafficShaperPerPolicy(d, v, "per_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy"] = t
		}
	}

	if v, ok := d.GetOk("diffserv"); ok {
		t, err := expandFirewallShaperTrafficShaperDiffserv(d, v, "diffserv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode"); ok {
		t, err := expandFirewallShaperTrafficShaperDiffservcode(d, v, "diffservcode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode"] = t
		}
	}

	return &obj, nil
}
