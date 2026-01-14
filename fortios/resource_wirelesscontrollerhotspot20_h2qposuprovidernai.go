// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure online sign up (OSU) provider NAI list.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpOsuProviderNai() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpOsuProviderNaiCreate,
		Read:   resourceWirelessControllerHotspot20H2QpOsuProviderNaiRead,
		Update: resourceWirelessControllerHotspot20H2QpOsuProviderNaiUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpOsuProviderNaiDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"nai_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"osu_nai": &schema.Schema{
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

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProviderNai(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProviderNai resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20H2QpOsuProviderNai(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpOsuProviderNai resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpOsuProviderNai")
	}

	return resourceWirelessControllerHotspot20H2QpOsuProviderNaiRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerHotspot20H2QpOsuProviderNai(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProviderNai resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20H2QpOsuProviderNai(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpOsuProviderNai resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpOsuProviderNai")
	}

	return resourceWirelessControllerHotspot20H2QpOsuProviderNaiRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20H2QpOsuProviderNai(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpOsuProviderNai resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpOsuProviderNaiRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerHotspot20H2QpOsuProviderNai(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProviderNai resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpOsuProviderNai(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpOsuProviderNai resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["osu-nai"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "osu_nai"
			}
			tmp["osu_nai"] = flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpOsuProviderNai(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpOsuProviderNaiName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("nai_list", flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(o["nai-list"], d, "nai_list", sv)); err != nil {
			if !fortiAPIPatch(o["nai-list"]) {
				return fmt.Errorf("Error reading nai_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nai_list"); ok {
			if err = d.Set("nai_list", flattenWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(o["nai-list"], d, "nai_list", sv)); err != nil {
				if !fortiAPIPatch(o["nai-list"]) {
					return fmt.Errorf("Error reading nai_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpOsuProviderNaiFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20H2QpOsuProviderNaiName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "osu_nai"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["osu-nai"], _ = expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai(d, i["osu_nai"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["osu-nai"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiListOsuNai(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpOsuProviderNai(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderNaiName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nai_list"); ok || d.HasChange("nai_list") {
		t, err := expandWirelessControllerHotspot20H2QpOsuProviderNaiNaiList(d, v, "nai_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nai-list"] = t
		}
	}

	return &obj, nil
}
