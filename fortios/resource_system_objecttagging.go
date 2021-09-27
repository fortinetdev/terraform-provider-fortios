// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure object tagging.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"category": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multiple": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemObjectTaggingCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemObjectTagging(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemObjectTagging resource while getting object: %v", err)
	}

	o, err := c.CreateSystemObjectTagging(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemObjectTagging(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemObjectTagging resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemObjectTagging(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemObjectTagging(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemObjectTagging(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTagging resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemObjectTagging(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemObjectTagging resource from API: %v", err)
	}
	return nil
}

func flattenSystemObjectTaggingCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemObjectTaggingAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemObjectTaggingDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemObjectTaggingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemObjectTaggingMultiple(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemObjectTaggingColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemObjectTaggingTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemObjectTaggingTagsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemObjectTaggingTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemObjectTagging(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("category", flattenSystemObjectTaggingCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemObjectTaggingAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemObjectTaggingDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemObjectTaggingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("multiple", flattenSystemObjectTaggingMultiple(o["multiple"], d, "multiple", sv)); err != nil {
		if !fortiAPIPatch(o["multiple"]) {
			return fmt.Errorf("Error reading multiple: %v", err)
		}
	}

	if err = d.Set("color", flattenSystemObjectTaggingColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tags", flattenSystemObjectTaggingTags(o["tags"], d, "tags", sv)); err != nil {
			if !fortiAPIPatch(o["tags"]) {
				return fmt.Errorf("Error reading tags: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tags"); ok {
			if err = d.Set("tags", flattenSystemObjectTaggingTags(o["tags"], d, "tags", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemObjectTaggingCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingMultiple(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemObjectTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemObjectTaggingTagsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemObjectTaggingTagsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemObjectTagging(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("category"); ok {

		t, err := expandSystemObjectTaggingCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {

		t, err := expandSystemObjectTaggingAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {

		t, err := expandSystemObjectTaggingDevice(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSystemObjectTaggingInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("multiple"); ok {

		t, err := expandSystemObjectTaggingMultiple(d, v, "multiple", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multiple"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandSystemObjectTaggingColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("tags"); ok {

		t, err := expandSystemObjectTaggingTags(d, v, "tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tags"] = t
		}
	}

	return &obj, nil
}
