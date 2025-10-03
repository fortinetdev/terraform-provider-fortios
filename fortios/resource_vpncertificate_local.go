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
				Optional:  true,
				Sensitive: true,
			},
			"certificate": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
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

func resourceVpnCertificateLocalCreate(d *schema.ResourceData, m interface{}) error {
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

func flattenVpnCertificateLocalComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
	return convintf2i(v)
}

func flattenVpnCertificateLocalAutoRegenerateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
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
	return convintf2i(v)
}

func flattenVpnCertificateLocalEnrollProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalPrivateKeyRetain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

func flattenVpnCertificateLocalAcmeEabKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalAcmeRsaKeySize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnCertificateLocalAcmeRenewWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnCertificateLocalEstServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstCaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstHttpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstHttpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstSrpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstSrpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateLocalEstRegenerationMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

	if err = d.Set("private_key_retain", flattenVpnCertificateLocalPrivateKeyRetain(o["private-key-retain"], d, "private_key_retain", sv)); err != nil {
		if !fortiAPIPatch(o["private-key-retain"]) {
			return fmt.Errorf("Error reading private_key_retain: %v", err)
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

	if err = d.Set("acme_eab_key_id", flattenVpnCertificateLocalAcmeEabKeyId(o["acme-eab-key-id"], d, "acme_eab_key_id", sv)); err != nil {
		if !fortiAPIPatch(o["acme-eab-key-id"]) {
			return fmt.Errorf("Error reading acme_eab_key_id: %v", err)
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

	if err = d.Set("est_server", flattenVpnCertificateLocalEstServer(o["est-server"], d, "est_server", sv)); err != nil {
		if !fortiAPIPatch(o["est-server"]) {
			return fmt.Errorf("Error reading est_server: %v", err)
		}
	}

	if err = d.Set("est_ca_id", flattenVpnCertificateLocalEstCaId(o["est-ca-id"], d, "est_ca_id", sv)); err != nil {
		if !fortiAPIPatch(o["est-ca-id"]) {
			return fmt.Errorf("Error reading est_ca_id: %v", err)
		}
	}

	if err = d.Set("est_http_username", flattenVpnCertificateLocalEstHttpUsername(o["est-http-username"], d, "est_http_username", sv)); err != nil {
		if !fortiAPIPatch(o["est-http-username"]) {
			return fmt.Errorf("Error reading est_http_username: %v", err)
		}
	}

	if err = d.Set("est_http_password", flattenVpnCertificateLocalEstHttpPassword(o["est-http-password"], d, "est_http_password", sv)); err != nil {
		if !fortiAPIPatch(o["est-http-password"]) {
			return fmt.Errorf("Error reading est_http_password: %v", err)
		}
	}

	if err = d.Set("est_client_cert", flattenVpnCertificateLocalEstClientCert(o["est-client-cert"], d, "est_client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["est-client-cert"]) {
			return fmt.Errorf("Error reading est_client_cert: %v", err)
		}
	}

	if err = d.Set("est_server_cert", flattenVpnCertificateLocalEstServerCert(o["est-server-cert"], d, "est_server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["est-server-cert"]) {
			return fmt.Errorf("Error reading est_server_cert: %v", err)
		}
	}

	if err = d.Set("est_srp_username", flattenVpnCertificateLocalEstSrpUsername(o["est-srp-username"], d, "est_srp_username", sv)); err != nil {
		if !fortiAPIPatch(o["est-srp-username"]) {
			return fmt.Errorf("Error reading est_srp_username: %v", err)
		}
	}

	if err = d.Set("est_srp_password", flattenVpnCertificateLocalEstSrpPassword(o["est-srp-password"], d, "est_srp_password", sv)); err != nil {
		if !fortiAPIPatch(o["est-srp-password"]) {
			return fmt.Errorf("Error reading est_srp_password: %v", err)
		}
	}

	if err = d.Set("est_regeneration_method", flattenVpnCertificateLocalEstRegenerationMethod(o["est-regeneration-method"], d, "est_regeneration_method", sv)); err != nil {
		if !fortiAPIPatch(o["est-regeneration-method"]) {
			return fmt.Errorf("Error reading est_regeneration_method: %v", err)
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
	new_version_map := map[string][]string{
		">=": []string{"7.4.4"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); !versionMatch {
		return toCertFormat(v), nil
	} else {
		return v, nil
	}
}

func expandVpnCertificateLocalCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"7.4.4"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); !versionMatch {
		return toCertFormat(v), nil
	} else {
		return v, nil
	}
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

func expandVpnCertificateLocalPrivateKeyRetain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandVpnCertificateLocalAcmeEabKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeEabKeyHmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeRsaKeySize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalAcmeRenewWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstCaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstHttpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstHttpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstSrpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstSrpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateLocalEstRegenerationMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
	} else if d.HasChange("password") {
		obj["password"] = nil
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
	} else if d.HasChange("private_key") {
		obj["private-key"] = nil
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandVpnCertificateLocalCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("csr"); ok {
		t, err := expandVpnCertificateLocalCsr(d, v, "csr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["csr"] = t
		}
	} else if d.HasChange("csr") {
		obj["csr"] = nil
	}

	if v, ok := d.GetOk("state"); ok {
		t, err := expandVpnCertificateLocalState(d, v, "state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["state"] = t
		}
	} else if d.HasChange("state") {
		obj["state"] = nil
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandVpnCertificateLocalScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	} else if d.HasChange("scep_url") {
		obj["scep-url"] = nil
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
	} else if d.HasChange("auto_regenerate_days") {
		obj["auto-regenerate-days"] = nil
	}

	if v, ok := d.GetOkExists("auto_regenerate_days_warning"); ok {
		t, err := expandVpnCertificateLocalAutoRegenerateDaysWarning(d, v, "auto_regenerate_days_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-regenerate-days-warning"] = t
		}
	} else if d.HasChange("auto_regenerate_days_warning") {
		obj["auto-regenerate-days-warning"] = nil
	}

	if v, ok := d.GetOk("scep_password"); ok {
		t, err := expandVpnCertificateLocalScepPassword(d, v, "scep_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-password"] = t
		}
	} else if d.HasChange("scep_password") {
		obj["scep-password"] = nil
	}

	if v, ok := d.GetOk("ca_identifier"); ok {
		t, err := expandVpnCertificateLocalCaIdentifier(d, v, "ca_identifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	} else if d.HasChange("ca_identifier") {
		obj["ca-identifier"] = nil
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
	} else if d.HasChange("ike_localid") {
		obj["ike-localid"] = nil
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

	if v, ok := d.GetOk("private_key_retain"); ok {
		t, err := expandVpnCertificateLocalPrivateKeyRetain(d, v, "private_key_retain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key-retain"] = t
		}
	}

	if v, ok := d.GetOk("cmp_server"); ok {
		t, err := expandVpnCertificateLocalCmpServer(d, v, "cmp_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server"] = t
		}
	} else if d.HasChange("cmp_server") {
		obj["cmp-server"] = nil
	}

	if v, ok := d.GetOk("cmp_path"); ok {
		t, err := expandVpnCertificateLocalCmpPath(d, v, "cmp_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-path"] = t
		}
	} else if d.HasChange("cmp_path") {
		obj["cmp-path"] = nil
	}

	if v, ok := d.GetOk("cmp_server_cert"); ok {
		t, err := expandVpnCertificateLocalCmpServerCert(d, v, "cmp_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-server-cert"] = t
		}
	} else if d.HasChange("cmp_server_cert") {
		obj["cmp-server-cert"] = nil
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
	} else if d.HasChange("acme_domain") {
		obj["acme-domain"] = nil
	}

	if v, ok := d.GetOk("acme_email"); ok {
		t, err := expandVpnCertificateLocalAcmeEmail(d, v, "acme_email", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-email"] = t
		}
	} else if d.HasChange("acme_email") {
		obj["acme-email"] = nil
	}

	if v, ok := d.GetOk("acme_eab_key_id"); ok {
		t, err := expandVpnCertificateLocalAcmeEabKeyId(d, v, "acme_eab_key_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-eab-key-id"] = t
		}
	} else if d.HasChange("acme_eab_key_id") {
		obj["acme-eab-key-id"] = nil
	}

	if v, ok := d.GetOk("acme_eab_key_hmac"); ok {
		t, err := expandVpnCertificateLocalAcmeEabKeyHmac(d, v, "acme_eab_key_hmac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acme-eab-key-hmac"] = t
		}
	} else if d.HasChange("acme_eab_key_hmac") {
		obj["acme-eab-key-hmac"] = nil
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

	if v, ok := d.GetOk("est_server"); ok {
		t, err := expandVpnCertificateLocalEstServer(d, v, "est_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server"] = t
		}
	} else if d.HasChange("est_server") {
		obj["est-server"] = nil
	}

	if v, ok := d.GetOk("est_ca_id"); ok {
		t, err := expandVpnCertificateLocalEstCaId(d, v, "est_ca_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-ca-id"] = t
		}
	} else if d.HasChange("est_ca_id") {
		obj["est-ca-id"] = nil
	}

	if v, ok := d.GetOk("est_http_username"); ok {
		t, err := expandVpnCertificateLocalEstHttpUsername(d, v, "est_http_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-username"] = t
		}
	} else if d.HasChange("est_http_username") {
		obj["est-http-username"] = nil
	}

	if v, ok := d.GetOk("est_http_password"); ok {
		t, err := expandVpnCertificateLocalEstHttpPassword(d, v, "est_http_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-http-password"] = t
		}
	} else if d.HasChange("est_http_password") {
		obj["est-http-password"] = nil
	}

	if v, ok := d.GetOk("est_client_cert"); ok {
		t, err := expandVpnCertificateLocalEstClientCert(d, v, "est_client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-client-cert"] = t
		}
	} else if d.HasChange("est_client_cert") {
		obj["est-client-cert"] = nil
	}

	if v, ok := d.GetOk("est_server_cert"); ok {
		t, err := expandVpnCertificateLocalEstServerCert(d, v, "est_server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-server-cert"] = t
		}
	} else if d.HasChange("est_server_cert") {
		obj["est-server-cert"] = nil
	}

	if v, ok := d.GetOk("est_srp_username"); ok {
		t, err := expandVpnCertificateLocalEstSrpUsername(d, v, "est_srp_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-username"] = t
		}
	} else if d.HasChange("est_srp_username") {
		obj["est-srp-username"] = nil
	}

	if v, ok := d.GetOk("est_srp_password"); ok {
		t, err := expandVpnCertificateLocalEstSrpPassword(d, v, "est_srp_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-srp-password"] = t
		}
	} else if d.HasChange("est_srp_password") {
		obj["est-srp-password"] = nil
	}

	if v, ok := d.GetOk("est_regeneration_method"); ok {
		t, err := expandVpnCertificateLocalEstRegenerationMethod(d, v, "est_regeneration_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-regeneration-method"] = t
		}
	}

	return &obj, nil
}
