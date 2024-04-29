// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SSL/SSH protocol options.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspect_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
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
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_probe_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"https": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
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
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_probe_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ftps": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
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
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sni_server_cert_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_allowed_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"imaps": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
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
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
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
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
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
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
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
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
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
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
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
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"proxy_after_tcp_handshake": &schema.Schema{
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
			"dot": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_after_tcp_handshake": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unsupported_ssl_negotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"expired_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revoked_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"untrusted_server_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_timeout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cert_validation_failure": &schema.Schema{
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
			"allowlist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_blocklisted_certificates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"whitelist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_blacklisted_certificates": &schema.Schema{
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
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"address6": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"wildcard_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"https_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"smtps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pop3s_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"imaps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ftps_client_certificate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_other_client_certificate": &schema.Schema{
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
			"ssl_exemption_ip_rating": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_anomaly_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_exemption_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"ssl_negotiation_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_cert_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_handshake_log": &schema.Schema{
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
			"supported_alpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallSslSshProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallSslSshProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSslSshProfile resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallSslSshProfile(obj, vdomparam)

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

	obj, err := getObjectFirewallSslSshProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSshProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSslSshProfile(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallSslSshProfile(mkey, vdomparam)
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

	o, err := c.ReadFirewallSslSshProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSshProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSshProfile resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSshProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSsl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallSslSshProfileSslInspectAll(i["inspect-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileSslClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileSslUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileSslUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileSslUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileSslExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileSslRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileSslClientCertRequest(i["client-cert-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileSslUnsupportedSsl(i["unsupported-ssl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileSslInvalidServerCert(i["invalid-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileSslUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileSslCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileSslCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileSslSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := i["cert-probe-failure"]; ok {
		result["cert_probe_failure"] = flattenFirewallSslSshProfileSslCertProbeFailure(i["cert-probe-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileSslMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSslInspectAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslInvalidServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslCertProbeFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileHttpsPorts(i["ports"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileHttpsStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "quic"
	if _, ok := i["quic"]; ok {
		result["quic"] = flattenFirewallSslSshProfileHttpsQuic(i["quic"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileHttpsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileHttpsClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileHttpsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileHttpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileHttpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileHttpsExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileHttpsRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileHttpsClientCertRequest(i["client-cert-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileHttpsUnsupportedSsl(i["unsupported-ssl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileHttpsInvalidServerCert(i["invalid-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileHttpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileHttpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileHttpsCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileHttpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := i["cert-probe-failure"]; ok {
		result["cert_probe_failure"] = flattenFirewallSslSshProfileHttpsCertProbeFailure(i["cert-probe-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileHttpsMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileHttpsPorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintflist2str(v)
}

func flattenFirewallSslSshProfileHttpsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsQuic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsCertProbeFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileHttpsMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileFtpsPorts(i["ports"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileFtpsStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileFtpsClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileFtpsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileFtpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileFtpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileFtpsExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileFtpsRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileFtpsClientCertRequest(i["client-cert-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileFtpsUnsupportedSsl(i["unsupported-ssl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileFtpsInvalidServerCert(i["invalid-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileFtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileFtpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileFtpsCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileFtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := i["min-allowed-ssl-version"]; ok {
		result["min_allowed_ssl_version"] = flattenFirewallSslSshProfileFtpsMinAllowedSslVersion(i["min-allowed-ssl-version"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileFtpsPorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintflist2str(v)
}

func flattenFirewallSslSshProfileFtpsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileFtpsMinAllowedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileImapsPorts(i["ports"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileImapsStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileImapsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileImapsClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileImapsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileImapsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileImapsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileImapsExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileImapsRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileImapsClientCertRequest(i["client-cert-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileImapsUnsupportedSsl(i["unsupported-ssl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileImapsInvalidServerCert(i["invalid-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileImapsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileImapsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileImapsCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileImapsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileImapsPorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintflist2str(v)
}

func flattenFirewallSslSshProfileImapsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileImapsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3S(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfilePop3SPorts(i["ports"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfilePop3SStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfilePop3SProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfilePop3SClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfilePop3SUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfilePop3SUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfilePop3SUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfilePop3SExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfilePop3SRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfilePop3SClientCertRequest(i["client-cert-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfilePop3SUnsupportedSsl(i["unsupported-ssl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfilePop3SInvalidServerCert(i["invalid-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfilePop3SUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfilePop3SCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfilePop3SCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfilePop3SSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfilePop3SPorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintflist2str(v)
}

func flattenFirewallSslSshProfilePop3SStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SInvalidServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfilePop3SSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileSmtpsPorts(i["ports"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSmtpsStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileSmtpsClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileSmtpsUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileSmtpsUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileSmtpsExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileSmtpsRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := i["client-cert-request"]; ok {
		result["client_cert_request"] = flattenFirewallSslSshProfileSmtpsClientCertRequest(i["client-cert-request"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := i["unsupported-ssl"]; ok {
		result["unsupported_ssl"] = flattenFirewallSslSshProfileSmtpsUnsupportedSsl(i["unsupported-ssl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := i["invalid-server-cert"]; ok {
		result["invalid_server_cert"] = flattenFirewallSslSshProfileSmtpsInvalidServerCert(i["invalid-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileSmtpsUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileSmtpsCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileSmtpsCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileSmtpsSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSmtpsPorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintflist2str(v)
}

func flattenFirewallSslSshProfileSmtpsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUnsupportedSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsInvalidServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSmtpsSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSsh(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := i["ports"]; ok {
		result["ports"] = flattenFirewallSslSshProfileSshPorts(i["ports"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileSshStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "inspect_all"
	if _, ok := i["inspect-all"]; ok {
		result["inspect_all"] = flattenFirewallSslSshProfileSshInspectAll(i["inspect-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileSshProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := i["unsupported-version"]; ok {
		result["unsupported_version"] = flattenFirewallSslSshProfileSshUnsupportedVersion(i["unsupported-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssh_policy_check"
	if _, ok := i["ssh-policy-check"]; ok {
		result["ssh_policy_check"] = flattenFirewallSslSshProfileSshSshPolicyCheck(i["ssh-policy-check"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := i["ssh-tun-policy-check"]; ok {
		result["ssh_tun_policy_check"] = flattenFirewallSslSshProfileSshSshTunPolicyCheck(i["ssh-tun-policy-check"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := i["ssh-algorithm"]; ok {
		result["ssh_algorithm"] = flattenFirewallSslSshProfileSshSshAlgorithm(i["ssh-algorithm"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileSshPorts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintflist2str(v)
}

func flattenFirewallSslSshProfileSshStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshInspectAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshUnsupportedVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshPolicyCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshTunPolicyCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSshSshAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDot(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenFirewallSslSshProfileDotStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "quic"
	if _, ok := i["quic"]; ok {
		result["quic"] = flattenFirewallSslSshProfileDotQuic(i["quic"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := i["proxy-after-tcp-handshake"]; ok {
		result["proxy_after_tcp_handshake"] = flattenFirewallSslSshProfileDotProxyAfterTcpHandshake(i["proxy-after-tcp-handshake"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "client_certificate"
	if _, ok := i["client-certificate"]; ok {
		result["client_certificate"] = flattenFirewallSslSshProfileDotClientCertificate(i["client-certificate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := i["unsupported-ssl-version"]; ok {
		result["unsupported_ssl_version"] = flattenFirewallSslSshProfileDotUnsupportedSslVersion(i["unsupported-ssl-version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := i["unsupported-ssl-cipher"]; ok {
		result["unsupported_ssl_cipher"] = flattenFirewallSslSshProfileDotUnsupportedSslCipher(i["unsupported-ssl-cipher"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := i["unsupported-ssl-negotiation"]; ok {
		result["unsupported_ssl_negotiation"] = flattenFirewallSslSshProfileDotUnsupportedSslNegotiation(i["unsupported-ssl-negotiation"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := i["expired-server-cert"]; ok {
		result["expired_server_cert"] = flattenFirewallSslSshProfileDotExpiredServerCert(i["expired-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := i["revoked-server-cert"]; ok {
		result["revoked_server_cert"] = flattenFirewallSslSshProfileDotRevokedServerCert(i["revoked-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := i["untrusted-server-cert"]; ok {
		result["untrusted_server_cert"] = flattenFirewallSslSshProfileDotUntrustedServerCert(i["untrusted-server-cert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := i["cert-validation-timeout"]; ok {
		result["cert_validation_timeout"] = flattenFirewallSslSshProfileDotCertValidationTimeout(i["cert-validation-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := i["cert-validation-failure"]; ok {
		result["cert_validation_failure"] = flattenFirewallSslSshProfileDotCertValidationFailure(i["cert-validation-failure"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := i["sni-server-cert-check"]; ok {
		result["sni_server_cert_check"] = flattenFirewallSslSshProfileDotSniServerCertCheck(i["sni-server-cert-check"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallSslSshProfileDotStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotQuic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotProxyAfterTcpHandshake(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUnsupportedSslNegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotExpiredServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotRevokedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotUntrustedServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotCertValidationTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotCertValidationFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileDotSniServerCertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileAllowlist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileBlockBlocklistedCertificates(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileWhitelist(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileBlockBlacklistedCertificates(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExempt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallSslSshProfileSslExemptId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenFirewallSslSshProfileSslExemptType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if cur_v, ok := i["fortiguard-category"]; ok {
			tmp["fortiguard_category"] = flattenFirewallSslSshProfileSslExemptFortiguardCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if cur_v, ok := i["address"]; ok {
			tmp["address"] = flattenFirewallSslSshProfileSslExemptAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if cur_v, ok := i["address6"]; ok {
			tmp["address6"] = flattenFirewallSslSshProfileSslExemptAddress6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if cur_v, ok := i["wildcard-fqdn"]; ok {
			tmp["wildcard_fqdn"] = flattenFirewallSslSshProfileSslExemptWildcardFqdn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if cur_v, ok := i["regex"]; ok {
			tmp["regex"] = flattenFirewallSslSshProfileSslExemptRegex(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSslSshProfileSslExemptId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptFortiguardCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptWildcardFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptRegex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileServerCertMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileUseSslServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileCaname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileUntrustedCaname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenFirewallSslSshProfileSslServerId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenFirewallSslSshProfileSslServerIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_certificate"
		if cur_v, ok := i["https-client-certificate"]; ok {
			tmp["https_client_certificate"] = flattenFirewallSslSshProfileSslServerHttpsClientCertificate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_certificate"
		if cur_v, ok := i["smtps-client-certificate"]; ok {
			tmp["smtps_client_certificate"] = flattenFirewallSslSshProfileSslServerSmtpsClientCertificate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_certificate"
		if cur_v, ok := i["pop3s-client-certificate"]; ok {
			tmp["pop3s_client_certificate"] = flattenFirewallSslSshProfileSslServerPop3SClientCertificate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_certificate"
		if cur_v, ok := i["imaps-client-certificate"]; ok {
			tmp["imaps_client_certificate"] = flattenFirewallSslSshProfileSslServerImapsClientCertificate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_certificate"
		if cur_v, ok := i["ftps-client-certificate"]; ok {
			tmp["ftps_client_certificate"] = flattenFirewallSslSshProfileSslServerFtpsClientCertificate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_certificate"
		if cur_v, ok := i["ssl-other-client-certificate"]; ok {
			tmp["ssl_other_client_certificate"] = flattenFirewallSslSshProfileSslServerSslOtherClientCertificate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_cert_request"
		if cur_v, ok := i["https-client-cert-request"]; ok {
			tmp["https_client_cert_request"] = flattenFirewallSslSshProfileSslServerHttpsClientCertRequest(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_cert_request"
		if cur_v, ok := i["smtps-client-cert-request"]; ok {
			tmp["smtps_client_cert_request"] = flattenFirewallSslSshProfileSslServerSmtpsClientCertRequest(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_cert_request"
		if cur_v, ok := i["pop3s-client-cert-request"]; ok {
			tmp["pop3s_client_cert_request"] = flattenFirewallSslSshProfileSslServerPop3SClientCertRequest(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_cert_request"
		if cur_v, ok := i["imaps-client-cert-request"]; ok {
			tmp["imaps_client_cert_request"] = flattenFirewallSslSshProfileSslServerImapsClientCertRequest(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_cert_request"
		if cur_v, ok := i["ftps-client-cert-request"]; ok {
			tmp["ftps_client_cert_request"] = flattenFirewallSslSshProfileSslServerFtpsClientCertRequest(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_cert_request"
		if cur_v, ok := i["ssl-other-client-cert-request"]; ok {
			tmp["ssl_other_client_cert_request"] = flattenFirewallSslSshProfileSslServerSslOtherClientCertRequest(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSslSshProfileSslServerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerHttpsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSmtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerPop3SClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerImapsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerFtpsClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSslOtherClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerHttpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSmtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerPop3SClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerImapsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerFtpsClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerSslOtherClientCertRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionIpRating(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslAnomalyLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslAnomaliesLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslExemptionsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslNegotiationLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslServerCertLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSslHandshakeLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileRpcOverHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileMapiOverHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSslSshProfileSupportedAlpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallSslSshProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenFirewallSslSshProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallSslSshProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o["ssl"], d, "ssl", sv)); err != nil {
			if !fortiAPIPatch(o["ssl"]) {
				return fmt.Errorf("Error reading ssl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl"); ok {
			if err = d.Set("ssl", flattenFirewallSslSshProfileSsl(o["ssl"], d, "ssl", sv)); err != nil {
				if !fortiAPIPatch(o["ssl"]) {
					return fmt.Errorf("Error reading ssl: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("https", flattenFirewallSslSshProfileHttps(o["https"], d, "https", sv)); err != nil {
			if !fortiAPIPatch(o["https"]) {
				return fmt.Errorf("Error reading https: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("https"); ok {
			if err = d.Set("https", flattenFirewallSslSshProfileHttps(o["https"], d, "https", sv)); err != nil {
				if !fortiAPIPatch(o["https"]) {
					return fmt.Errorf("Error reading https: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o["ftps"], d, "ftps", sv)); err != nil {
			if !fortiAPIPatch(o["ftps"]) {
				return fmt.Errorf("Error reading ftps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftps"); ok {
			if err = d.Set("ftps", flattenFirewallSslSshProfileFtps(o["ftps"], d, "ftps", sv)); err != nil {
				if !fortiAPIPatch(o["ftps"]) {
					return fmt.Errorf("Error reading ftps: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o["imaps"], d, "imaps", sv)); err != nil {
			if !fortiAPIPatch(o["imaps"]) {
				return fmt.Errorf("Error reading imaps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("imaps"); ok {
			if err = d.Set("imaps", flattenFirewallSslSshProfileImaps(o["imaps"], d, "imaps", sv)); err != nil {
				if !fortiAPIPatch(o["imaps"]) {
					return fmt.Errorf("Error reading imaps: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s", sv)); err != nil {
			if !fortiAPIPatch(o["pop3s"]) {
				return fmt.Errorf("Error reading pop3s: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pop3s"); ok {
			if err = d.Set("pop3s", flattenFirewallSslSshProfilePop3S(o["pop3s"], d, "pop3s", sv)); err != nil {
				if !fortiAPIPatch(o["pop3s"]) {
					return fmt.Errorf("Error reading pop3s: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o["smtps"], d, "smtps", sv)); err != nil {
			if !fortiAPIPatch(o["smtps"]) {
				return fmt.Errorf("Error reading smtps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("smtps"); ok {
			if err = d.Set("smtps", flattenFirewallSslSshProfileSmtps(o["smtps"], d, "smtps", sv)); err != nil {
				if !fortiAPIPatch(o["smtps"]) {
					return fmt.Errorf("Error reading smtps: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o["ssh"], d, "ssh", sv)); err != nil {
			if !fortiAPIPatch(o["ssh"]) {
				return fmt.Errorf("Error reading ssh: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssh"); ok {
			if err = d.Set("ssh", flattenFirewallSslSshProfileSsh(o["ssh"], d, "ssh", sv)); err != nil {
				if !fortiAPIPatch(o["ssh"]) {
					return fmt.Errorf("Error reading ssh: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dot", flattenFirewallSslSshProfileDot(o["dot"], d, "dot", sv)); err != nil {
			if !fortiAPIPatch(o["dot"]) {
				return fmt.Errorf("Error reading dot: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dot"); ok {
			if err = d.Set("dot", flattenFirewallSslSshProfileDot(o["dot"], d, "dot", sv)); err != nil {
				if !fortiAPIPatch(o["dot"]) {
					return fmt.Errorf("Error reading dot: %v", err)
				}
			}
		}
	}

	if err = d.Set("allowlist", flattenFirewallSslSshProfileAllowlist(o["allowlist"], d, "allowlist", sv)); err != nil {
		if !fortiAPIPatch(o["allowlist"]) {
			return fmt.Errorf("Error reading allowlist: %v", err)
		}
	}

	if err = d.Set("block_blocklisted_certificates", flattenFirewallSslSshProfileBlockBlocklistedCertificates(o["block-blocklisted-certificates"], d, "block_blocklisted_certificates", sv)); err != nil {
		if !fortiAPIPatch(o["block-blocklisted-certificates"]) {
			return fmt.Errorf("Error reading block_blocklisted_certificates: %v", err)
		}
	}

	if err = d.Set("whitelist", flattenFirewallSslSshProfileWhitelist(o["whitelist"], d, "whitelist", sv)); err != nil {
		if !fortiAPIPatch(o["whitelist"]) {
			return fmt.Errorf("Error reading whitelist: %v", err)
		}
	}

	if err = d.Set("block_blacklisted_certificates", flattenFirewallSslSshProfileBlockBlacklistedCertificates(o["block-blacklisted-certificates"], d, "block_blacklisted_certificates", sv)); err != nil {
		if !fortiAPIPatch(o["block-blacklisted-certificates"]) {
			return fmt.Errorf("Error reading block_blacklisted_certificates: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-exempt"]) {
				return fmt.Errorf("Error reading ssl_exempt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_exempt"); ok {
			if err = d.Set("ssl_exempt", flattenFirewallSslSshProfileSslExempt(o["ssl-exempt"], d, "ssl_exempt", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-exempt"]) {
					return fmt.Errorf("Error reading ssl_exempt: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_cert_mode", flattenFirewallSslSshProfileServerCertMode(o["server-cert-mode"], d, "server_cert_mode", sv)); err != nil {
		if !fortiAPIPatch(o["server-cert-mode"]) {
			return fmt.Errorf("Error reading server_cert_mode: %v", err)
		}
	}

	if err = d.Set("use_ssl_server", flattenFirewallSslSshProfileUseSslServer(o["use-ssl-server"], d, "use_ssl_server", sv)); err != nil {
		if !fortiAPIPatch(o["use-ssl-server"]) {
			return fmt.Errorf("Error reading use_ssl_server: %v", err)
		}
	}

	if err = d.Set("caname", flattenFirewallSslSshProfileCaname(o["caname"], d, "caname", sv)); err != nil {
		if !fortiAPIPatch(o["caname"]) {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("untrusted_caname", flattenFirewallSslSshProfileUntrustedCaname(o["untrusted-caname"], d, "untrusted_caname", sv)); err != nil {
		if !fortiAPIPatch(o["untrusted-caname"]) {
			return fmt.Errorf("Error reading untrusted_caname: %v", err)
		}
	}

	{
		v := flattenFirewallSslSshProfileServerCert(o["server-cert"], d, "server_cert", sv)
		vx := ""
		bstring := false
		new_version_map := map[string][]string{
			">=": []string{"7.0.0"},
		}
		if versionMatch, _ := checkVersionMatch(sv, new_version_map); versionMatch {
			l := v.([]interface{})
			if len(l) > 0 {
				for k, r := range l {
					i := r.(map[string]interface{})
					if _, ok := i["name"]; ok {
						if xv, ok := i["name"].(string); ok {
							vx += xv
							if k < len(l)-1 {
								vx += ", "
							}
						}
					}
				}
			}
			bstring = true
		}
		if bstring == true {
			if err = d.Set("server_cert", vx); err != nil {
				if !fortiAPIPatch(o["server-cert"]) {
					return fmt.Errorf("Error reading server_cert: %v", err)
				}
			}
		} else {
			if err = d.Set("server_cert", v); err != nil {
				if !fortiAPIPatch(o["server-cert"]) {
					return fmt.Errorf("Error reading server_cert: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-server"]) {
				return fmt.Errorf("Error reading ssl_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server"); ok {
			if err = d.Set("ssl_server", flattenFirewallSslSshProfileSslServer(o["ssl-server"], d, "ssl_server", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-server"]) {
					return fmt.Errorf("Error reading ssl_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_exemption_ip_rating", flattenFirewallSslSshProfileSslExemptionIpRating(o["ssl-exemption-ip-rating"], d, "ssl_exemption_ip_rating", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-exemption-ip-rating"]) {
			return fmt.Errorf("Error reading ssl_exemption_ip_rating: %v", err)
		}
	}

	if err = d.Set("ssl_anomaly_log", flattenFirewallSslSshProfileSslAnomalyLog(o["ssl-anomaly-log"], d, "ssl_anomaly_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-anomaly-log"]) {
			return fmt.Errorf("Error reading ssl_anomaly_log: %v", err)
		}
	}

	if err = d.Set("ssl_exemption_log", flattenFirewallSslSshProfileSslExemptionLog(o["ssl-exemption-log"], d, "ssl_exemption_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-exemption-log"]) {
			return fmt.Errorf("Error reading ssl_exemption_log: %v", err)
		}
	}

	if err = d.Set("ssl_anomalies_log", flattenFirewallSslSshProfileSslAnomaliesLog(o["ssl-anomalies-log"], d, "ssl_anomalies_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-anomalies-log"]) {
			return fmt.Errorf("Error reading ssl_anomalies_log: %v", err)
		}
	}

	if err = d.Set("ssl_exemptions_log", flattenFirewallSslSshProfileSslExemptionsLog(o["ssl-exemptions-log"], d, "ssl_exemptions_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-exemptions-log"]) {
			return fmt.Errorf("Error reading ssl_exemptions_log: %v", err)
		}
	}

	if err = d.Set("ssl_negotiation_log", flattenFirewallSslSshProfileSslNegotiationLog(o["ssl-negotiation-log"], d, "ssl_negotiation_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-negotiation-log"]) {
			return fmt.Errorf("Error reading ssl_negotiation_log: %v", err)
		}
	}

	if err = d.Set("ssl_server_cert_log", flattenFirewallSslSshProfileSslServerCertLog(o["ssl-server-cert-log"], d, "ssl_server_cert_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-cert-log"]) {
			return fmt.Errorf("Error reading ssl_server_cert_log: %v", err)
		}
	}

	if err = d.Set("ssl_handshake_log", flattenFirewallSslSshProfileSslHandshakeLog(o["ssl-handshake-log"], d, "ssl_handshake_log", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-handshake-log"]) {
			return fmt.Errorf("Error reading ssl_handshake_log: %v", err)
		}
	}

	if err = d.Set("rpc_over_https", flattenFirewallSslSshProfileRpcOverHttps(o["rpc-over-https"], d, "rpc_over_https", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-over-https"]) {
			return fmt.Errorf("Error reading rpc_over_https: %v", err)
		}
	}

	if err = d.Set("mapi_over_https", flattenFirewallSslSshProfileMapiOverHttps(o["mapi-over-https"], d, "mapi_over_https", sv)); err != nil {
		if !fortiAPIPatch(o["mapi-over-https"]) {
			return fmt.Errorf("Error reading mapi_over_https: %v", err)
		}
	}

	if err = d.Set("supported_alpn", flattenFirewallSslSshProfileSupportedAlpn(o["supported-alpn"], d, "supported_alpn", sv)); err != nil {
		if !fortiAPIPatch(o["supported-alpn"]) {
			return fmt.Errorf("Error reading supported_alpn: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSshProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallSslSshProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallSslSshProfileSslInspectAll(d, i["inspect_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfileSslClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileSslUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileSslUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileSslUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileSslExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileSslRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileSslClientCertRequest(d, i["client_cert_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileSslUnsupportedSsl(d, i["unsupported_ssl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileSslInvalidServerCert(d, i["invalid_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileSslUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileSslCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileSslCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileSslSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-probe-failure"], _ = expandFirewallSslSshProfileSslCertProbeFailure(d, i["cert_probe_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileSslMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileSslInspectAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslInvalidServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslCertProbeFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileHttpsPorts(d, i["ports"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileHttpsStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "quic"
	if _, ok := d.GetOk(pre_append); ok {
		result["quic"], _ = expandFirewallSslSshProfileHttpsQuic(d, i["quic"], pre_append, sv)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileHttpsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfileHttpsClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileHttpsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileHttpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileHttpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileHttpsExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileHttpsRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileHttpsClientCertRequest(d, i["client_cert_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileHttpsUnsupportedSsl(d, i["unsupported_ssl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileHttpsInvalidServerCert(d, i["invalid_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileHttpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileHttpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileHttpsCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileHttpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_probe_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-probe-failure"], _ = expandFirewallSslSshProfileHttpsCertProbeFailure(d, i["cert_probe_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileHttpsMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileHttpsPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsQuic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsCertProbeFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileHttpsMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileFtpsPorts(d, i["ports"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileFtpsStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfileFtpsClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileFtpsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileFtpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileFtpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileFtpsExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileFtpsRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileFtpsClientCertRequest(d, i["client_cert_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileFtpsUnsupportedSsl(d, i["unsupported_ssl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileFtpsInvalidServerCert(d, i["invalid_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileFtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileFtpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileFtpsCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileFtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}
	pre_append = pre + ".0." + "min_allowed_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["min-allowed-ssl-version"], _ = expandFirewallSslSshProfileFtpsMinAllowedSslVersion(d, i["min_allowed_ssl_version"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileFtpsPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileFtpsMinAllowedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileImapsPorts(d, i["ports"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileImapsStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileImapsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfileImapsClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileImapsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileImapsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileImapsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileImapsExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileImapsRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileImapsClientCertRequest(d, i["client_cert_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileImapsUnsupportedSsl(d, i["unsupported_ssl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileImapsInvalidServerCert(d, i["invalid_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileImapsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileImapsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileImapsCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileImapsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileImapsPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileImapsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3S(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfilePop3SPorts(d, i["ports"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfilePop3SStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfilePop3SProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfilePop3SClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfilePop3SUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfilePop3SUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfilePop3SUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfilePop3SExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfilePop3SRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfilePop3SClientCertRequest(d, i["client_cert_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfilePop3SUnsupportedSsl(d, i["unsupported_ssl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfilePop3SInvalidServerCert(d, i["invalid_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfilePop3SUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfilePop3SCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfilePop3SCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfilePop3SSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfilePop3SPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SInvalidServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfilePop3SSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileSmtpsPorts(d, i["ports"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileSmtpsStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfileSmtpsClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileSmtpsExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileSmtpsRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_cert_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-cert-request"], _ = expandFirewallSslSshProfileSmtpsClientCertRequest(d, i["client_cert_request"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl"], _ = expandFirewallSslSshProfileSmtpsUnsupportedSsl(d, i["unsupported_ssl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "invalid_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["invalid-server-cert"], _ = expandFirewallSslSshProfileSmtpsInvalidServerCert(d, i["invalid_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileSmtpsUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileSmtpsCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileSmtpsCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileSmtpsSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileSmtpsPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUnsupportedSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsInvalidServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSmtpsSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSsh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ports"
	if _, ok := d.GetOk(pre_append); ok {
		result["ports"], _ = expandFirewallSslSshProfileSshPorts(d, i["ports"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileSshStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "inspect_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["inspect-all"], _ = expandFirewallSslSshProfileSshInspectAll(d, i["inspect_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileSshProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-version"], _ = expandFirewallSslSshProfileSshUnsupportedVersion(d, i["unsupported_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssh_policy_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-policy-check"], _ = expandFirewallSslSshProfileSshSshPolicyCheck(d, i["ssh_policy_check"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssh_tun_policy_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-tun-policy-check"], _ = expandFirewallSslSshProfileSshSshTunPolicyCheck(d, i["ssh_tun_policy_check"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ssh_algorithm"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssh-algorithm"], _ = expandFirewallSslSshProfileSshSshAlgorithm(d, i["ssh_algorithm"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileSshPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshInspectAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshUnsupportedVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshPolicyCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshTunPolicyCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSshSshAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDot(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandFirewallSslSshProfileDotStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "quic"
	if _, ok := d.GetOk(pre_append); ok {
		result["quic"], _ = expandFirewallSslSshProfileDotQuic(d, i["quic"], pre_append, sv)
	}
	pre_append = pre + ".0." + "proxy_after_tcp_handshake"
	if _, ok := d.GetOk(pre_append); ok {
		result["proxy-after-tcp-handshake"], _ = expandFirewallSslSshProfileDotProxyAfterTcpHandshake(d, i["proxy_after_tcp_handshake"], pre_append, sv)
	}
	pre_append = pre + ".0." + "client_certificate"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-certificate"], _ = expandFirewallSslSshProfileDotClientCertificate(d, i["client_certificate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-version"], _ = expandFirewallSslSshProfileDotUnsupportedSslVersion(d, i["unsupported_ssl_version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_cipher"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-cipher"], _ = expandFirewallSslSshProfileDotUnsupportedSslCipher(d, i["unsupported_ssl_cipher"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unsupported_ssl_negotiation"
	if _, ok := d.GetOk(pre_append); ok {
		result["unsupported-ssl-negotiation"], _ = expandFirewallSslSshProfileDotUnsupportedSslNegotiation(d, i["unsupported_ssl_negotiation"], pre_append, sv)
	}
	pre_append = pre + ".0." + "expired_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["expired-server-cert"], _ = expandFirewallSslSshProfileDotExpiredServerCert(d, i["expired_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revoked_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["revoked-server-cert"], _ = expandFirewallSslSshProfileDotRevokedServerCert(d, i["revoked_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "untrusted_server_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["untrusted-server-cert"], _ = expandFirewallSslSshProfileDotUntrustedServerCert(d, i["untrusted_server_cert"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-timeout"], _ = expandFirewallSslSshProfileDotCertValidationTimeout(d, i["cert_validation_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "cert_validation_failure"
	if _, ok := d.GetOk(pre_append); ok {
		result["cert-validation-failure"], _ = expandFirewallSslSshProfileDotCertValidationFailure(d, i["cert_validation_failure"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sni_server_cert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["sni-server-cert-check"], _ = expandFirewallSslSshProfileDotSniServerCertCheck(d, i["sni_server_cert_check"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallSslSshProfileDotStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotQuic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotProxyAfterTcpHandshake(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUnsupportedSslNegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotExpiredServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotRevokedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotUntrustedServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotCertValidationTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotCertValidationFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileDotSniServerCertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileAllowlist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileBlockBlocklistedCertificates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileWhitelist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileBlockBlacklistedCertificates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFirewallSslSshProfileSslExemptId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandFirewallSslSshProfileSslExemptType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiguard_category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fortiguard-category"], _ = expandFirewallSslSshProfileSslExemptFortiguardCategory(d, i["fortiguard_category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandFirewallSslSshProfileSslExemptAddress(d, i["address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address6"], _ = expandFirewallSslSshProfileSslExemptAddress6(d, i["address6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wildcard_fqdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wildcard-fqdn"], _ = expandFirewallSslSshProfileSslExemptWildcardFqdn(d, i["wildcard_fqdn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regex"], _ = expandFirewallSslSshProfileSslExemptRegex(d, i["regex"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileSslExemptId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptFortiguardCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptWildcardFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptRegex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileServerCertMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileUseSslServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileCaname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileUntrustedCaname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandFirewallSslSshProfileSslServerId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFirewallSslSshProfileSslServerIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-client-certificate"], _ = expandFirewallSslSshProfileSslServerHttpsClientCertificate(d, i["https_client_certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["smtps-client-certificate"], _ = expandFirewallSslSshProfileSslServerSmtpsClientCertificate(d, i["smtps_client_certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pop3s-client-certificate"], _ = expandFirewallSslSshProfileSslServerPop3SClientCertificate(d, i["pop3s_client_certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imaps-client-certificate"], _ = expandFirewallSslSshProfileSslServerImapsClientCertificate(d, i["imaps_client_certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ftps-client-certificate"], _ = expandFirewallSslSshProfileSslServerFtpsClientCertificate(d, i["ftps_client_certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-other-client-certificate"], _ = expandFirewallSslSshProfileSslServerSslOtherClientCertificate(d, i["ssl_other_client_certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-client-cert-request"], _ = expandFirewallSslSshProfileSslServerHttpsClientCertRequest(d, i["https_client_cert_request"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "smtps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["smtps-client-cert-request"], _ = expandFirewallSslSshProfileSslServerSmtpsClientCertRequest(d, i["smtps_client_cert_request"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pop3s_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pop3s-client-cert-request"], _ = expandFirewallSslSshProfileSslServerPop3SClientCertRequest(d, i["pop3s_client_cert_request"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "imaps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["imaps-client-cert-request"], _ = expandFirewallSslSshProfileSslServerImapsClientCertRequest(d, i["imaps_client_cert_request"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftps_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ftps-client-cert-request"], _ = expandFirewallSslSshProfileSslServerFtpsClientCertRequest(d, i["ftps_client_cert_request"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_other_client_cert_request"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-other-client-cert-request"], _ = expandFirewallSslSshProfileSslServerSslOtherClientCertRequest(d, i["ssl_other_client_cert_request"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSslSshProfileSslServerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerHttpsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSmtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerPop3SClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerImapsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerFtpsClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSslOtherClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerHttpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSmtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerPop3SClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerImapsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerFtpsClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerSslOtherClientCertRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionIpRating(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslAnomalyLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslAnomaliesLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslExemptionsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslNegotiationLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslServerCertLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSslHandshakeLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileRpcOverHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileMapiOverHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSshProfileSupportedAlpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSshProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallSslSshProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallSslSshProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandFirewallSslSshProfileSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("https"); ok {
		t, err := expandFirewallSslSshProfileHttps(d, v, "https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https"] = t
		}
	}

	if v, ok := d.GetOk("ftps"); ok {
		t, err := expandFirewallSslSshProfileFtps(d, v, "ftps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftps"] = t
		}
	}

	if v, ok := d.GetOk("imaps"); ok {
		t, err := expandFirewallSslSshProfileImaps(d, v, "imaps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imaps"] = t
		}
	}

	if v, ok := d.GetOk("pop3s"); ok {
		t, err := expandFirewallSslSshProfilePop3S(d, v, "pop3s", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3s"] = t
		}
	}

	if v, ok := d.GetOk("smtps"); ok {
		t, err := expandFirewallSslSshProfileSmtps(d, v, "smtps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtps"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandFirewallSslSshProfileSsh(d, v, "ssh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("dot"); ok {
		t, err := expandFirewallSslSshProfileDot(d, v, "dot", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dot"] = t
		}
	}

	if v, ok := d.GetOk("allowlist"); ok {
		t, err := expandFirewallSslSshProfileAllowlist(d, v, "allowlist", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowlist"] = t
		}
	}

	if v, ok := d.GetOk("block_blocklisted_certificates"); ok {
		t, err := expandFirewallSslSshProfileBlockBlocklistedCertificates(d, v, "block_blocklisted_certificates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-blocklisted-certificates"] = t
		}
	}

	if v, ok := d.GetOk("whitelist"); ok {
		t, err := expandFirewallSslSshProfileWhitelist(d, v, "whitelist", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["whitelist"] = t
		}
	}

	if v, ok := d.GetOk("block_blacklisted_certificates"); ok {
		t, err := expandFirewallSslSshProfileBlockBlacklistedCertificates(d, v, "block_blacklisted_certificates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-blacklisted-certificates"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exempt"); ok || d.HasChange("ssl_exempt") {
		t, err := expandFirewallSslSshProfileSslExempt(d, v, "ssl_exempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exempt"] = t
		}
	}

	if v, ok := d.GetOk("server_cert_mode"); ok {
		t, err := expandFirewallSslSshProfileServerCertMode(d, v, "server_cert_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert-mode"] = t
		}
	}

	if v, ok := d.GetOk("use_ssl_server"); ok {
		t, err := expandFirewallSslSshProfileUseSslServer(d, v, "use_ssl_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("caname"); ok {
		t, err := expandFirewallSslSshProfileCaname(d, v, "caname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caname"] = t
		}
	}

	if v, ok := d.GetOk("untrusted_caname"); ok {
		t, err := expandFirewallSslSshProfileUntrustedCaname(d, v, "untrusted_caname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untrusted-caname"] = t
		}
	}

	if v, ok := d.GetOk("server_cert"); ok {
		t, err := expandFirewallSslSshProfileServerCert(d, v, "server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			new_version_map := map[string][]string{
				">=": []string{"7.0.0"},
			}
			if versionMatch, _ := checkVersionMatch(sv, new_version_map); versionMatch {
				vx := fmt.Sprintf("%v", t)
				vxx := strings.Split(vx, ", ")

				tmps := make([]map[string]interface{}, 0, len(vxx))

				for _, xv := range vxx {
					xtmp := make(map[string]interface{})
					xtmp["name"] = xv

					tmps = append(tmps, xtmp)
				}
				obj["server-cert"] = tmps
			} else {
				obj["server-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_server"); ok || d.HasChange("ssl_server") {
		t, err := expandFirewallSslSshProfileSslServer(d, v, "ssl_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemption_ip_rating"); ok {
		t, err := expandFirewallSslSshProfileSslExemptionIpRating(d, v, "ssl_exemption_ip_rating", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemption-ip-rating"] = t
		}
	}

	if v, ok := d.GetOk("ssl_anomaly_log"); ok {
		t, err := expandFirewallSslSshProfileSslAnomalyLog(d, v, "ssl_anomaly_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-anomaly-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemption_log"); ok {
		t, err := expandFirewallSslSshProfileSslExemptionLog(d, v, "ssl_exemption_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemption-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_anomalies_log"); ok {
		t, err := expandFirewallSslSshProfileSslAnomaliesLog(d, v, "ssl_anomalies_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-anomalies-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_exemptions_log"); ok {
		t, err := expandFirewallSslSshProfileSslExemptionsLog(d, v, "ssl_exemptions_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-exemptions-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_negotiation_log"); ok {
		t, err := expandFirewallSslSshProfileSslNegotiationLog(d, v, "ssl_negotiation_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-negotiation-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cert_log"); ok {
		t, err := expandFirewallSslSshProfileSslServerCertLog(d, v, "ssl_server_cert_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cert-log"] = t
		}
	}

	if v, ok := d.GetOk("ssl_handshake_log"); ok {
		t, err := expandFirewallSslSshProfileSslHandshakeLog(d, v, "ssl_handshake_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-handshake-log"] = t
		}
	}

	if v, ok := d.GetOk("rpc_over_https"); ok {
		t, err := expandFirewallSslSshProfileRpcOverHttps(d, v, "rpc_over_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-over-https"] = t
		}
	}

	if v, ok := d.GetOk("mapi_over_https"); ok {
		t, err := expandFirewallSslSshProfileMapiOverHttps(d, v, "mapi_over_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi-over-https"] = t
		}
	}

	if v, ok := d.GetOk("supported_alpn"); ok {
		t, err := expandFirewallSslSshProfileSupportedAlpn(d, v, "supported_alpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supported-alpn"] = t
		}
	}

	return &obj, nil
}
