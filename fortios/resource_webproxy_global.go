// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Web proxy global settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebProxyGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyGlobalUpdate,
		Read:   resourceWebProxyGlobalRead,
		Update: resourceWebProxyGlobalUpdate,
		Delete: resourceWebProxyGlobalDelete,

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
			"ssl_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fast_policy_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_user_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Required:     true,
			},
			"max_request_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 64),
				Optional:     true,
				Computed:     true,
			},
			"max_message_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(16, 256),
				Optional:     true,
				Computed:     true,
			},
			"http2_client_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"http2_server_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auth_sign_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 3600),
				Optional:     true,
				Computed:     true,
			},
			"strict_web_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_proxy_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_non_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forward_server_affinity_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 60),
				Optional:     true,
				Computed:     true,
			},
			"max_waf_body_cache_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1024),
				Optional:     true,
				Computed:     true,
			},
			"webproxy_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"learn_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"always_learn_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_client_ip_from_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"learn_client_ip_srcaddr": &schema.Schema{
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
			"learn_client_ip_srcaddr6": &schema.Schema{
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
			"src_affinity_exempt_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_affinity_exempt_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_partial_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_category_deep_inspect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_policy_pending": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_forward_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_app_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_transparent_cert_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_obs_fold": &schema.Schema{
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

func resourceWebProxyGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyGlobal")
	}

	return resourceWebProxyGlobalRead(d, m)
}

func resourceWebProxyGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WebProxyGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateWebProxyGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WebProxyGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyGlobal resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyGlobalSslCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalSslCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalFastPolicyMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLdapUserCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalProxyFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalMaxRequestLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalMaxMessageLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalHttp2ClientWindowSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalHttp2ServerWindowSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalAuthSignTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalStrictWebCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalForwardProxyAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalTunnelNonHttp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalForwardServerAffinityTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalMaxWafBodyCacheLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebProxyGlobalWebproxyProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalAlwaysLearnClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIpFromHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIpSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenWebProxyGlobalLearnClientIpSrcaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebProxyGlobalLearnClientIpSrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIpSrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenWebProxyGlobalLearnClientIpSrcaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebProxyGlobalLearnClientIpSrcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalSrcAffinityExemptAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalSrcAffinityExemptAddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalPolicyPartialMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalPolicyCategoryDeepInspect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLogPolicyPending(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLogForwardServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLogAppId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalProxyTransparentCertInspection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalRequestObsFold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("ssl_cert", flattenWebProxyGlobalSslCert(o["ssl-cert"], d, "ssl_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-cert"]) {
			return fmt.Errorf("Error reading ssl_cert: %v", err)
		}
	}

	if err = d.Set("ssl_ca_cert", flattenWebProxyGlobalSslCaCert(o["ssl-ca-cert"], d, "ssl_ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ca-cert"]) {
			return fmt.Errorf("Error reading ssl_ca_cert: %v", err)
		}
	}

	if err = d.Set("fast_policy_match", flattenWebProxyGlobalFastPolicyMatch(o["fast-policy-match"], d, "fast_policy_match", sv)); err != nil {
		if !fortiAPIPatch(o["fast-policy-match"]) {
			return fmt.Errorf("Error reading fast_policy_match: %v", err)
		}
	}

	if err = d.Set("ldap_user_cache", flattenWebProxyGlobalLdapUserCache(o["ldap-user-cache"], d, "ldap_user_cache", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-user-cache"]) {
			return fmt.Errorf("Error reading ldap_user_cache: %v", err)
		}
	}

	if err = d.Set("proxy_fqdn", flattenWebProxyGlobalProxyFqdn(o["proxy-fqdn"], d, "proxy_fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-fqdn"]) {
			return fmt.Errorf("Error reading proxy_fqdn: %v", err)
		}
	}

	if err = d.Set("max_request_length", flattenWebProxyGlobalMaxRequestLength(o["max-request-length"], d, "max_request_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-request-length"]) {
			return fmt.Errorf("Error reading max_request_length: %v", err)
		}
	}

	if err = d.Set("max_message_length", flattenWebProxyGlobalMaxMessageLength(o["max-message-length"], d, "max_message_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-message-length"]) {
			return fmt.Errorf("Error reading max_message_length: %v", err)
		}
	}

	if err = d.Set("http2_client_window_size", flattenWebProxyGlobalHttp2ClientWindowSize(o["http2-client-window-size"], d, "http2_client_window_size", sv)); err != nil {
		if !fortiAPIPatch(o["http2-client-window-size"]) {
			return fmt.Errorf("Error reading http2_client_window_size: %v", err)
		}
	}

	if err = d.Set("http2_server_window_size", flattenWebProxyGlobalHttp2ServerWindowSize(o["http2-server-window-size"], d, "http2_server_window_size", sv)); err != nil {
		if !fortiAPIPatch(o["http2-server-window-size"]) {
			return fmt.Errorf("Error reading http2_server_window_size: %v", err)
		}
	}

	if err = d.Set("auth_sign_timeout", flattenWebProxyGlobalAuthSignTimeout(o["auth-sign-timeout"], d, "auth_sign_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auth-sign-timeout"]) {
			return fmt.Errorf("Error reading auth_sign_timeout: %v", err)
		}
	}

	if err = d.Set("strict_web_check", flattenWebProxyGlobalStrictWebCheck(o["strict-web-check"], d, "strict_web_check", sv)); err != nil {
		if !fortiAPIPatch(o["strict-web-check"]) {
			return fmt.Errorf("Error reading strict_web_check: %v", err)
		}
	}

	if err = d.Set("forward_proxy_auth", flattenWebProxyGlobalForwardProxyAuth(o["forward-proxy-auth"], d, "forward_proxy_auth", sv)); err != nil {
		if !fortiAPIPatch(o["forward-proxy-auth"]) {
			return fmt.Errorf("Error reading forward_proxy_auth: %v", err)
		}
	}

	if err = d.Set("tunnel_non_http", flattenWebProxyGlobalTunnelNonHttp(o["tunnel-non-http"], d, "tunnel_non_http", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-non-http"]) {
			return fmt.Errorf("Error reading tunnel_non_http: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenWebProxyGlobalUnknownHttpVersion(o["unknown-http-version"], d, "unknown_http_version", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-http-version"]) {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("forward_server_affinity_timeout", flattenWebProxyGlobalForwardServerAffinityTimeout(o["forward-server-affinity-timeout"], d, "forward_server_affinity_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["forward-server-affinity-timeout"]) {
			return fmt.Errorf("Error reading forward_server_affinity_timeout: %v", err)
		}
	}

	if err = d.Set("max_waf_body_cache_length", flattenWebProxyGlobalMaxWafBodyCacheLength(o["max-waf-body-cache-length"], d, "max_waf_body_cache_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-waf-body-cache-length"]) {
			return fmt.Errorf("Error reading max_waf_body_cache_length: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenWebProxyGlobalWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-profile"]) {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("learn_client_ip", flattenWebProxyGlobalLearnClientIp(o["learn-client-ip"], d, "learn_client_ip", sv)); err != nil {
		if !fortiAPIPatch(o["learn-client-ip"]) {
			return fmt.Errorf("Error reading learn_client_ip: %v", err)
		}
	}

	if err = d.Set("always_learn_client_ip", flattenWebProxyGlobalAlwaysLearnClientIp(o["always-learn-client-ip"], d, "always_learn_client_ip", sv)); err != nil {
		if !fortiAPIPatch(o["always-learn-client-ip"]) {
			return fmt.Errorf("Error reading always_learn_client_ip: %v", err)
		}
	}

	if err = d.Set("learn_client_ip_from_header", flattenWebProxyGlobalLearnClientIpFromHeader(o["learn-client-ip-from-header"], d, "learn_client_ip_from_header", sv)); err != nil {
		if !fortiAPIPatch(o["learn-client-ip-from-header"]) {
			return fmt.Errorf("Error reading learn_client_ip_from_header: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("learn_client_ip_srcaddr", flattenWebProxyGlobalLearnClientIpSrcaddr(o["learn-client-ip-srcaddr"], d, "learn_client_ip_srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["learn-client-ip-srcaddr"]) {
				return fmt.Errorf("Error reading learn_client_ip_srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("learn_client_ip_srcaddr"); ok {
			if err = d.Set("learn_client_ip_srcaddr", flattenWebProxyGlobalLearnClientIpSrcaddr(o["learn-client-ip-srcaddr"], d, "learn_client_ip_srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["learn-client-ip-srcaddr"]) {
					return fmt.Errorf("Error reading learn_client_ip_srcaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("learn_client_ip_srcaddr6", flattenWebProxyGlobalLearnClientIpSrcaddr6(o["learn-client-ip-srcaddr6"], d, "learn_client_ip_srcaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["learn-client-ip-srcaddr6"]) {
				return fmt.Errorf("Error reading learn_client_ip_srcaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("learn_client_ip_srcaddr6"); ok {
			if err = d.Set("learn_client_ip_srcaddr6", flattenWebProxyGlobalLearnClientIpSrcaddr6(o["learn-client-ip-srcaddr6"], d, "learn_client_ip_srcaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["learn-client-ip-srcaddr6"]) {
					return fmt.Errorf("Error reading learn_client_ip_srcaddr6: %v", err)
				}
			}
		}
	}

	if err = d.Set("src_affinity_exempt_addr", flattenWebProxyGlobalSrcAffinityExemptAddr(o["src-affinity-exempt-addr"], d, "src_affinity_exempt_addr", sv)); err != nil {
		if !fortiAPIPatch(o["src-affinity-exempt-addr"]) {
			return fmt.Errorf("Error reading src_affinity_exempt_addr: %v", err)
		}
	}

	if err = d.Set("src_affinity_exempt_addr6", flattenWebProxyGlobalSrcAffinityExemptAddr6(o["src-affinity-exempt-addr6"], d, "src_affinity_exempt_addr6", sv)); err != nil {
		if !fortiAPIPatch(o["src-affinity-exempt-addr6"]) {
			return fmt.Errorf("Error reading src_affinity_exempt_addr6: %v", err)
		}
	}

	if err = d.Set("policy_partial_match", flattenWebProxyGlobalPolicyPartialMatch(o["policy-partial-match"], d, "policy_partial_match", sv)); err != nil {
		if !fortiAPIPatch(o["policy-partial-match"]) {
			return fmt.Errorf("Error reading policy_partial_match: %v", err)
		}
	}

	if err = d.Set("policy_category_deep_inspect", flattenWebProxyGlobalPolicyCategoryDeepInspect(o["policy-category-deep-inspect"], d, "policy_category_deep_inspect", sv)); err != nil {
		if !fortiAPIPatch(o["policy-category-deep-inspect"]) {
			return fmt.Errorf("Error reading policy_category_deep_inspect: %v", err)
		}
	}

	if err = d.Set("log_policy_pending", flattenWebProxyGlobalLogPolicyPending(o["log-policy-pending"], d, "log_policy_pending", sv)); err != nil {
		if !fortiAPIPatch(o["log-policy-pending"]) {
			return fmt.Errorf("Error reading log_policy_pending: %v", err)
		}
	}

	if err = d.Set("log_forward_server", flattenWebProxyGlobalLogForwardServer(o["log-forward-server"], d, "log_forward_server", sv)); err != nil {
		if !fortiAPIPatch(o["log-forward-server"]) {
			return fmt.Errorf("Error reading log_forward_server: %v", err)
		}
	}

	if err = d.Set("log_app_id", flattenWebProxyGlobalLogAppId(o["log-app-id"], d, "log_app_id", sv)); err != nil {
		if !fortiAPIPatch(o["log-app-id"]) {
			return fmt.Errorf("Error reading log_app_id: %v", err)
		}
	}

	if err = d.Set("proxy_transparent_cert_inspection", flattenWebProxyGlobalProxyTransparentCertInspection(o["proxy-transparent-cert-inspection"], d, "proxy_transparent_cert_inspection", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-transparent-cert-inspection"]) {
			return fmt.Errorf("Error reading proxy_transparent_cert_inspection: %v", err)
		}
	}

	if err = d.Set("request_obs_fold", flattenWebProxyGlobalRequestObsFold(o["request-obs-fold"], d, "request_obs_fold", sv)); err != nil {
		if !fortiAPIPatch(o["request-obs-fold"]) {
			return fmt.Errorf("Error reading request_obs_fold: %v", err)
		}
	}

	return nil
}

func flattenWebProxyGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyGlobalSslCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalSslCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalFastPolicyMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLdapUserCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalProxyFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxRequestLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxMessageLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalHttp2ClientWindowSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalHttp2ServerWindowSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalAuthSignTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalStrictWebCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalForwardProxyAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalTunnelNonHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalForwardServerAffinityTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxWafBodyCacheLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalWebproxyProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalAlwaysLearnClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIpFromHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebProxyGlobalLearnClientIpSrcaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebProxyGlobalLearnClientIpSrcaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalSrcAffinityExemptAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalSrcAffinityExemptAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalPolicyPartialMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalPolicyCategoryDeepInspect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLogPolicyPending(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLogForwardServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLogAppId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalProxyTransparentCertInspection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalRequestObsFold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ssl_cert"); ok {
		if setArgNil {
			obj["ssl-cert"] = nil
		} else {
			t, err := expandWebProxyGlobalSslCert(d, v, "ssl_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_ca_cert"); ok {
		if setArgNil {
			obj["ssl-ca-cert"] = nil
		} else {
			t, err := expandWebProxyGlobalSslCaCert(d, v, "ssl_ca_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-ca-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("fast_policy_match"); ok {
		if setArgNil {
			obj["fast-policy-match"] = nil
		} else {
			t, err := expandWebProxyGlobalFastPolicyMatch(d, v, "fast_policy_match", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fast-policy-match"] = t
			}
		}
	}

	if v, ok := d.GetOk("ldap_user_cache"); ok {
		if setArgNil {
			obj["ldap-user-cache"] = nil
		} else {
			t, err := expandWebProxyGlobalLdapUserCache(d, v, "ldap_user_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ldap-user-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_fqdn"); ok {
		if setArgNil {
			obj["proxy-fqdn"] = nil
		} else {
			t, err := expandWebProxyGlobalProxyFqdn(d, v, "proxy_fqdn", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-fqdn"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_request_length"); ok {
		if setArgNil {
			obj["max-request-length"] = nil
		} else {
			t, err := expandWebProxyGlobalMaxRequestLength(d, v, "max_request_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-request-length"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_message_length"); ok {
		if setArgNil {
			obj["max-message-length"] = nil
		} else {
			t, err := expandWebProxyGlobalMaxMessageLength(d, v, "max_message_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-message-length"] = t
			}
		}
	}

	if v, ok := d.GetOk("http2_client_window_size"); ok {
		if setArgNil {
			obj["http2-client-window-size"] = nil
		} else {
			t, err := expandWebProxyGlobalHttp2ClientWindowSize(d, v, "http2_client_window_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http2-client-window-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("http2_server_window_size"); ok {
		if setArgNil {
			obj["http2-server-window-size"] = nil
		} else {
			t, err := expandWebProxyGlobalHttp2ServerWindowSize(d, v, "http2_server_window_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http2-server-window-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_sign_timeout"); ok {
		if setArgNil {
			obj["auth-sign-timeout"] = nil
		} else {
			t, err := expandWebProxyGlobalAuthSignTimeout(d, v, "auth_sign_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-sign-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("strict_web_check"); ok {
		if setArgNil {
			obj["strict-web-check"] = nil
		} else {
			t, err := expandWebProxyGlobalStrictWebCheck(d, v, "strict_web_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strict-web-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("forward_proxy_auth"); ok {
		if setArgNil {
			obj["forward-proxy-auth"] = nil
		} else {
			t, err := expandWebProxyGlobalForwardProxyAuth(d, v, "forward_proxy_auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward-proxy-auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_non_http"); ok {
		if setArgNil {
			obj["tunnel-non-http"] = nil
		} else {
			t, err := expandWebProxyGlobalTunnelNonHttp(d, v, "tunnel_non_http", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-non-http"] = t
			}
		}
	} else if d.HasChange("tunnel_non_http") {
		obj["tunnel-non-http"] = nil
	}

	if v, ok := d.GetOk("unknown_http_version"); ok {
		if setArgNil {
			obj["unknown-http-version"] = nil
		} else {
			t, err := expandWebProxyGlobalUnknownHttpVersion(d, v, "unknown_http_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unknown-http-version"] = t
			}
		}
	} else if d.HasChange("unknown_http_version") {
		obj["unknown-http-version"] = nil
	}

	if v, ok := d.GetOk("forward_server_affinity_timeout"); ok {
		if setArgNil {
			obj["forward-server-affinity-timeout"] = nil
		} else {
			t, err := expandWebProxyGlobalForwardServerAffinityTimeout(d, v, "forward_server_affinity_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward-server-affinity-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_waf_body_cache_length"); ok {
		if setArgNil {
			obj["max-waf-body-cache-length"] = nil
		} else {
			t, err := expandWebProxyGlobalMaxWafBodyCacheLength(d, v, "max_waf_body_cache_length", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-waf-body-cache-length"] = t
			}
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok {
		if setArgNil {
			obj["webproxy-profile"] = nil
		} else {
			t, err := expandWebProxyGlobalWebproxyProfile(d, v, "webproxy_profile", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["webproxy-profile"] = t
			}
		}
	} else if d.HasChange("webproxy_profile") {
		obj["webproxy-profile"] = nil
	}

	if v, ok := d.GetOk("learn_client_ip"); ok {
		if setArgNil {
			obj["learn-client-ip"] = nil
		} else {
			t, err := expandWebProxyGlobalLearnClientIp(d, v, "learn_client_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["learn-client-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("always_learn_client_ip"); ok {
		if setArgNil {
			obj["always-learn-client-ip"] = nil
		} else {
			t, err := expandWebProxyGlobalAlwaysLearnClientIp(d, v, "always_learn_client_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["always-learn-client-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("learn_client_ip_from_header"); ok {
		if setArgNil {
			obj["learn-client-ip-from-header"] = nil
		} else {
			t, err := expandWebProxyGlobalLearnClientIpFromHeader(d, v, "learn_client_ip_from_header", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["learn-client-ip-from-header"] = t
			}
		}
	} else if d.HasChange("learn_client_ip_from_header") {
		obj["learn-client-ip-from-header"] = nil
	}

	if v, ok := d.GetOk("learn_client_ip_srcaddr"); ok || d.HasChange("learn_client_ip_srcaddr") {
		if setArgNil {
			obj["learn-client-ip-srcaddr"] = make([]struct{}, 0)
		} else {
			t, err := expandWebProxyGlobalLearnClientIpSrcaddr(d, v, "learn_client_ip_srcaddr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["learn-client-ip-srcaddr"] = t
			}
		}
	}

	if v, ok := d.GetOk("learn_client_ip_srcaddr6"); ok || d.HasChange("learn_client_ip_srcaddr6") {
		if setArgNil {
			obj["learn-client-ip-srcaddr6"] = make([]struct{}, 0)
		} else {
			t, err := expandWebProxyGlobalLearnClientIpSrcaddr6(d, v, "learn_client_ip_srcaddr6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["learn-client-ip-srcaddr6"] = t
			}
		}
	}

	if v, ok := d.GetOk("src_affinity_exempt_addr"); ok {
		if setArgNil {
			obj["src-affinity-exempt-addr"] = nil
		} else {
			t, err := expandWebProxyGlobalSrcAffinityExemptAddr(d, v, "src_affinity_exempt_addr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["src-affinity-exempt-addr"] = t
			}
		}
	} else if d.HasChange("src_affinity_exempt_addr") {
		obj["src-affinity-exempt-addr"] = nil
	}

	if v, ok := d.GetOk("src_affinity_exempt_addr6"); ok {
		if setArgNil {
			obj["src-affinity-exempt-addr6"] = nil
		} else {
			t, err := expandWebProxyGlobalSrcAffinityExemptAddr6(d, v, "src_affinity_exempt_addr6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["src-affinity-exempt-addr6"] = t
			}
		}
	} else if d.HasChange("src_affinity_exempt_addr6") {
		obj["src-affinity-exempt-addr6"] = nil
	}

	if v, ok := d.GetOk("policy_partial_match"); ok {
		if setArgNil {
			obj["policy-partial-match"] = nil
		} else {
			t, err := expandWebProxyGlobalPolicyPartialMatch(d, v, "policy_partial_match", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["policy-partial-match"] = t
			}
		}
	}

	if v, ok := d.GetOk("policy_category_deep_inspect"); ok {
		if setArgNil {
			obj["policy-category-deep-inspect"] = nil
		} else {
			t, err := expandWebProxyGlobalPolicyCategoryDeepInspect(d, v, "policy_category_deep_inspect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["policy-category-deep-inspect"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_policy_pending"); ok {
		if setArgNil {
			obj["log-policy-pending"] = nil
		} else {
			t, err := expandWebProxyGlobalLogPolicyPending(d, v, "log_policy_pending", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-policy-pending"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_forward_server"); ok {
		if setArgNil {
			obj["log-forward-server"] = nil
		} else {
			t, err := expandWebProxyGlobalLogForwardServer(d, v, "log_forward_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-forward-server"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_app_id"); ok {
		if setArgNil {
			obj["log-app-id"] = nil
		} else {
			t, err := expandWebProxyGlobalLogAppId(d, v, "log_app_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-app-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("proxy_transparent_cert_inspection"); ok {
		if setArgNil {
			obj["proxy-transparent-cert-inspection"] = nil
		} else {
			t, err := expandWebProxyGlobalProxyTransparentCertInspection(d, v, "proxy_transparent_cert_inspection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy-transparent-cert-inspection"] = t
			}
		}
	}

	if v, ok := d.GetOk("request_obs_fold"); ok {
		if setArgNil {
			obj["request-obs-fold"] = nil
		} else {
			t, err := expandWebProxyGlobalRequestObsFold(d, v, "request_obs_fold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["request-obs-fold"] = t
			}
		}
	}

	return &obj, nil
}
