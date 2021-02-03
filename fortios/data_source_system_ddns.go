// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DDNS.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemDdns() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemDdnsRead,
		Schema: map[string]*schema.Schema{
			"ddnsid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ddns_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_zone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_keyname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"ddns_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_sn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ddns_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"use_public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"clear_text": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bound_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitor_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemDdnsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("ddnsid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemDdns: type error")
	}

	o, err := c.ReadSystemDdns(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemDdns: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemDdns(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemDdns from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemDdnsDdnsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsZone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsKeyname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsDdnsPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsUsePublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsClearText(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsSslCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsBoundIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDdnsMonitorInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			tmp["interface_name"] = dataSourceFlattenSystemDdnsMonitorInterfaceInterfaceName(i["interface-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDdnsMonitorInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemDdns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ddnsid", dataSourceFlattenSystemDdnsDdnsid(o["ddnsid"], d, "ddnsid")); err != nil {
		if !fortiAPIPatch(o["ddnsid"]) {
			return fmt.Errorf("Error reading ddnsid: %v", err)
		}
	}

	if err = d.Set("ddns_server", dataSourceFlattenSystemDdnsDdnsServer(o["ddns-server"], d, "ddns_server")); err != nil {
		if !fortiAPIPatch(o["ddns-server"]) {
			return fmt.Errorf("Error reading ddns_server: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", dataSourceFlattenSystemDdnsDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip")); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_zone", dataSourceFlattenSystemDdnsDdnsZone(o["ddns-zone"], d, "ddns_zone")); err != nil {
		if !fortiAPIPatch(o["ddns-zone"]) {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", dataSourceFlattenSystemDdnsDdnsTtl(o["ddns-ttl"], d, "ddns_ttl")); err != nil {
		if !fortiAPIPatch(o["ddns-ttl"]) {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("ddns_auth", dataSourceFlattenSystemDdnsDdnsAuth(o["ddns-auth"], d, "ddns_auth")); err != nil {
		if !fortiAPIPatch(o["ddns-auth"]) {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", dataSourceFlattenSystemDdnsDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname")); err != nil {
		if !fortiAPIPatch(o["ddns-keyname"]) {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_domain", dataSourceFlattenSystemDdnsDdnsDomain(o["ddns-domain"], d, "ddns_domain")); err != nil {
		if !fortiAPIPatch(o["ddns-domain"]) {
			return fmt.Errorf("Error reading ddns_domain: %v", err)
		}
	}

	if err = d.Set("ddns_username", dataSourceFlattenSystemDdnsDdnsUsername(o["ddns-username"], d, "ddns_username")); err != nil {
		if !fortiAPIPatch(o["ddns-username"]) {
			return fmt.Errorf("Error reading ddns_username: %v", err)
		}
	}

	if err = d.Set("ddns_sn", dataSourceFlattenSystemDdnsDdnsSn(o["ddns-sn"], d, "ddns_sn")); err != nil {
		if !fortiAPIPatch(o["ddns-sn"]) {
			return fmt.Errorf("Error reading ddns_sn: %v", err)
		}
	}

	if err = d.Set("use_public_ip", dataSourceFlattenSystemDdnsUsePublicIp(o["use-public-ip"], d, "use_public_ip")); err != nil {
		if !fortiAPIPatch(o["use-public-ip"]) {
			return fmt.Errorf("Error reading use_public_ip: %v", err)
		}
	}

	if err = d.Set("update_interval", dataSourceFlattenSystemDdnsUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("clear_text", dataSourceFlattenSystemDdnsClearText(o["clear-text"], d, "clear_text")); err != nil {
		if !fortiAPIPatch(o["clear-text"]) {
			return fmt.Errorf("Error reading clear_text: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", dataSourceFlattenSystemDdnsSslCertificate(o["ssl-certificate"], d, "ssl_certificate")); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("bound_ip", dataSourceFlattenSystemDdnsBoundIp(o["bound-ip"], d, "bound_ip")); err != nil {
		if !fortiAPIPatch(o["bound-ip"]) {
			return fmt.Errorf("Error reading bound_ip: %v", err)
		}
	}

	if err = d.Set("monitor_interface", dataSourceFlattenSystemDdnsMonitorInterface(o["monitor-interface"], d, "monitor_interface")); err != nil {
		if !fortiAPIPatch(o["monitor-interface"]) {
			return fmt.Errorf("Error reading monitor_interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemDdnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
