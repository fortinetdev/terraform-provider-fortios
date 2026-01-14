// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemGeoipOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGeoipOverrideCreate,
		Read:   resourceSystemGeoipOverrideRead,
		Update: resourceSystemGeoipOverrideUpdate,
		Delete: resourceSystemGeoipOverrideDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Required:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"country_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				Optional:     true,
				Computed:     true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ip6_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
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

func resourceSystemGeoipOverrideCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemGeoipOverride(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeoipOverride resource while getting object: %v", err)
	}

	o, err := c.CreateSystemGeoipOverride(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemGeoipOverride resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGeoipOverride")
	}

	return resourceSystemGeoipOverrideRead(d, m)
}

func resourceSystemGeoipOverrideUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemGeoipOverride(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeoipOverride resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemGeoipOverride(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeoipOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGeoipOverride")
	}

	return resourceSystemGeoipOverrideRead(d, m)
}

func resourceSystemGeoipOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemGeoipOverride(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGeoipOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGeoipOverrideRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemGeoipOverride(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeoipOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGeoipOverride(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeoipOverride resource from API: %v", err)
	}
	return nil
}

func flattenSystemGeoipOverrideName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeoipOverrideDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeoipOverrideCountryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIpRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemGeoipOverrideIpRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
			}
			tmp["start_ip"] = flattenSystemGeoipOverrideIpRangeStartIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
			}
			tmp["end_ip"] = flattenSystemGeoipOverrideIpRangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemGeoipOverrideIpRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemGeoipOverrideIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6Range(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemGeoipOverrideIp6RangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
			}
			tmp["start_ip"] = flattenSystemGeoipOverrideIp6RangeStartIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
			}
			tmp["end_ip"] = flattenSystemGeoipOverrideIp6RangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemGeoipOverrideIp6RangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemGeoipOverrideIp6RangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeoipOverrideIp6RangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemGeoipOverride(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemGeoipOverrideName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemGeoipOverrideDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("country_id", flattenSystemGeoipOverrideCountryId(o["country-id"], d, "country_id", sv)); err != nil {
		if !fortiAPIPatch(o["country-id"]) {
			return fmt.Errorf("Error reading country_id: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip_range", flattenSystemGeoipOverrideIpRange(o["ip-range"], d, "ip_range", sv)); err != nil {
			if !fortiAPIPatch(o["ip-range"]) {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemGeoipOverrideIpRange(o["ip-range"], d, "ip_range", sv)); err != nil {
				if !fortiAPIPatch(o["ip-range"]) {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip6_range", flattenSystemGeoipOverrideIp6Range(o["ip6-range"], d, "ip6_range", sv)); err != nil {
			if !fortiAPIPatch(o["ip6-range"]) {
				return fmt.Errorf("Error reading ip6_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip6_range"); ok {
			if err = d.Set("ip6_range", flattenSystemGeoipOverrideIp6Range(o["ip6-range"], d, "ip6_range", sv)); err != nil {
				if !fortiAPIPatch(o["ip6-range"]) {
					return fmt.Errorf("Error reading ip6_range: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemGeoipOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemGeoipOverrideName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideCountryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemGeoipOverrideIpRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandSystemGeoipOverrideIpRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandSystemGeoipOverrideIpRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemGeoipOverrideIpRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6Range(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemGeoipOverrideIp6RangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandSystemGeoipOverrideIp6RangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandSystemGeoipOverrideIp6RangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemGeoipOverrideIp6RangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6RangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeoipOverrideIp6RangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGeoipOverride(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemGeoipOverrideName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemGeoipOverrideDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("country_id"); ok {
		t, err := expandSystemGeoipOverrideCountryId(d, v, "country_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-id"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemGeoipOverrideIpRange(d, v, "ip_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("ip6_range"); ok || d.HasChange("ip6_range") {
		t, err := expandSystemGeoipOverrideIp6Range(d, v, "ip6_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-range"] = t
		}
	}

	return &obj, nil
}
