// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure local users.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserLocalCreate,
		Read:   resourceUserLocalRead,
		Update: resourceUserLocalUpdate,
		Delete: resourceUserLocalDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 64),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"tacacs_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"two_factor_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortitoken": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
				Computed:     true,
			},
			"email_to": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"sms_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sms_custom_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"sms_phone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"passwd_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"passwd_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1440),
				Optional:     true,
				Computed:     true,
			},
			"workstation": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_concurrent_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_concurrent_value": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"ppk_secret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ppk_identity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"qkd_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"username_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_insensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username_case_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserLocalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserLocal resource while getting object: %v", err)
	}

	o, err := c.CreateUserLocal(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserLocal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLocal")
	}

	return resourceUserLocalRead(d, m)
}

func resourceUserLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserLocal resource while getting object: %v", err)
	}

	o, err := c.UpdateUserLocal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserLocal")
	}

	return resourceUserLocalRead(d, m)
}

func resourceUserLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserLocal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserLocal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLocal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserLocal resource from API: %v", err)
	}
	return nil
}

func flattenUserLocalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalPasswd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalTacacsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalTwoFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalTwoFactorAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalTwoFactorNotification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalFortitoken(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalEmailTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalSmsServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalSmsCustomServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalSmsPhone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalPasswdPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalPasswdTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalAuthtimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalWorkstation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalAuthConcurrentOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalAuthConcurrentValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalPpkSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalPpkIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalQkdProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalUsernameSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalUsernameCaseInsensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserLocalUsernameCaseSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserLocal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserLocalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserLocalId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenUserLocalStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenUserLocalType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserLocalLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenUserLocalRadiusServer(o["radius-server"], d, "radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["radius-server"]) {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("tacacs_server", flattenUserLocalTacacsServer(o["tacacs+-server"], d, "tacacs_server", sv)); err != nil {
		if !fortiAPIPatch(o["tacacs+-server"]) {
			return fmt.Errorf("Error reading tacacs_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserLocalTwoFactor(o["two-factor"], d, "two_factor", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	if err = d.Set("two_factor_authentication", flattenUserLocalTwoFactorAuthentication(o["two-factor-authentication"], d, "two_factor_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-authentication"]) {
			return fmt.Errorf("Error reading two_factor_authentication: %v", err)
		}
	}

	if err = d.Set("two_factor_notification", flattenUserLocalTwoFactorNotification(o["two-factor-notification"], d, "two_factor_notification", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor-notification"]) {
			return fmt.Errorf("Error reading two_factor_notification: %v", err)
		}
	}

	if err = d.Set("fortitoken", flattenUserLocalFortitoken(o["fortitoken"], d, "fortitoken", sv)); err != nil {
		if !fortiAPIPatch(o["fortitoken"]) {
			return fmt.Errorf("Error reading fortitoken: %v", err)
		}
	}

	if err = d.Set("email_to", flattenUserLocalEmailTo(o["email-to"], d, "email_to", sv)); err != nil {
		if !fortiAPIPatch(o["email-to"]) {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("sms_server", flattenUserLocalSmsServer(o["sms-server"], d, "sms_server", sv)); err != nil {
		if !fortiAPIPatch(o["sms-server"]) {
			return fmt.Errorf("Error reading sms_server: %v", err)
		}
	}

	if err = d.Set("sms_custom_server", flattenUserLocalSmsCustomServer(o["sms-custom-server"], d, "sms_custom_server", sv)); err != nil {
		if !fortiAPIPatch(o["sms-custom-server"]) {
			return fmt.Errorf("Error reading sms_custom_server: %v", err)
		}
	}

	if err = d.Set("sms_phone", flattenUserLocalSmsPhone(o["sms-phone"], d, "sms_phone", sv)); err != nil {
		if !fortiAPIPatch(o["sms-phone"]) {
			return fmt.Errorf("Error reading sms_phone: %v", err)
		}
	}

	if err = d.Set("passwd_policy", flattenUserLocalPasswdPolicy(o["passwd-policy"], d, "passwd_policy", sv)); err != nil {
		if !fortiAPIPatch(o["passwd-policy"]) {
			return fmt.Errorf("Error reading passwd_policy: %v", err)
		}
	}

	if err = d.Set("passwd_time", flattenUserLocalPasswdTime(o["passwd-time"], d, "passwd_time", sv)); err != nil {
		if !fortiAPIPatch(o["passwd-time"]) {
			return fmt.Errorf("Error reading passwd_time: %v", err)
		}
	}

	if err = d.Set("authtimeout", flattenUserLocalAuthtimeout(o["authtimeout"], d, "authtimeout", sv)); err != nil {
		if !fortiAPIPatch(o["authtimeout"]) {
			return fmt.Errorf("Error reading authtimeout: %v", err)
		}
	}

	if err = d.Set("workstation", flattenUserLocalWorkstation(o["workstation"], d, "workstation", sv)); err != nil {
		if !fortiAPIPatch(o["workstation"]) {
			return fmt.Errorf("Error reading workstation: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_override", flattenUserLocalAuthConcurrentOverride(o["auth-concurrent-override"], d, "auth_concurrent_override", sv)); err != nil {
		if !fortiAPIPatch(o["auth-concurrent-override"]) {
			return fmt.Errorf("Error reading auth_concurrent_override: %v", err)
		}
	}

	if err = d.Set("auth_concurrent_value", flattenUserLocalAuthConcurrentValue(o["auth-concurrent-value"], d, "auth_concurrent_value", sv)); err != nil {
		if !fortiAPIPatch(o["auth-concurrent-value"]) {
			return fmt.Errorf("Error reading auth_concurrent_value: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenUserLocalPpkIdentity(o["ppk-identity"], d, "ppk_identity", sv)); err != nil {
		if !fortiAPIPatch(o["ppk-identity"]) {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenUserLocalQkdProfile(o["qkd-profile"], d, "qkd_profile", sv)); err != nil {
		if !fortiAPIPatch(o["qkd-profile"]) {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("username_sensitivity", flattenUserLocalUsernameSensitivity(o["username-sensitivity"], d, "username_sensitivity", sv)); err != nil {
		if !fortiAPIPatch(o["username-sensitivity"]) {
			return fmt.Errorf("Error reading username_sensitivity: %v", err)
		}
	}

	if err = d.Set("username_case_insensitivity", flattenUserLocalUsernameCaseInsensitivity(o["username-case-insensitivity"], d, "username_case_insensitivity", sv)); err != nil {
		if !fortiAPIPatch(o["username-case-insensitivity"]) {
			return fmt.Errorf("Error reading username_case_insensitivity: %v", err)
		}
	}

	if err = d.Set("username_case_sensitivity", flattenUserLocalUsernameCaseSensitivity(o["username-case-sensitivity"], d, "username_case_sensitivity", sv)); err != nil {
		if !fortiAPIPatch(o["username-case-sensitivity"]) {
			return fmt.Errorf("Error reading username_case_sensitivity: %v", err)
		}
	}

	return nil
}

func flattenUserLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserLocalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTacacsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTwoFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTwoFactorAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalTwoFactorNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalFortitoken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalEmailTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalSmsServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalSmsCustomServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalSmsPhone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPasswdPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPasswdTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalAuthtimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalWorkstation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalAuthConcurrentOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalAuthConcurrentValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPpkSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalPpkIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalQkdProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalUsernameSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalUsernameCaseInsensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserLocalUsernameCaseSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserLocalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandUserLocalId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandUserLocalStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserLocalType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok {
		t, err := expandUserLocalPasswd(d, v, "passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserLocalLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok {
		t, err := expandUserLocalRadiusServer(d, v, "radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("tacacs_server"); ok {
		t, err := expandUserLocalTacacsServer(d, v, "tacacs_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs+-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok {
		t, err := expandUserLocalTwoFactor(d, v, "two_factor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_authentication"); ok {
		t, err := expandUserLocalTwoFactorAuthentication(d, v, "two_factor_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-authentication"] = t
		}
	}

	if v, ok := d.GetOk("two_factor_notification"); ok {
		t, err := expandUserLocalTwoFactorNotification(d, v, "two_factor_notification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor-notification"] = t
		}
	}

	if v, ok := d.GetOk("fortitoken"); ok {
		t, err := expandUserLocalFortitoken(d, v, "fortitoken", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortitoken"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok {
		t, err := expandUserLocalEmailTo(d, v, "email_to", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("sms_server"); ok {
		t, err := expandUserLocalSmsServer(d, v, "sms_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_custom_server"); ok {
		t, err := expandUserLocalSmsCustomServer(d, v, "sms_custom_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-custom-server"] = t
		}
	}

	if v, ok := d.GetOk("sms_phone"); ok {
		t, err := expandUserLocalSmsPhone(d, v, "sms_phone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sms-phone"] = t
		}
	}

	if v, ok := d.GetOk("passwd_policy"); ok {
		t, err := expandUserLocalPasswdPolicy(d, v, "passwd_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-policy"] = t
		}
	}

	if v, ok := d.GetOk("passwd_time"); ok {
		t, err := expandUserLocalPasswdTime(d, v, "passwd_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-time"] = t
		}
	}

	if v, ok := d.GetOkExists("authtimeout"); ok {
		t, err := expandUserLocalAuthtimeout(d, v, "authtimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authtimeout"] = t
		}
	}

	if v, ok := d.GetOk("workstation"); ok {
		t, err := expandUserLocalWorkstation(d, v, "workstation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["workstation"] = t
		}
	}

	if v, ok := d.GetOk("auth_concurrent_override"); ok {
		t, err := expandUserLocalAuthConcurrentOverride(d, v, "auth_concurrent_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-override"] = t
		}
	}

	if v, ok := d.GetOkExists("auth_concurrent_value"); ok {
		t, err := expandUserLocalAuthConcurrentValue(d, v, "auth_concurrent_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-concurrent-value"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok {
		t, err := expandUserLocalPpkSecret(d, v, "ppk_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok {
		t, err := expandUserLocalPpkIdentity(d, v, "ppk_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok {
		t, err := expandUserLocalQkdProfile(d, v, "qkd_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	}

	if v, ok := d.GetOk("username_sensitivity"); ok {
		t, err := expandUserLocalUsernameSensitivity(d, v, "username_sensitivity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_case_insensitivity"); ok {
		t, err := expandUserLocalUsernameCaseInsensitivity(d, v, "username_case_insensitivity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-insensitivity"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitivity"); ok {
		t, err := expandUserLocalUsernameCaseSensitivity(d, v, "username_case_sensitivity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitivity"] = t
		}
	}

	return &obj, nil
}
