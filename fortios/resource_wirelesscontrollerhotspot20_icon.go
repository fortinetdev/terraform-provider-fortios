// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure OSU provider icon.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20Icon() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20IconCreate,
		Read:   resourceWirelessControllerHotspot20IconRead,
		Update: resourceWirelessControllerHotspot20IconUpdate,
		Delete: resourceWirelessControllerHotspot20IconDelete,

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
			"icon_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"lang": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional:     true,
							Computed:     true,
						},
						"file": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"width": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"height": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
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

func resourceWirelessControllerHotspot20IconCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20Icon(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Icon resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20Icon(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Icon resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20Icon")
	}

	return resourceWirelessControllerHotspot20IconRead(d, m)
}

func resourceWirelessControllerHotspot20IconUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20Icon(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Icon resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20Icon(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Icon resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20Icon")
	}

	return resourceWirelessControllerHotspot20IconRead(d, m)
}

func resourceWirelessControllerHotspot20IconDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20Icon(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20Icon resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20IconRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerHotspot20Icon(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Icon resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20Icon(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Icon resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20IconName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenWirelessControllerHotspot20IconIconListName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {

			tmp["lang"] = flattenWirelessControllerHotspot20IconIconListLang(i["lang"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file"
		if _, ok := i["file"]; ok {

			tmp["file"] = flattenWirelessControllerHotspot20IconIconListFile(i["file"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenWirelessControllerHotspot20IconIconListType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := i["width"]; ok {

			tmp["width"] = flattenWirelessControllerHotspot20IconIconListWidth(i["width"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := i["height"]; ok {

			tmp["height"] = flattenWirelessControllerHotspot20IconIconListHeight(i["height"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerHotspot20IconIconListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListLang(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20IconIconListHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20Icon(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20IconName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("icon_list", flattenWirelessControllerHotspot20IconIconList(o["icon-list"], d, "icon_list", sv)); err != nil {
			if !fortiAPIPatch(o["icon-list"]) {
				return fmt.Errorf("Error reading icon_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icon_list"); ok {
			if err = d.Set("icon_list", flattenWirelessControllerHotspot20IconIconList(o["icon-list"], d, "icon_list", sv)); err != nil {
				if !fortiAPIPatch(o["icon-list"]) {
					return fmt.Errorf("Error reading icon_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20IconFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20IconName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerHotspot20IconIconListName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["lang"], _ = expandWirelessControllerHotspot20IconIconListLang(d, i["lang"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["file"], _ = expandWirelessControllerHotspot20IconIconListFile(d, i["file"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandWirelessControllerHotspot20IconIconListType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["width"], _ = expandWirelessControllerHotspot20IconIconListWidth(d, i["width"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["height"], _ = expandWirelessControllerHotspot20IconIconListHeight(d, i["height"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20IconIconListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListLang(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20IconIconListHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20Icon(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerHotspot20IconName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("icon_list"); ok {

		t, err := expandWirelessControllerHotspot20IconIconList(d, v, "icon_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-list"] = t
		}
	}

	return &obj, nil
}
