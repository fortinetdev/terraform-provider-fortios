// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 access proxy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallAccessProxy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAccessProxy6Create,
		Read:   resourceFirewallAccessProxy6Read,
		Update: resourceFirewallAccessProxy6Update,
		Delete: resourceFirewallAccessProxy6Delete,

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
				Optional:     true,
			},
			"vip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
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
			"add_vhost_domain_to_dnsdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_supported_max_version": &schema.Schema{
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
			"decrypted_traffic_mirror": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"api_gateway": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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
									"domain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
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
									"external_auth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tunnel_encryption": &schema.Schema{
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
									"ssh_client_cert": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"ssh_host_key_validation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_host_key": &schema.Schema{
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
								},
							},
						},
						"application": &schema.Schema{
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
						"saml_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
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
						"ssl_vpn_web_portal": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
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
									"domain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
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
									"external_auth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tunnel_encryption": &schema.Schema{
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
									"ssh_client_cert": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"ssh_host_key_validation": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ssh_host_key": &schema.Schema{
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
								},
							},
						},
						"application": &schema.Schema{
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
						"saml_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
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
						"ssl_vpn_web_portal": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
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

func resourceFirewallAccessProxy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxy6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy6 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallAccessProxy6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallAccessProxy6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxy6")
	}

	return resourceFirewallAccessProxy6Read(d, m)
}

func resourceFirewallAccessProxy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallAccessProxy6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAccessProxy6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAccessProxy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAccessProxy6")
	}

	return resourceFirewallAccessProxy6Read(d, m)
}

func resourceFirewallAccessProxy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallAccessProxy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAccessProxy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAccessProxy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallAccessProxy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAccessProxy6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAccessProxy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAccessProxy6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6Vip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6UserAgentDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6AuthPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6AuthVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6EmptyCertAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6LogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6AddVhostDomainToDnsdb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6HttpSupportedMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6SvrPoolMultiplex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6SvrPoolTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6SvrPoolServerMaxRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6DecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallAccessProxy6ApiGatewayId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if cur_v, ok := i["url-map"]; ok {
			tmp["url_map"] = flattenFirewallAccessProxy6ApiGatewayUrlMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if cur_v, ok := i["service"]; ok {
			tmp["service"] = flattenFirewallAccessProxy6ApiGatewayService(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if cur_v, ok := i["ldb-method"]; ok {
			tmp["ldb_method"] = flattenFirewallAccessProxy6ApiGatewayLdbMethod(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if cur_v, ok := i["virtual-host"]; ok {
			tmp["virtual_host"] = flattenFirewallAccessProxy6ApiGatewayVirtualHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if cur_v, ok := i["url-map-type"]; ok {
			tmp["url_map_type"] = flattenFirewallAccessProxy6ApiGatewayUrlMapType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if cur_v, ok := i["h2-support"]; ok {
			tmp["h2_support"] = flattenFirewallAccessProxy6ApiGatewayH2Support(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if cur_v, ok := i["h3-support"]; ok {
			tmp["h3_support"] = flattenFirewallAccessProxy6ApiGatewayH3Support(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if cur_v, ok := i["quic"]; ok {
			tmp["quic"] = flattenFirewallAccessProxy6ApiGatewayQuic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if cur_v, ok := i["realservers"]; ok {
			tmp["realservers"] = flattenFirewallAccessProxy6ApiGatewayRealservers(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if cur_v, ok := i["application"]; ok {
			tmp["application"] = flattenFirewallAccessProxy6ApiGatewayApplication(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if cur_v, ok := i["persistence"]; ok {
			tmp["persistence"] = flattenFirewallAccessProxy6ApiGatewayPersistence(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if cur_v, ok := i["http-cookie-domain-from-host"]; ok {
			tmp["http_cookie_domain_from_host"] = flattenFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if cur_v, ok := i["http-cookie-domain"]; ok {
			tmp["http_cookie_domain"] = flattenFirewallAccessProxy6ApiGatewayHttpCookieDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if cur_v, ok := i["http-cookie-path"]; ok {
			tmp["http_cookie_path"] = flattenFirewallAccessProxy6ApiGatewayHttpCookiePath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if cur_v, ok := i["http-cookie-generation"]; ok {
			tmp["http_cookie_generation"] = flattenFirewallAccessProxy6ApiGatewayHttpCookieGeneration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if cur_v, ok := i["http-cookie-age"]; ok {
			tmp["http_cookie_age"] = flattenFirewallAccessProxy6ApiGatewayHttpCookieAge(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if cur_v, ok := i["http-cookie-share"]; ok {
			tmp["http_cookie_share"] = flattenFirewallAccessProxy6ApiGatewayHttpCookieShare(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if cur_v, ok := i["https-cookie-secure"]; ok {
			tmp["https_cookie_secure"] = flattenFirewallAccessProxy6ApiGatewayHttpsCookieSecure(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if cur_v, ok := i["saml-server"]; ok {
			tmp["saml_server"] = flattenFirewallAccessProxy6ApiGatewaySamlServer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if cur_v, ok := i["saml-redirect"]; ok {
			tmp["saml_redirect"] = flattenFirewallAccessProxy6ApiGatewaySamlRedirect(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if cur_v, ok := i["ssl-dh-bits"]; ok {
			tmp["ssl_dh_bits"] = flattenFirewallAccessProxy6ApiGatewaySslDhBits(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if cur_v, ok := i["ssl-algorithm"]; ok {
			tmp["ssl_algorithm"] = flattenFirewallAccessProxy6ApiGatewaySslAlgorithm(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if cur_v, ok := i["ssl-cipher-suites"]; ok {
			tmp["ssl_cipher_suites"] = flattenFirewallAccessProxy6ApiGatewaySslCipherSuites(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if cur_v, ok := i["ssl-min-version"]; ok {
			tmp["ssl_min_version"] = flattenFirewallAccessProxy6ApiGatewaySslMinVersion(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if cur_v, ok := i["ssl-max-version"]; ok {
			tmp["ssl_max_version"] = flattenFirewallAccessProxy6ApiGatewaySslMaxVersion(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if cur_v, ok := i["ssl-renegotiation"]; ok {
			tmp["ssl_renegotiation"] = flattenFirewallAccessProxy6ApiGatewaySslRenegotiation(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if cur_v, ok := i["ssl-vpn-web-portal"]; ok {
			tmp["ssl_vpn_web_portal"] = flattenFirewallAccessProxy6ApiGatewaySslVpnWebPortal(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxy6ApiGatewayId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayUrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayUrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayH2Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayH3Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(i["max-ack-delay"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxy6ApiGatewayQuicActiveMigration(i["active-migration"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayQuicActiveMigration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallAccessProxy6ApiGatewayRealserversId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if cur_v, ok := i["addr-type"]; ok {
			tmp["addr_type"] = flattenFirewallAccessProxy6ApiGatewayRealserversAddrType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if cur_v, ok := i["address"]; ok {
			tmp["address"] = flattenFirewallAccessProxy6ApiGatewayRealserversAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenFirewallAccessProxy6ApiGatewayRealserversIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if cur_v, ok := i["domain"]; ok {
			tmp["domain"] = flattenFirewallAccessProxy6ApiGatewayRealserversDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenFirewallAccessProxy6ApiGatewayRealserversPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if cur_v, ok := i["mappedport"]; ok {
			tmp["mappedport"] = flattenFirewallAccessProxy6ApiGatewayRealserversMappedport(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenFirewallAccessProxy6ApiGatewayRealserversStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenFirewallAccessProxy6ApiGatewayRealserversType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if cur_v, ok := i["external-auth"]; ok {
			tmp["external_auth"] = flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if cur_v, ok := i["tunnel-encryption"]; ok {
			tmp["tunnel_encryption"] = flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if cur_v, ok := i["weight"]; ok {
			tmp["weight"] = flattenFirewallAccessProxy6ApiGatewayRealserversWeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if cur_v, ok := i["http-host"]; ok {
			tmp["http_host"] = flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if cur_v, ok := i["health-check"]; ok {
			tmp["health_check"] = flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if cur_v, ok := i["health-check-proto"]; ok {
			tmp["health_check_proto"] = flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if cur_v, ok := i["holddown-interval"]; ok {
			tmp["holddown_interval"] = flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if cur_v, ok := i["translate-host"]; ok {
			tmp["translate_host"] = flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if cur_v, ok := i["ssh-client-cert"]; ok {
			tmp["ssh_client_cert"] = flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if cur_v, ok := i["ssh-host-key-validation"]; ok {
			tmp["ssh_host_key_validation"] = flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if cur_v, ok := i["ssh-host-key"]; ok {
			tmp["ssh_host_key"] = flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxy6ApiGatewayRealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversMappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversExternalAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAccessProxy6ApiGatewayRealserversSshHostKeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallAccessProxy6ApiGatewayApplicationName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAccessProxy6ApiGatewayApplicationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewayHttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewayHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySamlServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySamlRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["priority"] = flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if cur_v, ok := i["cipher"]; ok {
			tmp["cipher"] = flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if cur_v, ok := i["versions"]; ok {
			tmp["versions"] = flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGatewaySslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallAccessProxy6ApiGateway6Id(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if cur_v, ok := i["url-map"]; ok {
			tmp["url_map"] = flattenFirewallAccessProxy6ApiGateway6UrlMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if cur_v, ok := i["service"]; ok {
			tmp["service"] = flattenFirewallAccessProxy6ApiGateway6Service(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if cur_v, ok := i["ldb-method"]; ok {
			tmp["ldb_method"] = flattenFirewallAccessProxy6ApiGateway6LdbMethod(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if cur_v, ok := i["virtual-host"]; ok {
			tmp["virtual_host"] = flattenFirewallAccessProxy6ApiGateway6VirtualHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if cur_v, ok := i["url-map-type"]; ok {
			tmp["url_map_type"] = flattenFirewallAccessProxy6ApiGateway6UrlMapType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if cur_v, ok := i["h2-support"]; ok {
			tmp["h2_support"] = flattenFirewallAccessProxy6ApiGateway6H2Support(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if cur_v, ok := i["h3-support"]; ok {
			tmp["h3_support"] = flattenFirewallAccessProxy6ApiGateway6H3Support(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if cur_v, ok := i["quic"]; ok {
			tmp["quic"] = flattenFirewallAccessProxy6ApiGateway6Quic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if cur_v, ok := i["realservers"]; ok {
			tmp["realservers"] = flattenFirewallAccessProxy6ApiGateway6Realservers(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if cur_v, ok := i["application"]; ok {
			tmp["application"] = flattenFirewallAccessProxy6ApiGateway6Application(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if cur_v, ok := i["persistence"]; ok {
			tmp["persistence"] = flattenFirewallAccessProxy6ApiGateway6Persistence(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if cur_v, ok := i["http-cookie-domain-from-host"]; ok {
			tmp["http_cookie_domain_from_host"] = flattenFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if cur_v, ok := i["http-cookie-domain"]; ok {
			tmp["http_cookie_domain"] = flattenFirewallAccessProxy6ApiGateway6HttpCookieDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if cur_v, ok := i["http-cookie-path"]; ok {
			tmp["http_cookie_path"] = flattenFirewallAccessProxy6ApiGateway6HttpCookiePath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if cur_v, ok := i["http-cookie-generation"]; ok {
			tmp["http_cookie_generation"] = flattenFirewallAccessProxy6ApiGateway6HttpCookieGeneration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if cur_v, ok := i["http-cookie-age"]; ok {
			tmp["http_cookie_age"] = flattenFirewallAccessProxy6ApiGateway6HttpCookieAge(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if cur_v, ok := i["http-cookie-share"]; ok {
			tmp["http_cookie_share"] = flattenFirewallAccessProxy6ApiGateway6HttpCookieShare(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if cur_v, ok := i["https-cookie-secure"]; ok {
			tmp["https_cookie_secure"] = flattenFirewallAccessProxy6ApiGateway6HttpsCookieSecure(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if cur_v, ok := i["saml-server"]; ok {
			tmp["saml_server"] = flattenFirewallAccessProxy6ApiGateway6SamlServer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if cur_v, ok := i["saml-redirect"]; ok {
			tmp["saml_redirect"] = flattenFirewallAccessProxy6ApiGateway6SamlRedirect(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if cur_v, ok := i["ssl-dh-bits"]; ok {
			tmp["ssl_dh_bits"] = flattenFirewallAccessProxy6ApiGateway6SslDhBits(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if cur_v, ok := i["ssl-algorithm"]; ok {
			tmp["ssl_algorithm"] = flattenFirewallAccessProxy6ApiGateway6SslAlgorithm(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if cur_v, ok := i["ssl-cipher-suites"]; ok {
			tmp["ssl_cipher_suites"] = flattenFirewallAccessProxy6ApiGateway6SslCipherSuites(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if cur_v, ok := i["ssl-min-version"]; ok {
			tmp["ssl_min_version"] = flattenFirewallAccessProxy6ApiGateway6SslMinVersion(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if cur_v, ok := i["ssl-max-version"]; ok {
			tmp["ssl_max_version"] = flattenFirewallAccessProxy6ApiGateway6SslMaxVersion(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if cur_v, ok := i["ssl-renegotiation"]; ok {
			tmp["ssl_renegotiation"] = flattenFirewallAccessProxy6ApiGateway6SslRenegotiation(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if cur_v, ok := i["ssl-vpn-web-portal"]; ok {
			tmp["ssl_vpn_web_portal"] = flattenFirewallAccessProxy6ApiGateway6SslVpnWebPortal(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxy6ApiGateway6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6UrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Service(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6LdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6VirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6UrlMapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6H2Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6H3Support(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Quic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := i["max-idle-timeout"]; ok {
		result["max_idle_timeout"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(i["max-idle-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := i["max-udp-payload-size"]; ok {
		result["max_udp_payload_size"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(i["max-udp-payload-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := i["active-connection-id-limit"]; ok {
		result["active_connection_id_limit"] = flattenFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(i["active-connection-id-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := i["ack-delay-exponent"]; ok {
		result["ack_delay_exponent"] = flattenFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(i["ack-delay-exponent"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := i["max-ack-delay"]; ok {
		result["max_ack_delay"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(i["max-ack-delay"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := i["max-datagram-frame-size"]; ok {
		result["max_datagram_frame_size"] = flattenFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(i["max-datagram-frame-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "active_migration"
	if _, ok := i["active-migration"]; ok {
		result["active_migration"] = flattenFirewallAccessProxy6ApiGateway6QuicActiveMigration(i["active-migration"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := i["grease-quic-bit"]; ok {
		result["grease_quic_bit"] = flattenFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(i["grease-quic-bit"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6QuicActiveMigration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Realservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallAccessProxy6ApiGateway6RealserversId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if cur_v, ok := i["addr-type"]; ok {
			tmp["addr_type"] = flattenFirewallAccessProxy6ApiGateway6RealserversAddrType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if cur_v, ok := i["address"]; ok {
			tmp["address"] = flattenFirewallAccessProxy6ApiGateway6RealserversAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenFirewallAccessProxy6ApiGateway6RealserversIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if cur_v, ok := i["domain"]; ok {
			tmp["domain"] = flattenFirewallAccessProxy6ApiGateway6RealserversDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenFirewallAccessProxy6ApiGateway6RealserversPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if cur_v, ok := i["mappedport"]; ok {
			tmp["mappedport"] = flattenFirewallAccessProxy6ApiGateway6RealserversMappedport(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenFirewallAccessProxy6ApiGateway6RealserversStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenFirewallAccessProxy6ApiGateway6RealserversType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if cur_v, ok := i["external-auth"]; ok {
			tmp["external_auth"] = flattenFirewallAccessProxy6ApiGateway6RealserversExternalAuth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if cur_v, ok := i["tunnel-encryption"]; ok {
			tmp["tunnel_encryption"] = flattenFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if cur_v, ok := i["weight"]; ok {
			tmp["weight"] = flattenFirewallAccessProxy6ApiGateway6RealserversWeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if cur_v, ok := i["http-host"]; ok {
			tmp["http_host"] = flattenFirewallAccessProxy6ApiGateway6RealserversHttpHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if cur_v, ok := i["health-check"]; ok {
			tmp["health_check"] = flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheck(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if cur_v, ok := i["health-check-proto"]; ok {
			tmp["health_check_proto"] = flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if cur_v, ok := i["holddown-interval"]; ok {
			tmp["holddown_interval"] = flattenFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if cur_v, ok := i["translate-host"]; ok {
			tmp["translate_host"] = flattenFirewallAccessProxy6ApiGateway6RealserversTranslateHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if cur_v, ok := i["ssh-client-cert"]; ok {
			tmp["ssh_client_cert"] = flattenFirewallAccessProxy6ApiGateway6RealserversSshClientCert(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if cur_v, ok := i["ssh-host-key-validation"]; ok {
			tmp["ssh_host_key_validation"] = flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if cur_v, ok := i["ssh-host-key"]; ok {
			tmp["ssh_host_key"] = flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallAccessProxy6ApiGateway6RealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversMappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversExternalAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversTranslateHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAccessProxy6ApiGateway6RealserversSshHostKeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Application(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallAccessProxy6ApiGateway6ApplicationName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAccessProxy6ApiGateway6ApplicationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6Persistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SamlServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SamlRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["priority"] = flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if cur_v, ok := i["cipher"]; ok {
			tmp["cipher"] = flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if cur_v, ok := i["versions"]; ok {
			tmp["versions"] = flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAccessProxy6ApiGateway6SslVpnWebPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAccessProxy6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenFirewallAccessProxy6Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vip", flattenFirewallAccessProxy6Vip(o["vip"], d, "vip", sv)); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("client_cert", flattenFirewallAccessProxy6ClientCert(o["client-cert"], d, "client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("user_agent_detect", flattenFirewallAccessProxy6UserAgentDetect(o["user-agent-detect"], d, "user_agent_detect", sv)); err != nil {
		if !fortiAPIPatch(o["user-agent-detect"]) {
			return fmt.Errorf("Error reading user_agent_detect: %v", err)
		}
	}

	if err = d.Set("auth_portal", flattenFirewallAccessProxy6AuthPortal(o["auth-portal"], d, "auth_portal", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal"]) {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenFirewallAccessProxy6AuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host", sv)); err != nil {
		if !fortiAPIPatch(o["auth-virtual-host"]) {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("empty_cert_action", flattenFirewallAccessProxy6EmptyCertAction(o["empty-cert-action"], d, "empty_cert_action", sv)); err != nil {
		if !fortiAPIPatch(o["empty-cert-action"]) {
			return fmt.Errorf("Error reading empty_cert_action: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenFirewallAccessProxy6LogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["log-blocked-traffic"]) {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("add_vhost_domain_to_dnsdb", flattenFirewallAccessProxy6AddVhostDomainToDnsdb(o["add-vhost-domain-to-dnsdb"], d, "add_vhost_domain_to_dnsdb", sv)); err != nil {
		if !fortiAPIPatch(o["add-vhost-domain-to-dnsdb"]) {
			return fmt.Errorf("Error reading add_vhost_domain_to_dnsdb: %v", err)
		}
	}

	if err = d.Set("http_supported_max_version", flattenFirewallAccessProxy6HttpSupportedMaxVersion(o["http-supported-max-version"], d, "http_supported_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["http-supported-max-version"]) {
			return fmt.Errorf("Error reading http_supported_max_version: %v", err)
		}
	}

	if err = d.Set("svr_pool_multiplex", flattenFirewallAccessProxy6SvrPoolMultiplex(o["svr-pool-multiplex"], d, "svr_pool_multiplex", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-multiplex"]) {
			return fmt.Errorf("Error reading svr_pool_multiplex: %v", err)
		}
	}

	if err = d.Set("svr_pool_ttl", flattenFirewallAccessProxy6SvrPoolTtl(o["svr-pool-ttl"], d, "svr_pool_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-ttl"]) {
			return fmt.Errorf("Error reading svr_pool_ttl: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_request", flattenFirewallAccessProxy6SvrPoolServerMaxRequest(o["svr-pool-server-max-request"], d, "svr_pool_server_max_request", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-server-max-request"]) {
			return fmt.Errorf("Error reading svr_pool_server_max_request: %v", err)
		}
	}

	if err = d.Set("svr_pool_server_max_concurrent_request", flattenFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(o["svr-pool-server-max-concurrent-request"], d, "svr_pool_server_max_concurrent_request", sv)); err != nil {
		if !fortiAPIPatch(o["svr-pool-server-max-concurrent-request"]) {
			return fmt.Errorf("Error reading svr_pool_server_max_concurrent_request: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallAccessProxy6DecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("api_gateway", flattenFirewallAccessProxy6ApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway"]) {
				return fmt.Errorf("Error reading api_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway"); ok {
			if err = d.Set("api_gateway", flattenFirewallAccessProxy6ApiGateway(o["api-gateway"], d, "api_gateway", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway"]) {
					return fmt.Errorf("Error reading api_gateway: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("api_gateway6", flattenFirewallAccessProxy6ApiGateway6(o["api-gateway6"], d, "api_gateway6", sv)); err != nil {
			if !fortiAPIPatch(o["api-gateway6"]) {
				return fmt.Errorf("Error reading api_gateway6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("api_gateway6"); ok {
			if err = d.Set("api_gateway6", flattenFirewallAccessProxy6ApiGateway6(o["api-gateway6"], d, "api_gateway6", sv)); err != nil {
				if !fortiAPIPatch(o["api-gateway6"]) {
					return fmt.Errorf("Error reading api_gateway6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallAccessProxy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallAccessProxy6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6Vip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6UserAgentDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6AuthPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6AuthVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6EmptyCertAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6LogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6AddVhostDomainToDnsdb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6HttpSupportedMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolMultiplex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolServerMaxRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6DecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallAccessProxy6ApiGatewayId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map"], _ = expandFirewallAccessProxy6ApiGatewayUrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["service"], _ = expandFirewallAccessProxy6ApiGatewayService(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ldb-method"], _ = expandFirewallAccessProxy6ApiGatewayLdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-host"], _ = expandFirewallAccessProxy6ApiGatewayVirtualHost(d, i["virtual_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["virtual-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map-type"], _ = expandFirewallAccessProxy6ApiGatewayUrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h2-support"], _ = expandFirewallAccessProxy6ApiGatewayH2Support(d, i["h2_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h3-support"], _ = expandFirewallAccessProxy6ApiGatewayH3Support(d, i["h3_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quic"], _ = expandFirewallAccessProxy6ApiGatewayQuic(d, i["quic"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["quic"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["realservers"], _ = expandFirewallAccessProxy6ApiGatewayRealservers(d, i["realservers"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["application"], _ = expandFirewallAccessProxy6ApiGatewayApplication(d, i["application"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["persistence"], _ = expandFirewallAccessProxy6ApiGatewayPersistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-path"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-path"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-generation"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-generation"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-age"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-share"], _ = expandFirewallAccessProxy6ApiGatewayHttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-cookie-secure"], _ = expandFirewallAccessProxy6ApiGatewayHttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["saml-server"], _ = expandFirewallAccessProxy6ApiGatewaySamlServer(d, i["saml_server"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["saml-server"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["saml-redirect"], _ = expandFirewallAccessProxy6ApiGatewaySamlRedirect(d, i["saml_redirect"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxy6ApiGatewaySslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-algorithm"], _ = expandFirewallAccessProxy6ApiGatewaySslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-cipher-suites"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-min-version"], _ = expandFirewallAccessProxy6ApiGatewaySslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-max-version"], _ = expandFirewallAccessProxy6ApiGatewaySslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-renegotiation"], _ = expandFirewallAccessProxy6ApiGatewaySslRenegotiation(d, i["ssl_renegotiation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxy6ApiGatewaySslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssl-vpn-web-portal"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayUrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayUrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayH2Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayH3Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-idle-timeout"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-ack-delay"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(d, i["max_ack_delay"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-migration"], _ = expandFirewallAccessProxy6ApiGatewayQuicActiveMigration(d, i["active_migration"], pre_append, sv)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok {
		result["grease-quic-bit"], _ = expandFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicActiveMigration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayQuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallAccessProxy6ApiGatewayRealserversId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandFirewallAccessProxy6ApiGatewayRealserversAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandFirewallAccessProxy6ApiGatewayRealserversAddress(d, i["address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFirewallAccessProxy6ApiGatewayRealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandFirewallAccessProxy6ApiGatewayRealserversDomain(d, i["domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandFirewallAccessProxy6ApiGatewayRealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mappedport"], _ = expandFirewallAccessProxy6ApiGatewayRealserversMappedport(d, i["mappedport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mappedport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandFirewallAccessProxy6ApiGatewayRealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandFirewallAccessProxy6ApiGatewayRealserversType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["external-auth"], _ = expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth(d, i["external_auth"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(d, i["tunnel_encryption"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandFirewallAccessProxy6ApiGatewayRealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-host"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHttpHost(d, i["http_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check-proto"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holddown-interval"], _ = expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["translate-host"], _ = expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost(d, i["translate_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert(d, i["ssh_client_cert"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssh-client-cert"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey(d, i["ssh_host_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssh-host-key"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversMappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversExternalAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTunnelEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayRealserversSshHostKeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallAccessProxy6ApiGatewayApplicationName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewayApplicationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewayHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySamlServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySamlRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["priority"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cipher"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cipher"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["versions"], _ = expandFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGatewaySslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallAccessProxy6ApiGateway6Id(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map"], _ = expandFirewallAccessProxy6ApiGateway6UrlMap(d, i["url_map"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["service"], _ = expandFirewallAccessProxy6ApiGateway6Service(d, i["service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ldb_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ldb-method"], _ = expandFirewallAccessProxy6ApiGateway6LdbMethod(d, i["ldb_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-host"], _ = expandFirewallAccessProxy6ApiGateway6VirtualHost(d, i["virtual_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["virtual-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_map_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-map-type"], _ = expandFirewallAccessProxy6ApiGateway6UrlMapType(d, i["url_map_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h2_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h2-support"], _ = expandFirewallAccessProxy6ApiGateway6H2Support(d, i["h2_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "h3_support"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["h3-support"], _ = expandFirewallAccessProxy6ApiGateway6H3Support(d, i["h3_support"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quic"], _ = expandFirewallAccessProxy6ApiGateway6Quic(d, i["quic"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["quic"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "realservers"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["realservers"], _ = expandFirewallAccessProxy6ApiGateway6Realservers(d, i["realservers"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["realservers"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["application"], _ = expandFirewallAccessProxy6ApiGateway6Application(d, i["application"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "persistence"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["persistence"], _ = expandFirewallAccessProxy6ApiGateway6Persistence(d, i["persistence"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain_from_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain-from-host"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(d, i["http_cookie_domain_from_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-domain"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieDomain(d, i["http_cookie_domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-path"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookiePath(d, i["http_cookie_path"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-path"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_generation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-generation"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieGeneration(d, i["http_cookie_generation"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-cookie-generation"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_age"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-age"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieAge(d, i["http_cookie_age"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_cookie_share"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-cookie-share"], _ = expandFirewallAccessProxy6ApiGateway6HttpCookieShare(d, i["http_cookie_share"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_cookie_secure"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-cookie-secure"], _ = expandFirewallAccessProxy6ApiGateway6HttpsCookieSecure(d, i["https_cookie_secure"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["saml-server"], _ = expandFirewallAccessProxy6ApiGateway6SamlServer(d, i["saml_server"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["saml-server"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "saml_redirect"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["saml-redirect"], _ = expandFirewallAccessProxy6ApiGateway6SamlRedirect(d, i["saml_redirect"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_dh_bits"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-dh-bits"], _ = expandFirewallAccessProxy6ApiGateway6SslDhBits(d, i["ssl_dh_bits"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_algorithm"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-algorithm"], _ = expandFirewallAccessProxy6ApiGateway6SslAlgorithm(d, i["ssl_algorithm"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_cipher_suites"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-cipher-suites"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuites(d, i["ssl_cipher_suites"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssl-cipher-suites"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_min_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-min-version"], _ = expandFirewallAccessProxy6ApiGateway6SslMinVersion(d, i["ssl_min_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_max_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-max-version"], _ = expandFirewallAccessProxy6ApiGateway6SslMaxVersion(d, i["ssl_max_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_renegotiation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-renegotiation"], _ = expandFirewallAccessProxy6ApiGateway6SslRenegotiation(d, i["ssl_renegotiation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssl_vpn_web_portal"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssl-vpn-web-portal"], _ = expandFirewallAccessProxy6ApiGateway6SslVpnWebPortal(d, i["ssl_vpn_web_portal"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssl-vpn-web-portal"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6UrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Service(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6LdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6VirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6UrlMapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6H2Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6H3Support(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Quic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "max_idle_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-idle-timeout"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(d, i["max_idle_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_udp_payload_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-udp-payload-size"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(d, i["max_udp_payload_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_connection_id_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-connection-id-limit"], _ = expandFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(d, i["active_connection_id_limit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ack_delay_exponent"
	if _, ok := d.GetOk(pre_append); ok {
		result["ack-delay-exponent"], _ = expandFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(d, i["ack_delay_exponent"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_ack_delay"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-ack-delay"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(d, i["max_ack_delay"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_datagram_frame_size"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-datagram-frame-size"], _ = expandFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(d, i["max_datagram_frame_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "active_migration"
	if _, ok := d.GetOk(pre_append); ok {
		result["active-migration"], _ = expandFirewallAccessProxy6ApiGateway6QuicActiveMigration(d, i["active_migration"], pre_append, sv)
	}
	pre_append = pre + ".0." + "grease_quic_bit"
	if _, ok := d.GetOk(pre_append); ok {
		result["grease-quic-bit"], _ = expandFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(d, i["grease_quic_bit"], pre_append, sv)
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxUdpPayloadSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicActiveConnectionIdLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicAckDelayExponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxAckDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicMaxDatagramFrameSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicActiveMigration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6QuicGreaseQuicBit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallAccessProxy6ApiGateway6RealserversId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandFirewallAccessProxy6ApiGateway6RealserversAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandFirewallAccessProxy6ApiGateway6RealserversAddress(d, i["address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandFirewallAccessProxy6ApiGateway6RealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandFirewallAccessProxy6ApiGateway6RealserversDomain(d, i["domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandFirewallAccessProxy6ApiGateway6RealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mappedport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mappedport"], _ = expandFirewallAccessProxy6ApiGateway6RealserversMappedport(d, i["mappedport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mappedport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandFirewallAccessProxy6ApiGateway6RealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandFirewallAccessProxy6ApiGateway6RealserversType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_auth"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["external-auth"], _ = expandFirewallAccessProxy6ApiGateway6RealserversExternalAuth(d, i["external_auth"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_encryption"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tunnel-encryption"], _ = expandFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(d, i["tunnel_encryption"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandFirewallAccessProxy6ApiGateway6RealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-host"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHttpHost(d, i["http_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check_proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check-proto"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(d, i["health_check_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holddown-interval"], _ = expandFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "translate_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["translate-host"], _ = expandFirewallAccessProxy6ApiGateway6RealserversTranslateHost(d, i["translate_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_client_cert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssh-client-cert"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshClientCert(d, i["ssh_client_cert"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssh-client-cert"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key_validation"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssh-host-key-validation"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(d, i["ssh_host_key_validation"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssh_host_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssh-host-key"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey(d, i["ssh_host_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ssh-host-key"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversMappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversExternalAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversTunnelEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHealthCheckProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversTranslateHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6RealserversSshHostKeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Application(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallAccessProxy6ApiGateway6ApplicationName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6ApplicationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6Persistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SamlServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SamlRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["priority"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cipher"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cipher"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["versions"], _ = expandFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAccessProxy6ApiGateway6SslVpnWebPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAccessProxy6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallAccessProxy6Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("vip"); ok {
		t, err := expandFirewallAccessProxy6Vip(d, v, "vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	} else if d.HasChange("vip") {
		obj["vip"] = nil
	}

	if v, ok := d.GetOk("client_cert"); ok {
		t, err := expandFirewallAccessProxy6ClientCert(d, v, "client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-cert"] = t
		}
	}

	if v, ok := d.GetOk("user_agent_detect"); ok {
		t, err := expandFirewallAccessProxy6UserAgentDetect(d, v, "user_agent_detect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-agent-detect"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok {
		t, err := expandFirewallAccessProxy6AuthPortal(d, v, "auth_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok {
		t, err := expandFirewallAccessProxy6AuthVirtualHost(d, v, "auth_virtual_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	} else if d.HasChange("auth_virtual_host") {
		obj["auth-virtual-host"] = nil
	}

	if v, ok := d.GetOk("empty_cert_action"); ok {
		t, err := expandFirewallAccessProxy6EmptyCertAction(d, v, "empty_cert_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["empty-cert-action"] = t
		}
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok {
		t, err := expandFirewallAccessProxy6LogBlockedTraffic(d, v, "log_blocked_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("add_vhost_domain_to_dnsdb"); ok {
		t, err := expandFirewallAccessProxy6AddVhostDomainToDnsdb(d, v, "add_vhost_domain_to_dnsdb", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-vhost-domain-to-dnsdb"] = t
		}
	}

	if v, ok := d.GetOk("http_supported_max_version"); ok {
		t, err := expandFirewallAccessProxy6HttpSupportedMaxVersion(d, v, "http_supported_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-supported-max-version"] = t
		}
	}

	if v, ok := d.GetOk("svr_pool_multiplex"); ok {
		t, err := expandFirewallAccessProxy6SvrPoolMultiplex(d, v, "svr_pool_multiplex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-multiplex"] = t
		}
	}

	if v, ok := d.GetOkExists("svr_pool_ttl"); ok {
		t, err := expandFirewallAccessProxy6SvrPoolTtl(d, v, "svr_pool_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-ttl"] = t
		}
	}

	if v, ok := d.GetOkExists("svr_pool_server_max_request"); ok {
		t, err := expandFirewallAccessProxy6SvrPoolServerMaxRequest(d, v, "svr_pool_server_max_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-request"] = t
		}
	} else if d.HasChange("svr_pool_server_max_request") {
		obj["svr-pool-server-max-request"] = nil
	}

	if v, ok := d.GetOkExists("svr_pool_server_max_concurrent_request"); ok {
		t, err := expandFirewallAccessProxy6SvrPoolServerMaxConcurrentRequest(d, v, "svr_pool_server_max_concurrent_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["svr-pool-server-max-concurrent-request"] = t
		}
	} else if d.HasChange("svr_pool_server_max_concurrent_request") {
		obj["svr-pool-server-max-concurrent-request"] = nil
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		t, err := expandFirewallAccessProxy6DecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	} else if d.HasChange("decrypted_traffic_mirror") {
		obj["decrypted-traffic-mirror"] = nil
	}

	if v, ok := d.GetOk("api_gateway"); ok || d.HasChange("api_gateway") {
		t, err := expandFirewallAccessProxy6ApiGateway(d, v, "api_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway"] = t
		}
	}

	if v, ok := d.GetOk("api_gateway6"); ok || d.HasChange("api_gateway6") {
		t, err := expandFirewallAccessProxy6ApiGateway6(d, v, "api_gateway6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-gateway6"] = t
		}
	}

	return &obj, nil
}
