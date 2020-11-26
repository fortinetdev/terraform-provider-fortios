// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallPolicyCreate,
		Read:   resourceFirewallPolicyRead,
		Update: resourceFirewallPolicyUpdate,
		Delete: resourceFirewallPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"policyid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147482999),
				Optional:     true,
				Computed:     true,
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"srcaddr": &schema.Schema{
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
			"dstaddr": &schema.Schema{
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
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"reputation_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reputation_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtp_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtp_addr": &schema.Schema{
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
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
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
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"schedule_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
			"permit_any_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_stun_host": &schema.Schema{
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
							ValidateFunc: validation.StringLenBetween(0, 64),
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
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntlm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntlm_guest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntlm_enabled_browsers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_agent_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"fsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"groups": &schema.Schema{
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
			"users": &schema.Schema{
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
			"auth_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disclaimer": &schema.Schema{
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
			"natip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_vip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"auth_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_redirect_addr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"redirect_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"identity_based_route": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"block_notification": &schema.Schema{
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
			"timeout_send_rst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallPolicy")
	}

	return resourceFirewallPolicyRead(d, m)
}

func resourceFirewallPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallPolicy")
	}

	return resourceFirewallPolicyRead(d, m)
}

func resourceFirewallPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySrcintfName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicySrcintfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDstintfName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyDstintfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySrcaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDstaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyInternetServiceIdId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceCustomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceCustomGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyInternetServiceSrcIdId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceSrcIdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceSrcGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcCustomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceSrcCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcCustomGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyInternetServiceSrcCustomGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReputationMinimum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReputationDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRtpNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRtpAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyRtpAddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyRtpAddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLearningMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyScheduleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyServiceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTosNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInspectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyHttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebproxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyProfileType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySpamfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWafProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCapturePacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanopt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptDetection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptPassiveOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebcacheHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyApplicationId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyApplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyAppCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyAppCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyUrlCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyUrlCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyAppGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyAppGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPermitStunHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFixedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyPoolname(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyPoolnameName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyPoolnameName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVlanCosFwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVlanCosRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyOutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNatinbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNatoutbound(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNtlmGuest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNtlmEnabledBrowsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["user_agent_string"] = flattenFirewallPolicyNtlmEnabledBrowsersUserAgentString(i["user-agent-string"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyNtlmEnabledBrowsersUserAgentString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyWsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRsso(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDevices(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDevicesName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyDevicesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAuthPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVpntunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyNatip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenFirewallPolicyMatchVip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpMssSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTcpMssReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyAuthRedirectAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyBlockNotification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCustomLogFields(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["field_id"] = flattenFirewallPolicyCustomLogFieldsFieldId(i["field-id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicyCustomLogFieldsFieldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyTimeoutSendRst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyCaptivePortalExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySslMirror(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicySslMirrorIntf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySslMirrorIntfName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenFirewallPolicySslMirrorIntfName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDsri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyRadiusMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallPolicyVlanFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("policyid", flattenFirewallPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcintf", flattenFirewallPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
			if !fortiAPIPatch(o["srcintf"]) {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf"); ok {
			if err = d.Set("srcintf", flattenFirewallPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
				if !fortiAPIPatch(o["srcintf"]) {
					return fmt.Errorf("Error reading srcintf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstintf", flattenFirewallPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
			if !fortiAPIPatch(o["dstintf"]) {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstintf"); ok {
			if err = d.Set("dstintf", flattenFirewallPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
				if !fortiAPIPatch(o["dstintf"]) {
					return fmt.Errorf("Error reading dstintf: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr", flattenFirewallPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr", flattenFirewallPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenFirewallPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service", flattenFirewallPolicyInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_id", flattenFirewallPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
			if !fortiAPIPatch(o["internet-service-id"]) {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_id"); ok {
			if err = d.Set("internet_service_id", flattenFirewallPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
				if !fortiAPIPatch(o["internet-service-id"]) {
					return fmt.Errorf("Error reading internet_service_id: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_group", flattenFirewallPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
			if !fortiAPIPatch(o["internet-service-group"]) {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_group"); ok {
			if err = d.Set("internet_service_group", flattenFirewallPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group")); err != nil {
				if !fortiAPIPatch(o["internet-service-group"]) {
					return fmt.Errorf("Error reading internet_service_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom", flattenFirewallPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenFirewallPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("Error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom_group", flattenFirewallPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
			if !fortiAPIPatch(o["internet-service-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom_group"); ok {
			if err = d.Set("internet_service_custom_group", flattenFirewallPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group")); err != nil {
				if !fortiAPIPatch(o["internet-service-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src")); err != nil {
		if !fortiAPIPatch(o["internet-service-src"]) {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_id", flattenFirewallPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id")); err != nil {
			if !fortiAPIPatch(o["internet-service-src-id"]) {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_id"); ok {
			if err = d.Set("internet_service_src_id", flattenFirewallPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id")); err != nil {
				if !fortiAPIPatch(o["internet-service-src-id"]) {
					return fmt.Errorf("Error reading internet_service_src_id: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_group", flattenFirewallPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
			if !fortiAPIPatch(o["internet-service-src-group"]) {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_group"); ok {
			if err = d.Set("internet_service_src_group", flattenFirewallPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group")); err != nil {
				if !fortiAPIPatch(o["internet-service-src-group"]) {
					return fmt.Errorf("Error reading internet_service_src_group: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_custom", flattenFirewallPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom"]) {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom"); ok {
			if err = d.Set("internet_service_src_custom", flattenFirewallPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom")); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom"]) {
					return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom_group"); ok {
			if err = d.Set("internet_service_src_custom_group", flattenFirewallPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group")); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("reputation_minimum", flattenFirewallPolicyReputationMinimum(o["reputation-minimum"], d, "reputation_minimum")); err != nil {
		if !fortiAPIPatch(o["reputation-minimum"]) {
			return fmt.Errorf("Error reading reputation_minimum: %v", err)
		}
	}

	if err = d.Set("reputation_direction", flattenFirewallPolicyReputationDirection(o["reputation-direction"], d, "reputation_direction")); err != nil {
		if !fortiAPIPatch(o["reputation-direction"]) {
			return fmt.Errorf("Error reading reputation_direction: %v", err)
		}
	}

	if err = d.Set("rtp_nat", flattenFirewallPolicyRtpNat(o["rtp-nat"], d, "rtp_nat")); err != nil {
		if !fortiAPIPatch(o["rtp-nat"]) {
			return fmt.Errorf("Error reading rtp_nat: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("rtp_addr", flattenFirewallPolicyRtpAddr(o["rtp-addr"], d, "rtp_addr")); err != nil {
			if !fortiAPIPatch(o["rtp-addr"]) {
				return fmt.Errorf("Error reading rtp_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rtp_addr"); ok {
			if err = d.Set("rtp_addr", flattenFirewallPolicyRtpAddr(o["rtp-addr"], d, "rtp_addr")); err != nil {
				if !fortiAPIPatch(o["rtp-addr"]) {
					return fmt.Errorf("Error reading rtp_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("learning_mode", flattenFirewallPolicyLearningMode(o["learning-mode"], d, "learning_mode")); err != nil {
		if !fortiAPIPatch(o["learning-mode"]) {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallPolicyAction(o["action"], d, "action")); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenFirewallPolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if !fortiAPIPatch(o["send-deny-packet"]) {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenFirewallPolicyFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if !fortiAPIPatch(o["firewall-session-dirty"]) {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallPolicyStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallPolicySchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("schedule_timeout", flattenFirewallPolicyScheduleTimeout(o["schedule-timeout"], d, "schedule_timeout")); err != nil {
		if !fortiAPIPatch(o["schedule-timeout"]) {
			return fmt.Errorf("Error reading schedule_timeout: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallPolicyService(o["service"], d, "service")); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallPolicyService(o["service"], d, "service")); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("tos", flattenFirewallPolicyTos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenFirewallPolicyTosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenFirewallPolicyTosNegate(o["tos-negate"], d, "tos_negate")); err != nil {
		if !fortiAPIPatch(o["tos-negate"]) {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenFirewallPolicyTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn")); err != nil {
		if !fortiAPIPatch(o["tcp-session-without-syn"]) {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallPolicyUtmStatus(o["utm-status"], d, "utm_status")); err != nil {
		if !fortiAPIPatch(o["utm-status"]) {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenFirewallPolicyInspectionMode(o["inspection-mode"], d, "inspection_mode")); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenFirewallPolicyHttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect")); err != nil {
		if !fortiAPIPatch(o["http-policy-redirect"]) {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenFirewallPolicySshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect")); err != nil {
		if !fortiAPIPatch(o["ssh-policy-redirect"]) {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenFirewallPolicyWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile")); err != nil {
		if !fortiAPIPatch(o["webproxy-profile"]) {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallPolicyProfileType(o["profile-type"], d, "profile_type")); err != nil {
		if !fortiAPIPatch(o["profile-type"]) {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallPolicyProfileGroup(o["profile-group"], d, "profile_group")); err != nil {
		if !fortiAPIPatch(o["profile-group"]) {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallPolicyAvProfile(o["av-profile"], d, "av_profile")); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile")); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallPolicySpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile")); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor")); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallPolicyApplicationList(o["application-list"], d, "application_list")); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallPolicyVoipProfile(o["voip-profile"], d, "voip_profile")); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallPolicyIcapProfile(o["icap-profile"], d, "icap_profile")); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallPolicyWafProfile(o["waf-profile"], d, "waf_profile")); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile")); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options")); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile")); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if !fortiAPIPatch(o["logtraffic-start"]) {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("capture_packet", flattenFirewallPolicyCapturePacket(o["capture-packet"], d, "capture_packet")); err != nil {
		if !fortiAPIPatch(o["capture-packet"]) {
			return fmt.Errorf("Error reading capture_packet: %v", err)
		}
	}

	if err = d.Set("wanopt", flattenFirewallPolicyWanopt(o["wanopt"], d, "wanopt")); err != nil {
		if !fortiAPIPatch(o["wanopt"]) {
			return fmt.Errorf("Error reading wanopt: %v", err)
		}
	}

	if err = d.Set("wanopt_detection", flattenFirewallPolicyWanoptDetection(o["wanopt-detection"], d, "wanopt_detection")); err != nil {
		if !fortiAPIPatch(o["wanopt-detection"]) {
			return fmt.Errorf("Error reading wanopt_detection: %v", err)
		}
	}

	if err = d.Set("wanopt_passive_opt", flattenFirewallPolicyWanoptPassiveOpt(o["wanopt-passive-opt"], d, "wanopt_passive_opt")); err != nil {
		if !fortiAPIPatch(o["wanopt-passive-opt"]) {
			return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
		}
	}

	if err = d.Set("wanopt_profile", flattenFirewallPolicyWanoptProfile(o["wanopt-profile"], d, "wanopt_profile")); err != nil {
		if !fortiAPIPatch(o["wanopt-profile"]) {
			return fmt.Errorf("Error reading wanopt_profile: %v", err)
		}
	}

	if err = d.Set("wanopt_peer", flattenFirewallPolicyWanoptPeer(o["wanopt-peer"], d, "wanopt_peer")); err != nil {
		if !fortiAPIPatch(o["wanopt-peer"]) {
			return fmt.Errorf("Error reading wanopt_peer: %v", err)
		}
	}

	if err = d.Set("webcache", flattenFirewallPolicyWebcache(o["webcache"], d, "webcache")); err != nil {
		if !fortiAPIPatch(o["webcache"]) {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenFirewallPolicyWebcacheHttps(o["webcache-https"], d, "webcache_https")); err != nil {
		if !fortiAPIPatch(o["webcache-https"]) {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenFirewallPolicyWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server")); err != nil {
		if !fortiAPIPatch(o["webproxy-forward-server"]) {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if !fortiAPIPatch(o["traffic-shaper"]) {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if !fortiAPIPatch(o["traffic-shaper-reverse"]) {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if !fortiAPIPatch(o["per-ip-shaper"]) {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("application", flattenFirewallPolicyApplication(o["application"], d, "application")); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallPolicyApplication(o["application"], d, "application")); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_category", flattenFirewallPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
			if !fortiAPIPatch(o["app-category"]) {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_category"); ok {
			if err = d.Set("app_category", flattenFirewallPolicyAppCategory(o["app-category"], d, "app_category")); err != nil {
				if !fortiAPIPatch(o["app-category"]) {
					return fmt.Errorf("Error reading app_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_category", flattenFirewallPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
			if !fortiAPIPatch(o["url-category"]) {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_category"); ok {
			if err = d.Set("url_category", flattenFirewallPolicyUrlCategory(o["url-category"], d, "url_category")); err != nil {
				if !fortiAPIPatch(o["url-category"]) {
					return fmt.Errorf("Error reading url_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("app_group", flattenFirewallPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
			if !fortiAPIPatch(o["app-group"]) {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_group"); ok {
			if err = d.Set("app_group", flattenFirewallPolicyAppGroup(o["app-group"], d, "app_group")); err != nil {
				if !fortiAPIPatch(o["app-group"]) {
					return fmt.Errorf("Error reading app_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("nat", flattenFirewallPolicyNat(o["nat"], d, "nat")); err != nil {
		if !fortiAPIPatch(o["nat"]) {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenFirewallPolicyPermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if !fortiAPIPatch(o["permit-any-host"]) {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("permit_stun_host", flattenFirewallPolicyPermitStunHost(o["permit-stun-host"], d, "permit_stun_host")); err != nil {
		if !fortiAPIPatch(o["permit-stun-host"]) {
			return fmt.Errorf("Error reading permit_stun_host: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenFirewallPolicyFixedport(o["fixedport"], d, "fixedport")); err != nil {
		if !fortiAPIPatch(o["fixedport"]) {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("ippool", flattenFirewallPolicyIppool(o["ippool"], d, "ippool")); err != nil {
		if !fortiAPIPatch(o["ippool"]) {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("poolname", flattenFirewallPolicyPoolname(o["poolname"], d, "poolname")); err != nil {
			if !fortiAPIPatch(o["poolname"]) {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("poolname"); ok {
			if err = d.Set("poolname", flattenFirewallPolicyPoolname(o["poolname"], d, "poolname")); err != nil {
				if !fortiAPIPatch(o["poolname"]) {
					return fmt.Errorf("Error reading poolname: %v", err)
				}
			}
		}
	}

	if err = d.Set("session_ttl", flattenFirewallPolicySessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if !fortiAPIPatch(o["session-ttl"]) {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", flattenFirewallPolicyVlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd")); err != nil {
		if !fortiAPIPatch(o["vlan-cos-fwd"]) {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", flattenFirewallPolicyVlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev")); err != nil {
		if !fortiAPIPatch(o["vlan-cos-rev"]) {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("inbound", flattenFirewallPolicyInbound(o["inbound"], d, "inbound")); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("outbound", flattenFirewallPolicyOutbound(o["outbound"], d, "outbound")); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("natinbound", flattenFirewallPolicyNatinbound(o["natinbound"], d, "natinbound")); err != nil {
		if !fortiAPIPatch(o["natinbound"]) {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natoutbound", flattenFirewallPolicyNatoutbound(o["natoutbound"], d, "natoutbound")); err != nil {
		if !fortiAPIPatch(o["natoutbound"]) {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("wccp", flattenFirewallPolicyWccp(o["wccp"], d, "wccp")); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("ntlm", flattenFirewallPolicyNtlm(o["ntlm"], d, "ntlm")); err != nil {
		if !fortiAPIPatch(o["ntlm"]) {
			return fmt.Errorf("Error reading ntlm: %v", err)
		}
	}

	if err = d.Set("ntlm_guest", flattenFirewallPolicyNtlmGuest(o["ntlm-guest"], d, "ntlm_guest")); err != nil {
		if !fortiAPIPatch(o["ntlm-guest"]) {
			return fmt.Errorf("Error reading ntlm_guest: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ntlm_enabled_browsers", flattenFirewallPolicyNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers")); err != nil {
			if !fortiAPIPatch(o["ntlm-enabled-browsers"]) {
				return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ntlm_enabled_browsers"); ok {
			if err = d.Set("ntlm_enabled_browsers", flattenFirewallPolicyNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers")); err != nil {
				if !fortiAPIPatch(o["ntlm-enabled-browsers"]) {
					return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
				}
			}
		}
	}

	if err = d.Set("fsso", flattenFirewallPolicyFsso(o["fsso"], d, "fsso")); err != nil {
		if !fortiAPIPatch(o["fsso"]) {
			return fmt.Errorf("Error reading fsso: %v", err)
		}
	}

	if err = d.Set("wsso", flattenFirewallPolicyWsso(o["wsso"], d, "wsso")); err != nil {
		if !fortiAPIPatch(o["wsso"]) {
			return fmt.Errorf("Error reading wsso: %v", err)
		}
	}

	if err = d.Set("rsso", flattenFirewallPolicyRsso(o["rsso"], d, "rsso")); err != nil {
		if !fortiAPIPatch(o["rsso"]) {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenFirewallPolicyFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm")); err != nil {
		if !fortiAPIPatch(o["fsso-agent-for-ntlm"]) {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenFirewallPolicyGroups(o["groups"], d, "groups")); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenFirewallPolicyGroups(o["groups"], d, "groups")); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenFirewallPolicyUsers(o["users"], d, "users")); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenFirewallPolicyUsers(o["users"], d, "users")); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("devices", flattenFirewallPolicyDevices(o["devices"], d, "devices")); err != nil {
			if !fortiAPIPatch(o["devices"]) {
				return fmt.Errorf("Error reading devices: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("devices"); ok {
			if err = d.Set("devices", flattenFirewallPolicyDevices(o["devices"], d, "devices")); err != nil {
				if !fortiAPIPatch(o["devices"]) {
					return fmt.Errorf("Error reading devices: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_path", flattenFirewallPolicyAuthPath(o["auth-path"], d, "auth_path")); err != nil {
		if !fortiAPIPatch(o["auth-path"]) {
			return fmt.Errorf("Error reading auth_path: %v", err)
		}
	}

	if err = d.Set("disclaimer", flattenFirewallPolicyDisclaimer(o["disclaimer"], d, "disclaimer")); err != nil {
		if !fortiAPIPatch(o["disclaimer"]) {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if err = d.Set("vpntunnel", flattenFirewallPolicyVpntunnel(o["vpntunnel"], d, "vpntunnel")); err != nil {
		if !fortiAPIPatch(o["vpntunnel"]) {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("natip", flattenFirewallPolicyNatip(o["natip"], d, "natip")); err != nil {
		if !fortiAPIPatch(o["natip"]) {
			return fmt.Errorf("Error reading natip: %v", err)
		}
	}

	if err = d.Set("match_vip", flattenFirewallPolicyMatchVip(o["match-vip"], d, "match_vip")); err != nil {
		if !fortiAPIPatch(o["match-vip"]) {
			return fmt.Errorf("Error reading match_vip: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenFirewallPolicyTcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender")); err != nil {
		if !fortiAPIPatch(o["tcp-mss-sender"]) {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenFirewallPolicyTcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver")); err != nil {
		if !fortiAPIPatch(o["tcp-mss-receiver"]) {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallPolicyComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallPolicyLabel(o["label"], d, "label")); err != nil {
		if !fortiAPIPatch(o["label"]) {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("global_label", flattenFirewallPolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if !fortiAPIPatch(o["global-label"]) {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenFirewallPolicyAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_redirect_addr", flattenFirewallPolicyAuthRedirectAddr(o["auth-redirect-addr"], d, "auth_redirect_addr")); err != nil {
		if !fortiAPIPatch(o["auth-redirect-addr"]) {
			return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenFirewallPolicyRedirectUrl(o["redirect-url"], d, "redirect_url")); err != nil {
		if !fortiAPIPatch(o["redirect-url"]) {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("identity_based_route", flattenFirewallPolicyIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route")); err != nil {
		if !fortiAPIPatch(o["identity-based-route"]) {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	if err = d.Set("block_notification", flattenFirewallPolicyBlockNotification(o["block-notification"], d, "block_notification")); err != nil {
		if !fortiAPIPatch(o["block-notification"]) {
			return fmt.Errorf("Error reading block_notification: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("custom_log_fields", flattenFirewallPolicyCustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
			if !fortiAPIPatch(o["custom-log-fields"]) {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_log_fields"); ok {
			if err = d.Set("custom_log_fields", flattenFirewallPolicyCustomLogFields(o["custom-log-fields"], d, "custom_log_fields")); err != nil {
				if !fortiAPIPatch(o["custom-log-fields"]) {
					return fmt.Errorf("Error reading custom_log_fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_override_group", flattenFirewallPolicyReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if !fortiAPIPatch(o["srcaddr-negate"]) {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if !fortiAPIPatch(o["dstaddr-negate"]) {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallPolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if !fortiAPIPatch(o["service-negate"]) {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate")); err != nil {
		if !fortiAPIPatch(o["internet-service-negate"]) {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate")); err != nil {
		if !fortiAPIPatch(o["internet-service-src-negate"]) {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("timeout_send_rst", flattenFirewallPolicyTimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst")); err != nil {
		if !fortiAPIPatch(o["timeout-send-rst"]) {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("captive_portal_exempt", flattenFirewallPolicyCaptivePortalExempt(o["captive-portal-exempt"], d, "captive_portal_exempt")); err != nil {
		if !fortiAPIPatch(o["captive-portal-exempt"]) {
			return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", flattenFirewallPolicySslMirror(o["ssl-mirror"], d, "ssl_mirror")); err != nil {
		if !fortiAPIPatch(o["ssl-mirror"]) {
			return fmt.Errorf("Error reading ssl_mirror: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ssl_mirror_intf", flattenFirewallPolicySslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf")); err != nil {
			if !fortiAPIPatch(o["ssl-mirror-intf"]) {
				return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_mirror_intf"); ok {
			if err = d.Set("ssl_mirror_intf", flattenFirewallPolicySslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf")); err != nil {
				if !fortiAPIPatch(o["ssl-mirror-intf"]) {
					return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
				}
			}
		}
	}

	if err = d.Set("scan_botnet_connections", flattenFirewallPolicyScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallPolicyDsri(o["dsri"], d, "dsri")); err != nil {
		if !fortiAPIPatch(o["dsri"]) {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_bypass", flattenFirewallPolicyRadiusMacAuthBypass(o["radius-mac-auth-bypass"], d, "radius_mac_auth_bypass")); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth-bypass"]) {
			return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", flattenFirewallPolicyDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session")); err != nil {
		if !fortiAPIPatch(o["delay-tcp-npu-session"]) {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenFirewallPolicyVlanFilter(o["vlan-filter"], d, "vlan_filter")); err != nil {
		if !fortiAPIPatch(o["vlan-filter"]) {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	return nil
}

func flattenFirewallPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicySrcintfName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySrcintfName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyDstintfName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDstintfName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicySrcaddrName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyDstaddrName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallPolicyInternetServiceIdId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceIdId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyInternetServiceGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyInternetServiceCustomName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyInternetServiceCustomGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceCustomGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallPolicyInternetServiceSrcIdId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcIdId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyInternetServiceSrcGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyInternetServiceSrcCustomName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyInternetServiceSrcCustomGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcCustomGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationMinimum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRtpNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRtpAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyRtpAddrName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyRtpAddrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLearningMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyScheduleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyServiceName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTosNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInspectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyHttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebproxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySpamfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWafProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCapturePacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanopt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptDetection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptPassiveOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebcacheHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallPolicyApplicationId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyApplicationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallPolicyAppCategoryId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyAppCategoryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallPolicyUrlCategoryId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyUrlCategoryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyAppGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyAppGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPermitStunHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFixedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPoolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyPoolnameName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyPoolnameName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanCosFwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanCosRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyOutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatinbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatoutbound(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWccp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlmGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlmEnabledBrowsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["user-agent-string"], _ = expandFirewallPolicyNtlmEnabledBrowsersUserAgentString(d, i["user_agent_string"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyNtlmEnabledBrowsersUserAgentString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRsso(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyGroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyUsersName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyUsersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDevices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyDevicesName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDevicesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVpntunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyMatchVip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpMssSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpMssReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthRedirectAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIdentityBasedRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyBlockNotification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCustomLogFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["field-id"], _ = expandFirewallPolicyCustomLogFieldsFieldId(d, i["field_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyCustomLogFieldsFieldId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTimeoutSendRst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCaptivePortalExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslMirror(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslMirrorIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicySslMirrorIntfName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySslMirrorIntfName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDsri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRadiusMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDelayTcpNpuSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("policyid"); ok {
		t, err := expandFirewallPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {
		t, err := expandFirewallPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {
		t, err := expandFirewallPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandFirewallPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandFirewallPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {
		t, err := expandFirewallPolicyInternetService(d, v, "internet_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok {
		t, err := expandFirewallPolicyInternetServiceId(d, v, "internet_service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok {
		t, err := expandFirewallPolicyInternetServiceGroup(d, v, "internet_service_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {
		t, err := expandFirewallPolicyInternetServiceCustom(d, v, "internet_service_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok {
		t, err := expandFirewallPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok {
		t, err := expandFirewallPolicyInternetServiceSrc(d, v, "internet_service_src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok {
		t, err := expandFirewallPolicyInternetServiceSrcId(d, v, "internet_service_src_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok {
		t, err := expandFirewallPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok {
		t, err := expandFirewallPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok {
		t, err := expandFirewallPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("reputation_minimum"); ok {
		t, err := expandFirewallPolicyReputationMinimum(d, v, "reputation_minimum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum"] = t
		}
	}

	if v, ok := d.GetOk("reputation_direction"); ok {
		t, err := expandFirewallPolicyReputationDirection(d, v, "reputation_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction"] = t
		}
	}

	if v, ok := d.GetOk("rtp_nat"); ok {
		t, err := expandFirewallPolicyRtpNat(d, v, "rtp_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-nat"] = t
		}
	}

	if v, ok := d.GetOk("rtp_addr"); ok {
		t, err := expandFirewallPolicyRtpAddr(d, v, "rtp_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-addr"] = t
		}
	}

	if v, ok := d.GetOk("learning_mode"); ok {
		t, err := expandFirewallPolicyLearningMode(d, v, "learning_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-mode"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandFirewallPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok {
		t, err := expandFirewallPolicySendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok {
		t, err := expandFirewallPolicyFirewallSessionDirty(d, v, "firewall_session_dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandFirewallPolicySchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("schedule_timeout"); ok {
		t, err := expandFirewallPolicyScheduleTimeout(d, v, "schedule_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-timeout"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandFirewallPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {
		t, err := expandFirewallPolicyTos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {
		t, err := expandFirewallPolicyTosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok {
		t, err := expandFirewallPolicyTosNegate(d, v, "tos_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok {
		t, err := expandFirewallPolicyTcpSessionWithoutSyn(d, v, "tcp_session_without_syn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok {
		t, err := expandFirewallPolicyUtmStatus(d, v, "utm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {
		t, err := expandFirewallPolicyInspectionMode(d, v, "inspection_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok {
		t, err := expandFirewallPolicyHttpPolicyRedirect(d, v, "http_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok {
		t, err := expandFirewallPolicySshPolicyRedirect(d, v, "ssh_policy_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok {
		t, err := expandFirewallPolicyWebproxyProfile(d, v, "webproxy_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_type"); ok {
		t, err := expandFirewallPolicyProfileType(d, v, "profile_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok {
		t, err := expandFirewallPolicyProfileGroup(d, v, "profile_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandFirewallPolicyAvProfile(d, v, "av_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandFirewallPolicyWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {
		t, err := expandFirewallPolicyDnsfilterProfile(d, v, "dnsfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {
		t, err := expandFirewallPolicySpamfilterProfile(d, v, "spamfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {
		t, err := expandFirewallPolicyDlpSensor(d, v, "dlp_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandFirewallPolicyIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandFirewallPolicyApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("voip_profile"); ok {
		t, err := expandFirewallPolicyVoipProfile(d, v, "voip_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	}

	if v, ok := d.GetOk("icap_profile"); ok {
		t, err := expandFirewallPolicyIcapProfile(d, v, "icap_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	}

	if v, ok := d.GetOk("waf_profile"); ok {
		t, err := expandFirewallPolicyWafProfile(d, v, "waf_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {
		t, err := expandFirewallPolicySshFilterProfile(d, v, "ssh_filter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {
		t, err := expandFirewallPolicyProfileProtocolOptions(d, v, "profile_protocol_options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {
		t, err := expandFirewallPolicySslSshProfile(d, v, "ssl_ssh_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandFirewallPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {
		t, err := expandFirewallPolicyLogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("capture_packet"); ok {
		t, err := expandFirewallPolicyCapturePacket(d, v, "capture_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capture-packet"] = t
		}
	}

	if v, ok := d.GetOk("wanopt"); ok {
		t, err := expandFirewallPolicyWanopt(d, v, "wanopt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_detection"); ok {
		t, err := expandFirewallPolicyWanoptDetection(d, v, "wanopt_detection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-detection"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_passive_opt"); ok {
		t, err := expandFirewallPolicyWanoptPassiveOpt(d, v, "wanopt_passive_opt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-passive-opt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_profile"); ok {
		t, err := expandFirewallPolicyWanoptProfile(d, v, "wanopt_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-profile"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_peer"); ok {
		t, err := expandFirewallPolicyWanoptPeer(d, v, "wanopt_peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-peer"] = t
		}
	}

	if v, ok := d.GetOk("webcache"); ok {
		t, err := expandFirewallPolicyWebcache(d, v, "webcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok {
		t, err := expandFirewallPolicyWebcacheHttps(d, v, "webcache_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok {
		t, err := expandFirewallPolicyWebproxyForwardServer(d, v, "webproxy_forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok {
		t, err := expandFirewallPolicyTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok {
		t, err := expandFirewallPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok {
		t, err := expandFirewallPolicyPerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandFirewallPolicyApplication(d, v, "application")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok {
		t, err := expandFirewallPolicyAppCategory(d, v, "app_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok {
		t, err := expandFirewallPolicyUrlCategory(d, v, "url_category")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok {
		t, err := expandFirewallPolicyAppGroup(d, v, "app_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok {
		t, err := expandFirewallPolicyNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok {
		t, err := expandFirewallPolicyPermitAnyHost(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("permit_stun_host"); ok {
		t, err := expandFirewallPolicyPermitStunHost(d, v, "permit_stun_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-stun-host"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok {
		t, err := expandFirewallPolicyFixedport(d, v, "fixedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok {
		t, err := expandFirewallPolicyIppool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok {
		t, err := expandFirewallPolicyPoolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOkExists("session_ttl"); ok {
		t, err := expandFirewallPolicySessionTtl(d, v, "session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_fwd"); ok {
		t, err := expandFirewallPolicyVlanCosFwd(d, v, "vlan_cos_fwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-fwd"] = t
		}
	}

	if v, ok := d.GetOk("vlan_cos_rev"); ok {
		t, err := expandFirewallPolicyVlanCosRev(d, v, "vlan_cos_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-rev"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok {
		t, err := expandFirewallPolicyInbound(d, v, "inbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok {
		t, err := expandFirewallPolicyOutbound(d, v, "outbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("natinbound"); ok {
		t, err := expandFirewallPolicyNatinbound(d, v, "natinbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natinbound"] = t
		}
	}

	if v, ok := d.GetOk("natoutbound"); ok {
		t, err := expandFirewallPolicyNatoutbound(d, v, "natoutbound")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natoutbound"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok {
		t, err := expandFirewallPolicyWccp(d, v, "wccp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("ntlm"); ok {
		t, err := expandFirewallPolicyNtlm(d, v, "ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_guest"); ok {
		t, err := expandFirewallPolicyNtlmGuest(d, v, "ntlm_guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-guest"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_enabled_browsers"); ok {
		t, err := expandFirewallPolicyNtlmEnabledBrowsers(d, v, "ntlm_enabled_browsers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-enabled-browsers"] = t
		}
	}

	if v, ok := d.GetOk("fsso"); ok {
		t, err := expandFirewallPolicyFsso(d, v, "fsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso"] = t
		}
	}

	if v, ok := d.GetOk("wsso"); ok {
		t, err := expandFirewallPolicyWsso(d, v, "wsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wsso"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok {
		t, err := expandFirewallPolicyRsso(d, v, "rsso")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok {
		t, err := expandFirewallPolicyFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok {
		t, err := expandFirewallPolicyGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandFirewallPolicyUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("devices"); ok {
		t, err := expandFirewallPolicyDevices(d, v, "devices")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devices"] = t
		}
	}

	if v, ok := d.GetOk("auth_path"); ok {
		t, err := expandFirewallPolicyAuthPath(d, v, "auth_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-path"] = t
		}
	}

	if v, ok := d.GetOk("disclaimer"); ok {
		t, err := expandFirewallPolicyDisclaimer(d, v, "disclaimer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok {
		t, err := expandFirewallPolicyVpntunnel(d, v, "vpntunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	}

	if v, ok := d.GetOk("natip"); ok {
		t, err := expandFirewallPolicyNatip(d, v, "natip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natip"] = t
		}
	}

	if v, ok := d.GetOk("match_vip"); ok {
		t, err := expandFirewallPolicyMatchVip(d, v, "match_vip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok {
		t, err := expandFirewallPolicyDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok {
		t, err := expandFirewallPolicyDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok {
		t, err := expandFirewallPolicyDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok {
		t, err := expandFirewallPolicyDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_sender"); ok {
		t, err := expandFirewallPolicyTcpMssSender(d, v, "tcp_mss_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_receiver"); ok {
		t, err := expandFirewallPolicyTcpMssReceiver(d, v, "tcp_mss_receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok {
		t, err := expandFirewallPolicyLabel(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok {
		t, err := expandFirewallPolicyGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("auth_cert"); ok {
		t, err := expandFirewallPolicyAuthCert(d, v, "auth_cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	}

	if v, ok := d.GetOk("auth_redirect_addr"); ok {
		t, err := expandFirewallPolicyAuthRedirectAddr(d, v, "auth_redirect_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-redirect-addr"] = t
		}
	}

	if v, ok := d.GetOk("redirect_url"); ok {
		t, err := expandFirewallPolicyRedirectUrl(d, v, "redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("identity_based_route"); ok {
		t, err := expandFirewallPolicyIdentityBasedRoute(d, v, "identity_based_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based-route"] = t
		}
	}

	if v, ok := d.GetOk("block_notification"); ok {
		t, err := expandFirewallPolicyBlockNotification(d, v, "block_notification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notification"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok {
		t, err := expandFirewallPolicyCustomLogFields(d, v, "custom_log_fields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok {
		t, err := expandFirewallPolicyReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok {
		t, err := expandFirewallPolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok {
		t, err := expandFirewallPolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok {
		t, err := expandFirewallPolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok {
		t, err := expandFirewallPolicyInternetServiceNegate(d, v, "internet_service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok {
		t, err := expandFirewallPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("timeout_send_rst"); ok {
		t, err := expandFirewallPolicyTimeoutSendRst(d, v, "timeout_send_rst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-send-rst"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_exempt"); ok {
		t, err := expandFirewallPolicyCaptivePortalExempt(d, v, "captive_portal_exempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-exempt"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror"); ok {
		t, err := expandFirewallPolicySslMirror(d, v, "ssl_mirror")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror"] = t
		}
	}

	if v, ok := d.GetOk("ssl_mirror_intf"); ok {
		t, err := expandFirewallPolicySslMirrorIntf(d, v, "ssl_mirror_intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror-intf"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandFirewallPolicyScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok {
		t, err := expandFirewallPolicyDsri(d, v, "dsri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_bypass"); ok {
		t, err := expandFirewallPolicyRadiusMacAuthBypass(d, v, "radius_mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("delay_tcp_npu_session"); ok {
		t, err := expandFirewallPolicyDelayTcpNpuSession(d, v, "delay_tcp_npu_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-tcp-npu-session"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok {
		t, err := expandFirewallPolicyVlanFilter(d, v, "vlan_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	return &obj, nil
}
