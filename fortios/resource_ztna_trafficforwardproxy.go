// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ZTNA traffic forward proxy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaTrafficForwardProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaTrafficForwardProxyCreate,
		Read:   resourceZtnaTrafficForwardProxyRead,
		Update: resourceZtnaTrafficForwardProxyUpdate,
		Delete: resourceZtnaTrafficForwardProxyDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_agent_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_virtual_host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"vip6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"empty_cert_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_blocked_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svr_pool_multiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svr_pool_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"svr_pool_server_max_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"svr_pool_server_max_concurrent_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"h3_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quic": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_idle_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 60000),
							Optional:     true,
							Computed:     true,
						},
						"max_udp_payload_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1200, 1500),
							Optional:     true,
							Computed:     true,
						},
						"active_connection_id_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 8),
							Optional:     true,
							Computed:     true,
						},
						"ack_delay_exponent": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),
							Optional:     true,
							Computed:     true,
						},
						"max_ack_delay": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16383),
							Optional:     true,
							Computed:     true,
						},
						"max_datagram_frame_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 1500),
							Optional:     true,
							Computed:     true,
						},
						"active_migration": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"grease_quic_bit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"versions": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_server_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_cipher_suites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"cipher": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"versions": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ssl_pfs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_accept_ffdhe_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_fallback": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_session_state_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 14400),
				Optional:     true,
				Computed:     true,
			},
			"ssl_client_session_state_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10000),
				Optional:     true,
				Computed:     true,
			},
			"ssl_client_rekey_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(200, 1048576),
				Optional:     true,
			},
			"ssl_server_renegotiation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_session_state_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_session_state_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 14400),
				Optional:     true,
				Computed:     true,
			},
			"ssl_server_session_state_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10000),
				Optional:     true,
				Computed:     true,
			},
			"ssl_http_location_conversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_http_match_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hpkp_primary": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ssl_hpkp_backup": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ssl_hpkp_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 157680000),
				Optional:     true,
				Computed:     true,
			},
			"ssl_hpkp_report_uri": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ssl_hpkp_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hsts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_hsts_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 157680000),
				Optional:     true,
				Computed:     true,
			},
			"ssl_hsts_include_subdomains": &schema.Schema{
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

func resourceZtnaTrafficForwardProxyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaTrafficForwardProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxy resource while getting object: %v", err)
	}

	o, err := c.CreateZtnaTrafficForwardProxy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaTrafficForwardProxy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaTrafficForwardProxy")
	}

	return resourceZtnaTrafficForwardProxyRead(d, m)
}

func resourceZtnaTrafficForwardProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaTrafficForwardProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxy resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaTrafficForwardProxy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaTrafficForwardProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaTrafficForwardProxy")
	}

	return resourceZtnaTrafficForwardProxyRead(d, m)
}

func resourceZtnaTrafficForwardProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteZtnaTrafficForwardProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaTrafficForwardProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaTrafficForwardProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaTrafficForwardProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaTrafficForwardProxy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaTrafficForwardProxy resource from API: %v", err)
	}
	return nil
}

func flattenZtnaTrafficForwardProxyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyUserAgentDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyAuthPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyVip6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySvrPoolTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyH3Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaTrafficForwardProxyQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaTrafficForwardProxyQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaTrafficForwardProxyQuicMaxAckDelay(i["max-ack-delay"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaTrafficForwardProxyQuicActiveMigration(i["active-migration"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaTrafficForwardProxyQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaTrafficForwardProxyQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxyQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxyQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenZtnaTrafficForwardProxySslCertificateName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenZtnaTrafficForwardProxySslCertificateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenZtnaTrafficForwardProxySslCipherSuitesPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if cur_v, ok := i["cipher"]; ok {
			tmp["cipher"] = flattenZtnaTrafficForwardProxySslCipherSuitesCipher(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if cur_v, ok := i["versions"]; ok {
			tmp["versions"] = flattenZtnaTrafficForwardProxySslCipherSuitesVersions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenZtnaTrafficForwardProxySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenZtnaTrafficForwardProxySslServerCipherSuitesPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if cur_v, ok := i["cipher"]; ok {
			tmp["cipher"] = flattenZtnaTrafficForwardProxySslServerCipherSuitesCipher(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if cur_v, ok := i["versions"]; ok {
			tmp["versions"] = flattenZtnaTrafficForwardProxySslServerCipherSuitesVersions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslPfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslServerRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHsts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaTrafficForwardProxySslHstsAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaTrafficForwardProxySslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectZtnaTrafficForwardProxy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenZtnaTrafficForwardProxyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vip", flattenZtnaTrafficForwardProxyVip(o["vip"], d, "vip", sv)); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("host", flattenZtnaTrafficForwardProxyHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenZtnaTrafficForwardProxyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("status", flattenZtnaTrafficForwardProxyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("interface", flattenZtnaTrafficForwardProxyInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("port", flattenZtnaTrafficForwardProxyPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenZtnaTrafficForwardProxyClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenZtnaTrafficForwardProxyUserAgentDetect(o["user-agent-detect"], d, "user_agent_detect", sv)); err != nil {
		if !fortiAPIPatch(o["user-agent-detect"]) {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	if err = d.Set("auth_portal", flattenZtnaTrafficForwardProxyAuthPortal(o["auth-portal"], d, "auth_portal", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal"]) {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenZtnaTrafficForwardProxyAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host", sv)); err != nil {
		if !fortiAPIPatch(o["auth-virtual-host"]) {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("vip6", flattenZtnaTrafficForwardProxyVip6(o["vip6"], d, "vip6", sv)); err != nil {
		if !fortiAPIPatch(o["vip6"]) {
			return fmt.Errorf("Error reading vip6: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenZtnaTrafficForwardProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action", sv)); err != nil {
		if !fortiAPIPatch(o["empty-cert-action"]) {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenZtnaTrafficForwardProxyLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["log-blocked-traffic"]) {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenZtnaTrafficForwardProxySvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-multiplex"]) {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenZtnaTrafficForwardProxySvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-ttl"]) {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenZtnaTrafficForwardProxySvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-server-max-request"]) {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-server-max-concurrent-request"]) {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("comment", flattenZtnaTrafficForwardProxyComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("h3_support", flattenZtnaTrafficForwardProxyH3Support(o["h3-support"], d, "h3_support", sv)); err != nil {
		if !fortiAPIPatch(o["h3-support"]) {
			return fmt.Errorf("Error reading h3_support: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("quic", flattenZtnaTrafficForwardProxyQuic(o["quic"], d, "quic", sv)); err != nil {
			if !fortiAPIPatch(o["quic"]) {
				return fmt.Errorf("Error reading quic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("quic"); ok {
			if err = d.Set("quic", flattenZtnaTrafficForwardProxyQuic(o["quic"], d, "quic", sv)); err != nil {
				if !fortiAPIPatch(o["quic"]) {
					return fmt.Errorf("Error reading quic: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_mode", flattenZtnaTrafficForwardProxySslMode(o["ssl-mode"], d, "ssl_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mode"]) {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_certificate", flattenZtnaTrafficForwardProxySslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-certificate"]) {
				return fmt.Errorf("Error reading ssl_certificate: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_certificate"); ok {
			if err = d.Set("ssl_certificate", flattenZtnaTrafficForwardProxySslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-certificate"]) {
					return fmt.Errorf("Error reading ssl_certificate: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_dh_bits", flattenZtnaTrafficForwardProxySslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-dh-bits"]) {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenZtnaTrafficForwardProxySslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_cipher_suites", flattenZtnaTrafficForwardProxySslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-cipher-suites"]) {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenZtnaTrafficForwardProxySslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-cipher-suites"]) {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenZtnaTrafficForwardProxySslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-algorithm"]) {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_server_cipher_suites", flattenZtnaTrafficForwardProxySslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-server-cipher-suites"]) {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenZtnaTrafficForwardProxySslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-server-cipher-suites"]) {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_pfs", flattenZtnaTrafficForwardProxySslPfs(o["ssl-pfs"], d, "ssl_pfs", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-pfs"]) {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenZtnaTrafficForwardProxySslMinVersion(o["ssl-min-version"], d, "ssl_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-version"]) {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaTrafficForwardProxySslMaxVersion(o["ssl-max-version"], d, "ssl_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenZtnaTrafficForwardProxySslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-min-version"]) {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_max_version", flattenZtnaTrafficForwardProxySslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-max-version"]) {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenZtnaTrafficForwardProxySslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-accept-ffdhe-groups"]) {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenZtnaTrafficForwardProxySslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-send-empty-frags"]) {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_client_fallback", flattenZtnaTrafficForwardProxySslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-fallback"]) {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenZtnaTrafficForwardProxySslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenZtnaTrafficForwardProxySslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-type"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenZtnaTrafficForwardProxySslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-timeout"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenZtnaTrafficForwardProxySslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-max"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenZtnaTrafficForwardProxySslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-rekey-count"]) {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_server_renegotiation", flattenZtnaTrafficForwardProxySslServerRenegotiation(o["ssl-server-renegotiation"], d, "ssl_server_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_server_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenZtnaTrafficForwardProxySslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-type"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenZtnaTrafficForwardProxySslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-timeout"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenZtnaTrafficForwardProxySslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-max"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenZtnaTrafficForwardProxySslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-http-location-conversion"]) {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenZtnaTrafficForwardProxySslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-http-match-host"]) {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenZtnaTrafficForwardProxySslHpkp(o["ssl-hpkp"], d, "ssl_hpkp", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp"]) {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenZtnaTrafficForwardProxySslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-primary"]) {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenZtnaTrafficForwardProxySslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-backup"]) {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenZtnaTrafficForwardProxySslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-age"]) {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenZtnaTrafficForwardProxySslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-report-uri"]) {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenZtnaTrafficForwardProxySslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-include-subdomains"]) {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenZtnaTrafficForwardProxySslHsts(o["ssl-hsts"], d, "ssl_hsts", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts"]) {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenZtnaTrafficForwardProxySslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts-age"]) {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenZtnaTrafficForwardProxySslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts-include-subdomains"]) {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	return nil
}

func flattenZtnaTrafficForwardProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaTrafficForwardProxyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyUserAgentDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyAuthPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyVip6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyH3Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-idle-timeout"], _ = expandZtnaTrafficForwardProxyQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-udp-payload-size"], _ = expandZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-connection-id-limit"], _ = expandZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-delay-exponent"], _ = expandZtnaTrafficForwardProxyQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-ack-delay"], _ = expandZtnaTrafficForwardProxyQuicMaxAckDelay(d, i["max_ack_delay"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-datagram-frame-size"], _ = expandZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-migration"], _ = expandZtnaTrafficForwardProxyQuicActiveMigration(d, i["active_migration"], pre_append, sv)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok {
		result["grease-quic-bit"], _ = expandZtnaTrafficForwardProxyQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append, sv)
	}

	return result, nil
}

func expandZtnaTrafficForwardProxyQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxyQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandZtnaTrafficForwardProxySslCertificateName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxySslCertificateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandZtnaTrafficForwardProxySslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cipher"], _ = expandZtnaTrafficForwardProxySslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cipher"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["versions"], _ = expandZtnaTrafficForwardProxySslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandZtnaTrafficForwardProxySslServerCipherSuitesPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cipher"], _ = expandZtnaTrafficForwardProxySslServerCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cipher"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["versions"], _ = expandZtnaTrafficForwardProxySslServerCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslPfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHsts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHstsAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaTrafficForwardProxySslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaTrafficForwardProxy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandZtnaTrafficForwardProxyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok {
		t, err := expandZtnaTrafficForwardProxyVip(d, v, "vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	} else if d.HasChange("vip") {
		obj["vip"] = nil
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandZtnaTrafficForwardProxyHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	} else if d.HasChange("host") {
		obj["host"] = nil
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		t, err := expandZtnaTrafficForwardProxyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	} else if d.HasChange("decrypted_traffic_mirror") {
		obj["decrypted-traffic-mirror"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandZtnaTrafficForwardProxyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandZtnaTrafficForwardProxyInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandZtnaTrafficForwardProxyPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	} else if d.HasChange("port") {
		obj["port"] = nil
	}

	if v, ok := d.GetOk("client_cert"); ok {
		t, err := expandZtnaTrafficForwardProxyClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok {
		t, err := expandZtnaTrafficForwardProxyUserAgentDetect(d, v, "user_agent_detect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok {
		t, err := expandZtnaTrafficForwardProxyAuthPortal(d, v, "auth_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok {
		t, err := expandZtnaTrafficForwardProxyAuthVirtualHost(d, v, "auth_virtual_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	} else if d.HasChange("auth_virtual_host") {
		obj["auth-virtual-host"] = nil
	}

	if v, ok := d.GetOk("vip6"); ok {
		t, err := expandZtnaTrafficForwardProxyVip6(d, v, "vip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip6"] = t
		}
	} else if d.HasChange("vip6") {
		obj["vip6"] = nil
	}

	if v, ok := d.GetOk("empty_cert_action"); ok {
		t, err := expandZtnaTrafficForwardProxyEmptyCertAction(d, v, "empty_cert_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok {
		t, err := expandZtnaTrafficForwardProxyLogBlockedTraffic(d, v, "log_blocked_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok {
		t, err := expandZtnaTrafficForwardProxySvrPoolMultiplex(d, v, "svr_pool_multiplex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOkExists("svr_pool_ttl"); ok {
		t, err := expandZtnaTrafficForwardProxySvrPoolTtl(d, v, "svr_pool_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOkExists("svr_pool_server_max_request"); ok {
		t, err := expandZtnaTrafficForwardProxySvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	} else if d.HasChange("svr_pool_server_max_request") {
		obj["svr-pool-server-max-request"] = nil
	}

	if v, ok := d.GetOkExists("svr_pool_server_max_concurrent_request"); ok {
		t, err := expandZtnaTrafficForwardProxySvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	} else if d.HasChange("svr_pool_server_max_concurrent_request") {
		obj["svr-pool-server-max-concurrent-request"] = nil
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandZtnaTrafficForwardProxyComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("h3_support"); ok {
		t, err := expandZtnaTrafficForwardProxyH3Support(d, v, "h3_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-support"] = t
		}
	}

	if v, ok := d.GetOk("quic"); ok {
		t, err := expandZtnaTrafficForwardProxyQuic(d, v, "quic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quic"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok {
		t, err := expandZtnaTrafficForwardProxySslMode(d, v, "ssl_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok || d.HasChange("ssl_certificate") {
		t, err := expandZtnaTrafficForwardProxySslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok {
		t, err := expandZtnaTrafficForwardProxySslDhBits(d, v, "ssl_dh_bits", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {
		t, err := expandZtnaTrafficForwardProxySslAlgorithm(d, v, "ssl_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok || d.HasChange("ssl_cipher_suites") {
		t, err := expandZtnaTrafficForwardProxySslCipherSuites(d, v, "ssl_cipher_suites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerAlgorithm(d, v, "ssl_server_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok || d.HasChange("ssl_server_cipher_suites") {
		t, err := expandZtnaTrafficForwardProxySslServerCipherSuites(d, v, "ssl_server_cipher_suites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok {
		t, err := expandZtnaTrafficForwardProxySslPfs(d, v, "ssl_pfs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok {
		t, err := expandZtnaTrafficForwardProxySslMinVersion(d, v, "ssl_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {
		t, err := expandZtnaTrafficForwardProxySslMaxVersion(d, v, "ssl_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerMinVersion(d, v, "ssl_server_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerMaxVersion(d, v, "ssl_server_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok {
		t, err := expandZtnaTrafficForwardProxySslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok {
		t, err := expandZtnaTrafficForwardProxySslSendEmptyFrags(d, v, "ssl_send_empty_frags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok {
		t, err := expandZtnaTrafficForwardProxySslClientFallback(d, v, "ssl_client_fallback", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok {
		t, err := expandZtnaTrafficForwardProxySslClientRenegotiation(d, v, "ssl_client_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok {
		t, err := expandZtnaTrafficForwardProxySslClientSessionStateType(d, v, "ssl_client_session_state_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok {
		t, err := expandZtnaTrafficForwardProxySslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok {
		t, err := expandZtnaTrafficForwardProxySslClientSessionStateMax(d, v, "ssl_client_session_state_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok {
		t, err := expandZtnaTrafficForwardProxySslClientRekeyCount(d, v, "ssl_client_rekey_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	} else if d.HasChange("ssl_client_rekey_count") {
		obj["ssl-client-rekey-count"] = nil
	}

	if v, ok := d.GetOk("ssl_server_renegotiation"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerRenegotiation(d, v, "ssl_server_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerSessionStateType(d, v, "ssl_server_session_state_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok {
		t, err := expandZtnaTrafficForwardProxySslServerSessionStateMax(d, v, "ssl_server_session_state_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok {
		t, err := expandZtnaTrafficForwardProxySslHttpLocationConversion(d, v, "ssl_http_location_conversion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok {
		t, err := expandZtnaTrafficForwardProxySslHttpMatchHost(d, v, "ssl_http_match_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok {
		t, err := expandZtnaTrafficForwardProxySslHpkp(d, v, "ssl_hpkp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok {
		t, err := expandZtnaTrafficForwardProxySslHpkpPrimary(d, v, "ssl_hpkp_primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	} else if d.HasChange("ssl_hpkp_primary") {
		obj["ssl-hpkp-primary"] = nil
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok {
		t, err := expandZtnaTrafficForwardProxySslHpkpBackup(d, v, "ssl_hpkp_backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	} else if d.HasChange("ssl_hpkp_backup") {
		obj["ssl-hpkp-backup"] = nil
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok {
		t, err := expandZtnaTrafficForwardProxySslHpkpAge(d, v, "ssl_hpkp_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok {
		t, err := expandZtnaTrafficForwardProxySslHpkpReportUri(d, v, "ssl_hpkp_report_uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	} else if d.HasChange("ssl_hpkp_report_uri") {
		obj["ssl-hpkp-report-uri"] = nil
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok {
		t, err := expandZtnaTrafficForwardProxySslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok {
		t, err := expandZtnaTrafficForwardProxySslHsts(d, v, "ssl_hsts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok {
		t, err := expandZtnaTrafficForwardProxySslHstsAge(d, v, "ssl_hsts_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok {
		t, err := expandZtnaTrafficForwardProxySslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	return &obj, nil
}
