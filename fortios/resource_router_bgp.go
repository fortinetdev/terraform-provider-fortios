// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure BGP.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterBgpUpdate,
		Read:   resourceRouterBgpRead,
		Update: resourceRouterBgpUpdate,
		Delete: resourceRouterBgpDelete,

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
			"as_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"as": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"keepalive_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"holdtime_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 65535),
				Optional:     true,
				Computed:     true,
			},
			"always_compare_med": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_as_path_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_cmp_confed_aspath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_cmp_routerid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_med_confed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bestpath_med_missing_as_worst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_to_client_reflection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dampening": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deterministic_med": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ebgp_multipath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibgp_multipath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_first_as": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_external_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_import_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ignore_optional_capability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"multipath_recursive_distance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recursive_next_hop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"recursive_inherit_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag_resolve_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"confederation_identifier": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"confederation_peers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dampening_route_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"dampening_reachability_half_life": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 45),
				Optional:     true,
				Computed:     true,
			},
			"dampening_reuse": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20000),
				Optional:     true,
				Computed:     true,
			},
			"dampening_suppress": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20000),
				Optional:     true,
				Computed:     true,
			},
			"dampening_max_suppress_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"dampening_unreachability_half_life": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 45),
				Optional:     true,
				Computed:     true,
			},
			"default_local_preference": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"scan_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),
				Optional:     true,
				Computed:     true,
			},
			"distance_external": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"distance_internal": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"distance_local": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"synchronization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graceful_restart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graceful_restart_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"graceful_stalepath_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"graceful_update_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"graceful_end_on_timer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_path_select": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"additional_path_select6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"additional_path_select_vpnv4": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"additional_path_select_vpnv6": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 255),
				Optional:     true,
				Computed:     true,
			},
			"cross_family_conditional_adv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"aggregate_address6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),
							Optional:     true,
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
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"holdtime_timer": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 65535),
							Optional:     true,
							Computed:     true,
						},
						"connect_timer": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
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
							ValidateFunc: validation.IntBetween(0, 65535),
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
					},
				},
			},
			"neighbor_group": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),
							Optional:     true,
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
						"remote_as_filter": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
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
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"holdtime_timer": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 65535),
							Optional:     true,
							Computed:     true,
						},
						"connect_timer": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
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
							ValidateFunc: validation.IntBetween(0, 65535),
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
					},
				},
			},
			"neighbor_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_neighbor_num": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 1000),
							Optional:     true,
						},
						"neighbor_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"neighbor_range6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_neighbor_num": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 1000),
							Optional:     true,
						},
						"neighbor_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_import_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"prefix_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"network6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_import_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"admin_distance": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"neighbour_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"route_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"distance": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
						},
					},
				},
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"export_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"import_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"import_route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"leak_target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 7),
										Optional:     true,
									},
									"route_map": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"vrf6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"export_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"import_rt": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"route_target": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"import_route_map": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"leak_target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 7),
										Optional:     true,
									},
									"route_map": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"vrf_leak": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 7),
										Optional:     true,
									},
									"route_map": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"vrf_leak6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 7),
										Optional:     true,
									},
									"route_map": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
								},
							},
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

func resourceRouterBgpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterBgp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterBgp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterBgp")
	}

	return resourceRouterBgpRead(d, m)
}

func resourceRouterBgpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterBgp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterBgp resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterBgp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterBgp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterBgp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterBgp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterBgp resource from API: %v", err)
	}
	return nil
}

func flattenRouterBgpAsString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpKeepaliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathAsPathIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpConfedAspath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathMedConfed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpBestpathMedMissingAsWorst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpClientToClientReflection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampening(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpEbgpMultipath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpIbgpMultipath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpFastExternalFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkImportCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpIgnoreOptionalCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpMultipathRecursiveDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRecursiveNextHop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRecursiveInheritPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpTagResolveMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpClusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpConfederationIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpConfederationPeers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if cur_v, ok := i["peer"]; ok {
			tmp["peer"] = flattenRouterBgpConfederationPeersPeer(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "peer", d)
	return result
}

func flattenRouterBgpConfederationPeersPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDampeningReuse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDampeningSuppress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDampeningMaxSuppressTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDefaultLocalPreference(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpScanTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDistanceExternal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDistanceInternal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpDistanceLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpSynchronization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpGracefulRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpGracefulStalepathTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpGracefulUpdateDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpGracefulEndOnTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdditionalPathSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAdditionalPathSelect6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAdditionalPathSelectVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAdditionalPathSelectVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpCrossFamilyConditionalAdv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpAggregateAddressId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if cur_v, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterBgpAggregateAddressPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if cur_v, ok := i["as-set"]; ok {
			tmp["as_set"] = flattenRouterBgpAggregateAddressAsSet(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if cur_v, ok := i["summary-only"]; ok {
			tmp["summary_only"] = flattenRouterBgpAggregateAddressSummaryOnly(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpAggregateAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAggregateAddressPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBgpAggregateAddressAsSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddressSummaryOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpAggregateAddress6Id(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if cur_v, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterBgpAggregateAddress6Prefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if cur_v, ok := i["as-set"]; ok {
			tmp["as_set"] = flattenRouterBgpAggregateAddress6AsSet(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if cur_v, ok := i["summary-only"]; ok {
			tmp["summary_only"] = flattenRouterBgpAggregateAddress6SummaryOnly(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpAggregateAddress6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAggregateAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6AsSet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAggregateAddress6SummaryOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenRouterBgpNeighborIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if cur_v, ok := i["advertisement-interval"]; ok {
			tmp["advertisement_interval"] = flattenRouterBgpNeighborAdvertisementInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if cur_v, ok := i["allowas-in-enable"]; ok {
			tmp["allowas_in_enable"] = flattenRouterBgpNeighborAllowasInEnable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if cur_v, ok := i["allowas-in-enable6"]; ok {
			tmp["allowas_in_enable6"] = flattenRouterBgpNeighborAllowasInEnable6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if cur_v, ok := i["allowas-in-enable-vpnv4"]; ok {
			tmp["allowas_in_enable_vpnv4"] = flattenRouterBgpNeighborAllowasInEnableVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if cur_v, ok := i["allowas-in-enable-vpnv6"]; ok {
			tmp["allowas_in_enable_vpnv6"] = flattenRouterBgpNeighborAllowasInEnableVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if cur_v, ok := i["allowas-in-enable-evpn"]; ok {
			tmp["allowas_in_enable_evpn"] = flattenRouterBgpNeighborAllowasInEnableEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if cur_v, ok := i["allowas-in"]; ok {
			tmp["allowas_in"] = flattenRouterBgpNeighborAllowasIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if cur_v, ok := i["allowas-in6"]; ok {
			tmp["allowas_in6"] = flattenRouterBgpNeighborAllowasIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if cur_v, ok := i["allowas-in-vpnv4"]; ok {
			tmp["allowas_in_vpnv4"] = flattenRouterBgpNeighborAllowasInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if cur_v, ok := i["allowas-in-vpnv6"]; ok {
			tmp["allowas_in_vpnv6"] = flattenRouterBgpNeighborAllowasInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if cur_v, ok := i["allowas-in-evpn"]; ok {
			tmp["allowas_in_evpn"] = flattenRouterBgpNeighborAllowasInEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if cur_v, ok := i["attribute-unchanged"]; ok {
			tmp["attribute_unchanged"] = flattenRouterBgpNeighborAttributeUnchanged(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if cur_v, ok := i["attribute-unchanged6"]; ok {
			tmp["attribute_unchanged6"] = flattenRouterBgpNeighborAttributeUnchanged6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if cur_v, ok := i["attribute-unchanged-vpnv4"]; ok {
			tmp["attribute_unchanged_vpnv4"] = flattenRouterBgpNeighborAttributeUnchangedVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if cur_v, ok := i["attribute-unchanged-vpnv6"]; ok {
			tmp["attribute_unchanged_vpnv6"] = flattenRouterBgpNeighborAttributeUnchangedVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if cur_v, ok := i["activate"]; ok {
			tmp["activate"] = flattenRouterBgpNeighborActivate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if cur_v, ok := i["activate6"]; ok {
			tmp["activate6"] = flattenRouterBgpNeighborActivate6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if cur_v, ok := i["activate-vpnv4"]; ok {
			tmp["activate_vpnv4"] = flattenRouterBgpNeighborActivateVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if cur_v, ok := i["activate-vpnv6"]; ok {
			tmp["activate_vpnv6"] = flattenRouterBgpNeighborActivateVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if cur_v, ok := i["activate-evpn"]; ok {
			tmp["activate_evpn"] = flattenRouterBgpNeighborActivateEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if cur_v, ok := i["bfd"]; ok {
			tmp["bfd"] = flattenRouterBgpNeighborBfd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if cur_v, ok := i["capability-dynamic"]; ok {
			tmp["capability_dynamic"] = flattenRouterBgpNeighborCapabilityDynamic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if cur_v, ok := i["capability-orf"]; ok {
			tmp["capability_orf"] = flattenRouterBgpNeighborCapabilityOrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if cur_v, ok := i["capability-orf6"]; ok {
			tmp["capability_orf6"] = flattenRouterBgpNeighborCapabilityOrf6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if cur_v, ok := i["capability-graceful-restart"]; ok {
			tmp["capability_graceful_restart"] = flattenRouterBgpNeighborCapabilityGracefulRestart(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if cur_v, ok := i["capability-graceful-restart6"]; ok {
			tmp["capability_graceful_restart6"] = flattenRouterBgpNeighborCapabilityGracefulRestart6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if cur_v, ok := i["capability-graceful-restart-vpnv4"]; ok {
			tmp["capability_graceful_restart_vpnv4"] = flattenRouterBgpNeighborCapabilityGracefulRestartVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if cur_v, ok := i["capability-graceful-restart-vpnv6"]; ok {
			tmp["capability_graceful_restart_vpnv6"] = flattenRouterBgpNeighborCapabilityGracefulRestartVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if cur_v, ok := i["capability-graceful-restart-evpn"]; ok {
			tmp["capability_graceful_restart_evpn"] = flattenRouterBgpNeighborCapabilityGracefulRestartEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if cur_v, ok := i["capability-route-refresh"]; ok {
			tmp["capability_route_refresh"] = flattenRouterBgpNeighborCapabilityRouteRefresh(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if cur_v, ok := i["capability-default-originate"]; ok {
			tmp["capability_default_originate"] = flattenRouterBgpNeighborCapabilityDefaultOriginate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if cur_v, ok := i["capability-default-originate6"]; ok {
			tmp["capability_default_originate6"] = flattenRouterBgpNeighborCapabilityDefaultOriginate6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if cur_v, ok := i["dont-capability-negotiate"]; ok {
			tmp["dont_capability_negotiate"] = flattenRouterBgpNeighborDontCapabilityNegotiate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if cur_v, ok := i["ebgp-enforce-multihop"]; ok {
			tmp["ebgp_enforce_multihop"] = flattenRouterBgpNeighborEbgpEnforceMultihop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if cur_v, ok := i["link-down-failover"]; ok {
			tmp["link_down_failover"] = flattenRouterBgpNeighborLinkDownFailover(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if cur_v, ok := i["stale-route"]; ok {
			tmp["stale_route"] = flattenRouterBgpNeighborStaleRoute(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if cur_v, ok := i["next-hop-self"]; ok {
			tmp["next_hop_self"] = flattenRouterBgpNeighborNextHopSelf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if cur_v, ok := i["next-hop-self6"]; ok {
			tmp["next_hop_self6"] = flattenRouterBgpNeighborNextHopSelf6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if cur_v, ok := i["next-hop-self-rr"]; ok {
			tmp["next_hop_self_rr"] = flattenRouterBgpNeighborNextHopSelfRr(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if cur_v, ok := i["next-hop-self-rr6"]; ok {
			tmp["next_hop_self_rr6"] = flattenRouterBgpNeighborNextHopSelfRr6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if cur_v, ok := i["next-hop-self-vpnv4"]; ok {
			tmp["next_hop_self_vpnv4"] = flattenRouterBgpNeighborNextHopSelfVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if cur_v, ok := i["next-hop-self-vpnv6"]; ok {
			tmp["next_hop_self_vpnv6"] = flattenRouterBgpNeighborNextHopSelfVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if cur_v, ok := i["override-capability"]; ok {
			tmp["override_capability"] = flattenRouterBgpNeighborOverrideCapability(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if cur_v, ok := i["passive"]; ok {
			tmp["passive"] = flattenRouterBgpNeighborPassive(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if cur_v, ok := i["remove-private-as"]; ok {
			tmp["remove_private_as"] = flattenRouterBgpNeighborRemovePrivateAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if cur_v, ok := i["remove-private-as6"]; ok {
			tmp["remove_private_as6"] = flattenRouterBgpNeighborRemovePrivateAs6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if cur_v, ok := i["remove-private-as-vpnv4"]; ok {
			tmp["remove_private_as_vpnv4"] = flattenRouterBgpNeighborRemovePrivateAsVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if cur_v, ok := i["remove-private-as-vpnv6"]; ok {
			tmp["remove_private_as_vpnv6"] = flattenRouterBgpNeighborRemovePrivateAsVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if cur_v, ok := i["remove-private-as-evpn"]; ok {
			tmp["remove_private_as_evpn"] = flattenRouterBgpNeighborRemovePrivateAsEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if cur_v, ok := i["route-reflector-client"]; ok {
			tmp["route_reflector_client"] = flattenRouterBgpNeighborRouteReflectorClient(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if cur_v, ok := i["route-reflector-client6"]; ok {
			tmp["route_reflector_client6"] = flattenRouterBgpNeighborRouteReflectorClient6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if cur_v, ok := i["route-reflector-client-vpnv4"]; ok {
			tmp["route_reflector_client_vpnv4"] = flattenRouterBgpNeighborRouteReflectorClientVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if cur_v, ok := i["route-reflector-client-vpnv6"]; ok {
			tmp["route_reflector_client_vpnv6"] = flattenRouterBgpNeighborRouteReflectorClientVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if cur_v, ok := i["route-reflector-client-evpn"]; ok {
			tmp["route_reflector_client_evpn"] = flattenRouterBgpNeighborRouteReflectorClientEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if cur_v, ok := i["route-server-client"]; ok {
			tmp["route_server_client"] = flattenRouterBgpNeighborRouteServerClient(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if cur_v, ok := i["route-server-client6"]; ok {
			tmp["route_server_client6"] = flattenRouterBgpNeighborRouteServerClient6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if cur_v, ok := i["route-server-client-vpnv4"]; ok {
			tmp["route_server_client_vpnv4"] = flattenRouterBgpNeighborRouteServerClientVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if cur_v, ok := i["route-server-client-vpnv6"]; ok {
			tmp["route_server_client_vpnv6"] = flattenRouterBgpNeighborRouteServerClientVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if cur_v, ok := i["route-server-client-evpn"]; ok {
			tmp["route_server_client_evpn"] = flattenRouterBgpNeighborRouteServerClientEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if cur_v, ok := i["shutdown"]; ok {
			tmp["shutdown"] = flattenRouterBgpNeighborShutdown(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if cur_v, ok := i["soft-reconfiguration"]; ok {
			tmp["soft_reconfiguration"] = flattenRouterBgpNeighborSoftReconfiguration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if cur_v, ok := i["soft-reconfiguration6"]; ok {
			tmp["soft_reconfiguration6"] = flattenRouterBgpNeighborSoftReconfiguration6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if cur_v, ok := i["soft-reconfiguration-vpnv4"]; ok {
			tmp["soft_reconfiguration_vpnv4"] = flattenRouterBgpNeighborSoftReconfigurationVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if cur_v, ok := i["soft-reconfiguration-vpnv6"]; ok {
			tmp["soft_reconfiguration_vpnv6"] = flattenRouterBgpNeighborSoftReconfigurationVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if cur_v, ok := i["soft-reconfiguration-evpn"]; ok {
			tmp["soft_reconfiguration_evpn"] = flattenRouterBgpNeighborSoftReconfigurationEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if cur_v, ok := i["as-override"]; ok {
			tmp["as_override"] = flattenRouterBgpNeighborAsOverride(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if cur_v, ok := i["as-override6"]; ok {
			tmp["as_override6"] = flattenRouterBgpNeighborAsOverride6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if cur_v, ok := i["strict-capability-match"]; ok {
			tmp["strict_capability_match"] = flattenRouterBgpNeighborStrictCapabilityMatch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if cur_v, ok := i["default-originate-routemap"]; ok {
			tmp["default_originate_routemap"] = flattenRouterBgpNeighborDefaultOriginateRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if cur_v, ok := i["default-originate-routemap6"]; ok {
			tmp["default_originate_routemap6"] = flattenRouterBgpNeighborDefaultOriginateRoutemap6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenRouterBgpNeighborDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if cur_v, ok := i["distribute-list-in"]; ok {
			tmp["distribute_list_in"] = flattenRouterBgpNeighborDistributeListIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if cur_v, ok := i["distribute-list-in6"]; ok {
			tmp["distribute_list_in6"] = flattenRouterBgpNeighborDistributeListIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if cur_v, ok := i["distribute-list-in-vpnv4"]; ok {
			tmp["distribute_list_in_vpnv4"] = flattenRouterBgpNeighborDistributeListInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if cur_v, ok := i["distribute-list-in-vpnv6"]; ok {
			tmp["distribute_list_in_vpnv6"] = flattenRouterBgpNeighborDistributeListInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if cur_v, ok := i["distribute-list-out"]; ok {
			tmp["distribute_list_out"] = flattenRouterBgpNeighborDistributeListOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if cur_v, ok := i["distribute-list-out6"]; ok {
			tmp["distribute_list_out6"] = flattenRouterBgpNeighborDistributeListOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if cur_v, ok := i["distribute-list-out-vpnv4"]; ok {
			tmp["distribute_list_out_vpnv4"] = flattenRouterBgpNeighborDistributeListOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if cur_v, ok := i["distribute-list-out-vpnv6"]; ok {
			tmp["distribute_list_out_vpnv6"] = flattenRouterBgpNeighborDistributeListOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if cur_v, ok := i["ebgp-multihop-ttl"]; ok {
			tmp["ebgp_multihop_ttl"] = flattenRouterBgpNeighborEbgpMultihopTtl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if cur_v, ok := i["filter-list-in"]; ok {
			tmp["filter_list_in"] = flattenRouterBgpNeighborFilterListIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if cur_v, ok := i["filter-list-in6"]; ok {
			tmp["filter_list_in6"] = flattenRouterBgpNeighborFilterListIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if cur_v, ok := i["filter-list-in-vpnv4"]; ok {
			tmp["filter_list_in_vpnv4"] = flattenRouterBgpNeighborFilterListInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if cur_v, ok := i["filter-list-in-vpnv6"]; ok {
			tmp["filter_list_in_vpnv6"] = flattenRouterBgpNeighborFilterListInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if cur_v, ok := i["filter-list-out"]; ok {
			tmp["filter_list_out"] = flattenRouterBgpNeighborFilterListOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if cur_v, ok := i["filter-list-out6"]; ok {
			tmp["filter_list_out6"] = flattenRouterBgpNeighborFilterListOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if cur_v, ok := i["filter-list-out-vpnv4"]; ok {
			tmp["filter_list_out_vpnv4"] = flattenRouterBgpNeighborFilterListOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if cur_v, ok := i["filter-list-out-vpnv6"]; ok {
			tmp["filter_list_out_vpnv6"] = flattenRouterBgpNeighborFilterListOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBgpNeighborInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if cur_v, ok := i["maximum-prefix"]; ok {
			tmp["maximum_prefix"] = flattenRouterBgpNeighborMaximumPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if cur_v, ok := i["maximum-prefix6"]; ok {
			tmp["maximum_prefix6"] = flattenRouterBgpNeighborMaximumPrefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if cur_v, ok := i["maximum-prefix-vpnv4"]; ok {
			tmp["maximum_prefix_vpnv4"] = flattenRouterBgpNeighborMaximumPrefixVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if cur_v, ok := i["maximum-prefix-vpnv6"]; ok {
			tmp["maximum_prefix_vpnv6"] = flattenRouterBgpNeighborMaximumPrefixVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if cur_v, ok := i["maximum-prefix-evpn"]; ok {
			tmp["maximum_prefix_evpn"] = flattenRouterBgpNeighborMaximumPrefixEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if cur_v, ok := i["maximum-prefix-threshold"]; ok {
			tmp["maximum_prefix_threshold"] = flattenRouterBgpNeighborMaximumPrefixThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if cur_v, ok := i["maximum-prefix-threshold6"]; ok {
			tmp["maximum_prefix_threshold6"] = flattenRouterBgpNeighborMaximumPrefixThreshold6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if cur_v, ok := i["maximum-prefix-threshold-vpnv4"]; ok {
			tmp["maximum_prefix_threshold_vpnv4"] = flattenRouterBgpNeighborMaximumPrefixThresholdVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if cur_v, ok := i["maximum-prefix-threshold-vpnv6"]; ok {
			tmp["maximum_prefix_threshold_vpnv6"] = flattenRouterBgpNeighborMaximumPrefixThresholdVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if cur_v, ok := i["maximum-prefix-threshold-evpn"]; ok {
			tmp["maximum_prefix_threshold_evpn"] = flattenRouterBgpNeighborMaximumPrefixThresholdEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if cur_v, ok := i["maximum-prefix-warning-only"]; ok {
			tmp["maximum_prefix_warning_only"] = flattenRouterBgpNeighborMaximumPrefixWarningOnly(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if cur_v, ok := i["maximum-prefix-warning-only6"]; ok {
			tmp["maximum_prefix_warning_only6"] = flattenRouterBgpNeighborMaximumPrefixWarningOnly6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if cur_v, ok := i["maximum-prefix-warning-only-vpnv4"]; ok {
			tmp["maximum_prefix_warning_only_vpnv4"] = flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if cur_v, ok := i["maximum-prefix-warning-only-vpnv6"]; ok {
			tmp["maximum_prefix_warning_only_vpnv6"] = flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if cur_v, ok := i["maximum-prefix-warning-only-evpn"]; ok {
			tmp["maximum_prefix_warning_only_evpn"] = flattenRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if cur_v, ok := i["prefix-list-in"]; ok {
			tmp["prefix_list_in"] = flattenRouterBgpNeighborPrefixListIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if cur_v, ok := i["prefix-list-in6"]; ok {
			tmp["prefix_list_in6"] = flattenRouterBgpNeighborPrefixListIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if cur_v, ok := i["prefix-list-in-vpnv4"]; ok {
			tmp["prefix_list_in_vpnv4"] = flattenRouterBgpNeighborPrefixListInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if cur_v, ok := i["prefix-list-in-vpnv6"]; ok {
			tmp["prefix_list_in_vpnv6"] = flattenRouterBgpNeighborPrefixListInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if cur_v, ok := i["prefix-list-out"]; ok {
			tmp["prefix_list_out"] = flattenRouterBgpNeighborPrefixListOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if cur_v, ok := i["prefix-list-out6"]; ok {
			tmp["prefix_list_out6"] = flattenRouterBgpNeighborPrefixListOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if cur_v, ok := i["prefix-list-out-vpnv4"]; ok {
			tmp["prefix_list_out_vpnv4"] = flattenRouterBgpNeighborPrefixListOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if cur_v, ok := i["prefix-list-out-vpnv6"]; ok {
			tmp["prefix_list_out_vpnv6"] = flattenRouterBgpNeighborPrefixListOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if cur_v, ok := i["remote-as"]; ok {
			tmp["remote_as"] = flattenRouterBgpNeighborRemoteAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if cur_v, ok := i["local-as"]; ok {
			tmp["local_as"] = flattenRouterBgpNeighborLocalAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if cur_v, ok := i["local-as-no-prepend"]; ok {
			tmp["local_as_no_prepend"] = flattenRouterBgpNeighborLocalAsNoPrepend(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if cur_v, ok := i["local-as-replace-as"]; ok {
			tmp["local_as_replace_as"] = flattenRouterBgpNeighborLocalAsReplaceAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if cur_v, ok := i["retain-stale-time"]; ok {
			tmp["retain_stale_time"] = flattenRouterBgpNeighborRetainStaleTime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if cur_v, ok := i["route-map-in"]; ok {
			tmp["route_map_in"] = flattenRouterBgpNeighborRouteMapIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if cur_v, ok := i["route-map-in6"]; ok {
			tmp["route_map_in6"] = flattenRouterBgpNeighborRouteMapIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if cur_v, ok := i["route-map-in-vpnv4"]; ok {
			tmp["route_map_in_vpnv4"] = flattenRouterBgpNeighborRouteMapInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if cur_v, ok := i["route-map-in-vpnv6"]; ok {
			tmp["route_map_in_vpnv6"] = flattenRouterBgpNeighborRouteMapInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if cur_v, ok := i["route-map-in-evpn"]; ok {
			tmp["route_map_in_evpn"] = flattenRouterBgpNeighborRouteMapInEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if cur_v, ok := i["route-map-out"]; ok {
			tmp["route_map_out"] = flattenRouterBgpNeighborRouteMapOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if cur_v, ok := i["route-map-out-preferable"]; ok {
			tmp["route_map_out_preferable"] = flattenRouterBgpNeighborRouteMapOutPreferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if cur_v, ok := i["route-map-out6"]; ok {
			tmp["route_map_out6"] = flattenRouterBgpNeighborRouteMapOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if cur_v, ok := i["route-map-out6-preferable"]; ok {
			tmp["route_map_out6_preferable"] = flattenRouterBgpNeighborRouteMapOut6Preferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if cur_v, ok := i["route-map-out-vpnv4"]; ok {
			tmp["route_map_out_vpnv4"] = flattenRouterBgpNeighborRouteMapOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if cur_v, ok := i["route-map-out-vpnv6"]; ok {
			tmp["route_map_out_vpnv6"] = flattenRouterBgpNeighborRouteMapOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if cur_v, ok := i["route-map-out-vpnv4-preferable"]; ok {
			tmp["route_map_out_vpnv4_preferable"] = flattenRouterBgpNeighborRouteMapOutVpnv4Preferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if cur_v, ok := i["route-map-out-vpnv6-preferable"]; ok {
			tmp["route_map_out_vpnv6_preferable"] = flattenRouterBgpNeighborRouteMapOutVpnv6Preferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if cur_v, ok := i["route-map-out-evpn"]; ok {
			tmp["route_map_out_evpn"] = flattenRouterBgpNeighborRouteMapOutEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if cur_v, ok := i["send-community"]; ok {
			tmp["send_community"] = flattenRouterBgpNeighborSendCommunity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if cur_v, ok := i["send-community6"]; ok {
			tmp["send_community6"] = flattenRouterBgpNeighborSendCommunity6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if cur_v, ok := i["send-community-vpnv4"]; ok {
			tmp["send_community_vpnv4"] = flattenRouterBgpNeighborSendCommunityVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if cur_v, ok := i["send-community-vpnv6"]; ok {
			tmp["send_community_vpnv6"] = flattenRouterBgpNeighborSendCommunityVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if cur_v, ok := i["send-community-evpn"]; ok {
			tmp["send_community_evpn"] = flattenRouterBgpNeighborSendCommunityEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if cur_v, ok := i["keep-alive-timer"]; ok {
			tmp["keep_alive_timer"] = flattenRouterBgpNeighborKeepAliveTimer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if cur_v, ok := i["holdtime-timer"]; ok {
			tmp["holdtime_timer"] = flattenRouterBgpNeighborHoldtimeTimer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if cur_v, ok := i["connect-timer"]; ok {
			tmp["connect_timer"] = flattenRouterBgpNeighborConnectTimer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if cur_v, ok := i["unsuppress-map"]; ok {
			tmp["unsuppress_map"] = flattenRouterBgpNeighborUnsuppressMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if cur_v, ok := i["unsuppress-map6"]; ok {
			tmp["unsuppress_map6"] = flattenRouterBgpNeighborUnsuppressMap6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if cur_v, ok := i["update-source"]; ok {
			tmp["update_source"] = flattenRouterBgpNeighborUpdateSource(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if cur_v, ok := i["weight"]; ok {
			tmp["weight"] = flattenRouterBgpNeighborWeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if cur_v, ok := i["restart-time"]; ok {
			tmp["restart_time"] = flattenRouterBgpNeighborRestartTime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if cur_v, ok := i["additional-path"]; ok {
			tmp["additional_path"] = flattenRouterBgpNeighborAdditionalPath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if cur_v, ok := i["additional-path6"]; ok {
			tmp["additional_path6"] = flattenRouterBgpNeighborAdditionalPath6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if cur_v, ok := i["additional-path-vpnv4"]; ok {
			tmp["additional_path_vpnv4"] = flattenRouterBgpNeighborAdditionalPathVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if cur_v, ok := i["additional-path-vpnv6"]; ok {
			tmp["additional_path_vpnv6"] = flattenRouterBgpNeighborAdditionalPathVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if cur_v, ok := i["adv-additional-path"]; ok {
			tmp["adv_additional_path"] = flattenRouterBgpNeighborAdvAdditionalPath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if cur_v, ok := i["adv-additional-path6"]; ok {
			tmp["adv_additional_path6"] = flattenRouterBgpNeighborAdvAdditionalPath6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if cur_v, ok := i["adv-additional-path-vpnv4"]; ok {
			tmp["adv_additional_path_vpnv4"] = flattenRouterBgpNeighborAdvAdditionalPathVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if cur_v, ok := i["adv-additional-path-vpnv6"]; ok {
			tmp["adv_additional_path_vpnv6"] = flattenRouterBgpNeighborAdvAdditionalPathVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if cur_v, ok := i["auth-options"]; ok {
			tmp["auth_options"] = flattenRouterBgpNeighborAuthOptions(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if cur_v, ok := i["conditional-advertise"]; ok {
			tmp["conditional_advertise"] = flattenRouterBgpNeighborConditionalAdvertise(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise6"
		if cur_v, ok := i["conditional-advertise6"]; ok {
			tmp["conditional_advertise6"] = flattenRouterBgpNeighborConditionalAdvertise6(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenRouterBgpNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasInEnableEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAllowasInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAllowasInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAllowasInEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchangedVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAttributeUnchangedVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborActivateEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityGracefulRestartEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborLinkDownFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborStaleRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborNextHopSelfVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemovePrivateAsEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteReflectorClientEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteServerClientEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSoftReconfigurationEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborDistributeListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborFilterListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixThresholdVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixThresholdVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixThresholdEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborPrefixListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborLocalAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRetainStaleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapInEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutVpnv4Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutVpnv6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRouteMapOutEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborSendCommunityEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAdvAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAdvAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborAuthOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["advertise_routemap"] = flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if cur_v, ok := i["condition-routemap"]; ok {
			tmp["condition_routemap"] = flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if cur_v, ok := i["condition-type"]; ok {
			tmp["condition_type"] = flattenRouterBgpNeighborConditionalAdvertiseConditionType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "advertise_routemap", d)
	return result
}

func flattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("condition_routemap"), "name")
}

func flattenRouterBgpNeighborConditionalAdvertiseConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["advertise_routemap"] = flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if cur_v, ok := i["condition-routemap"]; ok {
			tmp["condition_routemap"] = flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if cur_v, ok := i["condition-type"]; ok {
			tmp["condition_type"] = flattenRouterBgpNeighborConditionalAdvertise6ConditionType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "advertise_routemap", d)
	return result
}

func flattenRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("condition_routemap"), "name")
}

func flattenRouterBgpNeighborConditionalAdvertise6ConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterBgpNeighborGroupName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if cur_v, ok := i["advertisement-interval"]; ok {
			tmp["advertisement_interval"] = flattenRouterBgpNeighborGroupAdvertisementInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if cur_v, ok := i["allowas-in-enable"]; ok {
			tmp["allowas_in_enable"] = flattenRouterBgpNeighborGroupAllowasInEnable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if cur_v, ok := i["allowas-in-enable6"]; ok {
			tmp["allowas_in_enable6"] = flattenRouterBgpNeighborGroupAllowasInEnable6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if cur_v, ok := i["allowas-in-enable-vpnv4"]; ok {
			tmp["allowas_in_enable_vpnv4"] = flattenRouterBgpNeighborGroupAllowasInEnableVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if cur_v, ok := i["allowas-in-enable-vpnv6"]; ok {
			tmp["allowas_in_enable_vpnv6"] = flattenRouterBgpNeighborGroupAllowasInEnableVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if cur_v, ok := i["allowas-in-enable-evpn"]; ok {
			tmp["allowas_in_enable_evpn"] = flattenRouterBgpNeighborGroupAllowasInEnableEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if cur_v, ok := i["allowas-in"]; ok {
			tmp["allowas_in"] = flattenRouterBgpNeighborGroupAllowasIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if cur_v, ok := i["allowas-in6"]; ok {
			tmp["allowas_in6"] = flattenRouterBgpNeighborGroupAllowasIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if cur_v, ok := i["allowas-in-vpnv4"]; ok {
			tmp["allowas_in_vpnv4"] = flattenRouterBgpNeighborGroupAllowasInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if cur_v, ok := i["allowas-in-vpnv6"]; ok {
			tmp["allowas_in_vpnv6"] = flattenRouterBgpNeighborGroupAllowasInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if cur_v, ok := i["allowas-in-evpn"]; ok {
			tmp["allowas_in_evpn"] = flattenRouterBgpNeighborGroupAllowasInEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if cur_v, ok := i["attribute-unchanged"]; ok {
			tmp["attribute_unchanged"] = flattenRouterBgpNeighborGroupAttributeUnchanged(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if cur_v, ok := i["attribute-unchanged6"]; ok {
			tmp["attribute_unchanged6"] = flattenRouterBgpNeighborGroupAttributeUnchanged6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if cur_v, ok := i["attribute-unchanged-vpnv4"]; ok {
			tmp["attribute_unchanged_vpnv4"] = flattenRouterBgpNeighborGroupAttributeUnchangedVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if cur_v, ok := i["attribute-unchanged-vpnv6"]; ok {
			tmp["attribute_unchanged_vpnv6"] = flattenRouterBgpNeighborGroupAttributeUnchangedVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if cur_v, ok := i["activate"]; ok {
			tmp["activate"] = flattenRouterBgpNeighborGroupActivate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if cur_v, ok := i["activate6"]; ok {
			tmp["activate6"] = flattenRouterBgpNeighborGroupActivate6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if cur_v, ok := i["activate-vpnv4"]; ok {
			tmp["activate_vpnv4"] = flattenRouterBgpNeighborGroupActivateVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if cur_v, ok := i["activate-vpnv6"]; ok {
			tmp["activate_vpnv6"] = flattenRouterBgpNeighborGroupActivateVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if cur_v, ok := i["activate-evpn"]; ok {
			tmp["activate_evpn"] = flattenRouterBgpNeighborGroupActivateEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if cur_v, ok := i["bfd"]; ok {
			tmp["bfd"] = flattenRouterBgpNeighborGroupBfd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if cur_v, ok := i["capability-dynamic"]; ok {
			tmp["capability_dynamic"] = flattenRouterBgpNeighborGroupCapabilityDynamic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if cur_v, ok := i["capability-orf"]; ok {
			tmp["capability_orf"] = flattenRouterBgpNeighborGroupCapabilityOrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if cur_v, ok := i["capability-orf6"]; ok {
			tmp["capability_orf6"] = flattenRouterBgpNeighborGroupCapabilityOrf6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if cur_v, ok := i["capability-graceful-restart"]; ok {
			tmp["capability_graceful_restart"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestart(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if cur_v, ok := i["capability-graceful-restart6"]; ok {
			tmp["capability_graceful_restart6"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestart6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if cur_v, ok := i["capability-graceful-restart-vpnv4"]; ok {
			tmp["capability_graceful_restart_vpnv4"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if cur_v, ok := i["capability-graceful-restart-vpnv6"]; ok {
			tmp["capability_graceful_restart_vpnv6"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if cur_v, ok := i["capability-graceful-restart-evpn"]; ok {
			tmp["capability_graceful_restart_evpn"] = flattenRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if cur_v, ok := i["capability-route-refresh"]; ok {
			tmp["capability_route_refresh"] = flattenRouterBgpNeighborGroupCapabilityRouteRefresh(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if cur_v, ok := i["capability-default-originate"]; ok {
			tmp["capability_default_originate"] = flattenRouterBgpNeighborGroupCapabilityDefaultOriginate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if cur_v, ok := i["capability-default-originate6"]; ok {
			tmp["capability_default_originate6"] = flattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if cur_v, ok := i["dont-capability-negotiate"]; ok {
			tmp["dont_capability_negotiate"] = flattenRouterBgpNeighborGroupDontCapabilityNegotiate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if cur_v, ok := i["ebgp-enforce-multihop"]; ok {
			tmp["ebgp_enforce_multihop"] = flattenRouterBgpNeighborGroupEbgpEnforceMultihop(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if cur_v, ok := i["link-down-failover"]; ok {
			tmp["link_down_failover"] = flattenRouterBgpNeighborGroupLinkDownFailover(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if cur_v, ok := i["stale-route"]; ok {
			tmp["stale_route"] = flattenRouterBgpNeighborGroupStaleRoute(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if cur_v, ok := i["next-hop-self"]; ok {
			tmp["next_hop_self"] = flattenRouterBgpNeighborGroupNextHopSelf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if cur_v, ok := i["next-hop-self6"]; ok {
			tmp["next_hop_self6"] = flattenRouterBgpNeighborGroupNextHopSelf6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if cur_v, ok := i["next-hop-self-rr"]; ok {
			tmp["next_hop_self_rr"] = flattenRouterBgpNeighborGroupNextHopSelfRr(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if cur_v, ok := i["next-hop-self-rr6"]; ok {
			tmp["next_hop_self_rr6"] = flattenRouterBgpNeighborGroupNextHopSelfRr6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if cur_v, ok := i["next-hop-self-vpnv4"]; ok {
			tmp["next_hop_self_vpnv4"] = flattenRouterBgpNeighborGroupNextHopSelfVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if cur_v, ok := i["next-hop-self-vpnv6"]; ok {
			tmp["next_hop_self_vpnv6"] = flattenRouterBgpNeighborGroupNextHopSelfVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if cur_v, ok := i["override-capability"]; ok {
			tmp["override_capability"] = flattenRouterBgpNeighborGroupOverrideCapability(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if cur_v, ok := i["passive"]; ok {
			tmp["passive"] = flattenRouterBgpNeighborGroupPassive(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if cur_v, ok := i["remove-private-as"]; ok {
			tmp["remove_private_as"] = flattenRouterBgpNeighborGroupRemovePrivateAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if cur_v, ok := i["remove-private-as6"]; ok {
			tmp["remove_private_as6"] = flattenRouterBgpNeighborGroupRemovePrivateAs6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if cur_v, ok := i["remove-private-as-vpnv4"]; ok {
			tmp["remove_private_as_vpnv4"] = flattenRouterBgpNeighborGroupRemovePrivateAsVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if cur_v, ok := i["remove-private-as-vpnv6"]; ok {
			tmp["remove_private_as_vpnv6"] = flattenRouterBgpNeighborGroupRemovePrivateAsVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if cur_v, ok := i["remove-private-as-evpn"]; ok {
			tmp["remove_private_as_evpn"] = flattenRouterBgpNeighborGroupRemovePrivateAsEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if cur_v, ok := i["route-reflector-client"]; ok {
			tmp["route_reflector_client"] = flattenRouterBgpNeighborGroupRouteReflectorClient(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if cur_v, ok := i["route-reflector-client6"]; ok {
			tmp["route_reflector_client6"] = flattenRouterBgpNeighborGroupRouteReflectorClient6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if cur_v, ok := i["route-reflector-client-vpnv4"]; ok {
			tmp["route_reflector_client_vpnv4"] = flattenRouterBgpNeighborGroupRouteReflectorClientVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if cur_v, ok := i["route-reflector-client-vpnv6"]; ok {
			tmp["route_reflector_client_vpnv6"] = flattenRouterBgpNeighborGroupRouteReflectorClientVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if cur_v, ok := i["route-reflector-client-evpn"]; ok {
			tmp["route_reflector_client_evpn"] = flattenRouterBgpNeighborGroupRouteReflectorClientEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if cur_v, ok := i["route-server-client"]; ok {
			tmp["route_server_client"] = flattenRouterBgpNeighborGroupRouteServerClient(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if cur_v, ok := i["route-server-client6"]; ok {
			tmp["route_server_client6"] = flattenRouterBgpNeighborGroupRouteServerClient6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if cur_v, ok := i["route-server-client-vpnv4"]; ok {
			tmp["route_server_client_vpnv4"] = flattenRouterBgpNeighborGroupRouteServerClientVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if cur_v, ok := i["route-server-client-vpnv6"]; ok {
			tmp["route_server_client_vpnv6"] = flattenRouterBgpNeighborGroupRouteServerClientVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if cur_v, ok := i["route-server-client-evpn"]; ok {
			tmp["route_server_client_evpn"] = flattenRouterBgpNeighborGroupRouteServerClientEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if cur_v, ok := i["shutdown"]; ok {
			tmp["shutdown"] = flattenRouterBgpNeighborGroupShutdown(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if cur_v, ok := i["soft-reconfiguration"]; ok {
			tmp["soft_reconfiguration"] = flattenRouterBgpNeighborGroupSoftReconfiguration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if cur_v, ok := i["soft-reconfiguration6"]; ok {
			tmp["soft_reconfiguration6"] = flattenRouterBgpNeighborGroupSoftReconfiguration6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if cur_v, ok := i["soft-reconfiguration-vpnv4"]; ok {
			tmp["soft_reconfiguration_vpnv4"] = flattenRouterBgpNeighborGroupSoftReconfigurationVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if cur_v, ok := i["soft-reconfiguration-vpnv6"]; ok {
			tmp["soft_reconfiguration_vpnv6"] = flattenRouterBgpNeighborGroupSoftReconfigurationVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if cur_v, ok := i["soft-reconfiguration-evpn"]; ok {
			tmp["soft_reconfiguration_evpn"] = flattenRouterBgpNeighborGroupSoftReconfigurationEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if cur_v, ok := i["as-override"]; ok {
			tmp["as_override"] = flattenRouterBgpNeighborGroupAsOverride(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if cur_v, ok := i["as-override6"]; ok {
			tmp["as_override6"] = flattenRouterBgpNeighborGroupAsOverride6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if cur_v, ok := i["strict-capability-match"]; ok {
			tmp["strict_capability_match"] = flattenRouterBgpNeighborGroupStrictCapabilityMatch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if cur_v, ok := i["default-originate-routemap"]; ok {
			tmp["default_originate_routemap"] = flattenRouterBgpNeighborGroupDefaultOriginateRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if cur_v, ok := i["default-originate-routemap6"]; ok {
			tmp["default_originate_routemap6"] = flattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenRouterBgpNeighborGroupDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if cur_v, ok := i["distribute-list-in"]; ok {
			tmp["distribute_list_in"] = flattenRouterBgpNeighborGroupDistributeListIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if cur_v, ok := i["distribute-list-in6"]; ok {
			tmp["distribute_list_in6"] = flattenRouterBgpNeighborGroupDistributeListIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if cur_v, ok := i["distribute-list-in-vpnv4"]; ok {
			tmp["distribute_list_in_vpnv4"] = flattenRouterBgpNeighborGroupDistributeListInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if cur_v, ok := i["distribute-list-in-vpnv6"]; ok {
			tmp["distribute_list_in_vpnv6"] = flattenRouterBgpNeighborGroupDistributeListInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if cur_v, ok := i["distribute-list-out"]; ok {
			tmp["distribute_list_out"] = flattenRouterBgpNeighborGroupDistributeListOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if cur_v, ok := i["distribute-list-out6"]; ok {
			tmp["distribute_list_out6"] = flattenRouterBgpNeighborGroupDistributeListOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if cur_v, ok := i["distribute-list-out-vpnv4"]; ok {
			tmp["distribute_list_out_vpnv4"] = flattenRouterBgpNeighborGroupDistributeListOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if cur_v, ok := i["distribute-list-out-vpnv6"]; ok {
			tmp["distribute_list_out_vpnv6"] = flattenRouterBgpNeighborGroupDistributeListOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if cur_v, ok := i["ebgp-multihop-ttl"]; ok {
			tmp["ebgp_multihop_ttl"] = flattenRouterBgpNeighborGroupEbgpMultihopTtl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if cur_v, ok := i["filter-list-in"]; ok {
			tmp["filter_list_in"] = flattenRouterBgpNeighborGroupFilterListIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if cur_v, ok := i["filter-list-in6"]; ok {
			tmp["filter_list_in6"] = flattenRouterBgpNeighborGroupFilterListIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if cur_v, ok := i["filter-list-in-vpnv4"]; ok {
			tmp["filter_list_in_vpnv4"] = flattenRouterBgpNeighborGroupFilterListInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if cur_v, ok := i["filter-list-in-vpnv6"]; ok {
			tmp["filter_list_in_vpnv6"] = flattenRouterBgpNeighborGroupFilterListInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if cur_v, ok := i["filter-list-out"]; ok {
			tmp["filter_list_out"] = flattenRouterBgpNeighborGroupFilterListOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if cur_v, ok := i["filter-list-out6"]; ok {
			tmp["filter_list_out6"] = flattenRouterBgpNeighborGroupFilterListOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if cur_v, ok := i["filter-list-out-vpnv4"]; ok {
			tmp["filter_list_out_vpnv4"] = flattenRouterBgpNeighborGroupFilterListOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if cur_v, ok := i["filter-list-out-vpnv6"]; ok {
			tmp["filter_list_out_vpnv6"] = flattenRouterBgpNeighborGroupFilterListOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBgpNeighborGroupInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if cur_v, ok := i["maximum-prefix"]; ok {
			tmp["maximum_prefix"] = flattenRouterBgpNeighborGroupMaximumPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if cur_v, ok := i["maximum-prefix6"]; ok {
			tmp["maximum_prefix6"] = flattenRouterBgpNeighborGroupMaximumPrefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if cur_v, ok := i["maximum-prefix-vpnv4"]; ok {
			tmp["maximum_prefix_vpnv4"] = flattenRouterBgpNeighborGroupMaximumPrefixVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if cur_v, ok := i["maximum-prefix-vpnv6"]; ok {
			tmp["maximum_prefix_vpnv6"] = flattenRouterBgpNeighborGroupMaximumPrefixVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if cur_v, ok := i["maximum-prefix-evpn"]; ok {
			tmp["maximum_prefix_evpn"] = flattenRouterBgpNeighborGroupMaximumPrefixEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if cur_v, ok := i["maximum-prefix-threshold"]; ok {
			tmp["maximum_prefix_threshold"] = flattenRouterBgpNeighborGroupMaximumPrefixThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if cur_v, ok := i["maximum-prefix-threshold6"]; ok {
			tmp["maximum_prefix_threshold6"] = flattenRouterBgpNeighborGroupMaximumPrefixThreshold6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if cur_v, ok := i["maximum-prefix-threshold-vpnv4"]; ok {
			tmp["maximum_prefix_threshold_vpnv4"] = flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if cur_v, ok := i["maximum-prefix-threshold-vpnv6"]; ok {
			tmp["maximum_prefix_threshold_vpnv6"] = flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if cur_v, ok := i["maximum-prefix-threshold-evpn"]; ok {
			tmp["maximum_prefix_threshold_evpn"] = flattenRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if cur_v, ok := i["maximum-prefix-warning-only"]; ok {
			tmp["maximum_prefix_warning_only"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if cur_v, ok := i["maximum-prefix-warning-only6"]; ok {
			tmp["maximum_prefix_warning_only6"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if cur_v, ok := i["maximum-prefix-warning-only-vpnv4"]; ok {
			tmp["maximum_prefix_warning_only_vpnv4"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if cur_v, ok := i["maximum-prefix-warning-only-vpnv6"]; ok {
			tmp["maximum_prefix_warning_only_vpnv6"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if cur_v, ok := i["maximum-prefix-warning-only-evpn"]; ok {
			tmp["maximum_prefix_warning_only_evpn"] = flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if cur_v, ok := i["prefix-list-in"]; ok {
			tmp["prefix_list_in"] = flattenRouterBgpNeighborGroupPrefixListIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if cur_v, ok := i["prefix-list-in6"]; ok {
			tmp["prefix_list_in6"] = flattenRouterBgpNeighborGroupPrefixListIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if cur_v, ok := i["prefix-list-in-vpnv4"]; ok {
			tmp["prefix_list_in_vpnv4"] = flattenRouterBgpNeighborGroupPrefixListInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if cur_v, ok := i["prefix-list-in-vpnv6"]; ok {
			tmp["prefix_list_in_vpnv6"] = flattenRouterBgpNeighborGroupPrefixListInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if cur_v, ok := i["prefix-list-out"]; ok {
			tmp["prefix_list_out"] = flattenRouterBgpNeighborGroupPrefixListOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if cur_v, ok := i["prefix-list-out6"]; ok {
			tmp["prefix_list_out6"] = flattenRouterBgpNeighborGroupPrefixListOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if cur_v, ok := i["prefix-list-out-vpnv4"]; ok {
			tmp["prefix_list_out_vpnv4"] = flattenRouterBgpNeighborGroupPrefixListOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if cur_v, ok := i["prefix-list-out-vpnv6"]; ok {
			tmp["prefix_list_out_vpnv6"] = flattenRouterBgpNeighborGroupPrefixListOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if cur_v, ok := i["remote-as"]; ok {
			tmp["remote_as"] = flattenRouterBgpNeighborGroupRemoteAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as_filter"
		if cur_v, ok := i["remote-as-filter"]; ok {
			tmp["remote_as_filter"] = flattenRouterBgpNeighborGroupRemoteAsFilter(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if cur_v, ok := i["local-as"]; ok {
			tmp["local_as"] = flattenRouterBgpNeighborGroupLocalAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if cur_v, ok := i["local-as-no-prepend"]; ok {
			tmp["local_as_no_prepend"] = flattenRouterBgpNeighborGroupLocalAsNoPrepend(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if cur_v, ok := i["local-as-replace-as"]; ok {
			tmp["local_as_replace_as"] = flattenRouterBgpNeighborGroupLocalAsReplaceAs(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if cur_v, ok := i["retain-stale-time"]; ok {
			tmp["retain_stale_time"] = flattenRouterBgpNeighborGroupRetainStaleTime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if cur_v, ok := i["route-map-in"]; ok {
			tmp["route_map_in"] = flattenRouterBgpNeighborGroupRouteMapIn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if cur_v, ok := i["route-map-in6"]; ok {
			tmp["route_map_in6"] = flattenRouterBgpNeighborGroupRouteMapIn6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if cur_v, ok := i["route-map-in-vpnv4"]; ok {
			tmp["route_map_in_vpnv4"] = flattenRouterBgpNeighborGroupRouteMapInVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if cur_v, ok := i["route-map-in-vpnv6"]; ok {
			tmp["route_map_in_vpnv6"] = flattenRouterBgpNeighborGroupRouteMapInVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if cur_v, ok := i["route-map-in-evpn"]; ok {
			tmp["route_map_in_evpn"] = flattenRouterBgpNeighborGroupRouteMapInEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if cur_v, ok := i["route-map-out"]; ok {
			tmp["route_map_out"] = flattenRouterBgpNeighborGroupRouteMapOut(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if cur_v, ok := i["route-map-out-preferable"]; ok {
			tmp["route_map_out_preferable"] = flattenRouterBgpNeighborGroupRouteMapOutPreferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if cur_v, ok := i["route-map-out6"]; ok {
			tmp["route_map_out6"] = flattenRouterBgpNeighborGroupRouteMapOut6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if cur_v, ok := i["route-map-out6-preferable"]; ok {
			tmp["route_map_out6_preferable"] = flattenRouterBgpNeighborGroupRouteMapOut6Preferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if cur_v, ok := i["route-map-out-vpnv4"]; ok {
			tmp["route_map_out_vpnv4"] = flattenRouterBgpNeighborGroupRouteMapOutVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if cur_v, ok := i["route-map-out-vpnv6"]; ok {
			tmp["route_map_out_vpnv6"] = flattenRouterBgpNeighborGroupRouteMapOutVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if cur_v, ok := i["route-map-out-vpnv4-preferable"]; ok {
			tmp["route_map_out_vpnv4_preferable"] = flattenRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if cur_v, ok := i["route-map-out-vpnv6-preferable"]; ok {
			tmp["route_map_out_vpnv6_preferable"] = flattenRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if cur_v, ok := i["route-map-out-evpn"]; ok {
			tmp["route_map_out_evpn"] = flattenRouterBgpNeighborGroupRouteMapOutEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if cur_v, ok := i["send-community"]; ok {
			tmp["send_community"] = flattenRouterBgpNeighborGroupSendCommunity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if cur_v, ok := i["send-community6"]; ok {
			tmp["send_community6"] = flattenRouterBgpNeighborGroupSendCommunity6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if cur_v, ok := i["send-community-vpnv4"]; ok {
			tmp["send_community_vpnv4"] = flattenRouterBgpNeighborGroupSendCommunityVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if cur_v, ok := i["send-community-vpnv6"]; ok {
			tmp["send_community_vpnv6"] = flattenRouterBgpNeighborGroupSendCommunityVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if cur_v, ok := i["send-community-evpn"]; ok {
			tmp["send_community_evpn"] = flattenRouterBgpNeighborGroupSendCommunityEvpn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if cur_v, ok := i["keep-alive-timer"]; ok {
			tmp["keep_alive_timer"] = flattenRouterBgpNeighborGroupKeepAliveTimer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if cur_v, ok := i["holdtime-timer"]; ok {
			tmp["holdtime_timer"] = flattenRouterBgpNeighborGroupHoldtimeTimer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if cur_v, ok := i["connect-timer"]; ok {
			tmp["connect_timer"] = flattenRouterBgpNeighborGroupConnectTimer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if cur_v, ok := i["unsuppress-map"]; ok {
			tmp["unsuppress_map"] = flattenRouterBgpNeighborGroupUnsuppressMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if cur_v, ok := i["unsuppress-map6"]; ok {
			tmp["unsuppress_map6"] = flattenRouterBgpNeighborGroupUnsuppressMap6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if cur_v, ok := i["update-source"]; ok {
			tmp["update_source"] = flattenRouterBgpNeighborGroupUpdateSource(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if cur_v, ok := i["weight"]; ok {
			tmp["weight"] = flattenRouterBgpNeighborGroupWeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if cur_v, ok := i["restart-time"]; ok {
			tmp["restart_time"] = flattenRouterBgpNeighborGroupRestartTime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if cur_v, ok := i["additional-path"]; ok {
			tmp["additional_path"] = flattenRouterBgpNeighborGroupAdditionalPath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if cur_v, ok := i["additional-path6"]; ok {
			tmp["additional_path6"] = flattenRouterBgpNeighborGroupAdditionalPath6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if cur_v, ok := i["additional-path-vpnv4"]; ok {
			tmp["additional_path_vpnv4"] = flattenRouterBgpNeighborGroupAdditionalPathVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if cur_v, ok := i["additional-path-vpnv6"]; ok {
			tmp["additional_path_vpnv6"] = flattenRouterBgpNeighborGroupAdditionalPathVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if cur_v, ok := i["adv-additional-path"]; ok {
			tmp["adv_additional_path"] = flattenRouterBgpNeighborGroupAdvAdditionalPath(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if cur_v, ok := i["adv-additional-path6"]; ok {
			tmp["adv_additional_path6"] = flattenRouterBgpNeighborGroupAdvAdditionalPath6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if cur_v, ok := i["adv-additional-path-vpnv4"]; ok {
			tmp["adv_additional_path_vpnv4"] = flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv4(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if cur_v, ok := i["adv-additional-path-vpnv6"]; ok {
			tmp["adv_additional_path_vpnv6"] = flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if cur_v, ok := i["auth-options"]; ok {
			tmp["auth_options"] = flattenRouterBgpNeighborGroupAuthOptions(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterBgpNeighborGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAllowasInEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasInEnableEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAllowasIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAllowasIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAllowasInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAllowasInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAllowasInEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchangedVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAttributeUnchangedVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupActivateEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLinkDownFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStaleRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelf6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupNextHopSelfVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupOverrideCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemovePrivateAsEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteReflectorClientEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClient6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteServerClientEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupShutdown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSoftReconfigurationEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAsOverride6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupDistributeListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupFilterListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupFilterListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupPrefixListOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRemoteAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupRemoteAsFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRetainStaleTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapIn6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapInVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapInVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapInEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOut6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupRouteMapOutEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunity6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupSendCommunityEvpn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupConnectTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupUnsuppressMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupUpdateSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupRestartTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAdvAdditionalPathVpnv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborGroupAuthOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpNeighborRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if cur_v, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterBgpNeighborRangePrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if cur_v, ok := i["max-neighbor-num"]; ok {
			tmp["max_neighbor_num"] = flattenRouterBgpNeighborRangeMaxNeighborNum(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if cur_v, ok := i["neighbor-group"]; ok {
			tmp["neighbor_group"] = flattenRouterBgpNeighborRangeNeighborGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNeighborRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborRangePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBgpNeighborRangeMaxNeighborNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborRangeNeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpNeighborRange6Id(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if cur_v, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterBgpNeighborRange6Prefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if cur_v, ok := i["max-neighbor-num"]; ok {
			tmp["max_neighbor_num"] = flattenRouterBgpNeighborRange6MaxNeighborNum(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if cur_v, ok := i["neighbor-group"]; ok {
			tmp["neighbor_group"] = flattenRouterBgpNeighborRange6NeighborGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNeighborRange6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborRange6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNeighborRange6MaxNeighborNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNeighborRange6NeighborGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpNetworkId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if cur_v, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterBgpNetworkPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if cur_v, ok := i["network-import-check"]; ok {
			tmp["network_import_check"] = flattenRouterBgpNetworkNetworkImportCheck(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if cur_v, ok := i["backdoor"]; ok {
			tmp["backdoor"] = flattenRouterBgpNetworkBackdoor(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpNetworkRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_name"
		if cur_v, ok := i["prefix-name"]; ok {
			tmp["prefix_name"] = flattenRouterBgpNetworkPrefixName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBgpNetworkNetworkImportCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkBackdoor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetworkPrefixName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpNetwork6Id(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if cur_v, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterBgpNetwork6Prefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if cur_v, ok := i["network-import-check"]; ok {
			tmp["network_import_check"] = flattenRouterBgpNetwork6NetworkImportCheck(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if cur_v, ok := i["backdoor"]; ok {
			tmp["backdoor"] = flattenRouterBgpNetwork6Backdoor(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpNetwork6RouteMap(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpNetwork6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpNetwork6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6NetworkImportCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6Backdoor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpNetwork6RouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	conf_map := make(map[string]bool)
	for _, ele := range d.Get(pre).([]interface{}) {
		if ele_map, ok := ele.(map[string]interface{}); ok {
			conf_map[ele_map["name"].(string)] = true
		}
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			if !conf_map[cur_v.(string)] {
				continue
			}
			tmp["name"] = flattenRouterBgpRedistributeName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenRouterBgpRedistributeStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpRedistributeRouteMap(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterBgpRedistributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistributeRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	conf_map := make(map[string]bool)
	for _, ele := range d.Get(pre).([]interface{}) {
		if ele_map, ok := ele.(map[string]interface{}); ok {
			conf_map[ele_map["name"].(string)] = true
		}
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			if !conf_map[cur_v.(string)] {
				continue
			}
			tmp["name"] = flattenRouterBgpRedistribute6Name(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenRouterBgpRedistribute6Status(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpRedistribute6RouteMap(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterBgpRedistribute6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpRedistribute6RouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdminDistance(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterBgpAdminDistanceId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if cur_v, ok := i["neighbour-prefix"]; ok {
			tmp["neighbour_prefix"] = flattenRouterBgpAdminDistanceNeighbourPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if cur_v, ok := i["route-list"]; ok {
			tmp["route_list"] = flattenRouterBgpAdminDistanceRouteList(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if cur_v, ok := i["distance"]; ok {
			tmp["distance"] = flattenRouterBgpAdminDistanceDistance(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterBgpAdminDistanceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpAdminDistanceNeighbourPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterBgpAdminDistanceRouteList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpAdminDistanceDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterBgpVrf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrfVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if cur_v, ok := i["role"]; ok {
			tmp["role"] = flattenRouterBgpVrfRole(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if cur_v, ok := i["rd"]; ok {
			tmp["rd"] = flattenRouterBgpVrfRd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if cur_v, ok := i["export-rt"]; ok {
			tmp["export_rt"] = flattenRouterBgpVrfExportRt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if cur_v, ok := i["import-rt"]; ok {
			tmp["import_rt"] = flattenRouterBgpVrfImportRt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if cur_v, ok := i["import-route-map"]; ok {
			tmp["import_route_map"] = flattenRouterBgpVrfImportRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if cur_v, ok := i["leak-target"]; ok {
			tmp["leak_target"] = flattenRouterBgpVrfLeakTargetU(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfRd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfExportRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
		if cur_v, ok := i["route-target"]; ok {
			tmp["route_target"] = flattenRouterBgpVrfExportRtRouteTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "route_target", d)
	return result
}

func flattenRouterBgpVrfExportRtRouteTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfImportRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
		if cur_v, ok := i["route-target"]; ok {
			tmp["route_target"] = flattenRouterBgpVrfImportRtRouteTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "route_target", d)
	return result
}

func flattenRouterBgpVrfImportRtRouteTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfImportRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetU(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrfLeakTargetVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpVrfLeakTargetRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBgpVrfLeakTargetInterface(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeakTargetUVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetURouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetUInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrf6Vrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if cur_v, ok := i["role"]; ok {
			tmp["role"] = flattenRouterBgpVrf6Role(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if cur_v, ok := i["rd"]; ok {
			tmp["rd"] = flattenRouterBgpVrf6Rd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if cur_v, ok := i["export-rt"]; ok {
			tmp["export_rt"] = flattenRouterBgpVrf6ExportRt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if cur_v, ok := i["import-rt"]; ok {
			tmp["import_rt"] = flattenRouterBgpVrf6ImportRt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if cur_v, ok := i["import-route-map"]; ok {
			tmp["import_route_map"] = flattenRouterBgpVrf6ImportRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if cur_v, ok := i["leak-target"]; ok {
			tmp["leak_target"] = flattenRouterBgpVrf6LeakTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrf6Vrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6Role(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6Rd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6ExportRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
		if cur_v, ok := i["route-target"]; ok {
			tmp["route_target"] = flattenRouterBgpVrf6ExportRtRouteTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "route_target", d)
	return result
}

func flattenRouterBgpVrf6ExportRtRouteTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6ImportRt(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_target"
		if cur_v, ok := i["route-target"]; ok {
			tmp["route_target"] = flattenRouterBgpVrf6ImportRtRouteTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "route_target", d)
	return result
}

func flattenRouterBgpVrf6ImportRtRouteTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6ImportRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6LeakTarget(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrf6LeakTargetVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpVrf6LeakTargetRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBgpVrf6LeakTargetInterface(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrf6LeakTargetVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6LeakTargetRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrf6LeakTargetInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrfLeakVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if cur_v, ok := i["target"]; ok {
			tmp["target"] = flattenRouterBgpVrfLeakTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeakVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTarget(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrfLeakTargetVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpVrfLeakTargetRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBgpVrfLeakTargetInterface(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeakTargetVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeakTargetInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrfLeak6Vrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if cur_v, ok := i["target"]; ok {
			tmp["target"] = flattenRouterBgpVrfLeak6Target(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeak6Vrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6Target(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if cur_v, ok := i["vrf"]; ok {
			tmp["vrf"] = flattenRouterBgpVrfLeak6TargetVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if cur_v, ok := i["route-map"]; ok {
			tmp["route_map"] = flattenRouterBgpVrfLeak6TargetRouteMap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterBgpVrfLeak6TargetInterface(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterBgpVrfLeak6TargetVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6TargetRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterBgpVrfLeak6TargetInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if _, ok := o["as"].(string); ok {
		if err = d.Set("as_string", flattenRouterBgpAsString(o["as"], d, "as_string", sv)); err != nil {
			if !fortiAPIPatch(o["as"]) {
				return fmt.Errorf("Error reading as_string: %v", err)
			}
		}
	}

	if _, ok := o["as"].(float64); ok {
		if err = d.Set("as", flattenRouterBgpAs(o["as"], d, "as", sv)); err != nil {
			if !fortiAPIPatch(o["as"]) {
				return fmt.Errorf("Error reading as: %v", err)
			}
		}
	}

	if err = d.Set("router_id", flattenRouterBgpRouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("keepalive_timer", flattenRouterBgpKeepaliveTimer(o["keepalive-timer"], d, "keepalive_timer", sv)); err != nil {
		if !fortiAPIPatch(o["keepalive-timer"]) {
			return fmt.Errorf("Error reading keepalive_timer: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", flattenRouterBgpHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer", sv)); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("always_compare_med", flattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med", sv)); err != nil {
		if !fortiAPIPatch(o["always-compare-med"]) {
			return fmt.Errorf("Error reading always_compare_med: %v", err)
		}
	}

	if err = d.Set("bestpath_as_path_ignore", flattenRouterBgpBestpathAsPathIgnore(o["bestpath-as-path-ignore"], d, "bestpath_as_path_ignore", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-as-path-ignore"]) {
			return fmt.Errorf("Error reading bestpath_as_path_ignore: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_confed_aspath", flattenRouterBgpBestpathCmpConfedAspath(o["bestpath-cmp-confed-aspath"], d, "bestpath_cmp_confed_aspath", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-confed-aspath"]) {
			return fmt.Errorf("Error reading bestpath_cmp_confed_aspath: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", flattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-routerid"]) {
			return fmt.Errorf("Error reading bestpath_cmp_routerid: %v", err)
		}
	}

	if err = d.Set("bestpath_med_confed", flattenRouterBgpBestpathMedConfed(o["bestpath-med-confed"], d, "bestpath_med_confed", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-med-confed"]) {
			return fmt.Errorf("Error reading bestpath_med_confed: %v", err)
		}
	}

	if err = d.Set("bestpath_med_missing_as_worst", flattenRouterBgpBestpathMedMissingAsWorst(o["bestpath-med-missing-as-worst"], d, "bestpath_med_missing_as_worst", sv)); err != nil {
		if !fortiAPIPatch(o["bestpath-med-missing-as-worst"]) {
			return fmt.Errorf("Error reading bestpath_med_missing_as_worst: %v", err)
		}
	}

	if err = d.Set("client_to_client_reflection", flattenRouterBgpClientToClientReflection(o["client-to-client-reflection"], d, "client_to_client_reflection", sv)); err != nil {
		if !fortiAPIPatch(o["client-to-client-reflection"]) {
			return fmt.Errorf("Error reading client_to_client_reflection: %v", err)
		}
	}

	if err = d.Set("dampening", flattenRouterBgpDampening(o["dampening"], d, "dampening", sv)); err != nil {
		if !fortiAPIPatch(o["dampening"]) {
			return fmt.Errorf("Error reading dampening: %v", err)
		}
	}

	if err = d.Set("deterministic_med", flattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med", sv)); err != nil {
		if !fortiAPIPatch(o["deterministic-med"]) {
			return fmt.Errorf("Error reading deterministic_med: %v", err)
		}
	}

	if err = d.Set("ebgp_multipath", flattenRouterBgpEbgpMultipath(o["ebgp-multipath"], d, "ebgp_multipath", sv)); err != nil {
		if !fortiAPIPatch(o["ebgp-multipath"]) {
			return fmt.Errorf("Error reading ebgp_multipath: %v", err)
		}
	}

	if err = d.Set("ibgp_multipath", flattenRouterBgpIbgpMultipath(o["ibgp-multipath"], d, "ibgp_multipath", sv)); err != nil {
		if !fortiAPIPatch(o["ibgp-multipath"]) {
			return fmt.Errorf("Error reading ibgp_multipath: %v", err)
		}
	}

	if err = d.Set("enforce_first_as", flattenRouterBgpEnforceFirstAs(o["enforce-first-as"], d, "enforce_first_as", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-first-as"]) {
			return fmt.Errorf("Error reading enforce_first_as: %v", err)
		}
	}

	if err = d.Set("fast_external_failover", flattenRouterBgpFastExternalFailover(o["fast-external-failover"], d, "fast_external_failover", sv)); err != nil {
		if !fortiAPIPatch(o["fast-external-failover"]) {
			return fmt.Errorf("Error reading fast_external_failover: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterBgpLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes", sv)); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("network_import_check", flattenRouterBgpNetworkImportCheck(o["network-import-check"], d, "network_import_check", sv)); err != nil {
		if !fortiAPIPatch(o["network-import-check"]) {
			return fmt.Errorf("Error reading network_import_check: %v", err)
		}
	}

	if err = d.Set("ignore_optional_capability", flattenRouterBgpIgnoreOptionalCapability(o["ignore-optional-capability"], d, "ignore_optional_capability", sv)); err != nil {
		if !fortiAPIPatch(o["ignore-optional-capability"]) {
			return fmt.Errorf("Error reading ignore_optional_capability: %v", err)
		}
	}

	if err = d.Set("additional_path", flattenRouterBgpAdditionalPath(o["additional-path"], d, "additional_path", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("Error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path6", flattenRouterBgpAdditionalPath6(o["additional-path6"], d, "additional_path6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("Error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv4", flattenRouterBgpAdditionalPathVpnv4(o["additional-path-vpnv4"], d, "additional_path_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-vpnv4"]) {
			return fmt.Errorf("Error reading additional_path_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_vpnv6", flattenRouterBgpAdditionalPathVpnv6(o["additional-path-vpnv6"], d, "additional_path_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-vpnv6"]) {
			return fmt.Errorf("Error reading additional_path_vpnv6: %v", err)
		}
	}

	if err = d.Set("multipath_recursive_distance", flattenRouterBgpMultipathRecursiveDistance(o["multipath-recursive-distance"], d, "multipath_recursive_distance", sv)); err != nil {
		if !fortiAPIPatch(o["multipath-recursive-distance"]) {
			return fmt.Errorf("Error reading multipath_recursive_distance: %v", err)
		}
	}

	if err = d.Set("recursive_next_hop", flattenRouterBgpRecursiveNextHop(o["recursive-next-hop"], d, "recursive_next_hop", sv)); err != nil {
		if !fortiAPIPatch(o["recursive-next-hop"]) {
			return fmt.Errorf("Error reading recursive_next_hop: %v", err)
		}
	}

	if err = d.Set("recursive_inherit_priority", flattenRouterBgpRecursiveInheritPriority(o["recursive-inherit-priority"], d, "recursive_inherit_priority", sv)); err != nil {
		if !fortiAPIPatch(o["recursive-inherit-priority"]) {
			return fmt.Errorf("Error reading recursive_inherit_priority: %v", err)
		}
	}

	if err = d.Set("tag_resolve_mode", flattenRouterBgpTagResolveMode(o["tag-resolve-mode"], d, "tag_resolve_mode", sv)); err != nil {
		if !fortiAPIPatch(o["tag-resolve-mode"]) {
			return fmt.Errorf("Error reading tag_resolve_mode: %v", err)
		}
	}

	if err = d.Set("cluster_id", flattenRouterBgpClusterId(o["cluster-id"], d, "cluster_id", sv)); err != nil {
		if !fortiAPIPatch(o["cluster-id"]) {
			return fmt.Errorf("Error reading cluster_id: %v", err)
		}
	}

	if err = d.Set("confederation_identifier", flattenRouterBgpConfederationIdentifier(o["confederation-identifier"], d, "confederation_identifier", sv)); err != nil {
		if !fortiAPIPatch(o["confederation-identifier"]) {
			return fmt.Errorf("Error reading confederation_identifier: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("confederation_peers", flattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers", sv)); err != nil {
			if !fortiAPIPatch(o["confederation-peers"]) {
				return fmt.Errorf("Error reading confederation_peers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("confederation_peers"); ok {
			if err = d.Set("confederation_peers", flattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers", sv)); err != nil {
				if !fortiAPIPatch(o["confederation-peers"]) {
					return fmt.Errorf("Error reading confederation_peers: %v", err)
				}
			}
		}
	}

	if err = d.Set("dampening_route_map", flattenRouterBgpDampeningRouteMap(o["dampening-route-map"], d, "dampening_route_map", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-route-map"]) {
			return fmt.Errorf("Error reading dampening_route_map: %v", err)
		}
	}

	if err = d.Set("dampening_reachability_half_life", flattenRouterBgpDampeningReachabilityHalfLife(o["dampening-reachability-half-life"], d, "dampening_reachability_half_life", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-reachability-half-life"]) {
			return fmt.Errorf("Error reading dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("dampening_reuse", flattenRouterBgpDampeningReuse(o["dampening-reuse"], d, "dampening_reuse", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-reuse"]) {
			return fmt.Errorf("Error reading dampening_reuse: %v", err)
		}
	}

	if err = d.Set("dampening_suppress", flattenRouterBgpDampeningSuppress(o["dampening-suppress"], d, "dampening_suppress", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-suppress"]) {
			return fmt.Errorf("Error reading dampening_suppress: %v", err)
		}
	}

	if err = d.Set("dampening_max_suppress_time", flattenRouterBgpDampeningMaxSuppressTime(o["dampening-max-suppress-time"], d, "dampening_max_suppress_time", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-max-suppress-time"]) {
			return fmt.Errorf("Error reading dampening_max_suppress_time: %v", err)
		}
	}

	if err = d.Set("dampening_unreachability_half_life", flattenRouterBgpDampeningUnreachabilityHalfLife(o["dampening-unreachability-half-life"], d, "dampening_unreachability_half_life", sv)); err != nil {
		if !fortiAPIPatch(o["dampening-unreachability-half-life"]) {
			return fmt.Errorf("Error reading dampening_unreachability_half_life: %v", err)
		}
	}

	if err = d.Set("default_local_preference", flattenRouterBgpDefaultLocalPreference(o["default-local-preference"], d, "default_local_preference", sv)); err != nil {
		if !fortiAPIPatch(o["default-local-preference"]) {
			return fmt.Errorf("Error reading default_local_preference: %v", err)
		}
	}

	if err = d.Set("scan_time", flattenRouterBgpScanTime(o["scan-time"], d, "scan_time", sv)); err != nil {
		if !fortiAPIPatch(o["scan-time"]) {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("distance_external", flattenRouterBgpDistanceExternal(o["distance-external"], d, "distance_external", sv)); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_internal", flattenRouterBgpDistanceInternal(o["distance-internal"], d, "distance_internal", sv)); err != nil {
		if !fortiAPIPatch(o["distance-internal"]) {
			return fmt.Errorf("Error reading distance_internal: %v", err)
		}
	}

	if err = d.Set("distance_local", flattenRouterBgpDistanceLocal(o["distance-local"], d, "distance_local", sv)); err != nil {
		if !fortiAPIPatch(o["distance-local"]) {
			return fmt.Errorf("Error reading distance_local: %v", err)
		}
	}

	if err = d.Set("synchronization", flattenRouterBgpSynchronization(o["synchronization"], d, "synchronization", sv)); err != nil {
		if !fortiAPIPatch(o["synchronization"]) {
			return fmt.Errorf("Error reading synchronization: %v", err)
		}
	}

	if err = d.Set("graceful_restart", flattenRouterBgpGracefulRestart(o["graceful-restart"], d, "graceful_restart", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-restart"]) {
			return fmt.Errorf("Error reading graceful_restart: %v", err)
		}
	}

	if err = d.Set("graceful_restart_time", flattenRouterBgpGracefulRestartTime(o["graceful-restart-time"], d, "graceful_restart_time", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-restart-time"]) {
			return fmt.Errorf("Error reading graceful_restart_time: %v", err)
		}
	}

	if err = d.Set("graceful_stalepath_time", flattenRouterBgpGracefulStalepathTime(o["graceful-stalepath-time"], d, "graceful_stalepath_time", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-stalepath-time"]) {
			return fmt.Errorf("Error reading graceful_stalepath_time: %v", err)
		}
	}

	if err = d.Set("graceful_update_delay", flattenRouterBgpGracefulUpdateDelay(o["graceful-update-delay"], d, "graceful_update_delay", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-update-delay"]) {
			return fmt.Errorf("Error reading graceful_update_delay: %v", err)
		}
	}

	if err = d.Set("graceful_end_on_timer", flattenRouterBgpGracefulEndOnTimer(o["graceful-end-on-timer"], d, "graceful_end_on_timer", sv)); err != nil {
		if !fortiAPIPatch(o["graceful-end-on-timer"]) {
			return fmt.Errorf("Error reading graceful_end_on_timer: %v", err)
		}
	}

	if err = d.Set("additional_path_select", flattenRouterBgpAdditionalPathSelect(o["additional-path-select"], d, "additional_path_select", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-select"]) {
			return fmt.Errorf("Error reading additional_path_select: %v", err)
		}
	}

	if err = d.Set("additional_path_select6", flattenRouterBgpAdditionalPathSelect6(o["additional-path-select6"], d, "additional_path_select6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-select6"]) {
			return fmt.Errorf("Error reading additional_path_select6: %v", err)
		}
	}

	if err = d.Set("additional_path_select_vpnv4", flattenRouterBgpAdditionalPathSelectVpnv4(o["additional-path-select-vpnv4"], d, "additional_path_select_vpnv4", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-select-vpnv4"]) {
			return fmt.Errorf("Error reading additional_path_select_vpnv4: %v", err)
		}
	}

	if err = d.Set("additional_path_select_vpnv6", flattenRouterBgpAdditionalPathSelectVpnv6(o["additional-path-select-vpnv6"], d, "additional_path_select_vpnv6", sv)); err != nil {
		if !fortiAPIPatch(o["additional-path-select-vpnv6"]) {
			return fmt.Errorf("Error reading additional_path_select_vpnv6: %v", err)
		}
	}

	if err = d.Set("cross_family_conditional_adv", flattenRouterBgpCrossFamilyConditionalAdv(o["cross-family-conditional-adv"], d, "cross_family_conditional_adv", sv)); err != nil {
		if !fortiAPIPatch(o["cross-family-conditional-adv"]) {
			return fmt.Errorf("Error reading cross_family_conditional_adv: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address", sv)); err != nil {
			if !fortiAPIPatch(o["aggregate-address"]) {
				return fmt.Errorf("Error reading aggregate_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address"); ok {
			if err = d.Set("aggregate_address", flattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address", sv)); err != nil {
				if !fortiAPIPatch(o["aggregate-address"]) {
					return fmt.Errorf("Error reading aggregate_address: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6", sv)); err != nil {
			if !fortiAPIPatch(o["aggregate-address6"]) {
				return fmt.Errorf("Error reading aggregate_address6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregate_address6"); ok {
			if err = d.Set("aggregate_address6", flattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6", sv)); err != nil {
				if !fortiAPIPatch(o["aggregate-address6"]) {
					return fmt.Errorf("Error reading aggregate_address6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("neighbor", flattenRouterBgpNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterBgpNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor-group"]) {
				return fmt.Errorf("Error reading neighbor_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_group"); ok {
			if err = d.Set("neighbor_group", flattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor-group"]) {
					return fmt.Errorf("Error reading neighbor_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor-range"]) {
				return fmt.Errorf("Error reading neighbor_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_range"); ok {
			if err = d.Set("neighbor_range", flattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor-range"]) {
					return fmt.Errorf("Error reading neighbor_range: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor-range6"]) {
				return fmt.Errorf("Error reading neighbor_range6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor_range6"); ok {
			if err = d.Set("neighbor_range6", flattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor-range6"]) {
					return fmt.Errorf("Error reading neighbor_range6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("network", flattenRouterBgpNetwork(o["network"], d, "network", sv)); err != nil {
			if !fortiAPIPatch(o["network"]) {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterBgpNetwork(o["network"], d, "network", sv)); err != nil {
				if !fortiAPIPatch(o["network"]) {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("network6", flattenRouterBgpNetwork6(o["network6"], d, "network6", sv)); err != nil {
			if !fortiAPIPatch(o["network6"]) {
				return fmt.Errorf("Error reading network6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network6"); ok {
			if err = d.Set("network6", flattenRouterBgpNetwork6(o["network6"], d, "network6", sv)); err != nil {
				if !fortiAPIPatch(o["network6"]) {
					return fmt.Errorf("Error reading network6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("redistribute", flattenRouterBgpRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterBgpRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute"]) {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute6"]) {
				return fmt.Errorf("Error reading redistribute6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute6"); ok {
			if err = d.Set("redistribute6", flattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute6"]) {
					return fmt.Errorf("Error reading redistribute6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance", sv)); err != nil {
			if !fortiAPIPatch(o["admin-distance"]) {
				return fmt.Errorf("Error reading admin_distance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("admin_distance"); ok {
			if err = d.Set("admin_distance", flattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance", sv)); err != nil {
				if !fortiAPIPatch(o["admin-distance"]) {
					return fmt.Errorf("Error reading admin_distance: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("vrf", flattenRouterBgpVrf(o["vrf"], d, "vrf", sv)); err != nil {
			if !fortiAPIPatch(o["vrf"]) {
				return fmt.Errorf("Error reading vrf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf"); ok {
			if err = d.Set("vrf", flattenRouterBgpVrf(o["vrf"], d, "vrf", sv)); err != nil {
				if !fortiAPIPatch(o["vrf"]) {
					return fmt.Errorf("Error reading vrf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("vrf6", flattenRouterBgpVrf6(o["vrf6"], d, "vrf6", sv)); err != nil {
			if !fortiAPIPatch(o["vrf6"]) {
				return fmt.Errorf("Error reading vrf6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf6"); ok {
			if err = d.Set("vrf6", flattenRouterBgpVrf6(o["vrf6"], d, "vrf6", sv)); err != nil {
				if !fortiAPIPatch(o["vrf6"]) {
					return fmt.Errorf("Error reading vrf6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak", sv)); err != nil {
			if !fortiAPIPatch(o["vrf-leak"]) {
				return fmt.Errorf("Error reading vrf_leak: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf_leak"); ok {
			if err = d.Set("vrf_leak", flattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak", sv)); err != nil {
				if !fortiAPIPatch(o["vrf-leak"]) {
					return fmt.Errorf("Error reading vrf_leak: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(o["vrf-leak6"], d, "vrf_leak6", sv)); err != nil {
			if !fortiAPIPatch(o["vrf-leak6"]) {
				return fmt.Errorf("Error reading vrf_leak6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrf_leak6"); ok {
			if err = d.Set("vrf_leak6", flattenRouterBgpVrfLeak6(o["vrf-leak6"], d, "vrf_leak6", sv)); err != nil {
				if !fortiAPIPatch(o["vrf-leak6"]) {
					return fmt.Errorf("Error reading vrf_leak6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterBgpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterBgpAsString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpKeepaliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAlwaysCompareMed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathAsPathIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpConfedAspath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathCmpRouterid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathMedConfed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpBestpathMedMissingAsWorst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpClientToClientReflection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampening(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDeterministicMed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpEbgpMultipath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpIbgpMultipath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpEnforceFirstAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpFastExternalFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkImportCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpIgnoreOptionalCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpMultipathRecursiveDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRecursiveNextHop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRecursiveInheritPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpTagResolveMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpClusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpConfederationIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpConfederationPeers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["peer"], _ = expandRouterBgpConfederationPeersPeer(d, i["peer"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpConfederationPeersPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningReachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningReuse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningSuppress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningMaxSuppressTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDampeningUnreachabilityHalfLife(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDefaultLocalPreference(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpScanTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceExternal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceInternal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpDistanceLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpSynchronization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulStalepathTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulUpdateDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpGracefulEndOnTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelect6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelectVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdditionalPathSelectVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpCrossFamilyConditionalAdv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpAggregateAddressId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterBgpAggregateAddressPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["as-set"], _ = expandRouterBgpAggregateAddressAsSet(d, i["as_set"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["summary-only"], _ = expandRouterBgpAggregateAddressSummaryOnly(d, i["summary_only"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpAggregateAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressAsSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddressSummaryOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpAggregateAddress6Id(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterBgpAggregateAddress6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["as-set"], _ = expandRouterBgpAggregateAddress6AsSet(d, i["as_set"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["summary-only"], _ = expandRouterBgpAggregateAddress6SummaryOnly(d, i["summary_only"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpAggregateAddress6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6AsSet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAggregateAddress6SummaryOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandRouterBgpNeighborIp(d, i["ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertisement-interval"], _ = expandRouterBgpNeighborAdvertisementInterval(d, i["advertisement_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable"], _ = expandRouterBgpNeighborAllowasInEnable(d, i["allowas_in_enable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable6"], _ = expandRouterBgpNeighborAllowasInEnable6(d, i["allowas_in_enable6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable-vpnv4"], _ = expandRouterBgpNeighborAllowasInEnableVpnv4(d, i["allowas_in_enable_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable-vpnv6"], _ = expandRouterBgpNeighborAllowasInEnableVpnv6(d, i["allowas_in_enable_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable-evpn"], _ = expandRouterBgpNeighborAllowasInEnableEvpn(d, i["allowas_in_enable_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in"], _ = expandRouterBgpNeighborAllowasIn(d, i["allowas_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in6"], _ = expandRouterBgpNeighborAllowasIn6(d, i["allowas_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-vpnv4"], _ = expandRouterBgpNeighborAllowasInVpnv4(d, i["allowas_in_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-vpnv6"], _ = expandRouterBgpNeighborAllowasInVpnv6(d, i["allowas_in_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-evpn"], _ = expandRouterBgpNeighborAllowasInEvpn(d, i["allowas_in_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged"], _ = expandRouterBgpNeighborAttributeUnchanged(d, i["attribute_unchanged"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged6"], _ = expandRouterBgpNeighborAttributeUnchanged6(d, i["attribute_unchanged6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged-vpnv4"], _ = expandRouterBgpNeighborAttributeUnchangedVpnv4(d, i["attribute_unchanged_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged-vpnv6"], _ = expandRouterBgpNeighborAttributeUnchangedVpnv6(d, i["attribute_unchanged_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate"], _ = expandRouterBgpNeighborActivate(d, i["activate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate6"], _ = expandRouterBgpNeighborActivate6(d, i["activate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate-vpnv4"], _ = expandRouterBgpNeighborActivateVpnv4(d, i["activate_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate-vpnv6"], _ = expandRouterBgpNeighborActivateVpnv6(d, i["activate_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate-evpn"], _ = expandRouterBgpNeighborActivateEvpn(d, i["activate_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterBgpNeighborBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-dynamic"], _ = expandRouterBgpNeighborCapabilityDynamic(d, i["capability_dynamic"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-orf"], _ = expandRouterBgpNeighborCapabilityOrf(d, i["capability_orf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-orf6"], _ = expandRouterBgpNeighborCapabilityOrf6(d, i["capability_orf6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart"], _ = expandRouterBgpNeighborCapabilityGracefulRestart(d, i["capability_graceful_restart"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart6"], _ = expandRouterBgpNeighborCapabilityGracefulRestart6(d, i["capability_graceful_restart6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart-vpnv4"], _ = expandRouterBgpNeighborCapabilityGracefulRestartVpnv4(d, i["capability_graceful_restart_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart-vpnv6"], _ = expandRouterBgpNeighborCapabilityGracefulRestartVpnv6(d, i["capability_graceful_restart_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart-evpn"], _ = expandRouterBgpNeighborCapabilityGracefulRestartEvpn(d, i["capability_graceful_restart_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-route-refresh"], _ = expandRouterBgpNeighborCapabilityRouteRefresh(d, i["capability_route_refresh"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-default-originate"], _ = expandRouterBgpNeighborCapabilityDefaultOriginate(d, i["capability_default_originate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-default-originate6"], _ = expandRouterBgpNeighborCapabilityDefaultOriginate6(d, i["capability_default_originate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dont-capability-negotiate"], _ = expandRouterBgpNeighborDontCapabilityNegotiate(d, i["dont_capability_negotiate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ebgp-enforce-multihop"], _ = expandRouterBgpNeighborEbgpEnforceMultihop(d, i["ebgp_enforce_multihop"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["link-down-failover"], _ = expandRouterBgpNeighborLinkDownFailover(d, i["link_down_failover"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stale-route"], _ = expandRouterBgpNeighborStaleRoute(d, i["stale_route"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self"], _ = expandRouterBgpNeighborNextHopSelf(d, i["next_hop_self"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self6"], _ = expandRouterBgpNeighborNextHopSelf6(d, i["next_hop_self6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-rr"], _ = expandRouterBgpNeighborNextHopSelfRr(d, i["next_hop_self_rr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-rr6"], _ = expandRouterBgpNeighborNextHopSelfRr6(d, i["next_hop_self_rr6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-vpnv4"], _ = expandRouterBgpNeighborNextHopSelfVpnv4(d, i["next_hop_self_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-vpnv6"], _ = expandRouterBgpNeighborNextHopSelfVpnv6(d, i["next_hop_self_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-capability"], _ = expandRouterBgpNeighborOverrideCapability(d, i["override_capability"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passive"], _ = expandRouterBgpNeighborPassive(d, i["passive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as"], _ = expandRouterBgpNeighborRemovePrivateAs(d, i["remove_private_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as6"], _ = expandRouterBgpNeighborRemovePrivateAs6(d, i["remove_private_as6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as-vpnv4"], _ = expandRouterBgpNeighborRemovePrivateAsVpnv4(d, i["remove_private_as_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as-vpnv6"], _ = expandRouterBgpNeighborRemovePrivateAsVpnv6(d, i["remove_private_as_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as-evpn"], _ = expandRouterBgpNeighborRemovePrivateAsEvpn(d, i["remove_private_as_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client"], _ = expandRouterBgpNeighborRouteReflectorClient(d, i["route_reflector_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client6"], _ = expandRouterBgpNeighborRouteReflectorClient6(d, i["route_reflector_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client-vpnv4"], _ = expandRouterBgpNeighborRouteReflectorClientVpnv4(d, i["route_reflector_client_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client-vpnv6"], _ = expandRouterBgpNeighborRouteReflectorClientVpnv6(d, i["route_reflector_client_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client-evpn"], _ = expandRouterBgpNeighborRouteReflectorClientEvpn(d, i["route_reflector_client_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client"], _ = expandRouterBgpNeighborRouteServerClient(d, i["route_server_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client6"], _ = expandRouterBgpNeighborRouteServerClient6(d, i["route_server_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client-vpnv4"], _ = expandRouterBgpNeighborRouteServerClientVpnv4(d, i["route_server_client_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client-vpnv6"], _ = expandRouterBgpNeighborRouteServerClientVpnv6(d, i["route_server_client_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client-evpn"], _ = expandRouterBgpNeighborRouteServerClientEvpn(d, i["route_server_client_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shutdown"], _ = expandRouterBgpNeighborShutdown(d, i["shutdown"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration"], _ = expandRouterBgpNeighborSoftReconfiguration(d, i["soft_reconfiguration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration6"], _ = expandRouterBgpNeighborSoftReconfiguration6(d, i["soft_reconfiguration6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration-vpnv4"], _ = expandRouterBgpNeighborSoftReconfigurationVpnv4(d, i["soft_reconfiguration_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration-vpnv6"], _ = expandRouterBgpNeighborSoftReconfigurationVpnv6(d, i["soft_reconfiguration_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration-evpn"], _ = expandRouterBgpNeighborSoftReconfigurationEvpn(d, i["soft_reconfiguration_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["as-override"], _ = expandRouterBgpNeighborAsOverride(d, i["as_override"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["as-override6"], _ = expandRouterBgpNeighborAsOverride6(d, i["as_override6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["strict-capability-match"], _ = expandRouterBgpNeighborStrictCapabilityMatch(d, i["strict_capability_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-originate-routemap"], _ = expandRouterBgpNeighborDefaultOriginateRoutemap(d, i["default_originate_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["default-originate-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-originate-routemap6"], _ = expandRouterBgpNeighborDefaultOriginateRoutemap6(d, i["default_originate_routemap6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["default-originate-routemap6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandRouterBgpNeighborDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in"], _ = expandRouterBgpNeighborDistributeListIn(d, i["distribute_list_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in6"], _ = expandRouterBgpNeighborDistributeListIn6(d, i["distribute_list_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in-vpnv4"], _ = expandRouterBgpNeighborDistributeListInVpnv4(d, i["distribute_list_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in-vpnv6"], _ = expandRouterBgpNeighborDistributeListInVpnv6(d, i["distribute_list_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out"], _ = expandRouterBgpNeighborDistributeListOut(d, i["distribute_list_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out6"], _ = expandRouterBgpNeighborDistributeListOut6(d, i["distribute_list_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out-vpnv4"], _ = expandRouterBgpNeighborDistributeListOutVpnv4(d, i["distribute_list_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out-vpnv6"], _ = expandRouterBgpNeighborDistributeListOutVpnv6(d, i["distribute_list_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ebgp-multihop-ttl"], _ = expandRouterBgpNeighborEbgpMultihopTtl(d, i["ebgp_multihop_ttl"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in"], _ = expandRouterBgpNeighborFilterListIn(d, i["filter_list_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in6"], _ = expandRouterBgpNeighborFilterListIn6(d, i["filter_list_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in-vpnv4"], _ = expandRouterBgpNeighborFilterListInVpnv4(d, i["filter_list_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in-vpnv6"], _ = expandRouterBgpNeighborFilterListInVpnv6(d, i["filter_list_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out"], _ = expandRouterBgpNeighborFilterListOut(d, i["filter_list_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out6"], _ = expandRouterBgpNeighborFilterListOut6(d, i["filter_list_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out-vpnv4"], _ = expandRouterBgpNeighborFilterListOutVpnv4(d, i["filter_list_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out-vpnv6"], _ = expandRouterBgpNeighborFilterListOutVpnv6(d, i["filter_list_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBgpNeighborInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix"], _ = expandRouterBgpNeighborMaximumPrefix(d, i["maximum_prefix"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix6"], _ = expandRouterBgpNeighborMaximumPrefix6(d, i["maximum_prefix6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-vpnv4"], _ = expandRouterBgpNeighborMaximumPrefixVpnv4(d, i["maximum_prefix_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-vpnv6"], _ = expandRouterBgpNeighborMaximumPrefixVpnv6(d, i["maximum_prefix_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-evpn"], _ = expandRouterBgpNeighborMaximumPrefixEvpn(d, i["maximum_prefix_evpn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix-evpn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold"], _ = expandRouterBgpNeighborMaximumPrefixThreshold(d, i["maximum_prefix_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold6"], _ = expandRouterBgpNeighborMaximumPrefixThreshold6(d, i["maximum_prefix_threshold6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold-vpnv4"], _ = expandRouterBgpNeighborMaximumPrefixThresholdVpnv4(d, i["maximum_prefix_threshold_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold-vpnv6"], _ = expandRouterBgpNeighborMaximumPrefixThresholdVpnv6(d, i["maximum_prefix_threshold_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold-evpn"], _ = expandRouterBgpNeighborMaximumPrefixThresholdEvpn(d, i["maximum_prefix_threshold_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnly(d, i["maximum_prefix_warning_only"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only6"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnly6(d, i["maximum_prefix_warning_only6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only-vpnv4"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(d, i["maximum_prefix_warning_only_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only-vpnv6"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(d, i["maximum_prefix_warning_only_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only-evpn"], _ = expandRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(d, i["maximum_prefix_warning_only_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in"], _ = expandRouterBgpNeighborPrefixListIn(d, i["prefix_list_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in6"], _ = expandRouterBgpNeighborPrefixListIn6(d, i["prefix_list_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in-vpnv4"], _ = expandRouterBgpNeighborPrefixListInVpnv4(d, i["prefix_list_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in-vpnv6"], _ = expandRouterBgpNeighborPrefixListInVpnv6(d, i["prefix_list_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out"], _ = expandRouterBgpNeighborPrefixListOut(d, i["prefix_list_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out6"], _ = expandRouterBgpNeighborPrefixListOut6(d, i["prefix_list_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out-vpnv4"], _ = expandRouterBgpNeighborPrefixListOutVpnv4(d, i["prefix_list_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out-vpnv6"], _ = expandRouterBgpNeighborPrefixListOutVpnv6(d, i["prefix_list_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-as"], _ = expandRouterBgpNeighborRemoteAs(d, i["remote_as"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["remote-as"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["local-as"], _ = expandRouterBgpNeighborLocalAs(d, i["local_as"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["local-as"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["local-as-no-prepend"], _ = expandRouterBgpNeighborLocalAsNoPrepend(d, i["local_as_no_prepend"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["local-as-replace-as"], _ = expandRouterBgpNeighborLocalAsReplaceAs(d, i["local_as_replace_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retain-stale-time"], _ = expandRouterBgpNeighborRetainStaleTime(d, i["retain_stale_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["retain-stale-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in"], _ = expandRouterBgpNeighborRouteMapIn(d, i["route_map_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in6"], _ = expandRouterBgpNeighborRouteMapIn6(d, i["route_map_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in-vpnv4"], _ = expandRouterBgpNeighborRouteMapInVpnv4(d, i["route_map_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in-vpnv6"], _ = expandRouterBgpNeighborRouteMapInVpnv6(d, i["route_map_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in-evpn"], _ = expandRouterBgpNeighborRouteMapInEvpn(d, i["route_map_in_evpn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in-evpn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out"], _ = expandRouterBgpNeighborRouteMapOut(d, i["route_map_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-preferable"], _ = expandRouterBgpNeighborRouteMapOutPreferable(d, i["route_map_out_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out6"], _ = expandRouterBgpNeighborRouteMapOut6(d, i["route_map_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out6-preferable"], _ = expandRouterBgpNeighborRouteMapOut6Preferable(d, i["route_map_out6_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out6-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv4"], _ = expandRouterBgpNeighborRouteMapOutVpnv4(d, i["route_map_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv6"], _ = expandRouterBgpNeighborRouteMapOutVpnv6(d, i["route_map_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv4-preferable"], _ = expandRouterBgpNeighborRouteMapOutVpnv4Preferable(d, i["route_map_out_vpnv4_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv6-preferable"], _ = expandRouterBgpNeighborRouteMapOutVpnv6Preferable(d, i["route_map_out_vpnv6_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-evpn"], _ = expandRouterBgpNeighborRouteMapOutEvpn(d, i["route_map_out_evpn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-evpn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community"], _ = expandRouterBgpNeighborSendCommunity(d, i["send_community"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community6"], _ = expandRouterBgpNeighborSendCommunity6(d, i["send_community6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community-vpnv4"], _ = expandRouterBgpNeighborSendCommunityVpnv4(d, i["send_community_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community-vpnv6"], _ = expandRouterBgpNeighborSendCommunityVpnv6(d, i["send_community_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community-evpn"], _ = expandRouterBgpNeighborSendCommunityEvpn(d, i["send_community_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keep-alive-timer"], _ = expandRouterBgpNeighborKeepAliveTimer(d, i["keep_alive_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holdtime-timer"], _ = expandRouterBgpNeighborHoldtimeTimer(d, i["holdtime_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["connect-timer"], _ = expandRouterBgpNeighborConnectTimer(d, i["connect_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["unsuppress-map"], _ = expandRouterBgpNeighborUnsuppressMap(d, i["unsuppress_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["unsuppress-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["unsuppress-map6"], _ = expandRouterBgpNeighborUnsuppressMap6(d, i["unsuppress_map6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["unsuppress-map6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["update-source"], _ = expandRouterBgpNeighborUpdateSource(d, i["update_source"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["update-source"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandRouterBgpNeighborWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["restart-time"], _ = expandRouterBgpNeighborRestartTime(d, i["restart_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["restart-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path"], _ = expandRouterBgpNeighborAdditionalPath(d, i["additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path6"], _ = expandRouterBgpNeighborAdditionalPath6(d, i["additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path-vpnv4"], _ = expandRouterBgpNeighborAdditionalPathVpnv4(d, i["additional_path_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path-vpnv6"], _ = expandRouterBgpNeighborAdditionalPathVpnv6(d, i["additional_path_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path"], _ = expandRouterBgpNeighborAdvAdditionalPath(d, i["adv_additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path6"], _ = expandRouterBgpNeighborAdvAdditionalPath6(d, i["adv_additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path-vpnv4"], _ = expandRouterBgpNeighborAdvAdditionalPathVpnv4(d, i["adv_additional_path_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path-vpnv6"], _ = expandRouterBgpNeighborAdvAdditionalPathVpnv6(d, i["adv_additional_path_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandRouterBgpNeighborPassword(d, i["password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-options"], _ = expandRouterBgpNeighborAuthOptions(d, i["auth_options"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-options"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["conditional-advertise"], _ = expandRouterBgpNeighborConditionalAdvertise(d, i["conditional_advertise"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["conditional-advertise"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["conditional-advertise6"], _ = expandRouterBgpNeighborConditionalAdvertise6(d, i["conditional_advertise6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["conditional-advertise6"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEnableEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAllowasInEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchangedVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAttributeUnchangedVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborActivateEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityGracefulRestartEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLinkDownFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStaleRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborNextHopSelfVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborOverrideCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemovePrivateAsEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteReflectorClientEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteServerClientEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborShutdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSoftReconfigurationEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAsOverride6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborDistributeListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborFilterListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixThresholdEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborMaximumPrefixWarningOnlyEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPrefixListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterBgpNeighborLocalAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterBgpNeighborLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRetainStaleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapInEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutVpnv4Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutVpnv6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRouteMapOutEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunity6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborSendCommunityEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConnectTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUnsuppressMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborUpdateSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAdvAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborAuthOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(d, i["advertise_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["advertise-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d, i["condition_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["condition-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertiseConditionType(d, i["condition_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertiseConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandRouterBgpNeighborConditionalAdvertiseConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["advertise-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(d, i["advertise_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["advertise-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-routemap"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d, i["condition_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["condition-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["condition-type"], _ = expandRouterBgpNeighborConditionalAdvertise6ConditionType(d, i["condition_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborConditionalAdvertise6AdvertiseRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborConditionalAdvertise6ConditionRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandRouterBgpNeighborConditionalAdvertise6ConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterBgpNeighborGroupName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertisement-interval"], _ = expandRouterBgpNeighborGroupAdvertisementInterval(d, i["advertisement_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable"], _ = expandRouterBgpNeighborGroupAllowasInEnable(d, i["allowas_in_enable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable6"], _ = expandRouterBgpNeighborGroupAllowasInEnable6(d, i["allowas_in_enable6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable-vpnv4"], _ = expandRouterBgpNeighborGroupAllowasInEnableVpnv4(d, i["allowas_in_enable_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable-vpnv6"], _ = expandRouterBgpNeighborGroupAllowasInEnableVpnv6(d, i["allowas_in_enable_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-enable-evpn"], _ = expandRouterBgpNeighborGroupAllowasInEnableEvpn(d, i["allowas_in_enable_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in"], _ = expandRouterBgpNeighborGroupAllowasIn(d, i["allowas_in"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in6"], _ = expandRouterBgpNeighborGroupAllowasIn6(d, i["allowas_in6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-vpnv4"], _ = expandRouterBgpNeighborGroupAllowasInVpnv4(d, i["allowas_in_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-vpnv6"], _ = expandRouterBgpNeighborGroupAllowasInVpnv6(d, i["allowas_in_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowas-in-evpn"], _ = expandRouterBgpNeighborGroupAllowasInEvpn(d, i["allowas_in_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged"], _ = expandRouterBgpNeighborGroupAttributeUnchanged(d, i["attribute_unchanged"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged6"], _ = expandRouterBgpNeighborGroupAttributeUnchanged6(d, i["attribute_unchanged6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged-vpnv4"], _ = expandRouterBgpNeighborGroupAttributeUnchangedVpnv4(d, i["attribute_unchanged_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-unchanged-vpnv6"], _ = expandRouterBgpNeighborGroupAttributeUnchangedVpnv6(d, i["attribute_unchanged_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-unchanged-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate"], _ = expandRouterBgpNeighborGroupActivate(d, i["activate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate6"], _ = expandRouterBgpNeighborGroupActivate6(d, i["activate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate-vpnv4"], _ = expandRouterBgpNeighborGroupActivateVpnv4(d, i["activate_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate-vpnv6"], _ = expandRouterBgpNeighborGroupActivateVpnv6(d, i["activate_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["activate-evpn"], _ = expandRouterBgpNeighborGroupActivateEvpn(d, i["activate_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterBgpNeighborGroupBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-dynamic"], _ = expandRouterBgpNeighborGroupCapabilityDynamic(d, i["capability_dynamic"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-orf"], _ = expandRouterBgpNeighborGroupCapabilityOrf(d, i["capability_orf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-orf6"], _ = expandRouterBgpNeighborGroupCapabilityOrf6(d, i["capability_orf6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestart(d, i["capability_graceful_restart"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart6"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestart6(d, i["capability_graceful_restart6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart-vpnv4"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(d, i["capability_graceful_restart_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart-vpnv6"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(d, i["capability_graceful_restart_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-graceful-restart-evpn"], _ = expandRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(d, i["capability_graceful_restart_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-route-refresh"], _ = expandRouterBgpNeighborGroupCapabilityRouteRefresh(d, i["capability_route_refresh"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-default-originate"], _ = expandRouterBgpNeighborGroupCapabilityDefaultOriginate(d, i["capability_default_originate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["capability-default-originate6"], _ = expandRouterBgpNeighborGroupCapabilityDefaultOriginate6(d, i["capability_default_originate6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dont-capability-negotiate"], _ = expandRouterBgpNeighborGroupDontCapabilityNegotiate(d, i["dont_capability_negotiate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ebgp-enforce-multihop"], _ = expandRouterBgpNeighborGroupEbgpEnforceMultihop(d, i["ebgp_enforce_multihop"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["link-down-failover"], _ = expandRouterBgpNeighborGroupLinkDownFailover(d, i["link_down_failover"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stale-route"], _ = expandRouterBgpNeighborGroupStaleRoute(d, i["stale_route"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self"], _ = expandRouterBgpNeighborGroupNextHopSelf(d, i["next_hop_self"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self6"], _ = expandRouterBgpNeighborGroupNextHopSelf6(d, i["next_hop_self6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-rr"], _ = expandRouterBgpNeighborGroupNextHopSelfRr(d, i["next_hop_self_rr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-rr6"], _ = expandRouterBgpNeighborGroupNextHopSelfRr6(d, i["next_hop_self_rr6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-vpnv4"], _ = expandRouterBgpNeighborGroupNextHopSelfVpnv4(d, i["next_hop_self_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop-self-vpnv6"], _ = expandRouterBgpNeighborGroupNextHopSelfVpnv6(d, i["next_hop_self_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-capability"], _ = expandRouterBgpNeighborGroupOverrideCapability(d, i["override_capability"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passive"], _ = expandRouterBgpNeighborGroupPassive(d, i["passive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as"], _ = expandRouterBgpNeighborGroupRemovePrivateAs(d, i["remove_private_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as6"], _ = expandRouterBgpNeighborGroupRemovePrivateAs6(d, i["remove_private_as6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as-vpnv4"], _ = expandRouterBgpNeighborGroupRemovePrivateAsVpnv4(d, i["remove_private_as_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as-vpnv6"], _ = expandRouterBgpNeighborGroupRemovePrivateAsVpnv6(d, i["remove_private_as_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remove-private-as-evpn"], _ = expandRouterBgpNeighborGroupRemovePrivateAsEvpn(d, i["remove_private_as_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client"], _ = expandRouterBgpNeighborGroupRouteReflectorClient(d, i["route_reflector_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client6"], _ = expandRouterBgpNeighborGroupRouteReflectorClient6(d, i["route_reflector_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client-vpnv4"], _ = expandRouterBgpNeighborGroupRouteReflectorClientVpnv4(d, i["route_reflector_client_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client-vpnv6"], _ = expandRouterBgpNeighborGroupRouteReflectorClientVpnv6(d, i["route_reflector_client_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-reflector-client-evpn"], _ = expandRouterBgpNeighborGroupRouteReflectorClientEvpn(d, i["route_reflector_client_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client"], _ = expandRouterBgpNeighborGroupRouteServerClient(d, i["route_server_client"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client6"], _ = expandRouterBgpNeighborGroupRouteServerClient6(d, i["route_server_client6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client-vpnv4"], _ = expandRouterBgpNeighborGroupRouteServerClientVpnv4(d, i["route_server_client_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client-vpnv6"], _ = expandRouterBgpNeighborGroupRouteServerClientVpnv6(d, i["route_server_client_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-server-client-evpn"], _ = expandRouterBgpNeighborGroupRouteServerClientEvpn(d, i["route_server_client_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shutdown"], _ = expandRouterBgpNeighborGroupShutdown(d, i["shutdown"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration"], _ = expandRouterBgpNeighborGroupSoftReconfiguration(d, i["soft_reconfiguration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration6"], _ = expandRouterBgpNeighborGroupSoftReconfiguration6(d, i["soft_reconfiguration6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration-vpnv4"], _ = expandRouterBgpNeighborGroupSoftReconfigurationVpnv4(d, i["soft_reconfiguration_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration-vpnv6"], _ = expandRouterBgpNeighborGroupSoftReconfigurationVpnv6(d, i["soft_reconfiguration_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["soft-reconfiguration-evpn"], _ = expandRouterBgpNeighborGroupSoftReconfigurationEvpn(d, i["soft_reconfiguration_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["as-override"], _ = expandRouterBgpNeighborGroupAsOverride(d, i["as_override"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["as-override6"], _ = expandRouterBgpNeighborGroupAsOverride6(d, i["as_override6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["strict-capability-match"], _ = expandRouterBgpNeighborGroupStrictCapabilityMatch(d, i["strict_capability_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-originate-routemap"], _ = expandRouterBgpNeighborGroupDefaultOriginateRoutemap(d, i["default_originate_routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["default-originate-routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-originate-routemap6"], _ = expandRouterBgpNeighborGroupDefaultOriginateRoutemap6(d, i["default_originate_routemap6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["default-originate-routemap6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandRouterBgpNeighborGroupDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in"], _ = expandRouterBgpNeighborGroupDistributeListIn(d, i["distribute_list_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in6"], _ = expandRouterBgpNeighborGroupDistributeListIn6(d, i["distribute_list_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in-vpnv4"], _ = expandRouterBgpNeighborGroupDistributeListInVpnv4(d, i["distribute_list_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-in-vpnv6"], _ = expandRouterBgpNeighborGroupDistributeListInVpnv6(d, i["distribute_list_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out"], _ = expandRouterBgpNeighborGroupDistributeListOut(d, i["distribute_list_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out6"], _ = expandRouterBgpNeighborGroupDistributeListOut6(d, i["distribute_list_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out-vpnv4"], _ = expandRouterBgpNeighborGroupDistributeListOutVpnv4(d, i["distribute_list_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distribute-list-out-vpnv6"], _ = expandRouterBgpNeighborGroupDistributeListOutVpnv6(d, i["distribute_list_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distribute-list-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ebgp-multihop-ttl"], _ = expandRouterBgpNeighborGroupEbgpMultihopTtl(d, i["ebgp_multihop_ttl"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in"], _ = expandRouterBgpNeighborGroupFilterListIn(d, i["filter_list_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in6"], _ = expandRouterBgpNeighborGroupFilterListIn6(d, i["filter_list_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in-vpnv4"], _ = expandRouterBgpNeighborGroupFilterListInVpnv4(d, i["filter_list_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-in-vpnv6"], _ = expandRouterBgpNeighborGroupFilterListInVpnv6(d, i["filter_list_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out"], _ = expandRouterBgpNeighborGroupFilterListOut(d, i["filter_list_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out6"], _ = expandRouterBgpNeighborGroupFilterListOut6(d, i["filter_list_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out-vpnv4"], _ = expandRouterBgpNeighborGroupFilterListOutVpnv4(d, i["filter_list_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list-out-vpnv6"], _ = expandRouterBgpNeighborGroupFilterListOutVpnv6(d, i["filter_list_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBgpNeighborGroupInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix"], _ = expandRouterBgpNeighborGroupMaximumPrefix(d, i["maximum_prefix"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix6"], _ = expandRouterBgpNeighborGroupMaximumPrefix6(d, i["maximum_prefix6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-vpnv4"], _ = expandRouterBgpNeighborGroupMaximumPrefixVpnv4(d, i["maximum_prefix_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-vpnv6"], _ = expandRouterBgpNeighborGroupMaximumPrefixVpnv6(d, i["maximum_prefix_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-evpn"], _ = expandRouterBgpNeighborGroupMaximumPrefixEvpn(d, i["maximum_prefix_evpn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["maximum-prefix-evpn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold"], _ = expandRouterBgpNeighborGroupMaximumPrefixThreshold(d, i["maximum_prefix_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold6"], _ = expandRouterBgpNeighborGroupMaximumPrefixThreshold6(d, i["maximum_prefix_threshold6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold-vpnv4"], _ = expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(d, i["maximum_prefix_threshold_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold-vpnv6"], _ = expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(d, i["maximum_prefix_threshold_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-threshold-evpn"], _ = expandRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(d, i["maximum_prefix_threshold_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnly(d, i["maximum_prefix_warning_only"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only6"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnly6(d, i["maximum_prefix_warning_only6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only-vpnv4"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(d, i["maximum_prefix_warning_only_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only-vpnv6"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(d, i["maximum_prefix_warning_only_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["maximum-prefix-warning-only-evpn"], _ = expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(d, i["maximum_prefix_warning_only_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in"], _ = expandRouterBgpNeighborGroupPrefixListIn(d, i["prefix_list_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in6"], _ = expandRouterBgpNeighborGroupPrefixListIn6(d, i["prefix_list_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in-vpnv4"], _ = expandRouterBgpNeighborGroupPrefixListInVpnv4(d, i["prefix_list_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-in-vpnv6"], _ = expandRouterBgpNeighborGroupPrefixListInVpnv6(d, i["prefix_list_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out"], _ = expandRouterBgpNeighborGroupPrefixListOut(d, i["prefix_list_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out6"], _ = expandRouterBgpNeighborGroupPrefixListOut6(d, i["prefix_list_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out-vpnv4"], _ = expandRouterBgpNeighborGroupPrefixListOutVpnv4(d, i["prefix_list_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-list-out-vpnv6"], _ = expandRouterBgpNeighborGroupPrefixListOutVpnv6(d, i["prefix_list_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-list-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-as"], _ = expandRouterBgpNeighborGroupRemoteAs(d, i["remote_as"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["remote-as"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as_filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-as-filter"], _ = expandRouterBgpNeighborGroupRemoteAsFilter(d, i["remote_as_filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["remote-as-filter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["local-as"], _ = expandRouterBgpNeighborGroupLocalAs(d, i["local_as"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["local-as"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["local-as-no-prepend"], _ = expandRouterBgpNeighborGroupLocalAsNoPrepend(d, i["local_as_no_prepend"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["local-as-replace-as"], _ = expandRouterBgpNeighborGroupLocalAsReplaceAs(d, i["local_as_replace_as"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retain-stale-time"], _ = expandRouterBgpNeighborGroupRetainStaleTime(d, i["retain_stale_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["retain-stale-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in"], _ = expandRouterBgpNeighborGroupRouteMapIn(d, i["route_map_in"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in6"], _ = expandRouterBgpNeighborGroupRouteMapIn6(d, i["route_map_in6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in-vpnv4"], _ = expandRouterBgpNeighborGroupRouteMapInVpnv4(d, i["route_map_in_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in-vpnv6"], _ = expandRouterBgpNeighborGroupRouteMapInVpnv6(d, i["route_map_in_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-in-evpn"], _ = expandRouterBgpNeighborGroupRouteMapInEvpn(d, i["route_map_in_evpn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-in-evpn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out"], _ = expandRouterBgpNeighborGroupRouteMapOut(d, i["route_map_out"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutPreferable(d, i["route_map_out_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out6"], _ = expandRouterBgpNeighborGroupRouteMapOut6(d, i["route_map_out6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out6-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOut6Preferable(d, i["route_map_out6_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out6-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv4"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv4(d, i["route_map_out_vpnv4"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv6"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv6(d, i["route_map_out_vpnv6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv4_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv4-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(d, i["route_map_out_vpnv4_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv4-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_vpnv6_preferable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-vpnv6-preferable"], _ = expandRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(d, i["route_map_out_vpnv6_preferable"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-vpnv6-preferable"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map-out-evpn"], _ = expandRouterBgpNeighborGroupRouteMapOutEvpn(d, i["route_map_out_evpn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map-out-evpn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community"], _ = expandRouterBgpNeighborGroupSendCommunity(d, i["send_community"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community6"], _ = expandRouterBgpNeighborGroupSendCommunity6(d, i["send_community6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community-vpnv4"], _ = expandRouterBgpNeighborGroupSendCommunityVpnv4(d, i["send_community_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community-vpnv6"], _ = expandRouterBgpNeighborGroupSendCommunityVpnv6(d, i["send_community_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community_evpn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-community-evpn"], _ = expandRouterBgpNeighborGroupSendCommunityEvpn(d, i["send_community_evpn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keep-alive-timer"], _ = expandRouterBgpNeighborGroupKeepAliveTimer(d, i["keep_alive_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["holdtime-timer"], _ = expandRouterBgpNeighborGroupHoldtimeTimer(d, i["holdtime_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["connect-timer"], _ = expandRouterBgpNeighborGroupConnectTimer(d, i["connect_timer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["unsuppress-map"], _ = expandRouterBgpNeighborGroupUnsuppressMap(d, i["unsuppress_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["unsuppress-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["unsuppress-map6"], _ = expandRouterBgpNeighborGroupUnsuppressMap6(d, i["unsuppress_map6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["unsuppress-map6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["update-source"], _ = expandRouterBgpNeighborGroupUpdateSource(d, i["update_source"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["update-source"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandRouterBgpNeighborGroupWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["restart-time"], _ = expandRouterBgpNeighborGroupRestartTime(d, i["restart_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["restart-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path"], _ = expandRouterBgpNeighborGroupAdditionalPath(d, i["additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path6"], _ = expandRouterBgpNeighborGroupAdditionalPath6(d, i["additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path-vpnv4"], _ = expandRouterBgpNeighborGroupAdditionalPathVpnv4(d, i["additional_path_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-path-vpnv6"], _ = expandRouterBgpNeighborGroupAdditionalPathVpnv6(d, i["additional_path_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path"], _ = expandRouterBgpNeighborGroupAdvAdditionalPath(d, i["adv_additional_path"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path6"], _ = expandRouterBgpNeighborGroupAdvAdditionalPath6(d, i["adv_additional_path6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path-vpnv4"], _ = expandRouterBgpNeighborGroupAdvAdditionalPathVpnv4(d, i["adv_additional_path_vpnv4"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path_vpnv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-additional-path-vpnv6"], _ = expandRouterBgpNeighborGroupAdvAdditionalPathVpnv6(d, i["adv_additional_path_vpnv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandRouterBgpNeighborGroupPassword(d, i["password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_options"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-options"], _ = expandRouterBgpNeighborGroupAuthOptions(d, i["auth_options"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-options"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvertisementInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnable6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEnableEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAllowasInEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchanged6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchangedVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAttributeUnchangedVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupActivateEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDynamic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityOrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestart6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityGracefulRestartEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityRouteRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupCapabilityDefaultOriginate6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDontCapabilityNegotiate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpEnforceMultihop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLinkDownFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStaleRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfRr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupNextHopSelfVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupOverrideCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAs6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemovePrivateAsEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteReflectorClientEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClient6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteServerClientEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupShutdown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfiguration6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSoftReconfigurationEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAsOverride6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupStrictCapabilityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDefaultOriginateRoutemap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupDistributeListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupEbgpMultihopTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupFilterListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThreshold6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixThresholdEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnly6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupMaximumPrefixWarningOnlyEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPrefixListOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRemoteAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterBgpNeighborGroupRemoteAsFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return convintf2i(v), nil
}

func expandRouterBgpNeighborGroupLocalAsNoPrepend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupLocalAsReplaceAs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRetainStaleTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapIn6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapInVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapInVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapInEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutPreferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOut6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOut6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv4Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutVpnv6Preferable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRouteMapOutEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunity6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupSendCommunityEvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupKeepAliveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupHoldtimeTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupConnectTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUnsuppressMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUnsuppressMap6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupUpdateSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupRestartTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPath6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPathVpnv4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAdvAdditionalPathVpnv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborGroupAuthOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpNeighborRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterBgpNeighborRangePrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-neighbor-num"], _ = expandRouterBgpNeighborRangeMaxNeighborNum(d, i["max_neighbor_num"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["max-neighbor-num"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["neighbor-group"], _ = expandRouterBgpNeighborRangeNeighborGroup(d, i["neighbor_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["neighbor-group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeMaxNeighborNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRangeNeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpNeighborRange6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterBgpNeighborRange6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-neighbor-num"], _ = expandRouterBgpNeighborRange6MaxNeighborNum(d, i["max_neighbor_num"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["max-neighbor-num"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["neighbor-group"], _ = expandRouterBgpNeighborRange6NeighborGroup(d, i["neighbor_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["neighbor-group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNeighborRange6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6MaxNeighborNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNeighborRange6NeighborGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpNetworkId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterBgpNetworkPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["network-import-check"], _ = expandRouterBgpNetworkNetworkImportCheck(d, i["network_import_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["backdoor"], _ = expandRouterBgpNetworkBackdoor(d, i["backdoor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpNetworkRouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-name"], _ = expandRouterBgpNetworkPrefixName(d, i["prefix_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkNetworkImportCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkBackdoor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetworkPrefixName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpNetwork6Id(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterBgpNetwork6Prefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_import_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["network-import-check"], _ = expandRouterBgpNetwork6NetworkImportCheck(d, i["network_import_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["backdoor"], _ = expandRouterBgpNetwork6Backdoor(d, i["backdoor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpNetwork6RouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpNetwork6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6NetworkImportCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6Backdoor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpNetwork6RouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterBgpRedistributeName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterBgpRedistributeStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpRedistributeRouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpRedistributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistributeRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterBgpRedistribute6Name(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterBgpRedistribute6Status(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpRedistribute6RouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpRedistribute6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpRedistribute6RouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterBgpAdminDistanceId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["neighbour-prefix"], _ = expandRouterBgpAdminDistanceNeighbourPrefix(d, i["neighbour_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-list"], _ = expandRouterBgpAdminDistanceRouteList(d, i["route_list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-list"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distance"], _ = expandRouterBgpAdminDistanceDistance(d, i["distance"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["distance"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpAdminDistanceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceNeighbourPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceRouteList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpAdminDistanceDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrfVrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandRouterBgpVrfRole(d, i["role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rd"], _ = expandRouterBgpVrfRd(d, i["rd"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rd"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["export-rt"], _ = expandRouterBgpVrfExportRt(d, i["export_rt"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["export-rt"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["import-rt"], _ = expandRouterBgpVrfImportRt(d, i["import_rt"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["import-rt"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["import-route-map"], _ = expandRouterBgpVrfImportRouteMap(d, i["import_route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["import-route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["leak-target"], _ = expandRouterBgpVrfLeakTargetU(d, i["leak_target"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["leak-target"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfRd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfExportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["route-target"], _ = expandRouterBgpVrfExportRtRouteTarget(d, i["route_target"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfExportRtRouteTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfImportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["route-target"], _ = expandRouterBgpVrfImportRtRouteTarget(d, i["route_target"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfImportRtRouteTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfImportRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetU(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrfLeakTargetVrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpVrfLeakTargetRouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBgpVrfLeakTargetInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTargetUVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetURouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetUInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrf6Vrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandRouterBgpVrf6Role(d, i["role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rd"], _ = expandRouterBgpVrf6Rd(d, i["rd"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rd"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_rt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["export-rt"], _ = expandRouterBgpVrf6ExportRt(d, i["export_rt"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["export-rt"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_rt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["import-rt"], _ = expandRouterBgpVrf6ImportRt(d, i["import_rt"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["import-rt"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "import_route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["import-route-map"], _ = expandRouterBgpVrf6ImportRouteMap(d, i["import_route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["import-route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "leak_target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["leak-target"], _ = expandRouterBgpVrf6LeakTarget(d, i["leak_target"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["leak-target"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6Vrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Role(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6Rd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6ExportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["route-target"], _ = expandRouterBgpVrf6ExportRtRouteTarget(d, i["route_target"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6ExportRtRouteTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6ImportRt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["route-target"], _ = expandRouterBgpVrf6ImportRtRouteTarget(d, i["route_target"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6ImportRtRouteTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6ImportRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6LeakTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrf6LeakTargetVrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpVrf6LeakTargetRouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBgpVrf6LeakTargetInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrf6LeakTargetVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6LeakTargetRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrf6LeakTargetInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrfLeakVrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandRouterBgpVrfLeakTarget(d, i["target"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["target"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrfLeakTargetVrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpVrfLeakTargetRouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBgpVrfLeakTargetInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeakTargetVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeakTargetInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrfLeak6Vrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandRouterBgpVrfLeak6Target(d, i["target"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["target"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeak6Vrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6Target(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrf"], _ = expandRouterBgpVrfLeak6TargetVrf(d, i["vrf"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vrf"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-map"], _ = expandRouterBgpVrfLeak6TargetRouteMap(d, i["route_map"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-map"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterBgpVrfLeak6TargetInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterBgpVrfLeak6TargetVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6TargetRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterBgpVrfLeak6TargetInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterBgp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("as_string"); ok {
		if setArgNil {
			obj["as"] = nil
		} else {
			new_version_map := map[string][]string{
				">=": []string{"7.2.1"},
			}
			if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
				if _, ok := d.GetOk("as"); !ok && !d.HasChange("as") {
					err := fmt.Errorf("Argument 'as_string' %s.", err)
					return nil, err
				}
			} else {
				t, err := expandRouterBgpAsString(d, v, "as_string", sv)
				if err != nil {
					return &obj, err
				} else if t != nil {
					obj["as"] = t
				}
			}
		}
	} else if d.HasChange("as_string") {
		obj["as"] = nil
	}

	if v, ok := d.GetOkExists("as"); ok {
		if setArgNil {
			obj["as"] = nil
		} else {
			new_version_map := map[string][]string{
				"<=": []string{"7.2.0"},
			}
			if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
				if _, ok := d.GetOk("as_string"); !ok && !d.HasChange("as_string") {
					err := fmt.Errorf("Argument 'as' %s.", err)
					return nil, err
				}
			} else {
				t, err := expandRouterBgpAs(d, v, "as", sv)
				if err != nil {
					return &obj, err
				} else if t != nil {
					obj["as"] = t
				}
			}
		}
	} else if d.HasChange("as") {
		obj["as"] = nil
	}

	if v, ok := d.GetOk("router_id"); ok {
		if setArgNil {
			obj["router-id"] = nil
		} else {
			t, err := expandRouterBgpRouterId(d, v, "router_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router-id"] = t
			}
		}
	} else if d.HasChange("router_id") {
		obj["router-id"] = nil
	}

	if v, ok := d.GetOkExists("keepalive_timer"); ok {
		if setArgNil {
			obj["keepalive-timer"] = nil
		} else {
			t, err := expandRouterBgpKeepaliveTimer(d, v, "keepalive_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["keepalive-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("holdtime_timer"); ok {
		if setArgNil {
			obj["holdtime-timer"] = nil
		} else {
			t, err := expandRouterBgpHoldtimeTimer(d, v, "holdtime_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["holdtime-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("always_compare_med"); ok {
		if setArgNil {
			obj["always-compare-med"] = nil
		} else {
			t, err := expandRouterBgpAlwaysCompareMed(d, v, "always_compare_med", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["always-compare-med"] = t
			}
		}
	}

	if v, ok := d.GetOk("bestpath_as_path_ignore"); ok {
		if setArgNil {
			obj["bestpath-as-path-ignore"] = nil
		} else {
			t, err := expandRouterBgpBestpathAsPathIgnore(d, v, "bestpath_as_path_ignore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bestpath-as-path-ignore"] = t
			}
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_confed_aspath"); ok {
		if setArgNil {
			obj["bestpath-cmp-confed-aspath"] = nil
		} else {
			t, err := expandRouterBgpBestpathCmpConfedAspath(d, v, "bestpath_cmp_confed_aspath", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bestpath-cmp-confed-aspath"] = t
			}
		}
	}

	if v, ok := d.GetOk("bestpath_cmp_routerid"); ok {
		if setArgNil {
			obj["bestpath-cmp-routerid"] = nil
		} else {
			t, err := expandRouterBgpBestpathCmpRouterid(d, v, "bestpath_cmp_routerid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bestpath-cmp-routerid"] = t
			}
		}
	}

	if v, ok := d.GetOk("bestpath_med_confed"); ok {
		if setArgNil {
			obj["bestpath-med-confed"] = nil
		} else {
			t, err := expandRouterBgpBestpathMedConfed(d, v, "bestpath_med_confed", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bestpath-med-confed"] = t
			}
		}
	}

	if v, ok := d.GetOk("bestpath_med_missing_as_worst"); ok {
		if setArgNil {
			obj["bestpath-med-missing-as-worst"] = nil
		} else {
			t, err := expandRouterBgpBestpathMedMissingAsWorst(d, v, "bestpath_med_missing_as_worst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bestpath-med-missing-as-worst"] = t
			}
		}
	}

	if v, ok := d.GetOk("client_to_client_reflection"); ok {
		if setArgNil {
			obj["client-to-client-reflection"] = nil
		} else {
			t, err := expandRouterBgpClientToClientReflection(d, v, "client_to_client_reflection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["client-to-client-reflection"] = t
			}
		}
	}

	if v, ok := d.GetOk("dampening"); ok {
		if setArgNil {
			obj["dampening"] = nil
		} else {
			t, err := expandRouterBgpDampening(d, v, "dampening", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening"] = t
			}
		}
	}

	if v, ok := d.GetOk("deterministic_med"); ok {
		if setArgNil {
			obj["deterministic-med"] = nil
		} else {
			t, err := expandRouterBgpDeterministicMed(d, v, "deterministic_med", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["deterministic-med"] = t
			}
		}
	}

	if v, ok := d.GetOk("ebgp_multipath"); ok {
		if setArgNil {
			obj["ebgp-multipath"] = nil
		} else {
			t, err := expandRouterBgpEbgpMultipath(d, v, "ebgp_multipath", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ebgp-multipath"] = t
			}
		}
	}

	if v, ok := d.GetOk("ibgp_multipath"); ok {
		if setArgNil {
			obj["ibgp-multipath"] = nil
		} else {
			t, err := expandRouterBgpIbgpMultipath(d, v, "ibgp_multipath", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ibgp-multipath"] = t
			}
		}
	}

	if v, ok := d.GetOk("enforce_first_as"); ok {
		if setArgNil {
			obj["enforce-first-as"] = nil
		} else {
			t, err := expandRouterBgpEnforceFirstAs(d, v, "enforce_first_as", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enforce-first-as"] = t
			}
		}
	}

	if v, ok := d.GetOk("fast_external_failover"); ok {
		if setArgNil {
			obj["fast-external-failover"] = nil
		} else {
			t, err := expandRouterBgpFastExternalFailover(d, v, "fast_external_failover", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fast-external-failover"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		if setArgNil {
			obj["log-neighbour-changes"] = nil
		} else {
			t, err := expandRouterBgpLogNeighbourChanges(d, v, "log_neighbour_changes", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-neighbour-changes"] = t
			}
		}
	}

	if v, ok := d.GetOk("network_import_check"); ok {
		if setArgNil {
			obj["network-import-check"] = nil
		} else {
			t, err := expandRouterBgpNetworkImportCheck(d, v, "network_import_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network-import-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("ignore_optional_capability"); ok {
		if setArgNil {
			obj["ignore-optional-capability"] = nil
		} else {
			t, err := expandRouterBgpIgnoreOptionalCapability(d, v, "ignore_optional_capability", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ignore-optional-capability"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path"); ok {
		if setArgNil {
			obj["additional-path"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPath(d, v, "additional_path", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path6"); ok {
		if setArgNil {
			obj["additional-path6"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPath6(d, v, "additional_path6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path6"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv4"); ok {
		if setArgNil {
			obj["additional-path-vpnv4"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPathVpnv4(d, v, "additional_path_vpnv4", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path-vpnv4"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path_vpnv6"); ok {
		if setArgNil {
			obj["additional-path-vpnv6"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPathVpnv6(d, v, "additional_path_vpnv6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path-vpnv6"] = t
			}
		}
	}

	if v, ok := d.GetOk("multipath_recursive_distance"); ok {
		if setArgNil {
			obj["multipath-recursive-distance"] = nil
		} else {
			t, err := expandRouterBgpMultipathRecursiveDistance(d, v, "multipath_recursive_distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multipath-recursive-distance"] = t
			}
		}
	}

	if v, ok := d.GetOk("recursive_next_hop"); ok {
		if setArgNil {
			obj["recursive-next-hop"] = nil
		} else {
			t, err := expandRouterBgpRecursiveNextHop(d, v, "recursive_next_hop", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["recursive-next-hop"] = t
			}
		}
	}

	if v, ok := d.GetOk("recursive_inherit_priority"); ok {
		if setArgNil {
			obj["recursive-inherit-priority"] = nil
		} else {
			t, err := expandRouterBgpRecursiveInheritPriority(d, v, "recursive_inherit_priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["recursive-inherit-priority"] = t
			}
		}
	}

	if v, ok := d.GetOk("tag_resolve_mode"); ok {
		if setArgNil {
			obj["tag-resolve-mode"] = nil
		} else {
			t, err := expandRouterBgpTagResolveMode(d, v, "tag_resolve_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tag-resolve-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		if setArgNil {
			obj["cluster-id"] = nil
		} else {
			t, err := expandRouterBgpClusterId(d, v, "cluster_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cluster-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("confederation_identifier"); ok {
		if setArgNil {
			obj["confederation-identifier"] = nil
		} else {
			t, err := expandRouterBgpConfederationIdentifier(d, v, "confederation_identifier", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["confederation-identifier"] = t
			}
		}
	} else if d.HasChange("confederation_identifier") {
		obj["confederation-identifier"] = nil
	}

	if v, ok := d.GetOk("confederation_peers"); ok || d.HasChange("confederation_peers") {
		if setArgNil {
			obj["confederation-peers"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpConfederationPeers(d, v, "confederation_peers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["confederation-peers"] = t
			}
		}
	}

	if v, ok := d.GetOk("dampening_route_map"); ok {
		if setArgNil {
			obj["dampening-route-map"] = nil
		} else {
			t, err := expandRouterBgpDampeningRouteMap(d, v, "dampening_route_map", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening-route-map"] = t
			}
		}
	} else if d.HasChange("dampening_route_map") {
		obj["dampening-route-map"] = nil
	}

	if v, ok := d.GetOk("dampening_reachability_half_life"); ok {
		if setArgNil {
			obj["dampening-reachability-half-life"] = nil
		} else {
			t, err := expandRouterBgpDampeningReachabilityHalfLife(d, v, "dampening_reachability_half_life", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening-reachability-half-life"] = t
			}
		}
	}

	if v, ok := d.GetOk("dampening_reuse"); ok {
		if setArgNil {
			obj["dampening-reuse"] = nil
		} else {
			t, err := expandRouterBgpDampeningReuse(d, v, "dampening_reuse", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening-reuse"] = t
			}
		}
	}

	if v, ok := d.GetOk("dampening_suppress"); ok {
		if setArgNil {
			obj["dampening-suppress"] = nil
		} else {
			t, err := expandRouterBgpDampeningSuppress(d, v, "dampening_suppress", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening-suppress"] = t
			}
		}
	}

	if v, ok := d.GetOk("dampening_max_suppress_time"); ok {
		if setArgNil {
			obj["dampening-max-suppress-time"] = nil
		} else {
			t, err := expandRouterBgpDampeningMaxSuppressTime(d, v, "dampening_max_suppress_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening-max-suppress-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("dampening_unreachability_half_life"); ok {
		if setArgNil {
			obj["dampening-unreachability-half-life"] = nil
		} else {
			t, err := expandRouterBgpDampeningUnreachabilityHalfLife(d, v, "dampening_unreachability_half_life", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dampening-unreachability-half-life"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("default_local_preference"); ok {
		if setArgNil {
			obj["default-local-preference"] = nil
		} else {
			t, err := expandRouterBgpDefaultLocalPreference(d, v, "default_local_preference", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-local-preference"] = t
			}
		}
	}

	if v, ok := d.GetOk("scan_time"); ok {
		if setArgNil {
			obj["scan-time"] = nil
		} else {
			t, err := expandRouterBgpScanTime(d, v, "scan_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["scan-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance_external"); ok {
		if setArgNil {
			obj["distance-external"] = nil
		} else {
			t, err := expandRouterBgpDistanceExternal(d, v, "distance_external", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance-external"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance_internal"); ok {
		if setArgNil {
			obj["distance-internal"] = nil
		} else {
			t, err := expandRouterBgpDistanceInternal(d, v, "distance_internal", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance-internal"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance_local"); ok {
		if setArgNil {
			obj["distance-local"] = nil
		} else {
			t, err := expandRouterBgpDistanceLocal(d, v, "distance_local", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance-local"] = t
			}
		}
	}

	if v, ok := d.GetOk("synchronization"); ok {
		if setArgNil {
			obj["synchronization"] = nil
		} else {
			t, err := expandRouterBgpSynchronization(d, v, "synchronization", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["synchronization"] = t
			}
		}
	}

	if v, ok := d.GetOk("graceful_restart"); ok {
		if setArgNil {
			obj["graceful-restart"] = nil
		} else {
			t, err := expandRouterBgpGracefulRestart(d, v, "graceful_restart", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["graceful-restart"] = t
			}
		}
	}

	if v, ok := d.GetOk("graceful_restart_time"); ok {
		if setArgNil {
			obj["graceful-restart-time"] = nil
		} else {
			t, err := expandRouterBgpGracefulRestartTime(d, v, "graceful_restart_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["graceful-restart-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("graceful_stalepath_time"); ok {
		if setArgNil {
			obj["graceful-stalepath-time"] = nil
		} else {
			t, err := expandRouterBgpGracefulStalepathTime(d, v, "graceful_stalepath_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["graceful-stalepath-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("graceful_update_delay"); ok {
		if setArgNil {
			obj["graceful-update-delay"] = nil
		} else {
			t, err := expandRouterBgpGracefulUpdateDelay(d, v, "graceful_update_delay", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["graceful-update-delay"] = t
			}
		}
	}

	if v, ok := d.GetOk("graceful_end_on_timer"); ok {
		if setArgNil {
			obj["graceful-end-on-timer"] = nil
		} else {
			t, err := expandRouterBgpGracefulEndOnTimer(d, v, "graceful_end_on_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["graceful-end-on-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path_select"); ok {
		if setArgNil {
			obj["additional-path-select"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPathSelect(d, v, "additional_path_select", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path-select"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path_select6"); ok {
		if setArgNil {
			obj["additional-path-select6"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPathSelect6(d, v, "additional_path_select6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path-select6"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path_select_vpnv4"); ok {
		if setArgNil {
			obj["additional-path-select-vpnv4"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPathSelectVpnv4(d, v, "additional_path_select_vpnv4", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path-select-vpnv4"] = t
			}
		}
	}

	if v, ok := d.GetOk("additional_path_select_vpnv6"); ok {
		if setArgNil {
			obj["additional-path-select-vpnv6"] = nil
		} else {
			t, err := expandRouterBgpAdditionalPathSelectVpnv6(d, v, "additional_path_select_vpnv6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["additional-path-select-vpnv6"] = t
			}
		}
	}

	if v, ok := d.GetOk("cross_family_conditional_adv"); ok {
		if setArgNil {
			obj["cross-family-conditional-adv"] = nil
		} else {
			t, err := expandRouterBgpCrossFamilyConditionalAdv(d, v, "cross_family_conditional_adv", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cross-family-conditional-adv"] = t
			}
		}
	}

	if v, ok := d.GetOk("aggregate_address"); ok || d.HasChange("aggregate_address") {
		if setArgNil {
			obj["aggregate-address"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpAggregateAddress(d, v, "aggregate_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["aggregate-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("aggregate_address6"); ok || d.HasChange("aggregate_address6") {
		if setArgNil {
			obj["aggregate-address6"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpAggregateAddress6(d, v, "aggregate_address6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["aggregate-address6"] = t
			}
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		if setArgNil {
			obj["neighbor"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpNeighbor(d, v, "neighbor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	if v, ok := d.GetOk("neighbor_group"); ok || d.HasChange("neighbor_group") {
		if setArgNil {
			obj["neighbor-group"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpNeighborGroup(d, v, "neighbor_group", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor-group"] = t
			}
		}
	}

	if v, ok := d.GetOk("neighbor_range"); ok || d.HasChange("neighbor_range") {
		if setArgNil {
			obj["neighbor-range"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpNeighborRange(d, v, "neighbor_range", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor-range"] = t
			}
		}
	}

	if v, ok := d.GetOk("neighbor_range6"); ok || d.HasChange("neighbor_range6") {
		if setArgNil {
			obj["neighbor-range6"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpNeighborRange6(d, v, "neighbor_range6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor-range6"] = t
			}
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		if setArgNil {
			obj["network"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpNetwork(d, v, "network", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network"] = t
			}
		}
	}

	if v, ok := d.GetOk("network6"); ok || d.HasChange("network6") {
		if setArgNil {
			obj["network6"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpNetwork6(d, v, "network6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network6"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		if setArgNil {
			obj["redistribute"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpRedistribute(d, v, "redistribute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute6"); ok || d.HasChange("redistribute6") {
		if setArgNil {
			obj["redistribute6"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpRedistribute6(d, v, "redistribute6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute6"] = t
			}
		}
	}

	if v, ok := d.GetOk("admin_distance"); ok || d.HasChange("admin_distance") {
		if setArgNil {
			obj["admin-distance"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpAdminDistance(d, v, "admin_distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["admin-distance"] = t
			}
		}
	}

	if v, ok := d.GetOk("vrf"); ok || d.HasChange("vrf") {
		if setArgNil {
			obj["vrf"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpVrf(d, v, "vrf", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf"] = t
			}
		}
	}

	if v, ok := d.GetOk("vrf6"); ok || d.HasChange("vrf6") {
		if setArgNil {
			obj["vrf6"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpVrf6(d, v, "vrf6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf6"] = t
			}
		}
	}

	if v, ok := d.GetOk("vrf_leak"); ok || d.HasChange("vrf_leak") {
		if setArgNil {
			obj["vrf-leak"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpVrfLeak(d, v, "vrf_leak", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf-leak"] = t
			}
		}
	}

	if v, ok := d.GetOk("vrf_leak6"); ok || d.HasChange("vrf_leak6") {
		if setArgNil {
			obj["vrf-leak6"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterBgpVrfLeak6(d, v, "vrf_leak6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vrf-leak6"] = t
			}
		}
	}

	return &obj, nil
}
