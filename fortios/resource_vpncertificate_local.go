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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateLocalCreate,
		Read:   resourceVpnCertificateLocalRead,
		Update: resourceVpnCertificateLocalUpdate,
		Delete: resourceVpnCertificateLocalDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"cmp_regeneration_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acme_ca_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"acme_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"acme_email": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"acme_rsa_key_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2048, 4096),
				Optional:     true,
				Computed:     true,
			},
			"acme_renew_window": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceVpnCertificateLocalCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocal resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateLocal(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateLocal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateLocal")
	}

	return resourceVpnCertificateLocalRead(d, m)
}

func resourceVpnCertificateLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateLocal resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateLocal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateLocal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateLocal")
	}

	return resourceVpnCertificateLocalRead(d, m)
}

func resourceVpnCertificateLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnCertificateLocal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnCertificateLocal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateLocal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateLocal resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateLocalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalPrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCsr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAutoRegenerateDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalScepPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCaIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalNameEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalIkeLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalIkeLocalidType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalLastUpdated(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEnrollProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalCmpRegenerationMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeCaUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeEmail(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeRsaKeySize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeRenewWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnCertificateLocal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateLocalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnCertificateLocalComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("certificate", flattenVpnCertificateLocalCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("csr", flattenVpnCertificateLocalCsr(o["csr"], d, "csr", sv)); err != nil {
		if !fortiAPIPatch(o["csr"]) {
			return fmt.Errorf("Error reading csr: %v", err)
		}
	}

	if err = d.Set("state", flattenVpnCertificateLocalState(o["state"], d, "state", sv)); err != nil {
		if !fortiAPIPatch(o["state"]) {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenVpnCertificateLocalScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateLocalRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateLocalSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days", flattenVpnCertificateLocalAutoRegenerateDays(o["auto-regenerate-days"], d, "auto_regenerate_days", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days"]) {
			return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days_warning", flattenVpnCertificateLocalAutoRegenerateDaysWarning(o["auto-regenerate-days-warning"], d, "auto_regenerate_days_warning", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days-warning"]) {
			return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenVpnCertificateLocalCaIdentifier(o["ca-identifier"], d, "ca_identifier", sv)); err != nil {
		if !fortiAPIPatch(o["ca-identifier"]) {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("name_encoding", flattenVpnCertificateLocalNameEncoding(o["name-encoding"], d, "name_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["name-encoding"]) {
			return fmt.Errorf("Error reading name_encoding: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateLocalSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ike_localid", flattenVpnCertificateLocalIkeLocalid(o["ike-localid"], d, "ike_localid", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid"]) {
			return fmt.Errorf("Error reading ike_localid: %v", err)
		}
	}

	if err = d.Set("ike_localid_type", flattenVpnCertificateLocalIkeLocalidType(o["ike-localid-type"], d, "ike_localid_type", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid-type"]) {
			return fmt.Errorf("Error reading ike_localid_type: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenVpnCertificateLocalLastUpdated(o["last-updated"], d, "last_updated", sv)); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("enroll_protocol", flattenVpnCertificateLocalEnrollProtocol(o["enroll-protocol"], d, "enroll_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["enroll-protocol"]) {
			return fmt.Errorf("Error reading enroll_protocol: %v", err)
		}
	}

	if err = d.Set("cmp_server", flattenVpnCertificateLocalCmpServer(o["cmp-server"], d, "cmp_server", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-server"]) {
			return fmt.Errorf("Error reading cmp_server: %v", err)
		}
	}

	if err = d.Set("cmp_path", flattenVpnCertificateLocalCmpPath(o["cmp-path"], d, "cmp_path", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-path"]) {
			return fmt.Errorf("Error reading cmp_path: %v", err)
		}
	}

	if err = d.Set("cmp_server_cert", flattenVpnCertificateLocalCmpServerCert(o["cmp-server-cert"], d, "cmp_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-server-cert"]) {
			return fmt.Errorf("Error reading cmp_server_cert: %v", err)
		}
	}

	if err = d.Set("cmp_regeneration_method", flattenVpnCertificateLocalCmpRegenerationMethod(o["cmp-regeneration-method"], d, "cmp_regeneration_method", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-regeneration-method"]) {
			return fmt.Errorf("Error reading cmp_regeneration_method: %v", err)
		}
	}

	if err = d.Set("acme_ca_url", flattenVpnCertificateLocalAcmeCaUrl(o["acme-ca-url"], d, "acme_ca_url", sv)); err != nil {
		if !fortiAPIPatch(o["acme-ca-url"]) {
			return fmt.Errorf("Error reading acme_ca_url: %v", err)
		}
	}

	if err = d.Set("acme_domain", flattenVpnCertificateLocalAcmeDomain(o["acme-domain"], d, "acme_domain", sv)); err != nil {
		if !fortiAPIPatch(o["acme-domain"]) {
			return fmt.Errorf("Error reading acme_domain: %v", err)
		}
	}

	if err = d.Set("acme_email", flattenVpnCertificateLocalAcmeEmail(o["acme-email"], d, "acme_email", sv)); err != nil {
		if !fortiAPIPatch(o["acme-email"]) {
			return fmt.Errorf("Error reading acme_email: %v", err)
		}
	}

	if err = d.Set("acme_rsa_key_size", flattenVpnCertificateLocalAcmeRsaKeySize(o["acme-rsa-key-size"], d, "acme_rsa_key_size", sv)); err != nil {
		if !fortiAPIPatch(o["acme-rsa-key-size"]) {
			return fmt.Errorf("Error reading acme_rsa_key_size: %v", err)
		}
	}

	if err = d.Set("acme_renew_window", flattenVpnCertificateLocalAcmeRenewWindow(o["acme-renew-window"], d, "acme_renew_window", sv)); err != nil {
		if !fortiAPIPatch(o["acme-renew-window"]) {
			return fmt.Errorf("Error reading acme_renew_window: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnCertificateLocalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCsr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAutoRegenerateDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAutoRegenerateDaysWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalScepPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCaIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalNameEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalIkeLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalIkeLocalidType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalLastUpdated(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEnrollProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalCmpRegenerationMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeCaUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeEmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeRsaKeySize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeRenewWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnCertificateLocalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandVpnCertificateLocalPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandVpnCertificateLocalComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {

		t, err := expandVpnCertificateLocalPrivateKey(d, v, "private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {

		t, err := expandVpnCertificateLocalCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("csr"); ok {

		t, err := expandVpnCertificateLocalCsr(d, v, "csr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr"] = t
		}
	}

	if v, ok := d.GetOk("state"); ok {

		t, err := expandVpnCertificateLocalState(d, v, "state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {

		t, err := expandVpnCertificateLocalScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {

		t, err := expandVpnCertificateLocalRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {

		t, err := expandVpnCertificateLocalSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_regenerate_days"); ok {

		t, err := expandVpnCertificateLocalAutoRegenerateDays(d, v, "auto_regenerate_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_regenerate_days_warning"); ok {

		t, err := expandVpnCertificateLocalAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("scep_password"); ok {

		t, err := expandVpnCertificateLocalScepPassword(d, v, "scep_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	}

	if v, ok := d.GetOk("ca_identifier"); ok {

		t, err := expandVpnCertificateLocalCaIdentifier(d, v, "ca_identifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	}

	if v, ok := d.GetOk("name_encoding"); ok {

		t, err := expandVpnCertificateLocalNameEncoding(d, v, "name_encoding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name-encoding"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandVpnCertificateLocalSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid"); ok {

		t, err := expandVpnCertificateLocalIkeLocalid(d, v, "ike_localid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid_type"); ok {

		t, err := expandVpnCertificateLocalIkeLocalidType(d, v, "ike_localid_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid-type"] = t
		}
	}

	if v, ok := d.GetOkExists("last_updated"); ok {

		t, err := expandVpnCertificateLocalLastUpdated(d, v, "last_updated", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("enroll_protocol"); ok {

		t, err := expandVpnCertificateLocalEnrollProtocol(d, v, "enroll_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enroll-protocol"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server"); ok {

		t, err := expandVpnCertificateLocalCmpServer(d, v, "cmp_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server"] = t
		}
	}

	if v, ok := d.GetOk("cmp_path"); ok {

		t, err := expandVpnCertificateLocalCmpPath(d, v, "cmp_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-path"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server_cert"); ok {

		t, err := expandVpnCertificateLocalCmpServerCert(d, v, "cmp_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server-cert"] = t
		}
	}

	if v, ok := d.GetOk("cmp_regeneration_method"); ok {

		t, err := expandVpnCertificateLocalCmpRegenerationMethod(d, v, "cmp_regeneration_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-regeneration-method"] = t
		}
	}

	if v, ok := d.GetOk("acme_ca_url"); ok {

		t, err := expandVpnCertificateLocalAcmeCaUrl(d, v, "acme_ca_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-ca-url"] = t
		}
	}

	if v, ok := d.GetOk("acme_domain"); ok {

		t, err := expandVpnCertificateLocalAcmeDomain(d, v, "acme_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-domain"] = t
		}
	}

	if v, ok := d.GetOk("acme_email"); ok {

		t, err := expandVpnCertificateLocalAcmeEmail(d, v, "acme_email", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-email"] = t
		}
	}

	if v, ok := d.GetOk("acme_rsa_key_size"); ok {

		t, err := expandVpnCertificateLocalAcmeRsaKeySize(d, v, "acme_rsa_key_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-rsa-key-size"] = t
		}
	}

	if v, ok := d.GetOk("acme_renew_window"); ok {

		t, err := expandVpnCertificateLocalAcmeRenewWindow(d, v, "acme_renew_window", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-renew-window"] = t
		}
	}

	return &obj, nil
}
