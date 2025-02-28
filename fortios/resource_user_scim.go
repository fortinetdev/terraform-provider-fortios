// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SCIM client entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserScim() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserScimCreate,
		Read:   resourceUserScimRead,
		Update: resourceUserScimUpdate,
		Delete: resourceUserScimDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"token_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"client_authentication_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_secret_token": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"client_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserScimCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserScim(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserScim resource while getting object: %v", err)
	}

	o, err := c.CreateUserScim(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserScim resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserScim")
	}

	return resourceUserScimRead(d, m)
}

func resourceUserScimUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserScim(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserScim resource while getting object: %v", err)
	}

	o, err := c.UpdateUserScim(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserScim resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserScim")
	}

	return resourceUserScimRead(d, m)
}

func resourceUserScimDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserScim(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserScim resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserScimRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserScim(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserScim resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserScim(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserScim resource from API: %v", err)
	}
	return nil
}

func flattenUserScimName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserScimStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimBaseUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimAuthMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimTokenCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimClientAuthenticationMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimClientSecretToken(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserScimClientIdentityCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserScim(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserScimName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserScimId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenUserScimStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("base_url", flattenUserScimBaseUrl(o["base-url"], d, "base_url", sv)); err != nil {
		if !fortiAPIPatch(o["base-url"]) {
			return fmt.Errorf("Error reading base_url: %v", err)
		}
	}

	if err = d.Set("auth_method", flattenUserScimAuthMethod(o["auth-method"], d, "auth_method", sv)); err != nil {
		if !fortiAPIPatch(o["auth-method"]) {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("token_certificate", flattenUserScimTokenCertificate(o["token-certificate"], d, "token_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["token-certificate"]) {
			return fmt.Errorf("Error reading token_certificate: %v", err)
		}
	}

	if err = d.Set("client_authentication_method", flattenUserScimClientAuthenticationMethod(o["client-authentication-method"], d, "client_authentication_method", sv)); err != nil {
		if !fortiAPIPatch(o["client-authentication-method"]) {
			return fmt.Errorf("Error reading client_authentication_method: %v", err)
		}
	}

	if err = d.Set("client_secret_token", flattenUserScimClientSecretToken(o["client-secret-token"], d, "client_secret_token", sv)); err != nil {
		if !fortiAPIPatch(o["client-secret-token"]) {
			return fmt.Errorf("Error reading client_secret_token: %v", err)
		}
	}

	if err = d.Set("certificate", flattenUserScimCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("client_identity_check", flattenUserScimClientIdentityCheck(o["client-identity-check"], d, "client_identity_check", sv)); err != nil {
		if !fortiAPIPatch(o["client-identity-check"]) {
			return fmt.Errorf("Error reading client_identity_check: %v", err)
		}
	}

	return nil
}

func flattenUserScimFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserScimName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimBaseUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimAuthMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimTokenCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimClientAuthenticationMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimClientSecretToken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserScimClientIdentityCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserScim(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserScimName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandUserScimId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandUserScimStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("base_url"); ok {
		t, err := expandUserScimBaseUrl(d, v, "base_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base-url"] = t
		}
	} else if d.HasChange("base_url") {
		obj["base-url"] = nil
	}

	if v, ok := d.GetOk("auth_method"); ok {
		t, err := expandUserScimAuthMethod(d, v, "auth_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("token_certificate"); ok {
		t, err := expandUserScimTokenCertificate(d, v, "token_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-certificate"] = t
		}
	} else if d.HasChange("token_certificate") {
		obj["token-certificate"] = nil
	}

	if v, ok := d.GetOk("secret"); ok {
		t, err := expandUserScimSecret(d, v, "secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	} else if d.HasChange("secret") {
		obj["secret"] = nil
	}

	if v, ok := d.GetOk("client_authentication_method"); ok {
		t, err := expandUserScimClientAuthenticationMethod(d, v, "client_authentication_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-authentication-method"] = t
		}
	}

	if v, ok := d.GetOk("client_secret_token"); ok {
		t, err := expandUserScimClientSecretToken(d, v, "client_secret_token", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret-token"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandUserScimCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("client_identity_check"); ok {
		t, err := expandUserScimClientIdentityCheck(d, v, "client_identity_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-identity-check"] = t
		}
	}

	return &obj, nil
}
