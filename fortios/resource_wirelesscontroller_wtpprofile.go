// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerWtpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpProfileCreate,
		Read:   resourceWirelessControllerWtpProfileRead,
		Update: resourceWirelessControllerWtpProfileUpdate,
		Delete: resourceWirelessControllerWtpProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"control_message_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ble_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"wan_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port1_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port1_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port2_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port2_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port3_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port3_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port4_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port4_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port5_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port5_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port6_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port6_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port7_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port7_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"port8_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port8_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"energy_efficient_ethernet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"led_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"led_schedules": &schema.Schema{
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
					},
				},
			},
			"dtls_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtls_in_kernel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"handoff_rssi": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(20, 30),
				Optional:     true,
				Computed:     true,
			},
			"handoff_sta_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"handoff_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deny_mac_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ap_country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_fragment_preventing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tun_mtu_uplink": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 1500),
				Optional:     true,
				Computed:     true,
			},
			"tun_mtu_downlink": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 1500),
				Optional:     true,
				Computed:     true,
			},
			"split_tunneling_acl_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_acl_local_ap_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_acl": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dest_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_passwd_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Sensitive:    true,
			},
			"lldp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poe_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"radio_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2),
							Optional:     true,
							Computed:     true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"airtime_fairness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"powersave_optimize": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_optimize": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"amsdu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coexistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"short_guard_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_bonding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"dtim": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"beacon_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"rts_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(256, 2346),
							Optional:     true,
							Computed:     true,
						},
						"frag_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(800, 2346),
							Optional:     true,
							Computed:     true,
						},
						"ap_sniffer_bufsize": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),
							Optional:     true,
							Computed:     true,
						},
						"ap_sniffer_chan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_beacon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_probe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_ctl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_utilization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wids_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"darrp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_distance": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 54000),
							Optional:     true,
							Computed:     true,
						},
						"frequency_handoff": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_handoff": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vaps": &schema.Schema{
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
								},
							},
						},
						"channel": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"call_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_capacity": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 60),
							Optional:     true,
							Computed:     true,
						},
						"bandwidth_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_capacity": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 600000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"radio_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"radio_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2),
							Optional:     true,
							Computed:     true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"airtime_fairness": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protection_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"powersave_optimize": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"transmit_optimize": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"amsdu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"coexistence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"short_guard_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_bonding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auto_power_high": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"auto_power_low": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"dtim": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"beacon_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"rts_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(256, 2346),
							Optional:     true,
							Computed:     true,
						},
						"frag_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(800, 2346),
							Optional:     true,
							Computed:     true,
						},
						"ap_sniffer_bufsize": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 32),
							Optional:     true,
							Computed:     true,
						},
						"ap_sniffer_chan": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_beacon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_probe": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_mgmt_other": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_ctl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_sniffer_data": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"channel_utilization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"wids_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"darrp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_clients": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_distance": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 54000),
							Optional:     true,
							Computed:     true,
						},
						"frequency_handoff": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ap_handoff": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vap_all": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vaps": &schema.Schema{
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
								},
							},
						},
						"channel": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"chan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"call_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"call_capacity": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 60),
							Optional:     true,
							Computed:     true,
						},
						"bandwidth_admission_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth_capacity": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 600000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"lbs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ekahau_blink_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ekahau_tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"erc_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"erc_server_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1024, 65535),
							Optional:     true,
							Computed:     true,
						},
						"aeroscout": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_server_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1024, 65535),
							Optional:     true,
							Computed:     true,
						},
						"aeroscout_mu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_ap_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mmu_report": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mu_factor": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"aeroscout_mu_timeout": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"fortipresence": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(300, 65535),
							Optional:     true,
							Computed:     true,
						},
						"fortipresence_secret": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 123),
							Optional:     true,
							Sensitive:    true,
						},
						"fortipresence_project": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"fortipresence_frequency": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 65535),
							Optional:     true,
							Computed:     true,
						},
						"fortipresence_rogue": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_unassoc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"station_locate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ext_info_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerWtpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerWtpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWtpProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWtpProfile")
	}

	return resourceWirelessControllerWtpProfileRead(d, m)
}

func resourceWirelessControllerWtpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerWtpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWtpProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWtpProfile")
	}

	return resourceWirelessControllerWtpProfileRead(d, m)
}

func resourceWirelessControllerWtpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerWtpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerWtpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatform(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenWirelessControllerWtpProfilePlatformType(i["type"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfilePlatformType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileControlMessageOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileBleProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileWanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := i["port-mode"]; ok {
		result["port_mode"] = flattenWirelessControllerWtpProfileLanPortMode(i["port-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_ssid"
	if _, ok := i["port-ssid"]; ok {
		result["port_ssid"] = flattenWirelessControllerWtpProfileLanPortSsid(i["port-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_mode"
	if _, ok := i["port1-mode"]; ok {
		result["port1_mode"] = flattenWirelessControllerWtpProfileLanPort1Mode(i["port1-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := i["port1-ssid"]; ok {
		result["port1_ssid"] = flattenWirelessControllerWtpProfileLanPort1Ssid(i["port1-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_mode"
	if _, ok := i["port2-mode"]; ok {
		result["port2_mode"] = flattenWirelessControllerWtpProfileLanPort2Mode(i["port2-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := i["port2-ssid"]; ok {
		result["port2_ssid"] = flattenWirelessControllerWtpProfileLanPort2Ssid(i["port2-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_mode"
	if _, ok := i["port3-mode"]; ok {
		result["port3_mode"] = flattenWirelessControllerWtpProfileLanPort3Mode(i["port3-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := i["port3-ssid"]; ok {
		result["port3_ssid"] = flattenWirelessControllerWtpProfileLanPort3Ssid(i["port3-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_mode"
	if _, ok := i["port4-mode"]; ok {
		result["port4_mode"] = flattenWirelessControllerWtpProfileLanPort4Mode(i["port4-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := i["port4-ssid"]; ok {
		result["port4_ssid"] = flattenWirelessControllerWtpProfileLanPort4Ssid(i["port4-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_mode"
	if _, ok := i["port5-mode"]; ok {
		result["port5_mode"] = flattenWirelessControllerWtpProfileLanPort5Mode(i["port5-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := i["port5-ssid"]; ok {
		result["port5_ssid"] = flattenWirelessControllerWtpProfileLanPort5Ssid(i["port5-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_mode"
	if _, ok := i["port6-mode"]; ok {
		result["port6_mode"] = flattenWirelessControllerWtpProfileLanPort6Mode(i["port6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := i["port6-ssid"]; ok {
		result["port6_ssid"] = flattenWirelessControllerWtpProfileLanPort6Ssid(i["port6-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_mode"
	if _, ok := i["port7-mode"]; ok {
		result["port7_mode"] = flattenWirelessControllerWtpProfileLanPort7Mode(i["port7-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := i["port7-ssid"]; ok {
		result["port7_ssid"] = flattenWirelessControllerWtpProfileLanPort7Ssid(i["port7-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_mode"
	if _, ok := i["port8-mode"]; ok {
		result["port8_mode"] = flattenWirelessControllerWtpProfileLanPort8Mode(i["port8-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := i["port8-ssid"]; ok {
		result["port8_ssid"] = flattenWirelessControllerWtpProfileLanPort8Ssid(i["port8-ssid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileLanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPortSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort1Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort2Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort3Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort3Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort4Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort4Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort5Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort5Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort6Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort7Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort7Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort8Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort8Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEnergyEfficientEthernet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLedState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLedSchedules(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerWtpProfileLedSchedulesName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileLedSchedulesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDtlsPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDtlsInKernel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileHandoffRssi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileHandoffStaThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileHandoffRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDenyMacList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerWtpProfileDenyMacListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			tmp["mac"] = flattenWirelessControllerWtpProfileDenyMacListMac(i["mac"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileDenyMacListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDenyMacListMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileApCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileIpFragmentPreventing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileTunMtuUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileTunMtuDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAcl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerWtpProfileSplitTunnelingAclId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := i["dest-ip"]; ok {
			tmp["dest_ip"] = flattenWirelessControllerWtpProfileSplitTunnelingAclDestIp(i["dest-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclDestIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenWirelessControllerWtpProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLoginPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLldp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePoeMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpProfileRadio1RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfileRadio1Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpProfileRadio1Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenWirelessControllerWtpProfileRadio1AirtimeFairness(i["airtime-fairness"], d, pre_append)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenWirelessControllerWtpProfileRadio1ProtectionMode(i["protection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenWirelessControllerWtpProfileRadio1PowersaveOptimize(i["powersave-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenWirelessControllerWtpProfileRadio1TransmitOptimize(i["transmit-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenWirelessControllerWtpProfileRadio1Amsdu(i["amsdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenWirelessControllerWtpProfileRadio1Coexistence(i["coexistence"], d, pre_append)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenWirelessControllerWtpProfileRadio1ShortGuardInterval(i["short-guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenWirelessControllerWtpProfileRadio1ChannelBonding(i["channel-bonding"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpProfileRadio1AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpProfileRadio1AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpProfileRadio1AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpProfileRadio1PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenWirelessControllerWtpProfileRadio1Dtim(i["dtim"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenWirelessControllerWtpProfileRadio1BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenWirelessControllerWtpProfileRadio1RtsThreshold(i["rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenWirelessControllerWtpProfileRadio1FragThreshold(i["frag-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenWirelessControllerWtpProfileRadio1ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenWirelessControllerWtpProfileRadio1ApSnifferChan(i["ap-sniffer-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenWirelessControllerWtpProfileRadio1ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenWirelessControllerWtpProfileRadio1ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenWirelessControllerWtpProfileRadio1ApSnifferData(i["ap-sniffer-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenWirelessControllerWtpProfileRadio1ChannelUtilization(i["channel-utilization"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpProfileRadio1SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenWirelessControllerWtpProfileRadio1WidsProfile(i["wids-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenWirelessControllerWtpProfileRadio1Darrp(i["darrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenWirelessControllerWtpProfileRadio1MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenWirelessControllerWtpProfileRadio1MaxDistance(i["max-distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := i["frequency-handoff"]; ok {
		result["frequency_handoff"] = flattenWirelessControllerWtpProfileRadio1FrequencyHandoff(i["frequency-handoff"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := i["ap-handoff"]; ok {
		result["ap_handoff"] = flattenWirelessControllerWtpProfileRadio1ApHandoff(i["ap-handoff"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpProfileRadio1VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpProfileRadio1Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpProfileRadio1Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenWirelessControllerWtpProfileRadio1CallAdmissionControl(i["call-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenWirelessControllerWtpProfileRadio1CallCapacity(i["call-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenWirelessControllerWtpProfileRadio1BandwidthCapacity(i["bandwidth-capacity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileRadio1RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AirtimeFairness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ProtectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1TransmitOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Amsdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Coexistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ChannelBonding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Dtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1RtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1FragThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ChannelUtilization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1WidsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Darrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1MaxDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1FrequencyHandoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApHandoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Vaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerWtpProfileRadio1VapsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileRadio1VapsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Channel(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := i["chan"]; ok {
			tmp["chan"] = flattenWirelessControllerWtpProfileRadio1ChannelChan(i["chan"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileRadio1ChannelChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1CallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpProfileRadio2RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfileRadio2Mode(i["mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpProfileRadio2Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenWirelessControllerWtpProfileRadio2AirtimeFairness(i["airtime-fairness"], d, pre_append)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenWirelessControllerWtpProfileRadio2ProtectionMode(i["protection-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenWirelessControllerWtpProfileRadio2PowersaveOptimize(i["powersave-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenWirelessControllerWtpProfileRadio2TransmitOptimize(i["transmit-optimize"], d, pre_append)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenWirelessControllerWtpProfileRadio2Amsdu(i["amsdu"], d, pre_append)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenWirelessControllerWtpProfileRadio2Coexistence(i["coexistence"], d, pre_append)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenWirelessControllerWtpProfileRadio2ShortGuardInterval(i["short-guard-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenWirelessControllerWtpProfileRadio2ChannelBonding(i["channel-bonding"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpProfileRadio2AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpProfileRadio2AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpProfileRadio2AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpProfileRadio2PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenWirelessControllerWtpProfileRadio2Dtim(i["dtim"], d, pre_append)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenWirelessControllerWtpProfileRadio2BeaconInterval(i["beacon-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenWirelessControllerWtpProfileRadio2RtsThreshold(i["rts-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenWirelessControllerWtpProfileRadio2FragThreshold(i["frag-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenWirelessControllerWtpProfileRadio2ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenWirelessControllerWtpProfileRadio2ApSnifferChan(i["ap-sniffer-chan"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenWirelessControllerWtpProfileRadio2ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenWirelessControllerWtpProfileRadio2ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenWirelessControllerWtpProfileRadio2ApSnifferData(i["ap-sniffer-data"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenWirelessControllerWtpProfileRadio2ChannelUtilization(i["channel-utilization"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpProfileRadio2SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenWirelessControllerWtpProfileRadio2WidsProfile(i["wids-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenWirelessControllerWtpProfileRadio2Darrp(i["darrp"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenWirelessControllerWtpProfileRadio2MaxClients(i["max-clients"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenWirelessControllerWtpProfileRadio2MaxDistance(i["max-distance"], d, pre_append)
	}

	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := i["frequency-handoff"]; ok {
		result["frequency_handoff"] = flattenWirelessControllerWtpProfileRadio2FrequencyHandoff(i["frequency-handoff"], d, pre_append)
	}

	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := i["ap-handoff"]; ok {
		result["ap_handoff"] = flattenWirelessControllerWtpProfileRadio2ApHandoff(i["ap-handoff"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpProfileRadio2VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpProfileRadio2Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpProfileRadio2Channel(i["channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenWirelessControllerWtpProfileRadio2CallAdmissionControl(i["call-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenWirelessControllerWtpProfileRadio2CallCapacity(i["call-capacity"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenWirelessControllerWtpProfileRadio2BandwidthCapacity(i["bandwidth-capacity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileRadio2RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AirtimeFairness(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ProtectionMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2TransmitOptimize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Amsdu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Coexistence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ChannelBonding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Dtim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2RtsThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2FragThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferData(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ChannelUtilization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2WidsProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Darrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MaxDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2FrequencyHandoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApHandoff(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerWtpProfileRadio2VapsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileRadio2VapsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Channel(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := i["chan"]; ok {
			tmp["chan"] = flattenWirelessControllerWtpProfileRadio2ChannelChan(i["chan"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpProfileRadio2ChannelChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2CallCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ekahau_blink_mode"
	if _, ok := i["ekahau-blink-mode"]; ok {
		result["ekahau_blink_mode"] = flattenWirelessControllerWtpProfileLbsEkahauBlinkMode(i["ekahau-blink-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "ekahau_tag"
	if _, ok := i["ekahau-tag"]; ok {
		result["ekahau_tag"] = flattenWirelessControllerWtpProfileLbsEkahauTag(i["ekahau-tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "erc_server_ip"
	if _, ok := i["erc-server-ip"]; ok {
		result["erc_server_ip"] = flattenWirelessControllerWtpProfileLbsErcServerIp(i["erc-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "erc_server_port"
	if _, ok := i["erc-server-port"]; ok {
		result["erc_server_port"] = flattenWirelessControllerWtpProfileLbsErcServerPort(i["erc-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout"
	if _, ok := i["aeroscout"]; ok {
		result["aeroscout"] = flattenWirelessControllerWtpProfileLbsAeroscout(i["aeroscout"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_server_ip"
	if _, ok := i["aeroscout-server-ip"]; ok {
		result["aeroscout_server_ip"] = flattenWirelessControllerWtpProfileLbsAeroscoutServerIp(i["aeroscout-server-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_server_port"
	if _, ok := i["aeroscout-server-port"]; ok {
		result["aeroscout_server_port"] = flattenWirelessControllerWtpProfileLbsAeroscoutServerPort(i["aeroscout-server-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mu"
	if _, ok := i["aeroscout-mu"]; ok {
		result["aeroscout_mu"] = flattenWirelessControllerWtpProfileLbsAeroscoutMu(i["aeroscout-mu"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_ap_mac"
	if _, ok := i["aeroscout-ap-mac"]; ok {
		result["aeroscout_ap_mac"] = flattenWirelessControllerWtpProfileLbsAeroscoutApMac(i["aeroscout-ap-mac"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mmu_report"
	if _, ok := i["aeroscout-mmu-report"]; ok {
		result["aeroscout_mmu_report"] = flattenWirelessControllerWtpProfileLbsAeroscoutMmuReport(i["aeroscout-mmu-report"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mu_factor"
	if _, ok := i["aeroscout-mu-factor"]; ok {
		result["aeroscout_mu_factor"] = flattenWirelessControllerWtpProfileLbsAeroscoutMuFactor(i["aeroscout-mu-factor"], d, pre_append)
	}

	pre_append = pre + ".0." + "aeroscout_mu_timeout"
	if _, ok := i["aeroscout-mu-timeout"]; ok {
		result["aeroscout_mu_timeout"] = flattenWirelessControllerWtpProfileLbsAeroscoutMuTimeout(i["aeroscout-mu-timeout"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence"
	if _, ok := i["fortipresence"]; ok {
		result["fortipresence"] = flattenWirelessControllerWtpProfileLbsFortipresence(i["fortipresence"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_server"
	if _, ok := i["fortipresence-server"]; ok {
		result["fortipresence_server"] = flattenWirelessControllerWtpProfileLbsFortipresenceServer(i["fortipresence-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_port"
	if _, ok := i["fortipresence-port"]; ok {
		result["fortipresence_port"] = flattenWirelessControllerWtpProfileLbsFortipresencePort(i["fortipresence-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_secret"
	if _, ok := i["fortipresence-secret"]; ok {
		result["fortipresence_secret"] = flattenWirelessControllerWtpProfileLbsFortipresenceSecret(i["fortipresence-secret"], d, pre_append)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["fortipresence_secret"] = c
		}
	}

	pre_append = pre + ".0." + "fortipresence_project"
	if _, ok := i["fortipresence-project"]; ok {
		result["fortipresence_project"] = flattenWirelessControllerWtpProfileLbsFortipresenceProject(i["fortipresence-project"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_frequency"
	if _, ok := i["fortipresence-frequency"]; ok {
		result["fortipresence_frequency"] = flattenWirelessControllerWtpProfileLbsFortipresenceFrequency(i["fortipresence-frequency"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_rogue"
	if _, ok := i["fortipresence-rogue"]; ok {
		result["fortipresence_rogue"] = flattenWirelessControllerWtpProfileLbsFortipresenceRogue(i["fortipresence-rogue"], d, pre_append)
	}

	pre_append = pre + ".0." + "fortipresence_unassoc"
	if _, ok := i["fortipresence-unassoc"]; ok {
		result["fortipresence_unassoc"] = flattenWirelessControllerWtpProfileLbsFortipresenceUnassoc(i["fortipresence-unassoc"], d, pre_append)
	}

	pre_append = pre + ".0." + "station_locate"
	if _, ok := i["station-locate"]; ok {
		result["station_locate"] = flattenWirelessControllerWtpProfileLbsStationLocate(i["station-locate"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileLbsEkahauBlinkMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsEkahauTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsErcServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsErcServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutApMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMmuReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMuFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMuTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresencePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceFrequency(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceRogue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceUnassoc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsStationLocate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileExtInfoEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerWtpProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerWtpProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("platform", flattenWirelessControllerWtpProfilePlatform(o["platform"], d, "platform")); err != nil {
			if !fortiAPIPatch(o["platform"]) {
				return fmt.Errorf("Error reading platform: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("platform"); ok {
			if err = d.Set("platform", flattenWirelessControllerWtpProfilePlatform(o["platform"], d, "platform")); err != nil {
				if !fortiAPIPatch(o["platform"]) {
					return fmt.Errorf("Error reading platform: %v", err)
				}
			}
		}
	}

	if err = d.Set("control_message_offload", flattenWirelessControllerWtpProfileControlMessageOffload(o["control-message-offload"], d, "control_message_offload")); err != nil {
		if !fortiAPIPatch(o["control-message-offload"]) {
			return fmt.Errorf("Error reading control_message_offload: %v", err)
		}
	}

	if err = d.Set("ble_profile", flattenWirelessControllerWtpProfileBleProfile(o["ble-profile"], d, "ble_profile")); err != nil {
		if !fortiAPIPatch(o["ble-profile"]) {
			return fmt.Errorf("Error reading ble_profile: %v", err)
		}
	}

	if err = d.Set("wan_port_mode", flattenWirelessControllerWtpProfileWanPortMode(o["wan-port-mode"], d, "wan_port_mode")); err != nil {
		if !fortiAPIPatch(o["wan-port-mode"]) {
			return fmt.Errorf("Error reading wan_port_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan", flattenWirelessControllerWtpProfileLan(o["lan"], d, "lan")); err != nil {
			if !fortiAPIPatch(o["lan"]) {
				return fmt.Errorf("Error reading lan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan"); ok {
			if err = d.Set("lan", flattenWirelessControllerWtpProfileLan(o["lan"], d, "lan")); err != nil {
				if !fortiAPIPatch(o["lan"]) {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			}
		}
	}

	if err = d.Set("energy_efficient_ethernet", flattenWirelessControllerWtpProfileEnergyEfficientEthernet(o["energy-efficient-ethernet"], d, "energy_efficient_ethernet")); err != nil {
		if !fortiAPIPatch(o["energy-efficient-ethernet"]) {
			return fmt.Errorf("Error reading energy_efficient_ethernet: %v", err)
		}
	}

	if err = d.Set("led_state", flattenWirelessControllerWtpProfileLedState(o["led-state"], d, "led_state")); err != nil {
		if !fortiAPIPatch(o["led-state"]) {
			return fmt.Errorf("Error reading led_state: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("led_schedules", flattenWirelessControllerWtpProfileLedSchedules(o["led-schedules"], d, "led_schedules")); err != nil {
			if !fortiAPIPatch(o["led-schedules"]) {
				return fmt.Errorf("Error reading led_schedules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("led_schedules"); ok {
			if err = d.Set("led_schedules", flattenWirelessControllerWtpProfileLedSchedules(o["led-schedules"], d, "led_schedules")); err != nil {
				if !fortiAPIPatch(o["led-schedules"]) {
					return fmt.Errorf("Error reading led_schedules: %v", err)
				}
			}
		}
	}

	if err = d.Set("dtls_policy", flattenWirelessControllerWtpProfileDtlsPolicy(o["dtls-policy"], d, "dtls_policy")); err != nil {
		if !fortiAPIPatch(o["dtls-policy"]) {
			return fmt.Errorf("Error reading dtls_policy: %v", err)
		}
	}

	if err = d.Set("dtls_in_kernel", flattenWirelessControllerWtpProfileDtlsInKernel(o["dtls-in-kernel"], d, "dtls_in_kernel")); err != nil {
		if !fortiAPIPatch(o["dtls-in-kernel"]) {
			return fmt.Errorf("Error reading dtls_in_kernel: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerWtpProfileMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if !fortiAPIPatch(o["max-clients"]) {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("handoff_rssi", flattenWirelessControllerWtpProfileHandoffRssi(o["handoff-rssi"], d, "handoff_rssi")); err != nil {
		if !fortiAPIPatch(o["handoff-rssi"]) {
			return fmt.Errorf("Error reading handoff_rssi: %v", err)
		}
	}

	if err = d.Set("handoff_sta_thresh", flattenWirelessControllerWtpProfileHandoffStaThresh(o["handoff-sta-thresh"], d, "handoff_sta_thresh")); err != nil {
		if !fortiAPIPatch(o["handoff-sta-thresh"]) {
			return fmt.Errorf("Error reading handoff_sta_thresh: %v", err)
		}
	}

	if err = d.Set("handoff_roaming", flattenWirelessControllerWtpProfileHandoffRoaming(o["handoff-roaming"], d, "handoff_roaming")); err != nil {
		if !fortiAPIPatch(o["handoff-roaming"]) {
			return fmt.Errorf("Error reading handoff_roaming: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("deny_mac_list", flattenWirelessControllerWtpProfileDenyMacList(o["deny-mac-list"], d, "deny_mac_list")); err != nil {
			if !fortiAPIPatch(o["deny-mac-list"]) {
				return fmt.Errorf("Error reading deny_mac_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("deny_mac_list"); ok {
			if err = d.Set("deny_mac_list", flattenWirelessControllerWtpProfileDenyMacList(o["deny-mac-list"], d, "deny_mac_list")); err != nil {
				if !fortiAPIPatch(o["deny-mac-list"]) {
					return fmt.Errorf("Error reading deny_mac_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ap_country", flattenWirelessControllerWtpProfileApCountry(o["ap-country"], d, "ap_country")); err != nil {
		if !fortiAPIPatch(o["ap-country"]) {
			return fmt.Errorf("Error reading ap_country: %v", err)
		}
	}

	if err = d.Set("ip_fragment_preventing", flattenWirelessControllerWtpProfileIpFragmentPreventing(o["ip-fragment-preventing"], d, "ip_fragment_preventing")); err != nil {
		if !fortiAPIPatch(o["ip-fragment-preventing"]) {
			return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
		}
	}

	if err = d.Set("tun_mtu_uplink", flattenWirelessControllerWtpProfileTunMtuUplink(o["tun-mtu-uplink"], d, "tun_mtu_uplink")); err != nil {
		if !fortiAPIPatch(o["tun-mtu-uplink"]) {
			return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
		}
	}

	if err = d.Set("tun_mtu_downlink", flattenWirelessControllerWtpProfileTunMtuDownlink(o["tun-mtu-downlink"], d, "tun_mtu_downlink")); err != nil {
		if !fortiAPIPatch(o["tun-mtu-downlink"]) {
			return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_path", flattenWirelessControllerWtpProfileSplitTunnelingAclPath(o["split-tunneling-acl-path"], d, "split_tunneling_acl_path")); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-path"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_local_ap_subnet", flattenWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(o["split-tunneling-acl-local-ap-subnet"], d, "split_tunneling_acl_local_ap_subnet")); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-local-ap-subnet"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpProfileSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
			if !fortiAPIPatch(o["split-tunneling-acl"]) {
				return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_acl"); ok {
			if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpProfileSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
				if !fortiAPIPatch(o["split-tunneling-acl"]) {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			}
		}
	}

	if err = d.Set("allowaccess", flattenWirelessControllerWtpProfileAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("login_passwd_change", flattenWirelessControllerWtpProfileLoginPasswdChange(o["login-passwd-change"], d, "login_passwd_change")); err != nil {
		if !fortiAPIPatch(o["login-passwd-change"]) {
			return fmt.Errorf("Error reading login_passwd_change: %v", err)
		}
	}

	if err = d.Set("lldp", flattenWirelessControllerWtpProfileLldp(o["lldp"], d, "lldp")); err != nil {
		if !fortiAPIPatch(o["lldp"]) {
			return fmt.Errorf("Error reading lldp: %v", err)
		}
	}

	if err = d.Set("poe_mode", flattenWirelessControllerWtpProfilePoeMode(o["poe-mode"], d, "poe_mode")); err != nil {
		if !fortiAPIPatch(o["poe-mode"]) {
			return fmt.Errorf("Error reading poe_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radio_1", flattenWirelessControllerWtpProfileRadio1(o["radio-1"], d, "radio_1")); err != nil {
			if !fortiAPIPatch(o["radio-1"]) {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenWirelessControllerWtpProfileRadio1(o["radio-1"], d, "radio_1")); err != nil {
				if !fortiAPIPatch(o["radio-1"]) {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_2", flattenWirelessControllerWtpProfileRadio2(o["radio-2"], d, "radio_2")); err != nil {
			if !fortiAPIPatch(o["radio-2"]) {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenWirelessControllerWtpProfileRadio2(o["radio-2"], d, "radio_2")); err != nil {
				if !fortiAPIPatch(o["radio-2"]) {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("lbs", flattenWirelessControllerWtpProfileLbs(o["lbs"], d, "lbs")); err != nil {
			if !fortiAPIPatch(o["lbs"]) {
				return fmt.Errorf("Error reading lbs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lbs"); ok {
			if err = d.Set("lbs", flattenWirelessControllerWtpProfileLbs(o["lbs"], d, "lbs")); err != nil {
				if !fortiAPIPatch(o["lbs"]) {
					return fmt.Errorf("Error reading lbs: %v", err)
				}
			}
		}
	}

	if err = d.Set("ext_info_enable", flattenWirelessControllerWtpProfileExtInfoEnable(o["ext-info-enable"], d, "ext_info_enable")); err != nil {
		if !fortiAPIPatch(o["ext-info-enable"]) {
			return fmt.Errorf("Error reading ext_info_enable: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatform(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok {
		result["type"], _ = expandWirelessControllerWtpProfilePlatformType(d, i["type"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpProfilePlatformType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileControlMessageOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileBleProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileWanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-mode"], _ = expandWirelessControllerWtpProfileLanPortMode(d, i["port_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-ssid"], _ = expandWirelessControllerWtpProfileLanPortSsid(d, i["port_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port1_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-mode"], _ = expandWirelessControllerWtpProfileLanPort1Mode(d, i["port1_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-ssid"], _ = expandWirelessControllerWtpProfileLanPort1Ssid(d, i["port1_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port2_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-mode"], _ = expandWirelessControllerWtpProfileLanPort2Mode(d, i["port2_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-ssid"], _ = expandWirelessControllerWtpProfileLanPort2Ssid(d, i["port2_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port3_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-mode"], _ = expandWirelessControllerWtpProfileLanPort3Mode(d, i["port3_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-ssid"], _ = expandWirelessControllerWtpProfileLanPort3Ssid(d, i["port3_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port4_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-mode"], _ = expandWirelessControllerWtpProfileLanPort4Mode(d, i["port4_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-ssid"], _ = expandWirelessControllerWtpProfileLanPort4Ssid(d, i["port4_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port5_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-mode"], _ = expandWirelessControllerWtpProfileLanPort5Mode(d, i["port5_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-ssid"], _ = expandWirelessControllerWtpProfileLanPort5Ssid(d, i["port5_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port6_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-mode"], _ = expandWirelessControllerWtpProfileLanPort6Mode(d, i["port6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-ssid"], _ = expandWirelessControllerWtpProfileLanPort6Ssid(d, i["port6_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port7_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-mode"], _ = expandWirelessControllerWtpProfileLanPort7Mode(d, i["port7_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-ssid"], _ = expandWirelessControllerWtpProfileLanPort7Ssid(d, i["port7_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port8_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-mode"], _ = expandWirelessControllerWtpProfileLanPort8Mode(d, i["port8_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-ssid"], _ = expandWirelessControllerWtpProfileLanPort8Ssid(d, i["port8_ssid"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileLanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPortSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort1Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort2Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort3Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort3Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort4Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort4Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort5Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort5Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort6Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort7Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort7Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort8Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort8Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEnergyEfficientEthernet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLedState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLedSchedules(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileLedSchedulesName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileLedSchedulesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDtlsPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDtlsInKernel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileHandoffRssi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileHandoffStaThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileHandoffRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDenyMacList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerWtpProfileDenyMacListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandWirelessControllerWtpProfileDenyMacListMac(d, i["mac"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileDenyMacListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDenyMacListMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileApCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileIpFragmentPreventing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileTunMtuUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileTunMtuDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerWtpProfileSplitTunnelingAclId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dest-ip"], _ = expandWirelessControllerWtpProfileSplitTunnelingAclDestIp(d, i["dest_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclDestIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLoginPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLldp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePoeMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandWirelessControllerWtpProfileRadio1RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfileRadio1Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpProfileRadio1Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandWirelessControllerWtpProfileRadio1AirtimeFairness(d, i["airtime_fairness"], pre_append)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandWirelessControllerWtpProfileRadio1ProtectionMode(d, i["protection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandWirelessControllerWtpProfileRadio1PowersaveOptimize(d, i["powersave_optimize"], pre_append)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandWirelessControllerWtpProfileRadio1TransmitOptimize(d, i["transmit_optimize"], pre_append)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandWirelessControllerWtpProfileRadio1Amsdu(d, i["amsdu"], pre_append)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandWirelessControllerWtpProfileRadio1Coexistence(d, i["coexistence"], pre_append)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandWirelessControllerWtpProfileRadio1ShortGuardInterval(d, i["short_guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandWirelessControllerWtpProfileRadio1ChannelBonding(d, i["channel_bonding"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpProfileRadio1PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandWirelessControllerWtpProfileRadio1Dtim(d, i["dtim"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandWirelessControllerWtpProfileRadio1BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandWirelessControllerWtpProfileRadio1RtsThreshold(d, i["rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandWirelessControllerWtpProfileRadio1FragThreshold(d, i["frag_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferChan(d, i["ap_sniffer_chan"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferData(d, i["ap_sniffer_data"], pre_append)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandWirelessControllerWtpProfileRadio1ChannelUtilization(d, i["channel_utilization"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpProfileRadio1SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandWirelessControllerWtpProfileRadio1WidsProfile(d, i["wids_profile"], pre_append)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandWirelessControllerWtpProfileRadio1Darrp(d, i["darrp"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandWirelessControllerWtpProfileRadio1MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandWirelessControllerWtpProfileRadio1MaxDistance(d, i["max_distance"], pre_append)
	}
	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["frequency-handoff"], _ = expandWirelessControllerWtpProfileRadio1FrequencyHandoff(d, i["frequency_handoff"], pre_append)
	}
	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-handoff"], _ = expandWirelessControllerWtpProfileRadio1ApHandoff(d, i["ap_handoff"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpProfileRadio1VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpProfileRadio1Vaps(d, i["vaps"], pre_append)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpProfileRadio1Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandWirelessControllerWtpProfileRadio1CallAdmissionControl(d, i["call_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandWirelessControllerWtpProfileRadio1CallCapacity(d, i["call_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandWirelessControllerWtpProfileRadio1BandwidthCapacity(d, i["bandwidth_capacity"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio1RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AirtimeFairness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ProtectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1TransmitOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Amsdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Coexistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ChannelBonding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Dtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1RtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1FragThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ChannelUtilization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1WidsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Darrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1MaxDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1FrequencyHandoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApHandoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileRadio1VapsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio1VapsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chan"], _ = expandWirelessControllerWtpProfileRadio1ChannelChan(d, i["chan"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio1ChannelChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1CallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandWirelessControllerWtpProfileRadio2RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfileRadio2Mode(d, i["mode"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpProfileRadio2Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandWirelessControllerWtpProfileRadio2AirtimeFairness(d, i["airtime_fairness"], pre_append)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandWirelessControllerWtpProfileRadio2ProtectionMode(d, i["protection_mode"], pre_append)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandWirelessControllerWtpProfileRadio2PowersaveOptimize(d, i["powersave_optimize"], pre_append)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandWirelessControllerWtpProfileRadio2TransmitOptimize(d, i["transmit_optimize"], pre_append)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandWirelessControllerWtpProfileRadio2Amsdu(d, i["amsdu"], pre_append)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandWirelessControllerWtpProfileRadio2Coexistence(d, i["coexistence"], pre_append)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandWirelessControllerWtpProfileRadio2ShortGuardInterval(d, i["short_guard_interval"], pre_append)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandWirelessControllerWtpProfileRadio2ChannelBonding(d, i["channel_bonding"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpProfileRadio2PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandWirelessControllerWtpProfileRadio2Dtim(d, i["dtim"], pre_append)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandWirelessControllerWtpProfileRadio2BeaconInterval(d, i["beacon_interval"], pre_append)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandWirelessControllerWtpProfileRadio2RtsThreshold(d, i["rts_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandWirelessControllerWtpProfileRadio2FragThreshold(d, i["frag_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferChan(d, i["ap_sniffer_chan"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferData(d, i["ap_sniffer_data"], pre_append)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandWirelessControllerWtpProfileRadio2ChannelUtilization(d, i["channel_utilization"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpProfileRadio2SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandWirelessControllerWtpProfileRadio2WidsProfile(d, i["wids_profile"], pre_append)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandWirelessControllerWtpProfileRadio2Darrp(d, i["darrp"], pre_append)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandWirelessControllerWtpProfileRadio2MaxClients(d, i["max_clients"], pre_append)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandWirelessControllerWtpProfileRadio2MaxDistance(d, i["max_distance"], pre_append)
	}
	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["frequency-handoff"], _ = expandWirelessControllerWtpProfileRadio2FrequencyHandoff(d, i["frequency_handoff"], pre_append)
	}
	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-handoff"], _ = expandWirelessControllerWtpProfileRadio2ApHandoff(d, i["ap_handoff"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpProfileRadio2VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpProfileRadio2Vaps(d, i["vaps"], pre_append)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpProfileRadio2Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandWirelessControllerWtpProfileRadio2CallAdmissionControl(d, i["call_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandWirelessControllerWtpProfileRadio2CallCapacity(d, i["call_capacity"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandWirelessControllerWtpProfileRadio2BandwidthCapacity(d, i["bandwidth_capacity"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio2RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AirtimeFairness(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ProtectionMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2TransmitOptimize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Amsdu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Coexistence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelBonding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Dtim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2RtsThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2FragThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferData(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelUtilization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2WidsProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Darrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MaxDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2FrequencyHandoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApHandoff(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileRadio2VapsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio2VapsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chan"], _ = expandWirelessControllerWtpProfileRadio2ChannelChan(d, i["chan"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2CallCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ekahau_blink_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ekahau-blink-mode"], _ = expandWirelessControllerWtpProfileLbsEkahauBlinkMode(d, i["ekahau_blink_mode"], pre_append)
	}
	pre_append = pre + ".0." + "ekahau_tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ekahau-tag"], _ = expandWirelessControllerWtpProfileLbsEkahauTag(d, i["ekahau_tag"], pre_append)
	}
	pre_append = pre + ".0." + "erc_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["erc-server-ip"], _ = expandWirelessControllerWtpProfileLbsErcServerIp(d, i["erc_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "erc_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["erc-server-port"], _ = expandWirelessControllerWtpProfileLbsErcServerPort(d, i["erc_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout"], _ = expandWirelessControllerWtpProfileLbsAeroscout(d, i["aeroscout"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-server-ip"], _ = expandWirelessControllerWtpProfileLbsAeroscoutServerIp(d, i["aeroscout_server_ip"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-server-port"], _ = expandWirelessControllerWtpProfileLbsAeroscoutServerPort(d, i["aeroscout_server_port"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mu"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMu(d, i["aeroscout_mu"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_ap_mac"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-ap-mac"], _ = expandWirelessControllerWtpProfileLbsAeroscoutApMac(d, i["aeroscout_ap_mac"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mmu_report"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mmu-report"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMmuReport(d, i["aeroscout_mmu_report"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mu_factor"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu-factor"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMuFactor(d, i["aeroscout_mu_factor"], pre_append)
	}
	pre_append = pre + ".0." + "aeroscout_mu_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu-timeout"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMuTimeout(d, i["aeroscout_mu_timeout"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence"], _ = expandWirelessControllerWtpProfileLbsFortipresence(d, i["fortipresence"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-server"], _ = expandWirelessControllerWtpProfileLbsFortipresenceServer(d, i["fortipresence_server"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-port"], _ = expandWirelessControllerWtpProfileLbsFortipresencePort(d, i["fortipresence_port"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_secret"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-secret"], _ = expandWirelessControllerWtpProfileLbsFortipresenceSecret(d, i["fortipresence_secret"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_project"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-project"], _ = expandWirelessControllerWtpProfileLbsFortipresenceProject(d, i["fortipresence_project"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_frequency"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-frequency"], _ = expandWirelessControllerWtpProfileLbsFortipresenceFrequency(d, i["fortipresence_frequency"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_rogue"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-rogue"], _ = expandWirelessControllerWtpProfileLbsFortipresenceRogue(d, i["fortipresence_rogue"], pre_append)
	}
	pre_append = pre + ".0." + "fortipresence_unassoc"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-unassoc"], _ = expandWirelessControllerWtpProfileLbsFortipresenceUnassoc(d, i["fortipresence_unassoc"], pre_append)
	}
	pre_append = pre + ".0." + "station_locate"
	if _, ok := d.GetOk(pre_append); ok {
		result["station-locate"], _ = expandWirelessControllerWtpProfileLbsStationLocate(d, i["station_locate"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileLbsEkahauBlinkMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsEkahauTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsErcServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsErcServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutApMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMmuReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMuFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMuTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresencePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceProject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceFrequency(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceRogue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceUnassoc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsStationLocate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileExtInfoEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWtpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerWtpProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("platform"); ok {
		t, err := expandWirelessControllerWtpProfilePlatform(d, v, "platform")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform"] = t
		}
	}

	if v, ok := d.GetOk("control_message_offload"); ok {
		t, err := expandWirelessControllerWtpProfileControlMessageOffload(d, v, "control_message_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-message-offload"] = t
		}
	}

	if v, ok := d.GetOk("ble_profile"); ok {
		t, err := expandWirelessControllerWtpProfileBleProfile(d, v, "ble_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-profile"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_mode"); ok {
		t, err := expandWirelessControllerWtpProfileWanPortMode(d, v, "wan_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("lan"); ok {
		t, err := expandWirelessControllerWtpProfileLan(d, v, "lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan"] = t
		}
	}

	if v, ok := d.GetOk("energy_efficient_ethernet"); ok {
		t, err := expandWirelessControllerWtpProfileEnergyEfficientEthernet(d, v, "energy_efficient_ethernet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["energy-efficient-ethernet"] = t
		}
	}

	if v, ok := d.GetOk("led_state"); ok {
		t, err := expandWirelessControllerWtpProfileLedState(d, v, "led_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-state"] = t
		}
	}

	if v, ok := d.GetOk("led_schedules"); ok {
		t, err := expandWirelessControllerWtpProfileLedSchedules(d, v, "led_schedules")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-schedules"] = t
		}
	}

	if v, ok := d.GetOk("dtls_policy"); ok {
		t, err := expandWirelessControllerWtpProfileDtlsPolicy(d, v, "dtls_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-policy"] = t
		}
	}

	if v, ok := d.GetOk("dtls_in_kernel"); ok {
		t, err := expandWirelessControllerWtpProfileDtlsInKernel(d, v, "dtls_in_kernel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-in-kernel"] = t
		}
	}

	if v, ok := d.GetOk("max_clients"); ok {
		t, err := expandWirelessControllerWtpProfileMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("handoff_rssi"); ok {
		t, err := expandWirelessControllerWtpProfileHandoffRssi(d, v, "handoff_rssi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-rssi"] = t
		}
	}

	if v, ok := d.GetOk("handoff_sta_thresh"); ok {
		t, err := expandWirelessControllerWtpProfileHandoffStaThresh(d, v, "handoff_sta_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-sta-thresh"] = t
		}
	}

	if v, ok := d.GetOk("handoff_roaming"); ok {
		t, err := expandWirelessControllerWtpProfileHandoffRoaming(d, v, "handoff_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-roaming"] = t
		}
	}

	if v, ok := d.GetOk("deny_mac_list"); ok {
		t, err := expandWirelessControllerWtpProfileDenyMacList(d, v, "deny_mac_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny-mac-list"] = t
		}
	}

	if v, ok := d.GetOk("ap_country"); ok {
		t, err := expandWirelessControllerWtpProfileApCountry(d, v, "ap_country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-country"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_preventing"); ok {
		t, err := expandWirelessControllerWtpProfileIpFragmentPreventing(d, v, "ip_fragment_preventing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-preventing"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_uplink"); ok {
		t, err := expandWirelessControllerWtpProfileTunMtuUplink(d, v, "tun_mtu_uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-uplink"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_downlink"); ok {
		t, err := expandWirelessControllerWtpProfileTunMtuDownlink(d, v, "tun_mtu_downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-downlink"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_path"); ok {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAclPath(d, v, "split_tunneling_acl_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-path"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(d, v, "split_tunneling_acl_local_ap_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-local-ap-subnet"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl"); ok {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAcl(d, v, "split_tunneling_acl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandWirelessControllerWtpProfileAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_change"); ok {
		t, err := expandWirelessControllerWtpProfileLoginPasswdChange(d, v, "login_passwd_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok {
		t, err := expandWirelessControllerWtpProfileLoginPasswd(d, v, "login_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("lldp"); ok {
		t, err := expandWirelessControllerWtpProfileLldp(d, v, "lldp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp"] = t
		}
	}

	if v, ok := d.GetOk("poe_mode"); ok {
		t, err := expandWirelessControllerWtpProfilePoeMode(d, v, "poe_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-mode"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok {
		t, err := expandWirelessControllerWtpProfileRadio1(d, v, "radio_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok {
		t, err := expandWirelessControllerWtpProfileRadio2(d, v, "radio_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	if v, ok := d.GetOk("lbs"); ok {
		t, err := expandWirelessControllerWtpProfileLbs(d, v, "lbs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lbs"] = t
		}
	}

	if v, ok := d.GetOk("ext_info_enable"); ok {
		t, err := expandWirelessControllerWtpProfileExtInfoEnable(d, v, "ext_info_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-info-enable"] = t
		}
	}

	return &obj, nil
}
