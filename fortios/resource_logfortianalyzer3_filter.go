// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Filters for FortiAnalyzer.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzer3Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer3FilterUpdate,
		Read:   resourceLogFortianalyzer3FilterRead,
		Update: resourceLogFortianalyzer3FilterUpdate,
		Delete: resourceLogFortianalyzer3FilterDelete,

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

func resourceLogFortianalyzer3FilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogFortianalyzer3Filter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3Filter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzer3Filter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzer3Filter")
	}

	return resourceLogFortianalyzer3FilterRead(d, m)
}

func resourceLogFortianalyzer3FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzer3Filter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzer3Filter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzer3Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer3FilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogFortianalyzer3Filter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer3Filter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3Filter resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer3FilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterHttpTransaction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterDlpArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterGtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFortiSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterDebug(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFreeStyle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenLogFortianalyzer3FilterFreeStyleId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["category"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
			}
			tmp["category"] = flattenLogFortianalyzer3FilterFreeStyleCategory(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["filter"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
			}
			tmp["filter"] = flattenLogFortianalyzer3FilterFreeStyleFilter(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["filter-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
			}
			tmp["filter_type"] = flattenLogFortianalyzer3FilterFreeStyleFilterType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogFortianalyzer3FilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLogFortianalyzer3FilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterSsh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3FilterFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer3Filter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("severity", flattenLogFortianalyzer3FilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogFortianalyzer3FilterForwardTraffic(o["forward-traffic"], d, "forward_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogFortianalyzer3FilterLocalTraffic(o["local-traffic"], d, "local_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogFortianalyzer3FilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogFortianalyzer3FilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogFortianalyzer3FilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-traffic"]) {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	if err = d.Set("http_transaction", flattenLogFortianalyzer3FilterHttpTransaction(o["http-transaction"], d, "http_transaction", sv)); err != nil {
		if !fortiAPIPatch(o["http-transaction"]) {
			return fmt.Errorf("Error reading http_transaction: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogFortianalyzer3FilterAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogFortianalyzer3FilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogFortianalyzer3FilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogFortianalyzer3FilterVoip(o["voip"], d, "voip", sv)); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("dlp_archive", flattenLogFortianalyzer3FilterDlpArchive(o["dlp-archive"], d, "dlp_archive", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-archive"]) {
			return fmt.Errorf("Error reading dlp_archive: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogFortianalyzer3FilterGtp(o["gtp"], d, "gtp", sv)); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("forti_switch", flattenLogFortianalyzer3FilterFortiSwitch(o["forti-switch"], d, "forti_switch", sv)); err != nil {
		if !fortiAPIPatch(o["forti-switch"]) {
			return fmt.Errorf("Error reading forti_switch: %v", err)
		}
	}

	if err = d.Set("debug", flattenLogFortianalyzer3FilterDebug(o["debug"], d, "debug", sv)); err != nil {
		if !fortiAPIPatch(o["debug"]) {
			return fmt.Errorf("Error reading debug: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("free_style", flattenLogFortianalyzer3FilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
			if !fortiAPIPatch(o["free-style"]) {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogFortianalyzer3FilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
				if !fortiAPIPatch(o["free-style"]) {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns", flattenLogFortianalyzer3FilterDns(o["dns"], d, "dns", sv)); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogFortianalyzer3FilterSsh(o["ssh"], d, "ssh", sv)); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogFortianalyzer3FilterFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogFortianalyzer3FilterFilterType(o["filter-type"], d, "filter_type", sv)); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer3FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogFortianalyzer3FilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterHttpTransaction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterDlpArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterGtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFortiSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterDebug(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandLogFortianalyzer3FilterFreeStyleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandLogFortianalyzer3FilterFreeStyleCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter"], _ = expandLogFortianalyzer3FilterFreeStyleFilter(d, i["filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-type"], _ = expandLogFortianalyzer3FilterFreeStyleFilterType(d, i["filter_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogFortianalyzer3FilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3FilterFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer3Filter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {
			t, err := expandLogFortianalyzer3FilterSeverity(d, v, "severity", sv)
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
			t, err := expandLogFortianalyzer3FilterForwardTraffic(d, v, "forward_traffic", sv)
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
			t, err := expandLogFortianalyzer3FilterLocalTraffic(d, v, "local_traffic", sv)
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
			t, err := expandLogFortianalyzer3FilterMulticastTraffic(d, v, "multicast_traffic", sv)
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
			t, err := expandLogFortianalyzer3FilterSnifferTraffic(d, v, "sniffer_traffic", sv)
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
			t, err := expandLogFortianalyzer3FilterZtnaTraffic(d, v, "ztna_traffic", sv)
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
			t, err := expandLogFortianalyzer3FilterHttpTransaction(d, v, "http_transaction", sv)
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
			t, err := expandLogFortianalyzer3FilterAnomaly(d, v, "anomaly", sv)
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
			t, err := expandLogFortianalyzer3FilterNetscanDiscovery(d, v, "netscan_discovery", sv)
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
			t, err := expandLogFortianalyzer3FilterNetscanVulnerability(d, v, "netscan_vulnerability", sv)
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
			t, err := expandLogFortianalyzer3FilterVoip(d, v, "voip", sv)
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
			t, err := expandLogFortianalyzer3FilterDlpArchive(d, v, "dlp_archive", sv)
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
			t, err := expandLogFortianalyzer3FilterGtp(d, v, "gtp", sv)
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
			t, err := expandLogFortianalyzer3FilterFortiSwitch(d, v, "forti_switch", sv)
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
			t, err := expandLogFortianalyzer3FilterDebug(d, v, "debug", sv)
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
			t, err := expandLogFortianalyzer3FilterFreeStyle(d, v, "free_style", sv)
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
			t, err := expandLogFortianalyzer3FilterDns(d, v, "dns", sv)
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
			t, err := expandLogFortianalyzer3FilterSsh(d, v, "ssh", sv)
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
			t, err := expandLogFortianalyzer3FilterFilter(d, v, "filter", sv)
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
			t, err := expandLogFortianalyzer3FilterFilterType(d, v, "filter_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-type"] = t
			}
		}
	}

	return &obj, nil
}
