// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 OSPF.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterOspf6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterOspf6Read,
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
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
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
			"restart_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"restart_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"restart_on_topology_change": &schema.Schema{
				Type:     schema.TypeString,
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
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_rollover_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ipsec_auth_alg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipsec_enc_alg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipsec_keys": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spi": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"auth_key": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"enc_key": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
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
									"prefix6": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"advertise": &schema.Schema{
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
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"key_rollover_interval": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ipsec_auth_alg": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipsec_enc_alg": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipsec_keys": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"spi": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
												"auth_key": &schema.Schema{
													Type:      schema.TypeString,
													Sensitive: true,
													Computed:  true,
												},
												"enc_key": &schema.Schema{
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
					},
				},
			},
			"ospf6_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"area_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
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
						"status": &schema.Schema{
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
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mtu_ignore": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_rollover_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ipsec_auth_alg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipsec_enc_alg": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipsec_keys": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spi": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"auth_key": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"enc_key": &schema.Schema{
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
						},
						"neighbor": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip6": &schema.Schema{
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
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"advertise": &schema.Schema{
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

func dataSourceRouterOspf6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterOspf6"

	o, err := c.ReadRouterOspf6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterOspf6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterOspf6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterOspf6AbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6LogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6DefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RestartPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RestartOnTopologyChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Area(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspf6AreaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			tmp["default_cost"] = dataSourceFlattenRouterOspf6AreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = dataSourceFlattenRouterOspf6AreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			tmp["stub_type"] = dataSourceFlattenRouterOspf6AreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenRouterOspf6AreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {
			tmp["nssa_default_information_originate"] = dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {
			tmp["nssa_default_information_originate_metric"] = dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {
			tmp["nssa_default_information_originate_metric_type"] = dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {
			tmp["nssa_redistribution"] = dataSourceFlattenRouterOspf6AreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspf6AreaAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			tmp["key_rollover_interval"] = dataSourceFlattenRouterOspf6AreaKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			tmp["ipsec_auth_alg"] = dataSourceFlattenRouterOspf6AreaIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			tmp["ipsec_enc_alg"] = dataSourceFlattenRouterOspf6AreaIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			tmp["ipsec_keys"] = dataSourceFlattenRouterOspf6AreaIpsecKeys(i["ipsec-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = dataSourceFlattenRouterOspf6AreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = dataSourceFlattenRouterOspf6AreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			tmp["spi"] = dataSourceFlattenRouterOspf6AreaIpsecKeysSpi(i["spi"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecKeysAuthKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaIpsecKeysEncKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspf6AreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterOspf6AreaRangePrefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = dataSourceFlattenRouterOspf6AreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaRangePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspf6AreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspf6AreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = dataSourceFlattenRouterOspf6AreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspf6AreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			tmp["key_rollover_interval"] = dataSourceFlattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			tmp["ipsec_auth_alg"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			tmp["ipsec_enc_alg"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			tmp["ipsec_keys"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeys(i["ipsec-keys"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			tmp["spi"] = dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(i["spi"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeysAuthKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6AreaVirtualLinkIpsecKeysEncKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspf6Ospf6InterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := i["area-id"]; ok {
			tmp["area_id"] = dataSourceFlattenRouterOspf6Ospf6InterfaceAreaId(i["area-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterOspf6Ospf6InterfaceInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = dataSourceFlattenRouterOspf6Ospf6InterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = dataSourceFlattenRouterOspf6Ospf6InterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenRouterOspf6Ospf6InterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterOspf6Ospf6InterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			tmp["network_type"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNetworkType(i["network-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = dataSourceFlattenRouterOspf6Ospf6InterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			tmp["mtu"] = dataSourceFlattenRouterOspf6Ospf6InterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = dataSourceFlattenRouterOspf6Ospf6InterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenRouterOspf6Ospf6InterfaceAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := i["key-rollover-interval"]; ok {
			tmp["key_rollover_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(i["key-rollover-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := i["ipsec-auth-alg"]; ok {
			tmp["ipsec_auth_alg"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(i["ipsec-auth-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := i["ipsec-enc-alg"]; ok {
			tmp["ipsec_enc_alg"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecEncAlg(i["ipsec-enc-alg"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := i["ipsec-keys"]; ok {
			tmp["ipsec_keys"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeys(i["ipsec-keys"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := i["neighbor"]; ok {
			tmp["neighbor"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighbor(i["neighbor"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeys(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := i["spi"]; ok {
			tmp["spi"] = dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(i["spi"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeysAuthKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceIpsecKeysEncKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			tmp["ip6"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborIp6(i["ip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			tmp["poll_interval"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPollInterval(i["poll-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPriority(i["priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Ospf6InterfaceNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6Redistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspf6RedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterOspf6RedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterOspf6RedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterOspf6RedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = dataSourceFlattenRouterOspf6RedistributeMetricType(i["metric-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6RedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6RedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6PassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterOspf6PassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6PassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterOspf6SummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = dataSourceFlattenRouterOspf6SummaryAddressPrefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = dataSourceFlattenRouterOspf6SummaryAddressAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = dataSourceFlattenRouterOspf6SummaryAddressTag(i["tag"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterOspf6SummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterOspf6SummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterOspf6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("abr_type", dataSourceFlattenRouterOspf6AbrType(o["abr-type"], d, "abr_type")); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", dataSourceFlattenRouterOspf6AutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("default_information_originate", dataSourceFlattenRouterOspf6DefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", dataSourceFlattenRouterOspf6LogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("default_information_metric", dataSourceFlattenRouterOspf6DefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", dataSourceFlattenRouterOspf6DefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", dataSourceFlattenRouterOspf6DefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterOspf6DefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenRouterOspf6RouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", dataSourceFlattenRouterOspf6SpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterOspf6Bfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("restart_mode", dataSourceFlattenRouterOspf6RestartMode(o["restart-mode"], d, "restart_mode")); err != nil {
		if !fortiAPIPatch(o["restart-mode"]) {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_period", dataSourceFlattenRouterOspf6RestartPeriod(o["restart-period"], d, "restart_period")); err != nil {
		if !fortiAPIPatch(o["restart-period"]) {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if err = d.Set("restart_on_topology_change", dataSourceFlattenRouterOspf6RestartOnTopologyChange(o["restart-on-topology-change"], d, "restart_on_topology_change")); err != nil {
		if !fortiAPIPatch(o["restart-on-topology-change"]) {
			return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
		}
	}

	if err = d.Set("area", dataSourceFlattenRouterOspf6Area(o["area"], d, "area")); err != nil {
		if !fortiAPIPatch(o["area"]) {
			return fmt.Errorf("Error reading area: %v", err)
		}
	}

	if err = d.Set("ospf6_interface", dataSourceFlattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface")); err != nil {
		if !fortiAPIPatch(o["ospf6-interface"]) {
			return fmt.Errorf("Error reading ospf6_interface: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("summary_address", dataSourceFlattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
		if !fortiAPIPatch(o["summary-address"]) {
			return fmt.Errorf("Error reading summary_address: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterOspf6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
