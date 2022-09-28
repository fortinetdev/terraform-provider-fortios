// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure interfaces.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemInterfaceRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cli_conn_status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fortilink": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_options": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"code": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dhcp_relay_interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_link_selection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_request_all_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_agent_option": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_classless_route_addition": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gwdetect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ping_serv_status": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"detectserver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"detectprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fail_detect_option": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fail_alert_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fail_action_on_extender": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fail_alert_interfaces": &schema.Schema{
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
			"dhcp_client_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_renew_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipunnumbered": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pppoe_unnumbered_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"detected_peer_mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"padt_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ac_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"defaultgw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_server_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pptp_client": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pptp_user": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pptp_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"pptp_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pptp_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pptp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"arpforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ndiscforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"broadcast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"l2forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"icmp_send_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"icmp_accept_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reachable_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlanforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stpforward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stpforward_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ips_sniffer_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ident_accept": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipmac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"substitute_dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netbios_forward": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wins_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dedicated_to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip_1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip_2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip_3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip6_1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip6_2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trust_ip6_3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ring_rx": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ring_tx": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netflow_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sflow_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"drop_overlapped_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"drop_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sample_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"polling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sample_direction": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"explicit_web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"explicit_ftp_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy_captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_mss": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mediatype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"outbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"egress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_shaping_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"disconnect_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlan_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trunk": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_domain": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lacp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lacp_ha_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lacp_ha_slave": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lacp_speed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_links": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_links_down": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"link_up_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"aggregate_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aggregate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundant_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_device": &schema.Schema{
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
			"devindex": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vindex": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"switch": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"security_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_external_web": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_external_logout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_redirect_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_portal_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_exempt_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_groups": &schema.Schema{
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
			"ike_saml_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stp_ha_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_identification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_user_identification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_identification_active_scan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_netscan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lldp_network_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortiheartbeat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"broadcast_forticlient_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"endpoint_compliance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"estimated_upstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"estimated_downstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"measured_upstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"measured_downstream_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bandwidth_measure_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"monitor_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrrp_virtual_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrrp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrgrp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vrip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"adv_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"preempt": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"accept_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrdst": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrdst_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ignore_default_route": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"proxy_arp": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"snmp_index": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondaryip": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"allowaccess": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"gwdetect": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ping_serv_status": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"detectserver": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"detectprotocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ha_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"preserve_session_route": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_auth_extension_device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ap_discover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortilink_stacking": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortilink_neighbor_detect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_managed_by_fortiipam": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_subnetwork_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fortilink_split_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internal": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fortilink_backup_link": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"switch_controller_access_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_traffic_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_rspan_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_netflow_collect": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_mgmt_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"switch_controller_igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_igmp_snooping_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_igmp_snooping_fast_leave": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_dhcp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_dhcp_snooping_verify_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_dhcp_snooping_option82": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_snooping_server_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"switch_controller_arp_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_learning_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"switch_controller_nac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_dynamic": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_feature": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"switch_controller_iot_scanning": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"swc_vlan": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"swc_first_create": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": &schema.Schema{
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
					},
				},
			},
			"eap_supplicant": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"eap_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"eap_identity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"eap_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"eap_ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"eap_user_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forward_error_correction": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nd_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nd_cert": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nd_security_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nd_timestamp_delta": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nd_timestamp_fuzz": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nd_cga_modifier": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_dns_server_override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_extra_addr": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_send_adv": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"icmp6_send_redirect": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_manage_flag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_other_flag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_max_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_min_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_link_mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ra_send_mtu": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_reachable_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_retrans_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_default_life": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_hop_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"autoconf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unique_autoconf_addr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_identifier": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_prefix_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_upstream_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_delegated_prefix_iaid": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip6_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip6_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"valid_life_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"preferred_life_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rdnss": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dnssl": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"ip6_delegated_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"upstream_interface": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"delegated_prefix_iaid": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"subnet": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rdnss_service": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rdnss": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dhcp6_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_relay_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_client_options": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_prefix_delegation": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_information_request": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_iapd_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iaid": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"prefix_hint": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"prefix_hint_plt": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"prefix_hint_vlt": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"dhcp6_prefix_hint": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dhcp6_prefix_hint_plt": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dhcp6_prefix_hint_vlt": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cli_conn6_status": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"vrrp_virtual_mac6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrip6_link_local": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vrrp6": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrid": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vrgrp": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vrip6": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"priority": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"adv_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"start_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"preempt": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"accept_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"vrdst6": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
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

func dataSourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemInterface: type error")
	}

	o, err := c.ReadSystemInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterface: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemInterface from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceCliConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceClientOptions(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemInterfaceClientOptionsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {
			tmp["code"] = dataSourceFlattenSystemInterfaceClientOptionsCode(i["code"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenSystemInterfaceClientOptionsType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenSystemInterfaceClientOptionsValue(i["value"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemInterfaceClientOptionsIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceClientOptionsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceClientOptionsCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceClientOptionsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceClientOptionsValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceClientOptionsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	var temp_v = v.(string)
	temp_v = strings.TrimRight(temp_v, " ")
	temp_v = strings.ReplaceAll(temp_v, "\"", "")
	var rst_v interface{} = temp_v
	return rst_v
}

func dataSourceFlattenSystemInterfaceDhcpRelayLinkSelection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayRequestAllServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRelayAgentOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpClasslessRouteAddition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFailDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFailDetectOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFailAlertMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFailActionOnExtender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemInterfaceFailAlertInterfacesName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceFailAlertInterfacesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpClientIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpRenewTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpunnumbered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePppoeUnnumberedNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDetectedPeerMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePadtRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDefaultgw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDnsServerProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePptpClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePptpUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePptpPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePptpServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePptpAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePptpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceArpforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceNdiscforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBroadcastForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceL2Forward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIcmpSendRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIcmpAcceptRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVlanforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceStpforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceStpforwardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpsSnifferMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIdentAccept(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSubst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSubstituteDstMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceNetbiosForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceWinsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDedicatedTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTrustIp1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceTrustIp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceTrustIp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemInterfaceTrustIp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTrustIp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTrustIp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMtuOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRingRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRingTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceNetflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDropOverlappedFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDropFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSrcCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePollingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSampleDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceExplicitWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceExplicitFtpProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceProxyCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMediatype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceInbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceOutbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEgressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIngressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVlanProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTrunk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceForwardDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			tmp["interface_name"] = dataSourceFlattenSystemInterfaceMemberInterfaceName(i["interface-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceMemberInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLacpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLacpHaSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLacpHaSlave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSystemIdType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSystemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLacpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMinLinks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMinLinksDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLinkUpDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAggregateType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePriorityOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRedundantInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceManagedDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemInterfaceManagedDeviceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceManagedDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDevindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAuthCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecurityGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemInterfaceSecurityGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceSecurityGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIkeSamlServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceStp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceStpHaSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDeviceUserIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDeviceIdentificationActiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLldpReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLldpTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceLldpNetworkPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFortiheartbeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBroadcastForticlientDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEndpointCompliance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEstimatedUpstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEstimatedDownstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMeasuredUpstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMeasuredDownstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceBandwidthMeasureTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceMonitorBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVirtualMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			tmp["vrid"] = dataSourceFlattenSystemInterfaceVrrpVrid(i["vrid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			tmp["version"] = dataSourceFlattenSystemInterfaceVrrpVersion(i["version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			tmp["vrgrp"] = dataSourceFlattenSystemInterfaceVrrpVrgrp(i["vrgrp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := i["vrip"]; ok {
			tmp["vrip"] = dataSourceFlattenSystemInterfaceVrrpVrip(i["vrip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenSystemInterfaceVrrpPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			tmp["adv_interval"] = dataSourceFlattenSystemInterfaceVrrpAdvInterval(i["adv-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			tmp["start_time"] = dataSourceFlattenSystemInterfaceVrrpStartTime(i["start-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			tmp["preempt"] = dataSourceFlattenSystemInterfaceVrrpPreempt(i["preempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			tmp["accept_mode"] = dataSourceFlattenSystemInterfaceVrrpAcceptMode(i["accept-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := i["vrdst"]; ok {
			tmp["vrdst"] = dataSourceFlattenSystemInterfaceVrrpVrdst(i["vrdst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := i["vrdst-priority"]; ok {
			tmp["vrdst_priority"] = dataSourceFlattenSystemInterfaceVrrpVrdstPriority(i["vrdst-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ignore_default_route"
		if _, ok := i["ignore-default-route"]; ok {
			tmp["ignore_default_route"] = dataSourceFlattenSystemInterfaceVrrpIgnoreDefaultRoute(i["ignore-default-route"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenSystemInterfaceVrrpStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := i["proxy-arp"]; ok {
			tmp["proxy_arp"] = dataSourceFlattenSystemInterfaceVrrpProxyArp(i["proxy-arp"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceVrrpVrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpAdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpPreempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpAcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrdst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpVrdstPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpIgnoreDefaultRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpProxyArp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemInterfaceVrrpProxyArpId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemInterfaceVrrpProxyArpIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceVrrpProxyArpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceVrrpProxyArpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSnmpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryip(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemInterfaceSecondaryipId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemInterfaceSecondaryipIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := i["allowaccess"]; ok {
			tmp["allowaccess"] = dataSourceFlattenSystemInterfaceSecondaryipAllowaccess(i["allowaccess"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			tmp["gwdetect"] = dataSourceFlattenSystemInterfaceSecondaryipGwdetect(i["gwdetect"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			tmp["ping_serv_status"] = dataSourceFlattenSystemInterfaceSecondaryipPingServStatus(i["ping-serv-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			tmp["detectserver"] = dataSourceFlattenSystemInterfaceSecondaryipDetectserver(i["detectserver"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			tmp["detectprotocol"] = dataSourceFlattenSystemInterfaceSecondaryipDetectprotocol(i["detectprotocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			tmp["ha_priority"] = dataSourceFlattenSystemInterfaceSecondaryipHaPriority(i["ha-priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceSecondaryipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipPingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSecondaryipHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfacePreserveSessionRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceAutoAuthExtensionDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceApDiscover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFortilinkStacking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFortilinkNeighborDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpManagedByFortiipam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceManagedSubnetworkSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFortilinkSplitInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceInternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceFortilinkBackupLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerAccessVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerTrafficPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerRspanMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerNetflowCollect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerMgmtVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerIgmpSnoopingProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerIgmpSnoopingFastLeave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerDhcpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerDhcpSnoopingVerifyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerDhcpSnoopingOption82(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpSnoopingServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemInterfaceDhcpSnoopingServerListName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_ip"
		if _, ok := i["server-ip"]; ok {
			tmp["server_ip"] = dataSourceFlattenSystemInterfaceDhcpSnoopingServerListServerIp(i["server-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceDhcpSnoopingServerListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceDhcpSnoopingServerListServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerArpInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerLearningLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerNac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerDynamic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerFeature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwitchControllerIotScanning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwcVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceSwcFirstCreate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemInterfaceTaggingName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = dataSourceFlattenSystemInterfaceTaggingCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			tmp["tags"] = dataSourceFlattenSystemInterfaceTaggingTags(i["tags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceTaggingTags(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemInterfaceTaggingTagsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceTaggingTagsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEapSupplicant(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEapMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEapIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEapCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceEapUserCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceForwardErrorCorrection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = dataSourceFlattenSystemInterfaceIpv6Ip6Mode(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_mode"
	if _, ok := i["nd-mode"]; ok {
		result["nd_mode"] = dataSourceFlattenSystemInterfaceIpv6NdMode(i["nd-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cert"
	if _, ok := i["nd-cert"]; ok {
		result["nd_cert"] = dataSourceFlattenSystemInterfaceIpv6NdCert(i["nd-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := i["nd-security-level"]; ok {
		result["nd_security_level"] = dataSourceFlattenSystemInterfaceIpv6NdSecurityLevel(i["nd-security-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := i["nd-timestamp-delta"]; ok {
		result["nd_timestamp_delta"] = dataSourceFlattenSystemInterfaceIpv6NdTimestampDelta(i["nd-timestamp-delta"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := i["nd-timestamp-fuzz"]; ok {
		result["nd_timestamp_fuzz"] = dataSourceFlattenSystemInterfaceIpv6NdTimestampFuzz(i["nd-timestamp-fuzz"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := i["nd-cga-modifier"]; ok {
		result["nd_cga_modifier"] = dataSourceFlattenSystemInterfaceIpv6NdCgaModifier(i["nd-cga-modifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = dataSourceFlattenSystemInterfaceIpv6Ip6DnsServerOverride(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = dataSourceFlattenSystemInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddr(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = dataSourceFlattenSystemInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = dataSourceFlattenSystemInterfaceIpv6Ip6SendAdv(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "icmp6_send_redirect"
	if _, ok := i["icmp6-send-redirect"]; ok {
		result["icmp6_send_redirect"] = dataSourceFlattenSystemInterfaceIpv6Icmp6SendRedirect(i["icmp6-send-redirect"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6ManageFlag(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6OtherFlag(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = dataSourceFlattenSystemInterfaceIpv6Ip6MaxInterval(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = dataSourceFlattenSystemInterfaceIpv6Ip6MinInterval(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = dataSourceFlattenSystemInterfaceIpv6Ip6LinkMtu(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ra_send_mtu"
	if _, ok := i["ra-send-mtu"]; ok {
		result["ra_send_mtu"] = dataSourceFlattenSystemInterfaceIpv6RaSendMtu(i["ra-send-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6ReachableTime(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6RetransTime(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = dataSourceFlattenSystemInterfaceIpv6Ip6DefaultLife(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = dataSourceFlattenSystemInterfaceIpv6Ip6HopLimit(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = dataSourceFlattenSystemInterfaceIpv6Autoconf(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "unique_autoconf_addr"
	if _, ok := i["unique-autoconf-addr"]; ok {
		result["unique_autoconf_addr"] = dataSourceFlattenSystemInterfaceIpv6UniqueAutoconfAddr(i["unique-autoconf-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "interface_identifier"
	if _, ok := i["interface-identifier"]; ok {
		result["interface_identifier"] = dataSourceFlattenSystemInterfaceIpv6InterfaceIdentifier(i["interface-identifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_mode"
	if _, ok := i["ip6-prefix-mode"]; ok {
		result["ip6_prefix_mode"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixMode(i["ip6-prefix-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := i["ip6-upstream-interface"]; ok {
		result["ip6_upstream_interface"] = dataSourceFlattenSystemInterfaceIpv6Ip6UpstreamInterface(i["ip6-upstream-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_iaid"
	if _, ok := i["ip6-delegated-prefix-iaid"]; ok {
		result["ip6_delegated_prefix_iaid"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixIaid(i["ip6-delegated-prefix-iaid"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := i["ip6-subnet"]; ok {
		result["ip6_subnet"] = dataSourceFlattenSystemInterfaceIpv6Ip6Subnet(i["ip6-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixList(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := i["ip6-delegated-prefix-list"]; ok {
		result["ip6_delegated_prefix_list"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixList(i["ip6-delegated-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := i["dhcp6-relay-service"]; ok {
		result["dhcp6_relay_service"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6RelayService(i["dhcp6-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := i["dhcp6-relay-type"]; ok {
		result["dhcp6_relay_type"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6RelayType(i["dhcp6-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := i["dhcp6-relay-ip"]; ok {
		result["dhcp6_relay_ip"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6RelayIp(i["dhcp6-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := i["dhcp6-client-options"]; ok {
		result["dhcp6_client_options"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6ClientOptions(i["dhcp6-client-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := i["dhcp6-prefix-delegation"]; ok {
		result["dhcp6_prefix_delegation"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixDelegation(i["dhcp6-prefix-delegation"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6InformationRequest(i["dhcp6-information-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_iapd_list"
	if _, ok := i["dhcp6-iapd-list"]; ok {
		result["dhcp6_iapd_list"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdList(i["dhcp6-iapd-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := i["dhcp6-prefix-hint"]; ok {
		result["dhcp6_prefix_hint"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixHint(i["dhcp6-prefix-hint"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := i["dhcp6-prefix-hint-plt"]; ok {
		result["dhcp6_prefix_hint_plt"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixHintPlt(i["dhcp6-prefix-hint-plt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := i["dhcp6-prefix-hint-vlt"]; ok {
		result["dhcp6_prefix_hint_vlt"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixHintVlt(i["dhcp6-prefix-hint-vlt"], d, pre_append)
	}

	pre_append = pre + ".0." + "cli_conn6_status"
	if _, ok := i["cli-conn6-status"]; ok {
		result["cli_conn6_status"] = dataSourceFlattenSystemInterfaceIpv6CliConn6Status(i["cli-conn6-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = dataSourceFlattenSystemInterfaceIpv6VrrpVirtualMac6(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = dataSourceFlattenSystemInterfaceIpv6Vrip6_Link_Local(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6(i["vrrp6"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6NdMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6NdCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6NdSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6NdTimestampDelta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6NdTimestampFuzz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6NdCgaModifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddrPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ExtraAddrPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6SendAdv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Icmp6SendRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ManageFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6OtherFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6MaxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6MinInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6LinkMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6RaSendMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6ReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6RetransTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DefaultLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6HopLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6UniqueAutoconfAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6InterfaceIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6UpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			tmp["autonomous_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			tmp["onlink_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			tmp["valid_life_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime(i["valid-life-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			tmp["preferred_life_time"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(i["preferred-life-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			tmp["rdnss"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListRdnss(i["rdnss"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			tmp["dnssl"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListDnssl(i["dnssl"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListDnssl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			tmp["domain"] = dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListDnsslDomain(i["domain"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Ip6PrefixListDnsslDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			tmp["prefix_id"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId(i["prefix-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			tmp["upstream_interface"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(i["upstream-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delegated_prefix_iaid"
		if _, ok := i["delegated-prefix-iaid"]; ok {
			tmp["delegated_prefix_iaid"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(i["delegated-prefix-iaid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			tmp["autonomous_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			tmp["onlink_flag"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			tmp["subnet"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet(i["subnet"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			tmp["rdnss_service"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService(i["rdnss-service"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			tmp["rdnss"] = dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss(i["rdnss"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListDelegatedPrefixIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6ClientOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixDelegation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6InformationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "iaid"
		if _, ok := i["iaid"]; ok {
			tmp["iaid"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListIaid(i["iaid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint"
		if _, ok := i["prefix-hint"]; ok {
			tmp["prefix_hint"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListPrefixHint(i["prefix-hint"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint_plt"
		if _, ok := i["prefix-hint-plt"]; ok {
			tmp["prefix_hint_plt"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt(i["prefix-hint-plt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_hint_vlt"
		if _, ok := i["prefix-hint-vlt"]; ok {
			tmp["prefix_hint_vlt"] = dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt(i["prefix-hint-vlt"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListIaid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListPrefixHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintPlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6IapdListPrefixHintVlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixHintPlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Dhcp6PrefixHintVlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6CliConn6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6VrrpVirtualMac6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrip6_Link_Local(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			tmp["vrid"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrid(i["vrid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			tmp["vrgrp"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrgrp(i["vrgrp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			tmp["vrip6"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrip6(i["vrip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Priority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			tmp["adv_interval"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6AdvInterval(i["adv-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			tmp["start_time"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6StartTime(i["start-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			tmp["preempt"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Preempt(i["preempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			tmp["accept_mode"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6AcceptMode(i["accept-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			tmp["vrdst6"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrdst6(i["vrdst6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenSystemInterfaceIpv6Vrrp6Status(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6AdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6StartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Preempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6AcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Vrdst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemInterfaceIpv6Vrrp6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemInterfaceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemInterfaceVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vrf", dataSourceFlattenSystemInterfaceVrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("cli_conn_status", dataSourceFlattenSystemInterfaceCliConnStatus(o["cli-conn-status"], d, "cli_conn_status")); err != nil {
		if !fortiAPIPatch(o["cli-conn-status"]) {
			return fmt.Errorf("Error reading cli_conn_status: %v", err)
		}
	}

	if err = d.Set("fortilink", dataSourceFlattenSystemInterfaceFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("switch_controller_source_ip", dataSourceFlattenSystemInterfaceSwitchControllerSourceIp(o["switch-controller-source-ip"], d, "switch_controller_source_ip")); err != nil {
		if !fortiAPIPatch(o["switch-controller-source-ip"]) {
			return fmt.Errorf("Error reading switch_controller_source_ip: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemInterfaceMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("client_options", dataSourceFlattenSystemInterfaceClientOptions(o["client-options"], d, "client_options")); err != nil {
		if !fortiAPIPatch(o["client-options"]) {
			return fmt.Errorf("Error reading client_options: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenSystemInterfaceDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemInterfacePriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_interface_select_method", dataSourceFlattenSystemInterfaceDhcpRelayInterfaceSelectMethod(o["dhcp-relay-interface-select-method"], d, "dhcp_relay_interface_select_method")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-interface-select-method"]) {
			return fmt.Errorf("Error reading dhcp_relay_interface_select_method: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_interface", dataSourceFlattenSystemInterfaceDhcpRelayInterface(o["dhcp-relay-interface"], d, "dhcp_relay_interface")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-interface"]) {
			return fmt.Errorf("Error reading dhcp_relay_interface: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_service", dataSourceFlattenSystemInterfaceDhcpRelayService(o["dhcp-relay-service"], d, "dhcp_relay_service")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-service"]) {
			return fmt.Errorf("Error reading dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_ip", dataSourceFlattenSystemInterfaceDhcpRelayIp(o["dhcp-relay-ip"], d, "dhcp_relay_ip")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-ip"]) {
			return fmt.Errorf("Error reading dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_link_selection", dataSourceFlattenSystemInterfaceDhcpRelayLinkSelection(o["dhcp-relay-link-selection"], d, "dhcp_relay_link_selection")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-link-selection"]) {
			return fmt.Errorf("Error reading dhcp_relay_link_selection: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_request_all_server", dataSourceFlattenSystemInterfaceDhcpRelayRequestAllServer(o["dhcp-relay-request-all-server"], d, "dhcp_relay_request_all_server")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-request-all-server"]) {
			return fmt.Errorf("Error reading dhcp_relay_request_all_server: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_type", dataSourceFlattenSystemInterfaceDhcpRelayType(o["dhcp-relay-type"], d, "dhcp_relay_type")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-type"]) {
			return fmt.Errorf("Error reading dhcp_relay_type: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_agent_option", dataSourceFlattenSystemInterfaceDhcpRelayAgentOption(o["dhcp-relay-agent-option"], d, "dhcp_relay_agent_option")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-agent-option"]) {
			return fmt.Errorf("Error reading dhcp_relay_agent_option: %v", err)
		}
	}

	if err = d.Set("dhcp_classless_route_addition", dataSourceFlattenSystemInterfaceDhcpClasslessRouteAddition(o["dhcp-classless-route-addition"], d, "dhcp_classless_route_addition")); err != nil {
		if !fortiAPIPatch(o["dhcp-classless-route-addition"]) {
			return fmt.Errorf("Error reading dhcp_classless_route_addition: %v", err)
		}
	}

	if err = d.Set("management_ip", dataSourceFlattenSystemInterfaceManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemInterfaceIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", dataSourceFlattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("gwdetect", dataSourceFlattenSystemInterfaceGwdetect(o["gwdetect"], d, "gwdetect")); err != nil {
		if !fortiAPIPatch(o["gwdetect"]) {
			return fmt.Errorf("Error reading gwdetect: %v", err)
		}
	}

	if err = d.Set("ping_serv_status", dataSourceFlattenSystemInterfacePingServStatus(o["ping-serv-status"], d, "ping_serv_status")); err != nil {
		if !fortiAPIPatch(o["ping-serv-status"]) {
			return fmt.Errorf("Error reading ping_serv_status: %v", err)
		}
	}

	if err = d.Set("detectserver", dataSourceFlattenSystemInterfaceDetectserver(o["detectserver"], d, "detectserver")); err != nil {
		if !fortiAPIPatch(o["detectserver"]) {
			return fmt.Errorf("Error reading detectserver: %v", err)
		}
	}

	if err = d.Set("detectprotocol", dataSourceFlattenSystemInterfaceDetectprotocol(o["detectprotocol"], d, "detectprotocol")); err != nil {
		if !fortiAPIPatch(o["detectprotocol"]) {
			return fmt.Errorf("Error reading detectprotocol: %v", err)
		}
	}

	if err = d.Set("ha_priority", dataSourceFlattenSystemInterfaceHaPriority(o["ha-priority"], d, "ha_priority")); err != nil {
		if !fortiAPIPatch(o["ha-priority"]) {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("fail_detect", dataSourceFlattenSystemInterfaceFailDetect(o["fail-detect"], d, "fail_detect")); err != nil {
		if !fortiAPIPatch(o["fail-detect"]) {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if err = d.Set("fail_detect_option", dataSourceFlattenSystemInterfaceFailDetectOption(o["fail-detect-option"], d, "fail_detect_option")); err != nil {
		if !fortiAPIPatch(o["fail-detect-option"]) {
			return fmt.Errorf("Error reading fail_detect_option: %v", err)
		}
	}

	if err = d.Set("fail_alert_method", dataSourceFlattenSystemInterfaceFailAlertMethod(o["fail-alert-method"], d, "fail_alert_method")); err != nil {
		if !fortiAPIPatch(o["fail-alert-method"]) {
			return fmt.Errorf("Error reading fail_alert_method: %v", err)
		}
	}

	if err = d.Set("fail_action_on_extender", dataSourceFlattenSystemInterfaceFailActionOnExtender(o["fail-action-on-extender"], d, "fail_action_on_extender")); err != nil {
		if !fortiAPIPatch(o["fail-action-on-extender"]) {
			return fmt.Errorf("Error reading fail_action_on_extender: %v", err)
		}
	}

	if err = d.Set("fail_alert_interfaces", dataSourceFlattenSystemInterfaceFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
		if !fortiAPIPatch(o["fail-alert-interfaces"]) {
			return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
		}
	}

	if err = d.Set("dhcp_client_identifier", dataSourceFlattenSystemInterfaceDhcpClientIdentifier(o["dhcp-client-identifier"], d, "dhcp_client_identifier")); err != nil {
		if !fortiAPIPatch(o["dhcp-client-identifier"]) {
			return fmt.Errorf("Error reading dhcp_client_identifier: %v", err)
		}
	}

	if err = d.Set("dhcp_renew_time", dataSourceFlattenSystemInterfaceDhcpRenewTime(o["dhcp-renew-time"], d, "dhcp_renew_time")); err != nil {
		if !fortiAPIPatch(o["dhcp-renew-time"]) {
			return fmt.Errorf("Error reading dhcp_renew_time: %v", err)
		}
	}

	if err = d.Set("ipunnumbered", dataSourceFlattenSystemInterfaceIpunnumbered(o["ipunnumbered"], d, "ipunnumbered")); err != nil {
		if !fortiAPIPatch(o["ipunnumbered"]) {
			return fmt.Errorf("Error reading ipunnumbered: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemInterfaceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("pppoe_unnumbered_negotiate", dataSourceFlattenSystemInterfacePppoeUnnumberedNegotiate(o["pppoe-unnumbered-negotiate"], d, "pppoe_unnumbered_negotiate")); err != nil {
		if !fortiAPIPatch(o["pppoe-unnumbered-negotiate"]) {
			return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if err = d.Set("idle_timeout", dataSourceFlattenSystemInterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("detected_peer_mtu", dataSourceFlattenSystemInterfaceDetectedPeerMtu(o["detected-peer-mtu"], d, "detected_peer_mtu")); err != nil {
		if !fortiAPIPatch(o["detected-peer-mtu"]) {
			return fmt.Errorf("Error reading detected_peer_mtu: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", dataSourceFlattenSystemInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["disc-retry-timeout"]) {
			return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
		}
	}

	if err = d.Set("padt_retry_timeout", dataSourceFlattenSystemInterfacePadtRetryTimeout(o["padt-retry-timeout"], d, "padt_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["padt-retry-timeout"]) {
			return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
		}
	}

	if err = d.Set("service_name", dataSourceFlattenSystemInterfaceServiceName(o["service-name"], d, "service_name")); err != nil {
		if !fortiAPIPatch(o["service-name"]) {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("ac_name", dataSourceFlattenSystemInterfaceAcName(o["ac-name"], d, "ac_name")); err != nil {
		if !fortiAPIPatch(o["ac-name"]) {
			return fmt.Errorf("Error reading ac_name: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", dataSourceFlattenSystemInterfaceLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if !fortiAPIPatch(o["lcp-echo-interval"]) {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", dataSourceFlattenSystemInterfaceLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if !fortiAPIPatch(o["lcp-max-echo-fails"]) {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	if err = d.Set("defaultgw", dataSourceFlattenSystemInterfaceDefaultgw(o["defaultgw"], d, "defaultgw")); err != nil {
		if !fortiAPIPatch(o["defaultgw"]) {
			return fmt.Errorf("Error reading defaultgw: %v", err)
		}
	}

	if err = d.Set("dns_server_override", dataSourceFlattenSystemInterfaceDnsServerOverride(o["dns-server-override"], d, "dns_server_override")); err != nil {
		if !fortiAPIPatch(o["dns-server-override"]) {
			return fmt.Errorf("Error reading dns_server_override: %v", err)
		}
	}

	if err = d.Set("dns_server_protocol", dataSourceFlattenSystemInterfaceDnsServerProtocol(o["dns-server-protocol"], d, "dns_server_protocol")); err != nil {
		if !fortiAPIPatch(o["dns-server-protocol"]) {
			return fmt.Errorf("Error reading dns_server_protocol: %v", err)
		}
	}

	if err = d.Set("auth_type", dataSourceFlattenSystemInterfaceAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("pptp_client", dataSourceFlattenSystemInterfacePptpClient(o["pptp-client"], d, "pptp_client")); err != nil {
		if !fortiAPIPatch(o["pptp-client"]) {
			return fmt.Errorf("Error reading pptp_client: %v", err)
		}
	}

	if err = d.Set("pptp_user", dataSourceFlattenSystemInterfacePptpUser(o["pptp-user"], d, "pptp_user")); err != nil {
		if !fortiAPIPatch(o["pptp-user"]) {
			return fmt.Errorf("Error reading pptp_user: %v", err)
		}
	}

	if err = d.Set("pptp_server_ip", dataSourceFlattenSystemInterfacePptpServerIp(o["pptp-server-ip"], d, "pptp_server_ip")); err != nil {
		if !fortiAPIPatch(o["pptp-server-ip"]) {
			return fmt.Errorf("Error reading pptp_server_ip: %v", err)
		}
	}

	if err = d.Set("pptp_auth_type", dataSourceFlattenSystemInterfacePptpAuthType(o["pptp-auth-type"], d, "pptp_auth_type")); err != nil {
		if !fortiAPIPatch(o["pptp-auth-type"]) {
			return fmt.Errorf("Error reading pptp_auth_type: %v", err)
		}
	}

	if err = d.Set("pptp_timeout", dataSourceFlattenSystemInterfacePptpTimeout(o["pptp-timeout"], d, "pptp_timeout")); err != nil {
		if !fortiAPIPatch(o["pptp-timeout"]) {
			return fmt.Errorf("Error reading pptp_timeout: %v", err)
		}
	}

	if err = d.Set("arpforward", dataSourceFlattenSystemInterfaceArpforward(o["arpforward"], d, "arpforward")); err != nil {
		if !fortiAPIPatch(o["arpforward"]) {
			return fmt.Errorf("Error reading arpforward: %v", err)
		}
	}

	if err = d.Set("ndiscforward", dataSourceFlattenSystemInterfaceNdiscforward(o["ndiscforward"], d, "ndiscforward")); err != nil {
		if !fortiAPIPatch(o["ndiscforward"]) {
			return fmt.Errorf("Error reading ndiscforward: %v", err)
		}
	}

	if err = d.Set("broadcast_forward", dataSourceFlattenSystemInterfaceBroadcastForward(o["broadcast-forward"], d, "broadcast_forward")); err != nil {
		if !fortiAPIPatch(o["broadcast-forward"]) {
			return fmt.Errorf("Error reading broadcast_forward: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenSystemInterfaceBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", dataSourceFlattenSystemInterfaceBfdDesiredMinTx(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx")); err != nil {
		if !fortiAPIPatch(o["bfd-desired-min-tx"]) {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", dataSourceFlattenSystemInterfaceBfdDetectMult(o["bfd-detect-mult"], d, "bfd_detect_mult")); err != nil {
		if !fortiAPIPatch(o["bfd-detect-mult"]) {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", dataSourceFlattenSystemInterfaceBfdRequiredMinRx(o["bfd-required-min-rx"], d, "bfd_required_min_rx")); err != nil {
		if !fortiAPIPatch(o["bfd-required-min-rx"]) {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("l2forward", dataSourceFlattenSystemInterfaceL2Forward(o["l2forward"], d, "l2forward")); err != nil {
		if !fortiAPIPatch(o["l2forward"]) {
			return fmt.Errorf("Error reading l2forward: %v", err)
		}
	}

	if err = d.Set("icmp_send_redirect", dataSourceFlattenSystemInterfaceIcmpSendRedirect(o["icmp-send-redirect"], d, "icmp_send_redirect")); err != nil {
		if !fortiAPIPatch(o["icmp-send-redirect"]) {
			return fmt.Errorf("Error reading icmp_send_redirect: %v", err)
		}
	}

	if err = d.Set("icmp_accept_redirect", dataSourceFlattenSystemInterfaceIcmpAcceptRedirect(o["icmp-accept-redirect"], d, "icmp_accept_redirect")); err != nil {
		if !fortiAPIPatch(o["icmp-accept-redirect"]) {
			return fmt.Errorf("Error reading icmp_accept_redirect: %v", err)
		}
	}

	if err = d.Set("reachable_time", dataSourceFlattenSystemInterfaceReachableTime(o["reachable-time"], d, "reachable_time")); err != nil {
		if !fortiAPIPatch(o["reachable-time"]) {
			return fmt.Errorf("Error reading reachable_time: %v", err)
		}
	}

	if err = d.Set("vlanforward", dataSourceFlattenSystemInterfaceVlanforward(o["vlanforward"], d, "vlanforward")); err != nil {
		if !fortiAPIPatch(o["vlanforward"]) {
			return fmt.Errorf("Error reading vlanforward: %v", err)
		}
	}

	if err = d.Set("stpforward", dataSourceFlattenSystemInterfaceStpforward(o["stpforward"], d, "stpforward")); err != nil {
		if !fortiAPIPatch(o["stpforward"]) {
			return fmt.Errorf("Error reading stpforward: %v", err)
		}
	}

	if err = d.Set("stpforward_mode", dataSourceFlattenSystemInterfaceStpforwardMode(o["stpforward-mode"], d, "stpforward_mode")); err != nil {
		if !fortiAPIPatch(o["stpforward-mode"]) {
			return fmt.Errorf("Error reading stpforward_mode: %v", err)
		}
	}

	if err = d.Set("ips_sniffer_mode", dataSourceFlattenSystemInterfaceIpsSnifferMode(o["ips-sniffer-mode"], d, "ips_sniffer_mode")); err != nil {
		if !fortiAPIPatch(o["ips-sniffer-mode"]) {
			return fmt.Errorf("Error reading ips_sniffer_mode: %v", err)
		}
	}

	if err = d.Set("ident_accept", dataSourceFlattenSystemInterfaceIdentAccept(o["ident-accept"], d, "ident_accept")); err != nil {
		if !fortiAPIPatch(o["ident-accept"]) {
			return fmt.Errorf("Error reading ident_accept: %v", err)
		}
	}

	if err = d.Set("ipmac", dataSourceFlattenSystemInterfaceIpmac(o["ipmac"], d, "ipmac")); err != nil {
		if !fortiAPIPatch(o["ipmac"]) {
			return fmt.Errorf("Error reading ipmac: %v", err)
		}
	}

	if err = d.Set("subst", dataSourceFlattenSystemInterfaceSubst(o["subst"], d, "subst")); err != nil {
		if !fortiAPIPatch(o["subst"]) {
			return fmt.Errorf("Error reading subst: %v", err)
		}
	}

	if err = d.Set("macaddr", dataSourceFlattenSystemInterfaceMacaddr(o["macaddr"], d, "macaddr")); err != nil {
		if !fortiAPIPatch(o["macaddr"]) {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("substitute_dst_mac", dataSourceFlattenSystemInterfaceSubstituteDstMac(o["substitute-dst-mac"], d, "substitute_dst_mac")); err != nil {
		if !fortiAPIPatch(o["substitute-dst-mac"]) {
			return fmt.Errorf("Error reading substitute_dst_mac: %v", err)
		}
	}

	if err = d.Set("speed", dataSourceFlattenSystemInterfaceSpeed(o["speed"], d, "speed")); err != nil {
		if !fortiAPIPatch(o["speed"]) {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemInterfaceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("netbios_forward", dataSourceFlattenSystemInterfaceNetbiosForward(o["netbios-forward"], d, "netbios_forward")); err != nil {
		if !fortiAPIPatch(o["netbios-forward"]) {
			return fmt.Errorf("Error reading netbios_forward: %v", err)
		}
	}

	if err = d.Set("wins_ip", dataSourceFlattenSystemInterfaceWinsIp(o["wins-ip"], d, "wins_ip")); err != nil {
		if !fortiAPIPatch(o["wins-ip"]) {
			return fmt.Errorf("Error reading wins_ip: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemInterfaceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("dedicated_to", dataSourceFlattenSystemInterfaceDedicatedTo(o["dedicated-to"], d, "dedicated_to")); err != nil {
		if !fortiAPIPatch(o["dedicated-to"]) {
			return fmt.Errorf("Error reading dedicated_to: %v", err)
		}
	}

	if err = d.Set("trust_ip_1", dataSourceFlattenSystemInterfaceTrustIp1(o["trust-ip-1"], d, "trust_ip_1")); err != nil {
		if !fortiAPIPatch(o["trust-ip-1"]) {
			return fmt.Errorf("Error reading trust_ip_1: %v", err)
		}
	}

	if err = d.Set("trust_ip_2", dataSourceFlattenSystemInterfaceTrustIp2(o["trust-ip-2"], d, "trust_ip_2")); err != nil {
		if !fortiAPIPatch(o["trust-ip-2"]) {
			return fmt.Errorf("Error reading trust_ip_2: %v", err)
		}
	}

	if err = d.Set("trust_ip_3", dataSourceFlattenSystemInterfaceTrustIp3(o["trust-ip-3"], d, "trust_ip_3")); err != nil {
		if !fortiAPIPatch(o["trust-ip-3"]) {
			return fmt.Errorf("Error reading trust_ip_3: %v", err)
		}
	}

	if err = d.Set("trust_ip6_1", dataSourceFlattenSystemInterfaceTrustIp61(o["trust-ip6-1"], d, "trust_ip6_1")); err != nil {
		if !fortiAPIPatch(o["trust-ip6-1"]) {
			return fmt.Errorf("Error reading trust_ip6_1: %v", err)
		}
	}

	if err = d.Set("trust_ip6_2", dataSourceFlattenSystemInterfaceTrustIp62(o["trust-ip6-2"], d, "trust_ip6_2")); err != nil {
		if !fortiAPIPatch(o["trust-ip6-2"]) {
			return fmt.Errorf("Error reading trust_ip6_2: %v", err)
		}
	}

	if err = d.Set("trust_ip6_3", dataSourceFlattenSystemInterfaceTrustIp63(o["trust-ip6-3"], d, "trust_ip6_3")); err != nil {
		if !fortiAPIPatch(o["trust-ip6-3"]) {
			return fmt.Errorf("Error reading trust_ip6_3: %v", err)
		}
	}

	if err = d.Set("mtu_override", dataSourceFlattenSystemInterfaceMtuOverride(o["mtu-override"], d, "mtu_override")); err != nil {
		if !fortiAPIPatch(o["mtu-override"]) {
			return fmt.Errorf("Error reading mtu_override: %v", err)
		}
	}

	if err = d.Set("mtu", dataSourceFlattenSystemInterfaceMtu(o["mtu"], d, "mtu")); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("ring_rx", dataSourceFlattenSystemInterfaceRingRx(o["ring-rx"], d, "ring_rx")); err != nil {
		if !fortiAPIPatch(o["ring-rx"]) {
			return fmt.Errorf("Error reading ring_rx: %v", err)
		}
	}

	if err = d.Set("ring_tx", dataSourceFlattenSystemInterfaceRingTx(o["ring-tx"], d, "ring_tx")); err != nil {
		if !fortiAPIPatch(o["ring-tx"]) {
			return fmt.Errorf("Error reading ring_tx: %v", err)
		}
	}

	if err = d.Set("wccp", dataSourceFlattenSystemInterfaceWccp(o["wccp"], d, "wccp")); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("netflow_sampler", dataSourceFlattenSystemInterfaceNetflowSampler(o["netflow-sampler"], d, "netflow_sampler")); err != nil {
		if !fortiAPIPatch(o["netflow-sampler"]) {
			return fmt.Errorf("Error reading netflow_sampler: %v", err)
		}
	}

	if err = d.Set("sflow_sampler", dataSourceFlattenSystemInterfaceSflowSampler(o["sflow-sampler"], d, "sflow_sampler")); err != nil {
		if !fortiAPIPatch(o["sflow-sampler"]) {
			return fmt.Errorf("Error reading sflow_sampler: %v", err)
		}
	}

	if err = d.Set("drop_overlapped_fragment", dataSourceFlattenSystemInterfaceDropOverlappedFragment(o["drop-overlapped-fragment"], d, "drop_overlapped_fragment")); err != nil {
		if !fortiAPIPatch(o["drop-overlapped-fragment"]) {
			return fmt.Errorf("Error reading drop_overlapped_fragment: %v", err)
		}
	}

	if err = d.Set("drop_fragment", dataSourceFlattenSystemInterfaceDropFragment(o["drop-fragment"], d, "drop_fragment")); err != nil {
		if !fortiAPIPatch(o["drop-fragment"]) {
			return fmt.Errorf("Error reading drop_fragment: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", dataSourceFlattenSystemInterfaceScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("src_check", dataSourceFlattenSystemInterfaceSrcCheck(o["src-check"], d, "src_check")); err != nil {
		if !fortiAPIPatch(o["src-check"]) {
			return fmt.Errorf("Error reading src_check: %v", err)
		}
	}

	if err = d.Set("sample_rate", dataSourceFlattenSystemInterfaceSampleRate(o["sample-rate"], d, "sample_rate")); err != nil {
		if !fortiAPIPatch(o["sample-rate"]) {
			return fmt.Errorf("Error reading sample_rate: %v", err)
		}
	}

	if err = d.Set("polling_interval", dataSourceFlattenSystemInterfacePollingInterval(o["polling-interval"], d, "polling_interval")); err != nil {
		if !fortiAPIPatch(o["polling-interval"]) {
			return fmt.Errorf("Error reading polling_interval: %v", err)
		}
	}

	if err = d.Set("sample_direction", dataSourceFlattenSystemInterfaceSampleDirection(o["sample-direction"], d, "sample_direction")); err != nil {
		if !fortiAPIPatch(o["sample-direction"]) {
			return fmt.Errorf("Error reading sample_direction: %v", err)
		}
	}

	if err = d.Set("explicit_web_proxy", dataSourceFlattenSystemInterfaceExplicitWebProxy(o["explicit-web-proxy"], d, "explicit_web_proxy")); err != nil {
		if !fortiAPIPatch(o["explicit-web-proxy"]) {
			return fmt.Errorf("Error reading explicit_web_proxy: %v", err)
		}
	}

	if err = d.Set("explicit_ftp_proxy", dataSourceFlattenSystemInterfaceExplicitFtpProxy(o["explicit-ftp-proxy"], d, "explicit_ftp_proxy")); err != nil {
		if !fortiAPIPatch(o["explicit-ftp-proxy"]) {
			return fmt.Errorf("Error reading explicit_ftp_proxy: %v", err)
		}
	}

	if err = d.Set("proxy_captive_portal", dataSourceFlattenSystemInterfaceProxyCaptivePortal(o["proxy-captive-portal"], d, "proxy_captive_portal")); err != nil {
		if !fortiAPIPatch(o["proxy-captive-portal"]) {
			return fmt.Errorf("Error reading proxy_captive_portal: %v", err)
		}
	}

	if err = d.Set("tcp_mss", dataSourceFlattenSystemInterfaceTcpMss(o["tcp-mss"], d, "tcp_mss")); err != nil {
		if !fortiAPIPatch(o["tcp-mss"]) {
			return fmt.Errorf("Error reading tcp_mss: %v", err)
		}
	}

	if err = d.Set("mediatype", dataSourceFlattenSystemInterfaceMediatype(o["mediatype"], d, "mediatype")); err != nil {
		if !fortiAPIPatch(o["mediatype"]) {
			return fmt.Errorf("Error reading mediatype: %v", err)
		}
	}

	if err = d.Set("inbandwidth", dataSourceFlattenSystemInterfaceInbandwidth(o["inbandwidth"], d, "inbandwidth")); err != nil {
		if !fortiAPIPatch(o["inbandwidth"]) {
			return fmt.Errorf("Error reading inbandwidth: %v", err)
		}
	}

	if err = d.Set("outbandwidth", dataSourceFlattenSystemInterfaceOutbandwidth(o["outbandwidth"], d, "outbandwidth")); err != nil {
		if !fortiAPIPatch(o["outbandwidth"]) {
			return fmt.Errorf("Error reading outbandwidth: %v", err)
		}
	}

	if err = d.Set("egress_shaping_profile", dataSourceFlattenSystemInterfaceEgressShapingProfile(o["egress-shaping-profile"], d, "egress_shaping_profile")); err != nil {
		if !fortiAPIPatch(o["egress-shaping-profile"]) {
			return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("ingress_shaping_profile", dataSourceFlattenSystemInterfaceIngressShapingProfile(o["ingress-shaping-profile"], d, "ingress_shaping_profile")); err != nil {
		if !fortiAPIPatch(o["ingress-shaping-profile"]) {
			return fmt.Errorf("Error reading ingress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("disconnect_threshold", dataSourceFlattenSystemInterfaceDisconnectThreshold(o["disconnect-threshold"], d, "disconnect_threshold")); err != nil {
		if !fortiAPIPatch(o["disconnect-threshold"]) {
			return fmt.Errorf("Error reading disconnect_threshold: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", dataSourceFlattenSystemInterfaceSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if !fortiAPIPatch(o["spillover-threshold"]) {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", dataSourceFlattenSystemInterfaceIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if !fortiAPIPatch(o["ingress-spillover-threshold"]) {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenSystemInterfaceWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemInterfaceInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("external", dataSourceFlattenSystemInterfaceExternal(o["external"], d, "external")); err != nil {
		if !fortiAPIPatch(o["external"]) {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("vlan_protocol", dataSourceFlattenSystemInterfaceVlanProtocol(o["vlan-protocol"], d, "vlan_protocol")); err != nil {
		if !fortiAPIPatch(o["vlan-protocol"]) {
			return fmt.Errorf("Error reading vlan_protocol: %v", err)
		}
	}

	if err = d.Set("vlanid", dataSourceFlattenSystemInterfaceVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("trunk", dataSourceFlattenSystemInterfaceTrunk(o["trunk"], d, "trunk")); err != nil {
		if !fortiAPIPatch(o["trunk"]) {
			return fmt.Errorf("Error reading trunk: %v", err)
		}
	}

	if err = d.Set("forward_domain", dataSourceFlattenSystemInterfaceForwardDomain(o["forward-domain"], d, "forward_domain")); err != nil {
		if !fortiAPIPatch(o["forward-domain"]) {
			return fmt.Errorf("Error reading forward_domain: %v", err)
		}
	}

	if err = d.Set("remote_ip", dataSourceFlattenSystemInterfaceRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("member", dataSourceFlattenSystemInterfaceMember(o["member"], d, "member")); err != nil {
		if !fortiAPIPatch(o["member"]) {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("lacp_mode", dataSourceFlattenSystemInterfaceLacpMode(o["lacp-mode"], d, "lacp_mode")); err != nil {
		if !fortiAPIPatch(o["lacp-mode"]) {
			return fmt.Errorf("Error reading lacp_mode: %v", err)
		}
	}

	if err = d.Set("lacp_ha_secondary", dataSourceFlattenSystemInterfaceLacpHaSecondary(o["lacp-ha-secondary"], d, "lacp_ha_secondary")); err != nil {
		if !fortiAPIPatch(o["lacp-ha-secondary"]) {
			return fmt.Errorf("Error reading lacp_ha_secondary: %v", err)
		}
	}

	if err = d.Set("lacp_ha_slave", dataSourceFlattenSystemInterfaceLacpHaSlave(o["lacp-ha-slave"], d, "lacp_ha_slave")); err != nil {
		if !fortiAPIPatch(o["lacp-ha-slave"]) {
			return fmt.Errorf("Error reading lacp_ha_slave: %v", err)
		}
	}

	if err = d.Set("system_id_type", dataSourceFlattenSystemInterfaceSystemIdType(o["system-id-type"], d, "system_id_type")); err != nil {
		if !fortiAPIPatch(o["system-id-type"]) {
			return fmt.Errorf("Error reading system_id_type: %v", err)
		}
	}

	if err = d.Set("system_id", dataSourceFlattenSystemInterfaceSystemId(o["system-id"], d, "system_id")); err != nil {
		if !fortiAPIPatch(o["system-id"]) {
			return fmt.Errorf("Error reading system_id: %v", err)
		}
	}

	if err = d.Set("lacp_speed", dataSourceFlattenSystemInterfaceLacpSpeed(o["lacp-speed"], d, "lacp_speed")); err != nil {
		if !fortiAPIPatch(o["lacp-speed"]) {
			return fmt.Errorf("Error reading lacp_speed: %v", err)
		}
	}

	if err = d.Set("min_links", dataSourceFlattenSystemInterfaceMinLinks(o["min-links"], d, "min_links")); err != nil {
		if !fortiAPIPatch(o["min-links"]) {
			return fmt.Errorf("Error reading min_links: %v", err)
		}
	}

	if err = d.Set("min_links_down", dataSourceFlattenSystemInterfaceMinLinksDown(o["min-links-down"], d, "min_links_down")); err != nil {
		if !fortiAPIPatch(o["min-links-down"]) {
			return fmt.Errorf("Error reading min_links_down: %v", err)
		}
	}

	if err = d.Set("algorithm", dataSourceFlattenSystemInterfaceAlgorithm(o["algorithm"], d, "algorithm")); err != nil {
		if !fortiAPIPatch(o["algorithm"]) {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("link_up_delay", dataSourceFlattenSystemInterfaceLinkUpDelay(o["link-up-delay"], d, "link_up_delay")); err != nil {
		if !fortiAPIPatch(o["link-up-delay"]) {
			return fmt.Errorf("Error reading link_up_delay: %v", err)
		}
	}

	if err = d.Set("aggregate_type", dataSourceFlattenSystemInterfaceAggregateType(o["aggregate-type"], d, "aggregate_type")); err != nil {
		if !fortiAPIPatch(o["aggregate-type"]) {
			return fmt.Errorf("Error reading aggregate_type: %v", err)
		}
	}

	if err = d.Set("priority_override", dataSourceFlattenSystemInterfacePriorityOverride(o["priority-override"], d, "priority_override")); err != nil {
		if !fortiAPIPatch(o["priority-override"]) {
			return fmt.Errorf("Error reading priority_override: %v", err)
		}
	}

	if err = d.Set("aggregate", dataSourceFlattenSystemInterfaceAggregate(o["aggregate"], d, "aggregate")); err != nil {
		if !fortiAPIPatch(o["aggregate"]) {
			return fmt.Errorf("Error reading aggregate: %v", err)
		}
	}

	if err = d.Set("redundant_interface", dataSourceFlattenSystemInterfaceRedundantInterface(o["redundant-interface"], d, "redundant_interface")); err != nil {
		if !fortiAPIPatch(o["redundant-interface"]) {
			return fmt.Errorf("Error reading redundant_interface: %v", err)
		}
	}

	if err = d.Set("managed_device", dataSourceFlattenSystemInterfaceManagedDevice(o["managed-device"], d, "managed_device")); err != nil {
		if !fortiAPIPatch(o["managed-device"]) {
			return fmt.Errorf("Error reading managed_device: %v", err)
		}
	}

	if err = d.Set("devindex", dataSourceFlattenSystemInterfaceDevindex(o["devindex"], d, "devindex")); err != nil {
		if !fortiAPIPatch(o["devindex"]) {
			return fmt.Errorf("Error reading devindex: %v", err)
		}
	}

	if err = d.Set("vindex", dataSourceFlattenSystemInterfaceVindex(o["vindex"], d, "vindex")); err != nil {
		if !fortiAPIPatch(o["vindex"]) {
			return fmt.Errorf("Error reading vindex: %v", err)
		}
	}

	if err = d.Set("switch", dataSourceFlattenSystemInterfaceSwitch(o["switch"], d, "switch")); err != nil {
		if !fortiAPIPatch(o["switch"]) {
			return fmt.Errorf("Error reading switch: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenSystemInterfaceDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("alias", dataSourceFlattenSystemInterfaceAlias(o["alias"], d, "alias")); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("security_mode", dataSourceFlattenSystemInterfaceSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("captive_portal", dataSourceFlattenSystemInterfaceCaptivePortal(o["captive-portal"], d, "captive_portal")); err != nil {
		if !fortiAPIPatch(o["captive-portal"]) {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("security_mac_auth_bypass", dataSourceFlattenSystemInterfaceSecurityMacAuthBypass(o["security-mac-auth-bypass"], d, "security_mac_auth_bypass")); err != nil {
		if !fortiAPIPatch(o["security-mac-auth-bypass"]) {
			return fmt.Errorf("Error reading security_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("security_external_web", dataSourceFlattenSystemInterfaceSecurityExternalWeb(o["security-external-web"], d, "security_external_web")); err != nil {
		if !fortiAPIPatch(o["security-external-web"]) {
			return fmt.Errorf("Error reading security_external_web: %v", err)
		}
	}

	if err = d.Set("security_external_logout", dataSourceFlattenSystemInterfaceSecurityExternalLogout(o["security-external-logout"], d, "security_external_logout")); err != nil {
		if !fortiAPIPatch(o["security-external-logout"]) {
			return fmt.Errorf("Error reading security_external_logout: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", dataSourceFlattenSystemInterfaceReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", dataSourceFlattenSystemInterfaceSecurityRedirectUrl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if !fortiAPIPatch(o["security-redirect-url"]) {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("auth_cert", dataSourceFlattenSystemInterfaceAuthCert(o["auth-cert"], d, "auth_cert")); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal_addr", dataSourceFlattenSystemInterfaceAuthPortalAddr(o["auth-portal-addr"], d, "auth_portal_addr")); err != nil {
		if !fortiAPIPatch(o["auth-portal-addr"]) {
			return fmt.Errorf("Error reading auth_portal_addr: %v", err)
		}
	}

	if err = d.Set("security_exempt_list", dataSourceFlattenSystemInterfaceSecurityExemptList(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if !fortiAPIPatch(o["security-exempt-list"]) {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_groups", dataSourceFlattenSystemInterfaceSecurityGroups(o["security-groups"], d, "security_groups")); err != nil {
		if !fortiAPIPatch(o["security-groups"]) {
			return fmt.Errorf("Error reading security_groups: %v", err)
		}
	}

	if err = d.Set("ike_saml_server", dataSourceFlattenSystemInterfaceIkeSamlServer(o["ike-saml-server"], d, "ike_saml_server")); err != nil {
		if !fortiAPIPatch(o["ike-saml-server"]) {
			return fmt.Errorf("Error reading ike_saml_server: %v", err)
		}
	}

	if err = d.Set("stp", dataSourceFlattenSystemInterfaceStp(o["stp"], d, "stp")); err != nil {
		if !fortiAPIPatch(o["stp"]) {
			return fmt.Errorf("Error reading stp: %v", err)
		}
	}

	if err = d.Set("stp_ha_secondary", dataSourceFlattenSystemInterfaceStpHaSecondary(o["stp-ha-secondary"], d, "stp_ha_secondary")); err != nil {
		if !fortiAPIPatch(o["stp-ha-secondary"]) {
			return fmt.Errorf("Error reading stp_ha_secondary: %v", err)
		}
	}

	if err = d.Set("device_identification", dataSourceFlattenSystemInterfaceDeviceIdentification(o["device-identification"], d, "device_identification")); err != nil {
		if !fortiAPIPatch(o["device-identification"]) {
			return fmt.Errorf("Error reading device_identification: %v", err)
		}
	}

	if err = d.Set("device_user_identification", dataSourceFlattenSystemInterfaceDeviceUserIdentification(o["device-user-identification"], d, "device_user_identification")); err != nil {
		if !fortiAPIPatch(o["device-user-identification"]) {
			return fmt.Errorf("Error reading device_user_identification: %v", err)
		}
	}

	if err = d.Set("device_identification_active_scan", dataSourceFlattenSystemInterfaceDeviceIdentificationActiveScan(o["device-identification-active-scan"], d, "device_identification_active_scan")); err != nil {
		if !fortiAPIPatch(o["device-identification-active-scan"]) {
			return fmt.Errorf("Error reading device_identification_active_scan: %v", err)
		}
	}

	if err = d.Set("device_access_list", dataSourceFlattenSystemInterfaceDeviceAccessList(o["device-access-list"], d, "device_access_list")); err != nil {
		if !fortiAPIPatch(o["device-access-list"]) {
			return fmt.Errorf("Error reading device_access_list: %v", err)
		}
	}

	if err = d.Set("device_netscan", dataSourceFlattenSystemInterfaceDeviceNetscan(o["device-netscan"], d, "device_netscan")); err != nil {
		if !fortiAPIPatch(o["device-netscan"]) {
			return fmt.Errorf("Error reading device_netscan: %v", err)
		}
	}

	if err = d.Set("lldp_reception", dataSourceFlattenSystemInterfaceLldpReception(o["lldp-reception"], d, "lldp_reception")); err != nil {
		if !fortiAPIPatch(o["lldp-reception"]) {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", dataSourceFlattenSystemInterfaceLldpTransmission(o["lldp-transmission"], d, "lldp_transmission")); err != nil {
		if !fortiAPIPatch(o["lldp-transmission"]) {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("lldp_network_policy", dataSourceFlattenSystemInterfaceLldpNetworkPolicy(o["lldp-network-policy"], d, "lldp_network_policy")); err != nil {
		if !fortiAPIPatch(o["lldp-network-policy"]) {
			return fmt.Errorf("Error reading lldp_network_policy: %v", err)
		}
	}

	if err = d.Set("fortiheartbeat", dataSourceFlattenSystemInterfaceFortiheartbeat(o["fortiheartbeat"], d, "fortiheartbeat")); err != nil {
		if !fortiAPIPatch(o["fortiheartbeat"]) {
			return fmt.Errorf("Error reading fortiheartbeat: %v", err)
		}
	}

	if err = d.Set("broadcast_forticlient_discovery", dataSourceFlattenSystemInterfaceBroadcastForticlientDiscovery(o["broadcast-forticlient-discovery"], d, "broadcast_forticlient_discovery")); err != nil {
		if !fortiAPIPatch(o["broadcast-forticlient-discovery"]) {
			return fmt.Errorf("Error reading broadcast_forticlient_discovery: %v", err)
		}
	}

	if err = d.Set("endpoint_compliance", dataSourceFlattenSystemInterfaceEndpointCompliance(o["endpoint-compliance"], d, "endpoint_compliance")); err != nil {
		if !fortiAPIPatch(o["endpoint-compliance"]) {
			return fmt.Errorf("Error reading endpoint_compliance: %v", err)
		}
	}

	if err = d.Set("estimated_upstream_bandwidth", dataSourceFlattenSystemInterfaceEstimatedUpstreamBandwidth(o["estimated-upstream-bandwidth"], d, "estimated_upstream_bandwidth")); err != nil {
		if !fortiAPIPatch(o["estimated-upstream-bandwidth"]) {
			return fmt.Errorf("Error reading estimated_upstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("estimated_downstream_bandwidth", dataSourceFlattenSystemInterfaceEstimatedDownstreamBandwidth(o["estimated-downstream-bandwidth"], d, "estimated_downstream_bandwidth")); err != nil {
		if !fortiAPIPatch(o["estimated-downstream-bandwidth"]) {
			return fmt.Errorf("Error reading estimated_downstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("measured_upstream_bandwidth", dataSourceFlattenSystemInterfaceMeasuredUpstreamBandwidth(o["measured-upstream-bandwidth"], d, "measured_upstream_bandwidth")); err != nil {
		if !fortiAPIPatch(o["measured-upstream-bandwidth"]) {
			return fmt.Errorf("Error reading measured_upstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("measured_downstream_bandwidth", dataSourceFlattenSystemInterfaceMeasuredDownstreamBandwidth(o["measured-downstream-bandwidth"], d, "measured_downstream_bandwidth")); err != nil {
		if !fortiAPIPatch(o["measured-downstream-bandwidth"]) {
			return fmt.Errorf("Error reading measured_downstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_measure_time", dataSourceFlattenSystemInterfaceBandwidthMeasureTime(o["bandwidth-measure-time"], d, "bandwidth_measure_time")); err != nil {
		if !fortiAPIPatch(o["bandwidth-measure-time"]) {
			return fmt.Errorf("Error reading bandwidth_measure_time: %v", err)
		}
	}

	if err = d.Set("monitor_bandwidth", dataSourceFlattenSystemInterfaceMonitorBandwidth(o["monitor-bandwidth"], d, "monitor_bandwidth")); err != nil {
		if !fortiAPIPatch(o["monitor-bandwidth"]) {
			return fmt.Errorf("Error reading monitor_bandwidth: %v", err)
		}
	}

	if err = d.Set("vrrp_virtual_mac", dataSourceFlattenSystemInterfaceVrrpVirtualMac(o["vrrp-virtual-mac"], d, "vrrp_virtual_mac")); err != nil {
		if !fortiAPIPatch(o["vrrp-virtual-mac"]) {
			return fmt.Errorf("Error reading vrrp_virtual_mac: %v", err)
		}
	}

	if err = d.Set("vrrp", dataSourceFlattenSystemInterfaceVrrp(o["vrrp"], d, "vrrp")); err != nil {
		if !fortiAPIPatch(o["vrrp"]) {
			return fmt.Errorf("Error reading vrrp: %v", err)
		}
	}

	if err = d.Set("role", dataSourceFlattenSystemInterfaceRole(o["role"], d, "role")); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("snmp_index", dataSourceFlattenSystemInterfaceSnmpIndex(o["snmp-index"], d, "snmp_index")); err != nil {
		if !fortiAPIPatch(o["snmp-index"]) {
			return fmt.Errorf("Error reading snmp_index: %v", err)
		}
	}

	if err = d.Set("secondary_ip", dataSourceFlattenSystemInterfaceSecondaryIp(o["secondary-IP"], d, "secondary_ip")); err != nil {
		if !fortiAPIPatch(o["secondary-IP"]) {
			return fmt.Errorf("Error reading secondary_ip: %v", err)
		}
	}

	if err = d.Set("secondaryip", dataSourceFlattenSystemInterfaceSecondaryip(o["secondaryip"], d, "secondaryip")); err != nil {
		if !fortiAPIPatch(o["secondaryip"]) {
			return fmt.Errorf("Error reading secondaryip: %v", err)
		}
	}

	if err = d.Set("preserve_session_route", dataSourceFlattenSystemInterfacePreserveSessionRoute(o["preserve-session-route"], d, "preserve_session_route")); err != nil {
		if !fortiAPIPatch(o["preserve-session-route"]) {
			return fmt.Errorf("Error reading preserve_session_route: %v", err)
		}
	}

	if err = d.Set("auto_auth_extension_device", dataSourceFlattenSystemInterfaceAutoAuthExtensionDevice(o["auto-auth-extension-device"], d, "auto_auth_extension_device")); err != nil {
		if !fortiAPIPatch(o["auto-auth-extension-device"]) {
			return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
		}
	}

	if err = d.Set("ap_discover", dataSourceFlattenSystemInterfaceApDiscover(o["ap-discover"], d, "ap_discover")); err != nil {
		if !fortiAPIPatch(o["ap-discover"]) {
			return fmt.Errorf("Error reading ap_discover: %v", err)
		}
	}

	if err = d.Set("fortilink_stacking", dataSourceFlattenSystemInterfaceFortilinkStacking(o["fortilink-stacking"], d, "fortilink_stacking")); err != nil {
		if !fortiAPIPatch(o["fortilink-stacking"]) {
			return fmt.Errorf("Error reading fortilink_stacking: %v", err)
		}
	}

	if err = d.Set("fortilink_neighbor_detect", dataSourceFlattenSystemInterfaceFortilinkNeighborDetect(o["fortilink-neighbor-detect"], d, "fortilink_neighbor_detect")); err != nil {
		if !fortiAPIPatch(o["fortilink-neighbor-detect"]) {
			return fmt.Errorf("Error reading fortilink_neighbor_detect: %v", err)
		}
	}

	if err = d.Set("ip_managed_by_fortiipam", dataSourceFlattenSystemInterfaceIpManagedByFortiipam(o["ip-managed-by-fortiipam"], d, "ip_managed_by_fortiipam")); err != nil {
		if !fortiAPIPatch(o["ip-managed-by-fortiipam"]) {
			return fmt.Errorf("Error reading ip_managed_by_fortiipam: %v", err)
		}
	}

	if err = d.Set("managed_subnetwork_size", dataSourceFlattenSystemInterfaceManagedSubnetworkSize(o["managed-subnetwork-size"], d, "managed_subnetwork_size")); err != nil {
		if !fortiAPIPatch(o["managed-subnetwork-size"]) {
			return fmt.Errorf("Error reading managed_subnetwork_size: %v", err)
		}
	}

	if err = d.Set("fortilink_split_interface", dataSourceFlattenSystemInterfaceFortilinkSplitInterface(o["fortilink-split-interface"], d, "fortilink_split_interface")); err != nil {
		if !fortiAPIPatch(o["fortilink-split-interface"]) {
			return fmt.Errorf("Error reading fortilink_split_interface: %v", err)
		}
	}

	if err = d.Set("internal", dataSourceFlattenSystemInterfaceInternal(o["internal"], d, "internal")); err != nil {
		if !fortiAPIPatch(o["internal"]) {
			return fmt.Errorf("Error reading internal: %v", err)
		}
	}

	if err = d.Set("fortilink_backup_link", dataSourceFlattenSystemInterfaceFortilinkBackupLink(o["fortilink-backup-link"], d, "fortilink_backup_link")); err != nil {
		if !fortiAPIPatch(o["fortilink-backup-link"]) {
			return fmt.Errorf("Error reading fortilink_backup_link: %v", err)
		}
	}

	if err = d.Set("switch_controller_access_vlan", dataSourceFlattenSystemInterfaceSwitchControllerAccessVlan(o["switch-controller-access-vlan"], d, "switch_controller_access_vlan")); err != nil {
		if !fortiAPIPatch(o["switch-controller-access-vlan"]) {
			return fmt.Errorf("Error reading switch_controller_access_vlan: %v", err)
		}
	}

	if err = d.Set("switch_controller_traffic_policy", dataSourceFlattenSystemInterfaceSwitchControllerTrafficPolicy(o["switch-controller-traffic-policy"], d, "switch_controller_traffic_policy")); err != nil {
		if !fortiAPIPatch(o["switch-controller-traffic-policy"]) {
			return fmt.Errorf("Error reading switch_controller_traffic_policy: %v", err)
		}
	}

	if err = d.Set("switch_controller_rspan_mode", dataSourceFlattenSystemInterfaceSwitchControllerRspanMode(o["switch-controller-rspan-mode"], d, "switch_controller_rspan_mode")); err != nil {
		if !fortiAPIPatch(o["switch-controller-rspan-mode"]) {
			return fmt.Errorf("Error reading switch_controller_rspan_mode: %v", err)
		}
	}

	if err = d.Set("switch_controller_netflow_collect", dataSourceFlattenSystemInterfaceSwitchControllerNetflowCollect(o["switch-controller-netflow-collect"], d, "switch_controller_netflow_collect")); err != nil {
		if !fortiAPIPatch(o["switch-controller-netflow-collect"]) {
			return fmt.Errorf("Error reading switch_controller_netflow_collect: %v", err)
		}
	}

	if err = d.Set("switch_controller_mgmt_vlan", dataSourceFlattenSystemInterfaceSwitchControllerMgmtVlan(o["switch-controller-mgmt-vlan"], d, "switch_controller_mgmt_vlan")); err != nil {
		if !fortiAPIPatch(o["switch-controller-mgmt-vlan"]) {
			return fmt.Errorf("Error reading switch_controller_mgmt_vlan: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping", dataSourceFlattenSystemInterfaceSwitchControllerIgmpSnooping(o["switch-controller-igmp-snooping"], d, "switch_controller_igmp_snooping")); err != nil {
		if !fortiAPIPatch(o["switch-controller-igmp-snooping"]) {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping_proxy", dataSourceFlattenSystemInterfaceSwitchControllerIgmpSnoopingProxy(o["switch-controller-igmp-snooping-proxy"], d, "switch_controller_igmp_snooping_proxy")); err != nil {
		if !fortiAPIPatch(o["switch-controller-igmp-snooping-proxy"]) {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping_proxy: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping_fast_leave", dataSourceFlattenSystemInterfaceSwitchControllerIgmpSnoopingFastLeave(o["switch-controller-igmp-snooping-fast-leave"], d, "switch_controller_igmp_snooping_fast_leave")); err != nil {
		if !fortiAPIPatch(o["switch-controller-igmp-snooping-fast-leave"]) {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping_fast_leave: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping", dataSourceFlattenSystemInterfaceSwitchControllerDhcpSnooping(o["switch-controller-dhcp-snooping"], d, "switch_controller_dhcp_snooping")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dhcp-snooping"]) {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping_verify_mac", dataSourceFlattenSystemInterfaceSwitchControllerDhcpSnoopingVerifyMac(o["switch-controller-dhcp-snooping-verify-mac"], d, "switch_controller_dhcp_snooping_verify_mac")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dhcp-snooping-verify-mac"]) {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping_verify_mac: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping_option82", dataSourceFlattenSystemInterfaceSwitchControllerDhcpSnoopingOption82(o["switch-controller-dhcp-snooping-option82"], d, "switch_controller_dhcp_snooping_option82")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dhcp-snooping-option82"]) {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping_option82: %v", err)
		}
	}

	if err = d.Set("dhcp_snooping_server_list", dataSourceFlattenSystemInterfaceDhcpSnoopingServerList(o["dhcp-snooping-server-list"], d, "dhcp_snooping_server_list")); err != nil {
		if !fortiAPIPatch(o["dhcp-snooping-server-list"]) {
			return fmt.Errorf("Error reading dhcp_snooping_server_list: %v", err)
		}
	}

	if err = d.Set("switch_controller_arp_inspection", dataSourceFlattenSystemInterfaceSwitchControllerArpInspection(o["switch-controller-arp-inspection"], d, "switch_controller_arp_inspection")); err != nil {
		if !fortiAPIPatch(o["switch-controller-arp-inspection"]) {
			return fmt.Errorf("Error reading switch_controller_arp_inspection: %v", err)
		}
	}

	if err = d.Set("switch_controller_learning_limit", dataSourceFlattenSystemInterfaceSwitchControllerLearningLimit(o["switch-controller-learning-limit"], d, "switch_controller_learning_limit")); err != nil {
		if !fortiAPIPatch(o["switch-controller-learning-limit"]) {
			return fmt.Errorf("Error reading switch_controller_learning_limit: %v", err)
		}
	}

	if err = d.Set("switch_controller_nac", dataSourceFlattenSystemInterfaceSwitchControllerNac(o["switch-controller-nac"], d, "switch_controller_nac")); err != nil {
		if !fortiAPIPatch(o["switch-controller-nac"]) {
			return fmt.Errorf("Error reading switch_controller_nac: %v", err)
		}
	}

	if err = d.Set("switch_controller_dynamic", dataSourceFlattenSystemInterfaceSwitchControllerDynamic(o["switch-controller-dynamic"], d, "switch_controller_dynamic")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dynamic"]) {
			return fmt.Errorf("Error reading switch_controller_dynamic: %v", err)
		}
	}

	if err = d.Set("switch_controller_feature", dataSourceFlattenSystemInterfaceSwitchControllerFeature(o["switch-controller-feature"], d, "switch_controller_feature")); err != nil {
		if !fortiAPIPatch(o["switch-controller-feature"]) {
			return fmt.Errorf("Error reading switch_controller_feature: %v", err)
		}
	}

	if err = d.Set("switch_controller_iot_scanning", dataSourceFlattenSystemInterfaceSwitchControllerIotScanning(o["switch-controller-iot-scanning"], d, "switch_controller_iot_scanning")); err != nil {
		if !fortiAPIPatch(o["switch-controller-iot-scanning"]) {
			return fmt.Errorf("Error reading switch_controller_iot_scanning: %v", err)
		}
	}

	if err = d.Set("swc_vlan", dataSourceFlattenSystemInterfaceSwcVlan(o["swc-vlan"], d, "swc_vlan")); err != nil {
		if !fortiAPIPatch(o["swc-vlan"]) {
			return fmt.Errorf("Error reading swc_vlan: %v", err)
		}
	}

	if err = d.Set("swc_first_create", dataSourceFlattenSystemInterfaceSwcFirstCreate(o["swc-first-create"], d, "swc_first_create")); err != nil {
		if !fortiAPIPatch(o["swc-first-create"]) {
			return fmt.Errorf("Error reading swc_first_create: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenSystemInterfaceColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("tagging", dataSourceFlattenSystemInterfaceTagging(o["tagging"], d, "tagging")); err != nil {
		if !fortiAPIPatch(o["tagging"]) {
			return fmt.Errorf("Error reading tagging: %v", err)
		}
	}

	if err = d.Set("eap_supplicant", dataSourceFlattenSystemInterfaceEapSupplicant(o["eap-supplicant"], d, "eap_supplicant")); err != nil {
		if !fortiAPIPatch(o["eap-supplicant"]) {
			return fmt.Errorf("Error reading eap_supplicant: %v", err)
		}
	}

	if err = d.Set("eap_method", dataSourceFlattenSystemInterfaceEapMethod(o["eap-method"], d, "eap_method")); err != nil {
		if !fortiAPIPatch(o["eap-method"]) {
			return fmt.Errorf("Error reading eap_method: %v", err)
		}
	}

	if err = d.Set("eap_identity", dataSourceFlattenSystemInterfaceEapIdentity(o["eap-identity"], d, "eap_identity")); err != nil {
		if !fortiAPIPatch(o["eap-identity"]) {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("eap_ca_cert", dataSourceFlattenSystemInterfaceEapCaCert(o["eap-ca-cert"], d, "eap_ca_cert")); err != nil {
		if !fortiAPIPatch(o["eap-ca-cert"]) {
			return fmt.Errorf("Error reading eap_ca_cert: %v", err)
		}
	}

	if err = d.Set("eap_user_cert", dataSourceFlattenSystemInterfaceEapUserCert(o["eap-user-cert"], d, "eap_user_cert")); err != nil {
		if !fortiAPIPatch(o["eap-user-cert"]) {
			return fmt.Errorf("Error reading eap_user_cert: %v", err)
		}
	}

	if err = d.Set("forward_error_correction", dataSourceFlattenSystemInterfaceForwardErrorCorrection(o["forward-error-correction"], d, "forward_error_correction")); err != nil {
		if !fortiAPIPatch(o["forward-error-correction"]) {
			return fmt.Errorf("Error reading forward_error_correction: %v", err)
		}
	}

	if err = d.Set("ipv6", dataSourceFlattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
