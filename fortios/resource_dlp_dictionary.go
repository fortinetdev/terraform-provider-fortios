// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure dictionaries used by DLP blocking.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpDictionary() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpDictionaryCreate,
		Read:   resourceDlpDictionaryRead,
		Update: resourceDlpDictionaryUpdate,
		Delete: resourceDlpDictionaryDelete,

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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_around": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"entries": &schema.Schema{
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
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"pattern": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"ignore_case": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repeat": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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

func resourceDlpDictionaryCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDictionary(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating DlpDictionary resource while getting object: %v", err)
	}

	o, err := c.CreateDlpDictionary(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating DlpDictionary resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpDictionary")
	}

	return resourceDlpDictionaryRead(d, m)
}

func resourceDlpDictionaryUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDictionary(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DlpDictionary resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpDictionary(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DlpDictionary resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpDictionary")
	}

	return resourceDlpDictionaryRead(d, m)
}

func resourceDlpDictionaryDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteDlpDictionary(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting DlpDictionary resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpDictionaryRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpDictionary(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DlpDictionary resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpDictionary(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DlpDictionary resource from API: %v", err)
	}
	return nil
}

func flattenDlpDictionaryUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryMatchType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryMatchAround(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenDlpDictionaryEntriesId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenDlpDictionaryEntriesType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pattern"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
			}
			tmp["pattern"] = flattenDlpDictionaryEntriesPattern(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ignore-case"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_case"
			}
			tmp["ignore_case"] = flattenDlpDictionaryEntriesIgnoreCase(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["repeat"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "repeat"
			}
			tmp["repeat"] = flattenDlpDictionaryEntriesRepeat(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenDlpDictionaryEntriesStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comment"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
			}
			tmp["comment"] = flattenDlpDictionaryEntriesComment(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenDlpDictionaryEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpDictionaryEntriesType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesIgnoreCase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesRepeat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDictionaryEntriesComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDlpDictionary(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("uuid", flattenDlpDictionaryUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("name", flattenDlpDictionaryName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("match_type", flattenDlpDictionaryMatchType(o["match-type"], d, "match_type", sv)); err != nil {
		if !fortiAPIPatch(o["match-type"]) {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("match_around", flattenDlpDictionaryMatchAround(o["match-around"], d, "match_around", sv)); err != nil {
		if !fortiAPIPatch(o["match-around"]) {
			return fmt.Errorf("Error reading match_around: %v", err)
		}
	}

	if err = d.Set("comment", flattenDlpDictionaryComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("entries", flattenDlpDictionaryEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenDlpDictionaryEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenDlpDictionaryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDlpDictionaryUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryMatchType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryMatchAround(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandDlpDictionaryEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandDlpDictionaryEntriesType(d, i["type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandDlpDictionaryEntriesPattern(d, i["pattern"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["pattern"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_case"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ignore-case"], _ = expandDlpDictionaryEntriesIgnoreCase(d, i["ignore_case"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "repeat"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["repeat"], _ = expandDlpDictionaryEntriesRepeat(d, i["repeat"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandDlpDictionaryEntriesStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandDlpDictionaryEntriesComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpDictionaryEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesIgnoreCase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesRepeat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDictionaryEntriesComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDlpDictionary(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandDlpDictionaryUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpDictionaryName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok {
		t, err := expandDlpDictionaryMatchType(d, v, "match_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-type"] = t
		}
	}

	if v, ok := d.GetOk("match_around"); ok {
		t, err := expandDlpDictionaryMatchAround(d, v, "match_around", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-around"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandDlpDictionaryComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandDlpDictionaryEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	return &obj, nil
}
