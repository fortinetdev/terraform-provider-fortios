// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Designate logical storage for DLP fingerprint database.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSettingsUpdate,
		Read:   resourceDlpSettingsRead,
		Update: resourceDlpSettingsUpdate,
		Delete: resourceDlpSettingsDelete,

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
			"storage_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_mem_percent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"chunk_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 100000),
				Optional:     true,
				Computed:     true,
			},
			"config_builder_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 100000),
				Optional:     true,
				Computed:     true,
			},
			"ocr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"confidence": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"max_file_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1427456),
							Optional:     true,
						},
						"filetype_ignore_list": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 39),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceDlpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpSettings")
	}

	return resourceDlpSettingsRead(d, m)
}

func resourceDlpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDlpSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateDlpSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing DlpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSettingsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DlpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DlpSettings resource from API: %v", err)
	}
	return nil
}

func flattenDlpSettingsStorageDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSettingsSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSettingsDbMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSettingsCacheMemPercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSettingsChunkSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSettingsConfigBuilderTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSettingsOcr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "scan"
	if _, ok := i["scan"]; ok {
		result["scan"] = flattenDlpSettingsOcrScan(i["scan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "confidence"
	if _, ok := i["confidence"]; ok {
		result["confidence"] = flattenDlpSettingsOcrConfidence(i["confidence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_file_size"
	if _, ok := i["max-file-size"]; ok {
		result["max_file_size"] = flattenDlpSettingsOcrMaxFileSize(i["max-file-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "filetype_ignore_list"
	if _, ok := i["filetype-ignore-list"]; ok {
		result["filetype_ignore_list"] = flattenDlpSettingsOcrFiletypeIgnoreList(i["filetype-ignore-list"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDlpSettingsOcrScan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSettingsOcrConfidence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSettingsOcrMaxFileSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSettingsOcrFiletypeIgnoreList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenDlpSettingsOcrFiletypeIgnoreListName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenDlpSettingsOcrFiletypeIgnoreListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDlpSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("storage_device", flattenDlpSettingsStorageDevice(o["storage-device"], d, "storage_device", sv)); err != nil {
		if !fortiAPIPatch(o["storage-device"]) {
			return fmt.Errorf("Error reading storage_device: %v", err)
		}
	}

	if err = d.Set("size", flattenDlpSettingsSize(o["size"], d, "size", sv)); err != nil {
		if !fortiAPIPatch(o["size"]) {
			return fmt.Errorf("Error reading size: %v", err)
		}
	}

	if err = d.Set("db_mode", flattenDlpSettingsDbMode(o["db-mode"], d, "db_mode", sv)); err != nil {
		if !fortiAPIPatch(o["db-mode"]) {
			return fmt.Errorf("Error reading db_mode: %v", err)
		}
	}

	if err = d.Set("cache_mem_percent", flattenDlpSettingsCacheMemPercent(o["cache-mem-percent"], d, "cache_mem_percent", sv)); err != nil {
		if !fortiAPIPatch(o["cache-mem-percent"]) {
			return fmt.Errorf("Error reading cache_mem_percent: %v", err)
		}
	}

	if err = d.Set("chunk_size", flattenDlpSettingsChunkSize(o["chunk-size"], d, "chunk_size", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size"]) {
			return fmt.Errorf("Error reading chunk_size: %v", err)
		}
	}

	if err = d.Set("config_builder_timeout", flattenDlpSettingsConfigBuilderTimeout(o["config-builder-timeout"], d, "config_builder_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["config-builder-timeout"]) {
			return fmt.Errorf("Error reading config_builder_timeout: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ocr", flattenDlpSettingsOcr(o["ocr"], d, "ocr", sv)); err != nil {
			if !fortiAPIPatch(o["ocr"]) {
				return fmt.Errorf("Error reading ocr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ocr"); ok {
			if err = d.Set("ocr", flattenDlpSettingsOcr(o["ocr"], d, "ocr", sv)); err != nil {
				if !fortiAPIPatch(o["ocr"]) {
					return fmt.Errorf("Error reading ocr: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenDlpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDlpSettingsStorageDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsDbMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsCacheMemPercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsChunkSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsConfigBuilderTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcr(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "scan"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["scan"] = nil
		} else {
			result["scan"], _ = expandDlpSettingsOcrScan(d, i["scan"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "confidence"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["confidence"] = nil
		} else {
			result["confidence"], _ = expandDlpSettingsOcrConfidence(d, i["confidence"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "max_file_size"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["max-file-size"] = nil
		} else {
			result["max-file-size"], _ = expandDlpSettingsOcrMaxFileSize(d, i["max_file_size"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["max-file-size"] = nil
	}
	pre_append = pre + ".0." + "filetype_ignore_list"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["filetype-ignore-list"] = make([]struct{}, 0)
		} else {
			result["filetype-ignore-list"], _ = expandDlpSettingsOcrFiletypeIgnoreList(d, i["filetype_ignore_list"], pre_append, sv)
		}
	} else {
		result["filetype-ignore-list"] = make([]string, 0)
	}

	return result, nil
}

func expandDlpSettingsOcrScan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrConfidence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrMaxFileSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsOcrFiletypeIgnoreList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandDlpSettingsOcrFiletypeIgnoreListName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSettingsOcrFiletypeIgnoreListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("storage_device"); ok {
		if setArgNil {
			obj["storage-device"] = nil
		} else {
			t, err := expandDlpSettingsStorageDevice(d, v, "storage_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["storage-device"] = t
			}
		}
	} else if d.HasChange("storage_device") {
		obj["storage-device"] = nil
	}

	if v, ok := d.GetOk("size"); ok {
		if setArgNil {
			obj["size"] = nil
		} else {
			t, err := expandDlpSettingsSize(d, v, "size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["size"] = t
			}
		}
	}

	if v, ok := d.GetOk("db_mode"); ok {
		if setArgNil {
			obj["db-mode"] = nil
		} else {
			t, err := expandDlpSettingsDbMode(d, v, "db_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["db-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("cache_mem_percent"); ok {
		if setArgNil {
			obj["cache-mem-percent"] = nil
		} else {
			t, err := expandDlpSettingsCacheMemPercent(d, v, "cache_mem_percent", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cache-mem-percent"] = t
			}
		}
	}

	if v, ok := d.GetOk("chunk_size"); ok {
		if setArgNil {
			obj["chunk-size"] = nil
		} else {
			t, err := expandDlpSettingsChunkSize(d, v, "chunk_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["chunk-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_builder_timeout"); ok {
		if setArgNil {
			obj["config-builder-timeout"] = nil
		} else {
			t, err := expandDlpSettingsConfigBuilderTimeout(d, v, "config_builder_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-builder-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("ocr"); ok {
		t, err := expandDlpSettingsOcr(d, v, "ocr", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocr"] = t
		}
	}

	return &obj, nil
}
