// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch devices that are managed by this FortiGate.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerManagedSwitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerManagedSwitchCreate,
		Read:   resourceSwitchControllerManagedSwitchRead,
		Update: resourceSwitchControllerManagedSwitchUpdate,
		Delete: resourceSwitchControllerManagedSwitchDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"switch_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				ForceNew:     true,
				Required:     true,
			},
			"sn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"switch_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"access_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"purdue_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_wan1_peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"fsw_wan1_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_wan2_peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fsw_wan2_admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_pre_standard_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_detection_type": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"poe_lldp_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"directly_connected": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),
				Optional:     true,
				Computed:     true,
			},
			"version": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"max_allowed_trunk_members": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"pre_provisioned": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"l3_discovered": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),
				Optional:     true,
				Computed:     true,
			},
			"mgmt_mode": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"tunnel_discovered": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),
				Optional:     true,
				Computed:     true,
			},
			"tdr_supported": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_capability": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_device_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"switch_dhcp_opt43_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"mclag_igmp_snooping_aware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamically_discovered": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1),
				Optional:     true,
				Computed:     true,
			},
			"ptp_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"radius_nas_ip_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload_mclag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_offload_router": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"router_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"vlan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"assignment_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner_vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"flow_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"staged_image_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"delayed_restart_trigger": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"firmware_provision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port_owner": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"switch_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"speed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"speed_mask": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_source_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ptp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ptp_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"aggregator_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flapguard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flap_rate": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 30),
							Optional:     true,
							Computed:     true,
						},
						"flap_duration": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 300),
							Optional:     true,
							Computed:     true,
						},
						"flap_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 120),
							Optional:     true,
							Computed:     true,
						},
						"rpvst_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_pre_standard_detection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_number": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 64),
							Optional:     true,
							Computed:     true,
						},
						"port_prefix_type": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"fortilink_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"link_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_capable": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"stacking_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"p2p_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"mclag_icl_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"authenticated_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"restricted_auth_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"encrypted_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"fiber_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"media_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"poe_standard": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"poe_max_power": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"poe_mode_bt_cabable": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"poe_port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_port_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"poe_port_power": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"virtual_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"isl_local_trunk_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"isl_peer_port_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"isl_peer_device_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"isl_peer_device_sn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"fgt_peer_port_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"fgt_peer_device_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"allowed_vlans_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowed_vlans": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"untagged_vlans": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"matched_dpp_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"matched_dpp_intf_tags": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"acl_group": &schema.Schema{
							Type:     schema.TypeSet,
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
						"fortiswitch_acls": &schema.Schema{
							Type:     schema.TypeSet,
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
						"dhcp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_snoop_option82_trust": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp_snoop_option82_override": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
									"circuit_id": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 254),
										Optional:     true,
										Computed:     true,
									},
									"remote_id": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 254),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"arp_inspection_trust": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmp_snooping_flood_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mcast_snooping_flood_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmp_snooping": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmps_flood_reports": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"igmps_flood_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_root_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_bpdu_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stp_bpdu_guard_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 120),
							Optional:     true,
							Computed:     true,
						},
						"edge_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"discard_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_sample_rate": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 99999),
							Optional:     true,
							Computed:     true,
						},
						"sflow_sampler": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sflow_sample_rate": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 99999),
							Optional:     true,
							Computed:     true,
						},
						"sflow_counter_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sample_direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fec_capable": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"fec_state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flow_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pause_meter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pause_meter_resume": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"loop_guard": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"loop_guard_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 120),
							Optional:     true,
							Computed:     true,
						},
						"port_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"qos_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"storm_control_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"port_security_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"export_to_pool": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"interface_tags": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"export_tags": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"export_to_pool_flag": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1),
							Optional:     true,
							Computed:     true,
						},
						"learning_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),
							Optional:     true,
							Computed:     true,
						},
						"sticky_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lldp_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lldp_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"export_to": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"mac_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_selection_criteria": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"lacp_speed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bundle": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"member_withdrawal_behavior": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mclag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"min_bundle": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 24),
							Optional:     true,
							Computed:     true,
						},
						"max_bundle": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 24),
							Optional:     true,
							Computed:     true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"ip_source_guard": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"binding_entry": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"entry_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 16),
										Optional:     true,
										Computed:     true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"mac": &schema.Schema{
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
			"stp_settings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"revision": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),
							Optional:     true,
							Computed:     true,
						},
						"forward_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(4, 30),
							Optional:     true,
							Computed:     true,
						},
						"max_age": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(6, 40),
							Optional:     true,
							Computed:     true,
						},
						"max_hops": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 40),
							Optional:     true,
							Computed:     true,
						},
						"pending_timer": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"stp_instance": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"override_snmp_sysinfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_sysinfo": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"engine_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 24),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"contact_info": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"location": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"override_snmp_trap_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_trap_threshold": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trap_high_cpu_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_low_memory_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"trap_log_full_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"override_snmp_community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hosts": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"query_v1_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_v1_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"query_v2c_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_v2c_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"trap_v1_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_v1_lport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"trap_v1_rport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"trap_v2c_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_v2c_lport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"trap_v2c_rport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"events": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"override_snmp_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_user": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"queries": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"security_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_pwd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"priv_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priv_pwd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
					},
				},
			},
			"qos_drop_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_red_probability": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"switch_stp_settings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"switch_log": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"remote_log": &schema.Schema{
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"csv": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"facility": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"storm_control": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"unknown_unicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unknown_multicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"mirror": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"switching_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"src_ingress": &schema.Schema{
							Type:     schema.TypeSet,
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
						"src_egress": &schema.Schema{
							Type:     schema.TypeSet,
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
					},
				},
			},
			"static_mac": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"custom_command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"command_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dhcp_snooping_static_client": &schema.Schema{
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
						"vlan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aging_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(15, 3600),
							Optional:     true,
							Computed:     true,
						},
						"flood_unknown_multicast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlans": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
									"proxy": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"querier": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"querier_addr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(2, 3),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"n802_1x_settings": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"local_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"link_down_auth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reauth_period": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1440),
							Optional:     true,
							Computed:     true,
						},
						"max_reauth_attempt": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"tx_period": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(4, 60),
							Optional:     true,
							Computed:     true,
						},
						"mab_reauth": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_username_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_password_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_calling_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_called_station_delimiter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mac_case": &schema.Schema{
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

func resourceSwitchControllerManagedSwitchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerManagedSwitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitch resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerManagedSwitch(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerManagedSwitch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerManagedSwitch")
	}

	return resourceSwitchControllerManagedSwitchRead(d, m)
}

func resourceSwitchControllerManagedSwitchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerManagedSwitch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitch resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerManagedSwitch(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerManagedSwitch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerManagedSwitch")
	}

	return resourceSwitchControllerManagedSwitchRead(d, m)
}

func resourceSwitchControllerManagedSwitchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerManagedSwitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerManagedSwitch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerManagedSwitchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerManagedSwitch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerManagedSwitch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerManagedSwitch resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerManagedSwitchSwitchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchAccessProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPurdueLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan1Peer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan1Admin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan2Peer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFswWan2Admin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPoePreStandardDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPoeDetectionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPoeLldpDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDirectlyConnected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMaxAllowedTrunkMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPreProvisioned(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchL3Discovered(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMgmtMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchTunnelDiscovered(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchTdrSupported(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDynamicCapability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchDeviceTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchDhcp_Opt43_Key(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMclagIgmpSnoopingAware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDynamicallyDiscovered(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPtpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPtpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRadiusNasIpOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRadiusNasIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadMclag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if cur_v, ok := i["vlan-name"]; ok {
			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchRouteOffloadRouterVlanName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_ip"
		if cur_v, ok := i["router-ip"]; ok {
			tmp["router_ip"] = flattenSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchVlan(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if cur_v, ok := i["vlan-name"]; ok {
			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchVlanVlanName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assignment_priority"
		if cur_v, ok := i["assignment-priority"]; ok {
			tmp["assignment_priority"] = flattenSwitchControllerManagedSwitchVlanAssignmentPriority(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchVlanVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchVlanAssignmentPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOwnerVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFlowIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStagedImageVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDelayedRestartTrigger(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFirmwareProvision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFirmwareProvisionVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if cur_v, ok := i["port-name"]; ok {
			tmp["port_name"] = flattenSwitchControllerManagedSwitchPortsPortName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if cur_v, ok := i["port-owner"]; ok {
			tmp["port_owner"] = flattenSwitchControllerManagedSwitchPortsPortOwner(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if cur_v, ok := i["switch-id"]; ok {
			tmp["switch_id"] = flattenSwitchControllerManagedSwitchPortsSwitchId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed"
		if cur_v, ok := i["speed"]; ok {
			tmp["speed"] = flattenSwitchControllerManagedSwitchPortsSpeed(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed_mask"
		if cur_v, ok := i["speed-mask"]; ok {
			tmp["speed_mask"] = flattenSwitchControllerManagedSwitchPortsSpeedMask(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerManagedSwitchPortsStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if cur_v, ok := i["poe-status"]; ok {
			tmp["poe_status"] = flattenSwitchControllerManagedSwitchPortsPoeStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if cur_v, ok := i["ip-source-guard"]; ok {
			tmp["ip_source_guard"] = flattenSwitchControllerManagedSwitchPortsIpSourceGuard(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_status"
		if cur_v, ok := i["ptp-status"]; ok {
			tmp["ptp_status"] = flattenSwitchControllerManagedSwitchPortsPtpStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_policy"
		if cur_v, ok := i["ptp-policy"]; ok {
			tmp["ptp_policy"] = flattenSwitchControllerManagedSwitchPortsPtpPolicy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if cur_v, ok := i["aggregator-mode"]; ok {
			tmp["aggregator_mode"] = flattenSwitchControllerManagedSwitchPortsAggregatorMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flapguard"
		if cur_v, ok := i["flapguard"]; ok {
			tmp["flapguard"] = flattenSwitchControllerManagedSwitchPortsFlapguard(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_rate"
		if cur_v, ok := i["flap-rate"]; ok {
			tmp["flap_rate"] = flattenSwitchControllerManagedSwitchPortsFlapRate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_duration"
		if cur_v, ok := i["flap-duration"]; ok {
			tmp["flap_duration"] = flattenSwitchControllerManagedSwitchPortsFlapDuration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_timeout"
		if cur_v, ok := i["flap-timeout"]; ok {
			tmp["flap_timeout"] = flattenSwitchControllerManagedSwitchPortsFlapTimeout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if cur_v, ok := i["rpvst-port"]; ok {
			tmp["rpvst_port"] = flattenSwitchControllerManagedSwitchPortsRpvstPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if cur_v, ok := i["poe-pre-standard-detection"]; ok {
			tmp["poe_pre_standard_detection"] = flattenSwitchControllerManagedSwitchPortsPoePreStandardDetection(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_number"
		if cur_v, ok := i["port-number"]; ok {
			tmp["port_number"] = flattenSwitchControllerManagedSwitchPortsPortNumber(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_prefix_type"
		if cur_v, ok := i["port-prefix-type"]; ok {
			tmp["port_prefix_type"] = flattenSwitchControllerManagedSwitchPortsPortPrefixType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink_port"
		if cur_v, ok := i["fortilink-port"]; ok {
			tmp["fortilink_port"] = flattenSwitchControllerManagedSwitchPortsFortilinkPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_status"
		if cur_v, ok := i["link-status"]; ok {
			tmp["link_status"] = flattenSwitchControllerManagedSwitchPortsLinkStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_capable"
		if cur_v, ok := i["poe-capable"]; ok {
			tmp["poe_capable"] = flattenSwitchControllerManagedSwitchPortsPoeCapable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stacking_port"
		if cur_v, ok := i["stacking-port"]; ok {
			tmp["stacking_port"] = flattenSwitchControllerManagedSwitchPortsStackingPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if cur_v, ok := i["p2p-port"]; ok {
			tmp["p2p_port"] = flattenSwitchControllerManagedSwitchPortsP2PPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if cur_v, ok := i["mclag-icl-port"]; ok {
			tmp["mclag_icl_port"] = flattenSwitchControllerManagedSwitchPortsMclagIclPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authenticated_port"
		if cur_v, ok := i["authenticated-port"]; ok {
			tmp["authenticated_port"] = flattenSwitchControllerManagedSwitchPortsAuthenticatedPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_auth_port"
		if cur_v, ok := i["restricted-auth-port"]; ok {
			tmp["restricted_auth_port"] = flattenSwitchControllerManagedSwitchPortsRestrictedAuthPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypted_port"
		if cur_v, ok := i["encrypted-port"]; ok {
			tmp["encrypted_port"] = flattenSwitchControllerManagedSwitchPortsEncryptedPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fiber_port"
		if cur_v, ok := i["fiber-port"]; ok {
			tmp["fiber_port"] = flattenSwitchControllerManagedSwitchPortsFiberPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if cur_v, ok := i["media-type"]; ok {
			tmp["media_type"] = flattenSwitchControllerManagedSwitchPortsMediaType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_standard"
		if cur_v, ok := i["poe-standard"]; ok {
			tmp["poe_standard"] = flattenSwitchControllerManagedSwitchPortsPoeStandard(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_max_power"
		if cur_v, ok := i["poe-max-power"]; ok {
			tmp["poe_max_power"] = flattenSwitchControllerManagedSwitchPortsPoeMaxPower(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_mode_bt_cabable"
		if cur_v, ok := i["poe-mode-bt-cabable"]; ok {
			tmp["poe_mode_bt_cabable"] = flattenSwitchControllerManagedSwitchPortsPoeModeBtCabable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_mode"
		if cur_v, ok := i["poe-port-mode"]; ok {
			tmp["poe_port_mode"] = flattenSwitchControllerManagedSwitchPortsPoePortMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_priority"
		if cur_v, ok := i["poe-port-priority"]; ok {
			tmp["poe_port_priority"] = flattenSwitchControllerManagedSwitchPortsPoePortPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_power"
		if cur_v, ok := i["poe-port-power"]; ok {
			tmp["poe_port_power"] = flattenSwitchControllerManagedSwitchPortsPoePortPower(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if cur_v, ok := i["flags"]; ok {
			tmp["flags"] = flattenSwitchControllerManagedSwitchPortsFlags(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_port"
		if cur_v, ok := i["virtual-port"]; ok {
			tmp["virtual_port"] = flattenSwitchControllerManagedSwitchPortsVirtualPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_local_trunk_name"
		if cur_v, ok := i["isl-local-trunk-name"]; ok {
			tmp["isl_local_trunk_name"] = flattenSwitchControllerManagedSwitchPortsIslLocalTrunkName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_port_name"
		if cur_v, ok := i["isl-peer-port-name"]; ok {
			tmp["isl_peer_port_name"] = flattenSwitchControllerManagedSwitchPortsIslPeerPortName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_name"
		if cur_v, ok := i["isl-peer-device-name"]; ok {
			tmp["isl_peer_device_name"] = flattenSwitchControllerManagedSwitchPortsIslPeerDeviceName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_sn"
		if cur_v, ok := i["isl-peer-device-sn"]; ok {
			tmp["isl_peer_device_sn"] = flattenSwitchControllerManagedSwitchPortsIslPeerDeviceSn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_port_name"
		if cur_v, ok := i["fgt-peer-port-name"]; ok {
			tmp["fgt_peer_port_name"] = flattenSwitchControllerManagedSwitchPortsFgtPeerPortName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_device_name"
		if cur_v, ok := i["fgt-peer-device-name"]; ok {
			tmp["fgt_peer_device_name"] = flattenSwitchControllerManagedSwitchPortsFgtPeerDeviceName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if cur_v, ok := i["vlan"]; ok {
			tmp["vlan"] = flattenSwitchControllerManagedSwitchPortsVlan(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if cur_v, ok := i["allowed-vlans-all"]; ok {
			tmp["allowed_vlans_all"] = flattenSwitchControllerManagedSwitchPortsAllowedVlansAll(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if cur_v, ok := i["allowed-vlans"]; ok {
			tmp["allowed_vlans"] = flattenSwitchControllerManagedSwitchPortsAllowedVlans(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if cur_v, ok := i["untagged-vlans"]; ok {
			tmp["untagged_vlans"] = flattenSwitchControllerManagedSwitchPortsUntaggedVlans(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenSwitchControllerManagedSwitchPortsType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if cur_v, ok := i["access-mode"]; ok {
			tmp["access_mode"] = flattenSwitchControllerManagedSwitchPortsAccessMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_policy"
		if cur_v, ok := i["matched-dpp-policy"]; ok {
			tmp["matched_dpp_policy"] = flattenSwitchControllerManagedSwitchPortsMatchedDppPolicy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_intf_tags"
		if cur_v, ok := i["matched-dpp-intf-tags"]; ok {
			tmp["matched_dpp_intf_tags"] = flattenSwitchControllerManagedSwitchPortsMatchedDppIntfTags(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acl_group"
		if cur_v, ok := i["acl-group"]; ok {
			tmp["acl_group"] = flattenSwitchControllerManagedSwitchPortsAclGroup(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiswitch_acls"
		if cur_v, ok := i["fortiswitch-acls"]; ok {
			tmp["fortiswitch_acls"] = flattenSwitchControllerManagedSwitchPortsFortiswitchAcls(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if cur_v, ok := i["dhcp-snooping"]; ok {
			tmp["dhcp_snooping"] = flattenSwitchControllerManagedSwitchPortsDhcpSnooping(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if cur_v, ok := i["dhcp-snoop-option82-trust"]; ok {
			tmp["dhcp_snoop_option82_trust"] = flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_override"
		if cur_v, ok := i["dhcp-snoop-option82-override"]; ok {
			tmp["dhcp_snoop_option82_override"] = flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if cur_v, ok := i["arp-inspection-trust"]; ok {
			tmp["arp_inspection_trust"] = flattenSwitchControllerManagedSwitchPortsArpInspectionTrust(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping_flood_reports"
		if cur_v, ok := i["igmp-snooping-flood-reports"]; ok {
			tmp["igmp_snooping_flood_reports"] = flattenSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_snooping_flood_traffic"
		if cur_v, ok := i["mcast-snooping-flood-traffic"]; ok {
			tmp["mcast_snooping_flood_traffic"] = flattenSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if cur_v, ok := i["igmp-snooping"]; ok {
			tmp["igmp_snooping"] = flattenSwitchControllerManagedSwitchPortsIgmpSnooping(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if cur_v, ok := i["igmps-flood-reports"]; ok {
			tmp["igmps_flood_reports"] = flattenSwitchControllerManagedSwitchPortsIgmpsFloodReports(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if cur_v, ok := i["igmps-flood-traffic"]; ok {
			tmp["igmps_flood_traffic"] = flattenSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if cur_v, ok := i["stp-state"]; ok {
			tmp["stp_state"] = flattenSwitchControllerManagedSwitchPortsStpState(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if cur_v, ok := i["stp-root-guard"]; ok {
			tmp["stp_root_guard"] = flattenSwitchControllerManagedSwitchPortsStpRootGuard(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if cur_v, ok := i["stp-bpdu-guard"]; ok {
			tmp["stp_bpdu_guard"] = flattenSwitchControllerManagedSwitchPortsStpBpduGuard(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if cur_v, ok := i["stp-bpdu-guard-timeout"]; ok {
			tmp["stp_bpdu_guard_timeout"] = flattenSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if cur_v, ok := i["edge-port"]; ok {
			tmp["edge_port"] = flattenSwitchControllerManagedSwitchPortsEdgePort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if cur_v, ok := i["discard-mode"]; ok {
			tmp["discard_mode"] = flattenSwitchControllerManagedSwitchPortsDiscardMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if cur_v, ok := i["packet-sampler"]; ok {
			tmp["packet_sampler"] = flattenSwitchControllerManagedSwitchPortsPacketSampler(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if cur_v, ok := i["packet-sample-rate"]; ok {
			tmp["packet_sample_rate"] = flattenSwitchControllerManagedSwitchPortsPacketSampleRate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sampler"
		if cur_v, ok := i["sflow-sampler"]; ok {
			tmp["sflow_sampler"] = flattenSwitchControllerManagedSwitchPortsSflowSampler(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sample_rate"
		if cur_v, ok := i["sflow-sample-rate"]; ok {
			tmp["sflow_sample_rate"] = flattenSwitchControllerManagedSwitchPortsSflowSampleRate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if cur_v, ok := i["sflow-counter-interval"]; ok {
			tmp["sflow_counter_interval"] = flattenSwitchControllerManagedSwitchPortsSflowCounterInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if cur_v, ok := i["sample-direction"]; ok {
			tmp["sample_direction"] = flattenSwitchControllerManagedSwitchPortsSampleDirection(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if cur_v, ok := i["fec-capable"]; ok {
			tmp["fec_capable"] = flattenSwitchControllerManagedSwitchPortsFecCapable(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if cur_v, ok := i["fec-state"]; ok {
			tmp["fec_state"] = flattenSwitchControllerManagedSwitchPortsFecState(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if cur_v, ok := i["flow-control"]; ok {
			tmp["flow_control"] = flattenSwitchControllerManagedSwitchPortsFlowControl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if cur_v, ok := i["pause-meter"]; ok {
			tmp["pause_meter"] = flattenSwitchControllerManagedSwitchPortsPauseMeter(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if cur_v, ok := i["pause-meter-resume"]; ok {
			tmp["pause_meter_resume"] = flattenSwitchControllerManagedSwitchPortsPauseMeterResume(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if cur_v, ok := i["loop-guard"]; ok {
			tmp["loop_guard"] = flattenSwitchControllerManagedSwitchPortsLoopGuard(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if cur_v, ok := i["loop-guard-timeout"]; ok {
			tmp["loop_guard_timeout"] = flattenSwitchControllerManagedSwitchPortsLoopGuardTimeout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_policy"
		if cur_v, ok := i["port-policy"]; ok {
			tmp["port_policy"] = flattenSwitchControllerManagedSwitchPortsPortPolicy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if cur_v, ok := i["qos-policy"]; ok {
			tmp["qos_policy"] = flattenSwitchControllerManagedSwitchPortsQosPolicy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "storm_control_policy"
		if cur_v, ok := i["storm-control-policy"]; ok {
			tmp["storm_control_policy"] = flattenSwitchControllerManagedSwitchPortsStormControlPolicy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if cur_v, ok := i["port-security-policy"]; ok {
			tmp["port_security_policy"] = flattenSwitchControllerManagedSwitchPortsPortSecurityPolicy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool"
		if cur_v, ok := i["export-to-pool"]; ok {
			tmp["export_to_pool"] = flattenSwitchControllerManagedSwitchPortsExportToPool(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if cur_v, ok := i["interface-tags"]; ok {
			tmp["interface_tags"] = flattenSwitchControllerManagedSwitchPortsInterfaceTags(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_tags"
		if cur_v, ok := i["export-tags"]; ok {
			tmp["export_tags"] = flattenSwitchControllerManagedSwitchPortsExportTags(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool_flag"
		if cur_v, ok := i["export-to-pool_flag"]; ok {
			tmp["export_to_pool_flag"] = flattenSwitchControllerManagedSwitchPortsExportToPool_Flag(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if cur_v, ok := i["learning-limit"]; ok {
			tmp["learning_limit"] = flattenSwitchControllerManagedSwitchPortsLearningLimit(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if cur_v, ok := i["sticky-mac"]; ok {
			tmp["sticky_mac"] = flattenSwitchControllerManagedSwitchPortsStickyMac(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if cur_v, ok := i["lldp-status"]; ok {
			tmp["lldp_status"] = flattenSwitchControllerManagedSwitchPortsLldpStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if cur_v, ok := i["lldp-profile"]; ok {
			tmp["lldp_profile"] = flattenSwitchControllerManagedSwitchPortsLldpProfile(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to"
		if cur_v, ok := i["export-to"]; ok {
			tmp["export_to"] = flattenSwitchControllerManagedSwitchPortsExportTo(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if cur_v, ok := i["mac-addr"]; ok {
			tmp["mac_addr"] = flattenSwitchControllerManagedSwitchPortsMacAddr(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if cur_v, ok := i["port-selection-criteria"]; ok {
			tmp["port_selection_criteria"] = flattenSwitchControllerManagedSwitchPortsPortSelectionCriteria(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenSwitchControllerManagedSwitchPortsDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if cur_v, ok := i["lacp-speed"]; ok {
			tmp["lacp_speed"] = flattenSwitchControllerManagedSwitchPortsLacpSpeed(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if cur_v, ok := i["mode"]; ok {
			tmp["mode"] = flattenSwitchControllerManagedSwitchPortsMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if cur_v, ok := i["bundle"]; ok {
			tmp["bundle"] = flattenSwitchControllerManagedSwitchPortsBundle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if cur_v, ok := i["member-withdrawal-behavior"]; ok {
			tmp["member_withdrawal_behavior"] = flattenSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if cur_v, ok := i["mclag"]; ok {
			tmp["mclag"] = flattenSwitchControllerManagedSwitchPortsMclag(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if cur_v, ok := i["min-bundle"]; ok {
			tmp["min_bundle"] = flattenSwitchControllerManagedSwitchPortsMinBundle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if cur_v, ok := i["max-bundle"]; ok {
			tmp["max_bundle"] = flattenSwitchControllerManagedSwitchPortsMaxBundle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if cur_v, ok := i["members"]; ok {
			tmp["members"] = flattenSwitchControllerManagedSwitchPortsMembers(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "port_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortOwner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSwitchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSpeedMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIpSourceGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPtpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPtpPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAggregatorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapguard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlapTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsRpvstPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePreStandardDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortPrefixType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFortilinkPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLinkStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeCapable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStackingPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsP2PPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMclagIclPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAuthenticatedPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsRestrictedAuthPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsEncryptedPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFiberPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMediaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeStandard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeMaxPower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoeModeBtCabable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePortPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPoePortPower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsVirtualPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslLocalTrunkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslPeerPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslPeerDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIslPeerDeviceSn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFgtPeerPortName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFgtPeerDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlansAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if cur_v, ok := i["vlan-name"]; ok {
			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchPortsAllowedVlansVlanName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsUntaggedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if cur_v, ok := i["vlan-name"]; ok {
			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchPortsUntaggedVlansVlanName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsUntaggedVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAccessMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMatchedDppPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMatchedDppIntfTags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAclGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchPortsAclGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsAclGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFortiswitchAcls(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSwitchControllerManagedSwitchPortsFortiswitchAclsId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsFortiswitchAclsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if cur_v, ok := i["vlan-name"]; ok {
			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if cur_v, ok := i["circuit-id"]; ok {
			tmp["circuit_id"] = flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if cur_v, ok := i["remote-id"]; ok {
			tmp["remote_id"] = flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsArpInspectionTrust(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpsFloodReports(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpRootGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpBpduGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsEdgePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDiscardMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPacketSampler(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPacketSampleRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSflowSampler(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSflowSampleRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSflowCounterInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsSampleDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFecCapable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFecState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsFlowControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPauseMeter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPauseMeterResume(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLoopGuard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLoopGuardTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsQosPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStormControlPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortSecurityPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsExportToPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsInterfaceTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_name"
		if cur_v, ok := i["tag-name"]; ok {
			tmp["tag_name"] = flattenSwitchControllerManagedSwitchPortsInterfaceTagsTagName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "tag_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsInterfaceTagsTagName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsExportTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_name"
		if cur_v, ok := i["tag-name"]; ok {
			tmp["tag_name"] = flattenSwitchControllerManagedSwitchPortsExportTagsTagName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "tag_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsExportTagsTagName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsExportToPool_Flag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLearningLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsStickyMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLldpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLldpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsExportTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMacAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsPortSelectionCriteria(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsLacpSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsBundle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMclag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMinBundle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMaxBundle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if cur_v, ok := i["member-name"]; ok {
			tmp["member_name"] = flattenSwitchControllerManagedSwitchPortsMembersMemberName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "member_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchPortsMembersMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuard(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenSwitchControllerManagedSwitchIpSourceGuardPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenSwitchControllerManagedSwitchIpSourceGuardDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "binding_entry"
		if cur_v, ok := i["binding-entry"]; ok {
			tmp["binding_entry"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "port", d)
	return result
}

func flattenSwitchControllerManagedSwitchIpSourceGuardPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if cur_v, ok := i["entry-name"]; ok {
			tmp["entry_name"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if cur_v, ok := i["mac"]; ok {
			tmp["mac"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "entry_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettings(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchStpSettingsLocalOverride(i["local-override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenSwitchControllerManagedSwitchStpSettingsName(i["name"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchStpSettingsStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "revision"
	if _, ok := i["revision"]; ok {
		result["revision"] = flattenSwitchControllerManagedSwitchStpSettingsRevision(i["revision"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "hello_time"
	if _, ok := i["hello-time"]; ok {
		result["hello_time"] = flattenSwitchControllerManagedSwitchStpSettingsHelloTime(i["hello-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "forward_time"
	if _, ok := i["forward-time"]; ok {
		result["forward_time"] = flattenSwitchControllerManagedSwitchStpSettingsForwardTime(i["forward-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_age"
	if _, ok := i["max-age"]; ok {
		result["max_age"] = flattenSwitchControllerManagedSwitchStpSettingsMaxAge(i["max-age"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_hops"
	if _, ok := i["max-hops"]; ok {
		result["max_hops"] = flattenSwitchControllerManagedSwitchStpSettingsMaxHops(i["max-hops"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pending_timer"
	if _, ok := i["pending-timer"]; ok {
		result["pending_timer"] = flattenSwitchControllerManagedSwitchStpSettingsPendingTimer(i["pending-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchStpSettingsLocalOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsRevision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsHelloTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsForwardTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsMaxHops(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpSettingsPendingTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpInstance(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSwitchControllerManagedSwitchStpInstanceId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenSwitchControllerManagedSwitchStpInstancePriority(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerManagedSwitchStpInstanceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStpInstancePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpSysinfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfo(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchSnmpSysinfoStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "engine_id"
	if _, ok := i["engine-id"]; ok {
		result["engine_id"] = flattenSwitchControllerManagedSwitchSnmpSysinfoEngineId(i["engine-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "description"
	if _, ok := i["description"]; ok {
		result["description"] = flattenSwitchControllerManagedSwitchSnmpSysinfoDescription(i["description"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "contact_info"
	if _, ok := i["contact-info"]; ok {
		result["contact_info"] = flattenSwitchControllerManagedSwitchSnmpSysinfoContactInfo(i["contact-info"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "location"
	if _, ok := i["location"]; ok {
		result["location"] = flattenSwitchControllerManagedSwitchSnmpSysinfoLocation(i["location"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "trap_high_cpu_threshold"
	if _, ok := i["trap-high-cpu-threshold"]; ok {
		result["trap_high_cpu_threshold"] = flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(i["trap-high-cpu-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "trap_low_memory_threshold"
	if _, ok := i["trap-low-memory-threshold"]; ok {
		result["trap_low_memory_threshold"] = flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(i["trap-low-memory-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "trap_log_full_threshold"
	if _, ok := i["trap-log-full-threshold"]; ok {
		result["trap_log_full_threshold"] = flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(i["trap-log-full-threshold"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSwitchControllerManagedSwitchSnmpCommunityId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerManagedSwitchSnmpCommunityName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerManagedSwitchSnmpCommunityStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if cur_v, ok := i["hosts"]; ok {
			tmp["hosts"] = flattenSwitchControllerManagedSwitchSnmpCommunityHosts(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if cur_v, ok := i["query-v1-status"]; ok {
			tmp["query_v1_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_port"
		if cur_v, ok := i["query-v1-port"]; ok {
			tmp["query_v1_port"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if cur_v, ok := i["query-v2c-status"]; ok {
			tmp["query_v2c_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_port"
		if cur_v, ok := i["query-v2c-port"]; ok {
			tmp["query_v2c_port"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if cur_v, ok := i["trap-v1-status"]; ok {
			tmp["trap_v1_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_lport"
		if cur_v, ok := i["trap-v1-lport"]; ok {
			tmp["trap_v1_lport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_rport"
		if cur_v, ok := i["trap-v1-rport"]; ok {
			tmp["trap_v1_rport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if cur_v, ok := i["trap-v2c-status"]; ok {
			tmp["trap_v2c_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_lport"
		if cur_v, ok := i["trap-v2c-lport"]; ok {
			tmp["trap_v2c_lport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_rport"
		if cur_v, ok := i["trap-v2c-rport"]; ok {
			tmp["trap_v2c_rport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "events"
		if cur_v, ok := i["events"]; ok {
			tmp["events"] = flattenSwitchControllerManagedSwitchSnmpCommunityEvents(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerManagedSwitchSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSwitchControllerManagedSwitchSnmpCommunityHostsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpCommunityEvents(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchOverrideSnmpUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUser(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchSnmpUserName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if cur_v, ok := i["queries"]; ok {
			tmp["queries"] = flattenSwitchControllerManagedSwitchSnmpUserQueries(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_port"
		if cur_v, ok := i["query-port"]; ok {
			tmp["query_port"] = flattenSwitchControllerManagedSwitchSnmpUserQueryPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if cur_v, ok := i["security-level"]; ok {
			tmp["security_level"] = flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if cur_v, ok := i["auth-proto"]; ok {
			tmp["auth_proto"] = flattenSwitchControllerManagedSwitchSnmpUserAuthProto(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if cur_v, ok := i["auth-pwd"]; ok {
			tmp["auth_pwd"] = flattenSwitchControllerManagedSwitchSnmpUserAuthPwd(cur_v, d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_pwd"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if cur_v, ok := i["priv-proto"]; ok {
			tmp["priv_proto"] = flattenSwitchControllerManagedSwitchSnmpUserPrivProto(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if cur_v, ok := i["priv-pwd"]; ok {
			tmp["priv_pwd"] = flattenSwitchControllerManagedSwitchSnmpUserPrivPwd(cur_v, d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["priv_pwd"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchSnmpUserName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserAuthPwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSnmpUserPrivPwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchQosDropPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchQosRedProbability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchStpSettings(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchSwitchStpSettingsStatus(i["status"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSwitchStpSettingsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchLog(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchSwitchLogLocalOverride(i["local-override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSwitchControllerManagedSwitchSwitchLogStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenSwitchControllerManagedSwitchSwitchLogSeverity(i["severity"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchSwitchLogLocalOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchLogStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchSwitchLogSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLog(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchRemoteLogName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerManagedSwitchRemoteLogStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if cur_v, ok := i["server"]; ok {
			tmp["server"] = flattenSwitchControllerManagedSwitchRemoteLogServer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenSwitchControllerManagedSwitchRemoteLogPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if cur_v, ok := i["severity"]; ok {
			tmp["severity"] = flattenSwitchControllerManagedSwitchRemoteLogSeverity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csv"
		if cur_v, ok := i["csv"]; ok {
			tmp["csv"] = flattenSwitchControllerManagedSwitchRemoteLogCsv(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "facility"
		if cur_v, ok := i["facility"]; ok {
			tmp["facility"] = flattenSwitchControllerManagedSwitchRemoteLogFacility(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchRemoteLogName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogCsv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchRemoteLogFacility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchStormControlLocalOverride(i["local-override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rate"
	if _, ok := i["rate"]; ok {
		result["rate"] = flattenSwitchControllerManagedSwitchStormControlRate(i["rate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unknown_unicast"
	if _, ok := i["unknown-unicast"]; ok {
		result["unknown_unicast"] = flattenSwitchControllerManagedSwitchStormControlUnknownUnicast(i["unknown-unicast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unknown_multicast"
	if _, ok := i["unknown-multicast"]; ok {
		result["unknown_multicast"] = flattenSwitchControllerManagedSwitchStormControlUnknownMulticast(i["unknown-multicast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "broadcast"
	if _, ok := i["broadcast"]; ok {
		result["broadcast"] = flattenSwitchControllerManagedSwitchStormControlBroadcast(i["broadcast"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchStormControlLocalOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlUnknownUnicast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStormControlBroadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirror(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchMirrorName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerManagedSwitchMirrorStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switching_packet"
		if cur_v, ok := i["switching-packet"]; ok {
			tmp["switching_packet"] = flattenSwitchControllerManagedSwitchMirrorSwitchingPacket(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if cur_v, ok := i["dst"]; ok {
			tmp["dst"] = flattenSwitchControllerManagedSwitchMirrorDst(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ingress"
		if cur_v, ok := i["src-ingress"]; ok {
			tmp["src_ingress"] = flattenSwitchControllerManagedSwitchMirrorSrcIngress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_egress"
		if cur_v, ok := i["src-egress"]; ok {
			tmp["src_egress"] = flattenSwitchControllerManagedSwitchMirrorSrcEgress(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchMirrorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSwitchingPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSrcIngress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchMirrorSrcIngressName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchMirrorSrcIngressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchMirrorSrcEgressName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMac(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSwitchControllerManagedSwitchStaticMacId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenSwitchControllerManagedSwitchStaticMacType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if cur_v, ok := i["vlan"]; ok {
			tmp["vlan"] = flattenSwitchControllerManagedSwitchStaticMacVlan(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if cur_v, ok := i["mac"]; ok {
			tmp["mac"] = flattenSwitchControllerManagedSwitchStaticMacMac(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenSwitchControllerManagedSwitchStaticMacInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenSwitchControllerManagedSwitchStaticMacDescription(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerManagedSwitchStaticMacId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMacDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchCustomCommand(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if cur_v, ok := i["command-entry"]; ok {
			tmp["command_entry"] = flattenSwitchControllerManagedSwitchCustomCommandCommandEntry(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if cur_v, ok := i["command-name"]; ok {
			tmp["command_name"] = flattenSwitchControllerManagedSwitchCustomCommandCommandName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "command_entry", d)
	return result
}

func flattenSwitchControllerManagedSwitchCustomCommandCommandEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClient(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if cur_v, ok := i["vlan"]; ok {
			tmp["vlan"] = flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if cur_v, ok := i["mac"]; ok {
			tmp["mac"] = flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(i["local-override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aging_time"
	if _, ok := i["aging-time"]; ok {
		result["aging_time"] = flattenSwitchControllerManagedSwitchIgmpSnoopingAgingTime(i["aging-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "flood_unknown_multicast"
	if _, ok := i["flood-unknown-multicast"]; ok {
		result["flood_unknown_multicast"] = flattenSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(i["flood-unknown-multicast"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlans"
	if _, ok := i["vlans"]; ok {
		result["vlans"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlans(i["vlans"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingAgingTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if cur_v, ok := i["vlan-name"]; ok {
			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy"
		if cur_v, ok := i["proxy"]; ok {
			tmp["proxy"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier"
		if cur_v, ok := i["querier"]; ok {
			tmp["querier"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier_addr"
		if cur_v, ok := i["querier-addr"]; ok {
			tmp["querier_addr"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if cur_v, ok := i["version"]; ok {
			tmp["version"] = flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettings(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := i["local-override"]; ok {
		result["local_override"] = flattenSwitchControllerManagedSwitch8021XSettingsLocalOverride(i["local-override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "link_down_auth"
	if _, ok := i["link-down-auth"]; ok {
		result["link_down_auth"] = flattenSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(i["link-down-auth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "reauth_period"
	if _, ok := i["reauth-period"]; ok {
		result["reauth_period"] = flattenSwitchControllerManagedSwitch8021XSettingsReauthPeriod(i["reauth-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_reauth_attempt"
	if _, ok := i["max-reauth-attempt"]; ok {
		result["max_reauth_attempt"] = flattenSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(i["max-reauth-attempt"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tx_period"
	if _, ok := i["tx-period"]; ok {
		result["tx_period"] = flattenSwitchControllerManagedSwitch8021XSettingsTxPeriod(i["tx-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mab_reauth"
	if _, ok := i["mab-reauth"]; ok {
		result["mab_reauth"] = flattenSwitchControllerManagedSwitch8021XSettingsMabReauth(i["mab-reauth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_username_delimiter"
	if _, ok := i["mac-username-delimiter"]; ok {
		result["mac_username_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(i["mac-username-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_password_delimiter"
	if _, ok := i["mac-password-delimiter"]; ok {
		result["mac_password_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(i["mac-password-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_calling_station_delimiter"
	if _, ok := i["mac-calling-station-delimiter"]; ok {
		result["mac_calling_station_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(i["mac-calling-station-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_called_station_delimiter"
	if _, ok := i["mac-called-station-delimiter"]; ok {
		result["mac_called_station_delimiter"] = flattenSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(i["mac-called-station-delimiter"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mac_case"
	if _, ok := i["mac-case"]; ok {
		result["mac_case"] = flattenSwitchControllerManagedSwitch8021XSettingsMacCase(i["mac-case"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerManagedSwitch8021XSettingsLocalOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsReauthPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsTxPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMabReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitch8021XSettingsMacCase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerManagedSwitch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("switch_id", flattenSwitchControllerManagedSwitchSwitchId(o["switch-id"], d, "switch_id", sv)); err != nil {
		if !fortiAPIPatch(o["switch-id"]) {
			return fmt.Errorf("Error reading switch_id: %v", err)
		}
	}

	if err = d.Set("sn", flattenSwitchControllerManagedSwitchSn(o["sn"], d, "sn", sv)); err != nil {
		if !fortiAPIPatch(o["sn"]) {
			return fmt.Errorf("Error reading sn: %v", err)
		}
	}

	if err = d.Set("name", flattenSwitchControllerManagedSwitchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerManagedSwitchDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("switch_profile", flattenSwitchControllerManagedSwitchSwitchProfile(o["switch-profile"], d, "switch_profile", sv)); err != nil {
		if !fortiAPIPatch(o["switch-profile"]) {
			return fmt.Errorf("Error reading switch_profile: %v", err)
		}
	}

	if err = d.Set("access_profile", flattenSwitchControllerManagedSwitchAccessProfile(o["access-profile"], d, "access_profile", sv)); err != nil {
		if !fortiAPIPatch(o["access-profile"]) {
			return fmt.Errorf("Error reading access_profile: %v", err)
		}
	}

	if err = d.Set("purdue_level", flattenSwitchControllerManagedSwitchPurdueLevel(o["purdue-level"], d, "purdue_level", sv)); err != nil {
		if !fortiAPIPatch(o["purdue-level"]) {
			return fmt.Errorf("Error reading purdue_level: %v", err)
		}
	}

	if err = d.Set("fsw_wan1_peer", flattenSwitchControllerManagedSwitchFswWan1Peer(o["fsw-wan1-peer"], d, "fsw_wan1_peer", sv)); err != nil {
		if !fortiAPIPatch(o["fsw-wan1-peer"]) {
			return fmt.Errorf("Error reading fsw_wan1_peer: %v", err)
		}
	}

	if err = d.Set("fsw_wan1_admin", flattenSwitchControllerManagedSwitchFswWan1Admin(o["fsw-wan1-admin"], d, "fsw_wan1_admin", sv)); err != nil {
		if !fortiAPIPatch(o["fsw-wan1-admin"]) {
			return fmt.Errorf("Error reading fsw_wan1_admin: %v", err)
		}
	}

	if err = d.Set("fsw_wan2_peer", flattenSwitchControllerManagedSwitchFswWan2Peer(o["fsw-wan2-peer"], d, "fsw_wan2_peer", sv)); err != nil {
		if !fortiAPIPatch(o["fsw-wan2-peer"]) {
			return fmt.Errorf("Error reading fsw_wan2_peer: %v", err)
		}
	}

	if err = d.Set("fsw_wan2_admin", flattenSwitchControllerManagedSwitchFswWan2Admin(o["fsw-wan2-admin"], d, "fsw_wan2_admin", sv)); err != nil {
		if !fortiAPIPatch(o["fsw-wan2-admin"]) {
			return fmt.Errorf("Error reading fsw_wan2_admin: %v", err)
		}
	}

	if err = d.Set("poe_pre_standard_detection", flattenSwitchControllerManagedSwitchPoePreStandardDetection(o["poe-pre-standard-detection"], d, "poe_pre_standard_detection", sv)); err != nil {
		if !fortiAPIPatch(o["poe-pre-standard-detection"]) {
			return fmt.Errorf("Error reading poe_pre_standard_detection: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", flattenSwitchControllerManagedSwitchDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-server-access-list"]) {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if err = d.Set("poe_detection_type", flattenSwitchControllerManagedSwitchPoeDetectionType(o["poe-detection-type"], d, "poe_detection_type", sv)); err != nil {
		if !fortiAPIPatch(o["poe-detection-type"]) {
			return fmt.Errorf("Error reading poe_detection_type: %v", err)
		}
	}

	if err = d.Set("poe_lldp_detection", flattenSwitchControllerManagedSwitchPoeLldpDetection(o["poe-lldp-detection"], d, "poe_lldp_detection", sv)); err != nil {
		if !fortiAPIPatch(o["poe-lldp-detection"]) {
			return fmt.Errorf("Error reading poe_lldp_detection: %v", err)
		}
	}

	if err = d.Set("directly_connected", flattenSwitchControllerManagedSwitchDirectlyConnected(o["directly-connected"], d, "directly_connected", sv)); err != nil {
		if !fortiAPIPatch(o["directly-connected"]) {
			return fmt.Errorf("Error reading directly_connected: %v", err)
		}
	}

	if err = d.Set("version", flattenSwitchControllerManagedSwitchVersion(o["version"], d, "version", sv)); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("max_allowed_trunk_members", flattenSwitchControllerManagedSwitchMaxAllowedTrunkMembers(o["max-allowed-trunk-members"], d, "max_allowed_trunk_members", sv)); err != nil {
		if !fortiAPIPatch(o["max-allowed-trunk-members"]) {
			return fmt.Errorf("Error reading max_allowed_trunk_members: %v", err)
		}
	}

	if err = d.Set("pre_provisioned", flattenSwitchControllerManagedSwitchPreProvisioned(o["pre-provisioned"], d, "pre_provisioned", sv)); err != nil {
		if !fortiAPIPatch(o["pre-provisioned"]) {
			return fmt.Errorf("Error reading pre_provisioned: %v", err)
		}
	}

	if err = d.Set("l3_discovered", flattenSwitchControllerManagedSwitchL3Discovered(o["l3-discovered"], d, "l3_discovered", sv)); err != nil {
		if !fortiAPIPatch(o["l3-discovered"]) {
			return fmt.Errorf("Error reading l3_discovered: %v", err)
		}
	}

	if err = d.Set("mgmt_mode", flattenSwitchControllerManagedSwitchMgmtMode(o["mgmt-mode"], d, "mgmt_mode", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-mode"]) {
			return fmt.Errorf("Error reading mgmt_mode: %v", err)
		}
	}

	if err = d.Set("tunnel_discovered", flattenSwitchControllerManagedSwitchTunnelDiscovered(o["tunnel-discovered"], d, "tunnel_discovered", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-discovered"]) {
			return fmt.Errorf("Error reading tunnel_discovered: %v", err)
		}
	}

	if err = d.Set("tdr_supported", flattenSwitchControllerManagedSwitchTdrSupported(o["tdr-supported"], d, "tdr_supported", sv)); err != nil {
		if !fortiAPIPatch(o["tdr-supported"]) {
			return fmt.Errorf("Error reading tdr_supported: %v", err)
		}
	}

	{
		v := flattenSwitchControllerManagedSwitchDynamicCapability(o["dynamic-capability"], d, "dynamic_capability", sv)
		new_version_map := map[string][]string{
			">=": []string{"6.4.2"},
		}
		if i2ss2arrFortiAPIUpgrade(sv, new_version_map) == true {
			if vx, ok := v.(string); ok {
				vxx, err := strconv.Atoi(vx)
				if err == nil {
					v = vxx
				}
			}
		}

		if err = d.Set("dynamic_capability", v); err != nil {
			if !fortiAPIPatch(o["dynamic-capability"]) {
				return fmt.Errorf("Error reading dynamic_capability: %v", err)
			}
		}
	}

	if err = d.Set("switch_device_tag", flattenSwitchControllerManagedSwitchSwitchDeviceTag(o["switch-device-tag"], d, "switch_device_tag", sv)); err != nil {
		if !fortiAPIPatch(o["switch-device-tag"]) {
			return fmt.Errorf("Error reading switch_device_tag: %v", err)
		}
	}

	if err = d.Set("switch_dhcp_opt43_key", flattenSwitchControllerManagedSwitchSwitchDhcp_Opt43_Key(o["switch-dhcp_opt43_key"], d, "switch_dhcp_opt43_key", sv)); err != nil {
		if !fortiAPIPatch(o["switch-dhcp_opt43_key"]) {
			return fmt.Errorf("Error reading switch_dhcp_opt43_key: %v", err)
		}
	}

	if err = d.Set("mclag_igmp_snooping_aware", flattenSwitchControllerManagedSwitchMclagIgmpSnoopingAware(o["mclag-igmp-snooping-aware"], d, "mclag_igmp_snooping_aware", sv)); err != nil {
		if !fortiAPIPatch(o["mclag-igmp-snooping-aware"]) {
			return fmt.Errorf("Error reading mclag_igmp_snooping_aware: %v", err)
		}
	}

	if err = d.Set("dynamically_discovered", flattenSwitchControllerManagedSwitchDynamicallyDiscovered(o["dynamically-discovered"], d, "dynamically_discovered", sv)); err != nil {
		if !fortiAPIPatch(o["dynamically-discovered"]) {
			return fmt.Errorf("Error reading dynamically_discovered: %v", err)
		}
	}

	if err = d.Set("ptp_status", flattenSwitchControllerManagedSwitchPtpStatus(o["ptp-status"], d, "ptp_status", sv)); err != nil {
		if !fortiAPIPatch(o["ptp-status"]) {
			return fmt.Errorf("Error reading ptp_status: %v", err)
		}
	}

	if err = d.Set("ptp_profile", flattenSwitchControllerManagedSwitchPtpProfile(o["ptp-profile"], d, "ptp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ptp-profile"]) {
			return fmt.Errorf("Error reading ptp_profile: %v", err)
		}
	}

	if err = d.Set("radius_nas_ip_override", flattenSwitchControllerManagedSwitchRadiusNasIpOverride(o["radius-nas-ip-override"], d, "radius_nas_ip_override", sv)); err != nil {
		if !fortiAPIPatch(o["radius-nas-ip-override"]) {
			return fmt.Errorf("Error reading radius_nas_ip_override: %v", err)
		}
	}

	if err = d.Set("radius_nas_ip", flattenSwitchControllerManagedSwitchRadiusNasIp(o["radius-nas-ip"], d, "radius_nas_ip", sv)); err != nil {
		if !fortiAPIPatch(o["radius-nas-ip"]) {
			return fmt.Errorf("Error reading radius_nas_ip: %v", err)
		}
	}

	if err = d.Set("route_offload", flattenSwitchControllerManagedSwitchRouteOffload(o["route-offload"], d, "route_offload", sv)); err != nil {
		if !fortiAPIPatch(o["route-offload"]) {
			return fmt.Errorf("Error reading route_offload: %v", err)
		}
	}

	if err = d.Set("route_offload_mclag", flattenSwitchControllerManagedSwitchRouteOffloadMclag(o["route-offload-mclag"], d, "route_offload_mclag", sv)); err != nil {
		if !fortiAPIPatch(o["route-offload-mclag"]) {
			return fmt.Errorf("Error reading route_offload_mclag: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("route_offload_router", flattenSwitchControllerManagedSwitchRouteOffloadRouter(o["route-offload-router"], d, "route_offload_router", sv)); err != nil {
			if !fortiAPIPatch(o["route-offload-router"]) {
				return fmt.Errorf("Error reading route_offload_router: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_offload_router"); ok {
			if err = d.Set("route_offload_router", flattenSwitchControllerManagedSwitchRouteOffloadRouter(o["route-offload-router"], d, "route_offload_router", sv)); err != nil {
				if !fortiAPIPatch(o["route-offload-router"]) {
					return fmt.Errorf("Error reading route_offload_router: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("vlan", flattenSwitchControllerManagedSwitchVlan(o["vlan"], d, "vlan", sv)); err != nil {
			if !fortiAPIPatch(o["vlan"]) {
				return fmt.Errorf("Error reading vlan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan"); ok {
			if err = d.Set("vlan", flattenSwitchControllerManagedSwitchVlan(o["vlan"], d, "vlan", sv)); err != nil {
				if !fortiAPIPatch(o["vlan"]) {
					return fmt.Errorf("Error reading vlan: %v", err)
				}
			}
		}
	}

	if err = d.Set("type", flattenSwitchControllerManagedSwitchType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("owner_vdom", flattenSwitchControllerManagedSwitchOwnerVdom(o["owner-vdom"], d, "owner_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["owner-vdom"]) {
			return fmt.Errorf("Error reading owner_vdom: %v", err)
		}
	}

	if err = d.Set("flow_identity", flattenSwitchControllerManagedSwitchFlowIdentity(o["flow-identity"], d, "flow_identity", sv)); err != nil {
		if !fortiAPIPatch(o["flow-identity"]) {
			return fmt.Errorf("Error reading flow_identity: %v", err)
		}
	}

	if err = d.Set("staged_image_version", flattenSwitchControllerManagedSwitchStagedImageVersion(o["staged-image-version"], d, "staged_image_version", sv)); err != nil {
		if !fortiAPIPatch(o["staged-image-version"]) {
			return fmt.Errorf("Error reading staged_image_version: %v", err)
		}
	}

	if err = d.Set("delayed_restart_trigger", flattenSwitchControllerManagedSwitchDelayedRestartTrigger(o["delayed-restart-trigger"], d, "delayed_restart_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["delayed-restart-trigger"]) {
			return fmt.Errorf("Error reading delayed_restart_trigger: %v", err)
		}
	}

	if err = d.Set("firmware_provision", flattenSwitchControllerManagedSwitchFirmwareProvision(o["firmware-provision"], d, "firmware_provision", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision"]) {
			return fmt.Errorf("Error reading firmware_provision: %v", err)
		}
	}

	if err = d.Set("firmware_provision_version", flattenSwitchControllerManagedSwitchFirmwareProvisionVersion(o["firmware-provision-version"], d, "firmware_provision_version", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision-version"]) {
			return fmt.Errorf("Error reading firmware_provision_version: %v", err)
		}
	}

	if err = d.Set("firmware_provision_latest", flattenSwitchControllerManagedSwitchFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision-latest"]) {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ports", flattenSwitchControllerManagedSwitchPorts(o["ports"], d, "ports", sv)); err != nil {
			if !fortiAPIPatch(o["ports"]) {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ports"); ok {
			if err = d.Set("ports", flattenSwitchControllerManagedSwitchPorts(o["ports"], d, "ports", sv)); err != nil {
				if !fortiAPIPatch(o["ports"]) {
					return fmt.Errorf("Error reading ports: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip_source_guard", flattenSwitchControllerManagedSwitchIpSourceGuard(o["ip-source-guard"], d, "ip_source_guard", sv)); err != nil {
			if !fortiAPIPatch(o["ip-source-guard"]) {
				return fmt.Errorf("Error reading ip_source_guard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_source_guard"); ok {
			if err = d.Set("ip_source_guard", flattenSwitchControllerManagedSwitchIpSourceGuard(o["ip-source-guard"], d, "ip_source_guard", sv)); err != nil {
				if !fortiAPIPatch(o["ip-source-guard"]) {
					return fmt.Errorf("Error reading ip_source_guard: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("stp_settings", flattenSwitchControllerManagedSwitchStpSettings(o["stp-settings"], d, "stp_settings", sv)); err != nil {
			if !fortiAPIPatch(o["stp-settings"]) {
				return fmt.Errorf("Error reading stp_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("stp_settings"); ok {
			if err = d.Set("stp_settings", flattenSwitchControllerManagedSwitchStpSettings(o["stp-settings"], d, "stp_settings", sv)); err != nil {
				if !fortiAPIPatch(o["stp-settings"]) {
					return fmt.Errorf("Error reading stp_settings: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("stp_instance", flattenSwitchControllerManagedSwitchStpInstance(o["stp-instance"], d, "stp_instance", sv)); err != nil {
			if !fortiAPIPatch(o["stp-instance"]) {
				return fmt.Errorf("Error reading stp_instance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("stp_instance"); ok {
			if err = d.Set("stp_instance", flattenSwitchControllerManagedSwitchStpInstance(o["stp-instance"], d, "stp_instance", sv)); err != nil {
				if !fortiAPIPatch(o["stp-instance"]) {
					return fmt.Errorf("Error reading stp_instance: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_snmp_sysinfo", flattenSwitchControllerManagedSwitchOverrideSnmpSysinfo(o["override-snmp-sysinfo"], d, "override_snmp_sysinfo", sv)); err != nil {
		if !fortiAPIPatch(o["override-snmp-sysinfo"]) {
			return fmt.Errorf("Error reading override_snmp_sysinfo: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("snmp_sysinfo", flattenSwitchControllerManagedSwitchSnmpSysinfo(o["snmp-sysinfo"], d, "snmp_sysinfo", sv)); err != nil {
			if !fortiAPIPatch(o["snmp-sysinfo"]) {
				return fmt.Errorf("Error reading snmp_sysinfo: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_sysinfo"); ok {
			if err = d.Set("snmp_sysinfo", flattenSwitchControllerManagedSwitchSnmpSysinfo(o["snmp-sysinfo"], d, "snmp_sysinfo", sv)); err != nil {
				if !fortiAPIPatch(o["snmp-sysinfo"]) {
					return fmt.Errorf("Error reading snmp_sysinfo: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_snmp_trap_threshold", flattenSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(o["override-snmp-trap-threshold"], d, "override_snmp_trap_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["override-snmp-trap-threshold"]) {
			return fmt.Errorf("Error reading override_snmp_trap_threshold: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("snmp_trap_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThreshold(o["snmp-trap-threshold"], d, "snmp_trap_threshold", sv)); err != nil {
			if !fortiAPIPatch(o["snmp-trap-threshold"]) {
				return fmt.Errorf("Error reading snmp_trap_threshold: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_trap_threshold"); ok {
			if err = d.Set("snmp_trap_threshold", flattenSwitchControllerManagedSwitchSnmpTrapThreshold(o["snmp-trap-threshold"], d, "snmp_trap_threshold", sv)); err != nil {
				if !fortiAPIPatch(o["snmp-trap-threshold"]) {
					return fmt.Errorf("Error reading snmp_trap_threshold: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_snmp_community", flattenSwitchControllerManagedSwitchOverrideSnmpCommunity(o["override-snmp-community"], d, "override_snmp_community", sv)); err != nil {
		if !fortiAPIPatch(o["override-snmp-community"]) {
			return fmt.Errorf("Error reading override_snmp_community: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("snmp_community", flattenSwitchControllerManagedSwitchSnmpCommunity(o["snmp-community"], d, "snmp_community", sv)); err != nil {
			if !fortiAPIPatch(o["snmp-community"]) {
				return fmt.Errorf("Error reading snmp_community: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_community"); ok {
			if err = d.Set("snmp_community", flattenSwitchControllerManagedSwitchSnmpCommunity(o["snmp-community"], d, "snmp_community", sv)); err != nil {
				if !fortiAPIPatch(o["snmp-community"]) {
					return fmt.Errorf("Error reading snmp_community: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_snmp_user", flattenSwitchControllerManagedSwitchOverrideSnmpUser(o["override-snmp-user"], d, "override_snmp_user", sv)); err != nil {
		if !fortiAPIPatch(o["override-snmp-user"]) {
			return fmt.Errorf("Error reading override_snmp_user: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("snmp_user", flattenSwitchControllerManagedSwitchSnmpUser(o["snmp-user"], d, "snmp_user", sv)); err != nil {
			if !fortiAPIPatch(o["snmp-user"]) {
				return fmt.Errorf("Error reading snmp_user: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("snmp_user"); ok {
			if err = d.Set("snmp_user", flattenSwitchControllerManagedSwitchSnmpUser(o["snmp-user"], d, "snmp_user", sv)); err != nil {
				if !fortiAPIPatch(o["snmp-user"]) {
					return fmt.Errorf("Error reading snmp_user: %v", err)
				}
			}
		}
	}

	if err = d.Set("qos_drop_policy", flattenSwitchControllerManagedSwitchQosDropPolicy(o["qos-drop-policy"], d, "qos_drop_policy", sv)); err != nil {
		if !fortiAPIPatch(o["qos-drop-policy"]) {
			return fmt.Errorf("Error reading qos_drop_policy: %v", err)
		}
	}

	if err = d.Set("qos_red_probability", flattenSwitchControllerManagedSwitchQosRedProbability(o["qos-red-probability"], d, "qos_red_probability", sv)); err != nil {
		if !fortiAPIPatch(o["qos-red-probability"]) {
			return fmt.Errorf("Error reading qos_red_probability: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("switch_stp_settings", flattenSwitchControllerManagedSwitchSwitchStpSettings(o["switch-stp-settings"], d, "switch_stp_settings", sv)); err != nil {
			if !fortiAPIPatch(o["switch-stp-settings"]) {
				return fmt.Errorf("Error reading switch_stp_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("switch_stp_settings"); ok {
			if err = d.Set("switch_stp_settings", flattenSwitchControllerManagedSwitchSwitchStpSettings(o["switch-stp-settings"], d, "switch_stp_settings", sv)); err != nil {
				if !fortiAPIPatch(o["switch-stp-settings"]) {
					return fmt.Errorf("Error reading switch_stp_settings: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("switch_log", flattenSwitchControllerManagedSwitchSwitchLog(o["switch-log"], d, "switch_log", sv)); err != nil {
			if !fortiAPIPatch(o["switch-log"]) {
				return fmt.Errorf("Error reading switch_log: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("switch_log"); ok {
			if err = d.Set("switch_log", flattenSwitchControllerManagedSwitchSwitchLog(o["switch-log"], d, "switch_log", sv)); err != nil {
				if !fortiAPIPatch(o["switch-log"]) {
					return fmt.Errorf("Error reading switch_log: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("remote_log", flattenSwitchControllerManagedSwitchRemoteLog(o["remote-log"], d, "remote_log", sv)); err != nil {
			if !fortiAPIPatch(o["remote-log"]) {
				return fmt.Errorf("Error reading remote_log: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_log"); ok {
			if err = d.Set("remote_log", flattenSwitchControllerManagedSwitchRemoteLog(o["remote-log"], d, "remote_log", sv)); err != nil {
				if !fortiAPIPatch(o["remote-log"]) {
					return fmt.Errorf("Error reading remote_log: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("storm_control", flattenSwitchControllerManagedSwitchStormControl(o["storm-control"], d, "storm_control", sv)); err != nil {
			if !fortiAPIPatch(o["storm-control"]) {
				return fmt.Errorf("Error reading storm_control: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("storm_control"); ok {
			if err = d.Set("storm_control", flattenSwitchControllerManagedSwitchStormControl(o["storm-control"], d, "storm_control", sv)); err != nil {
				if !fortiAPIPatch(o["storm-control"]) {
					return fmt.Errorf("Error reading storm_control: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("mirror", flattenSwitchControllerManagedSwitchMirror(o["mirror"], d, "mirror", sv)); err != nil {
			if !fortiAPIPatch(o["mirror"]) {
				return fmt.Errorf("Error reading mirror: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mirror"); ok {
			if err = d.Set("mirror", flattenSwitchControllerManagedSwitchMirror(o["mirror"], d, "mirror", sv)); err != nil {
				if !fortiAPIPatch(o["mirror"]) {
					return fmt.Errorf("Error reading mirror: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("static_mac", flattenSwitchControllerManagedSwitchStaticMac(o["static-mac"], d, "static_mac", sv)); err != nil {
			if !fortiAPIPatch(o["static-mac"]) {
				return fmt.Errorf("Error reading static_mac: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("static_mac"); ok {
			if err = d.Set("static_mac", flattenSwitchControllerManagedSwitchStaticMac(o["static-mac"], d, "static_mac", sv)); err != nil {
				if !fortiAPIPatch(o["static-mac"]) {
					return fmt.Errorf("Error reading static_mac: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("custom_command", flattenSwitchControllerManagedSwitchCustomCommand(o["custom-command"], d, "custom_command", sv)); err != nil {
			if !fortiAPIPatch(o["custom-command"]) {
				return fmt.Errorf("Error reading custom_command: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_command"); ok {
			if err = d.Set("custom_command", flattenSwitchControllerManagedSwitchCustomCommand(o["custom-command"], d, "custom_command", sv)); err != nil {
				if !fortiAPIPatch(o["custom-command"]) {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dhcp_snooping_static_client", flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client", sv)); err != nil {
			if !fortiAPIPatch(o["dhcp-snooping-static-client"]) {
				return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dhcp_snooping_static_client"); ok {
			if err = d.Set("dhcp_snooping_static_client", flattenSwitchControllerManagedSwitchDhcpSnoopingStaticClient(o["dhcp-snooping-static-client"], d, "dhcp_snooping_static_client", sv)); err != nil {
				if !fortiAPIPatch(o["dhcp-snooping-static-client"]) {
					return fmt.Errorf("Error reading dhcp_snooping_static_client: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("igmp_snooping", flattenSwitchControllerManagedSwitchIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping", sv)); err != nil {
			if !fortiAPIPatch(o["igmp-snooping"]) {
				return fmt.Errorf("Error reading igmp_snooping: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("igmp_snooping"); ok {
			if err = d.Set("igmp_snooping", flattenSwitchControllerManagedSwitchIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping", sv)); err != nil {
				if !fortiAPIPatch(o["igmp-snooping"]) {
					return fmt.Errorf("Error reading igmp_snooping: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("n802_1x_settings", flattenSwitchControllerManagedSwitch8021XSettings(o["802-1X-settings"], d, "n802_1x_settings", sv)); err != nil {
			if !fortiAPIPatch(o["802-1X-settings"]) {
				return fmt.Errorf("Error reading n802_1x_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("n802_1x_settings"); ok {
			if err = d.Set("n802_1x_settings", flattenSwitchControllerManagedSwitch8021XSettings(o["802-1X-settings"], d, "n802_1x_settings", sv)); err != nil {
				if !fortiAPIPatch(o["802-1X-settings"]) {
					return fmt.Errorf("Error reading n802_1x_settings: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerManagedSwitchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerManagedSwitchSwitchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchAccessProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPurdueLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan1Peer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan1Admin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan2Peer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFswWan2Admin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPoePreStandardDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPoeDetectionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPoeLldpDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDirectlyConnected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMaxAllowedTrunkMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPreProvisioned(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchL3Discovered(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMgmtMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchTunnelDiscovered(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchTdrSupported(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDynamicCapability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchDeviceTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchDhcp_Opt43_Key(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMclagIgmpSnoopingAware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDynamicallyDiscovered(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPtpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPtpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRadiusNasIpOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRadiusNasIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadMclag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchRouteOffloadRouterVlanName(d, i["vlan_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "router_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["router-ip"], _ = expandSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(d, i["router_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouterVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRouteOffloadRouterRouterIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchVlanVlanName(d, i["vlan_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assignment_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["assignment-priority"], _ = expandSwitchControllerManagedSwitchVlanAssignmentPriority(d, i["assignment_priority"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchVlanVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchVlanAssignmentPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOwnerVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFlowIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStagedImageVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDelayedRestartTrigger(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFirmwareProvision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFirmwareProvisionVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-name"], _ = expandSwitchControllerManagedSwitchPortsPortName(d, i["port_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-owner"], _ = expandSwitchControllerManagedSwitchPortsPortOwner(d, i["port_owner"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["switch-id"], _ = expandSwitchControllerManagedSwitchPortsSwitchId(d, i["switch_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["speed"], _ = expandSwitchControllerManagedSwitchPortsSpeed(d, i["speed"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed_mask"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["speed-mask"], _ = expandSwitchControllerManagedSwitchPortsSpeedMask(d, i["speed_mask"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerManagedSwitchPortsStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-status"], _ = expandSwitchControllerManagedSwitchPortsPoeStatus(d, i["poe_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-source-guard"], _ = expandSwitchControllerManagedSwitchPortsIpSourceGuard(d, i["ip_source_guard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ptp-status"], _ = expandSwitchControllerManagedSwitchPortsPtpStatus(d, i["ptp_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ptp-policy"], _ = expandSwitchControllerManagedSwitchPortsPtpPolicy(d, i["ptp_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["aggregator-mode"], _ = expandSwitchControllerManagedSwitchPortsAggregatorMode(d, i["aggregator_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flapguard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flapguard"], _ = expandSwitchControllerManagedSwitchPortsFlapguard(d, i["flapguard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flap-rate"], _ = expandSwitchControllerManagedSwitchPortsFlapRate(d, i["flap_rate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flap-duration"], _ = expandSwitchControllerManagedSwitchPortsFlapDuration(d, i["flap_duration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flap_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flap-timeout"], _ = expandSwitchControllerManagedSwitchPortsFlapTimeout(d, i["flap_timeout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rpvst-port"], _ = expandSwitchControllerManagedSwitchPortsRpvstPort(d, i["rpvst_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-pre-standard-detection"], _ = expandSwitchControllerManagedSwitchPortsPoePreStandardDetection(d, i["poe_pre_standard_detection"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_number"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-number"], _ = expandSwitchControllerManagedSwitchPortsPortNumber(d, i["port_number"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_prefix_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-prefix-type"], _ = expandSwitchControllerManagedSwitchPortsPortPrefixType(d, i["port_prefix_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fortilink-port"], _ = expandSwitchControllerManagedSwitchPortsFortilinkPort(d, i["fortilink_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["link-status"], _ = expandSwitchControllerManagedSwitchPortsLinkStatus(d, i["link_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_capable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-capable"], _ = expandSwitchControllerManagedSwitchPortsPoeCapable(d, i["poe_capable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stacking_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stacking-port"], _ = expandSwitchControllerManagedSwitchPortsStackingPort(d, i["stacking_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["p2p-port"], _ = expandSwitchControllerManagedSwitchPortsP2PPort(d, i["p2p_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mclag-icl-port"], _ = expandSwitchControllerManagedSwitchPortsMclagIclPort(d, i["mclag_icl_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authenticated_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authenticated-port"], _ = expandSwitchControllerManagedSwitchPortsAuthenticatedPort(d, i["authenticated_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_auth_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["restricted-auth-port"], _ = expandSwitchControllerManagedSwitchPortsRestrictedAuthPort(d, i["restricted_auth_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encrypted_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["encrypted-port"], _ = expandSwitchControllerManagedSwitchPortsEncryptedPort(d, i["encrypted_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fiber_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fiber-port"], _ = expandSwitchControllerManagedSwitchPortsFiberPort(d, i["fiber_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["media-type"], _ = expandSwitchControllerManagedSwitchPortsMediaType(d, i["media_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_standard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-standard"], _ = expandSwitchControllerManagedSwitchPortsPoeStandard(d, i["poe_standard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_max_power"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-max-power"], _ = expandSwitchControllerManagedSwitchPortsPoeMaxPower(d, i["poe_max_power"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_mode_bt_cabable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-mode-bt-cabable"], _ = expandSwitchControllerManagedSwitchPortsPoeModeBtCabable(d, i["poe_mode_bt_cabable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-port-mode"], _ = expandSwitchControllerManagedSwitchPortsPoePortMode(d, i["poe_port_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-port-priority"], _ = expandSwitchControllerManagedSwitchPortsPoePortPriority(d, i["poe_port_priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_port_power"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poe-port-power"], _ = expandSwitchControllerManagedSwitchPortsPoePortPower(d, i["poe_port_power"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flags"], _ = expandSwitchControllerManagedSwitchPortsFlags(d, i["flags"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-port"], _ = expandSwitchControllerManagedSwitchPortsVirtualPort(d, i["virtual_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_local_trunk_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["isl-local-trunk-name"], _ = expandSwitchControllerManagedSwitchPortsIslLocalTrunkName(d, i["isl_local_trunk_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_port_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["isl-peer-port-name"], _ = expandSwitchControllerManagedSwitchPortsIslPeerPortName(d, i["isl_peer_port_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["isl-peer-device-name"], _ = expandSwitchControllerManagedSwitchPortsIslPeerDeviceName(d, i["isl_peer_device_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_sn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["isl-peer-device-sn"], _ = expandSwitchControllerManagedSwitchPortsIslPeerDeviceSn(d, i["isl_peer_device_sn"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_port_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fgt-peer-port-name"], _ = expandSwitchControllerManagedSwitchPortsFgtPeerPortName(d, i["fgt_peer_port_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_device_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fgt-peer-device-name"], _ = expandSwitchControllerManagedSwitchPortsFgtPeerDeviceName(d, i["fgt_peer_device_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan"], _ = expandSwitchControllerManagedSwitchPortsVlan(d, i["vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowed-vlans-all"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlansAll(d, i["allowed_vlans_all"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["allowed-vlans"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlans(d, i["allowed_vlans"], pre_append, sv)
		} else {
			tmp["allowed-vlans"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["untagged-vlans"], _ = expandSwitchControllerManagedSwitchPortsUntaggedVlans(d, i["untagged_vlans"], pre_append, sv)
		} else {
			tmp["untagged-vlans"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSwitchControllerManagedSwitchPortsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-mode"], _ = expandSwitchControllerManagedSwitchPortsAccessMode(d, i["access_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["matched-dpp-policy"], _ = expandSwitchControllerManagedSwitchPortsMatchedDppPolicy(d, i["matched_dpp_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "matched_dpp_intf_tags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["matched-dpp-intf-tags"], _ = expandSwitchControllerManagedSwitchPortsMatchedDppIntfTags(d, i["matched_dpp_intf_tags"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "acl_group"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["acl-group"], _ = expandSwitchControllerManagedSwitchPortsAclGroup(d, i["acl_group"], pre_append, sv)
		} else {
			tmp["acl-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortiswitch_acls"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["fortiswitch-acls"], _ = expandSwitchControllerManagedSwitchPortsFortiswitchAcls(d, i["fortiswitch_acls"], pre_append, sv)
		} else {
			tmp["fortiswitch-acls"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dhcp-snooping"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnooping(d, i["dhcp_snooping"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dhcp-snoop-option82-trust"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d, i["dhcp_snoop_option82_trust"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_override"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dhcp-snoop-option82-override"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d, i["dhcp_snoop_option82_override"], pre_append, sv)
		} else {
			tmp["dhcp-snoop-option82-override"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["arp-inspection-trust"], _ = expandSwitchControllerManagedSwitchPortsArpInspectionTrust(d, i["arp_inspection_trust"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping_flood_reports"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["igmp-snooping-flood-reports"], _ = expandSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(d, i["igmp_snooping_flood_reports"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcast_snooping_flood_traffic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mcast-snooping-flood-traffic"], _ = expandSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(d, i["mcast_snooping_flood_traffic"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["igmp-snooping"], _ = expandSwitchControllerManagedSwitchPortsIgmpSnooping(d, i["igmp_snooping"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["igmps-flood-reports"], _ = expandSwitchControllerManagedSwitchPortsIgmpsFloodReports(d, i["igmps_flood_reports"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["igmps-flood-traffic"], _ = expandSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(d, i["igmps_flood_traffic"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stp-state"], _ = expandSwitchControllerManagedSwitchPortsStpState(d, i["stp_state"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stp-root-guard"], _ = expandSwitchControllerManagedSwitchPortsStpRootGuard(d, i["stp_root_guard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stp-bpdu-guard"], _ = expandSwitchControllerManagedSwitchPortsStpBpduGuard(d, i["stp_bpdu_guard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stp-bpdu-guard-timeout"], _ = expandSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(d, i["stp_bpdu_guard_timeout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["edge-port"], _ = expandSwitchControllerManagedSwitchPortsEdgePort(d, i["edge_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["discard-mode"], _ = expandSwitchControllerManagedSwitchPortsDiscardMode(d, i["discard_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["packet-sampler"], _ = expandSwitchControllerManagedSwitchPortsPacketSampler(d, i["packet_sampler"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["packet-sample-rate"], _ = expandSwitchControllerManagedSwitchPortsPacketSampleRate(d, i["packet_sample_rate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sampler"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sflow-sampler"], _ = expandSwitchControllerManagedSwitchPortsSflowSampler(d, i["sflow_sampler"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sample_rate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sflow-sample-rate"], _ = expandSwitchControllerManagedSwitchPortsSflowSampleRate(d, i["sflow_sample_rate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sflow-counter-interval"], _ = expandSwitchControllerManagedSwitchPortsSflowCounterInterval(d, i["sflow_counter_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sample-direction"], _ = expandSwitchControllerManagedSwitchPortsSampleDirection(d, i["sample_direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fec-capable"], _ = expandSwitchControllerManagedSwitchPortsFecCapable(d, i["fec_capable"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fec-state"], _ = expandSwitchControllerManagedSwitchPortsFecState(d, i["fec_state"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flow-control"], _ = expandSwitchControllerManagedSwitchPortsFlowControl(d, i["flow_control"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pause-meter"], _ = expandSwitchControllerManagedSwitchPortsPauseMeter(d, i["pause_meter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pause-meter-resume"], _ = expandSwitchControllerManagedSwitchPortsPauseMeterResume(d, i["pause_meter_resume"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["loop-guard"], _ = expandSwitchControllerManagedSwitchPortsLoopGuard(d, i["loop_guard"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["loop-guard-timeout"], _ = expandSwitchControllerManagedSwitchPortsLoopGuardTimeout(d, i["loop_guard_timeout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-policy"], _ = expandSwitchControllerManagedSwitchPortsPortPolicy(d, i["port_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["qos-policy"], _ = expandSwitchControllerManagedSwitchPortsQosPolicy(d, i["qos_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "storm_control_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["storm-control-policy"], _ = expandSwitchControllerManagedSwitchPortsStormControlPolicy(d, i["storm_control_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-security-policy"], _ = expandSwitchControllerManagedSwitchPortsPortSecurityPolicy(d, i["port_security_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["export-to-pool"], _ = expandSwitchControllerManagedSwitchPortsExportToPool(d, i["export_to_pool"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["interface-tags"], _ = expandSwitchControllerManagedSwitchPortsInterfaceTags(d, i["interface_tags"], pre_append, sv)
		} else {
			tmp["interface-tags"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["export-tags"], _ = expandSwitchControllerManagedSwitchPortsExportTags(d, i["export_tags"], pre_append, sv)
		} else {
			tmp["export-tags"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool_flag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["export-to-pool_flag"], _ = expandSwitchControllerManagedSwitchPortsExportToPool_Flag(d, i["export_to_pool_flag"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["learning-limit"], _ = expandSwitchControllerManagedSwitchPortsLearningLimit(d, i["learning_limit"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sticky-mac"], _ = expandSwitchControllerManagedSwitchPortsStickyMac(d, i["sticky_mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lldp-status"], _ = expandSwitchControllerManagedSwitchPortsLldpStatus(d, i["lldp_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lldp-profile"], _ = expandSwitchControllerManagedSwitchPortsLldpProfile(d, i["lldp_profile"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["export-to"], _ = expandSwitchControllerManagedSwitchPortsExportTo(d, i["export_to"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac-addr"], _ = expandSwitchControllerManagedSwitchPortsMacAddr(d, i["mac_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-selection-criteria"], _ = expandSwitchControllerManagedSwitchPortsPortSelectionCriteria(d, i["port_selection_criteria"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSwitchControllerManagedSwitchPortsDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lacp-speed"], _ = expandSwitchControllerManagedSwitchPortsLacpSpeed(d, i["lacp_speed"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mode"], _ = expandSwitchControllerManagedSwitchPortsMode(d, i["mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bundle"], _ = expandSwitchControllerManagedSwitchPortsBundle(d, i["bundle"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["member-withdrawal-behavior"], _ = expandSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(d, i["member_withdrawal_behavior"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mclag"], _ = expandSwitchControllerManagedSwitchPortsMclag(d, i["mclag"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["min-bundle"], _ = expandSwitchControllerManagedSwitchPortsMinBundle(d, i["min_bundle"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-bundle"], _ = expandSwitchControllerManagedSwitchPortsMaxBundle(d, i["max_bundle"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["members"], _ = expandSwitchControllerManagedSwitchPortsMembers(d, i["members"], pre_append, sv)
		} else {
			tmp["members"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortOwner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSwitchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSpeedMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIpSourceGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPtpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPtpPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAggregatorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapguard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlapTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsRpvstPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePreStandardDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortPrefixType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFortilinkPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLinkStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeCapable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStackingPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsP2PPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMclagIclPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAuthenticatedPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsRestrictedAuthPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsEncryptedPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFiberPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMediaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeStandard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeMaxPower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoeModeBtCabable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePortPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPoePortPower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsVirtualPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslLocalTrunkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslPeerPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslPeerDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIslPeerDeviceSn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFgtPeerPortName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFgtPeerDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlansAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlansVlanName(d, i["vlan_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchPortsUntaggedVlansVlanName(d, i["vlan_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsUntaggedVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAccessMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMatchedDppPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMatchedDppIntfTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAclGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSwitchControllerManagedSwitchPortsAclGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsAclGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFortiswitchAcls(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandSwitchControllerManagedSwitchPortsFortiswitchAclsId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsFortiswitchAclsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Override(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(d, i["vlan_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["circuit-id"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(d, i["circuit_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-id"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(d, i["remote_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideCircuitId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82OverrideRemoteId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsArpInspectionTrust(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpSnoopingFloodReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMcastSnoopingFloodTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpsFloodReports(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpRootGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpBpduGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsEdgePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDiscardMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPacketSampler(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPacketSampleRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSflowSampler(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSflowSampleRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSflowCounterInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsSampleDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFecCapable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFecState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsFlowControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPauseMeter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPauseMeterResume(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLoopGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLoopGuardTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsQosPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStormControlPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortSecurityPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsExportToPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsInterfaceTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["tag-name"], _ = expandSwitchControllerManagedSwitchPortsInterfaceTagsTagName(d, i["tag_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsInterfaceTagsTagName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsExportTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["tag-name"], _ = expandSwitchControllerManagedSwitchPortsExportTagsTagName(d, i["tag_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsExportTagsTagName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsExportToPool_Flag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLearningLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsStickyMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLldpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLldpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsExportTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMacAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsPortSelectionCriteria(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsLacpSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsBundle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMclag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMinBundle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMaxBundle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["member-name"], _ = expandSwitchControllerManagedSwitchPortsMembersMemberName(d, i["member_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsMembersMemberName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandSwitchControllerManagedSwitchIpSourceGuardPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSwitchControllerManagedSwitchIpSourceGuardDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "binding_entry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["binding-entry"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d, i["binding_entry"], pre_append, sv)
		} else {
			tmp["binding-entry"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "entry_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["entry-name"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(d, i["entry_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(d, i["mac"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandSwitchControllerManagedSwitchStpSettingsLocalOverride(d, i["local_override"], pre_append, sv)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok {
		result["name"], _ = expandSwitchControllerManagedSwitchStpSettingsName(d, i["name"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSwitchControllerManagedSwitchStpSettingsStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "revision"
	if _, ok := d.GetOk(pre_append); ok {
		result["revision"], _ = expandSwitchControllerManagedSwitchStpSettingsRevision(d, i["revision"], pre_append, sv)
	}
	pre_append = pre + ".0." + "hello_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["hello-time"], _ = expandSwitchControllerManagedSwitchStpSettingsHelloTime(d, i["hello_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "forward_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["forward-time"], _ = expandSwitchControllerManagedSwitchStpSettingsForwardTime(d, i["forward_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_age"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-age"], _ = expandSwitchControllerManagedSwitchStpSettingsMaxAge(d, i["max_age"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_hops"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-hops"], _ = expandSwitchControllerManagedSwitchStpSettingsMaxHops(d, i["max_hops"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pending_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["pending-timer"], _ = expandSwitchControllerManagedSwitchStpSettingsPendingTimer(d, i["pending_timer"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStpSettingsLocalOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsRevision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsHelloTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsForwardTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsMaxHops(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpSettingsPendingTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpInstance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchStpInstanceId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSwitchControllerManagedSwitchStpInstancePriority(d, i["priority"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStpInstanceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStpInstancePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "engine_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["engine-id"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoEngineId(d, i["engine_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "description"
	if _, ok := d.GetOk(pre_append); ok {
		result["description"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoDescription(d, i["description"], pre_append, sv)
	}
	pre_append = pre + ".0." + "contact_info"
	if _, ok := d.GetOk(pre_append); ok {
		result["contact-info"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoContactInfo(d, i["contact_info"], pre_append, sv)
	}
	pre_append = pre + ".0." + "location"
	if _, ok := d.GetOk(pre_append); ok {
		result["location"], _ = expandSwitchControllerManagedSwitchSnmpSysinfoLocation(d, i["location"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "trap_high_cpu_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["trap-high-cpu-threshold"], _ = expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(d, i["trap_high_cpu_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "trap_low_memory_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["trap-low-memory-threshold"], _ = expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(d, i["trap_low_memory_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "trap_log_full_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["trap-log-full-threshold"], _ = expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(d, i["trap_log_full_threshold"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLowMemoryThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpTrapThresholdTrapLogFullThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchSnmpCommunityId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerManagedSwitchSnmpCommunityName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["hosts"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHosts(d, i["hosts"], pre_append, sv)
		} else {
			tmp["hosts"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["query-v1-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(d, i["query_v1_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["query-v1-port"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(d, i["query_v1_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["query-v2c-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(d, i["query_v2c_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["query-v2c-port"], _ = expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(d, i["query_v2c_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trap-v1-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(d, i["trap_v1_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_lport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trap-v1-lport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(d, i["trap_v1_lport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_rport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trap-v1-rport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(d, i["trap_v1_rport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trap-v2c-status"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(d, i["trap_v2c_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_lport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trap-v2c-lport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(d, i["trap_v2c_lport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_rport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["trap-v2c-rport"], _ = expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(d, i["trap_v2c_rport"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "events"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["events"], _ = expandSwitchControllerManagedSwitchSnmpCommunityEvents(d, i["events"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHostsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchSnmpCommunityHostsIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpCommunityEvents(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchOverrideSnmpUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerManagedSwitchSnmpUserName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["queries"], _ = expandSwitchControllerManagedSwitchSnmpUserQueries(d, i["queries"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["query-port"], _ = expandSwitchControllerManagedSwitchSnmpUserQueryPort(d, i["query_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["security-level"], _ = expandSwitchControllerManagedSwitchSnmpUserSecurityLevel(d, i["security_level"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-proto"], _ = expandSwitchControllerManagedSwitchSnmpUserAuthProto(d, i["auth_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-pwd"], _ = expandSwitchControllerManagedSwitchSnmpUserAuthPwd(d, i["auth_pwd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priv-proto"], _ = expandSwitchControllerManagedSwitchSnmpUserPrivProto(d, i["priv_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priv-pwd"], _ = expandSwitchControllerManagedSwitchSnmpUserPrivPwd(d, i["priv_pwd"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSnmpUserName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserQueryPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchQosDropPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchQosRedProbability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchStpSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSwitchControllerManagedSwitchSwitchStpSettingsStatus(d, i["status"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSwitchStpSettingsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandSwitchControllerManagedSwitchSwitchLogLocalOverride(d, i["local_override"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSwitchControllerManagedSwitchSwitchLogStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandSwitchControllerManagedSwitchSwitchLogSeverity(d, i["severity"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchSwitchLogLocalOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchLogStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchSwitchLogSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerManagedSwitchRemoteLogName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerManagedSwitchRemoteLogStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandSwitchControllerManagedSwitchRemoteLogServer(d, i["server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandSwitchControllerManagedSwitchRemoteLogPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandSwitchControllerManagedSwitchRemoteLogSeverity(d, i["severity"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csv"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["csv"], _ = expandSwitchControllerManagedSwitchRemoteLogCsv(d, i["csv"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "facility"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["facility"], _ = expandSwitchControllerManagedSwitchRemoteLogFacility(d, i["facility"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchRemoteLogName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogCsv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchRemoteLogFacility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandSwitchControllerManagedSwitchStormControlLocalOverride(d, i["local_override"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rate"
	if _, ok := d.GetOk(pre_append); ok {
		result["rate"], _ = expandSwitchControllerManagedSwitchStormControlRate(d, i["rate"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unknown_unicast"
	if _, ok := d.GetOk(pre_append); ok {
		result["unknown-unicast"], _ = expandSwitchControllerManagedSwitchStormControlUnknownUnicast(d, i["unknown_unicast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unknown_multicast"
	if _, ok := d.GetOk(pre_append); ok {
		result["unknown-multicast"], _ = expandSwitchControllerManagedSwitchStormControlUnknownMulticast(d, i["unknown_multicast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "broadcast"
	if _, ok := d.GetOk(pre_append); ok {
		result["broadcast"], _ = expandSwitchControllerManagedSwitchStormControlBroadcast(d, i["broadcast"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStormControlLocalOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlUnknownUnicast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStormControlBroadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerManagedSwitchMirrorName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerManagedSwitchMirrorStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switching_packet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["switching-packet"], _ = expandSwitchControllerManagedSwitchMirrorSwitchingPacket(d, i["switching_packet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandSwitchControllerManagedSwitchMirrorDst(d, i["dst"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ingress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-ingress"], _ = expandSwitchControllerManagedSwitchMirrorSrcIngress(d, i["src_ingress"], pre_append, sv)
		} else {
			tmp["src-ingress"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_egress"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["src-egress"], _ = expandSwitchControllerManagedSwitchMirrorSrcEgress(d, i["src_egress"], pre_append, sv)
		} else {
			tmp["src-egress"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchMirrorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSwitchingPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcIngress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSwitchControllerManagedSwitchMirrorSrcIngressName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcIngressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcEgress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSwitchControllerManagedSwitchMirrorSrcEgressName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcEgressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSwitchControllerManagedSwitchStaticMacId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSwitchControllerManagedSwitchStaticMacType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan"], _ = expandSwitchControllerManagedSwitchStaticMacVlan(d, i["vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchStaticMacMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSwitchControllerManagedSwitchStaticMacInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSwitchControllerManagedSwitchStaticMacDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchStaticMacId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchStaticMacDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchCustomCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["command-entry"], _ = expandSwitchControllerManagedSwitchCustomCommandCommandEntry(d, i["command_entry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["command-name"], _ = expandSwitchControllerManagedSwitchCustomCommandCommandName(d, i["command_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchCustomCommandCommandEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(d, i["vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(d, i["port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchDhcpSnoopingStaticClientPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(d, i["local_override"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aging_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["aging-time"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingAgingTime(d, i["aging_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "flood_unknown_multicast"
	if _, ok := d.GetOk(pre_append); ok {
		result["flood-unknown-multicast"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(d, i["flood_unknown_multicast"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlans"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlans"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlans(d, i["vlans"], pre_append, sv)
	} else {
		result["vlans"] = make([]string, 0)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingLocalOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingAgingTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingFloodUnknownMulticast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(d, i["vlan_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["proxy"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(d, i["proxy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["querier"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(d, i["querier"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "querier_addr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["querier-addr"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(d, i["querier_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["version"], _ = expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(d, i["version"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansQuerierAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchIgmpSnoopingVlansVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "local_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-override"], _ = expandSwitchControllerManagedSwitch8021XSettingsLocalOverride(d, i["local_override"], pre_append, sv)
	}
	pre_append = pre + ".0." + "link_down_auth"
	if _, ok := d.GetOk(pre_append); ok {
		result["link-down-auth"], _ = expandSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(d, i["link_down_auth"], pre_append, sv)
	}
	pre_append = pre + ".0." + "reauth_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["reauth-period"], _ = expandSwitchControllerManagedSwitch8021XSettingsReauthPeriod(d, i["reauth_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_reauth_attempt"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-reauth-attempt"], _ = expandSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(d, i["max_reauth_attempt"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tx_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["tx-period"], _ = expandSwitchControllerManagedSwitch8021XSettingsTxPeriod(d, i["tx_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mab_reauth"
	if _, ok := d.GetOk(pre_append); ok {
		result["mab-reauth"], _ = expandSwitchControllerManagedSwitch8021XSettingsMabReauth(d, i["mab_reauth"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_username_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-username-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(d, i["mac_username_delimiter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_password_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-password-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(d, i["mac_password_delimiter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_calling_station_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-calling-station-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(d, i["mac_calling_station_delimiter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_called_station_delimiter"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-called-station-delimiter"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(d, i["mac_called_station_delimiter"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mac_case"
	if _, ok := d.GetOk(pre_append); ok {
		result["mac-case"], _ = expandSwitchControllerManagedSwitch8021XSettingsMacCase(d, i["mac_case"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsLocalOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsLinkDownAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsReauthPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMaxReauthAttempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsTxPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMabReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitch8021XSettingsMacCase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerManagedSwitch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("switch_id"); ok {
		t, err := expandSwitchControllerManagedSwitchSwitchId(d, v, "switch_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-id"] = t
		}
	}

	if v, ok := d.GetOk("sn"); ok {
		t, err := expandSwitchControllerManagedSwitchSn(d, v, "sn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sn"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerManagedSwitchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerManagedSwitchDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("switch_profile"); ok {
		t, err := expandSwitchControllerManagedSwitchSwitchProfile(d, v, "switch_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-profile"] = t
		}
	}

	if v, ok := d.GetOk("access_profile"); ok {
		t, err := expandSwitchControllerManagedSwitchAccessProfile(d, v, "access_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-profile"] = t
		}
	}

	if v, ok := d.GetOk("purdue_level"); ok {
		t, err := expandSwitchControllerManagedSwitchPurdueLevel(d, v, "purdue_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["purdue-level"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan1_peer"); ok {
		t, err := expandSwitchControllerManagedSwitchFswWan1Peer(d, v, "fsw_wan1_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan1-peer"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan1_admin"); ok {
		t, err := expandSwitchControllerManagedSwitchFswWan1Admin(d, v, "fsw_wan1_admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan1-admin"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan2_peer"); ok {
		t, err := expandSwitchControllerManagedSwitchFswWan2Peer(d, v, "fsw_wan2_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan2-peer"] = t
		}
	}

	if v, ok := d.GetOk("fsw_wan2_admin"); ok {
		t, err := expandSwitchControllerManagedSwitchFswWan2Admin(d, v, "fsw_wan2_admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw-wan2-admin"] = t
		}
	}

	if v, ok := d.GetOk("poe_pre_standard_detection"); ok {
		t, err := expandSwitchControllerManagedSwitchPoePreStandardDetection(d, v, "poe_pre_standard_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-pre-standard-detection"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok {
		t, err := expandSwitchControllerManagedSwitchDhcpServerAccessList(d, v, "dhcp_server_access_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server-access-list"] = t
		}
	}

	if v, ok := d.GetOkExists("poe_detection_type"); ok {
		t, err := expandSwitchControllerManagedSwitchPoeDetectionType(d, v, "poe_detection_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-detection-type"] = t
		}
	}

	if v, ok := d.GetOk("poe_lldp_detection"); ok {
		t, err := expandSwitchControllerManagedSwitchPoeLldpDetection(d, v, "poe_lldp_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-lldp-detection"] = t
		}
	}

	if v, ok := d.GetOkExists("directly_connected"); ok {
		t, err := expandSwitchControllerManagedSwitchDirectlyConnected(d, v, "directly_connected", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["directly-connected"] = t
		}
	}

	if v, ok := d.GetOkExists("version"); ok {
		t, err := expandSwitchControllerManagedSwitchVersion(d, v, "version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOkExists("max_allowed_trunk_members"); ok {
		t, err := expandSwitchControllerManagedSwitchMaxAllowedTrunkMembers(d, v, "max_allowed_trunk_members", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-allowed-trunk-members"] = t
		}
	}

	if v, ok := d.GetOkExists("pre_provisioned"); ok {
		t, err := expandSwitchControllerManagedSwitchPreProvisioned(d, v, "pre_provisioned", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-provisioned"] = t
		}
	}

	if v, ok := d.GetOkExists("l3_discovered"); ok {
		t, err := expandSwitchControllerManagedSwitchL3Discovered(d, v, "l3_discovered", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-discovered"] = t
		}
	}

	if v, ok := d.GetOkExists("mgmt_mode"); ok {
		t, err := expandSwitchControllerManagedSwitchMgmtMode(d, v, "mgmt_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-mode"] = t
		}
	}

	if v, ok := d.GetOkExists("tunnel_discovered"); ok {
		t, err := expandSwitchControllerManagedSwitchTunnelDiscovered(d, v, "tunnel_discovered", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-discovered"] = t
		}
	}

	if v, ok := d.GetOk("tdr_supported"); ok {
		t, err := expandSwitchControllerManagedSwitchTdrSupported(d, v, "tdr_supported", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tdr-supported"] = t
		}
	}

	if v, ok := d.GetOkExists("dynamic_capability"); ok {
		t, err := expandSwitchControllerManagedSwitchDynamicCapability(d, v, "dynamic_capability", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			new_version_map := map[string][]string{
				">=": []string{"6.4.2"},
			}
			if i2ss2arrFortiAPIUpgrade(sv, new_version_map) == true {
				obj["dynamic-capability"] = fmt.Sprintf("%v", t)
			} else {
				obj["dynamic-capability"] = t
			}
		}
	}

	if v, ok := d.GetOk("switch_device_tag"); ok {
		t, err := expandSwitchControllerManagedSwitchSwitchDeviceTag(d, v, "switch_device_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-device-tag"] = t
		}
	}

	if v, ok := d.GetOk("switch_dhcp_opt43_key"); ok {
		t, err := expandSwitchControllerManagedSwitchSwitchDhcp_Opt43_Key(d, v, "switch_dhcp_opt43_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-dhcp_opt43_key"] = t
		}
	}

	if v, ok := d.GetOk("mclag_igmp_snooping_aware"); ok {
		t, err := expandSwitchControllerManagedSwitchMclagIgmpSnoopingAware(d, v, "mclag_igmp_snooping_aware", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mclag-igmp-snooping-aware"] = t
		}
	}

	if v, ok := d.GetOkExists("dynamically_discovered"); ok {
		t, err := expandSwitchControllerManagedSwitchDynamicallyDiscovered(d, v, "dynamically_discovered", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamically-discovered"] = t
		}
	}

	if v, ok := d.GetOk("ptp_status"); ok {
		t, err := expandSwitchControllerManagedSwitchPtpStatus(d, v, "ptp_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-status"] = t
		}
	}

	if v, ok := d.GetOk("ptp_profile"); ok {
		t, err := expandSwitchControllerManagedSwitchPtpProfile(d, v, "ptp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptp-profile"] = t
		}
	}

	if v, ok := d.GetOk("radius_nas_ip_override"); ok {
		t, err := expandSwitchControllerManagedSwitchRadiusNasIpOverride(d, v, "radius_nas_ip_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-nas-ip-override"] = t
		}
	}

	if v, ok := d.GetOk("radius_nas_ip"); ok {
		t, err := expandSwitchControllerManagedSwitchRadiusNasIp(d, v, "radius_nas_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("route_offload"); ok {
		t, err := expandSwitchControllerManagedSwitchRouteOffload(d, v, "route_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload"] = t
		}
	}

	if v, ok := d.GetOk("route_offload_mclag"); ok {
		t, err := expandSwitchControllerManagedSwitchRouteOffloadMclag(d, v, "route_offload_mclag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload-mclag"] = t
		}
	}

	if v, ok := d.GetOk("route_offload_router"); ok || d.HasChange("route_offload_router") {
		t, err := expandSwitchControllerManagedSwitchRouteOffloadRouter(d, v, "route_offload_router", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-offload-router"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok || d.HasChange("vlan") {
		t, err := expandSwitchControllerManagedSwitchVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSwitchControllerManagedSwitchType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("owner_vdom"); ok {
		t, err := expandSwitchControllerManagedSwitchOwnerVdom(d, v, "owner_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owner-vdom"] = t
		}
	}

	if v, ok := d.GetOk("flow_identity"); ok {
		t, err := expandSwitchControllerManagedSwitchFlowIdentity(d, v, "flow_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flow-identity"] = t
		}
	}

	if v, ok := d.GetOk("staged_image_version"); ok {
		t, err := expandSwitchControllerManagedSwitchStagedImageVersion(d, v, "staged_image_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["staged-image-version"] = t
		}
	}

	if v, ok := d.GetOkExists("delayed_restart_trigger"); ok {
		t, err := expandSwitchControllerManagedSwitchDelayedRestartTrigger(d, v, "delayed_restart_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delayed-restart-trigger"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision"); ok {
		t, err := expandSwitchControllerManagedSwitchFirmwareProvision(d, v, "firmware_provision", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_version"); ok {
		t, err := expandSwitchControllerManagedSwitchFirmwareProvisionVersion(d, v, "firmware_provision_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-version"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok {
		t, err := expandSwitchControllerManagedSwitchFirmwareProvisionLatest(d, v, "firmware_provision_latest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandSwitchControllerManagedSwitchPorts(d, v, "ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("ip_source_guard"); ok || d.HasChange("ip_source_guard") {
		t, err := expandSwitchControllerManagedSwitchIpSourceGuard(d, v, "ip_source_guard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-source-guard"] = t
		}
	}

	if v, ok := d.GetOk("stp_settings"); ok {
		t, err := expandSwitchControllerManagedSwitchStpSettings(d, v, "stp_settings", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-settings"] = t
		}
	}

	if v, ok := d.GetOk("stp_instance"); ok || d.HasChange("stp_instance") {
		t, err := expandSwitchControllerManagedSwitchStpInstance(d, v, "stp_instance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stp-instance"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_sysinfo"); ok {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpSysinfo(d, v, "override_snmp_sysinfo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("snmp_sysinfo"); ok {
		t, err := expandSwitchControllerManagedSwitchSnmpSysinfo(d, v, "snmp_sysinfo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-sysinfo"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_trap_threshold"); ok {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpTrapThreshold(d, v, "override_snmp_trap_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-trap-threshold"] = t
		}
	}

	if v, ok := d.GetOk("snmp_trap_threshold"); ok {
		t, err := expandSwitchControllerManagedSwitchSnmpTrapThreshold(d, v, "snmp_trap_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-trap-threshold"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_community"); ok {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpCommunity(d, v, "override_snmp_community", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-community"] = t
		}
	}

	if v, ok := d.GetOk("snmp_community"); ok || d.HasChange("snmp_community") {
		t, err := expandSwitchControllerManagedSwitchSnmpCommunity(d, v, "snmp_community", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-community"] = t
		}
	}

	if v, ok := d.GetOk("override_snmp_user"); ok {
		t, err := expandSwitchControllerManagedSwitchOverrideSnmpUser(d, v, "override_snmp_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-snmp-user"] = t
		}
	}

	if v, ok := d.GetOk("snmp_user"); ok || d.HasChange("snmp_user") {
		t, err := expandSwitchControllerManagedSwitchSnmpUser(d, v, "snmp_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-user"] = t
		}
	}

	if v, ok := d.GetOk("qos_drop_policy"); ok {
		t, err := expandSwitchControllerManagedSwitchQosDropPolicy(d, v, "qos_drop_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-drop-policy"] = t
		}
	}

	if v, ok := d.GetOkExists("qos_red_probability"); ok {
		t, err := expandSwitchControllerManagedSwitchQosRedProbability(d, v, "qos_red_probability", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-red-probability"] = t
		}
	}

	if v, ok := d.GetOk("switch_stp_settings"); ok {
		t, err := expandSwitchControllerManagedSwitchSwitchStpSettings(d, v, "switch_stp_settings", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-stp-settings"] = t
		}
	}

	if v, ok := d.GetOk("switch_log"); ok {
		t, err := expandSwitchControllerManagedSwitchSwitchLog(d, v, "switch_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-log"] = t
		}
	}

	if v, ok := d.GetOk("remote_log"); ok || d.HasChange("remote_log") {
		t, err := expandSwitchControllerManagedSwitchRemoteLog(d, v, "remote_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-log"] = t
		}
	}

	if v, ok := d.GetOk("storm_control"); ok {
		t, err := expandSwitchControllerManagedSwitchStormControl(d, v, "storm_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storm-control"] = t
		}
	}

	if v, ok := d.GetOk("mirror"); ok || d.HasChange("mirror") {
		t, err := expandSwitchControllerManagedSwitchMirror(d, v, "mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mirror"] = t
		}
	}

	if v, ok := d.GetOk("static_mac"); ok || d.HasChange("static_mac") {
		t, err := expandSwitchControllerManagedSwitchStaticMac(d, v, "static_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["static-mac"] = t
		}
	}

	if v, ok := d.GetOk("custom_command"); ok || d.HasChange("custom_command") {
		t, err := expandSwitchControllerManagedSwitchCustomCommand(d, v, "custom_command", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-command"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_snooping_static_client"); ok || d.HasChange("dhcp_snooping_static_client") {
		t, err := expandSwitchControllerManagedSwitchDhcpSnoopingStaticClient(d, v, "dhcp_snooping_static_client", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-snooping-static-client"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok {
		t, err := expandSwitchControllerManagedSwitchIgmpSnooping(d, v, "igmp_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("n802_1x_settings"); ok {
		t, err := expandSwitchControllerManagedSwitch8021XSettings(d, v, "n802_1x_settings", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802-1X-settings"] = t
		}
	}

	return &obj, nil
}
