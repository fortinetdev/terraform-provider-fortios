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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallProxyAddress() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallProxyAddressRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_regex": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"query": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"referrer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ua": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"header_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"case_sensitivity": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallProxyAddressRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing FirewallProxyAddress: type error")
	}

	o, err := c.ReadFirewallProxyAddress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallProxyAddress: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallProxyAddress(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallProxyAddress from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallProxyAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHostRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressQuery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressReferrer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallProxyAddressCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddressCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressUa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHeaderGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallProxyAddressHeaderGroupId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {
			tmp["header_name"] = dataSourceFlattenFirewallProxyAddressHeaderGroupHeaderName(i["header-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenFirewallProxyAddressHeaderGroupHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			tmp["case_sensitivity"] = dataSourceFlattenFirewallProxyAddressHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddressHeaderGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallProxyAddressTaggingName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = dataSourceFlattenFirewallProxyAddressTaggingCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			tmp["tags"] = dataSourceFlattenFirewallProxyAddressTaggingTags(i["tags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddressTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallProxyAddressTaggingTagsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallProxyAddressTaggingTagsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallProxyAddressVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallProxyAddress(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallProxyAddressName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenFirewallProxyAddressUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenFirewallProxyAddressType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("host", dataSourceFlattenFirewallProxyAddressHost(o["host"], d, "host")); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("host_regex", dataSourceFlattenFirewallProxyAddressHostRegex(o["host-regex"], d, "host_regex")); err != nil {
		if !fortiAPIPatch(o["host-regex"]) {
			return fmt.Errorf("Error reading host_regex: %v", err)
		}
	}

	if err = d.Set("path", dataSourceFlattenFirewallProxyAddressPath(o["path"], d, "path")); err != nil {
		if !fortiAPIPatch(o["path"]) {
			return fmt.Errorf("Error reading path: %v", err)
		}
	}

	if err = d.Set("query", dataSourceFlattenFirewallProxyAddressQuery(o["query"], d, "query")); err != nil {
		if !fortiAPIPatch(o["query"]) {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("referrer", dataSourceFlattenFirewallProxyAddressReferrer(o["referrer"], d, "referrer")); err != nil {
		if !fortiAPIPatch(o["referrer"]) {
			return fmt.Errorf("Error reading referrer: %v", err)
		}
	}

	if err = d.Set("category", dataSourceFlattenFirewallProxyAddressCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("method", dataSourceFlattenFirewallProxyAddressMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("ua", dataSourceFlattenFirewallProxyAddressUa(o["ua"], d, "ua")); err != nil {
		if !fortiAPIPatch(o["ua"]) {
			return fmt.Errorf("Error reading ua: %v", err)
		}
	}

	if err = d.Set("header_name", dataSourceFlattenFirewallProxyAddressHeaderName(o["header-name"], d, "header_name")); err != nil {
		if !fortiAPIPatch(o["header-name"]) {
			return fmt.Errorf("Error reading header_name: %v", err)
		}
	}

	if err = d.Set("header", dataSourceFlattenFirewallProxyAddressHeader(o["header"], d, "header")); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("case_sensitivity", dataSourceFlattenFirewallProxyAddressCaseSensitivity(o["case-sensitivity"], d, "case_sensitivity")); err != nil {
		if !fortiAPIPatch(o["case-sensitivity"]) {
			return fmt.Errorf("Error reading case_sensitivity: %v", err)
		}
	}

	if err = d.Set("header_group", dataSourceFlattenFirewallProxyAddressHeaderGroup(o["header-group"], d, "header_group")); err != nil {
		if !fortiAPIPatch(o["header-group"]) {
			return fmt.Errorf("Error reading header_group: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallProxyAddressColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("tagging", dataSourceFlattenFirewallProxyAddressTagging(o["tagging"], d, "tagging")); err != nil {
		if !fortiAPIPatch(o["tagging"]) {
			return fmt.Errorf("Error reading tagging: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenFirewallProxyAddressComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("visibility", dataSourceFlattenFirewallProxyAddressVisibility(o["visibility"], d, "visibility")); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallProxyAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
