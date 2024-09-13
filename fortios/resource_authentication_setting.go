// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure authentication setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAuthenticationSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationSettingUpdate,
		Read:   resourceAuthenticationSettingRead,
		Update: resourceAuthenticationSettingUpdate,
		Delete: resourceAuthenticationSettingDelete,

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
			"active_auth_scheme": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"sso_auth_scheme": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"update_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistent_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cookie_max_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 10080),
				Optional:     true,
				Computed:     true,
			},
			"cookie_refresh_div": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 4),
				Optional:     true,
				Computed:     true,
			},
			"captive_portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"captive_portal6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_captive_portal": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cert_captive_portal_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_captive_portal_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"captive_portal_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"auth_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ssl_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"user_cert_ca": &schema.Schema{
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
			"dev_range": &schema.Schema{
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

func resourceAuthenticationSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateAuthenticationSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AuthenticationSetting")
	}

	return resourceAuthenticationSettingRead(d, m)
}

func resourceAuthenticationSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectAuthenticationSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating AuthenticationSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateAuthenticationSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing AuthenticationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAuthenticationSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationSetting resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationSettingActiveAuthScheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingSsoAuthScheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingUpdateTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingPersistentCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingIpAuthCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCookieMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationSettingCookieRefreshDiv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationSettingCaptivePortalType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortal6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCertAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCertCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCertCaptivePortalIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCertCaptivePortalPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationSettingCaptivePortalPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationSettingAuthHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalSslPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationSettingUserCertCa(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationSettingUserCertCaName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationSettingUserCertCaName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationSettingDevRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationSettingDevRangeName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationSettingDevRangeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAuthenticationSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("active_auth_scheme", flattenAuthenticationSettingActiveAuthScheme(o["active-auth-scheme"], d, "active_auth_scheme", sv)); err != nil {
		if !fortiAPIPatch(o["active-auth-scheme"]) {
			return fmt.Errorf("Error reading active_auth_scheme: %v", err)
		}
	}

	if err = d.Set("sso_auth_scheme", flattenAuthenticationSettingSsoAuthScheme(o["sso-auth-scheme"], d, "sso_auth_scheme", sv)); err != nil {
		if !fortiAPIPatch(o["sso-auth-scheme"]) {
			return fmt.Errorf("Error reading sso_auth_scheme: %v", err)
		}
	}

	if err = d.Set("update_time", flattenAuthenticationSettingUpdateTime(o["update-time"], d, "update_time", sv)); err != nil {
		if !fortiAPIPatch(o["update-time"]) {
			return fmt.Errorf("Error reading update_time: %v", err)
		}
	}

	if err = d.Set("persistent_cookie", flattenAuthenticationSettingPersistentCookie(o["persistent-cookie"], d, "persistent_cookie", sv)); err != nil {
		if !fortiAPIPatch(o["persistent-cookie"]) {
			return fmt.Errorf("Error reading persistent_cookie: %v", err)
		}
	}

	if err = d.Set("ip_auth_cookie", flattenAuthenticationSettingIpAuthCookie(o["ip-auth-cookie"], d, "ip_auth_cookie", sv)); err != nil {
		if !fortiAPIPatch(o["ip-auth-cookie"]) {
			return fmt.Errorf("Error reading ip_auth_cookie: %v", err)
		}
	}

	if err = d.Set("cookie_max_age", flattenAuthenticationSettingCookieMaxAge(o["cookie-max-age"], d, "cookie_max_age", sv)); err != nil {
		if !fortiAPIPatch(o["cookie-max-age"]) {
			return fmt.Errorf("Error reading cookie_max_age: %v", err)
		}
	}

	if err = d.Set("cookie_refresh_div", flattenAuthenticationSettingCookieRefreshDiv(o["cookie-refresh-div"], d, "cookie_refresh_div", sv)); err != nil {
		if !fortiAPIPatch(o["cookie-refresh-div"]) {
			return fmt.Errorf("Error reading cookie_refresh_div: %v", err)
		}
	}

	if err = d.Set("captive_portal_type", flattenAuthenticationSettingCaptivePortalType(o["captive-portal-type"], d, "captive_portal_type", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-type"]) {
			return fmt.Errorf("Error reading captive_portal_type: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip", flattenAuthenticationSettingCaptivePortalIp(o["captive-portal-ip"], d, "captive_portal_ip", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-ip"]) {
			return fmt.Errorf("Error reading captive_portal_ip: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip6", flattenAuthenticationSettingCaptivePortalIp6(o["captive-portal-ip6"], d, "captive_portal_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-ip6"]) {
			return fmt.Errorf("Error reading captive_portal_ip6: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenAuthenticationSettingCaptivePortal(o["captive-portal"], d, "captive_portal", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal"]) {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("captive_portal6", flattenAuthenticationSettingCaptivePortal6(o["captive-portal6"], d, "captive_portal6", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal6"]) {
			return fmt.Errorf("Error reading captive_portal6: %v", err)
		}
	}

	if err = d.Set("cert_auth", flattenAuthenticationSettingCertAuth(o["cert-auth"], d, "cert_auth", sv)); err != nil {
		if !fortiAPIPatch(o["cert-auth"]) {
			return fmt.Errorf("Error reading cert_auth: %v", err)
		}
	}

	if err = d.Set("cert_captive_portal", flattenAuthenticationSettingCertCaptivePortal(o["cert-captive-portal"], d, "cert_captive_portal", sv)); err != nil {
		if !fortiAPIPatch(o["cert-captive-portal"]) {
			return fmt.Errorf("Error reading cert_captive_portal: %v", err)
		}
	}

	if err = d.Set("cert_captive_portal_ip", flattenAuthenticationSettingCertCaptivePortalIp(o["cert-captive-portal-ip"], d, "cert_captive_portal_ip", sv)); err != nil {
		if !fortiAPIPatch(o["cert-captive-portal-ip"]) {
			return fmt.Errorf("Error reading cert_captive_portal_ip: %v", err)
		}
	}

	if err = d.Set("cert_captive_portal_port", flattenAuthenticationSettingCertCaptivePortalPort(o["cert-captive-portal-port"], d, "cert_captive_portal_port", sv)); err != nil {
		if !fortiAPIPatch(o["cert-captive-portal-port"]) {
			return fmt.Errorf("Error reading cert_captive_portal_port: %v", err)
		}
	}

	if err = d.Set("captive_portal_port", flattenAuthenticationSettingCaptivePortalPort(o["captive-portal-port"], d, "captive_portal_port", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-port"]) {
			return fmt.Errorf("Error reading captive_portal_port: %v", err)
		}
	}

	if err = d.Set("auth_https", flattenAuthenticationSettingAuthHttps(o["auth-https"], d, "auth_https", sv)); err != nil {
		if !fortiAPIPatch(o["auth-https"]) {
			return fmt.Errorf("Error reading auth_https: %v", err)
		}
	}

	if err = d.Set("captive_portal_ssl_port", flattenAuthenticationSettingCaptivePortalSslPort(o["captive-portal-ssl-port"], d, "captive_portal_ssl_port", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-ssl-port"]) {
			return fmt.Errorf("Error reading captive_portal_ssl_port: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("user_cert_ca", flattenAuthenticationSettingUserCertCa(o["user-cert-ca"], d, "user_cert_ca", sv)); err != nil {
			if !fortiAPIPatch(o["user-cert-ca"]) {
				return fmt.Errorf("Error reading user_cert_ca: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user_cert_ca"); ok {
			if err = d.Set("user_cert_ca", flattenAuthenticationSettingUserCertCa(o["user-cert-ca"], d, "user_cert_ca", sv)); err != nil {
				if !fortiAPIPatch(o["user-cert-ca"]) {
					return fmt.Errorf("Error reading user_cert_ca: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dev_range", flattenAuthenticationSettingDevRange(o["dev-range"], d, "dev_range", sv)); err != nil {
			if !fortiAPIPatch(o["dev-range"]) {
				return fmt.Errorf("Error reading dev_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dev_range"); ok {
			if err = d.Set("dev_range", flattenAuthenticationSettingDevRange(o["dev-range"], d, "dev_range", sv)); err != nil {
				if !fortiAPIPatch(o["dev-range"]) {
					return fmt.Errorf("Error reading dev_range: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenAuthenticationSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAuthenticationSettingActiveAuthScheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingSsoAuthScheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingUpdateTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingPersistentCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingIpAuthCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCookieMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCookieRefreshDiv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortal6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCertAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCertCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCertCaptivePortalIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCertCaptivePortalPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingAuthHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalSslPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingUserCertCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationSettingUserCertCaName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationSettingUserCertCaName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingDevRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationSettingDevRangeName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationSettingDevRangeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAuthenticationSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("active_auth_scheme"); ok {
		if setArgNil {
			obj["active-auth-scheme"] = nil
		} else {
			t, err := expandAuthenticationSettingActiveAuthScheme(d, v, "active_auth_scheme", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["active-auth-scheme"] = t
			}
		}
	} else if d.HasChange("active_auth_scheme") {
		obj["active-auth-scheme"] = nil
	}

	if v, ok := d.GetOk("sso_auth_scheme"); ok {
		if setArgNil {
			obj["sso-auth-scheme"] = nil
		} else {
			t, err := expandAuthenticationSettingSsoAuthScheme(d, v, "sso_auth_scheme", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sso-auth-scheme"] = t
			}
		}
	} else if d.HasChange("sso_auth_scheme") {
		obj["sso-auth-scheme"] = nil
	}

	if v, ok := d.GetOk("update_time"); ok {
		if setArgNil {
			obj["update-time"] = nil
		} else {
			t, err := expandAuthenticationSettingUpdateTime(d, v, "update_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("persistent_cookie"); ok {
		if setArgNil {
			obj["persistent-cookie"] = nil
		} else {
			t, err := expandAuthenticationSettingPersistentCookie(d, v, "persistent_cookie", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["persistent-cookie"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_auth_cookie"); ok {
		if setArgNil {
			obj["ip-auth-cookie"] = nil
		} else {
			t, err := expandAuthenticationSettingIpAuthCookie(d, v, "ip_auth_cookie", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-auth-cookie"] = t
			}
		}
	}

	if v, ok := d.GetOk("cookie_max_age"); ok {
		if setArgNil {
			obj["cookie-max-age"] = nil
		} else {
			t, err := expandAuthenticationSettingCookieMaxAge(d, v, "cookie_max_age", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cookie-max-age"] = t
			}
		}
	}

	if v, ok := d.GetOk("cookie_refresh_div"); ok {
		if setArgNil {
			obj["cookie-refresh-div"] = nil
		} else {
			t, err := expandAuthenticationSettingCookieRefreshDiv(d, v, "cookie_refresh_div", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cookie-refresh-div"] = t
			}
		}
	}

	if v, ok := d.GetOk("captive_portal_type"); ok {
		if setArgNil {
			obj["captive-portal-type"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortalType(d, v, "captive_portal_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("captive_portal_ip"); ok {
		if setArgNil {
			obj["captive-portal-ip"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortalIp(d, v, "captive_portal_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("captive_portal_ip6"); ok {
		if setArgNil {
			obj["captive-portal-ip6"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortalIp6(d, v, "captive_portal_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok {
		if setArgNil {
			obj["captive-portal"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortal(d, v, "captive_portal", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal"] = t
			}
		}
	} else if d.HasChange("captive_portal") {
		obj["captive-portal"] = nil
	}

	if v, ok := d.GetOk("captive_portal6"); ok {
		if setArgNil {
			obj["captive-portal6"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortal6(d, v, "captive_portal6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal6"] = t
			}
		}
	} else if d.HasChange("captive_portal6") {
		obj["captive-portal6"] = nil
	}

	if v, ok := d.GetOk("cert_auth"); ok {
		if setArgNil {
			obj["cert-auth"] = nil
		} else {
			t, err := expandAuthenticationSettingCertAuth(d, v, "cert_auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert-auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("cert_captive_portal"); ok {
		if setArgNil {
			obj["cert-captive-portal"] = nil
		} else {
			t, err := expandAuthenticationSettingCertCaptivePortal(d, v, "cert_captive_portal", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert-captive-portal"] = t
			}
		}
	} else if d.HasChange("cert_captive_portal") {
		obj["cert-captive-portal"] = nil
	}

	if v, ok := d.GetOk("cert_captive_portal_ip"); ok {
		if setArgNil {
			obj["cert-captive-portal-ip"] = nil
		} else {
			t, err := expandAuthenticationSettingCertCaptivePortalIp(d, v, "cert_captive_portal_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert-captive-portal-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("cert_captive_portal_port"); ok {
		if setArgNil {
			obj["cert-captive-portal-port"] = nil
		} else {
			t, err := expandAuthenticationSettingCertCaptivePortalPort(d, v, "cert_captive_portal_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert-captive-portal-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("captive_portal_port"); ok {
		if setArgNil {
			obj["captive-portal-port"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortalPort(d, v, "captive_portal_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_https"); ok {
		if setArgNil {
			obj["auth-https"] = nil
		} else {
			t, err := expandAuthenticationSettingAuthHttps(d, v, "auth_https", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-https"] = t
			}
		}
	}

	if v, ok := d.GetOk("captive_portal_ssl_port"); ok {
		if setArgNil {
			obj["captive-portal-ssl-port"] = nil
		} else {
			t, err := expandAuthenticationSettingCaptivePortalSslPort(d, v, "captive_portal_ssl_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["captive-portal-ssl-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("user_cert_ca"); ok || d.HasChange("user_cert_ca") {
		if setArgNil {
			obj["user-cert-ca"] = make([]struct{}, 0)
		} else {
			t, err := expandAuthenticationSettingUserCertCa(d, v, "user_cert_ca", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["user-cert-ca"] = t
			}
		}
	}

	if v, ok := d.GetOk("dev_range"); ok || d.HasChange("dev_range") {
		if setArgNil {
			obj["dev-range"] = make([]struct{}, 0)
		} else {
			t, err := expandAuthenticationSettingDevRange(d, v, "dev_range", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dev-range"] = t
			}
		}
	}

	return &obj, nil
}
