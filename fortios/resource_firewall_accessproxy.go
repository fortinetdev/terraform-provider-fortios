// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 access proxy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxyCreate,
		Read:   resourceFirewallAccessProxyRead,
		Update: resourceFirewallAccessProxyUpdate,
		Delete: resourceFirewallAccessProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"vip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"client_cert": &schema.Schema{
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
				Computed:     true,
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
			"decrypted_traffic_mirror": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"api_gateway": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"url_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtual_host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"url_map_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"realservers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"addr_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"address": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"domain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"mappedport": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
										Computed:     true,
									},
									"http_host": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"health_check": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"health_check_proto": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"holddown_interval": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_client_cert": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
									"ssh_host_key_validation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_host_key": &schema.Schema{
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
								},
							},
						},
						"persistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_cookie_domain_from_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_cookie_domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_path": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_generation": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"http_cookie_age": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 525600),
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_share": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_cookie_secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"saml_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"saml_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
										Computed: true,
									},
									"cipher": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"versions": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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
						"ssl_vpn_web_portal": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"api_gateway6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"url_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
						"service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ldb_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtual_host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"url_map_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"realservers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"addr_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"address": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"domain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"mappedport": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
										Computed:     true,
									},
									"http_host": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"health_check": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"health_check_proto": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"holddown_interval": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_client_cert": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
									"ssh_host_key_validation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_host_key": &schema.Schema{
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
								},
							},
						},
						"persistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_cookie_domain_from_host": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_cookie_domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_path": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_generation": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"http_cookie_age": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 525600),
							Optional:     true,
							Computed:     true,
						},
						"http_cookie_share": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_cookie_secure": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"saml_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"saml_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
										Computed: true,
									},
									"cipher": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"versions": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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
						"ssl_vpn_web_portal": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
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

func resourceFirewallAccessProxyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallAccessProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallAccessProxy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxy")
	}

	return resourceFirewallAccessProxyRead(d, m)
}

func resourceFirewallAccessProxyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallAccessProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAccessProxy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxy")
	}

	return resourceFirewallAccessProxyRead(d, m)
}

func resourceFirewallAccessProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallAccessProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallAccessProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyAuthPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyEmptyCertAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallAccessProxyApiGatewayId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {

			tmp["url_map"] = flattenFirewallAccessProxyApiGatewayUrlMap(i["url-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {

			tmp["service"] = flattenFirewallAccessProxyApiGatewayService(i["service"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {

			tmp["ldb_method"] = flattenFirewallAccessProxyApiGatewayLdbMethod(i["ldb-method"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {

			tmp["virtual_host"] = flattenFirewallAccessProxyApiGatewayVirtualHost(i["virtual-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {

			tmp["url_map_type"] = flattenFirewallAccessProxyApiGatewayUrlMapType(i["url-map-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {

			tmp["realservers"] = flattenFirewallAccessProxyApiGatewayRealservers(i["realservers"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {

			tmp["persistence"] = flattenFirewallAccessProxyApiGatewayPersistence(i["persistence"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {

			tmp["http_cookie_domain_from_host"] = flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {

			tmp["http_cookie_domain"] = flattenFirewallAccessProxyApiGatewayHttpCookieDomain(i["http-cookie-domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {

			tmp["http_cookie_path"] = flattenFirewallAccessProxyApiGatewayHttpCookiePath(i["http-cookie-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {

			tmp["http_cookie_generation"] = flattenFirewallAccessProxyApiGatewayHttpCookieGeneration(i["http-cookie-generation"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {

			tmp["http_cookie_age"] = flattenFirewallAccessProxyApiGatewayHttpCookieAge(i["http-cookie-age"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {

			tmp["http_cookie_share"] = flattenFirewallAccessProxyApiGatewayHttpCookieShare(i["http-cookie-share"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {

			tmp["https_cookie_secure"] = flattenFirewallAccessProxyApiGatewayHttpsCookieSecure(i["https-cookie-secure"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {

			tmp["saml_server"] = flattenFirewallAccessProxyApiGatewaySamlServer(i["saml-server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := i["saml-redirect"]; ok {

			tmp["saml_redirect"] = flattenFirewallAccessProxyApiGatewaySamlRedirect(i["saml-redirect"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {

			tmp["ssl_dh_bits"] = flattenFirewallAccessProxyApiGatewaySslDhBits(i["ssl-dh-bits"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {

			tmp["ssl_algorithm"] = flattenFirewallAccessProxyApiGatewaySslAlgorithm(i["ssl-algorithm"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {

			tmp["ssl_cipher_suites"] = flattenFirewallAccessProxyApiGatewaySslCipherSuites(i["ssl-cipher-suites"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {

			tmp["ssl_min_version"] = flattenFirewallAccessProxyApiGatewaySslMinVersion(i["ssl-min-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {

			tmp["ssl_max_version"] = flattenFirewallAccessProxyApiGatewaySslMaxVersion(i["ssl-max-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := i["ssl-vpn-web-portal"]; ok {

			tmp["ssl_vpn_web_portal"] = flattenFirewallAccessProxyApiGatewaySslVpnWebPortal(i["ssl-vpn-web-portal"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxyApiGatewayId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallAccessProxyApiGatewayRealserversId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {

			tmp["addr_type"] = flattenFirewallAccessProxyApiGatewayRealserversAddrType(i["addr-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {

			tmp["address"] = flattenFirewallAccessProxyApiGatewayRealserversAddress(i["address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallAccessProxyApiGatewayRealserversIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {

			tmp["domain"] = flattenFirewallAccessProxyApiGatewayRealserversDomain(i["domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallAccessProxyApiGatewayRealserversPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {

			tmp["mappedport"] = flattenFirewallAccessProxyApiGatewayRealserversMappedport(i["mappedport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallAccessProxyApiGatewayRealserversStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenFirewallAccessProxyApiGatewayRealserversType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenFirewallAccessProxyApiGatewayRealserversWeight(i["weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {

			tmp["http_host"] = flattenFirewallAccessProxyApiGatewayRealserversHttpHost(i["http-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenFirewallAccessProxyApiGatewayRealserversHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {

			tmp["health_check_proto"] = flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(i["health-check-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {

			tmp["holddown_interval"] = flattenFirewallAccessProxyApiGatewayRealserversHolddownInterval(i["holddown-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {

			tmp["ssh_client_cert"] = flattenFirewallAccessProxyApiGatewayRealserversSshClientCert(i["ssh-client-cert"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {

			tmp["ssh_host_key_validation"] = flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(i["ssh-host-key-validation"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {

			tmp["ssh_host_key"] = flattenFirewallAccessProxyApiGatewayRealserversSshHostKey(i["ssh-host-key"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxyApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversMappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAccessProxyApiGatewayRealserversSshHostKeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySamlRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGatewaySslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallAccessProxyApiGateway6Id(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := i["url-map"]; ok {

			tmp["url_map"] = flattenFirewallAccessProxyApiGateway6UrlMap(i["url-map"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {

			tmp["service"] = flattenFirewallAccessProxyApiGateway6Service(i["service"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := i["ldb-method"]; ok {

			tmp["ldb_method"] = flattenFirewallAccessProxyApiGateway6LdbMethod(i["ldb-method"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := i["virtual-host"]; ok {

			tmp["virtual_host"] = flattenFirewallAccessProxyApiGateway6VirtualHost(i["virtual-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := i["url-map-type"]; ok {

			tmp["url_map_type"] = flattenFirewallAccessProxyApiGateway6UrlMapType(i["url-map-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := i["realservers"]; ok {

			tmp["realservers"] = flattenFirewallAccessProxyApiGateway6Realservers(i["realservers"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := i["persistence"]; ok {

			tmp["persistence"] = flattenFirewallAccessProxyApiGateway6Persistence(i["persistence"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := i["http-cookie-domain-from-host"]; ok {

			tmp["http_cookie_domain_from_host"] = flattenFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(i["http-cookie-domain-from-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := i["http-cookie-domain"]; ok {

			tmp["http_cookie_domain"] = flattenFirewallAccessProxyApiGateway6HttpCookieDomain(i["http-cookie-domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := i["http-cookie-path"]; ok {

			tmp["http_cookie_path"] = flattenFirewallAccessProxyApiGateway6HttpCookiePath(i["http-cookie-path"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := i["http-cookie-generation"]; ok {

			tmp["http_cookie_generation"] = flattenFirewallAccessProxyApiGateway6HttpCookieGeneration(i["http-cookie-generation"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := i["http-cookie-age"]; ok {

			tmp["http_cookie_age"] = flattenFirewallAccessProxyApiGateway6HttpCookieAge(i["http-cookie-age"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := i["http-cookie-share"]; ok {

			tmp["http_cookie_share"] = flattenFirewallAccessProxyApiGateway6HttpCookieShare(i["http-cookie-share"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := i["https-cookie-secure"]; ok {

			tmp["https_cookie_secure"] = flattenFirewallAccessProxyApiGateway6HttpsCookieSecure(i["https-cookie-secure"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := i["saml-server"]; ok {

			tmp["saml_server"] = flattenFirewallAccessProxyApiGateway6SamlServer(i["saml-server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := i["saml-redirect"]; ok {

			tmp["saml_redirect"] = flattenFirewallAccessProxyApiGateway6SamlRedirect(i["saml-redirect"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := i["ssl-dh-bits"]; ok {

			tmp["ssl_dh_bits"] = flattenFirewallAccessProxyApiGateway6SslDhBits(i["ssl-dh-bits"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := i["ssl-algorithm"]; ok {

			tmp["ssl_algorithm"] = flattenFirewallAccessProxyApiGateway6SslAlgorithm(i["ssl-algorithm"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := i["ssl-cipher-suites"]; ok {

			tmp["ssl_cipher_suites"] = flattenFirewallAccessProxyApiGateway6SslCipherSuites(i["ssl-cipher-suites"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := i["ssl-min-version"]; ok {

			tmp["ssl_min_version"] = flattenFirewallAccessProxyApiGateway6SslMinVersion(i["ssl-min-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := i["ssl-max-version"]; ok {

			tmp["ssl_max_version"] = flattenFirewallAccessProxyApiGateway6SslMaxVersion(i["ssl-max-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := i["ssl-vpn-web-portal"]; ok {

			tmp["ssl_vpn_web_portal"] = flattenFirewallAccessProxyApiGateway6SslVpnWebPortal(i["ssl-vpn-web-portal"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxyApiGateway6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6UrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Service(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6LdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6VirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6UrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Realservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallAccessProxyApiGateway6RealserversId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {

			tmp["addr_type"] = flattenFirewallAccessProxyApiGateway6RealserversAddrType(i["addr-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {

			tmp["address"] = flattenFirewallAccessProxyApiGateway6RealserversAddress(i["address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallAccessProxyApiGateway6RealserversIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {

			tmp["domain"] = flattenFirewallAccessProxyApiGateway6RealserversDomain(i["domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallAccessProxyApiGateway6RealserversPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := i["mappedport"]; ok {

			tmp["mappedport"] = flattenFirewallAccessProxyApiGateway6RealserversMappedport(i["mappedport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallAccessProxyApiGateway6RealserversStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenFirewallAccessProxyApiGateway6RealserversType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenFirewallAccessProxyApiGateway6RealserversWeight(i["weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {

			tmp["http_host"] = flattenFirewallAccessProxyApiGateway6RealserversHttpHost(i["http-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenFirewallAccessProxyApiGateway6RealserversHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := i["health-check-proto"]; ok {

			tmp["health_check_proto"] = flattenFirewallAccessProxyApiGateway6RealserversHealthCheckProto(i["health-check-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {

			tmp["holddown_interval"] = flattenFirewallAccessProxyApiGateway6RealserversHolddownInterval(i["holddown-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := i["ssh-client-cert"]; ok {

			tmp["ssh_client_cert"] = flattenFirewallAccessProxyApiGateway6RealserversSshClientCert(i["ssh-client-cert"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := i["ssh-host-key-validation"]; ok {

			tmp["ssh_host_key_validation"] = flattenFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(i["ssh-host-key-validation"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := i["ssh-host-key"]; ok {

			tmp["ssh_host_key"] = flattenFirewallAccessProxyApiGateway6RealserversSshHostKey(i["ssh-host-key"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxyApiGateway6RealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversMappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6RealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallAccessProxyApiGateway6RealserversSshHostKeyName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAccessProxyApiGateway6RealserversSshHostKeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6Persistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SamlServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SamlRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenFirewallAccessProxyApiGateway6SslCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallAccessProxyApiGateway6SslCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallAccessProxyApiGateway6SslCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxyApiGateway6SslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallAccessProxyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vip", flattenFirewallAccessProxyVip(o["vip"], d, "vip", sv)); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenFirewallAccessProxyClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal", flattenFirewallAccessProxyAuthPortal(o["auth-portal"], d, "auth_portal", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal"]) {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenFirewallAccessProxyAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host", sv)); err != nil {
		if !fortiAPIPatch(o["auth-virtual-host"]) {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallAccessProxyEmptyCertAction(o["empty-cert-action"], d, "empty_cert_action", sv)); err != nil {
		if !fortiAPIPatch(o["empty-cert-action"]) {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenFirewallAccessProxyLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["log-blocked-traffic"]) {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallAccessProxyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("api_gateway", flattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway"]) {
				return fmt.Errorf("Error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenFirewallAccessProxyApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway"]) {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("api_gateway6", flattenFirewallAccessProxyApiGateway6(o["api-gateway6"], d, "api_gateway6", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway6"]) {
				return fmt.Errorf("Error reading api_gateway6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway6"); ok {
			if err = d.Set("api_gateway6", flattenFirewallAccessProxyApiGateway6(o["api-gateway6"], d, "api_gateway6", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway6"]) {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallAccessProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallAccessProxyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyAuthPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyEmptyCertAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallAccessProxyApiGatewayId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url-map"], _ = expandFirewallAccessProxyApiGatewayUrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["service"], _ = expandFirewallAccessProxyApiGatewayService(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ldb-method"], _ = expandFirewallAccessProxyApiGatewayLdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["virtual-host"], _ = expandFirewallAccessProxyApiGatewayVirtualHost(d, i["virtual_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url-map-type"], _ = expandFirewallAccessProxyApiGatewayUrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["realservers"], _ = expandFirewallAccessProxyApiGatewayRealservers(d, i["realservers"], pre_append, sv)
		} else {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["persistence"], _ = expandFirewallAccessProxyApiGatewayPersistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-domain"], _ = expandFirewallAccessProxyApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-path"], _ = expandFirewallAccessProxyApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-generation"], _ = expandFirewallAccessProxyApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-age"], _ = expandFirewallAccessProxyApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-share"], _ = expandFirewallAccessProxyApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["https-cookie-secure"], _ = expandFirewallAccessProxyApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["saml-server"], _ = expandFirewallAccessProxyApiGatewaySamlServer(d, i["saml_server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["saml-redirect"], _ = expandFirewallAccessProxyApiGatewaySamlRedirect(d, i["saml_redirect"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxyApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-algorithm"], _ = expandFirewallAccessProxyApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["ssl-cipher-suites"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-min-version"], _ = expandFirewallAccessProxyApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-max-version"], _ = expandFirewallAccessProxyApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxyApiGatewaySslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallAccessProxyApiGatewayRealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-type"], _ = expandFirewallAccessProxyApiGatewayRealserversAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["address"], _ = expandFirewallAccessProxyApiGatewayRealserversAddress(d, i["address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandFirewallAccessProxyApiGatewayRealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["domain"], _ = expandFirewallAccessProxyApiGatewayRealserversDomain(d, i["domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallAccessProxyApiGatewayRealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mappedport"], _ = expandFirewallAccessProxyApiGatewayRealserversMappedport(d, i["mappedport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallAccessProxyApiGatewayRealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandFirewallAccessProxyApiGatewayRealserversType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandFirewallAccessProxyApiGatewayRealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-host"], _ = expandFirewallAccessProxyApiGatewayRealserversHttpHost(d, i["http_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check-proto"], _ = expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["holddown-interval"], _ = expandFirewallAccessProxyApiGatewayRealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssh-client-cert"], _ = expandFirewallAccessProxyApiGatewayRealserversSshClientCert(d, i["ssh_client_cert"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["ssh-host-key"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKey(d, i["ssh_host_key"], pre_append, sv)
		} else {
			tmp["ssh-host-key"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversMappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallAccessProxyApiGatewayRealserversSshHostKeyName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewayRealserversSshHostKeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySamlRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["priority"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGatewaySslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallAccessProxyApiGateway6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url-map"], _ = expandFirewallAccessProxyApiGateway6UrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["service"], _ = expandFirewallAccessProxyApiGateway6Service(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ldb-method"], _ = expandFirewallAccessProxyApiGateway6LdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["virtual-host"], _ = expandFirewallAccessProxyApiGateway6VirtualHost(d, i["virtual_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["url-map-type"], _ = expandFirewallAccessProxyApiGateway6UrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["realservers"], _ = expandFirewallAccessProxyApiGateway6Realservers(d, i["realservers"], pre_append, sv)
		} else {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["persistence"], _ = expandFirewallAccessProxyApiGateway6Persistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-domain"], _ = expandFirewallAccessProxyApiGateway6HttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-path"], _ = expandFirewallAccessProxyApiGateway6HttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-generation"], _ = expandFirewallAccessProxyApiGateway6HttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-age"], _ = expandFirewallAccessProxyApiGateway6HttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-cookie-share"], _ = expandFirewallAccessProxyApiGateway6HttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["https-cookie-secure"], _ = expandFirewallAccessProxyApiGateway6HttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["saml-server"], _ = expandFirewallAccessProxyApiGateway6SamlServer(d, i["saml_server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["saml-redirect"], _ = expandFirewallAccessProxyApiGateway6SamlRedirect(d, i["saml_redirect"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxyApiGateway6SslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-algorithm"], _ = expandFirewallAccessProxyApiGateway6SslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["ssl-cipher-suites"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-min-version"], _ = expandFirewallAccessProxyApiGateway6SslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-max-version"], _ = expandFirewallAccessProxyApiGateway6SslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxyApiGateway6SslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6UrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Service(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6LdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6VirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6UrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallAccessProxyApiGateway6RealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-type"], _ = expandFirewallAccessProxyApiGateway6RealserversAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["address"], _ = expandFirewallAccessProxyApiGateway6RealserversAddress(d, i["address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandFirewallAccessProxyApiGateway6RealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["domain"], _ = expandFirewallAccessProxyApiGateway6RealserversDomain(d, i["domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallAccessProxyApiGateway6RealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mappedport"], _ = expandFirewallAccessProxyApiGateway6RealserversMappedport(d, i["mappedport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallAccessProxyApiGateway6RealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandFirewallAccessProxyApiGateway6RealserversType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandFirewallAccessProxyApiGateway6RealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-host"], _ = expandFirewallAccessProxyApiGateway6RealserversHttpHost(d, i["http_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandFirewallAccessProxyApiGateway6RealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check-proto"], _ = expandFirewallAccessProxyApiGateway6RealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["holddown-interval"], _ = expandFirewallAccessProxyApiGateway6RealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssh-client-cert"], _ = expandFirewallAccessProxyApiGateway6RealserversSshClientCert(d, i["ssh_client_cert"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["ssh-host-key"], _ = expandFirewallAccessProxyApiGateway6RealserversSshHostKey(d, i["ssh_host_key"], pre_append, sv)
		} else {
			tmp["ssh-host-key"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6RealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversMappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallAccessProxyApiGateway6RealserversSshHostKeyName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6RealserversSshHostKeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6Persistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SamlServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SamlRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["priority"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallAccessProxyApiGateway6SslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxyApiGateway6SslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallAccessProxyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok {

		t, err := expandFirewallAccessProxyVip(d, v, "vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	}

	if v, ok := d.GetOk("client_cert"); ok {

		t, err := expandFirewallAccessProxyClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok {

		t, err := expandFirewallAccessProxyAuthPortal(d, v, "auth_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok {

		t, err := expandFirewallAccessProxyAuthVirtualHost(d, v, "auth_virtual_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("empty_cert_action"); ok {

		t, err := expandFirewallAccessProxyEmptyCertAction(d, v, "empty_cert_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok {

		t, err := expandFirewallAccessProxyLogBlockedTraffic(d, v, "log_blocked_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {

		t, err := expandFirewallAccessProxyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway"); ok || d.HasChange("api_gateway") {

		t, err := expandFirewallAccessProxyApiGateway(d, v, "api_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway6"); ok || d.HasChange("api_gateway6") {

		t, err := expandFirewallAccessProxyApiGateway6(d, v, "api_gateway6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway6"] = t
		}
	}

	return &obj, nil
}
