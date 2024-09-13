// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Show FMWP signatures.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRuleFmwp() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleFmwpCreate,
		Read:   resourceRuleFmwpRead,
		Update: resourceRuleFmwpUpdate,
		Delete: resourceRuleFmwpDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
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
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
						},
						"metaid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"valueid": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
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

func resourceRuleFmwpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRuleFmwp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RuleFmwp resource while getting object: %v", err)
	}

	o, err := c.CreateRuleFmwp(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RuleFmwp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RuleFmwp")
	}

	return resourceRuleFmwpRead(d, m)
}

func resourceRuleFmwpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRuleFmwp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RuleFmwp resource while getting object: %v", err)
	}

	o, err := c.UpdateRuleFmwp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RuleFmwp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RuleFmwp")
	}

	return resourceRuleFmwpRead(d, m)
}

func resourceRuleFmwpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRuleFmwp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RuleFmwp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRuleFmwpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRuleFmwp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RuleFmwp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRuleFmwp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RuleFmwp resource from API: %v", err)
	}
	return nil
}

func flattenRuleFmwpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRuleFmwpRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRuleFmwpRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRuleFmwpDate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRuleFmwpMetadata(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRuleFmwpMetadataId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metaid"
		if cur_v, ok := i["metaid"]; ok {
			tmp["metaid"] = flattenRuleFmwpMetadataMetaid(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valueid"
		if cur_v, ok := i["valueid"]; ok {
			tmp["valueid"] = flattenRuleFmwpMetadataValueid(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRuleFmwpMetadataId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRuleFmwpMetadataMetaid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRuleFmwpMetadataValueid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectRuleFmwp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenRuleFmwpName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("log", flattenRuleFmwpLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("log_packet", flattenRuleFmwpLogPacket(o["log-packet"], d, "log_packet", sv)); err != nil {
		if !fortiAPIPatch(o["log-packet"]) {
			return fmt.Errorf("Error reading log_packet: %v", err)
		}
	}

	if err = d.Set("action", flattenRuleFmwpAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("group", flattenRuleFmwpGroup(o["group"], d, "group", sv)); err != nil {
		if !fortiAPIPatch(o["group"]) {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}

	if err = d.Set("severity", flattenRuleFmwpSeverity(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("location", flattenRuleFmwpLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("os", flattenRuleFmwpOs(o["os"], d, "os", sv)); err != nil {
		if !fortiAPIPatch(o["os"]) {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("application", flattenRuleFmwpApplication(o["application"], d, "application", sv)); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("service", flattenRuleFmwpService(o["service"], d, "service", sv)); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("rule_id", flattenRuleFmwpRuleId(o["rule-id"], d, "rule_id", sv)); err != nil {
		if !fortiAPIPatch(o["rule-id"]) {
			return fmt.Errorf("Error reading rule_id: %v", err)
		}
	}

	if err = d.Set("rev", flattenRuleFmwpRev(o["rev"], d, "rev", sv)); err != nil {
		if !fortiAPIPatch(o["rev"]) {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("date", flattenRuleFmwpDate(o["date"], d, "date", sv)); err != nil {
		if !fortiAPIPatch(o["date"]) {
			return fmt.Errorf("Error reading date: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("metadata", flattenRuleFmwpMetadata(o["metadata"], d, "metadata", sv)); err != nil {
			if !fortiAPIPatch(o["metadata"]) {
				return fmt.Errorf("Error reading metadata: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("metadata"); ok {
			if err = d.Set("metadata", flattenRuleFmwpMetadata(o["metadata"], d, "metadata", sv)); err != nil {
				if !fortiAPIPatch(o["metadata"]) {
					return fmt.Errorf("Error reading metadata: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRuleFmwpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRuleFmwpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpDate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpMetadata(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRuleFmwpMetadataId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metaid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metaid"], _ = expandRuleFmwpMetadataMetaid(d, i["metaid"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["metaid"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valueid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["valueid"], _ = expandRuleFmwpMetadataValueid(d, i["valueid"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["valueid"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRuleFmwpMetadataId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpMetadataMetaid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRuleFmwpMetadataValueid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRuleFmwp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRuleFmwpName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {
		t, err := expandRuleFmwpLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("log_packet"); ok {
		t, err := expandRuleFmwpLogPacket(d, v, "log_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-packet"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandRuleFmwpAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("group"); ok {
		t, err := expandRuleFmwpGroup(d, v, "group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	} else if d.HasChange("group") {
		obj["group"] = nil
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandRuleFmwpSeverity(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	} else if d.HasChange("severity") {
		obj["severity"] = nil
	}

	if v, ok := d.GetOk("location"); ok {
		t, err := expandRuleFmwpLocation(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	} else if d.HasChange("location") {
		obj["location"] = nil
	}

	if v, ok := d.GetOk("os"); ok {
		t, err := expandRuleFmwpOs(d, v, "os", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	} else if d.HasChange("os") {
		obj["os"] = nil
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandRuleFmwpApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	} else if d.HasChange("application") {
		obj["application"] = nil
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandRuleFmwpService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	} else if d.HasChange("service") {
		obj["service"] = nil
	}

	if v, ok := d.GetOkExists("rule_id"); ok {
		t, err := expandRuleFmwpRuleId(d, v, "rule_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule-id"] = t
		}
	}

	if v, ok := d.GetOkExists("rev"); ok {
		t, err := expandRuleFmwpRev(d, v, "rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOkExists("date"); ok {
		t, err := expandRuleFmwpDate(d, v, "date", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["date"] = t
		}
	}

	if v, ok := d.GetOk("metadata"); ok || d.HasChange("metadata") {
		t, err := expandRuleFmwpMetadata(d, v, "metadata", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["metadata"] = t
		}
	}

	return &obj, nil
}
