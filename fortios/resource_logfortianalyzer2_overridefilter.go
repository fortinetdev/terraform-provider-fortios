// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Override filters for FortiAnalyzer.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzer2OverrideFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer2OverrideFilterUpdate,
		Read:   resourceLogFortianalyzer2OverrideFilterRead,
		Update: resourceLogFortianalyzer2OverrideFilterUpdate,
		Delete: resourceLogFortianalyzer2OverrideFilterDelete,

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
			"forti_switch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"debug": &schema.Schema{
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

func resourceLogFortianalyzer2OverrideFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogFortianalyzer2OverrideFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer2OverrideFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzer2OverrideFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer2OverrideFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzer2OverrideFilter")
	}

	return resourceLogFortianalyzer2OverrideFilterRead(d, m)
}

func resourceLogFortianalyzer2OverrideFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzer2OverrideFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer2OverrideFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzer2OverrideFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzer2OverrideFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer2OverrideFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortianalyzer2OverrideFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer2OverrideFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer2OverrideFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer2OverrideFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer2OverrideFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterDlpArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterGtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterDebug(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenLogFortianalyzer2OverrideFilterFreeStyleId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["category"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
			}
			tmp["category"] = flattenLogFortianalyzer2OverrideFilterFreeStyleCategory(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["filter"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
			}
			tmp["filter"] = flattenLogFortianalyzer2OverrideFilterFreeStyleFilter(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["filter-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
			}
			tmp["filter_type"] = flattenLogFortianalyzer2OverrideFilterFreeStyleFilterType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogFortianalyzer2OverrideFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzer2OverrideFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterSsh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer2OverrideFilterFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer2OverrideFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("severity", flattenLogFortianalyzer2OverrideFilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogFortianalyzer2OverrideFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogFortianalyzer2OverrideFilterLocalTraffic(o["local-traffic"], d, "local_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogFortianalyzer2OverrideFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogFortianalyzer2OverrideFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogFortianalyzer2OverrideFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-traffic"]) {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogFortianalyzer2OverrideFilterHttpTransaction(o["http-transaction"], d, "http_transaction", sv)); err != nil {
		if !fortiAPIPatch(o["http-transaction"]) {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogFortianalyzer2OverrideFilterAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogFortianalyzer2OverrideFilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogFortianalyzer2OverrideFilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogFortianalyzer2OverrideFilterVoip(o["voip"], d, "voip", sv)); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("dlp_archive", flattenLogFortianalyzer2OverrideFilterDlpArchive(o["dlp-archive"], d, "dlp_archive", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-archive"]) {
			return fmt.Errorf("Error reading dlp_archive: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogFortianalyzer2OverrideFilterGtp(o["gtp"], d, "gtp", sv)); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogFortianalyzer2OverrideFilterFortiSwitch(o["forti-switch"], d, "forti_switch", sv)); err != nil {
		if !fortiAPIPatch(o["forti-switch"]) {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("debug", flattenLogFortianalyzer2OverrideFilterDebug(o["debug"], d, "debug", sv)); err != nil {
		if !fortiAPIPatch(o["debug"]) {
			return fmt.Errorf("Error reading debug: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("free_style", flattenLogFortianalyzer2OverrideFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
			if !fortiAPIPatch(o["free-style"]) {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogFortianalyzer2OverrideFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
				if !fortiAPIPatch(o["free-style"]) {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns", flattenLogFortianalyzer2OverrideFilterDns(o["dns"], d, "dns", sv)); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogFortianalyzer2OverrideFilterSsh(o["ssh"], d, "ssh", sv)); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogFortianalyzer2OverrideFilterFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogFortianalyzer2OverrideFilterFilterType(o["filter-type"], d, "filter_type", sv)); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer2OverrideFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogFortianalyzer2OverrideFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterDlpArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterGtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterDebug(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandLogFortianalyzer2OverrideFilterFreeStyleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandLogFortianalyzer2OverrideFilterFreeStyleCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter"], _ = expandLogFortianalyzer2OverrideFilterFreeStyleFilter(d, i["filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-type"], _ = expandLogFortianalyzer2OverrideFilterFreeStyleFilterType(d, i["filter_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogFortianalyzer2OverrideFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer2OverrideFilterFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer2OverrideFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {
			t, err := expandLogFortianalyzer2OverrideFilterSeverity(d, v, "severity", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterForwardTraffic(d, v, "forward_traffic", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterLocalTraffic(d, v, "local_traffic", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterMulticastTraffic(d, v, "multicast_traffic", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterSnifferTraffic(d, v, "sniffer_traffic", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterZtnaTraffic(d, v, "ztna_traffic", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterHttpTransaction(d, v, "http_transaction", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterAnomaly(d, v, "anomaly", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterNetscanDiscovery(d, v, "netscan_discovery", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterNetscanVulnerability(d, v, "netscan_vulnerability", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterVoip(d, v, "voip", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterDlpArchive(d, v, "dlp_archive", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterGtp(d, v, "gtp", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterFortiSwitch(d, v, "forti_switch", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forti-switch"] = t
			}
		}
	}

	if v, ok := d.GetOk("debug"); ok {
		if setArgNil {
			obj["debug"] = nil
		} else {
			t, err := expandLogFortianalyzer2OverrideFilterDebug(d, v, "debug", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["debug"] = t
			}
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		if setArgNil {
			obj["free-style"] = make([]struct{}, 0)
		} else {
			t, err := expandLogFortianalyzer2OverrideFilterFreeStyle(d, v, "free_style", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterDns(d, v, "dns", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterSsh(d, v, "ssh", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh"] = t
			}
		}
	} else if d.HasChange("ssh") {
		obj["ssh"] = nil
	}

	if v, ok := d.GetOk("filter"); ok {
		if setArgNil {
			obj["filter"] = nil
		} else {
			t, err := expandLogFortianalyzer2OverrideFilterFilter(d, v, "filter", sv)
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
			t, err := expandLogFortianalyzer2OverrideFilterFilterType(d, v, "filter_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-type"] = t
			}
		}
	}

	return &obj, nil
}
