// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure router multicast.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterMulticastRead,
		Schema: map[string]*schema.Schema{
			"route_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"route_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"join_prune_holdtime": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"accept_register_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"accept_source_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bsr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_hash": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cisco_register_checksum": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cisco_register_checksum_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"cisco_ignore_rp_set_priority": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_rp_reachability": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_source": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_source_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_supression": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"null_register_retries": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_register_keepalive": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"spt_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"spt_threshold_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssm": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssm_range": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"register_rate_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ttl_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pim_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"neighbour_filter": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_holdtime": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cisco_exclude_genid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"propagation_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"state_refresh_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rp_candidate_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rp_candidate_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_candidate_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"multicast_flow": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"static_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"join_group": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"igmp": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_group": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"immediate_leave_group": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"last_member_query_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"last_member_query_count": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_max_response_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"query_timeout": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"router_alert_check": &schema.Schema{
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

func dataSourceRouterMulticastRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "RouterMulticast"

	o, err := c.ReadRouterMulticast(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterMulticast(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterMulticastRouteThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastRouteLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastMulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "message_interval"
	if _, ok := i["message-interval"]; ok {
		result["message_interval"] = dataSourceFlattenRouterMulticastPimSmGlobalMessageInterval(i["message-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "join_prune_holdtime"
	if _, ok := i["join-prune-holdtime"]; ok {
		result["join_prune_holdtime"] = dataSourceFlattenRouterMulticastPimSmGlobalJoinPruneHoldtime(i["join-prune-holdtime"], d, pre_append)
	}

	pre_append = pre + ".0." + "accept_register_list"
	if _, ok := i["accept-register-list"]; ok {
		result["accept_register_list"] = dataSourceFlattenRouterMulticastPimSmGlobalAcceptRegisterList(i["accept-register-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "accept_source_list"
	if _, ok := i["accept-source-list"]; ok {
		result["accept_source_list"] = dataSourceFlattenRouterMulticastPimSmGlobalAcceptSourceList(i["accept-source-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := i["bsr-candidate"]; ok {
		result["bsr_candidate"] = dataSourceFlattenRouterMulticastPimSmGlobalBsrCandidate(i["bsr-candidate"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := i["bsr-interface"]; ok {
		result["bsr_interface"] = dataSourceFlattenRouterMulticastPimSmGlobalBsrInterface(i["bsr-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := i["bsr-priority"]; ok {
		result["bsr_priority"] = dataSourceFlattenRouterMulticastPimSmGlobalBsrPriority(i["bsr-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := i["bsr-hash"]; ok {
		result["bsr_hash"] = dataSourceFlattenRouterMulticastPimSmGlobalBsrHash(i["bsr-hash"], d, pre_append)
	}

	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := i["bsr-allow-quick-refresh"]; ok {
		result["bsr_allow_quick_refresh"] = dataSourceFlattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_register_checksum"
	if _, ok := i["cisco-register-checksum"]; ok {
		result["cisco_register_checksum"] = dataSourceFlattenRouterMulticastPimSmGlobalCiscoRegisterChecksum(i["cisco-register-checksum"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_register_checksum_group"
	if _, ok := i["cisco-register-checksum-group"]; ok {
		result["cisco_register_checksum_group"] = dataSourceFlattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(i["cisco-register-checksum-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := i["cisco-crp-prefix"]; ok {
		result["cisco_crp_prefix"] = dataSourceFlattenRouterMulticastPimSmGlobalCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append)
	}

	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := i["cisco-ignore-rp-set-priority"]; ok {
		result["cisco_ignore_rp_set_priority"] = dataSourceFlattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(i["cisco-ignore-rp-set-priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rp_reachability"
	if _, ok := i["register-rp-reachability"]; ok {
		result["register_rp_reachability"] = dataSourceFlattenRouterMulticastPimSmGlobalRegisterRpReachability(i["register-rp-reachability"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_source"
	if _, ok := i["register-source"]; ok {
		result["register_source"] = dataSourceFlattenRouterMulticastPimSmGlobalRegisterSource(i["register-source"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_source_interface"
	if _, ok := i["register-source-interface"]; ok {
		result["register_source_interface"] = dataSourceFlattenRouterMulticastPimSmGlobalRegisterSourceInterface(i["register-source-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_source_ip"
	if _, ok := i["register-source-ip"]; ok {
		result["register_source_ip"] = dataSourceFlattenRouterMulticastPimSmGlobalRegisterSourceIp(i["register-source-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_supression"
	if _, ok := i["register-supression"]; ok {
		result["register_supression"] = dataSourceFlattenRouterMulticastPimSmGlobalRegisterSupression(i["register-supression"], d, pre_append)
	}

	pre_append = pre + ".0." + "null_register_retries"
	if _, ok := i["null-register-retries"]; ok {
		result["null_register_retries"] = dataSourceFlattenRouterMulticastPimSmGlobalNullRegisterRetries(i["null-register-retries"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_register_keepalive"
	if _, ok := i["rp-register-keepalive"]; ok {
		result["rp_register_keepalive"] = dataSourceFlattenRouterMulticastPimSmGlobalRpRegisterKeepalive(i["rp-register-keepalive"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := i["spt-threshold"]; ok {
		result["spt_threshold"] = dataSourceFlattenRouterMulticastPimSmGlobalSptThreshold(i["spt-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := i["spt-threshold-group"]; ok {
		result["spt_threshold_group"] = dataSourceFlattenRouterMulticastPimSmGlobalSptThresholdGroup(i["spt-threshold-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssm"
	if _, ok := i["ssm"]; ok {
		result["ssm"] = dataSourceFlattenRouterMulticastPimSmGlobalSsm(i["ssm"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssm_range"
	if _, ok := i["ssm-range"]; ok {
		result["ssm_range"] = dataSourceFlattenRouterMulticastPimSmGlobalSsmRange(i["ssm-range"], d, pre_append)
	}

	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = dataSourceFlattenRouterMulticastPimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = dataSourceFlattenRouterMulticastPimSmGlobalRpAddress(i["rp-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenRouterMulticastPimSmGlobalMessageInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalJoinPruneHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalAcceptRegisterList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalAcceptSourceList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalBsrCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalBsrInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalBsrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalBsrHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalCiscoRegisterChecksum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRegisterRpReachability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRegisterSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRegisterSourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRegisterSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRegisterSupression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalNullRegisterRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRpRegisterKeepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalSptThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalSptThresholdGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalSsm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalSsmRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterMulticastPimSmGlobalRpAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			tmp["ip_address"] = dataSourceFlattenRouterMulticastPimSmGlobalRpAddressIpAddress(i["ip-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := i["group"]; ok {
			tmp["group"] = dataSourceFlattenRouterMulticastPimSmGlobalRpAddressGroup(i["group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticastPimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRpAddressIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastPimSmGlobalRpAddressGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterMulticastInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl_threshold"
		if _, ok := i["ttl-threshold"]; ok {
			tmp["ttl_threshold"] = dataSourceFlattenRouterMulticastInterfaceTtlThreshold(i["ttl-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := i["pim-mode"]; ok {
			tmp["pim_mode"] = dataSourceFlattenRouterMulticastInterfacePimMode(i["pim-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := i["passive"]; ok {
			tmp["passive"] = dataSourceFlattenRouterMulticastInterfacePassive(i["passive"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterMulticastInterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_filter"
		if _, ok := i["neighbour-filter"]; ok {
			tmp["neighbour_filter"] = dataSourceFlattenRouterMulticastInterfaceNeighbourFilter(i["neighbour-filter"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterMulticastInterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := i["hello-holdtime"]; ok {
			tmp["hello_holdtime"] = dataSourceFlattenRouterMulticastInterfaceHelloHoldtime(i["hello-holdtime"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_exclude_genid"
		if _, ok := i["cisco-exclude-genid"]; ok {
			tmp["cisco_exclude_genid"] = dataSourceFlattenRouterMulticastInterfaceCiscoExcludeGenid(i["cisco-exclude-genid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := i["dr-priority"]; ok {
			tmp["dr_priority"] = dataSourceFlattenRouterMulticastInterfaceDrPriority(i["dr-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "propagation_delay"
		if _, ok := i["propagation-delay"]; ok {
			tmp["propagation_delay"] = dataSourceFlattenRouterMulticastInterfacePropagationDelay(i["propagation-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "state_refresh_interval"
		if _, ok := i["state-refresh-interval"]; ok {
			tmp["state_refresh_interval"] = dataSourceFlattenRouterMulticastInterfaceStateRefreshInterval(i["state-refresh-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := i["rp-candidate"]; ok {
			tmp["rp_candidate"] = dataSourceFlattenRouterMulticastInterfaceRpCandidate(i["rp-candidate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := i["rp-candidate-group"]; ok {
			tmp["rp_candidate_group"] = dataSourceFlattenRouterMulticastInterfaceRpCandidateGroup(i["rp-candidate-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := i["rp-candidate-priority"]; ok {
			tmp["rp_candidate_priority"] = dataSourceFlattenRouterMulticastInterfaceRpCandidatePriority(i["rp-candidate-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := i["rp-candidate-interval"]; ok {
			tmp["rp_candidate_interval"] = dataSourceFlattenRouterMulticastInterfaceRpCandidateInterval(i["rp-candidate-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := i["multicast-flow"]; ok {
			tmp["multicast_flow"] = dataSourceFlattenRouterMulticastInterfaceMulticastFlow(i["multicast-flow"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := i["static-group"]; ok {
			tmp["static_group"] = dataSourceFlattenRouterMulticastInterfaceStaticGroup(i["static-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "join_group"
		if _, ok := i["join-group"]; ok {
			tmp["join_group"] = dataSourceFlattenRouterMulticastInterfaceJoinGroup(i["join-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := i["igmp"]; ok {
			tmp["igmp"] = dataSourceFlattenRouterMulticastInterfaceIgmp(i["igmp"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticastInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceTtlThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfacePimMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfacePassive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceNeighbourFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceCiscoExcludeGenid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceDrPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfacePropagationDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceStateRefreshInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceRpCandidate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceRpCandidateGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceRpCandidatePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceRpCandidateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceMulticastFlow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceStaticGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceJoinGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = dataSourceFlattenRouterMulticastInterfaceJoinGroupAddress(i["address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticastInterfaceJoinGroupAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := i["access-group"]; ok {
		result["access_group"] = dataSourceFlattenRouterMulticastInterfaceIgmpAccessGroup(i["access-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = dataSourceFlattenRouterMulticastInterfaceIgmpVersion(i["version"], d, pre_append)
	}

	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := i["immediate-leave-group"]; ok {
		result["immediate_leave_group"] = dataSourceFlattenRouterMulticastInterfaceIgmpImmediateLeaveGroup(i["immediate-leave-group"], d, pre_append)
	}

	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := i["last-member-query-interval"]; ok {
		result["last_member_query_interval"] = dataSourceFlattenRouterMulticastInterfaceIgmpLastMemberQueryInterval(i["last-member-query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := i["last-member-query-count"]; ok {
		result["last_member_query_count"] = dataSourceFlattenRouterMulticastInterfaceIgmpLastMemberQueryCount(i["last-member-query-count"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := i["query-max-response-time"]; ok {
		result["query_max_response_time"] = dataSourceFlattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(i["query-max-response-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_interval"
	if _, ok := i["query-interval"]; ok {
		result["query_interval"] = dataSourceFlattenRouterMulticastInterfaceIgmpQueryInterval(i["query-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "query_timeout"
	if _, ok := i["query-timeout"]; ok {
		result["query_timeout"] = dataSourceFlattenRouterMulticastInterfaceIgmpQueryTimeout(i["query-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := i["router-alert-check"]; ok {
		result["router_alert_check"] = dataSourceFlattenRouterMulticastInterfaceIgmpRouterAlertCheck(i["router-alert-check"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenRouterMulticastInterfaceIgmpAccessGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpImmediateLeaveGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpLastMemberQueryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpLastMemberQueryCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpQueryInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpQueryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticastInterfaceIgmpRouterAlertCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterMulticast(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("route_threshold", dataSourceFlattenRouterMulticastRouteThreshold(o["route-threshold"], d, "route_threshold")); err != nil {
		if !fortiAPIPatch(o["route-threshold"]) {
			return fmt.Errorf("Error reading route_threshold: %v", err)
		}
	}

	if err = d.Set("route_limit", dataSourceFlattenRouterMulticastRouteLimit(o["route-limit"], d, "route_limit")); err != nil {
		if !fortiAPIPatch(o["route-limit"]) {
			return fmt.Errorf("Error reading route_limit: %v", err)
		}
	}

	if err = d.Set("multicast_routing", dataSourceFlattenRouterMulticastMulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if err = d.Set("pim_sm_global", dataSourceFlattenRouterMulticastPimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
		if !fortiAPIPatch(o["pim-sm-global"]) {
			return fmt.Errorf("Error reading pim_sm_global: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterMulticastInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterMulticastFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
