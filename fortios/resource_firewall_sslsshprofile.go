// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure SSL/SSH protocol options.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallSslSshProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSshProfileCreate,
		Read:   resourceFirewallSslSshProfileRead,
		Update: resourceFirewallSslSshProfileUpdate,
		Delete: resourceFirewallSslSshProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invalid_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"https": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invalid_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ftps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invalid_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imaps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invalid_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"pop3s": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invalid_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"smtps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"invalid_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssh_policy_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssh_tun_policy_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssh_algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exempt": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 512),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiguard_category": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"address6": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"wildcard_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"regex": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"server_cert_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_ssl_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"caname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"untrusted_caname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"smtps_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pop3s_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"imaps_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ftps_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_other_client_cert_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_anomalies_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exemptions_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_over_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mapi_over_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslSshProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSslSshProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfile resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallSslSshProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSslSshProfile")
	}

	return resourceFirewallSslSshProfileRead(d, m)
}

func resourceFirewallSslSshProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSslSshProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSslSshProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSslSshProfile")
	}

	return resourceFirewallSslSshProfileRead(d, m)
}

func resourceFirewallSslSshProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallSslSshProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSshProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSshProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallSslSshProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfile resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSsl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallSslSshProfileSslInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileSslClientCertRequest(i["client-cert-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileSslUnsupportedSsl(i["unsupported-ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileSslInvalidServerCert(i["invalid-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileSslUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileSslSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSslInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslInvalidServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileHttpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileHttpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileHttpsClientCertRequest(i["client-cert-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileHttpsUnsupportedSsl(i["unsupported-ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileHttpsInvalidServerCert(i["invalid-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileHttpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileHttpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileHttpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallSslSshProfileHttpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileFtpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileFtpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileFtpsClientCertRequest(i["client-cert-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileFtpsUnsupportedSsl(i["unsupported-ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileFtpsInvalidServerCert(i["invalid-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileFtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileFtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileFtpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallSslSshProfileFtpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileImapsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileImapsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileImapsClientCertRequest(i["client-cert-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileImapsUnsupportedSsl(i["unsupported-ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileImapsInvalidServerCert(i["invalid-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileImapsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileImapsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileImapsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallSslSshProfileImapsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3S(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfilePop3SPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfilePop3SStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfilePop3SClientCertRequest(i["client-cert-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfilePop3SUnsupportedSsl(i["unsupported-ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfilePop3SInvalidServerCert(i["invalid-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfilePop3SUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfilePop3SSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfilePop3SPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallSslSshProfilePop3SStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SInvalidServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileSmtpsPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSmtpsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileSmtpsClientCertRequest(i["client-cert-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileSmtpsUnsupportedSsl(i["unsupported-ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileSmtpsInvalidServerCert(i["invalid-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileSmtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileSmtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSmtpsPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallSslSshProfileSmtpsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSsh(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileSshPorts(i["ports"], d, pre_append)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSshStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallSslSshProfileSshInspectAll(i["inspect-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := i["unsupported-version"]; ok {
		result["unsupported_version"] = flattenFirewallSslSshProfileSshUnsupportedVersion(i["unsupported-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_policy_check"
	if _, ok := i["ssh-policy-check"]; ok {
		result["ssh_policy_check"] = flattenFirewallSslSshProfileSshSshPolicyCheck(i["ssh-policy-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := i["ssh-tun-policy-check"]; ok {
		result["ssh_tun_policy_check"] = flattenFirewallSslSshProfileSshSshTunPolicyCheck(i["ssh-tun-policy-check"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := i["ssh-algorithm"]; ok {
		result["ssh_algorithm"] = flattenFirewallSslSshProfileSshSshAlgorithm(i["ssh-algorithm"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSshPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenFirewallSslSshProfileSshStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshInspectAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshUnsupportedVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshPolicyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshTunPolicyCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileWhitelist(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExempt(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallSslSshProfileSslExemptId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenFirewallSslSshProfileSslExemptType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := i["fortiguard-category"]; ok {
			tmp["fortiguard_category"] = flattenFirewallSslSshProfileSslExemptFortiguardCategory(i["fortiguard-category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = flattenFirewallSslSshProfileSslExemptAddress(i["address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := i["address6"]; ok {
			tmp["address6"] = flattenFirewallSslSshProfileSslExemptAddress6(i["address6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := i["wildcard-fqdn"]; ok {
			tmp["wildcard_fqdn"] = flattenFirewallSslSshProfileSslExemptWildcardFqdn(i["wildcard-fqdn"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			tmp["regex"] = flattenFirewallSslSshProfileSslExemptRegex(i["regex"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallSslSshProfileSslExemptId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptFortiguardCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileServerCertMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileUseSslServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileUntrustedCaname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallSslSshProfileSslServerId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenFirewallSslSshProfileSslServerIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_cert_request"
		if _, ok := i["https-client-cert-request"]; ok {
			tmp["https_client_cert_request"] = flattenFirewallSslSshProfileSslServerHttpsClientCertRequest(i["https-client-cert-request"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_cert_request"
		if _, ok := i["smtps-client-cert-request"]; ok {
			tmp["smtps_client_cert_request"] = flattenFirewallSslSshProfileSslServerSmtpsClientCertRequest(i["smtps-client-cert-request"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_cert_request"
		if _, ok := i["pop3s-client-cert-request"]; ok {
			tmp["pop3s_client_cert_request"] = flattenFirewallSslSshProfileSslServerPop3SClientCertRequest(i["pop3s-client-cert-request"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_cert_request"
		if _, ok := i["imaps-client-cert-request"]; ok {
			tmp["imaps_client_cert_request"] = flattenFirewallSslSshProfileSslServerImapsClientCertRequest(i["imaps-client-cert-request"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_cert_request"
		if _, ok := i["ftps-client-cert-request"]; ok {
			tmp["ftps_client_cert_request"] = flattenFirewallSslSshProfileSslServerFtpsClientCertRequest(i["ftps-client-cert-request"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_cert_request"
		if _, ok := i["ssl-other-client-cert-request"]; ok {
			tmp["ssl_other_client_cert_request"] = flattenFirewallSslSshProfileSslServerSslOtherClientCertRequest(i["ssl-other-client-cert-request"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallSslSshProfileSslServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerHttpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSmtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerPop3SClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerImapsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerFtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSslOtherClientCertRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslAnomaliesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileRpcOverHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSshProfileMapiOverHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallSslSshProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallSslSshProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o["ssl"], d, "ssl")); err != nil {
			if !fortiAPIPatch(o["ssl"]) {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl"); ok {
			if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o["ssl"], d, "ssl")); err != nil {
				if !fortiAPIPatch(o["ssl"]) {
					return fmt.Errorf("Error reading ssl: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("https", flattenFirewallSslSshProfileHttps(o["https"], d, "https")); err != nil {
			if !fortiAPIPatch(o["https"]) {
				return fmt.Errorf("Error reading https: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("https"); ok {
			if err = d.Set("https", flattenFirewallSslSshProfileHttps(o["https"], d, "https")); err != nil {
				if !fortiAPIPatch(o["https"]) {
					return fmt.Errorf("Error reading https: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o["ftps"], d, "ftps")); err != nil {
			if !fortiAPIPatch(o["ftps"]) {
				return fmt.Errorf("Error reading ftps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftps"); ok {
			if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o["ftps"], d, "ftps")); err != nil {
				if !fortiAPIPatch(o["ftps"]) {
					return fmt.Errorf("Error reading ftps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o["imaps"], d, "imaps")); err != nil {
			if !fortiAPIPatch(o["imaps"]) {
				return fmt.Errorf("Error reading imaps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imaps"); ok {
			if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o["imaps"], d, "imaps")); err != nil {
				if !fortiAPIPatch(o["imaps"]) {
					return fmt.Errorf("Error reading imaps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s")); err != nil {
			if !fortiAPIPatch(o["pop3s"]) {
				return fmt.Errorf("Error reading pop3s: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3s"); ok {
			if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s")); err != nil {
				if !fortiAPIPatch(o["pop3s"]) {
					return fmt.Errorf("Error reading pop3s: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o["smtps"], d, "smtps")); err != nil {
			if !fortiAPIPatch(o["smtps"]) {
				return fmt.Errorf("Error reading smtps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtps"); ok {
			if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o["smtps"], d, "smtps")); err != nil {
				if !fortiAPIPatch(o["smtps"]) {
					return fmt.Errorf("Error reading smtps: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o["ssh"], d, "ssh")); err != nil {
			if !fortiAPIPatch(o["ssh"]) {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o["ssh"], d, "ssh")); err != nil {
				if !fortiAPIPatch(o["ssh"]) {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if err = d.Set("whitelist", flattenFirewallSslSshProfileWhitelist(o["whitelist"], d, "whitelist")); err != nil {
		if !fortiAPIPatch(o["whitelist"]) {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt")); err != nil {
			if !fortiAPIPatch(o["ssl-exempt"]) {
				return fmt.Errorf("Error reading ssl_exempt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_exempt"); ok {
			if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt")); err != nil {
				if !fortiAPIPatch(o["ssl-exempt"]) {
					return fmt.Errorf("Error reading ssl_exempt: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_cert_mode", flattenFirewallSslSshProfileServerCertMode(o["server-cert-mode"], d, "server_cert_mode")); err != nil {
		if !fortiAPIPatch(o["server-cert-mode"]) {
			return fmt.Errorf("Error reading server_cert_mode: %v", err)
		}
	}

	if err = d.Set("use_ssl_server", flattenFirewallSslSshProfileUseSslServer(o["use-ssl-server"], d, "use_ssl_server")); err != nil {
		if !fortiAPIPatch(o["use-ssl-server"]) {
			return fmt.Errorf("Error reading use_ssl_server: %v", err)
		}
	}

	if err = d.Set("caname", flattenFirewallSslSshProfileCaname(o["caname"], d, "caname")); err != nil {
		if !fortiAPIPatch(o["caname"]) {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("untrusted_caname", flattenFirewallSslSshProfileUntrustedCaname(o["untrusted-caname"], d, "untrusted_caname")); err != nil {
		if !fortiAPIPatch(o["untrusted-caname"]) {
			return fmt.Errorf("Error reading untrusted_caname: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenFirewallSslSshProfileServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if !fortiAPIPatch(o["server-cert"]) {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server")); err != nil {
			if !fortiAPIPatch(o["ssl-server"]) {
				return fmt.Errorf("Error reading ssl_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server"); ok {
			if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server")); err != nil {
				if !fortiAPIPatch(o["ssl-server"]) {
					return fmt.Errorf("Error reading ssl_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_anomalies_log", flattenFirewallSslSshProfileSslAnomaliesLog(o["ssl-anomalies-log"], d, "ssl_anomalies_log")); err != nil {
		if !fortiAPIPatch(o["ssl-anomalies-log"]) {
			return fmt.Errorf("Error reading ssl_anomalies_log: %v", err)
		}
	}

	if err = d.Set("ssl_exemptions_log", flattenFirewallSslSshProfileSslExemptionsLog(o["ssl-exemptions-log"], d, "ssl_exemptions_log")); err != nil {
		if !fortiAPIPatch(o["ssl-exemptions-log"]) {
			return fmt.Errorf("Error reading ssl_exemptions_log: %v", err)
		}
	}

	if err = d.Set("rpc_over_https", flattenFirewallSslSshProfileRpcOverHttps(o["rpc-over-https"], d, "rpc_over_https")); err != nil {
		if !fortiAPIPatch(o["rpc-over-https"]) {
			return fmt.Errorf("Error reading rpc_over_https: %v", err)
		}
	}

	if err = d.Set("mapi_over_https", flattenFirewallSslSshProfileMapiOverHttps(o["mapi-over-https"], d, "mapi_over_https")); err != nil {
		if !fortiAPIPatch(o["mapi-over-https"]) {
			return fmt.Errorf("Error reading mapi_over_https: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSshProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallSslSshProfileSslInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileSslClientCertRequest(d, i["client_cert_request"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileSslUnsupportedSsl(d, i["unsupported_ssl"], pre_append)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileSslInvalidServerCert(d, i["invalid_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileSslUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileSslSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSslInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslInvalidServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileHttpsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileHttpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileHttpsClientCertRequest(d, i["client_cert_request"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileHttpsUnsupportedSsl(d, i["unsupported_ssl"], pre_append)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileHttpsInvalidServerCert(d, i["invalid_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileHttpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileHttpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileHttpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileFtpsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileFtpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileFtpsClientCertRequest(d, i["client_cert_request"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileFtpsUnsupportedSsl(d, i["unsupported_ssl"], pre_append)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileFtpsInvalidServerCert(d, i["invalid_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileFtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileFtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileFtpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileImapsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileImapsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileImapsClientCertRequest(d, i["client_cert_request"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileImapsUnsupportedSsl(d, i["unsupported_ssl"], pre_append)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileImapsInvalidServerCert(d, i["invalid_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileImapsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileImapsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileImapsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3S(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfilePop3SPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfilePop3SStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfilePop3SClientCertRequest(d, i["client_cert_request"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfilePop3SUnsupportedSsl(d, i["unsupported_ssl"], pre_append)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfilePop3SInvalidServerCert(d, i["invalid_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfilePop3SUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfilePop3SSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfilePop3SPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SInvalidServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileSmtpsPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileSmtpsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileSmtpsClientCertRequest(d, i["client_cert_request"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSsl(d, i["unsupported_ssl"], pre_append)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileSmtpsInvalidServerCert(d, i["invalid_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileSmtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileSmtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSmtpsPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileSshPorts(d, i["ports"], pre_append)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileSshStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallSslSshProfileSshInspectAll(d, i["inspect_all"], pre_append)
	}
	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-version"], _ = expandFirewallSslSshProfileSshUnsupportedVersion(d, i["unsupported_version"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_policy_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-policy-check"], _ = expandFirewallSslSshProfileSshSshPolicyCheck(d, i["ssh_policy_check"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-tun-policy-check"], _ = expandFirewallSslSshProfileSshSshTunPolicyCheck(d, i["ssh_tun_policy_check"], pre_append)
	}
	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-algorithm"], _ = expandFirewallSslSshProfileSshSshAlgorithm(d, i["ssh_algorithm"], pre_append)
	}

	return result, nil
}

func expandFirewallSslSshProfileSshPorts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshInspectAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshUnsupportedVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshPolicyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshTunPolicyCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileWhitelist(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFirewallSslSshProfileSslExemptId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandFirewallSslSshProfileSslExemptType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fortiguard-category"], _ = expandFirewallSslSshProfileSslExemptFortiguardCategory(d, i["fortiguard_category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandFirewallSslSshProfileSslExemptAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address6"], _ = expandFirewallSslSshProfileSslExemptAddress6(d, i["address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wildcard-fqdn"], _ = expandFirewallSslSshProfileSslExemptWildcardFqdn(d, i["wildcard_fqdn"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regex"], _ = expandFirewallSslSshProfileSslExemptRegex(d, i["regex"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileSslExemptId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptFortiguardCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileServerCertMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileUseSslServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileUntrustedCaname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileServerCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFirewallSslSshProfileSslServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFirewallSslSshProfileSslServerIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-client-cert-request"], _ = expandFirewallSslSshProfileSslServerHttpsClientCertRequest(d, i["https_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["smtps-client-cert-request"], _ = expandFirewallSslSshProfileSslServerSmtpsClientCertRequest(d, i["smtps_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pop3s-client-cert-request"], _ = expandFirewallSslSshProfileSslServerPop3SClientCertRequest(d, i["pop3s_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imaps-client-cert-request"], _ = expandFirewallSslSshProfileSslServerImapsClientCertRequest(d, i["imaps_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ftps-client-cert-request"], _ = expandFirewallSslSshProfileSslServerFtpsClientCertRequest(d, i["ftps_client_cert_request"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-other-client-cert-request"], _ = expandFirewallSslSshProfileSslServerSslOtherClientCertRequest(d, i["ssl_other_client_cert_request"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileSslServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerHttpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSmtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerPop3SClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerImapsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerFtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSslOtherClientCertRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslAnomaliesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileRpcOverHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileMapiOverHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallSslSshProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallSslSshProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandFirewallSslSshProfileSsl(d, v, "ssl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("https"); ok {
		t, err := expandFirewallSslSshProfileHttps(d, v, "https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https"] = t
		}
	}

	if v, ok := d.GetOk("ftps"); ok {
		t, err := expandFirewallSslSshProfileFtps(d, v, "ftps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps"] = t
		}
	}

	if v, ok := d.GetOk("imaps"); ok {
		t, err := expandFirewallSslSshProfileImaps(d, v, "imaps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps"] = t
		}
	}

	if v, ok := d.GetOk("pop3s"); ok {
		t, err := expandFirewallSslSshProfilePop3S(d, v, "pop3s")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s"] = t
		}
	}

	if v, ok := d.GetOk("smtps"); ok {
		t, err := expandFirewallSslSshProfileSmtps(d, v, "smtps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandFirewallSslSshProfileSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok {
		t, err := expandFirewallSslSshProfileWhitelist(d, v, "whitelist")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exempt"); ok {
		t, err := expandFirewallSslSshProfileSslExempt(d, v, "ssl_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exempt"] = t
		}
	}

	if v, ok := d.GetOk("server_cert_mode"); ok {
		t, err := expandFirewallSslSshProfileServerCertMode(d, v, "server_cert_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert-mode"] = t
		}
	}

	if v, ok := d.GetOk("use_ssl_server"); ok {
		t, err := expandFirewallSslSshProfileUseSslServer(d, v, "use_ssl_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("caname"); ok {
		t, err := expandFirewallSslSshProfileCaname(d, v, "caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caname"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_caname"); ok {
		t, err := expandFirewallSslSshProfileUntrustedCaname(d, v, "untrusted_caname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-caname"] = t
		}
	}

	if v, ok := d.GetOk("server_cert"); ok {
		t, err := expandFirewallSslSshProfileServerCert(d, v, "server_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server"); ok {
		t, err := expandFirewallSslSshProfileSslServer(d, v, "ssl_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_anomalies_log"); ok {
		t, err := expandFirewallSslSshProfileSslAnomaliesLog(d, v, "ssl_anomalies_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-anomalies-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemptions_log"); ok {
		t, err := expandFirewallSslSshProfileSslExemptionsLog(d, v, "ssl_exemptions_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemptions-log"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_https"); ok {
		t, err := expandFirewallSslSshProfileRpcOverHttps(d, v, "rpc_over_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-https"] = t
		}
	}

	if v, ok := d.GetOk("mapi_over_https"); ok {
		t, err := expandFirewallSslSshProfileMapiOverHttps(d, v, "mapi_over_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi-over-https"] = t
		}
	}

	return &obj, nil
}
