// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Filters for FortiAnalyzer Cloud.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogFortianalyzerCloudFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzerCloudFilterUpdate,
		Read:   resourceLogFortianalyzerCloudFilterRead,
		Update: resourceLogFortianalyzerCloudFilterUpdate,
		Delete: resourceLogFortianalyzerCloudFilterDelete,

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

func resourceLogFortianalyzerCloudFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzerCloudFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerCloudFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzerCloudFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerCloudFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzerCloudFilter")
	}

	return resourceLogFortianalyzerCloudFilterRead(d, m)
}

func resourceLogFortianalyzerCloudFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzerCloudFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerCloudFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogFortianalyzerCloudFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogFortianalyzerCloudFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzerCloudFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogFortianalyzerCloudFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerCloudFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzerCloudFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerCloudFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzerCloudFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterZtnaTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterVoip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterDlpArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterGtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterFreeStyle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenLogFortianalyzerCloudFilterFreeStyleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenLogFortianalyzerCloudFilterFreeStyleCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {

			tmp["filter"] = flattenLogFortianalyzerCloudFilterFreeStyleFilter(i["filter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := i["filter-type"]; ok {

			tmp["filter_type"] = flattenLogFortianalyzerCloudFilterFreeStyleFilterType(i["filter-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenLogFortianalyzerCloudFilterFreeStyleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterFreeStyleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterFreeStyleFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterFreeStyleFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzerCloudFilterFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzerCloudFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("severity", flattenLogFortianalyzerCloudFilterSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogFortianalyzerCloudFilterForwardTraffic(o["forward-traffic"], d, "forward_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogFortianalyzerCloudFilterLocalTraffic(o["local-traffic"], d, "local_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogFortianalyzerCloudFilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogFortianalyzerCloudFilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("ztna_traffic", flattenLogFortianalyzerCloudFilterZtnaTraffic(o["ztna-traffic"], d, "ztna_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-traffic"]) {
			return fmt.Errorf("Error reading ztna_traffic: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogFortianalyzerCloudFilterAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogFortianalyzerCloudFilterVoip(o["voip"], d, "voip", sv)); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("dlp_archive", flattenLogFortianalyzerCloudFilterDlpArchive(o["dlp-archive"], d, "dlp_archive", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-archive"]) {
			return fmt.Errorf("Error reading dlp_archive: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogFortianalyzerCloudFilterGtp(o["gtp"], d, "gtp", sv)); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("free_style", flattenLogFortianalyzerCloudFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
			if !fortiAPIPatch(o["free-style"]) {
				return fmt.Errorf("Error reading free_style: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("free_style"); ok {
			if err = d.Set("free_style", flattenLogFortianalyzerCloudFilterFreeStyle(o["free-style"], d, "free_style", sv)); err != nil {
				if !fortiAPIPatch(o["free-style"]) {
					return fmt.Errorf("Error reading free_style: %v", err)
				}
			}
		}
	}

	if err = d.Set("filter", flattenLogFortianalyzerCloudFilterFilter(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogFortianalyzerCloudFilterFilterType(o["filter-type"], d, "filter_type", sv)); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzerCloudFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogFortianalyzerCloudFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterZtnaTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterVoip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterDlpArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterGtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterFreeStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandLogFortianalyzerCloudFilterFreeStyleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandLogFortianalyzerCloudFilterFreeStyleCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter"], _ = expandLogFortianalyzerCloudFilterFreeStyleFilter(d, i["filter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter-type"], _ = expandLogFortianalyzerCloudFilterFreeStyleFilterType(d, i["filter_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogFortianalyzerCloudFilterFreeStyleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterFreeStyleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterFreeStyleFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterFreeStyleFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerCloudFilterFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzerCloudFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("severity"); ok {
		if setArgNil {
			obj["severity"] = nil
		} else {

			t, err := expandLogFortianalyzerCloudFilterSeverity(d, v, "severity", sv)
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

			t, err := expandLogFortianalyzerCloudFilterForwardTraffic(d, v, "forward_traffic", sv)
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

			t, err := expandLogFortianalyzerCloudFilterLocalTraffic(d, v, "local_traffic", sv)
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

			t, err := expandLogFortianalyzerCloudFilterMulticastTraffic(d, v, "multicast_traffic", sv)
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

			t, err := expandLogFortianalyzerCloudFilterSnifferTraffic(d, v, "sniffer_traffic", sv)
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

			t, err := expandLogFortianalyzerCloudFilterZtnaTraffic(d, v, "ztna_traffic", sv)
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

			t, err := expandLogFortianalyzerCloudFilterAnomaly(d, v, "anomaly", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["anomaly"] = t
			}
		}
	}

	if v, ok := d.GetOk("voip"); ok {
		if setArgNil {
			obj["voip"] = nil
		} else {

			t, err := expandLogFortianalyzerCloudFilterVoip(d, v, "voip", sv)
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

			t, err := expandLogFortianalyzerCloudFilterDlpArchive(d, v, "dlp_archive", sv)
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

			t, err := expandLogFortianalyzerCloudFilterGtp(d, v, "gtp", sv)
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

			t, err := expandLogFortianalyzerCloudFilterFreeStyle(d, v, "free_style", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["free-style"] = t
			}
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		if setArgNil {
			obj["filter"] = nil
		} else {

			t, err := expandLogFortianalyzerCloudFilterFilter(d, v, "filter", sv)
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

			t, err := expandLogFortianalyzerCloudFilterFilterType(d, v, "filter_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["filter-type"] = t
			}
		}
	}

	return &obj, nil
}
