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

func resourceLogFortianalyzerOverrideFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzerOverrideFilterUpdate,
		Read:   resourceLogFortianalyzerOverrideFilterRead,
		Update: resourceLogFortianalyzerOverrideFilterUpdate,
		Delete: resourceLogFortianalyzerOverrideFilterDelete,

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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLogFortianalyzerOverrideFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzerOverrideFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzerOverrideFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzerOverrideFilter")
	}

	return resourceLogFortianalyzerOverrideFilterRead(d, m)
}

func resourceLogFortianalyzerOverrideFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzerOverrideFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzerOverrideFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzerOverrideFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzerOverrideFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogFortianalyzerOverrideFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerOverrideFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzerOverrideFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerOverrideFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzerOverrideFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterDlpArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterGtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenLogFortianalyzerOverrideFilterFreeStyleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenLogFortianalyzerOverrideFilterFreeStyleCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {

			tmp["filter"] = flattenLogFortianalyzerOverrideFilterFreeStyleFilter(i["filter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {

			tmp["filter_type"] = flattenLogFortianalyzerOverrideFilterFreeStyleFilterType(i["filter-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogFortianalyzerOverrideFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterSsh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideFilterFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzerOverrideFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("severity", flattenLogFortianalyzerOverrideFilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogFortianalyzerOverrideFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogFortianalyzerOverrideFilterLocalTraffic(o["local-traffic"], d, "local_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogFortianalyzerOverrideFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogFortianalyzerOverrideFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogFortianalyzerOverrideFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-traffic"]) {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogFortianalyzerOverrideFilterAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogFortianalyzerOverrideFilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogFortianalyzerOverrideFilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability", sv)); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogFortianalyzerOverrideFilterVoip(o["voip"], d, "voip", sv)); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("dlp_archive", flattenLogFortianalyzerOverrideFilterDlpArchive(o["dlp-archive"], d, "dlp_archive", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-archive"]) {
			return fmt.Errorf("Error reading dlp_archive: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogFortianalyzerOverrideFilterGtp(o["gtp"], d, "gtp", sv)); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogFortianalyzerOverrideFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
			if !fortiAPIPatch(o["free-style"]) {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogFortianalyzerOverrideFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
				if !fortiAPIPatch(o["free-style"]) {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns", flattenLogFortianalyzerOverrideFilterDns(o["dns"], d, "dns", sv)); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogFortianalyzerOverrideFilterSsh(o["ssh"], d, "ssh", sv)); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogFortianalyzerOverrideFilterFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogFortianalyzerOverrideFilterFilterType(o["filter-type"], d, "filter_type", sv)); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzerOverrideFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogFortianalyzerOverrideFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterDlpArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterGtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandLogFortianalyzerOverrideFilterFreeStyleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandLogFortianalyzerOverrideFilterFreeStyleCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter"], _ = expandLogFortianalyzerOverrideFilterFreeStyleFilter(d, i["filter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-type"], _ = expandLogFortianalyzerOverrideFilterFreeStyleFilterType(d, i["filter_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogFortianalyzerOverrideFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideFilterFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzerOverrideFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideFilterSeverity(d, v, "severity", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterForwardTraffic(d, v, "forward_traffic", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterLocalTraffic(d, v, "local_traffic", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterMulticastTraffic(d, v, "multicast_traffic", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterSnifferTraffic(d, v, "sniffer_traffic", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterZtnaTraffic(d, v, "ztna_traffic", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterAnomaly(d, v, "anomaly", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterNetscanDiscovery(d, v, "netscan_discovery", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterNetscanVulnerability(d, v, "netscan_vulnerability", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterVoip(d, v, "voip", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterDlpArchive(d, v, "dlp_archive", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterGtp(d, v, "gtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gtp"] = t
			}
		}
	}

	if v, ok := d.GetOk("free_style"); ok || d.HasChange("free_style") {
		if setArgNil {
			obj["free-style"] = make([]struct{}, 0)
		} else {

			t, err := expandLogFortianalyzerOverrideFilterFreeStyle(d, v, "free_style", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterDns(d, v, "dns", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterSsh(d, v, "ssh", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh"] = t
			}
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		if setArgNil {
			obj["filter"] = nil
		} else {

			t, err := expandLogFortianalyzerOverrideFilterFilter(d, v, "filter", sv)
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

			t, err := expandLogFortianalyzerOverrideFilterFilterType(d, v, "filter_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-type"] = t
			}
		}
	}

	return &obj, nil
}
