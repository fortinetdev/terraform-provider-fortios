// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure explicit Web proxy settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyExplicitUpdate,
		Read:   resourceWebProxyExplicitRead,
		Update: resourceWebProxyExplicitUpdate,
		Delete: resourceWebProxyExplicitDelete,

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
			"ftp_over_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"socks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_incoming_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_incoming_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_incoming_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"socks_incoming_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outgoing_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outgoing_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pref_dns_result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_http_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"realm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"sec_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_replacement_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message_upon_server_error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_file_server_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_file_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_file_server_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_file_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"pac_file_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pac_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policyid": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcaddr": &schema.Schema{
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
						"srcaddr6": &schema.Schema{
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
						"dstaddr": &schema.Schema{
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
						"pac_file_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"pac_file_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
							Optional:     true,
						},
					},
				},
			},
			"ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trace_auth_no_rsp": &schema.Schema{
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

func resourceWebProxyExplicitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyExplicit(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicit resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyExplicit(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyExplicit")
	}

	return resourceWebProxyExplicitRead(d, m)
}

func resourceWebProxyExplicitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyExplicit(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WebProxyExplicit resource while getting object: %v", err)
	}

	_, err = c.UpdateWebProxyExplicit(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WebProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyExplicitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebProxyExplicit(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyExplicit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyExplicit(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyExplicit resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyExplicitStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitFtpOverHttp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitSocks(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitHttpIncomingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitHttpsIncomingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitFtpIncomingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitSocksIncomingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitIncomingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitOutgoingIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitIpv6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitIncomingIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitOutgoingIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitStrictGuest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPrefDnsResult(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitSecDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitHttpsReplacementMessage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitMessageUponServerError(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacFileData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policyid"
		if _, ok := i["policyid"]; ok {

			tmp["policyid"] = flattenWebProxyExplicitPacPolicyPolicyid(i["policyid"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenWebProxyExplicitPacPolicyStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {

			tmp["srcaddr"] = flattenWebProxyExplicitPacPolicySrcaddr(i["srcaddr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {

			tmp["srcaddr6"] = flattenWebProxyExplicitPacPolicySrcaddr6(i["srcaddr6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {

			tmp["dstaddr"] = flattenWebProxyExplicitPacPolicyDstaddr(i["dstaddr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_name"
		if _, ok := i["pac-file-name"]; ok {

			tmp["pac_file_name"] = flattenWebProxyExplicitPacPolicyPacFileName(i["pac-file-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_data"
		if _, ok := i["pac-file-data"]; ok {

			tmp["pac_file_data"] = flattenWebProxyExplicitPacPolicyPacFileData(i["pac-file-data"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {

			tmp["comments"] = flattenWebProxyExplicitPacPolicyComments(i["comments"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "policyid", d)
	return result
}

func flattenWebProxyExplicitPacPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWebProxyExplicitPacPolicySrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebProxyExplicitPacPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWebProxyExplicitPacPolicySrcaddr6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebProxyExplicitPacPolicySrcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWebProxyExplicitPacPolicyDstaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebProxyExplicitPacPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyPacFileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyPacFileData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitPacPolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitSslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyExplicitTraceAuthNoRsp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyExplicit(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenWebProxyExplicitStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ftp_over_http", flattenWebProxyExplicitFtpOverHttp(o["ftp-over-http"], d, "ftp_over_http", sv)); err != nil {
		if !fortiAPIPatch(o["ftp-over-http"]) {
			return fmt.Errorf("Error reading ftp_over_http: %v", err)
		}
	}

	if err = d.Set("socks", flattenWebProxyExplicitSocks(o["socks"], d, "socks", sv)); err != nil {
		if !fortiAPIPatch(o["socks"]) {
			return fmt.Errorf("Error reading socks: %v", err)
		}
	}

	if err = d.Set("http_incoming_port", flattenWebProxyExplicitHttpIncomingPort(o["http-incoming-port"], d, "http_incoming_port", sv)); err != nil {
		if !fortiAPIPatch(o["http-incoming-port"]) {
			return fmt.Errorf("Error reading http_incoming_port: %v", err)
		}
	}

	if err = d.Set("https_incoming_port", flattenWebProxyExplicitHttpsIncomingPort(o["https-incoming-port"], d, "https_incoming_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-incoming-port"]) {
			return fmt.Errorf("Error reading https_incoming_port: %v", err)
		}
	}

	if err = d.Set("ftp_incoming_port", flattenWebProxyExplicitFtpIncomingPort(o["ftp-incoming-port"], d, "ftp_incoming_port", sv)); err != nil {
		if !fortiAPIPatch(o["ftp-incoming-port"]) {
			return fmt.Errorf("Error reading ftp_incoming_port: %v", err)
		}
	}

	if err = d.Set("socks_incoming_port", flattenWebProxyExplicitSocksIncomingPort(o["socks-incoming-port"], d, "socks_incoming_port", sv)); err != nil {
		if !fortiAPIPatch(o["socks-incoming-port"]) {
			return fmt.Errorf("Error reading socks_incoming_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenWebProxyExplicitIncomingIp(o["incoming-ip"], d, "incoming_ip", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-ip"]) {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenWebProxyExplicitOutgoingIp(o["outgoing-ip"], d, "outgoing_ip", sv)); err != nil {
		if !fortiAPIPatch(o["outgoing-ip"]) {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_status", flattenWebProxyExplicitIpv6Status(o["ipv6-status"], d, "ipv6_status", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-status"]) {
			return fmt.Errorf("Error reading ipv6_status: %v", err)
		}
	}

	if err = d.Set("incoming_ip6", flattenWebProxyExplicitIncomingIp6(o["incoming-ip6"], d, "incoming_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["incoming-ip6"]) {
			return fmt.Errorf("Error reading incoming_ip6: %v", err)
		}
	}

	if err = d.Set("outgoing_ip6", flattenWebProxyExplicitOutgoingIp6(o["outgoing-ip6"], d, "outgoing_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["outgoing-ip6"]) {
			return fmt.Errorf("Error reading outgoing_ip6: %v", err)
		}
	}

	if err = d.Set("strict_guest", flattenWebProxyExplicitStrictGuest(o["strict-guest"], d, "strict_guest", sv)); err != nil {
		if !fortiAPIPatch(o["strict-guest"]) {
			return fmt.Errorf("Error reading strict_guest: %v", err)
		}
	}

	if err = d.Set("pref_dns_result", flattenWebProxyExplicitPrefDnsResult(o["pref-dns-result"], d, "pref_dns_result", sv)); err != nil {
		if !fortiAPIPatch(o["pref-dns-result"]) {
			return fmt.Errorf("Error reading pref_dns_result: %v", err)
		}
	}

	if err = d.Set("unknown_http_version", flattenWebProxyExplicitUnknownHttpVersion(o["unknown-http-version"], d, "unknown_http_version", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-http-version"]) {
			return fmt.Errorf("Error reading unknown_http_version: %v", err)
		}
	}

	if err = d.Set("realm", flattenWebProxyExplicitRealm(o["realm"], d, "realm", sv)); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenWebProxyExplicitSecDefaultAction(o["sec-default-action"], d, "sec_default_action", sv)); err != nil {
		if !fortiAPIPatch(o["sec-default-action"]) {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	if err = d.Set("https_replacement_message", flattenWebProxyExplicitHttpsReplacementMessage(o["https-replacement-message"], d, "https_replacement_message", sv)); err != nil {
		if !fortiAPIPatch(o["https-replacement-message"]) {
			return fmt.Errorf("Error reading https_replacement_message: %v", err)
		}
	}

	if err = d.Set("message_upon_server_error", flattenWebProxyExplicitMessageUponServerError(o["message-upon-server-error"], d, "message_upon_server_error", sv)); err != nil {
		if !fortiAPIPatch(o["message-upon-server-error"]) {
			return fmt.Errorf("Error reading message_upon_server_error: %v", err)
		}
	}

	if err = d.Set("pac_file_server_status", flattenWebProxyExplicitPacFileServerStatus(o["pac-file-server-status"], d, "pac_file_server_status", sv)); err != nil {
		if !fortiAPIPatch(o["pac-file-server-status"]) {
			return fmt.Errorf("Error reading pac_file_server_status: %v", err)
		}
	}

	if err = d.Set("pac_file_url", flattenWebProxyExplicitPacFileUrl(o["pac-file-url"], d, "pac_file_url", sv)); err != nil {
		if !fortiAPIPatch(o["pac-file-url"]) {
			return fmt.Errorf("Error reading pac_file_url: %v", err)
		}
	}

	if err = d.Set("pac_file_server_port", flattenWebProxyExplicitPacFileServerPort(o["pac-file-server-port"], d, "pac_file_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["pac-file-server-port"]) {
			return fmt.Errorf("Error reading pac_file_server_port: %v", err)
		}
	}

	if err = d.Set("pac_file_name", flattenWebProxyExplicitPacFileName(o["pac-file-name"], d, "pac_file_name", sv)); err != nil {
		if !fortiAPIPatch(o["pac-file-name"]) {
			return fmt.Errorf("Error reading pac_file_name: %v", err)
		}
	}

	if err = d.Set("pac_file_data", flattenWebProxyExplicitPacFileData(o["pac-file-data"], d, "pac_file_data", sv)); err != nil {
		if !fortiAPIPatch(o["pac-file-data"]) {
			return fmt.Errorf("Error reading pac_file_data: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("pac_policy", flattenWebProxyExplicitPacPolicy(o["pac-policy"], d, "pac_policy", sv)); err != nil {
			if !fortiAPIPatch(o["pac-policy"]) {
				return fmt.Errorf("Error reading pac_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pac_policy"); ok {
			if err = d.Set("pac_policy", flattenWebProxyExplicitPacPolicy(o["pac-policy"], d, "pac_policy", sv)); err != nil {
				if !fortiAPIPatch(o["pac-policy"]) {
					return fmt.Errorf("Error reading pac_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_algorithm", flattenWebProxyExplicitSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("trace_auth_no_rsp", flattenWebProxyExplicitTraceAuthNoRsp(o["trace-auth-no-rsp"], d, "trace_auth_no_rsp", sv)); err != nil {
		if !fortiAPIPatch(o["trace-auth-no-rsp"]) {
			return fmt.Errorf("Error reading trace_auth_no_rsp: %v", err)
		}
	}

	return nil
}

func flattenWebProxyExplicitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyExplicitStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitFtpOverHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSocks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitHttpIncomingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitHttpsIncomingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitFtpIncomingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSocksIncomingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitIncomingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitOutgoingIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitIpv6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitIncomingIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitOutgoingIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitStrictGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPrefDnsResult(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSecDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitHttpsReplacementMessage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitMessageUponServerError(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacFileData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policyid"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["policyid"], _ = expandWebProxyExplicitPacPolicyPolicyid(d, i["policyid"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandWebProxyExplicitPacPolicyStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["srcaddr"], _ = expandWebProxyExplicitPacPolicySrcaddr(d, i["srcaddr"], pre_append, sv)
		} else {
			tmp["srcaddr"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["srcaddr6"], _ = expandWebProxyExplicitPacPolicySrcaddr6(d, i["srcaddr6"], pre_append, sv)
		} else {
			tmp["srcaddr6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["dstaddr"], _ = expandWebProxyExplicitPacPolicyDstaddr(d, i["dstaddr"], pre_append, sv)
		} else {
			tmp["dstaddr"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pac-file-name"], _ = expandWebProxyExplicitPacPolicyPacFileName(d, i["pac_file_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pac_file_data"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pac-file-data"], _ = expandWebProxyExplicitPacPolicyPacFileData(d, i["pac_file_data"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["comments"], _ = expandWebProxyExplicitPacPolicyComments(d, i["comments"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyExplicitPacPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWebProxyExplicitPacPolicySrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyExplicitPacPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWebProxyExplicitPacPolicySrcaddr6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyExplicitPacPolicySrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWebProxyExplicitPacPolicyDstaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyExplicitPacPolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyPacFileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyPacFileData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitPacPolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitSslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyExplicitTraceAuthNoRsp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyExplicit(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandWebProxyExplicitStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ftp_over_http"); ok {
		if setArgNil {
			obj["ftp-over-http"] = nil
		} else {

			t, err := expandWebProxyExplicitFtpOverHttp(d, v, "ftp_over_http", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ftp-over-http"] = t
			}
		}
	}

	if v, ok := d.GetOk("socks"); ok {
		if setArgNil {
			obj["socks"] = nil
		} else {

			t, err := expandWebProxyExplicitSocks(d, v, "socks", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["socks"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_incoming_port"); ok {
		if setArgNil {
			obj["http-incoming-port"] = nil
		} else {

			t, err := expandWebProxyExplicitHttpIncomingPort(d, v, "http_incoming_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-incoming-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_incoming_port"); ok {
		if setArgNil {
			obj["https-incoming-port"] = nil
		} else {

			t, err := expandWebProxyExplicitHttpsIncomingPort(d, v, "https_incoming_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-incoming-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("ftp_incoming_port"); ok {
		if setArgNil {
			obj["ftp-incoming-port"] = nil
		} else {

			t, err := expandWebProxyExplicitFtpIncomingPort(d, v, "ftp_incoming_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ftp-incoming-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("socks_incoming_port"); ok {
		if setArgNil {
			obj["socks-incoming-port"] = nil
		} else {

			t, err := expandWebProxyExplicitSocksIncomingPort(d, v, "socks_incoming_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["socks-incoming-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok {
		if setArgNil {
			obj["incoming-ip"] = nil
		} else {

			t, err := expandWebProxyExplicitIncomingIp(d, v, "incoming_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["incoming-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok {
		if setArgNil {
			obj["outgoing-ip"] = nil
		} else {

			t, err := expandWebProxyExplicitOutgoingIp(d, v, "outgoing_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outgoing-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipv6_status"); ok {
		if setArgNil {
			obj["ipv6-status"] = nil
		} else {

			t, err := expandWebProxyExplicitIpv6Status(d, v, "ipv6_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipv6-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("incoming_ip6"); ok {
		if setArgNil {
			obj["incoming-ip6"] = nil
		} else {

			t, err := expandWebProxyExplicitIncomingIp6(d, v, "incoming_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["incoming-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("outgoing_ip6"); ok {
		if setArgNil {
			obj["outgoing-ip6"] = nil
		} else {

			t, err := expandWebProxyExplicitOutgoingIp6(d, v, "outgoing_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["outgoing-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("strict_guest"); ok {
		if setArgNil {
			obj["strict-guest"] = nil
		} else {

			t, err := expandWebProxyExplicitStrictGuest(d, v, "strict_guest", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strict-guest"] = t
			}
		}
	}

	if v, ok := d.GetOk("pref_dns_result"); ok {
		if setArgNil {
			obj["pref-dns-result"] = nil
		} else {

			t, err := expandWebProxyExplicitPrefDnsResult(d, v, "pref_dns_result", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pref-dns-result"] = t
			}
		}
	}

	if v, ok := d.GetOk("unknown_http_version"); ok {
		if setArgNil {
			obj["unknown-http-version"] = nil
		} else {

			t, err := expandWebProxyExplicitUnknownHttpVersion(d, v, "unknown_http_version", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unknown-http-version"] = t
			}
		}
	}

	if v, ok := d.GetOk("realm"); ok {
		if setArgNil {
			obj["realm"] = nil
		} else {

			t, err := expandWebProxyExplicitRealm(d, v, "realm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["realm"] = t
			}
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok {
		if setArgNil {
			obj["sec-default-action"] = nil
		} else {

			t, err := expandWebProxyExplicitSecDefaultAction(d, v, "sec_default_action", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sec-default-action"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_replacement_message"); ok {
		if setArgNil {
			obj["https-replacement-message"] = nil
		} else {

			t, err := expandWebProxyExplicitHttpsReplacementMessage(d, v, "https_replacement_message", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-replacement-message"] = t
			}
		}
	}

	if v, ok := d.GetOk("message_upon_server_error"); ok {
		if setArgNil {
			obj["message-upon-server-error"] = nil
		} else {

			t, err := expandWebProxyExplicitMessageUponServerError(d, v, "message_upon_server_error", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["message-upon-server-error"] = t
			}
		}
	}

	if v, ok := d.GetOk("pac_file_server_status"); ok {
		if setArgNil {
			obj["pac-file-server-status"] = nil
		} else {

			t, err := expandWebProxyExplicitPacFileServerStatus(d, v, "pac_file_server_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pac-file-server-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("pac_file_url"); ok {
		if setArgNil {
			obj["pac-file-url"] = nil
		} else {

			t, err := expandWebProxyExplicitPacFileUrl(d, v, "pac_file_url", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pac-file-url"] = t
			}
		}
	}

	if v, ok := d.GetOk("pac_file_server_port"); ok {
		if setArgNil {
			obj["pac-file-server-port"] = nil
		} else {

			t, err := expandWebProxyExplicitPacFileServerPort(d, v, "pac_file_server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pac-file-server-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("pac_file_name"); ok {
		if setArgNil {
			obj["pac-file-name"] = nil
		} else {

			t, err := expandWebProxyExplicitPacFileName(d, v, "pac_file_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pac-file-name"] = t
			}
		}
	}

	if v, ok := d.GetOk("pac_file_data"); ok {
		if setArgNil {
			obj["pac-file-data"] = nil
		} else {

			t, err := expandWebProxyExplicitPacFileData(d, v, "pac_file_data", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pac-file-data"] = t
			}
		}
	}

	if v, ok := d.GetOk("pac_policy"); ok || d.HasChange("pac_policy") {
		if setArgNil {
			obj["pac-policy"] = make([]struct{}, 0)
		} else {

			t, err := expandWebProxyExplicitPacPolicy(d, v, "pac_policy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pac-policy"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {
		if setArgNil {
			obj["ssl-algorithm"] = nil
		} else {

			t, err := expandWebProxyExplicitSslAlgorithm(d, v, "ssl_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssl-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("trace_auth_no_rsp"); ok {
		if setArgNil {
			obj["trace-auth-no-rsp"] = nil
		} else {

			t, err := expandWebProxyExplicitTraceAuthNoRsp(d, v, "trace_auth_no_rsp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trace-auth-no-rsp"] = t
			}
		}
	}

	return &obj, nil
}
