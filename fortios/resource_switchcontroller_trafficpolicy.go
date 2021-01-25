// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch traffic policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerTrafficPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficPolicyCreate,
		Read:   resourceSwitchControllerTrafficPolicyRead,
		Update: resourceSwitchControllerTrafficPolicyUpdate,
		Delete: resourceSwitchControllerTrafficPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Required:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"policer_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guaranteed_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 524287000),
				Optional:     true,
				Computed:     true,
			},
			"guaranteed_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maximum_burst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cos": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerTrafficPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerTrafficPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerTrafficPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerTrafficPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerTrafficPolicy")
	}

	return resourceSwitchControllerTrafficPolicyRead(d, m)
}

func resourceSwitchControllerTrafficPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerTrafficPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerTrafficPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerTrafficPolicy")
	}

	return resourceSwitchControllerTrafficPolicyRead(d, m)
}

func resourceSwitchControllerTrafficPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerTrafficPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerTrafficPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyPolicerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyGuaranteedBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyGuaranteedBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyMaximumBurst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerTrafficPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerTrafficPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerTrafficPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerTrafficPolicyDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("policer_status", flattenSwitchControllerTrafficPolicyPolicerStatus(o["policer-status"], d, "policer_status")); err != nil {
		if !fortiAPIPatch(o["policer-status"]) {
			return fmt.Errorf("Error reading policer_status: %v", err)
		}
	}

	if err = d.Set("guaranteed_bandwidth", flattenSwitchControllerTrafficPolicyGuaranteedBandwidth(o["guaranteed-bandwidth"], d, "guaranteed_bandwidth")); err != nil {
		if !fortiAPIPatch(o["guaranteed-bandwidth"]) {
			return fmt.Errorf("Error reading guaranteed_bandwidth: %v", err)
		}
	}

	if err = d.Set("guaranteed_burst", flattenSwitchControllerTrafficPolicyGuaranteedBurst(o["guaranteed-burst"], d, "guaranteed_burst")); err != nil {
		if !fortiAPIPatch(o["guaranteed-burst"]) {
			return fmt.Errorf("Error reading guaranteed_burst: %v", err)
		}
	}

	if err = d.Set("maximum_burst", flattenSwitchControllerTrafficPolicyMaximumBurst(o["maximum-burst"], d, "maximum_burst")); err != nil {
		if !fortiAPIPatch(o["maximum-burst"]) {
			return fmt.Errorf("Error reading maximum_burst: %v", err)
		}
	}

	if err = d.Set("type", flattenSwitchControllerTrafficPolicyType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("cos", flattenSwitchControllerTrafficPolicyCos(o["cos"], d, "cos")); err != nil {
		if !fortiAPIPatch(o["cos"]) {
			return fmt.Errorf("Error reading cos: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSwitchControllerTrafficPolicyId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerTrafficPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerTrafficPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyPolicerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyGuaranteedBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyGuaranteedBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyMaximumBurst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyCos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerTrafficPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerTrafficPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerTrafficPolicyDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("policer_status"); ok {
		t, err := expandSwitchControllerTrafficPolicyPolicerStatus(d, v, "policer_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policer-status"] = t
		}
	}

	if v, ok := d.GetOkExists("guaranteed_bandwidth"); ok {
		t, err := expandSwitchControllerTrafficPolicyGuaranteedBandwidth(d, v, "guaranteed_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-bandwidth"] = t
		}
	}

	if v, ok := d.GetOkExists("guaranteed_burst"); ok {
		t, err := expandSwitchControllerTrafficPolicyGuaranteedBurst(d, v, "guaranteed_burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guaranteed-burst"] = t
		}
	}

	if v, ok := d.GetOkExists("maximum_burst"); ok {
		t, err := expandSwitchControllerTrafficPolicyMaximumBurst(d, v, "maximum_burst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-burst"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSwitchControllerTrafficPolicyType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOkExists("cos"); ok {
		t, err := expandSwitchControllerTrafficPolicyCos(d, v, "cos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSwitchControllerTrafficPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
