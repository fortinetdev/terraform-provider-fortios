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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"ports": &schema.Schema{
				Type:     schema.TypeList,
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
							Type:     schema.TypeList,
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
							Type:     schema.TypeList,
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
						"arp_inspection_trust": &schema.Schema{
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
							ValidateFunc: validation.IntBetween(1, 255),
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
						"export_tags": &schema.Schema{
							Type:     schema.TypeList,
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
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
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
							ValidateFunc: validation.IntBetween(1, 10000000),
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
						"src_egress": &schema.Schema{
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
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeList,
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
					},
				},
			},
			"n802_1x_settings": &schema.Schema{
				Type:     schema.TypeList,
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
							ValidateFunc: validation.IntBetween(1, 1440),
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

func flattenSwitchControllerManagedSwitchPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_name"
		if _, ok := i["port-name"]; ok {

			tmp["port_name"] = flattenSwitchControllerManagedSwitchPortsPortName(i["port-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_owner"
		if _, ok := i["port-owner"]; ok {

			tmp["port_owner"] = flattenSwitchControllerManagedSwitchPortsPortOwner(i["port-owner"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {

			tmp["switch_id"] = flattenSwitchControllerManagedSwitchPortsSwitchId(i["switch-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed"
		if _, ok := i["speed"]; ok {

			tmp["speed"] = flattenSwitchControllerManagedSwitchPortsSpeed(i["speed"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "speed_mask"
		if _, ok := i["speed-mask"]; ok {

			tmp["speed_mask"] = flattenSwitchControllerManagedSwitchPortsSpeedMask(i["speed-mask"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchControllerManagedSwitchPortsStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_status"
		if _, ok := i["poe-status"]; ok {

			tmp["poe_status"] = flattenSwitchControllerManagedSwitchPortsPoeStatus(i["poe-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_source_guard"
		if _, ok := i["ip-source-guard"]; ok {

			tmp["ip_source_guard"] = flattenSwitchControllerManagedSwitchPortsIpSourceGuard(i["ip-source-guard"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_policy"
		if _, ok := i["ptp-policy"]; ok {

			tmp["ptp_policy"] = flattenSwitchControllerManagedSwitchPortsPtpPolicy(i["ptp-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := i["aggregator-mode"]; ok {

			tmp["aggregator_mode"] = flattenSwitchControllerManagedSwitchPortsAggregatorMode(i["aggregator-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpvst_port"
		if _, ok := i["rpvst-port"]; ok {

			tmp["rpvst_port"] = flattenSwitchControllerManagedSwitchPortsRpvstPort(i["rpvst-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_pre_standard_detection"
		if _, ok := i["poe-pre-standard-detection"]; ok {

			tmp["poe_pre_standard_detection"] = flattenSwitchControllerManagedSwitchPortsPoePreStandardDetection(i["poe-pre-standard-detection"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_number"
		if _, ok := i["port-number"]; ok {

			tmp["port_number"] = flattenSwitchControllerManagedSwitchPortsPortNumber(i["port-number"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_prefix_type"
		if _, ok := i["port-prefix-type"]; ok {

			tmp["port_prefix_type"] = flattenSwitchControllerManagedSwitchPortsPortPrefixType(i["port-prefix-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fortilink_port"
		if _, ok := i["fortilink-port"]; ok {

			tmp["fortilink_port"] = flattenSwitchControllerManagedSwitchPortsFortilinkPort(i["fortilink-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poe_capable"
		if _, ok := i["poe-capable"]; ok {

			tmp["poe_capable"] = flattenSwitchControllerManagedSwitchPortsPoeCapable(i["poe-capable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stacking_port"
		if _, ok := i["stacking-port"]; ok {

			tmp["stacking_port"] = flattenSwitchControllerManagedSwitchPortsStackingPort(i["stacking-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "p2p_port"
		if _, ok := i["p2p-port"]; ok {

			tmp["p2p_port"] = flattenSwitchControllerManagedSwitchPortsP2PPort(i["p2p-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag_icl_port"
		if _, ok := i["mclag-icl-port"]; ok {

			tmp["mclag_icl_port"] = flattenSwitchControllerManagedSwitchPortsMclagIclPort(i["mclag-icl-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fiber_port"
		if _, ok := i["fiber-port"]; ok {

			tmp["fiber_port"] = flattenSwitchControllerManagedSwitchPortsFiberPort(i["fiber-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := i["media-type"]; ok {

			tmp["media_type"] = flattenSwitchControllerManagedSwitchPortsMediaType(i["media-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {

			tmp["flags"] = flattenSwitchControllerManagedSwitchPortsFlags(i["flags"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_port"
		if _, ok := i["virtual-port"]; ok {

			tmp["virtual_port"] = flattenSwitchControllerManagedSwitchPortsVirtualPort(i["virtual-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_local_trunk_name"
		if _, ok := i["isl-local-trunk-name"]; ok {

			tmp["isl_local_trunk_name"] = flattenSwitchControllerManagedSwitchPortsIslLocalTrunkName(i["isl-local-trunk-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_port_name"
		if _, ok := i["isl-peer-port-name"]; ok {

			tmp["isl_peer_port_name"] = flattenSwitchControllerManagedSwitchPortsIslPeerPortName(i["isl-peer-port-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "isl_peer_device_name"
		if _, ok := i["isl-peer-device-name"]; ok {

			tmp["isl_peer_device_name"] = flattenSwitchControllerManagedSwitchPortsIslPeerDeviceName(i["isl-peer-device-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_port_name"
		if _, ok := i["fgt-peer-port-name"]; ok {

			tmp["fgt_peer_port_name"] = flattenSwitchControllerManagedSwitchPortsFgtPeerPortName(i["fgt-peer-port-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fgt_peer_device_name"
		if _, ok := i["fgt-peer-device-name"]; ok {

			tmp["fgt_peer_device_name"] = flattenSwitchControllerManagedSwitchPortsFgtPeerDeviceName(i["fgt-peer-device-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {

			tmp["vlan"] = flattenSwitchControllerManagedSwitchPortsVlan(i["vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans_all"
		if _, ok := i["allowed-vlans-all"]; ok {

			tmp["allowed_vlans_all"] = flattenSwitchControllerManagedSwitchPortsAllowedVlansAll(i["allowed-vlans-all"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_vlans"
		if _, ok := i["allowed-vlans"]; ok {

			tmp["allowed_vlans"] = flattenSwitchControllerManagedSwitchPortsAllowedVlans(i["allowed-vlans"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := i["untagged-vlans"]; ok {

			tmp["untagged_vlans"] = flattenSwitchControllerManagedSwitchPortsUntaggedVlans(i["untagged-vlans"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSwitchControllerManagedSwitchPortsType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_mode"
		if _, ok := i["access-mode"]; ok {

			tmp["access_mode"] = flattenSwitchControllerManagedSwitchPortsAccessMode(i["access-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := i["dhcp-snooping"]; ok {

			tmp["dhcp_snooping"] = flattenSwitchControllerManagedSwitchPortsDhcpSnooping(i["dhcp-snooping"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := i["dhcp-snoop-option82-trust"]; ok {

			tmp["dhcp_snoop_option82_trust"] = flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(i["dhcp-snoop-option82-trust"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := i["arp-inspection-trust"]; ok {

			tmp["arp_inspection_trust"] = flattenSwitchControllerManagedSwitchPortsArpInspectionTrust(i["arp-inspection-trust"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp_snooping"
		if _, ok := i["igmp-snooping"]; ok {

			tmp["igmp_snooping"] = flattenSwitchControllerManagedSwitchPortsIgmpSnooping(i["igmp-snooping"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_reports"
		if _, ok := i["igmps-flood-reports"]; ok {

			tmp["igmps_flood_reports"] = flattenSwitchControllerManagedSwitchPortsIgmpsFloodReports(i["igmps-flood-reports"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmps_flood_traffic"
		if _, ok := i["igmps-flood-traffic"]; ok {

			tmp["igmps_flood_traffic"] = flattenSwitchControllerManagedSwitchPortsIgmpsFloodTraffic(i["igmps-flood-traffic"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_state"
		if _, ok := i["stp-state"]; ok {

			tmp["stp_state"] = flattenSwitchControllerManagedSwitchPortsStpState(i["stp-state"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_root_guard"
		if _, ok := i["stp-root-guard"]; ok {

			tmp["stp_root_guard"] = flattenSwitchControllerManagedSwitchPortsStpRootGuard(i["stp-root-guard"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard"
		if _, ok := i["stp-bpdu-guard"]; ok {

			tmp["stp_bpdu_guard"] = flattenSwitchControllerManagedSwitchPortsStpBpduGuard(i["stp-bpdu-guard"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stp_bpdu_guard_timeout"
		if _, ok := i["stp-bpdu-guard-timeout"]; ok {

			tmp["stp_bpdu_guard_timeout"] = flattenSwitchControllerManagedSwitchPortsStpBpduGuardTimeout(i["stp-bpdu-guard-timeout"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "edge_port"
		if _, ok := i["edge-port"]; ok {

			tmp["edge_port"] = flattenSwitchControllerManagedSwitchPortsEdgePort(i["edge-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "discard_mode"
		if _, ok := i["discard-mode"]; ok {

			tmp["discard_mode"] = flattenSwitchControllerManagedSwitchPortsDiscardMode(i["discard-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sampler"
		if _, ok := i["packet-sampler"]; ok {

			tmp["packet_sampler"] = flattenSwitchControllerManagedSwitchPortsPacketSampler(i["packet-sampler"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_sample_rate"
		if _, ok := i["packet-sample-rate"]; ok {

			tmp["packet_sample_rate"] = flattenSwitchControllerManagedSwitchPortsPacketSampleRate(i["packet-sample-rate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sampler"
		if _, ok := i["sflow-sampler"]; ok {

			tmp["sflow_sampler"] = flattenSwitchControllerManagedSwitchPortsSflowSampler(i["sflow-sampler"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_sample_rate"
		if _, ok := i["sflow-sample-rate"]; ok {

			tmp["sflow_sample_rate"] = flattenSwitchControllerManagedSwitchPortsSflowSampleRate(i["sflow-sample-rate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sflow_counter_interval"
		if _, ok := i["sflow-counter-interval"]; ok {

			tmp["sflow_counter_interval"] = flattenSwitchControllerManagedSwitchPortsSflowCounterInterval(i["sflow-counter-interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sample_direction"
		if _, ok := i["sample-direction"]; ok {

			tmp["sample_direction"] = flattenSwitchControllerManagedSwitchPortsSampleDirection(i["sample-direction"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_capable"
		if _, ok := i["fec-capable"]; ok {

			tmp["fec_capable"] = flattenSwitchControllerManagedSwitchPortsFecCapable(i["fec-capable"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fec_state"
		if _, ok := i["fec-state"]; ok {

			tmp["fec_state"] = flattenSwitchControllerManagedSwitchPortsFecState(i["fec-state"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flow_control"
		if _, ok := i["flow-control"]; ok {

			tmp["flow_control"] = flattenSwitchControllerManagedSwitchPortsFlowControl(i["flow-control"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter"
		if _, ok := i["pause-meter"]; ok {

			tmp["pause_meter"] = flattenSwitchControllerManagedSwitchPortsPauseMeter(i["pause-meter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pause_meter_resume"
		if _, ok := i["pause-meter-resume"]; ok {

			tmp["pause_meter_resume"] = flattenSwitchControllerManagedSwitchPortsPauseMeterResume(i["pause-meter-resume"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard"
		if _, ok := i["loop-guard"]; ok {

			tmp["loop_guard"] = flattenSwitchControllerManagedSwitchPortsLoopGuard(i["loop-guard"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "loop_guard_timeout"
		if _, ok := i["loop-guard-timeout"]; ok {

			tmp["loop_guard_timeout"] = flattenSwitchControllerManagedSwitchPortsLoopGuardTimeout(i["loop-guard-timeout"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := i["qos-policy"]; ok {

			tmp["qos_policy"] = flattenSwitchControllerManagedSwitchPortsQosPolicy(i["qos-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "storm_control_policy"
		if _, ok := i["storm-control-policy"]; ok {

			tmp["storm_control_policy"] = flattenSwitchControllerManagedSwitchPortsStormControlPolicy(i["storm-control-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_security_policy"
		if _, ok := i["port-security-policy"]; ok {

			tmp["port_security_policy"] = flattenSwitchControllerManagedSwitchPortsPortSecurityPolicy(i["port-security-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool"
		if _, ok := i["export-to-pool"]; ok {

			tmp["export_to_pool"] = flattenSwitchControllerManagedSwitchPortsExportToPool(i["export-to-pool"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_tags"
		if _, ok := i["export-tags"]; ok {

			tmp["export_tags"] = flattenSwitchControllerManagedSwitchPortsExportTags(i["export-tags"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to_pool_flag"
		if _, ok := i["export-to-pool_flag"]; ok {

			tmp["export_to_pool_flag"] = flattenSwitchControllerManagedSwitchPortsExportToPool_Flag(i["export-to-pool_flag"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "learning_limit"
		if _, ok := i["learning-limit"]; ok {

			tmp["learning_limit"] = flattenSwitchControllerManagedSwitchPortsLearningLimit(i["learning-limit"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sticky_mac"
		if _, ok := i["sticky-mac"]; ok {

			tmp["sticky_mac"] = flattenSwitchControllerManagedSwitchPortsStickyMac(i["sticky-mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_status"
		if _, ok := i["lldp-status"]; ok {

			tmp["lldp_status"] = flattenSwitchControllerManagedSwitchPortsLldpStatus(i["lldp-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := i["lldp-profile"]; ok {

			tmp["lldp_profile"] = flattenSwitchControllerManagedSwitchPortsLldpProfile(i["lldp-profile"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_to"
		if _, ok := i["export-to"]; ok {

			tmp["export_to"] = flattenSwitchControllerManagedSwitchPortsExportTo(i["export-to"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr"
		if _, ok := i["mac-addr"]; ok {

			tmp["mac_addr"] = flattenSwitchControllerManagedSwitchPortsMacAddr(i["mac-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_selection_criteria"
		if _, ok := i["port-selection-criteria"]; ok {

			tmp["port_selection_criteria"] = flattenSwitchControllerManagedSwitchPortsPortSelectionCriteria(i["port-selection-criteria"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerManagedSwitchPortsDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lacp_speed"
		if _, ok := i["lacp-speed"]; ok {

			tmp["lacp_speed"] = flattenSwitchControllerManagedSwitchPortsLacpSpeed(i["lacp-speed"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {

			tmp["mode"] = flattenSwitchControllerManagedSwitchPortsMode(i["mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bundle"
		if _, ok := i["bundle"]; ok {

			tmp["bundle"] = flattenSwitchControllerManagedSwitchPortsBundle(i["bundle"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_withdrawal_behavior"
		if _, ok := i["member-withdrawal-behavior"]; ok {

			tmp["member_withdrawal_behavior"] = flattenSwitchControllerManagedSwitchPortsMemberWithdrawalBehavior(i["member-withdrawal-behavior"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mclag"
		if _, ok := i["mclag"]; ok {

			tmp["mclag"] = flattenSwitchControllerManagedSwitchPortsMclag(i["mclag"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "min_bundle"
		if _, ok := i["min-bundle"]; ok {

			tmp["min_bundle"] = flattenSwitchControllerManagedSwitchPortsMinBundle(i["min-bundle"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_bundle"
		if _, ok := i["max-bundle"]; ok {

			tmp["max_bundle"] = flattenSwitchControllerManagedSwitchPortsMaxBundle(i["max-bundle"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {

			tmp["members"] = flattenSwitchControllerManagedSwitchPortsMembers(i["members"], d, pre_append, sv)
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

func flattenSwitchControllerManagedSwitchPortsPtpPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsAggregatorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

func flattenSwitchControllerManagedSwitchPortsFiberPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsMediaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
		if _, ok := i["vlan-name"]; ok {

			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchPortsAllowedVlansVlanName(i["vlan-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchPortsAllowedVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsUntaggedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {

			tmp["vlan_name"] = flattenSwitchControllerManagedSwitchPortsUntaggedVlansVlanName(i["vlan-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

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

func flattenSwitchControllerManagedSwitchPortsDhcpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchPortsArpInspectionTrust(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

func flattenSwitchControllerManagedSwitchPortsExportTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_name"
		if _, ok := i["tag-name"]; ok {

			tmp["tag_name"] = flattenSwitchControllerManagedSwitchPortsExportTagsTagName(i["tag-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

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
		if _, ok := i["member-name"]; ok {

			tmp["member_name"] = flattenSwitchControllerManagedSwitchPortsMembersMemberName(i["member-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchPortsMembersMemberName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchIpSourceGuard(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenSwitchControllerManagedSwitchIpSourceGuardPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerManagedSwitchIpSourceGuardDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "binding_entry"
		if _, ok := i["binding-entry"]; ok {

			tmp["binding_entry"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntry(i["binding-entry"], d, pre_append, sv)
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
		if _, ok := i["entry-name"]; ok {

			tmp["entry_name"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryEntryName(i["entry-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenSwitchControllerManagedSwitchIpSourceGuardBindingEntryMac(i["mac"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

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

			tmp["id"] = flattenSwitchControllerManagedSwitchStpInstanceId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenSwitchControllerManagedSwitchStpInstancePriority(i["priority"], d, pre_append, sv)
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

			tmp["id"] = flattenSwitchControllerManagedSwitchSnmpCommunityId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchControllerManagedSwitchSnmpCommunityName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchControllerManagedSwitchSnmpCommunityStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := i["hosts"]; ok {

			tmp["hosts"] = flattenSwitchControllerManagedSwitchSnmpCommunityHosts(i["hosts"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := i["query-v1-status"]; ok {

			tmp["query_v1_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Status(i["query-v1-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_port"
		if _, ok := i["query-v1-port"]; ok {

			tmp["query_v1_port"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV1Port(i["query-v1-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := i["query-v2c-status"]; ok {

			tmp["query_v2c_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CStatus(i["query-v2c-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_port"
		if _, ok := i["query-v2c-port"]; ok {

			tmp["query_v2c_port"] = flattenSwitchControllerManagedSwitchSnmpCommunityQueryV2CPort(i["query-v2c-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := i["trap-v1-status"]; ok {

			tmp["trap_v1_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Status(i["trap-v1-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_lport"
		if _, ok := i["trap-v1-lport"]; ok {

			tmp["trap_v1_lport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Lport(i["trap-v1-lport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_rport"
		if _, ok := i["trap-v1-rport"]; ok {

			tmp["trap_v1_rport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV1Rport(i["trap-v1-rport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := i["trap-v2c-status"]; ok {

			tmp["trap_v2c_status"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CStatus(i["trap-v2c-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_lport"
		if _, ok := i["trap-v2c-lport"]; ok {

			tmp["trap_v2c_lport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CLport(i["trap-v2c-lport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_rport"
		if _, ok := i["trap-v2c-rport"]; ok {

			tmp["trap_v2c_rport"] = flattenSwitchControllerManagedSwitchSnmpCommunityTrapV2CRport(i["trap-v2c-rport"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "events"
		if _, ok := i["events"]; ok {

			tmp["events"] = flattenSwitchControllerManagedSwitchSnmpCommunityEvents(i["events"], d, pre_append, sv)
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

			tmp["id"] = flattenSwitchControllerManagedSwitchSnmpCommunityHostsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSwitchControllerManagedSwitchSnmpCommunityHostsIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

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

			tmp["name"] = flattenSwitchControllerManagedSwitchSnmpUserName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := i["queries"]; ok {

			tmp["queries"] = flattenSwitchControllerManagedSwitchSnmpUserQueries(i["queries"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_port"
		if _, ok := i["query-port"]; ok {

			tmp["query_port"] = flattenSwitchControllerManagedSwitchSnmpUserQueryPort(i["query-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := i["security-level"]; ok {

			tmp["security_level"] = flattenSwitchControllerManagedSwitchSnmpUserSecurityLevel(i["security-level"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := i["auth-proto"]; ok {

			tmp["auth_proto"] = flattenSwitchControllerManagedSwitchSnmpUserAuthProto(i["auth-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if _, ok := i["auth-pwd"]; ok {

			tmp["auth_pwd"] = flattenSwitchControllerManagedSwitchSnmpUserAuthPwd(i["auth-pwd"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_pwd"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := i["priv-proto"]; ok {

			tmp["priv_proto"] = flattenSwitchControllerManagedSwitchSnmpUserPrivProto(i["priv-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if _, ok := i["priv-pwd"]; ok {

			tmp["priv_pwd"] = flattenSwitchControllerManagedSwitchSnmpUserPrivPwd(i["priv-pwd"], d, pre_append, sv)
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

			tmp["name"] = flattenSwitchControllerManagedSwitchRemoteLogName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchControllerManagedSwitchRemoteLogStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {

			tmp["server"] = flattenSwitchControllerManagedSwitchRemoteLogServer(i["server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenSwitchControllerManagedSwitchRemoteLogPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {

			tmp["severity"] = flattenSwitchControllerManagedSwitchRemoteLogSeverity(i["severity"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "csv"
		if _, ok := i["csv"]; ok {

			tmp["csv"] = flattenSwitchControllerManagedSwitchRemoteLogCsv(i["csv"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "facility"
		if _, ok := i["facility"]; ok {

			tmp["facility"] = flattenSwitchControllerManagedSwitchRemoteLogFacility(i["facility"], d, pre_append, sv)
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

			tmp["name"] = flattenSwitchControllerManagedSwitchMirrorName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchControllerManagedSwitchMirrorStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switching_packet"
		if _, ok := i["switching-packet"]; ok {

			tmp["switching_packet"] = flattenSwitchControllerManagedSwitchMirrorSwitchingPacket(i["switching-packet"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {

			tmp["dst"] = flattenSwitchControllerManagedSwitchMirrorDst(i["dst"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_ingress"
		if _, ok := i["src-ingress"]; ok {

			tmp["src_ingress"] = flattenSwitchControllerManagedSwitchMirrorSrcIngress(i["src-ingress"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_egress"
		if _, ok := i["src-egress"]; ok {

			tmp["src_egress"] = flattenSwitchControllerManagedSwitchMirrorSrcEgress(i["src-egress"], d, pre_append, sv)
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

			tmp["name"] = flattenSwitchControllerManagedSwitchMirrorSrcIngressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchMirrorSrcIngressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSwitchControllerManagedSwitchMirrorSrcEgressName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerManagedSwitchMirrorSrcEgressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerManagedSwitchStaticMac(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSwitchControllerManagedSwitchStaticMacId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSwitchControllerManagedSwitchStaticMacType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {

			tmp["vlan"] = flattenSwitchControllerManagedSwitchStaticMacVlan(i["vlan"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenSwitchControllerManagedSwitchStaticMacMac(i["mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenSwitchControllerManagedSwitchStaticMacInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerManagedSwitchStaticMacDescription(i["description"], d, pre_append, sv)
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
		if _, ok := i["command-entry"]; ok {

			tmp["command_entry"] = flattenSwitchControllerManagedSwitchCustomCommandCommandEntry(i["command-entry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {

			tmp["command_name"] = flattenSwitchControllerManagedSwitchCustomCommandCommandName(i["command-name"], d, pre_append, sv)
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

func refreshObjectSwitchControllerManagedSwitch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("switch_id", flattenSwitchControllerManagedSwitchSwitchId(o["switch-id"], d, "switch_id", sv)); err != nil {
		if !fortiAPIPatch(o["switch-id"]) {
			return fmt.Errorf("Error reading switch_id: %v", err)
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

	if err = d.Set("tdr_supported", flattenSwitchControllerManagedSwitchTdrSupported(o["tdr-supported"], d, "tdr_supported", sv)); err != nil {
		if !fortiAPIPatch(o["tdr-supported"]) {
			return fmt.Errorf("Error reading tdr_supported: %v", err)
		}
	}

	{
		v := flattenSwitchControllerManagedSwitchDynamicCapability(o["dynamic-capability"], d, "dynamic_capability", sv)
		if i2ss2arrFortiAPIUpgrade(sv, "6.4.2") == true {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

	if isImportTable() {
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

func expandSwitchControllerManagedSwitchPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ptp_policy"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ptp-policy"], _ = expandSwitchControllerManagedSwitchPortsPtpPolicy(d, i["ptp_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "aggregator_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["aggregator-mode"], _ = expandSwitchControllerManagedSwitchPortsAggregatorMode(d, i["aggregator_mode"], pre_append, sv)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fiber_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["fiber-port"], _ = expandSwitchControllerManagedSwitchPortsFiberPort(d, i["fiber_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "media_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["media-type"], _ = expandSwitchControllerManagedSwitchPortsMediaType(d, i["media_type"], pre_append, sv)
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
		if _, ok := d.GetOk(pre_append); ok {

			tmp["allowed-vlans"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlans(d, i["allowed_vlans"], pre_append, sv)
		} else {
			tmp["allowed-vlans"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "untagged_vlans"
		if _, ok := d.GetOk(pre_append); ok {

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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snooping"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dhcp-snooping"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnooping(d, i["dhcp_snooping"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dhcp_snoop_option82_trust"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dhcp-snoop-option82-trust"], _ = expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d, i["dhcp_snoop_option82_trust"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "arp_inspection_trust"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["arp-inspection-trust"], _ = expandSwitchControllerManagedSwitchPortsArpInspectionTrust(d, i["arp_inspection_trust"], pre_append, sv)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "export_tags"
		if _, ok := d.GetOk(pre_append); ok {

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
		if _, ok := d.GetOk(pre_append); ok {

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

func expandSwitchControllerManagedSwitchPortsPtpPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsAggregatorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandSwitchControllerManagedSwitchPortsFiberPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsMediaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchPortsAllowedVlansVlanName(d, i["vlan_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchPortsAllowedVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan-name"], _ = expandSwitchControllerManagedSwitchPortsUntaggedVlansVlanName(d, i["vlan_name"], pre_append, sv)
		}

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

func expandSwitchControllerManagedSwitchPortsDhcpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsDhcpSnoopOption82Trust(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchPortsArpInspectionTrust(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandSwitchControllerManagedSwitchPortsExportTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tag-name"], _ = expandSwitchControllerManagedSwitchPortsExportTagsTagName(d, i["tag_name"], pre_append, sv)
		}

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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["member-name"], _ = expandSwitchControllerManagedSwitchPortsMembersMemberName(d, i["member_name"], pre_append, sv)
		}

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
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

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
		if _, ok := d.GetOk(pre_append); ok {

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
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

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
		if _, ok := d.GetOk(pre_append); ok {

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
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src-ingress"], _ = expandSwitchControllerManagedSwitchMirrorSrcIngress(d, i["src_ingress"], pre_append, sv)
		} else {
			tmp["src-ingress"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_egress"
		if _, ok := d.GetOk(pre_append); ok {

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

			tmp["name"], _ = expandSwitchControllerManagedSwitchMirrorSrcIngressName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcIngressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerManagedSwitchMirrorSrcEgress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSwitchControllerManagedSwitchMirrorSrcEgressName(d, i["name"], pre_append, sv)
		}

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
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

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
			if i2ss2arrFortiAPIUpgrade(sv, "6.4.2") == true {
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

	if v, ok := d.GetOk("ports"); ok {

		t, err := expandSwitchControllerManagedSwitchPorts(d, v, "ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("ip_source_guard"); ok {

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

	if v, ok := d.GetOk("stp_instance"); ok {

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

	if v, ok := d.GetOk("snmp_community"); ok {

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

	if v, ok := d.GetOk("snmp_user"); ok {

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

	if v, ok := d.GetOk("remote_log"); ok {

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

	if v, ok := d.GetOk("mirror"); ok {

		t, err := expandSwitchControllerManagedSwitchMirror(d, v, "mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mirror"] = t
		}
	}

	if v, ok := d.GetOk("static_mac"); ok {

		t, err := expandSwitchControllerManagedSwitchStaticMac(d, v, "static_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["static-mac"] = t
		}
	}

	if v, ok := d.GetOk("custom_command"); ok {

		t, err := expandSwitchControllerManagedSwitchCustomCommand(d, v, "custom_command", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-command"] = t
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
