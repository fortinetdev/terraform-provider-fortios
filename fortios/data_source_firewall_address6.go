// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 firewall addresses.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceFirewallAddress6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallAddress6Read,
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
			"route_tag": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"macaddr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"start_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cache_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_segment": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"host_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"epg_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_tag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallAddress6Read(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallAddress6: type error")
	}

	o, err := c.ReadFirewallAddress6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallAddress6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallAddress6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallAddress6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6RouteTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Macaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "macaddr"
		if _, ok := i["macaddr"]; ok {
			tmp["macaddr"] = dataSourceFlattenFirewallAddress6MacaddrMacaddr(i["macaddr"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6MacaddrMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6StartMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6EndMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Sdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Ip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Wildcard(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Fqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Country(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6CacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Visibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6ObjId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6List(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["ip"] = dataSourceFlattenFirewallAddress6ListIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6ListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Tagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallAddress6TaggingName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = dataSourceFlattenFirewallAddress6TaggingCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			tmp["tags"] = dataSourceFlattenFirewallAddress6TaggingTags(i["tags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6TaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6TaggingTags(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallAddress6TaggingTagsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6TaggingTagsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Template(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6SubnetSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallAddress6SubnetSegmentName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenFirewallAddress6SubnetSegmentType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenFirewallAddress6SubnetSegmentValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallAddress6SubnetSegmentName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6SubnetSegmentType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6SubnetSegmentValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6HostType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Host(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Tenant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6EpgName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6SdnTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6Filter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6SdnAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallAddress6FabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallAddress6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallAddress6Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenFirewallAddress6Uuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenFirewallAddress6Type(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("route_tag", dataSourceFlattenFirewallAddress6RouteTag(o["route-tag"], d, "route_tag")); err != nil {
		if !fortiAPIPatch(o["route-tag"]) {
			return fmt.Errorf("Error reading route_tag: %v", err)
		}
	}

	if err = d.Set("macaddr", dataSourceFlattenFirewallAddress6Macaddr(o["macaddr"], d, "macaddr")); err != nil {
		if !fortiAPIPatch(o["macaddr"]) {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("start_mac", dataSourceFlattenFirewallAddress6StartMac(o["start-mac"], d, "start_mac")); err != nil {
		if !fortiAPIPatch(o["start-mac"]) {
			return fmt.Errorf("Error reading start_mac: %v", err)
		}
	}

	if err = d.Set("end_mac", dataSourceFlattenFirewallAddress6EndMac(o["end-mac"], d, "end_mac")); err != nil {
		if !fortiAPIPatch(o["end-mac"]) {
			return fmt.Errorf("Error reading end_mac: %v", err)
		}
	}

	if err = d.Set("sdn", dataSourceFlattenFirewallAddress6Sdn(o["sdn"], d, "sdn")); err != nil {
		if !fortiAPIPatch(o["sdn"]) {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}

	if err = d.Set("ip6", dataSourceFlattenFirewallAddress6Ip6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("wildcard", dataSourceFlattenFirewallAddress6Wildcard(o["wildcard"], d, "wildcard")); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}

	if err = d.Set("start_ip", dataSourceFlattenFirewallAddress6StartIp(o["start-ip"], d, "start_ip")); err != nil {
		if !fortiAPIPatch(o["start-ip"]) {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("end_ip", dataSourceFlattenFirewallAddress6EndIp(o["end-ip"], d, "end_ip")); err != nil {
		if !fortiAPIPatch(o["end-ip"]) {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("fqdn", dataSourceFlattenFirewallAddress6Fqdn(o["fqdn"], d, "fqdn")); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("country", dataSourceFlattenFirewallAddress6Country(o["country"], d, "country")); err != nil {
		if !fortiAPIPatch(o["country"]) {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("cache_ttl", dataSourceFlattenFirewallAddress6CacheTtl(o["cache-ttl"], d, "cache_ttl")); err != nil {
		if !fortiAPIPatch(o["cache-ttl"]) {
			return fmt.Errorf("Error reading cache_ttl: %v", err)
		}
	}

	if err = d.Set("visibility", dataSourceFlattenFirewallAddress6Visibility(o["visibility"], d, "visibility")); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallAddress6Color(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("obj_id", dataSourceFlattenFirewallAddress6ObjId(o["obj-id"], d, "obj_id")); err != nil {
		if !fortiAPIPatch(o["obj-id"]) {
			return fmt.Errorf("Error reading obj_id: %v", err)
		}
	}

	if err = d.Set("list", dataSourceFlattenFirewallAddress6List(o["list"], d, "list")); err != nil {
		if !fortiAPIPatch(o["list"]) {
			return fmt.Errorf("Error reading list: %v", err)
		}
	}

	if err = d.Set("tagging", dataSourceFlattenFirewallAddress6Tagging(o["tagging"], d, "tagging")); err != nil {
		if !fortiAPIPatch(o["tagging"]) {
			return fmt.Errorf("Error reading tagging: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenFirewallAddress6Comment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("template", dataSourceFlattenFirewallAddress6Template(o["template"], d, "template")); err != nil {
		if !fortiAPIPatch(o["template"]) {
			return fmt.Errorf("Error reading template: %v", err)
		}
	}

	if err = d.Set("subnet_segment", dataSourceFlattenFirewallAddress6SubnetSegment(o["subnet-segment"], d, "subnet_segment")); err != nil {
		if !fortiAPIPatch(o["subnet-segment"]) {
			return fmt.Errorf("Error reading subnet_segment: %v", err)
		}
	}

	if err = d.Set("host_type", dataSourceFlattenFirewallAddress6HostType(o["host-type"], d, "host_type")); err != nil {
		if !fortiAPIPatch(o["host-type"]) {
			return fmt.Errorf("Error reading host_type: %v", err)
		}
	}

	if err = d.Set("host", dataSourceFlattenFirewallAddress6Host(o["host"], d, "host")); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("tenant", dataSourceFlattenFirewallAddress6Tenant(o["tenant"], d, "tenant")); err != nil {
		if !fortiAPIPatch(o["tenant"]) {
			return fmt.Errorf("Error reading tenant: %v", err)
		}
	}

	if err = d.Set("epg_name", dataSourceFlattenFirewallAddress6EpgName(o["epg-name"], d, "epg_name")); err != nil {
		if !fortiAPIPatch(o["epg-name"]) {
			return fmt.Errorf("Error reading epg_name: %v", err)
		}
	}

	if err = d.Set("sdn_tag", dataSourceFlattenFirewallAddress6SdnTag(o["sdn-tag"], d, "sdn_tag")); err != nil {
		if !fortiAPIPatch(o["sdn-tag"]) {
			return fmt.Errorf("Error reading sdn_tag: %v", err)
		}
	}

	if err = d.Set("filter", dataSourceFlattenFirewallAddress6Filter(o["filter"], d, "filter")); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("sdn_addr_type", dataSourceFlattenFirewallAddress6SdnAddrType(o["sdn-addr-type"], d, "sdn_addr_type")); err != nil {
		if !fortiAPIPatch(o["sdn-addr-type"]) {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}

	if err = d.Set("fabric_object", dataSourceFlattenFirewallAddress6FabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallAddress6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
