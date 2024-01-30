// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure exact-data-match template used by DLP scan.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpExactDataMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpExactDataMatchCreate,
		Read:   resourceDlpExactDataMatchRead,
		Update: resourceDlpExactDataMatchUpdate,
		Delete: resourceDlpExactDataMatchDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"optional": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"data": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"columns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"optional": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceDlpExactDataMatchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDlpExactDataMatch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating DlpExactDataMatch resource while getting object: %v", err)
	}

	o, err := c.CreateDlpExactDataMatch(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating DlpExactDataMatch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpExactDataMatch")
	}

	return resourceDlpExactDataMatchRead(d, m)
}

func resourceDlpExactDataMatchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDlpExactDataMatch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DlpExactDataMatch resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpExactDataMatch(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DlpExactDataMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpExactDataMatch")
	}

	return resourceDlpExactDataMatchRead(d, m)
}

func resourceDlpExactDataMatchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteDlpExactDataMatch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting DlpExactDataMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpExactDataMatchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadDlpExactDataMatch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DlpExactDataMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpExactDataMatch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DlpExactDataMatch resource from API: %v", err)
	}
	return nil
}

func flattenDlpExactDataMatchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpExactDataMatchOptional(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpExactDataMatchData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumns(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if cur_v, ok := i["index"]; ok {
			tmp["index"] = flattenDlpExactDataMatchColumnsIndex(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenDlpExactDataMatchColumnsType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if cur_v, ok := i["optional"]; ok {
			tmp["optional"] = flattenDlpExactDataMatchColumnsOptional(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "index", d)
	return result
}

func flattenDlpExactDataMatchColumnsIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumnsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpExactDataMatchColumnsOptional(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDlpExactDataMatch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenDlpExactDataMatchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("optional", flattenDlpExactDataMatchOptional(o["optional"], d, "optional", sv)); err != nil {
		if !fortiAPIPatch(o["optional"]) {
			return fmt.Errorf("Error reading optional: %v", err)
		}
	}

	if err = d.Set("data", flattenDlpExactDataMatchData(o["data"], d, "data", sv)); err != nil {
		if !fortiAPIPatch(o["data"]) {
			return fmt.Errorf("Error reading data: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("columns", flattenDlpExactDataMatchColumns(o["columns"], d, "columns", sv)); err != nil {
			if !fortiAPIPatch(o["columns"]) {
				return fmt.Errorf("Error reading columns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("columns"); ok {
			if err = d.Set("columns", flattenDlpExactDataMatchColumns(o["columns"], d, "columns", sv)); err != nil {
				if !fortiAPIPatch(o["columns"]) {
					return fmt.Errorf("Error reading columns: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenDlpExactDataMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDlpExactDataMatchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchOptional(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandDlpExactDataMatchColumnsIndex(d, i["index"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandDlpExactDataMatchColumnsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "optional"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["optional"], _ = expandDlpExactDataMatchColumnsOptional(d, i["optional"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpExactDataMatchColumnsIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumnsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpExactDataMatchColumnsOptional(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDlpExactDataMatch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpExactDataMatchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("optional"); ok {
		t, err := expandDlpExactDataMatchOptional(d, v, "optional", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional"] = t
		}
	}

	if v, ok := d.GetOk("data"); ok {
		t, err := expandDlpExactDataMatchData(d, v, "data", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data"] = t
		}
	}

	if v, ok := d.GetOk("columns"); ok || d.HasChange("columns") {
		t, err := expandDlpExactDataMatchColumns(d, v, "columns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["columns"] = t
		}
	}

	return &obj, nil
}
