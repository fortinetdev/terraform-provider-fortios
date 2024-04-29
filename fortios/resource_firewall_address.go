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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"route_tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"macaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"macaddr": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
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
			},
			"country": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				Optional:     true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
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
			},
			"fsso_group": &schema.Schema{
				Type:     schema.TypeSet,
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
			},
			"tenant": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"organization": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"epg_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"subnet_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"sdn_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"policy_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"obj_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag_detection_level": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"tag_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"hw_vendor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"hw_model": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"os": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"sw_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
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
				Computed: true,
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
							Type:     schema.TypeSet,
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
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceFirewallAddressCreate(d *schema.ResourceData, m interface{}) error {
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

func flattenFirewallAddressRouteTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSubType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressClearpassSpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressMacaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if cur_v, ok := i["macaddr"]; ok {
			tmp["macaddr"] = flattenFirewallAddressMacaddrMacaddr(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "macaddr", d)
	return result
}

func flattenFirewallAddressMacaddrMacaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
			tmp["name"] = flattenFirewallAddressFssoGroupName(cur_v, d, pre_append, sv)
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

func flattenFirewallAddressTagDetectionLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressTagType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressHwVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressHwModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressSwVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenFirewallAddressListIp(cur_v, d, pre_append, sv)
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
			tmp["name"] = flattenFirewallAddressTaggingName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenFirewallAddressTaggingCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if cur_v, ok := i["tags"]; ok {
			tmp["tags"] = flattenFirewallAddressTaggingTags(cur_v, d, pre_append, sv)
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
			tmp["name"] = flattenFirewallAddressTaggingTagsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAddressTaggingTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressAllowRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAddressFabricObject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAddress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

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

	if err = d.Set("route_tag", flattenFirewallAddressRouteTag(o["route-tag"], d, "route_tag", sv)); err != nil {
		if !fortiAPIPatch(o["route-tag"]) {
			return fmt.Errorf("Error reading route_tag: %v", err)
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

	if b_get_all_tables {
		if err = d.Set("macaddr", flattenFirewallAddressMacaddr(o["macaddr"], d, "macaddr", sv)); err != nil {
			if !fortiAPIPatch(o["macaddr"]) {
				return fmt.Errorf("Error reading macaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("macaddr"); ok {
			if err = d.Set("macaddr", flattenFirewallAddressMacaddr(o["macaddr"], d, "macaddr", sv)); err != nil {
				if !fortiAPIPatch(o["macaddr"]) {
					return fmt.Errorf("Error reading macaddr: %v", err)
				}
			}
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

	if b_get_all_tables {
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

	if err = d.Set("tag_detection_level", flattenFirewallAddressTagDetectionLevel(o["tag-detection-level"], d, "tag_detection_level", sv)); err != nil {
		if !fortiAPIPatch(o["tag-detection-level"]) {
			return fmt.Errorf("Error reading tag_detection_level: %v", err)
		}
	}

	if err = d.Set("tag_type", flattenFirewallAddressTagType(o["tag-type"], d, "tag_type", sv)); err != nil {
		if !fortiAPIPatch(o["tag-type"]) {
			return fmt.Errorf("Error reading tag_type: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenFirewallAddressHwVendor(o["hw-vendor"], d, "hw_vendor", sv)); err != nil {
		if !fortiAPIPatch(o["hw-vendor"]) {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("hw_model", flattenFirewallAddressHwModel(o["hw-model"], d, "hw_model", sv)); err != nil {
		if !fortiAPIPatch(o["hw-model"]) {
			return fmt.Errorf("Error reading hw_model: %v", err)
		}
	}

	if err = d.Set("os", flattenFirewallAddressOs(o["os"], d, "os", sv)); err != nil {
		if !fortiAPIPatch(o["os"]) {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenFirewallAddressSwVersion(o["sw-version"], d, "sw_version", sv)); err != nil {
		if !fortiAPIPatch(o["sw-version"]) {
			return fmt.Errorf("Error reading sw_version: %v", err)
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

	if b_get_all_tables {
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

	if b_get_all_tables {
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

	if err = d.Set("fabric_object", flattenFirewallAddressFabricObject(o["fabric-object"], d, "fabric_object", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
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

func expandFirewallAddressRouteTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSubType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressClearpassSpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressMacaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["macaddr"], _ = expandFirewallAddressMacaddrMacaddr(d, i["macaddr"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAddressMacaddrMacaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallAddressFssoGroupName(d, i["name"], pre_append, sv)

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

func expandFirewallAddressTagDetectionLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressTagType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressHwVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressHwModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAddressSwVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

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
			tmp["name"], _ = expandFirewallAddressTaggingName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandFirewallAddressTaggingCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

		tmp["name"], _ = expandFirewallAddressTaggingTagsName(d, i["name"], pre_append, sv)

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

func expandFirewallAddressFabricObject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
	} else if d.HasChange("name") {
		obj["name"] = nil
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

	if v, ok := d.GetOkExists("route_tag"); ok {
		t, err := expandFirewallAddressRouteTag(d, v, "route_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-tag"] = t
		}
	} else if d.HasChange("route_tag") {
		obj["route-tag"] = nil
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

	if v, ok := d.GetOk("macaddr"); ok || d.HasChange("macaddr") {
		t, err := expandFirewallAddressMacaddr(d, v, "macaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
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
	} else if d.HasChange("fqdn") {
		obj["fqdn"] = nil
	}

	if v, ok := d.GetOk("country"); ok {
		t, err := expandFirewallAddressCountry(d, v, "country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	} else if d.HasChange("country") {
		obj["country"] = nil
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok {
		t, err := expandFirewallAddressWildcardFqdn(d, v, "wildcard_fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	} else if d.HasChange("wildcard_fqdn") {
		obj["wildcard-fqdn"] = nil
	}

	if v, ok := d.GetOkExists("cache_ttl"); ok {
		t, err := expandFirewallAddressCacheTtl(d, v, "cache_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-ttl"] = t
		}
	} else if d.HasChange("cache_ttl") {
		obj["cache-ttl"] = nil
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
	} else if d.HasChange("sdn") {
		obj["sdn"] = nil
	}

	if v, ok := d.GetOk("fsso_group"); ok || d.HasChange("fsso_group") {
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
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("tenant"); ok {
		t, err := expandFirewallAddressTenant(d, v, "tenant", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant"] = t
		}
	} else if d.HasChange("tenant") {
		obj["tenant"] = nil
	}

	if v, ok := d.GetOk("organization"); ok {
		t, err := expandFirewallAddressOrganization(d, v, "organization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	} else if d.HasChange("organization") {
		obj["organization"] = nil
	}

	if v, ok := d.GetOk("epg_name"); ok {
		t, err := expandFirewallAddressEpgName(d, v, "epg_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epg-name"] = t
		}
	} else if d.HasChange("epg_name") {
		obj["epg-name"] = nil
	}

	if v, ok := d.GetOk("subnet_name"); ok {
		t, err := expandFirewallAddressSubnetName(d, v, "subnet_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet-name"] = t
		}
	} else if d.HasChange("subnet_name") {
		obj["subnet-name"] = nil
	}

	if v, ok := d.GetOk("sdn_tag"); ok {
		t, err := expandFirewallAddressSdnTag(d, v, "sdn_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-tag"] = t
		}
	} else if d.HasChange("sdn_tag") {
		obj["sdn-tag"] = nil
	}

	if v, ok := d.GetOk("policy_group"); ok {
		t, err := expandFirewallAddressPolicyGroup(d, v, "policy_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-group"] = t
		}
	} else if d.HasChange("policy_group") {
		obj["policy-group"] = nil
	}

	if v, ok := d.GetOk("obj_tag"); ok {
		t, err := expandFirewallAddressObjTag(d, v, "obj_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-tag"] = t
		}
	} else if d.HasChange("obj_tag") {
		obj["obj-tag"] = nil
	}

	if v, ok := d.GetOk("obj_type"); ok {
		t, err := expandFirewallAddressObjType(d, v, "obj_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obj-type"] = t
		}
	}

	if v, ok := d.GetOk("tag_detection_level"); ok {
		t, err := expandFirewallAddressTagDetectionLevel(d, v, "tag_detection_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-detection-level"] = t
		}
	} else if d.HasChange("tag_detection_level") {
		obj["tag-detection-level"] = nil
	}

	if v, ok := d.GetOk("tag_type"); ok {
		t, err := expandFirewallAddressTagType(d, v, "tag_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag-type"] = t
		}
	} else if d.HasChange("tag_type") {
		obj["tag-type"] = nil
	}

	if v, ok := d.GetOk("hw_vendor"); ok {
		t, err := expandFirewallAddressHwVendor(d, v, "hw_vendor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	} else if d.HasChange("hw_vendor") {
		obj["hw-vendor"] = nil
	}

	if v, ok := d.GetOk("hw_model"); ok {
		t, err := expandFirewallAddressHwModel(d, v, "hw_model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-model"] = t
		}
	} else if d.HasChange("hw_model") {
		obj["hw-model"] = nil
	}

	if v, ok := d.GetOk("os"); ok {
		t, err := expandFirewallAddressOs(d, v, "os", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	} else if d.HasChange("os") {
		obj["os"] = nil
	}

	if v, ok := d.GetOk("sw_version"); ok {
		t, err := expandFirewallAddressSwVersion(d, v, "sw_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	} else if d.HasChange("sw_version") {
		obj["sw-version"] = nil
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallAddressComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("visibility"); ok {
		t, err := expandFirewallAddressVisibility(d, v, "visibility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	} else if d.HasChange("visibility") {
		obj["visibility"] = nil
	}

	if v, ok := d.GetOk("associated_interface"); ok {
		t, err := expandFirewallAddressAssociatedInterface(d, v, "associated_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["associated-interface"] = t
		}
	} else if d.HasChange("associated_interface") {
		obj["associated-interface"] = nil
	}

	if v, ok := d.GetOkExists("color"); ok {
		t, err := expandFirewallAddressColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	} else if d.HasChange("color") {
		obj["color"] = nil
	}

	if v, ok := d.GetOk("filter"); ok {
		t, err := expandFirewallAddressFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	} else if d.HasChange("filter") {
		obj["filter"] = nil
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
	} else if d.HasChange("obj_id") {
		obj["obj-id"] = nil
	}

	if v, ok := d.GetOk("tagging"); ok || d.HasChange("tagging") {
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

	if v, ok := d.GetOk("fabric_object"); ok {
		t, err := expandFirewallAddressFabricObject(d, v, "fabric_object", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	return &obj, nil
}
