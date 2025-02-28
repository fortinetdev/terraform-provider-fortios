// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SNMP community configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpCommunityCreate,
		Read:   resourceSystemSnmpCommunityRead,
		Update: resourceSystemSnmpCommunityUpdate,
		Delete: resourceSystemSnmpCommunityDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ha_direct": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"vrf_select": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 511),
							Optional:     true,
						},
					},
				},
			},
			"hosts6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"source_ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ha_direct": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"vrf_select": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 511),
							Optional:     true,
						},
					},
				},
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_v1_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"query_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_v2c_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v1_lport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"trap_v1_rport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"trap_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_v2c_lport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"trap_v2c_rport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mib_view": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
			},
			"vdoms": &schema.Schema{
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

func resourceSystemSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpCommunity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSnmpCommunity(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemSnmpCommunity")
	}

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpCommunity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnmpCommunity(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemSnmpCommunity")
	}

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSnmpCommunity(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSnmpCommunity(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunity(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemSnmpCommunityHostsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if cur_v, ok := i["source-ip"]; ok {
			tmp["source_ip"] = flattenSystemSnmpCommunityHostsSourceIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemSnmpCommunityHostsIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if cur_v, ok := i["ha-direct"]; ok {
			tmp["ha_direct"] = flattenSystemSnmpCommunityHostsHaDirect(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if cur_v, ok := i["host-type"]; ok {
			tmp["host_type"] = flattenSystemSnmpCommunityHostsHostType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if cur_v, ok := i["interface-select-method"]; ok {
			tmp["interface_select_method"] = flattenSystemSnmpCommunityHostsInterfaceSelectMethod(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenSystemSnmpCommunityHostsInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if cur_v, ok := i["vrf-select"]; ok {
			tmp["vrf_select"] = flattenSystemSnmpCommunityHostsVrfSelect(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityHostsSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemSnmpCommunityHostsHaDirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsHostType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHostsVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityHosts6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemSnmpCommunityHosts6Id(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ipv6"
		if cur_v, ok := i["source-ipv6"]; ok {
			tmp["source_ipv6"] = flattenSystemSnmpCommunityHosts6SourceIpv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if cur_v, ok := i["ipv6"]; ok {
			tmp["ipv6"] = flattenSystemSnmpCommunityHosts6Ipv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if cur_v, ok := i["ha-direct"]; ok {
			tmp["ha_direct"] = flattenSystemSnmpCommunityHosts6HaDirect(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if cur_v, ok := i["host-type"]; ok {
			tmp["host_type"] = flattenSystemSnmpCommunityHosts6HostType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if cur_v, ok := i["interface-select-method"]; ok {
			tmp["interface_select_method"] = flattenSystemSnmpCommunityHosts6InterfaceSelectMethod(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenSystemSnmpCommunityHosts6Interface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if cur_v, ok := i["vrf-select"]; ok {
			tmp["vrf_select"] = flattenSystemSnmpCommunityHosts6VrfSelect(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSnmpCommunityHosts6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityHosts6SourceIpv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Ipv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6HaDirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6HostType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6InterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6Interface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityHosts6VrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityMibView(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpCommunityVdoms(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSnmpCommunityVdomsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSnmpCommunityVdomsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpCommunity(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("fosid", flattenSystemSnmpCommunityId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSystemSnmpCommunityName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSnmpCommunityStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts", sv)); err != nil {
			if !fortiAPIPatch(o["hosts"]) {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts", sv)); err != nil {
				if !fortiAPIPatch(o["hosts"]) {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6", sv)); err != nil {
			if !fortiAPIPatch(o["hosts6"]) {
				return fmt.Errorf("Error reading hosts6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts6"); ok {
			if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6", sv)); err != nil {
				if !fortiAPIPatch(o["hosts6"]) {
					return fmt.Errorf("Error reading hosts6: %v", err)
				}
			}
		}
	}

	if err = d.Set("query_v1_status", flattenSystemSnmpCommunityQueryV1Status(o["query-v1-status"], d, "query_v1_status", sv)); err != nil {
		if !fortiAPIPatch(o["query-v1-status"]) {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v1_port", flattenSystemSnmpCommunityQueryV1Port(o["query-v1-port"], d, "query_v1_port", sv)); err != nil {
		if !fortiAPIPatch(o["query-v1-port"]) {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", flattenSystemSnmpCommunityQueryV2CStatus(o["query-v2c-status"], d, "query_v2c_status", sv)); err != nil {
		if !fortiAPIPatch(o["query-v2c-status"]) {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", flattenSystemSnmpCommunityQueryV2CPort(o["query-v2c-port"], d, "query_v2c_port", sv)); err != nil {
		if !fortiAPIPatch(o["query-v2c-port"]) {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", flattenSystemSnmpCommunityTrapV1Status(o["trap-v1-status"], d, "trap_v1_status", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v1-status"]) {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v1_lport", flattenSystemSnmpCommunityTrapV1Lport(o["trap-v1-lport"], d, "trap_v1_lport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v1-lport"]) {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", flattenSystemSnmpCommunityTrapV1Rport(o["trap-v1-rport"], d, "trap_v1_rport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v1-rport"]) {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", flattenSystemSnmpCommunityTrapV2CStatus(o["trap-v2c-status"], d, "trap_v2c_status", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v2c-status"]) {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_lport", flattenSystemSnmpCommunityTrapV2CLport(o["trap-v2c-lport"], d, "trap_v2c_lport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v2c-lport"]) {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", flattenSystemSnmpCommunityTrapV2CRport(o["trap-v2c-rport"], d, "trap_v2c_rport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v2c-rport"]) {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("events", flattenSystemSnmpCommunityEvents(o["events"], d, "events", sv)); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("mib_view", flattenSystemSnmpCommunityMibView(o["mib-view"], d, "mib_view", sv)); err != nil {
		if !fortiAPIPatch(o["mib-view"]) {
			return fmt.Errorf("Error reading mib_view: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vdoms", flattenSystemSnmpCommunityVdoms(o["vdoms"], d, "vdoms", sv)); err != nil {
			if !fortiAPIPatch(o["vdoms"]) {
				return fmt.Errorf("Error reading vdoms: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdoms"); ok {
			if err = d.Set("vdoms", flattenSystemSnmpCommunityVdoms(o["vdoms"], d, "vdoms", sv)); err != nil {
				if !fortiAPIPatch(o["vdoms"]) {
					return fmt.Errorf("Error reading vdoms: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSnmpCommunityHostsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip"], _ = expandSystemSnmpCommunityHostsSourceIp(d, i["source_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemSnmpCommunityHostsIp(d, i["ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-direct"], _ = expandSystemSnmpCommunityHostsHaDirect(d, i["ha_direct"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["host-type"], _ = expandSystemSnmpCommunityHostsHostType(d, i["host_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandSystemSnmpCommunityHostsInterfaceSelectMethod(d, i["interface_select_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemSnmpCommunityHostsInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf-select"], _ = expandSystemSnmpCommunityHostsVrfSelect(d, i["vrf_select"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf-select"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsHaDirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsHostType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHostsVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSnmpCommunityHosts6Id(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ipv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ipv6"], _ = expandSystemSnmpCommunityHosts6SourceIpv6(d, i["source_ipv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv6"], _ = expandSystemSnmpCommunityHosts6Ipv6(d, i["ipv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-direct"], _ = expandSystemSnmpCommunityHosts6HaDirect(d, i["ha_direct"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["host-type"], _ = expandSystemSnmpCommunityHosts6HostType(d, i["host_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandSystemSnmpCommunityHosts6InterfaceSelectMethod(d, i["interface_select_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemSnmpCommunityHosts6Interface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf_select"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf-select"], _ = expandSystemSnmpCommunityHosts6VrfSelect(d, i["vrf_select"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf-select"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityHosts6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6SourceIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Ipv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6HaDirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6HostType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6InterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityHosts6VrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV1Port(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityQueryV2CPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Lport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV1Rport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CLport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityTrapV2CRport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityEvents(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityMibView(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpCommunityVdoms(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemSnmpCommunityVdomsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSnmpCommunityVdomsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpCommunity(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemSnmpCommunityId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSnmpCommunityName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpCommunityStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandSystemSnmpCommunityHosts(d, v, "hosts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("hosts6"); ok || d.HasChange("hosts6") {
		t, err := expandSystemSnmpCommunityHosts6(d, v, "hosts6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts6"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_status"); ok {
		t, err := expandSystemSnmpCommunityQueryV1Status(d, v, "query_v1_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("query_v1_port"); ok {
		t, err := expandSystemSnmpCommunityQueryV1Port(d, v, "query_v1_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-port"] = t
		}
	}

	if v, ok := d.GetOk("query_v2c_status"); ok {
		t, err := expandSystemSnmpCommunityQueryV2CStatus(d, v, "query_v2c_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-status"] = t
		}
	}

	if v, ok := d.GetOkExists("query_v2c_port"); ok {
		t, err := expandSystemSnmpCommunityQueryV2CPort(d, v, "query_v2c_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-port"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_status"); ok {
		t, err := expandSystemSnmpCommunityTrapV1Status(d, v, "trap_v1_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_lport"); ok {
		t, err := expandSystemSnmpCommunityTrapV1Lport(d, v, "trap_v1_lport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v1_rport"); ok {
		t, err := expandSystemSnmpCommunityTrapV1Rport(d, v, "trap_v1_rport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-rport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_status"); ok {
		t, err := expandSystemSnmpCommunityTrapV2CStatus(d, v, "trap_v2c_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_lport"); ok {
		t, err := expandSystemSnmpCommunityTrapV2CLport(d, v, "trap_v2c_lport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_v2c_rport"); ok {
		t, err := expandSystemSnmpCommunityTrapV2CRport(d, v, "trap_v2c_rport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-rport"] = t
		}
	}

	if v, ok := d.GetOk("events"); ok {
		t, err := expandSystemSnmpCommunityEvents(d, v, "events", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("mib_view"); ok {
		t, err := expandSystemSnmpCommunityMibView(d, v, "mib_view", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mib-view"] = t
		}
	} else if d.HasChange("mib_view") {
		obj["mib-view"] = nil
	}

	if v, ok := d.GetOk("vdoms"); ok || d.HasChange("vdoms") {
		t, err := expandSystemSnmpCommunityVdoms(d, v, "vdoms", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdoms"] = t
		}
	}

	return &obj, nil
}
