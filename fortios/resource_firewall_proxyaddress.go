// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Web proxy address configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallProxyAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxyAddressCreate,
		Read:   resourceFirewallProxyAddressRead,
		Update: resourceFirewallProxyAddressUpdate,
		Delete: resourceFirewallProxyAddressDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"host_regex": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"query": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"referrer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ua": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ua_min_ver": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"ua_max_ver": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"header_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"header": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"header_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"category": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceFirewallProxyAddressCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyAddress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallProxyAddress resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallProxyAddress(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallProxyAddress resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallProxyAddress")
	}

	return resourceFirewallProxyAddressRead(d, m)
}

func resourceFirewallProxyAddressUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallProxyAddress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallProxyAddress(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallProxyAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallProxyAddress")
	}

	return resourceFirewallProxyAddressRead(d, m)
}

func resourceFirewallProxyAddressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallProxyAddress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallProxyAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyAddressRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallProxyAddress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxyAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallProxyAddress(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallProxyAddress resource from API: %v", err)
	}
	return nil
}

func flattenFirewallProxyAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHostRegex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressQuery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressReferrer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallProxyAddressCategoryId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallProxyAddressCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallProxyAddressMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressUa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressUaMinVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressUaMaxVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressCaseSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallProxyAddressHeaderGroupId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if cur_v, ok := i["header-name"]; ok {
			tmp["header_name"] = flattenFirewallProxyAddressHeaderGroupHeaderName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenFirewallProxyAddressHeaderGroupHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if cur_v, ok := i["case-sensitivity"]; ok {
			tmp["case_sensitivity"] = flattenFirewallProxyAddressHeaderGroupCaseSensitivity(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallProxyAddressHeaderGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallProxyAddressHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallProxyAddressTagging(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallProxyAddressTaggingName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenFirewallProxyAddressTaggingCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if cur_v, ok := i["tags"]; ok {
			tmp["tags"] = flattenFirewallProxyAddressTaggingTags(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallProxyAddressTaggingName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallProxyAddressTaggingTagsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallProxyAddressTaggingTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallProxyAddressApplicationName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallProxyAddressApplicationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressVisibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallProxyAddress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenFirewallProxyAddressName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallProxyAddressUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallProxyAddressType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("host", flattenFirewallProxyAddressHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_regex", flattenFirewallProxyAddressHostRegex(o["host-regex"], d, "host_regex", sv)); err != nil {
		if !fortiAPIPatch(o["host-regex"]) {
			return fmt.Errorf("Error reading host_regex: %v", err)
		}
	}

	if err = d.Set("path", flattenFirewallProxyAddressPath(o["path"], d, "path", sv)); err != nil {
		if !fortiAPIPatch(o["path"]) {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("query", flattenFirewallProxyAddressQuery(o["query"], d, "query", sv)); err != nil {
		if !fortiAPIPatch(o["query"]) {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("referrer", flattenFirewallProxyAddressReferrer(o["referrer"], d, "referrer", sv)); err != nil {
		if !fortiAPIPatch(o["referrer"]) {
			return fmt.Errorf("Error reading referrer: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("category", flattenFirewallProxyAddressCategory(o["category"], d, "category", sv)); err != nil {
			if !fortiAPIPatch(o["category"]) {
				return fmt.Errorf("Error reading category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("category"); ok {
			if err = d.Set("category", flattenFirewallProxyAddressCategory(o["category"], d, "category", sv)); err != nil {
				if !fortiAPIPatch(o["category"]) {
					return fmt.Errorf("Error reading category: %v", err)
				}
			}
		}
	}

	if err = d.Set("method", flattenFirewallProxyAddressMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("ua", flattenFirewallProxyAddressUa(o["ua"], d, "ua", sv)); err != nil {
		if !fortiAPIPatch(o["ua"]) {
			return fmt.Errorf("Error reading ua: %v", err)
		}
	}

	if err = d.Set("ua_min_ver", flattenFirewallProxyAddressUaMinVer(o["ua-min-ver"], d, "ua_min_ver", sv)); err != nil {
		if !fortiAPIPatch(o["ua-min-ver"]) {
			return fmt.Errorf("Error reading ua_min_ver: %v", err)
		}
	}

	if err = d.Set("ua_max_ver", flattenFirewallProxyAddressUaMaxVer(o["ua-max-ver"], d, "ua_max_ver", sv)); err != nil {
		if !fortiAPIPatch(o["ua-max-ver"]) {
			return fmt.Errorf("Error reading ua_max_ver: %v", err)
		}
	}

	if err = d.Set("header_name", flattenFirewallProxyAddressHeaderName(o["header-name"], d, "header_name", sv)); err != nil {
		if !fortiAPIPatch(o["header-name"]) {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("header", flattenFirewallProxyAddressHeader(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("case_sensitivity", flattenFirewallProxyAddressCaseSensitivity(o["case-sensitivity"], d, "case_sensitivity", sv)); err != nil {
		if !fortiAPIPatch(o["case-sensitivity"]) {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("header_group", flattenFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group", sv)); err != nil {
			if !fortiAPIPatch(o["header-group"]) {
				return fmt.Errorf("Error reading header_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("header_group"); ok {
			if err = d.Set("header_group", flattenFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group", sv)); err != nil {
				if !fortiAPIPatch(o["header-group"]) {
					return fmt.Errorf("Error reading header_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("color", flattenFirewallProxyAddressColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("tagging", flattenFirewallProxyAddressTagging(o["tagging"], d, "tagging", sv)); err != nil {
			if !fortiAPIPatch(o["tagging"]) {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallProxyAddressTagging(o["tagging"], d, "tagging", sv)); err != nil {
				if !fortiAPIPatch(o["tagging"]) {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenFirewallProxyAddressComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("application", flattenFirewallProxyAddressApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallProxyAddressApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if err = d.Set("visibility", flattenFirewallProxyAddressVisibility(o["visibility"], d, "visibility", sv)); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	return nil
}

func flattenFirewallProxyAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallProxyAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHostRegex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressQuery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressReferrer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallProxyAddressCategoryId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUaMinVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressUaMaxVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressCaseSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallProxyAddressHeaderGroupId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header-name"], _ = expandFirewallProxyAddressHeaderGroupHeaderName(d, i["header_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["header-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandFirewallProxyAddressHeaderGroupHeader(d, i["header"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["header"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["case-sensitivity"], _ = expandFirewallProxyAddressHeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressHeaderGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressHeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallProxyAddressTaggingName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandFirewallProxyAddressTaggingCategory(d, i["category"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["category"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tags"], _ = expandFirewallProxyAddressTaggingTags(d, i["tags"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tags"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressTaggingName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallProxyAddressTaggingTagsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressTaggingTagsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallProxyAddressApplicationName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallProxyAddressApplicationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallProxyAddressVisibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallProxyAddress(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallProxyAddressName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallProxyAddressUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandFirewallProxyAddressType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandFirewallProxyAddressHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	} else if d.HasChange("host") {
		obj["host"] = nil
	}

	if v, ok := d.GetOk("host_regex"); ok {
		t, err := expandFirewallProxyAddressHostRegex(d, v, "host_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-regex"] = t
		}
	} else if d.HasChange("host_regex") {
		obj["host-regex"] = nil
	}

	if v, ok := d.GetOk("path"); ok {
		t, err := expandFirewallProxyAddressPath(d, v, "path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	} else if d.HasChange("path") {
		obj["path"] = nil
	}

	if v, ok := d.GetOk("query"); ok {
		t, err := expandFirewallProxyAddressQuery(d, v, "query", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	} else if d.HasChange("query") {
		obj["query"] = nil
	}

	if v, ok := d.GetOk("referrer"); ok {
		t, err := expandFirewallProxyAddressReferrer(d, v, "referrer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok || d.HasChange("category") {
		t, err := expandFirewallProxyAddressCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandFirewallProxyAddressMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	} else if d.HasChange("method") {
		obj["method"] = nil
	}

	if v, ok := d.GetOk("ua"); ok {
		t, err := expandFirewallProxyAddressUa(d, v, "ua", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua"] = t
		}
	} else if d.HasChange("ua") {
		obj["ua"] = nil
	}

	if v, ok := d.GetOk("ua_min_ver"); ok {
		t, err := expandFirewallProxyAddressUaMinVer(d, v, "ua_min_ver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua-min-ver"] = t
		}
	} else if d.HasChange("ua_min_ver") {
		obj["ua-min-ver"] = nil
	}

	if v, ok := d.GetOk("ua_max_ver"); ok {
		t, err := expandFirewallProxyAddressUaMaxVer(d, v, "ua_max_ver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua-max-ver"] = t
		}
	} else if d.HasChange("ua_max_ver") {
		obj["ua-max-ver"] = nil
	}

	if v, ok := d.GetOk("header_name"); ok {
		t, err := expandFirewallProxyAddressHeaderName(d, v, "header_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	} else if d.HasChange("header_name") {
		obj["header-name"] = nil
	}

	if v, ok := d.GetOk("header"); ok {
		t, err := expandFirewallProxyAddressHeader(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	} else if d.HasChange("header") {
		obj["header"] = nil
	}

	if v, ok := d.GetOk("case_sensitivity"); ok {
		t, err := expandFirewallProxyAddressCaseSensitivity(d, v, "case_sensitivity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("header_group"); ok || d.HasChange("header_group") {
		t, err := expandFirewallProxyAddressHeaderGroup(d, v, "header_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-group"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {
		t, err := expandFirewallProxyAddressColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	} else if d.HasChange("color") {
		obj["color"] = nil
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
		t, err := expandFirewallProxyAddressTagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallProxyAddressComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallProxyAddressApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok {
		t, err := expandFirewallProxyAddressVisibility(d, v, "visibility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	} else if d.HasChange("visibility") {
		obj["visibility"] = nil
	}

	return &obj, nil
}
