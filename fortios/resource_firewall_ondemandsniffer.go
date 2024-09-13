// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure on-demand packet sniffer.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallOnDemandSniffer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallOnDemandSnifferCreate,
		Read:   resourceFirewallOnDemandSnifferRead,
		Update: resourceFirewallOnDemandSnifferUpdate,
		Delete: resourceFirewallOnDemandSnifferDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"max_packet_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20000),
				Optional:     true,
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65536),
							Optional:     true,
						},
					},
				},
			},
			"protocols": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"non_ip_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"advanced_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
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

func resourceFirewallOnDemandSnifferCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallOnDemandSniffer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallOnDemandSniffer resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallOnDemandSniffer(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallOnDemandSniffer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallOnDemandSniffer")
	}

	return resourceFirewallOnDemandSnifferRead(d, m)
}

func resourceFirewallOnDemandSnifferUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallOnDemandSniffer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallOnDemandSniffer resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallOnDemandSniffer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallOnDemandSniffer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallOnDemandSniffer")
	}

	return resourceFirewallOnDemandSnifferRead(d, m)
}

func resourceFirewallOnDemandSnifferDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallOnDemandSniffer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallOnDemandSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallOnDemandSnifferRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallOnDemandSniffer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallOnDemandSniffer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallOnDemandSniffer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallOnDemandSniffer resource from API: %v", err)
	}
	return nil
}

func flattenFirewallOnDemandSnifferName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferMaxPacketCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallOnDemandSnifferHosts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if cur_v, ok := i["host"]; ok {
			tmp["host"] = flattenFirewallOnDemandSnifferHostsHost(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "host", d)
	return result
}

func flattenFirewallOnDemandSnifferHostsHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenFirewallOnDemandSnifferPortsPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "port", d)
	return result
}

func flattenFirewallOnDemandSnifferPortsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallOnDemandSnifferProtocols(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if cur_v, ok := i["protocol"]; ok {
			tmp["protocol"] = flattenFirewallOnDemandSnifferProtocolsProtocol(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "protocol", d)
	return result
}

func flattenFirewallOnDemandSnifferProtocolsProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallOnDemandSnifferNonIpPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallOnDemandSnifferAdvancedFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallOnDemandSniffer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenFirewallOnDemandSnifferName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallOnDemandSnifferInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("max_packet_count", flattenFirewallOnDemandSnifferMaxPacketCount(o["max-packet-count"], d, "max_packet_count", sv)); err != nil {
		if !fortiAPIPatch(o["max-packet-count"]) {
			return fmt.Errorf("Error reading max_packet_count: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("hosts", flattenFirewallOnDemandSnifferHosts(o["hosts"], d, "hosts", sv)); err != nil {
			if !fortiAPIPatch(o["hosts"]) {
				return fmt.Errorf("Error reading hosts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("hosts"); ok {
			if err = d.Set("hosts", flattenFirewallOnDemandSnifferHosts(o["hosts"], d, "hosts", sv)); err != nil {
				if !fortiAPIPatch(o["hosts"]) {
					return fmt.Errorf("Error reading hosts: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ports", flattenFirewallOnDemandSnifferPorts(o["ports"], d, "ports", sv)); err != nil {
			if !fortiAPIPatch(o["ports"]) {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ports"); ok {
			if err = d.Set("ports", flattenFirewallOnDemandSnifferPorts(o["ports"], d, "ports", sv)); err != nil {
				if !fortiAPIPatch(o["ports"]) {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("protocols", flattenFirewallOnDemandSnifferProtocols(o["protocols"], d, "protocols", sv)); err != nil {
			if !fortiAPIPatch(o["protocols"]) {
				return fmt.Errorf("Error reading protocols: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("protocols"); ok {
			if err = d.Set("protocols", flattenFirewallOnDemandSnifferProtocols(o["protocols"], d, "protocols", sv)); err != nil {
				if !fortiAPIPatch(o["protocols"]) {
					return fmt.Errorf("Error reading protocols: %v", err)
				}
			}
		}
	}

	if err = d.Set("non_ip_packet", flattenFirewallOnDemandSnifferNonIpPacket(o["non-ip-packet"], d, "non_ip_packet", sv)); err != nil {
		if !fortiAPIPatch(o["non-ip-packet"]) {
			return fmt.Errorf("Error reading non_ip_packet: %v", err)
		}
	}

	if err = d.Set("advanced_filter", flattenFirewallOnDemandSnifferAdvancedFilter(o["advanced-filter"], d, "advanced_filter", sv)); err != nil {
		if !fortiAPIPatch(o["advanced-filter"]) {
			return fmt.Errorf("Error reading advanced_filter: %v", err)
		}
	}

	return nil
}

func flattenFirewallOnDemandSnifferFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallOnDemandSnifferName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferMaxPacketCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["host"], _ = expandFirewallOnDemandSnifferHostsHost(d, i["host"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallOnDemandSnifferHostsHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["port"], _ = expandFirewallOnDemandSnifferPortsPort(d, i["port"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallOnDemandSnifferPortsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferProtocols(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["protocol"], _ = expandFirewallOnDemandSnifferProtocolsProtocol(d, i["protocol"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallOnDemandSnifferProtocolsProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferNonIpPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallOnDemandSnifferAdvancedFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallOnDemandSniffer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallOnDemandSnifferName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandFirewallOnDemandSnifferInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("max_packet_count"); ok {
		t, err := expandFirewallOnDemandSnifferMaxPacketCount(d, v, "max_packet_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-packet-count"] = t
		}
	} else if d.HasChange("max_packet_count") {
		obj["max-packet-count"] = nil
	}

	if v, ok := d.GetOk("hosts"); ok || d.HasChange("hosts") {
		t, err := expandFirewallOnDemandSnifferHosts(d, v, "hosts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hosts"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandFirewallOnDemandSnifferPorts(d, v, "ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("protocols"); ok || d.HasChange("protocols") {
		t, err := expandFirewallOnDemandSnifferProtocols(d, v, "protocols", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocols"] = t
		}
	}

	if v, ok := d.GetOk("non_ip_packet"); ok {
		t, err := expandFirewallOnDemandSnifferNonIpPacket(d, v, "non_ip_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["non-ip-packet"] = t
		}
	}

	if v, ok := d.GetOk("advanced_filter"); ok {
		t, err := expandFirewallOnDemandSnifferAdvancedFilter(d, v, "advanced_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advanced-filter"] = t
		}
	} else if d.HasChange("advanced_filter") {
		obj["advanced-filter"] = nil
	}

	return &obj, nil
}
