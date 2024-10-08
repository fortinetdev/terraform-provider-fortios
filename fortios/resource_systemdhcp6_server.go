// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DHCPv6 servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcp6Server() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcp6ServerCreate,
		Read:   resourceSystemDhcp6ServerRead,
		Update: resourceSystemDhcp6ServerUpdate,
		Delete: resourceSystemDhcp6ServerDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rapid_commit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lease_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(300, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_search_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"code": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional:     true,
						},
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"option1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"option2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"option3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upstream_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"delegated_prefix_iaid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix_length": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 128),
							Optional:     true,
						},
					},
				},
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemDhcp6ServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcp6Server(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6Server resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDhcp6Server(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6Server resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDhcp6Server")
	}

	return resourceSystemDhcp6ServerRead(d, m)
}

func resourceSystemDhcp6ServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcp6Server(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6Server resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDhcp6Server(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6Server resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDhcp6Server")
	}

	return resourceSystemDhcp6ServerRead(d, m)
}

func resourceSystemDhcp6ServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemDhcp6Server(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcp6Server resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcp6ServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcp6Server(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6Server resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcp6Server(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6Server resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcp6ServerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerRapidCommit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerLeaseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerDnsService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsSearchList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptions(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcp6ServerOptionsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if cur_v, ok := i["code"]; ok {
			tmp["code"] = flattenSystemDhcp6ServerOptionsCode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenSystemDhcp6ServerOptionsType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenSystemDhcp6ServerOptionsValue(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if cur_v, ok := i["ip6"]; ok {
			tmp["ip6"] = flattenSystemDhcp6ServerOptionsIp6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if cur_v, ok := i["vci-match"]; ok {
			tmp["vci_match"] = flattenSystemDhcp6ServerOptionsVciMatch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if cur_v, ok := i["vci-string"]; ok {
			tmp["vci_string"] = flattenSystemDhcp6ServerOptionsVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcp6ServerOptionsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerOptionsCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerOptionsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOptionsVciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if cur_v, ok := i["vci-string"]; ok {
			tmp["vci_string"] = flattenSystemDhcp6ServerOptionsVciStringVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vci_string", d)
	return result
}

func flattenSystemDhcp6ServerOptionsVciStringVciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerUpstreamInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerIpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcp6ServerPrefixRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_prefix"
		if cur_v, ok := i["start-prefix"]; ok {
			tmp["start_prefix"] = flattenSystemDhcp6ServerPrefixRangeStartPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_prefix"
		if cur_v, ok := i["end-prefix"]; ok {
			tmp["end_prefix"] = flattenSystemDhcp6ServerPrefixRangeEndPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if cur_v, ok := i["prefix-length"]; ok {
			tmp["prefix_length"] = flattenSystemDhcp6ServerPrefixRangePrefixLength(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcp6ServerPrefixRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerPrefixRangeStartPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeEndPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangePrefixLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerIpRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcp6ServerIpRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if cur_v, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenSystemDhcp6ServerIpRangeStartIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if cur_v, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenSystemDhcp6ServerIpRangeEndIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if cur_v, ok := i["vci-match"]; ok {
			tmp["vci_match"] = flattenSystemDhcp6ServerIpRangeVciMatch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if cur_v, ok := i["vci-string"]; ok {
			tmp["vci_string"] = flattenSystemDhcp6ServerIpRangeVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcp6ServerIpRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcp6ServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeVciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if cur_v, ok := i["vci-string"]; ok {
			tmp["vci_string"] = flattenSystemDhcp6ServerIpRangeVciStringVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vci_string", d)
	return result
}

func flattenSystemDhcp6ServerIpRangeVciStringVciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDhcp6Server(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("fosid", flattenSystemDhcp6ServerId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDhcp6ServerStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("rapid_commit", flattenSystemDhcp6ServerRapidCommit(o["rapid-commit"], d, "rapid_commit", sv)); err != nil {
		if !fortiAPIPatch(o["rapid-commit"]) {
			return fmt.Errorf("Error reading rapid_commit: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcp6ServerLeaseTime(o["lease-time"], d, "lease_time", sv)); err != nil {
		if !fortiAPIPatch(o["lease-time"]) {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcp6ServerDnsService(o["dns-service"], d, "dns_service", sv)); err != nil {
		if !fortiAPIPatch(o["dns-service"]) {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("dns_search_list", flattenSystemDhcp6ServerDnsSearchList(o["dns-search-list"], d, "dns_search_list", sv)); err != nil {
		if !fortiAPIPatch(o["dns-search-list"]) {
			return fmt.Errorf("Error reading dns_search_list: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcp6ServerDnsServer1(o["dns-server1"], d, "dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcp6ServerDnsServer2(o["dns-server2"], d, "dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcp6ServerDnsServer3(o["dns-server3"], d, "dns_server3", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server3"]) {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenSystemDhcp6ServerDnsServer4(o["dns-server4"], d, "dns_server4", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server4"]) {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcp6ServerDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("subnet", flattenSystemDhcp6ServerSubnet(o["subnet"], d, "subnet", sv)); err != nil {
		if !fortiAPIPatch(o["subnet"]) {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDhcp6ServerInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("options", flattenSystemDhcp6ServerOptions(o["options"], d, "options", sv)); err != nil {
			if !fortiAPIPatch(o["options"]) {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenSystemDhcp6ServerOptions(o["options"], d, "options", sv)); err != nil {
				if !fortiAPIPatch(o["options"]) {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("option1", flattenSystemDhcp6ServerOption1(o["option1"], d, "option1", sv)); err != nil {
		if !fortiAPIPatch(o["option1"]) {
			return fmt.Errorf("Error reading option1: %v", err)
		}
	}

	if err = d.Set("option2", flattenSystemDhcp6ServerOption2(o["option2"], d, "option2", sv)); err != nil {
		if !fortiAPIPatch(o["option2"]) {
			return fmt.Errorf("Error reading option2: %v", err)
		}
	}

	if err = d.Set("option3", flattenSystemDhcp6ServerOption3(o["option3"], d, "option3", sv)); err != nil {
		if !fortiAPIPatch(o["option3"]) {
			return fmt.Errorf("Error reading option3: %v", err)
		}
	}

	if err = d.Set("upstream_interface", flattenSystemDhcp6ServerUpstreamInterface(o["upstream-interface"], d, "upstream_interface", sv)); err != nil {
		if !fortiAPIPatch(o["upstream-interface"]) {
			return fmt.Errorf("Error reading upstream_interface: %v", err)
		}
	}

	if err = d.Set("delegated_prefix_iaid", flattenSystemDhcp6ServerDelegatedPrefixIaid(o["delegated-prefix-iaid"], d, "delegated_prefix_iaid", sv)); err != nil {
		if !fortiAPIPatch(o["delegated-prefix-iaid"]) {
			return fmt.Errorf("Error reading delegated_prefix_iaid: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenSystemDhcp6ServerIpMode(o["ip-mode"], d, "ip_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ip-mode"]) {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("prefix_mode", flattenSystemDhcp6ServerPrefixMode(o["prefix-mode"], d, "prefix_mode", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-mode"]) {
			return fmt.Errorf("Error reading prefix_mode: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(o["prefix-range"], d, "prefix_range", sv)); err != nil {
			if !fortiAPIPatch(o["prefix-range"]) {
				return fmt.Errorf("Error reading prefix_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("prefix_range"); ok {
			if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(o["prefix-range"], d, "prefix_range", sv)); err != nil {
				if !fortiAPIPatch(o["prefix-range"]) {
					return fmt.Errorf("Error reading prefix_range: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(o["ip-range"], d, "ip_range", sv)); err != nil {
			if !fortiAPIPatch(o["ip-range"]) {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(o["ip-range"], d, "ip_range", sv)); err != nil {
				if !fortiAPIPatch(o["ip-range"]) {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemDhcp6ServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDhcp6ServerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerRapidCommit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerLeaseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsSearchList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcp6ServerOptionsId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["code"], _ = expandSystemDhcp6ServerOptionsCode(d, i["code"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["code"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemDhcp6ServerOptionsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSystemDhcp6ServerOptionsValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandSystemDhcp6ServerOptionsIp6(d, i["ip6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-match"], _ = expandSystemDhcp6ServerOptionsVciMatch(d, i["vci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-string"], _ = expandSystemDhcp6ServerOptionsVciString(d, i["vci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerOptionsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOptionsVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vci-string"], _ = expandSystemDhcp6ServerOptionsVciStringVciString(d, i["vci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerOptionsVciStringVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerUpstreamInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDelegatedPrefixIaid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcp6ServerPrefixRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-prefix"], _ = expandSystemDhcp6ServerPrefixRangeStartPrefix(d, i["start_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-prefix"], _ = expandSystemDhcp6ServerPrefixRangeEndPrefix(d, i["end_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-length"], _ = expandSystemDhcp6ServerPrefixRangePrefixLength(d, i["prefix_length"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-length"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerPrefixRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeStartPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeEndPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangePrefixLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcp6ServerIpRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandSystemDhcp6ServerIpRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandSystemDhcp6ServerIpRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-match"], _ = expandSystemDhcp6ServerIpRangeVciMatch(d, i["vci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-string"], _ = expandSystemDhcp6ServerIpRangeVciString(d, i["vci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerIpRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vci-string"], _ = expandSystemDhcp6ServerIpRangeVciStringVciString(d, i["vci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerIpRangeVciStringVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcp6Server(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemDhcp6ServerId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDhcp6ServerStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("rapid_commit"); ok {
		t, err := expandSystemDhcp6ServerRapidCommit(d, v, "rapid_commit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rapid-commit"] = t
		}
	}

	if v, ok := d.GetOkExists("lease_time"); ok {
		t, err := expandSystemDhcp6ServerLeaseTime(d, v, "lease_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok {
		t, err := expandSystemDhcp6ServerDnsService(d, v, "dns_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("dns_search_list"); ok {
		t, err := expandSystemDhcp6ServerDnsSearchList(d, v, "dns_search_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-search-list"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok {
		t, err := expandSystemDhcp6ServerDnsServer1(d, v, "dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok {
		t, err := expandSystemDhcp6ServerDnsServer2(d, v, "dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok {
		t, err := expandSystemDhcp6ServerDnsServer3(d, v, "dns_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok {
		t, err := expandSystemDhcp6ServerDnsServer4(d, v, "dns_server4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDhcp6ServerDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if d.HasChange("domain") {
		obj["domain"] = nil
	}

	if v, ok := d.GetOk("subnet"); ok {
		t, err := expandSystemDhcp6ServerSubnet(d, v, "subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemDhcp6ServerInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandSystemDhcp6ServerOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("option1"); ok {
		t, err := expandSystemDhcp6ServerOption1(d, v, "option1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option1"] = t
		}
	}

	if v, ok := d.GetOk("option2"); ok {
		t, err := expandSystemDhcp6ServerOption2(d, v, "option2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option2"] = t
		}
	}

	if v, ok := d.GetOk("option3"); ok {
		t, err := expandSystemDhcp6ServerOption3(d, v, "option3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option3"] = t
		}
	}

	if v, ok := d.GetOk("upstream_interface"); ok {
		t, err := expandSystemDhcp6ServerUpstreamInterface(d, v, "upstream_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-interface"] = t
		}
	} else if d.HasChange("upstream_interface") {
		obj["upstream-interface"] = nil
	}

	if v, ok := d.GetOkExists("delegated_prefix_iaid"); ok {
		t, err := expandSystemDhcp6ServerDelegatedPrefixIaid(d, v, "delegated_prefix_iaid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delegated-prefix-iaid"] = t
		}
	} else if d.HasChange("delegated_prefix_iaid") {
		obj["delegated-prefix-iaid"] = nil
	}

	if v, ok := d.GetOk("ip_mode"); ok {
		t, err := expandSystemDhcp6ServerIpMode(d, v, "ip_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("prefix_mode"); ok {
		t, err := expandSystemDhcp6ServerPrefixMode(d, v, "prefix_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-mode"] = t
		}
	}

	if v, ok := d.GetOk("prefix_range"); ok || d.HasChange("prefix_range") {
		t, err := expandSystemDhcp6ServerPrefixRange(d, v, "prefix_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-range"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemDhcp6ServerIpRange(d, v, "ip_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	return &obj, nil
}
