// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS egress queue policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerQosQueuePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosQueuePolicyCreate,
		Read:   resourceSwitchControllerQosQueuePolicyRead,
		Update: resourceSwitchControllerQosQueuePolicyUpdate,
		Delete: resourceSwitchControllerQosQueuePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required: true,
			},
			"schedule": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"rate_by": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"cos_queue": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"min_rate": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"max_rate": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"min_rate_percent": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"max_rate_percent": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"drop_policy": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerQosQueuePolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerQosQueuePolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerQosQueuePolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQueuePolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosQueuePolicy")
	}

	return resourceSwitchControllerQosQueuePolicyRead(d, m)
}

func resourceSwitchControllerQosQueuePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerQosQueuePolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQosQueuePolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQueuePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosQueuePolicy")
	}

	return resourceSwitchControllerQosQueuePolicyRead(d, m)
}

func resourceSwitchControllerQosQueuePolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerQosQueuePolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosQueuePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosQueuePolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerQosQueuePolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosQueuePolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy resource from API: %v", err)
	}
	return nil
}


func flattenSwitchControllerQosQueuePolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyRateBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueue(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerQosQueuePolicyCosQueueName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = flattenSwitchControllerQosQueuePolicyCosQueueDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := i["min-rate"]; ok {
			tmp["min_rate"] = flattenSwitchControllerQosQueuePolicyCosQueueMinRate(i["min-rate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := i["max-rate"]; ok {
			tmp["max_rate"] = flattenSwitchControllerQosQueuePolicyCosQueueMaxRate(i["max-rate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := i["min-rate-percent"]; ok {
			tmp["min_rate_percent"] = flattenSwitchControllerQosQueuePolicyCosQueueMinRatePercent(i["min-rate-percent"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := i["max-rate-percent"]; ok {
			tmp["max_rate_percent"] = flattenSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(i["max-rate-percent"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := i["drop-policy"]; ok {
			tmp["drop_policy"] = flattenSwitchControllerQosQueuePolicyCosQueueDropPolicy(i["drop-policy"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			tmp["weight"] = flattenSwitchControllerQosQueuePolicyCosQueueWeight(i["weight"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerQosQueuePolicyCosQueueName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMinRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMaxRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMinRatePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueDropPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenSwitchControllerQosQueuePolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSwitchControllerQosQueuePolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("rate_by", flattenSwitchControllerQosQueuePolicyRateBy(o["rate-by"], d, "rate_by")); err != nil {
		if !fortiAPIPatch(o["rate-by"]) {
			return fmt.Errorf("Error reading rate_by: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("cos_queue", flattenSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
            if !fortiAPIPatch(o["cos-queue"]) {
                return fmt.Errorf("Error reading cos_queue: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("cos_queue"); ok {
            if err = d.Set("cos_queue", flattenSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue")); err != nil {
                if !fortiAPIPatch(o["cos-queue"]) {
                    return fmt.Errorf("Error reading cos_queue: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSwitchControllerQosQueuePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSwitchControllerQosQueuePolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyRateBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerQosQueuePolicyCosQueueName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSwitchControllerQosQueuePolicyCosQueueDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-rate"], _ = expandSwitchControllerQosQueuePolicyCosQueueMinRate(d, i["min_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-rate"], _ = expandSwitchControllerQosQueuePolicyCosQueueMaxRate(d, i["max_rate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-rate-percent"], _ = expandSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d, i["min_rate_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-rate-percent"], _ = expandSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d, i["max_rate_percent"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["drop-policy"], _ = expandSwitchControllerQosQueuePolicyCosQueueDropPolicy(d, i["drop_policy"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSwitchControllerQosQueuePolicyCosQueueWeight(d, i["weight"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMinRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMaxRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueDropPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerQosQueuePolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandSwitchControllerQosQueuePolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("rate_by"); ok {
		t, err := expandSwitchControllerQosQueuePolicyRateBy(d, v, "rate_by")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-by"] = t
		}
	}

	if v, ok := d.GetOk("cos_queue"); ok {
		t, err := expandSwitchControllerQosQueuePolicyCosQueue(d, v, "cos_queue")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}


	return &obj, nil
}

