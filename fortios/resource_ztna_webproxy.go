// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ZTNA web-proxy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebProxyCreate,
		Read:   resourceZtnaWebProxyRead,
		Update: resourceZtnaWebProxyUpdate,
		Delete: resourceZtnaWebProxyDelete,

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
			"log_blocked_traffic": &schema.Schema{
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
						"url_map_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"h2_support": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
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
									"translate_host": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"verify_cert": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
						},
						"http_cookie_path": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"http_cookie_generation": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"ssl_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"url_map_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"h2_support": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
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
									"translate_host": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"verify_cert": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
						},
						"http_cookie_path": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"http_cookie_generation": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
						"ssl_renegotiation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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

func resourceZtnaWebProxyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaWebProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxy resource while getting object: %v", err)
	}

	o, err := c.CreateZtnaWebProxy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebProxy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaWebProxy")
	}

	return resourceZtnaWebProxyRead(d, m)
}

func resourceZtnaWebProxyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaWebProxy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxy resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaWebProxy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaWebProxy")
	}

	return resourceZtnaWebProxyRead(d, m)
}

func resourceZtnaWebProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteZtnaWebProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebProxyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebProxy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebProxy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebProxy resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebProxyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyAuthPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyVip6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxySvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxySvrPoolTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxySvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxySvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenZtnaWebProxyApiGatewayId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["url-map"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
			}
			tmp["url_map"] = flattenZtnaWebProxyApiGatewayUrlMap(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["service"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
			}
			tmp["service"] = flattenZtnaWebProxyApiGatewayService(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ldb-method"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
			}
			tmp["ldb_method"] = flattenZtnaWebProxyApiGatewayLdbMethod(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["url-map-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
			}
			tmp["url_map_type"] = flattenZtnaWebProxyApiGatewayUrlMapType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["h2-support"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
			}
			tmp["h2_support"] = flattenZtnaWebProxyApiGatewayH2Support(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["h3-support"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
			}
			tmp["h3_support"] = flattenZtnaWebProxyApiGatewayH3Support(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["quic"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
			}
			tmp["quic"] = flattenZtnaWebProxyApiGatewayQuic(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["realservers"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
			}
			tmp["realservers"] = flattenZtnaWebProxyApiGatewayRealservers(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["persistence"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
			}
			tmp["persistence"] = flattenZtnaWebProxyApiGatewayPersistence(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-domain-from-host"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
			}
			tmp["http_cookie_domain_from_host"] = flattenZtnaWebProxyApiGatewayHttpCookieDomainFromHost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-domain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
			}
			tmp["http_cookie_domain"] = flattenZtnaWebProxyApiGatewayHttpCookieDomain(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-path"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
			}
			tmp["http_cookie_path"] = flattenZtnaWebProxyApiGatewayHttpCookiePath(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-generation"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
			}
			tmp["http_cookie_generation"] = flattenZtnaWebProxyApiGatewayHttpCookieGeneration(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-age"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
			}
			tmp["http_cookie_age"] = flattenZtnaWebProxyApiGatewayHttpCookieAge(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-share"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
			}
			tmp["http_cookie_share"] = flattenZtnaWebProxyApiGatewayHttpCookieShare(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["https-cookie-secure"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
			}
			tmp["https_cookie_secure"] = flattenZtnaWebProxyApiGatewayHttpsCookieSecure(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-dh-bits"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
			}
			tmp["ssl_dh_bits"] = flattenZtnaWebProxyApiGatewaySslDhBits(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-algorithm"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
			}
			tmp["ssl_algorithm"] = flattenZtnaWebProxyApiGatewaySslAlgorithm(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-cipher-suites"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
			}
			tmp["ssl_cipher_suites"] = flattenZtnaWebProxyApiGatewaySslCipherSuites(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-min-version"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
			}
			tmp["ssl_min_version"] = flattenZtnaWebProxyApiGatewaySslMinVersion(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-max-version"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
			}
			tmp["ssl_max_version"] = flattenZtnaWebProxyApiGatewaySslMaxVersion(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-renegotiation"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
			}
			tmp["ssl_renegotiation"] = flattenZtnaWebProxyApiGatewaySslRenegotiation(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenZtnaWebProxyApiGatewayId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayH2Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayH3Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaWebProxyApiGatewayQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaWebProxyApiGatewayQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaWebProxyApiGatewayQuicMaxAckDelay(i["max-ack-delay"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaWebProxyApiGatewayQuicActiveMigration(i["active-migration"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaWebProxyApiGatewayQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaWebProxyApiGatewayQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenZtnaWebProxyApiGatewayRealserversId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
			}
			tmp["addr_type"] = flattenZtnaWebProxyApiGatewayRealserversAddrType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
			}
			tmp["address"] = flattenZtnaWebProxyApiGatewayRealserversAddress(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
			}
			tmp["ip"] = flattenZtnaWebProxyApiGatewayRealserversIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenZtnaWebProxyApiGatewayRealserversPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenZtnaWebProxyApiGatewayRealserversStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenZtnaWebProxyApiGatewayRealserversWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-host"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
			}
			tmp["http_host"] = flattenZtnaWebProxyApiGatewayRealserversHttpHost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
			}
			tmp["health_check"] = flattenZtnaWebProxyApiGatewayRealserversHealthCheck(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check-proto"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
			}
			tmp["health_check_proto"] = flattenZtnaWebProxyApiGatewayRealserversHealthCheckProto(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["holddown-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
			}
			tmp["holddown_interval"] = flattenZtnaWebProxyApiGatewayRealserversHolddownInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["translate-host"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
			}
			tmp["translate_host"] = flattenZtnaWebProxyApiGatewayRealserversTranslateHost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["verify-cert"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
			}
			tmp["verify_cert"] = flattenZtnaWebProxyApiGatewayRealserversVerifyCert(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenZtnaWebProxyApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayRealserversAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayRealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "priority", "priority")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
			}
			tmp["priority"] = flattenZtnaWebProxyApiGatewaySslCipherSuitesPriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cipher"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
			}
			tmp["cipher"] = flattenZtnaWebProxyApiGatewaySslCipherSuitesCipher(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["versions"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
			}
			tmp["versions"] = flattenZtnaWebProxyApiGatewaySslCipherSuitesVersions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGatewaySslRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenZtnaWebProxyApiGateway6Id(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["url-map"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
			}
			tmp["url_map"] = flattenZtnaWebProxyApiGateway6UrlMap(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["service"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
			}
			tmp["service"] = flattenZtnaWebProxyApiGateway6Service(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ldb-method"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
			}
			tmp["ldb_method"] = flattenZtnaWebProxyApiGateway6LdbMethod(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["url-map-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
			}
			tmp["url_map_type"] = flattenZtnaWebProxyApiGateway6UrlMapType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["h2-support"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
			}
			tmp["h2_support"] = flattenZtnaWebProxyApiGateway6H2Support(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["h3-support"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
			}
			tmp["h3_support"] = flattenZtnaWebProxyApiGateway6H3Support(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["quic"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
			}
			tmp["quic"] = flattenZtnaWebProxyApiGateway6Quic(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["realservers"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
			}
			tmp["realservers"] = flattenZtnaWebProxyApiGateway6Realservers(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["persistence"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
			}
			tmp["persistence"] = flattenZtnaWebProxyApiGateway6Persistence(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-domain-from-host"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
			}
			tmp["http_cookie_domain_from_host"] = flattenZtnaWebProxyApiGateway6HttpCookieDomainFromHost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-domain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
			}
			tmp["http_cookie_domain"] = flattenZtnaWebProxyApiGateway6HttpCookieDomain(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-path"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
			}
			tmp["http_cookie_path"] = flattenZtnaWebProxyApiGateway6HttpCookiePath(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-generation"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
			}
			tmp["http_cookie_generation"] = flattenZtnaWebProxyApiGateway6HttpCookieGeneration(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-age"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
			}
			tmp["http_cookie_age"] = flattenZtnaWebProxyApiGateway6HttpCookieAge(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-cookie-share"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
			}
			tmp["http_cookie_share"] = flattenZtnaWebProxyApiGateway6HttpCookieShare(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["https-cookie-secure"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
			}
			tmp["https_cookie_secure"] = flattenZtnaWebProxyApiGateway6HttpsCookieSecure(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-dh-bits"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
			}
			tmp["ssl_dh_bits"] = flattenZtnaWebProxyApiGateway6SslDhBits(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-algorithm"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
			}
			tmp["ssl_algorithm"] = flattenZtnaWebProxyApiGateway6SslAlgorithm(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-cipher-suites"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
			}
			tmp["ssl_cipher_suites"] = flattenZtnaWebProxyApiGateway6SslCipherSuites(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-min-version"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
			}
			tmp["ssl_min_version"] = flattenZtnaWebProxyApiGateway6SslMinVersion(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-max-version"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
			}
			tmp["ssl_max_version"] = flattenZtnaWebProxyApiGateway6SslMaxVersion(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ssl-renegotiation"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
			}
			tmp["ssl_renegotiation"] = flattenZtnaWebProxyApiGateway6SslRenegotiation(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenZtnaWebProxyApiGateway6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6UrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Service(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6LdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6UrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6H2Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6H3Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Quic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenZtnaWebProxyApiGateway6QuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenZtnaWebProxyApiGateway6QuicMaxAckDelay(i["max-ack-delay"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenZtnaWebProxyApiGateway6QuicActiveMigration(i["active-migration"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenZtnaWebProxyApiGateway6QuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6QuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6QuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6QuicActiveMigration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6QuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Realservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenZtnaWebProxyApiGateway6RealserversId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
			}
			tmp["addr_type"] = flattenZtnaWebProxyApiGateway6RealserversAddrType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["address"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
			}
			tmp["address"] = flattenZtnaWebProxyApiGateway6RealserversAddress(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
			}
			tmp["ip"] = flattenZtnaWebProxyApiGateway6RealserversIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenZtnaWebProxyApiGateway6RealserversPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenZtnaWebProxyApiGateway6RealserversStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenZtnaWebProxyApiGateway6RealserversWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-host"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
			}
			tmp["http_host"] = flattenZtnaWebProxyApiGateway6RealserversHttpHost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
			}
			tmp["health_check"] = flattenZtnaWebProxyApiGateway6RealserversHealthCheck(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check-proto"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
			}
			tmp["health_check_proto"] = flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["holddown-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
			}
			tmp["holddown_interval"] = flattenZtnaWebProxyApiGateway6RealserversHolddownInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["translate-host"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
			}
			tmp["translate_host"] = flattenZtnaWebProxyApiGateway6RealserversTranslateHost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["verify-cert"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
			}
			tmp["verify_cert"] = flattenZtnaWebProxyApiGateway6RealserversVerifyCert(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenZtnaWebProxyApiGateway6RealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6RealserversAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6RealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6RealserversVerifyCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6Persistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "priority", "priority")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
			}
			tmp["priority"] = flattenZtnaWebProxyApiGateway6SslCipherSuitesPriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cipher"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
			}
			tmp["cipher"] = flattenZtnaWebProxyApiGateway6SslCipherSuitesCipher(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["versions"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
			}
			tmp["versions"] = flattenZtnaWebProxyApiGateway6SslCipherSuitesVersions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebProxyApiGateway6SslRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectZtnaWebProxy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenZtnaWebProxyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vip", flattenZtnaWebProxyVip(o["vip"], d, "vip", sv)); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("host", flattenZtnaWebProxyHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenZtnaWebProxyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenZtnaWebProxyLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["log-blocked-traffic"]) {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("auth_portal", flattenZtnaWebProxyAuthPortal(o["auth-portal"], d, "auth_portal", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal"]) {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenZtnaWebProxyAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host", sv)); err != nil {
		if !fortiAPIPatch(o["auth-virtual-host"]) {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("vip6", flattenZtnaWebProxyVip6(o["vip6"], d, "vip6", sv)); err != nil {
		if !fortiAPIPatch(o["vip6"]) {
			return fmt.Errorf("Error reading vip6: %v", err)
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenZtnaWebProxySvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-multiplex"]) {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenZtnaWebProxySvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-ttl"]) {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenZtnaWebProxySvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-server-max-request"]) {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenZtnaWebProxySvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-server-max-concurrent-request"]) {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("api_gateway", flattenZtnaWebProxyApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway"]) {
				return fmt.Errorf("Error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenZtnaWebProxyApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway"]) {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("api_gateway6", flattenZtnaWebProxyApiGateway6(o["api-gateway6"], d, "api_gateway6", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway6"]) {
				return fmt.Errorf("Error reading api_gateway6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway6"); ok {
			if err = d.Set("api_gateway6", flattenZtnaWebProxyApiGateway6(o["api-gateway6"], d, "api_gateway6", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway6"]) {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenZtnaWebProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaWebProxyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyAuthPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyVip6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxySvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandZtnaWebProxyApiGatewayId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map"], _ = expandZtnaWebProxyApiGatewayUrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["service"], _ = expandZtnaWebProxyApiGatewayService(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ldb-method"], _ = expandZtnaWebProxyApiGatewayLdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map-type"], _ = expandZtnaWebProxyApiGatewayUrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h2-support"], _ = expandZtnaWebProxyApiGatewayH2Support(d, i["h2_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h3-support"], _ = expandZtnaWebProxyApiGatewayH3Support(d, i["h3_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quic"], _ = expandZtnaWebProxyApiGatewayQuic(d, i["quic"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["quic"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["realservers"], _ = expandZtnaWebProxyApiGatewayRealservers(d, i["realservers"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["persistence"], _ = expandZtnaWebProxyApiGatewayPersistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain-from-host"], _ = expandZtnaWebProxyApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain"], _ = expandZtnaWebProxyApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-path"], _ = expandZtnaWebProxyApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-path"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-generation"], _ = expandZtnaWebProxyApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-generation"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-age"], _ = expandZtnaWebProxyApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-share"], _ = expandZtnaWebProxyApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-cookie-secure"], _ = expandZtnaWebProxyApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-dh-bits"], _ = expandZtnaWebProxyApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-algorithm"], _ = expandZtnaWebProxyApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-cipher-suites"], _ = expandZtnaWebProxyApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-min-version"], _ = expandZtnaWebProxyApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-max-version"], _ = expandZtnaWebProxyApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-renegotiation"], _ = expandZtnaWebProxyApiGatewaySslRenegotiation(d, i["ssl_renegotiation"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewayId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayH2Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayH3Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-idle-timeout"], _ = expandZtnaWebProxyApiGatewayQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-udp-payload-size"], _ = expandZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-connection-id-limit"], _ = expandZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-delay-exponent"], _ = expandZtnaWebProxyApiGatewayQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-ack-delay"], _ = expandZtnaWebProxyApiGatewayQuicMaxAckDelay(d, i["max_ack_delay"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-datagram-frame-size"], _ = expandZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-migration"], _ = expandZtnaWebProxyApiGatewayQuicActiveMigration(d, i["active_migration"], pre_append, sv)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok {
		result["grease-quic-bit"], _ = expandZtnaWebProxyApiGatewayQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append, sv)
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandZtnaWebProxyApiGatewayRealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandZtnaWebProxyApiGatewayRealserversAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandZtnaWebProxyApiGatewayRealserversAddress(d, i["address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandZtnaWebProxyApiGatewayRealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandZtnaWebProxyApiGatewayRealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandZtnaWebProxyApiGatewayRealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandZtnaWebProxyApiGatewayRealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-host"], _ = expandZtnaWebProxyApiGatewayRealserversHttpHost(d, i["http_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandZtnaWebProxyApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check-proto"], _ = expandZtnaWebProxyApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holddown-interval"], _ = expandZtnaWebProxyApiGatewayRealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["translate-host"], _ = expandZtnaWebProxyApiGatewayRealserversTranslateHost(d, i["translate_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["verify-cert"], _ = expandZtnaWebProxyApiGatewayRealserversVerifyCert(d, i["verify_cert"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayRealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["priority"], _ = expandZtnaWebProxyApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cipher"], _ = expandZtnaWebProxyApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cipher"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["versions"], _ = expandZtnaWebProxyApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGatewaySslRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandZtnaWebProxyApiGateway6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map"], _ = expandZtnaWebProxyApiGateway6UrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["service"], _ = expandZtnaWebProxyApiGateway6Service(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ldb-method"], _ = expandZtnaWebProxyApiGateway6LdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map-type"], _ = expandZtnaWebProxyApiGateway6UrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h2-support"], _ = expandZtnaWebProxyApiGateway6H2Support(d, i["h2_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h3-support"], _ = expandZtnaWebProxyApiGateway6H3Support(d, i["h3_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quic"], _ = expandZtnaWebProxyApiGateway6Quic(d, i["quic"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["quic"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["realservers"], _ = expandZtnaWebProxyApiGateway6Realservers(d, i["realservers"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["persistence"], _ = expandZtnaWebProxyApiGateway6Persistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain-from-host"], _ = expandZtnaWebProxyApiGateway6HttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain"], _ = expandZtnaWebProxyApiGateway6HttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-path"], _ = expandZtnaWebProxyApiGateway6HttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-path"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-generation"], _ = expandZtnaWebProxyApiGateway6HttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-generation"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-age"], _ = expandZtnaWebProxyApiGateway6HttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-share"], _ = expandZtnaWebProxyApiGateway6HttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-cookie-secure"], _ = expandZtnaWebProxyApiGateway6HttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-dh-bits"], _ = expandZtnaWebProxyApiGateway6SslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-algorithm"], _ = expandZtnaWebProxyApiGateway6SslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-cipher-suites"], _ = expandZtnaWebProxyApiGateway6SslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-min-version"], _ = expandZtnaWebProxyApiGateway6SslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-max-version"], _ = expandZtnaWebProxyApiGateway6SslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-renegotiation"], _ = expandZtnaWebProxyApiGateway6SslRenegotiation(d, i["ssl_renegotiation"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6UrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Service(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6LdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6UrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6H2Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6H3Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Quic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-idle-timeout"], _ = expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-udp-payload-size"], _ = expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-connection-id-limit"], _ = expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-delay-exponent"], _ = expandZtnaWebProxyApiGateway6QuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-ack-delay"], _ = expandZtnaWebProxyApiGateway6QuicMaxAckDelay(d, i["max_ack_delay"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-datagram-frame-size"], _ = expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-migration"], _ = expandZtnaWebProxyApiGateway6QuicActiveMigration(d, i["active_migration"], pre_append, sv)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok {
		result["grease-quic-bit"], _ = expandZtnaWebProxyApiGateway6QuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append, sv)
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicActiveMigration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6QuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandZtnaWebProxyApiGateway6RealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandZtnaWebProxyApiGateway6RealserversAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandZtnaWebProxyApiGateway6RealserversAddress(d, i["address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandZtnaWebProxyApiGateway6RealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandZtnaWebProxyApiGateway6RealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandZtnaWebProxyApiGateway6RealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandZtnaWebProxyApiGateway6RealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-host"], _ = expandZtnaWebProxyApiGateway6RealserversHttpHost(d, i["http_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandZtnaWebProxyApiGateway6RealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check-proto"], _ = expandZtnaWebProxyApiGateway6RealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holddown-interval"], _ = expandZtnaWebProxyApiGateway6RealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["translate-host"], _ = expandZtnaWebProxyApiGateway6RealserversTranslateHost(d, i["translate_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "verify_cert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["verify-cert"], _ = expandZtnaWebProxyApiGateway6RealserversVerifyCert(d, i["verify_cert"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6RealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6RealserversVerifyCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6Persistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["priority"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cipher"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cipher"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["versions"], _ = expandZtnaWebProxyApiGateway6SslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebProxyApiGateway6SslRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebProxy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandZtnaWebProxyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok {
		t, err := expandZtnaWebProxyVip(d, v, "vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	} else if d.HasChange("vip") {
		obj["vip"] = nil
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandZtnaWebProxyHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	} else if d.HasChange("host") {
		obj["host"] = nil
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		t, err := expandZtnaWebProxyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	} else if d.HasChange("decrypted_traffic_mirror") {
		obj["decrypted-traffic-mirror"] = nil
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok {
		t, err := expandZtnaWebProxyLogBlockedTraffic(d, v, "log_blocked_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok {
		t, err := expandZtnaWebProxyAuthPortal(d, v, "auth_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok {
		t, err := expandZtnaWebProxyAuthVirtualHost(d, v, "auth_virtual_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	} else if d.HasChange("auth_virtual_host") {
		obj["auth-virtual-host"] = nil
	}

	if v, ok := d.GetOk("vip6"); ok {
		t, err := expandZtnaWebProxyVip6(d, v, "vip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip6"] = t
		}
	} else if d.HasChange("vip6") {
		obj["vip6"] = nil
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok {
		t, err := expandZtnaWebProxySvrPoolMultiplex(d, v, "svr_pool_multiplex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOkExists("svr_pool_ttl"); ok {
		t, err := expandZtnaWebProxySvrPoolTtl(d, v, "svr_pool_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOkExists("svr_pool_server_max_request"); ok {
		t, err := expandZtnaWebProxySvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	} else if d.HasChange("svr_pool_server_max_request") {
		obj["svr-pool-server-max-request"] = nil
	}

	if v, ok := d.GetOkExists("svr_pool_server_max_concurrent_request"); ok {
		t, err := expandZtnaWebProxySvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	} else if d.HasChange("svr_pool_server_max_concurrent_request") {
		obj["svr-pool-server-max-concurrent-request"] = nil
	}

	if v, ok := d.GetOk("api_gateway"); ok || d.HasChange("api_gateway") {
		t, err := expandZtnaWebProxyApiGateway(d, v, "api_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway6"); ok || d.HasChange("api_gateway6") {
		t, err := expandZtnaWebProxyApiGateway6(d, v, "api_gateway6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway6"] = t
		}
	}

	return &obj, nil
}
