// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SSL VPN.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceVpnSslSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVpnSslSettingsRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reqclientcert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_peer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_max_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_min_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tlsv1_0": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tlsv1_1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tlsv1_2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tlsv1_3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"banned_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ciphersuite": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_insert_empty_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"x_content_type_options": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_two_factor_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unsafe_legacy_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"servercert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"login_attempt_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"login_block_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"login_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dtls_hello_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dtls_heartbeat_idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dtls_heartbeat_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dtls_heartbeat_fail_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tunnel_ip_pools": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tunnel_ipv6_pools": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_source_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url_obscuration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_compression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_only_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"deflate_compression_level": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"deflate_min_data_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port_precedence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_tunnel_static_route": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"header_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_address_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_address6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_address6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_portal": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_rule": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"source_interface": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"source_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"source_address_negate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_address6": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"source_address6_negate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"users": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"portal": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"realm": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"client_cert": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_peer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"browser_language_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dtls_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dtls_max_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dtls_min_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"check_referer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_request_header_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"http_request_body_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"auth_session_check_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_connect_without_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_user_session_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"transform_backward_slashes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"encode_2f_sequence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"encrypt_and_store_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_sigalgs": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dual_stack_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_addr_assigned_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"saml_redirect_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"web_mode_snat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ztna_trusted_client": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVpnSslSettingsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "VpnSslSettings"

	o, err := c.ReadVpnSslSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing VpnSslSettings: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectVpnSslSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error describing VpnSslSettings from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenVpnSslSettingsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsReqclientcert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsUserPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSslMaxProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSslMinProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTlsv10(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTlsv11(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTlsv12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTlsv13(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsBannedCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsCiphersuite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSslInsertEmptyFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHttpsRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsXContentTypeOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsForceTwoFactorAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsUnsafeLegacyRenegotiation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsServercert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsLoginAttemptLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsLoginBlockTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsLoginTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsHelloTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsHeartbeatIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsHeartbeatFailCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTunnelIpPools(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsTunnelIpPoolsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsTunnelIpPoolsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTunnelIpv6Pools(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsTunnelIpv6PoolsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsTunnelIpv6PoolsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDnsSuffix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsWinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsWinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsIpv6WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsIpv6WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsRouteSourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsUrlObscuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHttpCompression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHttpOnlyCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDeflateCompressionLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDeflateMinDataSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsPortPrecedence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAutoTunnelStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSourceInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsSourceInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSourceAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsSourceAddressName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsSourceAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSourceAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsSourceAddress6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDefaultPortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_interface"
		if _, ok := i["source-interface"]; ok {
			tmp["source_interface"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceInterface(i["source-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := i["source-address"]; ok {
			tmp["source_address"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress(i["source-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address_negate"
		if _, ok := i["source-address-negate"]; ok {
			tmp["source_address_negate"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddressNegate(i["source-address-negate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6"
		if _, ok := i["source-address6"]; ok {
			tmp["source_address6"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress6(i["source-address6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6_negate"
		if _, ok := i["source-address6-negate"]; ok {
			tmp["source_address6_negate"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate(i["source-address6-negate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {
			tmp["users"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleUsers(i["users"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {
			tmp["groups"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleGroups(i["groups"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal"
		if _, ok := i["portal"]; ok {
			tmp["portal"] = dataSourceFlattenVpnSslSettingsAuthenticationRulePortal(i["portal"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := i["realm"]; ok {
			tmp["realm"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleRealm(i["realm"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {
			tmp["client_cert"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleClientCert(i["client-cert"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_peer"
		if _, ok := i["user-peer"]; ok {
			tmp["user_peer"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleUserPeer(i["user-peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {
			tmp["cipher"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleCipher(i["cipher"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := i["auth"]; ok {
			tmp["auth"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleAuth(i["auth"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddressName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenVpnSslSettingsAuthenticationRuleGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsBrowserLanguageDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsMaxProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDtlsMinProtoVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsCheckReferer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHttpRequestHeaderTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHttpRequestBodyTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsAuthSessionCheckSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTunnelConnectWithoutReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTunnelUserSessionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTransformBackwardSlashes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsEncode2FSequence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsEncryptAndStorePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsClientSigalgs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsDualStackMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsTunnelAddrAssignedMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsSamlRedirectPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsWebModeSnat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsZtnaTrustedClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenVpnSslSettingsServerHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectVpnSslSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenVpnSslSettingsStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("reqclientcert", dataSourceFlattenVpnSslSettingsReqclientcert(o["reqclientcert"], d, "reqclientcert")); err != nil {
		if !fortiAPIPatch(o["reqclientcert"]) {
			return fmt.Errorf("Error reading reqclientcert: %v", err)
		}
	}

	if err = d.Set("user_peer", dataSourceFlattenVpnSslSettingsUserPeer(o["user-peer"], d, "user_peer")); err != nil {
		if !fortiAPIPatch(o["user-peer"]) {
			return fmt.Errorf("Error reading user_peer: %v", err)
		}
	}

	if err = d.Set("ssl_max_proto_ver", dataSourceFlattenVpnSslSettingsSslMaxProtoVer(o["ssl-max-proto-ver"], d, "ssl_max_proto_ver")); err != nil {
		if !fortiAPIPatch(o["ssl-max-proto-ver"]) {
			return fmt.Errorf("Error reading ssl_max_proto_ver: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_ver", dataSourceFlattenVpnSslSettingsSslMinProtoVer(o["ssl-min-proto-ver"], d, "ssl_min_proto_ver")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-ver"]) {
			return fmt.Errorf("Error reading ssl_min_proto_ver: %v", err)
		}
	}

	if err = d.Set("tlsv1_0", dataSourceFlattenVpnSslSettingsTlsv10(o["tlsv1-0"], d, "tlsv1_0")); err != nil {
		if !fortiAPIPatch(o["tlsv1-0"]) {
			return fmt.Errorf("Error reading tlsv1_0: %v", err)
		}
	}

	if err = d.Set("tlsv1_1", dataSourceFlattenVpnSslSettingsTlsv11(o["tlsv1-1"], d, "tlsv1_1")); err != nil {
		if !fortiAPIPatch(o["tlsv1-1"]) {
			return fmt.Errorf("Error reading tlsv1_1: %v", err)
		}
	}

	if err = d.Set("tlsv1_2", dataSourceFlattenVpnSslSettingsTlsv12(o["tlsv1-2"], d, "tlsv1_2")); err != nil {
		if !fortiAPIPatch(o["tlsv1-2"]) {
			return fmt.Errorf("Error reading tlsv1_2: %v", err)
		}
	}

	if err = d.Set("tlsv1_3", dataSourceFlattenVpnSslSettingsTlsv13(o["tlsv1-3"], d, "tlsv1_3")); err != nil {
		if !fortiAPIPatch(o["tlsv1-3"]) {
			return fmt.Errorf("Error reading tlsv1_3: %v", err)
		}
	}

	if err = d.Set("banned_cipher", dataSourceFlattenVpnSslSettingsBannedCipher(o["banned-cipher"], d, "banned_cipher")); err != nil {
		if !fortiAPIPatch(o["banned-cipher"]) {
			return fmt.Errorf("Error reading banned_cipher: %v", err)
		}
	}

	if err = d.Set("ciphersuite", dataSourceFlattenVpnSslSettingsCiphersuite(o["ciphersuite"], d, "ciphersuite")); err != nil {
		if !fortiAPIPatch(o["ciphersuite"]) {
			return fmt.Errorf("Error reading ciphersuite: %v", err)
		}
	}

	if err = d.Set("ssl_insert_empty_fragment", dataSourceFlattenVpnSslSettingsSslInsertEmptyFragment(o["ssl-insert-empty-fragment"], d, "ssl_insert_empty_fragment")); err != nil {
		if !fortiAPIPatch(o["ssl-insert-empty-fragment"]) {
			return fmt.Errorf("Error reading ssl_insert_empty_fragment: %v", err)
		}
	}

	if err = d.Set("https_redirect", dataSourceFlattenVpnSslSettingsHttpsRedirect(o["https-redirect"], d, "https_redirect")); err != nil {
		if !fortiAPIPatch(o["https-redirect"]) {
			return fmt.Errorf("Error reading https_redirect: %v", err)
		}
	}

	if err = d.Set("x_content_type_options", dataSourceFlattenVpnSslSettingsXContentTypeOptions(o["x-content-type-options"], d, "x_content_type_options")); err != nil {
		if !fortiAPIPatch(o["x-content-type-options"]) {
			return fmt.Errorf("Error reading x_content_type_options: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", dataSourceFlattenVpnSslSettingsSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation")); err != nil {
		if !fortiAPIPatch(o["ssl-client-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("force_two_factor_auth", dataSourceFlattenVpnSslSettingsForceTwoFactorAuth(o["force-two-factor-auth"], d, "force_two_factor_auth")); err != nil {
		if !fortiAPIPatch(o["force-two-factor-auth"]) {
			return fmt.Errorf("Error reading force_two_factor_auth: %v", err)
		}
	}

	if err = d.Set("unsafe_legacy_renegotiation", dataSourceFlattenVpnSslSettingsUnsafeLegacyRenegotiation(o["unsafe-legacy-renegotiation"], d, "unsafe_legacy_renegotiation")); err != nil {
		if !fortiAPIPatch(o["unsafe-legacy-renegotiation"]) {
			return fmt.Errorf("Error reading unsafe_legacy_renegotiation: %v", err)
		}
	}

	if err = d.Set("servercert", dataSourceFlattenVpnSslSettingsServercert(o["servercert"], d, "servercert")); err != nil {
		if !fortiAPIPatch(o["servercert"]) {
			return fmt.Errorf("Error reading servercert: %v", err)
		}
	}

	if err = d.Set("algorithm", dataSourceFlattenVpnSslSettingsAlgorithm(o["algorithm"], d, "algorithm")); err != nil {
		if !fortiAPIPatch(o["algorithm"]) {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("idle_timeout", dataSourceFlattenVpnSslSettingsIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("auth_timeout", dataSourceFlattenVpnSslSettingsAuthTimeout(o["auth-timeout"], d, "auth_timeout")); err != nil {
		if !fortiAPIPatch(o["auth-timeout"]) {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if err = d.Set("login_attempt_limit", dataSourceFlattenVpnSslSettingsLoginAttemptLimit(o["login-attempt-limit"], d, "login_attempt_limit")); err != nil {
		if !fortiAPIPatch(o["login-attempt-limit"]) {
			return fmt.Errorf("Error reading login_attempt_limit: %v", err)
		}
	}

	if err = d.Set("login_block_time", dataSourceFlattenVpnSslSettingsLoginBlockTime(o["login-block-time"], d, "login_block_time")); err != nil {
		if !fortiAPIPatch(o["login-block-time"]) {
			return fmt.Errorf("Error reading login_block_time: %v", err)
		}
	}

	if err = d.Set("login_timeout", dataSourceFlattenVpnSslSettingsLoginTimeout(o["login-timeout"], d, "login_timeout")); err != nil {
		if !fortiAPIPatch(o["login-timeout"]) {
			return fmt.Errorf("Error reading login_timeout: %v", err)
		}
	}

	if err = d.Set("dtls_hello_timeout", dataSourceFlattenVpnSslSettingsDtlsHelloTimeout(o["dtls-hello-timeout"], d, "dtls_hello_timeout")); err != nil {
		if !fortiAPIPatch(o["dtls-hello-timeout"]) {
			return fmt.Errorf("Error reading dtls_hello_timeout: %v", err)
		}
	}

	if err = d.Set("dtls_heartbeat_idle_timeout", dataSourceFlattenVpnSslSettingsDtlsHeartbeatIdleTimeout(o["dtls-heartbeat-idle-timeout"], d, "dtls_heartbeat_idle_timeout")); err != nil {
		if !fortiAPIPatch(o["dtls-heartbeat-idle-timeout"]) {
			return fmt.Errorf("Error reading dtls_heartbeat_idle_timeout: %v", err)
		}
	}

	if err = d.Set("dtls_heartbeat_interval", dataSourceFlattenVpnSslSettingsDtlsHeartbeatInterval(o["dtls-heartbeat-interval"], d, "dtls_heartbeat_interval")); err != nil {
		if !fortiAPIPatch(o["dtls-heartbeat-interval"]) {
			return fmt.Errorf("Error reading dtls_heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("dtls_heartbeat_fail_count", dataSourceFlattenVpnSslSettingsDtlsHeartbeatFailCount(o["dtls-heartbeat-fail-count"], d, "dtls_heartbeat_fail_count")); err != nil {
		if !fortiAPIPatch(o["dtls-heartbeat-fail-count"]) {
			return fmt.Errorf("Error reading dtls_heartbeat_fail_count: %v", err)
		}
	}

	if err = d.Set("tunnel_ip_pools", dataSourceFlattenVpnSslSettingsTunnelIpPools(o["tunnel-ip-pools"], d, "tunnel_ip_pools")); err != nil {
		if !fortiAPIPatch(o["tunnel-ip-pools"]) {
			return fmt.Errorf("Error reading tunnel_ip_pools: %v", err)
		}
	}

	if err = d.Set("tunnel_ipv6_pools", dataSourceFlattenVpnSslSettingsTunnelIpv6Pools(o["tunnel-ipv6-pools"], d, "tunnel_ipv6_pools")); err != nil {
		if !fortiAPIPatch(o["tunnel-ipv6-pools"]) {
			return fmt.Errorf("Error reading tunnel_ipv6_pools: %v", err)
		}
	}

	if err = d.Set("dns_suffix", dataSourceFlattenVpnSslSettingsDnsSuffix(o["dns-suffix"], d, "dns_suffix")); err != nil {
		if !fortiAPIPatch(o["dns-suffix"]) {
			return fmt.Errorf("Error reading dns_suffix: %v", err)
		}
	}

	if err = d.Set("dns_server1", dataSourceFlattenVpnSslSettingsDnsServer1(o["dns-server1"], d, "dns_server1")); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", dataSourceFlattenVpnSslSettingsDnsServer2(o["dns-server2"], d, "dns_server2")); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("wins_server1", dataSourceFlattenVpnSslSettingsWinsServer1(o["wins-server1"], d, "wins_server1")); err != nil {
		if !fortiAPIPatch(o["wins-server1"]) {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", dataSourceFlattenVpnSslSettingsWinsServer2(o["wins-server2"], d, "wins_server2")); err != nil {
		if !fortiAPIPatch(o["wins-server2"]) {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", dataSourceFlattenVpnSslSettingsIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server1"]) {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", dataSourceFlattenVpnSslSettingsIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server2"]) {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server1", dataSourceFlattenVpnSslSettingsIpv6WinsServer1(o["ipv6-wins-server1"], d, "ipv6_wins_server1")); err != nil {
		if !fortiAPIPatch(o["ipv6-wins-server1"]) {
			return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server2", dataSourceFlattenVpnSslSettingsIpv6WinsServer2(o["ipv6-wins-server2"], d, "ipv6_wins_server2")); err != nil {
		if !fortiAPIPatch(o["ipv6-wins-server2"]) {
			return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
		}
	}

	if err = d.Set("route_source_interface", dataSourceFlattenVpnSslSettingsRouteSourceInterface(o["route-source-interface"], d, "route_source_interface")); err != nil {
		if !fortiAPIPatch(o["route-source-interface"]) {
			return fmt.Errorf("Error reading route_source_interface: %v", err)
		}
	}

	if err = d.Set("url_obscuration", dataSourceFlattenVpnSslSettingsUrlObscuration(o["url-obscuration"], d, "url_obscuration")); err != nil {
		if !fortiAPIPatch(o["url-obscuration"]) {
			return fmt.Errorf("Error reading url_obscuration: %v", err)
		}
	}

	if err = d.Set("http_compression", dataSourceFlattenVpnSslSettingsHttpCompression(o["http-compression"], d, "http_compression")); err != nil {
		if !fortiAPIPatch(o["http-compression"]) {
			return fmt.Errorf("Error reading http_compression: %v", err)
		}
	}

	if err = d.Set("http_only_cookie", dataSourceFlattenVpnSslSettingsHttpOnlyCookie(o["http-only-cookie"], d, "http_only_cookie")); err != nil {
		if !fortiAPIPatch(o["http-only-cookie"]) {
			return fmt.Errorf("Error reading http_only_cookie: %v", err)
		}
	}

	if err = d.Set("deflate_compression_level", dataSourceFlattenVpnSslSettingsDeflateCompressionLevel(o["deflate-compression-level"], d, "deflate_compression_level")); err != nil {
		if !fortiAPIPatch(o["deflate-compression-level"]) {
			return fmt.Errorf("Error reading deflate_compression_level: %v", err)
		}
	}

	if err = d.Set("deflate_min_data_size", dataSourceFlattenVpnSslSettingsDeflateMinDataSize(o["deflate-min-data-size"], d, "deflate_min_data_size")); err != nil {
		if !fortiAPIPatch(o["deflate-min-data-size"]) {
			return fmt.Errorf("Error reading deflate_min_data_size: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenVpnSslSettingsPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port_precedence", dataSourceFlattenVpnSslSettingsPortPrecedence(o["port-precedence"], d, "port_precedence")); err != nil {
		if !fortiAPIPatch(o["port-precedence"]) {
			return fmt.Errorf("Error reading port_precedence: %v", err)
		}
	}

	if err = d.Set("auto_tunnel_static_route", dataSourceFlattenVpnSslSettingsAutoTunnelStaticRoute(o["auto-tunnel-static-route"], d, "auto_tunnel_static_route")); err != nil {
		if !fortiAPIPatch(o["auto-tunnel-static-route"]) {
			return fmt.Errorf("Error reading auto_tunnel_static_route: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", dataSourceFlattenVpnSslSettingsHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for")); err != nil {
		if !fortiAPIPatch(o["header-x-forwarded-for"]) {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("source_interface", dataSourceFlattenVpnSslSettingsSourceInterface(o["source-interface"], d, "source_interface")); err != nil {
		if !fortiAPIPatch(o["source-interface"]) {
			return fmt.Errorf("Error reading source_interface: %v", err)
		}
	}

	if err = d.Set("source_address", dataSourceFlattenVpnSslSettingsSourceAddress(o["source-address"], d, "source_address")); err != nil {
		if !fortiAPIPatch(o["source-address"]) {
			return fmt.Errorf("Error reading source_address: %v", err)
		}
	}

	if err = d.Set("source_address_negate", dataSourceFlattenVpnSslSettingsSourceAddressNegate(o["source-address-negate"], d, "source_address_negate")); err != nil {
		if !fortiAPIPatch(o["source-address-negate"]) {
			return fmt.Errorf("Error reading source_address_negate: %v", err)
		}
	}

	if err = d.Set("source_address6", dataSourceFlattenVpnSslSettingsSourceAddress6(o["source-address6"], d, "source_address6")); err != nil {
		if !fortiAPIPatch(o["source-address6"]) {
			return fmt.Errorf("Error reading source_address6: %v", err)
		}
	}

	if err = d.Set("source_address6_negate", dataSourceFlattenVpnSslSettingsSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate")); err != nil {
		if !fortiAPIPatch(o["source-address6-negate"]) {
			return fmt.Errorf("Error reading source_address6_negate: %v", err)
		}
	}

	if err = d.Set("default_portal", dataSourceFlattenVpnSslSettingsDefaultPortal(o["default-portal"], d, "default_portal")); err != nil {
		if !fortiAPIPatch(o["default-portal"]) {
			return fmt.Errorf("Error reading default_portal: %v", err)
		}
	}

	if err = d.Set("authentication_rule", dataSourceFlattenVpnSslSettingsAuthenticationRule(o["authentication-rule"], d, "authentication_rule")); err != nil {
		if !fortiAPIPatch(o["authentication-rule"]) {
			return fmt.Errorf("Error reading authentication_rule: %v", err)
		}
	}

	if err = d.Set("browser_language_detection", dataSourceFlattenVpnSslSettingsBrowserLanguageDetection(o["browser-language-detection"], d, "browser_language_detection")); err != nil {
		if !fortiAPIPatch(o["browser-language-detection"]) {
			return fmt.Errorf("Error reading browser_language_detection: %v", err)
		}
	}

	if err = d.Set("dtls_tunnel", dataSourceFlattenVpnSslSettingsDtlsTunnel(o["dtls-tunnel"], d, "dtls_tunnel")); err != nil {
		if !fortiAPIPatch(o["dtls-tunnel"]) {
			return fmt.Errorf("Error reading dtls_tunnel: %v", err)
		}
	}

	if err = d.Set("dtls_max_proto_ver", dataSourceFlattenVpnSslSettingsDtlsMaxProtoVer(o["dtls-max-proto-ver"], d, "dtls_max_proto_ver")); err != nil {
		if !fortiAPIPatch(o["dtls-max-proto-ver"]) {
			return fmt.Errorf("Error reading dtls_max_proto_ver: %v", err)
		}
	}

	if err = d.Set("dtls_min_proto_ver", dataSourceFlattenVpnSslSettingsDtlsMinProtoVer(o["dtls-min-proto-ver"], d, "dtls_min_proto_ver")); err != nil {
		if !fortiAPIPatch(o["dtls-min-proto-ver"]) {
			return fmt.Errorf("Error reading dtls_min_proto_ver: %v", err)
		}
	}

	if err = d.Set("check_referer", dataSourceFlattenVpnSslSettingsCheckReferer(o["check-referer"], d, "check_referer")); err != nil {
		if !fortiAPIPatch(o["check-referer"]) {
			return fmt.Errorf("Error reading check_referer: %v", err)
		}
	}

	if err = d.Set("http_request_header_timeout", dataSourceFlattenVpnSslSettingsHttpRequestHeaderTimeout(o["http-request-header-timeout"], d, "http_request_header_timeout")); err != nil {
		if !fortiAPIPatch(o["http-request-header-timeout"]) {
			return fmt.Errorf("Error reading http_request_header_timeout: %v", err)
		}
	}

	if err = d.Set("http_request_body_timeout", dataSourceFlattenVpnSslSettingsHttpRequestBodyTimeout(o["http-request-body-timeout"], d, "http_request_body_timeout")); err != nil {
		if !fortiAPIPatch(o["http-request-body-timeout"]) {
			return fmt.Errorf("Error reading http_request_body_timeout: %v", err)
		}
	}

	if err = d.Set("auth_session_check_source_ip", dataSourceFlattenVpnSslSettingsAuthSessionCheckSourceIp(o["auth-session-check-source-ip"], d, "auth_session_check_source_ip")); err != nil {
		if !fortiAPIPatch(o["auth-session-check-source-ip"]) {
			return fmt.Errorf("Error reading auth_session_check_source_ip: %v", err)
		}
	}

	if err = d.Set("tunnel_connect_without_reauth", dataSourceFlattenVpnSslSettingsTunnelConnectWithoutReauth(o["tunnel-connect-without-reauth"], d, "tunnel_connect_without_reauth")); err != nil {
		if !fortiAPIPatch(o["tunnel-connect-without-reauth"]) {
			return fmt.Errorf("Error reading tunnel_connect_without_reauth: %v", err)
		}
	}

	if err = d.Set("tunnel_user_session_timeout", dataSourceFlattenVpnSslSettingsTunnelUserSessionTimeout(o["tunnel-user-session-timeout"], d, "tunnel_user_session_timeout")); err != nil {
		if !fortiAPIPatch(o["tunnel-user-session-timeout"]) {
			return fmt.Errorf("Error reading tunnel_user_session_timeout: %v", err)
		}
	}

	if err = d.Set("hsts_include_subdomains", dataSourceFlattenVpnSslSettingsHstsIncludeSubdomains(o["hsts-include-subdomains"], d, "hsts_include_subdomains")); err != nil {
		if !fortiAPIPatch(o["hsts-include-subdomains"]) {
			return fmt.Errorf("Error reading hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("transform_backward_slashes", dataSourceFlattenVpnSslSettingsTransformBackwardSlashes(o["transform-backward-slashes"], d, "transform_backward_slashes")); err != nil {
		if !fortiAPIPatch(o["transform-backward-slashes"]) {
			return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
		}
	}

	if err = d.Set("encode_2f_sequence", dataSourceFlattenVpnSslSettingsEncode2FSequence(o["encode-2f-sequence"], d, "encode_2f_sequence")); err != nil {
		if !fortiAPIPatch(o["encode-2f-sequence"]) {
			return fmt.Errorf("Error reading encode_2f_sequence: %v", err)
		}
	}

	if err = d.Set("encrypt_and_store_password", dataSourceFlattenVpnSslSettingsEncryptAndStorePassword(o["encrypt-and-store-password"], d, "encrypt_and_store_password")); err != nil {
		if !fortiAPIPatch(o["encrypt-and-store-password"]) {
			return fmt.Errorf("Error reading encrypt_and_store_password: %v", err)
		}
	}

	if err = d.Set("client_sigalgs", dataSourceFlattenVpnSslSettingsClientSigalgs(o["client-sigalgs"], d, "client_sigalgs")); err != nil {
		if !fortiAPIPatch(o["client-sigalgs"]) {
			return fmt.Errorf("Error reading client_sigalgs: %v", err)
		}
	}

	if err = d.Set("dual_stack_mode", dataSourceFlattenVpnSslSettingsDualStackMode(o["dual-stack-mode"], d, "dual_stack_mode")); err != nil {
		if !fortiAPIPatch(o["dual-stack-mode"]) {
			return fmt.Errorf("Error reading dual_stack_mode: %v", err)
		}
	}

	if err = d.Set("tunnel_addr_assigned_method", dataSourceFlattenVpnSslSettingsTunnelAddrAssignedMethod(o["tunnel-addr-assigned-method"], d, "tunnel_addr_assigned_method")); err != nil {
		if !fortiAPIPatch(o["tunnel-addr-assigned-method"]) {
			return fmt.Errorf("Error reading tunnel_addr_assigned_method: %v", err)
		}
	}

	if err = d.Set("saml_redirect_port", dataSourceFlattenVpnSslSettingsSamlRedirectPort(o["saml-redirect-port"], d, "saml_redirect_port")); err != nil {
		if !fortiAPIPatch(o["saml-redirect-port"]) {
			return fmt.Errorf("Error reading saml_redirect_port: %v", err)
		}
	}

	if err = d.Set("web_mode_snat", dataSourceFlattenVpnSslSettingsWebModeSnat(o["web-mode-snat"], d, "web_mode_snat")); err != nil {
		if !fortiAPIPatch(o["web-mode-snat"]) {
			return fmt.Errorf("Error reading web_mode_snat: %v", err)
		}
	}

	if err = d.Set("ztna_trusted_client", dataSourceFlattenVpnSslSettingsZtnaTrustedClient(o["ztna-trusted-client"], d, "ztna_trusted_client")); err != nil {
		if !fortiAPIPatch(o["ztna-trusted-client"]) {
			return fmt.Errorf("Error reading ztna_trusted_client: %v", err)
		}
	}

	if err = d.Set("server_hostname", dataSourceFlattenVpnSslSettingsServerHostname(o["server-hostname"], d, "server_hostname")); err != nil {
		if !fortiAPIPatch(o["server-hostname"]) {
			return fmt.Errorf("Error reading server_hostname: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenVpnSslSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
