// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Local keys and certificates.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateLocalCreate,
		Read:   resourceCertificateLocalRead,
		Update: resourceCertificateLocalUpdate,
		Delete: resourceCertificateLocalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
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
				Required:  true,
				Sensitive: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Optional: true,
				Computed: true,
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
		},
	}
}

func resourceCertificateLocalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectCertificateLocal(d)
	if err != nil {
		return fmt.Errorf("Error creating CertificateLocal resource while getting object: %v", err)
	}

	o, err := c.CreateCertificateLocal(obj)

	if err != nil {
		return fmt.Errorf("Error creating CertificateLocal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateLocal")
	}

	return resourceCertificateLocalRead(d, m)
}

func resourceCertificateLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectCertificateLocal(d)
	if err != nil {
		return fmt.Errorf("Error updating CertificateLocal resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateLocal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating CertificateLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateLocal")
	}

	return resourceCertificateLocalRead(d, m)
}

func resourceCertificateLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteCertificateLocal(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadCertificateLocal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading CertificateLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateLocal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CertificateLocal resource from API: %v", err)
	}
	return nil
}

func flattenCertificateLocalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalAutoRegenerateDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalScepPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCaIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalNameEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalIkeLocalid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalIkeLocalidType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalEnrollProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCmpServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCmpPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCmpServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateLocalCmpRegenerationMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCertificateLocal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenCertificateLocalName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenCertificateLocalComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("certificate", flattenCertificateLocalCertificate(o["certificate"], d, "certificate")); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("csr", flattenCertificateLocalCsr(o["csr"], d, "csr")); err != nil {
		if !fortiAPIPatch(o["csr"]) {
			return fmt.Errorf("Error reading csr: %v", err)
		}
	}

	if err = d.Set("state", flattenCertificateLocalState(o["state"], d, "state")); err != nil {
		if !fortiAPIPatch(o["state"]) {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenCertificateLocalScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateLocalRange(o["range"], d, "range")); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateLocalSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days", flattenCertificateLocalAutoRegenerateDays(o["auto-regenerate-days"], d, "auto_regenerate_days")); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days"]) {
			return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days_warning", flattenCertificateLocalAutoRegenerateDaysWarning(o["auto-regenerate-days-warning"], d, "auto_regenerate_days_warning")); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days-warning"]) {
			return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenCertificateLocalCaIdentifier(o["ca-identifier"], d, "ca_identifier")); err != nil {
		if !fortiAPIPatch(o["ca-identifier"]) {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("name_encoding", flattenCertificateLocalNameEncoding(o["name-encoding"], d, "name_encoding")); err != nil {
		if !fortiAPIPatch(o["name-encoding"]) {
			return fmt.Errorf("Error reading name_encoding: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenCertificateLocalSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ike_localid", flattenCertificateLocalIkeLocalid(o["ike-localid"], d, "ike_localid")); err != nil {
		if !fortiAPIPatch(o["ike-localid"]) {
			return fmt.Errorf("Error reading ike_localid: %v", err)
		}
	}

	if err = d.Set("ike_localid_type", flattenCertificateLocalIkeLocalidType(o["ike-localid-type"], d, "ike_localid_type")); err != nil {
		if !fortiAPIPatch(o["ike-localid-type"]) {
			return fmt.Errorf("Error reading ike_localid_type: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenCertificateLocalLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("enroll_protocol", flattenCertificateLocalEnrollProtocol(o["enroll-protocol"], d, "enroll_protocol")); err != nil {
		if !fortiAPIPatch(o["enroll-protocol"]) {
			return fmt.Errorf("Error reading enroll_protocol: %v", err)
		}
	}

	if err = d.Set("cmp_server", flattenCertificateLocalCmpServer(o["cmp-server"], d, "cmp_server")); err != nil {
		if !fortiAPIPatch(o["cmp-server"]) {
			return fmt.Errorf("Error reading cmp_server: %v", err)
		}
	}

	if err = d.Set("cmp_path", flattenCertificateLocalCmpPath(o["cmp-path"], d, "cmp_path")); err != nil {
		if !fortiAPIPatch(o["cmp-path"]) {
			return fmt.Errorf("Error reading cmp_path: %v", err)
		}
	}

	if err = d.Set("cmp_server_cert", flattenCertificateLocalCmpServerCert(o["cmp-server-cert"], d, "cmp_server_cert")); err != nil {
		if !fortiAPIPatch(o["cmp-server-cert"]) {
			return fmt.Errorf("Error reading cmp_server_cert: %v", err)
		}
	}

	if err = d.Set("cmp_regeneration_method", flattenCertificateLocalCmpRegenerationMethod(o["cmp-regeneration-method"], d, "cmp_regeneration_method")); err != nil {
		if !fortiAPIPatch(o["cmp-regeneration-method"]) {
			return fmt.Errorf("Error reading cmp_regeneration_method: %v", err)
		}
	}

	return nil
}

func flattenCertificateLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCertificateLocalName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalPrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAutoRegenerateDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAutoRegenerateDaysWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalScepPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCaIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalNameEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalIkeLocalid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalIkeLocalidType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEnrollProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpRegenerationMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateLocal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCertificateLocalName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandCertificateLocalPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandCertificateLocalComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandCertificateLocalPrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandCertificateLocalCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("csr"); ok {
		t, err := expandCertificateLocalCsr(d, v, "csr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok {
		t, err := expandCertificateLocalState(d, v, "state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandCertificateLocalScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandCertificateLocalRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandCertificateLocalSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("auto_regenerate_days"); ok {
		t, err := expandCertificateLocalAutoRegenerateDays(d, v, "auto_regenerate_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days"] = t
		}
	}

	if v, ok := d.GetOk("auto_regenerate_days_warning"); ok {
		t, err := expandCertificateLocalAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("scep_password"); ok {
		t, err := expandCertificateLocalScepPassword(d, v, "scep_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	}

	if v, ok := d.GetOk("ca_identifier"); ok {
		t, err := expandCertificateLocalCaIdentifier(d, v, "ca_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	}

	if v, ok := d.GetOk("name_encoding"); ok {
		t, err := expandCertificateLocalNameEncoding(d, v, "name_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name-encoding"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandCertificateLocalSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid"); ok {
		t, err := expandCertificateLocalIkeLocalid(d, v, "ike_localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid_type"); ok {
		t, err := expandCertificateLocalIkeLocalidType(d, v, "ike_localid_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid-type"] = t
		}
	}

	if v, ok := d.GetOk("last_updated"); ok {
		t, err := expandCertificateLocalLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("enroll_protocol"); ok {
		t, err := expandCertificateLocalEnrollProtocol(d, v, "enroll_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enroll-protocol"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server"); ok {
		t, err := expandCertificateLocalCmpServer(d, v, "cmp_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server"] = t
		}
	}

	if v, ok := d.GetOk("cmp_path"); ok {
		t, err := expandCertificateLocalCmpPath(d, v, "cmp_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-path"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server_cert"); ok {
		t, err := expandCertificateLocalCmpServerCert(d, v, "cmp_server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("cmp_regeneration_method"); ok {
		t, err := expandCertificateLocalCmpRegenerationMethod(d, v, "cmp_regeneration_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-regeneration-method"] = t
		}
	}

	return &obj, nil
}
