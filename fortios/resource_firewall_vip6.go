// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual IP for IPv6.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallVip6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVip6Create,
		Read:   resourceFirewallVip6Read,
		Update: resourceFirewallVip6Update,
		Delete: resourceFirewallVip6Delete,

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
			"src_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"extip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"nat_source_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arp_reply": &schema.Schema{
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
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"ldb_method": &schema.Schema{
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
			"nat66": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_nat64_route": &schema.Schema{
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
							ValidateFunc: validation.StringLenBetween(0, 79),
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
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"ssl_hpkp_backup": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
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
							ValidateFunc: validation.StringLenBetween(0, 79),
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
			"embedded_ipv4_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_mappedip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_mappedport": &schema.Schema{
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

func resourceFirewallVip6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallVip6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVip6 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallVip6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallVip6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVip6")
	}

	return resourceFirewallVip6Read(d, m)
}

func resourceFirewallVip6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallVip6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallVip6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVip6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallVip6")
	}

	return resourceFirewallVip6Read(d, m)
}

func resourceFirewallVip6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallVip6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVip6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVip6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallVip6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVip6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVip6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallVip6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Uuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Comment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Type(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SrcFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["range"] = flattenFirewallVip6SrcFilterRange(i["range"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "range", d)
	return result
}

func flattenFirewallVip6SrcFilterRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Extip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Mappedip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6NatSourceVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6ArpReply(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Portforward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Protocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Extport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Mappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Color(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6LdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6ServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Persistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Nat66(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Nat64(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6AddNat64Route(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Realservers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallVip6RealserversId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenFirewallVip6RealserversIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenFirewallVip6RealserversPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallVip6RealserversStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenFirewallVip6RealserversWeight(i["weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := i["holddown-interval"]; ok {

			tmp["holddown_interval"] = flattenFirewallVip6RealserversHolddownInterval(i["holddown-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := i["healthcheck"]; ok {

			tmp["healthcheck"] = flattenFirewallVip6RealserversHealthcheck(i["healthcheck"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := i["http-host"]; ok {

			tmp["http_host"] = flattenFirewallVip6RealserversHttpHost(i["http-host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := i["max-connections"]; ok {

			tmp["max_connections"] = flattenFirewallVip6RealserversMaxConnections(i["max-connections"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {

			v := flattenFirewallVip6RealserversMonitor(i["monitor"], d, pre_append, sv)
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

			tmp["client_ip"] = flattenFirewallVip6RealserversClientIp(i["client-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallVip6RealserversId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHolddownInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHealthcheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversHttpHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversMaxConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6RealserversClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieDomainFromHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookiePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpCookieShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpsCookieSecure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpMultiplex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpIpHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6HttpIpHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6OutlookWebAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6WeblogicServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6WebsphereServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslDhBits(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["priority"] = flattenFirewallVip6SslCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallVip6SslCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallVip6SslCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "priority", d)
	return result
}

func flattenFirewallVip6SslCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerCipherSuites(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["priority"] = flattenFirewallVip6SslServerCipherSuitesPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := i["cipher"]; ok {

			tmp["cipher"] = flattenFirewallVip6SslServerCipherSuitesCipher(i["cipher"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := i["versions"]; ok {

			tmp["versions"] = flattenFirewallVip6SslServerCipherSuitesVersions(i["versions"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "priority", d)
	return result
}

func flattenFirewallVip6SslServerCipherSuitesPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerCipherSuitesCipher(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerCipherSuitesVersions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslPfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslAcceptFfdheGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslClientFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslClientRenegotiation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslClientSessionStateType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslClientSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslClientSessionStateMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslClientRekeyCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerSessionStateType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerSessionStateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslServerSessionStateMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHttpLocationConversion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHttpMatchHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpBackup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpReportUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHpkpIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHsts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHstsAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6SslHstsIncludeSubdomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Monitor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallVip6MonitorName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallVip6MonitorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6MaxEmbryonicConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6EmbeddedIpv4Address(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Ipv4Mappedip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVip6Ipv4Mappedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallVip6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallVip6Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenFirewallVip6Id(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallVip6Uuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallVip6Comment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallVip6Type(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_filter", flattenFirewallVip6SrcFilter(o["src-filter"], d, "src_filter", sv)); err != nil {
			if !fortiAPIPatch(o["src-filter"]) {
				return fmt.Errorf("Error reading src_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_filter"); ok {
			if err = d.Set("src_filter", flattenFirewallVip6SrcFilter(o["src-filter"], d, "src_filter", sv)); err != nil {
				if !fortiAPIPatch(o["src-filter"]) {
					return fmt.Errorf("Error reading src_filter: %v", err)
				}
			}
		}
	}

	if err = d.Set("extip", flattenFirewallVip6Extip(o["extip"], d, "extip", sv)); err != nil {
		if !fortiAPIPatch(o["extip"]) {
			return fmt.Errorf("Error reading extip: %v", err)
		}
	}

	if err = d.Set("mappedip", flattenFirewallVip6Mappedip(o["mappedip"], d, "mappedip", sv)); err != nil {
		if !fortiAPIPatch(o["mappedip"]) {
			return fmt.Errorf("Error reading mappedip: %v", err)
		}
	}

	if err = d.Set("nat_source_vip", flattenFirewallVip6NatSourceVip(o["nat-source-vip"], d, "nat_source_vip", sv)); err != nil {
		if !fortiAPIPatch(o["nat-source-vip"]) {
			return fmt.Errorf("Error reading nat_source_vip: %v", err)
		}
	}

	if err = d.Set("arp_reply", flattenFirewallVip6ArpReply(o["arp-reply"], d, "arp_reply", sv)); err != nil {
		if !fortiAPIPatch(o["arp-reply"]) {
			return fmt.Errorf("Error reading arp_reply: %v", err)
		}
	}

	if err = d.Set("portforward", flattenFirewallVip6Portforward(o["portforward"], d, "portforward", sv)); err != nil {
		if !fortiAPIPatch(o["portforward"]) {
			return fmt.Errorf("Error reading portforward: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallVip6Protocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("extport", flattenFirewallVip6Extport(o["extport"], d, "extport", sv)); err != nil {
		if !fortiAPIPatch(o["extport"]) {
			return fmt.Errorf("Error reading extport: %v", err)
		}
	}

	if err = d.Set("mappedport", flattenFirewallVip6Mappedport(o["mappedport"], d, "mappedport", sv)); err != nil {
		if !fortiAPIPatch(o["mappedport"]) {
			return fmt.Errorf("Error reading mappedport: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallVip6Color(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenFirewallVip6LdbMethod(o["ldb-method"], d, "ldb_method", sv)); err != nil {
		if !fortiAPIPatch(o["ldb-method"]) {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("server_type", flattenFirewallVip6ServerType(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("http_redirect", flattenFirewallVip6HttpRedirect(o["http-redirect"], d, "http_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["http-redirect"]) {
			return fmt.Errorf("Error reading http_redirect: %v", err)
		}
	}

	if err = d.Set("persistence", flattenFirewallVip6Persistence(o["persistence"], d, "persistence", sv)); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}

	if err = d.Set("nat66", flattenFirewallVip6Nat66(o["nat66"], d, "nat66", sv)); err != nil {
		if !fortiAPIPatch(o["nat66"]) {
			return fmt.Errorf("Error reading nat66: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallVip6Nat64(o["nat64"], d, "nat64", sv)); err != nil {
		if !fortiAPIPatch(o["nat64"]) {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("add_nat64_route", flattenFirewallVip6AddNat64Route(o["add-nat64-route"], d, "add_nat64_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-nat64-route"]) {
			return fmt.Errorf("Error reading add_nat64_route: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("realservers", flattenFirewallVip6Realservers(o["realservers"], d, "realservers", sv)); err != nil {
			if !fortiAPIPatch(o["realservers"]) {
				return fmt.Errorf("Error reading realservers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("realservers"); ok {
			if err = d.Set("realservers", flattenFirewallVip6Realservers(o["realservers"], d, "realservers", sv)); err != nil {
				if !fortiAPIPatch(o["realservers"]) {
					return fmt.Errorf("Error reading realservers: %v", err)
				}
			}
		}
	}

	if err = d.Set("http_cookie_domain_from_host", flattenFirewallVip6HttpCookieDomainFromHost(o["http-cookie-domain-from-host"], d, "http_cookie_domain_from_host", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-domain-from-host"]) {
			return fmt.Errorf("Error reading http_cookie_domain_from_host: %v", err)
		}
	}

	if err = d.Set("http_cookie_domain", flattenFirewallVip6HttpCookieDomain(o["http-cookie-domain"], d, "http_cookie_domain", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-domain"]) {
			return fmt.Errorf("Error reading http_cookie_domain: %v", err)
		}
	}

	if err = d.Set("http_cookie_path", flattenFirewallVip6HttpCookiePath(o["http-cookie-path"], d, "http_cookie_path", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-path"]) {
			return fmt.Errorf("Error reading http_cookie_path: %v", err)
		}
	}

	if err = d.Set("http_cookie_generation", flattenFirewallVip6HttpCookieGeneration(o["http-cookie-generation"], d, "http_cookie_generation", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-generation"]) {
			return fmt.Errorf("Error reading http_cookie_generation: %v", err)
		}
	}

	if err = d.Set("http_cookie_age", flattenFirewallVip6HttpCookieAge(o["http-cookie-age"], d, "http_cookie_age", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-age"]) {
			return fmt.Errorf("Error reading http_cookie_age: %v", err)
		}
	}

	if err = d.Set("http_cookie_share", flattenFirewallVip6HttpCookieShare(o["http-cookie-share"], d, "http_cookie_share", sv)); err != nil {
		if !fortiAPIPatch(o["http-cookie-share"]) {
			return fmt.Errorf("Error reading http_cookie_share: %v", err)
		}
	}

	if err = d.Set("https_cookie_secure", flattenFirewallVip6HttpsCookieSecure(o["https-cookie-secure"], d, "https_cookie_secure", sv)); err != nil {
		if !fortiAPIPatch(o["https-cookie-secure"]) {
			return fmt.Errorf("Error reading https_cookie_secure: %v", err)
		}
	}

	if err = d.Set("http_multiplex", flattenFirewallVip6HttpMultiplex(o["http-multiplex"], d, "http_multiplex", sv)); err != nil {
		if !fortiAPIPatch(o["http-multiplex"]) {
			return fmt.Errorf("Error reading http_multiplex: %v", err)
		}
	}

	if err = d.Set("http_ip_header", flattenFirewallVip6HttpIpHeader(o["http-ip-header"], d, "http_ip_header", sv)); err != nil {
		if !fortiAPIPatch(o["http-ip-header"]) {
			return fmt.Errorf("Error reading http_ip_header: %v", err)
		}
	}

	if err = d.Set("http_ip_header_name", flattenFirewallVip6HttpIpHeaderName(o["http-ip-header-name"], d, "http_ip_header_name", sv)); err != nil {
		if !fortiAPIPatch(o["http-ip-header-name"]) {
			return fmt.Errorf("Error reading http_ip_header_name: %v", err)
		}
	}

	if err = d.Set("outlook_web_access", flattenFirewallVip6OutlookWebAccess(o["outlook-web-access"], d, "outlook_web_access", sv)); err != nil {
		if !fortiAPIPatch(o["outlook-web-access"]) {
			return fmt.Errorf("Error reading outlook_web_access: %v", err)
		}
	}

	if err = d.Set("weblogic_server", flattenFirewallVip6WeblogicServer(o["weblogic-server"], d, "weblogic_server", sv)); err != nil {
		if !fortiAPIPatch(o["weblogic-server"]) {
			return fmt.Errorf("Error reading weblogic_server: %v", err)
		}
	}

	if err = d.Set("websphere_server", flattenFirewallVip6WebsphereServer(o["websphere-server"], d, "websphere_server", sv)); err != nil {
		if !fortiAPIPatch(o["websphere-server"]) {
			return fmt.Errorf("Error reading websphere_server: %v", err)
		}
	}

	if err = d.Set("ssl_mode", flattenFirewallVip6SslMode(o["ssl-mode"], d, "ssl_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mode"]) {
			return fmt.Errorf("Error reading ssl_mode: %v", err)
		}
	}

	if err = d.Set("ssl_certificate", flattenFirewallVip6SslCertificate(o["ssl-certificate"], d, "ssl_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-certificate"]) {
			return fmt.Errorf("Error reading ssl_certificate: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallVip6SslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-dh-bits"]) {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_algorithm", flattenFirewallVip6SslAlgorithm(o["ssl-algorithm"], d, "ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-algorithm"]) {
			return fmt.Errorf("Error reading ssl_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_cipher_suites", flattenFirewallVip6SslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-cipher-suites"]) {
				return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_cipher_suites"); ok {
			if err = d.Set("ssl_cipher_suites", flattenFirewallVip6SslCipherSuites(o["ssl-cipher-suites"], d, "ssl_cipher_suites", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-cipher-suites"]) {
					return fmt.Errorf("Error reading ssl_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_server_algorithm", flattenFirewallVip6SslServerAlgorithm(o["ssl-server-algorithm"], d, "ssl_server_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-algorithm"]) {
			return fmt.Errorf("Error reading ssl_server_algorithm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_server_cipher_suites", flattenFirewallVip6SslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-server-cipher-suites"]) {
				return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_server_cipher_suites"); ok {
			if err = d.Set("ssl_server_cipher_suites", flattenFirewallVip6SslServerCipherSuites(o["ssl-server-cipher-suites"], d, "ssl_server_cipher_suites", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-server-cipher-suites"]) {
					return fmt.Errorf("Error reading ssl_server_cipher_suites: %v", err)
				}
			}
		}
	}

	if err = d.Set("ssl_pfs", flattenFirewallVip6SslPfs(o["ssl-pfs"], d, "ssl_pfs", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-pfs"]) {
			return fmt.Errorf("Error reading ssl_pfs: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenFirewallVip6SslMinVersion(o["ssl-min-version"], d, "ssl_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-version"]) {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenFirewallVip6SslMaxVersion(o["ssl-max-version"], d, "ssl_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_min_version", flattenFirewallVip6SslServerMinVersion(o["ssl-server-min-version"], d, "ssl_server_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-min-version"]) {
			return fmt.Errorf("Error reading ssl_server_min_version: %v", err)
		}
	}

	if err = d.Set("ssl_server_max_version", flattenFirewallVip6SslServerMaxVersion(o["ssl-server-max-version"], d, "ssl_server_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-max-version"]) {
			return fmt.Errorf("Error reading ssl_server_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_accept_ffdhe_groups", flattenFirewallVip6SslAcceptFfdheGroups(o["ssl-accept-ffdhe-groups"], d, "ssl_accept_ffdhe_groups", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-accept-ffdhe-groups"]) {
			return fmt.Errorf("Error reading ssl_accept_ffdhe_groups: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenFirewallVip6SslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-send-empty-frags"]) {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("ssl_client_fallback", flattenFirewallVip6SslClientFallback(o["ssl-client-fallback"], d, "ssl_client_fallback", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-fallback"]) {
			return fmt.Errorf("Error reading ssl_client_fallback: %v", err)
		}
	}

	if err = d.Set("ssl_client_renegotiation", flattenFirewallVip6SslClientRenegotiation(o["ssl-client-renegotiation"], d, "ssl_client_renegotiation", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-renegotiation"]) {
			return fmt.Errorf("Error reading ssl_client_renegotiation: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_type", flattenFirewallVip6SslClientSessionStateType(o["ssl-client-session-state-type"], d, "ssl_client_session_state_type", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-type"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_timeout", flattenFirewallVip6SslClientSessionStateTimeout(o["ssl-client-session-state-timeout"], d, "ssl_client_session_state_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-timeout"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_client_session_state_max", flattenFirewallVip6SslClientSessionStateMax(o["ssl-client-session-state-max"], d, "ssl_client_session_state_max", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-session-state-max"]) {
			return fmt.Errorf("Error reading ssl_client_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_client_rekey_count", flattenFirewallVip6SslClientRekeyCount(o["ssl-client-rekey-count"], d, "ssl_client_rekey_count", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-client-rekey-count"]) {
			return fmt.Errorf("Error reading ssl_client_rekey_count: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_type", flattenFirewallVip6SslServerSessionStateType(o["ssl-server-session-state-type"], d, "ssl_server_session_state_type", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-type"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_type: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_timeout", flattenFirewallVip6SslServerSessionStateTimeout(o["ssl-server-session-state-timeout"], d, "ssl_server_session_state_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-timeout"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_server_session_state_max", flattenFirewallVip6SslServerSessionStateMax(o["ssl-server-session-state-max"], d, "ssl_server_session_state_max", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-session-state-max"]) {
			return fmt.Errorf("Error reading ssl_server_session_state_max: %v", err)
		}
	}

	if err = d.Set("ssl_http_location_conversion", flattenFirewallVip6SslHttpLocationConversion(o["ssl-http-location-conversion"], d, "ssl_http_location_conversion", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-http-location-conversion"]) {
			return fmt.Errorf("Error reading ssl_http_location_conversion: %v", err)
		}
	}

	if err = d.Set("ssl_http_match_host", flattenFirewallVip6SslHttpMatchHost(o["ssl-http-match-host"], d, "ssl_http_match_host", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-http-match-host"]) {
			return fmt.Errorf("Error reading ssl_http_match_host: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp", flattenFirewallVip6SslHpkp(o["ssl-hpkp"], d, "ssl_hpkp", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp"]) {
			return fmt.Errorf("Error reading ssl_hpkp: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_primary", flattenFirewallVip6SslHpkpPrimary(o["ssl-hpkp-primary"], d, "ssl_hpkp_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-primary"]) {
			return fmt.Errorf("Error reading ssl_hpkp_primary: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_backup", flattenFirewallVip6SslHpkpBackup(o["ssl-hpkp-backup"], d, "ssl_hpkp_backup", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-backup"]) {
			return fmt.Errorf("Error reading ssl_hpkp_backup: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_age", flattenFirewallVip6SslHpkpAge(o["ssl-hpkp-age"], d, "ssl_hpkp_age", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-age"]) {
			return fmt.Errorf("Error reading ssl_hpkp_age: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_report_uri", flattenFirewallVip6SslHpkpReportUri(o["ssl-hpkp-report-uri"], d, "ssl_hpkp_report_uri", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-report-uri"]) {
			return fmt.Errorf("Error reading ssl_hpkp_report_uri: %v", err)
		}
	}

	if err = d.Set("ssl_hpkp_include_subdomains", flattenFirewallVip6SslHpkpIncludeSubdomains(o["ssl-hpkp-include-subdomains"], d, "ssl_hpkp_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hpkp-include-subdomains"]) {
			return fmt.Errorf("Error reading ssl_hpkp_include_subdomains: %v", err)
		}
	}

	if err = d.Set("ssl_hsts", flattenFirewallVip6SslHsts(o["ssl-hsts"], d, "ssl_hsts", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts"]) {
			return fmt.Errorf("Error reading ssl_hsts: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_age", flattenFirewallVip6SslHstsAge(o["ssl-hsts-age"], d, "ssl_hsts_age", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts-age"]) {
			return fmt.Errorf("Error reading ssl_hsts_age: %v", err)
		}
	}

	if err = d.Set("ssl_hsts_include_subdomains", flattenFirewallVip6SslHstsIncludeSubdomains(o["ssl-hsts-include-subdomains"], d, "ssl_hsts_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-hsts-include-subdomains"]) {
			return fmt.Errorf("Error reading ssl_hsts_include_subdomains: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("monitor", flattenFirewallVip6Monitor(o["monitor"], d, "monitor", sv)); err != nil {
			if !fortiAPIPatch(o["monitor"]) {
				return fmt.Errorf("Error reading monitor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitor"); ok {
			if err = d.Set("monitor", flattenFirewallVip6Monitor(o["monitor"], d, "monitor", sv)); err != nil {
				if !fortiAPIPatch(o["monitor"]) {
					return fmt.Errorf("Error reading monitor: %v", err)
				}
			}
		}
	}

	if err = d.Set("max_embryonic_connections", flattenFirewallVip6MaxEmbryonicConnections(o["max-embryonic-connections"], d, "max_embryonic_connections", sv)); err != nil {
		if !fortiAPIPatch(o["max-embryonic-connections"]) {
			return fmt.Errorf("Error reading max_embryonic_connections: %v", err)
		}
	}

	if err = d.Set("embedded_ipv4_address", flattenFirewallVip6EmbeddedIpv4Address(o["embedded-ipv4-address"], d, "embedded_ipv4_address", sv)); err != nil {
		if !fortiAPIPatch(o["embedded-ipv4-address"]) {
			return fmt.Errorf("Error reading embedded_ipv4_address: %v", err)
		}
	}

	if err = d.Set("ipv4_mappedip", flattenFirewallVip6Ipv4Mappedip(o["ipv4-mappedip"], d, "ipv4_mappedip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-mappedip"]) {
			return fmt.Errorf("Error reading ipv4_mappedip: %v", err)
		}
	}

	if err = d.Set("ipv4_mappedport", flattenFirewallVip6Ipv4Mappedport(o["ipv4-mappedport"], d, "ipv4_mappedport", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-mappedport"]) {
			return fmt.Errorf("Error reading ipv4_mappedport: %v", err)
		}
	}

	return nil
}

func flattenFirewallVip6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallVip6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Uuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Comment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Type(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SrcFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["range"], _ = expandFirewallVip6SrcFilterRange(d, i["range"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip6SrcFilterRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Extip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Mappedip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6NatSourceVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6ArpReply(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Portforward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Protocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Extport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Mappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Color(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6LdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6ServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Persistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Nat66(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Nat64(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6AddNat64Route(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Realservers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallVip6RealserversId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandFirewallVip6RealserversIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandFirewallVip6RealserversPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallVip6RealserversStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandFirewallVip6RealserversWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holddown_interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["holddown-interval"], _ = expandFirewallVip6RealserversHolddownInterval(d, i["holddown_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "healthcheck"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["healthcheck"], _ = expandFirewallVip6RealserversHealthcheck(d, i["healthcheck"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-host"], _ = expandFirewallVip6RealserversHttpHost(d, i["http_host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_connections"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["max-connections"], _ = expandFirewallVip6RealserversMaxConnections(d, i["max_connections"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok {

			bstring := false
			t, _ := expandFirewallVip6RealserversMonitor(d, i["monitor"], pre_append, sv)
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

			tmp["client-ip"], _ = expandFirewallVip6RealserversClientIp(d, i["client_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip6RealserversId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHolddownInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHealthcheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversHttpHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversMaxConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6RealserversClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieDomainFromHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookiePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpCookieShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpsCookieSecure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpMultiplex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpIpHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6HttpIpHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6OutlookWebAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6WeblogicServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6WebsphereServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslDhBits(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["priority"], _ = expandFirewallVip6SslCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallVip6SslCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallVip6SslCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip6SslCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerCipherSuites(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["priority"], _ = expandFirewallVip6SslServerCipherSuitesPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cipher"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cipher"], _ = expandFirewallVip6SslServerCipherSuitesCipher(d, i["cipher"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "versions"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["versions"], _ = expandFirewallVip6SslServerCipherSuitesVersions(d, i["versions"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip6SslServerCipherSuitesPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerCipherSuitesCipher(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerCipherSuitesVersions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslPfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslAcceptFfdheGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientRenegotiation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientSessionStateType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientSessionStateMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslClientRekeyCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerSessionStateType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerSessionStateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslServerSessionStateMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHttpLocationConversion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHttpMatchHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpBackup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpReportUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHpkpIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHsts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHstsAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6SslHstsIncludeSubdomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Monitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallVip6MonitorName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallVip6MonitorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6MaxEmbryonicConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6EmbeddedIpv4Address(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Ipv4Mappedip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVip6Ipv4Mappedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallVip6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallVip6Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandFirewallVip6Id(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {

		t, err := expandFirewallVip6Uuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallVip6Comment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallVip6Type(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("src_filter"); ok {

		t, err := expandFirewallVip6SrcFilter(d, v, "src_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-filter"] = t
		}
	}

	if v, ok := d.GetOk("extip"); ok {

		t, err := expandFirewallVip6Extip(d, v, "extip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extip"] = t
		}
	}

	if v, ok := d.GetOk("mappedip"); ok {

		t, err := expandFirewallVip6Mappedip(d, v, "mappedip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedip"] = t
		}
	}

	if v, ok := d.GetOk("nat_source_vip"); ok {

		t, err := expandFirewallVip6NatSourceVip(d, v, "nat_source_vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat-source-vip"] = t
		}
	}

	if v, ok := d.GetOk("arp_reply"); ok {

		t, err := expandFirewallVip6ArpReply(d, v, "arp_reply", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-reply"] = t
		}
	}

	if v, ok := d.GetOk("portforward"); ok {

		t, err := expandFirewallVip6Portforward(d, v, "portforward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portforward"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {

		t, err := expandFirewallVip6Protocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("extport"); ok {

		t, err := expandFirewallVip6Extport(d, v, "extport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extport"] = t
		}
	}

	if v, ok := d.GetOk("mappedport"); ok {

		t, err := expandFirewallVip6Mappedport(d, v, "mappedport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mappedport"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandFirewallVip6Color(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok {

		t, err := expandFirewallVip6LdbMethod(d, v, "ldb_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {

		t, err := expandFirewallVip6ServerType(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("http_redirect"); ok {

		t, err := expandFirewallVip6HttpRedirect(d, v, "http_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-redirect"] = t
		}
	}

	if v, ok := d.GetOk("persistence"); ok {

		t, err := expandFirewallVip6Persistence(d, v, "persistence", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}

	if v, ok := d.GetOk("nat66"); ok {

		t, err := expandFirewallVip6Nat66(d, v, "nat66", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat66"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok {

		t, err := expandFirewallVip6Nat64(d, v, "nat64", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("add_nat64_route"); ok {

		t, err := expandFirewallVip6AddNat64Route(d, v, "add_nat64_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat64-route"] = t
		}
	}

	if v, ok := d.GetOk("realservers"); ok {

		t, err := expandFirewallVip6Realservers(d, v, "realservers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realservers"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain_from_host"); ok {

		t, err := expandFirewallVip6HttpCookieDomainFromHost(d, v, "http_cookie_domain_from_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain-from-host"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_domain"); ok {

		t, err := expandFirewallVip6HttpCookieDomain(d, v, "http_cookie_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_path"); ok {

		t, err := expandFirewallVip6HttpCookiePath(d, v, "http_cookie_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-path"] = t
		}
	}

	if v, ok := d.GetOkExists("http_cookie_generation"); ok {

		t, err := expandFirewallVip6HttpCookieGeneration(d, v, "http_cookie_generation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-generation"] = t
		}
	}

	if v, ok := d.GetOkExists("http_cookie_age"); ok {

		t, err := expandFirewallVip6HttpCookieAge(d, v, "http_cookie_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("http_cookie_share"); ok {

		t, err := expandFirewallVip6HttpCookieShare(d, v, "http_cookie_share", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-cookie-share"] = t
		}
	}

	if v, ok := d.GetOk("https_cookie_secure"); ok {

		t, err := expandFirewallVip6HttpsCookieSecure(d, v, "https_cookie_secure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-cookie-secure"] = t
		}
	}

	if v, ok := d.GetOk("http_multiplex"); ok {

		t, err := expandFirewallVip6HttpMultiplex(d, v, "http_multiplex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-multiplex"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header"); ok {

		t, err := expandFirewallVip6HttpIpHeader(d, v, "http_ip_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header"] = t
		}
	}

	if v, ok := d.GetOk("http_ip_header_name"); ok {

		t, err := expandFirewallVip6HttpIpHeaderName(d, v, "http_ip_header_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-ip-header-name"] = t
		}
	}

	if v, ok := d.GetOk("outlook_web_access"); ok {

		t, err := expandFirewallVip6OutlookWebAccess(d, v, "outlook_web_access", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outlook-web-access"] = t
		}
	}

	if v, ok := d.GetOk("weblogic_server"); ok {

		t, err := expandFirewallVip6WeblogicServer(d, v, "weblogic_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weblogic-server"] = t
		}
	}

	if v, ok := d.GetOk("websphere_server"); ok {

		t, err := expandFirewallVip6WebsphereServer(d, v, "websphere_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["websphere-server"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mode"); ok {

		t, err := expandFirewallVip6SslMode(d, v, "ssl_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mode"] = t
		}
	}

	if v, ok := d.GetOk("ssl_certificate"); ok {

		t, err := expandFirewallVip6SslCertificate(d, v, "ssl_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-certificate"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok {

		t, err := expandFirewallVip6SslDhBits(d, v, "ssl_dh_bits", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_algorithm"); ok {

		t, err := expandFirewallVip6SslAlgorithm(d, v, "ssl_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_cipher_suites"); ok {

		t, err := expandFirewallVip6SslCipherSuites(d, v, "ssl_cipher_suites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_algorithm"); ok {

		t, err := expandFirewallVip6SslServerAlgorithm(d, v, "ssl_server_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_cipher_suites"); ok {

		t, err := expandFirewallVip6SslServerCipherSuites(d, v, "ssl_server_cipher_suites", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-cipher-suites"] = t
		}
	}

	if v, ok := d.GetOk("ssl_pfs"); ok {

		t, err := expandFirewallVip6SslPfs(d, v, "ssl_pfs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-pfs"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok {

		t, err := expandFirewallVip6SslMinVersion(d, v, "ssl_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {

		t, err := expandFirewallVip6SslMaxVersion(d, v, "ssl_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_min_version"); ok {

		t, err := expandFirewallVip6SslServerMinVersion(d, v, "ssl_server_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-min-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_max_version"); ok {

		t, err := expandFirewallVip6SslServerMaxVersion(d, v, "ssl_server_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_accept_ffdhe_groups"); ok {

		t, err := expandFirewallVip6SslAcceptFfdheGroups(d, v, "ssl_accept_ffdhe_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-accept-ffdhe-groups"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok {

		t, err := expandFirewallVip6SslSendEmptyFrags(d, v, "ssl_send_empty_frags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_fallback"); ok {

		t, err := expandFirewallVip6SslClientFallback(d, v, "ssl_client_fallback", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-fallback"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_renegotiation"); ok {

		t, err := expandFirewallVip6SslClientRenegotiation(d, v, "ssl_client_renegotiation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-renegotiation"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_type"); ok {

		t, err := expandFirewallVip6SslClientSessionStateType(d, v, "ssl_client_session_state_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_timeout"); ok {

		t, err := expandFirewallVip6SslClientSessionStateTimeout(d, v, "ssl_client_session_state_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_session_state_max"); ok {

		t, err := expandFirewallVip6SslClientSessionStateMax(d, v, "ssl_client_session_state_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_client_rekey_count"); ok {

		t, err := expandFirewallVip6SslClientRekeyCount(d, v, "ssl_client_rekey_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-client-rekey-count"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_type"); ok {

		t, err := expandFirewallVip6SslServerSessionStateType(d, v, "ssl_server_session_state_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-type"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_timeout"); ok {

		t, err := expandFirewallVip6SslServerSessionStateTimeout(d, v, "ssl_server_session_state_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_server_session_state_max"); ok {

		t, err := expandFirewallVip6SslServerSessionStateMax(d, v, "ssl_server_session_state_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-session-state-max"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_location_conversion"); ok {

		t, err := expandFirewallVip6SslHttpLocationConversion(d, v, "ssl_http_location_conversion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-location-conversion"] = t
		}
	}

	if v, ok := d.GetOk("ssl_http_match_host"); ok {

		t, err := expandFirewallVip6SslHttpMatchHost(d, v, "ssl_http_match_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-http-match-host"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp"); ok {

		t, err := expandFirewallVip6SslHpkp(d, v, "ssl_hpkp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_primary"); ok {

		t, err := expandFirewallVip6SslHpkpPrimary(d, v, "ssl_hpkp_primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-primary"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_backup"); ok {

		t, err := expandFirewallVip6SslHpkpBackup(d, v, "ssl_hpkp_backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-backup"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_age"); ok {

		t, err := expandFirewallVip6SslHpkpAge(d, v, "ssl_hpkp_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_report_uri"); ok {

		t, err := expandFirewallVip6SslHpkpReportUri(d, v, "ssl_hpkp_report_uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-report-uri"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hpkp_include_subdomains"); ok {

		t, err := expandFirewallVip6SslHpkpIncludeSubdomains(d, v, "ssl_hpkp_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hpkp-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts"); ok {

		t, err := expandFirewallVip6SslHsts(d, v, "ssl_hsts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_age"); ok {

		t, err := expandFirewallVip6SslHstsAge(d, v, "ssl_hsts_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-age"] = t
		}
	}

	if v, ok := d.GetOk("ssl_hsts_include_subdomains"); ok {

		t, err := expandFirewallVip6SslHstsIncludeSubdomains(d, v, "ssl_hsts_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-hsts-include-subdomains"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok {

		t, err := expandFirewallVip6Monitor(d, v, "monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOkExists("max_embryonic_connections"); ok {

		t, err := expandFirewallVip6MaxEmbryonicConnections(d, v, "max_embryonic_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-embryonic-connections"] = t
		}
	}

	if v, ok := d.GetOk("embedded_ipv4_address"); ok {

		t, err := expandFirewallVip6EmbeddedIpv4Address(d, v, "embedded_ipv4_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["embedded-ipv4-address"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_mappedip"); ok {

		t, err := expandFirewallVip6Ipv4Mappedip(d, v, "ipv4_mappedip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-mappedip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_mappedport"); ok {

		t, err := expandFirewallVip6Ipv4Mappedport(d, v, "ipv4_mappedport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-mappedport"] = t
		}
	}

	return &obj, nil
}
