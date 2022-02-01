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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSdwan() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdwanUpdate,
		Read:   resourceSystemSdwanRead,
		Update: resourceSystemSdwanUpdate,
		Delete: resourceSystemSdwanDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"load_balance_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speedtest_bypass_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duplication_max_num": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 4),
				Optional:     true,
				Computed:     true,
			},
			"neighbor_hold_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor_hold_down_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),
				Optional:     true,
				Computed:     true,
			},
			"neighbor_hold_boot_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000000),
				Optional:     true,
				Computed:     true,
			},
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_alert_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
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
							Computed:     true,
						},
						"service_sla_tie_break": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							ValidateFunc: validation.IntBetween(0, 512),
							Optional:     true,
							Computed:     true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"zone": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"priority6": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"spillover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16776000),
							Optional:     true,
							Computed:     true,
						},
						"ingress_spillover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 16776000),
							Optional:     true,
							Computed:     true,
						},
						"volume_ratio": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed:     true,
						},
						"probe_packets": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"system_dns": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"detect_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"quality_measured_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
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
							Computed:     true,
						},
						"ha_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 50),
							Optional:     true,
							Computed:     true,
						},
						"ftp_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ftp_file": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 254),
							Optional:     true,
							Computed:     true,
						},
						"http_get": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
							Computed:     true,
						},
						"http_agent": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
							Computed:     true,
						},
						"http_match": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
							Computed:     true,
						},
						"dns_request_domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"dns_match_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(500, 3600000),
							Optional:     true,
							Computed:     true,
						},
						"probe_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(500, 3600000),
							Optional:     true,
							Computed:     true,
						},
						"failtime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),
							Optional:     true,
							Computed:     true,
						},
						"recoverytime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),
							Optional:     true,
							Computed:     true,
						},
						"probe_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 30),
							Optional:     true,
							Computed:     true,
						},
						"diffservcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_cascade_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"update_static_route": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sla_fail_log_period": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
							Optional:     true,
							Computed:     true,
						},
						"sla_pass_log_period": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
							Optional:     true,
							Computed:     true,
						},
						"threshold_warning_packetloss": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"threshold_alert_packetloss": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"threshold_warning_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"threshold_alert_latency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"threshold_warning_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"threshold_alert_jitter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"members": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
										Computed:     true,
									},
									"link_cost_factor": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"latency_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"jitter_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 10000000),
										Optional:     true,
										Computed:     true,
									},
									"packetloss_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 100),
										Optional:     true,
										Computed:     true,
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
							Computed:     true,
						},
						"member": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"health_check": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sla_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
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
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"input_device": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"input_device_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"minimum_sla_meet_members": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"hash_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"role": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"standalone_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quality_link": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tos_mask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"start_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"end_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"route_tag": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dst_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dst6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"src6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"src_negate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"users": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"internet_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"internet_service_custom": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"internet_service_custom_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"internet_service_name": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"internet_service_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"internet_service_app_ctrl": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"internet_service_app_ctrl_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
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
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"link_cost_factor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_loss_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"latency_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"jitter_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"bandwidth_weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"link_cost_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"hold_down_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
							Computed:     true,
						},
						"dscp_forward": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_reverse": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_forward_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dscp_reverse_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
										Computed:     true,
									},
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"priority_members": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"seq_num": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"priority_zone": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sla_compare_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tie_break": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"use_shortcut_sla": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"passive_measurement": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"duplication": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"service_id": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"srcaddr6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"srcintf": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dstintf": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"service": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"packet_duplication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"packet_de_duplication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSdwanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSdwan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwan resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSdwan(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdwan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdwan")
	}

	return resourceSystemSdwanRead(d, m)
}

func resourceSystemSdwanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSdwan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdwan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdwanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemSdwan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdwan(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdwan resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdwanStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanLoadBalanceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanSpeedtestBypassRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationMaxNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHoldDown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHoldDownTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHoldBootTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanFailDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanFailAlertInterfacesName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdwanFailAlertInterfacesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanZone(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanZoneName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := i["service-sla-tie-break"]; ok {

			tmp["service_sla_tie_break"] = flattenSystemSdwanZoneServiceSlaTieBreak(i["service-sla-tie-break"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdwanZoneName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanZoneServiceSlaTieBreak(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {

			tmp["seq_num"] = flattenSystemSdwanMembersSeqNum(i["seq-num"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenSystemSdwanMembersInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := i["zone"]; ok {

			tmp["zone"] = flattenSystemSdwanMembersZone(i["zone"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {

			tmp["gateway"] = flattenSystemSdwanMembersGateway(i["gateway"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := i["source"]; ok {

			tmp["source"] = flattenSystemSdwanMembersSource(i["source"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {

			tmp["gateway6"] = flattenSystemSdwanMembersGateway6(i["gateway6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := i["source6"]; ok {

			tmp["source6"] = flattenSystemSdwanMembersSource6(i["source6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := i["cost"]; ok {

			tmp["cost"] = flattenSystemSdwanMembersCost(i["cost"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenSystemSdwanMembersWeight(i["weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {

			tmp["priority"] = flattenSystemSdwanMembersPriority(i["priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := i["priority6"]; ok {

			tmp["priority6"] = flattenSystemSdwanMembersPriority6(i["priority6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := i["spillover-threshold"]; ok {

			tmp["spillover_threshold"] = flattenSystemSdwanMembersSpilloverThreshold(i["spillover-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := i["ingress-spillover-threshold"]; ok {

			tmp["ingress_spillover_threshold"] = flattenSystemSdwanMembersIngressSpilloverThreshold(i["ingress-spillover-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := i["volume-ratio"]; ok {

			tmp["volume_ratio"] = flattenSystemSdwanMembersVolumeRatio(i["volume-ratio"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSystemSdwanMembersStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {

			tmp["comment"] = flattenSystemSdwanMembersComment(i["comment"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "seq_num", d)
	return result
}

func flattenSystemSdwanMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersSource6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersPriority6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersVolumeRatio(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanMembersComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanHealthCheckName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := i["probe-packets"]; ok {

			tmp["probe_packets"] = flattenSystemSdwanHealthCheckProbePackets(i["probe-packets"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {

			tmp["addr_mode"] = flattenSystemSdwanHealthCheckAddrMode(i["addr-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := i["system-dns"]; ok {

			tmp["system_dns"] = flattenSystemSdwanHealthCheckSystemDns(i["system-dns"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {

			tmp["server"] = flattenSystemSdwanHealthCheckServer(i["server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := i["detect-mode"]; ok {

			tmp["detect_mode"] = flattenSystemSdwanHealthCheckDetectMode(i["detect-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenSystemSdwanHealthCheckProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenSystemSdwanHealthCheckPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := i["quality-measured-method"]; ok {

			tmp["quality_measured_method"] = flattenSystemSdwanHealthCheckQualityMeasuredMethod(i["quality-measured-method"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := i["security-mode"]; ok {

			tmp["security_mode"] = flattenSystemSdwanHealthCheckSecurityMode(i["security-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {

			tmp["user"] = flattenSystemSdwanHealthCheckUser(i["user"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {

			tmp["password"] = flattenSystemSdwanHealthCheckPassword(i["password"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := i["packet-size"]; ok {

			tmp["packet_size"] = flattenSystemSdwanHealthCheckPacketSize(i["packet-size"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {

			tmp["ha_priority"] = flattenSystemSdwanHealthCheckHaPriority(i["ha-priority"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := i["ftp-mode"]; ok {

			tmp["ftp_mode"] = flattenSystemSdwanHealthCheckFtpMode(i["ftp-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := i["ftp-file"]; ok {

			tmp["ftp_file"] = flattenSystemSdwanHealthCheckFtpFile(i["ftp-file"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := i["http-get"]; ok {

			tmp["http_get"] = flattenSystemSdwanHealthCheckHttpGet(i["http-get"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := i["http-agent"]; ok {

			tmp["http_agent"] = flattenSystemSdwanHealthCheckHttpAgent(i["http-agent"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := i["http-match"]; ok {

			tmp["http_match"] = flattenSystemSdwanHealthCheckHttpMatch(i["http-match"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := i["dns-request-domain"]; ok {

			tmp["dns_request_domain"] = flattenSystemSdwanHealthCheckDnsRequestDomain(i["dns-request-domain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := i["dns-match-ip"]; ok {

			tmp["dns_match_ip"] = flattenSystemSdwanHealthCheckDnsMatchIp(i["dns-match-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := i["interval"]; ok {

			tmp["interval"] = flattenSystemSdwanHealthCheckInterval(i["interval"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := i["probe-timeout"]; ok {

			tmp["probe_timeout"] = flattenSystemSdwanHealthCheckProbeTimeout(i["probe-timeout"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := i["failtime"]; ok {

			tmp["failtime"] = flattenSystemSdwanHealthCheckFailtime(i["failtime"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := i["recoverytime"]; ok {

			tmp["recoverytime"] = flattenSystemSdwanHealthCheckRecoverytime(i["recoverytime"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := i["probe-count"]; ok {

			tmp["probe_count"] = flattenSystemSdwanHealthCheckProbeCount(i["probe-count"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := i["diffservcode"]; ok {

			tmp["diffservcode"] = flattenSystemSdwanHealthCheckDiffservcode(i["diffservcode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := i["update-cascade-interface"]; ok {

			tmp["update_cascade_interface"] = flattenSystemSdwanHealthCheckUpdateCascadeInterface(i["update-cascade-interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := i["update-static-route"]; ok {

			tmp["update_static_route"] = flattenSystemSdwanHealthCheckUpdateStaticRoute(i["update-static-route"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := i["sla-fail-log-period"]; ok {

			tmp["sla_fail_log_period"] = flattenSystemSdwanHealthCheckSlaFailLogPeriod(i["sla-fail-log-period"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := i["sla-pass-log-period"]; ok {

			tmp["sla_pass_log_period"] = flattenSystemSdwanHealthCheckSlaPassLogPeriod(i["sla-pass-log-period"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := i["threshold-warning-packetloss"]; ok {

			tmp["threshold_warning_packetloss"] = flattenSystemSdwanHealthCheckThresholdWarningPacketloss(i["threshold-warning-packetloss"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := i["threshold-alert-packetloss"]; ok {

			tmp["threshold_alert_packetloss"] = flattenSystemSdwanHealthCheckThresholdAlertPacketloss(i["threshold-alert-packetloss"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := i["threshold-warning-latency"]; ok {

			tmp["threshold_warning_latency"] = flattenSystemSdwanHealthCheckThresholdWarningLatency(i["threshold-warning-latency"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := i["threshold-alert-latency"]; ok {

			tmp["threshold_alert_latency"] = flattenSystemSdwanHealthCheckThresholdAlertLatency(i["threshold-alert-latency"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := i["threshold-warning-jitter"]; ok {

			tmp["threshold_warning_jitter"] = flattenSystemSdwanHealthCheckThresholdWarningJitter(i["threshold-warning-jitter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := i["threshold-alert-jitter"]; ok {

			tmp["threshold_alert_jitter"] = flattenSystemSdwanHealthCheckThresholdAlertJitter(i["threshold-alert-jitter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := i["members"]; ok {

			tmp["members"] = flattenSystemSdwanHealthCheckMembers(i["members"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {

			tmp["sla"] = flattenSystemSdwanHealthCheckSla(i["sla"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdwanHealthCheckName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbePackets(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSystemDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDetectMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckQualityMeasuredMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSecurityMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckPacketSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHaPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFtpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFtpFile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpGet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpAgent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckHttpMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDnsMatchIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbeTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckFailtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckRecoverytime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckProbeCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckDiffservcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaFailLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPassLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningPacketloss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertPacketloss(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningLatency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertLatency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdWarningJitter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckThresholdAlertJitter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {

			tmp["seq_num"] = flattenSystemSdwanHealthCheckMembersSeqNum(i["seq-num"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemSdwanHealthCheckSlaId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {

			tmp["link_cost_factor"] = flattenSystemSdwanHealthCheckSlaLinkCostFactor(i["link-cost-factor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := i["latency-threshold"]; ok {

			tmp["latency_threshold"] = flattenSystemSdwanHealthCheckSlaLatencyThreshold(i["latency-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := i["jitter-threshold"]; ok {

			tmp["jitter_threshold"] = flattenSystemSdwanHealthCheckSlaJitterThreshold(i["jitter-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := i["packetloss-threshold"]; ok {

			tmp["packetloss_threshold"] = flattenSystemSdwanHealthCheckSlaPacketlossThreshold(i["packetloss-threshold"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemSdwanNeighborIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := i["member"]; ok {

			tmp["member"] = flattenSystemSdwanNeighborMember(i["member"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {

			tmp["mode"] = flattenSystemSdwanNeighborMode(i["mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {

			tmp["role"] = flattenSystemSdwanNeighborRole(i["role"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenSystemSdwanNeighborHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := i["sla-id"]; ok {

			tmp["sla_id"] = flattenSystemSdwanNeighborSlaId(i["sla-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenSystemSdwanNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborMember(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanNeighborSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemSdwanServiceId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemSdwanServiceName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := i["addr-mode"]; ok {

			tmp["addr_mode"] = flattenSystemSdwanServiceAddrMode(i["addr-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := i["input-device"]; ok {

			tmp["input_device"] = flattenSystemSdwanServiceInputDevice(i["input-device"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := i["input-device-negate"]; ok {

			tmp["input_device_negate"] = flattenSystemSdwanServiceInputDeviceNegate(i["input-device-negate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := i["mode"]; ok {

			tmp["mode"] = flattenSystemSdwanServiceMode(i["mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := i["minimum-sla-meet-members"]; ok {

			tmp["minimum_sla_meet_members"] = flattenSystemSdwanServiceMinimumSlaMeetMembers(i["minimum-sla-meet-members"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := i["hash-mode"]; ok {

			tmp["hash_mode"] = flattenSystemSdwanServiceHashMode(i["hash-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {

			tmp["role"] = flattenSystemSdwanServiceRole(i["role"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := i["standalone-action"]; ok {

			tmp["standalone_action"] = flattenSystemSdwanServiceStandaloneAction(i["standalone-action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := i["quality-link"]; ok {

			tmp["quality_link"] = flattenSystemSdwanServiceQualityLink(i["quality-link"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := i["tos"]; ok {

			tmp["tos"] = flattenSystemSdwanServiceTos(i["tos"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := i["tos-mask"]; ok {

			tmp["tos_mask"] = flattenSystemSdwanServiceTosMask(i["tos-mask"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {

			tmp["protocol"] = flattenSystemSdwanServiceProtocol(i["protocol"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := i["start-port"]; ok {

			tmp["start_port"] = flattenSystemSdwanServiceStartPort(i["start-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := i["end-port"]; ok {

			tmp["end_port"] = flattenSystemSdwanServiceEndPort(i["end-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := i["route-tag"]; ok {

			tmp["route_tag"] = flattenSystemSdwanServiceRouteTag(i["route-tag"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {

			tmp["dst"] = flattenSystemSdwanServiceDst(i["dst"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := i["dst-negate"]; ok {

			tmp["dst_negate"] = flattenSystemSdwanServiceDstNegate(i["dst-negate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {

			tmp["src"] = flattenSystemSdwanServiceSrc(i["src"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := i["dst6"]; ok {

			tmp["dst6"] = flattenSystemSdwanServiceDst6(i["dst6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := i["src6"]; ok {

			tmp["src6"] = flattenSystemSdwanServiceSrc6(i["src6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := i["src-negate"]; ok {

			tmp["src_negate"] = flattenSystemSdwanServiceSrcNegate(i["src-negate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := i["users"]; ok {

			tmp["users"] = flattenSystemSdwanServiceUsers(i["users"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := i["groups"]; ok {

			tmp["groups"] = flattenSystemSdwanServiceGroups(i["groups"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := i["internet-service"]; ok {

			tmp["internet_service"] = flattenSystemSdwanServiceInternetService(i["internet-service"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := i["internet-service-custom"]; ok {

			tmp["internet_service_custom"] = flattenSystemSdwanServiceInternetServiceCustom(i["internet-service-custom"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := i["internet-service-custom-group"]; ok {

			tmp["internet_service_custom_group"] = flattenSystemSdwanServiceInternetServiceCustomGroup(i["internet-service-custom-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := i["internet-service-name"]; ok {

			tmp["internet_service_name"] = flattenSystemSdwanServiceInternetServiceName(i["internet-service-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := i["internet-service-group"]; ok {

			tmp["internet_service_group"] = flattenSystemSdwanServiceInternetServiceGroup(i["internet-service-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := i["internet-service-app-ctrl"]; ok {

			tmp["internet_service_app_ctrl"] = flattenSystemSdwanServiceInternetServiceAppCtrl(i["internet-service-app-ctrl"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := i["internet-service-app-ctrl-group"]; ok {

			tmp["internet_service_app_ctrl_group"] = flattenSystemSdwanServiceInternetServiceAppCtrlGroup(i["internet-service-app-ctrl-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenSystemSdwanServiceHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := i["link-cost-factor"]; ok {

			tmp["link_cost_factor"] = flattenSystemSdwanServiceLinkCostFactor(i["link-cost-factor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := i["packet-loss-weight"]; ok {

			tmp["packet_loss_weight"] = flattenSystemSdwanServicePacketLossWeight(i["packet-loss-weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := i["latency-weight"]; ok {

			tmp["latency_weight"] = flattenSystemSdwanServiceLatencyWeight(i["latency-weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := i["jitter-weight"]; ok {

			tmp["jitter_weight"] = flattenSystemSdwanServiceJitterWeight(i["jitter-weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := i["bandwidth-weight"]; ok {

			tmp["bandwidth_weight"] = flattenSystemSdwanServiceBandwidthWeight(i["bandwidth-weight"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := i["link-cost-threshold"]; ok {

			tmp["link_cost_threshold"] = flattenSystemSdwanServiceLinkCostThreshold(i["link-cost-threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := i["hold-down-time"]; ok {

			tmp["hold_down_time"] = flattenSystemSdwanServiceHoldDownTime(i["hold-down-time"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := i["dscp-forward"]; ok {

			tmp["dscp_forward"] = flattenSystemSdwanServiceDscpForward(i["dscp-forward"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := i["dscp-reverse"]; ok {

			tmp["dscp_reverse"] = flattenSystemSdwanServiceDscpReverse(i["dscp-reverse"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := i["dscp-forward-tag"]; ok {

			tmp["dscp_forward_tag"] = flattenSystemSdwanServiceDscpForwardTag(i["dscp-forward-tag"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := i["dscp-reverse-tag"]; ok {

			tmp["dscp_reverse_tag"] = flattenSystemSdwanServiceDscpReverseTag(i["dscp-reverse-tag"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := i["sla"]; ok {

			tmp["sla"] = flattenSystemSdwanServiceSla(i["sla"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := i["priority-members"]; ok {

			tmp["priority_members"] = flattenSystemSdwanServicePriorityMembers(i["priority-members"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := i["priority-zone"]; ok {

			tmp["priority_zone"] = flattenSystemSdwanServicePriorityZone(i["priority-zone"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSystemSdwanServiceStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {

			tmp["gateway"] = flattenSystemSdwanServiceGateway(i["gateway"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := i["default"]; ok {

			tmp["default"] = flattenSystemSdwanServiceDefault(i["default"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := i["sla-compare-method"]; ok {

			tmp["sla_compare_method"] = flattenSystemSdwanServiceSlaCompareMethod(i["sla-compare-method"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := i["tie-break"]; ok {

			tmp["tie_break"] = flattenSystemSdwanServiceTieBreak(i["tie-break"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := i["use-shortcut-sla"]; ok {

			tmp["use_shortcut_sla"] = flattenSystemSdwanServiceUseShortcutSla(i["use-shortcut-sla"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := i["passive-measurement"]; ok {

			tmp["passive_measurement"] = flattenSystemSdwanServicePassiveMeasurement(i["passive-measurement"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSdwanServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceInputDeviceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInputDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceHashMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceStandaloneAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceQualityLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceTosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceRouteTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceDstName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDstNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrc(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceSrcName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceSrcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDst6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceDst6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceDst6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrc6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceSrc6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceSrc6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceSrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceUsersName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceInternetServiceCustomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceCustomGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceInternetServiceCustomGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInternetServiceCustomGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceInternetServiceNameName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceInternetServiceGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInternetServiceGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceAppCtrl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemSdwanServiceInternetServiceAppCtrlId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInternetServiceAppCtrlId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceInternetServiceAppCtrlGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceInternetServiceAppCtrlGroupName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceInternetServiceAppCtrlGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServiceHealthCheckName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceHealthCheckName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceLinkCostFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServicePacketLossWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceLatencyWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceJitterWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceBandwidthWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceLinkCostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceHoldDownTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpForward(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpForwardTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDscpReverseTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceSla(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := i["health-check"]; ok {

			tmp["health_check"] = flattenSystemSdwanServiceSlaHealthCheck(i["health-check"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemSdwanServiceSlaId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServiceSlaHealthCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceSlaId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServicePriorityMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := i["seq-num"]; ok {

			tmp["seq_num"] = flattenSystemSdwanServicePriorityMembersSeqNum(i["seq-num"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServicePriorityMembersSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServicePriorityZone(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanServicePriorityZoneName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanServicePriorityZoneName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceSlaCompareMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceTieBreak(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServiceUseShortcutSla(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanServicePassiveMeasurement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemSdwanDuplicationId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := i["service-id"]; ok {

			tmp["service_id"] = flattenSystemSdwanDuplicationServiceId(i["service-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {

			tmp["srcaddr"] = flattenSystemSdwanDuplicationSrcaddr(i["srcaddr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {

			tmp["dstaddr"] = flattenSystemSdwanDuplicationDstaddr(i["dstaddr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := i["srcaddr6"]; ok {

			tmp["srcaddr6"] = flattenSystemSdwanDuplicationSrcaddr6(i["srcaddr6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {

			tmp["dstaddr6"] = flattenSystemSdwanDuplicationDstaddr6(i["dstaddr6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := i["srcintf"]; ok {

			tmp["srcintf"] = flattenSystemSdwanDuplicationSrcintf(i["srcintf"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := i["dstintf"]; ok {

			tmp["dstintf"] = flattenSystemSdwanDuplicationDstintf(i["dstintf"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := i["service"]; ok {

			tmp["service"] = flattenSystemSdwanDuplicationService(i["service"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := i["packet-duplication"]; ok {

			tmp["packet_duplication"] = flattenSystemSdwanDuplicationPacketDuplication(i["packet-duplication"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := i["packet-de-duplication"]; ok {

			tmp["packet_de_duplication"] = flattenSystemSdwanDuplicationPacketDeDuplication(i["packet-de-duplication"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSdwanDuplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemSdwanDuplicationServiceIdId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationServiceIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationSrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationSrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationDstaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationSrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationSrcaddr6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationSrcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationDstaddr6Name(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationDstaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationSrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationSrcintfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationSrcintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationDstintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationDstintfName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationDstintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSystemSdwanDuplicationServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemSdwanDuplicationServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationPacketDuplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdwanDuplicationPacketDeDuplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSdwan(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemSdwanStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("load_balance_mode", flattenSystemSdwanLoadBalanceMode(o["load-balance-mode"], d, "load_balance_mode", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-mode"]) {
			return fmt.Errorf("Error reading load_balance_mode: %v", err)
		}
	}

	if err = d.Set("speedtest_bypass_routing", flattenSystemSdwanSpeedtestBypassRouting(o["speedtest-bypass-routing"], d, "speedtest_bypass_routing", sv)); err != nil {
		if !fortiAPIPatch(o["speedtest-bypass-routing"]) {
			return fmt.Errorf("Error reading speedtest_bypass_routing: %v", err)
		}
	}

	if err = d.Set("duplication_max_num", flattenSystemSdwanDuplicationMaxNum(o["duplication-max-num"], d, "duplication_max_num", sv)); err != nil {
		if !fortiAPIPatch(o["duplication-max-num"]) {
			return fmt.Errorf("Error reading duplication_max_num: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down", flattenSystemSdwanNeighborHoldDown(o["neighbor-hold-down"], d, "neighbor_hold_down", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-hold-down"]) {
			return fmt.Errorf("Error reading neighbor_hold_down: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_down_time", flattenSystemSdwanNeighborHoldDownTime(o["neighbor-hold-down-time"], d, "neighbor_hold_down_time", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-hold-down-time"]) {
			return fmt.Errorf("Error reading neighbor_hold_down_time: %v", err)
		}
	}

	if err = d.Set("neighbor_hold_boot_time", flattenSystemSdwanNeighborHoldBootTime(o["neighbor-hold-boot-time"], d, "neighbor_hold_boot_time", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-hold-boot-time"]) {
			return fmt.Errorf("Error reading neighbor_hold_boot_time: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenSystemSdwanFailDetect(o["fail-detect"], d, "fail_detect", sv)); err != nil {
		if !fortiAPIPatch(o["fail-detect"]) {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fail_alert_interfaces", flattenSystemSdwanFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces", sv)); err != nil {
			if !fortiAPIPatch(o["fail-alert-interfaces"]) {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fail_alert_interfaces"); ok {
			if err = d.Set("fail_alert_interfaces", flattenSystemSdwanFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces", sv)); err != nil {
				if !fortiAPIPatch(o["fail-alert-interfaces"]) {
					return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("zone", flattenSystemSdwanZone(o["zone"], d, "zone", sv)); err != nil {
			if !fortiAPIPatch(o["zone"]) {
				return fmt.Errorf("Error reading zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("zone"); ok {
			if err = d.Set("zone", flattenSystemSdwanZone(o["zone"], d, "zone", sv)); err != nil {
				if !fortiAPIPatch(o["zone"]) {
					return fmt.Errorf("Error reading zone: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("members", flattenSystemSdwanMembers(o["members"], d, "members", sv)); err != nil {
			if !fortiAPIPatch(o["members"]) {
				return fmt.Errorf("Error reading members: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("members"); ok {
			if err = d.Set("members", flattenSystemSdwanMembers(o["members"], d, "members", sv)); err != nil {
				if !fortiAPIPatch(o["members"]) {
					return fmt.Errorf("Error reading members: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("health_check", flattenSystemSdwanHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
			if !fortiAPIPatch(o["health-check"]) {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("health_check"); ok {
			if err = d.Set("health_check", flattenSystemSdwanHealthCheck(o["health-check"], d, "health_check", sv)); err != nil {
				if !fortiAPIPatch(o["health-check"]) {
					return fmt.Errorf("Error reading health_check: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenSystemSdwanNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenSystemSdwanNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenSystemSdwanService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenSystemSdwanService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("duplication", flattenSystemSdwanDuplication(o["duplication"], d, "duplication", sv)); err != nil {
			if !fortiAPIPatch(o["duplication"]) {
				return fmt.Errorf("Error reading duplication: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("duplication"); ok {
			if err = d.Set("duplication", flattenSystemSdwanDuplication(o["duplication"], d, "duplication", sv)); err != nil {
				if !fortiAPIPatch(o["duplication"]) {
					return fmt.Errorf("Error reading duplication: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSdwanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSdwanStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanLoadBalanceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanSpeedtestBypassRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationMaxNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHoldDown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHoldDownTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHoldBootTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanFailDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanFailAlertInterfacesName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanFailAlertInterfacesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanZoneName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_sla_tie_break"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["service-sla-tie-break"], _ = expandSystemSdwanZoneServiceSlaTieBreak(d, i["service_sla_tie_break"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanZoneName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanZoneServiceSlaTieBreak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seq-num"], _ = expandSystemSdwanMembersSeqNum(d, i["seq_num"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandSystemSdwanMembersInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "zone"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["zone"], _ = expandSystemSdwanMembersZone(d, i["zone"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["gateway"], _ = expandSystemSdwanMembersGateway(d, i["gateway"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source"], _ = expandSystemSdwanMembersSource(d, i["source"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["gateway6"], _ = expandSystemSdwanMembersGateway6(d, i["gateway6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source6"], _ = expandSystemSdwanMembersSource6(d, i["source6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cost"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["cost"], _ = expandSystemSdwanMembersCost(d, i["cost"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandSystemSdwanMembersWeight(d, i["weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority"], _ = expandSystemSdwanMembersPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority6"], _ = expandSystemSdwanMembersPriority6(d, i["priority6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["spillover-threshold"], _ = expandSystemSdwanMembersSpilloverThreshold(d, i["spillover_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ingress_spillover_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ingress-spillover-threshold"], _ = expandSystemSdwanMembersIngressSpilloverThreshold(d, i["ingress_spillover_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "volume_ratio"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["volume-ratio"], _ = expandSystemSdwanMembersVolumeRatio(d, i["volume_ratio"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSystemSdwanMembersStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["comment"], _ = expandSystemSdwanMembersComment(d, i["comment"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSource6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersPriority6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersVolumeRatio(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanMembersComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanHealthCheckName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_packets"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["probe-packets"], _ = expandSystemSdwanHealthCheckProbePackets(d, i["probe_packets"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-mode"], _ = expandSystemSdwanHealthCheckAddrMode(d, i["addr_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "system_dns"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["system-dns"], _ = expandSystemSdwanHealthCheckSystemDns(d, i["system_dns"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server"], _ = expandSystemSdwanHealthCheckServer(d, i["server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detect_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["detect-mode"], _ = expandSystemSdwanHealthCheckDetectMode(d, i["detect_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandSystemSdwanHealthCheckProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandSystemSdwanHealthCheckPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_measured_method"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quality-measured-method"], _ = expandSystemSdwanHealthCheckQualityMeasuredMethod(d, i["quality_measured_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["security-mode"], _ = expandSystemSdwanHealthCheckSecurityMode(d, i["security_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["user"], _ = expandSystemSdwanHealthCheckUser(d, i["user"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["password"], _ = expandSystemSdwanHealthCheckPassword(d, i["password"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_size"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packet-size"], _ = expandSystemSdwanHealthCheckPacketSize(d, i["packet_size"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ha-priority"], _ = expandSystemSdwanHealthCheckHaPriority(d, i["ha_priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ftp-mode"], _ = expandSystemSdwanHealthCheckFtpMode(d, i["ftp_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ftp_file"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ftp-file"], _ = expandSystemSdwanHealthCheckFtpFile(d, i["ftp_file"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_get"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-get"], _ = expandSystemSdwanHealthCheckHttpGet(d, i["http_get"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_agent"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-agent"], _ = expandSystemSdwanHealthCheckHttpAgent(d, i["http_agent"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_match"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-match"], _ = expandSystemSdwanHealthCheckHttpMatch(d, i["http_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_request_domain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dns-request-domain"], _ = expandSystemSdwanHealthCheckDnsRequestDomain(d, i["dns_request_domain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_match_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dns-match-ip"], _ = expandSystemSdwanHealthCheckDnsMatchIp(d, i["dns_match_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interval"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interval"], _ = expandSystemSdwanHealthCheckInterval(d, i["interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_timeout"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["probe-timeout"], _ = expandSystemSdwanHealthCheckProbeTimeout(d, i["probe_timeout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "failtime"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["failtime"], _ = expandSystemSdwanHealthCheckFailtime(d, i["failtime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "recoverytime"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["recoverytime"], _ = expandSystemSdwanHealthCheckRecoverytime(d, i["recoverytime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "probe_count"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["probe-count"], _ = expandSystemSdwanHealthCheckProbeCount(d, i["probe_count"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffservcode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["diffservcode"], _ = expandSystemSdwanHealthCheckDiffservcode(d, i["diffservcode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_cascade_interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["update-cascade-interface"], _ = expandSystemSdwanHealthCheckUpdateCascadeInterface(d, i["update_cascade_interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "update_static_route"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["update-static-route"], _ = expandSystemSdwanHealthCheckUpdateStaticRoute(d, i["update_static_route"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_fail_log_period"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sla-fail-log-period"], _ = expandSystemSdwanHealthCheckSlaFailLogPeriod(d, i["sla_fail_log_period"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_pass_log_period"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sla-pass-log-period"], _ = expandSystemSdwanHealthCheckSlaPassLogPeriod(d, i["sla_pass_log_period"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_packetloss"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold-warning-packetloss"], _ = expandSystemSdwanHealthCheckThresholdWarningPacketloss(d, i["threshold_warning_packetloss"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_packetloss"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold-alert-packetloss"], _ = expandSystemSdwanHealthCheckThresholdAlertPacketloss(d, i["threshold_alert_packetloss"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_latency"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold-warning-latency"], _ = expandSystemSdwanHealthCheckThresholdWarningLatency(d, i["threshold_warning_latency"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_latency"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold-alert-latency"], _ = expandSystemSdwanHealthCheckThresholdAlertLatency(d, i["threshold_alert_latency"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_warning_jitter"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold-warning-jitter"], _ = expandSystemSdwanHealthCheckThresholdWarningJitter(d, i["threshold_warning_jitter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold_alert_jitter"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold-alert-jitter"], _ = expandSystemSdwanHealthCheckThresholdAlertJitter(d, i["threshold_alert_jitter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["members"], _ = expandSystemSdwanHealthCheckMembers(d, i["members"], pre_append, sv)
		} else {
			tmp["members"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sla"], _ = expandSystemSdwanHealthCheckSla(d, i["sla"], pre_append, sv)
		} else {
			tmp["sla"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbePackets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSystemDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDetectMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckQualityMeasuredMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSecurityMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckPacketSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHaPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFtpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFtpFile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpGet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpAgent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckHttpMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDnsMatchIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbeTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckFailtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckRecoverytime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckProbeCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckDiffservcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaFailLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPassLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningPacketloss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertPacketloss(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningLatency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertLatency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdWarningJitter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckThresholdAlertJitter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seq-num"], _ = expandSystemSdwanHealthCheckMembersSeqNum(d, i["seq_num"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemSdwanHealthCheckSlaId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["link-cost-factor"], _ = expandSystemSdwanHealthCheckSlaLinkCostFactor(d, i["link_cost_factor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["latency-threshold"], _ = expandSystemSdwanHealthCheckSlaLatencyThreshold(d, i["latency_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["jitter-threshold"], _ = expandSystemSdwanHealthCheckSlaJitterThreshold(d, i["jitter_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packetloss_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packetloss-threshold"], _ = expandSystemSdwanHealthCheckSlaPacketlossThreshold(d, i["packetloss_threshold"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemSdwanNeighborIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "member"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["member"], _ = expandSystemSdwanNeighborMember(d, i["member"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mode"], _ = expandSystemSdwanNeighborMode(d, i["mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["role"], _ = expandSystemSdwanNeighborRole(d, i["role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandSystemSdwanNeighborHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sla-id"], _ = expandSystemSdwanNeighborSlaId(d, i["sla_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanNeighborSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemSdwanServiceId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemSdwanServiceName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-mode"], _ = expandSystemSdwanServiceAddrMode(d, i["addr_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["input-device"], _ = expandSystemSdwanServiceInputDevice(d, i["input_device"], pre_append, sv)
		} else {
			tmp["input-device"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input_device_negate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["input-device-negate"], _ = expandSystemSdwanServiceInputDeviceNegate(d, i["input_device_negate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mode"], _ = expandSystemSdwanServiceMode(d, i["mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minimum_sla_meet_members"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["minimum-sla-meet-members"], _ = expandSystemSdwanServiceMinimumSlaMeetMembers(d, i["minimum_sla_meet_members"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hash_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hash-mode"], _ = expandSystemSdwanServiceHashMode(d, i["hash_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["role"], _ = expandSystemSdwanServiceRole(d, i["role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "standalone_action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["standalone-action"], _ = expandSystemSdwanServiceStandaloneAction(d, i["standalone_action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quality_link"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quality-link"], _ = expandSystemSdwanServiceQualityLink(d, i["quality_link"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tos"], _ = expandSystemSdwanServiceTos(d, i["tos"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tos_mask"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tos-mask"], _ = expandSystemSdwanServiceTosMask(d, i["tos_mask"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["protocol"], _ = expandSystemSdwanServiceProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["start-port"], _ = expandSystemSdwanServiceStartPort(d, i["start_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["end-port"], _ = expandSystemSdwanServiceEndPort(d, i["end_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route_tag"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["route-tag"], _ = expandSystemSdwanServiceRouteTag(d, i["route_tag"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dst"], _ = expandSystemSdwanServiceDst(d, i["dst"], pre_append, sv)
		} else {
			tmp["dst"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_negate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dst-negate"], _ = expandSystemSdwanServiceDstNegate(d, i["dst_negate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src"], _ = expandSystemSdwanServiceSrc(d, i["src"], pre_append, sv)
		} else {
			tmp["src"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dst6"], _ = expandSystemSdwanServiceDst6(d, i["dst6"], pre_append, sv)
		} else {
			tmp["dst6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src6"], _ = expandSystemSdwanServiceSrc6(d, i["src6"], pre_append, sv)
		} else {
			tmp["src6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_negate"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["src-negate"], _ = expandSystemSdwanServiceSrcNegate(d, i["src_negate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "users"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["users"], _ = expandSystemSdwanServiceUsers(d, i["users"], pre_append, sv)
		} else {
			tmp["users"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "groups"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["groups"], _ = expandSystemSdwanServiceGroups(d, i["groups"], pre_append, sv)
		} else {
			tmp["groups"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service"], _ = expandSystemSdwanServiceInternetService(d, i["internet_service"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service-custom"], _ = expandSystemSdwanServiceInternetServiceCustom(d, i["internet_service_custom"], pre_append, sv)
		} else {
			tmp["internet-service-custom"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_custom_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service-custom-group"], _ = expandSystemSdwanServiceInternetServiceCustomGroup(d, i["internet_service_custom_group"], pre_append, sv)
		} else {
			tmp["internet-service-custom-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service-name"], _ = expandSystemSdwanServiceInternetServiceName(d, i["internet_service_name"], pre_append, sv)
		} else {
			tmp["internet-service-name"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service-group"], _ = expandSystemSdwanServiceInternetServiceGroup(d, i["internet_service_group"], pre_append, sv)
		} else {
			tmp["internet-service-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service-app-ctrl"], _ = expandSystemSdwanServiceInternetServiceAppCtrl(d, i["internet_service_app_ctrl"], pre_append, sv)
		} else {
			tmp["internet-service-app-ctrl"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "internet_service_app_ctrl_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["internet-service-app-ctrl-group"], _ = expandSystemSdwanServiceInternetServiceAppCtrlGroup(d, i["internet_service_app_ctrl_group"], pre_append, sv)
		} else {
			tmp["internet-service-app-ctrl-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandSystemSdwanServiceHealthCheck(d, i["health_check"], pre_append, sv)
		} else {
			tmp["health-check"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_factor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["link-cost-factor"], _ = expandSystemSdwanServiceLinkCostFactor(d, i["link_cost_factor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_loss_weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packet-loss-weight"], _ = expandSystemSdwanServicePacketLossWeight(d, i["packet_loss_weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latency_weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["latency-weight"], _ = expandSystemSdwanServiceLatencyWeight(d, i["latency_weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jitter_weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["jitter-weight"], _ = expandSystemSdwanServiceJitterWeight(d, i["jitter_weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bandwidth_weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bandwidth-weight"], _ = expandSystemSdwanServiceBandwidthWeight(d, i["bandwidth_weight"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "link_cost_threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["link-cost-threshold"], _ = expandSystemSdwanServiceLinkCostThreshold(d, i["link_cost_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hold_down_time"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hold-down-time"], _ = expandSystemSdwanServiceHoldDownTime(d, i["hold_down_time"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dscp-forward"], _ = expandSystemSdwanServiceDscpForward(d, i["dscp_forward"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dscp-reverse"], _ = expandSystemSdwanServiceDscpReverse(d, i["dscp_reverse"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_forward_tag"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dscp-forward-tag"], _ = expandSystemSdwanServiceDscpForwardTag(d, i["dscp_forward_tag"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp_reverse_tag"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dscp-reverse-tag"], _ = expandSystemSdwanServiceDscpReverseTag(d, i["dscp_reverse_tag"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sla"], _ = expandSystemSdwanServiceSla(d, i["sla"], pre_append, sv)
		} else {
			tmp["sla"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_members"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority-members"], _ = expandSystemSdwanServicePriorityMembers(d, i["priority_members"], pre_append, sv)
		} else {
			tmp["priority-members"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority_zone"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priority-zone"], _ = expandSystemSdwanServicePriorityZone(d, i["priority_zone"], pre_append, sv)
		} else {
			tmp["priority-zone"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSystemSdwanServiceStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["gateway"], _ = expandSystemSdwanServiceGateway(d, i["gateway"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "default"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["default"], _ = expandSystemSdwanServiceDefault(d, i["default"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sla_compare_method"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["sla-compare-method"], _ = expandSystemSdwanServiceSlaCompareMethod(d, i["sla_compare_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tie_break"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tie-break"], _ = expandSystemSdwanServiceTieBreak(d, i["tie_break"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "use_shortcut_sla"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["use-shortcut-sla"], _ = expandSystemSdwanServiceUseShortcutSla(d, i["use_shortcut_sla"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passive_measurement"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["passive-measurement"], _ = expandSystemSdwanServicePassiveMeasurement(d, i["passive_measurement"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceInputDeviceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInputDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceHashMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStandaloneAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceQualityLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceRouteTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceDstName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDstNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceSrcName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceSrcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceDst6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceDst6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceSrc6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceSrc6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceUsersName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceInternetServiceCustomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceCustomGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceInternetServiceCustomGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInternetServiceCustomGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceInternetServiceNameName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceInternetServiceGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInternetServiceGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceAppCtrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemSdwanServiceInternetServiceAppCtrlId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceInternetServiceAppCtrlGroupName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceInternetServiceAppCtrlGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServiceHealthCheckName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceHealthCheckName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLinkCostFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePacketLossWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLatencyWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceJitterWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceBandwidthWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceLinkCostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceHoldDownTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpForward(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpForwardTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDscpReverseTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "health_check"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["health-check"], _ = expandSystemSdwanServiceSlaHealthCheck(d, i["health_check"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemSdwanServiceSlaId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServiceSlaHealthCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSlaId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePriorityMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "seq_num"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["seq-num"], _ = expandSystemSdwanServicePriorityMembersSeqNum(d, i["seq_num"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServicePriorityMembersSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePriorityZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanServicePriorityZoneName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanServicePriorityZoneName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceSlaCompareMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceTieBreak(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServiceUseShortcutSla(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanServicePassiveMeasurement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemSdwanDuplicationId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["service-id"], _ = expandSystemSdwanDuplicationServiceId(d, i["service_id"], pre_append, sv)
		} else {
			tmp["service-id"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["srcaddr"], _ = expandSystemSdwanDuplicationSrcaddr(d, i["srcaddr"], pre_append, sv)
		} else {
			tmp["srcaddr"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dstaddr"], _ = expandSystemSdwanDuplicationDstaddr(d, i["dstaddr"], pre_append, sv)
		} else {
			tmp["dstaddr"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["srcaddr6"], _ = expandSystemSdwanDuplicationSrcaddr6(d, i["srcaddr6"], pre_append, sv)
		} else {
			tmp["srcaddr6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dstaddr6"], _ = expandSystemSdwanDuplicationDstaddr6(d, i["dstaddr6"], pre_append, sv)
		} else {
			tmp["dstaddr6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcintf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["srcintf"], _ = expandSystemSdwanDuplicationSrcintf(d, i["srcintf"], pre_append, sv)
		} else {
			tmp["srcintf"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstintf"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dstintf"], _ = expandSystemSdwanDuplicationDstintf(d, i["dstintf"], pre_append, sv)
		} else {
			tmp["dstintf"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["service"], _ = expandSystemSdwanDuplicationService(d, i["service"], pre_append, sv)
		} else {
			tmp["service"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_duplication"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packet-duplication"], _ = expandSystemSdwanDuplicationPacketDuplication(d, i["packet_duplication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "packet_de_duplication"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["packet-de-duplication"], _ = expandSystemSdwanDuplicationPacketDeDuplication(d, i["packet_de_duplication"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemSdwanDuplicationServiceIdId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationServiceIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationSrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationSrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationDstaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationSrcaddr6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationSrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationDstaddr6Name(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationDstaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationSrcintfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationSrcintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationDstintfName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationDstintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSystemSdwanDuplicationServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdwanDuplicationServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationPacketDuplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdwanDuplicationPacketDeDuplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdwan(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSystemSdwanStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_mode"); ok {

		t, err := expandSystemSdwanLoadBalanceMode(d, v, "load_balance_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-mode"] = t
		}
	}

	if v, ok := d.GetOk("speedtest_bypass_routing"); ok {

		t, err := expandSystemSdwanSpeedtestBypassRouting(d, v, "speedtest_bypass_routing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speedtest-bypass-routing"] = t
		}
	}

	if v, ok := d.GetOk("duplication_max_num"); ok {

		t, err := expandSystemSdwanDuplicationMaxNum(d, v, "duplication_max_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication-max-num"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_hold_down"); ok {

		t, err := expandSystemSdwanNeighborHoldDown(d, v, "neighbor_hold_down", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down"] = t
		}
	}

	if v, ok := d.GetOkExists("neighbor_hold_down_time"); ok {

		t, err := expandSystemSdwanNeighborHoldDownTime(d, v, "neighbor_hold_down_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-down-time"] = t
		}
	}

	if v, ok := d.GetOkExists("neighbor_hold_boot_time"); ok {

		t, err := expandSystemSdwanNeighborHoldBootTime(d, v, "neighbor_hold_boot_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-hold-boot-time"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok {

		t, err := expandSystemSdwanFailDetect(d, v, "fail_detect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok {

		t, err := expandSystemSdwanFailAlertInterfaces(d, v, "fail_alert_interfaces", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("zone"); ok {

		t, err := expandSystemSdwanZone(d, v, "zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["zone"] = t
		}
	}

	if v, ok := d.GetOk("members"); ok {

		t, err := expandSystemSdwanMembers(d, v, "members", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["members"] = t
		}
	}

	if v, ok := d.GetOk("health_check"); ok {

		t, err := expandSystemSdwanHealthCheck(d, v, "health_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok {

		t, err := expandSystemSdwanNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandSystemSdwanService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("duplication"); ok {

		t, err := expandSystemSdwanDuplication(d, v, "duplication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplication"] = t
		}
	}

	return &obj, nil
}
