// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: VPN certificate setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
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
			"proxy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"proxy_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"proxy_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"proxy_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ssl_ocsp_source_ip": &schema.Schema{
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
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
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
			"subject_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cn_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cn_allow_multi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"crl_verification": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"leaf_crl_absence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"chain_crl_absence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"cmp_key_usage_checking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_expire_warning": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
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
			"certname_rsa4096": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
			"certname_ecdsa521": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"certname_ed25519": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"certname_ed448": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnCertificateSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateSetting(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnCertificateSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnCertificateSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing VpnCertificateSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateSetting resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateSettingOcspStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingOcspOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxyPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxyUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingProxyPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingSslOcspSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingOcspDefaultServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCheckCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCheckCaChain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingSubjectMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingSubjectSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCnMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCnAllowMulti(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerification(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "expiry"
	if _, ok := i["expiry"]; ok {
		result["expiry"] = flattenVpnCertificateSettingCrlVerificationExpiry(i["expiry"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "leaf_crl_absence"
	if _, ok := i["leaf-crl-absence"]; ok {
		result["leaf_crl_absence"] = flattenVpnCertificateSettingCrlVerificationLeafCrlAbsence(i["leaf-crl-absence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "chain_crl_absence"
	if _, ok := i["chain-crl-absence"]; ok {
		result["chain_crl_absence"] = flattenVpnCertificateSettingCrlVerificationChainCrlAbsence(i["chain-crl-absence"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnCertificateSettingCrlVerificationExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerificationLeafCrlAbsence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCrlVerificationChainCrlAbsence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingStrictCrlCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingStrictOcspCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCmpSaveExtraCerts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCmpKeyUsageChecking(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertExpireWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameRsa1024(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameRsa2048(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameRsa4096(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameDsa1024(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameDsa2048(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEcdsa256(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEcdsa384(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEcdsa521(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEd25519(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateSettingCertnameEd448(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnCertificateSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("ocsp_status", flattenVpnCertificateSettingOcspStatus(o["ocsp-status"], d, "ocsp_status", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-status"]) {
			return fmt.Errorf("Error reading ocsp_status: %v", err)
		}
	}

	if err = d.Set("ocsp_option", flattenVpnCertificateSettingOcspOption(o["ocsp-option"], d, "ocsp_option", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-option"]) {
			return fmt.Errorf("Error reading ocsp_option: %v", err)
		}
	}

	if err = d.Set("proxy", flattenVpnCertificateSettingProxy(o["proxy"], d, "proxy", sv)); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("proxy_port", flattenVpnCertificateSettingProxyPort(o["proxy-port"], d, "proxy_port", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-port"]) {
			return fmt.Errorf("Error reading proxy_port: %v", err)
		}
	}

	if err = d.Set("proxy_username", flattenVpnCertificateSettingProxyUsername(o["proxy-username"], d, "proxy_username", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-username"]) {
			return fmt.Errorf("Error reading proxy_username: %v", err)
		}
	}

	if err = d.Set("proxy_password", flattenVpnCertificateSettingProxyPassword(o["proxy-password"], d, "proxy_password", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-password"]) {
			return fmt.Errorf("Error reading proxy_password: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenVpnCertificateSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ssl_ocsp_source_ip", flattenVpnCertificateSettingSslOcspSourceIp(o["ssl-ocsp-source-ip"], d, "ssl_ocsp_source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ocsp-source-ip"]) {
			return fmt.Errorf("Error reading ssl_ocsp_source_ip: %v", err)
		}
	}

	if err = d.Set("ocsp_default_server", flattenVpnCertificateSettingOcspDefaultServer(o["ocsp-default-server"], d, "ocsp_default_server", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-default-server"]) {
			return fmt.Errorf("Error reading ocsp_default_server: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenVpnCertificateSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnCertificateSettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("check_ca_cert", flattenVpnCertificateSettingCheckCaCert(o["check-ca-cert"], d, "check_ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["check-ca-cert"]) {
			return fmt.Errorf("Error reading check_ca_cert: %v", err)
		}
	}

	if err = d.Set("check_ca_chain", flattenVpnCertificateSettingCheckCaChain(o["check-ca-chain"], d, "check_ca_chain", sv)); err != nil {
		if !fortiAPIPatch(o["check-ca-chain"]) {
			return fmt.Errorf("Error reading check_ca_chain: %v", err)
		}
	}

	if err = d.Set("subject_match", flattenVpnCertificateSettingSubjectMatch(o["subject-match"], d, "subject_match", sv)); err != nil {
		if !fortiAPIPatch(o["subject-match"]) {
			return fmt.Errorf("Error reading subject_match: %v", err)
		}
	}

	if err = d.Set("subject_set", flattenVpnCertificateSettingSubjectSet(o["subject-set"], d, "subject_set", sv)); err != nil {
		if !fortiAPIPatch(o["subject-set"]) {
			return fmt.Errorf("Error reading subject_set: %v", err)
		}
	}

	if err = d.Set("cn_match", flattenVpnCertificateSettingCnMatch(o["cn-match"], d, "cn_match", sv)); err != nil {
		if !fortiAPIPatch(o["cn-match"]) {
			return fmt.Errorf("Error reading cn_match: %v", err)
		}
	}

	if err = d.Set("cn_allow_multi", flattenVpnCertificateSettingCnAllowMulti(o["cn-allow-multi"], d, "cn_allow_multi", sv)); err != nil {
		if !fortiAPIPatch(o["cn-allow-multi"]) {
			return fmt.Errorf("Error reading cn_allow_multi: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("crl_verification", flattenVpnCertificateSettingCrlVerification(o["crl-verification"], d, "crl_verification", sv)); err != nil {
			if !fortiAPIPatch(o["crl-verification"]) {
				return fmt.Errorf("Error reading crl_verification: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("crl_verification"); ok {
			if err = d.Set("crl_verification", flattenVpnCertificateSettingCrlVerification(o["crl-verification"], d, "crl_verification", sv)); err != nil {
				if !fortiAPIPatch(o["crl-verification"]) {
					return fmt.Errorf("Error reading crl_verification: %v", err)
				}
			}
		}
	}

	if err = d.Set("strict_crl_check", flattenVpnCertificateSettingStrictCrlCheck(o["strict-crl-check"], d, "strict_crl_check", sv)); err != nil {
		if !fortiAPIPatch(o["strict-crl-check"]) {
			return fmt.Errorf("Error reading strict_crl_check: %v", err)
		}
	}

	if err = d.Set("strict_ocsp_check", flattenVpnCertificateSettingStrictOcspCheck(o["strict-ocsp-check"], d, "strict_ocsp_check", sv)); err != nil {
		if !fortiAPIPatch(o["strict-ocsp-check"]) {
			return fmt.Errorf("Error reading strict_ocsp_check: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenVpnCertificateSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("cmp_save_extra_certs", flattenVpnCertificateSettingCmpSaveExtraCerts(o["cmp-save-extra-certs"], d, "cmp_save_extra_certs", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-save-extra-certs"]) {
			return fmt.Errorf("Error reading cmp_save_extra_certs: %v", err)
		}
	}

	if err = d.Set("cmp_key_usage_checking", flattenVpnCertificateSettingCmpKeyUsageChecking(o["cmp-key-usage-checking"], d, "cmp_key_usage_checking", sv)); err != nil {
		if !fortiAPIPatch(o["cmp-key-usage-checking"]) {
			return fmt.Errorf("Error reading cmp_key_usage_checking: %v", err)
		}
	}

	if err = d.Set("cert_expire_warning", flattenVpnCertificateSettingCertExpireWarning(o["cert-expire-warning"], d, "cert_expire_warning", sv)); err != nil {
		if !fortiAPIPatch(o["cert-expire-warning"]) {
			return fmt.Errorf("Error reading cert_expire_warning: %v", err)
		}
	}

	if err = d.Set("certname_rsa1024", flattenVpnCertificateSettingCertnameRsa1024(o["certname-rsa1024"], d, "certname_rsa1024", sv)); err != nil {
		if !fortiAPIPatch(o["certname-rsa1024"]) {
			return fmt.Errorf("Error reading certname_rsa1024: %v", err)
		}
	}

	if err = d.Set("certname_rsa2048", flattenVpnCertificateSettingCertnameRsa2048(o["certname-rsa2048"], d, "certname_rsa2048", sv)); err != nil {
		if !fortiAPIPatch(o["certname-rsa2048"]) {
			return fmt.Errorf("Error reading certname_rsa2048: %v", err)
		}
	}

	if err = d.Set("certname_rsa4096", flattenVpnCertificateSettingCertnameRsa4096(o["certname-rsa4096"], d, "certname_rsa4096", sv)); err != nil {
		if !fortiAPIPatch(o["certname-rsa4096"]) {
			return fmt.Errorf("Error reading certname_rsa4096: %v", err)
		}
	}

	if err = d.Set("certname_dsa1024", flattenVpnCertificateSettingCertnameDsa1024(o["certname-dsa1024"], d, "certname_dsa1024", sv)); err != nil {
		if !fortiAPIPatch(o["certname-dsa1024"]) {
			return fmt.Errorf("Error reading certname_dsa1024: %v", err)
		}
	}

	if err = d.Set("certname_dsa2048", flattenVpnCertificateSettingCertnameDsa2048(o["certname-dsa2048"], d, "certname_dsa2048", sv)); err != nil {
		if !fortiAPIPatch(o["certname-dsa2048"]) {
			return fmt.Errorf("Error reading certname_dsa2048: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa256", flattenVpnCertificateSettingCertnameEcdsa256(o["certname-ecdsa256"], d, "certname_ecdsa256", sv)); err != nil {
		if !fortiAPIPatch(o["certname-ecdsa256"]) {
			return fmt.Errorf("Error reading certname_ecdsa256: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa384", flattenVpnCertificateSettingCertnameEcdsa384(o["certname-ecdsa384"], d, "certname_ecdsa384", sv)); err != nil {
		if !fortiAPIPatch(o["certname-ecdsa384"]) {
			return fmt.Errorf("Error reading certname_ecdsa384: %v", err)
		}
	}

	if err = d.Set("certname_ecdsa521", flattenVpnCertificateSettingCertnameEcdsa521(o["certname-ecdsa521"], d, "certname_ecdsa521", sv)); err != nil {
		if !fortiAPIPatch(o["certname-ecdsa521"]) {
			return fmt.Errorf("Error reading certname_ecdsa521: %v", err)
		}
	}

	if err = d.Set("certname_ed25519", flattenVpnCertificateSettingCertnameEd25519(o["certname-ed25519"], d, "certname_ed25519", sv)); err != nil {
		if !fortiAPIPatch(o["certname-ed25519"]) {
			return fmt.Errorf("Error reading certname_ed25519: %v", err)
		}
	}

	if err = d.Set("certname_ed448", flattenVpnCertificateSettingCertnameEd448(o["certname-ed448"], d, "certname_ed448", sv)); err != nil {
		if !fortiAPIPatch(o["certname-ed448"]) {
			return fmt.Errorf("Error reading certname_ed448: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnCertificateSettingOcspStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingOcspOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxyPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxyUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingProxyPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSslOcspSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingOcspDefaultServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCheckCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCheckCaChain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSubjectMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSubjectSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCnMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCnAllowMulti(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerification(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "expiry"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["expiry"] = nil
		} else {
			result["expiry"], _ = expandVpnCertificateSettingCrlVerificationExpiry(d, i["expiry"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "leaf_crl_absence"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["leaf-crl-absence"] = nil
		} else {
			result["leaf-crl-absence"], _ = expandVpnCertificateSettingCrlVerificationLeafCrlAbsence(d, i["leaf_crl_absence"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "chain_crl_absence"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["chain-crl-absence"] = nil
		} else {
			result["chain-crl-absence"], _ = expandVpnCertificateSettingCrlVerificationChainCrlAbsence(d, i["chain_crl_absence"], pre_append, sv)
		}
	}

	return result, nil
}

func expandVpnCertificateSettingCrlVerificationExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerificationLeafCrlAbsence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCrlVerificationChainCrlAbsence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingStrictCrlCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingStrictOcspCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCmpSaveExtraCerts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCmpKeyUsageChecking(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertExpireWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameRsa1024(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameRsa2048(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameRsa4096(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameDsa1024(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameDsa2048(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEcdsa256(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEcdsa384(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEcdsa521(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEd25519(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateSettingCertnameEd448(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ocsp_status"); ok {
		if setArgNil {
			obj["ocsp-status"] = nil
		} else {
			t, err := expandVpnCertificateSettingOcspStatus(d, v, "ocsp_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ocsp-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ocsp_option"); ok {
		if setArgNil {
			obj["ocsp-option"] = nil
		} else {
			t, err := expandVpnCertificateSettingOcspOption(d, v, "ocsp_option", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ocsp-option"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		if setArgNil {
			obj["proxy"] = nil
		} else {
			t, err := expandVpnCertificateSettingProxy(d, v, "proxy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_port"); ok {
		if setArgNil {
			obj["proxy-port"] = nil
		} else {
			t, err := expandVpnCertificateSettingProxyPort(d, v, "proxy_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_username"); ok {
		if setArgNil {
			obj["proxy-username"] = nil
		} else {
			t, err := expandVpnCertificateSettingProxyUsername(d, v, "proxy_username", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-username"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_password"); ok {
		if setArgNil {
			obj["proxy-password"] = nil
		} else {
			t, err := expandVpnCertificateSettingProxyPassword(d, v, "proxy_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-password"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandVpnCertificateSettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_ocsp_source_ip"); ok {
		if setArgNil {
			obj["ssl-ocsp-source-ip"] = nil
		} else {
			t, err := expandVpnCertificateSettingSslOcspSourceIp(d, v, "ssl_ocsp_source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-ocsp-source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("ocsp_default_server"); ok {
		if setArgNil {
			obj["ocsp-default-server"] = nil
		} else {
			t, err := expandVpnCertificateSettingOcspDefaultServer(d, v, "ocsp_default_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ocsp-default-server"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandVpnCertificateSettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandVpnCertificateSettingInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("check_ca_cert"); ok {
		if setArgNil {
			obj["check-ca-cert"] = nil
		} else {
			t, err := expandVpnCertificateSettingCheckCaCert(d, v, "check_ca_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["check-ca-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("check_ca_chain"); ok {
		if setArgNil {
			obj["check-ca-chain"] = nil
		} else {
			t, err := expandVpnCertificateSettingCheckCaChain(d, v, "check_ca_chain", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["check-ca-chain"] = t
			}
		}
	}

	if v, ok := d.GetOk("subject_match"); ok {
		if setArgNil {
			obj["subject-match"] = nil
		} else {
			t, err := expandVpnCertificateSettingSubjectMatch(d, v, "subject_match", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["subject-match"] = t
			}
		}
	}

	if v, ok := d.GetOk("subject_set"); ok {
		if setArgNil {
			obj["subject-set"] = nil
		} else {
			t, err := expandVpnCertificateSettingSubjectSet(d, v, "subject_set", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["subject-set"] = t
			}
		}
	}

	if v, ok := d.GetOk("cn_match"); ok {
		if setArgNil {
			obj["cn-match"] = nil
		} else {
			t, err := expandVpnCertificateSettingCnMatch(d, v, "cn_match", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cn-match"] = t
			}
		}
	}

	if v, ok := d.GetOk("cn_allow_multi"); ok {
		if setArgNil {
			obj["cn-allow-multi"] = nil
		} else {
			t, err := expandVpnCertificateSettingCnAllowMulti(d, v, "cn_allow_multi", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cn-allow-multi"] = t
			}
		}
	}

	if v, ok := d.GetOk("crl_verification"); ok {
		t, err := expandVpnCertificateSettingCrlVerification(d, v, "crl_verification", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl-verification"] = t
		}
	}

	if v, ok := d.GetOk("strict_crl_check"); ok {
		if setArgNil {
			obj["strict-crl-check"] = nil
		} else {
			t, err := expandVpnCertificateSettingStrictCrlCheck(d, v, "strict_crl_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strict-crl-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("strict_ocsp_check"); ok {
		if setArgNil {
			obj["strict-ocsp-check"] = nil
		} else {
			t, err := expandVpnCertificateSettingStrictOcspCheck(d, v, "strict_ocsp_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strict-ocsp-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		if setArgNil {
			obj["ssl-min-proto-version"] = nil
		} else {
			t, err := expandVpnCertificateSettingSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-min-proto-version"] = t
			}
		}
	}

	if v, ok := d.GetOk("cmp_save_extra_certs"); ok {
		if setArgNil {
			obj["cmp-save-extra-certs"] = nil
		} else {
			t, err := expandVpnCertificateSettingCmpSaveExtraCerts(d, v, "cmp_save_extra_certs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cmp-save-extra-certs"] = t
			}
		}
	}

	if v, ok := d.GetOk("cmp_key_usage_checking"); ok {
		if setArgNil {
			obj["cmp-key-usage-checking"] = nil
		} else {
			t, err := expandVpnCertificateSettingCmpKeyUsageChecking(d, v, "cmp_key_usage_checking", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cmp-key-usage-checking"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("cert_expire_warning"); ok {
		if setArgNil {
			obj["cert-expire-warning"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertExpireWarning(d, v, "cert_expire_warning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cert-expire-warning"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_rsa1024"); ok {
		if setArgNil {
			obj["certname-rsa1024"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameRsa1024(d, v, "certname_rsa1024", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-rsa1024"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_rsa2048"); ok {
		if setArgNil {
			obj["certname-rsa2048"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameRsa2048(d, v, "certname_rsa2048", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-rsa2048"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_rsa4096"); ok {
		if setArgNil {
			obj["certname-rsa4096"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameRsa4096(d, v, "certname_rsa4096", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-rsa4096"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_dsa1024"); ok {
		if setArgNil {
			obj["certname-dsa1024"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameDsa1024(d, v, "certname_dsa1024", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-dsa1024"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_dsa2048"); ok {
		if setArgNil {
			obj["certname-dsa2048"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameDsa2048(d, v, "certname_dsa2048", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-dsa2048"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_ecdsa256"); ok {
		if setArgNil {
			obj["certname-ecdsa256"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameEcdsa256(d, v, "certname_ecdsa256", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-ecdsa256"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_ecdsa384"); ok {
		if setArgNil {
			obj["certname-ecdsa384"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameEcdsa384(d, v, "certname_ecdsa384", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-ecdsa384"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_ecdsa521"); ok {
		if setArgNil {
			obj["certname-ecdsa521"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameEcdsa521(d, v, "certname_ecdsa521", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-ecdsa521"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_ed25519"); ok {
		if setArgNil {
			obj["certname-ed25519"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameEd25519(d, v, "certname_ed25519", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-ed25519"] = t
			}
		}
	}

	if v, ok := d.GetOk("certname_ed448"); ok {
		if setArgNil {
			obj["certname-ed448"] = nil
		} else {
			t, err := expandVpnCertificateSettingCertnameEd448(d, v, "certname_ed448", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certname-ed448"] = t
			}
		}
	}

	return &obj, nil
}
