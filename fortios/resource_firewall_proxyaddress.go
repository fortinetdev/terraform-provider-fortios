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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
			},
			"host_regex": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"query": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"referrer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ua": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"header": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
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
							Computed: true,
						},
						"header_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"header": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
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
				Computed:     true,
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
							Computed:     true,
						},
						"category": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
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
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

			tmp["id"] = flattenFirewallProxyAddressCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallProxyAddressCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressUa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

			tmp["id"] = flattenFirewallProxyAddressHeaderGroupId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {

			tmp["header_name"] = flattenFirewallProxyAddressHeaderGroupHeaderName(i["header-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {

			tmp["header"] = flattenFirewallProxyAddressHeaderGroupHeader(i["header"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {

			tmp["case_sensitivity"] = flattenFirewallProxyAddressHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallProxyAddressHeaderGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
	return v
}

func flattenFirewallProxyAddressTagging(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
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

			tmp["name"] = flattenFirewallProxyAddressTaggingName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenFirewallProxyAddressTaggingCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {

			tmp["tags"] = flattenFirewallProxyAddressTaggingTags(i["tags"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
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

			tmp["name"] = flattenFirewallProxyAddressTaggingTagsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallProxyAddressTaggingTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallProxyAddressVisibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallProxyAddress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

			tmp["id"], _ = expandFirewallProxyAddressCategoryId(d, i["id"], pre_append, sv)
		}

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

			tmp["id"], _ = expandFirewallProxyAddressHeaderGroupId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["header-name"], _ = expandFirewallProxyAddressHeaderGroupHeaderName(d, i["header_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["header"], _ = expandFirewallProxyAddressHeaderGroupHeader(d, i["header"], pre_append, sv)
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

			tmp["name"], _ = expandFirewallProxyAddressTaggingName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandFirewallProxyAddressTaggingCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tags"], _ = expandFirewallProxyAddressTaggingTags(d, i["tags"], pre_append, sv)
		} else {
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

			tmp["name"], _ = expandFirewallProxyAddressTaggingTagsName(d, i["name"], pre_append, sv)
		}

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
	}

	if v, ok := d.GetOk("host_regex"); ok {

		t, err := expandFirewallProxyAddressHostRegex(d, v, "host_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-regex"] = t
		}
	}

	if v, ok := d.GetOk("path"); ok {

		t, err := expandFirewallProxyAddressPath(d, v, "path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["path"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok {

		t, err := expandFirewallProxyAddressQuery(d, v, "query", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("referrer"); ok {

		t, err := expandFirewallProxyAddressReferrer(d, v, "referrer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["referrer"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {

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
	}

	if v, ok := d.GetOk("ua"); ok {

		t, err := expandFirewallProxyAddressUa(d, v, "ua", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ua"] = t
		}
	}

	if v, ok := d.GetOk("header_name"); ok {

		t, err := expandFirewallProxyAddressHeaderName(d, v, "header_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-name"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok {

		t, err := expandFirewallProxyAddressHeader(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("case_sensitivity"); ok {

		t, err := expandFirewallProxyAddressCaseSensitivity(d, v, "case_sensitivity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("header_group"); ok {

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
	}

	if v, ok := d.GetOk("tagging"); ok {

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
	}

	if v, ok := d.GetOk("visibility"); ok {

		t, err := expandFirewallProxyAddressVisibility(d, v, "visibility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	return &obj, nil
}
