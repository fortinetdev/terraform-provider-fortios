// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Report dataset configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceReportDataset() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportDatasetCreate,
		Read:   resourceReportDatasetRead,
		Update: resourceReportDatasetUpdate,
		Delete: resourceReportDatasetDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				ForceNew:     true,
				Required:     true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"query": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"field": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"displayname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"field": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"data_type": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceReportDatasetCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectReportDataset(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportDataset resource while getting object: %v", err)
	}

	o, err := c.CreateReportDataset(obj)

	if err != nil {
		return fmt.Errorf("Error creating ReportDataset resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportDataset")
	}

	return resourceReportDatasetRead(d, m)
}

func resourceReportDatasetUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectReportDataset(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportDataset resource while getting object: %v", err)
	}

	o, err := c.UpdateReportDataset(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ReportDataset resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportDataset")
	}

	return resourceReportDatasetRead(d, m)
}

func resourceReportDatasetDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteReportDataset(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ReportDataset resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportDatasetRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadReportDataset(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ReportDataset resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportDataset(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportDataset resource from API: %v", err)
	}
	return nil
}

func flattenReportDatasetName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetField(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenReportDatasetFieldId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenReportDatasetFieldType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenReportDatasetFieldName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "displayname"
		if _, ok := i["displayname"]; ok {
			tmp["displayname"] = flattenReportDatasetFieldDisplayname(i["displayname"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportDatasetFieldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetFieldType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetFieldName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetFieldDisplayname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetParameters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenReportDatasetParametersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "display_name"
		if _, ok := i["display-name"]; ok {
			tmp["display_name"] = flattenReportDatasetParametersDisplayName(i["display-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
		if _, ok := i["field"]; ok {
			tmp["field"] = flattenReportDatasetParametersField(i["field"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data_type"
		if _, ok := i["data-type"]; ok {
			tmp["data_type"] = flattenReportDatasetParametersDataType(i["data-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportDatasetParametersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetParametersDisplayName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetParametersField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportDatasetParametersDataType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportDataset(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenReportDatasetName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policy", flattenReportDatasetPolicy(o["policy"], d, "policy")); err != nil {
		if !fortiAPIPatch(o["policy"]) {
			return fmt.Errorf("Error reading policy: %v", err)
		}
	}

	if err = d.Set("query", flattenReportDatasetQuery(o["query"], d, "query")); err != nil {
		if !fortiAPIPatch(o["query"]) {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("field", flattenReportDatasetField(o["field"], d, "field")); err != nil {
			if !fortiAPIPatch(o["field"]) {
				return fmt.Errorf("Error reading field: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("field"); ok {
			if err = d.Set("field", flattenReportDatasetField(o["field"], d, "field")); err != nil {
				if !fortiAPIPatch(o["field"]) {
					return fmt.Errorf("Error reading field: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("parameters", flattenReportDatasetParameters(o["parameters"], d, "parameters")); err != nil {
			if !fortiAPIPatch(o["parameters"]) {
				return fmt.Errorf("Error reading parameters: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameters"); ok {
			if err = d.Set("parameters", flattenReportDatasetParameters(o["parameters"], d, "parameters")); err != nil {
				if !fortiAPIPatch(o["parameters"]) {
					return fmt.Errorf("Error reading parameters: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenReportDatasetFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportDatasetName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetQuery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandReportDatasetFieldId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandReportDatasetFieldType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandReportDatasetFieldName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "displayname"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["displayname"], _ = expandReportDatasetFieldDisplayname(d, i["displayname"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportDatasetFieldId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetFieldType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetFieldName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetFieldDisplayname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandReportDatasetParametersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "display_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["display-name"], _ = expandReportDatasetParametersDisplayName(d, i["display_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["field"], _ = expandReportDatasetParametersField(d, i["field"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "data_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["data-type"], _ = expandReportDatasetParametersDataType(d, i["data_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportDatasetParametersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetParametersDisplayName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetParametersField(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportDatasetParametersDataType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportDataset(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandReportDatasetName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("policy"); ok {
		t, err := expandReportDatasetPolicy(d, v, "policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok {
		t, err := expandReportDatasetQuery(d, v, "query")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("field"); ok {
		t, err := expandReportDatasetField(d, v, "field")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["field"] = t
		}
	}

	if v, ok := d.GetOk("parameters"); ok {
		t, err := expandReportDatasetParameters(d, v, "parameters")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameters"] = t
		}
	}

	return &obj, nil
}
