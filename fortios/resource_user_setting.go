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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Computed: true,
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
			},
			"auth_ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
			"auth_ssl_max_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_ssl_sigalgs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_user_password_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"cors": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cors_allowed_origins": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
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

func resourceUserSettingUpdate(d *schema.ResourceData, m interface{}) error {
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
	return convintf2i(v)
}

func flattenUserSettingAuthTimeoutType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortalTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingRadiusSesTimeoutAct(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthBlackoutTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingAuthInvalidMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingAuthLockoutThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingAuthLockoutDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingPerPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenUserSettingAuthPortsId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenUserSettingAuthPortsType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenUserSettingAuthPortsPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserSettingAuthPortsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingAuthPortsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthPortsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserSettingAuthSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSslMaxProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingAuthSslSigalgs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingDefaultUserPasswordPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingCors(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserSettingCorsAllowedOrigins(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenUserSettingCorsAllowedOriginsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenUserSettingCorsAllowedOriginsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

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

	if b_get_all_tables {
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

	if err = d.Set("auth_ssl_max_proto_version", flattenUserSettingAuthSslMaxProtoVersion(o["auth-ssl-max-proto-version"], d, "auth_ssl_max_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["auth-ssl-max-proto-version"]) {
			return fmt.Errorf("Error reading auth_ssl_max_proto_version: %v", err)
		}
	}

	if err = d.Set("auth_ssl_sigalgs", flattenUserSettingAuthSslSigalgs(o["auth-ssl-sigalgs"], d, "auth_ssl_sigalgs", sv)); err != nil {
		if !fortiAPIPatch(o["auth-ssl-sigalgs"]) {
			return fmt.Errorf("Error reading auth_ssl_sigalgs: %v", err)
		}
	}

	if err = d.Set("default_user_password_policy", flattenUserSettingDefaultUserPasswordPolicy(o["default-user-password-policy"], d, "default_user_password_policy", sv)); err != nil {
		if !fortiAPIPatch(o["default-user-password-policy"]) {
			return fmt.Errorf("Error reading default_user_password_policy: %v", err)
		}
	}

	if err = d.Set("cors", flattenUserSettingCors(o["cors"], d, "cors", sv)); err != nil {
		if !fortiAPIPatch(o["cors"]) {
			return fmt.Errorf("Error reading cors: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("cors_allowed_origins", flattenUserSettingCorsAllowedOrigins(o["cors-allowed-origins"], d, "cors_allowed_origins", sv)); err != nil {
			if !fortiAPIPatch(o["cors-allowed-origins"]) {
				return fmt.Errorf("Error reading cors_allowed_origins: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cors_allowed_origins"); ok {
			if err = d.Set("cors_allowed_origins", flattenUserSettingCorsAllowedOrigins(o["cors-allowed-origins"], d, "cors_allowed_origins", sv)); err != nil {
				if !fortiAPIPatch(o["cors-allowed-origins"]) {
					return fmt.Errorf("Error reading cors_allowed_origins: %v", err)
				}
			}
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
			tmp["id"], _ = expandUserSettingAuthPortsId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
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

func expandUserSettingAuthSslMaxProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingAuthSslSigalgs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingDefaultUserPasswordPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingCors(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserSettingCorsAllowedOrigins(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandUserSettingCorsAllowedOriginsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserSettingCorsAllowedOriginsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_type"); ok {
		if setArgNil {
			obj["auth-type"] = nil
		} else {
			t, err := expandUserSettingAuthType(d, v, "auth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok {
		if setArgNil {
			obj["auth-cert"] = nil
		} else {
			t, err := expandUserSettingAuthCert(d, v, "auth_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-cert"] = t
			}
		}
	} else if d.HasChange("auth_cert") {
		obj["auth-cert"] = nil
	}

	if v, ok := d.GetOk("auth_ca_cert"); ok {
		if setArgNil {
			obj["auth-ca-cert"] = nil
		} else {
			t, err := expandUserSettingAuthCaCert(d, v, "auth_ca_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ca-cert"] = t
			}
		}
	} else if d.HasChange("auth_ca_cert") {
		obj["auth-ca-cert"] = nil
	}

	if v, ok := d.GetOk("auth_secure_http"); ok {
		if setArgNil {
			obj["auth-secure-http"] = nil
		} else {
			t, err := expandUserSettingAuthSecureHttp(d, v, "auth_secure_http", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-secure-http"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_http_basic"); ok {
		if setArgNil {
			obj["auth-http-basic"] = nil
		} else {
			t, err := expandUserSettingAuthHttpBasic(d, v, "auth_http_basic", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-http-basic"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_ssl_allow_renegotiation"); ok {
		if setArgNil {
			obj["auth-ssl-allow-renegotiation"] = nil
		} else {
			t, err := expandUserSettingAuthSslAllowRenegotiation(d, v, "auth_ssl_allow_renegotiation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ssl-allow-renegotiation"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_src_mac"); ok {
		if setArgNil {
			obj["auth-src-mac"] = nil
		} else {
			t, err := expandUserSettingAuthSrcMac(d, v, "auth_src_mac", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-src-mac"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_on_demand"); ok {
		if setArgNil {
			obj["auth-on-demand"] = nil
		} else {
			t, err := expandUserSettingAuthOnDemand(d, v, "auth_on_demand", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-on-demand"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_timeout"); ok {
		if setArgNil {
			obj["auth-timeout"] = nil
		} else {
			t, err := expandUserSettingAuthTimeout(d, v, "auth_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_timeout_type"); ok {
		if setArgNil {
			obj["auth-timeout-type"] = nil
		} else {
			t, err := expandUserSettingAuthTimeoutType(d, v, "auth_timeout_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-timeout-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_portal_timeout"); ok {
		if setArgNil {
			obj["auth-portal-timeout"] = nil
		} else {
			t, err := expandUserSettingAuthPortalTimeout(d, v, "auth_portal_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-portal-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("radius_ses_timeout_act"); ok {
		if setArgNil {
			obj["radius-ses-timeout-act"] = nil
		} else {
			t, err := expandUserSettingRadiusSesTimeoutAct(d, v, "radius_ses_timeout_act", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["radius-ses-timeout-act"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auth_blackout_time"); ok {
		if setArgNil {
			obj["auth-blackout-time"] = nil
		} else {
			t, err := expandUserSettingAuthBlackoutTime(d, v, "auth_blackout_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-blackout-time"] = t
			}
		}
	} else if d.HasChange("auth_blackout_time") {
		obj["auth-blackout-time"] = nil
	}

	if v, ok := d.GetOk("auth_invalid_max"); ok {
		if setArgNil {
			obj["auth-invalid-max"] = nil
		} else {
			t, err := expandUserSettingAuthInvalidMax(d, v, "auth_invalid_max", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-invalid-max"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_lockout_threshold"); ok {
		if setArgNil {
			obj["auth-lockout-threshold"] = nil
		} else {
			t, err := expandUserSettingAuthLockoutThreshold(d, v, "auth_lockout_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-lockout-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auth_lockout_duration"); ok {
		if setArgNil {
			obj["auth-lockout-duration"] = nil
		} else {
			t, err := expandUserSettingAuthLockoutDuration(d, v, "auth_lockout_duration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-lockout-duration"] = t
			}
		}
	} else if d.HasChange("auth_lockout_duration") {
		obj["auth-lockout-duration"] = nil
	}

	if v, ok := d.GetOk("per_policy_disclaimer"); ok {
		if setArgNil {
			obj["per-policy-disclaimer"] = nil
		} else {
			t, err := expandUserSettingPerPolicyDisclaimer(d, v, "per_policy_disclaimer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["per-policy-disclaimer"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_ports"); ok || d.HasChange("auth_ports") {
		if setArgNil {
			obj["auth-ports"] = make([]struct{}, 0)
		} else {
			t, err := expandUserSettingAuthPorts(d, v, "auth_ports", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ports"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_ssl_min_proto_version"); ok {
		if setArgNil {
			obj["auth-ssl-min-proto-version"] = nil
		} else {
			t, err := expandUserSettingAuthSslMinProtoVersion(d, v, "auth_ssl_min_proto_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ssl-min-proto-version"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_ssl_max_proto_version"); ok {
		if setArgNil {
			obj["auth-ssl-max-proto-version"] = nil
		} else {
			t, err := expandUserSettingAuthSslMaxProtoVersion(d, v, "auth_ssl_max_proto_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ssl-max-proto-version"] = t
			}
		}
	} else if d.HasChange("auth_ssl_max_proto_version") {
		obj["auth-ssl-max-proto-version"] = nil
	}

	if v, ok := d.GetOk("auth_ssl_sigalgs"); ok {
		if setArgNil {
			obj["auth-ssl-sigalgs"] = nil
		} else {
			t, err := expandUserSettingAuthSslSigalgs(d, v, "auth_ssl_sigalgs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-ssl-sigalgs"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_user_password_policy"); ok {
		if setArgNil {
			obj["default-user-password-policy"] = nil
		} else {
			t, err := expandUserSettingDefaultUserPasswordPolicy(d, v, "default_user_password_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-user-password-policy"] = t
			}
		}
	} else if d.HasChange("default_user_password_policy") {
		obj["default-user-password-policy"] = nil
	}

	if v, ok := d.GetOk("cors"); ok {
		if setArgNil {
			obj["cors"] = nil
		} else {
			t, err := expandUserSettingCors(d, v, "cors", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cors"] = t
			}
		}
	}

	if v, ok := d.GetOk("cors_allowed_origins"); ok || d.HasChange("cors_allowed_origins") {
		if setArgNil {
			obj["cors-allowed-origins"] = make([]struct{}, 0)
		} else {
			t, err := expandUserSettingCorsAllowedOrigins(d, v, "cors_allowed_origins", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cors-allowed-origins"] = t
			}
		}
	}

	return &obj, nil
}
