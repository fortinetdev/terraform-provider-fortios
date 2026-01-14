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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"lsa_refresh_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),
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
			},
			"distribute_route_map_in": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
						"comments": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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
									"keychain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"md5_key": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"md5_keychain": &schema.Schema{
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
									"md5_keys": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 255),
													Optional:     true,
												},
												"key_string": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 16),
													Optional:     true,
													Sensitive:    true,
												},
											},
										},
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
						},
						"comments": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"linkdown_fast_failover": &schema.Schema{
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
						"keychain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"md5_key": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"md5_keychain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"prefix_length": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),
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
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"hello_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"hello_multiplier": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 10),
							Optional:     true,
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
						"md5_keys": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
									},
									"key_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 16),
										Optional:     true,
										Sensitive:    true,
									},
								},
							},
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
						"comments": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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

func resourceRouterOspfUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterOspf(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterOspf(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterOspf(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterOspf resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspf(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterOspf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterOspf(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspf(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspf resource from API: %v", err)
	}
	return nil
}

func flattenRouterOspfAbrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAutoCostRefBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDistanceExternal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDistanceInterArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDistanceIntraArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDatabaseOverflow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDatabaseOverflowMaxLsas(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDatabaseOverflowTimeToRecover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDefaultInformationMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultInformationRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDefaultMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfLsaRefreshInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfRfc1583Compatible(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfSpfTimers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfLogNeighbourChanges(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDistributeListIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDistributeRouteMapIn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRestartMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRestartPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfRestartOnTopologyChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfArea(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfAreaId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["shortcut"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
			}
			tmp["shortcut"] = flattenRouterOspfAreaShortcut(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["authentication"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
			}
			tmp["authentication"] = flattenRouterOspfAreaAuthentication(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["default-cost"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
			}
			tmp["default_cost"] = flattenRouterOspfAreaDefaultCost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["nssa-translator-role"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
			}
			tmp["nssa_translator_role"] = flattenRouterOspfAreaNssaTranslatorRole(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["stub-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
			}
			tmp["stub_type"] = flattenRouterOspfAreaStubType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenRouterOspfAreaType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["nssa-default-information-originate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
			}
			tmp["nssa_default_information_originate"] = flattenRouterOspfAreaNssaDefaultInformationOriginate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["nssa-default-information-originate-metric"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
			}
			tmp["nssa_default_information_originate_metric"] = flattenRouterOspfAreaNssaDefaultInformationOriginateMetric(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["nssa-default-information-originate-metric-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
			}
			tmp["nssa_default_information_originate_metric_type"] = flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["nssa-redistribution"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
			}
			tmp["nssa_redistribution"] = flattenRouterOspfAreaNssaRedistribution(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comments"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
			}
			tmp["comments"] = flattenRouterOspfAreaComments(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
			}
			tmp["range"] = flattenRouterOspfAreaRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["virtual-link"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
			}
			tmp["virtual_link"] = flattenRouterOspfAreaVirtualLink(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["filter-list"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
			}
			tmp["filter_list"] = flattenRouterOspfAreaFilterList(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfAreaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaShortcut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaDefaultCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaNssaTranslatorRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaStubType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaNssaDefaultInformationOriginateMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaNssaRedistribution(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfAreaRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["prefix"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
			}
			tmp["prefix"] = flattenRouterOspfAreaRangePrefix(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["advertise"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
			}
			tmp["advertise"] = flattenRouterOspfAreaRangeAdvertise(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["substitute"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
			}
			tmp["substitute"] = flattenRouterOspfAreaRangeSubstitute(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["substitute-status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
			}
			tmp["substitute_status"] = flattenRouterOspfAreaRangeSubstituteStatus(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfAreaRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaRangePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfAreaRangeAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaRangeSubstitute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfAreaRangeSubstituteStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLink(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenRouterOspfAreaVirtualLinkName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["authentication"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
			}
			tmp["authentication"] = flattenRouterOspfAreaVirtualLinkAuthentication(cur_v, d, pre_append, sv)
		}

		if _, ok := i["authentication-key"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "authentication_key"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["authentication_key"] = c
				}
			}
		}

		if cur_v, ok := i["keychain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
			}
			tmp["keychain"] = flattenRouterOspfAreaVirtualLinkKeychain(cur_v, d, pre_append, sv)
		}

		if _, ok := i["md5-key"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "md5_key"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["md5_key"] = c
				}
			}
		}

		if cur_v, ok := i["md5-keychain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
			}
			tmp["md5_keychain"] = flattenRouterOspfAreaVirtualLinkMd5Keychain(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dead-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
			}
			tmp["dead_interval"] = flattenRouterOspfAreaVirtualLinkDeadInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["hello-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
			}
			tmp["hello_interval"] = flattenRouterOspfAreaVirtualLinkHelloInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["retransmit-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
			}
			tmp["retransmit_interval"] = flattenRouterOspfAreaVirtualLinkRetransmitInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["transmit-delay"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
			}
			tmp["transmit_delay"] = flattenRouterOspfAreaVirtualLinkTransmitDelay(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["peer"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
			}
			tmp["peer"] = flattenRouterOspfAreaVirtualLinkPeer(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["md5-keys"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
			}
			tmp["md5_keys"] = flattenRouterOspfAreaVirtualLinkMd5Keys(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspfAreaVirtualLinkName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkKeychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkMd5Keychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaVirtualLinkHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaVirtualLinkRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaVirtualLinkTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaVirtualLinkPeer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaVirtualLinkMd5Keys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfAreaVirtualLinkMd5KeysId(cur_v, d, pre_append, sv)
		}

		if _, ok := i["key-string"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "key_string"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["key_string"] = c
				}
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfAreaVirtualLinkMd5KeysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaFilterList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfAreaFilterListId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["list"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
			}
			tmp["list"] = flattenRouterOspfAreaFilterListList(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["direction"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
			}
			tmp["direction"] = flattenRouterOspfAreaFilterListDirection(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfAreaFilterListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfAreaFilterListList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfAreaFilterListDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenRouterOspfOspfInterfaceName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comments"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
			}
			tmp["comments"] = flattenRouterOspfOspfInterfaceComments(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
			}
			tmp["interface"] = flattenRouterOspfOspfInterfaceInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
			}
			tmp["ip"] = flattenRouterOspfOspfInterfaceIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["linkdown-fast-failover"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "linkdown_fast_failover"
			}
			tmp["linkdown_fast_failover"] = flattenRouterOspfOspfInterfaceLinkdownFastFailover(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["authentication"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
			}
			tmp["authentication"] = flattenRouterOspfOspfInterfaceAuthentication(cur_v, d, pre_append, sv)
		}

		if _, ok := i["authentication-key"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "authentication_key"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["authentication_key"] = c
				}
			}
		}

		if cur_v, ok := i["keychain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
			}
			tmp["keychain"] = flattenRouterOspfOspfInterfaceKeychain(cur_v, d, pre_append, sv)
		}

		if _, ok := i["md5-key"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "md5_key"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["md5_key"] = c
				}
			}
		}

		if cur_v, ok := i["md5-keychain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
			}
			tmp["md5_keychain"] = flattenRouterOspfOspfInterfaceMd5Keychain(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["prefix-length"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
			}
			tmp["prefix_length"] = flattenRouterOspfOspfInterfacePrefixLength(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["retransmit-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
			}
			tmp["retransmit_interval"] = flattenRouterOspfOspfInterfaceRetransmitInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["transmit-delay"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
			}
			tmp["transmit_delay"] = flattenRouterOspfOspfInterfaceTransmitDelay(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cost"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
			}
			tmp["cost"] = flattenRouterOspfOspfInterfaceCost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
			}
			tmp["priority"] = flattenRouterOspfOspfInterfacePriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dead-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
			}
			tmp["dead_interval"] = flattenRouterOspfOspfInterfaceDeadInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["hello-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
			}
			tmp["hello_interval"] = flattenRouterOspfOspfInterfaceHelloInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["hello-multiplier"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
			}
			tmp["hello_multiplier"] = flattenRouterOspfOspfInterfaceHelloMultiplier(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["database-filter-out"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
			}
			tmp["database_filter_out"] = flattenRouterOspfOspfInterfaceDatabaseFilterOut(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mtu"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
			}
			tmp["mtu"] = flattenRouterOspfOspfInterfaceMtu(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mtu-ignore"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
			}
			tmp["mtu_ignore"] = flattenRouterOspfOspfInterfaceMtuIgnore(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["network-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
			}
			tmp["network_type"] = flattenRouterOspfOspfInterfaceNetworkType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bfd"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
			}
			tmp["bfd"] = flattenRouterOspfOspfInterfaceBfd(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenRouterOspfOspfInterfaceStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["resync-timeout"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
			}
			tmp["resync_timeout"] = flattenRouterOspfOspfInterfaceResyncTimeout(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["md5-keys"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
			}
			tmp["md5_keys"] = flattenRouterOspfOspfInterfaceMd5Keys(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspfOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceLinkdownFastFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceKeychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMd5Keychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfacePrefixLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceDatabaseFilterOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfOspfInterfaceResyncTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfOspfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfOspfInterfaceMd5KeysId(cur_v, d, pre_append, sv)
		}

		if _, ok := i["key-string"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "key_string"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["key_string"] = c
				}
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfOspfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfNetworkId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["prefix"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
			}
			tmp["prefix"] = flattenRouterOspfNetworkPrefix(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["area"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
			}
			tmp["area"] = flattenRouterOspfNetworkArea(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comments"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
			}
			tmp["comments"] = flattenRouterOspfNetworkComments(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfNetworkArea(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfNetworkComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfNeighborId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
			}
			tmp["ip"] = flattenRouterOspfNeighborIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["poll-interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
			}
			tmp["poll_interval"] = flattenRouterOspfNeighborPollInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cost"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
			}
			tmp["cost"] = flattenRouterOspfNeighborCost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
			}
			tmp["priority"] = flattenRouterOspfNeighborPriority(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfNeighborId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfNeighborCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfNeighborPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfPassiveInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenRouterOspfPassiveInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspfPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfSummaryAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfSummaryAddressId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["prefix"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
			}
			tmp["prefix"] = flattenRouterOspfSummaryAddressPrefix(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["tag"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
			}
			tmp["tag"] = flattenRouterOspfSummaryAddressTag(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["advertise"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
			}
			tmp["advertise"] = flattenRouterOspfSummaryAddressAdvertise(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfSummaryAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfSummaryAddressPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterOspfSummaryAddressTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfSummaryAddressAdvertise(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDistributeList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenRouterOspfDistributeListId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["access-list"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
			}
			tmp["access_list"] = flattenRouterOspfDistributeListAccessList(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenRouterOspfDistributeListProtocol(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterOspfDistributeListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfDistributeListAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfDistributeListProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenRouterOspfRedistributeName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenRouterOspfRedistributeStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["metric"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
			}
			tmp["metric"] = flattenRouterOspfRedistributeMetric(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["routemap"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
			}
			tmp["routemap"] = flattenRouterOspfRedistributeRoutemap(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["metric-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
			}
			tmp["metric_type"] = flattenRouterOspfRedistributeMetricType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["tag"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
			}
			tmp["tag"] = flattenRouterOspfRedistributeTag(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterOspfRedistributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterOspfRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeMetricType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfRedistributeTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectRouterOspf(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("abr_type", flattenRouterOspfAbrType(o["abr-type"], d, "abr_type", sv)); err != nil {
		if !fortiAPIPatch(o["abr-type"]) {
			return fmt.Errorf("Error reading abr_type: %v", err)
		}
	}

	if err = d.Set("auto_cost_ref_bandwidth", flattenRouterOspfAutoCostRefBandwidth(o["auto-cost-ref-bandwidth"], d, "auto_cost_ref_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["auto-cost-ref-bandwidth"]) {
			return fmt.Errorf("Error reading auto_cost_ref_bandwidth: %v", err)
		}
	}

	if err = d.Set("distance_external", flattenRouterOspfDistanceExternal(o["distance-external"], d, "distance_external", sv)); err != nil {
		if !fortiAPIPatch(o["distance-external"]) {
			return fmt.Errorf("Error reading distance_external: %v", err)
		}
	}

	if err = d.Set("distance_inter_area", flattenRouterOspfDistanceInterArea(o["distance-inter-area"], d, "distance_inter_area", sv)); err != nil {
		if !fortiAPIPatch(o["distance-inter-area"]) {
			return fmt.Errorf("Error reading distance_inter_area: %v", err)
		}
	}

	if err = d.Set("distance_intra_area", flattenRouterOspfDistanceIntraArea(o["distance-intra-area"], d, "distance_intra_area", sv)); err != nil {
		if !fortiAPIPatch(o["distance-intra-area"]) {
			return fmt.Errorf("Error reading distance_intra_area: %v", err)
		}
	}

	if err = d.Set("database_overflow", flattenRouterOspfDatabaseOverflow(o["database-overflow"], d, "database_overflow", sv)); err != nil {
		if !fortiAPIPatch(o["database-overflow"]) {
			return fmt.Errorf("Error reading database_overflow: %v", err)
		}
	}

	if err = d.Set("database_overflow_max_lsas", flattenRouterOspfDatabaseOverflowMaxLsas(o["database-overflow-max-lsas"], d, "database_overflow_max_lsas", sv)); err != nil {
		if !fortiAPIPatch(o["database-overflow-max-lsas"]) {
			return fmt.Errorf("Error reading database_overflow_max_lsas: %v", err)
		}
	}

	if err = d.Set("database_overflow_time_to_recover", flattenRouterOspfDatabaseOverflowTimeToRecover(o["database-overflow-time-to-recover"], d, "database_overflow_time_to_recover", sv)); err != nil {
		if !fortiAPIPatch(o["database-overflow-time-to-recover"]) {
			return fmt.Errorf("Error reading database_overflow_time_to_recover: %v", err)
		}
	}

	if err = d.Set("default_information_originate", flattenRouterOspfDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_information_metric", flattenRouterOspfDefaultInformationMetric(o["default-information-metric"], d, "default_information_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-metric"]) {
			return fmt.Errorf("Error reading default_information_metric: %v", err)
		}
	}

	if err = d.Set("default_information_metric_type", flattenRouterOspfDefaultInformationMetricType(o["default-information-metric-type"], d, "default_information_metric_type", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-metric-type"]) {
			return fmt.Errorf("Error reading default_information_metric_type: %v", err)
		}
	}

	if err = d.Set("default_information_route_map", flattenRouterOspfDefaultInformationRouteMap(o["default-information-route-map"], d, "default_information_route_map", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-route-map"]) {
			return fmt.Errorf("Error reading default_information_route_map: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterOspfDefaultMetric(o["default-metric"], d, "default_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterOspfDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("lsa_refresh_interval", flattenRouterOspfLsaRefreshInterval(o["lsa-refresh-interval"], d, "lsa_refresh_interval", sv)); err != nil {
		if !fortiAPIPatch(o["lsa-refresh-interval"]) {
			return fmt.Errorf("Error reading lsa_refresh_interval: %v", err)
		}
	}

	if err = d.Set("rfc1583_compatible", flattenRouterOspfRfc1583Compatible(o["rfc1583-compatible"], d, "rfc1583_compatible", sv)); err != nil {
		if !fortiAPIPatch(o["rfc1583-compatible"]) {
			return fmt.Errorf("Error reading rfc1583_compatible: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspfRouterId(o["router-id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("spf_timers", flattenRouterOspfSpfTimers(o["spf-timers"], d, "spf_timers", sv)); err != nil {
		if !fortiAPIPatch(o["spf-timers"]) {
			return fmt.Errorf("Error reading spf_timers: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterOspfBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("log_neighbour_changes", flattenRouterOspfLogNeighbourChanges(o["log-neighbour-changes"], d, "log_neighbour_changes", sv)); err != nil {
		if !fortiAPIPatch(o["log-neighbour-changes"]) {
			return fmt.Errorf("Error reading log_neighbour_changes: %v", err)
		}
	}

	if err = d.Set("distribute_list_in", flattenRouterOspfDistributeListIn(o["distribute-list-in"], d, "distribute_list_in", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-list-in"]) {
			return fmt.Errorf("Error reading distribute_list_in: %v", err)
		}
	}

	if err = d.Set("distribute_route_map_in", flattenRouterOspfDistributeRouteMapIn(o["distribute-route-map-in"], d, "distribute_route_map_in", sv)); err != nil {
		if !fortiAPIPatch(o["distribute-route-map-in"]) {
			return fmt.Errorf("Error reading distribute_route_map_in: %v", err)
		}
	}

	if err = d.Set("restart_mode", flattenRouterOspfRestartMode(o["restart-mode"], d, "restart_mode", sv)); err != nil {
		if !fortiAPIPatch(o["restart-mode"]) {
			return fmt.Errorf("Error reading restart_mode: %v", err)
		}
	}

	if err = d.Set("restart_period", flattenRouterOspfRestartPeriod(o["restart-period"], d, "restart_period", sv)); err != nil {
		if !fortiAPIPatch(o["restart-period"]) {
			return fmt.Errorf("Error reading restart_period: %v", err)
		}
	}

	if err = d.Set("restart_on_topology_change", flattenRouterOspfRestartOnTopologyChange(o["restart-on-topology-change"], d, "restart_on_topology_change", sv)); err != nil {
		if !fortiAPIPatch(o["restart-on-topology-change"]) {
			return fmt.Errorf("Error reading restart_on_topology_change: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("area", flattenRouterOspfArea(o["area"], d, "area", sv)); err != nil {
			if !fortiAPIPatch(o["area"]) {
				return fmt.Errorf("Error reading area: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("area"); ok {
			if err = d.Set("area", flattenRouterOspfArea(o["area"], d, "area", sv)); err != nil {
				if !fortiAPIPatch(o["area"]) {
					return fmt.Errorf("Error reading area: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface", sv)); err != nil {
			if !fortiAPIPatch(o["ospf-interface"]) {
				return fmt.Errorf("Error reading ospf_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ospf_interface"); ok {
			if err = d.Set("ospf_interface", flattenRouterOspfOspfInterface(o["ospf-interface"], d, "ospf_interface", sv)); err != nil {
				if !fortiAPIPatch(o["ospf-interface"]) {
					return fmt.Errorf("Error reading ospf_interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("network", flattenRouterOspfNetwork(o["network"], d, "network", sv)); err != nil {
			if !fortiAPIPatch(o["network"]) {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterOspfNetwork(o["network"], d, "network", sv)); err != nil {
				if !fortiAPIPatch(o["network"]) {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("neighbor", flattenRouterOspfNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterOspfNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("passive_interface", flattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface", sv)); err != nil {
			if !fortiAPIPatch(o["passive-interface"]) {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("passive_interface"); ok {
			if err = d.Set("passive_interface", flattenRouterOspfPassiveInterface(o["passive-interface"], d, "passive_interface", sv)); err != nil {
				if !fortiAPIPatch(o["passive-interface"]) {
					return fmt.Errorf("Error reading passive_interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address", sv)); err != nil {
			if !fortiAPIPatch(o["summary-address"]) {
				return fmt.Errorf("Error reading summary_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("summary_address"); ok {
			if err = d.Set("summary_address", flattenRouterOspfSummaryAddress(o["summary-address"], d, "summary_address", sv)); err != nil {
				if !fortiAPIPatch(o["summary-address"]) {
					return fmt.Errorf("Error reading summary_address: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list", sv)); err != nil {
			if !fortiAPIPatch(o["distribute-list"]) {
				return fmt.Errorf("Error reading distribute_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distribute_list"); ok {
			if err = d.Set("distribute_list", flattenRouterOspfDistributeList(o["distribute-list"], d, "distribute_list", sv)); err != nil {
				if !fortiAPIPatch(o["distribute-list"]) {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("redistribute", flattenRouterOspfRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterOspfRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterOspfAbrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAutoCostRefBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceExternal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceInterArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistanceIntraArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflowMaxLsas(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDatabaseOverflowTimeToRecover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultInformationRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDefaultMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfLsaRefreshInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRfc1583Compatible(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSpfTimers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfLogNeighbourChanges(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeRouteMapIn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRestartOnTopologyChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shortcut"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shortcut"], _ = expandRouterOspfAreaShortcut(d, i["shortcut"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspfAreaAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default_cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default-cost"], _ = expandRouterOspfAreaDefaultCost(d, i["default_cost"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_translator_role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-translator-role"], _ = expandRouterOspfAreaNssaTranslatorRole(d, i["nssa_translator_role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "stub_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["stub-type"], _ = expandRouterOspfAreaStubType(d, i["stub_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandRouterOspfAreaType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate"], _ = expandRouterOspfAreaNssaDefaultInformationOriginate(d, i["nssa_default_information_originate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric"], _ = expandRouterOspfAreaNssaDefaultInformationOriginateMetric(d, i["nssa_default_information_originate_metric"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_default_information_originate_metric_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-default-information-originate-metric-type"], _ = expandRouterOspfAreaNssaDefaultInformationOriginateMetricType(d, i["nssa_default_information_originate_metric_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nssa_redistribution"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["nssa-redistribution"], _ = expandRouterOspfAreaNssaRedistribution(d, i["nssa_redistribution"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comments"], _ = expandRouterOspfAreaComments(d, i["comments"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comments"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["range"], _ = expandRouterOspfAreaRange(d, i["range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "virtual_link"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["virtual-link"], _ = expandRouterOspfAreaVirtualLink(d, i["virtual_link"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["virtual-link"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["filter-list"], _ = expandRouterOspfAreaFilterList(d, i["filter_list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["filter-list"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaShortcut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaDefaultCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaTranslatorRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaStubType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaDefaultInformationOriginateMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaNssaRedistribution(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterOspfAreaRangePrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspfAreaRangeAdvertise(d, i["advertise"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["substitute"], _ = expandRouterOspfAreaRangeSubstitute(d, i["substitute"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "substitute_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["substitute-status"], _ = expandRouterOspfAreaRangeSubstituteStatus(d, i["substitute_status"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeSubstitute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaRangeSubstituteStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfAreaVirtualLinkName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspfAreaVirtualLinkAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication-key"], _ = expandRouterOspfAreaVirtualLinkAuthenticationKey(d, i["authentication_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["authentication-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keychain"], _ = expandRouterOspfAreaVirtualLinkKeychain(d, i["keychain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["keychain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-key"], _ = expandRouterOspfAreaVirtualLinkMd5Key(d, i["md5_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-keychain"], _ = expandRouterOspfAreaVirtualLinkMd5Keychain(d, i["md5_keychain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-keychain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspfAreaVirtualLinkDeadInterval(d, i["dead_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspfAreaVirtualLinkHelloInterval(d, i["hello_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspfAreaVirtualLinkRetransmitInterval(d, i["retransmit_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspfAreaVirtualLinkTransmitDelay(d, i["transmit_delay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peer"], _ = expandRouterOspfAreaVirtualLinkPeer(d, i["peer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-keys"], _ = expandRouterOspfAreaVirtualLinkMd5Keys(d, i["md5_keys"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-keys"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkAuthenticationKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkKeychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5Key(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5Keychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaVirtualLinkMd5KeysId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-string"], _ = expandRouterOspfAreaVirtualLinkMd5KeysKeyString(d, i["key_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key-string"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaVirtualLinkMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfAreaFilterListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["list"], _ = expandRouterOspfAreaFilterListList(d, i["list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["list"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandRouterOspfAreaFilterListDirection(d, i["direction"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfAreaFilterListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfAreaFilterListDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfOspfInterfaceName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comments"], _ = expandRouterOspfOspfInterfaceComments(d, i["comments"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comments"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterOspfOspfInterfaceInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandRouterOspfOspfInterfaceIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "linkdown_fast_failover"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["linkdown-fast-failover"], _ = expandRouterOspfOspfInterfaceLinkdownFastFailover(d, i["linkdown_fast_failover"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandRouterOspfOspfInterfaceAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication-key"], _ = expandRouterOspfOspfInterfaceAuthenticationKey(d, i["authentication_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["authentication-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keychain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keychain"], _ = expandRouterOspfOspfInterfaceKeychain(d, i["keychain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["keychain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-key"], _ = expandRouterOspfOspfInterfaceMd5Key(d, i["md5_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keychain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-keychain"], _ = expandRouterOspfOspfInterfaceMd5Keychain(d, i["md5_keychain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-keychain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-length"], _ = expandRouterOspfOspfInterfacePrefixLength(d, i["prefix_length"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["prefix-length"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "retransmit_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["retransmit-interval"], _ = expandRouterOspfOspfInterfaceRetransmitInterval(d, i["retransmit_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "transmit_delay"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["transmit-delay"], _ = expandRouterOspfOspfInterfaceTransmitDelay(d, i["transmit_delay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspfOspfInterfaceCost(d, i["cost"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cost"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspfOspfInterfacePriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dead_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dead-interval"], _ = expandRouterOspfOspfInterfaceDeadInterval(d, i["dead_interval"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dead-interval"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterOspfOspfInterfaceHelloInterval(d, i["hello_interval"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["hello-interval"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_multiplier"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-multiplier"], _ = expandRouterOspfOspfInterfaceHelloMultiplier(d, i["hello_multiplier"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["hello-multiplier"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "database_filter_out"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["database-filter-out"], _ = expandRouterOspfOspfInterfaceDatabaseFilterOut(d, i["database_filter_out"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu"], _ = expandRouterOspfOspfInterfaceMtu(d, i["mtu"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mtu"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mtu_ignore"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mtu-ignore"], _ = expandRouterOspfOspfInterfaceMtuIgnore(d, i["mtu_ignore"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "network_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["network-type"], _ = expandRouterOspfOspfInterfaceNetworkType(d, i["network_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bfd"], _ = expandRouterOspfOspfInterfaceBfd(d, i["bfd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterOspfOspfInterfaceStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resync_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["resync-timeout"], _ = expandRouterOspfOspfInterfaceResyncTimeout(d, i["resync_timeout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_keys"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["md5-keys"], _ = expandRouterOspfOspfInterfaceMd5Keys(d, i["md5_keys"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["md5-keys"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfOspfInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceLinkdownFastFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceAuthenticationKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceKeychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Key(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Keychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePrefixLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceHelloMultiplier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceDatabaseFilterOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceResyncTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfOspfInterfaceMd5KeysId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-string"], _ = expandRouterOspfOspfInterfaceMd5KeysKeyString(d, i["key_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key-string"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfOspfInterfaceMd5KeysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfOspfInterfaceMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfNetworkId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterOspfNetworkPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "area"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["area"], _ = expandRouterOspfNetworkArea(d, i["area"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comments"], _ = expandRouterOspfNetworkComments(d, i["comments"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comments"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkArea(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNetworkComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfNeighborId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandRouterOspfNeighborIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "poll_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["poll-interval"], _ = expandRouterOspfNeighborPollInterval(d, i["poll_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandRouterOspfNeighborCost(d, i["cost"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cost"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandRouterOspfNeighborPriority(d, i["priority"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfNeighborId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfNeighborPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfPassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterOspfPassiveInterfaceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfPassiveInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfSummaryAddressId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterOspfSummaryAddressPrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tag"], _ = expandRouterOspfSummaryAddressTag(d, i["tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advertise"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advertise"], _ = expandRouterOspfSummaryAddressAdvertise(d, i["advertise"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfSummaryAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfSummaryAddressAdvertise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterOspfDistributeListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-list"], _ = expandRouterOspfDistributeListAccessList(d, i["access_list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["access-list"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandRouterOspfDistributeListProtocol(d, i["protocol"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfDistributeListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfDistributeListProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
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
			tmp["name"], _ = expandRouterOspfRedistributeName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		if setArgNil {
			tmp["status"] = nil
		} else {
			pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			if _, ok := d.GetOk(pre_append); ok {
				if setArgNil {
					tmp["status"] = nil
				} else {
					tmp["status"], _ = expandRouterOspfRedistributeStatus(d, i["status"], pre_append, sv)
				}
			}
		}

		if setArgNil {
			tmp["metric"] = nil
		} else {
			pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
			if _, ok := d.GetOk(pre_append); ok {
				if setArgNil {
					tmp["metric"] = nil
				} else {
					tmp["metric"], _ = expandRouterOspfRedistributeMetric(d, i["metric"], pre_append, sv)
				}
			} else if d.HasChange(pre_append) {
				tmp["metric"] = nil
			}
		}

		if setArgNil {
			tmp["routemap"] = nil
		} else {
			pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
			if _, ok := d.GetOk(pre_append); ok {
				if setArgNil {
					tmp["routemap"] = nil
				} else {
					tmp["routemap"], _ = expandRouterOspfRedistributeRoutemap(d, i["routemap"], pre_append, sv)
				}
			} else if d.HasChange(pre_append) {
				tmp["routemap"] = nil
			}
		}

		if setArgNil {
			tmp["metric-type"] = nil
		} else {
			pre_append = pre + "." + strconv.Itoa(con) + "." + "metric_type"
			if _, ok := d.GetOk(pre_append); ok {
				if setArgNil {
					tmp["metric-type"] = nil
				} else {
					tmp["metric-type"], _ = expandRouterOspfRedistributeMetricType(d, i["metric_type"], pre_append, sv)
				}
			}
		}

		if setArgNil {
			tmp["tag"] = nil
		} else {
			pre_append = pre + "." + strconv.Itoa(con) + "." + "tag"
			if _, ok := d.GetOk(pre_append); ok {
				if setArgNil {
					tmp["tag"] = nil
				} else {
					tmp["tag"], _ = expandRouterOspfRedistributeTag(d, i["tag"], pre_append, sv)
				}
			} else if d.HasChange(pre_append) {
				tmp["tag"] = nil
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterOspfRedistributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeMetricType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfRedistributeTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspf(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("abr_type"); ok {
		if setArgNil {
			obj["abr-type"] = nil
		} else {
			t, err := expandRouterOspfAbrType(d, v, "abr_type", sv)
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
			t, err := expandRouterOspfAutoCostRefBandwidth(d, v, "auto_cost_ref_bandwidth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-cost-ref-bandwidth"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance_external"); ok {
		if setArgNil {
			obj["distance-external"] = nil
		} else {
			t, err := expandRouterOspfDistanceExternal(d, v, "distance_external", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance-external"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance_inter_area"); ok {
		if setArgNil {
			obj["distance-inter-area"] = nil
		} else {
			t, err := expandRouterOspfDistanceInterArea(d, v, "distance_inter_area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance-inter-area"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance_intra_area"); ok {
		if setArgNil {
			obj["distance-intra-area"] = nil
		} else {
			t, err := expandRouterOspfDistanceIntraArea(d, v, "distance_intra_area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance-intra-area"] = t
			}
		}
	}

	if v, ok := d.GetOk("database_overflow"); ok {
		if setArgNil {
			obj["database-overflow"] = nil
		} else {
			t, err := expandRouterOspfDatabaseOverflow(d, v, "database_overflow", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["database-overflow"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("database_overflow_max_lsas"); ok {
		if setArgNil {
			obj["database-overflow-max-lsas"] = nil
		} else {
			t, err := expandRouterOspfDatabaseOverflowMaxLsas(d, v, "database_overflow_max_lsas", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["database-overflow-max-lsas"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("database_overflow_time_to_recover"); ok {
		if setArgNil {
			obj["database-overflow-time-to-recover"] = nil
		} else {
			t, err := expandRouterOspfDatabaseOverflowTimeToRecover(d, v, "database_overflow_time_to_recover", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["database-overflow-time-to-recover"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_originate"); ok {
		if setArgNil {
			obj["default-information-originate"] = nil
		} else {
			t, err := expandRouterOspfDefaultInformationOriginate(d, v, "default_information_originate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-information-originate"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_information_metric"); ok {
		if setArgNil {
			obj["default-information-metric"] = nil
		} else {
			t, err := expandRouterOspfDefaultInformationMetric(d, v, "default_information_metric", sv)
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
			t, err := expandRouterOspfDefaultInformationMetricType(d, v, "default_information_metric_type", sv)
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
			t, err := expandRouterOspfDefaultInformationRouteMap(d, v, "default_information_route_map", sv)
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
			t, err := expandRouterOspfDefaultMetric(d, v, "default_metric", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-metric"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		if setArgNil {
			obj["distance"] = nil
		} else {
			t, err := expandRouterOspfDistance(d, v, "distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("lsa_refresh_interval"); ok {
		if setArgNil {
			obj["lsa-refresh-interval"] = nil
		} else {
			t, err := expandRouterOspfLsaRefreshInterval(d, v, "lsa_refresh_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lsa-refresh-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("rfc1583_compatible"); ok {
		if setArgNil {
			obj["rfc1583-compatible"] = nil
		} else {
			t, err := expandRouterOspfRfc1583Compatible(d, v, "rfc1583_compatible", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rfc1583-compatible"] = t
			}
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		if setArgNil {
			obj["router-id"] = nil
		} else {
			t, err := expandRouterOspfRouterId(d, v, "router_id", sv)
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
			t, err := expandRouterOspfSpfTimers(d, v, "spf_timers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spf-timers"] = t
			}
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		if setArgNil {
			obj["bfd"] = nil
		} else {
			t, err := expandRouterOspfBfd(d, v, "bfd", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bfd"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_neighbour_changes"); ok {
		if setArgNil {
			obj["log-neighbour-changes"] = nil
		} else {
			t, err := expandRouterOspfLogNeighbourChanges(d, v, "log_neighbour_changes", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-neighbour-changes"] = t
			}
		}
	}

	if v, ok := d.GetOk("distribute_list_in"); ok {
		if setArgNil {
			obj["distribute-list-in"] = nil
		} else {
			t, err := expandRouterOspfDistributeListIn(d, v, "distribute_list_in", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distribute-list-in"] = t
			}
		}
	} else if d.HasChange("distribute_list_in") {
		obj["distribute-list-in"] = nil
	}

	if v, ok := d.GetOk("distribute_route_map_in"); ok {
		if setArgNil {
			obj["distribute-route-map-in"] = nil
		} else {
			t, err := expandRouterOspfDistributeRouteMapIn(d, v, "distribute_route_map_in", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distribute-route-map-in"] = t
			}
		}
	} else if d.HasChange("distribute_route_map_in") {
		obj["distribute-route-map-in"] = nil
	}

	if v, ok := d.GetOk("restart_mode"); ok {
		if setArgNil {
			obj["restart-mode"] = nil
		} else {
			t, err := expandRouterOspfRestartMode(d, v, "restart_mode", sv)
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
			t, err := expandRouterOspfRestartPeriod(d, v, "restart_period", sv)
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
			t, err := expandRouterOspfRestartOnTopologyChange(d, v, "restart_on_topology_change", sv)
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
			t, err := expandRouterOspfArea(d, v, "area", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["area"] = t
			}
		}
	}

	if v, ok := d.GetOk("ospf_interface"); ok || d.HasChange("ospf_interface") {
		if setArgNil {
			obj["ospf-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspfOspfInterface(d, v, "ospf_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("network"); ok || d.HasChange("network") {
		if setArgNil {
			obj["network"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspfNetwork(d, v, "network", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network"] = t
			}
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		if setArgNil {
			obj["neighbor"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspfNeighbor(d, v, "neighbor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok || d.HasChange("passive_interface") {
		if setArgNil {
			obj["passive-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspfPassiveInterface(d, v, "passive_interface", sv)
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
			t, err := expandRouterOspfSummaryAddress(d, v, "summary_address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["summary-address"] = t
			}
		}
	}

	if v, ok := d.GetOk("distribute_list"); ok || d.HasChange("distribute_list") {
		if setArgNil {
			obj["distribute-list"] = make([]struct{}, 0)
		} else {
			t, err := expandRouterOspfDistributeList(d, v, "distribute_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distribute-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("redistribute"); ok || d.HasChange("redistribute") {
		t, err := expandRouterOspfRedistribute(d, v, "redistribute", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	return &obj, nil
}
