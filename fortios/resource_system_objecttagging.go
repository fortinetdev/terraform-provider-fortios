// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure object tagging.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemObjectTagging() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemObjectTaggingCreate,
		Read:   resourceSystemObjectTaggingRead,
		Update: resourceSystemObjectTaggingUpdate,
		Delete: resourceSystemObjectTaggingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multiple": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional: true,
				Computed: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemObjectTaggingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemObjectTagging(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemObjectTagging resource while getting object: %v", err)
	}

	o, err := c.CreateSystemObjectTagging(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemObjectTagging resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemObjectTagging")
	}

	return resourceSystemObjectTaggingRead(d, m)
}

func resourceSystemObjectTaggingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemObjectTagging(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTagging resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemObjectTagging(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTagging resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemObjectTagging")
	}

	return resourceSystemObjectTaggingRead(d, m)
}

func resourceSystemObjectTaggingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemObjectTagging(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemObjectTagging resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemObjectTaggingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemObjectTagging(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTagging resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemObjectTagging(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTagging resource from API: %v", err)
	}
	return nil
}


func flattenSystemObjectTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingMultiple(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemObjectTaggingTags(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemObjectTaggingTagsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemObjectTaggingTagsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemObjectTagging(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("category", flattenSystemObjectTaggingCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemObjectTaggingAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemObjectTaggingDevice(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemObjectTaggingInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("multiple", flattenSystemObjectTaggingMultiple(o["multiple"], d, "multiple")); err != nil {
		if !fortiAPIPatch(o["multiple"]) {
			return fmt.Errorf("Error reading multiple: %v", err)
		}
	}

	if err = d.Set("color", flattenSystemObjectTaggingColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("tags", flattenSystemObjectTaggingTags(o["tags"], d, "tags")); err != nil {
            if !fortiAPIPatch(o["tags"]) {
                return fmt.Errorf("Error reading tags: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("tags"); ok {
            if err = d.Set("tags", flattenSystemObjectTaggingTags(o["tags"], d, "tags")); err != nil {
                if !fortiAPIPatch(o["tags"]) {
                    return fmt.Errorf("Error reading tags: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSystemObjectTaggingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemObjectTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingMultiple(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemObjectTaggingTagsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemObjectTaggingTagsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemObjectTagging(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("category"); ok {
		t, err := expandSystemObjectTaggingCategory(d, v, "category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandSystemObjectTaggingAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandSystemObjectTaggingDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemObjectTaggingInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("multiple"); ok {
		t, err := expandSystemObjectTaggingMultiple(d, v, "multiple")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandSystemObjectTaggingColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		t, err := expandSystemObjectTaggingTags(d, v, "tags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}


	return &obj, nil
}

