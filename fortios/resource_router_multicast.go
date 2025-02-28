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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterMulticast() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastUpdate,
		Read:   resourceRouterMulticastRead,
		Update: resourceRouterMulticastUpdate,
		Delete: resourceRouterMulticastDelete,

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
			"route_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"route_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"join_prune_holdtime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"accept_register_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"accept_source_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"bsr_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"bsr_hash": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_register_checksum": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_register_checksum_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_ignore_rp_set_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_rp_reachability": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_source_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"register_source_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"register_supression": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"null_register_retries": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),
							Optional:     true,
							Computed:     true,
						},
						"rp_register_keepalive": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"spt_threshold": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spt_threshold_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"ssm": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssm_range": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"register_rate_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"pim_use_sdwan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"pim_sm_global_vrf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrf": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 511),
							Optional:     true,
							Computed:     true,
						},
						"bsr_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bsr_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"bsr_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"bsr_hash": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"bsr_allow_quick_refresh": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cisco_crp_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"ttl_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"pim_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"neighbour_filter": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"hello_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_holdtime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"cisco_exclude_genid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dr_priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"propagation_delay": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(100, 5000),
							Optional:     true,
							Computed:     true,
						},
						"state_refresh_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"rp_candidate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rp_candidate_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"rp_candidate_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"rp_candidate_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16383),
							Optional:     true,
							Computed:     true,
						},
						"multicast_flow": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"static_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"rpf_nbr_fail_back": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rpf_nbr_fail_back_filter": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"join_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"igmp": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"access_group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"immediate_leave_group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"last_member_query_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"last_member_query_count": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(2, 7),
										Optional:     true,
										Computed:     true,
									},
									"query_max_response_time": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 25),
										Optional:     true,
										Computed:     true,
									},
									"query_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"query_timeout": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(60, 900),
										Optional:     true,
										Computed:     true,
									},
									"router_alert_check": &schema.Schema{
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

func resourceRouterMulticastUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterMulticast(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterMulticast(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticast")
	}

	return resourceRouterMulticastRead(d, m)
}

func resourceRouterMulticastDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterMulticast(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterMulticast resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterMulticast(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastRouteThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastRouteLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastMulticastRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobal(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "message_interval"
	if _, ok := i["message-interval"]; ok {
		result["message_interval"] = flattenRouterMulticastPimSmGlobalMessageInterval(i["message-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "join_prune_holdtime"
	if _, ok := i["join-prune-holdtime"]; ok {
		result["join_prune_holdtime"] = flattenRouterMulticastPimSmGlobalJoinPruneHoldtime(i["join-prune-holdtime"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "accept_register_list"
	if _, ok := i["accept-register-list"]; ok {
		result["accept_register_list"] = flattenRouterMulticastPimSmGlobalAcceptRegisterList(i["accept-register-list"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "accept_source_list"
	if _, ok := i["accept-source-list"]; ok {
		result["accept_source_list"] = flattenRouterMulticastPimSmGlobalAcceptSourceList(i["accept-source-list"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := i["bsr-candidate"]; ok {
		result["bsr_candidate"] = flattenRouterMulticastPimSmGlobalBsrCandidate(i["bsr-candidate"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := i["bsr-interface"]; ok {
		result["bsr_interface"] = flattenRouterMulticastPimSmGlobalBsrInterface(i["bsr-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := i["bsr-priority"]; ok {
		result["bsr_priority"] = flattenRouterMulticastPimSmGlobalBsrPriority(i["bsr-priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := i["bsr-hash"]; ok {
		result["bsr_hash"] = flattenRouterMulticastPimSmGlobalBsrHash(i["bsr-hash"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := i["bsr-allow-quick-refresh"]; ok {
		result["bsr_allow_quick_refresh"] = flattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh(i["bsr-allow-quick-refresh"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cisco_register_checksum"
	if _, ok := i["cisco-register-checksum"]; ok {
		result["cisco_register_checksum"] = flattenRouterMulticastPimSmGlobalCiscoRegisterChecksum(i["cisco-register-checksum"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cisco_register_checksum_group"
	if _, ok := i["cisco-register-checksum-group"]; ok {
		result["cisco_register_checksum_group"] = flattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(i["cisco-register-checksum-group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := i["cisco-crp-prefix"]; ok {
		result["cisco_crp_prefix"] = flattenRouterMulticastPimSmGlobalCiscoCrpPrefix(i["cisco-crp-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := i["cisco-ignore-rp-set-priority"]; ok {
		result["cisco_ignore_rp_set_priority"] = flattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(i["cisco-ignore-rp-set-priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_rp_reachability"
	if _, ok := i["register-rp-reachability"]; ok {
		result["register_rp_reachability"] = flattenRouterMulticastPimSmGlobalRegisterRpReachability(i["register-rp-reachability"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_source"
	if _, ok := i["register-source"]; ok {
		result["register_source"] = flattenRouterMulticastPimSmGlobalRegisterSource(i["register-source"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_source_interface"
	if _, ok := i["register-source-interface"]; ok {
		result["register_source_interface"] = flattenRouterMulticastPimSmGlobalRegisterSourceInterface(i["register-source-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_source_ip"
	if _, ok := i["register-source-ip"]; ok {
		result["register_source_ip"] = flattenRouterMulticastPimSmGlobalRegisterSourceIp(i["register-source-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_supression"
	if _, ok := i["register-supression"]; ok {
		result["register_supression"] = flattenRouterMulticastPimSmGlobalRegisterSupression(i["register-supression"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "null_register_retries"
	if _, ok := i["null-register-retries"]; ok {
		result["null_register_retries"] = flattenRouterMulticastPimSmGlobalNullRegisterRetries(i["null-register-retries"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rp_register_keepalive"
	if _, ok := i["rp-register-keepalive"]; ok {
		result["rp_register_keepalive"] = flattenRouterMulticastPimSmGlobalRpRegisterKeepalive(i["rp-register-keepalive"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := i["spt-threshold"]; ok {
		result["spt_threshold"] = flattenRouterMulticastPimSmGlobalSptThreshold(i["spt-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := i["spt-threshold-group"]; ok {
		result["spt_threshold_group"] = flattenRouterMulticastPimSmGlobalSptThresholdGroup(i["spt-threshold-group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssm"
	if _, ok := i["ssm"]; ok {
		result["ssm"] = flattenRouterMulticastPimSmGlobalSsm(i["ssm"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ssm_range"
	if _, ok := i["ssm-range"]; ok {
		result["ssm_range"] = flattenRouterMulticastPimSmGlobalSsmRange(i["ssm-range"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = flattenRouterMulticastPimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := i["pim-use-sdwan"]; ok {
		result["pim_use_sdwan"] = flattenRouterMulticastPimSmGlobalPimUseSdwan(i["pim-use-sdwan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = flattenRouterMulticastPimSmGlobalRpAddress(i["rp-address"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticastPimSmGlobalMessageInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalJoinPruneHoldtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalAcceptRegisterList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalAcceptSourceList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrCandidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalBsrPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalBsrHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoRegisterChecksum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterRpReachability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSourceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterSupression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalNullRegisterRetries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalRpRegisterKeepalive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalSptThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSptThresholdGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSsm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalSsmRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalPimUseSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterMulticastPimSmGlobalRpAddressId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if cur_v, ok := i["ip-address"]; ok {
			tmp["ip_address"] = flattenRouterMulticastPimSmGlobalRpAddressIpAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if cur_v, ok := i["group"]; ok {
			tmp["group"] = flattenRouterMulticastPimSmGlobalRpAddressGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterMulticastPimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalRpAddressIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalRpAddressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["vrf"] = flattenRouterMulticastPimSmGlobalVrfVrf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
		if cur_v, ok := i["bsr-candidate"]; ok {
			tmp["bsr_candidate"] = flattenRouterMulticastPimSmGlobalVrfBsrCandidate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
		if cur_v, ok := i["bsr-interface"]; ok {
			tmp["bsr_interface"] = flattenRouterMulticastPimSmGlobalVrfBsrInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
		if cur_v, ok := i["bsr-priority"]; ok {
			tmp["bsr_priority"] = flattenRouterMulticastPimSmGlobalVrfBsrPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
		if cur_v, ok := i["bsr-hash"]; ok {
			tmp["bsr_hash"] = flattenRouterMulticastPimSmGlobalVrfBsrHash(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
		if cur_v, ok := i["bsr-allow-quick-refresh"]; ok {
			tmp["bsr_allow_quick_refresh"] = flattenRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
		if cur_v, ok := i["cisco-crp-prefix"]; ok {
			tmp["cisco_crp_prefix"] = flattenRouterMulticastPimSmGlobalVrfCiscoCrpPrefix(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
		if cur_v, ok := i["rp-address"]; ok {
			tmp["rp_address"] = flattenRouterMulticastPimSmGlobalVrfRpAddress(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vrf", d)
	return result
}

func flattenRouterMulticastPimSmGlobalVrfVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalVrfBsrCandidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfBsrInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfBsrPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalVrfBsrHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfCiscoCrpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfRpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterMulticastPimSmGlobalVrfRpAddressId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if cur_v, ok := i["ip-address"]; ok {
			tmp["ip_address"] = flattenRouterMulticastPimSmGlobalVrfRpAddressIpAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if cur_v, ok := i["group"]; ok {
			tmp["group"] = flattenRouterMulticastPimSmGlobalVrfRpAddressGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastPimSmGlobalVrfRpAddressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterMulticastInterfaceName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl_threshold"
		if cur_v, ok := i["ttl-threshold"]; ok {
			tmp["ttl_threshold"] = flattenRouterMulticastInterfaceTtlThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if cur_v, ok := i["pim-mode"]; ok {
			tmp["pim_mode"] = flattenRouterMulticastInterfacePimMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if cur_v, ok := i["passive"]; ok {
			tmp["passive"] = flattenRouterMulticastInterfacePassive(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if cur_v, ok := i["bfd"]; ok {
			tmp["bfd"] = flattenRouterMulticastInterfaceBfd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_filter"
		if cur_v, ok := i["neighbour-filter"]; ok {
			tmp["neighbour_filter"] = flattenRouterMulticastInterfaceNeighbourFilter(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if cur_v, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterMulticastInterfaceHelloInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if cur_v, ok := i["hello-holdtime"]; ok {
			tmp["hello_holdtime"] = flattenRouterMulticastInterfaceHelloHoldtime(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_exclude_genid"
		if cur_v, ok := i["cisco-exclude-genid"]; ok {
			tmp["cisco_exclude_genid"] = flattenRouterMulticastInterfaceCiscoExcludeGenid(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if cur_v, ok := i["dr-priority"]; ok {
			tmp["dr_priority"] = flattenRouterMulticastInterfaceDrPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "propagation_delay"
		if cur_v, ok := i["propagation-delay"]; ok {
			tmp["propagation_delay"] = flattenRouterMulticastInterfacePropagationDelay(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "state_refresh_interval"
		if cur_v, ok := i["state-refresh-interval"]; ok {
			tmp["state_refresh_interval"] = flattenRouterMulticastInterfaceStateRefreshInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if cur_v, ok := i["rp-candidate"]; ok {
			tmp["rp_candidate"] = flattenRouterMulticastInterfaceRpCandidate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if cur_v, ok := i["rp-candidate-group"]; ok {
			tmp["rp_candidate_group"] = flattenRouterMulticastInterfaceRpCandidateGroup(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if cur_v, ok := i["rp-candidate-priority"]; ok {
			tmp["rp_candidate_priority"] = flattenRouterMulticastInterfaceRpCandidatePriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if cur_v, ok := i["rp-candidate-interval"]; ok {
			tmp["rp_candidate_interval"] = flattenRouterMulticastInterfaceRpCandidateInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if cur_v, ok := i["multicast-flow"]; ok {
			tmp["multicast_flow"] = flattenRouterMulticastInterfaceMulticastFlow(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if cur_v, ok := i["static-group"]; ok {
			tmp["static_group"] = flattenRouterMulticastInterfaceStaticGroup(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back"
		if cur_v, ok := i["rpf-nbr-fail-back"]; ok {
			tmp["rpf_nbr_fail_back"] = flattenRouterMulticastInterfaceRpfNbrFailBack(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back_filter"
		if cur_v, ok := i["rpf-nbr-fail-back-filter"]; ok {
			tmp["rpf_nbr_fail_back_filter"] = flattenRouterMulticastInterfaceRpfNbrFailBackFilter(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "join_group"
		if cur_v, ok := i["join-group"]; ok {
			tmp["join_group"] = flattenRouterMulticastInterfaceJoinGroup(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if cur_v, ok := i["igmp"]; ok {
			tmp["igmp"] = flattenRouterMulticastInterfaceIgmp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterMulticastInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceTtlThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfacePimMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfacePassive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceNeighbourFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceCiscoExcludeGenid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceDrPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfacePropagationDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceStateRefreshInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceRpCandidate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidateGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpCandidatePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceRpCandidateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceMulticastFlow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceStaticGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpfNbrFailBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceRpfNbrFailBackFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceJoinGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["address"] = flattenRouterMulticastInterfaceJoinGroupAddress(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "address", d)
	return result
}

func flattenRouterMulticastInterfaceJoinGroupAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := i["access-group"]; ok {
		result["access_group"] = flattenRouterMulticastInterfaceIgmpAccessGroup(i["access-group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = flattenRouterMulticastInterfaceIgmpVersion(i["version"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := i["immediate-leave-group"]; ok {
		result["immediate_leave_group"] = flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup(i["immediate-leave-group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := i["last-member-query-interval"]; ok {
		result["last_member_query_interval"] = flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval(i["last-member-query-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := i["last-member-query-count"]; ok {
		result["last_member_query_count"] = flattenRouterMulticastInterfaceIgmpLastMemberQueryCount(i["last-member-query-count"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := i["query-max-response-time"]; ok {
		result["query_max_response_time"] = flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(i["query-max-response-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "query_interval"
	if _, ok := i["query-interval"]; ok {
		result["query_interval"] = flattenRouterMulticastInterfaceIgmpQueryInterval(i["query-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "query_timeout"
	if _, ok := i["query-timeout"]; ok {
		result["query_timeout"] = flattenRouterMulticastInterfaceIgmpQueryTimeout(i["query-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := i["router-alert-check"]; ok {
		result["router_alert_check"] = flattenRouterMulticastInterfaceIgmpRouterAlertCheck(i["router-alert-check"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticastInterfaceIgmpAccessGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpImmediateLeaveGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceIgmpLastMemberQueryCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceIgmpQueryMaxResponseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceIgmpQueryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceIgmpQueryTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterMulticastInterfaceIgmpRouterAlertCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMulticast(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("route_threshold", flattenRouterMulticastRouteThreshold(o["route-threshold"], d, "route_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["route-threshold"]) {
			return fmt.Errorf("Error reading route_threshold: %v", err)
		}
	}

	if err = d.Set("route_limit", flattenRouterMulticastRouteLimit(o["route-limit"], d, "route_limit", sv)); err != nil {
		if !fortiAPIPatch(o["route-limit"]) {
			return fmt.Errorf("Error reading route_limit: %v", err)
		}
	}

	if err = d.Set("multicast_routing", flattenRouterMulticastMulticastRouting(o["multicast-routing"], d, "multicast_routing", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("pim_sm_global", flattenRouterMulticastPimSmGlobal(o["pim-sm-global"], d, "pim_sm_global", sv)); err != nil {
			if !fortiAPIPatch(o["pim-sm-global"]) {
				return fmt.Errorf("Error reading pim_sm_global: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global"); ok {
			if err = d.Set("pim_sm_global", flattenRouterMulticastPimSmGlobal(o["pim-sm-global"], d, "pim_sm_global", sv)); err != nil {
				if !fortiAPIPatch(o["pim-sm-global"]) {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("pim_sm_global_vrf", flattenRouterMulticastPimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf", sv)); err != nil {
			if !fortiAPIPatch(o["pim-sm-global-vrf"]) {
				return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global_vrf"); ok {
			if err = d.Set("pim_sm_global_vrf", flattenRouterMulticastPimSmGlobalVrf(o["pim-sm-global-vrf"], d, "pim_sm_global_vrf", sv)); err != nil {
				if !fortiAPIPatch(o["pim-sm-global-vrf"]) {
					return fmt.Errorf("Error reading pim_sm_global_vrf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("interface", flattenRouterMulticastInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterMulticastInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterMulticastFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterMulticastRouteThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastRouteLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastMulticastRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobal(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "message_interval"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["message-interval"] = nil
		} else {
			result["message-interval"], _ = expandRouterMulticastPimSmGlobalMessageInterval(d, i["message_interval"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "join_prune_holdtime"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["join-prune-holdtime"] = nil
		} else {
			result["join-prune-holdtime"], _ = expandRouterMulticastPimSmGlobalJoinPruneHoldtime(d, i["join_prune_holdtime"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "accept_register_list"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["accept-register-list"] = nil
		} else {
			result["accept-register-list"], _ = expandRouterMulticastPimSmGlobalAcceptRegisterList(d, i["accept_register_list"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "accept_source_list"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["accept-source-list"] = nil
		} else {
			result["accept-source-list"], _ = expandRouterMulticastPimSmGlobalAcceptSourceList(d, i["accept_source_list"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_candidate"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-candidate"] = nil
		} else {
			result["bsr-candidate"], _ = expandRouterMulticastPimSmGlobalBsrCandidate(d, i["bsr_candidate"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_interface"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-interface"] = nil
		} else {
			result["bsr-interface"], _ = expandRouterMulticastPimSmGlobalBsrInterface(d, i["bsr_interface"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_priority"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-priority"] = nil
		} else {
			result["bsr-priority"], _ = expandRouterMulticastPimSmGlobalBsrPriority(d, i["bsr_priority"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_hash"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-hash"] = nil
		} else {
			result["bsr-hash"], _ = expandRouterMulticastPimSmGlobalBsrHash(d, i["bsr_hash"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bsr_allow_quick_refresh"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bsr-allow-quick-refresh"] = nil
		} else {
			result["bsr-allow-quick-refresh"], _ = expandRouterMulticastPimSmGlobalBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "cisco_register_checksum"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["cisco-register-checksum"] = nil
		} else {
			result["cisco-register-checksum"], _ = expandRouterMulticastPimSmGlobalCiscoRegisterChecksum(d, i["cisco_register_checksum"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "cisco_register_checksum_group"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["cisco-register-checksum-group"] = nil
		} else {
			result["cisco-register-checksum-group"], _ = expandRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(d, i["cisco_register_checksum_group"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "cisco_crp_prefix"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["cisco-crp-prefix"] = nil
		} else {
			result["cisco-crp-prefix"], _ = expandRouterMulticastPimSmGlobalCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "cisco_ignore_rp_set_priority"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["cisco-ignore-rp-set-priority"] = nil
		} else {
			result["cisco-ignore-rp-set-priority"], _ = expandRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(d, i["cisco_ignore_rp_set_priority"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_rp_reachability"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-rp-reachability"] = nil
		} else {
			result["register-rp-reachability"], _ = expandRouterMulticastPimSmGlobalRegisterRpReachability(d, i["register_rp_reachability"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_source"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-source"] = nil
		} else {
			result["register-source"], _ = expandRouterMulticastPimSmGlobalRegisterSource(d, i["register_source"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_source_interface"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-source-interface"] = nil
		} else {
			result["register-source-interface"], _ = expandRouterMulticastPimSmGlobalRegisterSourceInterface(d, i["register_source_interface"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_source_ip"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-source-ip"] = nil
		} else {
			result["register-source-ip"], _ = expandRouterMulticastPimSmGlobalRegisterSourceIp(d, i["register_source_ip"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_supression"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-supression"] = nil
		} else {
			result["register-supression"], _ = expandRouterMulticastPimSmGlobalRegisterSupression(d, i["register_supression"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "null_register_retries"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["null-register-retries"] = nil
		} else {
			result["null-register-retries"], _ = expandRouterMulticastPimSmGlobalNullRegisterRetries(d, i["null_register_retries"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "rp_register_keepalive"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["rp-register-keepalive"] = nil
		} else {
			result["rp-register-keepalive"], _ = expandRouterMulticastPimSmGlobalRpRegisterKeepalive(d, i["rp_register_keepalive"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "spt_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["spt-threshold"] = nil
		} else {
			result["spt-threshold"], _ = expandRouterMulticastPimSmGlobalSptThreshold(d, i["spt_threshold"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "spt_threshold_group"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["spt-threshold-group"] = nil
		} else {
			result["spt-threshold-group"], _ = expandRouterMulticastPimSmGlobalSptThresholdGroup(d, i["spt_threshold_group"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ssm"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ssm"] = nil
		} else {
			result["ssm"], _ = expandRouterMulticastPimSmGlobalSsm(d, i["ssm"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "ssm_range"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["ssm-range"] = nil
		} else {
			result["ssm-range"], _ = expandRouterMulticastPimSmGlobalSsmRange(d, i["ssm_range"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["register-rate-limit"] = nil
		} else {
			result["register-rate-limit"], _ = expandRouterMulticastPimSmGlobalRegisterRateLimit(d, i["register_rate_limit"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "pim_use_sdwan"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["pim-use-sdwan"] = nil
		} else {
			result["pim-use-sdwan"], _ = expandRouterMulticastPimSmGlobalPimUseSdwan(d, i["pim_use_sdwan"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "rp_address"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["rp-address"] = make([]struct{}, 0)
		} else {
			result["rp-address"], _ = expandRouterMulticastPimSmGlobalRpAddress(d, i["rp_address"], pre_append, sv)
		}
	} else {
		result["rp-address"] = make([]string, 0)
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalMessageInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalJoinPruneHoldtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalAcceptRegisterList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalAcceptSourceList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrCandidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoRegisterChecksum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoRegisterChecksumGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalCiscoIgnoreRpSetPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterRpReachability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSourceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterSupression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalNullRegisterRetries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpRegisterKeepalive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSptThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSptThresholdGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSsm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalSsmRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRegisterRateLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalPimUseSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterMulticastPimSmGlobalRpAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-address"], _ = expandRouterMulticastPimSmGlobalRpAddressIpAddress(d, i["ip_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["group"], _ = expandRouterMulticastPimSmGlobalRpAddressGroup(d, i["group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalRpAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddressIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalRpAddressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["vrf"], _ = expandRouterMulticastPimSmGlobalVrfVrf(d, i["vrf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_candidate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-candidate"], _ = expandRouterMulticastPimSmGlobalVrfBsrCandidate(d, i["bsr_candidate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-interface"], _ = expandRouterMulticastPimSmGlobalVrfBsrInterface(d, i["bsr_interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bsr-interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-priority"], _ = expandRouterMulticastPimSmGlobalVrfBsrPriority(d, i["bsr_priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bsr-priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_hash"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-hash"], _ = expandRouterMulticastPimSmGlobalVrfBsrHash(d, i["bsr_hash"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bsr_allow_quick_refresh"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bsr-allow-quick-refresh"], _ = expandRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh(d, i["bsr_allow_quick_refresh"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_crp_prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cisco-crp-prefix"], _ = expandRouterMulticastPimSmGlobalVrfCiscoCrpPrefix(d, i["cisco_crp_prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-address"], _ = expandRouterMulticastPimSmGlobalVrfRpAddress(d, i["rp_address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rp-address"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalVrfVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrCandidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfBsrAllowQuickRefresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfCiscoCrpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterMulticastPimSmGlobalVrfRpAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-address"], _ = expandRouterMulticastPimSmGlobalVrfRpAddressIpAddress(d, i["ip_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["group"], _ = expandRouterMulticastPimSmGlobalVrfRpAddressGroup(d, i["group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastPimSmGlobalVrfRpAddressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterMulticastInterfaceName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ttl-threshold"], _ = expandRouterMulticastInterfaceTtlThreshold(d, i["ttl_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pim_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pim-mode"], _ = expandRouterMulticastInterfacePimMode(d, i["pim_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passive"], _ = expandRouterMulticastInterfacePassive(d, i["passive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterMulticastInterfaceBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbour_filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["neighbour-filter"], _ = expandRouterMulticastInterfaceNeighbourFilter(d, i["neighbour_filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["neighbour-filter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterMulticastInterfaceHelloInterval(d, i["hello_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-holdtime"], _ = expandRouterMulticastInterfaceHelloHoldtime(d, i["hello_holdtime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cisco_exclude_genid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cisco-exclude-genid"], _ = expandRouterMulticastInterfaceCiscoExcludeGenid(d, i["cisco_exclude_genid"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dr_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dr-priority"], _ = expandRouterMulticastInterfaceDrPriority(d, i["dr_priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "propagation_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["propagation-delay"], _ = expandRouterMulticastInterfacePropagationDelay(d, i["propagation_delay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "state_refresh_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["state-refresh-interval"], _ = expandRouterMulticastInterfaceStateRefreshInterval(d, i["state_refresh_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate"], _ = expandRouterMulticastInterfaceRpCandidate(d, i["rp_candidate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate-group"], _ = expandRouterMulticastInterfaceRpCandidateGroup(d, i["rp_candidate_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rp-candidate-group"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate-priority"], _ = expandRouterMulticastInterfaceRpCandidatePriority(d, i["rp_candidate_priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rp_candidate_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rp-candidate-interval"], _ = expandRouterMulticastInterfaceRpCandidateInterval(d, i["rp_candidate_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "multicast_flow"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["multicast-flow"], _ = expandRouterMulticastInterfaceMulticastFlow(d, i["multicast_flow"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["multicast-flow"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "static_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["static-group"], _ = expandRouterMulticastInterfaceStaticGroup(d, i["static_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["static-group"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rpf-nbr-fail-back"], _ = expandRouterMulticastInterfaceRpfNbrFailBack(d, i["rpf_nbr_fail_back"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rpf_nbr_fail_back_filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rpf-nbr-fail-back-filter"], _ = expandRouterMulticastInterfaceRpfNbrFailBackFilter(d, i["rpf_nbr_fail_back_filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rpf-nbr-fail-back-filter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "join_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["join-group"], _ = expandRouterMulticastInterfaceJoinGroup(d, i["join_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["join-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "igmp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["igmp"], _ = expandRouterMulticastInterfaceIgmp(d, i["igmp"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["igmp"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceTtlThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePimMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePassive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceNeighbourFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceHelloHoldtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceCiscoExcludeGenid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceDrPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfacePropagationDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceStateRefreshInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidateGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidatePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpCandidateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceMulticastFlow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceStaticGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpfNbrFailBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceRpfNbrFailBackFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceJoinGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandRouterMulticastInterfaceJoinGroupAddress(d, i["address"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastInterfaceJoinGroupAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "access_group"
	if _, ok := d.GetOk(pre_append); ok {
		result["access-group"], _ = expandRouterMulticastInterfaceIgmpAccessGroup(d, i["access_group"], pre_append, sv)
	}
	pre_append = pre + ".0." + "version"
	if _, ok := d.GetOk(pre_append); ok {
		result["version"], _ = expandRouterMulticastInterfaceIgmpVersion(d, i["version"], pre_append, sv)
	}
	pre_append = pre + ".0." + "immediate_leave_group"
	if _, ok := d.GetOk(pre_append); ok {
		result["immediate-leave-group"], _ = expandRouterMulticastInterfaceIgmpImmediateLeaveGroup(d, i["immediate_leave_group"], pre_append, sv)
	}
	pre_append = pre + ".0." + "last_member_query_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["last-member-query-interval"], _ = expandRouterMulticastInterfaceIgmpLastMemberQueryInterval(d, i["last_member_query_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "last_member_query_count"
	if _, ok := d.GetOk(pre_append); ok {
		result["last-member-query-count"], _ = expandRouterMulticastInterfaceIgmpLastMemberQueryCount(d, i["last_member_query_count"], pre_append, sv)
	}
	pre_append = pre + ".0." + "query_max_response_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["query-max-response-time"], _ = expandRouterMulticastInterfaceIgmpQueryMaxResponseTime(d, i["query_max_response_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "query_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["query-interval"], _ = expandRouterMulticastInterfaceIgmpQueryInterval(d, i["query_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "query_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["query-timeout"], _ = expandRouterMulticastInterfaceIgmpQueryTimeout(d, i["query_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "router_alert_check"
	if _, ok := d.GetOk(pre_append); ok {
		result["router-alert-check"], _ = expandRouterMulticastInterfaceIgmpRouterAlertCheck(d, i["router_alert_check"], pre_append, sv)
	}

	return result, nil
}

func expandRouterMulticastInterfaceIgmpAccessGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpImmediateLeaveGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpLastMemberQueryCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryMaxResponseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpQueryTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastInterfaceIgmpRouterAlertCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("route_threshold"); ok {
		if setArgNil {
			obj["route-threshold"] = nil
		} else {
			t, err := expandRouterMulticastRouteThreshold(d, v, "route_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["route-threshold"] = t
			}
		}
	} else if d.HasChange("route_threshold") {
		obj["route-threshold"] = nil
	}

	if v, ok := d.GetOk("route_limit"); ok {
		if setArgNil {
			obj["route-limit"] = nil
		} else {
			t, err := expandRouterMulticastRouteLimit(d, v, "route_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["route-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_routing"); ok {
		if setArgNil {
			obj["multicast-routing"] = nil
		} else {
			t, err := expandRouterMulticastMulticastRouting(d, v, "multicast_routing", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-routing"] = t
			}
		}
	}

	if v, ok := d.GetOk("pim_sm_global"); ok {
		t, err := expandRouterMulticastPimSmGlobal(d, v, "pim_sm_global", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-sm-global"] = t
		}
	}

	if v, ok := d.GetOk("pim_sm_global_vrf"); ok || d.HasChange("pim_sm_global_vrf") {
		if setArgNil {
			obj["pim-sm-global-vrf"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterMulticastPimSmGlobalVrf(d, v, "pim_sm_global_vrf", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pim-sm-global-vrf"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		if setArgNil {
			obj["interface"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterMulticastInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	return &obj, nil
}
