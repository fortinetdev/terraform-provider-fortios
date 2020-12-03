// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 to IPv6 virtual IPs.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallVip46() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVip46Create,
		Read:   resourceFirewallVip46Read,
		Update: resourceFirewallVip46Update,
		Delete: resourceFirewallVip46Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realservers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"holddown_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(30, 65535),
							Optional:     true,
							Computed:     true,
						},
						"healthcheck": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"monitor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"client_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallVip46Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallVip46(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip46 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallVip46(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallVip46 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVip46")
	}

	return resourceFirewallVip46Read(d, m)
}

func resourceFirewallVip46Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallVip46(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip46 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallVip46(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip46 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVip46")
	}

	return resourceFirewallVip46Read(d, m)
}

func resourceFirewallVip46Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallVip46(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip46 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip46Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallVip46(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip46 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip46(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip46 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVip46Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46SrcFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = flattenFirewallVip46SrcFilterRange(i["range"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallVip46SrcFilterRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Extip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Mappedip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46ArpReply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Portforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Extport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Mappedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46LdbMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46ServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Realservers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallVip46RealserversId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenFirewallVip46RealserversIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			tmp["port"] = flattenFirewallVip46RealserversPort(i["port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenFirewallVip46RealserversStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			tmp["weight"] = flattenFirewallVip46RealserversWeight(i["weight"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {
			tmp["holddown_interval"] = flattenFirewallVip46RealserversHolddownInterval(i["holddown-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {
			tmp["healthcheck"] = flattenFirewallVip46RealserversHealthcheck(i["healthcheck"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {
			tmp["max_connections"] = flattenFirewallVip46RealserversMaxConnections(i["max-connections"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			tmp["monitor"] = flattenFirewallVip46RealserversMonitor(i["monitor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {
			tmp["client_ip"] = flattenFirewallVip46RealserversClientIp(i["client-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallVip46RealserversId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46RealserversClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallVip46Monitor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallVip46MonitorName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallVip46MonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallVip46(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallVip46Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVip46Id(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVip46Uuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallVip46Comment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallVip46Type(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_filter", flattenFirewallVip46SrcFilter(o["src-filter"], d, "src_filter")); err != nil {
			if !fortiAPIPatch(o["src-filter"]) {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_filter"); ok {
			if err = d.Set("src_filter", flattenFirewallVip46SrcFilter(o["src-filter"], d, "src_filter")); err != nil {
				if !fortiAPIPatch(o["src-filter"]) {
					return fmt.Errorf("Error reading src_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("extip", flattenFirewallVip46Extip(o["extip"], d, "extip")); err != nil {
		if !fortiAPIPatch(o["extip"]) {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenFirewallVip46Mappedip(o["mappedip"], d, "mappedip")); err != nil {
		if !fortiAPIPatch(o["mappedip"]) {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallVip46ArpReply(o["arp-reply"], d, "arp_reply")); err != nil {
		if !fortiAPIPatch(o["arp-reply"]) {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("portforward", flattenFirewallVip46Portforward(o["portforward"], d, "portforward")); err != nil {
		if !fortiAPIPatch(o["portforward"]) {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallVip46Protocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("extport", flattenFirewallVip46Extport(o["extport"], d, "extport")); err != nil {
		if !fortiAPIPatch(o["extport"]) {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallVip46Mappedport(o["mappedport"], d, "mappedport")); err != nil {
		if !fortiAPIPatch(o["mappedport"]) {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallVip46Color(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallVip46LdbMethod(o["ldb-method"], d, "ldb_method")); err != nil {
		if !fortiAPIPatch(o["ldb-method"]) {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("server_type", flattenFirewallVip46ServerType(o["server-type"], d, "server_type")); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenFirewallVip46Realservers(o["realservers"], d, "realservers")); err != nil {
			if !fortiAPIPatch(o["realservers"]) {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallVip46Realservers(o["realservers"], d, "realservers")); err != nil {
				if !fortiAPIPatch(o["realservers"]) {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("monitor", flattenFirewallVip46Monitor(o["monitor"], d, "monitor")); err != nil {
			if !fortiAPIPatch(o["monitor"]) {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitor"); ok {
			if err = d.Set("monitor", flattenFirewallVip46Monitor(o["monitor"], d, "monitor")); err != nil {
				if !fortiAPIPatch(o["monitor"]) {
					return fmt.Errorf("Error reading monitor: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallVip46FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallVip46Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46SrcFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["range"], _ = expandFirewallVip46SrcFilterRange(d, i["range"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip46SrcFilterRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Extip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Mappedip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46ArpReply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Portforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Extport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Mappedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46LdbMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46ServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Realservers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallVip46RealserversId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFirewallVip46RealserversIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandFirewallVip46RealserversPort(d, i["port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandFirewallVip46RealserversStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandFirewallVip46RealserversWeight(d, i["weight"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holddown-interval"], _ = expandFirewallVip46RealserversHolddownInterval(d, i["holddown_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["healthcheck"], _ = expandFirewallVip46RealserversHealthcheck(d, i["healthcheck"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-connections"], _ = expandFirewallVip46RealserversMaxConnections(d, i["max_connections"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["monitor"], _ = expandFirewallVip46RealserversMonitor(d, i["monitor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["client-ip"], _ = expandFirewallVip46RealserversClientIp(d, i["client_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip46RealserversId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46RealserversClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip46Monitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallVip46MonitorName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip46MonitorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip46(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallVip46Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandFirewallVip46Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallVip46Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallVip46Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandFirewallVip46Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok {
		t, err := expandFirewallVip46SrcFilter(d, v, "src_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok {
		t, err := expandFirewallVip46Extip(d, v, "extip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok {
		t, err := expandFirewallVip46Mappedip(d, v, "mappedip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok {
		t, err := expandFirewallVip46ArpReply(d, v, "arp_reply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok {
		t, err := expandFirewallVip46Portforward(d, v, "portforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandFirewallVip46Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok {
		t, err := expandFirewallVip46Extport(d, v, "extport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok {
		t, err := expandFirewallVip46Mappedport(d, v, "mappedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandFirewallVip46Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok {
		t, err := expandFirewallVip46LdbMethod(d, v, "ldb_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandFirewallVip46ServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok {
		t, err := expandFirewallVip46Realservers(d, v, "realservers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok {
		t, err := expandFirewallVip46Monitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	return &obj, nil
}
