// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure general log settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSettingUpdate,
		Read:   resourceLogSettingRead,
		Update: resourceLogSettingUpdate,
		Delete: resourceLogSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resolve_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_user_in_upper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwpolicy_implicit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fwpolicy6_implicit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_invalid_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_allow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_deny_unicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_in_deny_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"daemon_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor_event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"brief_traffic_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_anonymize": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expolicy_implicit_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_policy_comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_policy_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syslog_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_api_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_api_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"anonymization_hash": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLogSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogSetting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogSetting")
	}

	return resourceLogSettingRead(d, m)
}

func resourceLogSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLogSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LogSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogSettingResolveIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingResolvePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLogUserInUpper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingFwpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingFwpolicy6ImplicitLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLogInvalidPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLocalInAllow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLocalInDenyUnicast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLocalInDenyBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLocalOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingDaemonLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingNeighborEvent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingBriefTrafficFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingUserAnonymize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingExpolicyImplicitLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLogPolicyComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingLogPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingFazOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingSyslogOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingRestApiSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingRestApiGet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingCustomLogFields(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_id"
		if _, ok := i["field-id"]; ok {

			tmp["field_id"] = flattenLogSettingCustomLogFieldsFieldId(i["field-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "field_id", d)
	return result
}

func flattenLogSettingCustomLogFieldsFieldId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogSettingAnonymizationHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("resolve_ip", flattenLogSettingResolveIp(o["resolve-ip"], d, "resolve_ip", sv)); err != nil {
		if !fortiAPIPatch(o["resolve-ip"]) {
			return fmt.Errorf("Error reading resolve_ip: %v", err)
		}
	}

	if err = d.Set("resolve_port", flattenLogSettingResolvePort(o["resolve-port"], d, "resolve_port", sv)); err != nil {
		if !fortiAPIPatch(o["resolve-port"]) {
			return fmt.Errorf("Error reading resolve_port: %v", err)
		}
	}

	if err = d.Set("log_user_in_upper", flattenLogSettingLogUserInUpper(o["log-user-in-upper"], d, "log_user_in_upper", sv)); err != nil {
		if !fortiAPIPatch(o["log-user-in-upper"]) {
			return fmt.Errorf("Error reading log_user_in_upper: %v", err)
		}
	}

	if err = d.Set("fwpolicy_implicit_log", flattenLogSettingFwpolicyImplicitLog(o["fwpolicy-implicit-log"], d, "fwpolicy_implicit_log", sv)); err != nil {
		if !fortiAPIPatch(o["fwpolicy-implicit-log"]) {
			return fmt.Errorf("Error reading fwpolicy_implicit_log: %v", err)
		}
	}

	if err = d.Set("fwpolicy6_implicit_log", flattenLogSettingFwpolicy6ImplicitLog(o["fwpolicy6-implicit-log"], d, "fwpolicy6_implicit_log", sv)); err != nil {
		if !fortiAPIPatch(o["fwpolicy6-implicit-log"]) {
			return fmt.Errorf("Error reading fwpolicy6_implicit_log: %v", err)
		}
	}

	if err = d.Set("log_invalid_packet", flattenLogSettingLogInvalidPacket(o["log-invalid-packet"], d, "log_invalid_packet", sv)); err != nil {
		if !fortiAPIPatch(o["log-invalid-packet"]) {
			return fmt.Errorf("Error reading log_invalid_packet: %v", err)
		}
	}

	if err = d.Set("local_in_allow", flattenLogSettingLocalInAllow(o["local-in-allow"], d, "local_in_allow", sv)); err != nil {
		if !fortiAPIPatch(o["local-in-allow"]) {
			return fmt.Errorf("Error reading local_in_allow: %v", err)
		}
	}

	if err = d.Set("local_in_deny_unicast", flattenLogSettingLocalInDenyUnicast(o["local-in-deny-unicast"], d, "local_in_deny_unicast", sv)); err != nil {
		if !fortiAPIPatch(o["local-in-deny-unicast"]) {
			return fmt.Errorf("Error reading local_in_deny_unicast: %v", err)
		}
	}

	if err = d.Set("local_in_deny_broadcast", flattenLogSettingLocalInDenyBroadcast(o["local-in-deny-broadcast"], d, "local_in_deny_broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["local-in-deny-broadcast"]) {
			return fmt.Errorf("Error reading local_in_deny_broadcast: %v", err)
		}
	}

	if err = d.Set("local_out", flattenLogSettingLocalOut(o["local-out"], d, "local_out", sv)); err != nil {
		if !fortiAPIPatch(o["local-out"]) {
			return fmt.Errorf("Error reading local_out: %v", err)
		}
	}

	if err = d.Set("daemon_log", flattenLogSettingDaemonLog(o["daemon-log"], d, "daemon_log", sv)); err != nil {
		if !fortiAPIPatch(o["daemon-log"]) {
			return fmt.Errorf("Error reading daemon_log: %v", err)
		}
	}

	if err = d.Set("neighbor_event", flattenLogSettingNeighborEvent(o["neighbor-event"], d, "neighbor_event", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-event"]) {
			return fmt.Errorf("Error reading neighbor_event: %v", err)
		}
	}

	if err = d.Set("brief_traffic_format", flattenLogSettingBriefTrafficFormat(o["brief-traffic-format"], d, "brief_traffic_format", sv)); err != nil {
		if !fortiAPIPatch(o["brief-traffic-format"]) {
			return fmt.Errorf("Error reading brief_traffic_format: %v", err)
		}
	}

	if err = d.Set("user_anonymize", flattenLogSettingUserAnonymize(o["user-anonymize"], d, "user_anonymize", sv)); err != nil {
		if !fortiAPIPatch(o["user-anonymize"]) {
			return fmt.Errorf("Error reading user_anonymize: %v", err)
		}
	}

	if err = d.Set("expolicy_implicit_log", flattenLogSettingExpolicyImplicitLog(o["expolicy-implicit-log"], d, "expolicy_implicit_log", sv)); err != nil {
		if !fortiAPIPatch(o["expolicy-implicit-log"]) {
			return fmt.Errorf("Error reading expolicy_implicit_log: %v", err)
		}
	}

	if err = d.Set("log_policy_comment", flattenLogSettingLogPolicyComment(o["log-policy-comment"], d, "log_policy_comment", sv)); err != nil {
		if !fortiAPIPatch(o["log-policy-comment"]) {
			return fmt.Errorf("Error reading log_policy_comment: %v", err)
		}
	}

	if err = d.Set("log_policy_name", flattenLogSettingLogPolicyName(o["log-policy-name"], d, "log_policy_name", sv)); err != nil {
		if !fortiAPIPatch(o["log-policy-name"]) {
			return fmt.Errorf("Error reading log_policy_name: %v", err)
		}
	}

	if err = d.Set("faz_override", flattenLogSettingFazOverride(o["faz-override"], d, "faz_override", sv)); err != nil {
		if !fortiAPIPatch(o["faz-override"]) {
			return fmt.Errorf("Error reading faz_override: %v", err)
		}
	}

	if err = d.Set("syslog_override", flattenLogSettingSyslogOverride(o["syslog-override"], d, "syslog_override", sv)); err != nil {
		if !fortiAPIPatch(o["syslog-override"]) {
			return fmt.Errorf("Error reading syslog_override: %v", err)
		}
	}

	if err = d.Set("rest_api_set", flattenLogSettingRestApiSet(o["rest-api-set"], d, "rest_api_set", sv)); err != nil {
		if !fortiAPIPatch(o["rest-api-set"]) {
			return fmt.Errorf("Error reading rest_api_set: %v", err)
		}
	}

	if err = d.Set("rest_api_get", flattenLogSettingRestApiGet(o["rest-api-get"], d, "rest_api_get", sv)); err != nil {
		if !fortiAPIPatch(o["rest-api-get"]) {
			return fmt.Errorf("Error reading rest_api_get: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_log_fields", flattenLogSettingCustomLogFields(o["custom-log-fields"], d, "custom_log_fields", sv)); err != nil {
			if !fortiAPIPatch(o["custom-log-fields"]) {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_log_fields"); ok {
			if err = d.Set("custom_log_fields", flattenLogSettingCustomLogFields(o["custom-log-fields"], d, "custom_log_fields", sv)); err != nil {
				if !fortiAPIPatch(o["custom-log-fields"]) {
					return fmt.Errorf("Error reading custom_log_fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("anonymization_hash", flattenLogSettingAnonymizationHash(o["anonymization-hash"], d, "anonymization_hash", sv)); err != nil {
		if !fortiAPIPatch(o["anonymization-hash"]) {
			return fmt.Errorf("Error reading anonymization_hash: %v", err)
		}
	}

	return nil
}

func flattenLogSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogSettingResolveIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingResolvePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogUserInUpper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFwpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFwpolicy6ImplicitLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogInvalidPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInAllow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInDenyUnicast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalInDenyBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLocalOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingDaemonLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingNeighborEvent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingBriefTrafficFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingUserAnonymize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingExpolicyImplicitLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogPolicyComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingLogPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingFazOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingSyslogOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingRestApiSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingRestApiGet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingCustomLogFields(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["field-id"], _ = expandLogSettingCustomLogFieldsFieldId(d, i["field_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogSettingCustomLogFieldsFieldId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogSettingAnonymizationHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogSetting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("resolve_ip"); ok {

		t, err := expandLogSettingResolveIp(d, v, "resolve_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-ip"] = t
		}
	}

	if v, ok := d.GetOk("resolve_port"); ok {

		t, err := expandLogSettingResolvePort(d, v, "resolve_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-port"] = t
		}
	}

	if v, ok := d.GetOk("log_user_in_upper"); ok {

		t, err := expandLogSettingLogUserInUpper(d, v, "log_user_in_upper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-user-in-upper"] = t
		}
	}

	if v, ok := d.GetOk("fwpolicy_implicit_log"); ok {

		t, err := expandLogSettingFwpolicyImplicitLog(d, v, "fwpolicy_implicit_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwpolicy-implicit-log"] = t
		}
	}

	if v, ok := d.GetOk("fwpolicy6_implicit_log"); ok {

		t, err := expandLogSettingFwpolicy6ImplicitLog(d, v, "fwpolicy6_implicit_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwpolicy6-implicit-log"] = t
		}
	}

	if v, ok := d.GetOk("log_invalid_packet"); ok {

		t, err := expandLogSettingLogInvalidPacket(d, v, "log_invalid_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-invalid-packet"] = t
		}
	}

	if v, ok := d.GetOk("local_in_allow"); ok {

		t, err := expandLogSettingLocalInAllow(d, v, "local_in_allow", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-allow"] = t
		}
	}

	if v, ok := d.GetOk("local_in_deny_unicast"); ok {

		t, err := expandLogSettingLocalInDenyUnicast(d, v, "local_in_deny_unicast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-deny-unicast"] = t
		}
	}

	if v, ok := d.GetOk("local_in_deny_broadcast"); ok {

		t, err := expandLogSettingLocalInDenyBroadcast(d, v, "local_in_deny_broadcast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-in-deny-broadcast"] = t
		}
	}

	if v, ok := d.GetOk("local_out"); ok {

		t, err := expandLogSettingLocalOut(d, v, "local_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-out"] = t
		}
	}

	if v, ok := d.GetOk("daemon_log"); ok {

		t, err := expandLogSettingDaemonLog(d, v, "daemon_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daemon-log"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_event"); ok {

		t, err := expandLogSettingNeighborEvent(d, v, "neighbor_event", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-event"] = t
		}
	}

	if v, ok := d.GetOk("brief_traffic_format"); ok {

		t, err := expandLogSettingBriefTrafficFormat(d, v, "brief_traffic_format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["brief-traffic-format"] = t
		}
	}

	if v, ok := d.GetOk("user_anonymize"); ok {

		t, err := expandLogSettingUserAnonymize(d, v, "user_anonymize", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-anonymize"] = t
		}
	}

	if v, ok := d.GetOk("expolicy_implicit_log"); ok {

		t, err := expandLogSettingExpolicyImplicitLog(d, v, "expolicy_implicit_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expolicy-implicit-log"] = t
		}
	}

	if v, ok := d.GetOk("log_policy_comment"); ok {

		t, err := expandLogSettingLogPolicyComment(d, v, "log_policy_comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-policy-comment"] = t
		}
	}

	if v, ok := d.GetOk("log_policy_name"); ok {

		t, err := expandLogSettingLogPolicyName(d, v, "log_policy_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-policy-name"] = t
		}
	}

	if v, ok := d.GetOk("faz_override"); ok {

		t, err := expandLogSettingFazOverride(d, v, "faz_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-override"] = t
		}
	}

	if v, ok := d.GetOk("syslog_override"); ok {

		t, err := expandLogSettingSyslogOverride(d, v, "syslog_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-override"] = t
		}
	}

	if v, ok := d.GetOk("rest_api_set"); ok {

		t, err := expandLogSettingRestApiSet(d, v, "rest_api_set", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api-set"] = t
		}
	}

	if v, ok := d.GetOk("rest_api_get"); ok {

		t, err := expandLogSettingRestApiGet(d, v, "rest_api_get", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api-get"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok {

		t, err := expandLogSettingCustomLogFields(d, v, "custom_log_fields", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("anonymization_hash"); ok {

		t, err := expandLogSettingAnonymizationHash(d, v, "anonymization_hash", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anonymization-hash"] = t
		}
	}

	return &obj, nil
}
