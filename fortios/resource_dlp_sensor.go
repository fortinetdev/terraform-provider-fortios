// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DLP sensors.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"match_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eval": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),
							Optional:     true,
						},
						"dictionary": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"feature_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
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
						},
						"sensitivity": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
								},
							},
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
						},
						"regexp": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"full_archive_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceDlpSensorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpSensor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating DlpSensor resource while getting object: %v", err)
	}

	o, err := c.CreateDlpSensor(obj, vdomparam)

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

	obj, err := getObjectDlpSensor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DlpSensor resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpSensor(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteDlpSensor(mkey, vdomparam)
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

	o, err := c.ReadDlpSensor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSensor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DlpSensor resource from API: %v", err)
	}
	return nil
}

func flattenDlpSensorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorMatchType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorEval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenDlpSensorEntriesId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dictionary"
		if cur_v, ok := i["dictionary"]; ok {
			tmp["dictionary"] = flattenDlpSensorEntriesDictionary(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "count"
		if cur_v, ok := i["count"]; ok {
			tmp["count"] = flattenDlpSensorEntriesCount(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenDlpSensorEntriesStatus(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenDlpSensorEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSensorEntriesDictionary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorEntriesCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSensorEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFeatureSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenDlpSensorFilterId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenDlpSensorFilterName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if cur_v, ok := i["severity"]; ok {
			tmp["severity"] = flattenDlpSensorFilterSeverity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenDlpSensorFilterType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if cur_v, ok := i["proto"]; ok {
			tmp["proto"] = flattenDlpSensorFilterProto(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if cur_v, ok := i["filter-by"]; ok {
			tmp["filter_by"] = flattenDlpSensorFilterFilterBy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if cur_v, ok := i["file-size"]; ok {
			tmp["file_size"] = flattenDlpSensorFilterFileSize(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if cur_v, ok := i["company-identifier"]; ok {
			tmp["company_identifier"] = flattenDlpSensorFilterCompanyIdentifier(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if cur_v, ok := i["sensitivity"]; ok {
			tmp["sensitivity"] = flattenDlpSensorFilterSensitivity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fp_sensitivity"
		if cur_v, ok := i["fp-sensitivity"]; ok {
			tmp["fp_sensitivity"] = flattenDlpSensorFilterFpSensitivity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if cur_v, ok := i["match-percentage"]; ok {
			tmp["match_percentage"] = flattenDlpSensorFilterMatchPercentage(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if cur_v, ok := i["file-type"]; ok {
			tmp["file_type"] = flattenDlpSensorFilterFileType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if cur_v, ok := i["regexp"]; ok {
			tmp["regexp"] = flattenDlpSensorFilterRegexp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if cur_v, ok := i["archive"]; ok {
			tmp["archive"] = flattenDlpSensorFilterArchive(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenDlpSensorFilterAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if cur_v, ok := i["expiry"]; ok {
			tmp["expiry"] = flattenDlpSensorFilterExpiry(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenDlpSensorFilterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSensorFilterName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterFilterBy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterFileSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSensorFilterCompanyIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenDlpSensorFilterSensitivityName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenDlpSensorFilterSensitivityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterFpSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenDlpSensorFilterFpSensitivityName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenDlpSensorFilterFpSensitivityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterMatchPercentage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSensorFilterFileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpSensorFilterRegexp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFilterExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorDlpLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorExtendedLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorNacQuarLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFlowBased(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorFullArchiveProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpSensorSummaryProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDlpSensor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenDlpSensorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("match_type", flattenDlpSensorMatchType(o["match-type"], d, "match_type", sv)); err != nil {
		if !fortiAPIPatch(o["match-type"]) {
			return fmt.Errorf("Error reading match_type: %v", err)
		}
	}

	if err = d.Set("eval", flattenDlpSensorEval(o["eval"], d, "eval", sv)); err != nil {
		if !fortiAPIPatch(o["eval"]) {
			return fmt.Errorf("Error reading eval: %v", err)
		}
	}

	if err = d.Set("comment", flattenDlpSensorComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("entries", flattenDlpSensorEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenDlpSensorEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("feature_set", flattenDlpSensorFeatureSet(o["feature-set"], d, "feature_set", sv)); err != nil {
		if !fortiAPIPatch(o["feature-set"]) {
			return fmt.Errorf("Error reading feature_set: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenDlpSensorReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("filter", flattenDlpSensorFilter(o["filter"], d, "filter", sv)); err != nil {
			if !fortiAPIPatch(o["filter"]) {
				return fmt.Errorf("Error reading filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("filter"); ok {
			if err = d.Set("filter", flattenDlpSensorFilter(o["filter"], d, "filter", sv)); err != nil {
				if !fortiAPIPatch(o["filter"]) {
					return fmt.Errorf("Error reading filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("dlp_log", flattenDlpSensorDlpLog(o["dlp-log"], d, "dlp_log", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-log"]) {
			return fmt.Errorf("Error reading dlp_log: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenDlpSensorExtendedLog(o["extended-log"], d, "extended_log", sv)); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("nac_quar_log", flattenDlpSensorNacQuarLog(o["nac-quar-log"], d, "nac_quar_log", sv)); err != nil {
		if !fortiAPIPatch(o["nac-quar-log"]) {
			return fmt.Errorf("Error reading nac_quar_log: %v", err)
		}
	}

	if err = d.Set("flow_based", flattenDlpSensorFlowBased(o["flow-based"], d, "flow_based", sv)); err != nil {
		if !fortiAPIPatch(o["flow-based"]) {
			return fmt.Errorf("Error reading flow_based: %v", err)
		}
	}

	if err = d.Set("options", flattenDlpSensorOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("full_archive_proto", flattenDlpSensorFullArchiveProto(o["full-archive-proto"], d, "full_archive_proto", sv)); err != nil {
		if !fortiAPIPatch(o["full-archive-proto"]) {
			return fmt.Errorf("Error reading full_archive_proto: %v", err)
		}
	}

	if err = d.Set("summary_proto", flattenDlpSensorSummaryProto(o["summary-proto"], d, "summary_proto", sv)); err != nil {
		if !fortiAPIPatch(o["summary-proto"]) {
			return fmt.Errorf("Error reading summary_proto: %v", err)
		}
	}

	return nil
}

func flattenDlpSensorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDlpSensorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorMatchType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandDlpSensorEntriesId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dictionary"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dictionary"], _ = expandDlpSensorEntriesDictionary(d, i["dictionary"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dictionary"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["count"], _ = expandDlpSensorEntriesCount(d, i["count"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandDlpSensorEntriesStatus(d, i["status"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSensorEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesDictionary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFeatureSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandDlpSensorFilterId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandDlpSensorFilterName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandDlpSensorFilterSeverity(d, i["severity"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandDlpSensorFilterType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["proto"], _ = expandDlpSensorFilterProto(d, i["proto"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["proto"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_by"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-by"], _ = expandDlpSensorFilterFilterBy(d, i["filter_by"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_size"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file-size"], _ = expandDlpSensorFilterFileSize(d, i["file_size"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "company_identifier"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["company-identifier"], _ = expandDlpSensorFilterCompanyIdentifier(d, i["company_identifier"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["company-identifier"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sensitivity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sensitivity"], _ = expandDlpSensorFilterSensitivity(d, i["sensitivity"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sensitivity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fp_sensitivity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fp-sensitivity"], _ = expandDlpSensorFilterFpSensitivity(d, i["fp_sensitivity"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["fp-sensitivity"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_percentage"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-percentage"], _ = expandDlpSensorFilterMatchPercentage(d, i["match_percentage"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file-type"], _ = expandDlpSensorFilterFileType(d, i["file_type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["file-type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regexp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regexp"], _ = expandDlpSensorFilterRegexp(d, i["regexp"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["regexp"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "archive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["archive"], _ = expandDlpSensorFilterArchive(d, i["archive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandDlpSensorFilterAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["expiry"], _ = expandDlpSensorFilterExpiry(d, i["expiry"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSensorFilterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFilterBy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterCompanyIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandDlpSensorFilterSensitivityName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSensorFilterSensitivityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFpSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandDlpSensorFilterFpSensitivityName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDlpSensorFilterFpSensitivityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterMatchPercentage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterRegexp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFilterExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorDlpLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorExtendedLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorNacQuarLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFlowBased(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorFullArchiveProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpSensorSummaryProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSensor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpSensorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("match_type"); ok {
		t, err := expandDlpSensorMatchType(d, v, "match_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-type"] = t
		}
	}

	if v, ok := d.GetOk("eval"); ok {
		t, err := expandDlpSensorEval(d, v, "eval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eval"] = t
		}
	} else if d.HasChange("eval") {
		obj["eval"] = nil
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandDlpSensorComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandDlpSensorEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("feature_set"); ok {
		t, err := expandDlpSensorFeatureSet(d, v, "feature_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["feature-set"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandDlpSensorReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	} else if d.HasChange("replacemsg_group") {
		obj["replacemsg-group"] = nil
	}

	if v, ok := d.GetOk("filter"); ok || d.HasChange("filter") {
		t, err := expandDlpSensorFilter(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("dlp_log"); ok {
		t, err := expandDlpSensorDlpLog(d, v, "dlp_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-log"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandDlpSensorExtendedLog(d, v, "extended_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar_log"); ok {
		t, err := expandDlpSensorNacQuarLog(d, v, "nac_quar_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar-log"] = t
		}
	}

	if v, ok := d.GetOk("flow_based"); ok {
		t, err := expandDlpSensorFlowBased(d, v, "flow_based", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-based"] = t
		}
	} else if d.HasChange("flow_based") {
		obj["flow-based"] = nil
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandDlpSensorOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	} else if d.HasChange("options") {
		obj["options"] = nil
	}

	if v, ok := d.GetOk("full_archive_proto"); ok {
		t, err := expandDlpSensorFullArchiveProto(d, v, "full_archive_proto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-archive-proto"] = t
		}
	} else if d.HasChange("full_archive_proto") {
		obj["full-archive-proto"] = nil
	}

	if v, ok := d.GetOk("summary_proto"); ok {
		t, err := expandDlpSensorSummaryProto(d, v, "summary_proto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-proto"] = t
		}
	} else if d.HasChange("summary_proto") {
		obj["summary-proto"] = nil
	}

	return &obj, nil
}
