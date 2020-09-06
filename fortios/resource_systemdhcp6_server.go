// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure DHCPv6 servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
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
				ValidateFunc: validation.IntBetween(300, 8640000),
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
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
			},
			"ip_mode": &schema.Schema{
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
							Computed: true,
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
							Computed:     true,
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
					},
				},
			},
		},
	}
}

func resourceSystemDhcp6ServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDhcp6Server(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcp6Server resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDhcp6Server(obj)

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

	obj, err := getObjectSystemDhcp6Server(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcp6Server resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDhcp6Server(obj, mkey)
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

	err := c.DeleteSystemDhcp6Server(mkey)
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

	o, err := c.ReadSystemDhcp6Server(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6Server resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcp6Server(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcp6Server resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcp6ServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerRapidCommit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsSearchList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerOption3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcp6ServerPrefixRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_prefix"
		if _, ok := i["start-prefix"]; ok {
			tmp["start_prefix"] = flattenSystemDhcp6ServerPrefixRangeStartPrefix(i["start-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_prefix"
		if _, ok := i["end-prefix"]; ok {
			tmp["end_prefix"] = flattenSystemDhcp6ServerPrefixRangeEndPrefix(i["end-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := i["prefix-length"]; ok {
			tmp["prefix_length"] = flattenSystemDhcp6ServerPrefixRangePrefixLength(i["prefix-length"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcp6ServerPrefixRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeStartPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangeEndPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerPrefixRangePrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcp6ServerIpRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenSystemDhcp6ServerIpRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenSystemDhcp6ServerIpRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcp6ServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcp6ServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDhcp6Server(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemDhcp6ServerId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDhcp6ServerStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("rapid_commit", flattenSystemDhcp6ServerRapidCommit(o["rapid-commit"], d, "rapid_commit")); err != nil {
		if !fortiAPIPatch(o["rapid-commit"]) {
			return fmt.Errorf("Error reading rapid_commit: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcp6ServerLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if !fortiAPIPatch(o["lease-time"]) {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcp6ServerDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if !fortiAPIPatch(o["dns-service"]) {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("dns_search_list", flattenSystemDhcp6ServerDnsSearchList(o["dns-search-list"], d, "dns_search_list")); err != nil {
		if !fortiAPIPatch(o["dns-search-list"]) {
			return fmt.Errorf("Error reading dns_search_list: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcp6ServerDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcp6ServerDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcp6ServerDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if !fortiAPIPatch(o["dns-server3"]) {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcp6ServerDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("subnet", flattenSystemDhcp6ServerSubnet(o["subnet"], d, "subnet")); err != nil {
		if !fortiAPIPatch(o["subnet"]) {
			return fmt.Errorf("Error reading subnet: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDhcp6ServerInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("option1", flattenSystemDhcp6ServerOption1(o["option1"], d, "option1")); err != nil {
		if !fortiAPIPatch(o["option1"]) {
			return fmt.Errorf("Error reading option1: %v", err)
		}
	}

	if err = d.Set("option2", flattenSystemDhcp6ServerOption2(o["option2"], d, "option2")); err != nil {
		if !fortiAPIPatch(o["option2"]) {
			return fmt.Errorf("Error reading option2: %v", err)
		}
	}

	if err = d.Set("option3", flattenSystemDhcp6ServerOption3(o["option3"], d, "option3")); err != nil {
		if !fortiAPIPatch(o["option3"]) {
			return fmt.Errorf("Error reading option3: %v", err)
		}
	}

	if err = d.Set("upstream_interface", flattenSystemDhcp6ServerUpstreamInterface(o["upstream-interface"], d, "upstream_interface")); err != nil {
		if !fortiAPIPatch(o["upstream-interface"]) {
			return fmt.Errorf("Error reading upstream_interface: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenSystemDhcp6ServerIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if !fortiAPIPatch(o["ip-mode"]) {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(o["prefix-range"], d, "prefix_range")); err != nil {
			if !fortiAPIPatch(o["prefix-range"]) {
				return fmt.Errorf("Error reading prefix_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("prefix_range"); ok {
			if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(o["prefix-range"], d, "prefix_range")); err != nil {
				if !fortiAPIPatch(o["prefix-range"]) {
					return fmt.Errorf("Error reading prefix_range: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
			if !fortiAPIPatch(o["ip-range"]) {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
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
	log.Printf("ER List: %v", e)
}

func expandSystemDhcp6ServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerRapidCommit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsSearchList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerOption3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerUpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcp6ServerPrefixRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-prefix"], _ = expandSystemDhcp6ServerPrefixRangeStartPrefix(d, i["start_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-prefix"], _ = expandSystemDhcp6ServerPrefixRangeEndPrefix(d, i["end_prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-length"], _ = expandSystemDhcp6ServerPrefixRangePrefixLength(d, i["prefix_length"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerPrefixRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeStartPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangeEndPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerPrefixRangePrefixLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcp6ServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandSystemDhcp6ServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandSystemDhcp6ServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcp6ServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcp6ServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcp6Server(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemDhcp6ServerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDhcp6ServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("rapid_commit"); ok {
		t, err := expandSystemDhcp6ServerRapidCommit(d, v, "rapid_commit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rapid-commit"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok {
		t, err := expandSystemDhcp6ServerLeaseTime(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok {
		t, err := expandSystemDhcp6ServerDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("dns_search_list"); ok {
		t, err := expandSystemDhcp6ServerDnsSearchList(d, v, "dns_search_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-search-list"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok {
		t, err := expandSystemDhcp6ServerDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok {
		t, err := expandSystemDhcp6ServerDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok {
		t, err := expandSystemDhcp6ServerDnsServer3(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDhcp6ServerDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("subnet"); ok {
		t, err := expandSystemDhcp6ServerSubnet(d, v, "subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subnet"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemDhcp6ServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("option1"); ok {
		t, err := expandSystemDhcp6ServerOption1(d, v, "option1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option1"] = t
		}
	}

	if v, ok := d.GetOk("option2"); ok {
		t, err := expandSystemDhcp6ServerOption2(d, v, "option2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option2"] = t
		}
	}

	if v, ok := d.GetOk("option3"); ok {
		t, err := expandSystemDhcp6ServerOption3(d, v, "option3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["option3"] = t
		}
	}

	if v, ok := d.GetOk("upstream_interface"); ok {
		t, err := expandSystemDhcp6ServerUpstreamInterface(d, v, "upstream_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok {
		t, err := expandSystemDhcp6ServerIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("prefix_range"); ok {
		t, err := expandSystemDhcp6ServerPrefixRange(d, v, "prefix_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-range"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok {
		t, err := expandSystemDhcp6ServerIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	return &obj, nil
}
