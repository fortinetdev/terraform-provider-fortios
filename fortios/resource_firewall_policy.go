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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2147482999),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
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
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
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
			"srcaddr": &schema.Schema{
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
			"dstaddr": &schema.Schema{
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
			"srcaddr6": &schema.Schema{
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
			"dstaddr6": &schema.Schema{
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
			"ztna_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_device_ownership": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_ems_tag": &schema.Schema{
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
			"ztna_ems_tag_secondary": &schema.Schema{
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
			"ztna_tags_match_logic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_geo_tag": &schema.Schema{
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
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
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
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"internet_service_group": &schema.Schema{
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
			"internet_service_custom": &schema.Schema{
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
			"network_service_dynamic": &schema.Schema{
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
			"internet_service_custom_group": &schema.Schema{
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
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
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
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"internet_service_src_group": &schema.Schema{
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
			"internet_service_src_custom": &schema.Schema{
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
			"network_service_src_dynamic": &schema.Schema{
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
			"internet_service_src_custom_group": &schema.Schema{
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
			"reputation_minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reputation_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_vendor_mac": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"internet_service6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_name": &schema.Schema{
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
			"internet_service6_group": &schema.Schema{
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
			"internet_service6_custom": &schema.Schema{
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
			"internet_service6_custom_group": &schema.Schema{
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
			"internet_service6_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_name": &schema.Schema{
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
			"internet_service6_src_group": &schema.Schema{
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
			"internet_service6_src_custom": &schema.Schema{
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
			"internet_service6_src_custom_group": &schema.Schema{
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
			"reputation_minimum6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reputation_direction6": &schema.Schema{
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
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
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
				Default:      "always",
			},
			"schedule_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_expiry_date_utc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
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
			"anti_replay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_session_without_syn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geoip_anycast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geoip_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_shaping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passive_wan_health_measurement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_monitor": &schema.Schema{
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
			"ztna_policy_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webproxy_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"av_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"dlp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dlp_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"file_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"voip_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ips_voip_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"diameter_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"virtual_patch_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"icap_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"cifs_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"videofilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"waf_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"casb_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
				Computed:     true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
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
			"log_http_transaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"np_acceleration": &schema.Schema{
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
			},
			"wanopt_peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
			},
			"traffic_shaper": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"app_group": &schema.Schema{
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
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcp_outbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcp_inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcp_poolname": &schema.Schema{
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
			"port_preserve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_random": &schema.Schema{
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
			"poolname6": &schema.Schema{
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
			"session_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(300, 604800),
				Optional:     true,
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
			"fec": &schema.Schema{
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
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_agent_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"fsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fsso_agent_for_ntlm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"groups": &schema.Schema{
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
			"users": &schema.Schema{
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
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
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
			"email_collect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpntunnel": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
			"match_vip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_copy": &schema.Schema{
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
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
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
			},
			"global_label": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"auth_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"auth_redirect_addr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"redirect_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"identity_based_route": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"block_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_log_fields": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"replacemsg_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ztna_ems_tag_negate": &schema.Schema{
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
			"internet_service6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_negate": &schema.Schema{
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
			"decrypted_traffic_mirror": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ssl_mirror": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_mirror_intf": &schema.Schema{
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
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"radius_ip_auth_bypass": &schema.Schema{
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
			},
			"sgt_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sgt": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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

func resourceFirewallPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallPolicy(obj, vdomparam)

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

	obj, err := getObjectFirewallPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallPolicy(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallPolicy(mkey, vdomparam)
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

	o, err := c.ReadFirewallPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySrcintfName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicySrcintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDstintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDstintfName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyDstintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySrcaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDstaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySrcaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicySrcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDstaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyDstaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaDeviceOwnership(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaEmsTag(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyZtnaEmsTagName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyZtnaEmsTagName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaEmsTagSecondary(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyZtnaEmsTagSecondaryName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyZtnaEmsTagSecondaryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaTagsMatchLogic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaGeoTag(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyZtnaGeoTagName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyZtnaGeoTagName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceNameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyInternetServiceIdId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicyInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceCustomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNetworkServiceDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyNetworkServiceDynamicName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyNetworkServiceDynamicName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceCustomGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcNameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceSrcNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyInternetServiceSrcIdId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicyInternetServiceSrcIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyInternetServiceSrcGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceSrcGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcCustomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceSrcCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNetworkServiceSrcDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyNetworkServiceSrcDynamicName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyNetworkServiceSrcDynamicName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetServiceSrcCustomGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetServiceSrcCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyReputationMinimum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyReputationDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySrcVendorMac(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicySrcVendorMacId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicySrcVendorMacId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyInternetService6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Name(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6NameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6NameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Group(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6GroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6GroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Custom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6CustomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6CustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6CustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6CustomGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6CustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Src(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6SrcName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6SrcNameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6SrcNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6SrcGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6SrcGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6SrcGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6SrcCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6SrcCustomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6SrcCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6SrcCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyInternetService6SrcCustomGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyInternetService6SrcCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyReputationMinimum6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyReputationDirection6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyRtpNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyRtpAddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyRtpAddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyRtpAddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyLearningMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNat64(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNat46(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyScheduleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyExpiryDate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPolicyExpiryDateUtc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyServiceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTosNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAntiReplay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTcpSessionWithoutSyn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyGeoipAnycast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyGeoipMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDynamicShaping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPassiveWanHealthMeasurement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAppMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyUtmStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInspectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyHttpPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySshPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaPolicyRedirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWebproxyProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyProfileType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyProfileGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDnsfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDlpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySpamfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFileFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyVoipProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyIpsVoipFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySctpFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDiameterFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyVirtualPatchProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyIcapProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyCifsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyVideofilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWafProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySshFilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyCasbProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyProfileProtocolOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySslSshProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyLogtrafficStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyLogHttpTransaction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyCapturePacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNpAcceleration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWanopt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptPassiveOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWanoptPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWebcache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWebcacheHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWebproxyForwardServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPerIpShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyApplicationId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicyApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyAppCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyAppCategoryId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicyAppCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyUrlCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicyUrlCategoryId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicyUrlCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyAppGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyAppGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyAppGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPcpOutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPcpInbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPcpPoolname(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyPcpPoolnameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyPcpPoolnameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPermitAnyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPermitStunHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFixedport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPortPreserve(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPortRandom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyIppool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPoolname(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyPoolnameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyPoolnameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyPoolname6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyPoolname6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyPoolname6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySessionTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyVlanCosFwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyVlanCosRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyInbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyOutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNatinbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNatoutbound(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWccp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNtlm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNtlmGuest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNtlmEnabledBrowsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_agent_string"
		if cur_v, ok := i["user-agent-string"]; ok {
			tmp["user_agent_string"] = flattenFirewallPolicyNtlmEnabledBrowsersUserAgentString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "user_agent_string", d)
	return result
}

func flattenFirewallPolicyNtlmEnabledBrowsersUserAgentString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFsso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyWsso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyRsso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFssoAgentForNtlm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyGroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyUsersName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyFssoGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyFssoGroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyFssoGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDevices(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicyDevicesName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicyDevicesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAuthPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDisclaimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyEmailCollect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyVpntunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyNatip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenFirewallPolicyMatchVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyMatchVipOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservCopy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTcpMssSender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyTcpMssReceiver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallPolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyLabel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAuthCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyAuthRedirectAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyRedirectUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyBlockNotification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyCustomLogFields(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["field-id"]; ok {
			tmp["field_id"] = flattenFirewallPolicyCustomLogFieldsFieldId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "field_id", d)
	return result
}

func flattenFirewallPolicyCustomLogFieldsFieldId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySrcaddr6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDstaddr6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyZtnaEmsTagNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6Negate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyInternetService6SrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyTimeoutSendRst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyCaptivePortalExempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySslMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySslMirrorIntf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallPolicySslMirrorIntfName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallPolicySslMirrorIntfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDsri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyRadiusMacAuthBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyRadiusIpAuthBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicyVlanFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySgtCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallPolicySgt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallPolicySgtId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallPolicySgtId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectFirewallPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("policyid", flattenFirewallPolicyPolicyid(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallPolicyUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcintf", flattenFirewallPolicySrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
			if !fortiAPIPatch(o["srcintf"]) {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf"); ok {
			if err = d.Set("srcintf", flattenFirewallPolicySrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
				if !fortiAPIPatch(o["srcintf"]) {
					return fmt.Errorf("Error reading srcintf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstintf", flattenFirewallPolicyDstintf(o["dstintf"], d, "dstintf", sv)); err != nil {
			if !fortiAPIPatch(o["dstintf"]) {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstintf"); ok {
			if err = d.Set("dstintf", flattenFirewallPolicyDstintf(o["dstintf"], d, "dstintf", sv)); err != nil {
				if !fortiAPIPatch(o["dstintf"]) {
					return fmt.Errorf("Error reading dstintf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr", flattenFirewallPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr", flattenFirewallPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenFirewallPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr6", flattenFirewallPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr6"]) {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr6"); ok {
			if err = d.Set("srcaddr6", flattenFirewallPolicySrcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr6"]) {
					return fmt.Errorf("Error reading srcaddr6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr6", flattenFirewallPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr6"]) {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr6"); ok {
			if err = d.Set("dstaddr6", flattenFirewallPolicyDstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr6"]) {
					return fmt.Errorf("Error reading dstaddr6: %v", err)
				}
			}
		}
	}

	if err = d.Set("ztna_status", flattenFirewallPolicyZtnaStatus(o["ztna-status"], d, "ztna_status", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-status"]) {
			return fmt.Errorf("Error reading ztna_status: %v", err)
		}
	}

	if err = d.Set("ztna_device_ownership", flattenFirewallPolicyZtnaDeviceOwnership(o["ztna-device-ownership"], d, "ztna_device_ownership", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-device-ownership"]) {
			return fmt.Errorf("Error reading ztna_device_ownership: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ztna_ems_tag", flattenFirewallPolicyZtnaEmsTag(o["ztna-ems-tag"], d, "ztna_ems_tag", sv)); err != nil {
			if !fortiAPIPatch(o["ztna-ems-tag"]) {
				return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ztna_ems_tag"); ok {
			if err = d.Set("ztna_ems_tag", flattenFirewallPolicyZtnaEmsTag(o["ztna-ems-tag"], d, "ztna_ems_tag", sv)); err != nil {
				if !fortiAPIPatch(o["ztna-ems-tag"]) {
					return fmt.Errorf("Error reading ztna_ems_tag: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ztna_ems_tag_secondary", flattenFirewallPolicyZtnaEmsTagSecondary(o["ztna-ems-tag-secondary"], d, "ztna_ems_tag_secondary", sv)); err != nil {
			if !fortiAPIPatch(o["ztna-ems-tag-secondary"]) {
				return fmt.Errorf("Error reading ztna_ems_tag_secondary: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ztna_ems_tag_secondary"); ok {
			if err = d.Set("ztna_ems_tag_secondary", flattenFirewallPolicyZtnaEmsTagSecondary(o["ztna-ems-tag-secondary"], d, "ztna_ems_tag_secondary", sv)); err != nil {
				if !fortiAPIPatch(o["ztna-ems-tag-secondary"]) {
					return fmt.Errorf("Error reading ztna_ems_tag_secondary: %v", err)
				}
			}
		}
	}

	if err = d.Set("ztna_tags_match_logic", flattenFirewallPolicyZtnaTagsMatchLogic(o["ztna-tags-match-logic"], d, "ztna_tags_match_logic", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-tags-match-logic"]) {
			return fmt.Errorf("Error reading ztna_tags_match_logic: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ztna_geo_tag", flattenFirewallPolicyZtnaGeoTag(o["ztna-geo-tag"], d, "ztna_geo_tag", sv)); err != nil {
			if !fortiAPIPatch(o["ztna-geo-tag"]) {
				return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ztna_geo_tag"); ok {
			if err = d.Set("ztna_geo_tag", flattenFirewallPolicyZtnaGeoTag(o["ztna-geo-tag"], d, "ztna_geo_tag", sv)); err != nil {
				if !fortiAPIPatch(o["ztna-geo-tag"]) {
					return fmt.Errorf("Error reading ztna_geo_tag: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service", flattenFirewallPolicyInternetService(o["internet-service"], d, "internet_service", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_name", flattenFirewallPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-name"]) {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_name"); ok {
			if err = d.Set("internet_service_name", flattenFirewallPolicyInternetServiceName(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-name"]) {
					return fmt.Errorf("Error reading internet_service_name: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_id", flattenFirewallPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-id"]) {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_id"); ok {
			if err = d.Set("internet_service_id", flattenFirewallPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-id"]) {
					return fmt.Errorf("Error reading internet_service_id: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_group", flattenFirewallPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-group"]) {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_group"); ok {
			if err = d.Set("internet_service_group", flattenFirewallPolicyInternetServiceGroup(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-group"]) {
					return fmt.Errorf("Error reading internet_service_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_custom", flattenFirewallPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenFirewallPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("Error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("network_service_dynamic", flattenFirewallPolicyNetworkServiceDynamic(o["network-service-dynamic"], d, "network_service_dynamic", sv)); err != nil {
			if !fortiAPIPatch(o["network-service-dynamic"]) {
				return fmt.Errorf("Error reading network_service_dynamic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network_service_dynamic"); ok {
			if err = d.Set("network_service_dynamic", flattenFirewallPolicyNetworkServiceDynamic(o["network-service-dynamic"], d, "network_service_dynamic", sv)); err != nil {
				if !fortiAPIPatch(o["network-service-dynamic"]) {
					return fmt.Errorf("Error reading network_service_dynamic: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_custom_group", flattenFirewallPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom_group"); ok {
			if err = d.Set("internet_service_custom_group", flattenFirewallPolicyInternetServiceCustomGroup(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallPolicyInternetServiceSrc(o["internet-service-src"], d, "internet_service_src", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-src"]) {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_name", flattenFirewallPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-name"]) {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_name"); ok {
			if err = d.Set("internet_service_src_name", flattenFirewallPolicyInternetServiceSrcName(o["internet-service-src-name"], d, "internet_service_src_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-name"]) {
					return fmt.Errorf("Error reading internet_service_src_name: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_id", flattenFirewallPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-id"]) {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_id"); ok {
			if err = d.Set("internet_service_src_id", flattenFirewallPolicyInternetServiceSrcId(o["internet-service-src-id"], d, "internet_service_src_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-id"]) {
					return fmt.Errorf("Error reading internet_service_src_id: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_group", flattenFirewallPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-group"]) {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_group"); ok {
			if err = d.Set("internet_service_src_group", flattenFirewallPolicyInternetServiceSrcGroup(o["internet-service-src-group"], d, "internet_service_src_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-group"]) {
					return fmt.Errorf("Error reading internet_service_src_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_custom", flattenFirewallPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom"]) {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom"); ok {
			if err = d.Set("internet_service_src_custom", flattenFirewallPolicyInternetServiceSrcCustom(o["internet-service-src-custom"], d, "internet_service_src_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom"]) {
					return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("network_service_src_dynamic", flattenFirewallPolicyNetworkServiceSrcDynamic(o["network-service-src-dynamic"], d, "network_service_src_dynamic", sv)); err != nil {
			if !fortiAPIPatch(o["network-service-src-dynamic"]) {
				return fmt.Errorf("Error reading network_service_src_dynamic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network_service_src_dynamic"); ok {
			if err = d.Set("network_service_src_dynamic", flattenFirewallPolicyNetworkServiceSrcDynamic(o["network-service-src-dynamic"], d, "network_service_src_dynamic", sv)); err != nil {
				if !fortiAPIPatch(o["network-service-src-dynamic"]) {
					return fmt.Errorf("Error reading network_service_src_dynamic: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom_group"); ok {
			if err = d.Set("internet_service_src_custom_group", flattenFirewallPolicyInternetServiceSrcCustomGroup(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("reputation_minimum", flattenFirewallPolicyReputationMinimum(o["reputation-minimum"], d, "reputation_minimum", sv)); err != nil {
		if !fortiAPIPatch(o["reputation-minimum"]) {
			return fmt.Errorf("Error reading reputation_minimum: %v", err)
		}
	}

	if err = d.Set("reputation_direction", flattenFirewallPolicyReputationDirection(o["reputation-direction"], d, "reputation_direction", sv)); err != nil {
		if !fortiAPIPatch(o["reputation-direction"]) {
			return fmt.Errorf("Error reading reputation_direction: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("src_vendor_mac", flattenFirewallPolicySrcVendorMac(o["src-vendor-mac"], d, "src_vendor_mac", sv)); err != nil {
			if !fortiAPIPatch(o["src-vendor-mac"]) {
				return fmt.Errorf("Error reading src_vendor_mac: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_vendor_mac"); ok {
			if err = d.Set("src_vendor_mac", flattenFirewallPolicySrcVendorMac(o["src-vendor-mac"], d, "src_vendor_mac", sv)); err != nil {
				if !fortiAPIPatch(o["src-vendor-mac"]) {
					return fmt.Errorf("Error reading src_vendor_mac: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service6", flattenFirewallPolicyInternetService6(o["internet-service6"], d, "internet_service6", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6"]) {
			return fmt.Errorf("Error reading internet_service6: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_name", flattenFirewallPolicyInternetService6Name(o["internet-service6-name"], d, "internet_service6_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-name"]) {
				return fmt.Errorf("Error reading internet_service6_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_name"); ok {
			if err = d.Set("internet_service6_name", flattenFirewallPolicyInternetService6Name(o["internet-service6-name"], d, "internet_service6_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-name"]) {
					return fmt.Errorf("Error reading internet_service6_name: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_group", flattenFirewallPolicyInternetService6Group(o["internet-service6-group"], d, "internet_service6_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-group"]) {
				return fmt.Errorf("Error reading internet_service6_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_group"); ok {
			if err = d.Set("internet_service6_group", flattenFirewallPolicyInternetService6Group(o["internet-service6-group"], d, "internet_service6_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-group"]) {
					return fmt.Errorf("Error reading internet_service6_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_custom", flattenFirewallPolicyInternetService6Custom(o["internet-service6-custom"], d, "internet_service6_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-custom"]) {
				return fmt.Errorf("Error reading internet_service6_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_custom"); ok {
			if err = d.Set("internet_service6_custom", flattenFirewallPolicyInternetService6Custom(o["internet-service6-custom"], d, "internet_service6_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-custom"]) {
					return fmt.Errorf("Error reading internet_service6_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_custom_group", flattenFirewallPolicyInternetService6CustomGroup(o["internet-service6-custom-group"], d, "internet_service6_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-custom-group"]) {
				return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_custom_group"); ok {
			if err = d.Set("internet_service6_custom_group", flattenFirewallPolicyInternetService6CustomGroup(o["internet-service6-custom-group"], d, "internet_service6_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-custom-group"]) {
					return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service6_src", flattenFirewallPolicyInternetService6Src(o["internet-service6-src"], d, "internet_service6_src", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6-src"]) {
			return fmt.Errorf("Error reading internet_service6_src: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_name", flattenFirewallPolicyInternetService6SrcName(o["internet-service6-src-name"], d, "internet_service6_src_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-name"]) {
				return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_name"); ok {
			if err = d.Set("internet_service6_src_name", flattenFirewallPolicyInternetService6SrcName(o["internet-service6-src-name"], d, "internet_service6_src_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-name"]) {
					return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_group", flattenFirewallPolicyInternetService6SrcGroup(o["internet-service6-src-group"], d, "internet_service6_src_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-group"]) {
				return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_group"); ok {
			if err = d.Set("internet_service6_src_group", flattenFirewallPolicyInternetService6SrcGroup(o["internet-service6-src-group"], d, "internet_service6_src_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-group"]) {
					return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_custom", flattenFirewallPolicyInternetService6SrcCustom(o["internet-service6-src-custom"], d, "internet_service6_src_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-custom"]) {
				return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_custom"); ok {
			if err = d.Set("internet_service6_src_custom", flattenFirewallPolicyInternetService6SrcCustom(o["internet-service6-src-custom"], d, "internet_service6_src_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-custom"]) {
					return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_custom_group", flattenFirewallPolicyInternetService6SrcCustomGroup(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-custom-group"]) {
				return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_custom_group"); ok {
			if err = d.Set("internet_service6_src_custom_group", flattenFirewallPolicyInternetService6SrcCustomGroup(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-custom-group"]) {
					return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("reputation_minimum6", flattenFirewallPolicyReputationMinimum6(o["reputation-minimum6"], d, "reputation_minimum6", sv)); err != nil {
		if !fortiAPIPatch(o["reputation-minimum6"]) {
			return fmt.Errorf("Error reading reputation_minimum6: %v", err)
		}
	}

	if err = d.Set("reputation_direction6", flattenFirewallPolicyReputationDirection6(o["reputation-direction6"], d, "reputation_direction6", sv)); err != nil {
		if !fortiAPIPatch(o["reputation-direction6"]) {
			return fmt.Errorf("Error reading reputation_direction6: %v", err)
		}
	}

	if err = d.Set("rtp_nat", flattenFirewallPolicyRtpNat(o["rtp-nat"], d, "rtp_nat", sv)); err != nil {
		if !fortiAPIPatch(o["rtp-nat"]) {
			return fmt.Errorf("Error reading rtp_nat: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("rtp_addr", flattenFirewallPolicyRtpAddr(o["rtp-addr"], d, "rtp_addr", sv)); err != nil {
			if !fortiAPIPatch(o["rtp-addr"]) {
				return fmt.Errorf("Error reading rtp_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rtp_addr"); ok {
			if err = d.Set("rtp_addr", flattenFirewallPolicyRtpAddr(o["rtp-addr"], d, "rtp_addr", sv)); err != nil {
				if !fortiAPIPatch(o["rtp-addr"]) {
					return fmt.Errorf("Error reading rtp_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("learning_mode", flattenFirewallPolicyLearningMode(o["learning-mode"], d, "learning_mode", sv)); err != nil {
		if !fortiAPIPatch(o["learning-mode"]) {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallPolicyAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallPolicyNat64(o["nat64"], d, "nat64", sv)); err != nil {
		if !fortiAPIPatch(o["nat64"]) {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallPolicyNat46(o["nat46"], d, "nat46", sv)); err != nil {
		if !fortiAPIPatch(o["nat46"]) {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenFirewallPolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet", sv)); err != nil {
		if !fortiAPIPatch(o["send-deny-packet"]) {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenFirewallPolicyFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty", sv)); err != nil {
		if !fortiAPIPatch(o["firewall-session-dirty"]) {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallPolicySchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("schedule_timeout", flattenFirewallPolicyScheduleTimeout(o["schedule-timeout"], d, "schedule_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-timeout"]) {
			return fmt.Errorf("Error reading schedule_timeout: %v", err)
		}
	}

	if err = d.Set("policy_expiry", flattenFirewallPolicyPolicyExpiry(o["policy-expiry"], d, "policy_expiry", sv)); err != nil {
		if !fortiAPIPatch(o["policy-expiry"]) {
			return fmt.Errorf("Error reading policy_expiry: %v", err)
		}
	}

	if err = d.Set("policy_expiry_date", flattenFirewallPolicyPolicyExpiryDate(o["policy-expiry-date"], d, "policy_expiry_date", sv)); err != nil {
		if !fortiAPIPatch(o["policy-expiry-date"]) {
			return fmt.Errorf("Error reading policy_expiry_date: %v", err)
		}
	}

	if err = d.Set("policy_expiry_date_utc", flattenFirewallPolicyPolicyExpiryDateUtc(o["policy-expiry-date-utc"], d, "policy_expiry_date_utc", sv)); err != nil {
		if !fortiAPIPatch(o["policy-expiry-date-utc"]) {
			return fmt.Errorf("Error reading policy_expiry_date_utc: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("service", flattenFirewallPolicyService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallPolicyService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("tos", flattenFirewallPolicyTos(o["tos"], d, "tos", sv)); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenFirewallPolicyTosMask(o["tos-mask"], d, "tos_mask", sv)); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("tos_negate", flattenFirewallPolicyTosNegate(o["tos-negate"], d, "tos_negate", sv)); err != nil {
		if !fortiAPIPatch(o["tos-negate"]) {
			return fmt.Errorf("Error reading tos_negate: %v", err)
		}
	}

	if err = d.Set("anti_replay", flattenFirewallPolicyAntiReplay(o["anti-replay"], d, "anti_replay", sv)); err != nil {
		if !fortiAPIPatch(o["anti-replay"]) {
			return fmt.Errorf("Error reading anti_replay: %v", err)
		}
	}

	if err = d.Set("tcp_session_without_syn", flattenFirewallPolicyTcpSessionWithoutSyn(o["tcp-session-without-syn"], d, "tcp_session_without_syn", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-session-without-syn"]) {
			return fmt.Errorf("Error reading tcp_session_without_syn: %v", err)
		}
	}

	if err = d.Set("geoip_anycast", flattenFirewallPolicyGeoipAnycast(o["geoip-anycast"], d, "geoip_anycast", sv)); err != nil {
		if !fortiAPIPatch(o["geoip-anycast"]) {
			return fmt.Errorf("Error reading geoip_anycast: %v", err)
		}
	}

	if err = d.Set("geoip_match", flattenFirewallPolicyGeoipMatch(o["geoip-match"], d, "geoip_match", sv)); err != nil {
		if !fortiAPIPatch(o["geoip-match"]) {
			return fmt.Errorf("Error reading geoip_match: %v", err)
		}
	}

	if err = d.Set("dynamic_shaping", flattenFirewallPolicyDynamicShaping(o["dynamic-shaping"], d, "dynamic_shaping", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-shaping"]) {
			return fmt.Errorf("Error reading dynamic_shaping: %v", err)
		}
	}

	if err = d.Set("passive_wan_health_measurement", flattenFirewallPolicyPassiveWanHealthMeasurement(o["passive-wan-health-measurement"], d, "passive_wan_health_measurement", sv)); err != nil {
		if !fortiAPIPatch(o["passive-wan-health-measurement"]) {
			return fmt.Errorf("Error reading passive_wan_health_measurement: %v", err)
		}
	}

	if err = d.Set("app_monitor", flattenFirewallPolicyAppMonitor(o["app-monitor"], d, "app_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["app-monitor"]) {
			return fmt.Errorf("Error reading app_monitor: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenFirewallPolicyUtmStatus(o["utm-status"], d, "utm_status", sv)); err != nil {
		if !fortiAPIPatch(o["utm-status"]) {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("inspection_mode", flattenFirewallPolicyInspectionMode(o["inspection-mode"], d, "inspection_mode", sv)); err != nil {
		if !fortiAPIPatch(o["inspection-mode"]) {
			return fmt.Errorf("Error reading inspection_mode: %v", err)
		}
	}

	if err = d.Set("http_policy_redirect", flattenFirewallPolicyHttpPolicyRedirect(o["http-policy-redirect"], d, "http_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["http-policy-redirect"]) {
			return fmt.Errorf("Error reading http_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ssh_policy_redirect", flattenFirewallPolicySshPolicyRedirect(o["ssh-policy-redirect"], d, "ssh_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-policy-redirect"]) {
			return fmt.Errorf("Error reading ssh_policy_redirect: %v", err)
		}
	}

	if err = d.Set("ztna_policy_redirect", flattenFirewallPolicyZtnaPolicyRedirect(o["ztna-policy-redirect"], d, "ztna_policy_redirect", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-policy-redirect"]) {
			return fmt.Errorf("Error reading ztna_policy_redirect: %v", err)
		}
	}

	if err = d.Set("webproxy_profile", flattenFirewallPolicyWebproxyProfile(o["webproxy-profile"], d, "webproxy_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-profile"]) {
			return fmt.Errorf("Error reading webproxy_profile: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallPolicyProfileType(o["profile-type"], d, "profile_type", sv)); err != nil {
		if !fortiAPIPatch(o["profile-type"]) {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallPolicyProfileGroup(o["profile-group"], d, "profile_group", sv)); err != nil {
		if !fortiAPIPatch(o["profile-group"]) {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallPolicyAvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallPolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallPolicyDnsfilterProfile(o["dnsfilter-profile"], d, "dnsfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallPolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallPolicyDlpProfile(o["dlp-profile"], d, "dlp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-profile"]) {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallPolicySpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallPolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallPolicyFileFilterProfile(o["file-filter-profile"], d, "file_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["file-filter-profile"]) {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallPolicyIpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallPolicyApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallPolicyVoipProfile(o["voip-profile"], d, "voip_profile", sv)); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallPolicyIpsVoipFilter(o["ips-voip-filter"], d, "ips_voip_filter", sv)); err != nil {
		if !fortiAPIPatch(o["ips-voip-filter"]) {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallPolicySctpFilterProfile(o["sctp-filter-profile"], d, "sctp_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-filter-profile"]) {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallPolicyDiameterFilterProfile(o["diameter-filter-profile"], d, "diameter_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["diameter-filter-profile"]) {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallPolicyVirtualPatchProfile(o["virtual-patch-profile"], d, "virtual_patch_profile", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-patch-profile"]) {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallPolicyIcapProfile(o["icap-profile"], d, "icap_profile", sv)); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallPolicyCifsProfile(o["cifs-profile"], d, "cifs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["cifs-profile"]) {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallPolicyVideofilterProfile(o["videofilter-profile"], d, "videofilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["videofilter-profile"]) {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("waf_profile", flattenFirewallPolicyWafProfile(o["waf-profile"], d, "waf_profile", sv)); err != nil {
		if !fortiAPIPatch(o["waf-profile"]) {
			return fmt.Errorf("Error reading waf_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallPolicySshFilterProfile(o["ssh-filter-profile"], d, "ssh_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallPolicyCasbProfile(o["casb-profile"], d, "casb_profile", sv)); err != nil {
		if !fortiAPIPatch(o["casb-profile"]) {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallPolicyProfileProtocolOptions(o["profile-protocol-options"], d, "profile_protocol_options", sv)); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallPolicySslSshProfile(o["ssl-ssh-profile"], d, "ssl_ssh_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallPolicyLogtraffic(o["logtraffic"], d, "logtraffic", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallPolicyLogtrafficStart(o["logtraffic-start"], d, "logtraffic_start", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic-start"]) {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("log_http_transaction", flattenFirewallPolicyLogHttpTransaction(o["log-http-transaction"], d, "log_http_transaction", sv)); err != nil {
		if !fortiAPIPatch(o["log-http-transaction"]) {
			return fmt.Errorf("Error reading log_http_transaction: %v", err)
		}
	}

	if err = d.Set("capture_packet", flattenFirewallPolicyCapturePacket(o["capture-packet"], d, "capture_packet", sv)); err != nil {
		if !fortiAPIPatch(o["capture-packet"]) {
			return fmt.Errorf("Error reading capture_packet: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenFirewallPolicyAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload", sv)); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("np_acceleration", flattenFirewallPolicyNpAcceleration(o["np-acceleration"], d, "np_acceleration", sv)); err != nil {
		if !fortiAPIPatch(o["np-acceleration"]) {
			return fmt.Errorf("Error reading np_acceleration: %v", err)
		}
	}

	if err = d.Set("wanopt", flattenFirewallPolicyWanopt(o["wanopt"], d, "wanopt", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt"]) {
			return fmt.Errorf("Error reading wanopt: %v", err)
		}
	}

	if err = d.Set("wanopt_detection", flattenFirewallPolicyWanoptDetection(o["wanopt-detection"], d, "wanopt_detection", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-detection"]) {
			return fmt.Errorf("Error reading wanopt_detection: %v", err)
		}
	}

	if err = d.Set("wanopt_passive_opt", flattenFirewallPolicyWanoptPassiveOpt(o["wanopt-passive-opt"], d, "wanopt_passive_opt", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-passive-opt"]) {
			return fmt.Errorf("Error reading wanopt_passive_opt: %v", err)
		}
	}

	if err = d.Set("wanopt_profile", flattenFirewallPolicyWanoptProfile(o["wanopt-profile"], d, "wanopt_profile", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-profile"]) {
			return fmt.Errorf("Error reading wanopt_profile: %v", err)
		}
	}

	if err = d.Set("wanopt_peer", flattenFirewallPolicyWanoptPeer(o["wanopt-peer"], d, "wanopt_peer", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-peer"]) {
			return fmt.Errorf("Error reading wanopt_peer: %v", err)
		}
	}

	if err = d.Set("webcache", flattenFirewallPolicyWebcache(o["webcache"], d, "webcache", sv)); err != nil {
		if !fortiAPIPatch(o["webcache"]) {
			return fmt.Errorf("Error reading webcache: %v", err)
		}
	}

	if err = d.Set("webcache_https", flattenFirewallPolicyWebcacheHttps(o["webcache-https"], d, "webcache_https", sv)); err != nil {
		if !fortiAPIPatch(o["webcache-https"]) {
			return fmt.Errorf("Error reading webcache_https: %v", err)
		}
	}

	if err = d.Set("webproxy_forward_server", flattenFirewallPolicyWebproxyForwardServer(o["webproxy-forward-server"], d, "webproxy_forward_server", sv)); err != nil {
		if !fortiAPIPatch(o["webproxy-forward-server"]) {
			return fmt.Errorf("Error reading webproxy_forward_server: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenFirewallPolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-shaper"]) {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenFirewallPolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-shaper-reverse"]) {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenFirewallPolicyPerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper", sv)); err != nil {
		if !fortiAPIPatch(o["per-ip-shaper"]) {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("application", flattenFirewallPolicyApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallPolicyApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("app_category", flattenFirewallPolicyAppCategory(o["app-category"], d, "app_category", sv)); err != nil {
			if !fortiAPIPatch(o["app-category"]) {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_category"); ok {
			if err = d.Set("app_category", flattenFirewallPolicyAppCategory(o["app-category"], d, "app_category", sv)); err != nil {
				if !fortiAPIPatch(o["app-category"]) {
					return fmt.Errorf("Error reading app_category: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("url_category", flattenFirewallPolicyUrlCategory(o["url-category"], d, "url_category", sv)); err != nil {
			if !fortiAPIPatch(o["url-category"]) {
				return fmt.Errorf("Error reading url_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_category"); ok {
			if err = d.Set("url_category", flattenFirewallPolicyUrlCategory(o["url-category"], d, "url_category", sv)); err != nil {
				if !fortiAPIPatch(o["url-category"]) {
					return fmt.Errorf("Error reading url_category: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("app_group", flattenFirewallPolicyAppGroup(o["app-group"], d, "app_group", sv)); err != nil {
			if !fortiAPIPatch(o["app-group"]) {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_group"); ok {
			if err = d.Set("app_group", flattenFirewallPolicyAppGroup(o["app-group"], d, "app_group", sv)); err != nil {
				if !fortiAPIPatch(o["app-group"]) {
					return fmt.Errorf("Error reading app_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("nat", flattenFirewallPolicyNat(o["nat"], d, "nat", sv)); err != nil {
		if !fortiAPIPatch(o["nat"]) {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("pcp_outbound", flattenFirewallPolicyPcpOutbound(o["pcp-outbound"], d, "pcp_outbound", sv)); err != nil {
		if !fortiAPIPatch(o["pcp-outbound"]) {
			return fmt.Errorf("Error reading pcp_outbound: %v", err)
		}
	}

	if err = d.Set("pcp_inbound", flattenFirewallPolicyPcpInbound(o["pcp-inbound"], d, "pcp_inbound", sv)); err != nil {
		if !fortiAPIPatch(o["pcp-inbound"]) {
			return fmt.Errorf("Error reading pcp_inbound: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("pcp_poolname", flattenFirewallPolicyPcpPoolname(o["pcp-poolname"], d, "pcp_poolname", sv)); err != nil {
			if !fortiAPIPatch(o["pcp-poolname"]) {
				return fmt.Errorf("Error reading pcp_poolname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pcp_poolname"); ok {
			if err = d.Set("pcp_poolname", flattenFirewallPolicyPcpPoolname(o["pcp-poolname"], d, "pcp_poolname", sv)); err != nil {
				if !fortiAPIPatch(o["pcp-poolname"]) {
					return fmt.Errorf("Error reading pcp_poolname: %v", err)
				}
			}
		}
	}

	if err = d.Set("permit_any_host", flattenFirewallPolicyPermitAnyHost(o["permit-any-host"], d, "permit_any_host", sv)); err != nil {
		if !fortiAPIPatch(o["permit-any-host"]) {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("permit_stun_host", flattenFirewallPolicyPermitStunHost(o["permit-stun-host"], d, "permit_stun_host", sv)); err != nil {
		if !fortiAPIPatch(o["permit-stun-host"]) {
			return fmt.Errorf("Error reading permit_stun_host: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenFirewallPolicyFixedport(o["fixedport"], d, "fixedport", sv)); err != nil {
		if !fortiAPIPatch(o["fixedport"]) {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("port_preserve", flattenFirewallPolicyPortPreserve(o["port-preserve"], d, "port_preserve", sv)); err != nil {
		if !fortiAPIPatch(o["port-preserve"]) {
			return fmt.Errorf("Error reading port_preserve: %v", err)
		}
	}

	if err = d.Set("port_random", flattenFirewallPolicyPortRandom(o["port-random"], d, "port_random", sv)); err != nil {
		if !fortiAPIPatch(o["port-random"]) {
			return fmt.Errorf("Error reading port_random: %v", err)
		}
	}

	if err = d.Set("ippool", flattenFirewallPolicyIppool(o["ippool"], d, "ippool", sv)); err != nil {
		if !fortiAPIPatch(o["ippool"]) {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("poolname", flattenFirewallPolicyPoolname(o["poolname"], d, "poolname", sv)); err != nil {
			if !fortiAPIPatch(o["poolname"]) {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("poolname"); ok {
			if err = d.Set("poolname", flattenFirewallPolicyPoolname(o["poolname"], d, "poolname", sv)); err != nil {
				if !fortiAPIPatch(o["poolname"]) {
					return fmt.Errorf("Error reading poolname: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("poolname6", flattenFirewallPolicyPoolname6(o["poolname6"], d, "poolname6", sv)); err != nil {
			if !fortiAPIPatch(o["poolname6"]) {
				return fmt.Errorf("Error reading poolname6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("poolname6"); ok {
			if err = d.Set("poolname6", flattenFirewallPolicyPoolname6(o["poolname6"], d, "poolname6", sv)); err != nil {
				if !fortiAPIPatch(o["poolname6"]) {
					return fmt.Errorf("Error reading poolname6: %v", err)
				}
			}
		}
	}

	if err = d.Set("session_ttl", flattenFirewallPolicySessionTtl(o["session-ttl"], d, "session_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["session-ttl"]) {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("vlan_cos_fwd", flattenFirewallPolicyVlanCosFwd(o["vlan-cos-fwd"], d, "vlan_cos_fwd", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-cos-fwd"]) {
			return fmt.Errorf("Error reading vlan_cos_fwd: %v", err)
		}
	}

	if err = d.Set("vlan_cos_rev", flattenFirewallPolicyVlanCosRev(o["vlan-cos-rev"], d, "vlan_cos_rev", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-cos-rev"]) {
			return fmt.Errorf("Error reading vlan_cos_rev: %v", err)
		}
	}

	if err = d.Set("inbound", flattenFirewallPolicyInbound(o["inbound"], d, "inbound", sv)); err != nil {
		if !fortiAPIPatch(o["inbound"]) {
			return fmt.Errorf("Error reading inbound: %v", err)
		}
	}

	if err = d.Set("outbound", flattenFirewallPolicyOutbound(o["outbound"], d, "outbound", sv)); err != nil {
		if !fortiAPIPatch(o["outbound"]) {
			return fmt.Errorf("Error reading outbound: %v", err)
		}
	}

	if err = d.Set("natinbound", flattenFirewallPolicyNatinbound(o["natinbound"], d, "natinbound", sv)); err != nil {
		if !fortiAPIPatch(o["natinbound"]) {
			return fmt.Errorf("Error reading natinbound: %v", err)
		}
	}

	if err = d.Set("natoutbound", flattenFirewallPolicyNatoutbound(o["natoutbound"], d, "natoutbound", sv)); err != nil {
		if !fortiAPIPatch(o["natoutbound"]) {
			return fmt.Errorf("Error reading natoutbound: %v", err)
		}
	}

	if err = d.Set("fec", flattenFirewallPolicyFec(o["fec"], d, "fec", sv)); err != nil {
		if !fortiAPIPatch(o["fec"]) {
			return fmt.Errorf("Error reading fec: %v", err)
		}
	}

	if err = d.Set("wccp", flattenFirewallPolicyWccp(o["wccp"], d, "wccp", sv)); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("ntlm", flattenFirewallPolicyNtlm(o["ntlm"], d, "ntlm", sv)); err != nil {
		if !fortiAPIPatch(o["ntlm"]) {
			return fmt.Errorf("Error reading ntlm: %v", err)
		}
	}

	if err = d.Set("ntlm_guest", flattenFirewallPolicyNtlmGuest(o["ntlm-guest"], d, "ntlm_guest", sv)); err != nil {
		if !fortiAPIPatch(o["ntlm-guest"]) {
			return fmt.Errorf("Error reading ntlm_guest: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ntlm_enabled_browsers", flattenFirewallPolicyNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers", sv)); err != nil {
			if !fortiAPIPatch(o["ntlm-enabled-browsers"]) {
				return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ntlm_enabled_browsers"); ok {
			if err = d.Set("ntlm_enabled_browsers", flattenFirewallPolicyNtlmEnabledBrowsers(o["ntlm-enabled-browsers"], d, "ntlm_enabled_browsers", sv)); err != nil {
				if !fortiAPIPatch(o["ntlm-enabled-browsers"]) {
					return fmt.Errorf("Error reading ntlm_enabled_browsers: %v", err)
				}
			}
		}
	}

	if err = d.Set("fsso", flattenFirewallPolicyFsso(o["fsso"], d, "fsso", sv)); err != nil {
		if !fortiAPIPatch(o["fsso"]) {
			return fmt.Errorf("Error reading fsso: %v", err)
		}
	}

	if err = d.Set("wsso", flattenFirewallPolicyWsso(o["wsso"], d, "wsso", sv)); err != nil {
		if !fortiAPIPatch(o["wsso"]) {
			return fmt.Errorf("Error reading wsso: %v", err)
		}
	}

	if err = d.Set("rsso", flattenFirewallPolicyRsso(o["rsso"], d, "rsso", sv)); err != nil {
		if !fortiAPIPatch(o["rsso"]) {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("fsso_agent_for_ntlm", flattenFirewallPolicyFssoAgentForNtlm(o["fsso-agent-for-ntlm"], d, "fsso_agent_for_ntlm", sv)); err != nil {
		if !fortiAPIPatch(o["fsso-agent-for-ntlm"]) {
			return fmt.Errorf("Error reading fsso_agent_for_ntlm: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("groups", flattenFirewallPolicyGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenFirewallPolicyGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("users", flattenFirewallPolicyUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenFirewallPolicyUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("fsso_groups", flattenFirewallPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
			if !fortiAPIPatch(o["fsso-groups"]) {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fsso_groups"); ok {
			if err = d.Set("fsso_groups", flattenFirewallPolicyFssoGroups(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
				if !fortiAPIPatch(o["fsso-groups"]) {
					return fmt.Errorf("Error reading fsso_groups: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("devices", flattenFirewallPolicyDevices(o["devices"], d, "devices", sv)); err != nil {
			if !fortiAPIPatch(o["devices"]) {
				return fmt.Errorf("Error reading devices: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("devices"); ok {
			if err = d.Set("devices", flattenFirewallPolicyDevices(o["devices"], d, "devices", sv)); err != nil {
				if !fortiAPIPatch(o["devices"]) {
					return fmt.Errorf("Error reading devices: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth_path", flattenFirewallPolicyAuthPath(o["auth-path"], d, "auth_path", sv)); err != nil {
		if !fortiAPIPatch(o["auth-path"]) {
			return fmt.Errorf("Error reading auth_path: %v", err)
		}
	}

	if err = d.Set("disclaimer", flattenFirewallPolicyDisclaimer(o["disclaimer"], d, "disclaimer", sv)); err != nil {
		if !fortiAPIPatch(o["disclaimer"]) {
			return fmt.Errorf("Error reading disclaimer: %v", err)
		}
	}

	if err = d.Set("email_collect", flattenFirewallPolicyEmailCollect(o["email-collect"], d, "email_collect", sv)); err != nil {
		if !fortiAPIPatch(o["email-collect"]) {
			return fmt.Errorf("Error reading email_collect: %v", err)
		}
	}

	if err = d.Set("vpntunnel", flattenFirewallPolicyVpntunnel(o["vpntunnel"], d, "vpntunnel", sv)); err != nil {
		if !fortiAPIPatch(o["vpntunnel"]) {
			return fmt.Errorf("Error reading vpntunnel: %v", err)
		}
	}

	if err = d.Set("natip", flattenFirewallPolicyNatip(o["natip"], d, "natip", sv)); err != nil {
		if !fortiAPIPatch(o["natip"]) {
			return fmt.Errorf("Error reading natip: %v", err)
		}
	}

	if err = d.Set("match_vip", flattenFirewallPolicyMatchVip(o["match-vip"], d, "match_vip", sv)); err != nil {
		if !fortiAPIPatch(o["match-vip"]) {
			return fmt.Errorf("Error reading match_vip: %v", err)
		}
	}

	if err = d.Set("match_vip_only", flattenFirewallPolicyMatchVipOnly(o["match-vip-only"], d, "match_vip_only", sv)); err != nil {
		if !fortiAPIPatch(o["match-vip-only"]) {
			return fmt.Errorf("Error reading match_vip_only: %v", err)
		}
	}

	if err = d.Set("diffserv_copy", flattenFirewallPolicyDiffservCopy(o["diffserv-copy"], d, "diffserv_copy", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-copy"]) {
			return fmt.Errorf("Error reading diffserv_copy: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenFirewallPolicyDiffservForward(o["diffserv-forward"], d, "diffserv_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-forward"]) {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenFirewallPolicyDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse", sv)); err != nil {
		if !fortiAPIPatch(o["diffserv-reverse"]) {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenFirewallPolicyDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-forward"]) {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenFirewallPolicyDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev", sv)); err != nil {
		if !fortiAPIPatch(o["diffservcode-rev"]) {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenFirewallPolicyTcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-sender"]) {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenFirewallPolicyTcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-mss-receiver"]) {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallPolicyComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallPolicyLabel(o["label"], d, "label", sv)); err != nil {
		if !fortiAPIPatch(o["label"]) {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("global_label", flattenFirewallPolicyGlobalLabel(o["global-label"], d, "global_label", sv)); err != nil {
		if !fortiAPIPatch(o["global-label"]) {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenFirewallPolicyAuthCert(o["auth-cert"], d, "auth_cert", sv)); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_redirect_addr", flattenFirewallPolicyAuthRedirectAddr(o["auth-redirect-addr"], d, "auth_redirect_addr", sv)); err != nil {
		if !fortiAPIPatch(o["auth-redirect-addr"]) {
			return fmt.Errorf("Error reading auth_redirect_addr: %v", err)
		}
	}

	if err = d.Set("redirect_url", flattenFirewallPolicyRedirectUrl(o["redirect-url"], d, "redirect_url", sv)); err != nil {
		if !fortiAPIPatch(o["redirect-url"]) {
			return fmt.Errorf("Error reading redirect_url: %v", err)
		}
	}

	if err = d.Set("identity_based_route", flattenFirewallPolicyIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route", sv)); err != nil {
		if !fortiAPIPatch(o["identity-based-route"]) {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	if err = d.Set("block_notification", flattenFirewallPolicyBlockNotification(o["block-notification"], d, "block_notification", sv)); err != nil {
		if !fortiAPIPatch(o["block-notification"]) {
			return fmt.Errorf("Error reading block_notification: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("custom_log_fields", flattenFirewallPolicyCustomLogFields(o["custom-log-fields"], d, "custom_log_fields", sv)); err != nil {
			if !fortiAPIPatch(o["custom-log-fields"]) {
				return fmt.Errorf("Error reading custom_log_fields: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_log_fields"); ok {
			if err = d.Set("custom_log_fields", flattenFirewallPolicyCustomLogFields(o["custom-log-fields"], d, "custom_log_fields", sv)); err != nil {
				if !fortiAPIPatch(o["custom-log-fields"]) {
					return fmt.Errorf("Error reading custom_log_fields: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_override_group", flattenFirewallPolicyReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallPolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["srcaddr-negate"]) {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr6_negate", flattenFirewallPolicySrcaddr6Negate(o["srcaddr6-negate"], d, "srcaddr6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["srcaddr6-negate"]) {
			return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallPolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr-negate"]) {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr6_negate", flattenFirewallPolicyDstaddr6Negate(o["dstaddr6-negate"], d, "dstaddr6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr6-negate"]) {
			return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
		}
	}

	if err = d.Set("ztna_ems_tag_negate", flattenFirewallPolicyZtnaEmsTagNegate(o["ztna-ems-tag-negate"], d, "ztna_ems_tag_negate", sv)); err != nil {
		if !fortiAPIPatch(o["ztna-ems-tag-negate"]) {
			return fmt.Errorf("Error reading ztna_ems_tag_negate: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenFirewallPolicyServiceNegate(o["service-negate"], d, "service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["service-negate"]) {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallPolicyInternetServiceNegate(o["internet-service-negate"], d, "internet_service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-negate"]) {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallPolicyInternetServiceSrcNegate(o["internet-service-src-negate"], d, "internet_service_src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-src-negate"]) {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6_negate", flattenFirewallPolicyInternetService6Negate(o["internet-service6-negate"], d, "internet_service6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6-negate"]) {
			return fmt.Errorf("Error reading internet_service6_negate: %v", err)
		}
	}

	if err = d.Set("internet_service6_src_negate", flattenFirewallPolicyInternetService6SrcNegate(o["internet-service6-src-negate"], d, "internet_service6_src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6-src-negate"]) {
			return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
		}
	}

	if err = d.Set("timeout_send_rst", flattenFirewallPolicyTimeoutSendRst(o["timeout-send-rst"], d, "timeout_send_rst", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-send-rst"]) {
			return fmt.Errorf("Error reading timeout_send_rst: %v", err)
		}
	}

	if err = d.Set("captive_portal_exempt", flattenFirewallPolicyCaptivePortalExempt(o["captive-portal-exempt"], d, "captive_portal_exempt", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-exempt"]) {
			return fmt.Errorf("Error reading captive_portal_exempt: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenFirewallPolicyDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("ssl_mirror", flattenFirewallPolicySslMirror(o["ssl-mirror"], d, "ssl_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-mirror"]) {
			return fmt.Errorf("Error reading ssl_mirror: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ssl_mirror_intf", flattenFirewallPolicySslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf", sv)); err != nil {
			if !fortiAPIPatch(o["ssl-mirror-intf"]) {
				return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ssl_mirror_intf"); ok {
			if err = d.Set("ssl_mirror_intf", flattenFirewallPolicySslMirrorIntf(o["ssl-mirror-intf"], d, "ssl_mirror_intf", sv)); err != nil {
				if !fortiAPIPatch(o["ssl-mirror-intf"]) {
					return fmt.Errorf("Error reading ssl_mirror_intf: %v", err)
				}
			}
		}
	}

	if err = d.Set("scan_botnet_connections", flattenFirewallPolicyScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections", sv)); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallPolicyDsri(o["dsri"], d, "dsri", sv)); err != nil {
		if !fortiAPIPatch(o["dsri"]) {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_bypass", flattenFirewallPolicyRadiusMacAuthBypass(o["radius-mac-auth-bypass"], d, "radius_mac_auth_bypass", sv)); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth-bypass"]) {
			return fmt.Errorf("Error reading radius_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("radius_ip_auth_bypass", flattenFirewallPolicyRadiusIpAuthBypass(o["radius-ip-auth-bypass"], d, "radius_ip_auth_bypass", sv)); err != nil {
		if !fortiAPIPatch(o["radius-ip-auth-bypass"]) {
			return fmt.Errorf("Error reading radius_ip_auth_bypass: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", flattenFirewallPolicyDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session", sv)); err != nil {
		if !fortiAPIPatch(o["delay-tcp-npu-session"]) {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenFirewallPolicyVlanFilter(o["vlan-filter"], d, "vlan_filter", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-filter"]) {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	if err = d.Set("sgt_check", flattenFirewallPolicySgtCheck(o["sgt-check"], d, "sgt_check", sv)); err != nil {
		if !fortiAPIPatch(o["sgt-check"]) {
			return fmt.Errorf("Error reading sgt_check: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("sgt", flattenFirewallPolicySgt(o["sgt"], d, "sgt", sv)); err != nil {
			if !fortiAPIPatch(o["sgt"]) {
				return fmt.Errorf("Error reading sgt: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sgt"); ok {
			if err = d.Set("sgt", flattenFirewallPolicySgt(o["sgt"], d, "sgt", sv)); err != nil {
				if !fortiAPIPatch(o["sgt"]) {
					return fmt.Errorf("Error reading sgt: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicySrcintfName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySrcintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyDstintfName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDstintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicySrcaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyDstaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicySrcaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyDstaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDstaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaDeviceOwnership(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaEmsTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyZtnaEmsTagName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyZtnaEmsTagName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaEmsTagSecondary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyZtnaEmsTagSecondaryName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyZtnaEmsTagSecondaryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaTagsMatchLogic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaGeoTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyZtnaGeoTagName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyZtnaGeoTagName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceNameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicyInternetServiceIdId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceCustomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNetworkServiceDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyNetworkServiceDynamicName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyNetworkServiceDynamicName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceCustomGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceSrcNameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicyInternetServiceSrcIdId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceSrcGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceSrcCustomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNetworkServiceSrcDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyNetworkServiceSrcDynamicName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyNetworkServiceSrcDynamicName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetServiceSrcCustomGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetServiceSrcCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationMinimum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcVendorMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicySrcVendorMacId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySrcVendorMacId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6NameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6NameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Group(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6GroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6GroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Custom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6CustomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6CustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6CustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6CustomGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6CustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Src(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6SrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6SrcNameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6SrcNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6SrcGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6SrcGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6SrcGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6SrcCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6SrcCustomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6SrcCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6SrcCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyInternetService6SrcCustomGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyInternetService6SrcCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationMinimum6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReputationDirection6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRtpNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRtpAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyRtpAddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyRtpAddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLearningMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat64(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat46(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySendDenyPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyScheduleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyExpiryDate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPolicyExpiryDateUtc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyServiceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTosNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAntiReplay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpSessionWithoutSyn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGeoipAnycast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGeoipMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDynamicShaping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPassiveWanHealthMeasurement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAppMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUtmStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInspectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyHttpPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySshPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaPolicyRedirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebproxyProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDnsfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDlpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySpamfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFileFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVoipProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIpsVoipFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySctpFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiameterFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVirtualPatchProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIcapProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCifsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVideofilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWafProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySshFilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCasbProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyProfileProtocolOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslSshProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogtrafficStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLogHttpTransaction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCapturePacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNpAcceleration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanopt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptPassiveOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWanoptPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebcache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebcacheHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWebproxyForwardServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPerIpShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicyApplicationId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicyAppCategoryId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyAppCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUrlCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicyUrlCategoryId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyUrlCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAppGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyAppGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyAppGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPcpOutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPcpInbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPcpPoolname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyPcpPoolnameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyPcpPoolnameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPermitAnyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPermitStunHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFixedport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPortPreserve(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPortRandom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIppool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPoolname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyPoolnameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyPoolnameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyPoolname6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyPoolname6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyPoolname6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySessionTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandFirewallPolicyVlanCosFwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanCosRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyOutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatinbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatoutbound(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWccp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlmGuest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNtlmEnabledBrowsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["user-agent-string"], _ = expandFirewallPolicyNtlmEnabledBrowsersUserAgentString(d, i["user_agent_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyNtlmEnabledBrowsersUserAgentString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFsso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyWsso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRsso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFssoAgentForNtlm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyGroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyUsersName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyFssoGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicyFssoGroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyFssoGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDevices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandFirewallPolicyDevicesName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyDevicesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDisclaimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyEmailCollect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVpntunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyNatip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyMatchVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyMatchVipOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservCopy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpMssSender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTcpMssReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyLabel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyAuthRedirectAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRedirectUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyIdentityBasedRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyBlockNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCustomLogFields(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["field-id"], _ = expandFirewallPolicyCustomLogFieldsFieldId(d, i["field_id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicyCustomLogFieldsFieldId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySrcaddr6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDstaddr6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyZtnaEmsTagNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6Negate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyInternetService6SrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyTimeoutSendRst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyCaptivePortalExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySslMirrorIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallPolicySslMirrorIntfName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySslMirrorIntfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDsri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRadiusMacAuthBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyRadiusIpAuthBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyDelayTcpNpuSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicyVlanFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySgtCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallPolicySgt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandFirewallPolicySgtId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallPolicySgtId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("policyid"); ok {
		t, err := expandFirewallPolicyPolicyid(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallPolicyUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallPolicyDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallPolicySrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallPolicyDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("ztna_status"); ok {
		t, err := expandFirewallPolicyZtnaStatus(d, v, "ztna_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-status"] = t
		}
	}

	if v, ok := d.GetOk("ztna_device_ownership"); ok {
		t, err := expandFirewallPolicyZtnaDeviceOwnership(d, v, "ztna_device_ownership", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-device-ownership"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag"); ok || d.HasChange("ztna_ems_tag") {
		t, err := expandFirewallPolicyZtnaEmsTag(d, v, "ztna_ems_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag_secondary"); ok || d.HasChange("ztna_ems_tag_secondary") {
		t, err := expandFirewallPolicyZtnaEmsTagSecondary(d, v, "ztna_ems_tag_secondary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag-secondary"] = t
		}
	}

	if v, ok := d.GetOk("ztna_tags_match_logic"); ok {
		t, err := expandFirewallPolicyZtnaTagsMatchLogic(d, v, "ztna_tags_match_logic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-tags-match-logic"] = t
		}
	}

	if v, ok := d.GetOk("ztna_geo_tag"); ok || d.HasChange("ztna_geo_tag") {
		t, err := expandFirewallPolicyZtnaGeoTag(d, v, "ztna_geo_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-geo-tag"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {
		t, err := expandFirewallPolicyInternetService(d, v, "internet_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandFirewallPolicyInternetServiceName(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandFirewallPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandFirewallPolicyInternetServiceGroup(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandFirewallPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("network_service_dynamic"); ok || d.HasChange("network_service_dynamic") {
		t, err := expandFirewallPolicyNetworkServiceDynamic(d, v, "network_service_dynamic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-service-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandFirewallPolicyInternetServiceCustomGroup(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok {
		t, err := expandFirewallPolicyInternetServiceSrc(d, v, "internet_service_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandFirewallPolicyInternetServiceSrcName(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok || d.HasChange("internet_service_src_id") {
		t, err := expandFirewallPolicyInternetServiceSrcId(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandFirewallPolicyInternetServiceSrcGroup(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandFirewallPolicyInternetServiceSrcCustom(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("network_service_src_dynamic"); ok || d.HasChange("network_service_src_dynamic") {
		t, err := expandFirewallPolicyNetworkServiceSrcDynamic(d, v, "network_service_src_dynamic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-service-src-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandFirewallPolicyInternetServiceSrcCustomGroup(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOkExists("reputation_minimum"); ok {
		t, err := expandFirewallPolicyReputationMinimum(d, v, "reputation_minimum", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum"] = t
		}
	} else if d.HasChange("reputation_minimum") {
		obj["reputation-minimum"] = nil
	}

	if v, ok := d.GetOk("reputation_direction"); ok {
		t, err := expandFirewallPolicyReputationDirection(d, v, "reputation_direction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction"] = t
		}
	}

	if v, ok := d.GetOk("src_vendor_mac"); ok || d.HasChange("src_vendor_mac") {
		t, err := expandFirewallPolicySrcVendorMac(d, v, "src_vendor_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-vendor-mac"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6"); ok {
		t, err := expandFirewallPolicyInternetService6(d, v, "internet_service6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_name"); ok || d.HasChange("internet_service6_name") {
		t, err := expandFirewallPolicyInternetService6Name(d, v, "internet_service6_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_group"); ok || d.HasChange("internet_service6_group") {
		t, err := expandFirewallPolicyInternetService6Group(d, v, "internet_service6_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom"); ok || d.HasChange("internet_service6_custom") {
		t, err := expandFirewallPolicyInternetService6Custom(d, v, "internet_service6_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom_group"); ok || d.HasChange("internet_service6_custom_group") {
		t, err := expandFirewallPolicyInternetService6CustomGroup(d, v, "internet_service6_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src"); ok {
		t, err := expandFirewallPolicyInternetService6Src(d, v, "internet_service6_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_name"); ok || d.HasChange("internet_service6_src_name") {
		t, err := expandFirewallPolicyInternetService6SrcName(d, v, "internet_service6_src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_group"); ok || d.HasChange("internet_service6_src_group") {
		t, err := expandFirewallPolicyInternetService6SrcGroup(d, v, "internet_service6_src_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom"); ok || d.HasChange("internet_service6_src_custom") {
		t, err := expandFirewallPolicyInternetService6SrcCustom(d, v, "internet_service6_src_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom_group"); ok || d.HasChange("internet_service6_src_custom_group") {
		t, err := expandFirewallPolicyInternetService6SrcCustomGroup(d, v, "internet_service6_src_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOkExists("reputation_minimum6"); ok {
		t, err := expandFirewallPolicyReputationMinimum6(d, v, "reputation_minimum6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-minimum6"] = t
		}
	} else if d.HasChange("reputation_minimum6") {
		obj["reputation-minimum6"] = nil
	}

	if v, ok := d.GetOk("reputation_direction6"); ok {
		t, err := expandFirewallPolicyReputationDirection6(d, v, "reputation_direction6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation-direction6"] = t
		}
	}

	if v, ok := d.GetOk("rtp_nat"); ok {
		t, err := expandFirewallPolicyRtpNat(d, v, "rtp_nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-nat"] = t
		}
	}

	if v, ok := d.GetOk("rtp_addr"); ok || d.HasChange("rtp_addr") {
		t, err := expandFirewallPolicyRtpAddr(d, v, "rtp_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtp-addr"] = t
		}
	}

	if v, ok := d.GetOk("learning_mode"); ok {
		t, err := expandFirewallPolicyLearningMode(d, v, "learning_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-mode"] = t
		}
	} else if d.HasChange("learning_mode") {
		obj["learning-mode"] = nil
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandFirewallPolicyAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok {
		t, err := expandFirewallPolicyNat64(d, v, "nat64", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok {
		t, err := expandFirewallPolicyNat46(d, v, "nat46", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok {
		t, err := expandFirewallPolicySendDenyPacket(d, v, "send_deny_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok {
		t, err := expandFirewallPolicyFirewallSessionDirty(d, v, "firewall_session_dirty", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandFirewallPolicySchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	} else if d.HasChange("schedule") {
		obj["schedule"] = nil
	}

	if v, ok := d.GetOk("schedule_timeout"); ok {
		t, err := expandFirewallPolicyScheduleTimeout(d, v, "schedule_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-timeout"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry"); ok {
		t, err := expandFirewallPolicyPolicyExpiry(d, v, "policy_expiry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry_date"); ok {
		t, err := expandFirewallPolicyPolicyExpiryDate(d, v, "policy_expiry_date", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry-date"] = t
		}
	}

	if v, ok := d.GetOk("policy_expiry_date_utc"); ok {
		t, err := expandFirewallPolicyPolicyExpiryDateUtc(d, v, "policy_expiry_date_utc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-expiry-date-utc"] = t
		}
	} else if d.HasChange("policy_expiry_date_utc") {
		obj["policy-expiry-date-utc"] = nil
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {
		t, err := expandFirewallPolicyTos(d, v, "tos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {
		t, err := expandFirewallPolicyTosMask(d, v, "tos_mask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("tos_negate"); ok {
		t, err := expandFirewallPolicyTosNegate(d, v, "tos_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-negate"] = t
		}
	}

	if v, ok := d.GetOk("anti_replay"); ok {
		t, err := expandFirewallPolicyAntiReplay(d, v, "anti_replay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anti-replay"] = t
		}
	}

	if v, ok := d.GetOk("tcp_session_without_syn"); ok {
		t, err := expandFirewallPolicyTcpSessionWithoutSyn(d, v, "tcp_session_without_syn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-session-without-syn"] = t
		}
	}

	if v, ok := d.GetOk("geoip_anycast"); ok {
		t, err := expandFirewallPolicyGeoipAnycast(d, v, "geoip_anycast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip-anycast"] = t
		}
	}

	if v, ok := d.GetOk("geoip_match"); ok {
		t, err := expandFirewallPolicyGeoipMatch(d, v, "geoip_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geoip-match"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_shaping"); ok {
		t, err := expandFirewallPolicyDynamicShaping(d, v, "dynamic_shaping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-shaping"] = t
		}
	}

	if v, ok := d.GetOk("passive_wan_health_measurement"); ok {
		t, err := expandFirewallPolicyPassiveWanHealthMeasurement(d, v, "passive_wan_health_measurement", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-wan-health-measurement"] = t
		}
	}

	if v, ok := d.GetOk("app_monitor"); ok {
		t, err := expandFirewallPolicyAppMonitor(d, v, "app_monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-monitor"] = t
		}
	}

	if v, ok := d.GetOk("utm_status"); ok {
		t, err := expandFirewallPolicyUtmStatus(d, v, "utm_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("inspection_mode"); ok {
		t, err := expandFirewallPolicyInspectionMode(d, v, "inspection_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inspection-mode"] = t
		}
	}

	if v, ok := d.GetOk("http_policy_redirect"); ok {
		t, err := expandFirewallPolicyHttpPolicyRedirect(d, v, "http_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ssh_policy_redirect"); ok {
		t, err := expandFirewallPolicySshPolicyRedirect(d, v, "ssh_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("ztna_policy_redirect"); ok {
		t, err := expandFirewallPolicyZtnaPolicyRedirect(d, v, "ztna_policy_redirect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-policy-redirect"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_profile"); ok {
		t, err := expandFirewallPolicyWebproxyProfile(d, v, "webproxy_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-profile"] = t
		}
	} else if d.HasChange("webproxy_profile") {
		obj["webproxy-profile"] = nil
	}

	if v, ok := d.GetOk("profile_type"); ok {
		t, err := expandFirewallPolicyProfileType(d, v, "profile_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok {
		t, err := expandFirewallPolicyProfileGroup(d, v, "profile_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	} else if d.HasChange("profile_group") {
		obj["profile-group"] = nil
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandFirewallPolicyAvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	} else if d.HasChange("av_profile") {
		obj["av-profile"] = nil
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandFirewallPolicyWebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	} else if d.HasChange("webfilter_profile") {
		obj["webfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {
		t, err := expandFirewallPolicyDnsfilterProfile(d, v, "dnsfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	} else if d.HasChange("dnsfilter_profile") {
		obj["dnsfilter-profile"] = nil
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {
		t, err := expandFirewallPolicyEmailfilterProfile(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	} else if d.HasChange("emailfilter_profile") {
		obj["emailfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_profile"); ok {
		t, err := expandFirewallPolicyDlpProfile(d, v, "dlp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	} else if d.HasChange("dlp_profile") {
		obj["dlp-profile"] = nil
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {
		t, err := expandFirewallPolicySpamfilterProfile(d, v, "spamfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	} else if d.HasChange("spamfilter_profile") {
		obj["spamfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {
		t, err := expandFirewallPolicyDlpSensor(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	} else if d.HasChange("dlp_sensor") {
		obj["dlp-sensor"] = nil
	}

	if v, ok := d.GetOk("file_filter_profile"); ok {
		t, err := expandFirewallPolicyFileFilterProfile(d, v, "file_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	} else if d.HasChange("file_filter_profile") {
		obj["file-filter-profile"] = nil
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandFirewallPolicyIpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	} else if d.HasChange("ips_sensor") {
		obj["ips-sensor"] = nil
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandFirewallPolicyApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	} else if d.HasChange("application_list") {
		obj["application-list"] = nil
	}

	if v, ok := d.GetOk("voip_profile"); ok {
		t, err := expandFirewallPolicyVoipProfile(d, v, "voip_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	} else if d.HasChange("voip_profile") {
		obj["voip-profile"] = nil
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok {
		t, err := expandFirewallPolicyIpsVoipFilter(d, v, "ips_voip_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	} else if d.HasChange("ips_voip_filter") {
		obj["ips-voip-filter"] = nil
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok {
		t, err := expandFirewallPolicySctpFilterProfile(d, v, "sctp_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	} else if d.HasChange("sctp_filter_profile") {
		obj["sctp-filter-profile"] = nil
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok {
		t, err := expandFirewallPolicyDiameterFilterProfile(d, v, "diameter_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	} else if d.HasChange("diameter_filter_profile") {
		obj["diameter-filter-profile"] = nil
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok {
		t, err := expandFirewallPolicyVirtualPatchProfile(d, v, "virtual_patch_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	} else if d.HasChange("virtual_patch_profile") {
		obj["virtual-patch-profile"] = nil
	}

	if v, ok := d.GetOk("icap_profile"); ok {
		t, err := expandFirewallPolicyIcapProfile(d, v, "icap_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	} else if d.HasChange("icap_profile") {
		obj["icap-profile"] = nil
	}

	if v, ok := d.GetOk("cifs_profile"); ok {
		t, err := expandFirewallPolicyCifsProfile(d, v, "cifs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	} else if d.HasChange("cifs_profile") {
		obj["cifs-profile"] = nil
	}

	if v, ok := d.GetOk("videofilter_profile"); ok {
		t, err := expandFirewallPolicyVideofilterProfile(d, v, "videofilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	} else if d.HasChange("videofilter_profile") {
		obj["videofilter-profile"] = nil
	}

	if v, ok := d.GetOk("waf_profile"); ok {
		t, err := expandFirewallPolicyWafProfile(d, v, "waf_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["waf-profile"] = t
		}
	} else if d.HasChange("waf_profile") {
		obj["waf-profile"] = nil
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {
		t, err := expandFirewallPolicySshFilterProfile(d, v, "ssh_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	} else if d.HasChange("ssh_filter_profile") {
		obj["ssh-filter-profile"] = nil
	}

	if v, ok := d.GetOk("casb_profile"); ok {
		t, err := expandFirewallPolicyCasbProfile(d, v, "casb_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	} else if d.HasChange("casb_profile") {
		obj["casb-profile"] = nil
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {
		t, err := expandFirewallPolicyProfileProtocolOptions(d, v, "profile_protocol_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {
		t, err := expandFirewallPolicySslSshProfile(d, v, "ssl_ssh_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandFirewallPolicyLogtraffic(d, v, "logtraffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {
		t, err := expandFirewallPolicyLogtrafficStart(d, v, "logtraffic_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("log_http_transaction"); ok {
		t, err := expandFirewallPolicyLogHttpTransaction(d, v, "log_http_transaction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-http-transaction"] = t
		}
	}

	if v, ok := d.GetOk("capture_packet"); ok {
		t, err := expandFirewallPolicyCapturePacket(d, v, "capture_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capture-packet"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok {
		t, err := expandFirewallPolicyAutoAsicOffload(d, v, "auto_asic_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("np_acceleration"); ok {
		t, err := expandFirewallPolicyNpAcceleration(d, v, "np_acceleration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["np-acceleration"] = t
		}
	}

	if v, ok := d.GetOk("wanopt"); ok {
		t, err := expandFirewallPolicyWanopt(d, v, "wanopt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_detection"); ok {
		t, err := expandFirewallPolicyWanoptDetection(d, v, "wanopt_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-detection"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_passive_opt"); ok {
		t, err := expandFirewallPolicyWanoptPassiveOpt(d, v, "wanopt_passive_opt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-passive-opt"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_profile"); ok {
		t, err := expandFirewallPolicyWanoptProfile(d, v, "wanopt_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-profile"] = t
		}
	} else if d.HasChange("wanopt_profile") {
		obj["wanopt-profile"] = nil
	}

	if v, ok := d.GetOk("wanopt_peer"); ok {
		t, err := expandFirewallPolicyWanoptPeer(d, v, "wanopt_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-peer"] = t
		}
	} else if d.HasChange("wanopt_peer") {
		obj["wanopt-peer"] = nil
	}

	if v, ok := d.GetOk("webcache"); ok {
		t, err := expandFirewallPolicyWebcache(d, v, "webcache", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache"] = t
		}
	}

	if v, ok := d.GetOk("webcache_https"); ok {
		t, err := expandFirewallPolicyWebcacheHttps(d, v, "webcache_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webcache-https"] = t
		}
	}

	if v, ok := d.GetOk("webproxy_forward_server"); ok {
		t, err := expandFirewallPolicyWebproxyForwardServer(d, v, "webproxy_forward_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy-forward-server"] = t
		}
	} else if d.HasChange("webproxy_forward_server") {
		obj["webproxy-forward-server"] = nil
	}

	if v, ok := d.GetOk("traffic_shaper"); ok {
		t, err := expandFirewallPolicyTrafficShaper(d, v, "traffic_shaper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	} else if d.HasChange("traffic_shaper") {
		obj["traffic-shaper"] = nil
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok {
		t, err := expandFirewallPolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	} else if d.HasChange("traffic_shaper_reverse") {
		obj["traffic-shaper-reverse"] = nil
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok {
		t, err := expandFirewallPolicyPerIpShaper(d, v, "per_ip_shaper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	} else if d.HasChange("per_ip_shaper") {
		obj["per-ip-shaper"] = nil
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallPolicyApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandFirewallPolicyAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		t, err := expandFirewallPolicyUrlCategory(d, v, "url_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-category"] = t
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandFirewallPolicyAppGroup(d, v, "app_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok {
		t, err := expandFirewallPolicyNat(d, v, "nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("pcp_outbound"); ok {
		t, err := expandFirewallPolicyPcpOutbound(d, v, "pcp_outbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcp-outbound"] = t
		}
	}

	if v, ok := d.GetOk("pcp_inbound"); ok {
		t, err := expandFirewallPolicyPcpInbound(d, v, "pcp_inbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcp-inbound"] = t
		}
	}

	if v, ok := d.GetOk("pcp_poolname"); ok || d.HasChange("pcp_poolname") {
		t, err := expandFirewallPolicyPcpPoolname(d, v, "pcp_poolname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pcp-poolname"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok {
		t, err := expandFirewallPolicyPermitAnyHost(d, v, "permit_any_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("permit_stun_host"); ok {
		t, err := expandFirewallPolicyPermitStunHost(d, v, "permit_stun_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-stun-host"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok {
		t, err := expandFirewallPolicyFixedport(d, v, "fixedport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("port_preserve"); ok {
		t, err := expandFirewallPolicyPortPreserve(d, v, "port_preserve", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-preserve"] = t
		}
	}

	if v, ok := d.GetOk("port_random"); ok {
		t, err := expandFirewallPolicyPortRandom(d, v, "port_random", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-random"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok {
		t, err := expandFirewallPolicyIppool(d, v, "ippool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok || d.HasChange("poolname") {
		t, err := expandFirewallPolicyPoolname(d, v, "poolname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("poolname6"); ok || d.HasChange("poolname6") {
		t, err := expandFirewallPolicyPoolname6(d, v, "poolname6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname6"] = t
		}
	}

	if v, ok := d.GetOkExists("session_ttl"); ok {
		t, err := expandFirewallPolicySessionTtl(d, v, "session_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ttl"] = t
		}
	} else if d.HasChange("session_ttl") {
		obj["session-ttl"] = nil
	}

	if v, ok := d.GetOkExists("vlan_cos_fwd"); ok {
		t, err := expandFirewallPolicyVlanCosFwd(d, v, "vlan_cos_fwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-fwd"] = t
		}
	}

	if v, ok := d.GetOkExists("vlan_cos_rev"); ok {
		t, err := expandFirewallPolicyVlanCosRev(d, v, "vlan_cos_rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-cos-rev"] = t
		}
	}

	if v, ok := d.GetOk("inbound"); ok {
		t, err := expandFirewallPolicyInbound(d, v, "inbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound"] = t
		}
	}

	if v, ok := d.GetOk("outbound"); ok {
		t, err := expandFirewallPolicyOutbound(d, v, "outbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbound"] = t
		}
	}

	if v, ok := d.GetOk("natinbound"); ok {
		t, err := expandFirewallPolicyNatinbound(d, v, "natinbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natinbound"] = t
		}
	}

	if v, ok := d.GetOk("natoutbound"); ok {
		t, err := expandFirewallPolicyNatoutbound(d, v, "natoutbound", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natoutbound"] = t
		}
	}

	if v, ok := d.GetOk("fec"); ok {
		t, err := expandFirewallPolicyFec(d, v, "fec", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok {
		t, err := expandFirewallPolicyWccp(d, v, "wccp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("ntlm"); ok {
		t, err := expandFirewallPolicyNtlm(d, v, "ntlm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_guest"); ok {
		t, err := expandFirewallPolicyNtlmGuest(d, v, "ntlm_guest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-guest"] = t
		}
	}

	if v, ok := d.GetOk("ntlm_enabled_browsers"); ok || d.HasChange("ntlm_enabled_browsers") {
		t, err := expandFirewallPolicyNtlmEnabledBrowsers(d, v, "ntlm_enabled_browsers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntlm-enabled-browsers"] = t
		}
	}

	if v, ok := d.GetOk("fsso"); ok {
		t, err := expandFirewallPolicyFsso(d, v, "fsso", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso"] = t
		}
	} else if d.HasChange("fsso") {
		obj["fsso"] = nil
	}

	if v, ok := d.GetOk("wsso"); ok {
		t, err := expandFirewallPolicyWsso(d, v, "wsso", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wsso"] = t
		}
	} else if d.HasChange("wsso") {
		obj["wsso"] = nil
	}

	if v, ok := d.GetOk("rsso"); ok {
		t, err := expandFirewallPolicyRsso(d, v, "rsso", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	} else if d.HasChange("rsso") {
		obj["rsso"] = nil
	}

	if v, ok := d.GetOk("fsso_agent_for_ntlm"); ok {
		t, err := expandFirewallPolicyFssoAgentForNtlm(d, v, "fsso_agent_for_ntlm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-agent-for-ntlm"] = t
		}
	} else if d.HasChange("fsso_agent_for_ntlm") {
		obj["fsso-agent-for-ntlm"] = nil
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandFirewallPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok || d.HasChange("fsso_groups") {
		t, err := expandFirewallPolicyFssoGroups(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	if v, ok := d.GetOk("devices"); ok || d.HasChange("devices") {
		t, err := expandFirewallPolicyDevices(d, v, "devices", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devices"] = t
		}
	}

	if v, ok := d.GetOk("auth_path"); ok {
		t, err := expandFirewallPolicyAuthPath(d, v, "auth_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-path"] = t
		}
	}

	if v, ok := d.GetOk("disclaimer"); ok {
		t, err := expandFirewallPolicyDisclaimer(d, v, "disclaimer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disclaimer"] = t
		}
	}

	if v, ok := d.GetOk("email_collect"); ok {
		t, err := expandFirewallPolicyEmailCollect(d, v, "email_collect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-collect"] = t
		}
	}

	if v, ok := d.GetOk("vpntunnel"); ok {
		t, err := expandFirewallPolicyVpntunnel(d, v, "vpntunnel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpntunnel"] = t
		}
	} else if d.HasChange("vpntunnel") {
		obj["vpntunnel"] = nil
	}

	if v, ok := d.GetOk("natip"); ok {
		t, err := expandFirewallPolicyNatip(d, v, "natip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["natip"] = t
		}
	}

	if v, ok := d.GetOk("match_vip"); ok {
		t, err := expandFirewallPolicyMatchVip(d, v, "match_vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip"] = t
		}
	}

	if v, ok := d.GetOk("match_vip_only"); ok {
		t, err := expandFirewallPolicyMatchVipOnly(d, v, "match_vip_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-vip-only"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_copy"); ok {
		t, err := expandFirewallPolicyDiffservCopy(d, v, "diffserv_copy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-copy"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok {
		t, err := expandFirewallPolicyDiffservForward(d, v, "diffserv_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok {
		t, err := expandFirewallPolicyDiffservReverse(d, v, "diffserv_reverse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok {
		t, err := expandFirewallPolicyDiffservcodeForward(d, v, "diffservcode_forward", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok {
		t, err := expandFirewallPolicyDiffservcodeRev(d, v, "diffservcode_rev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_mss_sender"); ok {
		t, err := expandFirewallPolicyTcpMssSender(d, v, "tcp_mss_sender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	} else if d.HasChange("tcp_mss_sender") {
		obj["tcp-mss-sender"] = nil
	}

	if v, ok := d.GetOkExists("tcp_mss_receiver"); ok {
		t, err := expandFirewallPolicyTcpMssReceiver(d, v, "tcp_mss_receiver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	} else if d.HasChange("tcp_mss_receiver") {
		obj["tcp-mss-receiver"] = nil
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallPolicyComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("label"); ok {
		t, err := expandFirewallPolicyLabel(d, v, "label", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	} else if d.HasChange("label") {
		obj["label"] = nil
	}

	if v, ok := d.GetOk("global_label"); ok {
		t, err := expandFirewallPolicyGlobalLabel(d, v, "global_label", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	} else if d.HasChange("global_label") {
		obj["global-label"] = nil
	}

	if v, ok := d.GetOk("auth_cert"); ok {
		t, err := expandFirewallPolicyAuthCert(d, v, "auth_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	} else if d.HasChange("auth_cert") {
		obj["auth-cert"] = nil
	}

	if v, ok := d.GetOk("auth_redirect_addr"); ok {
		t, err := expandFirewallPolicyAuthRedirectAddr(d, v, "auth_redirect_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-redirect-addr"] = t
		}
	} else if d.HasChange("auth_redirect_addr") {
		obj["auth-redirect-addr"] = nil
	}

	if v, ok := d.GetOk("redirect_url"); ok {
		t, err := expandFirewallPolicyRedirectUrl(d, v, "redirect_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-url"] = t
		}
	} else if d.HasChange("redirect_url") {
		obj["redirect-url"] = nil
	}

	if v, ok := d.GetOk("identity_based_route"); ok {
		t, err := expandFirewallPolicyIdentityBasedRoute(d, v, "identity_based_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based-route"] = t
		}
	} else if d.HasChange("identity_based_route") {
		obj["identity-based-route"] = nil
	}

	if v, ok := d.GetOk("block_notification"); ok {
		t, err := expandFirewallPolicyBlockNotification(d, v, "block_notification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-notification"] = t
		}
	}

	if v, ok := d.GetOk("custom_log_fields"); ok || d.HasChange("custom_log_fields") {
		t, err := expandFirewallPolicyCustomLogFields(d, v, "custom_log_fields", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-log-fields"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok {
		t, err := expandFirewallPolicyReplacemsgOverrideGroup(d, v, "replacemsg_override_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	} else if d.HasChange("replacemsg_override_group") {
		obj["replacemsg-override-group"] = nil
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok {
		t, err := expandFirewallPolicySrcaddrNegate(d, v, "srcaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6_negate"); ok {
		t, err := expandFirewallPolicySrcaddr6Negate(d, v, "srcaddr6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok {
		t, err := expandFirewallPolicyDstaddrNegate(d, v, "dstaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6_negate"); ok {
		t, err := expandFirewallPolicyDstaddr6Negate(d, v, "dstaddr6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("ztna_ems_tag_negate"); ok {
		t, err := expandFirewallPolicyZtnaEmsTagNegate(d, v, "ztna_ems_tag_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ztna-ems-tag-negate"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok {
		t, err := expandFirewallPolicyServiceNegate(d, v, "service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok {
		t, err := expandFirewallPolicyInternetServiceNegate(d, v, "internet_service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok {
		t, err := expandFirewallPolicyInternetServiceSrcNegate(d, v, "internet_service_src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_negate"); ok {
		t, err := expandFirewallPolicyInternetService6Negate(d, v, "internet_service6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_negate"); ok {
		t, err := expandFirewallPolicyInternetService6SrcNegate(d, v, "internet_service6_src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("timeout_send_rst"); ok {
		t, err := expandFirewallPolicyTimeoutSendRst(d, v, "timeout_send_rst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-send-rst"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_exempt"); ok {
		t, err := expandFirewallPolicyCaptivePortalExempt(d, v, "captive_portal_exempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-exempt"] = t
		}
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		t, err := expandFirewallPolicyDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	} else if d.HasChange("decrypted_traffic_mirror") {
		obj["decrypted-traffic-mirror"] = nil
	}

	if v, ok := d.GetOk("ssl_mirror"); ok {
		t, err := expandFirewallPolicySslMirror(d, v, "ssl_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror"] = t
		}
	} else if d.HasChange("ssl_mirror") {
		obj["ssl-mirror"] = nil
	}

	if v, ok := d.GetOk("ssl_mirror_intf"); ok || d.HasChange("ssl_mirror_intf") {
		t, err := expandFirewallPolicySslMirrorIntf(d, v, "ssl_mirror_intf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-mirror-intf"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandFirewallPolicyScanBotnetConnections(d, v, "scan_botnet_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	} else if d.HasChange("scan_botnet_connections") {
		obj["scan-botnet-connections"] = nil
	}

	if v, ok := d.GetOk("dsri"); ok {
		t, err := expandFirewallPolicyDsri(d, v, "dsri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_bypass"); ok {
		t, err := expandFirewallPolicyRadiusMacAuthBypass(d, v, "radius_mac_auth_bypass", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("radius_ip_auth_bypass"); ok {
		t, err := expandFirewallPolicyRadiusIpAuthBypass(d, v, "radius_ip_auth_bypass", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-ip-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("delay_tcp_npu_session"); ok {
		t, err := expandFirewallPolicyDelayTcpNpuSession(d, v, "delay_tcp_npu_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-tcp-npu-session"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok {
		t, err := expandFirewallPolicyVlanFilter(d, v, "vlan_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	} else if d.HasChange("vlan_filter") {
		obj["vlan-filter"] = nil
	}

	if v, ok := d.GetOk("sgt_check"); ok {
		t, err := expandFirewallPolicySgtCheck(d, v, "sgt_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgt-check"] = t
		}
	}

	if v, ok := d.GetOk("sgt"); ok || d.HasChange("sgt") {
		t, err := expandFirewallPolicySgt(d, v, "sgt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sgt"] = t
		}
	}

	return &obj, nil
}
