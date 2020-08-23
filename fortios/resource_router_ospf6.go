// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 OSPF.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterOspf6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspf6Update,
		Read:   resourceRouterOspf6Read,
		Update: resourceRouterOspf6Update,
		Delete: resourceRouterOspf6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"abr_type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_cost_ref_bandwidth": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000000),
				Optional: true,
				Computed: true,
			},
			"default_information_originate": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_metric": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),
				Optional: true,
				Computed: true,
			},
			"default_information_metric_type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_route_map": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),
				Optional: true,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"spf_timers": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"area": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_cost": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777215),
							Optional: true,
							Computed: true,
						},
						"nssa_translator_role": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stub_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate_metric": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate_metric_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_redistribution": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"range": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
									"prefix6": &schema.Schema{
										Type: schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"advertise": &schema.Schema{
										Type: schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"virtual_link": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type: schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional: true,
										Computed: true,
									},
									"dead_interval": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional: true,
										Computed: true,
									},
									"hello_interval": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional: true,
										Computed: true,
									},
									"retransmit_interval": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional: true,
										Computed: true,
									},
									"transmit_delay": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional: true,
										Computed: true,
									},
									"peer": &schema.Schema{
										Type: schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"ospf6_interface": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"area_id": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional: true,
							Computed: true,
						},
						"retransmit_interval": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional: true,
							Computed: true,
						},
						"transmit_delay": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional: true,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional: true,
							Computed: true,
						},
						"dead_interval": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional: true,
							Computed: true,
						},
						"hello_interval": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mtu": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(576, 65535),
							Optional: true,
							Computed: true,
						},
						"mtu_ignore": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"neighbor": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip6": &schema.Schema{
										Type: schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"poll_interval": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional: true,
										Computed: true,
									},
									"cost": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),
							Optional: true,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"metric_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"summary_address": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"advertise": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}


func resourceRouterOspf6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterOspf6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterOspf6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterOspf6")
	}

	return resourceRouterOspf6Read(d, m)
}

func resourceRouterOspf6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterOspf6(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterOspf6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6 resource from API: %v", err)
	}
	return nil
}


func flattenRouterOspf6AbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6LogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6DefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Area(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspf6AreaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			tmp["default_cost"] = flattenRouterOspf6AreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = flattenRouterOspf6AreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			tmp["stub_type"] = flattenRouterOspf6AreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenRouterOspf6AreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {
			tmp["nssa_default_information_originate"] = flattenRouterOspf6AreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {
			tmp["nssa_default_information_originate_metric"] = flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {
			tmp["nssa_default_information_originate_metric_type"] = flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {
			tmp["nssa_redistribution"] = flattenRouterOspf6AreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = flattenRouterOspf6AreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = flattenRouterOspf6AreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspf6AreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterOspf6AreaRangePrefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = flattenRouterOspf6AreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6AreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = flattenRouterOspf6AreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterOspf6AreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = flattenRouterOspf6AreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = flattenRouterOspf6AreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = flattenRouterOspf6AreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6AreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6Ospf6InterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := i["area-id"]; ok {
			tmp["area_id"] = flattenRouterOspf6Ospf6InterfaceAreaId(i["area-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterOspf6Ospf6InterfaceInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = flattenRouterOspf6Ospf6InterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = flattenRouterOspf6Ospf6InterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = flattenRouterOspf6Ospf6InterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenRouterOspf6Ospf6InterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = flattenRouterOspf6Ospf6InterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterOspf6Ospf6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterOspf6Ospf6InterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			tmp["network_type"] = flattenRouterOspf6Ospf6InterfaceNetworkType(i["network-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = flattenRouterOspf6Ospf6InterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			tmp["mtu"] = flattenRouterOspf6Ospf6InterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = flattenRouterOspf6Ospf6InterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := i["neighbor"]; ok {
			tmp["neighbor"] = flattenRouterOspf6Ospf6InterfaceNeighbor(i["neighbor"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["ip6"] = flattenRouterOspf6Ospf6InterfaceNeighborIp6(i["ip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			tmp["poll_interval"] = flattenRouterOspf6Ospf6InterfaceNeighborPollInterval(i["poll-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = flattenRouterOspf6Ospf6InterfaceNeighborCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenRouterOspf6Ospf6InterfaceNeighborPriority(i["priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6Ospf6InterfaceNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6Redistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6RedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterOspf6RedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = flattenRouterOspf6RedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = flattenRouterOspf6RedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = flattenRouterOspf6RedistributeMetricType(i["metric-type"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6RedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6PassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6PassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6PassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspf6SummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterOspf6SummaryAddressPrefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = flattenRouterOspf6SummaryAddressAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = flattenRouterOspf6SummaryAddressTag(i["tag"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspf6SummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectRouterOspf6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("abr_type", flattenRouterOspf6AbrType(o["abr-type"], d, "abr_type")); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", flattenRouterOspf6AutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspf6DefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterOspf6LogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspf6DefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspf6DefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", flattenRouterOspf6DefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspf6DefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspf6RouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", flattenRouterOspf6SpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspf6Bfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("area", flattenRouterOspf6Area(o["area"], d, "area")); err != nil {
            if !fortiAPIPatch(o["area"]) {
                return fmt.Errorf("Error reading area: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("area"); ok {
            if err = d.Set("area", flattenRouterOspf6Area(o["area"], d, "area")); err != nil {
                if !fortiAPIPatch(o["area"]) {
                    return fmt.Errorf("Error reading area: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface")); err != nil {
            if !fortiAPIPatch(o["ospf6-interface"]) {
                return fmt.Errorf("Error reading ospf6_interface: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("ospf6_interface"); ok {
            if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface")); err != nil {
                if !fortiAPIPatch(o["ospf6-interface"]) {
                    return fmt.Errorf("Error reading ospf6_interface: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute")); err != nil {
            if !fortiAPIPatch(o["redistribute"]) {
                return fmt.Errorf("Error reading redistribute: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("redistribute"); ok {
            if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute")); err != nil {
                if !fortiAPIPatch(o["redistribute"]) {
                    return fmt.Errorf("Error reading redistribute: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("passive_interface", flattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
            if !fortiAPIPatch(o["passive-interface"]) {
                return fmt.Errorf("Error reading passive_interface: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("passive_interface"); ok {
            if err = d.Set("passive_interface", flattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
                if !fortiAPIPatch(o["passive-interface"]) {
                    return fmt.Errorf("Error reading passive_interface: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
            if !fortiAPIPatch(o["summary-address"]) {
                return fmt.Errorf("Error reading summary_address: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("summary_address"); ok {
            if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
                if !fortiAPIPatch(o["summary-address"]) {
                    return fmt.Errorf("Error reading summary_address: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenRouterOspf6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandRouterOspf6AbrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AutoCostRefBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6LogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RouterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SpfTimers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Bfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Area(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterOspf6AreaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-cost"], _ = expandRouterOspf6AreaDefaultCost(d, i["default_cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-translator-role"], _ = expandRouterOspf6AreaNssaTranslatorRole(d, i["nssa_translator_role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stub-type"], _ = expandRouterOspf6AreaStubType(d, i["stub_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandRouterOspf6AreaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginate(d, i["nssa_default_information_originate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginateMetric(d, i["nssa_default_information_originate_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric-type"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType(d, i["nssa_default_information_originate_metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-redistribution"], _ = expandRouterOspf6AreaNssaRedistribution(d, i["nssa_redistribution"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["range"], _ = expandRouterOspf6AreaRange(d, i["range"], pre_append)
		} else {
			tmp["range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-link"], _ = expandRouterOspf6AreaVirtualLink(d, i["virtual_link"], pre_append)
		} else {
			tmp["virtual-link"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaDefaultCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaStubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterOspf6AreaRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterOspf6AreaRangePrefix6(d, i["prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspf6AreaRangeAdvertise(d, i["advertise"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterOspf6AreaVirtualLinkName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspf6AreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspf6AreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspf6AreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspf6AreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peer"], _ = expandRouterOspf6AreaVirtualLinkPeer(d, i["peer"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterOspf6Ospf6InterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["area-id"], _ = expandRouterOspf6Ospf6InterfaceAreaId(d, i["area_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterOspf6Ospf6InterfaceInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspf6Ospf6InterfaceRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspf6Ospf6InterfaceTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfacePriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspf6Ospf6InterfaceDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspf6Ospf6InterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspf6Ospf6InterfaceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["network-type"], _ = expandRouterOspf6Ospf6InterfaceNetworkType(d, i["network_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterOspf6Ospf6InterfaceBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu"], _ = expandRouterOspf6Ospf6InterfaceMtu(d, i["mtu"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu-ignore"], _ = expandRouterOspf6Ospf6InterfaceMtuIgnore(d, i["mtu_ignore"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["neighbor"], _ = expandRouterOspf6Ospf6InterfaceNeighbor(d, i["neighbor"], pre_append)
		} else {
			tmp["neighbor"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceAreaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandRouterOspf6Ospf6InterfaceNeighborIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poll-interval"], _ = expandRouterOspf6Ospf6InterfaceNeighborPollInterval(d, i["poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceNeighborCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfaceNeighborPriority(d, i["priority"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Redistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterOspf6RedistributeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspf6RedistributeStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric"], _ = expandRouterOspf6RedistributeMetric(d, i["metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["routemap"], _ = expandRouterOspf6RedistributeRoutemap(d, i["routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric-type"], _ = expandRouterOspf6RedistributeMetricType(d, i["metric_type"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6RedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6PassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterOspf6PassiveInterfaceName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6PassiveInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterOspf6SummaryAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterOspf6SummaryAddressPrefix6(d, i["prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspf6SummaryAddressAdvertise(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tag"], _ = expandRouterOspf6SummaryAddressTag(d, i["tag"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6SummaryAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressPrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectRouterOspf6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("abr_type"); ok {
		t, err := expandRouterOspf6AbrType(d, v, "abr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["abr-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_cost_ref_bandwidth"); ok {
		t, err := expandRouterOspf6AutoCostRefBandwidth(d, v, "auto_cost_ref_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-cost-ref-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok {
		t, err := expandRouterOspf6DefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		t, err := expandRouterOspf6LogNeighbourChanges(d, v, "log_neighbour_changes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-neighbour-changes"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok {
		t, err := expandRouterOspf6DefaultInformationMetric(d, v, "default_information_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric_type"); ok {
		t, err := expandRouterOspf6DefaultInformationMetricType(d, v, "default_information_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("default_information_route_map"); ok {
		t, err := expandRouterOspf6DefaultInformationRouteMap(d, v, "default_information_route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-route-map"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok {
		t, err := expandRouterOspf6DefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		t, err := expandRouterOspf6RouterId(d, v, "router_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("spf_timers"); ok {
		t, err := expandRouterOspf6SpfTimers(d, v, "spf_timers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spf-timers"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterOspf6Bfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("area"); ok {
		t, err := expandRouterOspf6Area(d, v, "area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_interface"); ok {
		t, err := expandRouterOspf6Ospf6Interface(d, v, "ospf6_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6-interface"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok {
		t, err := expandRouterOspf6Redistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok {
		t, err := expandRouterOspf6PassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("summary_address"); ok {
		t, err := expandRouterOspf6SummaryAddress(d, v, "summary_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-address"] = t
		}
	}


	return &obj, nil
}

