// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch QoS egress queue policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Required:     true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rate_by": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cos_queue": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"min_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_rate_percent": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"drop_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ecn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerQosQueuePolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSwitchControllerQosQueuePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerQosQueuePolicy(obj, vdomparam)

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

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSwitchControllerQosQueuePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosQueuePolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQosQueuePolicy(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerQosQueuePolicy(mkey, vdomparam)
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

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSwitchControllerQosQueuePolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosQueuePolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosQueuePolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerQosQueuePolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicySchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyRateBy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueue(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSwitchControllerQosQueuePolicyCosQueueName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["description"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
			}
			tmp["description"] = flattenSwitchControllerQosQueuePolicyCosQueueDescription(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["min-rate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
			}
			tmp["min_rate"] = flattenSwitchControllerQosQueuePolicyCosQueueMinRate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["max-rate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
			}
			tmp["max_rate"] = flattenSwitchControllerQosQueuePolicyCosQueueMaxRate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["min-rate-percent"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
			}
			tmp["min_rate_percent"] = flattenSwitchControllerQosQueuePolicyCosQueueMinRatePercent(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["max-rate-percent"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
			}
			tmp["max_rate_percent"] = flattenSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["drop-policy"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
			}
			tmp["drop_policy"] = flattenSwitchControllerQosQueuePolicyCosQueueDropPolicy(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ecn"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
			}
			tmp["ecn"] = flattenSwitchControllerQosQueuePolicyCosQueueEcn(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenSwitchControllerQosQueuePolicyCosQueueWeight(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerQosQueuePolicyCosQueueName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueMinRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSwitchControllerQosQueuePolicyCosQueueMaxRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSwitchControllerQosQueuePolicyCosQueueMinRatePercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSwitchControllerQosQueuePolicyCosQueueDropPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueEcn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerQosQueuePolicyCosQueueWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSwitchControllerQosQueuePolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSwitchControllerQosQueuePolicySchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("rate_by", flattenSwitchControllerQosQueuePolicyRateBy(o["rate-by"], d, "rate_by", sv)); err != nil {
		if !fortiAPIPatch(o["rate-by"]) {
			return fmt.Errorf("Error reading rate_by: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("cos_queue", flattenSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue", sv)); err != nil {
			if !fortiAPIPatch(o["cos-queue"]) {
				return fmt.Errorf("Error reading cos_queue: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cos_queue"); ok {
			if err = d.Set("cos_queue", flattenSwitchControllerQosQueuePolicyCosQueue(o["cos-queue"], d, "cos_queue", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerQosQueuePolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicySchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyRateBy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerQosQueuePolicyCosQueueName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSwitchControllerQosQueuePolicyCosQueueDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-rate"], _ = expandSwitchControllerQosQueuePolicyCosQueueMinRate(d, i["min_rate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["min-rate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-rate"], _ = expandSwitchControllerQosQueuePolicyCosQueueMaxRate(d, i["max_rate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["max-rate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-rate-percent"], _ = expandSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d, i["min_rate_percent"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["min-rate-percent"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_rate_percent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-rate-percent"], _ = expandSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d, i["max_rate_percent"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["max-rate-percent"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drop_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["drop-policy"], _ = expandSwitchControllerQosQueuePolicyCosQueueDropPolicy(d, i["drop_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ecn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ecn"], _ = expandSwitchControllerQosQueuePolicyCosQueueEcn(d, i["ecn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSwitchControllerQosQueuePolicyCosQueueWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMinRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMaxRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMinRatePercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueMaxRatePercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueDropPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueEcn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosQueuePolicyCosQueueWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerQosQueuePolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerQosQueuePolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandSwitchControllerQosQueuePolicySchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("rate_by"); ok {
		t, err := expandSwitchControllerQosQueuePolicyRateBy(d, v, "rate_by", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rate-by"] = t
		}
	}

	if v, ok := d.GetOk("cos_queue"); ok || d.HasChange("cos_queue") {
		t, err := expandSwitchControllerQosQueuePolicyCosQueue(d, v, "cos_queue", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cos-queue"] = t
		}
	}

	return &obj, nil
}
