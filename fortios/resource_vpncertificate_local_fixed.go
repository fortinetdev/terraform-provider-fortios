// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Local keys and certificates.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	forticlient "github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateLocalFixed() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateLocalFixedCreate,
		Read:   resourceVpnCertificateLocalFixedRead,
		Update: resourceVpnCertificateLocalFixedUpdate,
		Delete: resourceVpnCertificateLocalFixedDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				ForceNew:     true,
				Optional:     true,
				Sensitive:    true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"private_key": &schema.Schema{
				Type:      schema.TypeString,
				ForceNew:  true,
				Optional:  true,
				Sensitive: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if old == strings.TrimSuffix(new, "\n") {
						return true
					}
					return false
				},
			},
			"csr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_regenerate_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_regenerate_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scep_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ca_identifier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"name_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_localid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ike_localid_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enroll_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmp_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"cmp_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"cmp_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"cmp_regeneration_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceVpnCertificateLocalFixedCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := d.Get("name").(string)

	i := &forticlient.JSONSystemVpnCertificateLocalImport{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Certificate: d.Get("certificate").(string),
		Password:    d.Get("password").(string),
		PrivateKey:  d.Get("private_key").(string),
		Scope:       d.Get("range").(string),
	}

	_, err := c.CreateSystemVpnCertificateLocalImport(i)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocalFixedImport resource: %v", err)
	}

	d.SetId(mkey)

	obj, err := getObjectVpnCertificateLocalFixed(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocalFixed resource while getting object: %v", err)
	}

	_, err = c.UpdateCertificateLocal(obj, mkey)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocalFixed resource: %v", err)
	}

	// if o["mkey"] != nil && o["mkey"] != "" {
	// 	d.SetId(o["mkey"].(string))
	// } else {
	// 	d.SetId("VpnCertificateLocalFixed")
	// }

	return resourceVpnCertificateLocalFixedRead(d, m)
}

func resourceVpnCertificateLocalFixedUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnCertificateLocalFixed(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateLocalFixed resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateLocal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateLocalFixed resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateLocalFixed")
	}

	return resourceVpnCertificateLocalFixedRead(d, m)
}

func resourceVpnCertificateLocalFixedDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteCertificateLocal(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateLocalFixed resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateLocalFixedRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	// get certificate from monitor endpoint

	i := &forticlient.JSONSystemCertificateDownload{
		Mkey: mkey,
		Type: "local-cer", // figure out how to set this dynamically
	}

	cert, err := c.ReadSystemCertificateDownload(i)
	if err != nil {
		cert = ""
		// return fmt.Errorf("Error reading VpnCertificateLocalImport resource: %v", err)
	}

	if cert == "" {
		log.Printf("[WARN] certificate (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	// get cert attributes from cmdb endpoint

	o, err := c.ReadCertificateLocal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateLocalFixed resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateLocalFixed(d, o, cert, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateLocalFixed resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateLocalFixedName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedPrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCsr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedAutoRegenerateDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedScepPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCaIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedNameEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedIkeLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedIkeLocalidType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedLastUpdated(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedEnrollProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCmpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCmpPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCmpServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalFixedCmpRegenerationMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnCertificateLocalFixed(d *schema.ResourceData, o map[string]interface{}, cert string, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateLocalFixedName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnCertificateLocalFixedComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("certificate", cert); err != nil {
		if !fortiAPIPatch(cert) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("csr", flattenVpnCertificateLocalFixedCsr(o["csr"], d, "csr", sv)); err != nil {
		if !fortiAPIPatch(o["csr"]) {
			return fmt.Errorf("Error reading csr: %v", err)
		}
	}

	if err = d.Set("state", flattenVpnCertificateLocalFixedState(o["state"], d, "state", sv)); err != nil {
		if !fortiAPIPatch(o["state"]) {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateLocalFixedScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateLocalFixedRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateLocalFixedSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days", flattenVpnCertificateLocalFixedAutoRegenerateDays(o["auto-regenerate-days"], d, "auto_regenerate_days", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days"]) {
			return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days_warning", flattenVpnCertificateLocalFixedAutoRegenerateDaysWarning(o["auto-regenerate-days-warning"], d, "auto_regenerate_days_warning", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days-warning"]) {
			return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenVpnCertificateLocalFixedCaIdentifier(o["ca-identifier"], d, "ca_identifier", sv)); err != nil {
		if !fortiAPIPatch(o["ca-identifier"]) {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("name_encoding", flattenVpnCertificateLocalFixedNameEncoding(o["name-encoding"], d, "name_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["name-encoding"]) {
			return fmt.Errorf("Error reading name_encoding: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateLocalFixedSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ike_localid", flattenVpnCertificateLocalFixedIkeLocalid(o["ike-localid"], d, "ike_localid", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid"]) {
			return fmt.Errorf("Error reading ike_localid: %v", err)
		}
	}

	if err = d.Set("ike_localid_type", flattenVpnCertificateLocalFixedIkeLocalidType(o["ike-localid-type"], d, "ike_localid_type", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid-type"]) {
			return fmt.Errorf("Error reading ike_localid_type: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateLocalFixedLastUpdated(o["last-updated"], d, "last_updated", sv)); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("enroll_protocol", flattenVpnCertificateLocalFixedEnrollProtocol(o["enroll-protocol"], d, "enroll_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["enroll-protocol"]) {
			return fmt.Errorf("Error reading enroll_protocol: %v", err)
		}
	}

	if err = d.Set("cmp_server", flattenVpnCertificateLocalFixedCmpServer(o["cmp-server"], d, "cmp_server", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-server"]) {
			return fmt.Errorf("Error reading cmp_server: %v", err)
		}
	}

	if err = d.Set("cmp_path", flattenVpnCertificateLocalFixedCmpPath(o["cmp-path"], d, "cmp_path", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-path"]) {
			return fmt.Errorf("Error reading cmp_path: %v", err)
		}
	}

	if err = d.Set("cmp_server_cert", flattenVpnCertificateLocalFixedCmpServerCert(o["cmp-server-cert"], d, "cmp_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-server-cert"]) {
			return fmt.Errorf("Error reading cmp_server_cert: %v", err)
		}
	}

	if err = d.Set("cmp_regeneration_method", flattenVpnCertificateLocalFixedCmpRegenerationMethod(o["cmp-regeneration-method"], d, "cmp_regeneration_method", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-regeneration-method"]) {
			return fmt.Errorf("Error reading cmp_regeneration_method: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateLocalFixedFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnCertificateLocalFixedName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCsr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedAutoRegenerateDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedAutoRegenerateDaysWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedScepPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCaIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedNameEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedIkeLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedIkeLocalidType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedLastUpdated(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedEnrollProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCmpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCmpPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCmpServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalFixedCmpRegenerationMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateLocalFixed(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	// if v, ok := d.GetOk("name"); ok {

	// 	t, err := expandVpnCertificateLocalFixedName(d, v, "name", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["name"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("password"); ok {

	// 	t, err := expandVpnCertificateLocalFixedPassword(d, v, "password", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["password"] = t
	// 	}
	// }

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandVpnCertificateLocalFixedComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	// if v, ok := d.GetOk("private_key"); ok {

	// 	t, err := expandVpnCertificateLocalFixedPrivateKey(d, v, "private_key", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["private-key"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("certificate"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCertificate(d, v, "certificate", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["certificate"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("csr"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCsr(d, v, "csr", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["csr"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("state"); ok {

	// 	t, err := expandVpnCertificateLocalFixedState(d, v, "state", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["state"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("scep_url"); ok {

	// 	t, err := expandVpnCertificateLocalFixedScepUrl(d, v, "scep_url", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["scep-url"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("range"); ok {

	// 	t, err := expandVpnCertificateLocalFixedRange(d, v, "range", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["range"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("source"); ok {

	// 	t, err := expandVpnCertificateLocalFixedSource(d, v, "source", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["source"] = t
	// 	}
	// }

	// if v, ok := d.GetOkExists("auto_regenerate_days"); ok {

	// 	t, err := expandVpnCertificateLocalFixedAutoRegenerateDays(d, v, "auto_regenerate_days", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["auto-regenerate-days"] = t
	// 	}
	// }

	// if v, ok := d.GetOkExists("auto_regenerate_days_warning"); ok {

	// 	t, err := expandVpnCertificateLocalFixedAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["auto-regenerate-days-warning"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("scep_password"); ok {

	// 	t, err := expandVpnCertificateLocalFixedScepPassword(d, v, "scep_password", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["scep-password"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("ca_identifier"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCaIdentifier(d, v, "ca_identifier", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["ca-identifier"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("name_encoding"); ok {

	// 	t, err := expandVpnCertificateLocalFixedNameEncoding(d, v, "name_encoding", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["name-encoding"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("source_ip"); ok {

	// 	t, err := expandVpnCertificateLocalFixedSourceIp(d, v, "source_ip", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["source-ip"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("ike_localid"); ok {

	// 	t, err := expandVpnCertificateLocalFixedIkeLocalid(d, v, "ike_localid", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["ike-localid"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("ike_localid_type"); ok {

	// 	t, err := expandVpnCertificateLocalFixedIkeLocalidType(d, v, "ike_localid_type", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["ike-localid-type"] = t
	// 	}
	// }

	// if v, ok := d.GetOkExists("last_updated"); ok {

	// 	t, err := expandVpnCertificateLocalFixedLastUpdated(d, v, "last_updated", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["last-updated"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("enroll_protocol"); ok {

	// 	t, err := expandVpnCertificateLocalFixedEnrollProtocol(d, v, "enroll_protocol", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["enroll-protocol"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("cmp_server"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCmpServer(d, v, "cmp_server", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["cmp-server"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("cmp_path"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCmpPath(d, v, "cmp_path", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["cmp-path"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("cmp_server_cert"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCmpServerCert(d, v, "cmp_server_cert", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["cmp-server-cert"] = t
	// 	}
	// }

	// if v, ok := d.GetOk("cmp_regeneration_method"); ok {

	// 	t, err := expandVpnCertificateLocalFixedCmpRegenerationMethod(d, v, "cmp_regeneration_method", sv)
	// 	if err != nil {
	// 		return &obj, err
	// 	} else if t != nil {
	// 		obj["cmp-regeneration-method"] = t
	// 	}
	// }

	return &obj, nil
}
