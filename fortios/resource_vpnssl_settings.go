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

func resourceVpnSslSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslSettingsUpdate,
		Read:   resourceVpnSslSettingsRead,
		Update: resourceVpnSslSettingsUpdate,
		Delete: resourceVpnSslSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reqclientcert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_max_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tlsv1_0": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tlsv1_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tlsv1_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tlsv1_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"banned_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ciphersuite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_insert_empty_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"x_content_type_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_two_factor_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unsafe_legacy_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servercert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 259200),
				Optional:     true,
				Computed:     true,
			},
			"auth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 259200),
				Optional:     true,
				Computed:     true,
			},
			"login_attempt_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"login_block_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"login_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 180),
				Optional:     true,
				Computed:     true,
			},
			"dtls_hello_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 60),
				Optional:     true,
				Computed:     true,
			},
			"tunnel_ip_pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"tunnel_ipv6_pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dns_suffix": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 253),
				Optional:     true,
			},
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_source_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_obscuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_compression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_only_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deflate_compression_level": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9),
				Optional:     true,
				Computed:     true,
			},
			"deflate_min_data_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(200, 65535),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"port_precedence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_tunnel_static_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"source_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"source_address_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_address6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"source_address6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_portal": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"authentication_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"source_interface": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"source_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"source_address_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_address6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"source_address6_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"users": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"portal": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"realm": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"client_cert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_peer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"browser_language_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_max_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_min_proto_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"check_referer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_request_header_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http_request_body_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_session_check_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_connect_without_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_user_session_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transform_backward_slashes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encode_2f_sequence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt_and_store_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_sigalgs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dual_stack_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_addr_assigned_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_redirect_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"web_mode_snat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_trusted_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceVpnSslSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslSettings(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslSettings")
	}

	return resourceVpnSslSettingsRead(d, m)
}

func resourceVpnSslSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslSettings(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating VpnSslSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnSslSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing VpnSslSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnSslSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslSettings resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsReqclientcert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsUserPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSslMaxProtoVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSslMinProtoVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTlsv10(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTlsv11(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTlsv12(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTlsv13(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsBannedCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsCiphersuite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSslInsertEmptyFragment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpsRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsXContentTypeOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsForceTwoFactorAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsUnsafeLegacyRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsServercert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsLoginAttemptLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsLoginBlockTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsLoginTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsHelloTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelIpPools(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsTunnelIpPoolsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsTunnelIpPoolsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelIpv6Pools(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsTunnelIpv6PoolsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsTunnelIpv6PoolsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDnsSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsWinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsWinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6WinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsIpv6WinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsRouteSourceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsUrlObscuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpCompression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpOnlyCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDeflateCompressionLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDeflateMinDataSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsPortPrecedence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAutoTunnelStaticRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsSourceInterfaceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsSourceAddressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsSourceAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsSourceAddress6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDefaultPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenVpnSslSettingsAuthenticationRuleId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_interface"
		if _, ok := i["source-interface"]; ok {

			tmp["source_interface"] = flattenVpnSslSettingsAuthenticationRuleSourceInterface(i["source-interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := i["source-address"]; ok {

			tmp["source_address"] = flattenVpnSslSettingsAuthenticationRuleSourceAddress(i["source-address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address_negate"
		if _, ok := i["source-address-negate"]; ok {

			tmp["source_address_negate"] = flattenVpnSslSettingsAuthenticationRuleSourceAddressNegate(i["source-address-negate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6"
		if _, ok := i["source-address6"]; ok {

			tmp["source_address6"] = flattenVpnSslSettingsAuthenticationRuleSourceAddress6(i["source-address6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6_negate"
		if _, ok := i["source-address6-negate"]; ok {

			tmp["source_address6_negate"] = flattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate(i["source-address6-negate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {

			tmp["users"] = flattenVpnSslSettingsAuthenticationRuleUsers(i["users"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {

			tmp["groups"] = flattenVpnSslSettingsAuthenticationRuleGroups(i["groups"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal"
		if _, ok := i["portal"]; ok {

			tmp["portal"] = flattenVpnSslSettingsAuthenticationRulePortal(i["portal"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := i["realm"]; ok {

			tmp["realm"] = flattenVpnSslSettingsAuthenticationRuleRealm(i["realm"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := i["client-cert"]; ok {

			tmp["client_cert"] = flattenVpnSslSettingsAuthenticationRuleClientCert(i["client-cert"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_peer"
		if _, ok := i["user-peer"]; ok {

			tmp["user_peer"] = flattenVpnSslSettingsAuthenticationRuleUserPeer(i["user-peer"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenVpnSslSettingsAuthenticationRuleCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := i["auth"]; ok {

			tmp["auth"] = flattenVpnSslSettingsAuthenticationRuleAuth(i["auth"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnSslSettingsAuthenticationRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsAuthenticationRuleSourceInterfaceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsAuthenticationRuleSourceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsAuthenticationRuleSourceAddressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddressNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsAuthenticationRuleSourceAddress6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleSourceAddress6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsAuthenticationRuleUsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsAuthenticationRuleUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnSslSettingsAuthenticationRuleGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslSettingsAuthenticationRuleGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRulePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleUserPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthenticationRuleAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsBrowserLanguageDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsTunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsMaxProtoVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDtlsMinProtoVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsCheckReferer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpRequestHeaderTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHttpRequestBodyTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsAuthSessionCheckSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelConnectWithoutReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelUserSessionTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTransformBackwardSlashes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsEncode2FSequence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsEncryptAndStorePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsClientSigalgs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsDualStackMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsTunnelAddrAssignedMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsSamlRedirectPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsWebModeSnat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslSettingsZtnaTrustedClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnSslSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenVpnSslSettingsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("reqclientcert", flattenVpnSslSettingsReqclientcert(o["reqclientcert"], d, "reqclientcert", sv)); err != nil {
		if !fortiAPIPatch(o["reqclientcert"]) {
			return fmt.Errorf("Error reading reqclientcert: %v", err)
		}
	}

	if err = d.Set("user_peer", flattenVpnSslSettingsUserPeer(o["user-peer"], d, "user_peer", sv)); err != nil {
		if !fortiAPIPatch(o["user-peer"]) {
			return fmt.Errorf("Error reading user_peer: %v", err)
		}
	}

	if err = d.Set("ssl_max_proto_ver", flattenVpnSslSettingsSslMaxProtoVer(o["ssl-max-proto-ver"], d, "ssl_max_proto_ver", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-proto-ver"]) {
			return fmt.Errorf("Error reading ssl_max_proto_ver: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_ver", flattenVpnSslSettingsSslMinProtoVer(o["ssl-min-proto-ver"], d, "ssl_min_proto_ver", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-ver"]) {
			return fmt.Errorf("Error reading ssl_min_proto_ver: %v", err)
		}
	}

	if err = d.Set("tlsv1_0", flattenVpnSslSettingsTlsv10(o["tlsv1-0"], d, "tlsv1_0", sv)); err != nil {
		if !fortiAPIPatch(o["tlsv1-0"]) {
			return fmt.Errorf("Error reading tlsv1_0: %v", err)
		}
	}

	if err = d.Set("tlsv1_1", flattenVpnSslSettingsTlsv11(o["tlsv1-1"], d, "tlsv1_1", sv)); err != nil {
		if !fortiAPIPatch(o["tlsv1-1"]) {
			return fmt.Errorf("Error reading tlsv1_1: %v", err)
		}
	}

	if err = d.Set("tlsv1_2", flattenVpnSslSettingsTlsv12(o["tlsv1-2"], d, "tlsv1_2", sv)); err != nil {
		if !fortiAPIPatch(o["tlsv1-2"]) {
			return fmt.Errorf("Error reading tlsv1_2: %v", err)
		}
	}

	if err = d.Set("tlsv1_3", flattenVpnSslSettingsTlsv13(o["tlsv1-3"], d, "tlsv1_3", sv)); err != nil {
		if !fortiAPIPatch(o["tlsv1-3"]) {
			return fmt.Errorf("Error reading tlsv1_3: %v", err)
		}
	}

	if err = d.Set("banned_cipher", flattenVpnSslSettingsBannedCipher(o["banned-cipher"], d, "banned_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["banned-cipher"]) {
			return fmt.Errorf("Error reading banned_cipher: %v", err)
		}
	}

	if err = d.Set("ciphersuite", flattenVpnSslSettingsCiphersuite(o["ciphersuite"], d, "ciphersuite", sv)); err != nil {
		if !fortiAPIPatch(o["ciphersuite"]) {
			return fmt.Errorf("Error reading ciphersuite: %v", err)
		}
	}

	if err = d.Set("ssl_insert_empty_fragment", flattenVpnSslSettingsSslInsertEmptyFragment(o["ssl-insert-empty-fragment"], d, "ssl_insert_empty_fragment", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-insert-empty-fragment"]) {
			return fmt.Errorf("Error reading ssl_insert_empty_fragment: %v", err)
		}
	}

	if err = d.Set("https_redirect", flattenVpnSslSettingsHttpsRedirect(o["https-redirect"], d, "https_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["https-redirect"]) {
			return fmt.Errorf("Error reading https_redirect: %v", err)
		}
	}

	if err = d.Set("x_content_type_options", flattenVpnSslSettingsXContentTypeOptions(o["x-content-type-options"], d, "x_content_type_options", sv)); err != nil {
		if !fortiAPIPatch(o["x-content-type-options"]) {
			return fmt.Errorf("Error reading x_content_type_options: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenVpnSslSettingsSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("force_two_factor_auth", flattenVpnSslSettingsForceTwoFactorAuth(o["force-two-factor-auth"], d, "force_two_factor_auth", sv)); err != nil {
		if !fortiAPIPatch(o["force-two-factor-auth"]) {
			return fmt.Errorf("Error reading force_two_factor_auth: %v", err)
		}
	}

	if err = d.Set("unsafe_legacy_renegotiation", flattenVpnSslSettingsUnsafeLegacyRenegotiation(o["unsafe-legacy-renegotiation"], d, "unsafe_legacy_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["unsafe-legacy-renegotiation"]) {
			return fmt.Errorf("Error reading unsafe_legacy_renegotiation: %v", err)
		}
	}

	if err = d.Set("servercert", flattenVpnSslSettingsServercert(o["servercert"], d, "servercert", sv)); err != nil {
		if !fortiAPIPatch(o["servercert"]) {
			return fmt.Errorf("Error reading servercert: %v", err)
		}
	}

	if err = d.Set("algorithm", flattenVpnSslSettingsAlgorithm(o["algorithm"], d, "algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["algorithm"]) {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnSslSettingsIdleTimeout(o["idle-timeout"], d, "idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("auth_timeout", flattenVpnSslSettingsAuthTimeout(o["auth-timeout"], d, "auth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auth-timeout"]) {
			return fmt.Errorf("Error reading auth_timeout: %v", err)
		}
	}

	if err = d.Set("login_attempt_limit", flattenVpnSslSettingsLoginAttemptLimit(o["login-attempt-limit"], d, "login_attempt_limit", sv)); err != nil {
		if !fortiAPIPatch(o["login-attempt-limit"]) {
			return fmt.Errorf("Error reading login_attempt_limit: %v", err)
		}
	}

	if err = d.Set("login_block_time", flattenVpnSslSettingsLoginBlockTime(o["login-block-time"], d, "login_block_time", sv)); err != nil {
		if !fortiAPIPatch(o["login-block-time"]) {
			return fmt.Errorf("Error reading login_block_time: %v", err)
		}
	}

	if err = d.Set("login_timeout", flattenVpnSslSettingsLoginTimeout(o["login-timeout"], d, "login_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["login-timeout"]) {
			return fmt.Errorf("Error reading login_timeout: %v", err)
		}
	}

	if err = d.Set("dtls_hello_timeout", flattenVpnSslSettingsDtlsHelloTimeout(o["dtls-hello-timeout"], d, "dtls_hello_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["dtls-hello-timeout"]) {
			return fmt.Errorf("Error reading dtls_hello_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tunnel_ip_pools", flattenVpnSslSettingsTunnelIpPools(o["tunnel-ip-pools"], d, "tunnel_ip_pools", sv)); err != nil {
			if !fortiAPIPatch(o["tunnel-ip-pools"]) {
				return fmt.Errorf("Error reading tunnel_ip_pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tunnel_ip_pools"); ok {
			if err = d.Set("tunnel_ip_pools", flattenVpnSslSettingsTunnelIpPools(o["tunnel-ip-pools"], d, "tunnel_ip_pools", sv)); err != nil {
				if !fortiAPIPatch(o["tunnel-ip-pools"]) {
					return fmt.Errorf("Error reading tunnel_ip_pools: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("tunnel_ipv6_pools", flattenVpnSslSettingsTunnelIpv6Pools(o["tunnel-ipv6-pools"], d, "tunnel_ipv6_pools", sv)); err != nil {
			if !fortiAPIPatch(o["tunnel-ipv6-pools"]) {
				return fmt.Errorf("Error reading tunnel_ipv6_pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tunnel_ipv6_pools"); ok {
			if err = d.Set("tunnel_ipv6_pools", flattenVpnSslSettingsTunnelIpv6Pools(o["tunnel-ipv6-pools"], d, "tunnel_ipv6_pools", sv)); err != nil {
				if !fortiAPIPatch(o["tunnel-ipv6-pools"]) {
					return fmt.Errorf("Error reading tunnel_ipv6_pools: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns_suffix", flattenVpnSslSettingsDnsSuffix(o["dns-suffix"], d, "dns_suffix", sv)); err != nil {
		if !fortiAPIPatch(o["dns-suffix"]) {
			return fmt.Errorf("Error reading dns_suffix: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenVpnSslSettingsDnsServer1(o["dns-server1"], d, "dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenVpnSslSettingsDnsServer2(o["dns-server2"], d, "dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenVpnSslSettingsWinsServer1(o["wins-server1"], d, "wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["wins-server1"]) {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenVpnSslSettingsWinsServer2(o["wins-server2"], d, "wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["wins-server2"]) {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnSslSettingsIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server1"]) {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnSslSettingsIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server2"]) {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server1", flattenVpnSslSettingsIpv6WinsServer1(o["ipv6-wins-server1"], d, "ipv6_wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-wins-server1"]) {
			return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server2", flattenVpnSslSettingsIpv6WinsServer2(o["ipv6-wins-server2"], d, "ipv6_wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-wins-server2"]) {
			return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
		}
	}

	if err = d.Set("route_source_interface", flattenVpnSslSettingsRouteSourceInterface(o["route-source-interface"], d, "route_source_interface", sv)); err != nil {
		if !fortiAPIPatch(o["route-source-interface"]) {
			return fmt.Errorf("Error reading route_source_interface: %v", err)
		}
	}

	if err = d.Set("url_obscuration", flattenVpnSslSettingsUrlObscuration(o["url-obscuration"], d, "url_obscuration", sv)); err != nil {
		if !fortiAPIPatch(o["url-obscuration"]) {
			return fmt.Errorf("Error reading url_obscuration: %v", err)
		}
	}

	if err = d.Set("http_compression", flattenVpnSslSettingsHttpCompression(o["http-compression"], d, "http_compression", sv)); err != nil {
		if !fortiAPIPatch(o["http-compression"]) {
			return fmt.Errorf("Error reading http_compression: %v", err)
		}
	}

	if err = d.Set("http_only_cookie", flattenVpnSslSettingsHttpOnlyCookie(o["http-only-cookie"], d, "http_only_cookie", sv)); err != nil {
		if !fortiAPIPatch(o["http-only-cookie"]) {
			return fmt.Errorf("Error reading http_only_cookie: %v", err)
		}
	}

	if err = d.Set("deflate_compression_level", flattenVpnSslSettingsDeflateCompressionLevel(o["deflate-compression-level"], d, "deflate_compression_level", sv)); err != nil {
		if !fortiAPIPatch(o["deflate-compression-level"]) {
			return fmt.Errorf("Error reading deflate_compression_level: %v", err)
		}
	}

	if err = d.Set("deflate_min_data_size", flattenVpnSslSettingsDeflateMinDataSize(o["deflate-min-data-size"], d, "deflate_min_data_size", sv)); err != nil {
		if !fortiAPIPatch(o["deflate-min-data-size"]) {
			return fmt.Errorf("Error reading deflate_min_data_size: %v", err)
		}
	}

	if err = d.Set("port", flattenVpnSslSettingsPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("port_precedence", flattenVpnSslSettingsPortPrecedence(o["port-precedence"], d, "port_precedence", sv)); err != nil {
		if !fortiAPIPatch(o["port-precedence"]) {
			return fmt.Errorf("Error reading port_precedence: %v", err)
		}
	}

	if err = d.Set("auto_tunnel_static_route", flattenVpnSslSettingsAutoTunnelStaticRoute(o["auto-tunnel-static-route"], d, "auto_tunnel_static_route", sv)); err != nil {
		if !fortiAPIPatch(o["auto-tunnel-static-route"]) {
			return fmt.Errorf("Error reading auto_tunnel_static_route: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", flattenVpnSslSettingsHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for", sv)); err != nil {
		if !fortiAPIPatch(o["header-x-forwarded-for"]) {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_interface", flattenVpnSslSettingsSourceInterface(o["source-interface"], d, "source_interface", sv)); err != nil {
			if !fortiAPIPatch(o["source-interface"]) {
				return fmt.Errorf("Error reading source_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_interface"); ok {
			if err = d.Set("source_interface", flattenVpnSslSettingsSourceInterface(o["source-interface"], d, "source_interface", sv)); err != nil {
				if !fortiAPIPatch(o["source-interface"]) {
					return fmt.Errorf("Error reading source_interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("source_address", flattenVpnSslSettingsSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
			if !fortiAPIPatch(o["source-address"]) {
				return fmt.Errorf("Error reading source_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_address"); ok {
			if err = d.Set("source_address", flattenVpnSslSettingsSourceAddress(o["source-address"], d, "source_address", sv)); err != nil {
				if !fortiAPIPatch(o["source-address"]) {
					return fmt.Errorf("Error reading source_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_address_negate", flattenVpnSslSettingsSourceAddressNegate(o["source-address-negate"], d, "source_address_negate", sv)); err != nil {
		if !fortiAPIPatch(o["source-address-negate"]) {
			return fmt.Errorf("Error reading source_address_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("source_address6", flattenVpnSslSettingsSourceAddress6(o["source-address6"], d, "source_address6", sv)); err != nil {
			if !fortiAPIPatch(o["source-address6"]) {
				return fmt.Errorf("Error reading source_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("source_address6"); ok {
			if err = d.Set("source_address6", flattenVpnSslSettingsSourceAddress6(o["source-address6"], d, "source_address6", sv)); err != nil {
				if !fortiAPIPatch(o["source-address6"]) {
					return fmt.Errorf("Error reading source_address6: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_address6_negate", flattenVpnSslSettingsSourceAddress6Negate(o["source-address6-negate"], d, "source_address6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["source-address6-negate"]) {
			return fmt.Errorf("Error reading source_address6_negate: %v", err)
		}
	}

	if err = d.Set("default_portal", flattenVpnSslSettingsDefaultPortal(o["default-portal"], d, "default_portal", sv)); err != nil {
		if !fortiAPIPatch(o["default-portal"]) {
			return fmt.Errorf("Error reading default_portal: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("authentication_rule", flattenVpnSslSettingsAuthenticationRule(o["authentication-rule"], d, "authentication_rule", sv)); err != nil {
			if !fortiAPIPatch(o["authentication-rule"]) {
				return fmt.Errorf("Error reading authentication_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("authentication_rule"); ok {
			if err = d.Set("authentication_rule", flattenVpnSslSettingsAuthenticationRule(o["authentication-rule"], d, "authentication_rule", sv)); err != nil {
				if !fortiAPIPatch(o["authentication-rule"]) {
					return fmt.Errorf("Error reading authentication_rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("browser_language_detection", flattenVpnSslSettingsBrowserLanguageDetection(o["browser-language-detection"], d, "browser_language_detection", sv)); err != nil {
		if !fortiAPIPatch(o["browser-language-detection"]) {
			return fmt.Errorf("Error reading browser_language_detection: %v", err)
		}
	}

	if err = d.Set("dtls_tunnel", flattenVpnSslSettingsDtlsTunnel(o["dtls-tunnel"], d, "dtls_tunnel", sv)); err != nil {
		if !fortiAPIPatch(o["dtls-tunnel"]) {
			return fmt.Errorf("Error reading dtls_tunnel: %v", err)
		}
	}

	if err = d.Set("dtls_max_proto_ver", flattenVpnSslSettingsDtlsMaxProtoVer(o["dtls-max-proto-ver"], d, "dtls_max_proto_ver", sv)); err != nil {
		if !fortiAPIPatch(o["dtls-max-proto-ver"]) {
			return fmt.Errorf("Error reading dtls_max_proto_ver: %v", err)
		}
	}

	if err = d.Set("dtls_min_proto_ver", flattenVpnSslSettingsDtlsMinProtoVer(o["dtls-min-proto-ver"], d, "dtls_min_proto_ver", sv)); err != nil {
		if !fortiAPIPatch(o["dtls-min-proto-ver"]) {
			return fmt.Errorf("Error reading dtls_min_proto_ver: %v", err)
		}
	}

	if err = d.Set("check_referer", flattenVpnSslSettingsCheckReferer(o["check-referer"], d, "check_referer", sv)); err != nil {
		if !fortiAPIPatch(o["check-referer"]) {
			return fmt.Errorf("Error reading check_referer: %v", err)
		}
	}

	if err = d.Set("http_request_header_timeout", flattenVpnSslSettingsHttpRequestHeaderTimeout(o["http-request-header-timeout"], d, "http_request_header_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["http-request-header-timeout"]) {
			return fmt.Errorf("Error reading http_request_header_timeout: %v", err)
		}
	}

	if err = d.Set("http_request_body_timeout", flattenVpnSslSettingsHttpRequestBodyTimeout(o["http-request-body-timeout"], d, "http_request_body_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["http-request-body-timeout"]) {
			return fmt.Errorf("Error reading http_request_body_timeout: %v", err)
		}
	}

	if err = d.Set("auth_session_check_source_ip", flattenVpnSslSettingsAuthSessionCheckSourceIp(o["auth-session-check-source-ip"], d, "auth_session_check_source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["auth-session-check-source-ip"]) {
			return fmt.Errorf("Error reading auth_session_check_source_ip: %v", err)
		}
	}

	if err = d.Set("tunnel_connect_without_reauth", flattenVpnSslSettingsTunnelConnectWithoutReauth(o["tunnel-connect-without-reauth"], d, "tunnel_connect_without_reauth", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-connect-without-reauth"]) {
			return fmt.Errorf("Error reading tunnel_connect_without_reauth: %v", err)
		}
	}

	if err = d.Set("tunnel_user_session_timeout", flattenVpnSslSettingsTunnelUserSessionTimeout(o["tunnel-user-session-timeout"], d, "tunnel_user_session_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-user-session-timeout"]) {
			return fmt.Errorf("Error reading tunnel_user_session_timeout: %v", err)
		}
	}

	if err = d.Set("hsts_include_subdomains", flattenVpnSslSettingsHstsIncludeSubdomains(o["hsts-include-subdomains"], d, "hsts_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["hsts-include-subdomains"]) {
			return fmt.Errorf("Error reading hsts_include_subdomains: %v", err)
		}
	}

	if err = d.Set("transform_backward_slashes", flattenVpnSslSettingsTransformBackwardSlashes(o["transform-backward-slashes"], d, "transform_backward_slashes", sv)); err != nil {
		if !fortiAPIPatch(o["transform-backward-slashes"]) {
			return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
		}
	}

	if err = d.Set("encode_2f_sequence", flattenVpnSslSettingsEncode2FSequence(o["encode-2f-sequence"], d, "encode_2f_sequence", sv)); err != nil {
		if !fortiAPIPatch(o["encode-2f-sequence"]) {
			return fmt.Errorf("Error reading encode_2f_sequence: %v", err)
		}
	}

	if err = d.Set("encrypt_and_store_password", flattenVpnSslSettingsEncryptAndStorePassword(o["encrypt-and-store-password"], d, "encrypt_and_store_password", sv)); err != nil {
		if !fortiAPIPatch(o["encrypt-and-store-password"]) {
			return fmt.Errorf("Error reading encrypt_and_store_password: %v", err)
		}
	}

	if err = d.Set("client_sigalgs", flattenVpnSslSettingsClientSigalgs(o["client-sigalgs"], d, "client_sigalgs", sv)); err != nil {
		if !fortiAPIPatch(o["client-sigalgs"]) {
			return fmt.Errorf("Error reading client_sigalgs: %v", err)
		}
	}

	if err = d.Set("dual_stack_mode", flattenVpnSslSettingsDualStackMode(o["dual-stack-mode"], d, "dual_stack_mode", sv)); err != nil {
		if !fortiAPIPatch(o["dual-stack-mode"]) {
			return fmt.Errorf("Error reading dual_stack_mode: %v", err)
		}
	}

	if err = d.Set("tunnel_addr_assigned_method", flattenVpnSslSettingsTunnelAddrAssignedMethod(o["tunnel-addr-assigned-method"], d, "tunnel_addr_assigned_method", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-addr-assigned-method"]) {
			return fmt.Errorf("Error reading tunnel_addr_assigned_method: %v", err)
		}
	}

	if err = d.Set("saml_redirect_port", flattenVpnSslSettingsSamlRedirectPort(o["saml-redirect-port"], d, "saml_redirect_port", sv)); err != nil {
		if !fortiAPIPatch(o["saml-redirect-port"]) {
			return fmt.Errorf("Error reading saml_redirect_port: %v", err)
		}
	}

	if err = d.Set("web_mode_snat", flattenVpnSslSettingsWebModeSnat(o["web-mode-snat"], d, "web_mode_snat", sv)); err != nil {
		if !fortiAPIPatch(o["web-mode-snat"]) {
			return fmt.Errorf("Error reading web_mode_snat: %v", err)
		}
	}

	if err = d.Set("ztna_trusted_client", flattenVpnSslSettingsZtnaTrustedClient(o["ztna-trusted-client"], d, "ztna_trusted_client", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-trusted-client"]) {
			return fmt.Errorf("Error reading ztna_trusted_client: %v", err)
		}
	}

	return nil
}

func flattenVpnSslSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnSslSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsReqclientcert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsUserPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslMaxProtoVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslMinProtoVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTlsv10(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTlsv11(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTlsv12(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTlsv13(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsBannedCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsCiphersuite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslInsertEmptyFragment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpsRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsXContentTypeOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsForceTwoFactorAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsUnsafeLegacyRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsServercert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsLoginAttemptLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsLoginBlockTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsLoginTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsHelloTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelIpPools(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsTunnelIpPoolsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsTunnelIpPoolsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelIpv6Pools(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsTunnelIpv6PoolsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsTunnelIpv6PoolsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDnsSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsWinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsWinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6WinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsIpv6WinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsRouteSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsUrlObscuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpCompression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpOnlyCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDeflateCompressionLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDeflateMinDataSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsPortPrecedence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAutoTunnelStaticRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHeaderXForwardedFor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsSourceInterfaceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsSourceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsSourceAddressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsSourceAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceAddressNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsSourceAddress6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsSourceAddress6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSourceAddress6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDefaultPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandVpnSslSettingsAuthenticationRuleId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_interface"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["source-interface"], _ = expandVpnSslSettingsAuthenticationRuleSourceInterface(d, i["source_interface"], pre_append, sv)
		} else {
			tmp["source-interface"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["source-address"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress(d, i["source_address"], pre_append, sv)
		} else {
			tmp["source-address"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address_negate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source-address-negate"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddressNegate(d, i["source_address_negate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["source-address6"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress6(d, i["source_address6"], pre_append, sv)
		} else {
			tmp["source-address6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_address6_negate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source-address6-negate"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress6Negate(d, i["source_address6_negate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["users"], _ = expandVpnSslSettingsAuthenticationRuleUsers(d, i["users"], pre_append, sv)
		} else {
			tmp["users"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["groups"], _ = expandVpnSslSettingsAuthenticationRuleGroups(d, i["groups"], pre_append, sv)
		} else {
			tmp["groups"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "portal"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["portal"], _ = expandVpnSslSettingsAuthenticationRulePortal(d, i["portal"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realm"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["realm"], _ = expandVpnSslSettingsAuthenticationRuleRealm(d, i["realm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_cert"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["client-cert"], _ = expandVpnSslSettingsAuthenticationRuleClientCert(d, i["client_cert"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_peer"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["user-peer"], _ = expandVpnSslSettingsAuthenticationRuleUserPeer(d, i["user_peer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandVpnSslSettingsAuthenticationRuleCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth"], _ = expandVpnSslSettingsAuthenticationRuleAuth(d, i["auth"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsAuthenticationRuleSourceInterfaceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddressNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsAuthenticationRuleSourceAddress6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleSourceAddress6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsAuthenticationRuleUsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnSslSettingsAuthenticationRuleGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslSettingsAuthenticationRuleGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRulePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleUserPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthenticationRuleAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsBrowserLanguageDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsTunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsMaxProtoVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDtlsMinProtoVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsCheckReferer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpRequestHeaderTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHttpRequestBodyTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsAuthSessionCheckSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelConnectWithoutReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelUserSessionTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTransformBackwardSlashes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsEncode2FSequence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsEncryptAndStorePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsClientSigalgs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsDualStackMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsTunnelAddrAssignedMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsSamlRedirectPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsWebModeSnat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslSettingsZtnaTrustedClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslSettings(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandVpnSslSettingsStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("reqclientcert"); ok {
		if setArgNil {
			obj["reqclientcert"] = nil
		} else {

			t, err := expandVpnSslSettingsReqclientcert(d, v, "reqclientcert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["reqclientcert"] = t
			}
		}
	}

	if v, ok := d.GetOk("user_peer"); ok {
		if setArgNil {
			obj["user-peer"] = nil
		} else {

			t, err := expandVpnSslSettingsUserPeer(d, v, "user_peer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["user-peer"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_max_proto_ver"); ok {
		if setArgNil {
			obj["ssl-max-proto-ver"] = nil
		} else {

			t, err := expandVpnSslSettingsSslMaxProtoVer(d, v, "ssl_max_proto_ver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-max-proto-ver"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_ver"); ok {
		if setArgNil {
			obj["ssl-min-proto-ver"] = nil
		} else {

			t, err := expandVpnSslSettingsSslMinProtoVer(d, v, "ssl_min_proto_ver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-min-proto-ver"] = t
			}
		}
	}

	if v, ok := d.GetOk("tlsv1_0"); ok {
		if setArgNil {
			obj["tlsv1-0"] = nil
		} else {

			t, err := expandVpnSslSettingsTlsv10(d, v, "tlsv1_0", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tlsv1-0"] = t
			}
		}
	}

	if v, ok := d.GetOk("tlsv1_1"); ok {
		if setArgNil {
			obj["tlsv1-1"] = nil
		} else {

			t, err := expandVpnSslSettingsTlsv11(d, v, "tlsv1_1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tlsv1-1"] = t
			}
		}
	}

	if v, ok := d.GetOk("tlsv1_2"); ok {
		if setArgNil {
			obj["tlsv1-2"] = nil
		} else {

			t, err := expandVpnSslSettingsTlsv12(d, v, "tlsv1_2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tlsv1-2"] = t
			}
		}
	}

	if v, ok := d.GetOk("tlsv1_3"); ok {
		if setArgNil {
			obj["tlsv1-3"] = nil
		} else {

			t, err := expandVpnSslSettingsTlsv13(d, v, "tlsv1_3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tlsv1-3"] = t
			}
		}
	}

	if v, ok := d.GetOk("banned_cipher"); ok {
		if setArgNil {
			obj["banned-cipher"] = nil
		} else {

			t, err := expandVpnSslSettingsBannedCipher(d, v, "banned_cipher", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["banned-cipher"] = t
			}
		}
	}

	if v, ok := d.GetOk("ciphersuite"); ok {
		if setArgNil {
			obj["ciphersuite"] = nil
		} else {

			t, err := expandVpnSslSettingsCiphersuite(d, v, "ciphersuite", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ciphersuite"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_insert_empty_fragment"); ok {
		if setArgNil {
			obj["ssl-insert-empty-fragment"] = nil
		} else {

			t, err := expandVpnSslSettingsSslInsertEmptyFragment(d, v, "ssl_insert_empty_fragment", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-insert-empty-fragment"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_redirect"); ok {
		if setArgNil {
			obj["https-redirect"] = nil
		} else {

			t, err := expandVpnSslSettingsHttpsRedirect(d, v, "https_redirect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-redirect"] = t
			}
		}
	}

	if v, ok := d.GetOk("x_content_type_options"); ok {
		if setArgNil {
			obj["x-content-type-options"] = nil
		} else {

			t, err := expandVpnSslSettingsXContentTypeOptions(d, v, "x_content_type_options", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["x-content-type-options"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok {
		if setArgNil {
			obj["ssl-client-renegotiation"] = nil
		} else {

			t, err := expandVpnSslSettingsSslClientRenegotiation(d, v, "ssl_client_renegotiation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-client-renegotiation"] = t
			}
		}
	}

	if v, ok := d.GetOk("force_two_factor_auth"); ok {
		if setArgNil {
			obj["force-two-factor-auth"] = nil
		} else {

			t, err := expandVpnSslSettingsForceTwoFactorAuth(d, v, "force_two_factor_auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["force-two-factor-auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("unsafe_legacy_renegotiation"); ok {
		if setArgNil {
			obj["unsafe-legacy-renegotiation"] = nil
		} else {

			t, err := expandVpnSslSettingsUnsafeLegacyRenegotiation(d, v, "unsafe_legacy_renegotiation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unsafe-legacy-renegotiation"] = t
			}
		}
	}

	if v, ok := d.GetOk("servercert"); ok {
		if setArgNil {
			obj["servercert"] = nil
		} else {

			t, err := expandVpnSslSettingsServercert(d, v, "servercert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["servercert"] = t
			}
		}
	}

	if v, ok := d.GetOk("algorithm"); ok {
		if setArgNil {
			obj["algorithm"] = nil
		} else {

			t, err := expandVpnSslSettingsAlgorithm(d, v, "algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("idle_timeout"); ok {
		if setArgNil {
			obj["idle-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsIdleTimeout(d, v, "idle_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["idle-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("auth_timeout"); ok {
		if setArgNil {
			obj["auth-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsAuthTimeout(d, v, "auth_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("login_attempt_limit"); ok {
		if setArgNil {
			obj["login-attempt-limit"] = nil
		} else {

			t, err := expandVpnSslSettingsLoginAttemptLimit(d, v, "login_attempt_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login-attempt-limit"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("login_block_time"); ok {
		if setArgNil {
			obj["login-block-time"] = nil
		} else {

			t, err := expandVpnSslSettingsLoginBlockTime(d, v, "login_block_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login-block-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("login_timeout"); ok {
		if setArgNil {
			obj["login-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsLoginTimeout(d, v, "login_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("dtls_hello_timeout"); ok {
		if setArgNil {
			obj["dtls-hello-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsDtlsHelloTimeout(d, v, "dtls_hello_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dtls-hello-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_ip_pools"); ok || d.HasChange("tunnel_ip_pools") {
		if setArgNil {
			obj["tunnel-ip-pools"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnSslSettingsTunnelIpPools(d, v, "tunnel_ip_pools", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-ip-pools"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_ipv6_pools"); ok || d.HasChange("tunnel_ipv6_pools") {
		if setArgNil {
			obj["tunnel-ipv6-pools"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnSslSettingsTunnelIpv6Pools(d, v, "tunnel_ipv6_pools", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-ipv6-pools"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_suffix"); ok {
		if setArgNil {
			obj["dns-suffix"] = nil
		} else {

			t, err := expandVpnSslSettingsDnsSuffix(d, v, "dns_suffix", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-suffix"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok {
		if setArgNil {
			obj["dns-server1"] = nil
		} else {

			t, err := expandVpnSslSettingsDnsServer1(d, v, "dns_server1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-server1"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok {
		if setArgNil {
			obj["dns-server2"] = nil
		} else {

			t, err := expandVpnSslSettingsDnsServer2(d, v, "dns_server2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns-server2"] = t
			}
		}
	}

	if v, ok := d.GetOk("wins_server1"); ok {
		if setArgNil {
			obj["wins-server1"] = nil
		} else {

			t, err := expandVpnSslSettingsWinsServer1(d, v, "wins_server1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wins-server1"] = t
			}
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok {
		if setArgNil {
			obj["wins-server2"] = nil
		} else {

			t, err := expandVpnSslSettingsWinsServer2(d, v, "wins_server2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wins-server2"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok {
		if setArgNil {
			obj["ipv6-dns-server1"] = nil
		} else {

			t, err := expandVpnSslSettingsIpv6DnsServer1(d, v, "ipv6_dns_server1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-dns-server1"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok {
		if setArgNil {
			obj["ipv6-dns-server2"] = nil
		} else {

			t, err := expandVpnSslSettingsIpv6DnsServer2(d, v, "ipv6_dns_server2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-dns-server2"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server1"); ok {
		if setArgNil {
			obj["ipv6-wins-server1"] = nil
		} else {

			t, err := expandVpnSslSettingsIpv6WinsServer1(d, v, "ipv6_wins_server1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-wins-server1"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server2"); ok {
		if setArgNil {
			obj["ipv6-wins-server2"] = nil
		} else {

			t, err := expandVpnSslSettingsIpv6WinsServer2(d, v, "ipv6_wins_server2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-wins-server2"] = t
			}
		}
	}

	if v, ok := d.GetOk("route_source_interface"); ok {
		if setArgNil {
			obj["route-source-interface"] = nil
		} else {

			t, err := expandVpnSslSettingsRouteSourceInterface(d, v, "route_source_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["route-source-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("url_obscuration"); ok {
		if setArgNil {
			obj["url-obscuration"] = nil
		} else {

			t, err := expandVpnSslSettingsUrlObscuration(d, v, "url_obscuration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["url-obscuration"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_compression"); ok {
		if setArgNil {
			obj["http-compression"] = nil
		} else {

			t, err := expandVpnSslSettingsHttpCompression(d, v, "http_compression", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-compression"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_only_cookie"); ok {
		if setArgNil {
			obj["http-only-cookie"] = nil
		} else {

			t, err := expandVpnSslSettingsHttpOnlyCookie(d, v, "http_only_cookie", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-only-cookie"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("deflate_compression_level"); ok {
		if setArgNil {
			obj["deflate-compression-level"] = nil
		} else {

			t, err := expandVpnSslSettingsDeflateCompressionLevel(d, v, "deflate_compression_level", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["deflate-compression-level"] = t
			}
		}
	}

	if v, ok := d.GetOk("deflate_min_data_size"); ok {
		if setArgNil {
			obj["deflate-min-data-size"] = nil
		} else {

			t, err := expandVpnSslSettingsDeflateMinDataSize(d, v, "deflate_min_data_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["deflate-min-data-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {

			t, err := expandVpnSslSettingsPort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	if v, ok := d.GetOk("port_precedence"); ok {
		if setArgNil {
			obj["port-precedence"] = nil
		} else {

			t, err := expandVpnSslSettingsPortPrecedence(d, v, "port_precedence", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port-precedence"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_tunnel_static_route"); ok {
		if setArgNil {
			obj["auto-tunnel-static-route"] = nil
		} else {

			t, err := expandVpnSslSettingsAutoTunnelStaticRoute(d, v, "auto_tunnel_static_route", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-tunnel-static-route"] = t
			}
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_for"); ok {
		if setArgNil {
			obj["header-x-forwarded-for"] = nil
		} else {

			t, err := expandVpnSslSettingsHeaderXForwardedFor(d, v, "header_x_forwarded_for", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["header-x-forwarded-for"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_interface"); ok || d.HasChange("source_interface") {
		if setArgNil {
			obj["source-interface"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnSslSettingsSourceInterface(d, v, "source_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_address"); ok || d.HasChange("source_address") {
		if setArgNil {
			obj["source-address"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnSslSettingsSourceAddress(d, v, "source_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_address_negate"); ok {
		if setArgNil {
			obj["source-address-negate"] = nil
		} else {

			t, err := expandVpnSslSettingsSourceAddressNegate(d, v, "source_address_negate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-address-negate"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_address6"); ok || d.HasChange("source_address6") {
		if setArgNil {
			obj["source-address6"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnSslSettingsSourceAddress6(d, v, "source_address6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-address6"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_address6_negate"); ok {
		if setArgNil {
			obj["source-address6-negate"] = nil
		} else {

			t, err := expandVpnSslSettingsSourceAddress6Negate(d, v, "source_address6_negate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-address6-negate"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_portal"); ok {
		if setArgNil {
			obj["default-portal"] = nil
		} else {

			t, err := expandVpnSslSettingsDefaultPortal(d, v, "default_portal", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-portal"] = t
			}
		}
	}

	if v, ok := d.GetOk("authentication_rule"); ok || d.HasChange("authentication_rule") {
		if setArgNil {
			obj["authentication-rule"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnSslSettingsAuthenticationRule(d, v, "authentication_rule", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authentication-rule"] = t
			}
		}
	}

	if v, ok := d.GetOk("browser_language_detection"); ok {
		if setArgNil {
			obj["browser-language-detection"] = nil
		} else {

			t, err := expandVpnSslSettingsBrowserLanguageDetection(d, v, "browser_language_detection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["browser-language-detection"] = t
			}
		}
	}

	if v, ok := d.GetOk("dtls_tunnel"); ok {
		if setArgNil {
			obj["dtls-tunnel"] = nil
		} else {

			t, err := expandVpnSslSettingsDtlsTunnel(d, v, "dtls_tunnel", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dtls-tunnel"] = t
			}
		}
	}

	if v, ok := d.GetOk("dtls_max_proto_ver"); ok {
		if setArgNil {
			obj["dtls-max-proto-ver"] = nil
		} else {

			t, err := expandVpnSslSettingsDtlsMaxProtoVer(d, v, "dtls_max_proto_ver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dtls-max-proto-ver"] = t
			}
		}
	}

	if v, ok := d.GetOk("dtls_min_proto_ver"); ok {
		if setArgNil {
			obj["dtls-min-proto-ver"] = nil
		} else {

			t, err := expandVpnSslSettingsDtlsMinProtoVer(d, v, "dtls_min_proto_ver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dtls-min-proto-ver"] = t
			}
		}
	}

	if v, ok := d.GetOk("check_referer"); ok {
		if setArgNil {
			obj["check-referer"] = nil
		} else {

			t, err := expandVpnSslSettingsCheckReferer(d, v, "check_referer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["check-referer"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("http_request_header_timeout"); ok {
		if setArgNil {
			obj["http-request-header-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsHttpRequestHeaderTimeout(d, v, "http_request_header_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-request-header-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("http_request_body_timeout"); ok {
		if setArgNil {
			obj["http-request-body-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsHttpRequestBodyTimeout(d, v, "http_request_body_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-request-body-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_session_check_source_ip"); ok {
		if setArgNil {
			obj["auth-session-check-source-ip"] = nil
		} else {

			t, err := expandVpnSslSettingsAuthSessionCheckSourceIp(d, v, "auth_session_check_source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-session-check-source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_connect_without_reauth"); ok {
		if setArgNil {
			obj["tunnel-connect-without-reauth"] = nil
		} else {

			t, err := expandVpnSslSettingsTunnelConnectWithoutReauth(d, v, "tunnel_connect_without_reauth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-connect-without-reauth"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_user_session_timeout"); ok {
		if setArgNil {
			obj["tunnel-user-session-timeout"] = nil
		} else {

			t, err := expandVpnSslSettingsTunnelUserSessionTimeout(d, v, "tunnel_user_session_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-user-session-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("hsts_include_subdomains"); ok {
		if setArgNil {
			obj["hsts-include-subdomains"] = nil
		} else {

			t, err := expandVpnSslSettingsHstsIncludeSubdomains(d, v, "hsts_include_subdomains", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hsts-include-subdomains"] = t
			}
		}
	}

	if v, ok := d.GetOk("transform_backward_slashes"); ok {
		if setArgNil {
			obj["transform-backward-slashes"] = nil
		} else {

			t, err := expandVpnSslSettingsTransformBackwardSlashes(d, v, "transform_backward_slashes", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["transform-backward-slashes"] = t
			}
		}
	}

	if v, ok := d.GetOk("encode_2f_sequence"); ok {
		if setArgNil {
			obj["encode-2f-sequence"] = nil
		} else {

			t, err := expandVpnSslSettingsEncode2FSequence(d, v, "encode_2f_sequence", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["encode-2f-sequence"] = t
			}
		}
	}

	if v, ok := d.GetOk("encrypt_and_store_password"); ok {
		if setArgNil {
			obj["encrypt-and-store-password"] = nil
		} else {

			t, err := expandVpnSslSettingsEncryptAndStorePassword(d, v, "encrypt_and_store_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["encrypt-and-store-password"] = t
			}
		}
	}

	if v, ok := d.GetOk("client_sigalgs"); ok {
		if setArgNil {
			obj["client-sigalgs"] = nil
		} else {

			t, err := expandVpnSslSettingsClientSigalgs(d, v, "client_sigalgs", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["client-sigalgs"] = t
			}
		}
	}

	if v, ok := d.GetOk("dual_stack_mode"); ok {
		if setArgNil {
			obj["dual-stack-mode"] = nil
		} else {

			t, err := expandVpnSslSettingsDualStackMode(d, v, "dual_stack_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dual-stack-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_addr_assigned_method"); ok {
		if setArgNil {
			obj["tunnel-addr-assigned-method"] = nil
		} else {

			t, err := expandVpnSslSettingsTunnelAddrAssignedMethod(d, v, "tunnel_addr_assigned_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-addr-assigned-method"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("saml_redirect_port"); ok {
		if setArgNil {
			obj["saml-redirect-port"] = nil
		} else {

			t, err := expandVpnSslSettingsSamlRedirectPort(d, v, "saml_redirect_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["saml-redirect-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("web_mode_snat"); ok {
		if setArgNil {
			obj["web-mode-snat"] = nil
		} else {

			t, err := expandVpnSslSettingsWebModeSnat(d, v, "web_mode_snat", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["web-mode-snat"] = t
			}
		}
	}

	if v, ok := d.GetOk("ztna_trusted_client"); ok {
		if setArgNil {
			obj["ztna-trusted-client"] = nil
		} else {

			t, err := expandVpnSslSettingsZtnaTrustedClient(d, v, "ztna_trusted_client", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ztna-trusted-client"] = t
			}
		}
	}

	return &obj, nil
}
