// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure filters for local disk logging. Use these filters to determine the log messages to record according to severity and type.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogDiskFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogDiskFilterUpdate,
		Read:   resourceLogDiskFilterRead,
		Update: resourceLogDiskFilterUpdate,
		Delete: resourceLogDiskFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"anomaly": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_vulnerability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp": &schema.Schema{
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
							Computed: true,
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
							Computed:     true,
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
				Computed: true,
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_log_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_log_adm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslvpn_log_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldb_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wireless_activity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_memory_usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
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
		},
	}
}

func resourceLogDiskFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogDiskFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogDiskFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogDiskFilter")
	}

	return resourceLogDiskFilterRead(d, m)
}

func resourceLogDiskFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogDiskFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogDiskFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogDiskFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogDiskFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogDiskFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogDiskFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogDiskFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterDlpArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterGtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenLogDiskFilterFreeStyleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenLogDiskFilterFreeStyleCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {

			tmp["filter"] = flattenLogDiskFilterFreeStyleFilter(i["filter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {

			tmp["filter_type"] = flattenLogDiskFilterFreeStyleFilterType(i["filter-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogDiskFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterSsh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterSystem(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterRadius(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterIpsec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterDhcp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterPpp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterSslvpnLogAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterSslvpnLogAdm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterSslvpnLogSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterVipSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterLdbMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterWanOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterWirelessActivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterCpuMemoryUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskFilterFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogDiskFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("severity", flattenLogDiskFilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogDiskFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogDiskFilterLocalTraffic(o["local-traffic"], d, "local_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogDiskFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogDiskFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogDiskFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-traffic"]) {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogDiskFilterAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogDiskFilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogDiskFilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogDiskFilterVoip(o["voip"], d, "voip", sv)); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("dlp_archive", flattenLogDiskFilterDlpArchive(o["dlp-archive"], d, "dlp_archive", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-archive"]) {
			return fmt.Errorf("Error reading dlp_archive: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogDiskFilterGtp(o["gtp"], d, "gtp", sv)); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogDiskFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
			if !fortiAPIPatch(o["free-style"]) {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogDiskFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
				if !fortiAPIPatch(o["free-style"]) {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns", flattenLogDiskFilterDns(o["dns"], d, "dns", sv)); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogDiskFilterSsh(o["ssh"], d, "ssh", sv)); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("event", flattenLogDiskFilterEvent(o["event"], d, "event", sv)); err != nil {
		if !fortiAPIPatch(o["event"]) {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("system", flattenLogDiskFilterSystem(o["system"], d, "system", sv)); err != nil {
		if !fortiAPIPatch(o["system"]) {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("radius", flattenLogDiskFilterRadius(o["radius"], d, "radius", sv)); err != nil {
		if !fortiAPIPatch(o["radius"]) {
			return fmt.Errorf("Error reading radius: %v", err)
		}
	}

	if err = d.Set("ipsec", flattenLogDiskFilterIpsec(o["ipsec"], d, "ipsec", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}

	if err = d.Set("dhcp", flattenLogDiskFilterDhcp(o["dhcp"], d, "dhcp", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp"]) {
			return fmt.Errorf("Error reading dhcp: %v", err)
		}
	}

	if err = d.Set("ppp", flattenLogDiskFilterPpp(o["ppp"], d, "ppp", sv)); err != nil {
		if !fortiAPIPatch(o["ppp"]) {
			return fmt.Errorf("Error reading ppp: %v", err)
		}
	}

	if err = d.Set("admin", flattenLogDiskFilterAdmin(o["admin"], d, "admin", sv)); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("ha", flattenLogDiskFilterHa(o["ha"], d, "ha", sv)); err != nil {
		if !fortiAPIPatch(o["ha"]) {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("auth", flattenLogDiskFilterAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("pattern", flattenLogDiskFilterPattern(o["pattern"], d, "pattern", sv)); err != nil {
		if !fortiAPIPatch(o["pattern"]) {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_auth", flattenLogDiskFilterSslvpnLogAuth(o["sslvpn-log-auth"], d, "sslvpn_log_auth", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-auth"]) {
			return fmt.Errorf("Error reading sslvpn_log_auth: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_adm", flattenLogDiskFilterSslvpnLogAdm(o["sslvpn-log-adm"], d, "sslvpn_log_adm", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-adm"]) {
			return fmt.Errorf("Error reading sslvpn_log_adm: %v", err)
		}
	}

	if err = d.Set("sslvpn_log_session", flattenLogDiskFilterSslvpnLogSession(o["sslvpn-log-session"], d, "sslvpn_log_session", sv)); err != nil {
		if !fortiAPIPatch(o["sslvpn-log-session"]) {
			return fmt.Errorf("Error reading sslvpn_log_session: %v", err)
		}
	}

	if err = d.Set("vip_ssl", flattenLogDiskFilterVipSsl(o["vip-ssl"], d, "vip_ssl", sv)); err != nil {
		if !fortiAPIPatch(o["vip-ssl"]) {
			return fmt.Errorf("Error reading vip_ssl: %v", err)
		}
	}

	if err = d.Set("ldb_monitor", flattenLogDiskFilterLdbMonitor(o["ldb-monitor"], d, "ldb_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["ldb-monitor"]) {
			return fmt.Errorf("Error reading ldb_monitor: %v", err)
		}
	}

	if err = d.Set("wan_opt", flattenLogDiskFilterWanOpt(o["wan-opt"], d, "wan_opt", sv)); err != nil {
		if !fortiAPIPatch(o["wan-opt"]) {
			return fmt.Errorf("Error reading wan_opt: %v", err)
		}
	}

	if err = d.Set("wireless_activity", flattenLogDiskFilterWirelessActivity(o["wireless-activity"], d, "wireless_activity", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-activity"]) {
			return fmt.Errorf("Error reading wireless_activity: %v", err)
		}
	}

	if err = d.Set("cpu_memory_usage", flattenLogDiskFilterCpuMemoryUsage(o["cpu-memory-usage"], d, "cpu_memory_usage", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-memory-usage"]) {
			return fmt.Errorf("Error reading cpu_memory_usage: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogDiskFilterFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogDiskFilterFilterType(o["filter-type"], d, "filter_type", sv)); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogDiskFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogDiskFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterDlpArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterGtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandLogDiskFilterFreeStyleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandLogDiskFilterFreeStyleCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter"], _ = expandLogDiskFilterFreeStyleFilter(d, i["filter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-type"], _ = expandLogDiskFilterFreeStyleFilterType(d, i["filter_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogDiskFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSystem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterRadius(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterIpsec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterDhcp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterPpp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSslvpnLogAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSslvpnLogAdm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterSslvpnLogSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterVipSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterLdbMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterWanOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterWirelessActivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterCpuMemoryUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskFilterFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogDiskFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {

			t, err := expandLogDiskFilterSeverity(d, v, "severity", sv)
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

			t, err := expandLogDiskFilterForwardTraffic(d, v, "forward_traffic", sv)
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

			t, err := expandLogDiskFilterLocalTraffic(d, v, "local_traffic", sv)
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

			t, err := expandLogDiskFilterMulticastTraffic(d, v, "multicast_traffic", sv)
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

			t, err := expandLogDiskFilterSnifferTraffic(d, v, "sniffer_traffic", sv)
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

			t, err := expandLogDiskFilterZtnaTraffic(d, v, "ztna_traffic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ztna-traffic"] = t
			}
		}
	}

	if v, ok := d.GetOk("anomaly"); ok {
		if setArgNil {
			obj["anomaly"] = nil
		} else {

			t, err := expandLogDiskFilterAnomaly(d, v, "anomaly", sv)
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

			t, err := expandLogDiskFilterNetscanDiscovery(d, v, "netscan_discovery", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["netscan-discovery"] = t
			}
		}
	}

	if v, ok := d.GetOk("netscan_vulnerability"); ok {
		if setArgNil {
			obj["netscan-vulnerability"] = nil
		} else {

			t, err := expandLogDiskFilterNetscanVulnerability(d, v, "netscan_vulnerability", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["netscan-vulnerability"] = t
			}
		}
	}

	if v, ok := d.GetOk("voip"); ok {
		if setArgNil {
			obj["voip"] = nil
		} else {

			t, err := expandLogDiskFilterVoip(d, v, "voip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["voip"] = t
			}
		}
	}

	if v, ok := d.GetOk("dlp_archive"); ok {
		if setArgNil {
			obj["dlp-archive"] = nil
		} else {

			t, err := expandLogDiskFilterDlpArchive(d, v, "dlp_archive", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dlp-archive"] = t
			}
		}
	}

	if v, ok := d.GetOk("gtp"); ok {
		if setArgNil {
			obj["gtp"] = nil
		} else {

			t, err := expandLogDiskFilterGtp(d, v, "gtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gtp"] = t
			}
		}
	}

	if v, ok := d.GetOk("free_style"); ok {
		if setArgNil {
			obj["free-style"] = make([]struct{}, 0)
		} else {

			t, err := expandLogDiskFilterFreeStyle(d, v, "free_style", sv)
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

			t, err := expandLogDiskFilterDns(d, v, "dns", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		if setArgNil {
			obj["ssh"] = nil
		} else {

			t, err := expandLogDiskFilterSsh(d, v, "ssh", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh"] = t
			}
		}
	}

	if v, ok := d.GetOk("event"); ok {
		if setArgNil {
			obj["event"] = nil
		} else {

			t, err := expandLogDiskFilterEvent(d, v, "event", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["event"] = t
			}
		}
	}

	if v, ok := d.GetOk("system"); ok {
		if setArgNil {
			obj["system"] = nil
		} else {

			t, err := expandLogDiskFilterSystem(d, v, "system", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["system"] = t
			}
		}
	}

	if v, ok := d.GetOk("radius"); ok {
		if setArgNil {
			obj["radius"] = nil
		} else {

			t, err := expandLogDiskFilterRadius(d, v, "radius", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radius"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec"); ok {
		if setArgNil {
			obj["ipsec"] = nil
		} else {

			t, err := expandLogDiskFilterIpsec(d, v, "ipsec", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp"); ok {
		if setArgNil {
			obj["dhcp"] = nil
		} else {

			t, err := expandLogDiskFilterDhcp(d, v, "dhcp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp"] = t
			}
		}
	}

	if v, ok := d.GetOk("ppp"); ok {
		if setArgNil {
			obj["ppp"] = nil
		} else {

			t, err := expandLogDiskFilterPpp(d, v, "ppp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ppp"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin"); ok {
		if setArgNil {
			obj["admin"] = nil
		} else {

			t, err := expandLogDiskFilterAdmin(d, v, "admin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha"); ok {
		if setArgNil {
			obj["ha"] = nil
		} else {

			t, err := expandLogDiskFilterHa(d, v, "ha", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		if setArgNil {
			obj["auth"] = nil
		} else {

			t, err := expandLogDiskFilterAuth(d, v, "auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("pattern"); ok {
		if setArgNil {
			obj["pattern"] = nil
		} else {

			t, err := expandLogDiskFilterPattern(d, v, "pattern", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pattern"] = t
			}
		}
	}

	if v, ok := d.GetOk("sslvpn_log_auth"); ok {
		if setArgNil {
			obj["sslvpn-log-auth"] = nil
		} else {

			t, err := expandLogDiskFilterSslvpnLogAuth(d, v, "sslvpn_log_auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-log-auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("sslvpn_log_adm"); ok {
		if setArgNil {
			obj["sslvpn-log-adm"] = nil
		} else {

			t, err := expandLogDiskFilterSslvpnLogAdm(d, v, "sslvpn_log_adm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-log-adm"] = t
			}
		}
	}

	if v, ok := d.GetOk("sslvpn_log_session"); ok {
		if setArgNil {
			obj["sslvpn-log-session"] = nil
		} else {

			t, err := expandLogDiskFilterSslvpnLogSession(d, v, "sslvpn_log_session", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sslvpn-log-session"] = t
			}
		}
	}

	if v, ok := d.GetOk("vip_ssl"); ok {
		if setArgNil {
			obj["vip-ssl"] = nil
		} else {

			t, err := expandLogDiskFilterVipSsl(d, v, "vip_ssl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vip-ssl"] = t
			}
		}
	}

	if v, ok := d.GetOk("ldb_monitor"); ok {
		if setArgNil {
			obj["ldb-monitor"] = nil
		} else {

			t, err := expandLogDiskFilterLdbMonitor(d, v, "ldb_monitor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ldb-monitor"] = t
			}
		}
	}

	if v, ok := d.GetOk("wan_opt"); ok {
		if setArgNil {
			obj["wan-opt"] = nil
		} else {

			t, err := expandLogDiskFilterWanOpt(d, v, "wan_opt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wan-opt"] = t
			}
		}
	}

	if v, ok := d.GetOk("wireless_activity"); ok {
		if setArgNil {
			obj["wireless-activity"] = nil
		} else {

			t, err := expandLogDiskFilterWirelessActivity(d, v, "wireless_activity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wireless-activity"] = t
			}
		}
	}

	if v, ok := d.GetOk("cpu_memory_usage"); ok {
		if setArgNil {
			obj["cpu-memory-usage"] = nil
		} else {

			t, err := expandLogDiskFilterCpuMemoryUsage(d, v, "cpu_memory_usage", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cpu-memory-usage"] = t
			}
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		if setArgNil {
			obj["filter"] = nil
		} else {

			t, err := expandLogDiskFilterFilter(d, v, "filter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter"] = t
			}
		}
	}

	if v, ok := d.GetOk("filter_type"); ok {
		if setArgNil {
			obj["filter-type"] = nil
		} else {

			t, err := expandLogDiskFilterFilterType(d, v, "filter_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-type"] = t
			}
		}
	}

	return &obj, nil
}
