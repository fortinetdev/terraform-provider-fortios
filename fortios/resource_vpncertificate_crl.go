// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Certificate Revocation List as a PEM file.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateCrlCreate,
		Read:   resourceVpnCertificateCrlRead,
		Update: resourceVpnCertificateCrlUpdate,
		Delete: resourceVpnCertificateCrlDelete,

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
				Required:     true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ldap_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"ldap_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"http_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"scep_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceVpnCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCrl resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateCrl(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateCrl resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateCrl")
	}

	return resourceVpnCertificateCrlRead(d, m)
}

func resourceVpnCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCrl resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateCrl(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateCrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateCrl")
	}

	return resourceVpnCertificateCrlRead(d, m)
}

func resourceVpnCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnCertificateCrl(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateCrl(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateCrl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateCrl resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateCrlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlCrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlUpdateVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlHttpUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlScepCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnCertificateCrlSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateCrlLastUpdated(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectVpnCertificateCrl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateCrlName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("crl", flattenVpnCertificateCrlCrl(o["crl"], d, "crl", sv)); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateCrlRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateCrlSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("update_vdom", flattenVpnCertificateCrlUpdateVdom(o["update-vdom"], d, "update_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["update-vdom"]) {
			return fmt.Errorf("Error reading update_vdom: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenVpnCertificateCrlLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenVpnCertificateCrlLdapUsername(o["ldap-username"], d, "ldap_username", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-username"]) {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("http_url", flattenVpnCertificateCrlHttpUrl(o["http-url"], d, "http_url", sv)); err != nil {
		if !fortiAPIPatch(o["http-url"]) {
			return fmt.Errorf("Error reading http_url: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateCrlScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("scep_cert", flattenVpnCertificateCrlScepCert(o["scep-cert"], d, "scep_cert", sv)); err != nil {
		if !fortiAPIPatch(o["scep-cert"]) {
			return fmt.Errorf("Error reading scep_cert: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenVpnCertificateCrlUpdateInterval(o["update-interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateCrlSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateCrlLastUpdated(o["last-updated"], d, "last_updated", sv)); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateCrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnCertificateCrlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlCrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return toCertFormat(v), nil
}

func expandVpnCertificateCrlRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlUpdateVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlHttpUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlScepCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateCrlLastUpdated(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateCrl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnCertificateCrlName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("crl"); ok {
		t, err := expandVpnCertificateCrlCrl(d, v, "crl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	} else if d.HasChange("crl") {
		obj["crl"] = nil
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandVpnCertificateCrlRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandVpnCertificateCrlSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("update_vdom"); ok {
		t, err := expandVpnCertificateCrlUpdateVdom(d, v, "update_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-vdom"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandVpnCertificateCrlLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	} else if d.HasChange("ldap_server") {
		obj["ldap-server"] = nil
	}

	if v, ok := d.GetOk("ldap_username"); ok {
		t, err := expandVpnCertificateCrlLdapUsername(d, v, "ldap_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	} else if d.HasChange("ldap_username") {
		obj["ldap-username"] = nil
	}

	if v, ok := d.GetOk("ldap_password"); ok {
		t, err := expandVpnCertificateCrlLdapPassword(d, v, "ldap_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	} else if d.HasChange("ldap_password") {
		obj["ldap-password"] = nil
	}

	if v, ok := d.GetOk("http_url"); ok {
		t, err := expandVpnCertificateCrlHttpUrl(d, v, "http_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-url"] = t
		}
	} else if d.HasChange("http_url") {
		obj["http-url"] = nil
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandVpnCertificateCrlScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	} else if d.HasChange("scep_url") {
		obj["scep-url"] = nil
	}

	if v, ok := d.GetOk("scep_cert"); ok {
		t, err := expandVpnCertificateCrlScepCert(d, v, "scep_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-cert"] = t
		}
	}

	if v, ok := d.GetOkExists("update_interval"); ok {
		t, err := expandVpnCertificateCrlUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	} else if d.HasChange("update_interval") {
		obj["update-interval"] = nil
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandVpnCertificateCrlSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("last_updated"); ok {
		t, err := expandVpnCertificateCrlLastUpdated(d, v, "last_updated", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	} else if d.HasChange("last_updated") {
		obj["last-updated"] = nil
	}

	return &obj, nil
}
