// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure authentication setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"active_auth_scheme": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"sso_auth_scheme": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"captive_portal_type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ip6": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
				Computed: true,
			},
			"captive_portal6": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
				Computed: true,
			},
			"captive_portal_port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional: true,
				Computed: true,
			},
			"auth_https": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_ssl_port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceAuthenticationSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectAuthenticationSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateAuthenticationSetting(obj, mkey)
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

	err := c.DeleteAuthenticationSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting AuthenticationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadAuthenticationSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationSetting resource from API: %v", err)
	}
	return nil
}


func flattenAuthenticationSettingActiveAuthScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingSsoAuthScheme(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortal6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingAuthHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenAuthenticationSettingCaptivePortalSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectAuthenticationSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("active_auth_scheme", flattenAuthenticationSettingActiveAuthScheme(o["active-auth-scheme"], d, "active_auth_scheme")); err != nil {
		if !fortiAPIPatch(o["active-auth-scheme"]) {
			return fmt.Errorf("Error reading active_auth_scheme: %v", err)
		}
	}

	if err = d.Set("sso_auth_scheme", flattenAuthenticationSettingSsoAuthScheme(o["sso-auth-scheme"], d, "sso_auth_scheme")); err != nil {
		if !fortiAPIPatch(o["sso-auth-scheme"]) {
			return fmt.Errorf("Error reading sso_auth_scheme: %v", err)
		}
	}

	if err = d.Set("captive_portal_type", flattenAuthenticationSettingCaptivePortalType(o["captive-portal-type"], d, "captive_portal_type")); err != nil {
		if !fortiAPIPatch(o["captive-portal-type"]) {
			return fmt.Errorf("Error reading captive_portal_type: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip", flattenAuthenticationSettingCaptivePortalIp(o["captive-portal-ip"], d, "captive_portal_ip")); err != nil {
		if !fortiAPIPatch(o["captive-portal-ip"]) {
			return fmt.Errorf("Error reading captive_portal_ip: %v", err)
		}
	}

	if err = d.Set("captive_portal_ip6", flattenAuthenticationSettingCaptivePortalIp6(o["captive-portal-ip6"], d, "captive_portal_ip6")); err != nil {
		if !fortiAPIPatch(o["captive-portal-ip6"]) {
			return fmt.Errorf("Error reading captive_portal_ip6: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenAuthenticationSettingCaptivePortal(o["captive-portal"], d, "captive_portal")); err != nil {
		if !fortiAPIPatch(o["captive-portal"]) {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("captive_portal6", flattenAuthenticationSettingCaptivePortal6(o["captive-portal6"], d, "captive_portal6")); err != nil {
		if !fortiAPIPatch(o["captive-portal6"]) {
			return fmt.Errorf("Error reading captive_portal6: %v", err)
		}
	}

	if err = d.Set("captive_portal_port", flattenAuthenticationSettingCaptivePortalPort(o["captive-portal-port"], d, "captive_portal_port")); err != nil {
		if !fortiAPIPatch(o["captive-portal-port"]) {
			return fmt.Errorf("Error reading captive_portal_port: %v", err)
		}
	}

	if err = d.Set("auth_https", flattenAuthenticationSettingAuthHttps(o["auth-https"], d, "auth_https")); err != nil {
		if !fortiAPIPatch(o["auth-https"]) {
			return fmt.Errorf("Error reading auth_https: %v", err)
		}
	}

	if err = d.Set("captive_portal_ssl_port", flattenAuthenticationSettingCaptivePortalSslPort(o["captive-portal-ssl-port"], d, "captive_portal_ssl_port")); err != nil {
		if !fortiAPIPatch(o["captive-portal-ssl-port"]) {
			return fmt.Errorf("Error reading captive_portal_ssl_port: %v", err)
		}
	}


	return nil
}

func flattenAuthenticationSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandAuthenticationSettingActiveAuthScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingSsoAuthScheme(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortal6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingAuthHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationSettingCaptivePortalSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectAuthenticationSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("active_auth_scheme"); ok {
		t, err := expandAuthenticationSettingActiveAuthScheme(d, v, "active_auth_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-scheme"] = t
		}
	}

	if v, ok := d.GetOk("sso_auth_scheme"); ok {
		t, err := expandAuthenticationSettingSsoAuthScheme(d, v, "sso_auth_scheme")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-scheme"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_type"); ok {
		t, err := expandAuthenticationSettingCaptivePortalType(d, v, "captive_portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-type"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ip"); ok {
		t, err := expandAuthenticationSettingCaptivePortalIp(d, v, "captive_portal_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ip"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ip6"); ok {
		t, err := expandAuthenticationSettingCaptivePortalIp6(d, v, "captive_portal_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ip6"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok {
		t, err := expandAuthenticationSettingCaptivePortal(d, v, "captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal6"); ok {
		t, err := expandAuthenticationSettingCaptivePortal6(d, v, "captive_portal6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal6"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_port"); ok {
		t, err := expandAuthenticationSettingCaptivePortalPort(d, v, "captive_portal_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_https"); ok {
		t, err := expandAuthenticationSettingAuthHttps(d, v, "auth_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-https"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ssl_port"); ok {
		t, err := expandAuthenticationSettingCaptivePortalSslPort(d, v, "captive_portal_ssl_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ssl-port"] = t
		}
	}


	return &obj, nil
}

