// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: VPN certificate setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateSettingUpdate,
		Read:   resourceVpnCertificateSettingRead,
		Update: resourceVpnCertificateSettingUpdate,
		Delete: resourceVpnCertificateSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ocsp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_default_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"check_ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_ca_chain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subject_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cn_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_crl_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_ocsp_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cmp_save_extra_certs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certname_rsa1024": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"certname_rsa2048": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"certname_dsa1024": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"certname_dsa2048": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"certname_ecdsa256": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"certname_ecdsa384": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
		},
	}
}

func resourceVpnCertificateSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnCertificateSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateSetting")
	}

	return resourceVpnCertificateSettingRead(d, m)
}

func resourceVpnCertificateSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnCertificateSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnCertificateSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSetting resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateSettingOcspStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingOcspOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingOcspDefaultServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCheckCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCheckCaChain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSubjectMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCnMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingStrictCrlCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingStrictOcspCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCmpSaveExtraCerts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameRsa1024(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameRsa2048(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameDsa1024(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameDsa2048(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEcdsa256(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEcdsa384(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ocsp_status", flattenVpnCertificateSettingOcspStatus(o["ocsp-status"], d, "ocsp_status")); err != nil {
		if !fortiAPIPatch(o["ocsp-status"]) {
			return fmt.Errorf("Error reading ocsp_status: %v", err)
		}
	}

	if err = d.Set("ocsp_option", flattenVpnCertificateSettingOcspOption(o["ocsp-option"], d, "ocsp_option")); err != nil {
		if !fortiAPIPatch(o["ocsp-option"]) {
			return fmt.Errorf("Error reading ocsp_option: %v", err)
		}
	}

	if err = d.Set("ocsp_default_server", flattenVpnCertificateSettingOcspDefaultServer(o["ocsp-default-server"], d, "ocsp_default_server")); err != nil {
		if !fortiAPIPatch(o["ocsp-default-server"]) {
			return fmt.Errorf("Error reading ocsp_default_server: %v", err)
		}
	}

	if err = d.Set("check_ca_cert", flattenVpnCertificateSettingCheckCaCert(o["check-ca-cert"], d, "check_ca_cert")); err != nil {
		if !fortiAPIPatch(o["check-ca-cert"]) {
			return fmt.Errorf("Error reading check_ca_cert: %v", err)
		}
	}

	if err = d.Set("check_ca_chain", flattenVpnCertificateSettingCheckCaChain(o["check-ca-chain"], d, "check_ca_chain")); err != nil {
		if !fortiAPIPatch(o["check-ca-chain"]) {
			return fmt.Errorf("Error reading check_ca_chain: %v", err)
		}
	}

	if err = d.Set("subject_match", flattenVpnCertificateSettingSubjectMatch(o["subject-match"], d, "subject_match")); err != nil {
		if !fortiAPIPatch(o["subject-match"]) {
			return fmt.Errorf("Error reading subject_match: %v", err)
		}
	}

	if err = d.Set("cn_match", flattenVpnCertificateSettingCnMatch(o["cn-match"], d, "cn_match")); err != nil {
		if !fortiAPIPatch(o["cn-match"]) {
			return fmt.Errorf("Error reading cn_match: %v", err)
		}
	}

	if err = d.Set("strict_crl_check", flattenVpnCertificateSettingStrictCrlCheck(o["strict-crl-check"], d, "strict_crl_check")); err != nil {
		if !fortiAPIPatch(o["strict-crl-check"]) {
			return fmt.Errorf("Error reading strict_crl_check: %v", err)
		}
	}

	if err = d.Set("strict_ocsp_check", flattenVpnCertificateSettingStrictOcspCheck(o["strict-ocsp-check"], d, "strict_ocsp_check")); err != nil {
		if !fortiAPIPatch(o["strict-ocsp-check"]) {
			return fmt.Errorf("Error reading strict_ocsp_check: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenVpnCertificateSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("cmp_save_extra_certs", flattenVpnCertificateSettingCmpSaveExtraCerts(o["cmp-save-extra-certs"], d, "cmp_save_extra_certs")); err != nil {
		if !fortiAPIPatch(o["cmp-save-extra-certs"]) {
			return fmt.Errorf("Error reading cmp_save_extra_certs: %v", err)
		}
	}

	if err = d.Set("certname_rsa1024", flattenVpnCertificateSettingCertnameRsa1024(o["certname-rsa1024"], d, "certname_rsa1024")); err != nil {
		if !fortiAPIPatch(o["certname-rsa1024"]) {
			return fmt.Errorf("Error reading certname_rsa1024: %v", err)
		}
	}

	if err = d.Set("certname_rsa2048", flattenVpnCertificateSettingCertnameRsa2048(o["certname-rsa2048"], d, "certname_rsa2048")); err != nil {
		if !fortiAPIPatch(o["certname-rsa2048"]) {
			return fmt.Errorf("Error reading certname_rsa2048: %v", err)
		}
	}

	if err = d.Set("certname_dsa1024", flattenVpnCertificateSettingCertnameDsa1024(o["certname-dsa1024"], d, "certname_dsa1024")); err != nil {
		if !fortiAPIPatch(o["certname-dsa1024"]) {
			return fmt.Errorf("Error reading certname_dsa1024: %v", err)
		}
	}

	if err = d.Set("certname_dsa2048", flattenVpnCertificateSettingCertnameDsa2048(o["certname-dsa2048"], d, "certname_dsa2048")); err != nil {
		if !fortiAPIPatch(o["certname-dsa2048"]) {
			return fmt.Errorf("Error reading certname_dsa2048: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa256", flattenVpnCertificateSettingCertnameEcdsa256(o["certname-ecdsa256"], d, "certname_ecdsa256")); err != nil {
		if !fortiAPIPatch(o["certname-ecdsa256"]) {
			return fmt.Errorf("Error reading certname_ecdsa256: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa384", flattenVpnCertificateSettingCertnameEcdsa384(o["certname-ecdsa384"], d, "certname_ecdsa384")); err != nil {
		if !fortiAPIPatch(o["certname-ecdsa384"]) {
			return fmt.Errorf("Error reading certname_ecdsa384: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateSettingOcspStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingOcspOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingOcspDefaultServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCheckCaCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCheckCaChain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSubjectMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCnMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingStrictCrlCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingStrictOcspCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCmpSaveExtraCerts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameRsa1024(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameRsa2048(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameDsa1024(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameDsa2048(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEcdsa256(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEcdsa384(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ocsp_status"); ok {
		t, err := expandVpnCertificateSettingOcspStatus(d, v, "ocsp_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-status"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_option"); ok {
		t, err := expandVpnCertificateSettingOcspOption(d, v, "ocsp_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-option"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_default_server"); ok {
		t, err := expandVpnCertificateSettingOcspDefaultServer(d, v, "ocsp_default_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-default-server"] = t
		}
	}

	if v, ok := d.GetOk("check_ca_cert"); ok {
		t, err := expandVpnCertificateSettingCheckCaCert(d, v, "check_ca_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("check_ca_chain"); ok {
		t, err := expandVpnCertificateSettingCheckCaChain(d, v, "check_ca_chain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-ca-chain"] = t
		}
	}

	if v, ok := d.GetOk("subject_match"); ok {
		t, err := expandVpnCertificateSettingSubjectMatch(d, v, "subject_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject-match"] = t
		}
	}

	if v, ok := d.GetOk("cn_match"); ok {
		t, err := expandVpnCertificateSettingCnMatch(d, v, "cn_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-match"] = t
		}
	}

	if v, ok := d.GetOk("strict_crl_check"); ok {
		t, err := expandVpnCertificateSettingStrictCrlCheck(d, v, "strict_crl_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-crl-check"] = t
		}
	}

	if v, ok := d.GetOk("strict_ocsp_check"); ok {
		t, err := expandVpnCertificateSettingStrictOcspCheck(d, v, "strict_ocsp_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-ocsp-check"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandVpnCertificateSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("cmp_save_extra_certs"); ok {
		t, err := expandVpnCertificateSettingCmpSaveExtraCerts(d, v, "cmp_save_extra_certs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cmp-save-extra-certs"] = t
		}
	}

	if v, ok := d.GetOk("certname_rsa1024"); ok {
		t, err := expandVpnCertificateSettingCertnameRsa1024(d, v, "certname_rsa1024")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-rsa1024"] = t
		}
	}

	if v, ok := d.GetOk("certname_rsa2048"); ok {
		t, err := expandVpnCertificateSettingCertnameRsa2048(d, v, "certname_rsa2048")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-rsa2048"] = t
		}
	}

	if v, ok := d.GetOk("certname_dsa1024"); ok {
		t, err := expandVpnCertificateSettingCertnameDsa1024(d, v, "certname_dsa1024")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-dsa1024"] = t
		}
	}

	if v, ok := d.GetOk("certname_dsa2048"); ok {
		t, err := expandVpnCertificateSettingCertnameDsa2048(d, v, "certname_dsa2048")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-dsa2048"] = t
		}
	}

	if v, ok := d.GetOk("certname_ecdsa256"); ok {
		t, err := expandVpnCertificateSettingCertnameEcdsa256(d, v, "certname_ecdsa256")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ecdsa256"] = t
		}
	}

	if v, ok := d.GetOk("certname_ecdsa384"); ok {
		t, err := expandVpnCertificateSettingCertnameEcdsa384(d, v, "certname_ecdsa384")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certname-ecdsa384"] = t
		}
	}

	return &obj, nil
}
