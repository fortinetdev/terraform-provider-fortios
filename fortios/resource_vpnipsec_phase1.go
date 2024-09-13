// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VPN remote gateway.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnIpsecPhase1() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase1Create,
		Read:   resourceVpnIpsecPhase1Read,
		Update: resourceVpnIpsecPhase1Update,
		Delete: resourceVpnIpsecPhase1Delete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"ike_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remotegw_ddns": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"keylife": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 172800),
				Optional:     true,
				Computed:     true,
			},
			"certificate": &schema.Schema{
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
			"authmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authmethod_remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peerid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"usrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"peergrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"mode_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode_cfg_allow_client_selector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_ip_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_ra_giaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_ra_linkaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal_domain_list": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"ipv4_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ipv4_split_include": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"split_include_service": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ipv4_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ipv6_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_prefix": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 128),
				Optional:     true,
				Computed:     true,
			},
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ipv6_split_include": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ipv6_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ip_delay_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 28800),
				Optional:     true,
			},
			"unity_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"banner": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Optional:     true,
			},
			"include_local_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_split_exclude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"ipv6_split_exclude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"save_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_gateway": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"add_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_gw_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"psksecret_remote": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"keepalive": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 900),
				Optional:     true,
				Computed:     true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"localid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"localid_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"negotiate_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"fragmentation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_retrycount": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),
				Optional:     true,
				Computed:     true,
			},
			"dpd_retryinterval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"npu_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_cert_chain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addke1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"addke7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"suite_b": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_exclude_peergrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"eap_cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk_secret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ppk_identity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"wizard_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xauthtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authusr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"authpasswd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"group_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_authentication_secret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"authusrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"mesh_selector_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shared_idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeoutinterval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 43200),
				Optional:     true,
				Computed:     true,
			},
			"ha_sync_esp_seqno": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgsp_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inbound_dscp_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nattraversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fragmentation_mtu": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 16000),
				Optional:     true,
				Computed:     true,
			},
			"childless_ike": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_ad_autoconnect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_resume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_resume_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 172800),
				Optional:     true,
				Computed:     true,
			},
			"rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digital_signature_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signature_hash_alg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsa_signature_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsa_signature_hash_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_unique_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_id_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_egress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_send_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000),
				Optional:     true,
				Computed:     true,
			},
			"fec_base": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"fec_codec_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_codec": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),
				Optional:     true,
				Computed:     true,
			},
			"fec_redundant": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"fec_ingress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fec_receive_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10000),
				Optional:     true,
				Computed:     true,
			},
			"fec_health_check": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"fec_mapping_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"network_overlay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"dev_id_notification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dev_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"loopback_asymroute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_cost": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"kms": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"exchange_fgt_device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_auto_linklocal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ems_sn_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_trust_store": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qkd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qkd_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortinet_esp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_transport_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"fallback_tcp_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"remote_gw_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw_country": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				Optional:     true,
			},
			"remote_gw_ztna_tags": &schema.Schema{
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
			"remote_gw6_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6_country": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),
				Optional:     true,
			},
			"cert_peer_username_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_peer_username_strip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceVpnIpsecPhase1Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecPhase1(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1 resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecPhase1(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase1")
	}

	return resourceVpnIpsecPhase1Read(d, m)
}

func resourceVpnIpsecPhase1Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecPhase1(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1 resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecPhase1(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase1")
	}

	return resourceVpnIpsecPhase1Read(d, m)
}

func resourceVpnIpsecPhase1Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnIpsecPhase1(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase1Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecPhase1(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase1(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1 resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase1Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Type(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IkeVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LocalGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemotegwDdns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Keylife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Certificate(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenVpnIpsecPhase1CertificateName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnIpsecPhase1CertificateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authmethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AuthmethodRemote(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peertype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peerid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Usrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peergrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ModeCfg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ModeCfgAllowClientSelector(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AssignIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AssignIpFrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4Netmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DnsMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InternalDomainList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_name"
		if cur_v, ok := i["domain-name"]; ok {
			tmp["domain_name"] = flattenVpnIpsecPhase1InternalDomainListDomainName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "domain_name", d)
	return result
}

func flattenVpnIpsecPhase1InternalDomainListDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4WinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4WinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenVpnIpsecPhase1Ipv4ExcludeRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if cur_v, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenVpnIpsecPhase1Ipv4ExcludeRangeStartIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if cur_v, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenVpnIpsecPhase1Ipv4ExcludeRangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4SplitInclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SplitIncludeService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Ipv6DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenVpnIpsecPhase1Ipv6ExcludeRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if cur_v, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenVpnIpsecPhase1Ipv6ExcludeRangeStartIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if cur_v, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenVpnIpsecPhase1Ipv6ExcludeRangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6SplitInclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IpDelayInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1UnitySupport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Domain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Banner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IncludeLocalLan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4SplitExclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6SplitExclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SavePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientAutoNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientKeepAlive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1BackupGateway(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if cur_v, ok := i["address"]; ok {
			tmp["address"] = flattenVpnIpsecPhase1BackupGatewayAddress(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "address", d)
	return result
}

func flattenVpnIpsecPhase1BackupGatewayAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Proposal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AddRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AddGwRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Keepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Distance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Priority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Localid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LocalidType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AutoNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NegotiateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Fragmentation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DpdRetrycount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1DpdRetryinterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ForticlientEnforcement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NpuOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SendCertChain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dhgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Addke7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SuiteB(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Eap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EapIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EapExcludePeergrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EapCertAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AcctVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ppk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1PpkIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1WizardType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Xauthtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Reauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authusr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1GroupAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authusrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1MeshSelectorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SharedIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IdleTimeoutinterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1HaSyncEspSeqno(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FgspSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InboundDscpCopy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Nattraversal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Esn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FragmentationMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1ChildlessIke(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AzureAdAutoconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientResume(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientResumeInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Rekey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DigitalSignatureAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SignatureHashAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RsaSignatureFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RsaSignatureHashOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EnforceUniqueId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertIdValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecEgress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecSendTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1FecBase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1FecCodecString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecCodec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1FecRedundant(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1FecIngress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecReceiveTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1FecHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FecMappingProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NetworkOverlay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1DevIdNotification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DevId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LoopbackAsymroute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LinkCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1Kms(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ExchangeFgtDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6AutoLinklocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EmsSnCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertTrustStore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Qkd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1QkdProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Transport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FortinetEsp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AutoTransportThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1FallbackTcpThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1RemoteGwMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnIpsecPhase1RemoteGwStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGwZtnaTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenVpnIpsecPhase1RemoteGwZtnaTagsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnIpsecPhase1RemoteGwZtnaTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6Match(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6Subnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw6Country(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertPeerUsernameValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertPeerUsernameStrip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase1(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenVpnIpsecPhase1Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnIpsecPhase1Type(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecPhase1Interface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ike_version", flattenVpnIpsecPhase1IkeVersion(o["ike-version"], d, "ike_version", sv)); err != nil {
		if !fortiAPIPatch(o["ike-version"]) {
			return fmt.Errorf("Error reading ike_version: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecPhase1RemoteGw(o["remote-gw"], d, "remote_gw", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecPhase1LocalGw(o["local-gw"], d, "local_gw", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("remotegw_ddns", flattenVpnIpsecPhase1RemotegwDdns(o["remotegw-ddns"], d, "remotegw_ddns", sv)); err != nil {
		if !fortiAPIPatch(o["remotegw-ddns"]) {
			return fmt.Errorf("Error reading remotegw_ddns: %v", err)
		}
	}

	if err = d.Set("keylife", flattenVpnIpsecPhase1Keylife(o["keylife"], d, "keylife", sv)); err != nil {
		if !fortiAPIPatch(o["keylife"]) {
			return fmt.Errorf("Error reading keylife: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("certificate", flattenVpnIpsecPhase1Certificate(o["certificate"], d, "certificate", sv)); err != nil {
			if !fortiAPIPatch(o["certificate"]) {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("certificate"); ok {
			if err = d.Set("certificate", flattenVpnIpsecPhase1Certificate(o["certificate"], d, "certificate", sv)); err != nil {
				if !fortiAPIPatch(o["certificate"]) {
					return fmt.Errorf("Error reading certificate: %v", err)
				}
			}
		}
	}

	if err = d.Set("authmethod", flattenVpnIpsecPhase1Authmethod(o["authmethod"], d, "authmethod", sv)); err != nil {
		if !fortiAPIPatch(o["authmethod"]) {
			return fmt.Errorf("Error reading authmethod: %v", err)
		}
	}

	if err = d.Set("authmethod_remote", flattenVpnIpsecPhase1AuthmethodRemote(o["authmethod-remote"], d, "authmethod_remote", sv)); err != nil {
		if !fortiAPIPatch(o["authmethod-remote"]) {
			return fmt.Errorf("Error reading authmethod_remote: %v", err)
		}
	}

	if err = d.Set("mode", flattenVpnIpsecPhase1Mode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("peertype", flattenVpnIpsecPhase1Peertype(o["peertype"], d, "peertype", sv)); err != nil {
		if !fortiAPIPatch(o["peertype"]) {
			return fmt.Errorf("Error reading peertype: %v", err)
		}
	}

	if err = d.Set("peerid", flattenVpnIpsecPhase1Peerid(o["peerid"], d, "peerid", sv)); err != nil {
		if !fortiAPIPatch(o["peerid"]) {
			return fmt.Errorf("Error reading peerid: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnIpsecPhase1Usrgrp(o["usrgrp"], d, "usrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["usrgrp"]) {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("peer", flattenVpnIpsecPhase1Peer(o["peer"], d, "peer", sv)); err != nil {
		if !fortiAPIPatch(o["peer"]) {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peergrp", flattenVpnIpsecPhase1Peergrp(o["peergrp"], d, "peergrp", sv)); err != nil {
		if !fortiAPIPatch(o["peergrp"]) {
			return fmt.Errorf("Error reading peergrp: %v", err)
		}
	}

	if err = d.Set("mode_cfg", flattenVpnIpsecPhase1ModeCfg(o["mode-cfg"], d, "mode_cfg", sv)); err != nil {
		if !fortiAPIPatch(o["mode-cfg"]) {
			return fmt.Errorf("Error reading mode_cfg: %v", err)
		}
	}

	if err = d.Set("mode_cfg_allow_client_selector", flattenVpnIpsecPhase1ModeCfgAllowClientSelector(o["mode-cfg-allow-client-selector"], d, "mode_cfg_allow_client_selector", sv)); err != nil {
		if !fortiAPIPatch(o["mode-cfg-allow-client-selector"]) {
			return fmt.Errorf("Error reading mode_cfg_allow_client_selector: %v", err)
		}
	}

	if err = d.Set("assign_ip", flattenVpnIpsecPhase1AssignIp(o["assign-ip"], d, "assign_ip", sv)); err != nil {
		if !fortiAPIPatch(o["assign-ip"]) {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("assign_ip_from", flattenVpnIpsecPhase1AssignIpFrom(o["assign-ip-from"], d, "assign_ip_from", sv)); err != nil {
		if !fortiAPIPatch(o["assign-ip-from"]) {
			return fmt.Errorf("Error reading assign_ip_from: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenVpnIpsecPhase1Ipv4StartIp(o["ipv4-start-ip"], d, "ipv4_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-start-ip"]) {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenVpnIpsecPhase1Ipv4EndIp(o["ipv4-end-ip"], d, "ipv4_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-end-ip"]) {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_netmask", flattenVpnIpsecPhase1Ipv4Netmask(o["ipv4-netmask"], d, "ipv4_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-netmask"]) {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenVpnIpsecPhase1DhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ra-giaddr"]) {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenVpnIpsecPhase1Dhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp6-ra-linkaddr"]) {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenVpnIpsecPhase1DnsMode(o["dns-mode"], d, "dns_mode", sv)); err != nil {
		if !fortiAPIPatch(o["dns-mode"]) {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server1", flattenVpnIpsecPhase1Ipv4DnsServer1(o["ipv4-dns-server1"], d, "ipv4_dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server1"]) {
			return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server2", flattenVpnIpsecPhase1Ipv4DnsServer2(o["ipv4-dns-server2"], d, "ipv4_dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server2"]) {
			return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server3", flattenVpnIpsecPhase1Ipv4DnsServer3(o["ipv4-dns-server3"], d, "ipv4_dns_server3", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server3"]) {
			return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internal_domain_list", flattenVpnIpsecPhase1InternalDomainList(o["internal-domain-list"], d, "internal_domain_list", sv)); err != nil {
			if !fortiAPIPatch(o["internal-domain-list"]) {
				return fmt.Errorf("Error reading internal_domain_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internal_domain_list"); ok {
			if err = d.Set("internal_domain_list", flattenVpnIpsecPhase1InternalDomainList(o["internal-domain-list"], d, "internal_domain_list", sv)); err != nil {
				if !fortiAPIPatch(o["internal-domain-list"]) {
					return fmt.Errorf("Error reading internal_domain_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_wins_server1", flattenVpnIpsecPhase1Ipv4WinsServer1(o["ipv4-wins-server1"], d, "ipv4_wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-wins-server1"]) {
			return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server2", flattenVpnIpsecPhase1Ipv4WinsServer2(o["ipv4-wins-server2"], d, "ipv4_wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-wins-server2"]) {
			return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range", sv)); err != nil {
			if !fortiAPIPatch(o["ipv4-exclude-range"]) {
				return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv4_exclude_range"); ok {
			if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range", sv)); err != nil {
				if !fortiAPIPatch(o["ipv4-exclude-range"]) {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_split_include", flattenVpnIpsecPhase1Ipv4SplitInclude(o["ipv4-split-include"], d, "ipv4_split_include", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-split-include"]) {
			return fmt.Errorf("Error reading ipv4_split_include: %v", err)
		}
	}

	if err = d.Set("split_include_service", flattenVpnIpsecPhase1SplitIncludeService(o["split-include-service"], d, "split_include_service", sv)); err != nil {
		if !fortiAPIPatch(o["split-include-service"]) {
			return fmt.Errorf("Error reading split_include_service: %v", err)
		}
	}

	if err = d.Set("ipv4_name", flattenVpnIpsecPhase1Ipv4Name(o["ipv4-name"], d, "ipv4_name", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-name"]) {
			return fmt.Errorf("Error reading ipv4_name: %v", err)
		}
	}

	if err = d.Set("ipv6_start_ip", flattenVpnIpsecPhase1Ipv6StartIp(o["ipv6-start-ip"], d, "ipv6_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-start-ip"]) {
			return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_end_ip", flattenVpnIpsecPhase1Ipv6EndIp(o["ipv6-end-ip"], d, "ipv6_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-end-ip"]) {
			return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix", flattenVpnIpsecPhase1Ipv6Prefix(o["ipv6-prefix"], d, "ipv6_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-prefix"]) {
			return fmt.Errorf("Error reading ipv6_prefix: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnIpsecPhase1Ipv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server1"]) {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnIpsecPhase1Ipv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server2"]) {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server3", flattenVpnIpsecPhase1Ipv6DnsServer3(o["ipv6-dns-server3"], d, "ipv6_dns_server3", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server3"]) {
			return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range", sv)); err != nil {
			if !fortiAPIPatch(o["ipv6-exclude-range"]) {
				return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_exclude_range"); ok {
			if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range", sv)); err != nil {
				if !fortiAPIPatch(o["ipv6-exclude-range"]) {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_split_include", flattenVpnIpsecPhase1Ipv6SplitInclude(o["ipv6-split-include"], d, "ipv6_split_include", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-split-include"]) {
			return fmt.Errorf("Error reading ipv6_split_include: %v", err)
		}
	}

	if err = d.Set("ipv6_name", flattenVpnIpsecPhase1Ipv6Name(o["ipv6-name"], d, "ipv6_name", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-name"]) {
			return fmt.Errorf("Error reading ipv6_name: %v", err)
		}
	}

	if err = d.Set("ip_delay_interval", flattenVpnIpsecPhase1IpDelayInterval(o["ip-delay-interval"], d, "ip_delay_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ip-delay-interval"]) {
			return fmt.Errorf("Error reading ip_delay_interval: %v", err)
		}
	}

	if err = d.Set("unity_support", flattenVpnIpsecPhase1UnitySupport(o["unity-support"], d, "unity_support", sv)); err != nil {
		if !fortiAPIPatch(o["unity-support"]) {
			return fmt.Errorf("Error reading unity_support: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnIpsecPhase1Domain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("banner", flattenVpnIpsecPhase1Banner(o["banner"], d, "banner", sv)); err != nil {
		if !fortiAPIPatch(o["banner"]) {
			return fmt.Errorf("Error reading banner: %v", err)
		}
	}

	if err = d.Set("include_local_lan", flattenVpnIpsecPhase1IncludeLocalLan(o["include-local-lan"], d, "include_local_lan", sv)); err != nil {
		if !fortiAPIPatch(o["include-local-lan"]) {
			return fmt.Errorf("Error reading include_local_lan: %v", err)
		}
	}

	if err = d.Set("ipv4_split_exclude", flattenVpnIpsecPhase1Ipv4SplitExclude(o["ipv4-split-exclude"], d, "ipv4_split_exclude", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-split-exclude"]) {
			return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv6_split_exclude", flattenVpnIpsecPhase1Ipv6SplitExclude(o["ipv6-split-exclude"], d, "ipv6_split_exclude", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-split-exclude"]) {
			return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnIpsecPhase1SavePassword(o["save-password"], d, "save_password", sv)); err != nil {
		if !fortiAPIPatch(o["save-password"]) {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("client_auto_negotiate", flattenVpnIpsecPhase1ClientAutoNegotiate(o["client-auto-negotiate"], d, "client_auto_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["client-auto-negotiate"]) {
			return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
		}
	}

	if err = d.Set("client_keep_alive", flattenVpnIpsecPhase1ClientKeepAlive(o["client-keep-alive"], d, "client_keep_alive", sv)); err != nil {
		if !fortiAPIPatch(o["client-keep-alive"]) {
			return fmt.Errorf("Error reading client_keep_alive: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("backup_gateway", flattenVpnIpsecPhase1BackupGateway(o["backup-gateway"], d, "backup_gateway", sv)); err != nil {
			if !fortiAPIPatch(o["backup-gateway"]) {
				return fmt.Errorf("Error reading backup_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("backup_gateway"); ok {
			if err = d.Set("backup_gateway", flattenVpnIpsecPhase1BackupGateway(o["backup-gateway"], d, "backup_gateway", sv)); err != nil {
				if !fortiAPIPatch(o["backup-gateway"]) {
					return fmt.Errorf("Error reading backup_gateway: %v", err)
				}
			}
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase1Proposal(o["proposal"], d, "proposal", sv)); err != nil {
		if !fortiAPIPatch(o["proposal"]) {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase1AddRoute(o["add-route"], d, "add_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-route"]) {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("add_gw_route", flattenVpnIpsecPhase1AddGwRoute(o["add-gw-route"], d, "add_gw_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-gw-route"]) {
			return fmt.Errorf("Error reading add_gw_route: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase1Keepalive(o["keepalive"], d, "keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("distance", flattenVpnIpsecPhase1Distance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", flattenVpnIpsecPhase1Priority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("localid", flattenVpnIpsecPhase1Localid(o["localid"], d, "localid", sv)); err != nil {
		if !fortiAPIPatch(o["localid"]) {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("localid_type", flattenVpnIpsecPhase1LocalidType(o["localid-type"], d, "localid_type", sv)); err != nil {
		if !fortiAPIPatch(o["localid-type"]) {
			return fmt.Errorf("Error reading localid_type: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase1AutoNegotiate(o["auto-negotiate"], d, "auto_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["auto-negotiate"]) {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("negotiate_timeout", flattenVpnIpsecPhase1NegotiateTimeout(o["negotiate-timeout"], d, "negotiate_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["negotiate-timeout"]) {
			return fmt.Errorf("Error reading negotiate_timeout: %v", err)
		}
	}

	if err = d.Set("fragmentation", flattenVpnIpsecPhase1Fragmentation(o["fragmentation"], d, "fragmentation", sv)); err != nil {
		if !fortiAPIPatch(o["fragmentation"]) {
			return fmt.Errorf("Error reading fragmentation: %v", err)
		}
	}

	if err = d.Set("dpd", flattenVpnIpsecPhase1Dpd(o["dpd"], d, "dpd", sv)); err != nil {
		if !fortiAPIPatch(o["dpd"]) {
			return fmt.Errorf("Error reading dpd: %v", err)
		}
	}

	if err = d.Set("dpd_retrycount", flattenVpnIpsecPhase1DpdRetrycount(o["dpd-retrycount"], d, "dpd_retrycount", sv)); err != nil {
		if !fortiAPIPatch(o["dpd-retrycount"]) {
			return fmt.Errorf("Error reading dpd_retrycount: %v", err)
		}
	}

	if err = d.Set("dpd_retryinterval", flattenVpnIpsecPhase1DpdRetryinterval(o["dpd-retryinterval"], d, "dpd_retryinterval", sv)); err != nil {
		if !fortiAPIPatch(o["dpd-retryinterval"]) {
			return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
		}
	}

	if err = d.Set("forticlient_enforcement", flattenVpnIpsecPhase1ForticlientEnforcement(o["forticlient-enforcement"], d, "forticlient_enforcement", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-enforcement"]) {
			return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase1Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenVpnIpsecPhase1NpuOffload(o["npu-offload"], d, "npu_offload", sv)); err != nil {
		if !fortiAPIPatch(o["npu-offload"]) {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	if err = d.Set("send_cert_chain", flattenVpnIpsecPhase1SendCertChain(o["send-cert-chain"], d, "send_cert_chain", sv)); err != nil {
		if !fortiAPIPatch(o["send-cert-chain"]) {
			return fmt.Errorf("Error reading send_cert_chain: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase1Dhgrp(o["dhgrp"], d, "dhgrp", sv)); err != nil {
		if !fortiAPIPatch(o["dhgrp"]) {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("addke1", flattenVpnIpsecPhase1Addke1(o["addke1"], d, "addke1", sv)); err != nil {
		if !fortiAPIPatch(o["addke1"]) {
			return fmt.Errorf("Error reading addke1: %v", err)
		}
	}

	if err = d.Set("addke2", flattenVpnIpsecPhase1Addke2(o["addke2"], d, "addke2", sv)); err != nil {
		if !fortiAPIPatch(o["addke2"]) {
			return fmt.Errorf("Error reading addke2: %v", err)
		}
	}

	if err = d.Set("addke3", flattenVpnIpsecPhase1Addke3(o["addke3"], d, "addke3", sv)); err != nil {
		if !fortiAPIPatch(o["addke3"]) {
			return fmt.Errorf("Error reading addke3: %v", err)
		}
	}

	if err = d.Set("addke4", flattenVpnIpsecPhase1Addke4(o["addke4"], d, "addke4", sv)); err != nil {
		if !fortiAPIPatch(o["addke4"]) {
			return fmt.Errorf("Error reading addke4: %v", err)
		}
	}

	if err = d.Set("addke5", flattenVpnIpsecPhase1Addke5(o["addke5"], d, "addke5", sv)); err != nil {
		if !fortiAPIPatch(o["addke5"]) {
			return fmt.Errorf("Error reading addke5: %v", err)
		}
	}

	if err = d.Set("addke6", flattenVpnIpsecPhase1Addke6(o["addke6"], d, "addke6", sv)); err != nil {
		if !fortiAPIPatch(o["addke6"]) {
			return fmt.Errorf("Error reading addke6: %v", err)
		}
	}

	if err = d.Set("addke7", flattenVpnIpsecPhase1Addke7(o["addke7"], d, "addke7", sv)); err != nil {
		if !fortiAPIPatch(o["addke7"]) {
			return fmt.Errorf("Error reading addke7: %v", err)
		}
	}

	if err = d.Set("suite_b", flattenVpnIpsecPhase1SuiteB(o["suite-b"], d, "suite_b", sv)); err != nil {
		if !fortiAPIPatch(o["suite-b"]) {
			return fmt.Errorf("Error reading suite_b: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnIpsecPhase1Eap(o["eap"], d, "eap", sv)); err != nil {
		if !fortiAPIPatch(o["eap"]) {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_identity", flattenVpnIpsecPhase1EapIdentity(o["eap-identity"], d, "eap_identity", sv)); err != nil {
		if !fortiAPIPatch(o["eap-identity"]) {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("eap_exclude_peergrp", flattenVpnIpsecPhase1EapExcludePeergrp(o["eap-exclude-peergrp"], d, "eap_exclude_peergrp", sv)); err != nil {
		if !fortiAPIPatch(o["eap-exclude-peergrp"]) {
			return fmt.Errorf("Error reading eap_exclude_peergrp: %v", err)
		}
	}

	if err = d.Set("eap_cert_auth", flattenVpnIpsecPhase1EapCertAuth(o["eap-cert-auth"], d, "eap_cert_auth", sv)); err != nil {
		if !fortiAPIPatch(o["eap-cert-auth"]) {
			return fmt.Errorf("Error reading eap_cert_auth: %v", err)
		}
	}

	if err = d.Set("acct_verify", flattenVpnIpsecPhase1AcctVerify(o["acct-verify"], d, "acct_verify", sv)); err != nil {
		if !fortiAPIPatch(o["acct-verify"]) {
			return fmt.Errorf("Error reading acct_verify: %v", err)
		}
	}

	if err = d.Set("ppk", flattenVpnIpsecPhase1Ppk(o["ppk"], d, "ppk", sv)); err != nil {
		if !fortiAPIPatch(o["ppk"]) {
			return fmt.Errorf("Error reading ppk: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenVpnIpsecPhase1PpkIdentity(o["ppk-identity"], d, "ppk_identity", sv)); err != nil {
		if !fortiAPIPatch(o["ppk-identity"]) {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("wizard_type", flattenVpnIpsecPhase1WizardType(o["wizard-type"], d, "wizard_type", sv)); err != nil {
		if !fortiAPIPatch(o["wizard-type"]) {
			return fmt.Errorf("Error reading wizard_type: %v", err)
		}
	}

	if err = d.Set("xauthtype", flattenVpnIpsecPhase1Xauthtype(o["xauthtype"], d, "xauthtype", sv)); err != nil {
		if !fortiAPIPatch(o["xauthtype"]) {
			return fmt.Errorf("Error reading xauthtype: %v", err)
		}
	}

	if err = d.Set("reauth", flattenVpnIpsecPhase1Reauth(o["reauth"], d, "reauth", sv)); err != nil {
		if !fortiAPIPatch(o["reauth"]) {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("authusr", flattenVpnIpsecPhase1Authusr(o["authusr"], d, "authusr", sv)); err != nil {
		if !fortiAPIPatch(o["authusr"]) {
			return fmt.Errorf("Error reading authusr: %v", err)
		}
	}

	if err = d.Set("group_authentication", flattenVpnIpsecPhase1GroupAuthentication(o["group-authentication"], d, "group_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["group-authentication"]) {
			return fmt.Errorf("Error reading group_authentication: %v", err)
		}
	}

	if err = d.Set("authusrgrp", flattenVpnIpsecPhase1Authusrgrp(o["authusrgrp"], d, "authusrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["authusrgrp"]) {
			return fmt.Errorf("Error reading authusrgrp: %v", err)
		}
	}

	if err = d.Set("mesh_selector_type", flattenVpnIpsecPhase1MeshSelectorType(o["mesh-selector-type"], d, "mesh_selector_type", sv)); err != nil {
		if !fortiAPIPatch(o["mesh-selector-type"]) {
			return fmt.Errorf("Error reading mesh_selector_type: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnIpsecPhase1IdleTimeout(o["idle-timeout"], d, "idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("shared_idle_timeout", flattenVpnIpsecPhase1SharedIdleTimeout(o["shared-idle-timeout"], d, "shared_idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["shared-idle-timeout"]) {
			return fmt.Errorf("Error reading shared_idle_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeoutinterval", flattenVpnIpsecPhase1IdleTimeoutinterval(o["idle-timeoutinterval"], d, "idle_timeoutinterval", sv)); err != nil {
		if !fortiAPIPatch(o["idle-timeoutinterval"]) {
			return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
		}
	}

	if err = d.Set("ha_sync_esp_seqno", flattenVpnIpsecPhase1HaSyncEspSeqno(o["ha-sync-esp-seqno"], d, "ha_sync_esp_seqno", sv)); err != nil {
		if !fortiAPIPatch(o["ha-sync-esp-seqno"]) {
			return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
		}
	}

	if err = d.Set("fgsp_sync", flattenVpnIpsecPhase1FgspSync(o["fgsp-sync"], d, "fgsp_sync", sv)); err != nil {
		if !fortiAPIPatch(o["fgsp-sync"]) {
			return fmt.Errorf("Error reading fgsp_sync: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenVpnIpsecPhase1InboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy", sv)); err != nil {
		if !fortiAPIPatch(o["inbound-dscp-copy"]) {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("nattraversal", flattenVpnIpsecPhase1Nattraversal(o["nattraversal"], d, "nattraversal", sv)); err != nil {
		if !fortiAPIPatch(o["nattraversal"]) {
			return fmt.Errorf("Error reading nattraversal: %v", err)
		}
	}

	if err = d.Set("esn", flattenVpnIpsecPhase1Esn(o["esn"], d, "esn", sv)); err != nil {
		if !fortiAPIPatch(o["esn"]) {
			return fmt.Errorf("Error reading esn: %v", err)
		}
	}

	if err = d.Set("fragmentation_mtu", flattenVpnIpsecPhase1FragmentationMtu(o["fragmentation-mtu"], d, "fragmentation_mtu", sv)); err != nil {
		if !fortiAPIPatch(o["fragmentation-mtu"]) {
			return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
		}
	}

	if err = d.Set("childless_ike", flattenVpnIpsecPhase1ChildlessIke(o["childless-ike"], d, "childless_ike", sv)); err != nil {
		if !fortiAPIPatch(o["childless-ike"]) {
			return fmt.Errorf("Error reading childless_ike: %v", err)
		}
	}

	if err = d.Set("azure_ad_autoconnect", flattenVpnIpsecPhase1AzureAdAutoconnect(o["azure-ad-autoconnect"], d, "azure_ad_autoconnect", sv)); err != nil {
		if !fortiAPIPatch(o["azure-ad-autoconnect"]) {
			return fmt.Errorf("Error reading azure_ad_autoconnect: %v", err)
		}
	}

	if err = d.Set("client_resume", flattenVpnIpsecPhase1ClientResume(o["client-resume"], d, "client_resume", sv)); err != nil {
		if !fortiAPIPatch(o["client-resume"]) {
			return fmt.Errorf("Error reading client_resume: %v", err)
		}
	}

	if err = d.Set("client_resume_interval", flattenVpnIpsecPhase1ClientResumeInterval(o["client-resume-interval"], d, "client_resume_interval", sv)); err != nil {
		if !fortiAPIPatch(o["client-resume-interval"]) {
			return fmt.Errorf("Error reading client_resume_interval: %v", err)
		}
	}

	if err = d.Set("rekey", flattenVpnIpsecPhase1Rekey(o["rekey"], d, "rekey", sv)); err != nil {
		if !fortiAPIPatch(o["rekey"]) {
			return fmt.Errorf("Error reading rekey: %v", err)
		}
	}

	if err = d.Set("digital_signature_auth", flattenVpnIpsecPhase1DigitalSignatureAuth(o["digital-signature-auth"], d, "digital_signature_auth", sv)); err != nil {
		if !fortiAPIPatch(o["digital-signature-auth"]) {
			return fmt.Errorf("Error reading digital_signature_auth: %v", err)
		}
	}

	if err = d.Set("signature_hash_alg", flattenVpnIpsecPhase1SignatureHashAlg(o["signature-hash-alg"], d, "signature_hash_alg", sv)); err != nil {
		if !fortiAPIPatch(o["signature-hash-alg"]) {
			return fmt.Errorf("Error reading signature_hash_alg: %v", err)
		}
	}

	if err = d.Set("rsa_signature_format", flattenVpnIpsecPhase1RsaSignatureFormat(o["rsa-signature-format"], d, "rsa_signature_format", sv)); err != nil {
		if !fortiAPIPatch(o["rsa-signature-format"]) {
			return fmt.Errorf("Error reading rsa_signature_format: %v", err)
		}
	}

	if err = d.Set("rsa_signature_hash_override", flattenVpnIpsecPhase1RsaSignatureHashOverride(o["rsa-signature-hash-override"], d, "rsa_signature_hash_override", sv)); err != nil {
		if !fortiAPIPatch(o["rsa-signature-hash-override"]) {
			return fmt.Errorf("Error reading rsa_signature_hash_override: %v", err)
		}
	}

	if err = d.Set("enforce_unique_id", flattenVpnIpsecPhase1EnforceUniqueId(o["enforce-unique-id"], d, "enforce_unique_id", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-unique-id"]) {
			return fmt.Errorf("Error reading enforce_unique_id: %v", err)
		}
	}

	if err = d.Set("cert_id_validation", flattenVpnIpsecPhase1CertIdValidation(o["cert-id-validation"], d, "cert_id_validation", sv)); err != nil {
		if !fortiAPIPatch(o["cert-id-validation"]) {
			return fmt.Errorf("Error reading cert_id_validation: %v", err)
		}
	}

	if err = d.Set("fec_egress", flattenVpnIpsecPhase1FecEgress(o["fec-egress"], d, "fec_egress", sv)); err != nil {
		if !fortiAPIPatch(o["fec-egress"]) {
			return fmt.Errorf("Error reading fec_egress: %v", err)
		}
	}

	if err = d.Set("fec_send_timeout", flattenVpnIpsecPhase1FecSendTimeout(o["fec-send-timeout"], d, "fec_send_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["fec-send-timeout"]) {
			return fmt.Errorf("Error reading fec_send_timeout: %v", err)
		}
	}

	if err = d.Set("fec_base", flattenVpnIpsecPhase1FecBase(o["fec-base"], d, "fec_base", sv)); err != nil {
		if !fortiAPIPatch(o["fec-base"]) {
			return fmt.Errorf("Error reading fec_base: %v", err)
		}
	}

	if _, ok := o["fec-codec"].(string); ok {
		if err = d.Set("fec_codec_string", flattenVpnIpsecPhase1FecCodecString(o["fec-codec"], d, "fec_codec_string", sv)); err != nil {
			if !fortiAPIPatch(o["fec-codec"]) {
				return fmt.Errorf("Error reading fec_codec_string: %v", err)
			}
		}
	}

	if _, ok := o["fec-codec"].(float64); ok {
		if err = d.Set("fec_codec", flattenVpnIpsecPhase1FecCodec(o["fec-codec"], d, "fec_codec", sv)); err != nil {
			if !fortiAPIPatch(o["fec-codec"]) {
				return fmt.Errorf("Error reading fec_codec: %v", err)
			}
		}
	}

	if err = d.Set("fec_redundant", flattenVpnIpsecPhase1FecRedundant(o["fec-redundant"], d, "fec_redundant", sv)); err != nil {
		if !fortiAPIPatch(o["fec-redundant"]) {
			return fmt.Errorf("Error reading fec_redundant: %v", err)
		}
	}

	if err = d.Set("fec_ingress", flattenVpnIpsecPhase1FecIngress(o["fec-ingress"], d, "fec_ingress", sv)); err != nil {
		if !fortiAPIPatch(o["fec-ingress"]) {
			return fmt.Errorf("Error reading fec_ingress: %v", err)
		}
	}

	if err = d.Set("fec_receive_timeout", flattenVpnIpsecPhase1FecReceiveTimeout(o["fec-receive-timeout"], d, "fec_receive_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["fec-receive-timeout"]) {
			return fmt.Errorf("Error reading fec_receive_timeout: %v", err)
		}
	}

	if err = d.Set("fec_health_check", flattenVpnIpsecPhase1FecHealthCheck(o["fec-health-check"], d, "fec_health_check", sv)); err != nil {
		if !fortiAPIPatch(o["fec-health-check"]) {
			return fmt.Errorf("Error reading fec_health_check: %v", err)
		}
	}

	if err = d.Set("fec_mapping_profile", flattenVpnIpsecPhase1FecMappingProfile(o["fec-mapping-profile"], d, "fec_mapping_profile", sv)); err != nil {
		if !fortiAPIPatch(o["fec-mapping-profile"]) {
			return fmt.Errorf("Error reading fec_mapping_profile: %v", err)
		}
	}

	if err = d.Set("network_overlay", flattenVpnIpsecPhase1NetworkOverlay(o["network-overlay"], d, "network_overlay", sv)); err != nil {
		if !fortiAPIPatch(o["network-overlay"]) {
			return fmt.Errorf("Error reading network_overlay: %v", err)
		}
	}

	if err = d.Set("network_id", flattenVpnIpsecPhase1NetworkId(o["network-id"], d, "network_id", sv)); err != nil {
		if !fortiAPIPatch(o["network-id"]) {
			return fmt.Errorf("Error reading network_id: %v", err)
		}
	}

	if err = d.Set("dev_id_notification", flattenVpnIpsecPhase1DevIdNotification(o["dev-id-notification"], d, "dev_id_notification", sv)); err != nil {
		if !fortiAPIPatch(o["dev-id-notification"]) {
			return fmt.Errorf("Error reading dev_id_notification: %v", err)
		}
	}

	if err = d.Set("dev_id", flattenVpnIpsecPhase1DevId(o["dev-id"], d, "dev_id", sv)); err != nil {
		if !fortiAPIPatch(o["dev-id"]) {
			return fmt.Errorf("Error reading dev_id: %v", err)
		}
	}

	if err = d.Set("loopback_asymroute", flattenVpnIpsecPhase1LoopbackAsymroute(o["loopback-asymroute"], d, "loopback_asymroute", sv)); err != nil {
		if !fortiAPIPatch(o["loopback-asymroute"]) {
			return fmt.Errorf("Error reading loopback_asymroute: %v", err)
		}
	}

	if err = d.Set("link_cost", flattenVpnIpsecPhase1LinkCost(o["link-cost"], d, "link_cost", sv)); err != nil {
		if !fortiAPIPatch(o["link-cost"]) {
			return fmt.Errorf("Error reading link_cost: %v", err)
		}
	}

	if err = d.Set("kms", flattenVpnIpsecPhase1Kms(o["kms"], d, "kms", sv)); err != nil {
		if !fortiAPIPatch(o["kms"]) {
			return fmt.Errorf("Error reading kms: %v", err)
		}
	}

	if err = d.Set("exchange_fgt_device_id", flattenVpnIpsecPhase1ExchangeFgtDeviceId(o["exchange-fgt-device-id"], d, "exchange_fgt_device_id", sv)); err != nil {
		if !fortiAPIPatch(o["exchange-fgt-device-id"]) {
			return fmt.Errorf("Error reading exchange_fgt_device_id: %v", err)
		}
	}

	if err = d.Set("ipv6_auto_linklocal", flattenVpnIpsecPhase1Ipv6AutoLinklocal(o["ipv6-auto-linklocal"], d, "ipv6_auto_linklocal", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-auto-linklocal"]) {
			return fmt.Errorf("Error reading ipv6_auto_linklocal: %v", err)
		}
	}

	if err = d.Set("ems_sn_check", flattenVpnIpsecPhase1EmsSnCheck(o["ems-sn-check"], d, "ems_sn_check", sv)); err != nil {
		if !fortiAPIPatch(o["ems-sn-check"]) {
			return fmt.Errorf("Error reading ems_sn_check: %v", err)
		}
	}

	if err = d.Set("cert_trust_store", flattenVpnIpsecPhase1CertTrustStore(o["cert-trust-store"], d, "cert_trust_store", sv)); err != nil {
		if !fortiAPIPatch(o["cert-trust-store"]) {
			return fmt.Errorf("Error reading cert_trust_store: %v", err)
		}
	}

	if err = d.Set("qkd", flattenVpnIpsecPhase1Qkd(o["qkd"], d, "qkd", sv)); err != nil {
		if !fortiAPIPatch(o["qkd"]) {
			return fmt.Errorf("Error reading qkd: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenVpnIpsecPhase1QkdProfile(o["qkd-profile"], d, "qkd_profile", sv)); err != nil {
		if !fortiAPIPatch(o["qkd-profile"]) {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("transport", flattenVpnIpsecPhase1Transport(o["transport"], d, "transport", sv)); err != nil {
		if !fortiAPIPatch(o["transport"]) {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	if err = d.Set("fortinet_esp", flattenVpnIpsecPhase1FortinetEsp(o["fortinet-esp"], d, "fortinet_esp", sv)); err != nil {
		if !fortiAPIPatch(o["fortinet-esp"]) {
			return fmt.Errorf("Error reading fortinet_esp: %v", err)
		}
	}

	if err = d.Set("auto_transport_threshold", flattenVpnIpsecPhase1AutoTransportThreshold(o["auto-transport-threshold"], d, "auto_transport_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["auto-transport-threshold"]) {
			return fmt.Errorf("Error reading auto_transport_threshold: %v", err)
		}
	}

	if err = d.Set("fallback_tcp_threshold", flattenVpnIpsecPhase1FallbackTcpThreshold(o["fallback-tcp-threshold"], d, "fallback_tcp_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["fallback-tcp-threshold"]) {
			return fmt.Errorf("Error reading fallback_tcp_threshold: %v", err)
		}
	}

	if err = d.Set("remote_gw_match", flattenVpnIpsecPhase1RemoteGwMatch(o["remote-gw-match"], d, "remote_gw_match", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-match"]) {
			return fmt.Errorf("Error reading remote_gw_match: %v", err)
		}
	}

	if err = d.Set("remote_gw_subnet", flattenVpnIpsecPhase1RemoteGwSubnet(o["remote-gw-subnet"], d, "remote_gw_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-subnet"]) {
			return fmt.Errorf("Error reading remote_gw_subnet: %v", err)
		}
	}

	if err = d.Set("remote_gw_start_ip", flattenVpnIpsecPhase1RemoteGwStartIp(o["remote-gw-start-ip"], d, "remote_gw_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-start-ip"]) {
			return fmt.Errorf("Error reading remote_gw_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_end_ip", flattenVpnIpsecPhase1RemoteGwEndIp(o["remote-gw-end-ip"], d, "remote_gw_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-end-ip"]) {
			return fmt.Errorf("Error reading remote_gw_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_country", flattenVpnIpsecPhase1RemoteGwCountry(o["remote-gw-country"], d, "remote_gw_country", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-country"]) {
			return fmt.Errorf("Error reading remote_gw_country: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("remote_gw_ztna_tags", flattenVpnIpsecPhase1RemoteGwZtnaTags(o["remote-gw-ztna-tags"], d, "remote_gw_ztna_tags", sv)); err != nil {
			if !fortiAPIPatch(o["remote-gw-ztna-tags"]) {
				return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_gw_ztna_tags"); ok {
			if err = d.Set("remote_gw_ztna_tags", flattenVpnIpsecPhase1RemoteGwZtnaTags(o["remote-gw-ztna-tags"], d, "remote_gw_ztna_tags", sv)); err != nil {
				if !fortiAPIPatch(o["remote-gw-ztna-tags"]) {
					return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
				}
			}
		}
	}

	if err = d.Set("remote_gw6_match", flattenVpnIpsecPhase1RemoteGw6Match(o["remote-gw6-match"], d, "remote_gw6_match", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-match"]) {
			return fmt.Errorf("Error reading remote_gw6_match: %v", err)
		}
	}

	if err = d.Set("remote_gw6_subnet", flattenVpnIpsecPhase1RemoteGw6Subnet(o["remote-gw6-subnet"], d, "remote_gw6_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-subnet"]) {
			return fmt.Errorf("Error reading remote_gw6_subnet: %v", err)
		}
	}

	if err = d.Set("remote_gw6_start_ip", flattenVpnIpsecPhase1RemoteGw6StartIp(o["remote-gw6-start-ip"], d, "remote_gw6_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-start-ip"]) {
			return fmt.Errorf("Error reading remote_gw6_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_end_ip", flattenVpnIpsecPhase1RemoteGw6EndIp(o["remote-gw6-end-ip"], d, "remote_gw6_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-end-ip"]) {
			return fmt.Errorf("Error reading remote_gw6_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_country", flattenVpnIpsecPhase1RemoteGw6Country(o["remote-gw6-country"], d, "remote_gw6_country", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-country"]) {
			return fmt.Errorf("Error reading remote_gw6_country: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_validation", flattenVpnIpsecPhase1CertPeerUsernameValidation(o["cert-peer-username-validation"], d, "cert_peer_username_validation", sv)); err != nil {
		if !fortiAPIPatch(o["cert-peer-username-validation"]) {
			return fmt.Errorf("Error reading cert_peer_username_validation: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_strip", flattenVpnIpsecPhase1CertPeerUsernameStrip(o["cert-peer-username-strip"], d, "cert_peer_username_strip", sv)); err != nil {
		if !fortiAPIPatch(o["cert-peer-username-strip"]) {
			return fmt.Errorf("Error reading cert_peer_username_strip: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecPhase1Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Type(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IkeVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LocalGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemotegwDdns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Keylife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Certificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnIpsecPhase1CertificateName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1CertificateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authmethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AuthmethodRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peertype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peerid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Usrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peergrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ModeCfg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ModeCfgAllowClientSelector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AssignIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AssignIpFrom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4Netmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DnsMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InternalDomainList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["domain-name"], _ = expandVpnIpsecPhase1InternalDomainListDomainName(d, i["domain_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InternalDomainListDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4WinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4WinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4SplitInclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SplitIncludeService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6SplitInclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IpDelayInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1UnitySupport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Domain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Banner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IncludeLocalLan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4SplitExclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6SplitExclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SavePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientAutoNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientKeepAlive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1BackupGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["address"], _ = expandVpnIpsecPhase1BackupGatewayAddress(d, i["address"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1BackupGatewayAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Proposal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AddRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AddGwRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Psksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PsksecretRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Keepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Distance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Priority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Localid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LocalidType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AutoNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NegotiateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Fragmentation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DpdRetrycount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DpdRetryinterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ForticlientEnforcement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NpuOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SendCertChain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dhgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Addke7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SuiteB(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Eap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EapIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EapExcludePeergrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EapCertAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AcctVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ppk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PpkSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PpkIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1WizardType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Xauthtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Reauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authusr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authpasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1GroupAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1GroupAuthenticationSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authusrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1MeshSelectorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SharedIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IdleTimeoutinterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1HaSyncEspSeqno(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FgspSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InboundDscpCopy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Nattraversal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Esn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FragmentationMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ChildlessIke(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AzureAdAutoconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientResume(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientResumeInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Rekey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DigitalSignatureAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SignatureHashAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RsaSignatureFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RsaSignatureHashOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EnforceUniqueId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertIdValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecEgress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecSendTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecBase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecCodecString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecCodec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecRedundant(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecIngress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecReceiveTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FecMappingProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NetworkOverlay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DevIdNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DevId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LoopbackAsymroute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LinkCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Kms(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ExchangeFgtDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6AutoLinklocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EmsSnCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertTrustStore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Qkd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1QkdProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Transport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FortinetEsp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AutoTransportThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FallbackTcpThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGwZtnaTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnIpsecPhase1RemoteGwZtnaTagsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1RemoteGwZtnaTagsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6Match(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6Subnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw6Country(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertPeerUsernameValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertPeerUsernameStrip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase1(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnIpsecPhase1Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandVpnIpsecPhase1Type(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandVpnIpsecPhase1Interface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("ike_version"); ok {
		t, err := expandVpnIpsecPhase1IkeVersion(d, v, "ike_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw(d, v, "remote_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandVpnIpsecPhase1LocalGw(d, v, "local_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("remotegw_ddns"); ok {
		t, err := expandVpnIpsecPhase1RemotegwDdns(d, v, "remotegw_ddns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotegw-ddns"] = t
		}
	} else if d.HasChange("remotegw_ddns") {
		obj["remotegw-ddns"] = nil
	}

	if v, ok := d.GetOk("keylife"); ok {
		t, err := expandVpnIpsecPhase1Keylife(d, v, "keylife", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandVpnIpsecPhase1Certificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("authmethod"); ok {
		t, err := expandVpnIpsecPhase1Authmethod(d, v, "authmethod", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod"] = t
		}
	}

	if v, ok := d.GetOk("authmethod_remote"); ok {
		t, err := expandVpnIpsecPhase1AuthmethodRemote(d, v, "authmethod_remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod-remote"] = t
		}
	} else if d.HasChange("authmethod_remote") {
		obj["authmethod-remote"] = nil
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandVpnIpsecPhase1Mode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("peertype"); ok {
		t, err := expandVpnIpsecPhase1Peertype(d, v, "peertype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peertype"] = t
		}
	}

	if v, ok := d.GetOk("peerid"); ok {
		t, err := expandVpnIpsecPhase1Peerid(d, v, "peerid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerid"] = t
		}
	} else if d.HasChange("peerid") {
		obj["peerid"] = nil
	}

	if v, ok := d.GetOk("usrgrp"); ok {
		t, err := expandVpnIpsecPhase1Usrgrp(d, v, "usrgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	} else if d.HasChange("usrgrp") {
		obj["usrgrp"] = nil
	}

	if v, ok := d.GetOk("peer"); ok {
		t, err := expandVpnIpsecPhase1Peer(d, v, "peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	} else if d.HasChange("peer") {
		obj["peer"] = nil
	}

	if v, ok := d.GetOk("peergrp"); ok {
		t, err := expandVpnIpsecPhase1Peergrp(d, v, "peergrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peergrp"] = t
		}
	} else if d.HasChange("peergrp") {
		obj["peergrp"] = nil
	}

	if v, ok := d.GetOk("mode_cfg"); ok {
		t, err := expandVpnIpsecPhase1ModeCfg(d, v, "mode_cfg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg_allow_client_selector"); ok {
		t, err := expandVpnIpsecPhase1ModeCfgAllowClientSelector(d, v, "mode_cfg_allow_client_selector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg-allow-client-selector"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip"); ok {
		t, err := expandVpnIpsecPhase1AssignIp(d, v, "assign_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip_from"); ok {
		t, err := expandVpnIpsecPhase1AssignIpFrom(d, v, "assign_ip_from", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip-from"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv4StartIp(d, v, "ipv4_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv4EndIp(d, v, "ipv4_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_netmask"); ok {
		t, err := expandVpnIpsecPhase1Ipv4Netmask(d, v, "ipv4_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok {
		t, err := expandVpnIpsecPhase1DhcpRaGiaddr(d, v, "dhcp_ra_giaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok {
		t, err := expandVpnIpsecPhase1Dhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok {
		t, err := expandVpnIpsecPhase1DnsMode(d, v, "dns_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server1"); ok {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer1(d, v, "ipv4_dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server2"); ok {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer2(d, v, "ipv4_dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server3"); ok {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer3(d, v, "ipv4_dns_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("internal_domain_list"); ok || d.HasChange("internal_domain_list") {
		t, err := expandVpnIpsecPhase1InternalDomainList(d, v, "internal_domain_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-domain-list"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server1"); ok {
		t, err := expandVpnIpsecPhase1Ipv4WinsServer1(d, v, "ipv4_wins_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server2"); ok {
		t, err := expandVpnIpsecPhase1Ipv4WinsServer2(d, v, "ipv4_wins_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_exclude_range"); ok || d.HasChange("ipv4_exclude_range") {
		t, err := expandVpnIpsecPhase1Ipv4ExcludeRange(d, v, "ipv4_exclude_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_include"); ok {
		t, err := expandVpnIpsecPhase1Ipv4SplitInclude(d, v, "ipv4_split_include", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-include"] = t
		}
	} else if d.HasChange("ipv4_split_include") {
		obj["ipv4-split-include"] = nil
	}

	if v, ok := d.GetOk("split_include_service"); ok {
		t, err := expandVpnIpsecPhase1SplitIncludeService(d, v, "split_include_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-include-service"] = t
		}
	} else if d.HasChange("split_include_service") {
		obj["split-include-service"] = nil
	}

	if v, ok := d.GetOk("ipv4_name"); ok {
		t, err := expandVpnIpsecPhase1Ipv4Name(d, v, "ipv4_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-name"] = t
		}
	} else if d.HasChange("ipv4_name") {
		obj["ipv4-name"] = nil
	}

	if v, ok := d.GetOk("ipv6_start_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv6StartIp(d, v, "ipv6_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_end_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv6EndIp(d, v, "ipv6_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix"); ok {
		t, err := expandVpnIpsecPhase1Ipv6Prefix(d, v, "ipv6_prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer1(d, v, "ipv6_dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer2(d, v, "ipv6_dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server3"); ok {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer3(d, v, "ipv6_dns_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclude_range"); ok || d.HasChange("ipv6_exclude_range") {
		t, err := expandVpnIpsecPhase1Ipv6ExcludeRange(d, v, "ipv6_exclude_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_include"); ok {
		t, err := expandVpnIpsecPhase1Ipv6SplitInclude(d, v, "ipv6_split_include", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-include"] = t
		}
	} else if d.HasChange("ipv6_split_include") {
		obj["ipv6-split-include"] = nil
	}

	if v, ok := d.GetOk("ipv6_name"); ok {
		t, err := expandVpnIpsecPhase1Ipv6Name(d, v, "ipv6_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-name"] = t
		}
	} else if d.HasChange("ipv6_name") {
		obj["ipv6-name"] = nil
	}

	if v, ok := d.GetOkExists("ip_delay_interval"); ok {
		t, err := expandVpnIpsecPhase1IpDelayInterval(d, v, "ip_delay_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-delay-interval"] = t
		}
	} else if d.HasChange("ip_delay_interval") {
		obj["ip-delay-interval"] = nil
	}

	if v, ok := d.GetOk("unity_support"); ok {
		t, err := expandVpnIpsecPhase1UnitySupport(d, v, "unity_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unity-support"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandVpnIpsecPhase1Domain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if d.HasChange("domain") {
		obj["domain"] = nil
	}

	if v, ok := d.GetOk("banner"); ok {
		t, err := expandVpnIpsecPhase1Banner(d, v, "banner", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner"] = t
		}
	} else if d.HasChange("banner") {
		obj["banner"] = nil
	}

	if v, ok := d.GetOk("include_local_lan"); ok {
		t, err := expandVpnIpsecPhase1IncludeLocalLan(d, v, "include_local_lan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-local-lan"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_exclude"); ok {
		t, err := expandVpnIpsecPhase1Ipv4SplitExclude(d, v, "ipv4_split_exclude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-exclude"] = t
		}
	} else if d.HasChange("ipv4_split_exclude") {
		obj["ipv4-split-exclude"] = nil
	}

	if v, ok := d.GetOk("ipv6_split_exclude"); ok {
		t, err := expandVpnIpsecPhase1Ipv6SplitExclude(d, v, "ipv6_split_exclude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-exclude"] = t
		}
	} else if d.HasChange("ipv6_split_exclude") {
		obj["ipv6-split-exclude"] = nil
	}

	if v, ok := d.GetOk("save_password"); ok {
		t, err := expandVpnIpsecPhase1SavePassword(d, v, "save_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("client_auto_negotiate"); ok {
		t, err := expandVpnIpsecPhase1ClientAutoNegotiate(d, v, "client_auto_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("client_keep_alive"); ok {
		t, err := expandVpnIpsecPhase1ClientKeepAlive(d, v, "client_keep_alive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("backup_gateway"); ok || d.HasChange("backup_gateway") {
		t, err := expandVpnIpsecPhase1BackupGateway(d, v, "backup_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-gateway"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok {
		t, err := expandVpnIpsecPhase1Proposal(d, v, "proposal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	} else if d.HasChange("proposal") {
		obj["proposal"] = nil
	}

	if v, ok := d.GetOk("add_route"); ok {
		t, err := expandVpnIpsecPhase1AddRoute(d, v, "add_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("add_gw_route"); ok {
		t, err := expandVpnIpsecPhase1AddGwRoute(d, v, "add_gw_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-gw-route"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		t, err := expandVpnIpsecPhase1Psksecret(d, v, "psksecret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	} else if d.HasChange("psksecret") {
		obj["psksecret"] = nil
	}

	if v, ok := d.GetOk("psksecret_remote"); ok {
		t, err := expandVpnIpsecPhase1PsksecretRemote(d, v, "psksecret_remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret-remote"] = t
		}
	} else if d.HasChange("psksecret_remote") {
		obj["psksecret-remote"] = nil
	}

	if v, ok := d.GetOk("keepalive"); ok {
		t, err := expandVpnIpsecPhase1Keepalive(d, v, "keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandVpnIpsecPhase1Distance(d, v, "distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {
		t, err := expandVpnIpsecPhase1Priority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("localid"); ok {
		t, err := expandVpnIpsecPhase1Localid(d, v, "localid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid"] = t
		}
	} else if d.HasChange("localid") {
		obj["localid"] = nil
	}

	if v, ok := d.GetOk("localid_type"); ok {
		t, err := expandVpnIpsecPhase1LocalidType(d, v, "localid_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok {
		t, err := expandVpnIpsecPhase1AutoNegotiate(d, v, "auto_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_timeout"); ok {
		t, err := expandVpnIpsecPhase1NegotiateTimeout(d, v, "negotiate_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation"); ok {
		t, err := expandVpnIpsecPhase1Fragmentation(d, v, "fragmentation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("dpd"); ok {
		t, err := expandVpnIpsecPhase1Dpd(d, v, "dpd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd"] = t
		}
	}

	if v, ok := d.GetOkExists("dpd_retrycount"); ok {
		t, err := expandVpnIpsecPhase1DpdRetrycount(d, v, "dpd_retrycount", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retrycount"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retryinterval"); ok {
		t, err := expandVpnIpsecPhase1DpdRetryinterval(d, v, "dpd_retryinterval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retryinterval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_enforcement"); ok {
		t, err := expandVpnIpsecPhase1ForticlientEnforcement(d, v, "forticlient_enforcement", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandVpnIpsecPhase1Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("npu_offload"); ok {
		t, err := expandVpnIpsecPhase1NpuOffload(d, v, "npu_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	if v, ok := d.GetOk("send_cert_chain"); ok {
		t, err := expandVpnIpsecPhase1SendCertChain(d, v, "send_cert_chain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-cert-chain"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok {
		t, err := expandVpnIpsecPhase1Dhgrp(d, v, "dhgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("addke1"); ok {
		t, err := expandVpnIpsecPhase1Addke1(d, v, "addke1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke1"] = t
		}
	} else if d.HasChange("addke1") {
		obj["addke1"] = nil
	}

	if v, ok := d.GetOk("addke2"); ok {
		t, err := expandVpnIpsecPhase1Addke2(d, v, "addke2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke2"] = t
		}
	} else if d.HasChange("addke2") {
		obj["addke2"] = nil
	}

	if v, ok := d.GetOk("addke3"); ok {
		t, err := expandVpnIpsecPhase1Addke3(d, v, "addke3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke3"] = t
		}
	} else if d.HasChange("addke3") {
		obj["addke3"] = nil
	}

	if v, ok := d.GetOk("addke4"); ok {
		t, err := expandVpnIpsecPhase1Addke4(d, v, "addke4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke4"] = t
		}
	} else if d.HasChange("addke4") {
		obj["addke4"] = nil
	}

	if v, ok := d.GetOk("addke5"); ok {
		t, err := expandVpnIpsecPhase1Addke5(d, v, "addke5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke5"] = t
		}
	} else if d.HasChange("addke5") {
		obj["addke5"] = nil
	}

	if v, ok := d.GetOk("addke6"); ok {
		t, err := expandVpnIpsecPhase1Addke6(d, v, "addke6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke6"] = t
		}
	} else if d.HasChange("addke6") {
		obj["addke6"] = nil
	}

	if v, ok := d.GetOk("addke7"); ok {
		t, err := expandVpnIpsecPhase1Addke7(d, v, "addke7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke7"] = t
		}
	} else if d.HasChange("addke7") {
		obj["addke7"] = nil
	}

	if v, ok := d.GetOk("suite_b"); ok {
		t, err := expandVpnIpsecPhase1SuiteB(d, v, "suite_b", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["suite-b"] = t
		}
	}

	if v, ok := d.GetOk("eap"); ok {
		t, err := expandVpnIpsecPhase1Eap(d, v, "eap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap"] = t
		}
	}

	if v, ok := d.GetOk("eap_identity"); ok {
		t, err := expandVpnIpsecPhase1EapIdentity(d, v, "eap_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-identity"] = t
		}
	}

	if v, ok := d.GetOk("eap_exclude_peergrp"); ok {
		t, err := expandVpnIpsecPhase1EapExcludePeergrp(d, v, "eap_exclude_peergrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-exclude-peergrp"] = t
		}
	} else if d.HasChange("eap_exclude_peergrp") {
		obj["eap-exclude-peergrp"] = nil
	}

	if v, ok := d.GetOk("eap_cert_auth"); ok {
		t, err := expandVpnIpsecPhase1EapCertAuth(d, v, "eap_cert_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("acct_verify"); ok {
		t, err := expandVpnIpsecPhase1AcctVerify(d, v, "acct_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-verify"] = t
		}
	}

	if v, ok := d.GetOk("ppk"); ok {
		t, err := expandVpnIpsecPhase1Ppk(d, v, "ppk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok {
		t, err := expandVpnIpsecPhase1PpkSecret(d, v, "ppk_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	} else if d.HasChange("ppk_secret") {
		obj["ppk-secret"] = nil
	}

	if v, ok := d.GetOk("ppk_identity"); ok {
		t, err := expandVpnIpsecPhase1PpkIdentity(d, v, "ppk_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	} else if d.HasChange("ppk_identity") {
		obj["ppk-identity"] = nil
	}

	if v, ok := d.GetOk("wizard_type"); ok {
		t, err := expandVpnIpsecPhase1WizardType(d, v, "wizard_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wizard-type"] = t
		}
	}

	if v, ok := d.GetOk("xauthtype"); ok {
		t, err := expandVpnIpsecPhase1Xauthtype(d, v, "xauthtype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xauthtype"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok {
		t, err := expandVpnIpsecPhase1Reauth(d, v, "reauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("authusr"); ok {
		t, err := expandVpnIpsecPhase1Authusr(d, v, "authusr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusr"] = t
		}
	} else if d.HasChange("authusr") {
		obj["authusr"] = nil
	}

	if v, ok := d.GetOk("authpasswd"); ok {
		t, err := expandVpnIpsecPhase1Authpasswd(d, v, "authpasswd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authpasswd"] = t
		}
	} else if d.HasChange("authpasswd") {
		obj["authpasswd"] = nil
	}

	if v, ok := d.GetOk("group_authentication"); ok {
		t, err := expandVpnIpsecPhase1GroupAuthentication(d, v, "group_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication_secret"); ok {
		t, err := expandVpnIpsecPhase1GroupAuthenticationSecret(d, v, "group_authentication_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication-secret"] = t
		}
	} else if d.HasChange("group_authentication_secret") {
		obj["group-authentication-secret"] = nil
	}

	if v, ok := d.GetOk("authusrgrp"); ok {
		t, err := expandVpnIpsecPhase1Authusrgrp(d, v, "authusrgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusrgrp"] = t
		}
	} else if d.HasChange("authusrgrp") {
		obj["authusrgrp"] = nil
	}

	if v, ok := d.GetOk("mesh_selector_type"); ok {
		t, err := expandVpnIpsecPhase1MeshSelectorType(d, v, "mesh_selector_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-selector-type"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok {
		t, err := expandVpnIpsecPhase1IdleTimeout(d, v, "idle_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("shared_idle_timeout"); ok {
		t, err := expandVpnIpsecPhase1SharedIdleTimeout(d, v, "shared_idle_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeoutinterval"); ok {
		t, err := expandVpnIpsecPhase1IdleTimeoutinterval(d, v, "idle_timeoutinterval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeoutinterval"] = t
		}
	}

	if v, ok := d.GetOk("ha_sync_esp_seqno"); ok {
		t, err := expandVpnIpsecPhase1HaSyncEspSeqno(d, v, "ha_sync_esp_seqno", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-sync-esp-seqno"] = t
		}
	}

	if v, ok := d.GetOk("fgsp_sync"); ok {
		t, err := expandVpnIpsecPhase1FgspSync(d, v, "fgsp_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgsp-sync"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok {
		t, err := expandVpnIpsecPhase1InboundDscpCopy(d, v, "inbound_dscp_copy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("nattraversal"); ok {
		t, err := expandVpnIpsecPhase1Nattraversal(d, v, "nattraversal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nattraversal"] = t
		}
	}

	if v, ok := d.GetOk("esn"); ok {
		t, err := expandVpnIpsecPhase1Esn(d, v, "esn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esn"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation_mtu"); ok {
		t, err := expandVpnIpsecPhase1FragmentationMtu(d, v, "fragmentation_mtu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation-mtu"] = t
		}
	}

	if v, ok := d.GetOk("childless_ike"); ok {
		t, err := expandVpnIpsecPhase1ChildlessIke(d, v, "childless_ike", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["childless-ike"] = t
		}
	}

	if v, ok := d.GetOk("azure_ad_autoconnect"); ok {
		t, err := expandVpnIpsecPhase1AzureAdAutoconnect(d, v, "azure_ad_autoconnect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-ad-autoconnect"] = t
		}
	}

	if v, ok := d.GetOk("client_resume"); ok {
		t, err := expandVpnIpsecPhase1ClientResume(d, v, "client_resume", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume"] = t
		}
	}

	if v, ok := d.GetOk("client_resume_interval"); ok {
		t, err := expandVpnIpsecPhase1ClientResumeInterval(d, v, "client_resume_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume-interval"] = t
		}
	}

	if v, ok := d.GetOk("rekey"); ok {
		t, err := expandVpnIpsecPhase1Rekey(d, v, "rekey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rekey"] = t
		}
	}

	if v, ok := d.GetOk("digital_signature_auth"); ok {
		t, err := expandVpnIpsecPhase1DigitalSignatureAuth(d, v, "digital_signature_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digital-signature-auth"] = t
		}
	}

	if v, ok := d.GetOk("signature_hash_alg"); ok {
		t, err := expandVpnIpsecPhase1SignatureHashAlg(d, v, "signature_hash_alg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-hash-alg"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_format"); ok {
		t, err := expandVpnIpsecPhase1RsaSignatureFormat(d, v, "rsa_signature_format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-format"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_hash_override"); ok {
		t, err := expandVpnIpsecPhase1RsaSignatureHashOverride(d, v, "rsa_signature_hash_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-hash-override"] = t
		}
	}

	if v, ok := d.GetOk("enforce_unique_id"); ok {
		t, err := expandVpnIpsecPhase1EnforceUniqueId(d, v, "enforce_unique_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-unique-id"] = t
		}
	}

	if v, ok := d.GetOk("cert_id_validation"); ok {
		t, err := expandVpnIpsecPhase1CertIdValidation(d, v, "cert_id_validation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-id-validation"] = t
		}
	}

	if v, ok := d.GetOk("fec_egress"); ok {
		t, err := expandVpnIpsecPhase1FecEgress(d, v, "fec_egress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-egress"] = t
		}
	}

	if v, ok := d.GetOk("fec_send_timeout"); ok {
		t, err := expandVpnIpsecPhase1FecSendTimeout(d, v, "fec_send_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-send-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fec_base"); ok {
		t, err := expandVpnIpsecPhase1FecBase(d, v, "fec_base", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-base"] = t
		}
	}

	if v, ok := d.GetOk("fec_codec_string"); ok {
		new_version_map := map[string][]string{
			">=": []string{"7.0.2"},
		}
		if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
			if _, ok := d.GetOk("fec_codec"); !ok && !d.HasChange("fec_codec") {
				err := fmt.Errorf("Argument 'fec_codec_string' %s.", err)
				return nil, err
			}
		} else {
			t, err := expandVpnIpsecPhase1FecCodecString(d, v, "fec_codec_string", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fec-codec"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("fec_codec"); ok {
		new_version_map := map[string][]string{
			"=": []string{"6.4.10", "6.4.11", "6.4.12", "6.4.13", "6.4.14", "6.4.15", "7.0.0", "7.0.1"},
		}
		if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
			if _, ok := d.GetOk("fec_codec_string"); !ok && !d.HasChange("fec_codec_string") {
				err := fmt.Errorf("Argument 'fec_codec' %s.", err)
				return nil, err
			}
		} else {
			t, err := expandVpnIpsecPhase1FecCodec(d, v, "fec_codec", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fec-codec"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("fec_redundant"); ok {
		t, err := expandVpnIpsecPhase1FecRedundant(d, v, "fec_redundant", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-redundant"] = t
		}
	}

	if v, ok := d.GetOk("fec_ingress"); ok {
		t, err := expandVpnIpsecPhase1FecIngress(d, v, "fec_ingress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-ingress"] = t
		}
	}

	if v, ok := d.GetOk("fec_receive_timeout"); ok {
		t, err := expandVpnIpsecPhase1FecReceiveTimeout(d, v, "fec_receive_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fec_health_check"); ok {
		t, err := expandVpnIpsecPhase1FecHealthCheck(d, v, "fec_health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-health-check"] = t
		}
	} else if d.HasChange("fec_health_check") {
		obj["fec-health-check"] = nil
	}

	if v, ok := d.GetOk("fec_mapping_profile"); ok {
		t, err := expandVpnIpsecPhase1FecMappingProfile(d, v, "fec_mapping_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-mapping-profile"] = t
		}
	} else if d.HasChange("fec_mapping_profile") {
		obj["fec-mapping-profile"] = nil
	}

	if v, ok := d.GetOk("network_overlay"); ok {
		t, err := expandVpnIpsecPhase1NetworkOverlay(d, v, "network_overlay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-overlay"] = t
		}
	}

	if v, ok := d.GetOkExists("network_id"); ok {
		t, err := expandVpnIpsecPhase1NetworkId(d, v, "network_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-id"] = t
		}
	} else if d.HasChange("network_id") {
		obj["network-id"] = nil
	}

	if v, ok := d.GetOk("dev_id_notification"); ok {
		t, err := expandVpnIpsecPhase1DevIdNotification(d, v, "dev_id_notification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id-notification"] = t
		}
	}

	if v, ok := d.GetOk("dev_id"); ok {
		t, err := expandVpnIpsecPhase1DevId(d, v, "dev_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id"] = t
		}
	} else if d.HasChange("dev_id") {
		obj["dev-id"] = nil
	}

	if v, ok := d.GetOk("loopback_asymroute"); ok {
		t, err := expandVpnIpsecPhase1LoopbackAsymroute(d, v, "loopback_asymroute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-asymroute"] = t
		}
	}

	if v, ok := d.GetOkExists("link_cost"); ok {
		t, err := expandVpnIpsecPhase1LinkCost(d, v, "link_cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost"] = t
		}
	} else if d.HasChange("link_cost") {
		obj["link-cost"] = nil
	}

	if v, ok := d.GetOk("kms"); ok {
		t, err := expandVpnIpsecPhase1Kms(d, v, "kms", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kms"] = t
		}
	} else if d.HasChange("kms") {
		obj["kms"] = nil
	}

	if v, ok := d.GetOk("exchange_fgt_device_id"); ok {
		t, err := expandVpnIpsecPhase1ExchangeFgtDeviceId(d, v, "exchange_fgt_device_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-fgt-device-id"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_auto_linklocal"); ok {
		t, err := expandVpnIpsecPhase1Ipv6AutoLinklocal(d, v, "ipv6_auto_linklocal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-auto-linklocal"] = t
		}
	}

	if v, ok := d.GetOk("ems_sn_check"); ok {
		t, err := expandVpnIpsecPhase1EmsSnCheck(d, v, "ems_sn_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("cert_trust_store"); ok {
		t, err := expandVpnIpsecPhase1CertTrustStore(d, v, "cert_trust_store", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-trust-store"] = t
		}
	}

	if v, ok := d.GetOk("qkd"); ok {
		t, err := expandVpnIpsecPhase1Qkd(d, v, "qkd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok {
		t, err := expandVpnIpsecPhase1QkdProfile(d, v, "qkd_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	} else if d.HasChange("qkd_profile") {
		obj["qkd-profile"] = nil
	}

	if v, ok := d.GetOk("transport"); ok {
		t, err := expandVpnIpsecPhase1Transport(d, v, "transport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_esp"); ok {
		t, err := expandVpnIpsecPhase1FortinetEsp(d, v, "fortinet_esp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-esp"] = t
		}
	}

	if v, ok := d.GetOk("auto_transport_threshold"); ok {
		t, err := expandVpnIpsecPhase1AutoTransportThreshold(d, v, "auto_transport_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-transport-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fallback_tcp_threshold"); ok {
		t, err := expandVpnIpsecPhase1FallbackTcpThreshold(d, v, "fallback_tcp_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-tcp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_match"); ok {
		t, err := expandVpnIpsecPhase1RemoteGwMatch(d, v, "remote_gw_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_subnet"); ok {
		t, err := expandVpnIpsecPhase1RemoteGwSubnet(d, v, "remote_gw_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_start_ip"); ok {
		t, err := expandVpnIpsecPhase1RemoteGwStartIp(d, v, "remote_gw_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_end_ip"); ok {
		t, err := expandVpnIpsecPhase1RemoteGwEndIp(d, v, "remote_gw_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_country"); ok {
		t, err := expandVpnIpsecPhase1RemoteGwCountry(d, v, "remote_gw_country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-country"] = t
		}
	} else if d.HasChange("remote_gw_country") {
		obj["remote-gw-country"] = nil
	}

	if v, ok := d.GetOk("remote_gw_ztna_tags"); ok || d.HasChange("remote_gw_ztna_tags") {
		t, err := expandVpnIpsecPhase1RemoteGwZtnaTags(d, v, "remote_gw_ztna_tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-ztna-tags"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_match"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw6Match(d, v, "remote_gw6_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_subnet"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw6Subnet(d, v, "remote_gw6_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_start_ip"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw6StartIp(d, v, "remote_gw6_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_end_ip"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw6EndIp(d, v, "remote_gw6_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_country"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw6Country(d, v, "remote_gw6_country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-country"] = t
		}
	} else if d.HasChange("remote_gw6_country") {
		obj["remote-gw6-country"] = nil
	}

	if v, ok := d.GetOk("cert_peer_username_validation"); ok {
		t, err := expandVpnIpsecPhase1CertPeerUsernameValidation(d, v, "cert_peer_username_validation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-validation"] = t
		}
	}

	if v, ok := d.GetOk("cert_peer_username_strip"); ok {
		t, err := expandVpnIpsecPhase1CertPeerUsernameStrip(d, v, "cert_peer_username_strip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-strip"] = t
		}
	}

	return &obj, nil
}
