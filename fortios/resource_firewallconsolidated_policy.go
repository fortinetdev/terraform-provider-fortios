// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure consolidated IPv4/IPv6 policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallConsolidatedPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallConsolidatedPolicyCreate,
		Read:   resourceFirewallConsolidatedPolicyRead,
		Update: resourceFirewallConsolidatedPolicyUpdate,
		Delete: resourceFirewallConsolidatedPolicyDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
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
			"dstintf": &schema.Schema{
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
			"srcaddr4": &schema.Schema{
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
			"dstaddr4": &schema.Schema{
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
			"dstaddr6": &schema.Schema{
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
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
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
			"internet_service_id": &schema.Schema{
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
			"internet_service_group": &schema.Schema{
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
			"internet_service_custom": &schema.Schema{
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
			"internet_service_custom_group": &schema.Schema{
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
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
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
			"internet_service_src_id": &schema.Schema{
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
			"internet_service_src_group": &schema.Schema{
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
			"internet_service_src_custom": &schema.Schema{
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
			"internet_service_src_custom_group": &schema.Schema{
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
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
			"webproxy_forward_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"wanopt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanopt_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanopt_passive_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanopt_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"wanopt_peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
			"poolname4": &schema.Schema{
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
			"poolname6": &schema.Schema{
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
				ValidateFunc: intBetweenWithZero(300, 2764800),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"vpntunnel": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
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
			"captive_portal_exempt": &schema.Schema{
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

func resourceFirewallConsolidatedPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallConsolidatedPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallConsolidatedPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallConsolidatedPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallConsolidatedPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallConsolidatedPolicy")
	}

	return resourceFirewallConsolidatedPolicyRead(d, m)
}

func resourceFirewallConsolidatedPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallConsolidatedPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallConsolidatedPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallConsolidatedPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallConsolidatedPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallConsolidatedPolicy")
	}

	return resourceFirewallConsolidatedPolicyRead(d, m)
}

func resourceFirewallConsolidatedPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallConsolidatedPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallConsolidatedPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallConsolidatedPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallConsolidatedPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallConsolidatedPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallConsolidatedPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallConsolidatedPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallConsolidatedPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicySrcintfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicySrcintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDstintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyDstintfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyDstintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySrcaddr4(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicySrcaddr4Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicySrcaddr4Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDstaddr4(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyDstaddr4Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyDstaddr4Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicySrcaddr6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicySrcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyDstaddr6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyDstaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceNameName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallConsolidatedPolicyInternetServiceIdId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceCustomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceCustomGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceSrcNameName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallConsolidatedPolicyInternetServiceSrcIdId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceSrcGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceSrcCustomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyInternetServiceSrcCustomGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInspectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyHttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWebproxyProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySpamfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWafProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyUsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyTcpMssSender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyTcpMssReceiver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWanopt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWanoptDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWanoptPassiveOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWanoptProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWanoptPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWebcache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyWebcacheHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyFixedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyIppool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyPoolname4(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyPoolname4Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyPoolname4Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyPoolname6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyPoolname6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyPoolname6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicySessionTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyVpntunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyInbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyOutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyCaptivePortalExempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyFssoGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyFssoGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyFssoGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallConsolidatedPolicyApplicationId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallConsolidatedPolicyApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallConsolidatedPolicyAppCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallConsolidatedPolicyAppCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenFirewallConsolidatedPolicyUrlCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenFirewallConsolidatedPolicyUrlCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallConsolidatedPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallConsolidatedPolicyAppGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallConsolidatedPolicyAppGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallConsolidatedPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("policyid", flattenFirewallConsolidatedPolicyPolicyid(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallConsolidatedPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallConsolidatedPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallConsolidatedPolicyUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcintf", flattenFirewallConsolidatedPolicySrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
			if !fortiAPIPatch(o["srcintf"]) {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf"); ok {
			if err = d.Set("srcintf", flattenFirewallConsolidatedPolicySrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
				if !fortiAPIPatch(o["srcintf"]) {
					return fmt.Errorf("Error reading srcintf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstintf", flattenFirewallConsolidatedPolicyDstintf(o["dstintf"], d, "dstintf", sv)); err != nil {
			if !fortiAPIPatch(o["dstintf"]) {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstintf"); ok {
			if err = d.Set("dstintf", flattenFirewallConsolidatedPolicyDstintf(o["dstintf"], d, "dstintf", sv)); err != nil {
				if !fortiAPIPatch(o["dstintf"]) {
					return fmt.Errorf("Error reading dstintf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr4", flattenFirewallConsolidatedPolicySrcaddr4(o["srcaddr4"], d, "srcaddr4", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr4"]) {
				return fmt.Errorf("Error reading srcaddr4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr4"); ok {
			if err = d.Set("srcaddr4", flattenFirewallConsolidatedPolicySrcaddr4(o["srcaddr4"], d, "srcaddr4", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr4"]) {
					return fmt.Errorf("Error reading srcaddr4: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr4", flattenFirewallConsolidatedPolicyDstaddr4(o["dstaddr4"], d, "dstaddr4", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr4"]) {
				return fmt.Errorf("Error reading dstaddr4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr4"); ok {
			if err = d.Set("dstaddr4", flattenFirewallConsolidatedPolicyDstaddr4(o["dstaddr4"], d, "dstaddr4", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr4"]) {
					return fmt.Errorf("Error reading dstaddr4: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr6", flattenFirewallConsolidatedPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr6"]) {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr6"); ok {
			if err = d.Set("srcaddr6", flattenFirewallConsolidatedPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr6"]) {
					return fmt.Errorf("Error reading srcaddr6: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr6", flattenFirewallConsolidatedPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr6"]) {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr6"); ok {
			if err = d.Set("dstaddr6", flattenFirewallConsolidatedPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr6"]) {
					return fmt.Errorf("Error reading dstaddr6: %v", err)
				}
			}
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallConsolidatedPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["srcaddr-negate"]) {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallConsolidatedPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr-negate"]) {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallConsolidatedPolicyServiceNegate(o["service-negate"], d, "service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["service-negate"]) {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenFirewallConsolidatedPolicyInternetService(o["internet-service"], d, "internet_service", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_name", flattenFirewallConsolidatedPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-name"]) {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_name"); ok {
			if err = d.Set("internet_service_name", flattenFirewallConsolidatedPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-name"]) {
					return fmt.Errorf("Error reading internet_service_name: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_id", flattenFirewallConsolidatedPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-id"]) {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_id"); ok {
			if err = d.Set("internet_service_id", flattenFirewallConsolidatedPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-id"]) {
					return fmt.Errorf("Error reading internet_service_id: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_group", flattenFirewallConsolidatedPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-group"]) {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_group"); ok {
			if err = d.Set("internet_service_group", flattenFirewallConsolidatedPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-group"]) {
					return fmt.Errorf("Error reading internet_service_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom", flattenFirewallConsolidatedPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenFirewallConsolidatedPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("Error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom_group", flattenFirewallConsolidatedPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom_group"); ok {
			if err = d.Set("internet_service_custom_group", flattenFirewallConsolidatedPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallConsolidatedPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-src"]) {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_name", flattenFirewallConsolidatedPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-name"]) {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_name"); ok {
			if err = d.Set("internet_service_src_name", flattenFirewallConsolidatedPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-name"]) {
					return fmt.Errorf("Error reading internet_service_src_name: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_id", flattenFirewallConsolidatedPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-id"]) {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_id"); ok {
			if err = d.Set("internet_service_src_id", flattenFirewallConsolidatedPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-id"]) {
					return fmt.Errorf("Error reading internet_service_src_id: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_group", flattenFirewallConsolidatedPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-group"]) {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_group"); ok {
			if err = d.Set("internet_service_src_group", flattenFirewallConsolidatedPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-group"]) {
					return fmt.Errorf("Error reading internet_service_src_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_custom", flattenFirewallConsolidatedPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom"]) {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom"); ok {
			if err = d.Set("internet_service_src_custom", flattenFirewallConsolidatedPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom"]) {
					return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallConsolidatedPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom_group"); ok {
			if err = d.Set("internet_service_src_custom_group", flattenFirewallConsolidatedPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallConsolidatedPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-negate"]) {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallConsolidatedPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-src-negate"]) {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallConsolidatedPolicyAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallConsolidatedPolicySchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallConsolidatedPolicyService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallConsolidatedPolicyService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("utm_status", flattenFirewallConsolidatedPolicyUtmStatus(o["utm-status"], d, "utm_status", sv)); err != nil {
		if !fortiAPIPatch(o["utm-status"]) {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenFirewallConsolidatedPolicyInspectionMode(o["inspection-mode"], d, "inspection_mode", sv)); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenFirewallConsolidatedPolicyHttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["http-policy-redirect"]) {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenFirewallConsolidatedPolicySshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-policy-redirect"]) {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenFirewallConsolidatedPolicyWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-profile"]) {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallConsolidatedPolicyProfileType(o["profile-type"], d, "profile_type", sv)); err != nil {
		if !fortiAPIPatch(o["profile-type"]) {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallConsolidatedPolicyProfileGroup(o["profile-group"], d, "profile_group", sv)); err != nil {
		if !fortiAPIPatch(o["profile-group"]) {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallConsolidatedPolicyAvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallConsolidatedPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallConsolidatedPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallConsolidatedPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallConsolidatedPolicySpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallConsolidatedPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallConsolidatedPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallConsolidatedPolicyApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallConsolidatedPolicyVoipProfile(o["voip-profile"], d, "voip_profile", sv)); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallConsolidatedPolicyIcapProfile(o["icap-profile"], d, "icap_profile", sv)); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallConsolidatedPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["cifs-profile"]) {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallConsolidatedPolicyWafProfile(o["waf-profile"], d, "waf_profile", sv)); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallConsolidatedPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallConsolidatedPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options", sv)); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallConsolidatedPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallConsolidatedPolicyLogtraffic(o["logtraffic"], d, "logtraffic", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallConsolidatedPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic-start"]) {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenFirewallConsolidatedPolicyAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload", sv)); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenFirewallConsolidatedPolicyGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenFirewallConsolidatedPolicyGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenFirewallConsolidatedPolicyUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenFirewallConsolidatedPolicyUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallConsolidatedPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallConsolidatedPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallConsolidatedPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallConsolidatedPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenFirewallConsolidatedPolicyTcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-sender"]) {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenFirewallConsolidatedPolicyTcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-receiver"]) {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenFirewallConsolidatedPolicyWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-forward-server"]) {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("wanopt", flattenFirewallConsolidatedPolicyWanopt(o["wanopt"], d, "wanopt", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt"]) {
			return fmt.Errorf("Error reading wanopt: %v", err)
		}
	}

	if err = d.Set("wanopt_detection", flattenFirewallConsolidatedPolicyWanoptDetection(o["wanopt-detection"], d, "wanopt_detection", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-detection"]) {
			return fmt.Errorf("Error reading wanopt_detection: %v", err)
		}
	}

	if err = d.Set("wanopt_passive_opt", flattenFirewallConsolidatedPolicyWanoptPassiveOpt(o["wanopt-passive-opt"], d, "wanopt_passive_opt", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-passive-opt"]) {
			return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
		}
	}

	if err = d.Set("wanopt_profile", flattenFirewallConsolidatedPolicyWanoptProfile(o["wanopt-profile"], d, "wanopt_profile", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-profile"]) {
			return fmt.Errorf("Error reading wanopt_profile: %v", err)
		}
	}

	if err = d.Set("wanopt_peer", flattenFirewallConsolidatedPolicyWanoptPeer(o["wanopt-peer"], d, "wanopt_peer", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-peer"]) {
			return fmt.Errorf("Error reading wanopt_peer: %v", err)
		}
	}

	if err = d.Set("webcache", flattenFirewallConsolidatedPolicyWebcache(o["webcache"], d, "webcache", sv)); err != nil {
		if !fortiAPIPatch(o["webcache"]) {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenFirewallConsolidatedPolicyWebcacheHttps(o["webcache-https"], d, "webcache_https", sv)); err != nil {
		if !fortiAPIPatch(o["webcache-https"]) {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallConsolidatedPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-shaper"]) {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallConsolidatedPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-shaper-reverse"]) {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallConsolidatedPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper", sv)); err != nil {
		if !fortiAPIPatch(o["per-ip-shaper"]) {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("nat", flattenFirewallConsolidatedPolicyNat(o["nat"], d, "nat", sv)); err != nil {
		if !fortiAPIPatch(o["nat"]) {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenFirewallConsolidatedPolicyFixedport(o["fixedport"], d, "fixedport", sv)); err != nil {
		if !fortiAPIPatch(o["fixedport"]) {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("ippool", flattenFirewallConsolidatedPolicyIppool(o["ippool"], d, "ippool", sv)); err != nil {
		if !fortiAPIPatch(o["ippool"]) {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("poolname4", flattenFirewallConsolidatedPolicyPoolname4(o["poolname4"], d, "poolname4", sv)); err != nil {
			if !fortiAPIPatch(o["poolname4"]) {
				return fmt.Errorf("Error reading poolname4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("poolname4"); ok {
			if err = d.Set("poolname4", flattenFirewallConsolidatedPolicyPoolname4(o["poolname4"], d, "poolname4", sv)); err != nil {
				if !fortiAPIPatch(o["poolname4"]) {
					return fmt.Errorf("Error reading poolname4: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("poolname6", flattenFirewallConsolidatedPolicyPoolname6(o["poolname6"], d, "poolname6", sv)); err != nil {
			if !fortiAPIPatch(o["poolname6"]) {
				return fmt.Errorf("Error reading poolname6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("poolname6"); ok {
			if err = d.Set("poolname6", flattenFirewallConsolidatedPolicyPoolname6(o["poolname6"], d, "poolname6", sv)); err != nil {
				if !fortiAPIPatch(o["poolname6"]) {
					return fmt.Errorf("Error reading poolname6: %v", err)
				}
			}
		}
	}

	if err = d.Set("session_ttl", flattenFirewallConsolidatedPolicySessionTtl(o["session-ttl"], d, "session_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["session-ttl"]) {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallConsolidatedPolicyComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("vpntunnel", flattenFirewallConsolidatedPolicyVpntunnel(o["vpntunnel"], d, "vpntunnel", sv)); err != nil {
		if !fortiAPIPatch(o["vpntunnel"]) {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("inbound", flattenFirewallConsolidatedPolicyInbound(o["inbound"], d, "inbound", sv)); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("outbound", flattenFirewallConsolidatedPolicyOutbound(o["outbound"], d, "outbound", sv)); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("captive_portal_exempt", flattenFirewallConsolidatedPolicyCaptivePortalExempt(o["captive-portal-exempt"], d, "captive_portal_exempt", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-exempt"]) {
			return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fsso_groups", flattenFirewallConsolidatedPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
			if !fortiAPIPatch(o["fsso-groups"]) {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fsso_groups"); ok {
			if err = d.Set("fsso_groups", flattenFirewallConsolidatedPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
				if !fortiAPIPatch(o["fsso-groups"]) {
					return fmt.Errorf("Error reading fsso_groups: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("application", flattenFirewallConsolidatedPolicyApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallConsolidatedPolicyApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_category", flattenFirewallConsolidatedPolicyAppCategory(o["app-category"], d, "app_category", sv)); err != nil {
			if !fortiAPIPatch(o["app-category"]) {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_category"); ok {
			if err = d.Set("app_category", flattenFirewallConsolidatedPolicyAppCategory(o["app-category"], d, "app_category", sv)); err != nil {
				if !fortiAPIPatch(o["app-category"]) {
					return fmt.Errorf("Error reading app_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_category", flattenFirewallConsolidatedPolicyUrlCategory(o["url-category"], d, "url_category", sv)); err != nil {
			if !fortiAPIPatch(o["url-category"]) {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_category"); ok {
			if err = d.Set("url_category", flattenFirewallConsolidatedPolicyUrlCategory(o["url-category"], d, "url_category", sv)); err != nil {
				if !fortiAPIPatch(o["url-category"]) {
					return fmt.Errorf("Error reading url_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_group", flattenFirewallConsolidatedPolicyAppGroup(o["app-group"], d, "app_group", sv)); err != nil {
			if !fortiAPIPatch(o["app-group"]) {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_group"); ok {
			if err = d.Set("app_group", flattenFirewallConsolidatedPolicyAppGroup(o["app-group"], d, "app_group", sv)); err != nil {
				if !fortiAPIPatch(o["app-group"]) {
					return fmt.Errorf("Error reading app_group: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallConsolidatedPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallConsolidatedPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicySrcintfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicySrcintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyDstintfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyDstintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySrcaddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicySrcaddr4Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicySrcaddr4Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDstaddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyDstaddr4Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyDstaddr4Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicySrcaddr6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicySrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyDstaddr6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyDstaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceNameName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallConsolidatedPolicyInternetServiceIdId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceCustomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceCustomGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceSrcNameName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallConsolidatedPolicyInternetServiceSrcIdId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceSrcGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceSrcCustomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyInternetServiceSrcCustomGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInspectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyHttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWebproxyProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySpamfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyCifsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWafProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyLogtrafficStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyUsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDiffservForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDiffservReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyTcpMssSender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyTcpMssReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWanopt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWanoptDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWanoptPassiveOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWanoptProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWanoptPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWebcache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyWebcacheHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyFixedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyIppool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyPoolname4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyPoolname4Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyPoolname4Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyPoolname6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyPoolname6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyPoolname6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicySessionTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyVpntunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyInbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyOutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyCaptivePortalExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyFssoGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyFssoGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallConsolidatedPolicyApplicationId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallConsolidatedPolicyAppCategoryId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyAppCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandFirewallConsolidatedPolicyUrlCategoryId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyUrlCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallConsolidatedPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallConsolidatedPolicyAppGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallConsolidatedPolicyAppGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallConsolidatedPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("policyid"); ok {

		t, err := expandFirewallConsolidatedPolicyPolicyid(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandFirewallConsolidatedPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallConsolidatedPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {

		t, err := expandFirewallConsolidatedPolicyUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {

		t, err := expandFirewallConsolidatedPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {

		t, err := expandFirewallConsolidatedPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr4"); ok {

		t, err := expandFirewallConsolidatedPolicySrcaddr4(d, v, "srcaddr4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr4"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr4"); ok {

		t, err := expandFirewallConsolidatedPolicyDstaddr4(d, v, "dstaddr4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr4"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok {

		t, err := expandFirewallConsolidatedPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok {

		t, err := expandFirewallConsolidatedPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok {

		t, err := expandFirewallConsolidatedPolicySrcaddrNegate(d, v, "srcaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok {

		t, err := expandFirewallConsolidatedPolicyDstaddrNegate(d, v, "dstaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok {

		t, err := expandFirewallConsolidatedPolicyServiceNegate(d, v, "service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetService(d, v, "internet_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrc(d, v, "internet_service_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrcName(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrcId(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceNegate(d, v, "internet_service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok {

		t, err := expandFirewallConsolidatedPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandFirewallConsolidatedPolicyAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {

		t, err := expandFirewallConsolidatedPolicySchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandFirewallConsolidatedPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok {

		t, err := expandFirewallConsolidatedPolicyUtmStatus(d, v, "utm_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {

		t, err := expandFirewallConsolidatedPolicyInspectionMode(d, v, "inspection_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok {

		t, err := expandFirewallConsolidatedPolicyHttpPolicyRedirect(d, v, "http_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok {

		t, err := expandFirewallConsolidatedPolicySshPolicyRedirect(d, v, "ssh_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyWebproxyProfile(d, v, "webproxy_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok {

		t, err := expandFirewallConsolidatedPolicyProfileType(d, v, "profile_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok {

		t, err := expandFirewallConsolidatedPolicyProfileGroup(d, v, "profile_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyAvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyWebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyDnsfilterProfile(d, v, "dnsfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyEmailfilterProfile(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {

		t, err := expandFirewallConsolidatedPolicySpamfilterProfile(d, v, "spamfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {

		t, err := expandFirewallConsolidatedPolicyDlpSensor(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {

		t, err := expandFirewallConsolidatedPolicyIpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {

		t, err := expandFirewallConsolidatedPolicyApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyVoipProfile(d, v, "voip_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyIcapProfile(d, v, "icap_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("cifs_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyCifsProfile(d, v, "cifs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyWafProfile(d, v, "waf_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {

		t, err := expandFirewallConsolidatedPolicySshFilterProfile(d, v, "ssh_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {

		t, err := expandFirewallConsolidatedPolicyProfileProtocolOptions(d, v, "profile_protocol_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {

		t, err := expandFirewallConsolidatedPolicySslSshProfile(d, v, "ssl_ssh_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {

		t, err := expandFirewallConsolidatedPolicyLogtraffic(d, v, "logtraffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {

		t, err := expandFirewallConsolidatedPolicyLogtrafficStart(d, v, "logtraffic_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok {

		t, err := expandFirewallConsolidatedPolicyAutoAsicOffload(d, v, "auto_asic_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {

		t, err := expandFirewallConsolidatedPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {

		t, err := expandFirewallConsolidatedPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok {

		t, err := expandFirewallConsolidatedPolicyDiffservForward(d, v, "diffserv_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok {

		t, err := expandFirewallConsolidatedPolicyDiffservReverse(d, v, "diffserv_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok {

		t, err := expandFirewallConsolidatedPolicyDiffservcodeForward(d, v, "diffservcode_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok {

		t, err := expandFirewallConsolidatedPolicyDiffservcodeRev(d, v, "diffservcode_rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_mss_sender"); ok {

		t, err := expandFirewallConsolidatedPolicyTcpMssSender(d, v, "tcp_mss_sender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_mss_receiver"); ok {

		t, err := expandFirewallConsolidatedPolicyTcpMssReceiver(d, v, "tcp_mss_receiver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok {

		t, err := expandFirewallConsolidatedPolicyWebproxyForwardServer(d, v, "webproxy_forward_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("wanopt"); ok {

		t, err := expandFirewallConsolidatedPolicyWanopt(d, v, "wanopt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_detection"); ok {

		t, err := expandFirewallConsolidatedPolicyWanoptDetection(d, v, "wanopt_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-detection"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_passive_opt"); ok {

		t, err := expandFirewallConsolidatedPolicyWanoptPassiveOpt(d, v, "wanopt_passive_opt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-passive-opt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_profile"); ok {

		t, err := expandFirewallConsolidatedPolicyWanoptProfile(d, v, "wanopt_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-profile"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_peer"); ok {

		t, err := expandFirewallConsolidatedPolicyWanoptPeer(d, v, "wanopt_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-peer"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok {

		t, err := expandFirewallConsolidatedPolicyWebcache(d, v, "webcache", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok {

		t, err := expandFirewallConsolidatedPolicyWebcacheHttps(d, v, "webcache_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok {

		t, err := expandFirewallConsolidatedPolicyTrafficShaper(d, v, "traffic_shaper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok {

		t, err := expandFirewallConsolidatedPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok {

		t, err := expandFirewallConsolidatedPolicyPerIpShaper(d, v, "per_ip_shaper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok {

		t, err := expandFirewallConsolidatedPolicyNat(d, v, "nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok {

		t, err := expandFirewallConsolidatedPolicyFixedport(d, v, "fixedport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok {

		t, err := expandFirewallConsolidatedPolicyIppool(d, v, "ippool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("poolname4"); ok {

		t, err := expandFirewallConsolidatedPolicyPoolname4(d, v, "poolname4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname4"] = t
		}
	}

	if v, ok := d.GetOk("poolname6"); ok {

		t, err := expandFirewallConsolidatedPolicyPoolname6(d, v, "poolname6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname6"] = t
		}
	}

	if v, ok := d.GetOkExists("session_ttl"); ok {

		t, err := expandFirewallConsolidatedPolicySessionTtl(d, v, "session_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallConsolidatedPolicyComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok {

		t, err := expandFirewallConsolidatedPolicyVpntunnel(d, v, "vpntunnel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok {

		t, err := expandFirewallConsolidatedPolicyInbound(d, v, "inbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok {

		t, err := expandFirewallConsolidatedPolicyOutbound(d, v, "outbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_exempt"); ok {

		t, err := expandFirewallConsolidatedPolicyCaptivePortalExempt(d, v, "captive_portal_exempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-exempt"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok {

		t, err := expandFirewallConsolidatedPolicyFssoGroups(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {

		t, err := expandFirewallConsolidatedPolicyApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok {

		t, err := expandFirewallConsolidatedPolicyAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok {

		t, err := expandFirewallConsolidatedPolicyUrlCategory(d, v, "url_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok {

		t, err := expandFirewallConsolidatedPolicyAppGroup(d, v, "app_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	return &obj, nil
}
