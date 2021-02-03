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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Computed: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_server_affinity_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 60),
				Optional:     true,
				Computed:     true,
			},
			"max_waf_body_cache_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1024),
				Optional:     true,
				Computed:     true,
			},
			"webproxy_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"learn_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_client_ip_from_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learn_client_ip_srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"learn_client_ip_srcaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
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

	obj, err := getObjectWebProxyGlobal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyGlobal(obj, mkey)
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

	err := c.DeleteWebProxyGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebProxyGlobal(mkey)
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

func flattenWebProxyGlobalProxyFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalMaxRequestLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalMaxMessageLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
	return v
}

func flattenWebProxyGlobalMaxWafBodyCacheLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalWebproxyProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIpFromHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyGlobalLearnClientIpSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWebProxyGlobalLearnClientIpSrcaddrName(i["name"], d, pre_append, sv)
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

			tmp["name"] = flattenWebProxyGlobalLearnClientIpSrcaddr6Name(i["name"], d, pre_append, sv)
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

func refreshObjectWebProxyGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

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

	if err = d.Set("learn_client_ip_from_header", flattenWebProxyGlobalLearnClientIpFromHeader(o["learn-client-ip-from-header"], d, "learn_client_ip_from_header", sv)); err != nil {
		if !fortiAPIPatch(o["learn-client-ip-from-header"]) {
			return fmt.Errorf("Error reading learn_client_ip_from_header: %v", err)
		}
	}

	if isImportTable() {
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

	if isImportTable() {
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

func expandWebProxyGlobalProxyFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxRequestLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalMaxMessageLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandWebProxyGlobalLearnClientIpFromHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandWebProxyGlobalLearnClientIpSrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandWebProxyGlobalLearnClientIpSrcaddr6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyGlobalLearnClientIpSrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyGlobal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ssl_cert"); ok {

		t, err := expandWebProxyGlobalSslCert(d, v, "ssl_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cert"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ca_cert"); ok {

		t, err := expandWebProxyGlobalSslCaCert(d, v, "ssl_ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ca-cert"] = t
		}
	}

	if v, ok := d.GetOk("fast_policy_match"); ok {

		t, err := expandWebProxyGlobalFastPolicyMatch(d, v, "fast_policy_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-policy-match"] = t
		}
	}

	if v, ok := d.GetOk("proxy_fqdn"); ok {

		t, err := expandWebProxyGlobalProxyFqdn(d, v, "proxy_fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("max_request_length"); ok {

		t, err := expandWebProxyGlobalMaxRequestLength(d, v, "max_request_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-request-length"] = t
		}
	}

	if v, ok := d.GetOk("max_message_length"); ok {

		t, err := expandWebProxyGlobalMaxMessageLength(d, v, "max_message_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-message-length"] = t
		}
	}

	if v, ok := d.GetOk("strict_web_check"); ok {

		t, err := expandWebProxyGlobalStrictWebCheck(d, v, "strict_web_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-web-check"] = t
		}
	}

	if v, ok := d.GetOk("forward_proxy_auth"); ok {

		t, err := expandWebProxyGlobalForwardProxyAuth(d, v, "forward_proxy_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-proxy-auth"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_non_http"); ok {

		t, err := expandWebProxyGlobalTunnelNonHttp(d, v, "tunnel_non_http", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-non-http"] = t
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok {

		t, err := expandWebProxyGlobalUnknownHttpVersion(d, v, "unknown_http_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-http-version"] = t
		}
	}

	if v, ok := d.GetOk("forward_server_affinity_timeout"); ok {

		t, err := expandWebProxyGlobalForwardServerAffinityTimeout(d, v, "forward_server_affinity_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-server-affinity-timeout"] = t
		}
	}

	if v, ok := d.GetOk("max_waf_body_cache_length"); ok {

		t, err := expandWebProxyGlobalMaxWafBodyCacheLength(d, v, "max_waf_body_cache_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-waf-body-cache-length"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok {

		t, err := expandWebProxyGlobalWebproxyProfile(d, v, "webproxy_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip"); ok {

		t, err := expandWebProxyGlobalLearnClientIp(d, v, "learn_client_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip_from_header"); ok {

		t, err := expandWebProxyGlobalLearnClientIpFromHeader(d, v, "learn_client_ip_from_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip-from-header"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip_srcaddr"); ok {

		t, err := expandWebProxyGlobalLearnClientIpSrcaddr(d, v, "learn_client_ip_srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip-srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("learn_client_ip_srcaddr6"); ok {

		t, err := expandWebProxyGlobalLearnClientIpSrcaddr6(d, v, "learn_client_ip_srcaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-client-ip-srcaddr6"] = t
		}
	}

	return &obj, nil
}
