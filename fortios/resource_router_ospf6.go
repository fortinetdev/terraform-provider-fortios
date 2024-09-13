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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
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
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_neighbour_changes": &schema.Schema{
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
			},
			"default_metric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777214),
				Optional:     true,
				Computed:     true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"spf_timers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"restart_on_topology_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
						"default_cost": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16777215),
							Optional:     true,
							Computed:     true,
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
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_rollover_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 216000),
							Optional:     true,
							Computed:     true,
						},
						"ipsec_auth_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipsec_enc_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipsec_keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spi": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"auth_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
									"enc_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
								},
							},
						},
						"range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"prefix6": &schema.Schema{
										Type:     schema.TypeString,
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
						"virtual_link": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
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
									"authentication": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key_rollover_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(300, 216000),
										Optional:     true,
										Computed:     true,
									},
									"ipsec_auth_alg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipsec_enc_alg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipsec_keys": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"spi": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"auth_key": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 128),
													Optional:     true,
													Sensitive:    true,
												},
												"enc_key": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 128),
													Optional:     true,
													Sensitive:    true,
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
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"area_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
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
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"dead_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
						"hello_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
						"status": &schema.Schema{
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
						"mtu": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(576, 65535),
							Optional:     true,
						},
						"mtu_ignore": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_rollover_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 216000),
							Optional:     true,
							Computed:     true,
						},
						"ipsec_auth_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipsec_enc_alg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipsec_keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spi": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"auth_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
									"enc_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
								},
							},
						},
						"neighbor": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip6": &schema.Schema{
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
						},
						"routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"metric_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
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
						},
						"prefix6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"advertise": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
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

func resourceRouterOspf6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterOspf6(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterOspf6(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterOspf6(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterOspf6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterOspf6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspf6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspf6AbrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6DefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6LogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6DefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6DefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6DefaultMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6RouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6SpfTimers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Bfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6RestartMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6RestartPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6RestartOnTopologyChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Area(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspf6AreaId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if cur_v, ok := i["default-cost"]; ok {
			tmp["default_cost"] = flattenRouterOspf6AreaDefaultCost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if cur_v, ok := i["nssa-translator-role"]; ok {
			tmp["nssa_translator_role"] = flattenRouterOspf6AreaNssaTranslatorRole(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if cur_v, ok := i["stub-type"]; ok {
			tmp["stub_type"] = flattenRouterOspf6AreaStubType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenRouterOspf6AreaType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if cur_v, ok := i["nssa-default-information-originate"]; ok {
			tmp["nssa_default_information_originate"] = flattenRouterOspf6AreaNssaDefaultInformationOriginate(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if cur_v, ok := i["nssa-default-information-originate-metric"]; ok {
			tmp["nssa_default_information_originate_metric"] = flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if cur_v, ok := i["nssa-default-information-originate-metric-type"]; ok {
			tmp["nssa_default_information_originate_metric_type"] = flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if cur_v, ok := i["nssa-redistribution"]; ok {
			tmp["nssa_redistribution"] = flattenRouterOspf6AreaNssaRedistribution(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if cur_v, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenRouterOspf6AreaAuthentication(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if cur_v, ok := i["key-rollover-interval"]; ok {
			tmp["key_rollover_interval"] = flattenRouterOspf6AreaKeyRolloverInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if cur_v, ok := i["ipsec-auth-alg"]; ok {
			tmp["ipsec_auth_alg"] = flattenRouterOspf6AreaIpsecAuthAlg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if cur_v, ok := i["ipsec-enc-alg"]; ok {
			tmp["ipsec_enc_alg"] = flattenRouterOspf6AreaIpsecEncAlg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if cur_v, ok := i["ipsec-keys"]; ok {
			tmp["ipsec_keys"] = flattenRouterOspf6AreaIpsecKeys(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if cur_v, ok := i["range"]; ok {
			tmp["range"] = flattenRouterOspf6AreaRange(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if cur_v, ok := i["virtual-link"]; ok {
			tmp["virtual_link"] = flattenRouterOspf6AreaVirtualLink(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspf6AreaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaDefaultCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaStubType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaIpsecKeys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if cur_v, ok := i["spi"]; ok {
			tmp["spi"] = flattenRouterOspf6AreaIpsecKeysSpi(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := i["auth-key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := i["enc-key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["enc_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "spi", d)
	return result
}

func flattenRouterOspf6AreaIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspf6AreaRangeId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if cur_v, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterOspf6AreaRangePrefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if cur_v, ok := i["advertise"]; ok {
			tmp["advertise"] = flattenRouterOspf6AreaRangeAdvertise(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspf6AreaRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaRangePrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLink(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6AreaVirtualLinkName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if cur_v, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = flattenRouterOspf6AreaVirtualLinkDeadInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if cur_v, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterOspf6AreaVirtualLinkHelloInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if cur_v, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = flattenRouterOspf6AreaVirtualLinkRetransmitInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if cur_v, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = flattenRouterOspf6AreaVirtualLinkTransmitDelay(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if cur_v, ok := i["peer"]; ok {
			tmp["peer"] = flattenRouterOspf6AreaVirtualLinkPeer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if cur_v, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenRouterOspf6AreaVirtualLinkAuthentication(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if cur_v, ok := i["key-rollover-interval"]; ok {
			tmp["key_rollover_interval"] = flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if cur_v, ok := i["ipsec-auth-alg"]; ok {
			tmp["ipsec_auth_alg"] = flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if cur_v, ok := i["ipsec-enc-alg"]; ok {
			tmp["ipsec_enc_alg"] = flattenRouterOspf6AreaVirtualLinkIpsecEncAlg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if cur_v, ok := i["ipsec-keys"]; ok {
			tmp["ipsec_keys"] = flattenRouterOspf6AreaVirtualLinkIpsecKeys(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspf6AreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6AreaVirtualLinkIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if cur_v, ok := i["spi"]; ok {
			tmp["spi"] = flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := i["auth-key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := i["enc-key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["enc_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "spi", d)
	return result
}

func flattenRouterOspf6AreaVirtualLinkIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6Interface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6Ospf6InterfaceName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if cur_v, ok := i["area-id"]; ok {
			tmp["area_id"] = flattenRouterOspf6Ospf6InterfaceAreaId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterOspf6Ospf6InterfaceInterface(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if cur_v, ok := i["retransmit-interval"]; ok {
			tmp["retransmit_interval"] = flattenRouterOspf6Ospf6InterfaceRetransmitInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if cur_v, ok := i["transmit-delay"]; ok {
			tmp["transmit_delay"] = flattenRouterOspf6Ospf6InterfaceTransmitDelay(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if cur_v, ok := i["cost"]; ok {
			tmp["cost"] = flattenRouterOspf6Ospf6InterfaceCost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenRouterOspf6Ospf6InterfacePriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if cur_v, ok := i["dead-interval"]; ok {
			tmp["dead_interval"] = flattenRouterOspf6Ospf6InterfaceDeadInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if cur_v, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterOspf6Ospf6InterfaceHelloInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenRouterOspf6Ospf6InterfaceStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if cur_v, ok := i["network-type"]; ok {
			tmp["network_type"] = flattenRouterOspf6Ospf6InterfaceNetworkType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if cur_v, ok := i["bfd"]; ok {
			tmp["bfd"] = flattenRouterOspf6Ospf6InterfaceBfd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if cur_v, ok := i["mtu"]; ok {
			tmp["mtu"] = flattenRouterOspf6Ospf6InterfaceMtu(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if cur_v, ok := i["mtu-ignore"]; ok {
			tmp["mtu_ignore"] = flattenRouterOspf6Ospf6InterfaceMtuIgnore(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if cur_v, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenRouterOspf6Ospf6InterfaceAuthentication(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if cur_v, ok := i["key-rollover-interval"]; ok {
			tmp["key_rollover_interval"] = flattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if cur_v, ok := i["ipsec-auth-alg"]; ok {
			tmp["ipsec_auth_alg"] = flattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if cur_v, ok := i["ipsec-enc-alg"]; ok {
			tmp["ipsec_enc_alg"] = flattenRouterOspf6Ospf6InterfaceIpsecEncAlg(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if cur_v, ok := i["ipsec-keys"]; ok {
			tmp["ipsec_keys"] = flattenRouterOspf6Ospf6InterfaceIpsecKeys(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if cur_v, ok := i["neighbor"]; ok {
			tmp["neighbor"] = flattenRouterOspf6Ospf6InterfaceNeighbor(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspf6Ospf6InterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceAreaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceKeyRolloverInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceIpsecAuthAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceIpsecEncAlg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if cur_v, ok := i["spi"]; ok {
			tmp["spi"] = flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := i["auth-key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := i["enc-key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["enc_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "spi", d)
	return result
}

func flattenRouterOspf6Ospf6InterfaceIpsecKeysSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if cur_v, ok := i["ip6"]; ok {
			tmp["ip6"] = flattenRouterOspf6Ospf6InterfaceNeighborIp6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if cur_v, ok := i["poll-interval"]; ok {
			tmp["poll_interval"] = flattenRouterOspf6Ospf6InterfaceNeighborPollInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if cur_v, ok := i["cost"]; ok {
			tmp["cost"] = flattenRouterOspf6Ospf6InterfaceNeighborCost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenRouterOspf6Ospf6InterfaceNeighborPriority(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip6", d)
	return result
}

func flattenRouterOspf6Ospf6InterfaceNeighborIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6Ospf6InterfaceNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceNeighborCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Ospf6InterfaceNeighborPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6Redistribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6RedistributeName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenRouterOspf6RedistributeStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if cur_v, ok := i["metric"]; ok {
			tmp["metric"] = flattenRouterOspf6RedistributeMetric(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if cur_v, ok := i["routemap"]; ok {
			tmp["routemap"] = flattenRouterOspf6RedistributeRoutemap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if cur_v, ok := i["metric-type"]; ok {
			tmp["metric_type"] = flattenRouterOspf6RedistributeMetricType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspf6RedistributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6RedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6RedistributeMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6PassiveInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterOspf6PassiveInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspf6PassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterOspf6SummaryAddressId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if cur_v, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterOspf6SummaryAddressPrefix6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if cur_v, ok := i["advertise"]; ok {
			tmp["advertise"] = flattenRouterOspf6SummaryAddressAdvertise(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if cur_v, ok := i["tag"]; ok {
			tmp["tag"] = flattenRouterOspf6SummaryAddressTag(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspf6SummaryAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspf6SummaryAddressPrefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspf6SummaryAddressTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectRouterOspf6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("abr_type", flattenRouterOspf6AbrType(o["abr-type"], d, "abr_type", sv)); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", flattenRouterOspf6AutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspf6DefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterOspf6LogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes", sv)); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspf6DefaultInformationMetric(o["default-information-metric"], d, "default_information_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspf6DefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", flattenRouterOspf6DefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspf6DefaultMetric(o["default-metric"], d, "default_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspf6RouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", flattenRouterOspf6SpfTimers(o["spf-timers"], d, "spf_timers", sv)); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspf6Bfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("restart_mode", flattenRouterOspf6RestartMode(o["restart-mode"], d, "restart_mode", sv)); err != nil {
		if !fortiAPIPatch(o["restart-mode"]) {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_period", flattenRouterOspf6RestartPeriod(o["restart-period"], d, "restart_period", sv)); err != nil {
		if !fortiAPIPatch(o["restart-period"]) {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if err = d.Set("restart_on_topology_change", flattenRouterOspf6RestartOnTopologyChange(o["restart-on-topology-change"], d, "restart_on_topology_change", sv)); err != nil {
		if !fortiAPIPatch(o["restart-on-topology-change"]) {
			return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("area", flattenRouterOspf6Area(o["area"], d, "area", sv)); err != nil {
			if !fortiAPIPatch(o["area"]) {
				return fmt.Errorf("Error reading area: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("area"); ok {
			if err = d.Set("area", flattenRouterOspf6Area(o["area"], d, "area", sv)); err != nil {
				if !fortiAPIPatch(o["area"]) {
					return fmt.Errorf("Error reading area: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface", sv)); err != nil {
			if !fortiAPIPatch(o["ospf6-interface"]) {
				return fmt.Errorf("Error reading ospf6_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ospf6_interface"); ok {
			if err = d.Set("ospf6_interface", flattenRouterOspf6Ospf6Interface(o["ospf6-interface"], d, "ospf6_interface", sv)); err != nil {
				if !fortiAPIPatch(o["ospf6-interface"]) {
					return fmt.Errorf("Error reading ospf6_interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterOspf6Redistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute"]) {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("passive_interface", flattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface", sv)); err != nil {
			if !fortiAPIPatch(o["passive-interface"]) {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("passive_interface"); ok {
			if err = d.Set("passive_interface", flattenRouterOspf6PassiveInterface(o["passive-interface"], d, "passive_interface", sv)); err != nil {
				if !fortiAPIPatch(o["passive-interface"]) {
					return fmt.Errorf("Error reading passive_interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address", sv)); err != nil {
			if !fortiAPIPatch(o["summary-address"]) {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterOspf6SummaryAddress(o["summary-address"], d, "summary_address", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterOspf6AbrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AutoCostRefBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6LogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultInformationRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6DefaultMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SpfTimers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Bfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RestartMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RestartPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RestartOnTopologyChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Area(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspf6AreaId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-cost"], _ = expandRouterOspf6AreaDefaultCost(d, i["default_cost"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-translator-role"], _ = expandRouterOspf6AreaNssaTranslatorRole(d, i["nssa_translator_role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stub-type"], _ = expandRouterOspf6AreaStubType(d, i["stub_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandRouterOspf6AreaType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginate(d, i["nssa_default_information_originate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginateMetric(d, i["nssa_default_information_originate_metric"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric-type"], _ = expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType(d, i["nssa_default_information_originate_metric_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-redistribution"], _ = expandRouterOspf6AreaNssaRedistribution(d, i["nssa_redistribution"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspf6AreaAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-rollover-interval"], _ = expandRouterOspf6AreaKeyRolloverInterval(d, i["key_rollover_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6AreaIpsecAuthAlg(d, i["ipsec_auth_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6AreaIpsecEncAlg(d, i["ipsec_enc_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-keys"], _ = expandRouterOspf6AreaIpsecKeys(d, i["ipsec_keys"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ipsec-keys"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["range"], _ = expandRouterOspf6AreaRange(d, i["range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-link"], _ = expandRouterOspf6AreaVirtualLink(d, i["virtual_link"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["virtual-link"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaDefaultCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaStubType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaKeyRolloverInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecAuthAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecEncAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecKeys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["spi"], _ = expandRouterOspf6AreaIpsecKeysSpi(d, i["spi"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["spi"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-key"], _ = expandRouterOspf6AreaIpsecKeysAuthKey(d, i["auth_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["enc-key"], _ = expandRouterOspf6AreaIpsecKeysEncKey(d, i["enc_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["enc-key"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaIpsecKeysSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecKeysAuthKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaIpsecKeysEncKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspf6AreaRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterOspf6AreaRangePrefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspf6AreaRangeAdvertise(d, i["advertise"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangePrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspf6AreaVirtualLinkName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspf6AreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspf6AreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspf6AreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspf6AreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peer"], _ = expandRouterOspf6AreaVirtualLinkPeer(d, i["peer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspf6AreaVirtualLinkAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-rollover-interval"], _ = expandRouterOspf6AreaVirtualLinkKeyRolloverInterval(d, i["key_rollover_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6AreaVirtualLinkIpsecAuthAlg(d, i["ipsec_auth_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6AreaVirtualLinkIpsecEncAlg(d, i["ipsec_enc_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-keys"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeys(d, i["ipsec_keys"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ipsec-keys"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkKeyRolloverInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecAuthAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecEncAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["spi"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysSpi(d, i["spi"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["spi"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey(d, i["auth_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["enc-key"], _ = expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey(d, i["enc_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["enc-key"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysAuthKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6AreaVirtualLinkIpsecKeysEncKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspf6Ospf6InterfaceName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["area-id"], _ = expandRouterOspf6Ospf6InterfaceAreaId(d, i["area_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterOspf6Ospf6InterfaceInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspf6Ospf6InterfaceRetransmitInterval(d, i["retransmit_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspf6Ospf6InterfaceTransmitDelay(d, i["transmit_delay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceCost(d, i["cost"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cost"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfacePriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspf6Ospf6InterfaceDeadInterval(d, i["dead_interval"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dead-interval"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspf6Ospf6InterfaceHelloInterval(d, i["hello_interval"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["hello-interval"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspf6Ospf6InterfaceStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["network-type"], _ = expandRouterOspf6Ospf6InterfaceNetworkType(d, i["network_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterOspf6Ospf6InterfaceBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu"], _ = expandRouterOspf6Ospf6InterfaceMtu(d, i["mtu"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mtu"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu-ignore"], _ = expandRouterOspf6Ospf6InterfaceMtuIgnore(d, i["mtu_ignore"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspf6Ospf6InterfaceAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_rollover_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-rollover-interval"], _ = expandRouterOspf6Ospf6InterfaceKeyRolloverInterval(d, i["key_rollover_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_auth_alg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-auth-alg"], _ = expandRouterOspf6Ospf6InterfaceIpsecAuthAlg(d, i["ipsec_auth_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_enc_alg"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-enc-alg"], _ = expandRouterOspf6Ospf6InterfaceIpsecEncAlg(d, i["ipsec_enc_alg"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_keys"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-keys"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeys(d, i["ipsec_keys"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ipsec-keys"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "neighbor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["neighbor"], _ = expandRouterOspf6Ospf6InterfaceNeighbor(d, i["neighbor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["neighbor"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceAreaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceKeyRolloverInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecAuthAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecEncAlg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spi"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["spi"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysSpi(d, i["spi"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["spi"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-key"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey(d, i["auth_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "enc_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["enc-key"], _ = expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey(d, i["enc_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["enc-key"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysAuthKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceIpsecKeysEncKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandRouterOspf6Ospf6InterfaceNeighborIp6(d, i["ip6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poll-interval"], _ = expandRouterOspf6Ospf6InterfaceNeighborPollInterval(d, i["poll_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspf6Ospf6InterfaceNeighborCost(d, i["cost"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cost"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspf6Ospf6InterfaceNeighborPriority(d, i["priority"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Ospf6InterfaceNeighborPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6Redistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspf6RedistributeName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspf6RedistributeStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric"], _ = expandRouterOspf6RedistributeMetric(d, i["metric"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["metric"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["routemap"], _ = expandRouterOspf6RedistributeRoutemap(d, i["routemap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["routemap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric-type"], _ = expandRouterOspf6RedistributeMetricType(d, i["metric_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6RedistributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6RedistributeMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6PassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterOspf6PassiveInterfaceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6PassiveInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspf6SummaryAddressId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterOspf6SummaryAddressPrefix6(d, i["prefix6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspf6SummaryAddressAdvertise(d, i["advertise"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tag"], _ = expandRouterOspf6SummaryAddressTag(d, i["tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tag"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspf6SummaryAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressPrefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspf6SummaryAddressTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf6(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("abr_type"); ok {
		if setArgNil {
			obj["abr-type"] = nil
		} else {
			t, err := expandRouterOspf6AbrType(d, v, "abr_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["abr-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_cost_ref_bandwidth"); ok {
		if setArgNil {
			obj["auto-cost-ref-bandwidth"] = nil
		} else {
			t, err := expandRouterOspf6AutoCostRefBandwidth(d, v, "auto_cost_ref_bandwidth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-cost-ref-bandwidth"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok {
		if setArgNil {
			obj["default-information-originate"] = nil
		} else {
			t, err := expandRouterOspf6DefaultInformationOriginate(d, v, "default_information_originate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-originate"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		if setArgNil {
			obj["log-neighbour-changes"] = nil
		} else {
			t, err := expandRouterOspf6LogNeighbourChanges(d, v, "log_neighbour_changes", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-neighbour-changes"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok {
		if setArgNil {
			obj["default-information-metric"] = nil
		} else {
			t, err := expandRouterOspf6DefaultInformationMetric(d, v, "default_information_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_metric_type"); ok {
		if setArgNil {
			obj["default-information-metric-type"] = nil
		} else {
			t, err := expandRouterOspf6DefaultInformationMetricType(d, v, "default_information_metric_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-metric-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_route_map"); ok {
		if setArgNil {
			obj["default-information-route-map"] = nil
		} else {
			t, err := expandRouterOspf6DefaultInformationRouteMap(d, v, "default_information_route_map", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-route-map"] = t
			}
		}
	} else if d.HasChange("default_information_route_map") {
		obj["default-information-route-map"] = nil
	}

	if v, ok := d.GetOk("default_metric"); ok {
		if setArgNil {
			obj["default-metric"] = nil
		} else {
			t, err := expandRouterOspf6DefaultMetric(d, v, "default_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		if setArgNil {
			obj["router-id"] = nil
		} else {
			t, err := expandRouterOspf6RouterId(d, v, "router_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["router-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("spf_timers"); ok {
		if setArgNil {
			obj["spf-timers"] = nil
		} else {
			t, err := expandRouterOspf6SpfTimers(d, v, "spf_timers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spf-timers"] = t
			}
		}
	} else if d.HasChange("spf_timers") {
		obj["spf-timers"] = nil
	}

	if v, ok := d.GetOk("bfd"); ok {
		if setArgNil {
			obj["bfd"] = nil
		} else {
			t, err := expandRouterOspf6Bfd(d, v, "bfd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd"] = t
			}
		}
	}

	if v, ok := d.GetOk("restart_mode"); ok {
		if setArgNil {
			obj["restart-mode"] = nil
		} else {
			t, err := expandRouterOspf6RestartMode(d, v, "restart_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["restart-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("restart_period"); ok {
		if setArgNil {
			obj["restart-period"] = nil
		} else {
			t, err := expandRouterOspf6RestartPeriod(d, v, "restart_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["restart-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("restart_on_topology_change"); ok {
		if setArgNil {
			obj["restart-on-topology-change"] = nil
		} else {
			t, err := expandRouterOspf6RestartOnTopologyChange(d, v, "restart_on_topology_change", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["restart-on-topology-change"] = t
			}
		}
	}

	if v, ok := d.GetOk("area"); ok || d.HasChange("area") {
		if setArgNil {
			obj["area"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspf6Area(d, v, "area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["area"] = t
			}
		}
	}

	if v, ok := d.GetOk("ospf6_interface"); ok || d.HasChange("ospf6_interface") {
		if setArgNil {
			obj["ospf6-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspf6Ospf6Interface(d, v, "ospf6_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		if setArgNil {
			obj["redistribute"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspf6Redistribute(d, v, "redistribute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute"] = t
			}
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok || d.HasChange("passive_interface") {
		if setArgNil {
			obj["passive-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspf6PassiveInterface(d, v, "passive_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["passive-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("summary_address"); ok || d.HasChange("summary_address") {
		if setArgNil {
			obj["summary-address"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspf6SummaryAddress(d, v, "summary_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["summary-address"] = t
			}
		}
	}

	return &obj, nil
}
