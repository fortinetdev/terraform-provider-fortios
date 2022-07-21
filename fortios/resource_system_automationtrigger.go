// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Trigger for automation stitches.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutomationTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationTriggerCreate,
		Read:   resourceSystemAutomationTriggerRead,
		Update: resourceSystemAutomationTriggerUpdate,
		Delete: resourceSystemAutomationTriggerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"trigger_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"license_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ioc_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99999),
				Optional:     true,
				Computed:     true,
			},
			"trigger_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trigger_day": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),
				Optional:     true,
				Computed:     true,
			},
			"trigger_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"trigger_minute": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 60),
				Optional:     true,
				Computed:     true,
			},
			"fields": &schema.Schema{
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
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"faz_event_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"faz_event_severity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"faz_event_tags": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"serial": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"fabric_event_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"fabric_event_severity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"logid_block": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemAutomationTriggerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAutomationTrigger(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationTrigger(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationTrigger resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAutomationTrigger(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationTrigger(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationTrigger resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationTrigger")
	}

	return resourceSystemAutomationTriggerRead(d, m)
}

func resourceSystemAutomationTriggerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemAutomationTrigger(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationTrigger resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationTriggerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemAutomationTrigger(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationTrigger(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationTrigger resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationTriggerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerEventType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemAutomationTriggerVdomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemAutomationTriggerVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLicenseType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerIocLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerReportType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLogid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerFrequency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerWeekday(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerHour(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerTriggerMinute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFields(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemAutomationTriggerFieldsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemAutomationTriggerFieldsName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenSystemAutomationTriggerFieldsValue(i["value"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAutomationTriggerFieldsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFieldsValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFazEventName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFazEventSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFazEventTags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerSerial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFabricEventName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerFabricEventSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationTriggerLogidBlock(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemAutomationTriggerLogidBlockId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAutomationTriggerLogidBlockId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutomationTrigger(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemAutomationTriggerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAutomationTriggerDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("trigger_type", flattenSystemAutomationTriggerTriggerType(o["trigger-type"], d, "trigger_type", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-type"]) {
			return fmt.Errorf("Error reading trigger_type: %v", err)
		}
	}

	if err = d.Set("event_type", flattenSystemAutomationTriggerEventType(o["event-type"], d, "event_type", sv)); err != nil {
		if !fortiAPIPatch(o["event-type"]) {
			return fmt.Errorf("Error reading event_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vdom", flattenSystemAutomationTriggerVdom(o["vdom"], d, "vdom", sv)); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemAutomationTriggerVdom(o["vdom"], d, "vdom", sv)); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	if err = d.Set("license_type", flattenSystemAutomationTriggerLicenseType(o["license-type"], d, "license_type", sv)); err != nil {
		if !fortiAPIPatch(o["license-type"]) {
			return fmt.Errorf("Error reading license_type: %v", err)
		}
	}

	if err = d.Set("ioc_level", flattenSystemAutomationTriggerIocLevel(o["ioc-level"], d, "ioc_level", sv)); err != nil {
		if !fortiAPIPatch(o["ioc-level"]) {
			return fmt.Errorf("Error reading ioc_level: %v", err)
		}
	}

	if err = d.Set("report_type", flattenSystemAutomationTriggerReportType(o["report-type"], d, "report_type", sv)); err != nil {
		if !fortiAPIPatch(o["report-type"]) {
			return fmt.Errorf("Error reading report_type: %v", err)
		}
	}

	if err = d.Set("logid", flattenSystemAutomationTriggerLogid(o["logid"], d, "logid", sv)); err != nil {
		if !fortiAPIPatch(o["logid"]) {
			return fmt.Errorf("Error reading logid: %v", err)
		}
	}

	if err = d.Set("trigger_frequency", flattenSystemAutomationTriggerTriggerFrequency(o["trigger-frequency"], d, "trigger_frequency", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-frequency"]) {
			return fmt.Errorf("Error reading trigger_frequency: %v", err)
		}
	}

	if err = d.Set("trigger_weekday", flattenSystemAutomationTriggerTriggerWeekday(o["trigger-weekday"], d, "trigger_weekday", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-weekday"]) {
			return fmt.Errorf("Error reading trigger_weekday: %v", err)
		}
	}

	if err = d.Set("trigger_day", flattenSystemAutomationTriggerTriggerDay(o["trigger-day"], d, "trigger_day", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-day"]) {
			return fmt.Errorf("Error reading trigger_day: %v", err)
		}
	}

	if err = d.Set("trigger_hour", flattenSystemAutomationTriggerTriggerHour(o["trigger-hour"], d, "trigger_hour", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-hour"]) {
			return fmt.Errorf("Error reading trigger_hour: %v", err)
		}
	}

	if err = d.Set("trigger_minute", flattenSystemAutomationTriggerTriggerMinute(o["trigger-minute"], d, "trigger_minute", sv)); err != nil {
		if !fortiAPIPatch(o["trigger-minute"]) {
			return fmt.Errorf("Error reading trigger_minute: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fields", flattenSystemAutomationTriggerFields(o["fields"], d, "fields", sv)); err != nil {
			if !fortiAPIPatch(o["fields"]) {
				return fmt.Errorf("Error reading fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fields"); ok {
			if err = d.Set("fields", flattenSystemAutomationTriggerFields(o["fields"], d, "fields", sv)); err != nil {
				if !fortiAPIPatch(o["fields"]) {
					return fmt.Errorf("Error reading fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("faz_event_name", flattenSystemAutomationTriggerFazEventName(o["faz-event-name"], d, "faz_event_name", sv)); err != nil {
		if !fortiAPIPatch(o["faz-event-name"]) {
			return fmt.Errorf("Error reading faz_event_name: %v", err)
		}
	}

	if err = d.Set("faz_event_severity", flattenSystemAutomationTriggerFazEventSeverity(o["faz-event-severity"], d, "faz_event_severity", sv)); err != nil {
		if !fortiAPIPatch(o["faz-event-severity"]) {
			return fmt.Errorf("Error reading faz_event_severity: %v", err)
		}
	}

	if err = d.Set("faz_event_tags", flattenSystemAutomationTriggerFazEventTags(o["faz-event-tags"], d, "faz_event_tags", sv)); err != nil {
		if !fortiAPIPatch(o["faz-event-tags"]) {
			return fmt.Errorf("Error reading faz_event_tags: %v", err)
		}
	}

	if err = d.Set("serial", flattenSystemAutomationTriggerSerial(o["serial"], d, "serial", sv)); err != nil {
		if !fortiAPIPatch(o["serial"]) {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("fabric_event_name", flattenSystemAutomationTriggerFabricEventName(o["fabric-event-name"], d, "fabric_event_name", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-event-name"]) {
			return fmt.Errorf("Error reading fabric_event_name: %v", err)
		}
	}

	if err = d.Set("fabric_event_severity", flattenSystemAutomationTriggerFabricEventSeverity(o["fabric-event-severity"], d, "fabric_event_severity", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-event-severity"]) {
			return fmt.Errorf("Error reading fabric_event_severity: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("logid_block", flattenSystemAutomationTriggerLogidBlock(o["logid"], d, "logid_block", sv)); err != nil {
			if !fortiAPIPatch(o["logid"]) {
				return fmt.Errorf("Error reading logid_block: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("logid_block"); ok {
			if err = d.Set("logid_block", flattenSystemAutomationTriggerLogidBlock(o["logid"], d, "logid_block", sv)); err != nil {
				if !fortiAPIPatch(o["logid"]) {
					return fmt.Errorf("Error reading logid_block: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemAutomationTriggerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutomationTriggerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerEventType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemAutomationTriggerVdomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationTriggerVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLicenseType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerIocLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerReportType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLogid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerFrequency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerWeekday(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerHour(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerTriggerMinute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFields(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemAutomationTriggerFieldsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemAutomationTriggerFieldsName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandSystemAutomationTriggerFieldsValue(d, i["value"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationTriggerFieldsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFieldsValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFazEventName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFazEventSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFazEventTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFabricEventName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerFabricEventSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationTriggerLogidBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemAutomationTriggerLogidBlockId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationTriggerLogidBlockId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationTrigger(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemAutomationTriggerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSystemAutomationTriggerDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("trigger_type"); ok {

		t, err := expandSystemAutomationTriggerTriggerType(d, v, "trigger_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-type"] = t
		}
	}

	if v, ok := d.GetOk("event_type"); ok {

		t, err := expandSystemAutomationTriggerEventType(d, v, "event_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event-type"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {

		t, err := expandSystemAutomationTriggerVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("license_type"); ok {

		t, err := expandSystemAutomationTriggerLicenseType(d, v, "license_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license-type"] = t
		}
	}

	if v, ok := d.GetOk("ioc_level"); ok {

		t, err := expandSystemAutomationTriggerIocLevel(d, v, "ioc_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ioc-level"] = t
		}
	}

	if v, ok := d.GetOk("report_type"); ok {

		t, err := expandSystemAutomationTriggerReportType(d, v, "report_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-type"] = t
		}
	}

	if v, ok := d.GetOk("logid"); ok {

		t, err := expandSystemAutomationTriggerLogid(d, v, "logid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logid"] = t
		}
	}

	if v, ok := d.GetOk("trigger_frequency"); ok {

		t, err := expandSystemAutomationTriggerTriggerFrequency(d, v, "trigger_frequency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-frequency"] = t
		}
	}

	if v, ok := d.GetOk("trigger_weekday"); ok {

		t, err := expandSystemAutomationTriggerTriggerWeekday(d, v, "trigger_weekday", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-weekday"] = t
		}
	}

	if v, ok := d.GetOk("trigger_day"); ok {

		t, err := expandSystemAutomationTriggerTriggerDay(d, v, "trigger_day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-day"] = t
		}
	}

	if v, ok := d.GetOkExists("trigger_hour"); ok {

		t, err := expandSystemAutomationTriggerTriggerHour(d, v, "trigger_hour", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-hour"] = t
		}
	}

	if v, ok := d.GetOkExists("trigger_minute"); ok {

		t, err := expandSystemAutomationTriggerTriggerMinute(d, v, "trigger_minute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger-minute"] = t
		}
	}

	if v, ok := d.GetOk("fields"); ok {

		t, err := expandSystemAutomationTriggerFields(d, v, "fields", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fields"] = t
		}
	}

	if v, ok := d.GetOk("faz_event_name"); ok {

		t, err := expandSystemAutomationTriggerFazEventName(d, v, "faz_event_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-event-name"] = t
		}
	}

	if v, ok := d.GetOk("faz_event_severity"); ok {

		t, err := expandSystemAutomationTriggerFazEventSeverity(d, v, "faz_event_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-event-severity"] = t
		}
	}

	if v, ok := d.GetOk("faz_event_tags"); ok {

		t, err := expandSystemAutomationTriggerFazEventTags(d, v, "faz_event_tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-event-tags"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok {

		t, err := expandSystemAutomationTriggerSerial(d, v, "serial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("fabric_event_name"); ok {

		t, err := expandSystemAutomationTriggerFabricEventName(d, v, "fabric_event_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-event-name"] = t
		}
	}

	if v, ok := d.GetOk("fabric_event_severity"); ok {

		t, err := expandSystemAutomationTriggerFabricEventSeverity(d, v, "fabric_event_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-event-severity"] = t
		}
	}

	if v, ok := d.GetOk("logid_block"); ok {

		t, err := expandSystemAutomationTriggerLogidBlock(d, v, "logid_block", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logid"] = t
		}
	}

	return &obj, nil
}
