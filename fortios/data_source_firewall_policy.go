// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallPolicyRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ztna_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ztna_ems_tag": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ztna_geo_tag": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"reputation_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reputation_direction": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src_vendor_mac": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"rtp_nat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rtp_addr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tos_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"geoip_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"geoip_match": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"passive_wan_health_measurement": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inspection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"voip_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"icap_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cifs_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"videofilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"np_acceleration": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanopt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanopt_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanopt_passive_opt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanopt_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wanopt_peer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webcache": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webcache_https": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webproxy_forward_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"permit_any_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"permit_stun_host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"poolname6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_cos_fwd": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_cos_rev": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"outbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"natinbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"natoutbound": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fec": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntlm_guest": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntlm_enabled_browsers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_agent_string": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fsso": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wsso": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"devices": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"auth_path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"disclaimer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_collect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpntunnel": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"natip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"match_vip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"match_vip_only": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_mss_sender": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_redirect_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity_based_route": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"block_notification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeout_send_rst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"captive_portal_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_mirror_intf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"radius_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"delay_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sgt_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sgt": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFirewallPolicyRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("policyid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing FirewallPolicy: type error")
	}

	o, err := c.ReadFirewallPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallPolicy: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallPolicy from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicySrcintfName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicySrcintfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyDstintfName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyDstintfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicySrcaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyDstaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicySrcaddr6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicySrcaddr6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyDstaddr6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyDstaddr6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyZtnaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyZtnaEmsTag(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyZtnaEmsTagName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyZtnaEmsTagName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyZtnaGeoTag(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyZtnaGeoTagName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyZtnaGeoTagName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceNameName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicyInternetServiceIdId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceCustomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceCustomGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceSrcNameName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcNameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicyInternetServiceSrcIdId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcIdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceSrcGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceSrcCustomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyInternetServiceSrcCustomGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcCustomGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyReputationMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyReputationDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySrcVendorMac(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicySrcVendorMacId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicySrcVendorMacId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyRtpNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyRtpAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyRtpAddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyRtpAddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyLearningMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNat64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNat46(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyScheduleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyServiceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAntiReplay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyGeoipAnycast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyGeoipMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDynamicShaping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyPassiveWanHealthMeasurement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyHttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySpamfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFileFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySctpFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyVideofilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyCapturePacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNpAcceleration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWanopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWanoptDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWanoptPassiveOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWanoptProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWanoptPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWebcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWebcacheHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicyApplicationId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyApplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicyAppCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyAppCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicyUrlCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyUrlCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyAppGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyAppGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyPermitStunHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFixedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyPoolname(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyPoolnameName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyPoolnameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyPoolname6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyPoolname6Name(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyPoolname6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyVlanCosFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyVlanCosRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNatinbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNatoutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNtlmGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNtlmEnabledBrowsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_string"
		if _, ok := i["user-agent-string"]; ok {
			tmp["user_agent_string"] = dataSourceFlattenFirewallPolicyNtlmEnabledBrowsersUserAgentString(i["user-agent-string"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyNtlmEnabledBrowsersUserAgentString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyWsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyFssoGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyFssoGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyFssoGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDevices(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicyDevicesName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyDevicesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAuthPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyEmailCollect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyVpntunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyNatip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenFirewallPolicyMatchVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyMatchVipOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTcpMssSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTcpMssReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyAuthRedirectAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyBlockNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyCustomLogFields(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "field_id"
		if _, ok := i["field-id"]; ok {
			tmp["field_id"] = dataSourceFlattenFirewallPolicyCustomLogFieldsFieldId(i["field-id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicyCustomLogFieldsFieldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyTimeoutSendRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyCaptivePortalExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySslMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySslMirrorIntf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenFirewallPolicySslMirrorIntfName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicySslMirrorIntfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyRadiusMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicyVlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySgtCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallPolicySgt(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallPolicySgtId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallPolicySgtId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("policyid", dataSourceFlattenFirewallPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenFirewallPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenFirewallPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("srcintf", dataSourceFlattenFirewallPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if !fortiAPIPatch(o["srcintf"]) {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("dstintf", dataSourceFlattenFirewallPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if !fortiAPIPatch(o["dstintf"]) {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("srcaddr", dataSourceFlattenFirewallPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if !fortiAPIPatch(o["srcaddr"]) {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr", dataSourceFlattenFirewallPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if !fortiAPIPatch(o["dstaddr"]) {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr6", dataSourceFlattenFirewallPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if !fortiAPIPatch(o["srcaddr6"]) {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("dstaddr6", dataSourceFlattenFirewallPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if !fortiAPIPatch(o["dstaddr6"]) {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("ztna_status", dataSourceFlattenFirewallPolicyZtnaStatus(o["ztna-status"], d, "ztna_status")); err != nil {
		if !fortiAPIPatch(o["ztna-status"]) {
			return fmt.Errorf("Error reading ztna_status: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag", dataSourceFlattenFirewallPolicyZtnaEmsTag(o["ztna-ems-tag"], d, "ztna_ems_tag")); err != nil {
		if !fortiAPIPatch(o["ztna-ems-tag"]) {
			return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
		}
	}

	if err = d.Set("ztna_geo_tag", dataSourceFlattenFirewallPolicyZtnaGeoTag(o["ztna-geo-tag"], d, "ztna_geo_tag")); err != nil {
		if !fortiAPIPatch(o["ztna-geo-tag"]) {
			return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
		}
	}

	if err = d.Set("internet_service", dataSourceFlattenFirewallPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_name", dataSourceFlattenFirewallPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name")); err != nil {
		if !fortiAPIPatch(o["internet-service-name"]) {
			return fmt.Errorf("Error reading internet_service_name: %v", err)
		}
	}

	if err = d.Set("internet_service_id", dataSourceFlattenFirewallPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if !fortiAPIPatch(o["internet-service-id"]) {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("internet_service_group", dataSourceFlattenFirewallPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-group"]) {
			return fmt.Errorf("Error reading internet_service_group: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", dataSourceFlattenFirewallPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if !fortiAPIPatch(o["internet-service-custom"]) {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_custom_group", dataSourceFlattenFirewallPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-custom-group"]) {
			return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src", dataSourceFlattenFirewallPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if !fortiAPIPatch(o["internet-service-src"]) {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if err = d.Set("internet_service_src_name", dataSourceFlattenFirewallPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-name"]) {
			return fmt.Errorf("Error reading internet_service_src_name: %v", err)
		}
	}

	if err = d.Set("internet_service_src_id", dataSourceFlattenFirewallPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-id"]) {
			return fmt.Errorf("Error reading internet_service_src_id: %v", err)
		}
	}

	if err = d.Set("internet_service_src_group", dataSourceFlattenFirewallPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-group"]) {
			return fmt.Errorf("Error reading internet_service_src_group: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom", dataSourceFlattenFirewallPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-custom"]) {
			return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
		}
	}

	if err = d.Set("internet_service_src_custom_group", dataSourceFlattenFirewallPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
			return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
		}
	}

	if err = d.Set("reputation_minimum", dataSourceFlattenFirewallPolicyReputationMinimum(o["reputation-minimum"], d, "reputation_minimum")); err != nil {
		if !fortiAPIPatch(o["reputation-minimum"]) {
			return fmt.Errorf("Error reading reputation_minimum: %v", err)
		}
	}

	if err = d.Set("reputation_direction", dataSourceFlattenFirewallPolicyReputationDirection(o["reputation-direction"], d, "reputation_direction")); err != nil {
		if !fortiAPIPatch(o["reputation-direction"]) {
			return fmt.Errorf("Error reading reputation_direction: %v", err)
		}
	}

	if err = d.Set("src_vendor_mac", dataSourceFlattenFirewallPolicySrcVendorMac(o["src-vendor-mac"], d, "src_vendor_mac")); err != nil {
		if !fortiAPIPatch(o["src-vendor-mac"]) {
			return fmt.Errorf("Error reading src_vendor_mac: %v", err)
		}
	}

	if err = d.Set("rtp_nat", dataSourceFlattenFirewallPolicyRtpNat(o["rtp-nat"], d, "rtp_nat")); err != nil {
		if !fortiAPIPatch(o["rtp-nat"]) {
			return fmt.Errorf("Error reading rtp_nat: %v", err)
		}
	}

	if err = d.Set("rtp_addr", dataSourceFlattenFirewallPolicyRtpAddr(o["rtp-addr"], d, "rtp_addr")); err != nil {
		if !fortiAPIPatch(o["rtp-addr"]) {
			return fmt.Errorf("Error reading rtp_addr: %v", err)
		}
	}

	if err = d.Set("learning_mode", dataSourceFlattenFirewallPolicyLearningMode(o["learning-mode"], d, "learning_mode")); err != nil {
		if !fortiAPIPatch(o["learning-mode"]) {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("action", dataSourceFlattenFirewallPolicyAction(o["action"], d, "action")); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("nat64", dataSourceFlattenFirewallPolicyNat64(o["nat64"], d, "nat64")); err != nil {
		if !fortiAPIPatch(o["nat64"]) {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("nat46", dataSourceFlattenFirewallPolicyNat46(o["nat46"], d, "nat46")); err != nil {
		if !fortiAPIPatch(o["nat46"]) {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", dataSourceFlattenFirewallPolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if !fortiAPIPatch(o["send-deny-packet"]) {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", dataSourceFlattenFirewallPolicyFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if !fortiAPIPatch(o["firewall-session-dirty"]) {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenFirewallPolicyStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("schedule", dataSourceFlattenFirewallPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("schedule_timeout", dataSourceFlattenFirewallPolicyScheduleTimeout(o["schedule-timeout"], d, "schedule_timeout")); err != nil {
		if !fortiAPIPatch(o["schedule-timeout"]) {
			return fmt.Errorf("Error reading schedule_timeout: %v", err)
		}
	}

	if err = d.Set("service", dataSourceFlattenFirewallPolicyService(o["service"], d, "service")); err != nil {
		if !fortiAPIPatch(o["service"]) {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("tos", dataSourceFlattenFirewallPolicyTos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", dataSourceFlattenFirewallPolicyTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", dataSourceFlattenFirewallPolicyTosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if !fortiAPIPatch(o["tos-negate"]) {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("anti_replay", dataSourceFlattenFirewallPolicyAntiReplay(o["anti-replay"], d, "anti_replay")); err != nil {
		if !fortiAPIPatch(o["anti-replay"]) {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", dataSourceFlattenFirewallPolicyTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn")); err != nil {
		if !fortiAPIPatch(o["tcp-session-without-syn"]) {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("geoip_anycast", dataSourceFlattenFirewallPolicyGeoipAnycast(o["geoip-anycast"], d, "geoip_anycast")); err != nil {
		if !fortiAPIPatch(o["geoip-anycast"]) {
			return fmt.Errorf("Error reading geoip_anycast: %v", err)
		}
	}

	if err = d.Set("geoip_match", dataSourceFlattenFirewallPolicyGeoipMatch(o["geoip-match"], d, "geoip_match")); err != nil {
		if !fortiAPIPatch(o["geoip-match"]) {
			return fmt.Errorf("Error reading geoip_match: %v", err)
		}
	}

	if err = d.Set("dynamic_shaping", dataSourceFlattenFirewallPolicyDynamicShaping(o["dynamic-shaping"], d, "dynamic_shaping")); err != nil {
		if !fortiAPIPatch(o["dynamic-shaping"]) {
			return fmt.Errorf("Error reading dynamic_shaping: %v", err)
		}
	}

	if err = d.Set("passive_wan_health_measurement", dataSourceFlattenFirewallPolicyPassiveWanHealthMeasurement(o["passive-wan-health-measurement"], d, "passive_wan_health_measurement")); err != nil {
		if !fortiAPIPatch(o["passive-wan-health-measurement"]) {
			return fmt.Errorf("Error reading passive_wan_health_measurement: %v", err)
		}
	}

	if err = d.Set("utm_status", dataSourceFlattenFirewallPolicyUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if !fortiAPIPatch(o["utm-status"]) {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("inspection_mode", dataSourceFlattenFirewallPolicyInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", dataSourceFlattenFirewallPolicyHttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect")); err != nil {
		if !fortiAPIPatch(o["http-policy-redirect"]) {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", dataSourceFlattenFirewallPolicySshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect")); err != nil {
		if !fortiAPIPatch(o["ssh-policy-redirect"]) {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", dataSourceFlattenFirewallPolicyWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if !fortiAPIPatch(o["webproxy-profile"]) {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("profile_type", dataSourceFlattenFirewallPolicyProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if !fortiAPIPatch(o["profile-type"]) {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("profile_group", dataSourceFlattenFirewallPolicyProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if !fortiAPIPatch(o["profile-group"]) {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("av_profile", dataSourceFlattenFirewallPolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", dataSourceFlattenFirewallPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", dataSourceFlattenFirewallPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", dataSourceFlattenFirewallPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile")); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", dataSourceFlattenFirewallPolicySpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile")); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", dataSourceFlattenFirewallPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", dataSourceFlattenFirewallPolicyFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile")); err != nil {
		if !fortiAPIPatch(o["file-filter-profile"]) {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("ips_sensor", dataSourceFlattenFirewallPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", dataSourceFlattenFirewallPolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", dataSourceFlattenFirewallPolicyVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", dataSourceFlattenFirewallPolicySctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile")); err != nil {
		if !fortiAPIPatch(o["sctp-filter-profile"]) {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", dataSourceFlattenFirewallPolicyIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", dataSourceFlattenFirewallPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile")); err != nil {
		if !fortiAPIPatch(o["cifs-profile"]) {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", dataSourceFlattenFirewallPolicyVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile")); err != nil {
		if !fortiAPIPatch(o["videofilter-profile"]) {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", dataSourceFlattenFirewallPolicyWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", dataSourceFlattenFirewallPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", dataSourceFlattenFirewallPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", dataSourceFlattenFirewallPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("logtraffic", dataSourceFlattenFirewallPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", dataSourceFlattenFirewallPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if !fortiAPIPatch(o["logtraffic-start"]) {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("capture_packet", dataSourceFlattenFirewallPolicyCapturePacket(o["capture-packet"], d, "capture_packet")); err != nil {
		if !fortiAPIPatch(o["capture-packet"]) {
			return fmt.Errorf("Error reading capture_packet: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", dataSourceFlattenFirewallPolicyAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("np_acceleration", dataSourceFlattenFirewallPolicyNpAcceleration(o["np-acceleration"], d, "np_acceleration")); err != nil {
		if !fortiAPIPatch(o["np-acceleration"]) {
			return fmt.Errorf("Error reading np_acceleration: %v", err)
		}
	}

	if err = d.Set("wanopt", dataSourceFlattenFirewallPolicyWanopt(o["wanopt"], d, "wanopt")); err != nil {
		if !fortiAPIPatch(o["wanopt"]) {
			return fmt.Errorf("Error reading wanopt: %v", err)
		}
	}

	if err = d.Set("wanopt_detection", dataSourceFlattenFirewallPolicyWanoptDetection(o["wanopt-detection"], d, "wanopt_detection")); err != nil {
		if !fortiAPIPatch(o["wanopt-detection"]) {
			return fmt.Errorf("Error reading wanopt_detection: %v", err)
		}
	}

	if err = d.Set("wanopt_passive_opt", dataSourceFlattenFirewallPolicyWanoptPassiveOpt(o["wanopt-passive-opt"], d, "wanopt_passive_opt")); err != nil {
		if !fortiAPIPatch(o["wanopt-passive-opt"]) {
			return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
		}
	}

	if err = d.Set("wanopt_profile", dataSourceFlattenFirewallPolicyWanoptProfile(o["wanopt-profile"], d, "wanopt_profile")); err != nil {
		if !fortiAPIPatch(o["wanopt-profile"]) {
			return fmt.Errorf("Error reading wanopt_profile: %v", err)
		}
	}

	if err = d.Set("wanopt_peer", dataSourceFlattenFirewallPolicyWanoptPeer(o["wanopt-peer"], d, "wanopt_peer")); err != nil {
		if !fortiAPIPatch(o["wanopt-peer"]) {
			return fmt.Errorf("Error reading wanopt_peer: %v", err)
		}
	}

	if err = d.Set("webcache", dataSourceFlattenFirewallPolicyWebcache(o["webcache"], d, "webcache")); err != nil {
		if !fortiAPIPatch(o["webcache"]) {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", dataSourceFlattenFirewallPolicyWebcacheHttps(o["webcache-https"], d, "webcache_https")); err != nil {
		if !fortiAPIPatch(o["webcache-https"]) {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", dataSourceFlattenFirewallPolicyWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server")); err != nil {
		if !fortiAPIPatch(o["webproxy-forward-server"]) {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", dataSourceFlattenFirewallPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if !fortiAPIPatch(o["traffic-shaper"]) {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", dataSourceFlattenFirewallPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if !fortiAPIPatch(o["traffic-shaper-reverse"]) {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", dataSourceFlattenFirewallPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if !fortiAPIPatch(o["per-ip-shaper"]) {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("application", dataSourceFlattenFirewallPolicyApplication(o["application"], d, "application")); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("app_category", dataSourceFlattenFirewallPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
		if !fortiAPIPatch(o["app-category"]) {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("url_category", dataSourceFlattenFirewallPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
		if !fortiAPIPatch(o["url-category"]) {
			return fmt.Errorf("Error reading url_category: %v", err)
		}
	}

	if err = d.Set("app_group", dataSourceFlattenFirewallPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
		if !fortiAPIPatch(o["app-group"]) {
			return fmt.Errorf("Error reading app_group: %v", err)
		}
	}

	if err = d.Set("nat", dataSourceFlattenFirewallPolicyNat(o["nat"], d, "nat")); err != nil {
		if !fortiAPIPatch(o["nat"]) {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("permit_any_host", dataSourceFlattenFirewallPolicyPermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if !fortiAPIPatch(o["permit-any-host"]) {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("permit_stun_host", dataSourceFlattenFirewallPolicyPermitStunHost(o["permit-stun-host"], d, "permit_stun_host")); err != nil {
		if !fortiAPIPatch(o["permit-stun-host"]) {
			return fmt.Errorf("Error reading permit_stun_host: %v", err)
		}
	}

	if err = d.Set("fixedport", dataSourceFlattenFirewallPolicyFixedport(o["fixedport"], d, "fixedport")); err != nil {
		if !fortiAPIPatch(o["fixedport"]) {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("ippool", dataSourceFlattenFirewallPolicyIppool(o["ippool"], d, "ippool")); err != nil {
		if !fortiAPIPatch(o["ippool"]) {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("poolname", dataSourceFlattenFirewallPolicyPoolname(o["poolname"], d, "poolname")); err != nil {
		if !fortiAPIPatch(o["poolname"]) {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("poolname6", dataSourceFlattenFirewallPolicyPoolname6(o["poolname6"], d, "poolname6")); err != nil {
		if !fortiAPIPatch(o["poolname6"]) {
			return fmt.Errorf("Error reading poolname6: %v", err)
		}
	}

	if err = d.Set("session_ttl", dataSourceFlattenFirewallPolicySessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if !fortiAPIPatch(o["session-ttl"]) {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", dataSourceFlattenFirewallPolicyVlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd")); err != nil {
		if !fortiAPIPatch(o["vlan-cos-fwd"]) {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", dataSourceFlattenFirewallPolicyVlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev")); err != nil {
		if !fortiAPIPatch(o["vlan-cos-rev"]) {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("inbound", dataSourceFlattenFirewallPolicyInbound(o["inbound"], d, "inbound")); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("outbound", dataSourceFlattenFirewallPolicyOutbound(o["outbound"], d, "outbound")); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("natinbound", dataSourceFlattenFirewallPolicyNatinbound(o["natinbound"], d, "natinbound")); err != nil {
		if !fortiAPIPatch(o["natinbound"]) {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natoutbound", dataSourceFlattenFirewallPolicyNatoutbound(o["natoutbound"], d, "natoutbound")); err != nil {
		if !fortiAPIPatch(o["natoutbound"]) {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("fec", dataSourceFlattenFirewallPolicyFec(o["fec"], d, "fec")); err != nil {
		if !fortiAPIPatch(o["fec"]) {
			return fmt.Errorf("Error reading fec: %v", err)
		}
	}

	if err = d.Set("wccp", dataSourceFlattenFirewallPolicyWccp(o["wccp"], d, "wccp")); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("ntlm", dataSourceFlattenFirewallPolicyNtlm(o["ntlm"], d, "ntlm")); err != nil {
		if !fortiAPIPatch(o["ntlm"]) {
			return fmt.Errorf("Error reading ntlm: %v", err)
		}
	}

	if err = d.Set("ntlm_guest", dataSourceFlattenFirewallPolicyNtlmGuest(o["ntlm-guest"], d, "ntlm_guest")); err != nil {
		if !fortiAPIPatch(o["ntlm-guest"]) {
			return fmt.Errorf("Error reading ntlm_guest: %v", err)
		}
	}

	if err = d.Set("ntlm_enabled_browsers", dataSourceFlattenFirewallPolicyNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers")); err != nil {
		if !fortiAPIPatch(o["ntlm-enabled-browsers"]) {
			return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
		}
	}

	if err = d.Set("fsso", dataSourceFlattenFirewallPolicyFsso(o["fsso"], d, "fsso")); err != nil {
		if !fortiAPIPatch(o["fsso"]) {
			return fmt.Errorf("Error reading fsso: %v", err)
		}
	}

	if err = d.Set("wsso", dataSourceFlattenFirewallPolicyWsso(o["wsso"], d, "wsso")); err != nil {
		if !fortiAPIPatch(o["wsso"]) {
			return fmt.Errorf("Error reading wsso: %v", err)
		}
	}

	if err = d.Set("rsso", dataSourceFlattenFirewallPolicyRsso(o["rsso"], d, "rsso")); err != nil {
		if !fortiAPIPatch(o["rsso"]) {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", dataSourceFlattenFirewallPolicyFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if !fortiAPIPatch(o["fsso-agent-for-ntlm"]) {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if err = d.Set("groups", dataSourceFlattenFirewallPolicyGroups(o["groups"], d, "groups")); err != nil {
		if !fortiAPIPatch(o["groups"]) {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	if err = d.Set("users", dataSourceFlattenFirewallPolicyUsers(o["users"], d, "users")); err != nil {
		if !fortiAPIPatch(o["users"]) {
			return fmt.Errorf("Error reading users: %v", err)
		}
	}

	if err = d.Set("fsso_groups", dataSourceFlattenFirewallPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups")); err != nil {
		if !fortiAPIPatch(o["fsso-groups"]) {
			return fmt.Errorf("Error reading fsso_groups: %v", err)
		}
	}

	if err = d.Set("devices", dataSourceFlattenFirewallPolicyDevices(o["devices"], d, "devices")); err != nil {
		if !fortiAPIPatch(o["devices"]) {
			return fmt.Errorf("Error reading devices: %v", err)
		}
	}

	if err = d.Set("auth_path", dataSourceFlattenFirewallPolicyAuthPath(o["auth-path"], d, "auth_path")); err != nil {
		if !fortiAPIPatch(o["auth-path"]) {
			return fmt.Errorf("Error reading auth_path: %v", err)
		}
	}

	if err = d.Set("disclaimer", dataSourceFlattenFirewallPolicyDisclaimer(o["disclaimer"], d, "disclaimer")); err != nil {
		if !fortiAPIPatch(o["disclaimer"]) {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if err = d.Set("email_collect", dataSourceFlattenFirewallPolicyEmailCollect(o["email-collect"], d, "email_collect")); err != nil {
		if !fortiAPIPatch(o["email-collect"]) {
			return fmt.Errorf("Error reading email_collect: %v", err)
		}
	}

	if err = d.Set("vpntunnel", dataSourceFlattenFirewallPolicyVpntunnel(o["vpntunnel"], d, "vpntunnel")); err != nil {
		if !fortiAPIPatch(o["vpntunnel"]) {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("natip", dataSourceFlattenFirewallPolicyNatip(o["natip"], d, "natip")); err != nil {
		if !fortiAPIPatch(o["natip"]) {
			return fmt.Errorf("Error reading natip: %v", err)
		}
	}

	if err = d.Set("match_vip", dataSourceFlattenFirewallPolicyMatchVip(o["match-vip"], d, "match_vip")); err != nil {
		if !fortiAPIPatch(o["match-vip"]) {
			return fmt.Errorf("Error reading match_vip: %v", err)
		}
	}

	if err = d.Set("match_vip_only", dataSourceFlattenFirewallPolicyMatchVipOnly(o["match-vip-only"], d, "match_vip_only")); err != nil {
		if !fortiAPIPatch(o["match-vip-only"]) {
			return fmt.Errorf("Error reading match_vip_only: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", dataSourceFlattenFirewallPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", dataSourceFlattenFirewallPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", dataSourceFlattenFirewallPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", dataSourceFlattenFirewallPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", dataSourceFlattenFirewallPolicyTcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender")); err != nil {
		if !fortiAPIPatch(o["tcp-mss-sender"]) {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", dataSourceFlattenFirewallPolicyTcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver")); err != nil {
		if !fortiAPIPatch(o["tcp-mss-receiver"]) {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenFirewallPolicyComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("label", dataSourceFlattenFirewallPolicyLabel(o["label"], d, "label")); err != nil {
		if !fortiAPIPatch(o["label"]) {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("global_label", dataSourceFlattenFirewallPolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if !fortiAPIPatch(o["global-label"]) {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("auth_cert", dataSourceFlattenFirewallPolicyAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_redirect_addr", dataSourceFlattenFirewallPolicyAuthRedirectAddr(o["auth-redirect-addr"], d, "auth_redirect_addr")); err != nil {
		if !fortiAPIPatch(o["auth-redirect-addr"]) {
			return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
		}
	}

	if err = d.Set("redirect_url", dataSourceFlattenFirewallPolicyRedirectUrl(o["redirect-url"], d, "redirect_url")); err != nil {
		if !fortiAPIPatch(o["redirect-url"]) {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("identity_based_route", dataSourceFlattenFirewallPolicyIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route")); err != nil {
		if !fortiAPIPatch(o["identity-based-route"]) {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	if err = d.Set("block_notification", dataSourceFlattenFirewallPolicyBlockNotification(o["block-notification"], d, "block_notification")); err != nil {
		if !fortiAPIPatch(o["block-notification"]) {
			return fmt.Errorf("Error reading block_notification: %v", err)
		}
	}

	if err = d.Set("custom_log_fields", dataSourceFlattenFirewallPolicyCustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
		if !fortiAPIPatch(o["custom-log-fields"]) {
			return fmt.Errorf("Error reading custom_log_fields: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", dataSourceFlattenFirewallPolicyReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", dataSourceFlattenFirewallPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if !fortiAPIPatch(o["srcaddr-negate"]) {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", dataSourceFlattenFirewallPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if !fortiAPIPatch(o["dstaddr-negate"]) {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("service_negate", dataSourceFlattenFirewallPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if !fortiAPIPatch(o["service-negate"]) {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", dataSourceFlattenFirewallPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if !fortiAPIPatch(o["internet-service-negate"]) {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", dataSourceFlattenFirewallPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-negate"]) {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("timeout_send_rst", dataSourceFlattenFirewallPolicyTimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst")); err != nil {
		if !fortiAPIPatch(o["timeout-send-rst"]) {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("captive_portal_exempt", dataSourceFlattenFirewallPolicyCaptivePortalExempt(o["captive-portal-exempt"], d, "captive_portal_exempt")); err != nil {
		if !fortiAPIPatch(o["captive-portal-exempt"]) {
			return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", dataSourceFlattenFirewallPolicyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror")); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", dataSourceFlattenFirewallPolicySslMirror(o["ssl-mirror"], d, "ssl_mirror")); err != nil {
		if !fortiAPIPatch(o["ssl-mirror"]) {
			return fmt.Errorf("Error reading ssl_mirror: %v", err)
		}
	}

	if err = d.Set("ssl_mirror_intf", dataSourceFlattenFirewallPolicySslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf")); err != nil {
		if !fortiAPIPatch(o["ssl-mirror-intf"]) {
			return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", dataSourceFlattenFirewallPolicyScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("dsri", dataSourceFlattenFirewallPolicyDsri(o["dsri"], d, "dsri")); err != nil {
		if !fortiAPIPatch(o["dsri"]) {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_bypass", dataSourceFlattenFirewallPolicyRadiusMacAuthBypass(o["radius-mac-auth-bypass"], d, "radius_mac_auth_bypass")); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth-bypass"]) {
			return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", dataSourceFlattenFirewallPolicyDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session")); err != nil {
		if !fortiAPIPatch(o["delay-tcp-npu-session"]) {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("vlan_filter", dataSourceFlattenFirewallPolicyVlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if !fortiAPIPatch(o["vlan-filter"]) {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("sgt_check", dataSourceFlattenFirewallPolicySgtCheck(o["sgt-check"], d, "sgt_check")); err != nil {
		if !fortiAPIPatch(o["sgt-check"]) {
			return fmt.Errorf("Error reading sgt_check: %v", err)
		}
	}

	if err = d.Set("sgt", dataSourceFlattenFirewallPolicySgt(o["sgt"], d, "sgt")); err != nil {
		if !fortiAPIPatch(o["sgt"]) {
			return fmt.Errorf("Error reading sgt: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
