// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure redundant internet connections using SD-WAN (formerly virtual WAN link).

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVirtualWanLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVirtualWanLinkUpdate,
		Read:   resourceSystemVirtualWanLinkRead,
		Update: resourceSystemVirtualWanLinkUpdate,
		Delete: resourceSystemVirtualWanLinkDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"load_balance_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"neighbor_hold_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"neighbor_hold_down_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),
				Optional:     true,
			},
			"neighbor_hold_boot_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),
				Optional:     true,
			},
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fail_alert_interfaces": &schema.Schema{
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
			"zone": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"seq_num": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"spillover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16776000),
							Optional:     true,
						},
						"ingress_spillover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16776000),
							Optional:     true,
						},
						"volume_ratio": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"probe_packets": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"system_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
						"security_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"packet_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(64, 1024),
							Optional:     true,
						},
						"ha_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 50),
							Optional:     true,
						},
						"http_get": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
						},
						"http_agent": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
						},
						"http_match": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
						},
						"dns_request_domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600000),
							Optional:     true,
						},
						"probe_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(500, 5000),
							Optional:     true,
						},
						"failtime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),
							Optional:     true,
						},
						"recoverytime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),
							Optional:     true,
						},
						"probe_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 30),
							Optional:     true,
						},
						"diffservcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_cascade_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_static_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sla_fail_log_period": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
							Optional:     true,
						},
						"sla_pass_log_period": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
							Optional:     true,
						},
						"threshold_warning_packetloss": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
						},
						"threshold_alert_packetloss": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
						},
						"threshold_warning_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_alert_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_warning_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"threshold_alert_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sla": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 32),
										Optional:     true,
									},
									"link_cost_factor": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"latency_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
									},
									"jitter_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
									},
									"packetloss_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),
							Optional:     true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"health_check": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"sla_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4000),
							Optional:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"input_device": &schema.Schema{
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
						"input_device_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"standalone_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"quality_link": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"tos_mask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"start_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"end_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"route_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dst": &schema.Schema{
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
						"dst_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"src": &schema.Schema{
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
						"dst6": &schema.Schema{
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
						"src6": &schema.Schema{
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
						"src_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"users": &schema.Schema{
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
						"groups": &schema.Schema{
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
						"internet_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"internet_service_custom": &schema.Schema{
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
						"internet_service_custom_group": &schema.Schema{
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
						"internet_service_name": &schema.Schema{
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
						"internet_service_id": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"internet_service_group": &schema.Schema{
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
						"internet_service_app_ctrl": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"internet_service_app_ctrl_group": &schema.Schema{
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
						"internet_service_ctrl": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"internet_service_ctrl_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
									},
								},
							},
						},
						"health_check": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"link_cost_factor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"packet_loss_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"latency_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"jitter_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"bandwidth_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"link_cost_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"hold_down_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"dscp_forward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp_reverse": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp_forward_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dscp_reverse_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sla": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"health_check": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"priority_members": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sla_compare_method": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceSystemVirtualWanLinkUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemVirtualWanLink(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWanLink resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVirtualWanLink(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWanLink resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVirtualWanLink")
	}

	return resourceSystemVirtualWanLinkRead(d, m)
}

func resourceSystemVirtualWanLinkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVirtualWanLink(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWanLink resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVirtualWanLink(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemVirtualWanLink resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVirtualWanLinkRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemVirtualWanLink(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualWanLink resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVirtualWanLink(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualWanLink resource from API: %v", err)
	}
	return nil
}

func flattenSystemVirtualWanLinkStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkLoadBalanceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkNeighborHoldDown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkNeighborHoldDownTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkNeighborHoldBootTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkFailDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkFailAlertInterfacesName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkFailAlertInterfacesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkZone(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkZoneName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkZoneName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "seq-num", "seq_num")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["seq-num"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
			}
			tmp["seq_num"] = flattenSystemVirtualWanLinkMembersSeqNum(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
			}
			tmp["interface"] = flattenSystemVirtualWanLinkMembersInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["gateway"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
			}
			tmp["gateway"] = flattenSystemVirtualWanLinkMembersGateway(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["source"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
			}
			tmp["source"] = flattenSystemVirtualWanLinkMembersSource(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["gateway6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
			}
			tmp["gateway6"] = flattenSystemVirtualWanLinkMembersGateway6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["source6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
			}
			tmp["source6"] = flattenSystemVirtualWanLinkMembersSource6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["cost"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
			}
			tmp["cost"] = flattenSystemVirtualWanLinkMembersCost(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
			}
			tmp["weight"] = flattenSystemVirtualWanLinkMembersWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
			}
			tmp["priority"] = flattenSystemVirtualWanLinkMembersPriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["spillover-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
			}
			tmp["spillover_threshold"] = flattenSystemVirtualWanLinkMembersSpilloverThreshold(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ingress-spillover-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
			}
			tmp["ingress_spillover_threshold"] = flattenSystemVirtualWanLinkMembersIngressSpilloverThreshold(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["volume-ratio"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
			}
			tmp["volume_ratio"] = flattenSystemVirtualWanLinkMembersVolumeRatio(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenSystemVirtualWanLinkMembersStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comment"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
			}
			tmp["comment"] = flattenSystemVirtualWanLinkMembersComment(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq_num", d)
	return result
}

func flattenSystemVirtualWanLinkMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembersGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembersSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembersGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembersSource6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembersCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkMembersStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkMembersComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkHealthCheckName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["probe-packets"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
			}
			tmp["probe_packets"] = flattenSystemVirtualWanLinkHealthCheckProbePackets(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
			}
			tmp["addr_mode"] = flattenSystemVirtualWanLinkHealthCheckAddrMode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["system-dns"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
			}
			tmp["system_dns"] = flattenSystemVirtualWanLinkHealthCheckSystemDns(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["server"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
			}
			tmp["server"] = flattenSystemVirtualWanLinkHealthCheckServer(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenSystemVirtualWanLinkHealthCheckProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenSystemVirtualWanLinkHealthCheckPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["security-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
			}
			tmp["security_mode"] = flattenSystemVirtualWanLinkHealthCheckSecurityMode(cur_v, d, pre_append, sv)
		}

		if _, ok := i["password"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "password"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["password"] = c
				}
			}
		}

		if cur_v, ok := i["packet-size"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
			}
			tmp["packet_size"] = flattenSystemVirtualWanLinkHealthCheckPacketSize(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ha-priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
			}
			tmp["ha_priority"] = flattenSystemVirtualWanLinkHealthCheckHaPriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-get"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
			}
			tmp["http_get"] = flattenSystemVirtualWanLinkHealthCheckHttpGet(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-agent"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
			}
			tmp["http_agent"] = flattenSystemVirtualWanLinkHealthCheckHttpAgent(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["http-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
			}
			tmp["http_match"] = flattenSystemVirtualWanLinkHealthCheckHttpMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dns-request-domain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
			}
			tmp["dns_request_domain"] = flattenSystemVirtualWanLinkHealthCheckDnsRequestDomain(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interval"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
			}
			tmp["interval"] = flattenSystemVirtualWanLinkHealthCheckInterval(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["probe-timeout"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
			}
			tmp["probe_timeout"] = flattenSystemVirtualWanLinkHealthCheckProbeTimeout(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["failtime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
			}
			tmp["failtime"] = flattenSystemVirtualWanLinkHealthCheckFailtime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["recoverytime"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
			}
			tmp["recoverytime"] = flattenSystemVirtualWanLinkHealthCheckRecoverytime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["probe-count"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
			}
			tmp["probe_count"] = flattenSystemVirtualWanLinkHealthCheckProbeCount(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["diffservcode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
			}
			tmp["diffservcode"] = flattenSystemVirtualWanLinkHealthCheckDiffservcode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["update-cascade-interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
			}
			tmp["update_cascade_interface"] = flattenSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["update-static-route"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
			}
			tmp["update_static_route"] = flattenSystemVirtualWanLinkHealthCheckUpdateStaticRoute(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla-fail-log-period"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
			}
			tmp["sla_fail_log_period"] = flattenSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla-pass-log-period"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
			}
			tmp["sla_pass_log_period"] = flattenSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["threshold-warning-packetloss"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
			}
			tmp["threshold_warning_packetloss"] = flattenSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["threshold-alert-packetloss"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
			}
			tmp["threshold_alert_packetloss"] = flattenSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["threshold-warning-latency"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
			}
			tmp["threshold_warning_latency"] = flattenSystemVirtualWanLinkHealthCheckThresholdWarningLatency(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["threshold-alert-latency"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
			}
			tmp["threshold_alert_latency"] = flattenSystemVirtualWanLinkHealthCheckThresholdAlertLatency(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["threshold-warning-jitter"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
			}
			tmp["threshold_warning_jitter"] = flattenSystemVirtualWanLinkHealthCheckThresholdWarningJitter(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["threshold-alert-jitter"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
			}
			tmp["threshold_alert_jitter"] = flattenSystemVirtualWanLinkHealthCheckThresholdAlertJitter(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["members"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
			}
			tmp["members"] = flattenSystemVirtualWanLinkHealthCheckMembers(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
			}
			tmp["sla"] = flattenSystemVirtualWanLinkHealthCheckSla(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkHealthCheckName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckSystemDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckProbeCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "seq-num", "seq_num")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["seq-num"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
			}
			tmp["seq_num"] = flattenSystemVirtualWanLinkHealthCheckMembersSeqNum(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq_num", d)
	return result
}

func flattenSystemVirtualWanLinkHealthCheckMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemVirtualWanLinkHealthCheckSlaId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["link-cost-factor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
			}
			tmp["link_cost_factor"] = flattenSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["latency-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
			}
			tmp["latency_threshold"] = flattenSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["jitter-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
			}
			tmp["jitter_threshold"] = flattenSystemVirtualWanLinkHealthCheckSlaJitterThreshold(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["packetloss-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
			}
			tmp["packetloss_threshold"] = flattenSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemVirtualWanLinkHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "ip", "ip")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
			}
			tmp["ip"] = flattenSystemVirtualWanLinkNeighborIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["member"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
			}
			tmp["member"] = flattenSystemVirtualWanLinkNeighborMember(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["role"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
			}
			tmp["role"] = flattenSystemVirtualWanLinkNeighborRole(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
			}
			tmp["health_check"] = flattenSystemVirtualWanLinkNeighborHealthCheck(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
			}
			tmp["sla_id"] = flattenSystemVirtualWanLinkNeighborSlaId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenSystemVirtualWanLinkNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkNeighborMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkNeighborRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkNeighborSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemVirtualWanLinkServiceId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemVirtualWanLinkServiceName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
			}
			tmp["addr_mode"] = flattenSystemVirtualWanLinkServiceAddrMode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["input-device"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
			}
			tmp["input_device"] = flattenSystemVirtualWanLinkServiceInputDevice(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["input-device-negate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
			}
			tmp["input_device_negate"] = flattenSystemVirtualWanLinkServiceInputDeviceNegate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
			}
			tmp["mode"] = flattenSystemVirtualWanLinkServiceMode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["role"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
			}
			tmp["role"] = flattenSystemVirtualWanLinkServiceRole(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["standalone-action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
			}
			tmp["standalone_action"] = flattenSystemVirtualWanLinkServiceStandaloneAction(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["quality-link"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
			}
			tmp["quality_link"] = flattenSystemVirtualWanLinkServiceQualityLink(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["member"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
			}
			tmp["member"] = flattenSystemVirtualWanLinkServiceMember(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["tos"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
			}
			tmp["tos"] = flattenSystemVirtualWanLinkServiceTos(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["tos-mask"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
			}
			tmp["tos_mask"] = flattenSystemVirtualWanLinkServiceTosMask(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenSystemVirtualWanLinkServiceProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
			}
			tmp["start_port"] = flattenSystemVirtualWanLinkServiceStartPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
			}
			tmp["end_port"] = flattenSystemVirtualWanLinkServiceEndPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["route-tag"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
			}
			tmp["route_tag"] = flattenSystemVirtualWanLinkServiceRouteTag(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
			}
			tmp["dst"] = flattenSystemVirtualWanLinkServiceDst(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst-negate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
			}
			tmp["dst_negate"] = flattenSystemVirtualWanLinkServiceDstNegate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["src"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
			}
			tmp["src"] = flattenSystemVirtualWanLinkServiceSrc(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
			}
			tmp["dst6"] = flattenSystemVirtualWanLinkServiceDst6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["src6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
			}
			tmp["src6"] = flattenSystemVirtualWanLinkServiceSrc6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["src-negate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
			}
			tmp["src_negate"] = flattenSystemVirtualWanLinkServiceSrcNegate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["users"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
			}
			tmp["users"] = flattenSystemVirtualWanLinkServiceUsers(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["groups"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
			}
			tmp["groups"] = flattenSystemVirtualWanLinkServiceGroups(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
			}
			tmp["internet_service"] = flattenSystemVirtualWanLinkServiceInternetService(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-custom"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
			}
			tmp["internet_service_custom"] = flattenSystemVirtualWanLinkServiceInternetServiceCustom(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-custom-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
			}
			tmp["internet_service_custom_group"] = flattenSystemVirtualWanLinkServiceInternetServiceCustomGroup(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
			}
			tmp["internet_service_name"] = flattenSystemVirtualWanLinkServiceInternetServiceName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
			}
			tmp["internet_service_id"] = flattenSystemVirtualWanLinkServiceInternetServiceId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
			}
			tmp["internet_service_group"] = flattenSystemVirtualWanLinkServiceInternetServiceGroup(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-app-ctrl"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
			}
			tmp["internet_service_app_ctrl"] = flattenSystemVirtualWanLinkServiceInternetServiceAppCtrl(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-app-ctrl-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
			}
			tmp["internet_service_app_ctrl_group"] = flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-ctrl"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl"
			}
			tmp["internet_service_ctrl"] = flattenSystemVirtualWanLinkServiceInternetServiceCtrl(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["internet-service-ctrl-group"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl_group"
			}
			tmp["internet_service_ctrl_group"] = flattenSystemVirtualWanLinkServiceInternetServiceCtrlGroup(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["health-check"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
			}
			tmp["health_check"] = flattenSystemVirtualWanLinkServiceHealthCheck(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["link-cost-factor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
			}
			tmp["link_cost_factor"] = flattenSystemVirtualWanLinkServiceLinkCostFactor(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["packet-loss-weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
			}
			tmp["packet_loss_weight"] = flattenSystemVirtualWanLinkServicePacketLossWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["latency-weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
			}
			tmp["latency_weight"] = flattenSystemVirtualWanLinkServiceLatencyWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["jitter-weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
			}
			tmp["jitter_weight"] = flattenSystemVirtualWanLinkServiceJitterWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["bandwidth-weight"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
			}
			tmp["bandwidth_weight"] = flattenSystemVirtualWanLinkServiceBandwidthWeight(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["link-cost-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
			}
			tmp["link_cost_threshold"] = flattenSystemVirtualWanLinkServiceLinkCostThreshold(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["hold-down-time"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
			}
			tmp["hold_down_time"] = flattenSystemVirtualWanLinkServiceHoldDownTime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp-forward"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
			}
			tmp["dscp_forward"] = flattenSystemVirtualWanLinkServiceDscpForward(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp-reverse"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
			}
			tmp["dscp_reverse"] = flattenSystemVirtualWanLinkServiceDscpReverse(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp-forward-tag"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
			}
			tmp["dscp_forward_tag"] = flattenSystemVirtualWanLinkServiceDscpForwardTag(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dscp-reverse-tag"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
			}
			tmp["dscp_reverse_tag"] = flattenSystemVirtualWanLinkServiceDscpReverseTag(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
			}
			tmp["sla"] = flattenSystemVirtualWanLinkServiceSla(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["priority-members"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
			}
			tmp["priority_members"] = flattenSystemVirtualWanLinkServicePriorityMembers(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenSystemVirtualWanLinkServiceStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["gateway"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
			}
			tmp["gateway"] = flattenSystemVirtualWanLinkServiceGateway(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["default"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
			}
			tmp["default"] = flattenSystemVirtualWanLinkServiceDefault(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["sla-compare-method"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
			}
			tmp["sla_compare_method"] = flattenSystemVirtualWanLinkServiceSlaCompareMethod(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemVirtualWanLinkServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInputDeviceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInputDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceQualityLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceTosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceRouteTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceDstName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceDstNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceSrc(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceSrcName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceSrcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceDst6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceDst6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceDst6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceSrc6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceSrc6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceSrc6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceUsersName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceGroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInternetServiceCustomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInternetServiceCustomGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInternetServiceNameName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemVirtualWanLinkServiceInternetServiceIdId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInternetServiceGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceInternetServiceCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemVirtualWanLinkServiceInternetServiceCtrlId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceCtrlId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceInternetServiceCtrlGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVirtualWanLinkServiceInternetServiceCtrlGroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemVirtualWanLinkServiceInternetServiceCtrlGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("health_check"), "name")
}

func flattenSystemVirtualWanLinkServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceDscpForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "health-check", "health_check")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["health-check"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
			}
			tmp["health_check"] = flattenSystemVirtualWanLinkServiceSlaHealthCheck(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenSystemVirtualWanLinkServiceSlaId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "health_check", d)
	return result
}

func flattenSystemVirtualWanLinkServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "seq-num", "seq_num")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["seq-num"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
			}
			tmp["seq_num"] = flattenSystemVirtualWanLinkServicePriorityMembersSeqNum(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq_num", d)
	return result
}

func flattenSystemVirtualWanLinkServicePriorityMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemVirtualWanLinkServiceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWanLinkServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVirtualWanLink(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemVirtualWanLinkStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("load_balance_mode", flattenSystemVirtualWanLinkLoadBalanceMode(o["load-balance-mode"], d, "load_balance_mode", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-mode"]) {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenSystemVirtualWanLinkNeighborHoldDown(o["neighbor-hold-down"], d, "neighbor_hold_down", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-hold-down"]) {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenSystemVirtualWanLinkNeighborHoldDownTime(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-hold-down-time"]) {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_boot_time", flattenSystemVirtualWanLinkNeighborHoldBootTime(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-hold-boot-time"]) {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenSystemVirtualWanLinkFailDetect(o["fail-detect"], d, "fail_detect", sv)); err != nil {
		if !fortiAPIPatch(o["fail-detect"]) {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("fail_alert_interfaces", flattenSystemVirtualWanLinkFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces", sv)); err != nil {
			if !fortiAPIPatch(o["fail-alert-interfaces"]) {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fail_alert_interfaces"); ok {
			if err = d.Set("fail_alert_interfaces", flattenSystemVirtualWanLinkFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces", sv)); err != nil {
				if !fortiAPIPatch(o["fail-alert-interfaces"]) {
					return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("zone", flattenSystemVirtualWanLinkZone(o["zone"], d, "zone", sv)); err != nil {
			if !fortiAPIPatch(o["zone"]) {
				return fmt.Errorf("Error reading zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("zone"); ok {
			if err = d.Set("zone", flattenSystemVirtualWanLinkZone(o["zone"], d, "zone", sv)); err != nil {
				if !fortiAPIPatch(o["zone"]) {
					return fmt.Errorf("Error reading zone: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("members", flattenSystemVirtualWanLinkMembers(o["members"], d, "members", sv)); err != nil {
			if !fortiAPIPatch(o["members"]) {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenSystemVirtualWanLinkMembers(o["members"], d, "members", sv)); err != nil {
				if !fortiAPIPatch(o["members"]) {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("health_check", flattenSystemVirtualWanLinkHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
			if !fortiAPIPatch(o["health-check"]) {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check"); ok {
			if err = d.Set("health_check", flattenSystemVirtualWanLinkHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
				if !fortiAPIPatch(o["health-check"]) {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("neighbor", flattenSystemVirtualWanLinkNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenSystemVirtualWanLinkNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("service", flattenSystemVirtualWanLinkService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenSystemVirtualWanLinkService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemVirtualWanLinkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVirtualWanLinkStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkLoadBalanceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborHoldDown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborHoldDownTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborHoldBootTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkFailDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkFailAlertInterfacesName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkFailAlertInterfacesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemVirtualWanLinkZoneName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkZoneName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["seq-num"], _ = expandSystemVirtualWanLinkMembersSeqNum(d, i["seq_num"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["seq-num"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemVirtualWanLinkMembersInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway"], _ = expandSystemVirtualWanLinkMembersGateway(d, i["gateway"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["gateway"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source"], _ = expandSystemVirtualWanLinkMembersSource(d, i["source"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["source"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway6"], _ = expandSystemVirtualWanLinkMembersGateway6(d, i["gateway6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["gateway6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source6"], _ = expandSystemVirtualWanLinkMembersSource6(d, i["source6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["source6"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cost"], _ = expandSystemVirtualWanLinkMembersCost(d, i["cost"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["cost"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandSystemVirtualWanLinkMembersWeight(d, i["weight"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["weight"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSystemVirtualWanLinkMembersPriority(d, i["priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["spillover-threshold"], _ = expandSystemVirtualWanLinkMembersSpilloverThreshold(d, i["spillover_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["spillover-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ingress-spillover-threshold"], _ = expandSystemVirtualWanLinkMembersIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ingress-spillover-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["volume-ratio"], _ = expandSystemVirtualWanLinkMembersVolumeRatio(d, i["volume_ratio"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["volume-ratio"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSystemVirtualWanLinkMembersStatus(d, i["status"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["status"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandSystemVirtualWanLinkMembersComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersSource6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkMembersComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemVirtualWanLinkHealthCheckName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["probe-packets"], _ = expandSystemVirtualWanLinkHealthCheckProbePackets(d, i["probe_packets"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["probe-packets"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-mode"], _ = expandSystemVirtualWanLinkHealthCheckAddrMode(d, i["addr_mode"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["addr-mode"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["system-dns"], _ = expandSystemVirtualWanLinkHealthCheckSystemDns(d, i["system_dns"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["system-dns"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandSystemVirtualWanLinkHealthCheckServer(d, i["server"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["server"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandSystemVirtualWanLinkHealthCheckProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandSystemVirtualWanLinkHealthCheckPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["security-mode"], _ = expandSystemVirtualWanLinkHealthCheckSecurityMode(d, i["security_mode"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["security-mode"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandSystemVirtualWanLinkHealthCheckPassword(d, i["password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["packet-size"], _ = expandSystemVirtualWanLinkHealthCheckPacketSize(d, i["packet_size"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["packet-size"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-priority"], _ = expandSystemVirtualWanLinkHealthCheckHaPriority(d, i["ha_priority"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ha-priority"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-get"], _ = expandSystemVirtualWanLinkHealthCheckHttpGet(d, i["http_get"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-get"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-agent"], _ = expandSystemVirtualWanLinkHealthCheckHttpAgent(d, i["http_agent"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-agent"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["http-match"], _ = expandSystemVirtualWanLinkHealthCheckHttpMatch(d, i["http_match"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["http-match"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dns-request-domain"], _ = expandSystemVirtualWanLinkHealthCheckDnsRequestDomain(d, i["dns_request_domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dns-request-domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interval"], _ = expandSystemVirtualWanLinkHealthCheckInterval(d, i["interval"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interval"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["probe-timeout"], _ = expandSystemVirtualWanLinkHealthCheckProbeTimeout(d, i["probe_timeout"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["probe-timeout"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["failtime"], _ = expandSystemVirtualWanLinkHealthCheckFailtime(d, i["failtime"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["failtime"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["recoverytime"], _ = expandSystemVirtualWanLinkHealthCheckRecoverytime(d, i["recoverytime"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["recoverytime"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["probe-count"], _ = expandSystemVirtualWanLinkHealthCheckProbeCount(d, i["probe_count"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["probe-count"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["diffservcode"], _ = expandSystemVirtualWanLinkHealthCheckDiffservcode(d, i["diffservcode"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["diffservcode"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["update-cascade-interface"], _ = expandSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["update-cascade-interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["update-static-route"], _ = expandSystemVirtualWanLinkHealthCheckUpdateStaticRoute(d, i["update_static_route"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["update-static-route"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla-fail-log-period"], _ = expandSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla-fail-log-period"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla-pass-log-period"], _ = expandSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla-pass-log-period"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold-warning-packetloss"], _ = expandSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold-warning-packetloss"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold-alert-packetloss"], _ = expandSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold-alert-packetloss"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold-warning-latency"], _ = expandSystemVirtualWanLinkHealthCheckThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold-warning-latency"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold-alert-latency"], _ = expandSystemVirtualWanLinkHealthCheckThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold-alert-latency"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold-warning-jitter"], _ = expandSystemVirtualWanLinkHealthCheckThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold-warning-jitter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold-alert-jitter"], _ = expandSystemVirtualWanLinkHealthCheckThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["threshold-alert-jitter"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["members"], _ = expandSystemVirtualWanLinkHealthCheckMembers(d, i["members"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["members"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla"], _ = expandSystemVirtualWanLinkHealthCheckSla(d, i["sla"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkHealthCheckName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckProbePackets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSystemDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSecurityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckPacketSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckHaPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckHttpGet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckHttpAgent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckHttpMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckProbeTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckFailtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckRecoverytime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckProbeCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckDiffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["seq-num"], _ = expandSystemVirtualWanLinkHealthCheckMembersSeqNum(d, i["seq_num"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkHealthCheckMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemVirtualWanLinkHealthCheckSlaId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["link-cost-factor"], _ = expandSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(d, i["link_cost_factor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["link-cost-factor"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["latency-threshold"], _ = expandSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(d, i["latency_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["latency-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["jitter-threshold"], _ = expandSystemVirtualWanLinkHealthCheckSlaJitterThreshold(d, i["jitter_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["jitter-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["packetloss-threshold"], _ = expandSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["packetloss-threshold"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemVirtualWanLinkNeighborIp(d, i["ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["member"], _ = expandSystemVirtualWanLinkNeighborMember(d, i["member"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["member"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandSystemVirtualWanLinkNeighborRole(d, i["role"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["role"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandSystemVirtualWanLinkNeighborHealthCheck(d, i["health_check"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["health-check"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla-id"], _ = expandSystemVirtualWanLinkNeighborSlaId(d, i["sla_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla-id"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkNeighborSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemVirtualWanLinkServiceId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemVirtualWanLinkServiceName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-mode"], _ = expandSystemVirtualWanLinkServiceAddrMode(d, i["addr_mode"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["addr-mode"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["input-device"], _ = expandSystemVirtualWanLinkServiceInputDevice(d, i["input_device"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["input-device"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["input-device-negate"], _ = expandSystemVirtualWanLinkServiceInputDeviceNegate(d, i["input_device_negate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["input-device-negate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mode"], _ = expandSystemVirtualWanLinkServiceMode(d, i["mode"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mode"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandSystemVirtualWanLinkServiceRole(d, i["role"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["role"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["standalone-action"], _ = expandSystemVirtualWanLinkServiceStandaloneAction(d, i["standalone_action"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["standalone-action"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quality-link"], _ = expandSystemVirtualWanLinkServiceQualityLink(d, i["quality_link"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["quality-link"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["member"], _ = expandSystemVirtualWanLinkServiceMember(d, i["member"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["member"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tos"], _ = expandSystemVirtualWanLinkServiceTos(d, i["tos"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tos"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tos-mask"], _ = expandSystemVirtualWanLinkServiceTosMask(d, i["tos_mask"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tos-mask"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandSystemVirtualWanLinkServiceProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-port"], _ = expandSystemVirtualWanLinkServiceStartPort(d, i["start_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["start-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-port"], _ = expandSystemVirtualWanLinkServiceEndPort(d, i["end_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["end-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route-tag"], _ = expandSystemVirtualWanLinkServiceRouteTag(d, i["route_tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route-tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandSystemVirtualWanLinkServiceDst(d, i["dst"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-negate"], _ = expandSystemVirtualWanLinkServiceDstNegate(d, i["dst_negate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst-negate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src"], _ = expandSystemVirtualWanLinkServiceSrc(d, i["src"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["src"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst6"], _ = expandSystemVirtualWanLinkServiceDst6(d, i["dst6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src6"], _ = expandSystemVirtualWanLinkServiceSrc6(d, i["src6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["src6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-negate"], _ = expandSystemVirtualWanLinkServiceSrcNegate(d, i["src_negate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["src-negate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["users"], _ = expandSystemVirtualWanLinkServiceUsers(d, i["users"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["users"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["groups"], _ = expandSystemVirtualWanLinkServiceGroups(d, i["groups"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["groups"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service"], _ = expandSystemVirtualWanLinkServiceInternetService(d, i["internet_service"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-custom"], _ = expandSystemVirtualWanLinkServiceInternetServiceCustom(d, i["internet_service_custom"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-custom"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-custom-group"], _ = expandSystemVirtualWanLinkServiceInternetServiceCustomGroup(d, i["internet_service_custom_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-custom-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-name"], _ = expandSystemVirtualWanLinkServiceInternetServiceName(d, i["internet_service_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-name"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-id"], _ = expandSystemVirtualWanLinkServiceInternetServiceId(d, i["internet_service_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-id"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-group"], _ = expandSystemVirtualWanLinkServiceInternetServiceGroup(d, i["internet_service_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-app-ctrl"], _ = expandSystemVirtualWanLinkServiceInternetServiceAppCtrl(d, i["internet_service_app_ctrl"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-app-ctrl-group"], _ = expandSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d, i["internet_service_app_ctrl_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-app-ctrl-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-ctrl"], _ = expandSystemVirtualWanLinkServiceInternetServiceCtrl(d, i["internet_service_ctrl"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-ctrl"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["internet-service-ctrl-group"], _ = expandSystemVirtualWanLinkServiceInternetServiceCtrlGroup(d, i["internet_service_ctrl_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["internet-service-ctrl-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandSystemVirtualWanLinkServiceHealthCheck(d, i["health_check"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["health-check"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["link-cost-factor"], _ = expandSystemVirtualWanLinkServiceLinkCostFactor(d, i["link_cost_factor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["link-cost-factor"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["packet-loss-weight"], _ = expandSystemVirtualWanLinkServicePacketLossWeight(d, i["packet_loss_weight"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["packet-loss-weight"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["latency-weight"], _ = expandSystemVirtualWanLinkServiceLatencyWeight(d, i["latency_weight"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["latency-weight"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["jitter-weight"], _ = expandSystemVirtualWanLinkServiceJitterWeight(d, i["jitter_weight"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["jitter-weight"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bandwidth-weight"], _ = expandSystemVirtualWanLinkServiceBandwidthWeight(d, i["bandwidth_weight"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bandwidth-weight"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["link-cost-threshold"], _ = expandSystemVirtualWanLinkServiceLinkCostThreshold(d, i["link_cost_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["link-cost-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hold-down-time"], _ = expandSystemVirtualWanLinkServiceHoldDownTime(d, i["hold_down_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["hold-down-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp-forward"], _ = expandSystemVirtualWanLinkServiceDscpForward(d, i["dscp_forward"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dscp-forward"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp-reverse"], _ = expandSystemVirtualWanLinkServiceDscpReverse(d, i["dscp_reverse"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dscp-reverse"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp-forward-tag"], _ = expandSystemVirtualWanLinkServiceDscpForwardTag(d, i["dscp_forward_tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dscp-forward-tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp-reverse-tag"], _ = expandSystemVirtualWanLinkServiceDscpReverseTag(d, i["dscp_reverse_tag"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dscp-reverse-tag"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla"], _ = expandSystemVirtualWanLinkServiceSla(d, i["sla"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority-members"], _ = expandSystemVirtualWanLinkServicePriorityMembers(d, i["priority_members"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["priority-members"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSystemVirtualWanLinkServiceStatus(d, i["status"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["status"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway"], _ = expandSystemVirtualWanLinkServiceGateway(d, i["gateway"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["gateway"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["default"], _ = expandSystemVirtualWanLinkServiceDefault(d, i["default"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["default"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sla-compare-method"], _ = expandSystemVirtualWanLinkServiceSlaCompareMethod(d, i["sla_compare_method"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sla-compare-method"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceInputDeviceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInputDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceStandaloneAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceQualityLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceTosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceRouteTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceDstName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDstNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceSrcName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceDst6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceDst6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceSrc6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceSrc6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceUsersName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceGroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceInternetServiceCustomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceInternetServiceCustomGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceInternetServiceNameName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandSystemVirtualWanLinkServiceInternetServiceIdId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceInternetServiceGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandSystemVirtualWanLinkServiceInternetServiceAppCtrlId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceAppCtrlId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceAppCtrlGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemVirtualWanLinkServiceInternetServiceCtrlId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCtrlId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemVirtualWanLinkServiceInternetServiceCtrlGroupName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceInternetServiceCtrlGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		"=": []string{"6.4.0"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); versionMatch {
		vx := fmt.Sprintf("%v", v)
		vxx := strings.Split(vx, ", ")

		tmps := make([]map[string]interface{}, 0, len(vxx))

		for _, xv := range vxx {
			xtmp := make(map[string]interface{})
			xtmp["name"] = xv

			tmps = append(tmps, xtmp)
		}
		return tmps, nil
	} else {
		return v, nil
	}
}

func expandSystemVirtualWanLinkServiceLinkCostFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServicePacketLossWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceLatencyWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceJitterWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceBandwidthWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceLinkCostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceHoldDownTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDscpForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDscpReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDscpForwardTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDscpReverseTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["health-check"], _ = expandSystemVirtualWanLinkServiceSlaHealthCheck(d, i["health_check"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["health-check"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemVirtualWanLinkServiceSlaId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServiceSlaHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["seq-num"], _ = expandSystemVirtualWanLinkServicePriorityMembersSeqNum(d, i["seq_num"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWanLinkServicePriorityMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWanLinkServiceSlaCompareMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVirtualWanLink(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemVirtualWanLinkStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	} else if d.HasChange("status") {
		obj["status"] = nil
	}

	if v, ok := d.GetOk("load_balance_mode"); ok {
		if setArgNil {
			obj["load-balance-mode"] = nil
		} else {
			t, err := expandSystemVirtualWanLinkLoadBalanceMode(d, v, "load_balance_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["load-balance-mode"] = t
			}
		}
	} else if d.HasChange("load_balance_mode") {
		obj["load-balance-mode"] = nil
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok {
		if setArgNil {
			obj["neighbor-hold-down"] = nil
		} else {
			t, err := expandSystemVirtualWanLinkNeighborHoldDown(d, v, "neighbor_hold_down", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor-hold-down"] = t
			}
		}
	} else if d.HasChange("neighbor_hold_down") {
		obj["neighbor-hold-down"] = nil
	}

	if v, ok := d.GetOkExists("neighbor_hold_down_time"); ok {
		if setArgNil {
			obj["neighbor-hold-down-time"] = nil
		} else {
			t, err := expandSystemVirtualWanLinkNeighborHoldDownTime(d, v, "neighbor_hold_down_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor-hold-down-time"] = t
			}
		}
	} else if d.HasChange("neighbor_hold_down_time") {
		obj["neighbor-hold-down-time"] = nil
	}

	if v, ok := d.GetOkExists("neighbor_hold_boot_time"); ok {
		if setArgNil {
			obj["neighbor-hold-boot-time"] = nil
		} else {
			t, err := expandSystemVirtualWanLinkNeighborHoldBootTime(d, v, "neighbor_hold_boot_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor-hold-boot-time"] = t
			}
		}
	} else if d.HasChange("neighbor_hold_boot_time") {
		obj["neighbor-hold-boot-time"] = nil
	}

	if v, ok := d.GetOk("fail_detect"); ok {
		if setArgNil {
			obj["fail-detect"] = nil
		} else {
			t, err := expandSystemVirtualWanLinkFailDetect(d, v, "fail_detect", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fail-detect"] = t
			}
		}
	} else if d.HasChange("fail_detect") {
		obj["fail-detect"] = nil
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok || d.HasChange("fail_alert_interfaces") {
		if setArgNil {
			obj["fail-alert-interfaces"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVirtualWanLinkFailAlertInterfaces(d, v, "fail_alert_interfaces", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fail-alert-interfaces"] = t
			}
		}
	}

	if v, ok := d.GetOk("zone"); ok || d.HasChange("zone") {
		if setArgNil {
			obj["zone"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVirtualWanLinkZone(d, v, "zone", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["zone"] = t
			}
		}
	}

	if v, ok := d.GetOk("members"); ok || d.HasChange("members") {
		if setArgNil {
			obj["members"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVirtualWanLinkMembers(d, v, "members", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["members"] = t
			}
		}
	}

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		if setArgNil {
			obj["health-check"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVirtualWanLinkHealthCheck(d, v, "health_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["health-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("neighbor"); ok || d.HasChange("neighbor") {
		if setArgNil {
			obj["neighbor"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVirtualWanLinkNeighbor(d, v, "neighbor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		if setArgNil {
			obj["service"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemVirtualWanLinkService(d, v, "service", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["service"] = t
			}
		}
	}

	return &obj, nil
}
