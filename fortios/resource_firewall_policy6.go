// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallPolicy6Create,
		Read:   resourceFirewallPolicy6Read,
		Update: resourceFirewallPolicy6Update,
		Delete: resourceFirewallPolicy6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
			"dstintf": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
				Required: true,
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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_cos_fwd": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"vlan_cos_rev": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"service": &schema.Schema{
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
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webcache_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"av_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dlp_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"voip_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"icap_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"cifs_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"waf_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webproxy_forward_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"traffic_shaper": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"app_group": &schema.Schema{
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
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poolname": &schema.Schema{
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
			"session_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(300, 604800),
				Optional:     true,
				Computed:     true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natinbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"natoutbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpntunnel": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_sender": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"label": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"global_label": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"replacemsg_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"devices": &schema.Schema{
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
			"timeout_send_rst": &schema.Schema{
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
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_mirror_intf": &schema.Schema{
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
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
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

func resourceFirewallPolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallPolicy6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicy6 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallPolicy6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicy6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallPolicy6")
	}

	return resourceFirewallPolicy6Read(d, m)
}

func resourceFirewallPolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallPolicy6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallPolicy6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallPolicy6")
	}

	return resourceFirewallPolicy6Read(d, m)
}

func resourceFirewallPolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallPolicy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallPolicy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicy6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Uuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Srcintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6SrcintfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6SrcintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Dstintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6DstintfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6DstintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6SrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6SrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6DstaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6DstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Action(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6FirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6VlanCosFwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6VlanCosRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Schedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Service(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6ServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6ServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Tos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TosNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6AntiReplay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6UtmStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6InspectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Webcache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6WebcacheHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6HttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6WebproxyProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6ProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6ProfileGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6AvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6WebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DnsfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6EmailfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SpamfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DlpSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6IpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6ApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6VoipProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6IcapProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6CifsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6WafProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SshFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6ProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SslSshProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6LogtrafficStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6AutoAsicOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6WebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TrafficShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6PerIpShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Application(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallPolicy6ApplicationId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicy6ApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6AppCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallPolicy6AppCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicy6AppCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6UrlCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallPolicy6UrlCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicy6UrlCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6AppGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6AppGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6AppGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Nat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Fixedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Ippool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Poolname(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6PoolnameName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6PoolnameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SessionTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Inbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Outbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Natinbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Natoutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SendDenyPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Vpntunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DiffservForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DiffservReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DiffservcodeForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DiffservcodeRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TcpMssSender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TcpMssReceiver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Label(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6GlobalLabel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Rsso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6CustomLogFields(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_id"
		if _, ok := i["field-id"]; ok {

			tmp["field_id"] = flattenFirewallPolicy6CustomLogFieldsFieldId(i["field-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "field_id", d)
	return result
}

func flattenFirewallPolicy6CustomLogFieldsFieldId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6ReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SrcaddrNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DstaddrNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6ServiceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Groups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6GroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6GroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Users(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6UsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6UsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Devices(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6DevicesName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6DevicesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6TimeoutSendRst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6DecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SslMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6SslMirrorIntf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6SslMirrorIntfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6SslMirrorIntfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6Dsri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6VlanFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicy6FssoGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallPolicy6FssoGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicy6FssoGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallPolicy6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("policyid", flattenFirewallPolicy6Policyid(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallPolicy6Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallPolicy6Uuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcintf", flattenFirewallPolicy6Srcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
			if !fortiAPIPatch(o["srcintf"]) {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf"); ok {
			if err = d.Set("srcintf", flattenFirewallPolicy6Srcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
				if !fortiAPIPatch(o["srcintf"]) {
					return fmt.Errorf("Error reading srcintf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstintf", flattenFirewallPolicy6Dstintf(o["dstintf"], d, "dstintf", sv)); err != nil {
			if !fortiAPIPatch(o["dstintf"]) {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstintf"); ok {
			if err = d.Set("dstintf", flattenFirewallPolicy6Dstintf(o["dstintf"], d, "dstintf", sv)); err != nil {
				if !fortiAPIPatch(o["dstintf"]) {
					return fmt.Errorf("Error reading dstintf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr", flattenFirewallPolicy6Srcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallPolicy6Srcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr", flattenFirewallPolicy6Dstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenFirewallPolicy6Dstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if err = d.Set("action", flattenFirewallPolicy6Action(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenFirewallPolicy6FirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-session-dirty"]) {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallPolicy6Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", flattenFirewallPolicy6VlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-cos-fwd"]) {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", flattenFirewallPolicy6VlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-cos-rev"]) {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallPolicy6Schedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallPolicy6Service(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallPolicy6Service(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("tos", flattenFirewallPolicy6Tos(o["tos"], d, "tos", sv)); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenFirewallPolicy6TosMask(o["tos-mask"], d, "tos_mask", sv)); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenFirewallPolicy6TosNegate(o["tos-negate"], d, "tos_negate", sv)); err != nil {
		if !fortiAPIPatch(o["tos-negate"]) {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenFirewallPolicy6TcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-session-without-syn"]) {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenFirewallPolicy6AntiReplay(o["anti-replay"], d, "anti_replay", sv)); err != nil {
		if !fortiAPIPatch(o["anti-replay"]) {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallPolicy6UtmStatus(o["utm-status"], d, "utm_status", sv)); err != nil {
		if !fortiAPIPatch(o["utm-status"]) {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenFirewallPolicy6InspectionMode(o["inspection-mode"], d, "inspection_mode", sv)); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("webcache", flattenFirewallPolicy6Webcache(o["webcache"], d, "webcache", sv)); err != nil {
		if !fortiAPIPatch(o["webcache"]) {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenFirewallPolicy6WebcacheHttps(o["webcache-https"], d, "webcache_https", sv)); err != nil {
		if !fortiAPIPatch(o["webcache-https"]) {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenFirewallPolicy6HttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["http-policy-redirect"]) {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenFirewallPolicy6SshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-policy-redirect"]) {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenFirewallPolicy6WebproxyProfile(o["webproxy-profile"], d, "webproxy_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-profile"]) {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallPolicy6ProfileType(o["profile-type"], d, "profile_type", sv)); err != nil {
		if !fortiAPIPatch(o["profile-type"]) {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallPolicy6ProfileGroup(o["profile-group"], d, "profile_group", sv)); err != nil {
		if !fortiAPIPatch(o["profile-group"]) {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallPolicy6AvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallPolicy6WebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallPolicy6DnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallPolicy6EmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallPolicy6SpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallPolicy6DlpSensor(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallPolicy6IpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallPolicy6ApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallPolicy6VoipProfile(o["voip-profile"], d, "voip_profile", sv)); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallPolicy6IcapProfile(o["icap-profile"], d, "icap_profile", sv)); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallPolicy6CifsProfile(o["cifs-profile"], d, "cifs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["cifs-profile"]) {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallPolicy6WafProfile(o["waf-profile"], d, "waf_profile", sv)); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallPolicy6SshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallPolicy6ProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options", sv)); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallPolicy6SslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallPolicy6Logtraffic(o["logtraffic"], d, "logtraffic", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallPolicy6LogtrafficStart(o["logtraffic-start"], d, "logtraffic_start", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic-start"]) {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenFirewallPolicy6AutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload", sv)); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenFirewallPolicy6WebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-forward-server"]) {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallPolicy6TrafficShaper(o["traffic-shaper"], d, "traffic_shaper", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-shaper"]) {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallPolicy6TrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-shaper-reverse"]) {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallPolicy6PerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper", sv)); err != nil {
		if !fortiAPIPatch(o["per-ip-shaper"]) {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("application", flattenFirewallPolicy6Application(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallPolicy6Application(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_category", flattenFirewallPolicy6AppCategory(o["app-category"], d, "app_category", sv)); err != nil {
			if !fortiAPIPatch(o["app-category"]) {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_category"); ok {
			if err = d.Set("app_category", flattenFirewallPolicy6AppCategory(o["app-category"], d, "app_category", sv)); err != nil {
				if !fortiAPIPatch(o["app-category"]) {
					return fmt.Errorf("Error reading app_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_category", flattenFirewallPolicy6UrlCategory(o["url-category"], d, "url_category", sv)); err != nil {
			if !fortiAPIPatch(o["url-category"]) {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_category"); ok {
			if err = d.Set("url_category", flattenFirewallPolicy6UrlCategory(o["url-category"], d, "url_category", sv)); err != nil {
				if !fortiAPIPatch(o["url-category"]) {
					return fmt.Errorf("Error reading url_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_group", flattenFirewallPolicy6AppGroup(o["app-group"], d, "app_group", sv)); err != nil {
			if !fortiAPIPatch(o["app-group"]) {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_group"); ok {
			if err = d.Set("app_group", flattenFirewallPolicy6AppGroup(o["app-group"], d, "app_group", sv)); err != nil {
				if !fortiAPIPatch(o["app-group"]) {
					return fmt.Errorf("Error reading app_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("nat", flattenFirewallPolicy6Nat(o["nat"], d, "nat", sv)); err != nil {
		if !fortiAPIPatch(o["nat"]) {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenFirewallPolicy6Fixedport(o["fixedport"], d, "fixedport", sv)); err != nil {
		if !fortiAPIPatch(o["fixedport"]) {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("ippool", flattenFirewallPolicy6Ippool(o["ippool"], d, "ippool", sv)); err != nil {
		if !fortiAPIPatch(o["ippool"]) {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("poolname", flattenFirewallPolicy6Poolname(o["poolname"], d, "poolname", sv)); err != nil {
			if !fortiAPIPatch(o["poolname"]) {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("poolname"); ok {
			if err = d.Set("poolname", flattenFirewallPolicy6Poolname(o["poolname"], d, "poolname", sv)); err != nil {
				if !fortiAPIPatch(o["poolname"]) {
					return fmt.Errorf("Error reading poolname: %v", err)
				}
			}
		}
	}

	{
		v := flattenFirewallPolicy6SessionTtl(o["session-ttl"], d, "session_ttl", sv)
		if i2ss2arrFortiAPIUpgrade(sv, "6.2.4") == true {
			if vx, ok := v.(string); ok {
				vxx, err := strconv.Atoi(vx)
				if err == nil {
					v = vxx
				}
			}
		}

		if err = d.Set("session_ttl", v); err != nil {
			if !fortiAPIPatch(o["session-ttl"]) {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		}
	}

	if err = d.Set("inbound", flattenFirewallPolicy6Inbound(o["inbound"], d, "inbound", sv)); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("outbound", flattenFirewallPolicy6Outbound(o["outbound"], d, "outbound", sv)); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("natinbound", flattenFirewallPolicy6Natinbound(o["natinbound"], d, "natinbound", sv)); err != nil {
		if !fortiAPIPatch(o["natinbound"]) {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natoutbound", flattenFirewallPolicy6Natoutbound(o["natoutbound"], d, "natoutbound", sv)); err != nil {
		if !fortiAPIPatch(o["natoutbound"]) {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenFirewallPolicy6SendDenyPacket(o["send-deny-packet"], d, "send_deny_packet", sv)); err != nil {
		if !fortiAPIPatch(o["send-deny-packet"]) {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("vpntunnel", flattenFirewallPolicy6Vpntunnel(o["vpntunnel"], d, "vpntunnel", sv)); err != nil {
		if !fortiAPIPatch(o["vpntunnel"]) {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallPolicy6DiffservForward(o["diffserv-forward"], d, "diffserv_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallPolicy6DiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallPolicy6DiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallPolicy6DiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenFirewallPolicy6TcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-sender"]) {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenFirewallPolicy6TcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-receiver"]) {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallPolicy6Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallPolicy6Label(o["label"], d, "label", sv)); err != nil {
		if !fortiAPIPatch(o["label"]) {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("global_label", flattenFirewallPolicy6GlobalLabel(o["global-label"], d, "global_label", sv)); err != nil {
		if !fortiAPIPatch(o["global-label"]) {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("rsso", flattenFirewallPolicy6Rsso(o["rsso"], d, "rsso", sv)); err != nil {
		if !fortiAPIPatch(o["rsso"]) {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_log_fields", flattenFirewallPolicy6CustomLogFields(o["custom-log-fields"], d, "custom_log_fields", sv)); err != nil {
			if !fortiAPIPatch(o["custom-log-fields"]) {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_log_fields"); ok {
			if err = d.Set("custom_log_fields", flattenFirewallPolicy6CustomLogFields(o["custom-log-fields"], d, "custom_log_fields", sv)); err != nil {
				if !fortiAPIPatch(o["custom-log-fields"]) {
					return fmt.Errorf("Error reading custom_log_fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_override_group", flattenFirewallPolicy6ReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallPolicy6SrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["srcaddr-negate"]) {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallPolicy6DstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr-negate"]) {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallPolicy6ServiceNegate(o["service-negate"], d, "service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["service-negate"]) {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenFirewallPolicy6Groups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenFirewallPolicy6Groups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenFirewallPolicy6Users(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenFirewallPolicy6Users(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("devices", flattenFirewallPolicy6Devices(o["devices"], d, "devices", sv)); err != nil {
			if !fortiAPIPatch(o["devices"]) {
				return fmt.Errorf("Error reading devices: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("devices"); ok {
			if err = d.Set("devices", flattenFirewallPolicy6Devices(o["devices"], d, "devices", sv)); err != nil {
				if !fortiAPIPatch(o["devices"]) {
					return fmt.Errorf("Error reading devices: %v", err)
				}
			}
		}
	}

	if err = d.Set("timeout_send_rst", flattenFirewallPolicy6TimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-send-rst"]) {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallPolicy6DecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", flattenFirewallPolicy6SslMirror(o["ssl-mirror"], d, "ssl_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mirror"]) {
			return fmt.Errorf("Error reading ssl_mirror: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_mirror_intf", flattenFirewallPolicy6SslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-mirror-intf"]) {
				return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_mirror_intf"); ok {
			if err = d.Set("ssl_mirror_intf", flattenFirewallPolicy6SslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-mirror-intf"]) {
					return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
				}
			}
		}
	}

	if err = d.Set("dsri", flattenFirewallPolicy6Dsri(o["dsri"], d, "dsri", sv)); err != nil {
		if !fortiAPIPatch(o["dsri"]) {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenFirewallPolicy6VlanFilter(o["vlan-filter"], d, "vlan_filter", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-filter"]) {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fsso_groups", flattenFirewallPolicy6FssoGroups(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
			if !fortiAPIPatch(o["fsso-groups"]) {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fsso_groups"); ok {
			if err = d.Set("fsso_groups", flattenFirewallPolicy6FssoGroups(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
				if !fortiAPIPatch(o["fsso-groups"]) {
					return fmt.Errorf("Error reading fsso_groups: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallPolicy6Policyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Uuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Srcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6SrcintfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6SrcintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Dstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6DstintfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6DstintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6SrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6SrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6DstaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6DstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Action(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6FirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6VlanCosFwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6VlanCosRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Schedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Service(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6ServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6ServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Tos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TosNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6AntiReplay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6UtmStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6InspectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Webcache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6WebcacheHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6HttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6WebproxyProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6ProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6ProfileGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6AvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6WebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DnsfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6EmailfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SpamfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DlpSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6IpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6ApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6VoipProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6IcapProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6CifsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6WafProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SshFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6ProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SslSshProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6LogtrafficStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6AutoAsicOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6WebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TrafficShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6PerIpShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Application(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallPolicy6ApplicationId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6ApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6AppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallPolicy6AppCategoryId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6AppCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6UrlCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallPolicy6UrlCategoryId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6UrlCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6AppGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6AppGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6AppGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Nat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Fixedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Ippool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Poolname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6PoolnameName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6PoolnameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SessionTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Inbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Outbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Natinbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Natoutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SendDenyPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Vpntunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DiffservForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DiffservReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DiffservcodeForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DiffservcodeRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TcpMssSender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TcpMssReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Label(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6GlobalLabel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Rsso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6CustomLogFields(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["field-id"], _ = expandFirewallPolicy6CustomLogFieldsFieldId(d, i["field_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6CustomLogFieldsFieldId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6ReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SrcaddrNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DstaddrNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6ServiceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Groups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6GroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6GroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Users(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6UsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6UsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Devices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6DevicesName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6DevicesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6TimeoutSendRst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6DecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SslMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6SslMirrorIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6SslMirrorIntfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6SslMirrorIntfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6Dsri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6VlanFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicy6FssoGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallPolicy6FssoGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicy6FssoGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallPolicy6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("policyid"); ok {

		t, err := expandFirewallPolicy6Policyid(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallPolicy6Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {

		t, err := expandFirewallPolicy6Uuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {

		t, err := expandFirewallPolicy6Srcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {

		t, err := expandFirewallPolicy6Dstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {

		t, err := expandFirewallPolicy6Srcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {

		t, err := expandFirewallPolicy6Dstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandFirewallPolicy6Action(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok {

		t, err := expandFirewallPolicy6FirewallSessionDirty(d, v, "firewall_session_dirty", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandFirewallPolicy6Status(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOkExists("vlan_cos_fwd"); ok {

		t, err := expandFirewallPolicy6VlanCosFwd(d, v, "vlan_cos_fwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-fwd"] = t
		}
	}

	if v, ok := d.GetOkExists("vlan_cos_rev"); ok {

		t, err := expandFirewallPolicy6VlanCosRev(d, v, "vlan_cos_rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-rev"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {

		t, err := expandFirewallPolicy6Schedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandFirewallPolicy6Service(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {

		t, err := expandFirewallPolicy6Tos(d, v, "tos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {

		t, err := expandFirewallPolicy6TosMask(d, v, "tos_mask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok {

		t, err := expandFirewallPolicy6TosNegate(d, v, "tos_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok {

		t, err := expandFirewallPolicy6TcpSessionWithoutSyn(d, v, "tcp_session_without_syn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok {

		t, err := expandFirewallPolicy6AntiReplay(d, v, "anti_replay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok {

		t, err := expandFirewallPolicy6UtmStatus(d, v, "utm_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {

		t, err := expandFirewallPolicy6InspectionMode(d, v, "inspection_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok {

		t, err := expandFirewallPolicy6Webcache(d, v, "webcache", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok {

		t, err := expandFirewallPolicy6WebcacheHttps(d, v, "webcache_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok {

		t, err := expandFirewallPolicy6HttpPolicyRedirect(d, v, "http_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok {

		t, err := expandFirewallPolicy6SshPolicyRedirect(d, v, "ssh_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok {

		t, err := expandFirewallPolicy6WebproxyProfile(d, v, "webproxy_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok {

		t, err := expandFirewallPolicy6ProfileType(d, v, "profile_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok {

		t, err := expandFirewallPolicy6ProfileGroup(d, v, "profile_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {

		t, err := expandFirewallPolicy6AvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {

		t, err := expandFirewallPolicy6WebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {

		t, err := expandFirewallPolicy6DnsfilterProfile(d, v, "dnsfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {

		t, err := expandFirewallPolicy6EmailfilterProfile(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {

		t, err := expandFirewallPolicy6SpamfilterProfile(d, v, "spamfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {

		t, err := expandFirewallPolicy6DlpSensor(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {

		t, err := expandFirewallPolicy6IpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {

		t, err := expandFirewallPolicy6ApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok {

		t, err := expandFirewallPolicy6VoipProfile(d, v, "voip_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok {

		t, err := expandFirewallPolicy6IcapProfile(d, v, "icap_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok {

		t, err := expandFirewallPolicy6CifsProfile(d, v, "cifs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok {

		t, err := expandFirewallPolicy6WafProfile(d, v, "waf_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {

		t, err := expandFirewallPolicy6SshFilterProfile(d, v, "ssh_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {

		t, err := expandFirewallPolicy6ProfileProtocolOptions(d, v, "profile_protocol_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {

		t, err := expandFirewallPolicy6SslSshProfile(d, v, "ssl_ssh_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {

		t, err := expandFirewallPolicy6Logtraffic(d, v, "logtraffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {

		t, err := expandFirewallPolicy6LogtrafficStart(d, v, "logtraffic_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok {

		t, err := expandFirewallPolicy6AutoAsicOffload(d, v, "auto_asic_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok {

		t, err := expandFirewallPolicy6WebproxyForwardServer(d, v, "webproxy_forward_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok {

		t, err := expandFirewallPolicy6TrafficShaper(d, v, "traffic_shaper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok {

		t, err := expandFirewallPolicy6TrafficShaperReverse(d, v, "traffic_shaper_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok {

		t, err := expandFirewallPolicy6PerIpShaper(d, v, "per_ip_shaper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {

		t, err := expandFirewallPolicy6Application(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok {

		t, err := expandFirewallPolicy6AppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok {

		t, err := expandFirewallPolicy6UrlCategory(d, v, "url_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok {

		t, err := expandFirewallPolicy6AppGroup(d, v, "app_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok {

		t, err := expandFirewallPolicy6Nat(d, v, "nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok {

		t, err := expandFirewallPolicy6Fixedport(d, v, "fixedport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok {

		t, err := expandFirewallPolicy6Ippool(d, v, "ippool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok {

		t, err := expandFirewallPolicy6Poolname(d, v, "poolname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOkExists("session_ttl"); ok {

		t, err := expandFirewallPolicy6SessionTtl(d, v, "session_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			if i2ss2arrFortiAPIUpgrade(sv, "6.2.4") == true {
				obj["session-ttl"] = fmt.Sprintf("%v", t)
			} else {
				obj["session-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("inbound"); ok {

		t, err := expandFirewallPolicy6Inbound(d, v, "inbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok {

		t, err := expandFirewallPolicy6Outbound(d, v, "outbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("natinbound"); ok {

		t, err := expandFirewallPolicy6Natinbound(d, v, "natinbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natinbound"] = t
		}
	}

	if v, ok := d.GetOk("natoutbound"); ok {

		t, err := expandFirewallPolicy6Natoutbound(d, v, "natoutbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natoutbound"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok {

		t, err := expandFirewallPolicy6SendDenyPacket(d, v, "send_deny_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok {

		t, err := expandFirewallPolicy6Vpntunnel(d, v, "vpntunnel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok {

		t, err := expandFirewallPolicy6DiffservForward(d, v, "diffserv_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok {

		t, err := expandFirewallPolicy6DiffservReverse(d, v, "diffserv_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok {

		t, err := expandFirewallPolicy6DiffservcodeForward(d, v, "diffservcode_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok {

		t, err := expandFirewallPolicy6DiffservcodeRev(d, v, "diffservcode_rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_mss_sender"); ok {

		t, err := expandFirewallPolicy6TcpMssSender(d, v, "tcp_mss_sender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_mss_receiver"); ok {

		t, err := expandFirewallPolicy6TcpMssReceiver(d, v, "tcp_mss_receiver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallPolicy6Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok {

		t, err := expandFirewallPolicy6Label(d, v, "label", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok {

		t, err := expandFirewallPolicy6GlobalLabel(d, v, "global_label", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok {

		t, err := expandFirewallPolicy6Rsso(d, v, "rsso", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok {

		t, err := expandFirewallPolicy6CustomLogFields(d, v, "custom_log_fields", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok {

		t, err := expandFirewallPolicy6ReplacemsgOverrideGroup(d, v, "replacemsg_override_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok {

		t, err := expandFirewallPolicy6SrcaddrNegate(d, v, "srcaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok {

		t, err := expandFirewallPolicy6DstaddrNegate(d, v, "dstaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok {

		t, err := expandFirewallPolicy6ServiceNegate(d, v, "service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {

		t, err := expandFirewallPolicy6Groups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {

		t, err := expandFirewallPolicy6Users(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("devices"); ok {

		t, err := expandFirewallPolicy6Devices(d, v, "devices", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devices"] = t
		}
	}

	if v, ok := d.GetOk("timeout_send_rst"); ok {

		t, err := expandFirewallPolicy6TimeoutSendRst(d, v, "timeout_send_rst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-send-rst"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {

		t, err := expandFirewallPolicy6DecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror"); ok {

		t, err := expandFirewallPolicy6SslMirror(d, v, "ssl_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror_intf"); ok {

		t, err := expandFirewallPolicy6SslMirrorIntf(d, v, "ssl_mirror_intf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror-intf"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok {

		t, err := expandFirewallPolicy6Dsri(d, v, "dsri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok {

		t, err := expandFirewallPolicy6VlanFilter(d, v, "vlan_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok {

		t, err := expandFirewallPolicy6FssoGroups(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	return &obj, nil
}
