// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure peer users.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPeerCreate,
		Read:   resourceUserPeerRead,
		Update: resourceUserPeerUpdate,
		Delete: resourceUserPeerDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"mandatory_ca_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"subject": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cn_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mfa_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mfa_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"mfa_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"mfa_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ldap_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ldap_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ldap_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_override_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceUserPeerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserPeer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserPeer resource while getting object: %v", err)
	}

	o, err := c.CreateUserPeer(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserPeer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserPeer")
	}

	return resourceUserPeerRead(d, m)
}

func resourceUserPeerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserPeer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserPeer resource while getting object: %v", err)
	}

	o, err := c.UpdateUserPeer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserPeer")
	}

	return resourceUserPeerRead(d, m)
}

func resourceUserPeerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserPeer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserPeer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserPeer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserPeer resource from API: %v", err)
	}
	return nil
}

func flattenUserPeerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerMandatoryCaVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerSubject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerCn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerCnType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerMfaMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerMfaServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerMfaUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerLdapMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerOcspOverrideServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserPeerTwoFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserPeer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserPeerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mandatory_ca_verify", flattenUserPeerMandatoryCaVerify(o["mandatory-ca-verify"], d, "mandatory_ca_verify", sv)); err != nil {
		if !fortiAPIPatch(o["mandatory-ca-verify"]) {
			return fmt.Errorf("Error reading mandatory_ca_verify: %v", err)
		}
	}

	if err = d.Set("ca", flattenUserPeerCa(o["ca"], d, "ca", sv)); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("subject", flattenUserPeerSubject(o["subject"], d, "subject", sv)); err != nil {
		if !fortiAPIPatch(o["subject"]) {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("cn", flattenUserPeerCn(o["cn"], d, "cn", sv)); err != nil {
		if !fortiAPIPatch(o["cn"]) {
			return fmt.Errorf("Error reading cn: %v", err)
		}
	}

	if err = d.Set("cn_type", flattenUserPeerCnType(o["cn-type"], d, "cn_type", sv)); err != nil {
		if !fortiAPIPatch(o["cn-type"]) {
			return fmt.Errorf("Error reading cn_type: %v", err)
		}
	}

	if err = d.Set("mfa_mode", flattenUserPeerMfaMode(o["mfa-mode"], d, "mfa_mode", sv)); err != nil {
		if !fortiAPIPatch(o["mfa-mode"]) {
			return fmt.Errorf("Error reading mfa_mode: %v", err)
		}
	}

	if err = d.Set("mfa_server", flattenUserPeerMfaServer(o["mfa-server"], d, "mfa_server", sv)); err != nil {
		if !fortiAPIPatch(o["mfa-server"]) {
			return fmt.Errorf("Error reading mfa_server: %v", err)
		}
	}

	if err = d.Set("mfa_username", flattenUserPeerMfaUsername(o["mfa-username"], d, "mfa_username", sv)); err != nil {
		if !fortiAPIPatch(o["mfa-username"]) {
			return fmt.Errorf("Error reading mfa_username: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserPeerLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenUserPeerLdapUsername(o["ldap-username"], d, "ldap_username", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-username"]) {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("ldap_mode", flattenUserPeerLdapMode(o["ldap-mode"], d, "ldap_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-mode"]) {
			return fmt.Errorf("Error reading ldap_mode: %v", err)
		}
	}

	if err = d.Set("ocsp_override_server", flattenUserPeerOcspOverrideServer(o["ocsp-override-server"], d, "ocsp_override_server", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-override-server"]) {
			return fmt.Errorf("Error reading ocsp_override_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserPeerTwoFactor(o["two-factor"], d, "two_factor", sv)); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	return nil
}

func flattenUserPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserPeerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMandatoryCaVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerSubject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCnType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMfaMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMfaServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMfaUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMfaPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerOcspOverrideServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerTwoFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserPeerPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserPeer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserPeerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mandatory_ca_verify"); ok {
		t, err := expandUserPeerMandatoryCaVerify(d, v, "mandatory_ca_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mandatory-ca-verify"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandUserPeerCa(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	} else if d.HasChange("ca") {
		obj["ca"] = nil
	}

	if v, ok := d.GetOk("subject"); ok {
		t, err := expandUserPeerSubject(d, v, "subject", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject"] = t
		}
	} else if d.HasChange("subject") {
		obj["subject"] = nil
	}

	if v, ok := d.GetOk("cn"); ok {
		t, err := expandUserPeerCn(d, v, "cn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn"] = t
		}
	} else if d.HasChange("cn") {
		obj["cn"] = nil
	}

	if v, ok := d.GetOk("cn_type"); ok {
		t, err := expandUserPeerCnType(d, v, "cn_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-type"] = t
		}
	}

	if v, ok := d.GetOk("mfa_mode"); ok {
		t, err := expandUserPeerMfaMode(d, v, "mfa_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-mode"] = t
		}
	}

	if v, ok := d.GetOk("mfa_server"); ok {
		t, err := expandUserPeerMfaServer(d, v, "mfa_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-server"] = t
		}
	} else if d.HasChange("mfa_server") {
		obj["mfa-server"] = nil
	}

	if v, ok := d.GetOk("mfa_username"); ok {
		t, err := expandUserPeerMfaUsername(d, v, "mfa_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-username"] = t
		}
	} else if d.HasChange("mfa_username") {
		obj["mfa-username"] = nil
	}

	if v, ok := d.GetOk("mfa_password"); ok {
		t, err := expandUserPeerMfaPassword(d, v, "mfa_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mfa-password"] = t
		}
	} else if d.HasChange("mfa_password") {
		obj["mfa-password"] = nil
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserPeerLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	} else if d.HasChange("ldap_server") {
		obj["ldap-server"] = nil
	}

	if v, ok := d.GetOk("ldap_username"); ok {
		t, err := expandUserPeerLdapUsername(d, v, "ldap_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	} else if d.HasChange("ldap_username") {
		obj["ldap-username"] = nil
	}

	if v, ok := d.GetOk("ldap_password"); ok {
		t, err := expandUserPeerLdapPassword(d, v, "ldap_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	} else if d.HasChange("ldap_password") {
		obj["ldap-password"] = nil
	}

	if v, ok := d.GetOk("ldap_mode"); ok {
		t, err := expandUserPeerLdapMode(d, v, "ldap_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-mode"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_override_server"); ok {
		t, err := expandUserPeerOcspOverrideServer(d, v, "ocsp_override_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-override-server"] = t
		}
	} else if d.HasChange("ocsp_override_server") {
		obj["ocsp-override-server"] = nil
	}

	if v, ok := d.GetOk("two_factor"); ok {
		t, err := expandUserPeerTwoFactor(d, v, "two_factor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok {
		t, err := expandUserPeerPasswd(d, v, "passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	} else if d.HasChange("passwd") {
		obj["passwd"] = nil
	}

	return &obj, nil
}
