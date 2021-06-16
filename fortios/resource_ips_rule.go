// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS rules.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsRuleCreate,
		Read:   resourceIpsRuleRead,
		Update: resourceIpsRuleUpdate,
		Delete: resourceIpsRuleDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"date": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"metaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"valueid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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

func resourceIpsRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating IpsRule resource while getting object: %v", err)
	}

	o, err := c.CreateIpsRule(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating IpsRule resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsRule")
	}

	return resourceIpsRuleRead(d, m)
}

func resourceIpsRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIpsRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IpsRule resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsRule(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IpsRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsRule")
	}

	return resourceIpsRuleRead(d, m)
}

func resourceIpsRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteIpsRule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting IpsRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadIpsRule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IpsRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IpsRule resource from API: %v", err)
	}
	return nil
}

func flattenIpsRuleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleDate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleMetadata(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenIpsRuleMetadataId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metaid"
		if _, ok := i["metaid"]; ok {

			tmp["metaid"] = flattenIpsRuleMetadataMetaid(i["metaid"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valueid"
		if _, ok := i["valueid"]; ok {

			tmp["valueid"] = flattenIpsRuleMetadataValueid(i["valueid"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenIpsRuleMetadataId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleMetadataMetaid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsRuleMetadataValueid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIpsRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenIpsRuleName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenIpsRuleStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("log", flattenIpsRuleLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenIpsRuleLogPacket(o["log-packet"], d, "log_packet", sv)); err != nil {
		if !fortiAPIPatch(o["log-packet"]) {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("action", flattenIpsRuleAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("group", flattenIpsRuleGroup(o["group"], d, "group", sv)); err != nil {
		if !fortiAPIPatch(o["group"]) {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("severity", flattenIpsRuleSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("location", flattenIpsRuleLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("os", flattenIpsRuleOs(o["os"], d, "os", sv)); err != nil {
		if !fortiAPIPatch(o["os"]) {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("application", flattenIpsRuleApplication(o["application"], d, "application", sv)); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("service", flattenIpsRuleService(o["service"], d, "service", sv)); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("rule_id", flattenIpsRuleRuleId(o["rule-id"], d, "rule_id", sv)); err != nil {
		if !fortiAPIPatch(o["rule-id"]) {
			return fmt.Errorf("Error reading rule_id: %v", err)
		}
	}

	if err = d.Set("rev", flattenIpsRuleRev(o["rev"], d, "rev", sv)); err != nil {
		if !fortiAPIPatch(o["rev"]) {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("date", flattenIpsRuleDate(o["date"], d, "date", sv)); err != nil {
		if !fortiAPIPatch(o["date"]) {
			return fmt.Errorf("Error reading date: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("metadata", flattenIpsRuleMetadata(o["metadata"], d, "metadata", sv)); err != nil {
			if !fortiAPIPatch(o["metadata"]) {
				return fmt.Errorf("Error reading metadata: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("metadata"); ok {
			if err = d.Set("metadata", flattenIpsRuleMetadata(o["metadata"], d, "metadata", sv)); err != nil {
				if !fortiAPIPatch(o["metadata"]) {
					return fmt.Errorf("Error reading metadata: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIpsRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIpsRuleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleDate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleMetadata(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIpsRuleMetadataId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metaid"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metaid"], _ = expandIpsRuleMetadataMetaid(d, i["metaid"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valueid"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["valueid"], _ = expandIpsRuleMetadataValueid(d, i["valueid"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsRuleMetadataId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleMetadataMetaid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsRuleMetadataValueid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIpsRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandIpsRuleName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandIpsRuleStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {

		t, err := expandIpsRuleLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok {

		t, err := expandIpsRuleLogPacket(d, v, "log_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandIpsRuleAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok {

		t, err := expandIpsRuleGroup(d, v, "group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {

		t, err := expandIpsRuleSeverity(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {

		t, err := expandIpsRuleLocation(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok {

		t, err := expandIpsRuleOs(d, v, "os", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {

		t, err := expandIpsRuleApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandIpsRuleService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOkExists("rule_id"); ok {

		t, err := expandIpsRuleRuleId(d, v, "rule_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-id"] = t
		}
	}

	if v, ok := d.GetOkExists("rev"); ok {

		t, err := expandIpsRuleRev(d, v, "rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOkExists("date"); ok {

		t, err := expandIpsRuleDate(d, v, "date", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["date"] = t
		}
	}

	if v, ok := d.GetOk("metadata"); ok {

		t, err := expandIpsRuleMetadata(d, v, "metadata", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metadata"] = t
		}
	}

	return &obj, nil
}
