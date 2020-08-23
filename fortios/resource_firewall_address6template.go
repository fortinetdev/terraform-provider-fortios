// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 address templates.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallAddress6Template() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddress6TemplateCreate,
		Read:   resourceFirewallAddress6TemplateRead,
		Update: resourceFirewallAddress6TemplateUpdate,
		Delete: resourceFirewallAddress6TemplateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required: true,
			},
			"ip6": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_segment_count": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 6),
				Required: true,
			},
			"subnet_segment": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"bits": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),
							Optional: true,
							Computed: true,
						},
						"exclusive": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"values": &schema.Schema{
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
									"value": &schema.Schema{
										Type: schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceFirewallAddress6TemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallAddress6Template(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6Template resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallAddress6Template(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress6Template resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAddress6Template")
	}

	return resourceFirewallAddress6TemplateRead(d, m)
}

func resourceFirewallAddress6TemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallAddress6Template(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6Template resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAddress6Template(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress6Template resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAddress6Template")
	}

	return resourceFirewallAddress6TemplateRead(d, m)
}

func resourceFirewallAddress6TemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallAddress6Template(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress6Template resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6TemplateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallAddress6Template(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6Template resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress6Template(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress6Template resource from API: %v", err)
	}
	return nil
}


func flattenFirewallAddress6TemplateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallAddress6TemplateSubnetSegmentId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallAddress6TemplateSubnetSegmentName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := i["bits"]; ok {
			tmp["bits"] = flattenFirewallAddress6TemplateSubnetSegmentBits(i["bits"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := i["exclusive"]; ok {
			tmp["exclusive"] = flattenFirewallAddress6TemplateSubnetSegmentExclusive(i["exclusive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := i["values"]; ok {
			tmp["values"] = flattenFirewallAddress6TemplateSubnetSegmentValues(i["values"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallAddress6TemplateSubnetSegmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentExclusive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValues(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallAddress6TemplateSubnetSegmentValuesName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenFirewallAddress6TemplateSubnetSegmentValuesValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallAddress6TemplateSubnetSegmentValuesValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectFirewallAddress6Template(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenFirewallAddress6TemplateName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ip6", flattenFirewallAddress6TemplateIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("subnet_segment_count", flattenFirewallAddress6TemplateSubnetSegmentCount(o["subnet-segment-count"], d, "subnet_segment_count")); err != nil {
		if !fortiAPIPatch(o["subnet-segment-count"]) {
			return fmt.Errorf("Error reading subnet_segment_count: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("subnet_segment", flattenFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
            if !fortiAPIPatch(o["subnet-segment"]) {
                return fmt.Errorf("Error reading subnet_segment: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("subnet_segment"); ok {
            if err = d.Set("subnet_segment", flattenFirewallAddress6TemplateSubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
                if !fortiAPIPatch(o["subnet-segment"]) {
                    return fmt.Errorf("Error reading subnet_segment: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenFirewallAddress6TemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandFirewallAddress6TemplateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFirewallAddress6TemplateSubnetSegmentId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandFirewallAddress6TemplateSubnetSegmentName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bits"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bits"], _ = expandFirewallAddress6TemplateSubnetSegmentBits(d, i["bits"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["exclusive"], _ = expandFirewallAddress6TemplateSubnetSegmentExclusive(d, i["exclusive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["values"], _ = expandFirewallAddress6TemplateSubnetSegmentValues(d, i["values"], pre_append)
		} else {
			tmp["values"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6TemplateSubnetSegmentId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentExclusive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallAddress6TemplateSubnetSegmentValuesName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandFirewallAddress6TemplateSubnetSegmentValuesValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValuesValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallAddress6Template(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallAddress6TemplateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandFirewallAddress6TemplateIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment_count"); ok {
		t, err := expandFirewallAddress6TemplateSubnetSegmentCount(d, v, "subnet_segment_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment-count"] = t
		}
	}

	if v, ok := d.GetOk("subnet_segment"); ok {
		t, err := expandFirewallAddress6TemplateSubnetSegment(d, v, "subnet_segment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-segment"] = t
		}
	}


	return &obj, nil
}

