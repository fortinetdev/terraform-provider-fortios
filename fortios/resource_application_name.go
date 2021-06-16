// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure application signatures.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceApplicationName() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationNameCreate,
		Read:   resourceApplicationNameRead,
		Update: resourceApplicationNameUpdate,
		Delete: resourceApplicationNameDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"sub_category": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"popularity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"risk": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"technology": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"parameter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"metaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"valueid": &schema.Schema{
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
		},
	}
}

func resourceApplicationNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectApplicationName(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationName resource while getting object: %v", err)
	}

	o, err := c.CreateApplicationName(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ApplicationName resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationName")
	}

	return resourceApplicationNameRead(d, m)
}

func resourceApplicationNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectApplicationName(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationName resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationName(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationName")
	}

	return resourceApplicationNameRead(d, m)
}

func resourceApplicationNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteApplicationName(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationNameRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadApplicationName(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationName(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationName resource from API: %v", err)
	}
	return nil
}

func flattenApplicationNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameSubCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNamePopularity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameRisk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameTechnology(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameBehavior(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameParameters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenApplicationNameParametersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenApplicationNameParametersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameParameter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameMetadata(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenApplicationNameMetadataId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metaid"
		if _, ok := i["metaid"]; ok {

			tmp["metaid"] = flattenApplicationNameMetadataMetaid(i["metaid"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valueid"
		if _, ok := i["valueid"]; ok {

			tmp["valueid"] = flattenApplicationNameMetadataValueid(i["valueid"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationNameMetadataId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameMetadataMetaid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationNameMetadataValueid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectApplicationName(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenApplicationNameName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenApplicationNameId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationNameCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("sub_category", flattenApplicationNameSubCategory(o["sub-category"], d, "sub_category", sv)); err != nil {
		if !fortiAPIPatch(o["sub-category"]) {
			return fmt.Errorf("Error reading sub_category: %v", err)
		}
	}

	if err = d.Set("popularity", flattenApplicationNamePopularity(o["popularity"], d, "popularity", sv)); err != nil {
		if !fortiAPIPatch(o["popularity"]) {
			return fmt.Errorf("Error reading popularity: %v", err)
		}
	}

	if err = d.Set("risk", flattenApplicationNameRisk(o["risk"], d, "risk", sv)); err != nil {
		if !fortiAPIPatch(o["risk"]) {
			return fmt.Errorf("Error reading risk: %v", err)
		}
	}

	if err = d.Set("weight", flattenApplicationNameWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("protocol", flattenApplicationNameProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("technology", flattenApplicationNameTechnology(o["technology"], d, "technology", sv)); err != nil {
		if !fortiAPIPatch(o["technology"]) {
			return fmt.Errorf("Error reading technology: %v", err)
		}
	}

	if err = d.Set("behavior", flattenApplicationNameBehavior(o["behavior"], d, "behavior", sv)); err != nil {
		if !fortiAPIPatch(o["behavior"]) {
			return fmt.Errorf("Error reading behavior: %v", err)
		}
	}

	if err = d.Set("vendor", flattenApplicationNameVendor(o["vendor"], d, "vendor", sv)); err != nil {
		if !fortiAPIPatch(o["vendor"]) {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("parameters", flattenApplicationNameParameters(o["parameters"], d, "parameters", sv)); err != nil {
			if !fortiAPIPatch(o["parameters"]) {
				return fmt.Errorf("Error reading parameters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameters"); ok {
			if err = d.Set("parameters", flattenApplicationNameParameters(o["parameters"], d, "parameters", sv)); err != nil {
				if !fortiAPIPatch(o["parameters"]) {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			}
		}
	}

	if err = d.Set("parameter", flattenApplicationNameParameter(o["parameter"], d, "parameter", sv)); err != nil {
		if !fortiAPIPatch(o["parameter"]) {
			return fmt.Errorf("Error reading parameter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("metadata", flattenApplicationNameMetadata(o["metadata"], d, "metadata", sv)); err != nil {
			if !fortiAPIPatch(o["metadata"]) {
				return fmt.Errorf("Error reading metadata: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("metadata"); ok {
			if err = d.Set("metadata", flattenApplicationNameMetadata(o["metadata"], d, "metadata", sv)); err != nil {
				if !fortiAPIPatch(o["metadata"]) {
					return fmt.Errorf("Error reading metadata: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenApplicationNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandApplicationNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameSubCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNamePopularity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameTechnology(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameBehavior(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandApplicationNameParametersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationNameParametersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameParameter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameMetadata(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandApplicationNameMetadataId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metaid"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metaid"], _ = expandApplicationNameMetadataMetaid(d, i["metaid"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valueid"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["valueid"], _ = expandApplicationNameMetadataValueid(d, i["valueid"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationNameMetadataId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameMetadataMetaid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationNameMetadataValueid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationName(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandApplicationNameName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandApplicationNameId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOkExists("category"); ok {

		t, err := expandApplicationNameCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOkExists("sub_category"); ok {

		t, err := expandApplicationNameSubCategory(d, v, "sub_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-category"] = t
		}
	}

	if v, ok := d.GetOkExists("popularity"); ok {

		t, err := expandApplicationNamePopularity(d, v, "popularity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["popularity"] = t
		}
	}

	if v, ok := d.GetOkExists("risk"); ok {

		t, err := expandApplicationNameRisk(d, v, "risk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk"] = t
		}
	}

	if v, ok := d.GetOkExists("weight"); ok {

		t, err := expandApplicationNameWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {

		t, err := expandApplicationNameProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("technology"); ok {

		t, err := expandApplicationNameTechnology(d, v, "technology", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["technology"] = t
		}
	}

	if v, ok := d.GetOk("behavior"); ok {

		t, err := expandApplicationNameBehavior(d, v, "behavior", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["behavior"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok {

		t, err := expandApplicationNameVendor(d, v, "vendor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	if v, ok := d.GetOk("parameters"); ok {

		t, err := expandApplicationNameParameters(d, v, "parameters", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameters"] = t
		}
	}

	if v, ok := d.GetOk("parameter"); ok {

		t, err := expandApplicationNameParameter(d, v, "parameter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter"] = t
		}
	}

	if v, ok := d.GetOk("metadata"); ok {

		t, err := expandApplicationNameMetadata(d, v, "metadata", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metadata"] = t
		}
	}

	return &obj, nil
}
