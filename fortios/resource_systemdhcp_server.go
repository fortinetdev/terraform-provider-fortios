// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure DHCP servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemDhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpServerCreate,
		Read:   resourceSystemDhcpServerRead,
		Update: resourceSystemDhcpServerUpdate,
		Delete: resourceSystemDhcpServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Required: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lease_time": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),
				Optional: true,
				Computed: true,
			},
			"mac_acl_default_action": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_on_net_status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_service": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server3": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac1": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac2": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac3": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_service": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server1": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server2": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server3": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"wins_server1": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_gateway": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_server": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required: true,
			},
			"ip_range": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"timezone_option": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"tftp_server": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tftp_server": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"filename": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"code": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_mode": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conflicted_ip_timeout": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 8640000),
				Optional: true,
				Computed: true,
			},
			"ipsec_lease_hold": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 8640000),
				Optional: true,
				Computed: true,
			},
			"auto_configuration": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_update": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_update_override": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_zone": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional: true,
				Computed: true,
			},
			"ddns_auth": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_keyname": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional: true,
				Computed: true,
			},
			"ddns_key": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_ttl": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional: true,
				Computed: true,
			},
			"vci_match": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vci_string": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"exclude_range": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"reserved_address": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"circuit_id": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional: true,
							Computed: true,
						},
						"remote_id": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemDhcpServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpServer resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDhcpServer(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDhcpServer")
	}

	return resourceSystemDhcpServerRead(d, m)
}

func resourceSystemDhcpServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDhcpServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServer resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDhcpServer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDhcpServer")
	}

	return resourceSystemDhcpServerRead(d, m)
}

func resourceSystemDhcpServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemDhcpServer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemDhcpServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpServer resource from API: %v", err)
	}
	return nil
}


func flattenSystemDhcpServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWifiAc1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWifiAc2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWifiAc3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNtpServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDefaultGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNextServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpServerIpRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenSystemDhcpServerIpRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenSystemDhcpServerIpRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcpServerIpRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerTimezoneOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerTimezone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tftp_server"
		if _, ok := i["tftp-server"]; ok {
			tmp["tftp_server"] = flattenSystemDhcpServerTftpServerTftpServer(i["tftp-server"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcpServerTftpServerTftpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerFilename(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpServerOptionsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			tmp["code"] = flattenSystemDhcpServerOptionsCode(i["code"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenSystemDhcpServerOptionsType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenSystemDhcpServerOptionsValue(i["value"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemDhcpServerOptionsIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcpServerOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerAutoConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsUpdate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerVciMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerVciString(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := i["vci-string"]; ok {
			tmp["vci_string"] = flattenSystemDhcpServerVciStringVciString(i["vci-string"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcpServerVciStringVciString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpServerExcludeRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenSystemDhcpServerExcludeRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenSystemDhcpServerExcludeRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcpServerExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpServerReservedAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenSystemDhcpServerReservedAddressType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemDhcpServerReservedAddressIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			tmp["mac"] = flattenSystemDhcpServerReservedAddressMac(i["mac"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenSystemDhcpServerReservedAddressAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := i["circuit-id"]; ok {
			tmp["circuit_id"] = flattenSystemDhcpServerReservedAddressCircuitId(i["circuit-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := i["remote-id"]; ok {
			tmp["remote_id"] = flattenSystemDhcpServerReservedAddressRemoteId(i["remote-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = flattenSystemDhcpServerReservedAddressDescription(i["description"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDhcpServerReservedAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDhcpServerReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemDhcpServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("fosid", flattenSystemDhcpServerId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDhcpServerStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpServerLeaseTime(o["lease-time"], d, "lease_time")); err != nil {
		if !fortiAPIPatch(o["lease-time"]) {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenSystemDhcpServerMacAclDefaultAction(o["mac-acl-default-action"], d, "mac_acl_default_action")); err != nil {
		if !fortiAPIPatch(o["mac-acl-default-action"]) {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenSystemDhcpServerForticlientOnNetStatus(o["forticlient-on-net-status"], d, "forticlient_on_net_status")); err != nil {
		if !fortiAPIPatch(o["forticlient-on-net-status"]) {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcpServerDnsService(o["dns-service"], d, "dns_service")); err != nil {
		if !fortiAPIPatch(o["dns-service"]) {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcpServerDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcpServerDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcpServerDnsServer3(o["dns-server3"], d, "dns_server3")); err != nil {
		if !fortiAPIPatch(o["dns-server3"]) {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenSystemDhcpServerWifiAc1(o["wifi-ac1"], d, "wifi_ac1")); err != nil {
		if !fortiAPIPatch(o["wifi-ac1"]) {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenSystemDhcpServerWifiAc2(o["wifi-ac2"], d, "wifi_ac2")); err != nil {
		if !fortiAPIPatch(o["wifi-ac2"]) {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenSystemDhcpServerWifiAc3(o["wifi-ac3"], d, "wifi_ac3")); err != nil {
		if !fortiAPIPatch(o["wifi-ac3"]) {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenSystemDhcpServerNtpService(o["ntp-service"], d, "ntp_service")); err != nil {
		if !fortiAPIPatch(o["ntp-service"]) {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenSystemDhcpServerNtpServer1(o["ntp-server1"], d, "ntp_server1")); err != nil {
		if !fortiAPIPatch(o["ntp-server1"]) {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenSystemDhcpServerNtpServer2(o["ntp-server2"], d, "ntp_server2")); err != nil {
		if !fortiAPIPatch(o["ntp-server2"]) {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenSystemDhcpServerNtpServer3(o["ntp-server3"], d, "ntp_server3")); err != nil {
		if !fortiAPIPatch(o["ntp-server3"]) {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcpServerDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenSystemDhcpServerWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if !fortiAPIPatch(o["wins-server1"]) {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenSystemDhcpServerWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if !fortiAPIPatch(o["wins-server2"]) {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("default_gateway", flattenSystemDhcpServerDefaultGateway(o["default-gateway"], d, "default_gateway")); err != nil {
		if !fortiAPIPatch(o["default-gateway"]) {
			return fmt.Errorf("Error reading default_gateway: %v", err)
		}
	}

	if err = d.Set("next_server", flattenSystemDhcpServerNextServer(o["next-server"], d, "next_server")); err != nil {
		if !fortiAPIPatch(o["next-server"]) {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("netmask", flattenSystemDhcpServerNetmask(o["netmask"], d, "netmask")); err != nil {
		if !fortiAPIPatch(o["netmask"]) {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemDhcpServerInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("ip_range", flattenSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
            if !fortiAPIPatch(o["ip-range"]) {
                return fmt.Errorf("Error reading ip_range: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("ip_range"); ok {
            if err = d.Set("ip_range", flattenSystemDhcpServerIpRange(o["ip-range"], d, "ip_range")); err != nil {
                if !fortiAPIPatch(o["ip-range"]) {
                    return fmt.Errorf("Error reading ip_range: %v", err)
                }
            }
        }
    }

	if err = d.Set("timezone_option", flattenSystemDhcpServerTimezoneOption(o["timezone-option"], d, "timezone_option")); err != nil {
		if !fortiAPIPatch(o["timezone-option"]) {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemDhcpServerTimezone(o["timezone"], d, "timezone")); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("tftp_server", flattenSystemDhcpServerTftpServer(o["tftp-server"], d, "tftp_server")); err != nil {
            if !fortiAPIPatch(o["tftp-server"]) {
                return fmt.Errorf("Error reading tftp_server: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("tftp_server"); ok {
            if err = d.Set("tftp_server", flattenSystemDhcpServerTftpServer(o["tftp-server"], d, "tftp_server")); err != nil {
                if !fortiAPIPatch(o["tftp-server"]) {
                    return fmt.Errorf("Error reading tftp_server: %v", err)
                }
            }
        }
    }

	if err = d.Set("filename", flattenSystemDhcpServerFilename(o["filename"], d, "filename")); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("options", flattenSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
            if !fortiAPIPatch(o["options"]) {
                return fmt.Errorf("Error reading options: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("options"); ok {
            if err = d.Set("options", flattenSystemDhcpServerOptions(o["options"], d, "options")); err != nil {
                if !fortiAPIPatch(o["options"]) {
                    return fmt.Errorf("Error reading options: %v", err)
                }
            }
        }
    }

	if err = d.Set("server_type", flattenSystemDhcpServerServerType(o["server-type"], d, "server_type")); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenSystemDhcpServerIpMode(o["ip-mode"], d, "ip_mode")); err != nil {
		if !fortiAPIPatch(o["ip-mode"]) {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenSystemDhcpServerConflictedIpTimeout(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout")); err != nil {
		if !fortiAPIPatch(o["conflicted-ip-timeout"]) {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenSystemDhcpServerIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold")); err != nil {
		if !fortiAPIPatch(o["ipsec-lease-hold"]) {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("auto_configuration", flattenSystemDhcpServerAutoConfiguration(o["auto-configuration"], d, "auto_configuration")); err != nil {
		if !fortiAPIPatch(o["auto-configuration"]) {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenSystemDhcpServerDdnsUpdate(o["ddns-update"], d, "ddns_update")); err != nil {
		if !fortiAPIPatch(o["ddns-update"]) {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenSystemDhcpServerDdnsUpdateOverride(o["ddns-update-override"], d, "ddns_update_override")); err != nil {
		if !fortiAPIPatch(o["ddns-update-override"]) {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemDhcpServerDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenSystemDhcpServerDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if !fortiAPIPatch(o["ddns-zone"]) {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenSystemDhcpServerDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if !fortiAPIPatch(o["ddns-auth"]) {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenSystemDhcpServerDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if !fortiAPIPatch(o["ddns-keyname"]) {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_key", flattenSystemDhcpServerDdnsKey(o["ddns-key"], d, "ddns_key")); err != nil {
		if !fortiAPIPatch(o["ddns-key"]) {
			return fmt.Errorf("Error reading ddns_key: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenSystemDhcpServerDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if !fortiAPIPatch(o["ddns-ttl"]) {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpServerVciMatch(o["vci-match"], d, "vci_match")); err != nil {
		if !fortiAPIPatch(o["vci-match"]) {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("vci_string", flattenSystemDhcpServerVciString(o["vci-string"], d, "vci_string")); err != nil {
            if !fortiAPIPatch(o["vci-string"]) {
                return fmt.Errorf("Error reading vci_string: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("vci_string"); ok {
            if err = d.Set("vci_string", flattenSystemDhcpServerVciString(o["vci-string"], d, "vci_string")); err != nil {
                if !fortiAPIPatch(o["vci-string"]) {
                    return fmt.Errorf("Error reading vci_string: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("exclude_range", flattenSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
            if !fortiAPIPatch(o["exclude-range"]) {
                return fmt.Errorf("Error reading exclude_range: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("exclude_range"); ok {
            if err = d.Set("exclude_range", flattenSystemDhcpServerExcludeRange(o["exclude-range"], d, "exclude_range")); err != nil {
                if !fortiAPIPatch(o["exclude-range"]) {
                    return fmt.Errorf("Error reading exclude_range: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("reserved_address", flattenSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
            if !fortiAPIPatch(o["reserved-address"]) {
                return fmt.Errorf("Error reading reserved_address: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("reserved_address"); ok {
            if err = d.Set("reserved_address", flattenSystemDhcpServerReservedAddress(o["reserved-address"], d, "reserved_address")); err != nil {
                if !fortiAPIPatch(o["reserved-address"]) {
                    return fmt.Errorf("Error reading reserved_address: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSystemDhcpServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemDhcpServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWifiAc1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWifiAc2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWifiAc3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNtpServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerWinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDefaultGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNextServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemDhcpServerIpRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandSystemDhcpServerIpRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandSystemDhcpServerIpRangeEndIp(d, i["end_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerIpRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerTimezoneOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerTimezone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tftp_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tftp-server"], _ = expandSystemDhcpServerTftpServerTftpServer(d, i["tftp_server"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerTftpServerTftpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerFilename(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemDhcpServerOptionsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["code"], _ = expandSystemDhcpServerOptionsCode(d, i["code"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemDhcpServerOptionsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSystemDhcpServerOptionsValue(d, i["value"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemDhcpServerOptionsIp(d, i["ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerOptionsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerOptionsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerAutoConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsUpdate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsZone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsKeyname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerDdnsTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerVciMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-string"], _ = expandSystemDhcpServerVciStringVciString(d, i["vci_string"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerVciStringVciString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemDhcpServerExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandSystemDhcpServerExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandSystemDhcpServerExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemDhcpServerReservedAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemDhcpServerReservedAddressType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemDhcpServerReservedAddressIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandSystemDhcpServerReservedAddressMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemDhcpServerReservedAddressAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["circuit-id"], _ = expandSystemDhcpServerReservedAddressCircuitId(d, i["circuit_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-id"], _ = expandSystemDhcpServerReservedAddressRemoteId(d, i["remote_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSystemDhcpServerReservedAddressDescription(d, i["description"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpServerReservedAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpServerReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemDhcpServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemDhcpServerId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDhcpServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("lease_time"); ok {
		t, err := expandSystemDhcpServerLeaseTime(d, v, "lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok {
		t, err := expandSystemDhcpServerMacAclDefaultAction(d, v, "mac_acl_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok {
		t, err := expandSystemDhcpServerForticlientOnNetStatus(d, v, "forticlient_on_net_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok {
		t, err := expandSystemDhcpServerDnsService(d, v, "dns_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok {
		t, err := expandSystemDhcpServerDnsServer1(d, v, "dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok {
		t, err := expandSystemDhcpServerDnsServer2(d, v, "dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok {
		t, err := expandSystemDhcpServerDnsServer3(d, v, "dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok {
		t, err := expandSystemDhcpServerWifiAc1(d, v, "wifi_ac1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok {
		t, err := expandSystemDhcpServerWifiAc2(d, v, "wifi_ac2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok {
		t, err := expandSystemDhcpServerWifiAc3(d, v, "wifi_ac3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok {
		t, err := expandSystemDhcpServerNtpService(d, v, "ntp_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok {
		t, err := expandSystemDhcpServerNtpServer1(d, v, "ntp_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok {
		t, err := expandSystemDhcpServerNtpServer2(d, v, "ntp_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok {
		t, err := expandSystemDhcpServerNtpServer3(d, v, "ntp_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDhcpServerDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok {
		t, err := expandSystemDhcpServerWinsServer1(d, v, "wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok {
		t, err := expandSystemDhcpServerWinsServer2(d, v, "wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("default_gateway"); ok {
		t, err := expandSystemDhcpServerDefaultGateway(d, v, "default_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gateway"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok {
		t, err := expandSystemDhcpServerNextServer(d, v, "next_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok {
		t, err := expandSystemDhcpServerNetmask(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemDhcpServerInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok {
		t, err := expandSystemDhcpServerIpRange(d, v, "ip_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok {
		t, err := expandSystemDhcpServerTimezoneOption(d, v, "timezone_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok {
		t, err := expandSystemDhcpServerTimezone(d, v, "timezone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	}

	if v, ok := d.GetOk("tftp_server"); ok {
		t, err := expandSystemDhcpServerTftpServer(d, v, "tftp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok {
		t, err := expandSystemDhcpServerFilename(d, v, "filename")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandSystemDhcpServerOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandSystemDhcpServerServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok {
		t, err := expandSystemDhcpServerIpMode(d, v, "ip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok {
		t, err := expandSystemDhcpServerConflictedIpTimeout(d, v, "conflicted_ip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_lease_hold"); ok {
		t, err := expandSystemDhcpServerIpsecLeaseHold(d, v, "ipsec_lease_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("auto_configuration"); ok {
		t, err := expandSystemDhcpServerAutoConfiguration(d, v, "auto_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok {
		t, err := expandSystemDhcpServerDdnsUpdate(d, v, "ddns_update")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok {
		t, err := expandSystemDhcpServerDdnsUpdateOverride(d, v, "ddns_update_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok {
		t, err := expandSystemDhcpServerDdnsServerIp(d, v, "ddns_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok {
		t, err := expandSystemDhcpServerDdnsZone(d, v, "ddns_zone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	}

	if v, ok := d.GetOk("ddns_auth"); ok {
		t, err := expandSystemDhcpServerDdnsAuth(d, v, "ddns_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok {
		t, err := expandSystemDhcpServerDdnsKeyname(d, v, "ddns_keyname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	}

	if v, ok := d.GetOk("ddns_key"); ok {
		t, err := expandSystemDhcpServerDdnsKey(d, v, "ddns_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	}

	if v, ok := d.GetOk("ddns_ttl"); ok {
		t, err := expandSystemDhcpServerDdnsTtl(d, v, "ddns_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok {
		t, err := expandSystemDhcpServerVciMatch(d, v, "vci_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok {
		t, err := expandSystemDhcpServerVciString(d, v, "vci_string")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok {
		t, err := expandSystemDhcpServerExcludeRange(d, v, "exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok {
		t, err := expandSystemDhcpServerReservedAddress(d, v, "reserved_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}


	return &obj, nil
}

