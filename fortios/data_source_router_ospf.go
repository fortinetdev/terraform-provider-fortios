// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure OSPF.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspfRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"abr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_cost_ref_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_external": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_inter_area": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance_intra_area": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"database_overflow": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_overflow_max_lsas": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"database_overflow_time_to_recover": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_information_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_information_route_map": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"rfc1583_compatible": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"spf_timers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_list_in": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distribute_route_map_in": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"restart_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"restart_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"area": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"shortcut": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"default_cost": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nssa_translator_role": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"stub_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nssa_default_information_originate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nssa_default_information_originate_metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"nssa_default_information_originate_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nssa_redistribution": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"range": &schema.Schema{
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
									"advertise": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"substitute": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"substitute_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"virtual_link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"authentication_key": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"keychain": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"md5_key": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"md5_keychain": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dead_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"hello_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"retransmit_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"transmit_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"peer": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"md5_keys": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"key_string": &schema.Schema{
													Type:      schema.TypeString,
													Sensitive: true,
													Computed:  true,
												},
											},
										},
									},
								},
							},
						},
						"filter_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"list": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"ospf_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication_key": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"keychain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"md5_key": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"md5_keychain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_length": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"retransmit_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"transmit_delay": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dead_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_multiplier": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"database_filter_out": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mtu_ignore": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resync_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"md5_keys": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"key_string": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
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
						"area": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"comments": &schema.Schema{
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
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"poll_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
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
			"summary_address": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"advertise": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
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
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterOspf"

	o, err := c.ReadRouterOspf(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterOspf(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterOspfAbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistanceInterArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistanceIntraArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDatabaseOverflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDatabaseOverflowMaxLsas(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDatabaseOverflowTimeToRecover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRfc1583Compatible(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfSpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRestartPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfArea(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfAreaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			tmp["shortcut"] = dataSourceFlattenRouterOspfAreaShortcut(i["shortcut"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfAreaAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			tmp["default_cost"] = dataSourceFlattenRouterOspfAreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = dataSourceFlattenRouterOspfAreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			tmp["stub_type"] = dataSourceFlattenRouterOspfAreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenRouterOspfAreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {
			tmp["nssa_default_information_originate"] = dataSourceFlattenRouterOspfAreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {
			tmp["nssa_default_information_originate_metric"] = dataSourceFlattenRouterOspfAreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {
			tmp["nssa_default_information_originate_metric_type"] = dataSourceFlattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {
			tmp["nssa_redistribution"] = dataSourceFlattenRouterOspfAreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			tmp["comments"] = dataSourceFlattenRouterOspfAreaComments(i["comments"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = dataSourceFlattenRouterOspfAreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = dataSourceFlattenRouterOspfAreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := i["filter-list"]; ok {
			tmp["filter_list"] = dataSourceFlattenRouterOspfAreaFilterList(i["filter-list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfAreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterOspfAreaRangePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = dataSourceFlattenRouterOspfAreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {
			tmp["substitute"] = dataSourceFlattenRouterOspfAreaRangeSubstitute(i["substitute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {
			tmp["substitute_status"] = dataSourceFlattenRouterOspfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspfAreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = dataSourceFlattenRouterOspfAreaVirtualLinkAuthenticationKey(i["authentication-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["authentication_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {
			tmp["keychain"] = dataSourceFlattenRouterOspfAreaVirtualLinkKeychain(i["keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := i["md5-key"]; ok {
			tmp["md5_key"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5Key(i["md5-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["md5_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			tmp["md5_keychain"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5Keychain(i["md5-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = dataSourceFlattenRouterOspfAreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			tmp["md5_keys"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5Keys(i["md5-keys"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5Keychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := i["key-string"]; ok {
			tmp["key_string"] = dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysKeyString(i["key-string"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key_string"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaVirtualLinkMd5KeysKeyString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfAreaFilterListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			tmp["list"] = dataSourceFlattenRouterOspfAreaFilterListList(i["list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterOspfAreaFilterListDirection(i["direction"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspfOspfInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			tmp["comments"] = dataSourceFlattenRouterOspfOspfInterfaceComments(i["comments"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterOspfOspfInterfaceInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterOspfOspfInterfaceIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspfOspfInterfaceAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = dataSourceFlattenRouterOspfOspfInterfaceAuthenticationKey(i["authentication-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["authentication_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := i["keychain"]; ok {
			tmp["keychain"] = dataSourceFlattenRouterOspfOspfInterfaceKeychain(i["keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := i["md5-key"]; ok {
			tmp["md5_key"] = dataSourceFlattenRouterOspfOspfInterfaceMd5Key(i["md5-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["md5_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			tmp["md5_keychain"] = dataSourceFlattenRouterOspfOspfInterfaceMd5Keychain(i["md5-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := i["prefix-length"]; ok {
			tmp["prefix_length"] = dataSourceFlattenRouterOspfOspfInterfacePrefixLength(i["prefix-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspfOspfInterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspfOspfInterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = dataSourceFlattenRouterOspfOspfInterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenRouterOspfOspfInterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspfOspfInterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspfOspfInterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := i["hello-multiplier"]; ok {
			tmp["hello_multiplier"] = dataSourceFlattenRouterOspfOspfInterfaceHelloMultiplier(i["hello-multiplier"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
		if _, ok := i["database-filter-out"]; ok {
			tmp["database_filter_out"] = dataSourceFlattenRouterOspfOspfInterfaceDatabaseFilterOut(i["database-filter-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			tmp["mtu"] = dataSourceFlattenRouterOspfOspfInterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = dataSourceFlattenRouterOspfOspfInterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			tmp["network_type"] = dataSourceFlattenRouterOspfOspfInterfaceNetworkType(i["network-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterOspfOspfInterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterOspfOspfInterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
		if _, ok := i["resync-timeout"]; ok {
			tmp["resync_timeout"] = dataSourceFlattenRouterOspfOspfInterfaceResyncTimeout(i["resync-timeout"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := i["md5-keys"]; ok {
			tmp["md5_keys"] = dataSourceFlattenRouterOspfOspfInterfaceMd5Keys(i["md5-keys"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceMd5Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceMd5Keychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfacePrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceDatabaseFilterOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceResyncTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfOspfInterfaceMd5KeysId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := i["key-string"]; ok {
			tmp["key_string"] = dataSourceFlattenRouterOspfOspfInterfaceMd5KeysKeyString(i["key-string"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key_string"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfOspfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfOspfInterfaceMd5KeysKeyString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterOspfNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := i["area"]; ok {
			tmp["area"] = dataSourceFlattenRouterOspfNetworkArea(i["area"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			tmp["comments"] = dataSourceFlattenRouterOspfNetworkComments(i["comments"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfNetworkArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNetworkComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfNeighborId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterOspfNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			tmp["poll_interval"] = dataSourceFlattenRouterOspfNeighborPollInterval(i["poll-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = dataSourceFlattenRouterOspfNeighborCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenRouterOspfNeighborPriority(i["priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspfPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfSummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterOspfSummaryAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspfSummaryAddressTag(i["tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = dataSourceFlattenRouterOspfSummaryAddressAdvertise(i["advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterOspfSummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfSummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspfDistributeListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterOspfDistributeListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = dataSourceFlattenRouterOspfDistributeListProtocol(i["protocol"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfDistributeListProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspfRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterOspfRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterOspfRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterOspfRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterOspfRedistributeMetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspfRedistributeTag(i["tag"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspfRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspfRedistributeTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("abr_type", dataSourceFlattenRouterOspfAbrType(o["abr-type"], d, "abr_type")); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", dataSourceFlattenRouterOspfAutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("distance_external", dataSourceFlattenRouterOspfDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_inter_area", dataSourceFlattenRouterOspfDistanceInterArea(o["distance-inter-area"], d, "distance_inter_area")); err != nil {
		if !fortiAPIPatch(o["distance-inter-area"]) {
			return fmt.Errorf("Error reading distance_inter_area: %v", err)
		}
	}

	if err = d.Set("distance_intra_area", dataSourceFlattenRouterOspfDistanceIntraArea(o["distance-intra-area"], d, "distance_intra_area")); err != nil {
		if !fortiAPIPatch(o["distance-intra-area"]) {
			return fmt.Errorf("Error reading distance_intra_area: %v", err)
		}
	}

	if err = d.Set("database_overflow", dataSourceFlattenRouterOspfDatabaseOverflow(o["database-overflow"], d, "database_overflow")); err != nil {
		if !fortiAPIPatch(o["database-overflow"]) {
			return fmt.Errorf("Error reading database_overflow: %v", err)
		}
	}

	if err = d.Set("database_overflow_max_lsas", dataSourceFlattenRouterOspfDatabaseOverflowMaxLsas(o["database-overflow-max-lsas"], d, "database_overflow_max_lsas")); err != nil {
		if !fortiAPIPatch(o["database-overflow-max-lsas"]) {
			return fmt.Errorf("Error reading database_overflow_max_lsas: %v", err)
		}
	}

	if err = d.Set("database_overflow_time_to_recover", dataSourceFlattenRouterOspfDatabaseOverflowTimeToRecover(o["database-overflow-time-to-recover"], d, "database_overflow_time_to_recover")); err != nil {
		if !fortiAPIPatch(o["database-overflow-time-to-recover"]) {
			return fmt.Errorf("Error reading database_overflow_time_to_recover: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterOspfDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_information_metric", dataSourceFlattenRouterOspfDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", dataSourceFlattenRouterOspfDefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", dataSourceFlattenRouterOspfDefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterOspfDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterOspfDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("rfc1583_compatible", dataSourceFlattenRouterOspfRfc1583Compatible(o["rfc1583-compatible"], d, "rfc1583_compatible")); err != nil {
		if !fortiAPIPatch(o["rfc1583-compatible"]) {
			return fmt.Errorf("Error reading rfc1583_compatible: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterOspfRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", dataSourceFlattenRouterOspfSpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterOspfBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterOspfLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", dataSourceFlattenRouterOspfDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_route_map_in", dataSourceFlattenRouterOspfDistributeRouteMapIn(o["distribute-route-map-in"], d, "distribute_route_map_in")); err != nil {
		if !fortiAPIPatch(o["distribute-route-map-in"]) {
			return fmt.Errorf("Error reading distribute_route_map_in: %v", err)
		}
	}

	if err = d.Set("restart_mode", dataSourceFlattenRouterOspfRestartMode(o["restart-mode"], d, "restart_mode")); err != nil {
		if !fortiAPIPatch(o["restart-mode"]) {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_period", dataSourceFlattenRouterOspfRestartPeriod(o["restart-period"], d, "restart_period")); err != nil {
		if !fortiAPIPatch(o["restart-period"]) {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if err = d.Set("area", dataSourceFlattenRouterOspfArea(o["area"], d, "area")); err != nil {
		if !fortiAPIPatch(o["area"]) {
			return fmt.Errorf("Error reading area: %v", err)
		}
	}

	if err = d.Set("ospf_interface", dataSourceFlattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface")); err != nil {
		if !fortiAPIPatch(o["ospf-interface"]) {
			return fmt.Errorf("Error reading ospf_interface: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterOspfNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterOspfNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("summary_address", dataSourceFlattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
		if !fortiAPIPatch(o["summary-address"]) {
			return fmt.Errorf("Error reading summary_address: %v", err)
		}
	}

	if err = d.Set("distribute_list", dataSourceFlattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
		if !fortiAPIPatch(o["distribute-list"]) {
			return fmt.Errorf("Error reading distribute_list: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterOspfRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterOspfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
