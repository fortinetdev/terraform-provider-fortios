// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual IP for IPv4.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallVip() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVipCreate,
		Read:   resourceFirewallVipRead,
		Update: resourceFirewallVipUpdate,
		Delete: resourceFirewallVipDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_mapping_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 604800),
				Optional:     true,
				Computed:     true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"service": &schema.Schema{
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
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extaddr": &schema.Schema{
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
			"mappedip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"mapped_addr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"extintf": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"arp_reply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat_source_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mappedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gratuitous_arp_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"srcintf_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"portmapping_type": &schema.Schema{
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
						"type": &schema.Schema{
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
						"holddown_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(30, 65535),
							Optional:     true,
							Computed:     true,
						},
						"healthcheck": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"max_connections": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"monitor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"client_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
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
			"http_multiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_ip_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_ip_header_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"outlook_web_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weblogic_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"websphere_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
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
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_hpkp_backup": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
			"monitor": &schema.Schema{
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
			"max_embryonic_connections": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),
				Optional:     true,
				Computed:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallVipCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallVip(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallVip(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallVip resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVip")
	}

	return resourceFirewallVipRead(d, m)
}

func resourceFirewallVipUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallVip(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallVip(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVip")
	}

	return resourceFirewallVipRead(d, m)
}

func resourceFirewallVipDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallVip(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVipRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallVip(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVipName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipDnsMappingTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSrcFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {

			tmp["range"] = flattenFirewallVipSrcFilterRange(i["range"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "range", d)
	return result
}

func flattenFirewallVipSrcFilterRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallVipServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallVipServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallVipExtaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallVipExtaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedip(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {

			tmp["range"] = flattenFirewallVipMappedipRange(i["range"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "range", d)
	return result
}

func flattenFirewallVipMappedipRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtintf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipArpReply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipPersistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipNatSourceVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipPortforward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipExtport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipGratuitousArpInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSrcintfFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {

			tmp["interface_name"] = flattenFirewallVipSrcintfFilterInterfaceName(i["interface-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenFirewallVipSrcintfFilterInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipPortmappingType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallVipRealserversId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenFirewallVipRealserversType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {

			tmp["address"] = flattenFirewallVipRealserversAddress(i["address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallVipRealserversIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallVipRealserversPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallVipRealserversStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenFirewallVipRealserversWeight(i["weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {

			tmp["holddown_interval"] = flattenFirewallVipRealserversHolddownInterval(i["holddown-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {

			tmp["healthcheck"] = flattenFirewallVipRealserversHealthcheck(i["healthcheck"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {

			tmp["http_host"] = flattenFirewallVipRealserversHttpHost(i["http-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {

			tmp["max_connections"] = flattenFirewallVipRealserversMaxConnections(i["max-connections"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {

			v := flattenFirewallVipRealserversMonitor(i["monitor"], d, pre_append, sv)
			vx := ""
			bstring := false
			if i2ss2arrFortiAPIUpgrade(sv, "6.4.2") == true {
				l := v.([]interface{})
				if len(l) > 0 {
					for k, r := range l {
						i := r.(map[string]interface{})
						if _, ok := i["name"]; ok {
							if xv, ok := i["name"].(string); ok {
								vx += "\"" + xv + "\""
								if k < len(l)-1 {
									vx += " "
								}
							}
						}
					}
					bstring = true
				}
			}
			if bstring == true {
				tmp["monitor"] = vx
			} else {
				tmp["monitor"] = v
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := i["client-ip"]; ok {

			tmp["client_ip"] = flattenFirewallVipRealserversClientIp(i["client-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallVipRealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipRealserversClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpMultiplex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpIpHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipHttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipOutlookWebAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipWeblogicServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipWebsphereServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenFirewallVipSslCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallVipSslCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallVipSslCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenFirewallVipSslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenFirewallVipSslServerCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallVipSslServerCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallVipSslServerCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "priority", d)
	return result
}

func flattenFirewallVipSslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslPfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslClientFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHpkp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHsts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHstsAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipSslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallVipMonitorName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallVipMonitorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipMaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVipColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallVip(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallVipName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVipId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVipUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallVipComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallVipType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("dns_mapping_ttl", flattenFirewallVipDnsMappingTtl(o["dns-mapping-ttl"], d, "dns_mapping_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["dns-mapping-ttl"]) {
			return fmt.Errorf("Error reading dns_mapping_ttl: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallVipLdbMethod(o["ldb-method"], d, "ldb_method", sv)); err != nil {
		if !fortiAPIPatch(o["ldb-method"]) {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_filter", flattenFirewallVipSrcFilter(o["src-filter"], d, "src_filter", sv)); err != nil {
			if !fortiAPIPatch(o["src-filter"]) {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_filter"); ok {
			if err = d.Set("src_filter", flattenFirewallVipSrcFilter(o["src-filter"], d, "src_filter", sv)); err != nil {
				if !fortiAPIPatch(o["src-filter"]) {
					return fmt.Errorf("Error reading src_filter: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallVipService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallVipService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("extip", flattenFirewallVipExtip(o["extip"], d, "extip", sv)); err != nil {
		if !fortiAPIPatch(o["extip"]) {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("extaddr", flattenFirewallVipExtaddr(o["extaddr"], d, "extaddr", sv)); err != nil {
			if !fortiAPIPatch(o["extaddr"]) {
				return fmt.Errorf("Error reading extaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("extaddr"); ok {
			if err = d.Set("extaddr", flattenFirewallVipExtaddr(o["extaddr"], d, "extaddr", sv)); err != nil {
				if !fortiAPIPatch(o["extaddr"]) {
					return fmt.Errorf("Error reading extaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mappedip", flattenFirewallVipMappedip(o["mappedip"], d, "mappedip", sv)); err != nil {
			if !fortiAPIPatch(o["mappedip"]) {
				return fmt.Errorf("Error reading mappedip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mappedip"); ok {
			if err = d.Set("mappedip", flattenFirewallVipMappedip(o["mappedip"], d, "mappedip", sv)); err != nil {
				if !fortiAPIPatch(o["mappedip"]) {
					return fmt.Errorf("Error reading mappedip: %v", err)
				}
			}
		}
	}

	if err = d.Set("mapped_addr", flattenFirewallVipMappedAddr(o["mapped-addr"], d, "mapped_addr", sv)); err != nil {
		if !fortiAPIPatch(o["mapped-addr"]) {
			return fmt.Errorf("Error reading mapped_addr: %v", err)
		}
	}

	if err = d.Set("extintf", flattenFirewallVipExtintf(o["extintf"], d, "extintf", sv)); err != nil {
		if !fortiAPIPatch(o["extintf"]) {
			return fmt.Errorf("Error reading extintf: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallVipArpReply(o["arp-reply"], d, "arp_reply", sv)); err != nil {
		if !fortiAPIPatch(o["arp-reply"]) {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("server_type", flattenFirewallVipServerType(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("http_redirect", flattenFirewallVipHttpRedirect(o["http-redirect"], d, "http_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["http-redirect"]) {
			return fmt.Errorf("Error reading http_redirect: %v", err)
		}
	}

	if err = d.Set("persistence", flattenFirewallVipPersistence(o["persistence"], d, "persistence", sv)); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("nat_source_vip", flattenFirewallVipNatSourceVip(o["nat-source-vip"], d, "nat_source_vip", sv)); err != nil {
		if !fortiAPIPatch(o["nat-source-vip"]) {
			return fmt.Errorf("Error reading nat_source_vip: %v", err)
		}
	}

	if err = d.Set("portforward", flattenFirewallVipPortforward(o["portforward"], d, "portforward", sv)); err != nil {
		if !fortiAPIPatch(o["portforward"]) {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallVipProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("extport", flattenFirewallVipExtport(o["extport"], d, "extport", sv)); err != nil {
		if !fortiAPIPatch(o["extport"]) {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallVipMappedport(o["mappedport"], d, "mappedport", sv)); err != nil {
		if !fortiAPIPatch(o["mappedport"]) {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("gratuitous_arp_interval", flattenFirewallVipGratuitousArpInterval(o["gratuitous-arp-interval"], d, "gratuitous_arp_interval", sv)); err != nil {
		if !fortiAPIPatch(o["gratuitous-arp-interval"]) {
			return fmt.Errorf("Error reading gratuitous_arp_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcintf_filter", flattenFirewallVipSrcintfFilter(o["srcintf-filter"], d, "srcintf_filter", sv)); err != nil {
			if !fortiAPIPatch(o["srcintf-filter"]) {
				return fmt.Errorf("Error reading srcintf_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf_filter"); ok {
			if err = d.Set("srcintf_filter", flattenFirewallVipSrcintfFilter(o["srcintf-filter"], d, "srcintf_filter", sv)); err != nil {
				if !fortiAPIPatch(o["srcintf-filter"]) {
					return fmt.Errorf("Error reading srcintf_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("portmapping_type", flattenFirewallVipPortmappingType(o["portmapping-type"], d, "portmapping_type", sv)); err != nil {
		if !fortiAPIPatch(o["portmapping-type"]) {
			return fmt.Errorf("Error reading portmapping_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenFirewallVipRealservers(o["realservers"], d, "realservers", sv)); err != nil {
			if !fortiAPIPatch(o["realservers"]) {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallVipRealservers(o["realservers"], d, "realservers", sv)); err != nil {
				if !fortiAPIPatch(o["realservers"]) {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenFirewallVipHttpCookieDomainFromHost(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-domain-from-host"]) {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenFirewallVipHttpCookieDomain(o["http-cookie-domain"], d, "http_cookie_domain", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-domain"]) {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenFirewallVipHttpCookiePath(o["http-cookie-path"], d, "http_cookie_path", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-path"]) {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenFirewallVipHttpCookieGeneration(o["http-cookie-generation"], d, "http_cookie_generation", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-generation"]) {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenFirewallVipHttpCookieAge(o["http-cookie-age"], d, "http_cookie_age", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-age"]) {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenFirewallVipHttpCookieShare(o["http-cookie-share"], d, "http_cookie_share", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-share"]) {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenFirewallVipHttpsCookieSecure(o["https-cookie-secure"], d, "https_cookie_secure", sv)); err != nil {
		if !fortiAPIPatch(o["https-cookie-secure"]) {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("http_multiplex", flattenFirewallVipHttpMultiplex(o["http-multiplex"], d, "http_multiplex", sv)); err != nil {
		if !fortiAPIPatch(o["http-multiplex"]) {
			return fmt.Errorf("Error reading http_multiplex: %v", err)
		}
	}

	if err = d.Set("http_ip_header", flattenFirewallVipHttpIpHeader(o["http-ip-header"], d, "http_ip_header", sv)); err != nil {
		if !fortiAPIPatch(o["http-ip-header"]) {
			return fmt.Errorf("Error reading http_ip_header: %v", err)
		}
	}

	if err = d.Set("http_ip_header_name", flattenFirewallVipHttpIpHeaderName(o["http-ip-header-name"], d, "http_ip_header_name", sv)); err != nil {
		if !fortiAPIPatch(o["http-ip-header-name"]) {
			return fmt.Errorf("Error reading http_ip_header_name: %v", err)
		}
	}

	if err = d.Set("outlook_web_access", flattenFirewallVipOutlookWebAccess(o["outlook-web-access"], d, "outlook_web_access", sv)); err != nil {
		if !fortiAPIPatch(o["outlook-web-access"]) {
			return fmt.Errorf("Error reading outlook_web_access: %v", err)
		}
	}

	if err = d.Set("weblogic_server", flattenFirewallVipWeblogicServer(o["weblogic-server"], d, "weblogic_server", sv)); err != nil {
		if !fortiAPIPatch(o["weblogic-server"]) {
			return fmt.Errorf("Error reading weblogic_server: %v", err)
		}
	}

	if err = d.Set("websphere_server", flattenFirewallVipWebsphereServer(o["websphere-server"], d, "websphere_server", sv)); err != nil {
		if !fortiAPIPatch(o["websphere-server"]) {
			return fmt.Errorf("Error reading websphere_server: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenFirewallVipSslMode(o["ssl-mode"], d, "ssl_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mode"]) {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenFirewallVipSslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallVipSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-dh-bits"]) {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallVipSslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenFirewallVipSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-cipher-suites"]) {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenFirewallVipSslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-cipher-suites"]) {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenFirewallVipSslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-algorithm"]) {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenFirewallVipSslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-server-cipher-suites"]) {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenFirewallVipSslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-server-cipher-suites"]) {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_pfs", flattenFirewallVipSslPfs(o["ssl-pfs"], d, "ssl_pfs", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-pfs"]) {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallVipSslMinVersion(o["ssl-min-version"], d, "ssl_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-version"]) {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallVipSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenFirewallVipSslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-min-version"]) {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_max_version", flattenFirewallVipSslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-max-version"]) {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenFirewallVipSslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-send-empty-frags"]) {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_client_fallback", flattenFirewallVipSslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-fallback"]) {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenFirewallVipSslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenFirewallVipSslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-type"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenFirewallVipSslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-timeout"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenFirewallVipSslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-max"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenFirewallVipSslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-rekey-count"]) {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenFirewallVipSslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-type"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenFirewallVipSslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-timeout"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenFirewallVipSslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-max"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenFirewallVipSslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-http-location-conversion"]) {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenFirewallVipSslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-http-match-host"]) {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenFirewallVipSslHpkp(o["ssl-hpkp"], d, "ssl_hpkp", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp"]) {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenFirewallVipSslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-primary"]) {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenFirewallVipSslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-backup"]) {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenFirewallVipSslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-age"]) {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenFirewallVipSslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-report-uri"]) {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenFirewallVipSslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-include-subdomains"]) {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenFirewallVipSslHsts(o["ssl-hsts"], d, "ssl_hsts", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts"]) {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenFirewallVipSslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts-age"]) {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenFirewallVipSslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts-include-subdomains"]) {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("monitor", flattenFirewallVipMonitor(o["monitor"], d, "monitor", sv)); err != nil {
			if !fortiAPIPatch(o["monitor"]) {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitor"); ok {
			if err = d.Set("monitor", flattenFirewallVipMonitor(o["monitor"], d, "monitor", sv)); err != nil {
				if !fortiAPIPatch(o["monitor"]) {
					return fmt.Errorf("Error reading monitor: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_embryonic_connections", flattenFirewallVipMaxEmbryonicConnections(o["max-embryonic-connections"], d, "max_embryonic_connections", sv)); err != nil {
		if !fortiAPIPatch(o["max-embryonic-connections"]) {
			return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallVipColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	return nil
}

func flattenFirewallVipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallVipName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipDnsMappingTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSrcFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["range"], _ = expandFirewallVipSrcFilterRange(d, i["range"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipSrcFilterRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallVipServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallVipExtaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipExtaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["range"], _ = expandFirewallVipMappedipRange(d, i["range"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipMappedipRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipArpReply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPersistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipNatSourceVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPortforward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipExtport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipGratuitousArpInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSrcintfFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface-name"], _ = expandFirewallVipSrcintfFilterInterfaceName(d, i["interface_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipSrcintfFilterInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipPortmappingType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallVipRealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandFirewallVipRealserversType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["address"], _ = expandFirewallVipRealserversAddress(d, i["address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandFirewallVipRealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallVipRealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallVipRealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandFirewallVipRealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["holddown-interval"], _ = expandFirewallVipRealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["healthcheck"], _ = expandFirewallVipRealserversHealthcheck(d, i["healthcheck"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-host"], _ = expandFirewallVipRealserversHttpHost(d, i["http_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max-connections"], _ = expandFirewallVipRealserversMaxConnections(d, i["max_connections"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok {

			bstring := false
			t, _ := expandFirewallVipRealserversMonitor(d, i["monitor"], pre_append, sv)
			if t != nil {
				if i2ss2arrFortiAPIUpgrade(sv, "6.4.2") == true {
					bstring = true
				}
			}

			if bstring == true {
				vx := fmt.Sprintf("%v", t)
				vx = strings.Replace(vx, "\"", "", -1)
				vxx := strings.Split(vx, " ")

				tmps := make([]map[string]interface{}, 0, len(vxx))

				for _, xv := range vxx {
					xtmp := make(map[string]interface{})
					xtmp["name"] = xv

					tmps = append(tmps, xtmp)
				}
				tmp["monitor"] = tmps
			} else {
				tmp["monitor"] = t
			}

		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "client_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["client-ip"], _ = expandFirewallVipRealserversClientIp(d, i["client_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipRealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipRealserversClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpMultiplex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpIpHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipHttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipOutlookWebAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipWeblogicServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipWebsphereServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandFirewallVipSslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallVipSslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallVipSslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipSslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandFirewallVipSslServerCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallVipSslServerCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallVipSslServerCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipSslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslPfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHsts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHstsAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipSslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallVipMonitorName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVipMonitorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipMaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVipColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallVipName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandFirewallVipId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {

		t, err := expandFirewallVipUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallVipComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallVipType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOkExists("dns_mapping_ttl"); ok {

		t, err := expandFirewallVipDnsMappingTtl(d, v, "dns_mapping_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mapping-ttl"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok {

		t, err := expandFirewallVipLdbMethod(d, v, "ldb_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok {

		t, err := expandFirewallVipSrcFilter(d, v, "src_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandFirewallVipService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok {

		t, err := expandFirewallVipExtip(d, v, "extip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("extaddr"); ok {

		t, err := expandFirewallVipExtaddr(d, v, "extaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extaddr"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok {

		t, err := expandFirewallVipMappedip(d, v, "mappedip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("mapped_addr"); ok {

		t, err := expandFirewallVipMappedAddr(d, v, "mapped_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapped-addr"] = t
		}
	}

	if v, ok := d.GetOk("extintf"); ok {

		t, err := expandFirewallVipExtintf(d, v, "extintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extintf"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok {

		t, err := expandFirewallVipArpReply(d, v, "arp_reply", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {

		t, err := expandFirewallVipServerType(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("http_redirect"); ok {

		t, err := expandFirewallVipHttpRedirect(d, v, "http_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-redirect"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok {

		t, err := expandFirewallVipPersistence(d, v, "persistence", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("nat_source_vip"); ok {

		t, err := expandFirewallVipNatSourceVip(d, v, "nat_source_vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-source-vip"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok {

		t, err := expandFirewallVipPortforward(d, v, "portforward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {

		t, err := expandFirewallVipProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok {

		t, err := expandFirewallVipExtport(d, v, "extport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok {

		t, err := expandFirewallVipMappedport(d, v, "mappedport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOk("gratuitous_arp_interval"); ok {

		t, err := expandFirewallVipGratuitousArpInterval(d, v, "gratuitous_arp_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gratuitous-arp-interval"] = t
		}
	}

	if v, ok := d.GetOk("srcintf_filter"); ok {

		t, err := expandFirewallVipSrcintfFilter(d, v, "srcintf_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf-filter"] = t
		}
	}

	if v, ok := d.GetOk("portmapping_type"); ok {

		t, err := expandFirewallVipPortmappingType(d, v, "portmapping_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portmapping-type"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok {

		t, err := expandFirewallVipRealservers(d, v, "realservers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok {

		t, err := expandFirewallVipHttpCookieDomainFromHost(d, v, "http_cookie_domain_from_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok {

		t, err := expandFirewallVipHttpCookieDomain(d, v, "http_cookie_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok {

		t, err := expandFirewallVipHttpCookiePath(d, v, "http_cookie_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOkExists("http_cookie_generation"); ok {

		t, err := expandFirewallVipHttpCookieGeneration(d, v, "http_cookie_generation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOkExists("http_cookie_age"); ok {

		t, err := expandFirewallVipHttpCookieAge(d, v, "http_cookie_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok {

		t, err := expandFirewallVipHttpCookieShare(d, v, "http_cookie_share", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok {

		t, err := expandFirewallVipHttpsCookieSecure(d, v, "https_cookie_secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex"); ok {

		t, err := expandFirewallVipHttpMultiplex(d, v, "http_multiplex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header"); ok {

		t, err := expandFirewallVipHttpIpHeader(d, v, "http_ip_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header_name"); ok {

		t, err := expandFirewallVipHttpIpHeaderName(d, v, "http_ip_header_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header-name"] = t
		}
	}

	if v, ok := d.GetOk("outlook_web_access"); ok {

		t, err := expandFirewallVipOutlookWebAccess(d, v, "outlook_web_access", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outlook-web-access"] = t
		}
	}

	if v, ok := d.GetOk("weblogic_server"); ok {

		t, err := expandFirewallVipWeblogicServer(d, v, "weblogic_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weblogic-server"] = t
		}
	}

	if v, ok := d.GetOk("websphere_server"); ok {

		t, err := expandFirewallVipWebsphereServer(d, v, "websphere_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websphere-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok {

		t, err := expandFirewallVipSslMode(d, v, "ssl_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {

		t, err := expandFirewallVipSslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok {

		t, err := expandFirewallVipSslDhBits(d, v, "ssl_dh_bits", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {

		t, err := expandFirewallVipSslAlgorithm(d, v, "ssl_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok {

		t, err := expandFirewallVipSslCipherSuites(d, v, "ssl_cipher_suites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok {

		t, err := expandFirewallVipSslServerAlgorithm(d, v, "ssl_server_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok {

		t, err := expandFirewallVipSslServerCipherSuites(d, v, "ssl_server_cipher_suites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok {

		t, err := expandFirewallVipSslPfs(d, v, "ssl_pfs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok {

		t, err := expandFirewallVipSslMinVersion(d, v, "ssl_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {

		t, err := expandFirewallVipSslMaxVersion(d, v, "ssl_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok {

		t, err := expandFirewallVipSslServerMinVersion(d, v, "ssl_server_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok {

		t, err := expandFirewallVipSslServerMaxVersion(d, v, "ssl_server_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok {

		t, err := expandFirewallVipSslSendEmptyFrags(d, v, "ssl_send_empty_frags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok {

		t, err := expandFirewallVipSslClientFallback(d, v, "ssl_client_fallback", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok {

		t, err := expandFirewallVipSslClientRenegotiation(d, v, "ssl_client_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok {

		t, err := expandFirewallVipSslClientSessionStateType(d, v, "ssl_client_session_state_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok {

		t, err := expandFirewallVipSslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok {

		t, err := expandFirewallVipSslClientSessionStateMax(d, v, "ssl_client_session_state_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok {

		t, err := expandFirewallVipSslClientRekeyCount(d, v, "ssl_client_rekey_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok {

		t, err := expandFirewallVipSslServerSessionStateType(d, v, "ssl_server_session_state_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok {

		t, err := expandFirewallVipSslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok {

		t, err := expandFirewallVipSslServerSessionStateMax(d, v, "ssl_server_session_state_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok {

		t, err := expandFirewallVipSslHttpLocationConversion(d, v, "ssl_http_location_conversion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok {

		t, err := expandFirewallVipSslHttpMatchHost(d, v, "ssl_http_match_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok {

		t, err := expandFirewallVipSslHpkp(d, v, "ssl_hpkp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok {

		t, err := expandFirewallVipSslHpkpPrimary(d, v, "ssl_hpkp_primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok {

		t, err := expandFirewallVipSslHpkpBackup(d, v, "ssl_hpkp_backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok {

		t, err := expandFirewallVipSslHpkpAge(d, v, "ssl_hpkp_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok {

		t, err := expandFirewallVipSslHpkpReportUri(d, v, "ssl_hpkp_report_uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok {

		t, err := expandFirewallVipSslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok {

		t, err := expandFirewallVipSslHsts(d, v, "ssl_hsts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok {

		t, err := expandFirewallVipSslHstsAge(d, v, "ssl_hsts_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok {

		t, err := expandFirewallVipSslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok {

		t, err := expandFirewallVipMonitor(d, v, "monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOkExists("max_embryonic_connections"); ok {

		t, err := expandFirewallVipMaxEmbryonicConnections(d, v, "max_embryonic_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-embryonic-connections"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandFirewallVipColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	return &obj, nil
}
