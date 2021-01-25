// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure DLP sensors.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDlpSensor() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSensorCreate,
		Read:   resourceDlpSensorRead,
		Update: resourceDlpSensorUpdate,
		Delete: resourceDlpSensorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filter_by": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"file_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"company_identifier": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"fp_sensitivity": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"match_percentage": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"file_type": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"regexp": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"archive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dlp_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_quar_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flow_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_archive_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"summary_proto": &schema.Schema{
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

func resourceDlpSensorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpSensor(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpSensor resource while getting object: %v", err)
	}

	o, err := c.CreateDlpSensor(obj)

	if err != nil {
		return fmt.Errorf("Error creating DlpSensor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpSensor")
	}

	return resourceDlpSensorRead(d, m)
}

func resourceDlpSensorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpSensor(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensor resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpSensor(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpSensor")
	}

	return resourceDlpSensorRead(d, m)
}

func resourceDlpSensorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteDlpSensor(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting DlpSensor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSensorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadDlpSensor(mkey)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSensor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensor resource from API: %v", err)
	}
	return nil
}

func flattenDlpSensorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenDlpSensorFilterId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenDlpSensorFilterName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			tmp["severity"] = flattenDlpSensorFilterSeverity(i["severity"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenDlpSensorFilterType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := i["proto"]; ok {
			tmp["proto"] = flattenDlpSensorFilterProto(i["proto"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := i["filter-by"]; ok {
			tmp["filter_by"] = flattenDlpSensorFilterFilterBy(i["filter-by"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := i["file-size"]; ok {
			tmp["file_size"] = flattenDlpSensorFilterFileSize(i["file-size"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := i["company-identifier"]; ok {
			tmp["company_identifier"] = flattenDlpSensorFilterCompanyIdentifier(i["company-identifier"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fp_sensitivity"
		if _, ok := i["fp-sensitivity"]; ok {
			tmp["fp_sensitivity"] = flattenDlpSensorFilterFpSensitivity(i["fp-sensitivity"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := i["match-percentage"]; ok {
			tmp["match_percentage"] = flattenDlpSensorFilterMatchPercentage(i["match-percentage"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {
			tmp["file_type"] = flattenDlpSensorFilterFileType(i["file-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := i["regexp"]; ok {
			tmp["regexp"] = flattenDlpSensorFilterRegexp(i["regexp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := i["archive"]; ok {
			tmp["archive"] = flattenDlpSensorFilterArchive(i["archive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenDlpSensorFilterAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := i["expiry"]; ok {
			tmp["expiry"] = flattenDlpSensorFilterExpiry(i["expiry"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenDlpSensorFilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFilterBy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterCompanyIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFpSensitivity(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenDlpSensorFilterFpSensitivityName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenDlpSensorFilterFpSensitivityName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterMatchPercentage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterFileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterRegexp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFilterExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorDlpLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorNacQuarLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFlowBased(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorFullArchiveProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSensorSummaryProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpSensor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenDlpSensorName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenDlpSensorComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenDlpSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("filter", flattenDlpSensorFilter(o["filter"], d, "filter")); err != nil {
			if !fortiAPIPatch(o["filter"]) {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter"); ok {
			if err = d.Set("filter", flattenDlpSensorFilter(o["filter"], d, "filter")); err != nil {
				if !fortiAPIPatch(o["filter"]) {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("dlp_log", flattenDlpSensorDlpLog(o["dlp-log"], d, "dlp_log")); err != nil {
		if !fortiAPIPatch(o["dlp-log"]) {
			return fmt.Errorf("Error reading dlp_log: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenDlpSensorExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("nac_quar_log", flattenDlpSensorNacQuarLog(o["nac-quar-log"], d, "nac_quar_log")); err != nil {
		if !fortiAPIPatch(o["nac-quar-log"]) {
			return fmt.Errorf("Error reading nac_quar_log: %v", err)
		}
	}

	if err = d.Set("flow_based", flattenDlpSensorFlowBased(o["flow-based"], d, "flow_based")); err != nil {
		if !fortiAPIPatch(o["flow-based"]) {
			return fmt.Errorf("Error reading flow_based: %v", err)
		}
	}

	if err = d.Set("options", flattenDlpSensorOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("full_archive_proto", flattenDlpSensorFullArchiveProto(o["full-archive-proto"], d, "full_archive_proto")); err != nil {
		if !fortiAPIPatch(o["full-archive-proto"]) {
			return fmt.Errorf("Error reading full_archive_proto: %v", err)
		}
	}

	if err = d.Set("summary_proto", flattenDlpSensorSummaryProto(o["summary-proto"], d, "summary_proto")); err != nil {
		if !fortiAPIPatch(o["summary-proto"]) {
			return fmt.Errorf("Error reading summary_proto: %v", err)
		}
	}

	return nil
}

func flattenDlpSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSensorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandDlpSensorFilterId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandDlpSensorFilterName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandDlpSensorFilterSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandDlpSensorFilterType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["proto"], _ = expandDlpSensorFilterProto(d, i["proto"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-by"], _ = expandDlpSensorFilterFilterBy(d, i["filter_by"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file-size"], _ = expandDlpSensorFilterFileSize(d, i["file_size"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["company-identifier"], _ = expandDlpSensorFilterCompanyIdentifier(d, i["company_identifier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fp_sensitivity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fp-sensitivity"], _ = expandDlpSensorFilterFpSensitivity(d, i["fp_sensitivity"], pre_append)
		} else {
			tmp["fp-sensitivity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-percentage"], _ = expandDlpSensorFilterMatchPercentage(d, i["match_percentage"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file-type"], _ = expandDlpSensorFilterFileType(d, i["file_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regexp"], _ = expandDlpSensorFilterRegexp(d, i["regexp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["archive"], _ = expandDlpSensorFilterArchive(d, i["archive"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandDlpSensorFilterAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["expiry"], _ = expandDlpSensorFilterExpiry(d, i["expiry"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSensorFilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFilterBy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterCompanyIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFpSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandDlpSensorFilterFpSensitivityName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSensorFilterFpSensitivityName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterMatchPercentage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterRegexp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorDlpLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorNacQuarLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFlowBased(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFullArchiveProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorSummaryProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSensor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpSensorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandDlpSensorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandDlpSensorReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		t, err := expandDlpSensorFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("dlp_log"); ok {
		t, err := expandDlpSensorDlpLog(d, v, "dlp_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandDlpSensorExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar_log"); ok {
		t, err := expandDlpSensorNacQuarLog(d, v, "nac_quar_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar-log"] = t
		}
	}

	if v, ok := d.GetOk("flow_based"); ok {
		t, err := expandDlpSensorFlowBased(d, v, "flow_based")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-based"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandDlpSensorOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("full_archive_proto"); ok {
		t, err := expandDlpSensorFullArchiveProto(d, v, "full_archive_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-archive-proto"] = t
		}
	}

	if v, ok := d.GetOk("summary_proto"); ok {
		t, err := expandDlpSensorSummaryProto(d, v, "summary_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-proto"] = t
		}
	}

	return &obj, nil
}
