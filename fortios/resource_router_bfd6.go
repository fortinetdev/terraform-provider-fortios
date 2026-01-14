// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 BFD.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBfd6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBfd6Update,
		Read:   resourceRouterBfd6Read,
		Update: resourceRouterBfd6Update,
		Delete: resourceRouterBfd6Delete,

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
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
					},
				},
			},
			"multihop_template": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd_desired_min_tx": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 30000),
							Optional:     true,
							Computed:     true,
						},
						"bfd_required_min_rx": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 30000),
							Optional:     true,
							Computed:     true,
						},
						"bfd_detect_mult": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 50),
							Optional:     true,
							Computed:     true,
						},
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"md5_key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Sensitive:    true,
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

func resourceRouterBfd6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBfd6(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterBfd6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterBfd6")
	}

	return resourceRouterBfd6Read(d, m)
}

func resourceRouterBfd6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterBfd6(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterBfd6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBfd6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterBfd6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfd6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBfd6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBfd6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterBfd6Neighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "ip6-address", "ip6_address")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["ip6-address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
			}
			tmp["ip6_address"] = flattenRouterBfd6NeighborIp6Address(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
			}
			tmp["interface"] = flattenRouterBfd6NeighborInterface(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip6_address", d)
	return result
}

func flattenRouterBfd6NeighborIp6Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfd6NeighborInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplate(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBfd6MultihopTemplateId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["src"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
			}
			tmp["src"] = flattenRouterBfd6MultihopTemplateSrc(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
			}
			tmp["dst"] = flattenRouterBfd6MultihopTemplateDst(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bfd-desired-min-tx"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
			}
			tmp["bfd_desired_min_tx"] = flattenRouterBfd6MultihopTemplateBfdDesiredMinTx(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bfd-required-min-rx"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
			}
			tmp["bfd_required_min_rx"] = flattenRouterBfd6MultihopTemplateBfdRequiredMinRx(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bfd-detect-mult"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
			}
			tmp["bfd_detect_mult"] = flattenRouterBfd6MultihopTemplateBfdDetectMult(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["auth-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
			}
			tmp["auth_mode"] = flattenRouterBfd6MultihopTemplateAuthMode(cur_v, d, pre_append, sv)
		}

		if _, ok := i["md5-key"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "md5_key"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["md5_key"] = c
				}
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBfd6MultihopTemplateId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBfd6MultihopTemplateSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfd6MultihopTemplateBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBfd6MultihopTemplateBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBfd6MultihopTemplateBfdDetectMult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBfd6MultihopTemplateAuthMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBfd6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if b_get_all_tables {
		if err = d.Set("neighbor", flattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("multihop_template", flattenRouterBfd6MultihopTemplate(o["multihop-template"], d, "multihop_template", sv)); err != nil {
			if !fortiAPIPatch(o["multihop-template"]) {
				return fmt.Errorf("Error reading multihop_template: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("multihop_template"); ok {
			if err = d.Set("multihop_template", flattenRouterBfd6MultihopTemplate(o["multihop-template"], d, "multihop_template", sv)); err != nil {
				if !fortiAPIPatch(o["multihop-template"]) {
					return fmt.Errorf("Error reading multihop_template: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterBfd6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterBfd6Neighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6-address"], _ = expandRouterBfd6NeighborIp6Address(d, i["ip6_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBfd6NeighborInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBfd6NeighborIp6Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6NeighborInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBfd6MultihopTemplateId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src"], _ = expandRouterBfd6MultihopTemplateSrc(d, i["src"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandRouterBfd6MultihopTemplateDst(d, i["dst"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd-desired-min-tx"], _ = expandRouterBfd6MultihopTemplateBfdDesiredMinTx(d, i["bfd_desired_min_tx"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd-required-min-rx"], _ = expandRouterBfd6MultihopTemplateBfdRequiredMinRx(d, i["bfd_required_min_rx"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd-detect-mult"], _ = expandRouterBfd6MultihopTemplateBfdDetectMult(d, i["bfd_detect_mult"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-mode"], _ = expandRouterBfd6MultihopTemplateAuthMode(d, i["auth_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-key"], _ = expandRouterBfd6MultihopTemplateMd5Key(d, i["md5_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-key"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBfd6MultihopTemplateId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateBfdDetectMult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateAuthMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfd6MultihopTemplateMd5Key(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBfd6(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		if setArgNil {
			obj["neighbor"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBfd6Neighbor(d, v, "neighbor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	if v, ok := d.GetOk("multihop_template"); ok || d.HasChange("multihop_template") {
		if setArgNil {
			obj["multihop-template"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBfd6MultihopTemplate(d, v, "multihop_template", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multihop-template"] = t
			}
		}
	}

	return &obj, nil
}
