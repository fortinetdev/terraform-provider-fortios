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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateCrlCreate,
		Read:   resourceCertificateCrlRead,
		Update: resourceCertificateCrlUpdate,
		Delete: resourceCertificateCrlDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed:     true,
			},
			"ldap_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
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
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CertificateCrl resource while getting object: %v", err)
	}

	o, err := c.CreateCertificateCrl(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CertificateCrl resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateCrl")
	}

	return resourceCertificateCrlRead(d, m)
}

func resourceCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CertificateCrl resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateCrl(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CertificateCrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateCrl")
	}

	return resourceCertificateCrlRead(d, m)
}

func resourceCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCertificateCrl(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadCertificateCrl(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateCrl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCrl resource from API: %v", err)
	}
	return nil
}

func flattenCertificateCrlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlCrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlUpdateVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlLdapUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlLdapPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlHttpUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlScepCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCrlLastUpdated(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateCrl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenCertificateCrlName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("crl", flattenCertificateCrlCrl(o["crl"], d, "crl", sv)); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateCrlRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateCrlSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("update_vdom", flattenCertificateCrlUpdateVdom(o["update-vdom"], d, "update_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["update-vdom"]) {
			return fmt.Errorf("Error reading update_vdom: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenCertificateCrlLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenCertificateCrlLdapUsername(o["ldap-username"], d, "ldap_username", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-username"]) {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("http_url", flattenCertificateCrlHttpUrl(o["http-url"], d, "http_url", sv)); err != nil {
		if !fortiAPIPatch(o["http-url"]) {
			return fmt.Errorf("Error reading http_url: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenCertificateCrlScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("scep_cert", flattenCertificateCrlScepCert(o["scep-cert"], d, "scep_cert", sv)); err != nil {
		if !fortiAPIPatch(o["scep-cert"]) {
			return fmt.Errorf("Error reading scep_cert: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenCertificateCrlUpdateInterval(o["update-interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenCertificateCrlSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenCertificateCrlLastUpdated(o["last-updated"], d, "last_updated", sv)); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	return nil
}

func flattenCertificateCrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCertificateCrlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlCrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlUpdateVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlLdapUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlLdapPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlHttpUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlScepCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlLastUpdated(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateCrl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandCertificateCrlName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("crl"); ok {

		t, err := expandCertificateCrlCrl(d, v, "crl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {

		t, err := expandCertificateCrlRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {

		t, err := expandCertificateCrlSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("update_vdom"); ok {

		t, err := expandCertificateCrlUpdateVdom(d, v, "update_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-vdom"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {

		t, err := expandCertificateCrlLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("ldap_username"); ok {

		t, err := expandCertificateCrlLdapUsername(d, v, "ldap_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	}

	if v, ok := d.GetOk("ldap_password"); ok {

		t, err := expandCertificateCrlLdapPassword(d, v, "ldap_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	}

	if v, ok := d.GetOk("http_url"); ok {

		t, err := expandCertificateCrlHttpUrl(d, v, "http_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-url"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {

		t, err := expandCertificateCrlScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("scep_cert"); ok {

		t, err := expandCertificateCrlScepCert(d, v, "scep_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-cert"] = t
		}
	}

	if v, ok := d.GetOkExists("update_interval"); ok {

		t, err := expandCertificateCrlUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandCertificateCrlSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("last_updated"); ok {

		t, err := expandCertificateCrlLastUpdated(d, v, "last_updated", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	return &obj, nil
}
