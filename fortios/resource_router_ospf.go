// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure OSPF.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterOspf() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterOspfUpdate,
		Read:   resourceRouterOspfRead,
		Update: resourceRouterOspfUpdate,
		Delete: resourceRouterOspfDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"abr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_cost_ref_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1000000),
				Optional:     true,
				Computed:     true,
			},
			"distance_external": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"distance_inter_area": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"distance_intra_area": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"database_overflow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_overflow_max_lsas": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"database_overflow_time_to_recover": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_metric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),
				Optional:     true,
				Computed:     true,
			},
			"default_information_metric_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_information_route_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"default_metric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),
				Optional:     true,
				Computed:     true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"rfc1583_compatible": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"spf_timers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distribute_list_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"distribute_route_map_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"restart_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restart_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"area": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shortcut": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"nssa_translator_role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"stub_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_default_information_originate_metric": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),
							Optional:     true,
							Computed:     true,
						},
						"nssa_default_information_originate_metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nssa_redistribution": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"range": &schema.Schema{
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
									"advertise": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"substitute": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"substitute_status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"virtual_link": &schema.Schema{
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
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"authentication_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 8),
										Optional:     true,
										Sensitive:    true,
									},
									"md5_key": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
										Computed:  true,
									},
									"md5_keychain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"dead_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"hello_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"retransmit_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"transmit_delay": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"peer": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"filter_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"list": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"direction": &schema.Schema{
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
			"ospf_interface": &schema.Schema{
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
						"interface": &schema.Schema{
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
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication_key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 8),
							Optional:     true,
							Sensitive:    true,
						},
						"md5_key": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"md5_keychain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"prefix_length": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"retransmit_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"transmit_delay": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"cost": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"dead_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_multiplier": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 10),
							Optional:     true,
							Computed:     true,
						},
						"database_filter_out": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mtu": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(576, 65535),
							Optional:     true,
							Computed:     true,
						},
						"mtu_ignore": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resync_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),
							Optional:     true,
							Computed:     true,
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
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"area": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
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
						"poll_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"cost": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
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
			"summary_address": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"advertise": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"access_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777214),
							Optional:     true,
							Computed:     true,
						},
						"routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterOspfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterOspf(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterOspf(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterOspf")
	}

	return resourceRouterOspfRead(d, m)
}

func resourceRouterOspfDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterOspf(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterOspf(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAbrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistanceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistanceInterArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistanceIntraArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflowMaxLsas(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflowTimeToRecover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRfc1583Compatible(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSpfTimers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeListIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeRouteMapIn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRestartMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRestartPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfArea(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfAreaId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := i["shortcut"]; ok {
			tmp["shortcut"] = flattenRouterOspfAreaShortcut(i["shortcut"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenRouterOspfAreaAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := i["default-cost"]; ok {
			tmp["default_cost"] = flattenRouterOspfAreaDefaultCost(i["default-cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = flattenRouterOspfAreaNssaTranslatorRole(i["nssa-translator-role"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := i["stub-type"]; ok {
			tmp["stub_type"] = flattenRouterOspfAreaStubType(i["stub-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenRouterOspfAreaType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := i["nssa-default-information-originate"]; ok {
			tmp["nssa_default_information_originate"] = flattenRouterOspfAreaNssaDefaultInformationOriginate(i["nssa-default-information-originate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := i["nssa-default-information-originate-metric"]; ok {
			tmp["nssa_default_information_originate_metric"] = flattenRouterOspfAreaNssaDefaultInformationOriginateMetric(i["nssa-default-information-originate-metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := i["nssa-default-information-originate-metric-type"]; ok {
			tmp["nssa_default_information_originate_metric_type"] = flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(i["nssa-default-information-originate-metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := i["nssa-redistribution"]; ok {
			tmp["nssa_redistribution"] = flattenRouterOspfAreaNssaRedistribution(i["nssa-redistribution"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := i["range"]; ok {
			tmp["range"] = flattenRouterOspfAreaRange(i["range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = flattenRouterOspfAreaVirtualLink(i["virtual-link"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := i["filter-list"]; ok {
			tmp["filter_list"] = flattenRouterOspfAreaFilterList(i["filter-list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfAreaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaShortcut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaStubType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfAreaRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterOspfAreaRangePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = flattenRouterOspfAreaRangeAdvertise(i["advertise"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := i["substitute"]; ok {
			tmp["substitute"] = flattenRouterOspfAreaRangeSubstitute(i["substitute"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := i["substitute-status"]; ok {
			tmp["substitute_status"] = flattenRouterOspfAreaRangeSubstituteStatus(i["substitute-status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfAreaRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspfAreaVirtualLinkName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenRouterOspfAreaVirtualLinkAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = flattenRouterOspfAreaVirtualLinkAuthenticationKey(i["authentication-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["authentication_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := i["md5-key"]; ok {
			tmp["md5_key"] = flattenRouterOspfAreaVirtualLinkMd5Key(i["md5-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["md5_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			tmp["md5_keychain"] = flattenRouterOspfAreaVirtualLinkMd5Keychain(i["md5-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = flattenRouterOspfAreaVirtualLinkDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterOspfAreaVirtualLinkHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = flattenRouterOspfAreaVirtualLinkRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = flattenRouterOspfAreaVirtualLinkTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := i["peer"]; ok {
			tmp["peer"] = flattenRouterOspfAreaVirtualLinkPeer(i["peer"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkMd5Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkMd5Keychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfAreaFilterListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := i["list"]; ok {
			tmp["list"] = flattenRouterOspfAreaFilterListList(i["list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = flattenRouterOspfAreaFilterListDirection(i["direction"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspfOspfInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterOspfOspfInterfaceInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenRouterOspfOspfInterfaceIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenRouterOspfOspfInterfaceAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := i["authentication-key"]; ok {
			tmp["authentication_key"] = flattenRouterOspfOspfInterfaceAuthenticationKey(i["authentication-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["authentication_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := i["md5-key"]; ok {
			tmp["md5_key"] = flattenRouterOspfOspfInterfaceMd5Key(i["md5-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["md5_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := i["md5-keychain"]; ok {
			tmp["md5_keychain"] = flattenRouterOspfOspfInterfaceMd5Keychain(i["md5-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := i["prefix-length"]; ok {
			tmp["prefix_length"] = flattenRouterOspfOspfInterfacePrefixLength(i["prefix-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = flattenRouterOspfOspfInterfaceRetransmitInterval(i["retransmit-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = flattenRouterOspfOspfInterfaceTransmitDelay(i["transmit-delay"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = flattenRouterOspfOspfInterfaceCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenRouterOspfOspfInterfacePriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = flattenRouterOspfOspfInterfaceDeadInterval(i["dead-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterOspfOspfInterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := i["hello-multiplier"]; ok {
			tmp["hello_multiplier"] = flattenRouterOspfOspfInterfaceHelloMultiplier(i["hello-multiplier"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
		if _, ok := i["database-filter-out"]; ok {
			tmp["database_filter_out"] = flattenRouterOspfOspfInterfaceDatabaseFilterOut(i["database-filter-out"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := i["mtu"]; ok {
			tmp["mtu"] = flattenRouterOspfOspfInterfaceMtu(i["mtu"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = flattenRouterOspfOspfInterfaceMtuIgnore(i["mtu-ignore"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := i["network-type"]; ok {
			tmp["network_type"] = flattenRouterOspfOspfInterfaceNetworkType(i["network-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := i["bfd"]; ok {
			tmp["bfd"] = flattenRouterOspfOspfInterfaceBfd(i["bfd"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterOspfOspfInterfaceStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
		if _, ok := i["resync-timeout"]; ok {
			tmp["resync_timeout"] = flattenRouterOspfOspfInterfaceResyncTimeout(i["resync-timeout"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceAuthenticationKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMd5Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMd5Keychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePrefixLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceDatabaseFilterOut(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceResyncTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterOspfNetworkPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := i["area"]; ok {
			tmp["area"] = flattenRouterOspfNetworkArea(i["area"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfNetworkArea(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfNeighborId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenRouterOspfNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := i["poll-interval"]; ok {
			tmp["poll_interval"] = flattenRouterOspfNeighborPollInterval(i["poll-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {
			tmp["cost"] = flattenRouterOspfNeighborCost(i["cost"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenRouterOspfNeighborPriority(i["priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborCost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfNeighborPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspfPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfSummaryAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterOspfSummaryAddressPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = flattenRouterOspfSummaryAddressTag(i["tag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := i["advertise"]; ok {
			tmp["advertise"] = flattenRouterOspfSummaryAddressAdvertise(i["advertise"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfSummaryAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfSummaryAddressTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspfDistributeListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = flattenRouterOspfDistributeListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = flattenRouterOspfDistributeListProtocol(i["protocol"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfDistributeListProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspfRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterOspfRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = flattenRouterOspfRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = flattenRouterOspfRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := i["metric-type"]; ok {
			tmp["metric_type"] = flattenRouterOspfRedistributeMetricType(i["metric-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := i["tag"]; ok {
			tmp["tag"] = flattenRouterOspfRedistributeTag(i["tag"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterOspfRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterOspfRedistributeTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("abr_type", flattenRouterOspfAbrType(o["abr-type"], d, "abr_type")); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", flattenRouterOspfAutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth")); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("distance_external", flattenRouterOspfDistanceExternal(o["distance-external"], d, "distance_external")); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_inter_area", flattenRouterOspfDistanceInterArea(o["distance-inter-area"], d, "distance_inter_area")); err != nil {
		if !fortiAPIPatch(o["distance-inter-area"]) {
			return fmt.Errorf("Error reading distance_inter_area: %v", err)
		}
	}

	if err = d.Set("distance_intra_area", flattenRouterOspfDistanceIntraArea(o["distance-intra-area"], d, "distance_intra_area")); err != nil {
		if !fortiAPIPatch(o["distance-intra-area"]) {
			return fmt.Errorf("Error reading distance_intra_area: %v", err)
		}
	}

	if err = d.Set("database_overflow", flattenRouterOspfDatabaseOverflow(o["database-overflow"], d, "database_overflow")); err != nil {
		if !fortiAPIPatch(o["database-overflow"]) {
			return fmt.Errorf("Error reading database_overflow: %v", err)
		}
	}

	if err = d.Set("database_overflow_max_lsas", flattenRouterOspfDatabaseOverflowMaxLsas(o["database-overflow-max-lsas"], d, "database_overflow_max_lsas")); err != nil {
		if !fortiAPIPatch(o["database-overflow-max-lsas"]) {
			return fmt.Errorf("Error reading database_overflow_max_lsas: %v", err)
		}
	}

	if err = d.Set("database_overflow_time_to_recover", flattenRouterOspfDatabaseOverflowTimeToRecover(o["database-overflow-time-to-recover"], d, "database_overflow_time_to_recover")); err != nil {
		if !fortiAPIPatch(o["database-overflow-time-to-recover"]) {
			return fmt.Errorf("Error reading database_overflow_time_to_recover: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspfDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspfDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric")); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspfDefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type")); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", flattenRouterOspfDefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map")); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspfDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterOspfDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("rfc1583_compatible", flattenRouterOspfRfc1583Compatible(o["rfc1583-compatible"], d, "rfc1583_compatible")); err != nil {
		if !fortiAPIPatch(o["rfc1583-compatible"]) {
			return fmt.Errorf("Error reading rfc1583_compatible: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspfRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", flattenRouterOspfSpfTimers(o["spf-timers"], d, "spf_timers")); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspfBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterOspfLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes")); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterOspfDistributeListIn(o["distribute-list-in"], d, "distribute_list_in")); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_route_map_in", flattenRouterOspfDistributeRouteMapIn(o["distribute-route-map-in"], d, "distribute_route_map_in")); err != nil {
		if !fortiAPIPatch(o["distribute-route-map-in"]) {
			return fmt.Errorf("Error reading distribute_route_map_in: %v", err)
		}
	}

	if err = d.Set("restart_mode", flattenRouterOspfRestartMode(o["restart-mode"], d, "restart_mode")); err != nil {
		if !fortiAPIPatch(o["restart-mode"]) {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_period", flattenRouterOspfRestartPeriod(o["restart-period"], d, "restart_period")); err != nil {
		if !fortiAPIPatch(o["restart-period"]) {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("area", flattenRouterOspfArea(o["area"], d, "area")); err != nil {
			if !fortiAPIPatch(o["area"]) {
				return fmt.Errorf("Error reading area: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("area"); ok {
			if err = d.Set("area", flattenRouterOspfArea(o["area"], d, "area")); err != nil {
				if !fortiAPIPatch(o["area"]) {
					return fmt.Errorf("Error reading area: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface")); err != nil {
			if !fortiAPIPatch(o["ospf-interface"]) {
				return fmt.Errorf("Error reading ospf_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ospf_interface"); ok {
			if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface")); err != nil {
				if !fortiAPIPatch(o["ospf-interface"]) {
					return fmt.Errorf("Error reading ospf_interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenRouterOspfNetwork(o["network"], d, "network")); err != nil {
			if !fortiAPIPatch(o["network"]) {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterOspfNetwork(o["network"], d, "network")); err != nil {
				if !fortiAPIPatch(o["network"]) {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterOspfNeighbor(o["neighbor"], d, "neighbor")); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterOspfNeighbor(o["neighbor"], d, "neighbor")); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("passive_interface", flattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
			if !fortiAPIPatch(o["passive-interface"]) {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("passive_interface"); ok {
			if err = d.Set("passive_interface", flattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
				if !fortiAPIPatch(o["passive-interface"]) {
					return fmt.Errorf("Error reading passive_interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
			if !fortiAPIPatch(o["summary-address"]) {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address")); err != nil {
				if !fortiAPIPatch(o["summary-address"]) {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
			if !fortiAPIPatch(o["distribute-list"]) {
				return fmt.Errorf("Error reading distribute_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distribute_list"); ok {
			if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
				if !fortiAPIPatch(o["distribute-list"]) {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterOspfRedistribute(o["redistribute"], d, "redistribute")); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterOspfRedistribute(o["redistribute"], d, "redistribute")); err != nil {
				if !fortiAPIPatch(o["redistribute"]) {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterOspfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterOspfAbrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAutoCostRefBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceInterArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceIntraArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflowMaxLsas(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflowTimeToRecover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationRouteMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRfc1583Compatible(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRouterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSpfTimers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeRouteMapIn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shortcut"], _ = expandRouterOspfAreaShortcut(d, i["shortcut"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspfAreaAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-cost"], _ = expandRouterOspfAreaDefaultCost(d, i["default_cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-translator-role"], _ = expandRouterOspfAreaNssaTranslatorRole(d, i["nssa_translator_role"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stub-type"], _ = expandRouterOspfAreaStubType(d, i["stub_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandRouterOspfAreaType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate"], _ = expandRouterOspfAreaNssaDefaultInformationOriginate(d, i["nssa_default_information_originate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric"], _ = expandRouterOspfAreaNssaDefaultInformationOriginateMetric(d, i["nssa_default_information_originate_metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric-type"], _ = expandRouterOspfAreaNssaDefaultInformationOriginateMetricType(d, i["nssa_default_information_originate_metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-redistribution"], _ = expandRouterOspfAreaNssaRedistribution(d, i["nssa_redistribution"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["range"], _ = expandRouterOspfAreaRange(d, i["range"], pre_append)
		} else {
			tmp["range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-link"], _ = expandRouterOspfAreaVirtualLink(d, i["virtual_link"], pre_append)
		} else {
			tmp["virtual-link"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list"], _ = expandRouterOspfAreaFilterList(d, i["filter_list"], pre_append)
		} else {
			tmp["filter-list"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaShortcut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaDefaultCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaStubType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterOspfAreaRangePrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspfAreaRangeAdvertise(d, i["advertise"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["substitute"], _ = expandRouterOspfAreaRangeSubstitute(d, i["substitute"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["substitute-status"], _ = expandRouterOspfAreaRangeSubstituteStatus(d, i["substitute_status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangePrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeSubstitute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeSubstituteStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfAreaVirtualLinkName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspfAreaVirtualLinkAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication-key"], _ = expandRouterOspfAreaVirtualLinkAuthenticationKey(d, i["authentication_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-key"], _ = expandRouterOspfAreaVirtualLinkMd5Key(d, i["md5_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-keychain"], _ = expandRouterOspfAreaVirtualLinkMd5Keychain(d, i["md5_keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspfAreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspfAreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspfAreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspfAreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peer"], _ = expandRouterOspfAreaVirtualLinkPeer(d, i["peer"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthenticationKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5Key(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5Keychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaFilterListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["list"], _ = expandRouterOspfAreaFilterListList(d, i["list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandRouterOspfAreaFilterListDirection(d, i["direction"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaFilterListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfOspfInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterOspfOspfInterfaceInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandRouterOspfOspfInterfaceIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspfOspfInterfaceAuthentication(d, i["authentication"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication-key"], _ = expandRouterOspfOspfInterfaceAuthenticationKey(d, i["authentication_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-key"], _ = expandRouterOspfOspfInterfaceMd5Key(d, i["md5_key"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-keychain"], _ = expandRouterOspfOspfInterfaceMd5Keychain(d, i["md5_keychain"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-length"], _ = expandRouterOspfOspfInterfacePrefixLength(d, i["prefix_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspfOspfInterfaceRetransmitInterval(d, i["retransmit_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspfOspfInterfaceTransmitDelay(d, i["transmit_delay"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspfOspfInterfaceCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspfOspfInterfacePriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspfOspfInterfaceDeadInterval(d, i["dead_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspfOspfInterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-multiplier"], _ = expandRouterOspfOspfInterfaceHelloMultiplier(d, i["hello_multiplier"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["database-filter-out"], _ = expandRouterOspfOspfInterfaceDatabaseFilterOut(d, i["database_filter_out"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu"], _ = expandRouterOspfOspfInterfaceMtu(d, i["mtu"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu-ignore"], _ = expandRouterOspfOspfInterfaceMtuIgnore(d, i["mtu_ignore"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["network-type"], _ = expandRouterOspfOspfInterfaceNetworkType(d, i["network_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterOspfOspfInterfaceBfd(d, i["bfd"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspfOspfInterfaceStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["resync-timeout"], _ = expandRouterOspfOspfInterfaceResyncTimeout(d, i["resync_timeout"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfOspfInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceAuthenticationKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Key(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Keychain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePrefixLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloMultiplier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDatabaseFilterOut(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceResyncTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterOspfNetworkPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["area"], _ = expandRouterOspfNetworkArea(d, i["area"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkArea(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfNeighborId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandRouterOspfNeighborIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poll-interval"], _ = expandRouterOspfNeighborPollInterval(d, i["poll_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspfNeighborCost(d, i["cost"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspfNeighborPriority(d, i["priority"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfNeighborId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborCost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfPassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfPassiveInterfaceName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfPassiveInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfSummaryAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterOspfSummaryAddressPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tag"], _ = expandRouterOspfSummaryAddressTag(d, i["tag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspfSummaryAddressAdvertise(d, i["advertise"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfSummaryAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressAdvertise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfDistributeListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-list"], _ = expandRouterOspfDistributeListAccessList(d, i["access_list"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandRouterOspfDistributeListProtocol(d, i["protocol"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfDistributeListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfRedistributeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspfRedistributeStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric"], _ = expandRouterOspfRedistributeMetric(d, i["metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["routemap"], _ = expandRouterOspfRedistributeRoutemap(d, i["routemap"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric-type"], _ = expandRouterOspfRedistributeMetricType(d, i["metric_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tag"], _ = expandRouterOspfRedistributeTag(d, i["tag"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfRedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeMetricType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("abr_type"); ok {
		t, err := expandRouterOspfAbrType(d, v, "abr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["abr-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_cost_ref_bandwidth"); ok {
		t, err := expandRouterOspfAutoCostRefBandwidth(d, v, "auto_cost_ref_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-cost-ref-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("distance_external"); ok {
		t, err := expandRouterOspfDistanceExternal(d, v, "distance_external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-external"] = t
		}
	}

	if v, ok := d.GetOk("distance_inter_area"); ok {
		t, err := expandRouterOspfDistanceInterArea(d, v, "distance_inter_area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-inter-area"] = t
		}
	}

	if v, ok := d.GetOk("distance_intra_area"); ok {
		t, err := expandRouterOspfDistanceIntraArea(d, v, "distance_intra_area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance-intra-area"] = t
		}
	}

	if v, ok := d.GetOk("database_overflow"); ok {
		t, err := expandRouterOspfDatabaseOverflow(d, v, "database_overflow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-overflow"] = t
		}
	}

	if v, ok := d.GetOk("database_overflow_max_lsas"); ok {
		t, err := expandRouterOspfDatabaseOverflowMaxLsas(d, v, "database_overflow_max_lsas")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-overflow-max-lsas"] = t
		}
	}

	if v, ok := d.GetOk("database_overflow_time_to_recover"); ok {
		t, err := expandRouterOspfDatabaseOverflowTimeToRecover(d, v, "database_overflow_time_to_recover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-overflow-time-to-recover"] = t
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok {
		t, err := expandRouterOspfDefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok {
		t, err := expandRouterOspfDefaultInformationMetric(d, v, "default_information_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric"] = t
		}
	}

	if v, ok := d.GetOk("default_information_metric_type"); ok {
		t, err := expandRouterOspfDefaultInformationMetricType(d, v, "default_information_metric_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-metric-type"] = t
		}
	}

	if v, ok := d.GetOk("default_information_route_map"); ok {
		t, err := expandRouterOspfDefaultInformationRouteMap(d, v, "default_information_route_map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-route-map"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok {
		t, err := expandRouterOspfDefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandRouterOspfDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("rfc1583_compatible"); ok {
		t, err := expandRouterOspfRfc1583Compatible(d, v, "rfc1583_compatible")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc1583-compatible"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		t, err := expandRouterOspfRouterId(d, v, "router_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router-id"] = t
		}
	}

	if v, ok := d.GetOk("spf_timers"); ok {
		t, err := expandRouterOspfSpfTimers(d, v, "spf_timers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spf-timers"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterOspfBfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		t, err := expandRouterOspfLogNeighbourChanges(d, v, "log_neighbour_changes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-neighbour-changes"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok {
		t, err := expandRouterOspfDistributeListIn(d, v, "distribute_list_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list-in"] = t
		}
	}

	if v, ok := d.GetOk("distribute_route_map_in"); ok {
		t, err := expandRouterOspfDistributeRouteMapIn(d, v, "distribute_route_map_in")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-route-map-in"] = t
		}
	}

	if v, ok := d.GetOk("restart_mode"); ok {
		t, err := expandRouterOspfRestartMode(d, v, "restart_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-mode"] = t
		}
	}

	if v, ok := d.GetOk("restart_period"); ok {
		t, err := expandRouterOspfRestartPeriod(d, v, "restart_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["restart-period"] = t
		}
	}

	if v, ok := d.GetOk("area"); ok {
		t, err := expandRouterOspfArea(d, v, "area")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["area"] = t
		}
	}

	if v, ok := d.GetOk("ospf_interface"); ok {
		t, err := expandRouterOspfOspfInterface(d, v, "ospf_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf-interface"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok {
		t, err := expandRouterOspfNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok {
		t, err := expandRouterOspfNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok {
		t, err := expandRouterOspfPassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("summary_address"); ok {
		t, err := expandRouterOspfSummaryAddress(d, v, "summary_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["summary-address"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list"); ok {
		t, err := expandRouterOspfDistributeList(d, v, "distribute_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok {
		t, err := expandRouterOspfRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	return &obj, nil
}
