// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: BGP neighbor table.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterbgpNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterbgpNeighborCreate,
		Read:   resourceRouterbgpNeighborRead,
		Update: resourceRouterbgpNeighborUpdate,
		Delete: resourceRouterbgpNeighborDelete,

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
			"ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 45),
				ForceNew:     true,
				Required:     true,
			},
			"advertisement_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 600),
				Optional:     true,
				Computed:     true,
			},
			"allowas_in_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowas_in_enable6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowas_in_enable_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowas_in_enable_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowas_in_enable_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowas_in": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"allowas_in6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"allowas_in_vpnv4": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"allowas_in_vpnv6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"allowas_in_evpn": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"attribute_unchanged": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_unchanged6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_unchanged_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_unchanged_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"activate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"activate6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"activate_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"activate_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"activate_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_dynamic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_orf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_orf6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_graceful_restart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_graceful_restart6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_graceful_restart_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_graceful_restart_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_graceful_restart_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_route_refresh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_default_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capability_default_originate6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dont_capability_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ebgp_enforce_multihop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_down_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stale_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_self": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_self6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_self_rr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_self_rr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_self_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"next_hop_self_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_capability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_private_as": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_private_as6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_private_as_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_private_as_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_private_as_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_reflector_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_reflector_client6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_reflector_client_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_reflector_client_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_reflector_client_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_server_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_server_client6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_server_client_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_server_client_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_server_client_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rr_attr_allow_change_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shutdown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"soft_reconfiguration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"soft_reconfiguration6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"soft_reconfiguration_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"soft_reconfiguration_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"soft_reconfiguration_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"as_override6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strict_capability_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_originate_routemap": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"default_originate_routemap6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"distribute_list_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_in6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_in_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_in_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_out": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_out6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_out_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distribute_list_out_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ebgp_multihop_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"filter_list_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_in6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_in_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_in_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_out": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_out6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_out_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"filter_list_out_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"maximum_prefix": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_prefix6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_prefix_vpnv4": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_prefix_vpnv6": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_prefix_evpn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maximum_prefix_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_threshold6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_threshold_vpnv4": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_threshold_vpnv6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_threshold_evpn": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"maximum_prefix_warning_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_prefix_warning_only6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_prefix_warning_only_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_prefix_warning_only_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_prefix_warning_only_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefix_list_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_in6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_in_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_in_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_out": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_out6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_out_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_list_out_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"remote_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"local_as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"local_as_no_prepend": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_as_replace_as": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retain_stale_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"route_map_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_in6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_in_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_in_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_in_evpn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out_preferable": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out6_preferable": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out_vpnv4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out_vpnv6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out_vpnv4_preferable": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out_vpnv6_preferable": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"route_map_out_evpn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"send_community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_community6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_community_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_community_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_community_evpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keep_alive_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithMinus1(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"holdtime_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithMinus1(3, 65535),
				Optional:     true,
				Computed:     true,
			},
			"connect_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithMinus1(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"unsuppress_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"unsuppress_map6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"update_source": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithMinus1(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"restart_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
			},
			"additional_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path_vpnv4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path_vpnv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_additional_path": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"adv_additional_path6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"adv_additional_path_vpnv4": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"adv_additional_path_vpnv6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"auth_options": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"conditional_advertise": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise_routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"condition_routemap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"conditional_advertise6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise_routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"condition_routemap": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"condition_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceRouterbgpNeighborCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterbgpNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterbgpNeighbor resource while getting object: %v", err)
	}

	o, err := c.CreateRouterbgpNeighbor(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterbgpNeighbor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterbgpNeighbor")
	}

	return resourceRouterbgpNeighborRead(d, m)
}

func resourceRouterbgpNeighborUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterbgpNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterbgpNeighbor resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterbgpNeighbor(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterbgpNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterbgpNeighbor")
	}

	return resourceRouterbgpNeighborRead(d, m)
}

func resourceRouterbgpNeighborDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterbgpNeighbor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterbgpNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterbgpNeighborRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterbgpNeighbor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterbgpNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterbgpNeighbor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterbgpNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenRouterbgpNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAllowasInEnableVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAllowasInEnableVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAllowasInEnableEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAllowasInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAllowasInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAllowasInEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAttributeUnchangedVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAttributeUnchangedVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborActivateVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborActivateVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborActivateEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityGracefulRestartVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityGracefulRestartVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityGracefulRestartEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborLinkDownFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborStaleRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborNextHopSelfVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborNextHopSelfVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRemovePrivateAsVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRemovePrivateAsVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRemovePrivateAsEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteReflectorClientVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteReflectorClientVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteReflectorClientEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteServerClientVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteServerClientVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteServerClientEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRrAttrAllowChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRrAttrAllowChange6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRrAttrAllowChangeVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRrAttrAllowChangeVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRrAttrAllowChangeEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSoftReconfigurationVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSoftReconfigurationVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSoftReconfigurationEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborDistributeListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborFilterListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixThresholdVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixThresholdVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixThresholdEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborMaximumPrefixWarningOnlyVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborMaximumPrefixWarningOnlyVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborMaximumPrefixWarningOnlyEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborPrefixListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborLocalAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRetainStaleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapInEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOutVpnv4Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOutVpnv6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborRouteMapOutEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSendCommunityVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSendCommunityVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborSendCommunityEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if vf, ok := v.(float64); ok && (vf < 0 || vf == 4294967295) {
		var vi interface{}
		vi = -1
		return vi
	}
	return v
}

func flattenRouterbgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if vf, ok := v.(float64); ok && (vf < 0 || vf == 4294967295) {
		var vi interface{}
		vi = -1
		return vi
	}
	return v
}

func flattenRouterbgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if vf, ok := v.(float64); ok && (vf < 0 || vf == 4294967295) {
		var vi interface{}
		vi = -1
		return vi
	}
	return v
}

func flattenRouterbgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if vf, ok := v.(float64); ok && (vf < 0 || vf == 4294967295) {
		var vi interface{}
		vi = -1
		return vi
	}
	return v
}

func flattenRouterbgpNeighborRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAdvAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAdvAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterbgpNeighborAuthOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborConditionalAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if cur_v, ok := i["advertise-routemap"]; ok {
			tmp["advertise_routemap"] = flattenRouterbgpNeighborConditionalAdvertiseAdvertiseRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if cur_v, ok := i["condition-routemap"]; ok {
			tmp["condition_routemap"] = flattenRouterbgpNeighborConditionalAdvertiseConditionRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if cur_v, ok := i["condition-type"]; ok {
			tmp["condition_type"] = flattenRouterbgpNeighborConditionalAdvertiseConditionType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "advertise_routemap", d)
	return result
}

func flattenRouterbgpNeighborConditionalAdvertiseAdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborConditionalAdvertiseConditionRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("condition_routemap"), "name")
}

func flattenRouterbgpNeighborConditionalAdvertiseConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborConditionalAdvertise6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if cur_v, ok := i["advertise-routemap"]; ok {
			tmp["advertise_routemap"] = flattenRouterbgpNeighborConditionalAdvertise6AdvertiseRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if cur_v, ok := i["condition-routemap"]; ok {
			tmp["condition_routemap"] = flattenRouterbgpNeighborConditionalAdvertise6ConditionRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if cur_v, ok := i["condition-type"]; ok {
			tmp["condition_type"] = flattenRouterbgpNeighborConditionalAdvertise6ConditionType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "advertise_routemap", d)
	return result
}

func flattenRouterbgpNeighborConditionalAdvertise6AdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNeighborConditionalAdvertise6ConditionRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("condition_routemap"), "name")
}

func flattenRouterbgpNeighborConditionalAdvertise6ConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterbgpNeighbor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("ip", flattenRouterbgpNeighborIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("advertisement_interval", flattenRouterbgpNeighborAdvertisementInterval(o["advertisement-interval"], d, "advertisement_interval", sv)); err != nil {
		if !fortiAPIPatch(o["advertisement-interval"]) {
			return fmt.Errorf("Error reading advertisement_interval: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable", flattenRouterbgpNeighborAllowasInEnable(o["allowas-in-enable"], d, "allowas_in_enable", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable"]) {
			return fmt.Errorf("Error reading allowas_in_enable: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable6", flattenRouterbgpNeighborAllowasInEnable6(o["allowas-in-enable6"], d, "allowas_in_enable6", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable6"]) {
			return fmt.Errorf("Error reading allowas_in_enable6: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_vpnv4", flattenRouterbgpNeighborAllowasInEnableVpnv4(o["allowas-in-enable-vpnv4"], d, "allowas_in_enable_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable-vpnv4"]) {
			return fmt.Errorf("Error reading allowas_in_enable_vpnv4: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_vpnv6", flattenRouterbgpNeighborAllowasInEnableVpnv6(o["allowas-in-enable-vpnv6"], d, "allowas_in_enable_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable-vpnv6"]) {
			return fmt.Errorf("Error reading allowas_in_enable_vpnv6: %v", err)
		}
	}

	if err = d.Set("allowas_in_enable_evpn", flattenRouterbgpNeighborAllowasInEnableEvpn(o["allowas-in-enable-evpn"], d, "allowas_in_enable_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-enable-evpn"]) {
			return fmt.Errorf("Error reading allowas_in_enable_evpn: %v", err)
		}
	}

	if err = d.Set("allowas_in", flattenRouterbgpNeighborAllowasIn(o["allowas-in"], d, "allowas_in", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in"]) {
			return fmt.Errorf("Error reading allowas_in: %v", err)
		}
	}

	if err = d.Set("allowas_in6", flattenRouterbgpNeighborAllowasIn6(o["allowas-in6"], d, "allowas_in6", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in6"]) {
			return fmt.Errorf("Error reading allowas_in6: %v", err)
		}
	}

	if err = d.Set("allowas_in_vpnv4", flattenRouterbgpNeighborAllowasInVpnv4(o["allowas-in-vpnv4"], d, "allowas_in_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-vpnv4"]) {
			return fmt.Errorf("Error reading allowas_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("allowas_in_vpnv6", flattenRouterbgpNeighborAllowasInVpnv6(o["allowas-in-vpnv6"], d, "allowas_in_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-vpnv6"]) {
			return fmt.Errorf("Error reading allowas_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("allowas_in_evpn", flattenRouterbgpNeighborAllowasInEvpn(o["allowas-in-evpn"], d, "allowas_in_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["allowas-in-evpn"]) {
			return fmt.Errorf("Error reading allowas_in_evpn: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged", flattenRouterbgpNeighborAttributeUnchanged(o["attribute-unchanged"], d, "attribute_unchanged", sv)); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged"]) {
			return fmt.Errorf("Error reading attribute_unchanged: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged6", flattenRouterbgpNeighborAttributeUnchanged6(o["attribute-unchanged6"], d, "attribute_unchanged6", sv)); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged6"]) {
			return fmt.Errorf("Error reading attribute_unchanged6: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_vpnv4", flattenRouterbgpNeighborAttributeUnchangedVpnv4(o["attribute-unchanged-vpnv4"], d, "attribute_unchanged_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged-vpnv4"]) {
			return fmt.Errorf("Error reading attribute_unchanged_vpnv4: %v", err)
		}
	}

	if err = d.Set("attribute_unchanged_vpnv6", flattenRouterbgpNeighborAttributeUnchangedVpnv6(o["attribute-unchanged-vpnv6"], d, "attribute_unchanged_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["attribute-unchanged-vpnv6"]) {
			return fmt.Errorf("Error reading attribute_unchanged_vpnv6: %v", err)
		}
	}

	if err = d.Set("activate", flattenRouterbgpNeighborActivate(o["activate"], d, "activate", sv)); err != nil {
		if !fortiAPIPatch(o["activate"]) {
			return fmt.Errorf("Error reading activate: %v", err)
		}
	}

	if err = d.Set("activate6", flattenRouterbgpNeighborActivate6(o["activate6"], d, "activate6", sv)); err != nil {
		if !fortiAPIPatch(o["activate6"]) {
			return fmt.Errorf("Error reading activate6: %v", err)
		}
	}

	if err = d.Set("activate_vpnv4", flattenRouterbgpNeighborActivateVpnv4(o["activate-vpnv4"], d, "activate_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["activate-vpnv4"]) {
			return fmt.Errorf("Error reading activate_vpnv4: %v", err)
		}
	}

	if err = d.Set("activate_vpnv6", flattenRouterbgpNeighborActivateVpnv6(o["activate-vpnv6"], d, "activate_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["activate-vpnv6"]) {
			return fmt.Errorf("Error reading activate_vpnv6: %v", err)
		}
	}

	if err = d.Set("activate_evpn", flattenRouterbgpNeighborActivateEvpn(o["activate-evpn"], d, "activate_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["activate-evpn"]) {
			return fmt.Errorf("Error reading activate_evpn: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterbgpNeighborBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("capability_dynamic", flattenRouterbgpNeighborCapabilityDynamic(o["capability-dynamic"], d, "capability_dynamic", sv)); err != nil {
		if !fortiAPIPatch(o["capability-dynamic"]) {
			return fmt.Errorf("Error reading capability_dynamic: %v", err)
		}
	}

	if err = d.Set("capability_orf", flattenRouterbgpNeighborCapabilityOrf(o["capability-orf"], d, "capability_orf", sv)); err != nil {
		if !fortiAPIPatch(o["capability-orf"]) {
			return fmt.Errorf("Error reading capability_orf: %v", err)
		}
	}

	if err = d.Set("capability_orf6", flattenRouterbgpNeighborCapabilityOrf6(o["capability-orf6"], d, "capability_orf6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-orf6"]) {
			return fmt.Errorf("Error reading capability_orf6: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart", flattenRouterbgpNeighborCapabilityGracefulRestart(o["capability-graceful-restart"], d, "capability_graceful_restart", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart"]) {
			return fmt.Errorf("Error reading capability_graceful_restart: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart6", flattenRouterbgpNeighborCapabilityGracefulRestart6(o["capability-graceful-restart6"], d, "capability_graceful_restart6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart6"]) {
			return fmt.Errorf("Error reading capability_graceful_restart6: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_vpnv4", flattenRouterbgpNeighborCapabilityGracefulRestartVpnv4(o["capability-graceful-restart-vpnv4"], d, "capability_graceful_restart_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart-vpnv4"]) {
			return fmt.Errorf("Error reading capability_graceful_restart_vpnv4: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_vpnv6", flattenRouterbgpNeighborCapabilityGracefulRestartVpnv6(o["capability-graceful-restart-vpnv6"], d, "capability_graceful_restart_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart-vpnv6"]) {
			return fmt.Errorf("Error reading capability_graceful_restart_vpnv6: %v", err)
		}
	}

	if err = d.Set("capability_graceful_restart_evpn", flattenRouterbgpNeighborCapabilityGracefulRestartEvpn(o["capability-graceful-restart-evpn"], d, "capability_graceful_restart_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["capability-graceful-restart-evpn"]) {
			return fmt.Errorf("Error reading capability_graceful_restart_evpn: %v", err)
		}
	}

	if err = d.Set("capability_route_refresh", flattenRouterbgpNeighborCapabilityRouteRefresh(o["capability-route-refresh"], d, "capability_route_refresh", sv)); err != nil {
		if !fortiAPIPatch(o["capability-route-refresh"]) {
			return fmt.Errorf("Error reading capability_route_refresh: %v", err)
		}
	}

	if err = d.Set("capability_default_originate", flattenRouterbgpNeighborCapabilityDefaultOriginate(o["capability-default-originate"], d, "capability_default_originate", sv)); err != nil {
		if !fortiAPIPatch(o["capability-default-originate"]) {
			return fmt.Errorf("Error reading capability_default_originate: %v", err)
		}
	}

	if err = d.Set("capability_default_originate6", flattenRouterbgpNeighborCapabilityDefaultOriginate6(o["capability-default-originate6"], d, "capability_default_originate6", sv)); err != nil {
		if !fortiAPIPatch(o["capability-default-originate6"]) {
			return fmt.Errorf("Error reading capability_default_originate6: %v", err)
		}
	}

	if err = d.Set("dont_capability_negotiate", flattenRouterbgpNeighborDontCapabilityNegotiate(o["dont-capability-negotiate"], d, "dont_capability_negotiate", sv)); err != nil {
		if !fortiAPIPatch(o["dont-capability-negotiate"]) {
			return fmt.Errorf("Error reading dont_capability_negotiate: %v", err)
		}
	}

	if err = d.Set("ebgp_enforce_multihop", flattenRouterbgpNeighborEbgpEnforceMultihop(o["ebgp-enforce-multihop"], d, "ebgp_enforce_multihop", sv)); err != nil {
		if !fortiAPIPatch(o["ebgp-enforce-multihop"]) {
			return fmt.Errorf("Error reading ebgp_enforce_multihop: %v", err)
		}
	}

	if err = d.Set("link_down_failover", flattenRouterbgpNeighborLinkDownFailover(o["link-down-failover"], d, "link_down_failover", sv)); err != nil {
		if !fortiAPIPatch(o["link-down-failover"]) {
			return fmt.Errorf("Error reading link_down_failover: %v", err)
		}
	}

	if err = d.Set("stale_route", flattenRouterbgpNeighborStaleRoute(o["stale-route"], d, "stale_route", sv)); err != nil {
		if !fortiAPIPatch(o["stale-route"]) {
			return fmt.Errorf("Error reading stale_route: %v", err)
		}
	}

	if err = d.Set("next_hop_self", flattenRouterbgpNeighborNextHopSelf(o["next-hop-self"], d, "next_hop_self", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self"]) {
			return fmt.Errorf("Error reading next_hop_self: %v", err)
		}
	}

	if err = d.Set("next_hop_self6", flattenRouterbgpNeighborNextHopSelf6(o["next-hop-self6"], d, "next_hop_self6", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self6"]) {
			return fmt.Errorf("Error reading next_hop_self6: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr", flattenRouterbgpNeighborNextHopSelfRr(o["next-hop-self-rr"], d, "next_hop_self_rr", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr"]) {
			return fmt.Errorf("Error reading next_hop_self_rr: %v", err)
		}
	}

	if err = d.Set("next_hop_self_rr6", flattenRouterbgpNeighborNextHopSelfRr6(o["next-hop-self-rr6"], d, "next_hop_self_rr6", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self-rr6"]) {
			return fmt.Errorf("Error reading next_hop_self_rr6: %v", err)
		}
	}

	if err = d.Set("next_hop_self_vpnv4", flattenRouterbgpNeighborNextHopSelfVpnv4(o["next-hop-self-vpnv4"], d, "next_hop_self_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self-vpnv4"]) {
			return fmt.Errorf("Error reading next_hop_self_vpnv4: %v", err)
		}
	}

	if err = d.Set("next_hop_self_vpnv6", flattenRouterbgpNeighborNextHopSelfVpnv6(o["next-hop-self-vpnv6"], d, "next_hop_self_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["next-hop-self-vpnv6"]) {
			return fmt.Errorf("Error reading next_hop_self_vpnv6: %v", err)
		}
	}

	if err = d.Set("override_capability", flattenRouterbgpNeighborOverrideCapability(o["override-capability"], d, "override_capability", sv)); err != nil {
		if !fortiAPIPatch(o["override-capability"]) {
			return fmt.Errorf("Error reading override_capability: %v", err)
		}
	}

	if err = d.Set("passive", flattenRouterbgpNeighborPassive(o["passive"], d, "passive", sv)); err != nil {
		if !fortiAPIPatch(o["passive"]) {
			return fmt.Errorf("Error reading passive: %v", err)
		}
	}

	if err = d.Set("remove_private_as", flattenRouterbgpNeighborRemovePrivateAs(o["remove-private-as"], d, "remove_private_as", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as"]) {
			return fmt.Errorf("Error reading remove_private_as: %v", err)
		}
	}

	if err = d.Set("remove_private_as6", flattenRouterbgpNeighborRemovePrivateAs6(o["remove-private-as6"], d, "remove_private_as6", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as6"]) {
			return fmt.Errorf("Error reading remove_private_as6: %v", err)
		}
	}

	if err = d.Set("remove_private_as_vpnv4", flattenRouterbgpNeighborRemovePrivateAsVpnv4(o["remove-private-as-vpnv4"], d, "remove_private_as_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as-vpnv4"]) {
			return fmt.Errorf("Error reading remove_private_as_vpnv4: %v", err)
		}
	}

	if err = d.Set("remove_private_as_vpnv6", flattenRouterbgpNeighborRemovePrivateAsVpnv6(o["remove-private-as-vpnv6"], d, "remove_private_as_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as-vpnv6"]) {
			return fmt.Errorf("Error reading remove_private_as_vpnv6: %v", err)
		}
	}

	if err = d.Set("remove_private_as_evpn", flattenRouterbgpNeighborRemovePrivateAsEvpn(o["remove-private-as-evpn"], d, "remove_private_as_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["remove-private-as-evpn"]) {
			return fmt.Errorf("Error reading remove_private_as_evpn: %v", err)
		}
	}

	if err = d.Set("route_reflector_client", flattenRouterbgpNeighborRouteReflectorClient(o["route-reflector-client"], d, "route_reflector_client", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client"]) {
			return fmt.Errorf("Error reading route_reflector_client: %v", err)
		}
	}

	if err = d.Set("route_reflector_client6", flattenRouterbgpNeighborRouteReflectorClient6(o["route-reflector-client6"], d, "route_reflector_client6", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client6"]) {
			return fmt.Errorf("Error reading route_reflector_client6: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_vpnv4", flattenRouterbgpNeighborRouteReflectorClientVpnv4(o["route-reflector-client-vpnv4"], d, "route_reflector_client_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client-vpnv4"]) {
			return fmt.Errorf("Error reading route_reflector_client_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_vpnv6", flattenRouterbgpNeighborRouteReflectorClientVpnv6(o["route-reflector-client-vpnv6"], d, "route_reflector_client_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client-vpnv6"]) {
			return fmt.Errorf("Error reading route_reflector_client_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_reflector_client_evpn", flattenRouterbgpNeighborRouteReflectorClientEvpn(o["route-reflector-client-evpn"], d, "route_reflector_client_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["route-reflector-client-evpn"]) {
			return fmt.Errorf("Error reading route_reflector_client_evpn: %v", err)
		}
	}

	if err = d.Set("route_server_client", flattenRouterbgpNeighborRouteServerClient(o["route-server-client"], d, "route_server_client", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client"]) {
			return fmt.Errorf("Error reading route_server_client: %v", err)
		}
	}

	if err = d.Set("route_server_client6", flattenRouterbgpNeighborRouteServerClient6(o["route-server-client6"], d, "route_server_client6", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client6"]) {
			return fmt.Errorf("Error reading route_server_client6: %v", err)
		}
	}

	if err = d.Set("route_server_client_vpnv4", flattenRouterbgpNeighborRouteServerClientVpnv4(o["route-server-client-vpnv4"], d, "route_server_client_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client-vpnv4"]) {
			return fmt.Errorf("Error reading route_server_client_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_server_client_vpnv6", flattenRouterbgpNeighborRouteServerClientVpnv6(o["route-server-client-vpnv6"], d, "route_server_client_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client-vpnv6"]) {
			return fmt.Errorf("Error reading route_server_client_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_server_client_evpn", flattenRouterbgpNeighborRouteServerClientEvpn(o["route-server-client-evpn"], d, "route_server_client_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["route-server-client-evpn"]) {
			return fmt.Errorf("Error reading route_server_client_evpn: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change", flattenRouterbgpNeighborRrAttrAllowChange(o["rr-attr-allow-change"], d, "rr_attr_allow_change", sv)); err != nil {
		if !fortiAPIPatch(o["rr-attr-allow-change"]) {
			return fmt.Errorf("Error reading rr_attr_allow_change: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change6", flattenRouterbgpNeighborRrAttrAllowChange6(o["rr-attr-allow-change6"], d, "rr_attr_allow_change6", sv)); err != nil {
		if !fortiAPIPatch(o["rr-attr-allow-change6"]) {
			return fmt.Errorf("Error reading rr_attr_allow_change6: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change_vpnv4", flattenRouterbgpNeighborRrAttrAllowChangeVpnv4(o["rr-attr-allow-change-vpnv4"], d, "rr_attr_allow_change_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["rr-attr-allow-change-vpnv4"]) {
			return fmt.Errorf("Error reading rr_attr_allow_change_vpnv4: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change_vpnv6", flattenRouterbgpNeighborRrAttrAllowChangeVpnv6(o["rr-attr-allow-change-vpnv6"], d, "rr_attr_allow_change_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["rr-attr-allow-change-vpnv6"]) {
			return fmt.Errorf("Error reading rr_attr_allow_change_vpnv6: %v", err)
		}
	}

	if err = d.Set("rr_attr_allow_change_evpn", flattenRouterbgpNeighborRrAttrAllowChangeEvpn(o["rr-attr-allow-change-evpn"], d, "rr_attr_allow_change_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["rr-attr-allow-change-evpn"]) {
			return fmt.Errorf("Error reading rr_attr_allow_change_evpn: %v", err)
		}
	}

	if err = d.Set("shutdown", flattenRouterbgpNeighborShutdown(o["shutdown"], d, "shutdown", sv)); err != nil {
		if !fortiAPIPatch(o["shutdown"]) {
			return fmt.Errorf("Error reading shutdown: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration", flattenRouterbgpNeighborSoftReconfiguration(o["soft-reconfiguration"], d, "soft_reconfiguration", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration"]) {
			return fmt.Errorf("Error reading soft_reconfiguration: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration6", flattenRouterbgpNeighborSoftReconfiguration6(o["soft-reconfiguration6"], d, "soft_reconfiguration6", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration6"]) {
			return fmt.Errorf("Error reading soft_reconfiguration6: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_vpnv4", flattenRouterbgpNeighborSoftReconfigurationVpnv4(o["soft-reconfiguration-vpnv4"], d, "soft_reconfiguration_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration-vpnv4"]) {
			return fmt.Errorf("Error reading soft_reconfiguration_vpnv4: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_vpnv6", flattenRouterbgpNeighborSoftReconfigurationVpnv6(o["soft-reconfiguration-vpnv6"], d, "soft_reconfiguration_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration-vpnv6"]) {
			return fmt.Errorf("Error reading soft_reconfiguration_vpnv6: %v", err)
		}
	}

	if err = d.Set("soft_reconfiguration_evpn", flattenRouterbgpNeighborSoftReconfigurationEvpn(o["soft-reconfiguration-evpn"], d, "soft_reconfiguration_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["soft-reconfiguration-evpn"]) {
			return fmt.Errorf("Error reading soft_reconfiguration_evpn: %v", err)
		}
	}

	if err = d.Set("as_override", flattenRouterbgpNeighborAsOverride(o["as-override"], d, "as_override", sv)); err != nil {
		if !fortiAPIPatch(o["as-override"]) {
			return fmt.Errorf("Error reading as_override: %v", err)
		}
	}

	if err = d.Set("as_override6", flattenRouterbgpNeighborAsOverride6(o["as-override6"], d, "as_override6", sv)); err != nil {
		if !fortiAPIPatch(o["as-override6"]) {
			return fmt.Errorf("Error reading as_override6: %v", err)
		}
	}

	if err = d.Set("strict_capability_match", flattenRouterbgpNeighborStrictCapabilityMatch(o["strict-capability-match"], d, "strict_capability_match", sv)); err != nil {
		if !fortiAPIPatch(o["strict-capability-match"]) {
			return fmt.Errorf("Error reading strict_capability_match: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap", flattenRouterbgpNeighborDefaultOriginateRoutemap(o["default-originate-routemap"], d, "default_originate_routemap", sv)); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap"]) {
			return fmt.Errorf("Error reading default_originate_routemap: %v", err)
		}
	}

	if err = d.Set("default_originate_routemap6", flattenRouterbgpNeighborDefaultOriginateRoutemap6(o["default-originate-routemap6"], d, "default_originate_routemap6", sv)); err != nil {
		if !fortiAPIPatch(o["default-originate-routemap6"]) {
			return fmt.Errorf("Error reading default_originate_routemap6: %v", err)
		}
	}

	if err = d.Set("description", flattenRouterbgpNeighborDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterbgpNeighborDistributeListIn(o["distribute-list-in"], d, "distribute_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_list_in6", flattenRouterbgpNeighborDistributeListIn6(o["distribute-list-in6"], d, "distribute_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in6"]) {
			return fmt.Errorf("Error reading distribute_list_in6: %v", err)
		}
	}

	if err = d.Set("distribute_list_in_vpnv4", flattenRouterbgpNeighborDistributeListInVpnv4(o["distribute-list-in-vpnv4"], d, "distribute_list_in_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in-vpnv4"]) {
			return fmt.Errorf("Error reading distribute_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("distribute_list_in_vpnv6", flattenRouterbgpNeighborDistributeListInVpnv6(o["distribute-list-in-vpnv6"], d, "distribute_list_in_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in-vpnv6"]) {
			return fmt.Errorf("Error reading distribute_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out", flattenRouterbgpNeighborDistributeListOut(o["distribute-list-out"], d, "distribute_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out"]) {
			return fmt.Errorf("Error reading distribute_list_out: %v", err)
		}
	}

	if err = d.Set("distribute_list_out6", flattenRouterbgpNeighborDistributeListOut6(o["distribute-list-out6"], d, "distribute_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out6"]) {
			return fmt.Errorf("Error reading distribute_list_out6: %v", err)
		}
	}

	if err = d.Set("distribute_list_out_vpnv4", flattenRouterbgpNeighborDistributeListOutVpnv4(o["distribute-list-out-vpnv4"], d, "distribute_list_out_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out-vpnv4"]) {
			return fmt.Errorf("Error reading distribute_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("distribute_list_out_vpnv6", flattenRouterbgpNeighborDistributeListOutVpnv6(o["distribute-list-out-vpnv6"], d, "distribute_list_out_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-out-vpnv6"]) {
			return fmt.Errorf("Error reading distribute_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("ebgp_multihop_ttl", flattenRouterbgpNeighborEbgpMultihopTtl(o["ebgp-multihop-ttl"], d, "ebgp_multihop_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ebgp-multihop-ttl"]) {
			return fmt.Errorf("Error reading ebgp_multihop_ttl: %v", err)
		}
	}

	if err = d.Set("filter_list_in", flattenRouterbgpNeighborFilterListIn(o["filter-list-in"], d, "filter_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-in"]) {
			return fmt.Errorf("Error reading filter_list_in: %v", err)
		}
	}

	if err = d.Set("filter_list_in6", flattenRouterbgpNeighborFilterListIn6(o["filter-list-in6"], d, "filter_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-in6"]) {
			return fmt.Errorf("Error reading filter_list_in6: %v", err)
		}
	}

	if err = d.Set("filter_list_in_vpnv4", flattenRouterbgpNeighborFilterListInVpnv4(o["filter-list-in-vpnv4"], d, "filter_list_in_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-in-vpnv4"]) {
			return fmt.Errorf("Error reading filter_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("filter_list_in_vpnv6", flattenRouterbgpNeighborFilterListInVpnv6(o["filter-list-in-vpnv6"], d, "filter_list_in_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-in-vpnv6"]) {
			return fmt.Errorf("Error reading filter_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("filter_list_out", flattenRouterbgpNeighborFilterListOut(o["filter-list-out"], d, "filter_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-out"]) {
			return fmt.Errorf("Error reading filter_list_out: %v", err)
		}
	}

	if err = d.Set("filter_list_out6", flattenRouterbgpNeighborFilterListOut6(o["filter-list-out6"], d, "filter_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-out6"]) {
			return fmt.Errorf("Error reading filter_list_out6: %v", err)
		}
	}

	if err = d.Set("filter_list_out_vpnv4", flattenRouterbgpNeighborFilterListOutVpnv4(o["filter-list-out-vpnv4"], d, "filter_list_out_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-out-vpnv4"]) {
			return fmt.Errorf("Error reading filter_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("filter_list_out_vpnv6", flattenRouterbgpNeighborFilterListOutVpnv6(o["filter-list-out-vpnv6"], d, "filter_list_out_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["filter-list-out-vpnv6"]) {
			return fmt.Errorf("Error reading filter_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterbgpNeighborInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("maximum_prefix", flattenRouterbgpNeighborMaximumPrefix(o["maximum-prefix"], d, "maximum_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix"]) {
			return fmt.Errorf("Error reading maximum_prefix: %v", err)
		}
	}

	if err = d.Set("maximum_prefix6", flattenRouterbgpNeighborMaximumPrefix6(o["maximum-prefix6"], d, "maximum_prefix6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix6"]) {
			return fmt.Errorf("Error reading maximum_prefix6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_vpnv4", flattenRouterbgpNeighborMaximumPrefixVpnv4(o["maximum-prefix-vpnv4"], d, "maximum_prefix_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-vpnv4"]) {
			return fmt.Errorf("Error reading maximum_prefix_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_vpnv6", flattenRouterbgpNeighborMaximumPrefixVpnv6(o["maximum-prefix-vpnv6"], d, "maximum_prefix_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-vpnv6"]) {
			return fmt.Errorf("Error reading maximum_prefix_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_evpn", flattenRouterbgpNeighborMaximumPrefixEvpn(o["maximum-prefix-evpn"], d, "maximum_prefix_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-evpn"]) {
			return fmt.Errorf("Error reading maximum_prefix_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold", flattenRouterbgpNeighborMaximumPrefixThreshold(o["maximum-prefix-threshold"], d, "maximum_prefix_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold6", flattenRouterbgpNeighborMaximumPrefixThreshold6(o["maximum-prefix-threshold6"], d, "maximum_prefix_threshold6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold6"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_vpnv4", flattenRouterbgpNeighborMaximumPrefixThresholdVpnv4(o["maximum-prefix-threshold-vpnv4"], d, "maximum_prefix_threshold_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold-vpnv4"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_vpnv6", flattenRouterbgpNeighborMaximumPrefixThresholdVpnv6(o["maximum-prefix-threshold-vpnv6"], d, "maximum_prefix_threshold_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold-vpnv6"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_threshold_evpn", flattenRouterbgpNeighborMaximumPrefixThresholdEvpn(o["maximum-prefix-threshold-evpn"], d, "maximum_prefix_threshold_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-threshold-evpn"]) {
			return fmt.Errorf("Error reading maximum_prefix_threshold_evpn: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only", flattenRouterbgpNeighborMaximumPrefixWarningOnly(o["maximum-prefix-warning-only"], d, "maximum_prefix_warning_only", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only6", flattenRouterbgpNeighborMaximumPrefixWarningOnly6(o["maximum-prefix-warning-only6"], d, "maximum_prefix_warning_only6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only6"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_vpnv4", flattenRouterbgpNeighborMaximumPrefixWarningOnlyVpnv4(o["maximum-prefix-warning-only-vpnv4"], d, "maximum_prefix_warning_only_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only-vpnv4"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv4: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_vpnv6", flattenRouterbgpNeighborMaximumPrefixWarningOnlyVpnv6(o["maximum-prefix-warning-only-vpnv6"], d, "maximum_prefix_warning_only_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only-vpnv6"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_vpnv6: %v", err)
		}
	}

	if err = d.Set("maximum_prefix_warning_only_evpn", flattenRouterbgpNeighborMaximumPrefixWarningOnlyEvpn(o["maximum-prefix-warning-only-evpn"], d, "maximum_prefix_warning_only_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-prefix-warning-only-evpn"]) {
			return fmt.Errorf("Error reading maximum_prefix_warning_only_evpn: %v", err)
		}
	}

	if err = d.Set("prefix_list_in", flattenRouterbgpNeighborPrefixListIn(o["prefix-list-in"], d, "prefix_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in"]) {
			return fmt.Errorf("Error reading prefix_list_in: %v", err)
		}
	}

	if err = d.Set("prefix_list_in6", flattenRouterbgpNeighborPrefixListIn6(o["prefix-list-in6"], d, "prefix_list_in6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in6"]) {
			return fmt.Errorf("Error reading prefix_list_in6: %v", err)
		}
	}

	if err = d.Set("prefix_list_in_vpnv4", flattenRouterbgpNeighborPrefixListInVpnv4(o["prefix-list-in-vpnv4"], d, "prefix_list_in_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in-vpnv4"]) {
			return fmt.Errorf("Error reading prefix_list_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("prefix_list_in_vpnv6", flattenRouterbgpNeighborPrefixListInVpnv6(o["prefix-list-in-vpnv6"], d, "prefix_list_in_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-in-vpnv6"]) {
			return fmt.Errorf("Error reading prefix_list_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out", flattenRouterbgpNeighborPrefixListOut(o["prefix-list-out"], d, "prefix_list_out", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out"]) {
			return fmt.Errorf("Error reading prefix_list_out: %v", err)
		}
	}

	if err = d.Set("prefix_list_out6", flattenRouterbgpNeighborPrefixListOut6(o["prefix-list-out6"], d, "prefix_list_out6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out6"]) {
			return fmt.Errorf("Error reading prefix_list_out6: %v", err)
		}
	}

	if err = d.Set("prefix_list_out_vpnv4", flattenRouterbgpNeighborPrefixListOutVpnv4(o["prefix-list-out-vpnv4"], d, "prefix_list_out_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out-vpnv4"]) {
			return fmt.Errorf("Error reading prefix_list_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("prefix_list_out_vpnv6", flattenRouterbgpNeighborPrefixListOutVpnv6(o["prefix-list-out-vpnv6"], d, "prefix_list_out_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-list-out-vpnv6"]) {
			return fmt.Errorf("Error reading prefix_list_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("remote_as", flattenRouterbgpNeighborRemoteAs(o["remote-as"], d, "remote_as", sv)); err != nil {
		if !fortiAPIPatch(o["remote-as"]) {
			return fmt.Errorf("Error reading remote_as: %v", err)
		}
	}

	if err = d.Set("local_as", flattenRouterbgpNeighborLocalAs(o["local-as"], d, "local_as", sv)); err != nil {
		if !fortiAPIPatch(o["local-as"]) {
			return fmt.Errorf("Error reading local_as: %v", err)
		}
	}

	if err = d.Set("local_as_no_prepend", flattenRouterbgpNeighborLocalAsNoPrepend(o["local-as-no-prepend"], d, "local_as_no_prepend", sv)); err != nil {
		if !fortiAPIPatch(o["local-as-no-prepend"]) {
			return fmt.Errorf("Error reading local_as_no_prepend: %v", err)
		}
	}

	if err = d.Set("local_as_replace_as", flattenRouterbgpNeighborLocalAsReplaceAs(o["local-as-replace-as"], d, "local_as_replace_as", sv)); err != nil {
		if !fortiAPIPatch(o["local-as-replace-as"]) {
			return fmt.Errorf("Error reading local_as_replace_as: %v", err)
		}
	}

	if err = d.Set("retain_stale_time", flattenRouterbgpNeighborRetainStaleTime(o["retain-stale-time"], d, "retain_stale_time", sv)); err != nil {
		if !fortiAPIPatch(o["retain-stale-time"]) {
			return fmt.Errorf("Error reading retain_stale_time: %v", err)
		}
	}

	if err = d.Set("route_map_in", flattenRouterbgpNeighborRouteMapIn(o["route-map-in"], d, "route_map_in", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in"]) {
			return fmt.Errorf("Error reading route_map_in: %v", err)
		}
	}

	if err = d.Set("route_map_in6", flattenRouterbgpNeighborRouteMapIn6(o["route-map-in6"], d, "route_map_in6", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in6"]) {
			return fmt.Errorf("Error reading route_map_in6: %v", err)
		}
	}

	if err = d.Set("route_map_in_vpnv4", flattenRouterbgpNeighborRouteMapInVpnv4(o["route-map-in-vpnv4"], d, "route_map_in_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in-vpnv4"]) {
			return fmt.Errorf("Error reading route_map_in_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_map_in_vpnv6", flattenRouterbgpNeighborRouteMapInVpnv6(o["route-map-in-vpnv6"], d, "route_map_in_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in-vpnv6"]) {
			return fmt.Errorf("Error reading route_map_in_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_map_in_evpn", flattenRouterbgpNeighborRouteMapInEvpn(o["route-map-in-evpn"], d, "route_map_in_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-in-evpn"]) {
			return fmt.Errorf("Error reading route_map_in_evpn: %v", err)
		}
	}

	if err = d.Set("route_map_out", flattenRouterbgpNeighborRouteMapOut(o["route-map-out"], d, "route_map_out", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out"]) {
			return fmt.Errorf("Error reading route_map_out: %v", err)
		}
	}

	if err = d.Set("route_map_out_preferable", flattenRouterbgpNeighborRouteMapOutPreferable(o["route-map-out-preferable"], d, "route_map_out_preferable", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-preferable"]) {
			return fmt.Errorf("Error reading route_map_out_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out6", flattenRouterbgpNeighborRouteMapOut6(o["route-map-out6"], d, "route_map_out6", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out6"]) {
			return fmt.Errorf("Error reading route_map_out6: %v", err)
		}
	}

	if err = d.Set("route_map_out6_preferable", flattenRouterbgpNeighborRouteMapOut6Preferable(o["route-map-out6-preferable"], d, "route_map_out6_preferable", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out6-preferable"]) {
			return fmt.Errorf("Error reading route_map_out6_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv4", flattenRouterbgpNeighborRouteMapOutVpnv4(o["route-map-out-vpnv4"], d, "route_map_out_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-vpnv4"]) {
			return fmt.Errorf("Error reading route_map_out_vpnv4: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv6", flattenRouterbgpNeighborRouteMapOutVpnv6(o["route-map-out-vpnv6"], d, "route_map_out_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-vpnv6"]) {
			return fmt.Errorf("Error reading route_map_out_vpnv6: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv4_preferable", flattenRouterbgpNeighborRouteMapOutVpnv4Preferable(o["route-map-out-vpnv4-preferable"], d, "route_map_out_vpnv4_preferable", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-vpnv4-preferable"]) {
			return fmt.Errorf("Error reading route_map_out_vpnv4_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_vpnv6_preferable", flattenRouterbgpNeighborRouteMapOutVpnv6Preferable(o["route-map-out-vpnv6-preferable"], d, "route_map_out_vpnv6_preferable", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-vpnv6-preferable"]) {
			return fmt.Errorf("Error reading route_map_out_vpnv6_preferable: %v", err)
		}
	}

	if err = d.Set("route_map_out_evpn", flattenRouterbgpNeighborRouteMapOutEvpn(o["route-map-out-evpn"], d, "route_map_out_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["route-map-out-evpn"]) {
			return fmt.Errorf("Error reading route_map_out_evpn: %v", err)
		}
	}

	if err = d.Set("send_community", flattenRouterbgpNeighborSendCommunity(o["send-community"], d, "send_community", sv)); err != nil {
		if !fortiAPIPatch(o["send-community"]) {
			return fmt.Errorf("Error reading send_community: %v", err)
		}
	}

	if err = d.Set("send_community6", flattenRouterbgpNeighborSendCommunity6(o["send-community6"], d, "send_community6", sv)); err != nil {
		if !fortiAPIPatch(o["send-community6"]) {
			return fmt.Errorf("Error reading send_community6: %v", err)
		}
	}

	if err = d.Set("send_community_vpnv4", flattenRouterbgpNeighborSendCommunityVpnv4(o["send-community-vpnv4"], d, "send_community_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["send-community-vpnv4"]) {
			return fmt.Errorf("Error reading send_community_vpnv4: %v", err)
		}
	}

	if err = d.Set("send_community_vpnv6", flattenRouterbgpNeighborSendCommunityVpnv6(o["send-community-vpnv6"], d, "send_community_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["send-community-vpnv6"]) {
			return fmt.Errorf("Error reading send_community_vpnv6: %v", err)
		}
	}

	if err = d.Set("send_community_evpn", flattenRouterbgpNeighborSendCommunityEvpn(o["send-community-evpn"], d, "send_community_evpn", sv)); err != nil {
		if !fortiAPIPatch(o["send-community-evpn"]) {
			return fmt.Errorf("Error reading send_community_evpn: %v", err)
		}
	}

	if err = d.Set("keep_alive_timer", flattenRouterbgpNeighborKeepAliveTimer(o["keep-alive-timer"], d, "keep_alive_timer", sv)); err != nil {
		if !fortiAPIPatch(o["keep-alive-timer"]) {
			return fmt.Errorf("Error reading keep_alive_timer: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterbgpNeighborHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer", sv)); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("connect_timer", flattenRouterbgpNeighborConnectTimer(o["connect-timer"], d, "connect_timer", sv)); err != nil {
		if !fortiAPIPatch(o["connect-timer"]) {
			return fmt.Errorf("Error reading connect_timer: %v", err)
		}
	}

	if err = d.Set("unsuppress_map", flattenRouterbgpNeighborUnsuppressMap(o["unsuppress-map"], d, "unsuppress_map", sv)); err != nil {
		if !fortiAPIPatch(o["unsuppress-map"]) {
			return fmt.Errorf("Error reading unsuppress_map: %v", err)
		}
	}

	if err = d.Set("unsuppress_map6", flattenRouterbgpNeighborUnsuppressMap6(o["unsuppress-map6"], d, "unsuppress_map6", sv)); err != nil {
		if !fortiAPIPatch(o["unsuppress-map6"]) {
			return fmt.Errorf("Error reading unsuppress_map6: %v", err)
		}
	}

	if err = d.Set("update_source", flattenRouterbgpNeighborUpdateSource(o["update-source"], d, "update_source", sv)); err != nil {
		if !fortiAPIPatch(o["update-source"]) {
			return fmt.Errorf("Error reading update_source: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterbgpNeighborWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("restart_time", flattenRouterbgpNeighborRestartTime(o["restart-time"], d, "restart_time", sv)); err != nil {
		if !fortiAPIPatch(o["restart-time"]) {
			return fmt.Errorf("Error reading restart_time: %v", err)
		}
	}

	if err = d.Set("additional_path", flattenRouterbgpNeighborAdditionalPath(o["additional-path"], d, "additional_path", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("Error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterbgpNeighborAdditionalPath6(o["additional-path6"], d, "additional_path6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("Error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv4", flattenRouterbgpNeighborAdditionalPathVpnv4(o["additional-path-vpnv4"], d, "additional_path_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-vpnv4"]) {
			return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv6", flattenRouterbgpNeighborAdditionalPathVpnv6(o["additional-path-vpnv6"], d, "additional_path_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-vpnv6"]) {
			return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path", flattenRouterbgpNeighborAdvAdditionalPath(o["adv-additional-path"], d, "adv_additional_path", sv)); err != nil {
		if !fortiAPIPatch(o["adv-additional-path"]) {
			return fmt.Errorf("Error reading adv_additional_path: %v", err)
		}
	}

	if err = d.Set("adv_additional_path6", flattenRouterbgpNeighborAdvAdditionalPath6(o["adv-additional-path6"], d, "adv_additional_path6", sv)); err != nil {
		if !fortiAPIPatch(o["adv-additional-path6"]) {
			return fmt.Errorf("Error reading adv_additional_path6: %v", err)
		}
	}

	if err = d.Set("adv_additional_path_vpnv4", flattenRouterbgpNeighborAdvAdditionalPathVpnv4(o["adv-additional-path-vpnv4"], d, "adv_additional_path_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["adv-additional-path-vpnv4"]) {
			return fmt.Errorf("Error reading adv_additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("adv_additional_path_vpnv6", flattenRouterbgpNeighborAdvAdditionalPathVpnv6(o["adv-additional-path-vpnv6"], d, "adv_additional_path_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["adv-additional-path-vpnv6"]) {
			return fmt.Errorf("Error reading adv_additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("auth_options", flattenRouterbgpNeighborAuthOptions(o["auth-options"], d, "auth_options", sv)); err != nil {
		if !fortiAPIPatch(o["auth-options"]) {
			return fmt.Errorf("Error reading auth_options: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("conditional_advertise", flattenRouterbgpNeighborConditionalAdvertise(o["conditional-advertise"], d, "conditional_advertise", sv)); err != nil {
			if !fortiAPIPatch(o["conditional-advertise"]) {
				return fmt.Errorf("Error reading conditional_advertise: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("conditional_advertise"); ok {
			if err = d.Set("conditional_advertise", flattenRouterbgpNeighborConditionalAdvertise(o["conditional-advertise"], d, "conditional_advertise", sv)); err != nil {
				if !fortiAPIPatch(o["conditional-advertise"]) {
					return fmt.Errorf("Error reading conditional_advertise: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("conditional_advertise6", flattenRouterbgpNeighborConditionalAdvertise6(o["conditional-advertise6"], d, "conditional_advertise6", sv)); err != nil {
			if !fortiAPIPatch(o["conditional-advertise6"]) {
				return fmt.Errorf("Error reading conditional_advertise6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("conditional_advertise6"); ok {
			if err = d.Set("conditional_advertise6", flattenRouterbgpNeighborConditionalAdvertise6(o["conditional-advertise6"], d, "conditional_advertise6", sv)); err != nil {
				if !fortiAPIPatch(o["conditional-advertise6"]) {
					return fmt.Errorf("Error reading conditional_advertise6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterbgpNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterbgpNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInEnableVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInEnableVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInEnableEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAllowasInEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAttributeUnchangedVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAttributeUnchangedVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborActivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborActivate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborActivateVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborActivateVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborActivateEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityOrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityGracefulRestartVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityGracefulRestartVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityGracefulRestartEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborLinkDownFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborStaleRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborNextHopSelf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborNextHopSelf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborNextHopSelfVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborNextHopSelfVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborOverrideCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRemovePrivateAsVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRemovePrivateAsVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRemovePrivateAsEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteReflectorClientVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteReflectorClientVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteReflectorClientEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteServerClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteServerClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteServerClientVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteServerClientVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteServerClientEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRrAttrAllowChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRrAttrAllowChange6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRrAttrAllowChangeVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRrAttrAllowChangeVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRrAttrAllowChangeEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborShutdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSoftReconfigurationVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSoftReconfigurationVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSoftReconfigurationEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAsOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAsOverride6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborDistributeListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborFilterListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixThresholdVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixThresholdVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixThresholdEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixWarningOnlyVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixWarningOnlyVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborMaximumPrefixWarningOnlyEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPrefixListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterbgpNeighborLocalAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterbgpNeighborLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRetainStaleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapInEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOutVpnv4Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOutVpnv6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborRouteMapOutEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSendCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSendCommunity6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSendCommunityVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSendCommunityVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborSendCommunityEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if vi, ok := v.(int); ok && vi < 0 {
		return nil, nil
	}
	return v, nil
}

func expandRouterbgpNeighborHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if vi, ok := v.(int); ok && vi < 0 {
		return nil, nil
	}
	return v, nil
}

func expandRouterbgpNeighborConnectTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if vi, ok := v.(int); ok && vi < 0 {
		return nil, nil
	}
	return v, nil
}

func expandRouterbgpNeighborUnsuppressMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborUpdateSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if vi, ok := v.(int); ok && vi < 0 {
		return nil, nil
	}
	return v, nil
}

func expandRouterbgpNeighborRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdvAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAdvAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborAuthOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborConditionalAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise-routemap"], _ = expandRouterbgpNeighborConditionalAdvertiseAdvertiseRoutemap(d, i["advertise_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["advertise-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-routemap"], _ = expandRouterbgpNeighborConditionalAdvertiseConditionRoutemap(d, i["condition_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["condition-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-type"], _ = expandRouterbgpNeighborConditionalAdvertiseConditionType(d, i["condition_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterbgpNeighborConditionalAdvertiseAdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborConditionalAdvertiseConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"7.0.4"},
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

func expandRouterbgpNeighborConditionalAdvertiseConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborConditionalAdvertise6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise-routemap"], _ = expandRouterbgpNeighborConditionalAdvertise6AdvertiseRoutemap(d, i["advertise_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["advertise-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-routemap"], _ = expandRouterbgpNeighborConditionalAdvertise6ConditionRoutemap(d, i["condition_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["condition-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-type"], _ = expandRouterbgpNeighborConditionalAdvertise6ConditionType(d, i["condition_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterbgpNeighborConditionalAdvertise6AdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNeighborConditionalAdvertise6ConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"7.0.4"},
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

func expandRouterbgpNeighborConditionalAdvertise6ConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterbgpNeighbor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandRouterbgpNeighborIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOkExists("advertisement_interval"); ok {
		t, err := expandRouterbgpNeighborAdvertisementInterval(d, v, "advertisement_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertisement-interval"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable"); ok {
		t, err := expandRouterbgpNeighborAllowasInEnable(d, v, "allowas_in_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable6"); ok {
		t, err := expandRouterbgpNeighborAllowasInEnable6(d, v, "allowas_in_enable6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_vpnv4"); ok {
		t, err := expandRouterbgpNeighborAllowasInEnableVpnv4(d, v, "allowas_in_enable_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_vpnv6"); ok {
		t, err := expandRouterbgpNeighborAllowasInEnableVpnv6(d, v, "allowas_in_enable_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_enable_evpn"); ok {
		t, err := expandRouterbgpNeighborAllowasInEnableEvpn(d, v, "allowas_in_enable_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-enable-evpn"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in"); ok {
		t, err := expandRouterbgpNeighborAllowasIn(d, v, "allowas_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in6"); ok {
		t, err := expandRouterbgpNeighborAllowasIn6(d, v, "allowas_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_vpnv4"); ok {
		t, err := expandRouterbgpNeighborAllowasInVpnv4(d, v, "allowas_in_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_vpnv6"); ok {
		t, err := expandRouterbgpNeighborAllowasInVpnv6(d, v, "allowas_in_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("allowas_in_evpn"); ok {
		t, err := expandRouterbgpNeighborAllowasInEvpn(d, v, "allowas_in_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowas-in-evpn"] = t
		}
	}

	if v, ok := d.GetOk("attribute_unchanged"); ok {
		t, err := expandRouterbgpNeighborAttributeUnchanged(d, v, "attribute_unchanged", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged"] = t
		}
	} else if d.HasChange("attribute_unchanged") {
		obj["attribute-unchanged"] = nil
	}

	if v, ok := d.GetOk("attribute_unchanged6"); ok {
		t, err := expandRouterbgpNeighborAttributeUnchanged6(d, v, "attribute_unchanged6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged6"] = t
		}
	} else if d.HasChange("attribute_unchanged6") {
		obj["attribute-unchanged6"] = nil
	}

	if v, ok := d.GetOk("attribute_unchanged_vpnv4"); ok {
		t, err := expandRouterbgpNeighborAttributeUnchangedVpnv4(d, v, "attribute_unchanged_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged-vpnv4"] = t
		}
	} else if d.HasChange("attribute_unchanged_vpnv4") {
		obj["attribute-unchanged-vpnv4"] = nil
	}

	if v, ok := d.GetOk("attribute_unchanged_vpnv6"); ok {
		t, err := expandRouterbgpNeighborAttributeUnchangedVpnv6(d, v, "attribute_unchanged_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute-unchanged-vpnv6"] = t
		}
	} else if d.HasChange("attribute_unchanged_vpnv6") {
		obj["attribute-unchanged-vpnv6"] = nil
	}

	if v, ok := d.GetOk("activate"); ok {
		t, err := expandRouterbgpNeighborActivate(d, v, "activate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate"] = t
		}
	}

	if v, ok := d.GetOk("activate6"); ok {
		t, err := expandRouterbgpNeighborActivate6(d, v, "activate6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate6"] = t
		}
	}

	if v, ok := d.GetOk("activate_vpnv4"); ok {
		t, err := expandRouterbgpNeighborActivateVpnv4(d, v, "activate_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("activate_vpnv6"); ok {
		t, err := expandRouterbgpNeighborActivateVpnv6(d, v, "activate_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("activate_evpn"); ok {
		t, err := expandRouterbgpNeighborActivateEvpn(d, v, "activate_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activate-evpn"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterbgpNeighborBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("capability_dynamic"); ok {
		t, err := expandRouterbgpNeighborCapabilityDynamic(d, v, "capability_dynamic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-dynamic"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf"); ok {
		t, err := expandRouterbgpNeighborCapabilityOrf(d, v, "capability_orf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf"] = t
		}
	}

	if v, ok := d.GetOk("capability_orf6"); ok {
		t, err := expandRouterbgpNeighborCapabilityOrf6(d, v, "capability_orf6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-orf6"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart"); ok {
		t, err := expandRouterbgpNeighborCapabilityGracefulRestart(d, v, "capability_graceful_restart", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart6"); ok {
		t, err := expandRouterbgpNeighborCapabilityGracefulRestart6(d, v, "capability_graceful_restart6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart6"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_vpnv4"); ok {
		t, err := expandRouterbgpNeighborCapabilityGracefulRestartVpnv4(d, v, "capability_graceful_restart_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_vpnv6"); ok {
		t, err := expandRouterbgpNeighborCapabilityGracefulRestartVpnv6(d, v, "capability_graceful_restart_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("capability_graceful_restart_evpn"); ok {
		t, err := expandRouterbgpNeighborCapabilityGracefulRestartEvpn(d, v, "capability_graceful_restart_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-graceful-restart-evpn"] = t
		}
	}

	if v, ok := d.GetOk("capability_route_refresh"); ok {
		t, err := expandRouterbgpNeighborCapabilityRouteRefresh(d, v, "capability_route_refresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-route-refresh"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate"); ok {
		t, err := expandRouterbgpNeighborCapabilityDefaultOriginate(d, v, "capability_default_originate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate"] = t
		}
	}

	if v, ok := d.GetOk("capability_default_originate6"); ok {
		t, err := expandRouterbgpNeighborCapabilityDefaultOriginate6(d, v, "capability_default_originate6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["capability-default-originate6"] = t
		}
	}

	if v, ok := d.GetOk("dont_capability_negotiate"); ok {
		t, err := expandRouterbgpNeighborDontCapabilityNegotiate(d, v, "dont_capability_negotiate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dont-capability-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("ebgp_enforce_multihop"); ok {
		t, err := expandRouterbgpNeighborEbgpEnforceMultihop(d, v, "ebgp_enforce_multihop", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-enforce-multihop"] = t
		}
	}

	if v, ok := d.GetOk("link_down_failover"); ok {
		t, err := expandRouterbgpNeighborLinkDownFailover(d, v, "link_down_failover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-failover"] = t
		}
	}

	if v, ok := d.GetOk("stale_route"); ok {
		t, err := expandRouterbgpNeighborStaleRoute(d, v, "stale_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stale-route"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self"); ok {
		t, err := expandRouterbgpNeighborNextHopSelf(d, v, "next_hop_self", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self6"); ok {
		t, err := expandRouterbgpNeighborNextHopSelf6(d, v, "next_hop_self6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr"); ok {
		t, err := expandRouterbgpNeighborNextHopSelfRr(d, v, "next_hop_self_rr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_rr6"); ok {
		t, err := expandRouterbgpNeighborNextHopSelfRr6(d, v, "next_hop_self_rr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-rr6"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_vpnv4"); ok {
		t, err := expandRouterbgpNeighborNextHopSelfVpnv4(d, v, "next_hop_self_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("next_hop_self_vpnv6"); ok {
		t, err := expandRouterbgpNeighborNextHopSelfVpnv6(d, v, "next_hop_self_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-hop-self-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("override_capability"); ok {
		t, err := expandRouterbgpNeighborOverrideCapability(d, v, "override_capability", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-capability"] = t
		}
	}

	if v, ok := d.GetOk("passive"); ok {
		t, err := expandRouterbgpNeighborPassive(d, v, "passive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as"); ok {
		t, err := expandRouterbgpNeighborRemovePrivateAs(d, v, "remove_private_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as6"); ok {
		t, err := expandRouterbgpNeighborRemovePrivateAs6(d, v, "remove_private_as6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as6"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_vpnv4"); ok {
		t, err := expandRouterbgpNeighborRemovePrivateAsVpnv4(d, v, "remove_private_as_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_vpnv6"); ok {
		t, err := expandRouterbgpNeighborRemovePrivateAsVpnv6(d, v, "remove_private_as_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("remove_private_as_evpn"); ok {
		t, err := expandRouterbgpNeighborRemovePrivateAsEvpn(d, v, "remove_private_as_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-private-as-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client"); ok {
		t, err := expandRouterbgpNeighborRouteReflectorClient(d, v, "route_reflector_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client6"); ok {
		t, err := expandRouterbgpNeighborRouteReflectorClient6(d, v, "route_reflector_client6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client6"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_vpnv4"); ok {
		t, err := expandRouterbgpNeighborRouteReflectorClientVpnv4(d, v, "route_reflector_client_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_vpnv6"); ok {
		t, err := expandRouterbgpNeighborRouteReflectorClientVpnv6(d, v, "route_reflector_client_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_reflector_client_evpn"); ok {
		t, err := expandRouterbgpNeighborRouteReflectorClientEvpn(d, v, "route_reflector_client_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-reflector-client-evpn"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client"); ok {
		t, err := expandRouterbgpNeighborRouteServerClient(d, v, "route_server_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client6"); ok {
		t, err := expandRouterbgpNeighborRouteServerClient6(d, v, "route_server_client6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_vpnv4"); ok {
		t, err := expandRouterbgpNeighborRouteServerClientVpnv4(d, v, "route_server_client_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_vpnv6"); ok {
		t, err := expandRouterbgpNeighborRouteServerClientVpnv6(d, v, "route_server_client_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("route_server_client_evpn"); ok {
		t, err := expandRouterbgpNeighborRouteServerClientEvpn(d, v, "route_server_client_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-server-client-evpn"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change"); ok {
		t, err := expandRouterbgpNeighborRrAttrAllowChange(d, v, "rr_attr_allow_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change6"); ok {
		t, err := expandRouterbgpNeighborRrAttrAllowChange6(d, v, "rr_attr_allow_change6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change6"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change_vpnv4"); ok {
		t, err := expandRouterbgpNeighborRrAttrAllowChangeVpnv4(d, v, "rr_attr_allow_change_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change_vpnv6"); ok {
		t, err := expandRouterbgpNeighborRrAttrAllowChangeVpnv6(d, v, "rr_attr_allow_change_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("rr_attr_allow_change_evpn"); ok {
		t, err := expandRouterbgpNeighborRrAttrAllowChangeEvpn(d, v, "rr_attr_allow_change_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-attr-allow-change-evpn"] = t
		}
	}

	if v, ok := d.GetOk("shutdown"); ok {
		t, err := expandRouterbgpNeighborShutdown(d, v, "shutdown", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shutdown"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration"); ok {
		t, err := expandRouterbgpNeighborSoftReconfiguration(d, v, "soft_reconfiguration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration6"); ok {
		t, err := expandRouterbgpNeighborSoftReconfiguration6(d, v, "soft_reconfiguration6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration6"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_vpnv4"); ok {
		t, err := expandRouterbgpNeighborSoftReconfigurationVpnv4(d, v, "soft_reconfiguration_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_vpnv6"); ok {
		t, err := expandRouterbgpNeighborSoftReconfigurationVpnv6(d, v, "soft_reconfiguration_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("soft_reconfiguration_evpn"); ok {
		t, err := expandRouterbgpNeighborSoftReconfigurationEvpn(d, v, "soft_reconfiguration_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["soft-reconfiguration-evpn"] = t
		}
	}

	if v, ok := d.GetOk("as_override"); ok {
		t, err := expandRouterbgpNeighborAsOverride(d, v, "as_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override"] = t
		}
	}

	if v, ok := d.GetOk("as_override6"); ok {
		t, err := expandRouterbgpNeighborAsOverride6(d, v, "as_override6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["as-override6"] = t
		}
	}

	if v, ok := d.GetOk("strict_capability_match"); ok {
		t, err := expandRouterbgpNeighborStrictCapabilityMatch(d, v, "strict_capability_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strict-capability-match"] = t
		}
	}

	if v, ok := d.GetOk("default_originate_routemap"); ok {
		t, err := expandRouterbgpNeighborDefaultOriginateRoutemap(d, v, "default_originate_routemap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap"] = t
		}
	} else if d.HasChange("default_originate_routemap") {
		obj["default-originate-routemap"] = nil
	}

	if v, ok := d.GetOk("default_originate_routemap6"); ok {
		t, err := expandRouterbgpNeighborDefaultOriginateRoutemap6(d, v, "default_originate_routemap6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-originate-routemap6"] = t
		}
	} else if d.HasChange("default_originate_routemap6") {
		obj["default-originate-routemap6"] = nil
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandRouterbgpNeighborDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("distribute_list_in"); ok {
		t, err := expandRouterbgpNeighborDistributeListIn(d, v, "distribute_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	} else if d.HasChange("distribute_list_in") {
		obj["distribute-list-in"] = nil
	}

	if v, ok := d.GetOk("distribute_list_in6"); ok {
		t, err := expandRouterbgpNeighborDistributeListIn6(d, v, "distribute_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in6"] = t
		}
	} else if d.HasChange("distribute_list_in6") {
		obj["distribute-list-in6"] = nil
	}

	if v, ok := d.GetOk("distribute_list_in_vpnv4"); ok {
		t, err := expandRouterbgpNeighborDistributeListInVpnv4(d, v, "distribute_list_in_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in-vpnv4"] = t
		}
	} else if d.HasChange("distribute_list_in_vpnv4") {
		obj["distribute-list-in-vpnv4"] = nil
	}

	if v, ok := d.GetOk("distribute_list_in_vpnv6"); ok {
		t, err := expandRouterbgpNeighborDistributeListInVpnv6(d, v, "distribute_list_in_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in-vpnv6"] = t
		}
	} else if d.HasChange("distribute_list_in_vpnv6") {
		obj["distribute-list-in-vpnv6"] = nil
	}

	if v, ok := d.GetOk("distribute_list_out"); ok {
		t, err := expandRouterbgpNeighborDistributeListOut(d, v, "distribute_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out"] = t
		}
	} else if d.HasChange("distribute_list_out") {
		obj["distribute-list-out"] = nil
	}

	if v, ok := d.GetOk("distribute_list_out6"); ok {
		t, err := expandRouterbgpNeighborDistributeListOut6(d, v, "distribute_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out6"] = t
		}
	} else if d.HasChange("distribute_list_out6") {
		obj["distribute-list-out6"] = nil
	}

	if v, ok := d.GetOk("distribute_list_out_vpnv4"); ok {
		t, err := expandRouterbgpNeighborDistributeListOutVpnv4(d, v, "distribute_list_out_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out-vpnv4"] = t
		}
	} else if d.HasChange("distribute_list_out_vpnv4") {
		obj["distribute-list-out-vpnv4"] = nil
	}

	if v, ok := d.GetOk("distribute_list_out_vpnv6"); ok {
		t, err := expandRouterbgpNeighborDistributeListOutVpnv6(d, v, "distribute_list_out_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-out-vpnv6"] = t
		}
	} else if d.HasChange("distribute_list_out_vpnv6") {
		obj["distribute-list-out-vpnv6"] = nil
	}

	if v, ok := d.GetOk("ebgp_multihop_ttl"); ok {
		t, err := expandRouterbgpNeighborEbgpMultihopTtl(d, v, "ebgp_multihop_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ebgp-multihop-ttl"] = t
		}
	}

	if v, ok := d.GetOk("filter_list_in"); ok {
		t, err := expandRouterbgpNeighborFilterListIn(d, v, "filter_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in"] = t
		}
	} else if d.HasChange("filter_list_in") {
		obj["filter-list-in"] = nil
	}

	if v, ok := d.GetOk("filter_list_in6"); ok {
		t, err := expandRouterbgpNeighborFilterListIn6(d, v, "filter_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in6"] = t
		}
	} else if d.HasChange("filter_list_in6") {
		obj["filter-list-in6"] = nil
	}

	if v, ok := d.GetOk("filter_list_in_vpnv4"); ok {
		t, err := expandRouterbgpNeighborFilterListInVpnv4(d, v, "filter_list_in_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in-vpnv4"] = t
		}
	} else if d.HasChange("filter_list_in_vpnv4") {
		obj["filter-list-in-vpnv4"] = nil
	}

	if v, ok := d.GetOk("filter_list_in_vpnv6"); ok {
		t, err := expandRouterbgpNeighborFilterListInVpnv6(d, v, "filter_list_in_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-in-vpnv6"] = t
		}
	} else if d.HasChange("filter_list_in_vpnv6") {
		obj["filter-list-in-vpnv6"] = nil
	}

	if v, ok := d.GetOk("filter_list_out"); ok {
		t, err := expandRouterbgpNeighborFilterListOut(d, v, "filter_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out"] = t
		}
	} else if d.HasChange("filter_list_out") {
		obj["filter-list-out"] = nil
	}

	if v, ok := d.GetOk("filter_list_out6"); ok {
		t, err := expandRouterbgpNeighborFilterListOut6(d, v, "filter_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out6"] = t
		}
	} else if d.HasChange("filter_list_out6") {
		obj["filter-list-out6"] = nil
	}

	if v, ok := d.GetOk("filter_list_out_vpnv4"); ok {
		t, err := expandRouterbgpNeighborFilterListOutVpnv4(d, v, "filter_list_out_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out-vpnv4"] = t
		}
	} else if d.HasChange("filter_list_out_vpnv4") {
		obj["filter-list-out-vpnv4"] = nil
	}

	if v, ok := d.GetOk("filter_list_out_vpnv6"); ok {
		t, err := expandRouterbgpNeighborFilterListOutVpnv6(d, v, "filter_list_out_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-list-out-vpnv6"] = t
		}
	} else if d.HasChange("filter_list_out_vpnv6") {
		obj["filter-list-out-vpnv6"] = nil
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterbgpNeighborInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("maximum_prefix"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefix(d, v, "maximum_prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix"] = t
		}
	} else if d.HasChange("maximum_prefix") {
		obj["maximum-prefix"] = nil
	}

	if v, ok := d.GetOk("maximum_prefix6"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefix6(d, v, "maximum_prefix6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix6"] = t
		}
	} else if d.HasChange("maximum_prefix6") {
		obj["maximum-prefix6"] = nil
	}

	if v, ok := d.GetOk("maximum_prefix_vpnv4"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixVpnv4(d, v, "maximum_prefix_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-vpnv4"] = t
		}
	} else if d.HasChange("maximum_prefix_vpnv4") {
		obj["maximum-prefix-vpnv4"] = nil
	}

	if v, ok := d.GetOk("maximum_prefix_vpnv6"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixVpnv6(d, v, "maximum_prefix_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-vpnv6"] = t
		}
	} else if d.HasChange("maximum_prefix_vpnv6") {
		obj["maximum-prefix-vpnv6"] = nil
	}

	if v, ok := d.GetOk("maximum_prefix_evpn"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixEvpn(d, v, "maximum_prefix_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-evpn"] = t
		}
	} else if d.HasChange("maximum_prefix_evpn") {
		obj["maximum-prefix-evpn"] = nil
	}

	if v, ok := d.GetOk("maximum_prefix_threshold"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixThreshold(d, v, "maximum_prefix_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold6"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixThreshold6(d, v, "maximum_prefix_threshold6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_vpnv4"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixThresholdVpnv4(d, v, "maximum_prefix_threshold_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_vpnv6"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixThresholdVpnv6(d, v, "maximum_prefix_threshold_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_threshold_evpn"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixThresholdEvpn(d, v, "maximum_prefix_threshold_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-threshold-evpn"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixWarningOnly(d, v, "maximum_prefix_warning_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only6"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixWarningOnly6(d, v, "maximum_prefix_warning_only6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_vpnv4"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixWarningOnlyVpnv4(d, v, "maximum_prefix_warning_only_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_vpnv6"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixWarningOnlyVpnv6(d, v, "maximum_prefix_warning_only_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("maximum_prefix_warning_only_evpn"); ok {
		t, err := expandRouterbgpNeighborMaximumPrefixWarningOnlyEvpn(d, v, "maximum_prefix_warning_only_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-prefix-warning-only-evpn"] = t
		}
	}

	if v, ok := d.GetOk("prefix_list_in"); ok {
		t, err := expandRouterbgpNeighborPrefixListIn(d, v, "prefix_list_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in"] = t
		}
	} else if d.HasChange("prefix_list_in") {
		obj["prefix-list-in"] = nil
	}

	if v, ok := d.GetOk("prefix_list_in6"); ok {
		t, err := expandRouterbgpNeighborPrefixListIn6(d, v, "prefix_list_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in6"] = t
		}
	} else if d.HasChange("prefix_list_in6") {
		obj["prefix-list-in6"] = nil
	}

	if v, ok := d.GetOk("prefix_list_in_vpnv4"); ok {
		t, err := expandRouterbgpNeighborPrefixListInVpnv4(d, v, "prefix_list_in_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in-vpnv4"] = t
		}
	} else if d.HasChange("prefix_list_in_vpnv4") {
		obj["prefix-list-in-vpnv4"] = nil
	}

	if v, ok := d.GetOk("prefix_list_in_vpnv6"); ok {
		t, err := expandRouterbgpNeighborPrefixListInVpnv6(d, v, "prefix_list_in_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-in-vpnv6"] = t
		}
	} else if d.HasChange("prefix_list_in_vpnv6") {
		obj["prefix-list-in-vpnv6"] = nil
	}

	if v, ok := d.GetOk("prefix_list_out"); ok {
		t, err := expandRouterbgpNeighborPrefixListOut(d, v, "prefix_list_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out"] = t
		}
	} else if d.HasChange("prefix_list_out") {
		obj["prefix-list-out"] = nil
	}

	if v, ok := d.GetOk("prefix_list_out6"); ok {
		t, err := expandRouterbgpNeighborPrefixListOut6(d, v, "prefix_list_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out6"] = t
		}
	} else if d.HasChange("prefix_list_out6") {
		obj["prefix-list-out6"] = nil
	}

	if v, ok := d.GetOk("prefix_list_out_vpnv4"); ok {
		t, err := expandRouterbgpNeighborPrefixListOutVpnv4(d, v, "prefix_list_out_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out-vpnv4"] = t
		}
	} else if d.HasChange("prefix_list_out_vpnv4") {
		obj["prefix-list-out-vpnv4"] = nil
	}

	if v, ok := d.GetOk("prefix_list_out_vpnv6"); ok {
		t, err := expandRouterbgpNeighborPrefixListOutVpnv6(d, v, "prefix_list_out_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-list-out-vpnv6"] = t
		}
	} else if d.HasChange("prefix_list_out_vpnv6") {
		obj["prefix-list-out-vpnv6"] = nil
	}

	if v, ok := d.GetOk("remote_as"); ok {
		t, err := expandRouterbgpNeighborRemoteAs(d, v, "remote_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-as"] = t
		}
	} else if d.HasChange("remote_as") {
		obj["remote-as"] = nil
	}

	if v, ok := d.GetOkExists("local_as"); ok {
		t, err := expandRouterbgpNeighborLocalAs(d, v, "local_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as"] = t
		}
	} else if d.HasChange("local_as") {
		obj["local-as"] = nil
	}

	if v, ok := d.GetOk("local_as_no_prepend"); ok {
		t, err := expandRouterbgpNeighborLocalAsNoPrepend(d, v, "local_as_no_prepend", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-no-prepend"] = t
		}
	}

	if v, ok := d.GetOk("local_as_replace_as"); ok {
		t, err := expandRouterbgpNeighborLocalAsReplaceAs(d, v, "local_as_replace_as", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-as-replace-as"] = t
		}
	}

	if v, ok := d.GetOkExists("retain_stale_time"); ok {
		t, err := expandRouterbgpNeighborRetainStaleTime(d, v, "retain_stale_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retain-stale-time"] = t
		}
	} else if d.HasChange("retain_stale_time") {
		obj["retain-stale-time"] = nil
	}

	if v, ok := d.GetOk("route_map_in"); ok {
		t, err := expandRouterbgpNeighborRouteMapIn(d, v, "route_map_in", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in"] = t
		}
	} else if d.HasChange("route_map_in") {
		obj["route-map-in"] = nil
	}

	if v, ok := d.GetOk("route_map_in6"); ok {
		t, err := expandRouterbgpNeighborRouteMapIn6(d, v, "route_map_in6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in6"] = t
		}
	} else if d.HasChange("route_map_in6") {
		obj["route-map-in6"] = nil
	}

	if v, ok := d.GetOk("route_map_in_vpnv4"); ok {
		t, err := expandRouterbgpNeighborRouteMapInVpnv4(d, v, "route_map_in_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-vpnv4"] = t
		}
	} else if d.HasChange("route_map_in_vpnv4") {
		obj["route-map-in-vpnv4"] = nil
	}

	if v, ok := d.GetOk("route_map_in_vpnv6"); ok {
		t, err := expandRouterbgpNeighborRouteMapInVpnv6(d, v, "route_map_in_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-vpnv6"] = t
		}
	} else if d.HasChange("route_map_in_vpnv6") {
		obj["route-map-in-vpnv6"] = nil
	}

	if v, ok := d.GetOk("route_map_in_evpn"); ok {
		t, err := expandRouterbgpNeighborRouteMapInEvpn(d, v, "route_map_in_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-in-evpn"] = t
		}
	} else if d.HasChange("route_map_in_evpn") {
		obj["route-map-in-evpn"] = nil
	}

	if v, ok := d.GetOk("route_map_out"); ok {
		t, err := expandRouterbgpNeighborRouteMapOut(d, v, "route_map_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out"] = t
		}
	} else if d.HasChange("route_map_out") {
		obj["route-map-out"] = nil
	}

	if v, ok := d.GetOk("route_map_out_preferable"); ok {
		t, err := expandRouterbgpNeighborRouteMapOutPreferable(d, v, "route_map_out_preferable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-preferable"] = t
		}
	} else if d.HasChange("route_map_out_preferable") {
		obj["route-map-out-preferable"] = nil
	}

	if v, ok := d.GetOk("route_map_out6"); ok {
		t, err := expandRouterbgpNeighborRouteMapOut6(d, v, "route_map_out6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6"] = t
		}
	} else if d.HasChange("route_map_out6") {
		obj["route-map-out6"] = nil
	}

	if v, ok := d.GetOk("route_map_out6_preferable"); ok {
		t, err := expandRouterbgpNeighborRouteMapOut6Preferable(d, v, "route_map_out6_preferable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out6-preferable"] = t
		}
	} else if d.HasChange("route_map_out6_preferable") {
		obj["route-map-out6-preferable"] = nil
	}

	if v, ok := d.GetOk("route_map_out_vpnv4"); ok {
		t, err := expandRouterbgpNeighborRouteMapOutVpnv4(d, v, "route_map_out_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv4"] = t
		}
	} else if d.HasChange("route_map_out_vpnv4") {
		obj["route-map-out-vpnv4"] = nil
	}

	if v, ok := d.GetOk("route_map_out_vpnv6"); ok {
		t, err := expandRouterbgpNeighborRouteMapOutVpnv6(d, v, "route_map_out_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv6"] = t
		}
	} else if d.HasChange("route_map_out_vpnv6") {
		obj["route-map-out-vpnv6"] = nil
	}

	if v, ok := d.GetOk("route_map_out_vpnv4_preferable"); ok {
		t, err := expandRouterbgpNeighborRouteMapOutVpnv4Preferable(d, v, "route_map_out_vpnv4_preferable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv4-preferable"] = t
		}
	} else if d.HasChange("route_map_out_vpnv4_preferable") {
		obj["route-map-out-vpnv4-preferable"] = nil
	}

	if v, ok := d.GetOk("route_map_out_vpnv6_preferable"); ok {
		t, err := expandRouterbgpNeighborRouteMapOutVpnv6Preferable(d, v, "route_map_out_vpnv6_preferable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-vpnv6-preferable"] = t
		}
	} else if d.HasChange("route_map_out_vpnv6_preferable") {
		obj["route-map-out-vpnv6-preferable"] = nil
	}

	if v, ok := d.GetOk("route_map_out_evpn"); ok {
		t, err := expandRouterbgpNeighborRouteMapOutEvpn(d, v, "route_map_out_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map-out-evpn"] = t
		}
	} else if d.HasChange("route_map_out_evpn") {
		obj["route-map-out-evpn"] = nil
	}

	if v, ok := d.GetOk("send_community"); ok {
		t, err := expandRouterbgpNeighborSendCommunity(d, v, "send_community", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community"] = t
		}
	}

	if v, ok := d.GetOk("send_community6"); ok {
		t, err := expandRouterbgpNeighborSendCommunity6(d, v, "send_community6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community6"] = t
		}
	}

	if v, ok := d.GetOk("send_community_vpnv4"); ok {
		t, err := expandRouterbgpNeighborSendCommunityVpnv4(d, v, "send_community_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("send_community_vpnv6"); ok {
		t, err := expandRouterbgpNeighborSendCommunityVpnv6(d, v, "send_community_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("send_community_evpn"); ok {
		t, err := expandRouterbgpNeighborSendCommunityEvpn(d, v, "send_community_evpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-community-evpn"] = t
		}
	}

	if v, ok := d.GetOkExists("keep_alive_timer"); ok {
		t, err := expandRouterbgpNeighborKeepAliveTimer(d, v, "keep_alive_timer", sv)
		if err != nil {
			return &obj, err
		}
		obj["keep-alive-timer"] = t
	}

	if v, ok := d.GetOk("holdtime_timer"); ok {
		t, err := expandRouterbgpNeighborHoldtimeTimer(d, v, "holdtime_timer", sv)
		if err != nil {
			return &obj, err
		}
		obj["holdtime-timer"] = t
	}

	if v, ok := d.GetOkExists("connect_timer"); ok {
		t, err := expandRouterbgpNeighborConnectTimer(d, v, "connect_timer", sv)
		if err != nil {
			return &obj, err
		}
		obj["connect-timer"] = t
	}

	if v, ok := d.GetOk("unsuppress_map"); ok {
		t, err := expandRouterbgpNeighborUnsuppressMap(d, v, "unsuppress_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map"] = t
		}
	} else if d.HasChange("unsuppress_map") {
		obj["unsuppress-map"] = nil
	}

	if v, ok := d.GetOk("unsuppress_map6"); ok {
		t, err := expandRouterbgpNeighborUnsuppressMap6(d, v, "unsuppress_map6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unsuppress-map6"] = t
		}
	} else if d.HasChange("unsuppress_map6") {
		obj["unsuppress-map6"] = nil
	}

	if v, ok := d.GetOk("update_source"); ok {
		t, err := expandRouterbgpNeighborUpdateSource(d, v, "update_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-source"] = t
		}
	} else if d.HasChange("update_source") {
		obj["update-source"] = nil
	}

	if v, ok := d.GetOkExists("weight"); ok {
		t, err := expandRouterbgpNeighborWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		}
		obj["weight"] = t
	}

	if v, ok := d.GetOkExists("restart_time"); ok {
		t, err := expandRouterbgpNeighborRestartTime(d, v, "restart_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-time"] = t
		}
	} else if d.HasChange("restart_time") {
		obj["restart-time"] = nil
	}

	if v, ok := d.GetOk("additional_path"); ok {
		t, err := expandRouterbgpNeighborAdditionalPath(d, v, "additional_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path"] = t
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok {
		t, err := expandRouterbgpNeighborAdditionalPath6(d, v, "additional_path6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv4"); ok {
		t, err := expandRouterbgpNeighborAdditionalPathVpnv4(d, v, "additional_path_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv6"); ok {
		t, err := expandRouterbgpNeighborAdditionalPathVpnv6(d, v, "additional_path_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path"); ok {
		t, err := expandRouterbgpNeighborAdvAdditionalPath(d, v, "adv_additional_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path6"); ok {
		t, err := expandRouterbgpNeighborAdvAdditionalPath6(d, v, "adv_additional_path6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path6"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path_vpnv4"); ok {
		t, err := expandRouterbgpNeighborAdvAdditionalPathVpnv4(d, v, "adv_additional_path_vpnv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path-vpnv4"] = t
		}
	}

	if v, ok := d.GetOk("adv_additional_path_vpnv6"); ok {
		t, err := expandRouterbgpNeighborAdvAdditionalPathVpnv6(d, v, "adv_additional_path_vpnv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-additional-path-vpnv6"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandRouterbgpNeighborPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("auth_options"); ok {
		t, err := expandRouterbgpNeighborAuthOptions(d, v, "auth_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-options"] = t
		}
	} else if d.HasChange("auth_options") {
		obj["auth-options"] = nil
	}

	if v, ok := d.GetOk("conditional_advertise"); ok || d.HasChange("conditional_advertise") {
		t, err := expandRouterbgpNeighborConditionalAdvertise(d, v, "conditional_advertise", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conditional-advertise"] = t
		}
	}

	if v, ok := d.GetOk("conditional_advertise6"); ok || d.HasChange("conditional_advertise6") {
		t, err := expandRouterbgpNeighborConditionalAdvertise6(d, v, "conditional_advertise6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conditional-advertise6"] = t
		}
	}

	return &obj, nil
}
