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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterBgp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBgpRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"as": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"keepalive_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"holdtime_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"always_compare_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_as_path_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_cmp_confed_aspath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_cmp_routerid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_med_confed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bestpath_med_missing_as_worst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_to_client_reflection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dampening": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"deterministic_med": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ebgp_multipath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ibgp_multipath": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enforce_first_as": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fast_external_failover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_import_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_optional_capability": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_path6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"multipath_recursive_distance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"recursive_next_hop": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"confederation_identifier": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"confederation_peers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dampening_route_map": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dampening_reachability_half_life": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_suppress": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_max_suppress_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_unreachability_half_life": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_local_preference": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scan_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_external": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_internal": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_local": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"synchronization": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"graceful_restart": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"graceful_restart_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"graceful_stalepath_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"graceful_update_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"graceful_end_on_timer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_path_select": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"additional_path_select6": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"aggregate_address": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"aggregate_address6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_set": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"advertisement_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allowas_in_enable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in_enable6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allowas_in6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"attribute_unchanged": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_unchanged6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activate6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_orf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_orf6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_graceful_restart": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_graceful_restart6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_route_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_default_originate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_default_originate6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dont_capability_negotiate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_enforce_multihop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link_down_failover": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"stale_route": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self_rr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self_rr6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"override_capability": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_private_as": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_private_as6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_reflector_client": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_reflector_client6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_server_client": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_server_client6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"shutdown": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_reconfiguration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_reconfiguration6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_override6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"strict_capability_match": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_originate_routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_originate_routemap6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_multihop_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"filter_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_threshold6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_warning_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix_warning_only6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_as": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_as": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_as_no_prepend": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_as_replace_as": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"retain_stale_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"route_map_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out_preferable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out6_preferable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_community6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"keep_alive_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"holdtime_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"connect_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"unsuppress_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unsuppress_map6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"restart_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"additional_path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"additional_path6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"adv_additional_path": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"adv_additional_path6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"conditional_advertise": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advertise_routemap": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"condition_routemap": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"condition_type": &schema.Schema{
										Type:     schema.TypeString,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"advertisement_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allowas_in_enable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in_enable6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowas_in": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"allowas_in6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"attribute_unchanged": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attribute_unchanged6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"activate6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_dynamic": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_orf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_orf6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_graceful_restart": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_graceful_restart6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_route_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_default_originate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"capability_default_originate6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dont_capability_negotiate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_enforce_multihop": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link_down_failover": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"stale_route": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self_rr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_hop_self_rr6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"override_capability": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_private_as": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remove_private_as6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_reflector_client": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_reflector_client6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_server_client": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_server_client6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"shutdown": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_reconfiguration": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"soft_reconfiguration6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"as_override6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"strict_capability_match": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_originate_routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_originate_routemap6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distribute_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ebgp_multihop_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"filter_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"filter_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_threshold6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"maximum_prefix_warning_only": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"maximum_prefix_warning_only6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_as": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_as": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"local_as_no_prepend": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"local_as_replace_as": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"retain_stale_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"route_map_in": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_in6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out_preferable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map_out6_preferable": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_community": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_community6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"keep_alive_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"holdtime_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"connect_timer": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"unsuppress_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unsuppress_map6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"restart_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"additional_path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"additional_path6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"adv_additional_path": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"adv_additional_path6": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"neighbor_range": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_neighbor_num": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"neighbor_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"neighbor_range6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"max_neighbor_num": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"neighbor_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"network6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"backdoor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"redistribute6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_map": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"admin_distance": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"neighbour_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"vrf_leak": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrf": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"route_map": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"interface": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterBgpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterBgp"

	o, err := c.ReadRouterBgp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterBgp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBgp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterBgpAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpKeepaliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAlwaysCompareMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathAsPathIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathCmpConfedAspath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathCmpRouterid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathMedConfed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpBestpathMedMissingAsWorst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpClientToClientReflection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampening(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDeterministicMed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpEbgpMultipath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpIbgpMultipath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpEnforceFirstAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpFastExternalFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetworkImportCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpIgnoreOptionalCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpMultipathRecursiveDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRecursiveNextHop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpClusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpConfederationIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpConfederationPeers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = dataSourceFlattenRouterBgpConfederationPeersPeer(i["peer"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpConfederationPeersPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningReachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningReuse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningSuppress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningMaxSuppressTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDampeningUnreachabilityHalfLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDefaultLocalPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpScanTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDistanceInternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpDistanceLocal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpSynchronization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpGracefulRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpGracefulStalepathTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpGracefulUpdateDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpGracefulEndOnTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdditionalPathSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdditionalPathSelect6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpAggregateAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterBgpAggregateAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {
			tmp["as_set"] = dataSourceFlattenRouterBgpAggregateAddressAsSet(i["as-set"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {
			tmp["summary_only"] = dataSourceFlattenRouterBgpAggregateAddressSummaryOnly(i["summary-only"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAggregateAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpAggregateAddressAsSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddressSummaryOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpAggregateAddress6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterBgpAggregateAddress6Prefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_set"
		if _, ok := i["as-set"]; ok {
			tmp["as_set"] = dataSourceFlattenRouterBgpAggregateAddress6AsSet(i["as-set"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "summary_only"
		if _, ok := i["summary-only"]; ok {
			tmp["summary_only"] = dataSourceFlattenRouterBgpAggregateAddress6SummaryOnly(i["summary-only"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAggregateAddress6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6AsSet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAggregateAddress6SummaryOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterBgpNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {
			tmp["advertisement_interval"] = dataSourceFlattenRouterBgpNeighborAdvertisementInterval(i["advertisement-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {
			tmp["allowas_in_enable"] = dataSourceFlattenRouterBgpNeighborAllowasInEnable(i["allowas-in-enable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {
			tmp["allowas_in_enable6"] = dataSourceFlattenRouterBgpNeighborAllowasInEnable6(i["allowas-in-enable6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {
			tmp["allowas_in"] = dataSourceFlattenRouterBgpNeighborAllowasIn(i["allowas-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {
			tmp["allowas_in6"] = dataSourceFlattenRouterBgpNeighborAllowasIn6(i["allowas-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {
			tmp["attribute_unchanged"] = dataSourceFlattenRouterBgpNeighborAttributeUnchanged(i["attribute-unchanged"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {
			tmp["attribute_unchanged6"] = dataSourceFlattenRouterBgpNeighborAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {
			tmp["activate"] = dataSourceFlattenRouterBgpNeighborActivate(i["activate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {
			tmp["activate6"] = dataSourceFlattenRouterBgpNeighborActivate6(i["activate6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterBgpNeighborBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {
			tmp["capability_dynamic"] = dataSourceFlattenRouterBgpNeighborCapabilityDynamic(i["capability-dynamic"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {
			tmp["capability_orf"] = dataSourceFlattenRouterBgpNeighborCapabilityOrf(i["capability-orf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {
			tmp["capability_orf6"] = dataSourceFlattenRouterBgpNeighborCapabilityOrf6(i["capability-orf6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := i["capability-graceful-restart"]; ok {
			tmp["capability_graceful_restart"] = dataSourceFlattenRouterBgpNeighborCapabilityGracefulRestart(i["capability-graceful-restart"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := i["capability-graceful-restart6"]; ok {
			tmp["capability_graceful_restart6"] = dataSourceFlattenRouterBgpNeighborCapabilityGracefulRestart6(i["capability-graceful-restart6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := i["capability-route-refresh"]; ok {
			tmp["capability_route_refresh"] = dataSourceFlattenRouterBgpNeighborCapabilityRouteRefresh(i["capability-route-refresh"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {
			tmp["capability_default_originate"] = dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {
			tmp["capability_default_originate6"] = dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {
			tmp["dont_capability_negotiate"] = dataSourceFlattenRouterBgpNeighborDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {
			tmp["ebgp_enforce_multihop"] = dataSourceFlattenRouterBgpNeighborEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := i["link-down-failover"]; ok {
			tmp["link_down_failover"] = dataSourceFlattenRouterBgpNeighborLinkDownFailover(i["link-down-failover"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := i["stale-route"]; ok {
			tmp["stale_route"] = dataSourceFlattenRouterBgpNeighborStaleRoute(i["stale-route"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {
			tmp["next_hop_self"] = dataSourceFlattenRouterBgpNeighborNextHopSelf(i["next-hop-self"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {
			tmp["next_hop_self6"] = dataSourceFlattenRouterBgpNeighborNextHopSelf6(i["next-hop-self6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := i["next-hop-self-rr"]; ok {
			tmp["next_hop_self_rr"] = dataSourceFlattenRouterBgpNeighborNextHopSelfRr(i["next-hop-self-rr"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := i["next-hop-self-rr6"]; ok {
			tmp["next_hop_self_rr6"] = dataSourceFlattenRouterBgpNeighborNextHopSelfRr6(i["next-hop-self-rr6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {
			tmp["override_capability"] = dataSourceFlattenRouterBgpNeighborOverrideCapability(i["override-capability"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			tmp["passive"] = dataSourceFlattenRouterBgpNeighborPassive(i["passive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {
			tmp["remove_private_as"] = dataSourceFlattenRouterBgpNeighborRemovePrivateAs(i["remove-private-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {
			tmp["remove_private_as6"] = dataSourceFlattenRouterBgpNeighborRemovePrivateAs6(i["remove-private-as6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {
			tmp["route_reflector_client"] = dataSourceFlattenRouterBgpNeighborRouteReflectorClient(i["route-reflector-client"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {
			tmp["route_reflector_client6"] = dataSourceFlattenRouterBgpNeighborRouteReflectorClient6(i["route-reflector-client6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {
			tmp["route_server_client"] = dataSourceFlattenRouterBgpNeighborRouteServerClient(i["route-server-client"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {
			tmp["route_server_client6"] = dataSourceFlattenRouterBgpNeighborRouteServerClient6(i["route-server-client6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {
			tmp["shutdown"] = dataSourceFlattenRouterBgpNeighborShutdown(i["shutdown"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {
			tmp["soft_reconfiguration"] = dataSourceFlattenRouterBgpNeighborSoftReconfiguration(i["soft-reconfiguration"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {
			tmp["soft_reconfiguration6"] = dataSourceFlattenRouterBgpNeighborSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {
			tmp["as_override"] = dataSourceFlattenRouterBgpNeighborAsOverride(i["as-override"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {
			tmp["as_override6"] = dataSourceFlattenRouterBgpNeighborAsOverride6(i["as-override6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {
			tmp["strict_capability_match"] = dataSourceFlattenRouterBgpNeighborStrictCapabilityMatch(i["strict-capability-match"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {
			tmp["default_originate_routemap"] = dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {
			tmp["default_originate_routemap6"] = dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = dataSourceFlattenRouterBgpNeighborDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {
			tmp["distribute_list_in"] = dataSourceFlattenRouterBgpNeighborDistributeListIn(i["distribute-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {
			tmp["distribute_list_in6"] = dataSourceFlattenRouterBgpNeighborDistributeListIn6(i["distribute-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {
			tmp["distribute_list_out"] = dataSourceFlattenRouterBgpNeighborDistributeListOut(i["distribute-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {
			tmp["distribute_list_out6"] = dataSourceFlattenRouterBgpNeighborDistributeListOut6(i["distribute-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {
			tmp["ebgp_multihop_ttl"] = dataSourceFlattenRouterBgpNeighborEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {
			tmp["filter_list_in"] = dataSourceFlattenRouterBgpNeighborFilterListIn(i["filter-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {
			tmp["filter_list_in6"] = dataSourceFlattenRouterBgpNeighborFilterListIn6(i["filter-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {
			tmp["filter_list_out"] = dataSourceFlattenRouterBgpNeighborFilterListOut(i["filter-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {
			tmp["filter_list_out6"] = dataSourceFlattenRouterBgpNeighborFilterListOut6(i["filter-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBgpNeighborInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {
			tmp["maximum_prefix"] = dataSourceFlattenRouterBgpNeighborMaximumPrefix(i["maximum-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {
			tmp["maximum_prefix6"] = dataSourceFlattenRouterBgpNeighborMaximumPrefix6(i["maximum-prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {
			tmp["maximum_prefix_threshold"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {
			tmp["maximum_prefix_threshold6"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {
			tmp["maximum_prefix_warning_only"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {
			tmp["maximum_prefix_warning_only6"] = dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {
			tmp["prefix_list_in"] = dataSourceFlattenRouterBgpNeighborPrefixListIn(i["prefix-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {
			tmp["prefix_list_in6"] = dataSourceFlattenRouterBgpNeighborPrefixListIn6(i["prefix-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {
			tmp["prefix_list_out"] = dataSourceFlattenRouterBgpNeighborPrefixListOut(i["prefix-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {
			tmp["prefix_list_out6"] = dataSourceFlattenRouterBgpNeighborPrefixListOut6(i["prefix-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {
			tmp["remote_as"] = dataSourceFlattenRouterBgpNeighborRemoteAs(i["remote-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := i["local-as"]; ok {
			tmp["local_as"] = dataSourceFlattenRouterBgpNeighborLocalAs(i["local-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := i["local-as-no-prepend"]; ok {
			tmp["local_as_no_prepend"] = dataSourceFlattenRouterBgpNeighborLocalAsNoPrepend(i["local-as-no-prepend"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := i["local-as-replace-as"]; ok {
			tmp["local_as_replace_as"] = dataSourceFlattenRouterBgpNeighborLocalAsReplaceAs(i["local-as-replace-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := i["retain-stale-time"]; ok {
			tmp["retain_stale_time"] = dataSourceFlattenRouterBgpNeighborRetainStaleTime(i["retain-stale-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {
			tmp["route_map_in"] = dataSourceFlattenRouterBgpNeighborRouteMapIn(i["route-map-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {
			tmp["route_map_in6"] = dataSourceFlattenRouterBgpNeighborRouteMapIn6(i["route-map-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {
			tmp["route_map_out"] = dataSourceFlattenRouterBgpNeighborRouteMapOut(i["route-map-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := i["route-map-out-preferable"]; ok {
			tmp["route_map_out_preferable"] = dataSourceFlattenRouterBgpNeighborRouteMapOutPreferable(i["route-map-out-preferable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {
			tmp["route_map_out6"] = dataSourceFlattenRouterBgpNeighborRouteMapOut6(i["route-map-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := i["route-map-out6-preferable"]; ok {
			tmp["route_map_out6_preferable"] = dataSourceFlattenRouterBgpNeighborRouteMapOut6Preferable(i["route-map-out6-preferable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {
			tmp["send_community"] = dataSourceFlattenRouterBgpNeighborSendCommunity(i["send-community"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {
			tmp["send_community6"] = dataSourceFlattenRouterBgpNeighborSendCommunity6(i["send-community6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {
			tmp["keep_alive_timer"] = dataSourceFlattenRouterBgpNeighborKeepAliveTimer(i["keep-alive-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {
			tmp["holdtime_timer"] = dataSourceFlattenRouterBgpNeighborHoldtimeTimer(i["holdtime-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {
			tmp["connect_timer"] = dataSourceFlattenRouterBgpNeighborConnectTimer(i["connect-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {
			tmp["unsuppress_map"] = dataSourceFlattenRouterBgpNeighborUnsuppressMap(i["unsuppress-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {
			tmp["unsuppress_map6"] = dataSourceFlattenRouterBgpNeighborUnsuppressMap6(i["unsuppress-map6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {
			tmp["update_source"] = dataSourceFlattenRouterBgpNeighborUpdateSource(i["update-source"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			tmp["weight"] = dataSourceFlattenRouterBgpNeighborWeight(i["weight"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := i["restart-time"]; ok {
			tmp["restart_time"] = dataSourceFlattenRouterBgpNeighborRestartTime(i["restart-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := i["additional-path"]; ok {
			tmp["additional_path"] = dataSourceFlattenRouterBgpNeighborAdditionalPath(i["additional-path"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := i["additional-path6"]; ok {
			tmp["additional_path6"] = dataSourceFlattenRouterBgpNeighborAdditionalPath6(i["additional-path6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := i["adv-additional-path"]; ok {
			tmp["adv_additional_path"] = dataSourceFlattenRouterBgpNeighborAdvAdditionalPath(i["adv-additional-path"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := i["adv-additional-path6"]; ok {
			tmp["adv_additional_path6"] = dataSourceFlattenRouterBgpNeighborAdvAdditionalPath6(i["adv-additional-path6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			tmp["password"] = dataSourceFlattenRouterBgpNeighborPassword(i["password"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "conditional_advertise"
		if _, ok := i["conditional-advertise"]; ok {
			tmp["conditional_advertise"] = dataSourceFlattenRouterBgpNeighborConditionalAdvertise(i["conditional-advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborLinkDownFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborStaleRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborLocalAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRetainStaleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborConditionalAdvertise(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise_routemap"
		if _, ok := i["advertise-routemap"]; ok {
			tmp["advertise_routemap"] = dataSourceFlattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(i["advertise-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_routemap"
		if _, ok := i["condition-routemap"]; ok {
			tmp["condition_routemap"] = dataSourceFlattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(i["condition-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "condition_type"
		if _, ok := i["condition-type"]; ok {
			tmp["condition_type"] = dataSourceFlattenRouterBgpNeighborConditionalAdvertiseConditionType(i["condition-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNeighborConditionalAdvertiseAdvertiseRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborConditionalAdvertiseConditionRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborConditionalAdvertiseConditionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterBgpNeighborGroupName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertisement_interval"
		if _, ok := i["advertisement-interval"]; ok {
			tmp["advertisement_interval"] = dataSourceFlattenRouterBgpNeighborGroupAdvertisementInterval(i["advertisement-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable"
		if _, ok := i["allowas-in-enable"]; ok {
			tmp["allowas_in_enable"] = dataSourceFlattenRouterBgpNeighborGroupAllowasInEnable(i["allowas-in-enable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in_enable6"
		if _, ok := i["allowas-in-enable6"]; ok {
			tmp["allowas_in_enable6"] = dataSourceFlattenRouterBgpNeighborGroupAllowasInEnable6(i["allowas-in-enable6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in"
		if _, ok := i["allowas-in"]; ok {
			tmp["allowas_in"] = dataSourceFlattenRouterBgpNeighborGroupAllowasIn(i["allowas-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowas_in6"
		if _, ok := i["allowas-in6"]; ok {
			tmp["allowas_in6"] = dataSourceFlattenRouterBgpNeighborGroupAllowasIn6(i["allowas-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged"
		if _, ok := i["attribute-unchanged"]; ok {
			tmp["attribute_unchanged"] = dataSourceFlattenRouterBgpNeighborGroupAttributeUnchanged(i["attribute-unchanged"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_unchanged6"
		if _, ok := i["attribute-unchanged6"]; ok {
			tmp["attribute_unchanged6"] = dataSourceFlattenRouterBgpNeighborGroupAttributeUnchanged6(i["attribute-unchanged6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate"
		if _, ok := i["activate"]; ok {
			tmp["activate"] = dataSourceFlattenRouterBgpNeighborGroupActivate(i["activate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "activate6"
		if _, ok := i["activate6"]; ok {
			tmp["activate6"] = dataSourceFlattenRouterBgpNeighborGroupActivate6(i["activate6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterBgpNeighborGroupBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_dynamic"
		if _, ok := i["capability-dynamic"]; ok {
			tmp["capability_dynamic"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityDynamic(i["capability-dynamic"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf"
		if _, ok := i["capability-orf"]; ok {
			tmp["capability_orf"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityOrf(i["capability-orf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_orf6"
		if _, ok := i["capability-orf6"]; ok {
			tmp["capability_orf6"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityOrf6(i["capability-orf6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart"
		if _, ok := i["capability-graceful-restart"]; ok {
			tmp["capability_graceful_restart"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityGracefulRestart(i["capability-graceful-restart"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_graceful_restart6"
		if _, ok := i["capability-graceful-restart6"]; ok {
			tmp["capability_graceful_restart6"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityGracefulRestart6(i["capability-graceful-restart6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_route_refresh"
		if _, ok := i["capability-route-refresh"]; ok {
			tmp["capability_route_refresh"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityRouteRefresh(i["capability-route-refresh"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate"
		if _, ok := i["capability-default-originate"]; ok {
			tmp["capability_default_originate"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityDefaultOriginate(i["capability-default-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "capability_default_originate6"
		if _, ok := i["capability-default-originate6"]; ok {
			tmp["capability_default_originate6"] = dataSourceFlattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(i["capability-default-originate6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dont_capability_negotiate"
		if _, ok := i["dont-capability-negotiate"]; ok {
			tmp["dont_capability_negotiate"] = dataSourceFlattenRouterBgpNeighborGroupDontCapabilityNegotiate(i["dont-capability-negotiate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_enforce_multihop"
		if _, ok := i["ebgp-enforce-multihop"]; ok {
			tmp["ebgp_enforce_multihop"] = dataSourceFlattenRouterBgpNeighborGroupEbgpEnforceMultihop(i["ebgp-enforce-multihop"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_down_failover"
		if _, ok := i["link-down-failover"]; ok {
			tmp["link_down_failover"] = dataSourceFlattenRouterBgpNeighborGroupLinkDownFailover(i["link-down-failover"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stale_route"
		if _, ok := i["stale-route"]; ok {
			tmp["stale_route"] = dataSourceFlattenRouterBgpNeighborGroupStaleRoute(i["stale-route"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self"
		if _, ok := i["next-hop-self"]; ok {
			tmp["next_hop_self"] = dataSourceFlattenRouterBgpNeighborGroupNextHopSelf(i["next-hop-self"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self6"
		if _, ok := i["next-hop-self6"]; ok {
			tmp["next_hop_self6"] = dataSourceFlattenRouterBgpNeighborGroupNextHopSelf6(i["next-hop-self6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr"
		if _, ok := i["next-hop-self-rr"]; ok {
			tmp["next_hop_self_rr"] = dataSourceFlattenRouterBgpNeighborGroupNextHopSelfRr(i["next-hop-self-rr"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop_self_rr6"
		if _, ok := i["next-hop-self-rr6"]; ok {
			tmp["next_hop_self_rr6"] = dataSourceFlattenRouterBgpNeighborGroupNextHopSelfRr6(i["next-hop-self-rr6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_capability"
		if _, ok := i["override-capability"]; ok {
			tmp["override_capability"] = dataSourceFlattenRouterBgpNeighborGroupOverrideCapability(i["override-capability"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			tmp["passive"] = dataSourceFlattenRouterBgpNeighborGroupPassive(i["passive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as"
		if _, ok := i["remove-private-as"]; ok {
			tmp["remove_private_as"] = dataSourceFlattenRouterBgpNeighborGroupRemovePrivateAs(i["remove-private-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remove_private_as6"
		if _, ok := i["remove-private-as6"]; ok {
			tmp["remove_private_as6"] = dataSourceFlattenRouterBgpNeighborGroupRemovePrivateAs6(i["remove-private-as6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client"
		if _, ok := i["route-reflector-client"]; ok {
			tmp["route_reflector_client"] = dataSourceFlattenRouterBgpNeighborGroupRouteReflectorClient(i["route-reflector-client"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_reflector_client6"
		if _, ok := i["route-reflector-client6"]; ok {
			tmp["route_reflector_client6"] = dataSourceFlattenRouterBgpNeighborGroupRouteReflectorClient6(i["route-reflector-client6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client"
		if _, ok := i["route-server-client"]; ok {
			tmp["route_server_client"] = dataSourceFlattenRouterBgpNeighborGroupRouteServerClient(i["route-server-client"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_server_client6"
		if _, ok := i["route-server-client6"]; ok {
			tmp["route_server_client6"] = dataSourceFlattenRouterBgpNeighborGroupRouteServerClient6(i["route-server-client6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shutdown"
		if _, ok := i["shutdown"]; ok {
			tmp["shutdown"] = dataSourceFlattenRouterBgpNeighborGroupShutdown(i["shutdown"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration"
		if _, ok := i["soft-reconfiguration"]; ok {
			tmp["soft_reconfiguration"] = dataSourceFlattenRouterBgpNeighborGroupSoftReconfiguration(i["soft-reconfiguration"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "soft_reconfiguration6"
		if _, ok := i["soft-reconfiguration6"]; ok {
			tmp["soft_reconfiguration6"] = dataSourceFlattenRouterBgpNeighborGroupSoftReconfiguration6(i["soft-reconfiguration6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override"
		if _, ok := i["as-override"]; ok {
			tmp["as_override"] = dataSourceFlattenRouterBgpNeighborGroupAsOverride(i["as-override"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "as_override6"
		if _, ok := i["as-override6"]; ok {
			tmp["as_override6"] = dataSourceFlattenRouterBgpNeighborGroupAsOverride6(i["as-override6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strict_capability_match"
		if _, ok := i["strict-capability-match"]; ok {
			tmp["strict_capability_match"] = dataSourceFlattenRouterBgpNeighborGroupStrictCapabilityMatch(i["strict-capability-match"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap"
		if _, ok := i["default-originate-routemap"]; ok {
			tmp["default_originate_routemap"] = dataSourceFlattenRouterBgpNeighborGroupDefaultOriginateRoutemap(i["default-originate-routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_originate_routemap6"
		if _, ok := i["default-originate-routemap6"]; ok {
			tmp["default_originate_routemap6"] = dataSourceFlattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(i["default-originate-routemap6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = dataSourceFlattenRouterBgpNeighborGroupDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in"
		if _, ok := i["distribute-list-in"]; ok {
			tmp["distribute_list_in"] = dataSourceFlattenRouterBgpNeighborGroupDistributeListIn(i["distribute-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_in6"
		if _, ok := i["distribute-list-in6"]; ok {
			tmp["distribute_list_in6"] = dataSourceFlattenRouterBgpNeighborGroupDistributeListIn6(i["distribute-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out"
		if _, ok := i["distribute-list-out"]; ok {
			tmp["distribute_list_out"] = dataSourceFlattenRouterBgpNeighborGroupDistributeListOut(i["distribute-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distribute_list_out6"
		if _, ok := i["distribute-list-out6"]; ok {
			tmp["distribute_list_out6"] = dataSourceFlattenRouterBgpNeighborGroupDistributeListOut6(i["distribute-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ebgp_multihop_ttl"
		if _, ok := i["ebgp-multihop-ttl"]; ok {
			tmp["ebgp_multihop_ttl"] = dataSourceFlattenRouterBgpNeighborGroupEbgpMultihopTtl(i["ebgp-multihop-ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in"
		if _, ok := i["filter-list-in"]; ok {
			tmp["filter_list_in"] = dataSourceFlattenRouterBgpNeighborGroupFilterListIn(i["filter-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_in6"
		if _, ok := i["filter-list-in6"]; ok {
			tmp["filter_list_in6"] = dataSourceFlattenRouterBgpNeighborGroupFilterListIn6(i["filter-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out"
		if _, ok := i["filter-list-out"]; ok {
			tmp["filter_list_out"] = dataSourceFlattenRouterBgpNeighborGroupFilterListOut(i["filter-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list_out6"
		if _, ok := i["filter-list-out6"]; ok {
			tmp["filter_list_out6"] = dataSourceFlattenRouterBgpNeighborGroupFilterListOut6(i["filter-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBgpNeighborGroupInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix"
		if _, ok := i["maximum-prefix"]; ok {
			tmp["maximum_prefix"] = dataSourceFlattenRouterBgpNeighborGroupMaximumPrefix(i["maximum-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix6"
		if _, ok := i["maximum-prefix6"]; ok {
			tmp["maximum_prefix6"] = dataSourceFlattenRouterBgpNeighborGroupMaximumPrefix6(i["maximum-prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold"
		if _, ok := i["maximum-prefix-threshold"]; ok {
			tmp["maximum_prefix_threshold"] = dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixThreshold(i["maximum-prefix-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_threshold6"
		if _, ok := i["maximum-prefix-threshold6"]; ok {
			tmp["maximum_prefix_threshold6"] = dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixThreshold6(i["maximum-prefix-threshold6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only"
		if _, ok := i["maximum-prefix-warning-only"]; ok {
			tmp["maximum_prefix_warning_only"] = dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(i["maximum-prefix-warning-only"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "maximum_prefix_warning_only6"
		if _, ok := i["maximum-prefix-warning-only6"]; ok {
			tmp["maximum_prefix_warning_only6"] = dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(i["maximum-prefix-warning-only6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in"
		if _, ok := i["prefix-list-in"]; ok {
			tmp["prefix_list_in"] = dataSourceFlattenRouterBgpNeighborGroupPrefixListIn(i["prefix-list-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_in6"
		if _, ok := i["prefix-list-in6"]; ok {
			tmp["prefix_list_in6"] = dataSourceFlattenRouterBgpNeighborGroupPrefixListIn6(i["prefix-list-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out"
		if _, ok := i["prefix-list-out"]; ok {
			tmp["prefix_list_out"] = dataSourceFlattenRouterBgpNeighborGroupPrefixListOut(i["prefix-list-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_list_out6"
		if _, ok := i["prefix-list-out6"]; ok {
			tmp["prefix_list_out6"] = dataSourceFlattenRouterBgpNeighborGroupPrefixListOut6(i["prefix-list-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_as"
		if _, ok := i["remote-as"]; ok {
			tmp["remote_as"] = dataSourceFlattenRouterBgpNeighborGroupRemoteAs(i["remote-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as"
		if _, ok := i["local-as"]; ok {
			tmp["local_as"] = dataSourceFlattenRouterBgpNeighborGroupLocalAs(i["local-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_no_prepend"
		if _, ok := i["local-as-no-prepend"]; ok {
			tmp["local_as_no_prepend"] = dataSourceFlattenRouterBgpNeighborGroupLocalAsNoPrepend(i["local-as-no-prepend"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "local_as_replace_as"
		if _, ok := i["local-as-replace-as"]; ok {
			tmp["local_as_replace_as"] = dataSourceFlattenRouterBgpNeighborGroupLocalAsReplaceAs(i["local-as-replace-as"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retain_stale_time"
		if _, ok := i["retain-stale-time"]; ok {
			tmp["retain_stale_time"] = dataSourceFlattenRouterBgpNeighborGroupRetainStaleTime(i["retain-stale-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in"
		if _, ok := i["route-map-in"]; ok {
			tmp["route_map_in"] = dataSourceFlattenRouterBgpNeighborGroupRouteMapIn(i["route-map-in"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_in6"
		if _, ok := i["route-map-in6"]; ok {
			tmp["route_map_in6"] = dataSourceFlattenRouterBgpNeighborGroupRouteMapIn6(i["route-map-in6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out"
		if _, ok := i["route-map-out"]; ok {
			tmp["route_map_out"] = dataSourceFlattenRouterBgpNeighborGroupRouteMapOut(i["route-map-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out_preferable"
		if _, ok := i["route-map-out-preferable"]; ok {
			tmp["route_map_out_preferable"] = dataSourceFlattenRouterBgpNeighborGroupRouteMapOutPreferable(i["route-map-out-preferable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6"
		if _, ok := i["route-map-out6"]; ok {
			tmp["route_map_out6"] = dataSourceFlattenRouterBgpNeighborGroupRouteMapOut6(i["route-map-out6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map_out6_preferable"
		if _, ok := i["route-map-out6-preferable"]; ok {
			tmp["route_map_out6_preferable"] = dataSourceFlattenRouterBgpNeighborGroupRouteMapOut6Preferable(i["route-map-out6-preferable"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community"
		if _, ok := i["send-community"]; ok {
			tmp["send_community"] = dataSourceFlattenRouterBgpNeighborGroupSendCommunity(i["send-community"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_community6"
		if _, ok := i["send-community6"]; ok {
			tmp["send_community6"] = dataSourceFlattenRouterBgpNeighborGroupSendCommunity6(i["send-community6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keep_alive_timer"
		if _, ok := i["keep-alive-timer"]; ok {
			tmp["keep_alive_timer"] = dataSourceFlattenRouterBgpNeighborGroupKeepAliveTimer(i["keep-alive-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "holdtime_timer"
		if _, ok := i["holdtime-timer"]; ok {
			tmp["holdtime_timer"] = dataSourceFlattenRouterBgpNeighborGroupHoldtimeTimer(i["holdtime-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "connect_timer"
		if _, ok := i["connect-timer"]; ok {
			tmp["connect_timer"] = dataSourceFlattenRouterBgpNeighborGroupConnectTimer(i["connect-timer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map"
		if _, ok := i["unsuppress-map"]; ok {
			tmp["unsuppress_map"] = dataSourceFlattenRouterBgpNeighborGroupUnsuppressMap(i["unsuppress-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "unsuppress_map6"
		if _, ok := i["unsuppress-map6"]; ok {
			tmp["unsuppress_map6"] = dataSourceFlattenRouterBgpNeighborGroupUnsuppressMap6(i["unsuppress-map6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_source"
		if _, ok := i["update-source"]; ok {
			tmp["update_source"] = dataSourceFlattenRouterBgpNeighborGroupUpdateSource(i["update-source"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			tmp["weight"] = dataSourceFlattenRouterBgpNeighborGroupWeight(i["weight"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restart_time"
		if _, ok := i["restart-time"]; ok {
			tmp["restart_time"] = dataSourceFlattenRouterBgpNeighborGroupRestartTime(i["restart-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path"
		if _, ok := i["additional-path"]; ok {
			tmp["additional_path"] = dataSourceFlattenRouterBgpNeighborGroupAdditionalPath(i["additional-path"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_path6"
		if _, ok := i["additional-path6"]; ok {
			tmp["additional_path6"] = dataSourceFlattenRouterBgpNeighborGroupAdditionalPath6(i["additional-path6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path"
		if _, ok := i["adv-additional-path"]; ok {
			tmp["adv_additional_path"] = dataSourceFlattenRouterBgpNeighborGroupAdvAdditionalPath(i["adv-additional-path"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_additional_path6"
		if _, ok := i["adv-additional-path6"]; ok {
			tmp["adv_additional_path6"] = dataSourceFlattenRouterBgpNeighborGroupAdvAdditionalPath6(i["adv-additional-path6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNeighborGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAdvertisementInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAllowasInEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAllowasInEnable6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAllowasIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAllowasIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAttributeUnchanged(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAttributeUnchanged6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupActivate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupActivate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityOrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityOrf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityGracefulRestart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityGracefulRestart6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityRouteRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityDefaultOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupCapabilityDefaultOriginate6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDontCapabilityNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupEbgpEnforceMultihop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupLinkDownFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupStaleRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupNextHopSelf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupNextHopSelf6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupNextHopSelfRr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupNextHopSelfRr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupOverrideCapability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupPassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRemovePrivateAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRemovePrivateAs6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteReflectorClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteReflectorClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteServerClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteServerClient6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupShutdown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupSoftReconfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupSoftReconfiguration6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAsOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAsOverride6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupStrictCapabilityMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDefaultOriginateRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDefaultOriginateRoutemap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDistributeListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDistributeListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupDistributeListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupEbgpMultihopTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupFilterListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupFilterListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupFilterListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupFilterListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupMaximumPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupMaximumPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixThreshold6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixWarningOnly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupMaximumPrefixWarningOnly6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupPrefixListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupPrefixListIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupPrefixListOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupPrefixListOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRemoteAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupLocalAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupLocalAsNoPrepend(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupLocalAsReplaceAs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRetainStaleTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteMapIn6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteMapOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteMapOutPreferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteMapOut6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRouteMapOut6Preferable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupSendCommunity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupSendCommunity6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupKeepAliveTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupHoldtimeTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupConnectTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupUnsuppressMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupUnsuppressMap6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupUpdateSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupRestartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAdvAdditionalPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborGroupAdvAdditionalPath6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpNeighborRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterBgpNeighborRangePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := i["max-neighbor-num"]; ok {
			tmp["max_neighbor_num"] = dataSourceFlattenRouterBgpNeighborRangeMaxNeighborNum(i["max-neighbor-num"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := i["neighbor-group"]; ok {
			tmp["neighbor_group"] = dataSourceFlattenRouterBgpNeighborRangeNeighborGroup(i["neighbor-group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNeighborRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpNeighborRangeMaxNeighborNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRangeNeighborGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRange6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpNeighborRange6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterBgpNeighborRange6Prefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_neighbor_num"
		if _, ok := i["max-neighbor-num"]; ok {
			tmp["max_neighbor_num"] = dataSourceFlattenRouterBgpNeighborRange6MaxNeighborNum(i["max-neighbor-num"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor_group"
		if _, ok := i["neighbor-group"]; ok {
			tmp["neighbor_group"] = dataSourceFlattenRouterBgpNeighborRange6NeighborGroup(i["neighbor-group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNeighborRange6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRange6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRange6MaxNeighborNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNeighborRange6NeighborGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterBgpNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {
			tmp["backdoor"] = dataSourceFlattenRouterBgpNetworkBackdoor(i["backdoor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpNetworkRouteMap(i["route-map"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpNetworkBackdoor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetworkRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpNetwork6Id(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterBgpNetwork6Prefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "backdoor"
		if _, ok := i["backdoor"]; ok {
			tmp["backdoor"] = dataSourceFlattenRouterBgpNetwork6Backdoor(i["backdoor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpNetwork6RouteMap(i["route-map"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpNetwork6Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6Prefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6Backdoor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpNetwork6RouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterBgpRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterBgpRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpRedistributeRouteMap(i["route-map"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistributeRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterBgpRedistribute6Name(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterBgpRedistribute6Status(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpRedistribute6RouteMap(i["route-map"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpRedistribute6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpRedistribute6RouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBgpAdminDistanceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_prefix"
		if _, ok := i["neighbour-prefix"]; ok {
			tmp["neighbour_prefix"] = dataSourceFlattenRouterBgpAdminDistanceNeighbourPrefix(i["neighbour-prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_list"
		if _, ok := i["route-list"]; ok {
			tmp["route_list"] = dataSourceFlattenRouterBgpAdminDistanceRouteList(i["route-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterBgpAdminDistanceDistance(i["distance"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpAdminDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistanceNeighbourPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBgpAdminDistanceRouteList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpAdminDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpVrfLeak(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			tmp["vrf"] = dataSourceFlattenRouterBgpVrfLeakVrf(i["vrf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			tmp["target"] = dataSourceFlattenRouterBgpVrfLeakTarget(i["target"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpVrfLeakVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpVrfLeakTarget(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrf"
		if _, ok := i["vrf"]; ok {
			tmp["vrf"] = dataSourceFlattenRouterBgpVrfLeakTargetVrf(i["vrf"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_map"
		if _, ok := i["route-map"]; ok {
			tmp["route_map"] = dataSourceFlattenRouterBgpVrfLeakTargetRouteMap(i["route-map"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBgpVrfLeakTargetInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBgpVrfLeakTargetVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpVrfLeakTargetRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBgpVrfLeakTargetInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBgp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("as", dataSourceFlattenRouterBgpAs(o["as"], d, "as")); err != nil {
		if !fortiAPIPatch(o["as"]) {
			return fmt.Errorf("Error reading as: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterBgpRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("keepalive_timer", dataSourceFlattenRouterBgpKeepaliveTimer(o["keepalive-timer"], d, "keepalive_timer")); err != nil {
		if !fortiAPIPatch(o["keepalive-timer"]) {
			return fmt.Errorf("Error reading keepalive_timer: %v", err)
		}
	}

	if err = d.Set("holdtime_timer", dataSourceFlattenRouterBgpHoldtimeTimer(o["holdtime-timer"], d, "holdtime_timer")); err != nil {
		if !fortiAPIPatch(o["holdtime-timer"]) {
			return fmt.Errorf("Error reading holdtime_timer: %v", err)
		}
	}

	if err = d.Set("always_compare_med", dataSourceFlattenRouterBgpAlwaysCompareMed(o["always-compare-med"], d, "always_compare_med")); err != nil {
		if !fortiAPIPatch(o["always-compare-med"]) {
			return fmt.Errorf("Error reading always_compare_med: %v", err)
		}
	}

	if err = d.Set("bestpath_as_path_ignore", dataSourceFlattenRouterBgpBestpathAsPathIgnore(o["bestpath-as-path-ignore"], d, "bestpath_as_path_ignore")); err != nil {
		if !fortiAPIPatch(o["bestpath-as-path-ignore"]) {
			return fmt.Errorf("Error reading bestpath_as_path_ignore: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_confed_aspath", dataSourceFlattenRouterBgpBestpathCmpConfedAspath(o["bestpath-cmp-confed-aspath"], d, "bestpath_cmp_confed_aspath")); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-confed-aspath"]) {
			return fmt.Errorf("Error reading bestpath_cmp_confed_aspath: %v", err)
		}
	}

	if err = d.Set("bestpath_cmp_routerid", dataSourceFlattenRouterBgpBestpathCmpRouterid(o["bestpath-cmp-routerid"], d, "bestpath_cmp_routerid")); err != nil {
		if !fortiAPIPatch(o["bestpath-cmp-routerid"]) {
			return fmt.Errorf("Error reading bestpath_cmp_routerid: %v", err)
		}
	}

	if err = d.Set("bestpath_med_confed", dataSourceFlattenRouterBgpBestpathMedConfed(o["bestpath-med-confed"], d, "bestpath_med_confed")); err != nil {
		if !fortiAPIPatch(o["bestpath-med-confed"]) {
			return fmt.Errorf("Error reading bestpath_med_confed: %v", err)
		}
	}

	if err = d.Set("bestpath_med_missing_as_worst", dataSourceFlattenRouterBgpBestpathMedMissingAsWorst(o["bestpath-med-missing-as-worst"], d, "bestpath_med_missing_as_worst")); err != nil {
		if !fortiAPIPatch(o["bestpath-med-missing-as-worst"]) {
			return fmt.Errorf("Error reading bestpath_med_missing_as_worst: %v", err)
		}
	}

	if err = d.Set("client_to_client_reflection", dataSourceFlattenRouterBgpClientToClientReflection(o["client-to-client-reflection"], d, "client_to_client_reflection")); err != nil {
		if !fortiAPIPatch(o["client-to-client-reflection"]) {
			return fmt.Errorf("Error reading client_to_client_reflection: %v", err)
		}
	}

	if err = d.Set("dampening", dataSourceFlattenRouterBgpDampening(o["dampening"], d, "dampening")); err != nil {
		if !fortiAPIPatch(o["dampening"]) {
			return fmt.Errorf("Error reading dampening: %v", err)
		}
	}

	if err = d.Set("deterministic_med", dataSourceFlattenRouterBgpDeterministicMed(o["deterministic-med"], d, "deterministic_med")); err != nil {
		if !fortiAPIPatch(o["deterministic-med"]) {
			return fmt.Errorf("Error reading deterministic_med: %v", err)
		}
	}

	if err = d.Set("ebgp_multipath", dataSourceFlattenRouterBgpEbgpMultipath(o["ebgp-multipath"], d, "ebgp_multipath")); err != nil {
		if !fortiAPIPatch(o["ebgp-multipath"]) {
			return fmt.Errorf("Error reading ebgp_multipath: %v", err)
		}
	}

	if err = d.Set("ibgp_multipath", dataSourceFlattenRouterBgpIbgpMultipath(o["ibgp-multipath"], d, "ibgp_multipath")); err != nil {
		if !fortiAPIPatch(o["ibgp-multipath"]) {
			return fmt.Errorf("Error reading ibgp_multipath: %v", err)
		}
	}

	if err = d.Set("enforce_first_as", dataSourceFlattenRouterBgpEnforceFirstAs(o["enforce-first-as"], d, "enforce_first_as")); err != nil {
		if !fortiAPIPatch(o["enforce-first-as"]) {
			return fmt.Errorf("Error reading enforce_first_as: %v", err)
		}
	}

	if err = d.Set("fast_external_failover", dataSourceFlattenRouterBgpFastExternalFailover(o["fast-external-failover"], d, "fast_external_failover")); err != nil {
		if !fortiAPIPatch(o["fast-external-failover"]) {
			return fmt.Errorf("Error reading fast_external_failover: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterBgpLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("network_import_check", dataSourceFlattenRouterBgpNetworkImportCheck(o["network-import-check"], d, "network_import_check")); err != nil {
		if !fortiAPIPatch(o["network-import-check"]) {
			return fmt.Errorf("Error reading network_import_check: %v", err)
		}
	}

	if err = d.Set("ignore_optional_capability", dataSourceFlattenRouterBgpIgnoreOptionalCapability(o["ignore-optional-capability"], d, "ignore_optional_capability")); err != nil {
		if !fortiAPIPatch(o["ignore-optional-capability"]) {
			return fmt.Errorf("Error reading ignore_optional_capability: %v", err)
		}
	}

	if err = d.Set("additional_path", dataSourceFlattenRouterBgpAdditionalPath(o["additional-path"], d, "additional_path")); err != nil {
		if !fortiAPIPatch(o["additional-path"]) {
			return fmt.Errorf("Error reading additional_path: %v", err)
		}
	}

	if err = d.Set("additional_path6", dataSourceFlattenRouterBgpAdditionalPath6(o["additional-path6"], d, "additional_path6")); err != nil {
		if !fortiAPIPatch(o["additional-path6"]) {
			return fmt.Errorf("Error reading additional_path6: %v", err)
		}
	}

	if err = d.Set("multipath_recursive_distance", dataSourceFlattenRouterBgpMultipathRecursiveDistance(o["multipath-recursive-distance"], d, "multipath_recursive_distance")); err != nil {
		if !fortiAPIPatch(o["multipath-recursive-distance"]) {
			return fmt.Errorf("Error reading multipath_recursive_distance: %v", err)
		}
	}

	if err = d.Set("recursive_next_hop", dataSourceFlattenRouterBgpRecursiveNextHop(o["recursive-next-hop"], d, "recursive_next_hop")); err != nil {
		if !fortiAPIPatch(o["recursive-next-hop"]) {
			return fmt.Errorf("Error reading recursive_next_hop: %v", err)
		}
	}

	if err = d.Set("cluster_id", dataSourceFlattenRouterBgpClusterId(o["cluster-id"], d, "cluster_id")); err != nil {
		if !fortiAPIPatch(o["cluster-id"]) {
			return fmt.Errorf("Error reading cluster_id: %v", err)
		}
	}

	if err = d.Set("confederation_identifier", dataSourceFlattenRouterBgpConfederationIdentifier(o["confederation-identifier"], d, "confederation_identifier")); err != nil {
		if !fortiAPIPatch(o["confederation-identifier"]) {
			return fmt.Errorf("Error reading confederation_identifier: %v", err)
		}
	}

	if err = d.Set("confederation_peers", dataSourceFlattenRouterBgpConfederationPeers(o["confederation-peers"], d, "confederation_peers")); err != nil {
		if !fortiAPIPatch(o["confederation-peers"]) {
			return fmt.Errorf("Error reading confederation_peers: %v", err)
		}
	}

	if err = d.Set("dampening_route_map", dataSourceFlattenRouterBgpDampeningRouteMap(o["dampening-route-map"], d, "dampening_route_map")); err != nil {
		if !fortiAPIPatch(o["dampening-route-map"]) {
			return fmt.Errorf("Error reading dampening_route_map: %v", err)
		}
	}

	if err = d.Set("dampening_reachability_half_life", dataSourceFlattenRouterBgpDampeningReachabilityHalfLife(o["dampening-reachability-half-life"], d, "dampening_reachability_half_life")); err != nil {
		if !fortiAPIPatch(o["dampening-reachability-half-life"]) {
			return fmt.Errorf("Error reading dampening_reachability_half_life: %v", err)
		}
	}

	if err = d.Set("dampening_reuse", dataSourceFlattenRouterBgpDampeningReuse(o["dampening-reuse"], d, "dampening_reuse")); err != nil {
		if !fortiAPIPatch(o["dampening-reuse"]) {
			return fmt.Errorf("Error reading dampening_reuse: %v", err)
		}
	}

	if err = d.Set("dampening_suppress", dataSourceFlattenRouterBgpDampeningSuppress(o["dampening-suppress"], d, "dampening_suppress")); err != nil {
		if !fortiAPIPatch(o["dampening-suppress"]) {
			return fmt.Errorf("Error reading dampening_suppress: %v", err)
		}
	}

	if err = d.Set("dampening_max_suppress_time", dataSourceFlattenRouterBgpDampeningMaxSuppressTime(o["dampening-max-suppress-time"], d, "dampening_max_suppress_time")); err != nil {
		if !fortiAPIPatch(o["dampening-max-suppress-time"]) {
			return fmt.Errorf("Error reading dampening_max_suppress_time: %v", err)
		}
	}

	if err = d.Set("dampening_unreachability_half_life", dataSourceFlattenRouterBgpDampeningUnreachabilityHalfLife(o["dampening-unreachability-half-life"], d, "dampening_unreachability_half_life")); err != nil {
		if !fortiAPIPatch(o["dampening-unreachability-half-life"]) {
			return fmt.Errorf("Error reading dampening_unreachability_half_life: %v", err)
		}
	}

	if err = d.Set("default_local_preference", dataSourceFlattenRouterBgpDefaultLocalPreference(o["default-local-preference"], d, "default_local_preference")); err != nil {
		if !fortiAPIPatch(o["default-local-preference"]) {
			return fmt.Errorf("Error reading default_local_preference: %v", err)
		}
	}

	if err = d.Set("scan_time", dataSourceFlattenRouterBgpScanTime(o["scan-time"], d, "scan_time")); err != nil {
		if !fortiAPIPatch(o["scan-time"]) {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("distance_external", dataSourceFlattenRouterBgpDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_internal", dataSourceFlattenRouterBgpDistanceInternal(o["distance-internal"], d, "distance_internal")); err != nil {
		if !fortiAPIPatch(o["distance-internal"]) {
			return fmt.Errorf("Error reading distance_internal: %v", err)
		}
	}

	if err = d.Set("distance_local", dataSourceFlattenRouterBgpDistanceLocal(o["distance-local"], d, "distance_local")); err != nil {
		if !fortiAPIPatch(o["distance-local"]) {
			return fmt.Errorf("Error reading distance_local: %v", err)
		}
	}

	if err = d.Set("synchronization", dataSourceFlattenRouterBgpSynchronization(o["synchronization"], d, "synchronization")); err != nil {
		if !fortiAPIPatch(o["synchronization"]) {
			return fmt.Errorf("Error reading synchronization: %v", err)
		}
	}

	if err = d.Set("graceful_restart", dataSourceFlattenRouterBgpGracefulRestart(o["graceful-restart"], d, "graceful_restart")); err != nil {
		if !fortiAPIPatch(o["graceful-restart"]) {
			return fmt.Errorf("Error reading graceful_restart: %v", err)
		}
	}

	if err = d.Set("graceful_restart_time", dataSourceFlattenRouterBgpGracefulRestartTime(o["graceful-restart-time"], d, "graceful_restart_time")); err != nil {
		if !fortiAPIPatch(o["graceful-restart-time"]) {
			return fmt.Errorf("Error reading graceful_restart_time: %v", err)
		}
	}

	if err = d.Set("graceful_stalepath_time", dataSourceFlattenRouterBgpGracefulStalepathTime(o["graceful-stalepath-time"], d, "graceful_stalepath_time")); err != nil {
		if !fortiAPIPatch(o["graceful-stalepath-time"]) {
			return fmt.Errorf("Error reading graceful_stalepath_time: %v", err)
		}
	}

	if err = d.Set("graceful_update_delay", dataSourceFlattenRouterBgpGracefulUpdateDelay(o["graceful-update-delay"], d, "graceful_update_delay")); err != nil {
		if !fortiAPIPatch(o["graceful-update-delay"]) {
			return fmt.Errorf("Error reading graceful_update_delay: %v", err)
		}
	}

	if err = d.Set("graceful_end_on_timer", dataSourceFlattenRouterBgpGracefulEndOnTimer(o["graceful-end-on-timer"], d, "graceful_end_on_timer")); err != nil {
		if !fortiAPIPatch(o["graceful-end-on-timer"]) {
			return fmt.Errorf("Error reading graceful_end_on_timer: %v", err)
		}
	}

	if err = d.Set("additional_path_select", dataSourceFlattenRouterBgpAdditionalPathSelect(o["additional-path-select"], d, "additional_path_select")); err != nil {
		if !fortiAPIPatch(o["additional-path-select"]) {
			return fmt.Errorf("Error reading additional_path_select: %v", err)
		}
	}

	if err = d.Set("additional_path_select6", dataSourceFlattenRouterBgpAdditionalPathSelect6(o["additional-path-select6"], d, "additional_path_select6")); err != nil {
		if !fortiAPIPatch(o["additional-path-select6"]) {
			return fmt.Errorf("Error reading additional_path_select6: %v", err)
		}
	}

	if err = d.Set("aggregate_address", dataSourceFlattenRouterBgpAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
		if !fortiAPIPatch(o["aggregate-address"]) {
			return fmt.Errorf("Error reading aggregate_address: %v", err)
		}
	}

	if err = d.Set("aggregate_address6", dataSourceFlattenRouterBgpAggregateAddress6(o["aggregate-address6"], d, "aggregate_address6")); err != nil {
		if !fortiAPIPatch(o["aggregate-address6"]) {
			return fmt.Errorf("Error reading aggregate_address6: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterBgpNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("neighbor_group", dataSourceFlattenRouterBgpNeighborGroup(o["neighbor-group"], d, "neighbor_group")); err != nil {
		if !fortiAPIPatch(o["neighbor-group"]) {
			return fmt.Errorf("Error reading neighbor_group: %v", err)
		}
	}

	if err = d.Set("neighbor_range", dataSourceFlattenRouterBgpNeighborRange(o["neighbor-range"], d, "neighbor_range")); err != nil {
		if !fortiAPIPatch(o["neighbor-range"]) {
			return fmt.Errorf("Error reading neighbor_range: %v", err)
		}
	}

	if err = d.Set("neighbor_range6", dataSourceFlattenRouterBgpNeighborRange6(o["neighbor-range6"], d, "neighbor_range6")); err != nil {
		if !fortiAPIPatch(o["neighbor-range6"]) {
			return fmt.Errorf("Error reading neighbor_range6: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterBgpNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	if err = d.Set("network6", dataSourceFlattenRouterBgpNetwork6(o["network6"], d, "network6")); err != nil {
		if !fortiAPIPatch(o["network6"]) {
			return fmt.Errorf("Error reading network6: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterBgpRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("redistribute6", dataSourceFlattenRouterBgpRedistribute6(o["redistribute6"], d, "redistribute6")); err != nil {
		if !fortiAPIPatch(o["redistribute6"]) {
			return fmt.Errorf("Error reading redistribute6: %v", err)
		}
	}

	if err = d.Set("admin_distance", dataSourceFlattenRouterBgpAdminDistance(o["admin-distance"], d, "admin_distance")); err != nil {
		if !fortiAPIPatch(o["admin-distance"]) {
			return fmt.Errorf("Error reading admin_distance: %v", err)
		}
	}

	if err = d.Set("vrf_leak", dataSourceFlattenRouterBgpVrfLeak(o["vrf-leak"], d, "vrf_leak")); err != nil {
		if !fortiAPIPatch(o["vrf-leak"]) {
			return fmt.Errorf("Error reading vrf_leak: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterBgpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
