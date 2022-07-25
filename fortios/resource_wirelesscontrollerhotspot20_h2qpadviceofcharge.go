// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure advice of charge.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpAdviceOfCharge() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpAdviceOfChargeCreate,
		Read:   resourceWirelessControllerHotspot20H2QpAdviceOfChargeRead,
		Update: resourceWirelessControllerHotspot20H2QpAdviceOfChargeUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpAdviceOfChargeDelete,

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
			"aoc_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nai_realm_encoding": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"nai_realm": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"plan_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"lang": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),
										Optional:     true,
										Computed:     true,
									},
									"currency": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),
										Optional:     true,
										Computed:     true,
									},
									"info_file": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
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
		},
	}
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpAdviceOfCharge resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20H2QpAdviceOfCharge(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpAdviceOfCharge")
	}

	return resourceWirelessControllerHotspot20H2QpAdviceOfChargeRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpAdviceOfCharge resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20H2QpAdviceOfCharge(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpAdviceOfCharge")
	}

	return resourceWirelessControllerHotspot20H2QpAdviceOfChargeRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20H2QpAdviceOfCharge(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpAdviceOfChargeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerHotspot20H2QpAdviceOfCharge(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpAdviceOfCharge resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpAdviceOfCharge resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm_encoding"
		if _, ok := i["nai-realm-encoding"]; ok {

			tmp["nai_realm_encoding"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(i["nai-realm-encoding"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := i["nai-realm"]; ok {

			tmp["nai_realm"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(i["nai-realm"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "plan_info"
		if _, ok := i["plan-info"]; ok {

			tmp["plan_info"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(i["plan-info"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {

			tmp["lang"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(i["lang"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "currency"
		if _, ok := i["currency"]; ok {

			tmp["currency"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(i["currency"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "info_file"
		if _, ok := i["info-file"]; ok {

			tmp["info_file"] = flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(i["info-file"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpAdviceOfChargeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("aoc_list", flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocList(o["aoc-list"], d, "aoc_list", sv)); err != nil {
			if !fortiAPIPatch(o["aoc-list"]) {
				return fmt.Errorf("Error reading aoc_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aoc_list"); ok {
			if err = d.Set("aoc_list", flattenWirelessControllerHotspot20H2QpAdviceOfChargeAocList(o["aoc-list"], d, "aoc_list", sv)); err != nil {
				if !fortiAPIPatch(o["aoc-list"]) {
					return fmt.Errorf("Error reading aoc_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpAdviceOfChargeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm_encoding"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["nai-realm-encoding"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(d, i["nai_realm_encoding"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nai_realm"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["nai-realm"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(d, i["nai_realm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "plan_info"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["plan-info"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d, i["plan_info"], pre_append, sv)
		} else {
			tmp["plan-info"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealmEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListNaiRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["lang"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(d, i["lang"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "currency"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["currency"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(d, i["currency"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "info_file"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["info-file"], _ = expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(d, i["info_file"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoLang(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoCurrency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpAdviceOfChargeAocListPlanInfoInfoFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpAdviceOfCharge(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerHotspot20H2QpAdviceOfChargeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("aoc_list"); ok || d.HasChange("aoc_list") {

		t, err := expandWirelessControllerHotspot20H2QpAdviceOfChargeAocList(d, v, "aoc_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aoc-list"] = t
		}
	}

	return &obj, nil
}
