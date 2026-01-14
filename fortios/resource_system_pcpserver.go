// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure PCP server information.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPcpServerUpdate,
		Read:   resourceSystemPcpServerRead,
		Update: resourceSystemPcpServerUpdate,
		Delete: resourceSystemPcpServerDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"client_subnet": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"ext_intf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"arp_reply": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"extport": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"minimal_lifetime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 300),
							Optional:     true,
							Computed:     true,
						},
						"maximal_lifetime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3600, 604800),
							Optional:     true,
							Computed:     true,
						},
						"client_mapping_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"mapping_filter_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 5),
							Optional:     true,
							Computed:     true,
						},
						"allow_opcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"third_party": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"third_party_subnet": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"multicast_announcement": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"announcement_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 10),
							Optional:     true,
							Computed:     true,
						},
						"intl_intf": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"recycle_delay": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
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

func resourceSystemPcpServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemPcpServer(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServer resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPcpServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPcpServer")
	}

	return resourceSystemPcpServerRead(d, m)
}

func resourceSystemPcpServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemPcpServer(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemPcpServer resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPcpServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemPcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPcpServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPcpServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemPcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPcpServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemPcpServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemPcpServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPools(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemPcpServerPoolsName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["description"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
			}
			tmp["description"] = flattenSystemPcpServerPoolsDescription(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenSystemPcpServerPoolsId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["client-subnet"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "client_subnet"
			}
			tmp["client_subnet"] = flattenSystemPcpServerPoolsClientSubnet(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ext-intf"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ext_intf"
			}
			tmp["ext_intf"] = flattenSystemPcpServerPoolsExtIntf(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["arp-reply"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
			}
			tmp["arp_reply"] = flattenSystemPcpServerPoolsArpReply(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["extip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
			}
			tmp["extip"] = flattenSystemPcpServerPoolsExtip(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["extport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
			}
			tmp["extport"] = flattenSystemPcpServerPoolsExtport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["minimal-lifetime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "minimal_lifetime"
			}
			tmp["minimal_lifetime"] = flattenSystemPcpServerPoolsMinimalLifetime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["maximal-lifetime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "maximal_lifetime"
			}
			tmp["maximal_lifetime"] = flattenSystemPcpServerPoolsMaximalLifetime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["client-mapping-limit"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "client_mapping_limit"
			}
			tmp["client_mapping_limit"] = flattenSystemPcpServerPoolsClientMappingLimit(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mapping-filter-limit"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mapping_filter_limit"
			}
			tmp["mapping_filter_limit"] = flattenSystemPcpServerPoolsMappingFilterLimit(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["allow-opcode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_opcode"
			}
			tmp["allow_opcode"] = flattenSystemPcpServerPoolsAllowOpcode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["third-party"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party"
			}
			tmp["third_party"] = flattenSystemPcpServerPoolsThirdParty(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["third-party-subnet"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party_subnet"
			}
			tmp["third_party_subnet"] = flattenSystemPcpServerPoolsThirdPartySubnet(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["multicast-announcement"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_announcement"
			}
			tmp["multicast_announcement"] = flattenSystemPcpServerPoolsMulticastAnnouncement(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["announcement-count"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "announcement_count"
			}
			tmp["announcement_count"] = flattenSystemPcpServerPoolsAnnouncementCount(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["intl-intf"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "intl_intf"
			}
			tmp["intl_intf"] = flattenSystemPcpServerPoolsIntlIntf(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["recycle-delay"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "recycle_delay"
			}
			tmp["recycle_delay"] = flattenSystemPcpServerPoolsRecycleDelay(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemPcpServerPoolsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPcpServerPoolsClientSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "subnet", "subnet")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["subnet"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
			}
			tmp["subnet"] = flattenSystemPcpServerPoolsClientSubnetSubnet(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subnet", d)
	return result
}

func flattenSystemPcpServerPoolsClientSubnetSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsArpReply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsExtport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMinimalLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPcpServerPoolsMaximalLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPcpServerPoolsClientMappingLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPcpServerPoolsMappingFilterLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPcpServerPoolsAllowOpcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsThirdParty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsThirdPartySubnet(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "subnet", "subnet")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["subnet"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
			}
			tmp["subnet"] = flattenSystemPcpServerPoolsThirdPartySubnetSubnet(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subnet", d)
	return result
}

func flattenSystemPcpServerPoolsThirdPartySubnetSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsMulticastAnnouncement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsAnnouncementCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPcpServerPoolsIntlIntf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "interface-name", "interface_name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["interface-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
			}
			tmp["interface_name"] = flattenSystemPcpServerPoolsIntlIntfInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemPcpServerPoolsIntlIntfInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPcpServerPoolsRecycleDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemPcpServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemPcpServerStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("pools", flattenSystemPcpServerPools(o["pools"], d, "pools", sv)); err != nil {
			if !fortiAPIPatch(o["pools"]) {
				return fmt.Errorf("Error reading pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pools"); ok {
			if err = d.Set("pools", flattenSystemPcpServerPools(o["pools"], d, "pools", sv)); err != nil {
				if !fortiAPIPatch(o["pools"]) {
					return fmt.Errorf("Error reading pools: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemPcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemPcpServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPools(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemPcpServerPoolsName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSystemPcpServerPoolsDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemPcpServerPoolsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_subnet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["client-subnet"], _ = expandSystemPcpServerPoolsClientSubnet(d, i["client_subnet"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["client-subnet"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ext_intf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ext-intf"], _ = expandSystemPcpServerPoolsExtIntf(d, i["ext_intf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ext-intf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_reply"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["arp-reply"], _ = expandSystemPcpServerPoolsArpReply(d, i["arp_reply"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["extip"], _ = expandSystemPcpServerPoolsExtip(d, i["extip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["extip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "extport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["extport"], _ = expandSystemPcpServerPoolsExtport(d, i["extport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["extport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimal_lifetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["minimal-lifetime"], _ = expandSystemPcpServerPoolsMinimalLifetime(d, i["minimal_lifetime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximal_lifetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximal-lifetime"], _ = expandSystemPcpServerPoolsMaximalLifetime(d, i["maximal_lifetime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_mapping_limit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["client-mapping-limit"], _ = expandSystemPcpServerPoolsClientMappingLimit(d, i["client_mapping_limit"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["client-mapping-limit"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapping_filter_limit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mapping-filter-limit"], _ = expandSystemPcpServerPoolsMappingFilterLimit(d, i["mapping_filter_limit"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allow_opcode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allow-opcode"], _ = expandSystemPcpServerPoolsAllowOpcode(d, i["allow_opcode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["third-party"], _ = expandSystemPcpServerPoolsThirdParty(d, i["third_party"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "third_party_subnet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["third-party-subnet"], _ = expandSystemPcpServerPoolsThirdPartySubnet(d, i["third_party_subnet"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["third-party-subnet"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_announcement"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["multicast-announcement"], _ = expandSystemPcpServerPoolsMulticastAnnouncement(d, i["multicast_announcement"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "announcement_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["announcement-count"], _ = expandSystemPcpServerPoolsAnnouncementCount(d, i["announcement_count"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "intl_intf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["intl-intf"], _ = expandSystemPcpServerPoolsIntlIntf(d, i["intl_intf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["intl-intf"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recycle_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["recycle-delay"], _ = expandSystemPcpServerPoolsRecycleDelay(d, i["recycle_delay"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["recycle-delay"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemPcpServerPoolsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsClientSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["subnet"], _ = expandSystemPcpServerPoolsClientSubnetSubnet(d, i["subnet"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemPcpServerPoolsClientSubnetSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsArpReply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsExtport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMinimalLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMaximalLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsClientMappingLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMappingFilterLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsAllowOpcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsThirdParty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsThirdPartySubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["subnet"], _ = expandSystemPcpServerPoolsThirdPartySubnetSubnet(d, i["subnet"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemPcpServerPoolsThirdPartySubnetSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsMulticastAnnouncement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsAnnouncementCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsIntlIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["interface-name"], _ = expandSystemPcpServerPoolsIntlIntfInterfaceName(d, i["interface_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemPcpServerPoolsIntlIntfInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPcpServerPoolsRecycleDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPcpServer(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemPcpServerStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("pools"); ok || d.HasChange("pools") {
		if setArgNil {
			obj["pools"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemPcpServerPools(d, v, "pools", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pools"] = t
			}
		}
	}

	return &obj, nil
}
