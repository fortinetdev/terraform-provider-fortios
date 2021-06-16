// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure user authentication setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserSettingUpdate,
		Read:   resourceUserSettingRead,
		Update: resourceUserSettingUpdate,
		Delete: resourceUserSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_secure_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_http_basic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_ssl_allow_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_on_demand": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"auth_timeout_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_portal_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"radius_ses_timeout_act": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_blackout_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"auth_invalid_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"auth_lockout_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"auth_lockout_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"per_policy_disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"auth_ssl_min_proto_version": &schema.Schema{
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

func resourceUserSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateUserSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserSetting")
	}

	return resourceUserSettingRead(d, m)
}

func resourceUserSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating UserSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateUserSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing UserSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserSetting resource from API: %v", err)
	}
	return nil
}

func flattenUserSettingAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSecureHttp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthHttpBasic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSslAllowRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthOnDemand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthTimeoutType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortalTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingRadiusSesTimeoutAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthBlackoutTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthInvalidMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthLockoutThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthLockoutDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingPerPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenUserSettingAuthPortsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenUserSettingAuthPortsType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenUserSettingAuthPortsPort(i["port"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserSettingAuthPortsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auth_type", flattenUserSettingAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenUserSettingAuthCert(o["auth-cert"], d, "auth_cert", sv)); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_ca_cert", flattenUserSettingAuthCaCert(o["auth-ca-cert"], d, "auth_ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["auth-ca-cert"]) {
			return fmt.Errorf("Error reading auth_ca_cert: %v", err)
		}
	}

	if err = d.Set("auth_secure_http", flattenUserSettingAuthSecureHttp(o["auth-secure-http"], d, "auth_secure_http", sv)); err != nil {
		if !fortiAPIPatch(o["auth-secure-http"]) {
			return fmt.Errorf("Error reading auth_secure_http: %v", err)
		}
	}

	if err = d.Set("auth_http_basic", flattenUserSettingAuthHttpBasic(o["auth-http-basic"], d, "auth_http_basic", sv)); err != nil {
		if !fortiAPIPatch(o["auth-http-basic"]) {
			return fmt.Errorf("Error reading auth_http_basic: %v", err)
		}
	}

	if err = d.Set("auth_ssl_allow_renegotiation", flattenUserSettingAuthSslAllowRenegotiation(o["auth-ssl-allow-renegotiation"], d, "auth_ssl_allow_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["auth-ssl-allow-renegotiation"]) {
			return fmt.Errorf("Error reading auth_ssl_allow_renegotiation: %v", err)
		}
	}

	if err = d.Set("auth_src_mac", flattenUserSettingAuthSrcMac(o["auth-src-mac"], d, "auth_src_mac", sv)); err != nil {
		if !fortiAPIPatch(o["auth-src-mac"]) {
			return fmt.Errorf("Error reading auth_src_mac: %v", err)
		}
	}

	if err = d.Set("auth_on_demand", flattenUserSettingAuthOnDemand(o["auth-on-demand"], d, "auth_on_demand", sv)); err != nil {
		if !fortiAPIPatch(o["auth-on-demand"]) {
			return fmt.Errorf("Error reading auth_on_demand: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenUserSettingAuthTimeout(o["auth-timeout"], d, "auth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auth-timeout"]) {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if err = d.Set("auth_timeout_type", flattenUserSettingAuthTimeoutType(o["auth-timeout-type"], d, "auth_timeout_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-timeout-type"]) {
			return fmt.Errorf("Error reading auth_timeout_type: %v", err)
		}
	}

	if err = d.Set("auth_portal_timeout", flattenUserSettingAuthPortalTimeout(o["auth-portal-timeout"], d, "auth_portal_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal-timeout"]) {
			return fmt.Errorf("Error reading auth_portal_timeout: %v", err)
		}
	}

	if err = d.Set("radius_ses_timeout_act", flattenUserSettingRadiusSesTimeoutAct(o["radius-ses-timeout-act"], d, "radius_ses_timeout_act", sv)); err != nil {
		if !fortiAPIPatch(o["radius-ses-timeout-act"]) {
			return fmt.Errorf("Error reading radius_ses_timeout_act: %v", err)
		}
	}

	if err = d.Set("auth_blackout_time", flattenUserSettingAuthBlackoutTime(o["auth-blackout-time"], d, "auth_blackout_time", sv)); err != nil {
		if !fortiAPIPatch(o["auth-blackout-time"]) {
			return fmt.Errorf("Error reading auth_blackout_time: %v", err)
		}
	}

	if err = d.Set("auth_invalid_max", flattenUserSettingAuthInvalidMax(o["auth-invalid-max"], d, "auth_invalid_max", sv)); err != nil {
		if !fortiAPIPatch(o["auth-invalid-max"]) {
			return fmt.Errorf("Error reading auth_invalid_max: %v", err)
		}
	}

	if err = d.Set("auth_lockout_threshold", flattenUserSettingAuthLockoutThreshold(o["auth-lockout-threshold"], d, "auth_lockout_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["auth-lockout-threshold"]) {
			return fmt.Errorf("Error reading auth_lockout_threshold: %v", err)
		}
	}

	if err = d.Set("auth_lockout_duration", flattenUserSettingAuthLockoutDuration(o["auth-lockout-duration"], d, "auth_lockout_duration", sv)); err != nil {
		if !fortiAPIPatch(o["auth-lockout-duration"]) {
			return fmt.Errorf("Error reading auth_lockout_duration: %v", err)
		}
	}

	if err = d.Set("per_policy_disclaimer", flattenUserSettingPerPolicyDisclaimer(o["per-policy-disclaimer"], d, "per_policy_disclaimer", sv)); err != nil {
		if !fortiAPIPatch(o["per-policy-disclaimer"]) {
			return fmt.Errorf("Error reading per_policy_disclaimer: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o["auth-ports"], d, "auth_ports", sv)); err != nil {
			if !fortiAPIPatch(o["auth-ports"]) {
				return fmt.Errorf("Error reading auth_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth_ports"); ok {
			if err = d.Set("auth_ports", flattenUserSettingAuthPorts(o["auth-ports"], d, "auth_ports", sv)); err != nil {
				if !fortiAPIPatch(o["auth-ports"]) {
					return fmt.Errorf("Error reading auth_ports: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_ssl_min_proto_version", flattenUserSettingAuthSslMinProtoVersion(o["auth-ssl-min-proto-version"], d, "auth_ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["auth-ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading auth_ssl_min_proto_version: %v", err)
		}
	}

	return nil
}

func flattenUserSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserSettingAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSecureHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthHttpBasic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslAllowRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthOnDemand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthTimeoutType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortalTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingRadiusSesTimeoutAct(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthBlackoutTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthInvalidMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthLockoutThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthLockoutDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingPerPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandUserSettingAuthPortsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandUserSettingAuthPortsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandUserSettingAuthPortsPort(d, i["port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserSettingAuthPortsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthPortsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserSetting(d *schema.ResourceData, bemptysontable bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_type"); ok {

		t, err := expandUserSettingAuthType(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok {

		t, err := expandUserSettingAuthCert(d, v, "auth_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_ca_cert"); ok {

		t, err := expandUserSettingAuthCaCert(d, v, "auth_ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_secure_http"); ok {

		t, err := expandUserSettingAuthSecureHttp(d, v, "auth_secure_http", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-secure-http"] = t
		}
	}

	if v, ok := d.GetOk("auth_http_basic"); ok {

		t, err := expandUserSettingAuthHttpBasic(d, v, "auth_http_basic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-http-basic"] = t
		}
	}

	if v, ok := d.GetOk("auth_ssl_allow_renegotiation"); ok {

		t, err := expandUserSettingAuthSslAllowRenegotiation(d, v, "auth_ssl_allow_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ssl-allow-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("auth_src_mac"); ok {

		t, err := expandUserSettingAuthSrcMac(d, v, "auth_src_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-src-mac"] = t
		}
	}

	if v, ok := d.GetOk("auth_on_demand"); ok {

		t, err := expandUserSettingAuthOnDemand(d, v, "auth_on_demand", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-on-demand"] = t
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok {

		t, err := expandUserSettingAuthTimeout(d, v, "auth_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("auth_timeout_type"); ok {

		t, err := expandUserSettingAuthTimeoutType(d, v, "auth_timeout_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-timeout-type"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal_timeout"); ok {

		t, err := expandUserSettingAuthPortalTimeout(d, v, "auth_portal_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-timeout"] = t
		}
	}

	if v, ok := d.GetOk("radius_ses_timeout_act"); ok {

		t, err := expandUserSettingRadiusSesTimeoutAct(d, v, "radius_ses_timeout_act", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-ses-timeout-act"] = t
		}
	}

	if v, ok := d.GetOkExists("auth_blackout_time"); ok {

		t, err := expandUserSettingAuthBlackoutTime(d, v, "auth_blackout_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-blackout-time"] = t
		}
	}

	if v, ok := d.GetOk("auth_invalid_max"); ok {

		t, err := expandUserSettingAuthInvalidMax(d, v, "auth_invalid_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-invalid-max"] = t
		}
	}

	if v, ok := d.GetOk("auth_lockout_threshold"); ok {

		t, err := expandUserSettingAuthLockoutThreshold(d, v, "auth_lockout_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-lockout-threshold"] = t
		}
	}

	if v, ok := d.GetOkExists("auth_lockout_duration"); ok {

		t, err := expandUserSettingAuthLockoutDuration(d, v, "auth_lockout_duration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-lockout-duration"] = t
		}
	}

	if v, ok := d.GetOk("per_policy_disclaimer"); ok {

		t, err := expandUserSettingPerPolicyDisclaimer(d, v, "per_policy_disclaimer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-policy-disclaimer"] = t
		}
	}

	if bemptysontable {
		obj["auth-ports"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("auth_ports"); ok {

			t, err := expandUserSettingAuthPorts(d, v, "auth_ports", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ports"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_ssl_min_proto_version"); ok {

		t, err := expandUserSettingAuthSslMinProtoVersion(d, v, "auth_ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-ssl-min-proto-version"] = t
		}
	}

	return &obj, nil
}
