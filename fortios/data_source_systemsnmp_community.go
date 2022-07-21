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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSnmpCommunityRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ha_direct": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hosts6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"source_ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ha_direct": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"query_v1_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"query_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"query_v2c_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_v1_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_v1_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_v2c_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_v2c_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mib_view": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdoms": &schema.Schema{
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
	}
}

func dataSourceSystemSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("fosid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemSnmpCommunity: type error")
	}

	o, err := c.ReadSystemSnmpCommunity(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpCommunity: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSnmpCommunity(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpCommunity from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemSnmpCommunityHostsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			tmp["source_ip"] = dataSourceFlattenSystemSnmpCommunityHostsSourceIp(i["source-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemSnmpCommunityHostsIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := i["ha-direct"]; ok {
			tmp["ha_direct"] = dataSourceFlattenSystemSnmpCommunityHostsHaDirect(i["ha-direct"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := i["host-type"]; ok {
			tmp["host_type"] = dataSourceFlattenSystemSnmpCommunityHostsHostType(i["host-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHostsHostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemSnmpCommunityHosts6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ipv6"
		if _, ok := i["source-ipv6"]; ok {
			tmp["source_ipv6"] = dataSourceFlattenSystemSnmpCommunityHosts6SourceIpv6(i["source-ipv6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := i["ipv6"]; ok {
			tmp["ipv6"] = dataSourceFlattenSystemSnmpCommunityHosts6Ipv6(i["ipv6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_direct"
		if _, ok := i["ha-direct"]; ok {
			tmp["ha_direct"] = dataSourceFlattenSystemSnmpCommunityHosts6HaDirect(i["ha-direct"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host_type"
		if _, ok := i["host-type"]; ok {
			tmp["host_type"] = dataSourceFlattenSystemSnmpCommunityHosts6HostType(i["host-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSnmpCommunityHosts6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6SourceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6Ipv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6HaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityHosts6HostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityMibView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpCommunityVdoms(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSnmpCommunityVdomsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSnmpCommunityVdomsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSnmpCommunity(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", dataSourceFlattenSystemSnmpCommunityId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenSystemSnmpCommunityName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemSnmpCommunityStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("hosts", dataSourceFlattenSystemSnmpCommunityHosts(o["hosts"], d, "hosts")); err != nil {
		if !fortiAPIPatch(o["hosts"]) {
			return fmt.Errorf("Error reading hosts: %v", err)
		}
	}

	if err = d.Set("hosts6", dataSourceFlattenSystemSnmpCommunityHosts6(o["hosts6"], d, "hosts6")); err != nil {
		if !fortiAPIPatch(o["hosts6"]) {
			return fmt.Errorf("Error reading hosts6: %v", err)
		}
	}

	if err = d.Set("query_v1_status", dataSourceFlattenSystemSnmpCommunityQueryV1Status(o["query-v1-status"], d, "query_v1_status")); err != nil {
		if !fortiAPIPatch(o["query-v1-status"]) {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}

	if err = d.Set("query_v1_port", dataSourceFlattenSystemSnmpCommunityQueryV1Port(o["query-v1-port"], d, "query_v1_port")); err != nil {
		if !fortiAPIPatch(o["query-v1-port"]) {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}

	if err = d.Set("query_v2c_status", dataSourceFlattenSystemSnmpCommunityQueryV2CStatus(o["query-v2c-status"], d, "query_v2c_status")); err != nil {
		if !fortiAPIPatch(o["query-v2c-status"]) {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}

	if err = d.Set("query_v2c_port", dataSourceFlattenSystemSnmpCommunityQueryV2CPort(o["query-v2c-port"], d, "query_v2c_port")); err != nil {
		if !fortiAPIPatch(o["query-v2c-port"]) {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}

	if err = d.Set("trap_v1_status", dataSourceFlattenSystemSnmpCommunityTrapV1Status(o["trap-v1-status"], d, "trap_v1_status")); err != nil {
		if !fortiAPIPatch(o["trap-v1-status"]) {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}

	if err = d.Set("trap_v1_lport", dataSourceFlattenSystemSnmpCommunityTrapV1Lport(o["trap-v1-lport"], d, "trap_v1_lport")); err != nil {
		if !fortiAPIPatch(o["trap-v1-lport"]) {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}

	if err = d.Set("trap_v1_rport", dataSourceFlattenSystemSnmpCommunityTrapV1Rport(o["trap-v1-rport"], d, "trap_v1_rport")); err != nil {
		if !fortiAPIPatch(o["trap-v1-rport"]) {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_status", dataSourceFlattenSystemSnmpCommunityTrapV2CStatus(o["trap-v2c-status"], d, "trap_v2c_status")); err != nil {
		if !fortiAPIPatch(o["trap-v2c-status"]) {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}

	if err = d.Set("trap_v2c_lport", dataSourceFlattenSystemSnmpCommunityTrapV2CLport(o["trap-v2c-lport"], d, "trap_v2c_lport")); err != nil {
		if !fortiAPIPatch(o["trap-v2c-lport"]) {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}

	if err = d.Set("trap_v2c_rport", dataSourceFlattenSystemSnmpCommunityTrapV2CRport(o["trap-v2c-rport"], d, "trap_v2c_rport")); err != nil {
		if !fortiAPIPatch(o["trap-v2c-rport"]) {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}

	if err = d.Set("events", dataSourceFlattenSystemSnmpCommunityEvents(o["events"], d, "events")); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("mib_view", dataSourceFlattenSystemSnmpCommunityMibView(o["mib-view"], d, "mib_view")); err != nil {
		if !fortiAPIPatch(o["mib-view"]) {
			return fmt.Errorf("Error reading mib_view: %v", err)
		}
	}

	if err = d.Set("vdoms", dataSourceFlattenSystemSnmpCommunityVdoms(o["vdoms"], d, "vdoms")); err != nil {
		if !fortiAPIPatch(o["vdoms"]) {
			return fmt.Errorf("Error reading vdoms: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSnmpCommunityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
