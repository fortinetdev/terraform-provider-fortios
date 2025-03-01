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

func resourceVpnIpsecPhase1Interface() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase1InterfaceCreate,
		Read:   resourceVpnIpsecPhase1InterfaceRead,
		Update: resourceVpnIpsecPhase1InterfaceUpdate,
		Delete: resourceVpnIpsecPhase1InterfaceDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
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
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw6": &schema.Schema{
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
			"default_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_gw_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"monitor_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"monitor_hold_down_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_hold_down_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31536000),
				Optional:     true,
			},
			"monitor_hold_down_weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor_hold_down_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"net_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passive_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_interface_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_ip_addr4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exchange_ip_addr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate_member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate_weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"packet_redistribution": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Optional:  true,
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
			"ip_fragmentation": &schema.Schema{
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
			"auto_discovery_sender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_receiver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_forwarder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_psk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_shortcuts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_crossover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_offer_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"auto_discovery_dialup_placeholder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encapsulation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encapsulation_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_local_gw4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_local_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_remote_gw4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encap_remote_gw6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vni": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777215),
				Optional:     true,
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

func resourceVpnIpsecPhase1InterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecPhase1Interface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1Interface resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecPhase1Interface(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1Interface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase1Interface")
	}

	return resourceVpnIpsecPhase1InterfaceRead(d, m)
}

func resourceVpnIpsecPhase1InterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnIpsecPhase1Interface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1Interface resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecPhase1Interface(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1Interface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase1Interface")
	}

	return resourceVpnIpsecPhase1InterfaceRead(d, m)
}

func resourceVpnIpsecPhase1InterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnIpsecPhase1Interface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase1Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase1InterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnIpsecPhase1Interface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1Interface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase1Interface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1Interface resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase1InterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIkeVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemotegwDdns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceKeylife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenVpnIpsecPhase1InterfaceCertificateName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnIpsecPhase1InterfaceCertificateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthmethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthmethodRemote(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePeertype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePeerid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDefaultGw(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDefaultGwPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceUsrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePeergrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("monitor"), "name")
}

func flattenVpnIpsecPhase1InterfaceMonitorMin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMonitorHoldDownTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNetDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceTunnelSearch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePassiveMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeInterfaceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeIpAddr4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeIpAddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAggregateMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAggregateWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfacePacketRedistribution(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceModeCfg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAssignIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAssignIpFrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4Netmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDnsMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4DnsServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceInternalDomainList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["domain_name"] = flattenVpnIpsecPhase1InterfaceInternalDomainListDomainName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "domain_name", d)
	return result
}

func flattenVpnIpsecPhase1InterfaceInternalDomainListDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4WinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4WinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if cur_v, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if cur_v, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4SplitInclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSplitIncludeService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6DnsServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if cur_v, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if cur_v, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6SplitInclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpDelayInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceUnitySupport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceBanner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIncludeLocalLan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv4SplitExclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6SplitExclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSavePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientAutoNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientKeepAlive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceBackupGateway(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["address"] = flattenVpnIpsecPhase1InterfaceBackupGatewayAddress(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "address", d)
	return result
}

func flattenVpnIpsecPhase1InterfaceBackupGatewayAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceProposal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddGwRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceLocalid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLocalidType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNegotiateTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFragmentation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpFragmentation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDpd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDpdRetrycount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceDpdRetryinterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceForticlientEnforcement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNpuOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSendCertChain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDhgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAddke7(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSuiteB(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEapIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEapExcludePeergrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEapCertAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAcctVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePpk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfacePpkIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceWizardType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceXauthtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthusr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceGroupAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAuthusrgrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceMeshSelectorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSharedIdleTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIdleTimeoutinterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceHaSyncEspSeqno(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFgspSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceInboundDscpCopy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoverySender(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryPsk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceAutoDiscoveryDialupPlaceholder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapsulation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapsulationAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapLocalGw4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapLocalGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapRemoteGw4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEncapRemoteGw6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceVni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceNattraversal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEsn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFragmentationMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceChildlessIke(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAzureAdAutoconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientResume(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceClientResumeInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceRekey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDigitalSignatureAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceSignatureHashAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRsaSignatureFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRsaSignatureHashOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEnforceUniqueId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertIdValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecEgress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecSendTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFecBase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFecCodecString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecCodec(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFecRedundant(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFecIngress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecReceiveTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFecHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFecMappingProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNetworkOverlay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceDevIdNotification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceDevId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLoopbackAsymroute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceLinkCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceKms(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceExchangeFgtDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceIpv6AutoLinklocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceEmsSnCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTagsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTagsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6Match(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6Subnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceRemoteGw6Country(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertPeerUsernameValidation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertPeerUsernameStrip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceCertTrustStore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceQkd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceQkdProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceTransport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceFortinetEsp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnIpsecPhase1InterfaceAutoTransportThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnIpsecPhase1InterfaceFallbackTcpThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectVpnIpsecPhase1Interface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenVpnIpsecPhase1InterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnIpsecPhase1InterfaceType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecPhase1InterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenVpnIpsecPhase1InterfaceIpVersion(o["ip-version"], d, "ip_version", sv)); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("ike_version", flattenVpnIpsecPhase1InterfaceIkeVersion(o["ike-version"], d, "ike_version", sv)); err != nil {
		if !fortiAPIPatch(o["ike-version"]) {
			return fmt.Errorf("Error reading ike_version: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecPhase1InterfaceLocalGw(o["local-gw"], d, "local_gw", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("local_gw6", flattenVpnIpsecPhase1InterfaceLocalGw6(o["local-gw6"], d, "local_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["local-gw6"]) {
			return fmt.Errorf("Error reading local_gw6: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecPhase1InterfaceRemoteGw(o["remote-gw"], d, "remote_gw", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("remote_gw6", flattenVpnIpsecPhase1InterfaceRemoteGw6(o["remote-gw6"], d, "remote_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6"]) {
			return fmt.Errorf("Error reading remote_gw6: %v", err)
		}
	}

	if err = d.Set("remotegw_ddns", flattenVpnIpsecPhase1InterfaceRemotegwDdns(o["remotegw-ddns"], d, "remotegw_ddns", sv)); err != nil {
		if !fortiAPIPatch(o["remotegw-ddns"]) {
			return fmt.Errorf("Error reading remotegw_ddns: %v", err)
		}
	}

	if err = d.Set("keylife", flattenVpnIpsecPhase1InterfaceKeylife(o["keylife"], d, "keylife", sv)); err != nil {
		if !fortiAPIPatch(o["keylife"]) {
			return fmt.Errorf("Error reading keylife: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("certificate", flattenVpnIpsecPhase1InterfaceCertificate(o["certificate"], d, "certificate", sv)); err != nil {
			if !fortiAPIPatch(o["certificate"]) {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("certificate"); ok {
			if err = d.Set("certificate", flattenVpnIpsecPhase1InterfaceCertificate(o["certificate"], d, "certificate", sv)); err != nil {
				if !fortiAPIPatch(o["certificate"]) {
					return fmt.Errorf("Error reading certificate: %v", err)
				}
			}
		}
	}

	if err = d.Set("authmethod", flattenVpnIpsecPhase1InterfaceAuthmethod(o["authmethod"], d, "authmethod", sv)); err != nil {
		if !fortiAPIPatch(o["authmethod"]) {
			return fmt.Errorf("Error reading authmethod: %v", err)
		}
	}

	if err = d.Set("authmethod_remote", flattenVpnIpsecPhase1InterfaceAuthmethodRemote(o["authmethod-remote"], d, "authmethod_remote", sv)); err != nil {
		if !fortiAPIPatch(o["authmethod-remote"]) {
			return fmt.Errorf("Error reading authmethod_remote: %v", err)
		}
	}

	if err = d.Set("mode", flattenVpnIpsecPhase1InterfaceMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("peertype", flattenVpnIpsecPhase1InterfacePeertype(o["peertype"], d, "peertype", sv)); err != nil {
		if !fortiAPIPatch(o["peertype"]) {
			return fmt.Errorf("Error reading peertype: %v", err)
		}
	}

	if err = d.Set("peerid", flattenVpnIpsecPhase1InterfacePeerid(o["peerid"], d, "peerid", sv)); err != nil {
		if !fortiAPIPatch(o["peerid"]) {
			return fmt.Errorf("Error reading peerid: %v", err)
		}
	}

	if err = d.Set("default_gw", flattenVpnIpsecPhase1InterfaceDefaultGw(o["default-gw"], d, "default_gw", sv)); err != nil {
		if !fortiAPIPatch(o["default-gw"]) {
			return fmt.Errorf("Error reading default_gw: %v", err)
		}
	}

	if err = d.Set("default_gw_priority", flattenVpnIpsecPhase1InterfaceDefaultGwPriority(o["default-gw-priority"], d, "default_gw_priority", sv)); err != nil {
		if !fortiAPIPatch(o["default-gw-priority"]) {
			return fmt.Errorf("Error reading default_gw_priority: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnIpsecPhase1InterfaceUsrgrp(o["usrgrp"], d, "usrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["usrgrp"]) {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("peer", flattenVpnIpsecPhase1InterfacePeer(o["peer"], d, "peer", sv)); err != nil {
		if !fortiAPIPatch(o["peer"]) {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peergrp", flattenVpnIpsecPhase1InterfacePeergrp(o["peergrp"], d, "peergrp", sv)); err != nil {
		if !fortiAPIPatch(o["peergrp"]) {
			return fmt.Errorf("Error reading peergrp: %v", err)
		}
	}

	if err = d.Set("monitor", flattenVpnIpsecPhase1InterfaceMonitor(o["monitor"], d, "monitor", sv)); err != nil {
		if !fortiAPIPatch(o["monitor"]) {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("monitor_min", flattenVpnIpsecPhase1InterfaceMonitorMin(o["monitor-min"], d, "monitor_min", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-min"]) {
			return fmt.Errorf("Error reading monitor_min: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_type", flattenVpnIpsecPhase1InterfaceMonitorHoldDownType(o["monitor-hold-down-type"], d, "monitor_hold_down_type", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-hold-down-type"]) {
			return fmt.Errorf("Error reading monitor_hold_down_type: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_delay", flattenVpnIpsecPhase1InterfaceMonitorHoldDownDelay(o["monitor-hold-down-delay"], d, "monitor_hold_down_delay", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-hold-down-delay"]) {
			return fmt.Errorf("Error reading monitor_hold_down_delay: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_weekday", flattenVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(o["monitor-hold-down-weekday"], d, "monitor_hold_down_weekday", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-hold-down-weekday"]) {
			return fmt.Errorf("Error reading monitor_hold_down_weekday: %v", err)
		}
	}

	if err = d.Set("monitor_hold_down_time", flattenVpnIpsecPhase1InterfaceMonitorHoldDownTime(o["monitor-hold-down-time"], d, "monitor_hold_down_time", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-hold-down-time"]) {
			return fmt.Errorf("Error reading monitor_hold_down_time: %v", err)
		}
	}

	if err = d.Set("net_device", flattenVpnIpsecPhase1InterfaceNetDevice(o["net-device"], d, "net_device", sv)); err != nil {
		if !fortiAPIPatch(o["net-device"]) {
			return fmt.Errorf("Error reading net_device: %v", err)
		}
	}

	if err = d.Set("tunnel_search", flattenVpnIpsecPhase1InterfaceTunnelSearch(o["tunnel-search"], d, "tunnel_search", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-search"]) {
			return fmt.Errorf("Error reading tunnel_search: %v", err)
		}
	}

	if err = d.Set("passive_mode", flattenVpnIpsecPhase1InterfacePassiveMode(o["passive-mode"], d, "passive_mode", sv)); err != nil {
		if !fortiAPIPatch(o["passive-mode"]) {
			return fmt.Errorf("Error reading passive_mode: %v", err)
		}
	}

	if err = d.Set("exchange_interface_ip", flattenVpnIpsecPhase1InterfaceExchangeInterfaceIp(o["exchange-interface-ip"], d, "exchange_interface_ip", sv)); err != nil {
		if !fortiAPIPatch(o["exchange-interface-ip"]) {
			return fmt.Errorf("Error reading exchange_interface_ip: %v", err)
		}
	}

	if err = d.Set("exchange_ip_addr4", flattenVpnIpsecPhase1InterfaceExchangeIpAddr4(o["exchange-ip-addr4"], d, "exchange_ip_addr4", sv)); err != nil {
		if !fortiAPIPatch(o["exchange-ip-addr4"]) {
			return fmt.Errorf("Error reading exchange_ip_addr4: %v", err)
		}
	}

	if err = d.Set("exchange_ip_addr6", flattenVpnIpsecPhase1InterfaceExchangeIpAddr6(o["exchange-ip-addr6"], d, "exchange_ip_addr6", sv)); err != nil {
		if !fortiAPIPatch(o["exchange-ip-addr6"]) {
			return fmt.Errorf("Error reading exchange_ip_addr6: %v", err)
		}
	}

	if err = d.Set("aggregate_member", flattenVpnIpsecPhase1InterfaceAggregateMember(o["aggregate-member"], d, "aggregate_member", sv)); err != nil {
		if !fortiAPIPatch(o["aggregate-member"]) {
			return fmt.Errorf("Error reading aggregate_member: %v", err)
		}
	}

	if err = d.Set("aggregate_weight", flattenVpnIpsecPhase1InterfaceAggregateWeight(o["aggregate-weight"], d, "aggregate_weight", sv)); err != nil {
		if !fortiAPIPatch(o["aggregate-weight"]) {
			return fmt.Errorf("Error reading aggregate_weight: %v", err)
		}
	}

	if err = d.Set("packet_redistribution", flattenVpnIpsecPhase1InterfacePacketRedistribution(o["packet-redistribution"], d, "packet_redistribution", sv)); err != nil {
		if !fortiAPIPatch(o["packet-redistribution"]) {
			return fmt.Errorf("Error reading packet_redistribution: %v", err)
		}
	}

	if err = d.Set("mode_cfg", flattenVpnIpsecPhase1InterfaceModeCfg(o["mode-cfg"], d, "mode_cfg", sv)); err != nil {
		if !fortiAPIPatch(o["mode-cfg"]) {
			return fmt.Errorf("Error reading mode_cfg: %v", err)
		}
	}

	if err = d.Set("mode_cfg_allow_client_selector", flattenVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(o["mode-cfg-allow-client-selector"], d, "mode_cfg_allow_client_selector", sv)); err != nil {
		if !fortiAPIPatch(o["mode-cfg-allow-client-selector"]) {
			return fmt.Errorf("Error reading mode_cfg_allow_client_selector: %v", err)
		}
	}

	if err = d.Set("assign_ip", flattenVpnIpsecPhase1InterfaceAssignIp(o["assign-ip"], d, "assign_ip", sv)); err != nil {
		if !fortiAPIPatch(o["assign-ip"]) {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("assign_ip_from", flattenVpnIpsecPhase1InterfaceAssignIpFrom(o["assign-ip-from"], d, "assign_ip_from", sv)); err != nil {
		if !fortiAPIPatch(o["assign-ip-from"]) {
			return fmt.Errorf("Error reading assign_ip_from: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenVpnIpsecPhase1InterfaceIpv4StartIp(o["ipv4-start-ip"], d, "ipv4_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-start-ip"]) {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenVpnIpsecPhase1InterfaceIpv4EndIp(o["ipv4-end-ip"], d, "ipv4_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-end-ip"]) {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_netmask", flattenVpnIpsecPhase1InterfaceIpv4Netmask(o["ipv4-netmask"], d, "ipv4_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-netmask"]) {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenVpnIpsecPhase1InterfaceDhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ra-giaddr"]) {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp6-ra-linkaddr"]) {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenVpnIpsecPhase1InterfaceDnsMode(o["dns-mode"], d, "dns_mode", sv)); err != nil {
		if !fortiAPIPatch(o["dns-mode"]) {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server1", flattenVpnIpsecPhase1InterfaceIpv4DnsServer1(o["ipv4-dns-server1"], d, "ipv4_dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server1"]) {
			return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server2", flattenVpnIpsecPhase1InterfaceIpv4DnsServer2(o["ipv4-dns-server2"], d, "ipv4_dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server2"]) {
			return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server3", flattenVpnIpsecPhase1InterfaceIpv4DnsServer3(o["ipv4-dns-server3"], d, "ipv4_dns_server3", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server3"]) {
			return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internal_domain_list", flattenVpnIpsecPhase1InterfaceInternalDomainList(o["internal-domain-list"], d, "internal_domain_list", sv)); err != nil {
			if !fortiAPIPatch(o["internal-domain-list"]) {
				return fmt.Errorf("Error reading internal_domain_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internal_domain_list"); ok {
			if err = d.Set("internal_domain_list", flattenVpnIpsecPhase1InterfaceInternalDomainList(o["internal-domain-list"], d, "internal_domain_list", sv)); err != nil {
				if !fortiAPIPatch(o["internal-domain-list"]) {
					return fmt.Errorf("Error reading internal_domain_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_wins_server1", flattenVpnIpsecPhase1InterfaceIpv4WinsServer1(o["ipv4-wins-server1"], d, "ipv4_wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-wins-server1"]) {
			return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server2", flattenVpnIpsecPhase1InterfaceIpv4WinsServer2(o["ipv4-wins-server2"], d, "ipv4_wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-wins-server2"]) {
			return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1InterfaceIpv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range", sv)); err != nil {
			if !fortiAPIPatch(o["ipv4-exclude-range"]) {
				return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv4_exclude_range"); ok {
			if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1InterfaceIpv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range", sv)); err != nil {
				if !fortiAPIPatch(o["ipv4-exclude-range"]) {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_split_include", flattenVpnIpsecPhase1InterfaceIpv4SplitInclude(o["ipv4-split-include"], d, "ipv4_split_include", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-split-include"]) {
			return fmt.Errorf("Error reading ipv4_split_include: %v", err)
		}
	}

	if err = d.Set("split_include_service", flattenVpnIpsecPhase1InterfaceSplitIncludeService(o["split-include-service"], d, "split_include_service", sv)); err != nil {
		if !fortiAPIPatch(o["split-include-service"]) {
			return fmt.Errorf("Error reading split_include_service: %v", err)
		}
	}

	if err = d.Set("ipv4_name", flattenVpnIpsecPhase1InterfaceIpv4Name(o["ipv4-name"], d, "ipv4_name", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-name"]) {
			return fmt.Errorf("Error reading ipv4_name: %v", err)
		}
	}

	if err = d.Set("ipv6_start_ip", flattenVpnIpsecPhase1InterfaceIpv6StartIp(o["ipv6-start-ip"], d, "ipv6_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-start-ip"]) {
			return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_end_ip", flattenVpnIpsecPhase1InterfaceIpv6EndIp(o["ipv6-end-ip"], d, "ipv6_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-end-ip"]) {
			return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix", flattenVpnIpsecPhase1InterfaceIpv6Prefix(o["ipv6-prefix"], d, "ipv6_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-prefix"]) {
			return fmt.Errorf("Error reading ipv6_prefix: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnIpsecPhase1InterfaceIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server1"]) {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnIpsecPhase1InterfaceIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server2"]) {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server3", flattenVpnIpsecPhase1InterfaceIpv6DnsServer3(o["ipv6-dns-server3"], d, "ipv6_dns_server3", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server3"]) {
			return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1InterfaceIpv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range", sv)); err != nil {
			if !fortiAPIPatch(o["ipv6-exclude-range"]) {
				return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_exclude_range"); ok {
			if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1InterfaceIpv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range", sv)); err != nil {
				if !fortiAPIPatch(o["ipv6-exclude-range"]) {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_split_include", flattenVpnIpsecPhase1InterfaceIpv6SplitInclude(o["ipv6-split-include"], d, "ipv6_split_include", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-split-include"]) {
			return fmt.Errorf("Error reading ipv6_split_include: %v", err)
		}
	}

	if err = d.Set("ipv6_name", flattenVpnIpsecPhase1InterfaceIpv6Name(o["ipv6-name"], d, "ipv6_name", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-name"]) {
			return fmt.Errorf("Error reading ipv6_name: %v", err)
		}
	}

	if err = d.Set("ip_delay_interval", flattenVpnIpsecPhase1InterfaceIpDelayInterval(o["ip-delay-interval"], d, "ip_delay_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ip-delay-interval"]) {
			return fmt.Errorf("Error reading ip_delay_interval: %v", err)
		}
	}

	if err = d.Set("unity_support", flattenVpnIpsecPhase1InterfaceUnitySupport(o["unity-support"], d, "unity_support", sv)); err != nil {
		if !fortiAPIPatch(o["unity-support"]) {
			return fmt.Errorf("Error reading unity_support: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnIpsecPhase1InterfaceDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("banner", flattenVpnIpsecPhase1InterfaceBanner(o["banner"], d, "banner", sv)); err != nil {
		if !fortiAPIPatch(o["banner"]) {
			return fmt.Errorf("Error reading banner: %v", err)
		}
	}

	if err = d.Set("include_local_lan", flattenVpnIpsecPhase1InterfaceIncludeLocalLan(o["include-local-lan"], d, "include_local_lan", sv)); err != nil {
		if !fortiAPIPatch(o["include-local-lan"]) {
			return fmt.Errorf("Error reading include_local_lan: %v", err)
		}
	}

	if err = d.Set("ipv4_split_exclude", flattenVpnIpsecPhase1InterfaceIpv4SplitExclude(o["ipv4-split-exclude"], d, "ipv4_split_exclude", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-split-exclude"]) {
			return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv6_split_exclude", flattenVpnIpsecPhase1InterfaceIpv6SplitExclude(o["ipv6-split-exclude"], d, "ipv6_split_exclude", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-split-exclude"]) {
			return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnIpsecPhase1InterfaceSavePassword(o["save-password"], d, "save_password", sv)); err != nil {
		if !fortiAPIPatch(o["save-password"]) {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("client_auto_negotiate", flattenVpnIpsecPhase1InterfaceClientAutoNegotiate(o["client-auto-negotiate"], d, "client_auto_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["client-auto-negotiate"]) {
			return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
		}
	}

	if err = d.Set("client_keep_alive", flattenVpnIpsecPhase1InterfaceClientKeepAlive(o["client-keep-alive"], d, "client_keep_alive", sv)); err != nil {
		if !fortiAPIPatch(o["client-keep-alive"]) {
			return fmt.Errorf("Error reading client_keep_alive: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("backup_gateway", flattenVpnIpsecPhase1InterfaceBackupGateway(o["backup-gateway"], d, "backup_gateway", sv)); err != nil {
			if !fortiAPIPatch(o["backup-gateway"]) {
				return fmt.Errorf("Error reading backup_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("backup_gateway"); ok {
			if err = d.Set("backup_gateway", flattenVpnIpsecPhase1InterfaceBackupGateway(o["backup-gateway"], d, "backup_gateway", sv)); err != nil {
				if !fortiAPIPatch(o["backup-gateway"]) {
					return fmt.Errorf("Error reading backup_gateway: %v", err)
				}
			}
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase1InterfaceProposal(o["proposal"], d, "proposal", sv)); err != nil {
		if !fortiAPIPatch(o["proposal"]) {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase1InterfaceAddRoute(o["add-route"], d, "add_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-route"]) {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("add_gw_route", flattenVpnIpsecPhase1InterfaceAddGwRoute(o["add-gw-route"], d, "add_gw_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-gw-route"]) {
			return fmt.Errorf("Error reading add_gw_route: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase1InterfaceKeepalive(o["keepalive"], d, "keepalive", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("distance", flattenVpnIpsecPhase1InterfaceDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", flattenVpnIpsecPhase1InterfacePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("localid", flattenVpnIpsecPhase1InterfaceLocalid(o["localid"], d, "localid", sv)); err != nil {
		if !fortiAPIPatch(o["localid"]) {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("localid_type", flattenVpnIpsecPhase1InterfaceLocalidType(o["localid-type"], d, "localid_type", sv)); err != nil {
		if !fortiAPIPatch(o["localid-type"]) {
			return fmt.Errorf("Error reading localid_type: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase1InterfaceAutoNegotiate(o["auto-negotiate"], d, "auto_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["auto-negotiate"]) {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("negotiate_timeout", flattenVpnIpsecPhase1InterfaceNegotiateTimeout(o["negotiate-timeout"], d, "negotiate_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["negotiate-timeout"]) {
			return fmt.Errorf("Error reading negotiate_timeout: %v", err)
		}
	}

	if err = d.Set("fragmentation", flattenVpnIpsecPhase1InterfaceFragmentation(o["fragmentation"], d, "fragmentation", sv)); err != nil {
		if !fortiAPIPatch(o["fragmentation"]) {
			return fmt.Errorf("Error reading fragmentation: %v", err)
		}
	}

	if err = d.Set("ip_fragmentation", flattenVpnIpsecPhase1InterfaceIpFragmentation(o["ip-fragmentation"], d, "ip_fragmentation", sv)); err != nil {
		if !fortiAPIPatch(o["ip-fragmentation"]) {
			return fmt.Errorf("Error reading ip_fragmentation: %v", err)
		}
	}

	if err = d.Set("dpd", flattenVpnIpsecPhase1InterfaceDpd(o["dpd"], d, "dpd", sv)); err != nil {
		if !fortiAPIPatch(o["dpd"]) {
			return fmt.Errorf("Error reading dpd: %v", err)
		}
	}

	if err = d.Set("dpd_retrycount", flattenVpnIpsecPhase1InterfaceDpdRetrycount(o["dpd-retrycount"], d, "dpd_retrycount", sv)); err != nil {
		if !fortiAPIPatch(o["dpd-retrycount"]) {
			return fmt.Errorf("Error reading dpd_retrycount: %v", err)
		}
	}

	if err = d.Set("dpd_retryinterval", flattenVpnIpsecPhase1InterfaceDpdRetryinterval(o["dpd-retryinterval"], d, "dpd_retryinterval", sv)); err != nil {
		if !fortiAPIPatch(o["dpd-retryinterval"]) {
			return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
		}
	}

	if err = d.Set("forticlient_enforcement", flattenVpnIpsecPhase1InterfaceForticlientEnforcement(o["forticlient-enforcement"], d, "forticlient_enforcement", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-enforcement"]) {
			return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase1InterfaceComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("npu_offload", flattenVpnIpsecPhase1InterfaceNpuOffload(o["npu-offload"], d, "npu_offload", sv)); err != nil {
		if !fortiAPIPatch(o["npu-offload"]) {
			return fmt.Errorf("Error reading npu_offload: %v", err)
		}
	}

	if err = d.Set("send_cert_chain", flattenVpnIpsecPhase1InterfaceSendCertChain(o["send-cert-chain"], d, "send_cert_chain", sv)); err != nil {
		if !fortiAPIPatch(o["send-cert-chain"]) {
			return fmt.Errorf("Error reading send_cert_chain: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase1InterfaceDhgrp(o["dhgrp"], d, "dhgrp", sv)); err != nil {
		if !fortiAPIPatch(o["dhgrp"]) {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("addke1", flattenVpnIpsecPhase1InterfaceAddke1(o["addke1"], d, "addke1", sv)); err != nil {
		if !fortiAPIPatch(o["addke1"]) {
			return fmt.Errorf("Error reading addke1: %v", err)
		}
	}

	if err = d.Set("addke2", flattenVpnIpsecPhase1InterfaceAddke2(o["addke2"], d, "addke2", sv)); err != nil {
		if !fortiAPIPatch(o["addke2"]) {
			return fmt.Errorf("Error reading addke2: %v", err)
		}
	}

	if err = d.Set("addke3", flattenVpnIpsecPhase1InterfaceAddke3(o["addke3"], d, "addke3", sv)); err != nil {
		if !fortiAPIPatch(o["addke3"]) {
			return fmt.Errorf("Error reading addke3: %v", err)
		}
	}

	if err = d.Set("addke4", flattenVpnIpsecPhase1InterfaceAddke4(o["addke4"], d, "addke4", sv)); err != nil {
		if !fortiAPIPatch(o["addke4"]) {
			return fmt.Errorf("Error reading addke4: %v", err)
		}
	}

	if err = d.Set("addke5", flattenVpnIpsecPhase1InterfaceAddke5(o["addke5"], d, "addke5", sv)); err != nil {
		if !fortiAPIPatch(o["addke5"]) {
			return fmt.Errorf("Error reading addke5: %v", err)
		}
	}

	if err = d.Set("addke6", flattenVpnIpsecPhase1InterfaceAddke6(o["addke6"], d, "addke6", sv)); err != nil {
		if !fortiAPIPatch(o["addke6"]) {
			return fmt.Errorf("Error reading addke6: %v", err)
		}
	}

	if err = d.Set("addke7", flattenVpnIpsecPhase1InterfaceAddke7(o["addke7"], d, "addke7", sv)); err != nil {
		if !fortiAPIPatch(o["addke7"]) {
			return fmt.Errorf("Error reading addke7: %v", err)
		}
	}

	if err = d.Set("suite_b", flattenVpnIpsecPhase1InterfaceSuiteB(o["suite-b"], d, "suite_b", sv)); err != nil {
		if !fortiAPIPatch(o["suite-b"]) {
			return fmt.Errorf("Error reading suite_b: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnIpsecPhase1InterfaceEap(o["eap"], d, "eap", sv)); err != nil {
		if !fortiAPIPatch(o["eap"]) {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_identity", flattenVpnIpsecPhase1InterfaceEapIdentity(o["eap-identity"], d, "eap_identity", sv)); err != nil {
		if !fortiAPIPatch(o["eap-identity"]) {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("eap_exclude_peergrp", flattenVpnIpsecPhase1InterfaceEapExcludePeergrp(o["eap-exclude-peergrp"], d, "eap_exclude_peergrp", sv)); err != nil {
		if !fortiAPIPatch(o["eap-exclude-peergrp"]) {
			return fmt.Errorf("Error reading eap_exclude_peergrp: %v", err)
		}
	}

	if err = d.Set("eap_cert_auth", flattenVpnIpsecPhase1InterfaceEapCertAuth(o["eap-cert-auth"], d, "eap_cert_auth", sv)); err != nil {
		if !fortiAPIPatch(o["eap-cert-auth"]) {
			return fmt.Errorf("Error reading eap_cert_auth: %v", err)
		}
	}

	if err = d.Set("acct_verify", flattenVpnIpsecPhase1InterfaceAcctVerify(o["acct-verify"], d, "acct_verify", sv)); err != nil {
		if !fortiAPIPatch(o["acct-verify"]) {
			return fmt.Errorf("Error reading acct_verify: %v", err)
		}
	}

	if err = d.Set("ppk", flattenVpnIpsecPhase1InterfacePpk(o["ppk"], d, "ppk", sv)); err != nil {
		if !fortiAPIPatch(o["ppk"]) {
			return fmt.Errorf("Error reading ppk: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenVpnIpsecPhase1InterfacePpkIdentity(o["ppk-identity"], d, "ppk_identity", sv)); err != nil {
		if !fortiAPIPatch(o["ppk-identity"]) {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("wizard_type", flattenVpnIpsecPhase1InterfaceWizardType(o["wizard-type"], d, "wizard_type", sv)); err != nil {
		if !fortiAPIPatch(o["wizard-type"]) {
			return fmt.Errorf("Error reading wizard_type: %v", err)
		}
	}

	if err = d.Set("xauthtype", flattenVpnIpsecPhase1InterfaceXauthtype(o["xauthtype"], d, "xauthtype", sv)); err != nil {
		if !fortiAPIPatch(o["xauthtype"]) {
			return fmt.Errorf("Error reading xauthtype: %v", err)
		}
	}

	if err = d.Set("reauth", flattenVpnIpsecPhase1InterfaceReauth(o["reauth"], d, "reauth", sv)); err != nil {
		if !fortiAPIPatch(o["reauth"]) {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("authusr", flattenVpnIpsecPhase1InterfaceAuthusr(o["authusr"], d, "authusr", sv)); err != nil {
		if !fortiAPIPatch(o["authusr"]) {
			return fmt.Errorf("Error reading authusr: %v", err)
		}
	}

	if err = d.Set("group_authentication", flattenVpnIpsecPhase1InterfaceGroupAuthentication(o["group-authentication"], d, "group_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["group-authentication"]) {
			return fmt.Errorf("Error reading group_authentication: %v", err)
		}
	}

	if err = d.Set("authusrgrp", flattenVpnIpsecPhase1InterfaceAuthusrgrp(o["authusrgrp"], d, "authusrgrp", sv)); err != nil {
		if !fortiAPIPatch(o["authusrgrp"]) {
			return fmt.Errorf("Error reading authusrgrp: %v", err)
		}
	}

	if err = d.Set("mesh_selector_type", flattenVpnIpsecPhase1InterfaceMeshSelectorType(o["mesh-selector-type"], d, "mesh_selector_type", sv)); err != nil {
		if !fortiAPIPatch(o["mesh-selector-type"]) {
			return fmt.Errorf("Error reading mesh_selector_type: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnIpsecPhase1InterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("shared_idle_timeout", flattenVpnIpsecPhase1InterfaceSharedIdleTimeout(o["shared-idle-timeout"], d, "shared_idle_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["shared-idle-timeout"]) {
			return fmt.Errorf("Error reading shared_idle_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeoutinterval", flattenVpnIpsecPhase1InterfaceIdleTimeoutinterval(o["idle-timeoutinterval"], d, "idle_timeoutinterval", sv)); err != nil {
		if !fortiAPIPatch(o["idle-timeoutinterval"]) {
			return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
		}
	}

	if err = d.Set("ha_sync_esp_seqno", flattenVpnIpsecPhase1InterfaceHaSyncEspSeqno(o["ha-sync-esp-seqno"], d, "ha_sync_esp_seqno", sv)); err != nil {
		if !fortiAPIPatch(o["ha-sync-esp-seqno"]) {
			return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
		}
	}

	if err = d.Set("fgsp_sync", flattenVpnIpsecPhase1InterfaceFgspSync(o["fgsp-sync"], d, "fgsp_sync", sv)); err != nil {
		if !fortiAPIPatch(o["fgsp-sync"]) {
			return fmt.Errorf("Error reading fgsp_sync: %v", err)
		}
	}

	if err = d.Set("inbound_dscp_copy", flattenVpnIpsecPhase1InterfaceInboundDscpCopy(o["inbound-dscp-copy"], d, "inbound_dscp_copy", sv)); err != nil {
		if !fortiAPIPatch(o["inbound-dscp-copy"]) {
			return fmt.Errorf("Error reading inbound_dscp_copy: %v", err)
		}
	}

	if err = d.Set("auto_discovery_sender", flattenVpnIpsecPhase1InterfaceAutoDiscoverySender(o["auto-discovery-sender"], d, "auto_discovery_sender", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-sender"]) {
			return fmt.Errorf("Error reading auto_discovery_sender: %v", err)
		}
	}

	if err = d.Set("auto_discovery_receiver", flattenVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(o["auto-discovery-receiver"], d, "auto_discovery_receiver", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-receiver"]) {
			return fmt.Errorf("Error reading auto_discovery_receiver: %v", err)
		}
	}

	if err = d.Set("auto_discovery_forwarder", flattenVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(o["auto-discovery-forwarder"], d, "auto_discovery_forwarder", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-forwarder"]) {
			return fmt.Errorf("Error reading auto_discovery_forwarder: %v", err)
		}
	}

	if err = d.Set("auto_discovery_psk", flattenVpnIpsecPhase1InterfaceAutoDiscoveryPsk(o["auto-discovery-psk"], d, "auto_discovery_psk", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-psk"]) {
			return fmt.Errorf("Error reading auto_discovery_psk: %v", err)
		}
	}

	if err = d.Set("auto_discovery_shortcuts", flattenVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(o["auto-discovery-shortcuts"], d, "auto_discovery_shortcuts", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-shortcuts"]) {
			return fmt.Errorf("Error reading auto_discovery_shortcuts: %v", err)
		}
	}

	if err = d.Set("auto_discovery_crossover", flattenVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(o["auto-discovery-crossover"], d, "auto_discovery_crossover", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-crossover"]) {
			return fmt.Errorf("Error reading auto_discovery_crossover: %v", err)
		}
	}

	if err = d.Set("auto_discovery_offer_interval", flattenVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(o["auto-discovery-offer-interval"], d, "auto_discovery_offer_interval", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-offer-interval"]) {
			return fmt.Errorf("Error reading auto_discovery_offer_interval: %v", err)
		}
	}

	if err = d.Set("auto_discovery_dialup_placeholder", flattenVpnIpsecPhase1InterfaceAutoDiscoveryDialupPlaceholder(o["auto-discovery-dialup-placeholder"], d, "auto_discovery_dialup_placeholder", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-dialup-placeholder"]) {
			return fmt.Errorf("Error reading auto_discovery_dialup_placeholder: %v", err)
		}
	}

	if err = d.Set("encapsulation", flattenVpnIpsecPhase1InterfaceEncapsulation(o["encapsulation"], d, "encapsulation", sv)); err != nil {
		if !fortiAPIPatch(o["encapsulation"]) {
			return fmt.Errorf("Error reading encapsulation: %v", err)
		}
	}

	if err = d.Set("encapsulation_address", flattenVpnIpsecPhase1InterfaceEncapsulationAddress(o["encapsulation-address"], d, "encapsulation_address", sv)); err != nil {
		if !fortiAPIPatch(o["encapsulation-address"]) {
			return fmt.Errorf("Error reading encapsulation_address: %v", err)
		}
	}

	if err = d.Set("encap_local_gw4", flattenVpnIpsecPhase1InterfaceEncapLocalGw4(o["encap-local-gw4"], d, "encap_local_gw4", sv)); err != nil {
		if !fortiAPIPatch(o["encap-local-gw4"]) {
			return fmt.Errorf("Error reading encap_local_gw4: %v", err)
		}
	}

	if err = d.Set("encap_local_gw6", flattenVpnIpsecPhase1InterfaceEncapLocalGw6(o["encap-local-gw6"], d, "encap_local_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["encap-local-gw6"]) {
			return fmt.Errorf("Error reading encap_local_gw6: %v", err)
		}
	}

	if err = d.Set("encap_remote_gw4", flattenVpnIpsecPhase1InterfaceEncapRemoteGw4(o["encap-remote-gw4"], d, "encap_remote_gw4", sv)); err != nil {
		if !fortiAPIPatch(o["encap-remote-gw4"]) {
			return fmt.Errorf("Error reading encap_remote_gw4: %v", err)
		}
	}

	if err = d.Set("encap_remote_gw6", flattenVpnIpsecPhase1InterfaceEncapRemoteGw6(o["encap-remote-gw6"], d, "encap_remote_gw6", sv)); err != nil {
		if !fortiAPIPatch(o["encap-remote-gw6"]) {
			return fmt.Errorf("Error reading encap_remote_gw6: %v", err)
		}
	}

	if err = d.Set("vni", flattenVpnIpsecPhase1InterfaceVni(o["vni"], d, "vni", sv)); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("nattraversal", flattenVpnIpsecPhase1InterfaceNattraversal(o["nattraversal"], d, "nattraversal", sv)); err != nil {
		if !fortiAPIPatch(o["nattraversal"]) {
			return fmt.Errorf("Error reading nattraversal: %v", err)
		}
	}

	if err = d.Set("esn", flattenVpnIpsecPhase1InterfaceEsn(o["esn"], d, "esn", sv)); err != nil {
		if !fortiAPIPatch(o["esn"]) {
			return fmt.Errorf("Error reading esn: %v", err)
		}
	}

	if err = d.Set("fragmentation_mtu", flattenVpnIpsecPhase1InterfaceFragmentationMtu(o["fragmentation-mtu"], d, "fragmentation_mtu", sv)); err != nil {
		if !fortiAPIPatch(o["fragmentation-mtu"]) {
			return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
		}
	}

	if err = d.Set("childless_ike", flattenVpnIpsecPhase1InterfaceChildlessIke(o["childless-ike"], d, "childless_ike", sv)); err != nil {
		if !fortiAPIPatch(o["childless-ike"]) {
			return fmt.Errorf("Error reading childless_ike: %v", err)
		}
	}

	if err = d.Set("azure_ad_autoconnect", flattenVpnIpsecPhase1InterfaceAzureAdAutoconnect(o["azure-ad-autoconnect"], d, "azure_ad_autoconnect", sv)); err != nil {
		if !fortiAPIPatch(o["azure-ad-autoconnect"]) {
			return fmt.Errorf("Error reading azure_ad_autoconnect: %v", err)
		}
	}

	if err = d.Set("client_resume", flattenVpnIpsecPhase1InterfaceClientResume(o["client-resume"], d, "client_resume", sv)); err != nil {
		if !fortiAPIPatch(o["client-resume"]) {
			return fmt.Errorf("Error reading client_resume: %v", err)
		}
	}

	if err = d.Set("client_resume_interval", flattenVpnIpsecPhase1InterfaceClientResumeInterval(o["client-resume-interval"], d, "client_resume_interval", sv)); err != nil {
		if !fortiAPIPatch(o["client-resume-interval"]) {
			return fmt.Errorf("Error reading client_resume_interval: %v", err)
		}
	}

	if err = d.Set("rekey", flattenVpnIpsecPhase1InterfaceRekey(o["rekey"], d, "rekey", sv)); err != nil {
		if !fortiAPIPatch(o["rekey"]) {
			return fmt.Errorf("Error reading rekey: %v", err)
		}
	}

	if err = d.Set("digital_signature_auth", flattenVpnIpsecPhase1InterfaceDigitalSignatureAuth(o["digital-signature-auth"], d, "digital_signature_auth", sv)); err != nil {
		if !fortiAPIPatch(o["digital-signature-auth"]) {
			return fmt.Errorf("Error reading digital_signature_auth: %v", err)
		}
	}

	if err = d.Set("signature_hash_alg", flattenVpnIpsecPhase1InterfaceSignatureHashAlg(o["signature-hash-alg"], d, "signature_hash_alg", sv)); err != nil {
		if !fortiAPIPatch(o["signature-hash-alg"]) {
			return fmt.Errorf("Error reading signature_hash_alg: %v", err)
		}
	}

	if err = d.Set("rsa_signature_format", flattenVpnIpsecPhase1InterfaceRsaSignatureFormat(o["rsa-signature-format"], d, "rsa_signature_format", sv)); err != nil {
		if !fortiAPIPatch(o["rsa-signature-format"]) {
			return fmt.Errorf("Error reading rsa_signature_format: %v", err)
		}
	}

	if err = d.Set("rsa_signature_hash_override", flattenVpnIpsecPhase1InterfaceRsaSignatureHashOverride(o["rsa-signature-hash-override"], d, "rsa_signature_hash_override", sv)); err != nil {
		if !fortiAPIPatch(o["rsa-signature-hash-override"]) {
			return fmt.Errorf("Error reading rsa_signature_hash_override: %v", err)
		}
	}

	if err = d.Set("enforce_unique_id", flattenVpnIpsecPhase1InterfaceEnforceUniqueId(o["enforce-unique-id"], d, "enforce_unique_id", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-unique-id"]) {
			return fmt.Errorf("Error reading enforce_unique_id: %v", err)
		}
	}

	if err = d.Set("cert_id_validation", flattenVpnIpsecPhase1InterfaceCertIdValidation(o["cert-id-validation"], d, "cert_id_validation", sv)); err != nil {
		if !fortiAPIPatch(o["cert-id-validation"]) {
			return fmt.Errorf("Error reading cert_id_validation: %v", err)
		}
	}

	if err = d.Set("fec_egress", flattenVpnIpsecPhase1InterfaceFecEgress(o["fec-egress"], d, "fec_egress", sv)); err != nil {
		if !fortiAPIPatch(o["fec-egress"]) {
			return fmt.Errorf("Error reading fec_egress: %v", err)
		}
	}

	if err = d.Set("fec_send_timeout", flattenVpnIpsecPhase1InterfaceFecSendTimeout(o["fec-send-timeout"], d, "fec_send_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["fec-send-timeout"]) {
			return fmt.Errorf("Error reading fec_send_timeout: %v", err)
		}
	}

	if err = d.Set("fec_base", flattenVpnIpsecPhase1InterfaceFecBase(o["fec-base"], d, "fec_base", sv)); err != nil {
		if !fortiAPIPatch(o["fec-base"]) {
			return fmt.Errorf("Error reading fec_base: %v", err)
		}
	}

	if _, ok := o["fec-codec"].(string); ok {
		if err = d.Set("fec_codec_string", flattenVpnIpsecPhase1InterfaceFecCodecString(o["fec-codec"], d, "fec_codec_string", sv)); err != nil {
			if !fortiAPIPatch(o["fec-codec"]) {
				return fmt.Errorf("Error reading fec_codec_string: %v", err)
			}
		}
	}

	if _, ok := o["fec-codec"].(float64); ok {
		if err = d.Set("fec_codec", flattenVpnIpsecPhase1InterfaceFecCodec(o["fec-codec"], d, "fec_codec", sv)); err != nil {
			if !fortiAPIPatch(o["fec-codec"]) {
				return fmt.Errorf("Error reading fec_codec: %v", err)
			}
		}
	}

	if err = d.Set("fec_redundant", flattenVpnIpsecPhase1InterfaceFecRedundant(o["fec-redundant"], d, "fec_redundant", sv)); err != nil {
		if !fortiAPIPatch(o["fec-redundant"]) {
			return fmt.Errorf("Error reading fec_redundant: %v", err)
		}
	}

	if err = d.Set("fec_ingress", flattenVpnIpsecPhase1InterfaceFecIngress(o["fec-ingress"], d, "fec_ingress", sv)); err != nil {
		if !fortiAPIPatch(o["fec-ingress"]) {
			return fmt.Errorf("Error reading fec_ingress: %v", err)
		}
	}

	if err = d.Set("fec_receive_timeout", flattenVpnIpsecPhase1InterfaceFecReceiveTimeout(o["fec-receive-timeout"], d, "fec_receive_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["fec-receive-timeout"]) {
			return fmt.Errorf("Error reading fec_receive_timeout: %v", err)
		}
	}

	if err = d.Set("fec_health_check", flattenVpnIpsecPhase1InterfaceFecHealthCheck(o["fec-health-check"], d, "fec_health_check", sv)); err != nil {
		if !fortiAPIPatch(o["fec-health-check"]) {
			return fmt.Errorf("Error reading fec_health_check: %v", err)
		}
	}

	if err = d.Set("fec_mapping_profile", flattenVpnIpsecPhase1InterfaceFecMappingProfile(o["fec-mapping-profile"], d, "fec_mapping_profile", sv)); err != nil {
		if !fortiAPIPatch(o["fec-mapping-profile"]) {
			return fmt.Errorf("Error reading fec_mapping_profile: %v", err)
		}
	}

	if err = d.Set("network_overlay", flattenVpnIpsecPhase1InterfaceNetworkOverlay(o["network-overlay"], d, "network_overlay", sv)); err != nil {
		if !fortiAPIPatch(o["network-overlay"]) {
			return fmt.Errorf("Error reading network_overlay: %v", err)
		}
	}

	if err = d.Set("network_id", flattenVpnIpsecPhase1InterfaceNetworkId(o["network-id"], d, "network_id", sv)); err != nil {
		if !fortiAPIPatch(o["network-id"]) {
			return fmt.Errorf("Error reading network_id: %v", err)
		}
	}

	if err = d.Set("dev_id_notification", flattenVpnIpsecPhase1InterfaceDevIdNotification(o["dev-id-notification"], d, "dev_id_notification", sv)); err != nil {
		if !fortiAPIPatch(o["dev-id-notification"]) {
			return fmt.Errorf("Error reading dev_id_notification: %v", err)
		}
	}

	if err = d.Set("dev_id", flattenVpnIpsecPhase1InterfaceDevId(o["dev-id"], d, "dev_id", sv)); err != nil {
		if !fortiAPIPatch(o["dev-id"]) {
			return fmt.Errorf("Error reading dev_id: %v", err)
		}
	}

	if err = d.Set("loopback_asymroute", flattenVpnIpsecPhase1InterfaceLoopbackAsymroute(o["loopback-asymroute"], d, "loopback_asymroute", sv)); err != nil {
		if !fortiAPIPatch(o["loopback-asymroute"]) {
			return fmt.Errorf("Error reading loopback_asymroute: %v", err)
		}
	}

	if err = d.Set("link_cost", flattenVpnIpsecPhase1InterfaceLinkCost(o["link-cost"], d, "link_cost", sv)); err != nil {
		if !fortiAPIPatch(o["link-cost"]) {
			return fmt.Errorf("Error reading link_cost: %v", err)
		}
	}

	if err = d.Set("kms", flattenVpnIpsecPhase1InterfaceKms(o["kms"], d, "kms", sv)); err != nil {
		if !fortiAPIPatch(o["kms"]) {
			return fmt.Errorf("Error reading kms: %v", err)
		}
	}

	if err = d.Set("exchange_fgt_device_id", flattenVpnIpsecPhase1InterfaceExchangeFgtDeviceId(o["exchange-fgt-device-id"], d, "exchange_fgt_device_id", sv)); err != nil {
		if !fortiAPIPatch(o["exchange-fgt-device-id"]) {
			return fmt.Errorf("Error reading exchange_fgt_device_id: %v", err)
		}
	}

	if err = d.Set("ipv6_auto_linklocal", flattenVpnIpsecPhase1InterfaceIpv6AutoLinklocal(o["ipv6-auto-linklocal"], d, "ipv6_auto_linklocal", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-auto-linklocal"]) {
			return fmt.Errorf("Error reading ipv6_auto_linklocal: %v", err)
		}
	}

	if err = d.Set("ems_sn_check", flattenVpnIpsecPhase1InterfaceEmsSnCheck(o["ems-sn-check"], d, "ems_sn_check", sv)); err != nil {
		if !fortiAPIPatch(o["ems-sn-check"]) {
			return fmt.Errorf("Error reading ems_sn_check: %v", err)
		}
	}

	if err = d.Set("remote_gw_match", flattenVpnIpsecPhase1InterfaceRemoteGwMatch(o["remote-gw-match"], d, "remote_gw_match", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-match"]) {
			return fmt.Errorf("Error reading remote_gw_match: %v", err)
		}
	}

	if err = d.Set("remote_gw_subnet", flattenVpnIpsecPhase1InterfaceRemoteGwSubnet(o["remote-gw-subnet"], d, "remote_gw_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-subnet"]) {
			return fmt.Errorf("Error reading remote_gw_subnet: %v", err)
		}
	}

	if err = d.Set("remote_gw_start_ip", flattenVpnIpsecPhase1InterfaceRemoteGwStartIp(o["remote-gw-start-ip"], d, "remote_gw_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-start-ip"]) {
			return fmt.Errorf("Error reading remote_gw_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_end_ip", flattenVpnIpsecPhase1InterfaceRemoteGwEndIp(o["remote-gw-end-ip"], d, "remote_gw_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-end-ip"]) {
			return fmt.Errorf("Error reading remote_gw_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw_country", flattenVpnIpsecPhase1InterfaceRemoteGwCountry(o["remote-gw-country"], d, "remote_gw_country", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw-country"]) {
			return fmt.Errorf("Error reading remote_gw_country: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("remote_gw_ztna_tags", flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTags(o["remote-gw-ztna-tags"], d, "remote_gw_ztna_tags", sv)); err != nil {
			if !fortiAPIPatch(o["remote-gw-ztna-tags"]) {
				return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_gw_ztna_tags"); ok {
			if err = d.Set("remote_gw_ztna_tags", flattenVpnIpsecPhase1InterfaceRemoteGwZtnaTags(o["remote-gw-ztna-tags"], d, "remote_gw_ztna_tags", sv)); err != nil {
				if !fortiAPIPatch(o["remote-gw-ztna-tags"]) {
					return fmt.Errorf("Error reading remote_gw_ztna_tags: %v", err)
				}
			}
		}
	}

	if err = d.Set("remote_gw6_match", flattenVpnIpsecPhase1InterfaceRemoteGw6Match(o["remote-gw6-match"], d, "remote_gw6_match", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-match"]) {
			return fmt.Errorf("Error reading remote_gw6_match: %v", err)
		}
	}

	if err = d.Set("remote_gw6_subnet", flattenVpnIpsecPhase1InterfaceRemoteGw6Subnet(o["remote-gw6-subnet"], d, "remote_gw6_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-subnet"]) {
			return fmt.Errorf("Error reading remote_gw6_subnet: %v", err)
		}
	}

	if err = d.Set("remote_gw6_start_ip", flattenVpnIpsecPhase1InterfaceRemoteGw6StartIp(o["remote-gw6-start-ip"], d, "remote_gw6_start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-start-ip"]) {
			return fmt.Errorf("Error reading remote_gw6_start_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_end_ip", flattenVpnIpsecPhase1InterfaceRemoteGw6EndIp(o["remote-gw6-end-ip"], d, "remote_gw6_end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-end-ip"]) {
			return fmt.Errorf("Error reading remote_gw6_end_ip: %v", err)
		}
	}

	if err = d.Set("remote_gw6_country", flattenVpnIpsecPhase1InterfaceRemoteGw6Country(o["remote-gw6-country"], d, "remote_gw6_country", sv)); err != nil {
		if !fortiAPIPatch(o["remote-gw6-country"]) {
			return fmt.Errorf("Error reading remote_gw6_country: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_validation", flattenVpnIpsecPhase1InterfaceCertPeerUsernameValidation(o["cert-peer-username-validation"], d, "cert_peer_username_validation", sv)); err != nil {
		if !fortiAPIPatch(o["cert-peer-username-validation"]) {
			return fmt.Errorf("Error reading cert_peer_username_validation: %v", err)
		}
	}

	if err = d.Set("cert_peer_username_strip", flattenVpnIpsecPhase1InterfaceCertPeerUsernameStrip(o["cert-peer-username-strip"], d, "cert_peer_username_strip", sv)); err != nil {
		if !fortiAPIPatch(o["cert-peer-username-strip"]) {
			return fmt.Errorf("Error reading cert_peer_username_strip: %v", err)
		}
	}

	if err = d.Set("cert_trust_store", flattenVpnIpsecPhase1InterfaceCertTrustStore(o["cert-trust-store"], d, "cert_trust_store", sv)); err != nil {
		if !fortiAPIPatch(o["cert-trust-store"]) {
			return fmt.Errorf("Error reading cert_trust_store: %v", err)
		}
	}

	if err = d.Set("qkd", flattenVpnIpsecPhase1InterfaceQkd(o["qkd"], d, "qkd", sv)); err != nil {
		if !fortiAPIPatch(o["qkd"]) {
			return fmt.Errorf("Error reading qkd: %v", err)
		}
	}

	if err = d.Set("qkd_profile", flattenVpnIpsecPhase1InterfaceQkdProfile(o["qkd-profile"], d, "qkd_profile", sv)); err != nil {
		if !fortiAPIPatch(o["qkd-profile"]) {
			return fmt.Errorf("Error reading qkd_profile: %v", err)
		}
	}

	if err = d.Set("transport", flattenVpnIpsecPhase1InterfaceTransport(o["transport"], d, "transport", sv)); err != nil {
		if !fortiAPIPatch(o["transport"]) {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	if err = d.Set("fortinet_esp", flattenVpnIpsecPhase1InterfaceFortinetEsp(o["fortinet-esp"], d, "fortinet_esp", sv)); err != nil {
		if !fortiAPIPatch(o["fortinet-esp"]) {
			return fmt.Errorf("Error reading fortinet_esp: %v", err)
		}
	}

	if err = d.Set("auto_transport_threshold", flattenVpnIpsecPhase1InterfaceAutoTransportThreshold(o["auto-transport-threshold"], d, "auto_transport_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["auto-transport-threshold"]) {
			return fmt.Errorf("Error reading auto_transport_threshold: %v", err)
		}
	}

	if err = d.Set("fallback_tcp_threshold", flattenVpnIpsecPhase1InterfaceFallbackTcpThreshold(o["fallback-tcp-threshold"], d, "fallback_tcp_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["fallback-tcp-threshold"]) {
			return fmt.Errorf("Error reading fallback_tcp_threshold: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase1InterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnIpsecPhase1InterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIkeVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemotegwDdns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceKeylife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnIpsecPhase1InterfaceCertificateName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceCertificateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthmethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthmethodRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePeertype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePeerid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDefaultGw(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDefaultGwPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceUsrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePeergrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"7.4.1"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); versionMatch {
		vx := fmt.Sprintf("%v", v)
		vxx := strings.Split(vx, ", ")

		tmps := make([]map[string]interface{}, 0, len(vxx))

		for _, xv := range vxx {
			xtmp := make(map[string]interface{})
			xtmp["name"] = xv

			tmps = append(tmps, xtmp)
		}
		return tmps, nil
	} else {
		return v, nil
	}
}

func expandVpnIpsecPhase1InterfaceMonitorMin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMonitorHoldDownTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNetDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceTunnelSearch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePassiveMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeInterfaceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeIpAddr4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeIpAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAggregateMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAggregateWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePacketRedistribution(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceModeCfg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAssignIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAssignIpFrom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4Netmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDnsMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4DnsServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceInternalDomainList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["domain-name"], _ = expandVpnIpsecPhase1InterfaceInternalDomainListDomainName(d, i["domain_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceInternalDomainListDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4WinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4WinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4SplitInclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSplitIncludeService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6DnsServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6SplitInclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpDelayInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceUnitySupport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceBanner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIncludeLocalLan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv4SplitExclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6SplitExclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSavePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientAutoNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientKeepAlive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceBackupGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["address"], _ = expandVpnIpsecPhase1InterfaceBackupGatewayAddress(d, i["address"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceBackupGatewayAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceProposal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddGwRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePsksecretRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLocalidType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNegotiateTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFragmentation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpFragmentation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDpd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDpdRetrycount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDpdRetryinterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceForticlientEnforcement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNpuOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSendCertChain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDhgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAddke7(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSuiteB(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEapIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEapExcludePeergrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEapCertAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAcctVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePpk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePpkSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfacePpkIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceWizardType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceXauthtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthusr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthpasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceGroupAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceGroupAuthenticationSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAuthusrgrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceMeshSelectorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSharedIdleTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIdleTimeoutinterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceHaSyncEspSeqno(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFgspSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceInboundDscpCopy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoverySender(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryPsk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoDiscoveryDialupPlaceholder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapsulation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapsulationAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapLocalGw4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapLocalGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapRemoteGw4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEncapRemoteGw6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceVni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNattraversal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEsn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFragmentationMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceChildlessIke(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAzureAdAutoconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientResume(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceClientResumeInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRekey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDigitalSignatureAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceSignatureHashAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRsaSignatureFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRsaSignatureHashOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEnforceUniqueId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertIdValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecEgress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecSendTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecBase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecCodecString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecCodec(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecRedundant(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecIngress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecReceiveTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFecMappingProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNetworkOverlay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDevIdNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceDevId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLoopbackAsymroute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceLinkCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceKms(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceExchangeFgtDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceIpv6AutoLinklocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceEmsSnCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwZtnaTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnIpsecPhase1InterfaceRemoteGwZtnaTagsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGwZtnaTagsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6Match(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6Subnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceRemoteGw6Country(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertPeerUsernameValidation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertPeerUsernameStrip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceCertTrustStore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceQkd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceQkdProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceTransport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFortinetEsp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceAutoTransportThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1InterfaceFallbackTcpThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase1Interface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnIpsecPhase1InterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandVpnIpsecPhase1InterfaceType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandVpnIpsecPhase1InterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpVersion(d, v, "ip_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("ike_version"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIkeVersion(d, v, "ike_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-version"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandVpnIpsecPhase1InterfaceLocalGw(d, v, "local_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw6"); ok {
		t, err := expandVpnIpsecPhase1InterfaceLocalGw6(d, v, "local_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw(d, v, "remote_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6(d, v, "remote_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("remotegw_ddns"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemotegwDdns(d, v, "remotegw_ddns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotegw-ddns"] = t
		}
	} else if d.HasChange("remotegw_ddns") {
		obj["remotegw-ddns"] = nil
	}

	if v, ok := d.GetOk("keylife"); ok {
		t, err := expandVpnIpsecPhase1InterfaceKeylife(d, v, "keylife", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok || d.HasChange("certificate") {
		t, err := expandVpnIpsecPhase1InterfaceCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("authmethod"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAuthmethod(d, v, "authmethod", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod"] = t
		}
	}

	if v, ok := d.GetOk("authmethod_remote"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAuthmethodRemote(d, v, "authmethod_remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod-remote"] = t
		}
	} else if d.HasChange("authmethod_remote") {
		obj["authmethod-remote"] = nil
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("peertype"); ok {
		t, err := expandVpnIpsecPhase1InterfacePeertype(d, v, "peertype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peertype"] = t
		}
	}

	if v, ok := d.GetOk("peerid"); ok {
		t, err := expandVpnIpsecPhase1InterfacePeerid(d, v, "peerid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerid"] = t
		}
	} else if d.HasChange("peerid") {
		obj["peerid"] = nil
	}

	if v, ok := d.GetOk("default_gw"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDefaultGw(d, v, "default_gw", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gw"] = t
		}
	}

	if v, ok := d.GetOkExists("default_gw_priority"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDefaultGwPriority(d, v, "default_gw_priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-gw-priority"] = t
		}
	} else if d.HasChange("default_gw_priority") {
		obj["default-gw-priority"] = nil
	}

	if v, ok := d.GetOk("usrgrp"); ok {
		t, err := expandVpnIpsecPhase1InterfaceUsrgrp(d, v, "usrgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	} else if d.HasChange("usrgrp") {
		obj["usrgrp"] = nil
	}

	if v, ok := d.GetOk("peer"); ok {
		t, err := expandVpnIpsecPhase1InterfacePeer(d, v, "peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	} else if d.HasChange("peer") {
		obj["peer"] = nil
	}

	if v, ok := d.GetOk("peergrp"); ok {
		t, err := expandVpnIpsecPhase1InterfacePeergrp(d, v, "peergrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peergrp"] = t
		}
	} else if d.HasChange("peergrp") {
		obj["peergrp"] = nil
	}

	if v, ok := d.GetOk("monitor"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMonitor(d, v, "monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	} else if d.HasChange("monitor") {
		obj["monitor"] = nil
	}

	if v, ok := d.GetOkExists("monitor_min"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMonitorMin(d, v, "monitor_min", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-min"] = t
		}
	} else if d.HasChange("monitor_min") {
		obj["monitor-min"] = nil
	}

	if v, ok := d.GetOk("monitor_hold_down_type"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownType(d, v, "monitor_hold_down_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-type"] = t
		}
	}

	if v, ok := d.GetOkExists("monitor_hold_down_delay"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownDelay(d, v, "monitor_hold_down_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-delay"] = t
		}
	} else if d.HasChange("monitor_hold_down_delay") {
		obj["monitor-hold-down-delay"] = nil
	}

	if v, ok := d.GetOk("monitor_hold_down_weekday"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownWeekday(d, v, "monitor_hold_down_weekday", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-weekday"] = t
		}
	}

	if v, ok := d.GetOk("monitor_hold_down_time"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMonitorHoldDownTime(d, v, "monitor_hold_down_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOk("net_device"); ok {
		t, err := expandVpnIpsecPhase1InterfaceNetDevice(d, v, "net_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["net-device"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_search"); ok {
		t, err := expandVpnIpsecPhase1InterfaceTunnelSearch(d, v, "tunnel_search", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-search"] = t
		}
	}

	if v, ok := d.GetOk("passive_mode"); ok {
		t, err := expandVpnIpsecPhase1InterfacePassiveMode(d, v, "passive_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-mode"] = t
		}
	}

	if v, ok := d.GetOk("exchange_interface_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceExchangeInterfaceIp(d, v, "exchange_interface_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-interface-ip"] = t
		}
	}

	if v, ok := d.GetOk("exchange_ip_addr4"); ok {
		t, err := expandVpnIpsecPhase1InterfaceExchangeIpAddr4(d, v, "exchange_ip_addr4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-ip-addr4"] = t
		}
	}

	if v, ok := d.GetOk("exchange_ip_addr6"); ok {
		t, err := expandVpnIpsecPhase1InterfaceExchangeIpAddr6(d, v, "exchange_ip_addr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-ip-addr6"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_member"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAggregateMember(d, v, "aggregate_member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-member"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_weight"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAggregateWeight(d, v, "aggregate_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-weight"] = t
		}
	}

	if v, ok := d.GetOk("packet_redistribution"); ok {
		t, err := expandVpnIpsecPhase1InterfacePacketRedistribution(d, v, "packet_redistribution", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-redistribution"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg"); ok {
		t, err := expandVpnIpsecPhase1InterfaceModeCfg(d, v, "mode_cfg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg_allow_client_selector"); ok {
		t, err := expandVpnIpsecPhase1InterfaceModeCfgAllowClientSelector(d, v, "mode_cfg_allow_client_selector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg-allow-client-selector"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAssignIp(d, v, "assign_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip_from"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAssignIpFrom(d, v, "assign_ip_from", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip-from"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4StartIp(d, v, "ipv4_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4EndIp(d, v, "ipv4_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_netmask"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4Netmask(d, v, "ipv4_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDhcpRaGiaddr(d, v, "dhcp_ra_giaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDnsMode(d, v, "dns_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server1"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4DnsServer1(d, v, "ipv4_dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server2"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4DnsServer2(d, v, "ipv4_dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server3"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4DnsServer3(d, v, "ipv4_dns_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("internal_domain_list"); ok || d.HasChange("internal_domain_list") {
		t, err := expandVpnIpsecPhase1InterfaceInternalDomainList(d, v, "internal_domain_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-domain-list"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server1"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4WinsServer1(d, v, "ipv4_wins_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server2"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4WinsServer2(d, v, "ipv4_wins_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_exclude_range"); ok || d.HasChange("ipv4_exclude_range") {
		t, err := expandVpnIpsecPhase1InterfaceIpv4ExcludeRange(d, v, "ipv4_exclude_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_include"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4SplitInclude(d, v, "ipv4_split_include", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-include"] = t
		}
	} else if d.HasChange("ipv4_split_include") {
		obj["ipv4-split-include"] = nil
	}

	if v, ok := d.GetOk("split_include_service"); ok {
		t, err := expandVpnIpsecPhase1InterfaceSplitIncludeService(d, v, "split_include_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-include-service"] = t
		}
	} else if d.HasChange("split_include_service") {
		obj["split-include-service"] = nil
	}

	if v, ok := d.GetOk("ipv4_name"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4Name(d, v, "ipv4_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-name"] = t
		}
	} else if d.HasChange("ipv4_name") {
		obj["ipv4-name"] = nil
	}

	if v, ok := d.GetOk("ipv6_start_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6StartIp(d, v, "ipv6_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_end_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6EndIp(d, v, "ipv6_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6Prefix(d, v, "ipv6_prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6DnsServer1(d, v, "ipv6_dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6DnsServer2(d, v, "ipv6_dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server3"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6DnsServer3(d, v, "ipv6_dns_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclude_range"); ok || d.HasChange("ipv6_exclude_range") {
		t, err := expandVpnIpsecPhase1InterfaceIpv6ExcludeRange(d, v, "ipv6_exclude_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_include"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6SplitInclude(d, v, "ipv6_split_include", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-include"] = t
		}
	} else if d.HasChange("ipv6_split_include") {
		obj["ipv6-split-include"] = nil
	}

	if v, ok := d.GetOk("ipv6_name"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6Name(d, v, "ipv6_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-name"] = t
		}
	} else if d.HasChange("ipv6_name") {
		obj["ipv6-name"] = nil
	}

	if v, ok := d.GetOkExists("ip_delay_interval"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpDelayInterval(d, v, "ip_delay_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-delay-interval"] = t
		}
	} else if d.HasChange("ip_delay_interval") {
		obj["ip-delay-interval"] = nil
	}

	if v, ok := d.GetOk("unity_support"); ok {
		t, err := expandVpnIpsecPhase1InterfaceUnitySupport(d, v, "unity_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unity-support"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if d.HasChange("domain") {
		obj["domain"] = nil
	}

	if v, ok := d.GetOk("banner"); ok {
		t, err := expandVpnIpsecPhase1InterfaceBanner(d, v, "banner", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner"] = t
		}
	} else if d.HasChange("banner") {
		obj["banner"] = nil
	}

	if v, ok := d.GetOk("include_local_lan"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIncludeLocalLan(d, v, "include_local_lan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-local-lan"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_exclude"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv4SplitExclude(d, v, "ipv4_split_exclude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-exclude"] = t
		}
	} else if d.HasChange("ipv4_split_exclude") {
		obj["ipv4-split-exclude"] = nil
	}

	if v, ok := d.GetOk("ipv6_split_exclude"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6SplitExclude(d, v, "ipv6_split_exclude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-exclude"] = t
		}
	} else if d.HasChange("ipv6_split_exclude") {
		obj["ipv6-split-exclude"] = nil
	}

	if v, ok := d.GetOk("save_password"); ok {
		t, err := expandVpnIpsecPhase1InterfaceSavePassword(d, v, "save_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("client_auto_negotiate"); ok {
		t, err := expandVpnIpsecPhase1InterfaceClientAutoNegotiate(d, v, "client_auto_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("client_keep_alive"); ok {
		t, err := expandVpnIpsecPhase1InterfaceClientKeepAlive(d, v, "client_keep_alive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("backup_gateway"); ok || d.HasChange("backup_gateway") {
		t, err := expandVpnIpsecPhase1InterfaceBackupGateway(d, v, "backup_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-gateway"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok {
		t, err := expandVpnIpsecPhase1InterfaceProposal(d, v, "proposal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	} else if d.HasChange("proposal") {
		obj["proposal"] = nil
	}

	if v, ok := d.GetOk("add_route"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddRoute(d, v, "add_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("add_gw_route"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddGwRoute(d, v, "add_gw_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-gw-route"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		t, err := expandVpnIpsecPhase1InterfacePsksecret(d, v, "psksecret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	} else if d.HasChange("psksecret") {
		obj["psksecret"] = nil
	}

	if v, ok := d.GetOk("psksecret_remote"); ok {
		t, err := expandVpnIpsecPhase1InterfacePsksecretRemote(d, v, "psksecret_remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret-remote"] = t
		}
	} else if d.HasChange("psksecret_remote") {
		obj["psksecret-remote"] = nil
	}

	if v, ok := d.GetOk("keepalive"); ok {
		t, err := expandVpnIpsecPhase1InterfaceKeepalive(d, v, "keepalive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDistance(d, v, "distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {
		t, err := expandVpnIpsecPhase1InterfacePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("localid"); ok {
		t, err := expandVpnIpsecPhase1InterfaceLocalid(d, v, "localid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid"] = t
		}
	} else if d.HasChange("localid") {
		obj["localid"] = nil
	}

	if v, ok := d.GetOk("localid_type"); ok {
		t, err := expandVpnIpsecPhase1InterfaceLocalidType(d, v, "localid_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoNegotiate(d, v, "auto_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_timeout"); ok {
		t, err := expandVpnIpsecPhase1InterfaceNegotiateTimeout(d, v, "negotiate_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFragmentation(d, v, "fragmentation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragmentation"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpFragmentation(d, v, "ip_fragmentation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("dpd"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDpd(d, v, "dpd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd"] = t
		}
	}

	if v, ok := d.GetOkExists("dpd_retrycount"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDpdRetrycount(d, v, "dpd_retrycount", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retrycount"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retryinterval"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDpdRetryinterval(d, v, "dpd_retryinterval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retryinterval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_enforcement"); ok {
		t, err := expandVpnIpsecPhase1InterfaceForticlientEnforcement(d, v, "forticlient_enforcement", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandVpnIpsecPhase1InterfaceComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("npu_offload"); ok {
		t, err := expandVpnIpsecPhase1InterfaceNpuOffload(d, v, "npu_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["npu-offload"] = t
		}
	}

	if v, ok := d.GetOk("send_cert_chain"); ok {
		t, err := expandVpnIpsecPhase1InterfaceSendCertChain(d, v, "send_cert_chain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-cert-chain"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDhgrp(d, v, "dhgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("addke1"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke1(d, v, "addke1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke1"] = t
		}
	} else if d.HasChange("addke1") {
		obj["addke1"] = nil
	}

	if v, ok := d.GetOk("addke2"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke2(d, v, "addke2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke2"] = t
		}
	} else if d.HasChange("addke2") {
		obj["addke2"] = nil
	}

	if v, ok := d.GetOk("addke3"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke3(d, v, "addke3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke3"] = t
		}
	} else if d.HasChange("addke3") {
		obj["addke3"] = nil
	}

	if v, ok := d.GetOk("addke4"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke4(d, v, "addke4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke4"] = t
		}
	} else if d.HasChange("addke4") {
		obj["addke4"] = nil
	}

	if v, ok := d.GetOk("addke5"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke5(d, v, "addke5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke5"] = t
		}
	} else if d.HasChange("addke5") {
		obj["addke5"] = nil
	}

	if v, ok := d.GetOk("addke6"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke6(d, v, "addke6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke6"] = t
		}
	} else if d.HasChange("addke6") {
		obj["addke6"] = nil
	}

	if v, ok := d.GetOk("addke7"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAddke7(d, v, "addke7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addke7"] = t
		}
	} else if d.HasChange("addke7") {
		obj["addke7"] = nil
	}

	if v, ok := d.GetOk("suite_b"); ok {
		t, err := expandVpnIpsecPhase1InterfaceSuiteB(d, v, "suite_b", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["suite-b"] = t
		}
	}

	if v, ok := d.GetOk("eap"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEap(d, v, "eap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap"] = t
		}
	}

	if v, ok := d.GetOk("eap_identity"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEapIdentity(d, v, "eap_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-identity"] = t
		}
	}

	if v, ok := d.GetOk("eap_exclude_peergrp"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEapExcludePeergrp(d, v, "eap_exclude_peergrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-exclude-peergrp"] = t
		}
	} else if d.HasChange("eap_exclude_peergrp") {
		obj["eap-exclude-peergrp"] = nil
	}

	if v, ok := d.GetOk("eap_cert_auth"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEapCertAuth(d, v, "eap_cert_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-cert-auth"] = t
		}
	}

	if v, ok := d.GetOk("acct_verify"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAcctVerify(d, v, "acct_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-verify"] = t
		}
	}

	if v, ok := d.GetOk("ppk"); ok {
		t, err := expandVpnIpsecPhase1InterfacePpk(d, v, "ppk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok {
		t, err := expandVpnIpsecPhase1InterfacePpkSecret(d, v, "ppk_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	} else if d.HasChange("ppk_secret") {
		obj["ppk-secret"] = nil
	}

	if v, ok := d.GetOk("ppk_identity"); ok {
		t, err := expandVpnIpsecPhase1InterfacePpkIdentity(d, v, "ppk_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	} else if d.HasChange("ppk_identity") {
		obj["ppk-identity"] = nil
	}

	if v, ok := d.GetOk("wizard_type"); ok {
		t, err := expandVpnIpsecPhase1InterfaceWizardType(d, v, "wizard_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wizard-type"] = t
		}
	}

	if v, ok := d.GetOk("xauthtype"); ok {
		t, err := expandVpnIpsecPhase1InterfaceXauthtype(d, v, "xauthtype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xauthtype"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok {
		t, err := expandVpnIpsecPhase1InterfaceReauth(d, v, "reauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("authusr"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAuthusr(d, v, "authusr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusr"] = t
		}
	} else if d.HasChange("authusr") {
		obj["authusr"] = nil
	}

	if v, ok := d.GetOk("authpasswd"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAuthpasswd(d, v, "authpasswd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authpasswd"] = t
		}
	} else if d.HasChange("authpasswd") {
		obj["authpasswd"] = nil
	}

	if v, ok := d.GetOk("group_authentication"); ok {
		t, err := expandVpnIpsecPhase1InterfaceGroupAuthentication(d, v, "group_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication_secret"); ok {
		t, err := expandVpnIpsecPhase1InterfaceGroupAuthenticationSecret(d, v, "group_authentication_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication-secret"] = t
		}
	} else if d.HasChange("group_authentication_secret") {
		obj["group-authentication-secret"] = nil
	}

	if v, ok := d.GetOk("authusrgrp"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAuthusrgrp(d, v, "authusrgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusrgrp"] = t
		}
	} else if d.HasChange("authusrgrp") {
		obj["authusrgrp"] = nil
	}

	if v, ok := d.GetOk("mesh_selector_type"); ok {
		t, err := expandVpnIpsecPhase1InterfaceMeshSelectorType(d, v, "mesh_selector_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-selector-type"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIdleTimeout(d, v, "idle_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("shared_idle_timeout"); ok {
		t, err := expandVpnIpsecPhase1InterfaceSharedIdleTimeout(d, v, "shared_idle_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeoutinterval"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIdleTimeoutinterval(d, v, "idle_timeoutinterval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeoutinterval"] = t
		}
	}

	if v, ok := d.GetOk("ha_sync_esp_seqno"); ok {
		t, err := expandVpnIpsecPhase1InterfaceHaSyncEspSeqno(d, v, "ha_sync_esp_seqno", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-sync-esp-seqno"] = t
		}
	}

	if v, ok := d.GetOk("fgsp_sync"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFgspSync(d, v, "fgsp_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgsp-sync"] = t
		}
	}

	if v, ok := d.GetOk("inbound_dscp_copy"); ok {
		t, err := expandVpnIpsecPhase1InterfaceInboundDscpCopy(d, v, "inbound_dscp_copy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbound-dscp-copy"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_sender"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoverySender(d, v, "auto_discovery_sender", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-sender"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_receiver"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryReceiver(d, v, "auto_discovery_receiver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-receiver"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_forwarder"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryForwarder(d, v, "auto_discovery_forwarder", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-forwarder"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_psk"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryPsk(d, v, "auto_discovery_psk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-psk"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_shortcuts"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryShortcuts(d, v, "auto_discovery_shortcuts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-shortcuts"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_crossover"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryCrossover(d, v, "auto_discovery_crossover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-crossover"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_offer_interval"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryOfferInterval(d, v, "auto_discovery_offer_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-offer-interval"] = t
		}
	}

	if v, ok := d.GetOk("auto_discovery_dialup_placeholder"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoDiscoveryDialupPlaceholder(d, v, "auto_discovery_dialup_placeholder", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-discovery-dialup-placeholder"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEncapsulation(d, v, "encapsulation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation"] = t
		}
	}

	if v, ok := d.GetOk("encapsulation_address"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEncapsulationAddress(d, v, "encapsulation_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encapsulation-address"] = t
		}
	}

	if v, ok := d.GetOk("encap_local_gw4"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEncapLocalGw4(d, v, "encap_local_gw4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-local-gw4"] = t
		}
	}

	if v, ok := d.GetOk("encap_local_gw6"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEncapLocalGw6(d, v, "encap_local_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-local-gw6"] = t
		}
	}

	if v, ok := d.GetOk("encap_remote_gw4"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEncapRemoteGw4(d, v, "encap_remote_gw4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-remote-gw4"] = t
		}
	}

	if v, ok := d.GetOk("encap_remote_gw6"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEncapRemoteGw6(d, v, "encap_remote_gw6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encap-remote-gw6"] = t
		}
	}

	if v, ok := d.GetOk("vni"); ok {
		t, err := expandVpnIpsecPhase1InterfaceVni(d, v, "vni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	} else if d.HasChange("vni") {
		obj["vni"] = nil
	}

	if v, ok := d.GetOk("nattraversal"); ok {
		t, err := expandVpnIpsecPhase1InterfaceNattraversal(d, v, "nattraversal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nattraversal"] = t
		}
	}

	if v, ok := d.GetOk("esn"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEsn(d, v, "esn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esn"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation_mtu"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFragmentationMtu(d, v, "fragmentation_mtu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation-mtu"] = t
		}
	}

	if v, ok := d.GetOk("childless_ike"); ok {
		t, err := expandVpnIpsecPhase1InterfaceChildlessIke(d, v, "childless_ike", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["childless-ike"] = t
		}
	}

	if v, ok := d.GetOk("azure_ad_autoconnect"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAzureAdAutoconnect(d, v, "azure_ad_autoconnect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-ad-autoconnect"] = t
		}
	}

	if v, ok := d.GetOk("client_resume"); ok {
		t, err := expandVpnIpsecPhase1InterfaceClientResume(d, v, "client_resume", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume"] = t
		}
	}

	if v, ok := d.GetOk("client_resume_interval"); ok {
		t, err := expandVpnIpsecPhase1InterfaceClientResumeInterval(d, v, "client_resume_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-resume-interval"] = t
		}
	}

	if v, ok := d.GetOk("rekey"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRekey(d, v, "rekey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rekey"] = t
		}
	}

	if v, ok := d.GetOk("digital_signature_auth"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDigitalSignatureAuth(d, v, "digital_signature_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digital-signature-auth"] = t
		}
	}

	if v, ok := d.GetOk("signature_hash_alg"); ok {
		t, err := expandVpnIpsecPhase1InterfaceSignatureHashAlg(d, v, "signature_hash_alg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-hash-alg"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_format"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRsaSignatureFormat(d, v, "rsa_signature_format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-format"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_hash_override"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRsaSignatureHashOverride(d, v, "rsa_signature_hash_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-hash-override"] = t
		}
	}

	if v, ok := d.GetOk("enforce_unique_id"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEnforceUniqueId(d, v, "enforce_unique_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-unique-id"] = t
		}
	}

	if v, ok := d.GetOk("cert_id_validation"); ok {
		t, err := expandVpnIpsecPhase1InterfaceCertIdValidation(d, v, "cert_id_validation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-id-validation"] = t
		}
	}

	if v, ok := d.GetOk("fec_egress"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecEgress(d, v, "fec_egress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-egress"] = t
		}
	}

	if v, ok := d.GetOk("fec_send_timeout"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecSendTimeout(d, v, "fec_send_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-send-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fec_base"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecBase(d, v, "fec_base", sv)
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
			t, err := expandVpnIpsecPhase1InterfaceFecCodecString(d, v, "fec_codec_string", sv)
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
			t, err := expandVpnIpsecPhase1InterfaceFecCodec(d, v, "fec_codec", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fec-codec"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("fec_redundant"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecRedundant(d, v, "fec_redundant", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-redundant"] = t
		}
	}

	if v, ok := d.GetOk("fec_ingress"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecIngress(d, v, "fec_ingress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-ingress"] = t
		}
	}

	if v, ok := d.GetOk("fec_receive_timeout"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecReceiveTimeout(d, v, "fec_receive_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fec_health_check"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecHealthCheck(d, v, "fec_health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-health-check"] = t
		}
	} else if d.HasChange("fec_health_check") {
		obj["fec-health-check"] = nil
	}

	if v, ok := d.GetOk("fec_mapping_profile"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFecMappingProfile(d, v, "fec_mapping_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fec-mapping-profile"] = t
		}
	} else if d.HasChange("fec_mapping_profile") {
		obj["fec-mapping-profile"] = nil
	}

	if v, ok := d.GetOk("network_overlay"); ok {
		t, err := expandVpnIpsecPhase1InterfaceNetworkOverlay(d, v, "network_overlay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-overlay"] = t
		}
	}

	if v, ok := d.GetOkExists("network_id"); ok {
		t, err := expandVpnIpsecPhase1InterfaceNetworkId(d, v, "network_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-id"] = t
		}
	} else if d.HasChange("network_id") {
		obj["network-id"] = nil
	}

	if v, ok := d.GetOk("dev_id_notification"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDevIdNotification(d, v, "dev_id_notification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id-notification"] = t
		}
	}

	if v, ok := d.GetOk("dev_id"); ok {
		t, err := expandVpnIpsecPhase1InterfaceDevId(d, v, "dev_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev-id"] = t
		}
	} else if d.HasChange("dev_id") {
		obj["dev-id"] = nil
	}

	if v, ok := d.GetOk("loopback_asymroute"); ok {
		t, err := expandVpnIpsecPhase1InterfaceLoopbackAsymroute(d, v, "loopback_asymroute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loopback-asymroute"] = t
		}
	}

	if v, ok := d.GetOkExists("link_cost"); ok {
		t, err := expandVpnIpsecPhase1InterfaceLinkCost(d, v, "link_cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost"] = t
		}
	} else if d.HasChange("link_cost") {
		obj["link-cost"] = nil
	}

	if v, ok := d.GetOk("kms"); ok {
		t, err := expandVpnIpsecPhase1InterfaceKms(d, v, "kms", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kms"] = t
		}
	} else if d.HasChange("kms") {
		obj["kms"] = nil
	}

	if v, ok := d.GetOk("exchange_fgt_device_id"); ok {
		t, err := expandVpnIpsecPhase1InterfaceExchangeFgtDeviceId(d, v, "exchange_fgt_device_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-fgt-device-id"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_auto_linklocal"); ok {
		t, err := expandVpnIpsecPhase1InterfaceIpv6AutoLinklocal(d, v, "ipv6_auto_linklocal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-auto-linklocal"] = t
		}
	}

	if v, ok := d.GetOk("ems_sn_check"); ok {
		t, err := expandVpnIpsecPhase1InterfaceEmsSnCheck(d, v, "ems_sn_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-sn-check"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_match"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwMatch(d, v, "remote_gw_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_subnet"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwSubnet(d, v, "remote_gw_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_start_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwStartIp(d, v, "remote_gw_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_end_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwEndIp(d, v, "remote_gw_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw_country"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwCountry(d, v, "remote_gw_country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-country"] = t
		}
	} else if d.HasChange("remote_gw_country") {
		obj["remote-gw-country"] = nil
	}

	if v, ok := d.GetOk("remote_gw_ztna_tags"); ok || d.HasChange("remote_gw_ztna_tags") {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGwZtnaTags(d, v, "remote_gw_ztna_tags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw-ztna-tags"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_match"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6Match(d, v, "remote_gw6_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-match"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_subnet"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6Subnet(d, v, "remote_gw6_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-subnet"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_start_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6StartIp(d, v, "remote_gw6_start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_end_ip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6EndIp(d, v, "remote_gw6_end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw6_country"); ok {
		t, err := expandVpnIpsecPhase1InterfaceRemoteGw6Country(d, v, "remote_gw6_country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw6-country"] = t
		}
	} else if d.HasChange("remote_gw6_country") {
		obj["remote-gw6-country"] = nil
	}

	if v, ok := d.GetOk("cert_peer_username_validation"); ok {
		t, err := expandVpnIpsecPhase1InterfaceCertPeerUsernameValidation(d, v, "cert_peer_username_validation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-validation"] = t
		}
	}

	if v, ok := d.GetOk("cert_peer_username_strip"); ok {
		t, err := expandVpnIpsecPhase1InterfaceCertPeerUsernameStrip(d, v, "cert_peer_username_strip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-peer-username-strip"] = t
		}
	}

	if v, ok := d.GetOk("cert_trust_store"); ok {
		t, err := expandVpnIpsecPhase1InterfaceCertTrustStore(d, v, "cert_trust_store", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-trust-store"] = t
		}
	}

	if v, ok := d.GetOk("qkd"); ok {
		t, err := expandVpnIpsecPhase1InterfaceQkd(d, v, "qkd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd"] = t
		}
	}

	if v, ok := d.GetOk("qkd_profile"); ok {
		t, err := expandVpnIpsecPhase1InterfaceQkdProfile(d, v, "qkd_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qkd-profile"] = t
		}
	} else if d.HasChange("qkd_profile") {
		obj["qkd-profile"] = nil
	}

	if v, ok := d.GetOk("transport"); ok {
		t, err := expandVpnIpsecPhase1InterfaceTransport(d, v, "transport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transport"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_esp"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFortinetEsp(d, v, "fortinet_esp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-esp"] = t
		}
	}

	if v, ok := d.GetOk("auto_transport_threshold"); ok {
		t, err := expandVpnIpsecPhase1InterfaceAutoTransportThreshold(d, v, "auto_transport_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-transport-threshold"] = t
		}
	}

	if v, ok := d.GetOk("fallback_tcp_threshold"); ok {
		t, err := expandVpnIpsecPhase1InterfaceFallbackTcpThreshold(d, v, "fallback_tcp_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fallback-tcp-threshold"] = t
		}
	}

	return &obj, nil
}
