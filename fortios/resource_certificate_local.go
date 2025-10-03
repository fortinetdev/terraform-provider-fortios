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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			},
			"csr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
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
			},
			"auto_regenerate_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"private_key_retain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmp_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"cmp_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"cmp_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
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
			},
			"acme_email": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"acme_eab_key_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"acme_eab_key_hmac": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"acme_rsa_key_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2048, 4096),
				Optional:     true,
				Computed:     true,
			},
			"acme_renew_window": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"est_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"est_ca_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"est_http_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"est_http_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"est_client_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"est_server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"est_srp_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"est_srp_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"est_regeneration_method": &schema.Schema{
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

	obj, err := getObjectCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CertificateLocal resource while getting object: %v", err)
	}

	o, err := c.CreateCertificateLocal(obj, vdomparam)

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

	obj, err := getObjectCertificateLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CertificateLocal resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateLocal(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCertificateLocal(mkey, vdomparam)
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

	o, err := c.ReadCertificateLocal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CertificateLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateLocal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CertificateLocal resource from API: %v", err)
	}
	return nil
}

func flattenCertificateLocalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalCsr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalAutoRegenerateDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCertificateLocalAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCertificateLocalCaIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalNameEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalIkeLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalIkeLocalidType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalLastUpdated(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCertificateLocalEnrollProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalPrivateKeyRetain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalCmpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalCmpPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalCmpServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalCmpRegenerationMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalAcmeCaUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalAcmeDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalAcmeEmail(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalAcmeEabKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalAcmeRsaKeySize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCertificateLocalAcmeRenewWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCertificateLocalEstServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstCaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstHttpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstHttpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstSrpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstSrpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateLocalEstRegenerationMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateLocal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenCertificateLocalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenCertificateLocalComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("certificate", flattenCertificateLocalCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("csr", flattenCertificateLocalCsr(o["csr"], d, "csr", sv)); err != nil {
		if !fortiAPIPatch(o["csr"]) {
			return fmt.Errorf("Error reading csr: %v", err)
		}
	}

	if err = d.Set("state", flattenCertificateLocalState(o["state"], d, "state", sv)); err != nil {
		if !fortiAPIPatch(o["state"]) {
			return fmt.Errorf("Error reading state: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenCertificateLocalScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateLocalRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateLocalSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days", flattenCertificateLocalAutoRegenerateDays(o["auto-regenerate-days"], d, "auto_regenerate_days", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days"]) {
			return fmt.Errorf("Error reading auto_regenerate_days: %v", err)
		}
	}

	if err = d.Set("auto_regenerate_days_warning", flattenCertificateLocalAutoRegenerateDaysWarning(o["auto-regenerate-days-warning"], d, "auto_regenerate_days_warning", sv)); err != nil {
		if !fortiAPIPatch(o["auto-regenerate-days-warning"]) {
			return fmt.Errorf("Error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenCertificateLocalCaIdentifier(o["ca-identifier"], d, "ca_identifier", sv)); err != nil {
		if !fortiAPIPatch(o["ca-identifier"]) {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("name_encoding", flattenCertificateLocalNameEncoding(o["name-encoding"], d, "name_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["name-encoding"]) {
			return fmt.Errorf("Error reading name_encoding: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenCertificateLocalSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ike_localid", flattenCertificateLocalIkeLocalid(o["ike-localid"], d, "ike_localid", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid"]) {
			return fmt.Errorf("Error reading ike_localid: %v", err)
		}
	}

	if err = d.Set("ike_localid_type", flattenCertificateLocalIkeLocalidType(o["ike-localid-type"], d, "ike_localid_type", sv)); err != nil {
		if !fortiAPIPatch(o["ike-localid-type"]) {
			return fmt.Errorf("Error reading ike_localid_type: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenCertificateLocalLastUpdated(o["last-updated"], d, "last_updated", sv)); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	if err = d.Set("enroll_protocol", flattenCertificateLocalEnrollProtocol(o["enroll-protocol"], d, "enroll_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["enroll-protocol"]) {
			return fmt.Errorf("Error reading enroll_protocol: %v", err)
		}
	}

	if err = d.Set("private_key_retain", flattenCertificateLocalPrivateKeyRetain(o["private-key-retain"], d, "private_key_retain", sv)); err != nil {
		if !fortiAPIPatch(o["private-key-retain"]) {
			return fmt.Errorf("Error reading private_key_retain: %v", err)
		}
	}

	if err = d.Set("cmp_server", flattenCertificateLocalCmpServer(o["cmp-server"], d, "cmp_server", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-server"]) {
			return fmt.Errorf("Error reading cmp_server: %v", err)
		}
	}

	if err = d.Set("cmp_path", flattenCertificateLocalCmpPath(o["cmp-path"], d, "cmp_path", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-path"]) {
			return fmt.Errorf("Error reading cmp_path: %v", err)
		}
	}

	if err = d.Set("cmp_server_cert", flattenCertificateLocalCmpServerCert(o["cmp-server-cert"], d, "cmp_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-server-cert"]) {
			return fmt.Errorf("Error reading cmp_server_cert: %v", err)
		}
	}

	if err = d.Set("cmp_regeneration_method", flattenCertificateLocalCmpRegenerationMethod(o["cmp-regeneration-method"], d, "cmp_regeneration_method", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-regeneration-method"]) {
			return fmt.Errorf("Error reading cmp_regeneration_method: %v", err)
		}
	}

	if err = d.Set("acme_ca_url", flattenCertificateLocalAcmeCaUrl(o["acme-ca-url"], d, "acme_ca_url", sv)); err != nil {
		if !fortiAPIPatch(o["acme-ca-url"]) {
			return fmt.Errorf("Error reading acme_ca_url: %v", err)
		}
	}

	if err = d.Set("acme_domain", flattenCertificateLocalAcmeDomain(o["acme-domain"], d, "acme_domain", sv)); err != nil {
		if !fortiAPIPatch(o["acme-domain"]) {
			return fmt.Errorf("Error reading acme_domain: %v", err)
		}
	}

	if err = d.Set("acme_email", flattenCertificateLocalAcmeEmail(o["acme-email"], d, "acme_email", sv)); err != nil {
		if !fortiAPIPatch(o["acme-email"]) {
			return fmt.Errorf("Error reading acme_email: %v", err)
		}
	}

	if err = d.Set("acme_eab_key_id", flattenCertificateLocalAcmeEabKeyId(o["acme-eab-key-id"], d, "acme_eab_key_id", sv)); err != nil {
		if !fortiAPIPatch(o["acme-eab-key-id"]) {
			return fmt.Errorf("Error reading acme_eab_key_id: %v", err)
		}
	}

	if err = d.Set("acme_rsa_key_size", flattenCertificateLocalAcmeRsaKeySize(o["acme-rsa-key-size"], d, "acme_rsa_key_size", sv)); err != nil {
		if !fortiAPIPatch(o["acme-rsa-key-size"]) {
			return fmt.Errorf("Error reading acme_rsa_key_size: %v", err)
		}
	}

	if err = d.Set("acme_renew_window", flattenCertificateLocalAcmeRenewWindow(o["acme-renew-window"], d, "acme_renew_window", sv)); err != nil {
		if !fortiAPIPatch(o["acme-renew-window"]) {
			return fmt.Errorf("Error reading acme_renew_window: %v", err)
		}
	}

	if err = d.Set("est_server", flattenCertificateLocalEstServer(o["est-server"], d, "est_server", sv)); err != nil {
		if !fortiAPIPatch(o["est-server"]) {
			return fmt.Errorf("Error reading est_server: %v", err)
		}
	}

	if err = d.Set("est_ca_id", flattenCertificateLocalEstCaId(o["est-ca-id"], d, "est_ca_id", sv)); err != nil {
		if !fortiAPIPatch(o["est-ca-id"]) {
			return fmt.Errorf("Error reading est_ca_id: %v", err)
		}
	}

	if err = d.Set("est_http_username", flattenCertificateLocalEstHttpUsername(o["est-http-username"], d, "est_http_username", sv)); err != nil {
		if !fortiAPIPatch(o["est-http-username"]) {
			return fmt.Errorf("Error reading est_http_username: %v", err)
		}
	}

	if err = d.Set("est_http_password", flattenCertificateLocalEstHttpPassword(o["est-http-password"], d, "est_http_password", sv)); err != nil {
		if !fortiAPIPatch(o["est-http-password"]) {
			return fmt.Errorf("Error reading est_http_password: %v", err)
		}
	}

	if err = d.Set("est_client_cert", flattenCertificateLocalEstClientCert(o["est-client-cert"], d, "est_client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["est-client-cert"]) {
			return fmt.Errorf("Error reading est_client_cert: %v", err)
		}
	}

	if err = d.Set("est_server_cert", flattenCertificateLocalEstServerCert(o["est-server-cert"], d, "est_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["est-server-cert"]) {
			return fmt.Errorf("Error reading est_server_cert: %v", err)
		}
	}

	if err = d.Set("est_srp_username", flattenCertificateLocalEstSrpUsername(o["est-srp-username"], d, "est_srp_username", sv)); err != nil {
		if !fortiAPIPatch(o["est-srp-username"]) {
			return fmt.Errorf("Error reading est_srp_username: %v", err)
		}
	}

	if err = d.Set("est_srp_password", flattenCertificateLocalEstSrpPassword(o["est-srp-password"], d, "est_srp_password", sv)); err != nil {
		if !fortiAPIPatch(o["est-srp-password"]) {
			return fmt.Errorf("Error reading est_srp_password: %v", err)
		}
	}

	if err = d.Set("est_regeneration_method", flattenCertificateLocalEstRegenerationMethod(o["est-regeneration-method"], d, "est_regeneration_method", sv)); err != nil {
		if !fortiAPIPatch(o["est-regeneration-method"]) {
			return fmt.Errorf("Error reading est_regeneration_method: %v", err)
		}
	}

	return nil
}

func flattenCertificateLocalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCertificateLocalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"7.4.4"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); !versionMatch {
		return toCertFormat(v), nil
	} else {
		return v, nil
	}
}

func expandCertificateLocalCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"7.4.4"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); !versionMatch {
		return toCertFormat(v), nil
	} else {
		return v, nil
	}
}

func expandCertificateLocalCsr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAutoRegenerateDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAutoRegenerateDaysWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalScepPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCaIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalNameEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalIkeLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalIkeLocalidType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalLastUpdated(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEnrollProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalPrivateKeyRetain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalCmpRegenerationMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeCaUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeEmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeEabKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeEabKeyHmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeRsaKeySize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalAcmeRenewWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstCaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstHttpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstHttpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstSrpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstSrpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateLocalEstRegenerationMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCertificateLocalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandCertificateLocalPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandCertificateLocalComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandCertificateLocalPrivateKey(d, v, "private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	} else if d.HasChange("private_key") {
		obj["private-key"] = nil
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandCertificateLocalCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("csr"); ok {
		t, err := expandCertificateLocalCsr(d, v, "csr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr"] = t
		}
	} else if d.HasChange("csr") {
		obj["csr"] = nil
	}

	if v, ok := d.GetOk("state"); ok {
		t, err := expandCertificateLocalState(d, v, "state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	} else if d.HasChange("state") {
		obj["state"] = nil
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandCertificateLocalScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	} else if d.HasChange("scep_url") {
		obj["scep-url"] = nil
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandCertificateLocalRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandCertificateLocalSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_regenerate_days"); ok {
		t, err := expandCertificateLocalAutoRegenerateDays(d, v, "auto_regenerate_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days"] = t
		}
	} else if d.HasChange("auto_regenerate_days") {
		obj["auto-regenerate-days"] = nil
	}

	if v, ok := d.GetOkExists("auto_regenerate_days_warning"); ok {
		t, err := expandCertificateLocalAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days-warning"] = t
		}
	} else if d.HasChange("auto_regenerate_days_warning") {
		obj["auto-regenerate-days-warning"] = nil
	}

	if v, ok := d.GetOk("scep_password"); ok {
		t, err := expandCertificateLocalScepPassword(d, v, "scep_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	} else if d.HasChange("scep_password") {
		obj["scep-password"] = nil
	}

	if v, ok := d.GetOk("ca_identifier"); ok {
		t, err := expandCertificateLocalCaIdentifier(d, v, "ca_identifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	} else if d.HasChange("ca_identifier") {
		obj["ca-identifier"] = nil
	}

	if v, ok := d.GetOk("name_encoding"); ok {
		t, err := expandCertificateLocalNameEncoding(d, v, "name_encoding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name-encoding"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandCertificateLocalSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ike_localid"); ok {
		t, err := expandCertificateLocalIkeLocalid(d, v, "ike_localid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid"] = t
		}
	} else if d.HasChange("ike_localid") {
		obj["ike-localid"] = nil
	}

	if v, ok := d.GetOk("ike_localid_type"); ok {
		t, err := expandCertificateLocalIkeLocalidType(d, v, "ike_localid_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-localid-type"] = t
		}
	}

	if v, ok := d.GetOkExists("last_updated"); ok {
		t, err := expandCertificateLocalLastUpdated(d, v, "last_updated", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	if v, ok := d.GetOk("enroll_protocol"); ok {
		t, err := expandCertificateLocalEnrollProtocol(d, v, "enroll_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enroll-protocol"] = t
		}
	}

	if v, ok := d.GetOk("private_key_retain"); ok {
		t, err := expandCertificateLocalPrivateKeyRetain(d, v, "private_key_retain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key-retain"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server"); ok {
		t, err := expandCertificateLocalCmpServer(d, v, "cmp_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server"] = t
		}
	} else if d.HasChange("cmp_server") {
		obj["cmp-server"] = nil
	}

	if v, ok := d.GetOk("cmp_path"); ok {
		t, err := expandCertificateLocalCmpPath(d, v, "cmp_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-path"] = t
		}
	} else if d.HasChange("cmp_path") {
		obj["cmp-path"] = nil
	}

	if v, ok := d.GetOk("cmp_server_cert"); ok {
		t, err := expandCertificateLocalCmpServerCert(d, v, "cmp_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server-cert"] = t
		}
	} else if d.HasChange("cmp_server_cert") {
		obj["cmp-server-cert"] = nil
	}

	if v, ok := d.GetOk("cmp_regeneration_method"); ok {
		t, err := expandCertificateLocalCmpRegenerationMethod(d, v, "cmp_regeneration_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-regeneration-method"] = t
		}
	}

	if v, ok := d.GetOk("acme_ca_url"); ok {
		t, err := expandCertificateLocalAcmeCaUrl(d, v, "acme_ca_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-ca-url"] = t
		}
	}

	if v, ok := d.GetOk("acme_domain"); ok {
		t, err := expandCertificateLocalAcmeDomain(d, v, "acme_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-domain"] = t
		}
	} else if d.HasChange("acme_domain") {
		obj["acme-domain"] = nil
	}

	if v, ok := d.GetOk("acme_email"); ok {
		t, err := expandCertificateLocalAcmeEmail(d, v, "acme_email", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-email"] = t
		}
	} else if d.HasChange("acme_email") {
		obj["acme-email"] = nil
	}

	if v, ok := d.GetOk("acme_eab_key_id"); ok {
		t, err := expandCertificateLocalAcmeEabKeyId(d, v, "acme_eab_key_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-eab-key-id"] = t
		}
	} else if d.HasChange("acme_eab_key_id") {
		obj["acme-eab-key-id"] = nil
	}

	if v, ok := d.GetOk("acme_eab_key_hmac"); ok {
		t, err := expandCertificateLocalAcmeEabKeyHmac(d, v, "acme_eab_key_hmac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-eab-key-hmac"] = t
		}
	} else if d.HasChange("acme_eab_key_hmac") {
		obj["acme-eab-key-hmac"] = nil
	}

	if v, ok := d.GetOk("acme_rsa_key_size"); ok {
		t, err := expandCertificateLocalAcmeRsaKeySize(d, v, "acme_rsa_key_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-rsa-key-size"] = t
		}
	}

	if v, ok := d.GetOk("acme_renew_window"); ok {
		t, err := expandCertificateLocalAcmeRenewWindow(d, v, "acme_renew_window", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-renew-window"] = t
		}
	}

	if v, ok := d.GetOk("est_server"); ok {
		t, err := expandCertificateLocalEstServer(d, v, "est_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server"] = t
		}
	} else if d.HasChange("est_server") {
		obj["est-server"] = nil
	}

	if v, ok := d.GetOk("est_ca_id"); ok {
		t, err := expandCertificateLocalEstCaId(d, v, "est_ca_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-ca-id"] = t
		}
	} else if d.HasChange("est_ca_id") {
		obj["est-ca-id"] = nil
	}

	if v, ok := d.GetOk("est_http_username"); ok {
		t, err := expandCertificateLocalEstHttpUsername(d, v, "est_http_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-username"] = t
		}
	} else if d.HasChange("est_http_username") {
		obj["est-http-username"] = nil
	}

	if v, ok := d.GetOk("est_http_password"); ok {
		t, err := expandCertificateLocalEstHttpPassword(d, v, "est_http_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-password"] = t
		}
	} else if d.HasChange("est_http_password") {
		obj["est-http-password"] = nil
	}

	if v, ok := d.GetOk("est_client_cert"); ok {
		t, err := expandCertificateLocalEstClientCert(d, v, "est_client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-client-cert"] = t
		}
	} else if d.HasChange("est_client_cert") {
		obj["est-client-cert"] = nil
	}

	if v, ok := d.GetOk("est_server_cert"); ok {
		t, err := expandCertificateLocalEstServerCert(d, v, "est_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server-cert"] = t
		}
	} else if d.HasChange("est_server_cert") {
		obj["est-server-cert"] = nil
	}

	if v, ok := d.GetOk("est_srp_username"); ok {
		t, err := expandCertificateLocalEstSrpUsername(d, v, "est_srp_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-username"] = t
		}
	} else if d.HasChange("est_srp_username") {
		obj["est-srp-username"] = nil
	}

	if v, ok := d.GetOk("est_srp_password"); ok {
		t, err := expandCertificateLocalEstSrpPassword(d, v, "est_srp_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-password"] = t
		}
	} else if d.HasChange("est_srp_password") {
		obj["est-srp-password"] = nil
	}

	if v, ok := d.GetOk("est_regeneration_method"); ok {
		t, err := expandCertificateLocalEstRegenerationMethod(d, v, "est_regeneration_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-regeneration-method"] = t
		}
	}

	return &obj, nil
}
