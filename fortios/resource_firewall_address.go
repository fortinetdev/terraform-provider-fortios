// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 addresses.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAddressCreate,
		Read:   resourceFirewallAddressRead,
		Update: resourceFirewallAddressUpdate,
		Delete: resourceFirewallAddressDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sub_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clearpass_spt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"country": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				Optional:     true,
				Computed:     true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fsso_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"tenant": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"organization": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"epg_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"subnet_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"sdn_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"policy_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"obj_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"associated_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
			},
			"sdn_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_ip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
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
			"allow_routing": &schema.Schema{
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

func resourceFirewallAddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallAddress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallAddress(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallAddress resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAddress")
	}

	return resourceFirewallAddressRead(d, m)
}

func resourceFirewallAddressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallAddress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAddress(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAddress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAddress")
	}

	return resourceFirewallAddressRead(d, m)
}

func resourceFirewallAddressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallAddress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallAddress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAddress(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAddress resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenFirewallAddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSubType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressClearpassSpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressStartMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressEndMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressWildcardFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressCacheTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressWildcard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenFirewallAddressSdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressFssoGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallAddressFssoGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAddressFssoGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressTenant(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressOrganization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressEpgName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSubnetName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSdnTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressPolicyGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressObjTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressObjType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressVisibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressAssociatedInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSdnAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressNodeIpOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressObjId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallAddressListIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenFirewallAddressListIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressTagging(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallAddressTaggingName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenFirewallAddressTaggingCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {

			tmp["tags"] = flattenFirewallAddressTaggingTags(i["tags"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAddressTaggingName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressTaggingCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressTaggingTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallAddressTaggingTagsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallAddressTaggingTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressAllowRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAddress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallAddressName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallAddressUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("subnet", flattenFirewallAddressSubnet(o["subnet"], d, "subnet", sv)); err != nil {
		if !fortiAPIPatch(o["subnet"]) {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallAddressType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("sub_type", flattenFirewallAddressSubType(o["sub-type"], d, "sub_type", sv)); err != nil {
		if !fortiAPIPatch(o["sub-type"]) {
			return fmt.Errorf("Error reading sub_type: %v", err)
		}
	}

	if err = d.Set("clearpass_spt", flattenFirewallAddressClearpassSpt(o["clearpass-spt"], d, "clearpass_spt", sv)); err != nil {
		if !fortiAPIPatch(o["clearpass-spt"]) {
			return fmt.Errorf("Error reading clearpass_spt: %v", err)
		}
	}

	if err = d.Set("start_mac", flattenFirewallAddressStartMac(o["start-mac"], d, "start_mac", sv)); err != nil {
		if !fortiAPIPatch(o["start-mac"]) {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if err = d.Set("end_mac", flattenFirewallAddressEndMac(o["end-mac"], d, "end_mac", sv)); err != nil {
		if !fortiAPIPatch(o["end-mac"]) {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenFirewallAddressStartIp(o["start-ip"], d, "start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["start-ip"]) {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("end_ip", flattenFirewallAddressEndIp(o["end-ip"], d, "end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["end-ip"]) {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenFirewallAddressFqdn(o["fqdn"], d, "fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("country", flattenFirewallAddressCountry(o["country"], d, "country", sv)); err != nil {
		if !fortiAPIPatch(o["country"]) {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenFirewallAddressWildcardFqdn(o["wildcard-fqdn"], d, "wildcard_fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["wildcard-fqdn"]) {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	if err = d.Set("cache_ttl", flattenFirewallAddressCacheTtl(o["cache-ttl"], d, "cache_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["cache-ttl"]) {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("wildcard", flattenFirewallAddressWildcard(o["wildcard"], d, "wildcard", sv)); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("sdn", flattenFirewallAddressSdn(o["sdn"], d, "sdn", sv)); err != nil {
		if !fortiAPIPatch(o["sdn"]) {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fsso_group", flattenFirewallAddressFssoGroup(o["fsso-group"], d, "fsso_group", sv)); err != nil {
			if !fortiAPIPatch(o["fsso-group"]) {
				return fmt.Errorf("Error reading fsso_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fsso_group"); ok {
			if err = d.Set("fsso_group", flattenFirewallAddressFssoGroup(o["fsso-group"], d, "fsso_group", sv)); err != nil {
				if !fortiAPIPatch(o["fsso-group"]) {
					return fmt.Errorf("Error reading fsso_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("interface", flattenFirewallAddressInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("tenant", flattenFirewallAddressTenant(o["tenant"], d, "tenant", sv)); err != nil {
		if !fortiAPIPatch(o["tenant"]) {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("organization", flattenFirewallAddressOrganization(o["organization"], d, "organization", sv)); err != nil {
		if !fortiAPIPatch(o["organization"]) {
			return fmt.Errorf("Error reading organization: %v", err)
		}
	}

	if err = d.Set("epg_name", flattenFirewallAddressEpgName(o["epg-name"], d, "epg_name", sv)); err != nil {
		if !fortiAPIPatch(o["epg-name"]) {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("subnet_name", flattenFirewallAddressSubnetName(o["subnet-name"], d, "subnet_name", sv)); err != nil {
		if !fortiAPIPatch(o["subnet-name"]) {
			return fmt.Errorf("Error reading subnet_name: %v", err)
		}
	}

	if err = d.Set("sdn_tag", flattenFirewallAddressSdnTag(o["sdn-tag"], d, "sdn_tag", sv)); err != nil {
		if !fortiAPIPatch(o["sdn-tag"]) {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("policy_group", flattenFirewallAddressPolicyGroup(o["policy-group"], d, "policy_group", sv)); err != nil {
		if !fortiAPIPatch(o["policy-group"]) {
			return fmt.Errorf("Error reading policy_group: %v", err)
		}
	}

	if err = d.Set("obj_tag", flattenFirewallAddressObjTag(o["obj-tag"], d, "obj_tag", sv)); err != nil {
		if !fortiAPIPatch(o["obj-tag"]) {
			return fmt.Errorf("Error reading obj_tag: %v", err)
		}
	}

	if err = d.Set("obj_type", flattenFirewallAddressObjType(o["obj-type"], d, "obj_type", sv)); err != nil {
		if !fortiAPIPatch(o["obj-type"]) {
			return fmt.Errorf("Error reading obj_type: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallAddressComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("visibility", flattenFirewallAddressVisibility(o["visibility"], d, "visibility", sv)); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("associated_interface", flattenFirewallAddressAssociatedInterface(o["associated-interface"], d, "associated_interface", sv)); err != nil {
		if !fortiAPIPatch(o["associated-interface"]) {
			return fmt.Errorf("Error reading associated_interface: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallAddressColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("filter", flattenFirewallAddressFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", flattenFirewallAddressSdnAddrType(o["sdn-addr-type"], d, "sdn_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["sdn-addr-type"]) {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("node_ip_only", flattenFirewallAddressNodeIpOnly(o["node-ip-only"], d, "node_ip_only", sv)); err != nil {
		if !fortiAPIPatch(o["node-ip-only"]) {
			return fmt.Errorf("Error reading node_ip_only: %v", err)
		}
	}

	if err = d.Set("obj_id", flattenFirewallAddressObjId(o["obj-id"], d, "obj_id", sv)); err != nil {
		if !fortiAPIPatch(o["obj-id"]) {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("list", flattenFirewallAddressList(o["list"], d, "list", sv)); err != nil {
			if !fortiAPIPatch(o["list"]) {
				return fmt.Errorf("Error reading list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list"); ok {
			if err = d.Set("list", flattenFirewallAddressList(o["list"], d, "list", sv)); err != nil {
				if !fortiAPIPatch(o["list"]) {
					return fmt.Errorf("Error reading list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenFirewallAddressTagging(o["tagging"], d, "tagging", sv)); err != nil {
			if !fortiAPIPatch(o["tagging"]) {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenFirewallAddressTagging(o["tagging"], d, "tagging", sv)); err != nil {
				if !fortiAPIPatch(o["tagging"]) {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if err = d.Set("allow_routing", flattenFirewallAddressAllowRouting(o["allow-routing"], d, "allow_routing", sv)); err != nil {
		if !fortiAPIPatch(o["allow-routing"]) {
			return fmt.Errorf("Error reading allow_routing: %v", err)
		}
	}

	return nil
}

func flattenFirewallAddressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSubType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressClearpassSpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressStartMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressEndMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressWildcardFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressCacheTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressWildcard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFssoGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallAddressFssoGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddressFssoGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTenant(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressOrganization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressEpgName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSubnetName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSdnTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressPolicyGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressObjTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressObjType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressVisibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressAssociatedInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSdnAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressNodeIpOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressObjId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["ip"], _ = expandFirewallAddressListIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddressListIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallAddressTaggingName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandFirewallAddressTaggingCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tags"], _ = expandFirewallAddressTaggingTags(d, i["tags"], pre_append, sv)
		} else {
			tmp["tags"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddressTaggingName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTaggingCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallAddressTaggingTagsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddressTaggingTagsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressAllowRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAddress(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallAddressName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {

		t, err := expandFirewallAddressUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok {

		t, err := expandFirewallAddressSubnet(d, v, "subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallAddressType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("sub_type"); ok {

		t, err := expandFirewallAddressSubType(d, v, "sub_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub-type"] = t
		}
	}

	if v, ok := d.GetOk("clearpass_spt"); ok {

		t, err := expandFirewallAddressClearpassSpt(d, v, "clearpass_spt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clearpass-spt"] = t
		}
	}

	if v, ok := d.GetOk("start_mac"); ok {

		t, err := expandFirewallAddressStartMac(d, v, "start_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-mac"] = t
		}
	}

	if v, ok := d.GetOk("end_mac"); ok {

		t, err := expandFirewallAddressEndMac(d, v, "end_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-mac"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok {

		t, err := expandFirewallAddressStartIp(d, v, "start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok {

		t, err := expandFirewallAddressEndIp(d, v, "end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {

		t, err := expandFirewallAddressFqdn(d, v, "fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok {

		t, err := expandFirewallAddressCountry(d, v, "country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok {

		t, err := expandFirewallAddressWildcardFqdn(d, v, "wildcard_fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	if v, ok := d.GetOkExists("cache_ttl"); ok {

		t, err := expandFirewallAddressCacheTtl(d, v, "cache_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("wildcard"); ok {

		t, err := expandFirewallAddressWildcard(d, v, "wildcard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard"] = t
		}
	}

	if v, ok := d.GetOk("sdn"); ok {

		t, err := expandFirewallAddressSdn(d, v, "sdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}

	if v, ok := d.GetOk("fsso_group"); ok {

		t, err := expandFirewallAddressFssoGroup(d, v, "fsso_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-group"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandFirewallAddressInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("tenant"); ok {

		t, err := expandFirewallAddressTenant(d, v, "tenant", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	}

	if v, ok := d.GetOk("organization"); ok {

		t, err := expandFirewallAddressOrganization(d, v, "organization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}

	if v, ok := d.GetOk("epg_name"); ok {

		t, err := expandFirewallAddressEpgName(d, v, "epg_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	}

	if v, ok := d.GetOk("subnet_name"); ok {

		t, err := expandFirewallAddressSubnetName(d, v, "subnet_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-name"] = t
		}
	}

	if v, ok := d.GetOk("sdn_tag"); ok {

		t, err := expandFirewallAddressSdnTag(d, v, "sdn_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	}

	if v, ok := d.GetOk("policy_group"); ok {

		t, err := expandFirewallAddressPolicyGroup(d, v, "policy_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-group"] = t
		}
	}

	if v, ok := d.GetOk("obj_tag"); ok {

		t, err := expandFirewallAddressObjTag(d, v, "obj_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-tag"] = t
		}
	}

	if v, ok := d.GetOk("obj_type"); ok {

		t, err := expandFirewallAddressObjType(d, v, "obj_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-type"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallAddressComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok {

		t, err := expandFirewallAddressVisibility(d, v, "visibility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	if v, ok := d.GetOk("associated_interface"); ok {

		t, err := expandFirewallAddressAssociatedInterface(d, v, "associated_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandFirewallAddressColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {

		t, err := expandFirewallAddressFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("sdn_addr_type"); ok {

		t, err := expandFirewallAddressSdnAddrType(d, v, "sdn_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}

	if v, ok := d.GetOk("node_ip_only"); ok {

		t, err := expandFirewallAddressNodeIpOnly(d, v, "node_ip_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("obj_id"); ok {

		t, err := expandFirewallAddressObjId(d, v, "obj_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-id"] = t
		}
	}

	if v, ok := d.GetOk("list"); ok {

		t, err := expandFirewallAddressList(d, v, "list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok {

		t, err := expandFirewallAddressTagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("allow_routing"); ok {

		t, err := expandFirewallAddressAllowRouting(d, v, "allow_routing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-routing"] = t
		}
	}

	return &obj, nil
}
