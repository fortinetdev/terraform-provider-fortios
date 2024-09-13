// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Filters for memory buffer.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogMemoryFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemoryFilterUpdate,
		Read:   resourceLogMemoryFilterRead,
		Update: resourceLogMemoryFilterUpdate,
		Delete: resourceLogMemoryFilterDelete,

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
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sniffer_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_transaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anomaly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netscan_vulnerability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"voip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forti_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"free_style": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
							Optional:     true,
						},
						"filter_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"radius": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ppp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslvpn_log_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslvpn_log_adm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslvpn_log_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldb_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wan_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wireless_activity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_memory_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"filter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceLogMemoryFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogMemoryFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemoryFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemoryFilter")
	}

	return resourceLogMemoryFilterRead(d, m)
}

func resourceLogMemoryFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogMemoryFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogMemoryFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogMemoryFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogMemoryFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemoryFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogMemoryFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemoryFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogMemoryFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterGtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenLogMemoryFilterFreeStyleId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenLogMemoryFilterFreeStyleCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if cur_v, ok := i["filter"]; ok {
			tmp["filter"] = flattenLogMemoryFilterFreeStyleFilter(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if cur_v, ok := i["filter-type"]; ok {
			tmp["filter_type"] = flattenLogMemoryFilterFreeStyleFilterType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogMemoryFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogMemoryFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSsh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSystem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterRadius(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterDhcp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterPpp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSslvpnLogAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSslvpnLogAdm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterSslvpnLogSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterVipSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterLdbMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterWanOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterWirelessActivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterCpuMemoryUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryFilterFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogMemoryFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("severity", flattenLogMemoryFilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogMemoryFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogMemoryFilterLocalTraffic(o["local-traffic"], d, "local_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogMemoryFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogMemoryFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogMemoryFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-traffic"]) {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogMemoryFilterHttpTransaction(o["http-transaction"], d, "http_transaction", sv)); err != nil {
		if !fortiAPIPatch(o["http-transaction"]) {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogMemoryFilterAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogMemoryFilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogMemoryFilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogMemoryFilterVoip(o["voip"], d, "voip", sv)); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogMemoryFilterGtp(o["gtp"], d, "gtp", sv)); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogMemoryFilterFortiSwitch(o["forti-switch"], d, "forti_switch", sv)); err != nil {
		if !fortiAPIPatch(o["forti-switch"]) {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("free_style", flattenLogMemoryFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
			if !fortiAPIPatch(o["free-style"]) {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogMemoryFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
				if !fortiAPIPatch(o["free-style"]) {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns", flattenLogMemoryFilterDns(o["dns"], d, "dns", sv)); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogMemoryFilterSsh(o["ssh"], d, "ssh", sv)); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("event", flattenLogMemoryFilterEvent(o["event"], d, "event", sv)); err != nil {
		if !fortiAPIPatch(o["event"]) {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("system", flattenLogMemoryFilterSystem(o["system"], d, "system", sv)); err != nil {
		if !fortiAPIPatch(o["system"]) {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("radius", flattenLogMemoryFilterRadius(o["radius"], d, "radius", sv)); err != nil {
		if !fortiAPIPatch(o["radius"]) {
			return fmt.Errorf("Error reading radius: %v", err)
		}
	}

	if err = d.Set("ipsec", flattenLogMemoryFilterIpsec(o["ipsec"], d, "ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}

	if err = d.Set("dhcp", flattenLogMemoryFilterDhcp(o["dhcp"], d, "dhcp", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp"]) {
			return fmt.Errorf("Error reading dhcp: %v", err)
		}
	}

	if err = d.Set("ppp", flattenLogMemoryFilterPpp(o["ppp"], d, "ppp", sv)); err != nil {
		if !fortiAPIPatch(o["ppp"]) {
			return fmt.Errorf("Error reading ppp: %v", err)
		}
	}

	if err = d.Set("admin", flattenLogMemoryFilterAdmin(o["admin"], d, "admin", sv)); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("ha", flattenLogMemoryFilterHa(o["ha"], d, "ha", sv)); err != nil {
		if !fortiAPIPatch(o["ha"]) {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("auth", flattenLogMemoryFilterAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("pattern", flattenLogMemoryFilterPattern(o["pattern"], d, "pattern", sv)); err != nil {
		if !fortiAPIPatch(o["pattern"]) {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_auth", flattenLogMemoryFilterSslvpnLogAuth(o["sslvpn-log-auth"], d, "sslvpn_log_auth", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-auth"]) {
			return fmt.Errorf("Error reading sslvpn_log_auth: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_adm", flattenLogMemoryFilterSslvpnLogAdm(o["sslvpn-log-adm"], d, "sslvpn_log_adm", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-adm"]) {
			return fmt.Errorf("Error reading sslvpn_log_adm: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_session", flattenLogMemoryFilterSslvpnLogSession(o["sslvpn-log-session"], d, "sslvpn_log_session", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-session"]) {
			return fmt.Errorf("Error reading sslvpn_log_session: %v", err)
		}
	}

	if err = d.Set("vip_ssl", flattenLogMemoryFilterVipSsl(o["vip-ssl"], d, "vip_ssl", sv)); err != nil {
		if !fortiAPIPatch(o["vip-ssl"]) {
			return fmt.Errorf("Error reading vip_ssl: %v", err)
		}
	}

	if err = d.Set("ldb_monitor", flattenLogMemoryFilterLdbMonitor(o["ldb-monitor"], d, "ldb_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["ldb-monitor"]) {
			return fmt.Errorf("Error reading ldb_monitor: %v", err)
		}
	}

	if err = d.Set("wan_opt", flattenLogMemoryFilterWanOpt(o["wan-opt"], d, "wan_opt", sv)); err != nil {
		if !fortiAPIPatch(o["wan-opt"]) {
			return fmt.Errorf("Error reading wan_opt: %v", err)
		}
	}

	if err = d.Set("wireless_activity", flattenLogMemoryFilterWirelessActivity(o["wireless-activity"], d, "wireless_activity", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-activity"]) {
			return fmt.Errorf("Error reading wireless_activity: %v", err)
		}
	}

	if err = d.Set("cpu_memory_usage", flattenLogMemoryFilterCpuMemoryUsage(o["cpu-memory-usage"], d, "cpu_memory_usage", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-memory-usage"]) {
			return fmt.Errorf("Error reading cpu_memory_usage: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogMemoryFilterFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogMemoryFilterFilterType(o["filter-type"], d, "filter_type", sv)); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogMemoryFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogMemoryFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterGtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandLogMemoryFilterFreeStyleId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandLogMemoryFilterFreeStyleCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter"], _ = expandLogMemoryFilterFreeStyleFilter(d, i["filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-type"], _ = expandLogMemoryFilterFreeStyleFilterType(d, i["filter_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogMemoryFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSystem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterRadius(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterDhcp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterPpp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSslvpnLogAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSslvpnLogAdm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterSslvpnLogSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterVipSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterLdbMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterWanOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterWirelessActivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterCpuMemoryUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryFilterFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemoryFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {
			t, err := expandLogMemoryFilterSeverity(d, v, "severity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["severity"] = t
			}
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok {
		if setArgNil {
			obj["forward-traffic"] = nil
		} else {
			t, err := expandLogMemoryFilterForwardTraffic(d, v, "forward_traffic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward-traffic"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok {
		if setArgNil {
			obj["local-traffic"] = nil
		} else {
			t, err := expandLogMemoryFilterLocalTraffic(d, v, "local_traffic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-traffic"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok {
		if setArgNil {
			obj["multicast-traffic"] = nil
		} else {
			t, err := expandLogMemoryFilterMulticastTraffic(d, v, "multicast_traffic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-traffic"] = t
			}
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok {
		if setArgNil {
			obj["sniffer-traffic"] = nil
		} else {
			t, err := expandLogMemoryFilterSnifferTraffic(d, v, "sniffer_traffic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sniffer-traffic"] = t
			}
		}
	}

	if v, ok := d.GetOk("ztna_traffic"); ok {
		if setArgNil {
			obj["ztna-traffic"] = nil
		} else {
			t, err := expandLogMemoryFilterZtnaTraffic(d, v, "ztna_traffic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ztna-traffic"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_transaction"); ok {
		if setArgNil {
			obj["http-transaction"] = nil
		} else {
			t, err := expandLogMemoryFilterHttpTransaction(d, v, "http_transaction", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-transaction"] = t
			}
		}
	}

	if v, ok := d.GetOk("anomaly"); ok {
		if setArgNil {
			obj["anomaly"] = nil
		} else {
			t, err := expandLogMemoryFilterAnomaly(d, v, "anomaly", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["anomaly"] = t
			}
		}
	}

	if v, ok := d.GetOk("netscan_discovery"); ok {
		if setArgNil {
			obj["netscan-discovery"] = nil
		} else {
			t, err := expandLogMemoryFilterNetscanDiscovery(d, v, "netscan_discovery", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["netscan-discovery"] = t
			}
		}
	} else if d.HasChange("netscan_discovery") {
		obj["netscan-discovery"] = nil
	}

	if v, ok := d.GetOk("netscan_vulnerability"); ok {
		if setArgNil {
			obj["netscan-vulnerability"] = nil
		} else {
			t, err := expandLogMemoryFilterNetscanVulnerability(d, v, "netscan_vulnerability", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["netscan-vulnerability"] = t
			}
		}
	} else if d.HasChange("netscan_vulnerability") {
		obj["netscan-vulnerability"] = nil
	}

	if v, ok := d.GetOk("voip"); ok {
		if setArgNil {
			obj["voip"] = nil
		} else {
			t, err := expandLogMemoryFilterVoip(d, v, "voip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["voip"] = t
			}
		}
	}

	if v, ok := d.GetOk("gtp"); ok {
		if setArgNil {
			obj["gtp"] = nil
		} else {
			t, err := expandLogMemoryFilterGtp(d, v, "gtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gtp"] = t
			}
		}
	}

	if v, ok := d.GetOk("forti_switch"); ok {
		if setArgNil {
			obj["forti-switch"] = nil
		} else {
			t, err := expandLogMemoryFilterFortiSwitch(d, v, "forti_switch", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forti-switch"] = t
			}
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		if setArgNil {
			obj["free-style"] = make([]struct{}, 0)
		} else {
			t, err := expandLogMemoryFilterFreeStyle(d, v, "free_style", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["free-style"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns"); ok {
		if setArgNil {
			obj["dns"] = nil
		} else {
			t, err := expandLogMemoryFilterDns(d, v, "dns", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns"] = t
			}
		}
	} else if d.HasChange("dns") {
		obj["dns"] = nil
	}

	if v, ok := d.GetOk("ssh"); ok {
		if setArgNil {
			obj["ssh"] = nil
		} else {
			t, err := expandLogMemoryFilterSsh(d, v, "ssh", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh"] = t
			}
		}
	} else if d.HasChange("ssh") {
		obj["ssh"] = nil
	}

	if v, ok := d.GetOk("event"); ok {
		if setArgNil {
			obj["event"] = nil
		} else {
			t, err := expandLogMemoryFilterEvent(d, v, "event", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["event"] = t
			}
		}
	} else if d.HasChange("event") {
		obj["event"] = nil
	}

	if v, ok := d.GetOk("system"); ok {
		if setArgNil {
			obj["system"] = nil
		} else {
			t, err := expandLogMemoryFilterSystem(d, v, "system", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["system"] = t
			}
		}
	} else if d.HasChange("system") {
		obj["system"] = nil
	}

	if v, ok := d.GetOk("radius"); ok {
		if setArgNil {
			obj["radius"] = nil
		} else {
			t, err := expandLogMemoryFilterRadius(d, v, "radius", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radius"] = t
			}
		}
	} else if d.HasChange("radius") {
		obj["radius"] = nil
	}

	if v, ok := d.GetOk("ipsec"); ok {
		if setArgNil {
			obj["ipsec"] = nil
		} else {
			t, err := expandLogMemoryFilterIpsec(d, v, "ipsec", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec"] = t
			}
		}
	} else if d.HasChange("ipsec") {
		obj["ipsec"] = nil
	}

	if v, ok := d.GetOk("dhcp"); ok {
		if setArgNil {
			obj["dhcp"] = nil
		} else {
			t, err := expandLogMemoryFilterDhcp(d, v, "dhcp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp"] = t
			}
		}
	} else if d.HasChange("dhcp") {
		obj["dhcp"] = nil
	}

	if v, ok := d.GetOk("ppp"); ok {
		if setArgNil {
			obj["ppp"] = nil
		} else {
			t, err := expandLogMemoryFilterPpp(d, v, "ppp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ppp"] = t
			}
		}
	} else if d.HasChange("ppp") {
		obj["ppp"] = nil
	}

	if v, ok := d.GetOk("admin"); ok {
		if setArgNil {
			obj["admin"] = nil
		} else {
			t, err := expandLogMemoryFilterAdmin(d, v, "admin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin"] = t
			}
		}
	} else if d.HasChange("admin") {
		obj["admin"] = nil
	}

	if v, ok := d.GetOk("ha"); ok {
		if setArgNil {
			obj["ha"] = nil
		} else {
			t, err := expandLogMemoryFilterHa(d, v, "ha", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha"] = t
			}
		}
	} else if d.HasChange("ha") {
		obj["ha"] = nil
	}

	if v, ok := d.GetOk("auth"); ok {
		if setArgNil {
			obj["auth"] = nil
		} else {
			t, err := expandLogMemoryFilterAuth(d, v, "auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth"] = t
			}
		}
	} else if d.HasChange("auth") {
		obj["auth"] = nil
	}

	if v, ok := d.GetOk("pattern"); ok {
		if setArgNil {
			obj["pattern"] = nil
		} else {
			t, err := expandLogMemoryFilterPattern(d, v, "pattern", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pattern"] = t
			}
		}
	} else if d.HasChange("pattern") {
		obj["pattern"] = nil
	}

	if v, ok := d.GetOk("sslvpn_log_auth"); ok {
		if setArgNil {
			obj["sslvpn-log-auth"] = nil
		} else {
			t, err := expandLogMemoryFilterSslvpnLogAuth(d, v, "sslvpn_log_auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-log-auth"] = t
			}
		}
	} else if d.HasChange("sslvpn_log_auth") {
		obj["sslvpn-log-auth"] = nil
	}

	if v, ok := d.GetOk("sslvpn_log_adm"); ok {
		if setArgNil {
			obj["sslvpn-log-adm"] = nil
		} else {
			t, err := expandLogMemoryFilterSslvpnLogAdm(d, v, "sslvpn_log_adm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-log-adm"] = t
			}
		}
	} else if d.HasChange("sslvpn_log_adm") {
		obj["sslvpn-log-adm"] = nil
	}

	if v, ok := d.GetOk("sslvpn_log_session"); ok {
		if setArgNil {
			obj["sslvpn-log-session"] = nil
		} else {
			t, err := expandLogMemoryFilterSslvpnLogSession(d, v, "sslvpn_log_session", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-log-session"] = t
			}
		}
	} else if d.HasChange("sslvpn_log_session") {
		obj["sslvpn-log-session"] = nil
	}

	if v, ok := d.GetOk("vip_ssl"); ok {
		if setArgNil {
			obj["vip-ssl"] = nil
		} else {
			t, err := expandLogMemoryFilterVipSsl(d, v, "vip_ssl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vip-ssl"] = t
			}
		}
	} else if d.HasChange("vip_ssl") {
		obj["vip-ssl"] = nil
	}

	if v, ok := d.GetOk("ldb_monitor"); ok {
		if setArgNil {
			obj["ldb-monitor"] = nil
		} else {
			t, err := expandLogMemoryFilterLdbMonitor(d, v, "ldb_monitor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ldb-monitor"] = t
			}
		}
	} else if d.HasChange("ldb_monitor") {
		obj["ldb-monitor"] = nil
	}

	if v, ok := d.GetOk("wan_opt"); ok {
		if setArgNil {
			obj["wan-opt"] = nil
		} else {
			t, err := expandLogMemoryFilterWanOpt(d, v, "wan_opt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wan-opt"] = t
			}
		}
	} else if d.HasChange("wan_opt") {
		obj["wan-opt"] = nil
	}

	if v, ok := d.GetOk("wireless_activity"); ok {
		if setArgNil {
			obj["wireless-activity"] = nil
		} else {
			t, err := expandLogMemoryFilterWirelessActivity(d, v, "wireless_activity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wireless-activity"] = t
			}
		}
	} else if d.HasChange("wireless_activity") {
		obj["wireless-activity"] = nil
	}

	if v, ok := d.GetOk("cpu_memory_usage"); ok {
		if setArgNil {
			obj["cpu-memory-usage"] = nil
		} else {
			t, err := expandLogMemoryFilterCpuMemoryUsage(d, v, "cpu_memory_usage", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cpu-memory-usage"] = t
			}
		}
	} else if d.HasChange("cpu_memory_usage") {
		obj["cpu-memory-usage"] = nil
	}

	if v, ok := d.GetOk("filter"); ok {
		if setArgNil {
			obj["filter"] = nil
		} else {
			t, err := expandLogMemoryFilterFilter(d, v, "filter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter"] = t
			}
		}
	} else if d.HasChange("filter") {
		obj["filter"] = nil
	}

	if v, ok := d.GetOk("filter_type"); ok {
		if setArgNil {
			obj["filter-type"] = nil
		} else {
			t, err := expandLogMemoryFilterFilterType(d, v, "filter_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-type"] = t
			}
		}
	}

	return &obj, nil
}
