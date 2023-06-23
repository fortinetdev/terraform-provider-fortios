// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SSL-VPN host check software.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebHostCheckSoftware() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebHostCheckSoftwareCreate,
		Read:   resourceVpnSslWebHostCheckSoftwareRead,
		Update: resourceVpnSslWebHostCheckSoftwareUpdate,
		Delete: resourceVpnSslWebHostCheckSoftwareDelete,

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
			"os_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_item_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"version": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"md5s": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 32),
										Optional:     true,
										Computed:     true,
									},
								},
							},
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

func resourceVpnSslWebHostCheckSoftwareCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslWebHostCheckSoftware(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebHostCheckSoftware resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebHostCheckSoftware(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebHostCheckSoftware resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebHostCheckSoftware")
	}

	return resourceVpnSslWebHostCheckSoftwareRead(d, m)
}

func resourceVpnSslWebHostCheckSoftwareUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslWebHostCheckSoftware(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebHostCheckSoftware resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebHostCheckSoftware(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebHostCheckSoftware resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebHostCheckSoftware")
	}

	return resourceVpnSslWebHostCheckSoftwareRead(d, m)
}

func resourceVpnSslWebHostCheckSoftwareDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnSslWebHostCheckSoftware(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebHostCheckSoftware resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebHostCheckSoftwareRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnSslWebHostCheckSoftware(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebHostCheckSoftware resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebHostCheckSoftware(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebHostCheckSoftware resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebHostCheckSoftwareName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareOsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareGuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareCheckItemList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenVpnSslWebHostCheckSoftwareCheckItemListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenVpnSslWebHostCheckSoftwareCheckItemListAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenVpnSslWebHostCheckSoftwareCheckItemListType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			tmp["target"] = flattenVpnSslWebHostCheckSoftwareCheckItemListTarget(i["target"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			tmp["version"] = flattenVpnSslWebHostCheckSoftwareCheckItemListVersion(i["version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5s"
		if _, ok := i["md5s"]; ok {
			tmp["md5s"] = flattenVpnSslWebHostCheckSoftwareCheckItemListMd5S(i["md5s"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListMd5S(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenVpnSslWebHostCheckSoftwareCheckItemListMd5SId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnSslWebHostCheckSoftwareCheckItemListMd5SId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnSslWebHostCheckSoftware(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenVpnSslWebHostCheckSoftwareName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("os_type", flattenVpnSslWebHostCheckSoftwareOsType(o["os-type"], d, "os_type", sv)); err != nil {
		if !fortiAPIPatch(o["os-type"]) {
			return fmt.Errorf("Error reading os_type: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnSslWebHostCheckSoftwareType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("version", flattenVpnSslWebHostCheckSoftwareVersion(o["version"], d, "version", sv)); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("guid", flattenVpnSslWebHostCheckSoftwareGuid(o["guid"], d, "guid", sv)); err != nil {
		if !fortiAPIPatch(o["guid"]) {
			return fmt.Errorf("Error reading guid: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("check_item_list", flattenVpnSslWebHostCheckSoftwareCheckItemList(o["check-item-list"], d, "check_item_list", sv)); err != nil {
			if !fortiAPIPatch(o["check-item-list"]) {
				return fmt.Errorf("Error reading check_item_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("check_item_list"); ok {
			if err = d.Set("check_item_list", flattenVpnSslWebHostCheckSoftwareCheckItemList(o["check-item-list"], d, "check_item_list", sv)); err != nil {
				if !fortiAPIPatch(o["check-item-list"]) {
					return fmt.Errorf("Error reading check_item_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnSslWebHostCheckSoftwareFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnSslWebHostCheckSoftwareName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareOsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareGuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListTarget(d, i["target"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["version"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListVersion(d, i["version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5s"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["md5s"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListMd5S(d, i["md5s"], pre_append, sv)
		} else {
			tmp["md5s"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListMd5S(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnSslWebHostCheckSoftwareCheckItemListMd5SId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebHostCheckSoftwareCheckItemListMd5SId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebHostCheckSoftware(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnSslWebHostCheckSoftwareName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_type"); ok {
		t, err := expandVpnSslWebHostCheckSoftwareOsType(d, v, "os_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-type"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandVpnSslWebHostCheckSoftwareType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok {
		t, err := expandVpnSslWebHostCheckSoftwareVersion(d, v, "version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("guid"); ok {
		t, err := expandVpnSslWebHostCheckSoftwareGuid(d, v, "guid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guid"] = t
		}
	}

	if v, ok := d.GetOk("check_item_list"); ok || d.HasChange("check_item_list") {
		t, err := expandVpnSslWebHostCheckSoftwareCheckItemList(d, v, "check_item_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-item-list"] = t
		}
	}

	return &obj, nil
}
