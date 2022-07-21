// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure BFD.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterBfd() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBfdUpdate,
		Read:   resourceRouterBfdRead,
		Update: resourceRouterBfdUpdate,
		Delete: resourceRouterBfdDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
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
							Computed: true,
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
		},
	}
}

func resourceRouterBfdUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterBfd(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterBfd(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterBfd resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterBfd")
	}

	return resourceRouterBfdRead(d, m)
}

func resourceRouterBfdDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterBfd(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterBfd resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBfd(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterBfd resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBfdRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterBfd(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBfd(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBfd resource from API: %v", err)
	}
	return nil
}

func flattenRouterBfdNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenRouterBfdNeighborIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterBfdNeighborInterface(i["interface"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenRouterBfdNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdNeighborInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdMultihopTemplate(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterBfdMultihopTemplateId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {

			tmp["src"] = flattenRouterBfdMultihopTemplateSrc(i["src"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {

			tmp["dst"] = flattenRouterBfdMultihopTemplateDst(i["dst"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := i["bfd-desired-min-tx"]; ok {

			tmp["bfd_desired_min_tx"] = flattenRouterBfdMultihopTemplateBfdDesiredMinTx(i["bfd-desired-min-tx"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := i["bfd-required-min-rx"]; ok {

			tmp["bfd_required_min_rx"] = flattenRouterBfdMultihopTemplateBfdRequiredMinRx(i["bfd-required-min-rx"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := i["bfd-detect-mult"]; ok {

			tmp["bfd_detect_mult"] = flattenRouterBfdMultihopTemplateBfdDetectMult(i["bfd-detect-mult"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {

			tmp["auth_mode"] = flattenRouterBfdMultihopTemplateAuthMode(i["auth-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := i["md5-key"]; ok {

			tmp["md5_key"] = flattenRouterBfdMultihopTemplateMd5Key(i["md5-key"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["md5_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBfdMultihopTemplateId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdMultihopTemplateSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBfdMultihopTemplateDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBfdMultihopTemplateBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdMultihopTemplateBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdMultihopTemplateBfdDetectMult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdMultihopTemplateAuthMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBfdMultihopTemplateMd5Key(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBfd(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterBfdNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBfdNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("multihop_template", flattenRouterBfdMultihopTemplate(o["multihop-template"], d, "multihop_template", sv)); err != nil {
			if !fortiAPIPatch(o["multihop-template"]) {
				return fmt.Errorf("Error reading multihop_template: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("multihop_template"); ok {
			if err = d.Set("multihop_template", flattenRouterBfdMultihopTemplate(o["multihop-template"], d, "multihop_template", sv)); err != nil {
				if !fortiAPIPatch(o["multihop-template"]) {
					return fmt.Errorf("Error reading multihop_template: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterBfdFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterBfdNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandRouterBfdNeighborIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterBfdNeighborInterface(d, i["interface"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBfdNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdNeighborInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterBfdMultihopTemplateId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src"], _ = expandRouterBfdMultihopTemplateSrc(d, i["src"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dst"], _ = expandRouterBfdMultihopTemplateDst(d, i["dst"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd-desired-min-tx"], _ = expandRouterBfdMultihopTemplateBfdDesiredMinTx(d, i["bfd_desired_min_tx"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd-required-min-rx"], _ = expandRouterBfdMultihopTemplateBfdRequiredMinRx(d, i["bfd_required_min_rx"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bfd-detect-mult"], _ = expandRouterBfdMultihopTemplateBfdDetectMult(d, i["bfd_detect_mult"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-mode"], _ = expandRouterBfdMultihopTemplateAuthMode(d, i["auth_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["md5-key"], _ = expandRouterBfdMultihopTemplateMd5Key(d, i["md5_key"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBfdMultihopTemplateId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateBfdDetectMult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateAuthMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBfdMultihopTemplateMd5Key(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBfd(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("neighbor"); ok {
		if setArgNil {
			obj["neighbor"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterBfdNeighbor(d, v, "neighbor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	if v, ok := d.GetOk("multihop_template"); ok {
		if setArgNil {
			obj["multihop-template"] = make([]struct{}, 0)
		} else {

			t, err := expandRouterBfdMultihopTemplate(d, v, "multihop_template", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multihop-template"] = t
			}
		}
	}

	return &obj, nil
}
