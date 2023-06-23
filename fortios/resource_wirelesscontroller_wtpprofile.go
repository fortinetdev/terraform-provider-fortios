// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WTP profiles or FortiAP profiles that define radio settings for manageable FortiAP platforms.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ddscan": &schema.Schema{
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
			"apcfg_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ble_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"syslog_profile": &schema.Schema{
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
				Computed: true,
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
						"port_esl_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port_esl_ssid": &schema.Schema{
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
				ValidateFunc: validation.StringLenBetween(0, 128),
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
			"radio_1": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
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
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
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
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"bss_color_mode": &schema.Schema{
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
						"optional_antenna": &schema.Schema{
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
						"auto_power_target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"power_value": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),
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
						"n80211d": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"sam_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_cwp_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_cwp_test_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_match_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_success_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_failure_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"iperf_server_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),
							Optional:     true,
							Computed:     true,
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
						"arrp_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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
				Computed: true,
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
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
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
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"bss_color_mode": &schema.Schema{
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
						"optional_antenna": &schema.Schema{
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
						"auto_power_target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"power_value": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),
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
						"n80211d": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"sam_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_cwp_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_cwp_test_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_match_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_success_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_failure_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"iperf_server_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),
							Optional:     true,
							Computed:     true,
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
						"arrp_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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
			"radio_3": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
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
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"bss_color_mode": &schema.Schema{
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
						"optional_antenna": &schema.Schema{
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
						"auto_power_target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"power_value": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),
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
						"n80211d": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"sam_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_cwp_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_cwp_test_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_match_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_success_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_failure_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"iperf_server_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),
							Optional:     true,
							Computed:     true,
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
						"arrp_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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
			"radio_4": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"band_5g_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drma_sensitivity": &schema.Schema{
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
						"zero_wait_dfs": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bss_color": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"bss_color_mode": &schema.Schema{
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
						"optional_antenna": &schema.Schema{
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
						"auto_power_target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"power_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"power_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 100),
							Optional:     true,
							Computed:     true,
						},
						"power_value": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 33),
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
						"n80211d": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"sam_ssid": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"sam_bssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_security_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_captive_portal": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_cwp_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_cwp_test_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_match_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_success_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_cwp_failure_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"sam_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"sam_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
						},
						"sam_test": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_server_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"iperf_server_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"iperf_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sam_report_intv": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(60, 864000),
							Optional:     true,
							Computed:     true,
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
						"arrp_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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
				Computed: true,
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
						"fortipresence_server_addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_server": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortipresence_server_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
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
						"fortipresence_ble": &schema.Schema{
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
			"indoor_outdoor_deployment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"esl_ses_dongle": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scd_enable": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"esl_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"output_power": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"apc_addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"apc_fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"apc_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"apc_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"coex_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tls_cert_verification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tls_fqdn_verification": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"console_login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_port_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_port_auth_usrname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"wan_port_auth_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"wan_port_auth_methods": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unii_4_5ghz_band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceWirelessControllerWtpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerWtpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWtpProfile(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerWtpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWtpProfile(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerWtpProfile(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerWtpProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatform(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenWirelessControllerWtpProfilePlatformType(i["type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfilePlatformMode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ddscan"
	if _, ok := i["ddscan"]; ok {
		result["ddscan"] = flattenWirelessControllerWtpProfilePlatformDdscan(i["ddscan"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfilePlatformType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatformMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePlatformDdscan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileControlMessageOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileApcfgProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileBleProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSyslogProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileWanPortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLan(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := i["port-mode"]; ok {
		result["port_mode"] = flattenWirelessControllerWtpProfileLanPortMode(i["port-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_ssid"
	if _, ok := i["port-ssid"]; ok {
		result["port_ssid"] = flattenWirelessControllerWtpProfileLanPortSsid(i["port-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port1_mode"
	if _, ok := i["port1-mode"]; ok {
		result["port1_mode"] = flattenWirelessControllerWtpProfileLanPort1Mode(i["port1-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := i["port1-ssid"]; ok {
		result["port1_ssid"] = flattenWirelessControllerWtpProfileLanPort1Ssid(i["port1-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port2_mode"
	if _, ok := i["port2-mode"]; ok {
		result["port2_mode"] = flattenWirelessControllerWtpProfileLanPort2Mode(i["port2-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := i["port2-ssid"]; ok {
		result["port2_ssid"] = flattenWirelessControllerWtpProfileLanPort2Ssid(i["port2-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port3_mode"
	if _, ok := i["port3-mode"]; ok {
		result["port3_mode"] = flattenWirelessControllerWtpProfileLanPort3Mode(i["port3-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := i["port3-ssid"]; ok {
		result["port3_ssid"] = flattenWirelessControllerWtpProfileLanPort3Ssid(i["port3-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port4_mode"
	if _, ok := i["port4-mode"]; ok {
		result["port4_mode"] = flattenWirelessControllerWtpProfileLanPort4Mode(i["port4-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := i["port4-ssid"]; ok {
		result["port4_ssid"] = flattenWirelessControllerWtpProfileLanPort4Ssid(i["port4-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port5_mode"
	if _, ok := i["port5-mode"]; ok {
		result["port5_mode"] = flattenWirelessControllerWtpProfileLanPort5Mode(i["port5-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := i["port5-ssid"]; ok {
		result["port5_ssid"] = flattenWirelessControllerWtpProfileLanPort5Ssid(i["port5-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port6_mode"
	if _, ok := i["port6-mode"]; ok {
		result["port6_mode"] = flattenWirelessControllerWtpProfileLanPort6Mode(i["port6-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := i["port6-ssid"]; ok {
		result["port6_ssid"] = flattenWirelessControllerWtpProfileLanPort6Ssid(i["port6-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port7_mode"
	if _, ok := i["port7-mode"]; ok {
		result["port7_mode"] = flattenWirelessControllerWtpProfileLanPort7Mode(i["port7-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := i["port7-ssid"]; ok {
		result["port7_ssid"] = flattenWirelessControllerWtpProfileLanPort7Ssid(i["port7-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port8_mode"
	if _, ok := i["port8-mode"]; ok {
		result["port8_mode"] = flattenWirelessControllerWtpProfileLanPort8Mode(i["port8-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := i["port8-ssid"]; ok {
		result["port8_ssid"] = flattenWirelessControllerWtpProfileLanPort8Ssid(i["port8-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := i["port-esl-mode"]; ok {
		result["port_esl_mode"] = flattenWirelessControllerWtpProfileLanPortEslMode(i["port-esl-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := i["port-esl-ssid"]; ok {
		result["port_esl_ssid"] = flattenWirelessControllerWtpProfileLanPortEslSsid(i["port-esl-ssid"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileLanPortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPortSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort1Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort1Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort2Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort2Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort3Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort3Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort4Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort4Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort5Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort5Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort6Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort6Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort7Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort7Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort8Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPort8Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPortEslMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLanPortEslSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEnergyEfficientEthernet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLedState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLedSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerWtpProfileLedSchedulesName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpProfileLedSchedulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDtlsPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDtlsInKernel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileMaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileHandoffRssi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileHandoffStaThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileHandoffRoaming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDenyMacList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenWirelessControllerWtpProfileDenyMacListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			tmp["mac"] = flattenWirelessControllerWtpProfileDenyMacListMac(i["mac"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerWtpProfileDenyMacListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileDenyMacListMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileApCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileIpFragmentPreventing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileTunMtuUplink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileTunMtuDownlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAcl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenWirelessControllerWtpProfileSplitTunnelingAclId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := i["dest-ip"]; ok {
			tmp["dest_ip"] = flattenWirelessControllerWtpProfileSplitTunnelingAclDestIp(i["dest-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileSplitTunnelingAclDestIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenWirelessControllerWtpProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLoginPasswd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLldp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfilePoeMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileFrequencyHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileApHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpProfileRadio1RadioId(i["radio-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfileRadio1Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpProfileRadio1Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenWirelessControllerWtpProfileRadio1Band5GType(i["band-5g-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenWirelessControllerWtpProfileRadio1Drma(i["drma"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenWirelessControllerWtpProfileRadio1DrmaSensitivity(i["drma-sensitivity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenWirelessControllerWtpProfileRadio1AirtimeFairness(i["airtime-fairness"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenWirelessControllerWtpProfileRadio1ProtectionMode(i["protection-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenWirelessControllerWtpProfileRadio1PowersaveOptimize(i["powersave-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenWirelessControllerWtpProfileRadio1TransmitOptimize(i["transmit-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenWirelessControllerWtpProfileRadio1Amsdu(i["amsdu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenWirelessControllerWtpProfileRadio1Coexistence(i["coexistence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenWirelessControllerWtpProfileRadio1ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenWirelessControllerWtpProfileRadio1BssColor(i["bss-color"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenWirelessControllerWtpProfileRadio1BssColorMode(i["bss-color-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenWirelessControllerWtpProfileRadio1ShortGuardInterval(i["short-guard-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenWirelessControllerWtpProfileRadio1ChannelBonding(i["channel-bonding"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := i["optional-antenna"]; ok {
		result["optional_antenna"] = flattenWirelessControllerWtpProfileRadio1OptionalAntenna(i["optional-antenna"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpProfileRadio1AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpProfileRadio1AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpProfileRadio1AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpProfileRadio1AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpProfileRadio1PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpProfileRadio1PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpProfileRadio1PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenWirelessControllerWtpProfileRadio1Dtim(i["dtim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenWirelessControllerWtpProfileRadio1BeaconInterval(i["beacon-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenWirelessControllerWtpProfileRadio180211D(i["80211d"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenWirelessControllerWtpProfileRadio1RtsThreshold(i["rts-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenWirelessControllerWtpProfileRadio1FragThreshold(i["frag-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenWirelessControllerWtpProfileRadio1ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenWirelessControllerWtpProfileRadio1ApSnifferChan(i["ap-sniffer-chan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenWirelessControllerWtpProfileRadio1ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenWirelessControllerWtpProfileRadio1ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenWirelessControllerWtpProfileRadio1ApSnifferData(i["ap-sniffer-data"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenWirelessControllerWtpProfileRadio1SamSsid(i["sam-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenWirelessControllerWtpProfileRadio1SamBssid(i["sam-bssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenWirelessControllerWtpProfileRadio1SamSecurityType(i["sam-security-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenWirelessControllerWtpProfileRadio1SamCaptivePortal(i["sam-captive-portal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := i["sam-cwp-username"]; ok {
		result["sam_cwp_username"] = flattenWirelessControllerWtpProfileRadio1SamCwpUsername(i["sam-cwp-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := i["sam-cwp-password"]; ok {
		result["sam_cwp_password"] = flattenWirelessControllerWtpProfileRadio1SamCwpPassword(i["sam-cwp-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := i["sam-cwp-test-url"]; ok {
		result["sam_cwp_test_url"] = flattenWirelessControllerWtpProfileRadio1SamCwpTestUrl(i["sam-cwp-test-url"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := i["sam-cwp-match-string"]; ok {
		result["sam_cwp_match_string"] = flattenWirelessControllerWtpProfileRadio1SamCwpMatchString(i["sam-cwp-match-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := i["sam-cwp-success-string"]; ok {
		result["sam_cwp_success_string"] = flattenWirelessControllerWtpProfileRadio1SamCwpSuccessString(i["sam-cwp-success-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := i["sam-cwp-failure-string"]; ok {
		result["sam_cwp_failure_string"] = flattenWirelessControllerWtpProfileRadio1SamCwpFailureString(i["sam-cwp-failure-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenWirelessControllerWtpProfileRadio1SamUsername(i["sam-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenWirelessControllerWtpProfileRadio1SamPassword(i["sam-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenWirelessControllerWtpProfileRadio1SamTest(i["sam-test"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := i["sam-server-type"]; ok {
		result["sam_server_type"] = flattenWirelessControllerWtpProfileRadio1SamServerType(i["sam-server-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := i["sam-server-ip"]; ok {
		result["sam_server_ip"] = flattenWirelessControllerWtpProfileRadio1SamServerIp(i["sam-server-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := i["sam-server-fqdn"]; ok {
		result["sam_server_fqdn"] = flattenWirelessControllerWtpProfileRadio1SamServerFqdn(i["sam-server-fqdn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenWirelessControllerWtpProfileRadio1IperfServerPort(i["iperf-server-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenWirelessControllerWtpProfileRadio1IperfProtocol(i["iperf-protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenWirelessControllerWtpProfileRadio1SamReportIntv(i["sam-report-intv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenWirelessControllerWtpProfileRadio1ChannelUtilization(i["channel-utilization"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpProfileRadio1SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenWirelessControllerWtpProfileRadio1WidsProfile(i["wids-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenWirelessControllerWtpProfileRadio1Darrp(i["darrp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := i["arrp-profile"]; ok {
		result["arrp_profile"] = flattenWirelessControllerWtpProfileRadio1ArrpProfile(i["arrp-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenWirelessControllerWtpProfileRadio1MaxClients(i["max-clients"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenWirelessControllerWtpProfileRadio1MaxDistance(i["max-distance"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := i["frequency-handoff"]; ok {
		result["frequency_handoff"] = flattenWirelessControllerWtpProfileRadio1FrequencyHandoff(i["frequency-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := i["ap-handoff"]; ok {
		result["ap_handoff"] = flattenWirelessControllerWtpProfileRadio1ApHandoff(i["ap-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpProfileRadio1VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpProfileRadio1Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpProfileRadio1Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenWirelessControllerWtpProfileRadio1CallAdmissionControl(i["call-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenWirelessControllerWtpProfileRadio1CallCapacity(i["call-capacity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenWirelessControllerWtpProfileRadio1BandwidthCapacity(i["bandwidth-capacity"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileRadio1RadioId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Band5GType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Drma(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AirtimeFairness(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ProtectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1TransmitOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Amsdu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Coexistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BssColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BssColorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ChannelBonding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1OptionalAntenna(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Dtim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio180211D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1RtsThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1FragThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApSnifferData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamBssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamSecurityType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCwpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCwpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCwpTestUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCwpMatchString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCwpSuccessString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamCwpFailureString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamTest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1IperfServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1IperfProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SamReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ChannelUtilization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1WidsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Darrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ArrpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1MaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1MaxDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1FrequencyHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1ApHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerWtpProfileRadio1VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio1VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := i["chan"]; ok {
			tmp["chan"] = flattenWirelessControllerWtpProfileRadio1ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio1ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1CallCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio1BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpProfileRadio2RadioId(i["radio-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfileRadio2Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpProfileRadio2Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenWirelessControllerWtpProfileRadio2Band5GType(i["band-5g-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenWirelessControllerWtpProfileRadio2Drma(i["drma"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenWirelessControllerWtpProfileRadio2DrmaSensitivity(i["drma-sensitivity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenWirelessControllerWtpProfileRadio2AirtimeFairness(i["airtime-fairness"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenWirelessControllerWtpProfileRadio2ProtectionMode(i["protection-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenWirelessControllerWtpProfileRadio2PowersaveOptimize(i["powersave-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenWirelessControllerWtpProfileRadio2TransmitOptimize(i["transmit-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenWirelessControllerWtpProfileRadio2Amsdu(i["amsdu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenWirelessControllerWtpProfileRadio2Coexistence(i["coexistence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenWirelessControllerWtpProfileRadio2ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenWirelessControllerWtpProfileRadio2BssColor(i["bss-color"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenWirelessControllerWtpProfileRadio2BssColorMode(i["bss-color-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenWirelessControllerWtpProfileRadio2ShortGuardInterval(i["short-guard-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenWirelessControllerWtpProfileRadio2ChannelBonding(i["channel-bonding"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := i["optional-antenna"]; ok {
		result["optional_antenna"] = flattenWirelessControllerWtpProfileRadio2OptionalAntenna(i["optional-antenna"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpProfileRadio2AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpProfileRadio2AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpProfileRadio2AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpProfileRadio2AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpProfileRadio2PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpProfileRadio2PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpProfileRadio2PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenWirelessControllerWtpProfileRadio2Dtim(i["dtim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenWirelessControllerWtpProfileRadio2BeaconInterval(i["beacon-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenWirelessControllerWtpProfileRadio280211D(i["80211d"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenWirelessControllerWtpProfileRadio2RtsThreshold(i["rts-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenWirelessControllerWtpProfileRadio2FragThreshold(i["frag-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenWirelessControllerWtpProfileRadio2ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenWirelessControllerWtpProfileRadio2ApSnifferChan(i["ap-sniffer-chan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenWirelessControllerWtpProfileRadio2ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenWirelessControllerWtpProfileRadio2ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenWirelessControllerWtpProfileRadio2ApSnifferData(i["ap-sniffer-data"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenWirelessControllerWtpProfileRadio2SamSsid(i["sam-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenWirelessControllerWtpProfileRadio2SamBssid(i["sam-bssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenWirelessControllerWtpProfileRadio2SamSecurityType(i["sam-security-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenWirelessControllerWtpProfileRadio2SamCaptivePortal(i["sam-captive-portal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := i["sam-cwp-username"]; ok {
		result["sam_cwp_username"] = flattenWirelessControllerWtpProfileRadio2SamCwpUsername(i["sam-cwp-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := i["sam-cwp-password"]; ok {
		result["sam_cwp_password"] = flattenWirelessControllerWtpProfileRadio2SamCwpPassword(i["sam-cwp-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := i["sam-cwp-test-url"]; ok {
		result["sam_cwp_test_url"] = flattenWirelessControllerWtpProfileRadio2SamCwpTestUrl(i["sam-cwp-test-url"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := i["sam-cwp-match-string"]; ok {
		result["sam_cwp_match_string"] = flattenWirelessControllerWtpProfileRadio2SamCwpMatchString(i["sam-cwp-match-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := i["sam-cwp-success-string"]; ok {
		result["sam_cwp_success_string"] = flattenWirelessControllerWtpProfileRadio2SamCwpSuccessString(i["sam-cwp-success-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := i["sam-cwp-failure-string"]; ok {
		result["sam_cwp_failure_string"] = flattenWirelessControllerWtpProfileRadio2SamCwpFailureString(i["sam-cwp-failure-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenWirelessControllerWtpProfileRadio2SamUsername(i["sam-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenWirelessControllerWtpProfileRadio2SamPassword(i["sam-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenWirelessControllerWtpProfileRadio2SamTest(i["sam-test"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := i["sam-server-type"]; ok {
		result["sam_server_type"] = flattenWirelessControllerWtpProfileRadio2SamServerType(i["sam-server-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := i["sam-server-ip"]; ok {
		result["sam_server_ip"] = flattenWirelessControllerWtpProfileRadio2SamServerIp(i["sam-server-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := i["sam-server-fqdn"]; ok {
		result["sam_server_fqdn"] = flattenWirelessControllerWtpProfileRadio2SamServerFqdn(i["sam-server-fqdn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenWirelessControllerWtpProfileRadio2IperfServerPort(i["iperf-server-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenWirelessControllerWtpProfileRadio2IperfProtocol(i["iperf-protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenWirelessControllerWtpProfileRadio2SamReportIntv(i["sam-report-intv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenWirelessControllerWtpProfileRadio2ChannelUtilization(i["channel-utilization"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpProfileRadio2SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenWirelessControllerWtpProfileRadio2WidsProfile(i["wids-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenWirelessControllerWtpProfileRadio2Darrp(i["darrp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := i["arrp-profile"]; ok {
		result["arrp_profile"] = flattenWirelessControllerWtpProfileRadio2ArrpProfile(i["arrp-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenWirelessControllerWtpProfileRadio2MaxClients(i["max-clients"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenWirelessControllerWtpProfileRadio2MaxDistance(i["max-distance"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := i["frequency-handoff"]; ok {
		result["frequency_handoff"] = flattenWirelessControllerWtpProfileRadio2FrequencyHandoff(i["frequency-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := i["ap-handoff"]; ok {
		result["ap_handoff"] = flattenWirelessControllerWtpProfileRadio2ApHandoff(i["ap-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpProfileRadio2VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpProfileRadio2Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpProfileRadio2Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenWirelessControllerWtpProfileRadio2CallAdmissionControl(i["call-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenWirelessControllerWtpProfileRadio2CallCapacity(i["call-capacity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenWirelessControllerWtpProfileRadio2BandwidthCapacity(i["bandwidth-capacity"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileRadio2RadioId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Band5GType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Drma(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AirtimeFairness(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ProtectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2TransmitOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Amsdu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Coexistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BssColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BssColorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ChannelBonding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2OptionalAntenna(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Dtim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio280211D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2RtsThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2FragThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApSnifferData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamBssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamSecurityType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpTestUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpMatchString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpSuccessString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamCwpFailureString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamTest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2IperfServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2IperfProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SamReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ChannelUtilization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2WidsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Darrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ArrpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2MaxDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2FrequencyHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2ApHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerWtpProfileRadio2VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio2VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := i["chan"]; ok {
			tmp["chan"] = flattenWirelessControllerWtpProfileRadio2ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio2ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2CallCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio2BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfileRadio3Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpProfileRadio3Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenWirelessControllerWtpProfileRadio3Band5GType(i["band-5g-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenWirelessControllerWtpProfileRadio3Drma(i["drma"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenWirelessControllerWtpProfileRadio3DrmaSensitivity(i["drma-sensitivity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenWirelessControllerWtpProfileRadio3AirtimeFairness(i["airtime-fairness"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenWirelessControllerWtpProfileRadio3ProtectionMode(i["protection-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenWirelessControllerWtpProfileRadio3PowersaveOptimize(i["powersave-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenWirelessControllerWtpProfileRadio3TransmitOptimize(i["transmit-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenWirelessControllerWtpProfileRadio3Amsdu(i["amsdu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenWirelessControllerWtpProfileRadio3Coexistence(i["coexistence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenWirelessControllerWtpProfileRadio3ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenWirelessControllerWtpProfileRadio3BssColor(i["bss-color"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenWirelessControllerWtpProfileRadio3BssColorMode(i["bss-color-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenWirelessControllerWtpProfileRadio3ShortGuardInterval(i["short-guard-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenWirelessControllerWtpProfileRadio3ChannelBonding(i["channel-bonding"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := i["optional-antenna"]; ok {
		result["optional_antenna"] = flattenWirelessControllerWtpProfileRadio3OptionalAntenna(i["optional-antenna"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpProfileRadio3AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpProfileRadio3AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpProfileRadio3AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpProfileRadio3AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpProfileRadio3PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpProfileRadio3PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpProfileRadio3PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenWirelessControllerWtpProfileRadio3Dtim(i["dtim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenWirelessControllerWtpProfileRadio3BeaconInterval(i["beacon-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenWirelessControllerWtpProfileRadio380211D(i["80211d"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenWirelessControllerWtpProfileRadio3RtsThreshold(i["rts-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenWirelessControllerWtpProfileRadio3FragThreshold(i["frag-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenWirelessControllerWtpProfileRadio3ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenWirelessControllerWtpProfileRadio3ApSnifferChan(i["ap-sniffer-chan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenWirelessControllerWtpProfileRadio3ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenWirelessControllerWtpProfileRadio3ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenWirelessControllerWtpProfileRadio3ApSnifferData(i["ap-sniffer-data"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenWirelessControllerWtpProfileRadio3SamSsid(i["sam-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenWirelessControllerWtpProfileRadio3SamBssid(i["sam-bssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenWirelessControllerWtpProfileRadio3SamSecurityType(i["sam-security-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenWirelessControllerWtpProfileRadio3SamCaptivePortal(i["sam-captive-portal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := i["sam-cwp-username"]; ok {
		result["sam_cwp_username"] = flattenWirelessControllerWtpProfileRadio3SamCwpUsername(i["sam-cwp-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := i["sam-cwp-password"]; ok {
		result["sam_cwp_password"] = flattenWirelessControllerWtpProfileRadio3SamCwpPassword(i["sam-cwp-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := i["sam-cwp-test-url"]; ok {
		result["sam_cwp_test_url"] = flattenWirelessControllerWtpProfileRadio3SamCwpTestUrl(i["sam-cwp-test-url"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := i["sam-cwp-match-string"]; ok {
		result["sam_cwp_match_string"] = flattenWirelessControllerWtpProfileRadio3SamCwpMatchString(i["sam-cwp-match-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := i["sam-cwp-success-string"]; ok {
		result["sam_cwp_success_string"] = flattenWirelessControllerWtpProfileRadio3SamCwpSuccessString(i["sam-cwp-success-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := i["sam-cwp-failure-string"]; ok {
		result["sam_cwp_failure_string"] = flattenWirelessControllerWtpProfileRadio3SamCwpFailureString(i["sam-cwp-failure-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenWirelessControllerWtpProfileRadio3SamUsername(i["sam-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenWirelessControllerWtpProfileRadio3SamPassword(i["sam-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenWirelessControllerWtpProfileRadio3SamTest(i["sam-test"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := i["sam-server-type"]; ok {
		result["sam_server_type"] = flattenWirelessControllerWtpProfileRadio3SamServerType(i["sam-server-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := i["sam-server-ip"]; ok {
		result["sam_server_ip"] = flattenWirelessControllerWtpProfileRadio3SamServerIp(i["sam-server-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := i["sam-server-fqdn"]; ok {
		result["sam_server_fqdn"] = flattenWirelessControllerWtpProfileRadio3SamServerFqdn(i["sam-server-fqdn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenWirelessControllerWtpProfileRadio3IperfServerPort(i["iperf-server-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenWirelessControllerWtpProfileRadio3IperfProtocol(i["iperf-protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenWirelessControllerWtpProfileRadio3SamReportIntv(i["sam-report-intv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenWirelessControllerWtpProfileRadio3ChannelUtilization(i["channel-utilization"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpProfileRadio3SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenWirelessControllerWtpProfileRadio3WidsProfile(i["wids-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenWirelessControllerWtpProfileRadio3Darrp(i["darrp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := i["arrp-profile"]; ok {
		result["arrp_profile"] = flattenWirelessControllerWtpProfileRadio3ArrpProfile(i["arrp-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenWirelessControllerWtpProfileRadio3MaxClients(i["max-clients"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenWirelessControllerWtpProfileRadio3MaxDistance(i["max-distance"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := i["frequency-handoff"]; ok {
		result["frequency_handoff"] = flattenWirelessControllerWtpProfileRadio3FrequencyHandoff(i["frequency-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := i["ap-handoff"]; ok {
		result["ap_handoff"] = flattenWirelessControllerWtpProfileRadio3ApHandoff(i["ap-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpProfileRadio3VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpProfileRadio3Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpProfileRadio3Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenWirelessControllerWtpProfileRadio3CallAdmissionControl(i["call-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenWirelessControllerWtpProfileRadio3CallCapacity(i["call-capacity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenWirelessControllerWtpProfileRadio3BandwidthCapacity(i["bandwidth-capacity"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileRadio3Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Band5GType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Drma(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3AirtimeFairness(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ProtectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3TransmitOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Amsdu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Coexistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3BssColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3BssColorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ChannelBonding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3OptionalAntenna(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Dtim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3BeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio380211D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3RtsThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3FragThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApSnifferData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamBssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamSecurityType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCwpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCwpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCwpTestUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCwpMatchString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCwpSuccessString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamCwpFailureString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamTest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3IperfServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3IperfProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SamReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ChannelUtilization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3WidsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Darrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ArrpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3MaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3MaxDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3FrequencyHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3ApHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerWtpProfileRadio3VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio3VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := i["chan"]; ok {
			tmp["chan"] = flattenWirelessControllerWtpProfileRadio3ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio3ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3CallCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio3BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenWirelessControllerWtpProfileRadio4Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpProfileRadio4Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := i["band-5g-type"]; ok {
		result["band_5g_type"] = flattenWirelessControllerWtpProfileRadio4Band5GType(i["band-5g-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma"
	if _, ok := i["drma"]; ok {
		result["drma"] = flattenWirelessControllerWtpProfileRadio4Drma(i["drma"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := i["drma-sensitivity"]; ok {
		result["drma_sensitivity"] = flattenWirelessControllerWtpProfileRadio4DrmaSensitivity(i["drma-sensitivity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := i["airtime-fairness"]; ok {
		result["airtime_fairness"] = flattenWirelessControllerWtpProfileRadio4AirtimeFairness(i["airtime-fairness"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "protection_mode"
	if _, ok := i["protection-mode"]; ok {
		result["protection_mode"] = flattenWirelessControllerWtpProfileRadio4ProtectionMode(i["protection-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := i["powersave-optimize"]; ok {
		result["powersave_optimize"] = flattenWirelessControllerWtpProfileRadio4PowersaveOptimize(i["powersave-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := i["transmit-optimize"]; ok {
		result["transmit_optimize"] = flattenWirelessControllerWtpProfileRadio4TransmitOptimize(i["transmit-optimize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "amsdu"
	if _, ok := i["amsdu"]; ok {
		result["amsdu"] = flattenWirelessControllerWtpProfileRadio4Amsdu(i["amsdu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "coexistence"
	if _, ok := i["coexistence"]; ok {
		result["coexistence"] = flattenWirelessControllerWtpProfileRadio4Coexistence(i["coexistence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := i["zero-wait-dfs"]; ok {
		result["zero_wait_dfs"] = flattenWirelessControllerWtpProfileRadio4ZeroWaitDfs(i["zero-wait-dfs"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenWirelessControllerWtpProfileRadio4BssColor(i["bss-color"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenWirelessControllerWtpProfileRadio4BssColorMode(i["bss-color-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := i["short-guard-interval"]; ok {
		result["short_guard_interval"] = flattenWirelessControllerWtpProfileRadio4ShortGuardInterval(i["short-guard-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := i["channel-bonding"]; ok {
		result["channel_bonding"] = flattenWirelessControllerWtpProfileRadio4ChannelBonding(i["channel-bonding"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := i["optional-antenna"]; ok {
		result["optional_antenna"] = flattenWirelessControllerWtpProfileRadio4OptionalAntenna(i["optional-antenna"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpProfileRadio4AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpProfileRadio4AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpProfileRadio4AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {
		result["auto_power_target"] = flattenWirelessControllerWtpProfileRadio4AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {
		result["power_mode"] = flattenWirelessControllerWtpProfileRadio4PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpProfileRadio4PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {
		result["power_value"] = flattenWirelessControllerWtpProfileRadio4PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dtim"
	if _, ok := i["dtim"]; ok {
		result["dtim"] = flattenWirelessControllerWtpProfileRadio4Dtim(i["dtim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenWirelessControllerWtpProfileRadio4BeaconInterval(i["beacon-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenWirelessControllerWtpProfileRadio480211D(i["80211d"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := i["rts-threshold"]; ok {
		result["rts_threshold"] = flattenWirelessControllerWtpProfileRadio4RtsThreshold(i["rts-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := i["frag-threshold"]; ok {
		result["frag_threshold"] = flattenWirelessControllerWtpProfileRadio4FragThreshold(i["frag-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := i["ap-sniffer-bufsize"]; ok {
		result["ap_sniffer_bufsize"] = flattenWirelessControllerWtpProfileRadio4ApSnifferBufsize(i["ap-sniffer-bufsize"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := i["ap-sniffer-chan"]; ok {
		result["ap_sniffer_chan"] = flattenWirelessControllerWtpProfileRadio4ApSnifferChan(i["ap-sniffer-chan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := i["ap-sniffer-addr"]; ok {
		result["ap_sniffer_addr"] = flattenWirelessControllerWtpProfileRadio4ApSnifferAddr(i["ap-sniffer-addr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := i["ap-sniffer-mgmt-beacon"]; ok {
		result["ap_sniffer_mgmt_beacon"] = flattenWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(i["ap-sniffer-mgmt-beacon"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := i["ap-sniffer-mgmt-probe"]; ok {
		result["ap_sniffer_mgmt_probe"] = flattenWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(i["ap-sniffer-mgmt-probe"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := i["ap-sniffer-mgmt-other"]; ok {
		result["ap_sniffer_mgmt_other"] = flattenWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(i["ap-sniffer-mgmt-other"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := i["ap-sniffer-ctl"]; ok {
		result["ap_sniffer_ctl"] = flattenWirelessControllerWtpProfileRadio4ApSnifferCtl(i["ap-sniffer-ctl"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := i["ap-sniffer-data"]; ok {
		result["ap_sniffer_data"] = flattenWirelessControllerWtpProfileRadio4ApSnifferData(i["ap-sniffer-data"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := i["sam-ssid"]; ok {
		result["sam_ssid"] = flattenWirelessControllerWtpProfileRadio4SamSsid(i["sam-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := i["sam-bssid"]; ok {
		result["sam_bssid"] = flattenWirelessControllerWtpProfileRadio4SamBssid(i["sam-bssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := i["sam-security-type"]; ok {
		result["sam_security_type"] = flattenWirelessControllerWtpProfileRadio4SamSecurityType(i["sam-security-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := i["sam-captive-portal"]; ok {
		result["sam_captive_portal"] = flattenWirelessControllerWtpProfileRadio4SamCaptivePortal(i["sam-captive-portal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := i["sam-cwp-username"]; ok {
		result["sam_cwp_username"] = flattenWirelessControllerWtpProfileRadio4SamCwpUsername(i["sam-cwp-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := i["sam-cwp-password"]; ok {
		result["sam_cwp_password"] = flattenWirelessControllerWtpProfileRadio4SamCwpPassword(i["sam-cwp-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := i["sam-cwp-test-url"]; ok {
		result["sam_cwp_test_url"] = flattenWirelessControllerWtpProfileRadio4SamCwpTestUrl(i["sam-cwp-test-url"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := i["sam-cwp-match-string"]; ok {
		result["sam_cwp_match_string"] = flattenWirelessControllerWtpProfileRadio4SamCwpMatchString(i["sam-cwp-match-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := i["sam-cwp-success-string"]; ok {
		result["sam_cwp_success_string"] = flattenWirelessControllerWtpProfileRadio4SamCwpSuccessString(i["sam-cwp-success-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := i["sam-cwp-failure-string"]; ok {
		result["sam_cwp_failure_string"] = flattenWirelessControllerWtpProfileRadio4SamCwpFailureString(i["sam-cwp-failure-string"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_username"
	if _, ok := i["sam-username"]; ok {
		result["sam_username"] = flattenWirelessControllerWtpProfileRadio4SamUsername(i["sam-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_password"
	if _, ok := i["sam-password"]; ok {
		result["sam_password"] = flattenWirelessControllerWtpProfileRadio4SamPassword(i["sam-password"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_test"
	if _, ok := i["sam-test"]; ok {
		result["sam_test"] = flattenWirelessControllerWtpProfileRadio4SamTest(i["sam-test"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := i["sam-server-type"]; ok {
		result["sam_server_type"] = flattenWirelessControllerWtpProfileRadio4SamServerType(i["sam-server-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := i["sam-server-ip"]; ok {
		result["sam_server_ip"] = flattenWirelessControllerWtpProfileRadio4SamServerIp(i["sam-server-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := i["sam-server-fqdn"]; ok {
		result["sam_server_fqdn"] = flattenWirelessControllerWtpProfileRadio4SamServerFqdn(i["sam-server-fqdn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := i["iperf-server-port"]; ok {
		result["iperf_server_port"] = flattenWirelessControllerWtpProfileRadio4IperfServerPort(i["iperf-server-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := i["iperf-protocol"]; ok {
		result["iperf_protocol"] = flattenWirelessControllerWtpProfileRadio4IperfProtocol(i["iperf-protocol"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := i["sam-report-intv"]; ok {
		result["sam_report_intv"] = flattenWirelessControllerWtpProfileRadio4SamReportIntv(i["sam-report-intv"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := i["channel-utilization"]; ok {
		result["channel_utilization"] = flattenWirelessControllerWtpProfileRadio4ChannelUtilization(i["channel-utilization"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpProfileRadio4SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "wids_profile"
	if _, ok := i["wids-profile"]; ok {
		result["wids_profile"] = flattenWirelessControllerWtpProfileRadio4WidsProfile(i["wids-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "darrp"
	if _, ok := i["darrp"]; ok {
		result["darrp"] = flattenWirelessControllerWtpProfileRadio4Darrp(i["darrp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := i["arrp-profile"]; ok {
		result["arrp_profile"] = flattenWirelessControllerWtpProfileRadio4ArrpProfile(i["arrp-profile"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenWirelessControllerWtpProfileRadio4MaxClients(i["max-clients"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_distance"
	if _, ok := i["max-distance"]; ok {
		result["max_distance"] = flattenWirelessControllerWtpProfileRadio4MaxDistance(i["max-distance"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := i["frequency-handoff"]; ok {
		result["frequency_handoff"] = flattenWirelessControllerWtpProfileRadio4FrequencyHandoff(i["frequency-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := i["ap-handoff"]; ok {
		result["ap_handoff"] = flattenWirelessControllerWtpProfileRadio4ApHandoff(i["ap-handoff"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpProfileRadio4VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpProfileRadio4Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpProfileRadio4Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := i["call-admission-control"]; ok {
		result["call_admission_control"] = flattenWirelessControllerWtpProfileRadio4CallAdmissionControl(i["call-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "call_capacity"
	if _, ok := i["call-capacity"]; ok {
		result["call_capacity"] = flattenWirelessControllerWtpProfileRadio4CallCapacity(i["call-capacity"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := i["bandwidth-admission-control"]; ok {
		result["bandwidth_admission_control"] = flattenWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(i["bandwidth-admission-control"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := i["bandwidth-capacity"]; ok {
		result["bandwidth_capacity"] = flattenWirelessControllerWtpProfileRadio4BandwidthCapacity(i["bandwidth-capacity"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileRadio4Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Band5GType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Drma(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4DrmaSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4AirtimeFairness(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ProtectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4PowersaveOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4TransmitOptimize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Amsdu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Coexistence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ZeroWaitDfs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4BssColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4BssColorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ShortGuardInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ChannelBonding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4OptionalAntenna(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Dtim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4BeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio480211D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4RtsThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4FragThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferBufsize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferCtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApSnifferData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamBssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamSecurityType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCwpUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCwpPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCwpTestUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCwpMatchString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCwpSuccessString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamCwpFailureString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamTest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4IperfServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4IperfProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SamReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ChannelUtilization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4WidsProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Darrp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ArrpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4MaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4MaxDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4FrequencyHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4ApHandoff(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWirelessControllerWtpProfileRadio4VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio4VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := i["chan"]; ok {
			tmp["chan"] = flattenWirelessControllerWtpProfileRadio4ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpProfileRadio4ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4CallAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4CallCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileRadio4BandwidthCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbs(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ekahau_blink_mode"
	if _, ok := i["ekahau-blink-mode"]; ok {
		result["ekahau_blink_mode"] = flattenWirelessControllerWtpProfileLbsEkahauBlinkMode(i["ekahau-blink-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ekahau_tag"
	if _, ok := i["ekahau-tag"]; ok {
		result["ekahau_tag"] = flattenWirelessControllerWtpProfileLbsEkahauTag(i["ekahau-tag"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "erc_server_ip"
	if _, ok := i["erc-server-ip"]; ok {
		result["erc_server_ip"] = flattenWirelessControllerWtpProfileLbsErcServerIp(i["erc-server-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "erc_server_port"
	if _, ok := i["erc-server-port"]; ok {
		result["erc_server_port"] = flattenWirelessControllerWtpProfileLbsErcServerPort(i["erc-server-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout"
	if _, ok := i["aeroscout"]; ok {
		result["aeroscout"] = flattenWirelessControllerWtpProfileLbsAeroscout(i["aeroscout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_server_ip"
	if _, ok := i["aeroscout-server-ip"]; ok {
		result["aeroscout_server_ip"] = flattenWirelessControllerWtpProfileLbsAeroscoutServerIp(i["aeroscout-server-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_server_port"
	if _, ok := i["aeroscout-server-port"]; ok {
		result["aeroscout_server_port"] = flattenWirelessControllerWtpProfileLbsAeroscoutServerPort(i["aeroscout-server-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_mu"
	if _, ok := i["aeroscout-mu"]; ok {
		result["aeroscout_mu"] = flattenWirelessControllerWtpProfileLbsAeroscoutMu(i["aeroscout-mu"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_ap_mac"
	if _, ok := i["aeroscout-ap-mac"]; ok {
		result["aeroscout_ap_mac"] = flattenWirelessControllerWtpProfileLbsAeroscoutApMac(i["aeroscout-ap-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_mmu_report"
	if _, ok := i["aeroscout-mmu-report"]; ok {
		result["aeroscout_mmu_report"] = flattenWirelessControllerWtpProfileLbsAeroscoutMmuReport(i["aeroscout-mmu-report"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_mu_factor"
	if _, ok := i["aeroscout-mu-factor"]; ok {
		result["aeroscout_mu_factor"] = flattenWirelessControllerWtpProfileLbsAeroscoutMuFactor(i["aeroscout-mu-factor"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "aeroscout_mu_timeout"
	if _, ok := i["aeroscout-mu-timeout"]; ok {
		result["aeroscout_mu_timeout"] = flattenWirelessControllerWtpProfileLbsAeroscoutMuTimeout(i["aeroscout-mu-timeout"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence"
	if _, ok := i["fortipresence"]; ok {
		result["fortipresence"] = flattenWirelessControllerWtpProfileLbsFortipresence(i["fortipresence"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_server_addr_type"
	if _, ok := i["fortipresence-server-addr-type"]; ok {
		result["fortipresence_server_addr_type"] = flattenWirelessControllerWtpProfileLbsFortipresenceServerAddrType(i["fortipresence-server-addr-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_server"
	if _, ok := i["fortipresence-server"]; ok {
		result["fortipresence_server"] = flattenWirelessControllerWtpProfileLbsFortipresenceServer(i["fortipresence-server"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_server_fqdn"
	if _, ok := i["fortipresence-server-fqdn"]; ok {
		result["fortipresence_server_fqdn"] = flattenWirelessControllerWtpProfileLbsFortipresenceServerFqdn(i["fortipresence-server-fqdn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_port"
	if _, ok := i["fortipresence-port"]; ok {
		result["fortipresence_port"] = flattenWirelessControllerWtpProfileLbsFortipresencePort(i["fortipresence-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_secret"
	if _, ok := i["fortipresence-secret"]; ok {
		result["fortipresence_secret"] = flattenWirelessControllerWtpProfileLbsFortipresenceSecret(i["fortipresence-secret"], d, pre_append, sv)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["fortipresence_secret"] = c
		}
	}

	pre_append = pre + ".0." + "fortipresence_project"
	if _, ok := i["fortipresence-project"]; ok {
		result["fortipresence_project"] = flattenWirelessControllerWtpProfileLbsFortipresenceProject(i["fortipresence-project"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_frequency"
	if _, ok := i["fortipresence-frequency"]; ok {
		result["fortipresence_frequency"] = flattenWirelessControllerWtpProfileLbsFortipresenceFrequency(i["fortipresence-frequency"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_rogue"
	if _, ok := i["fortipresence-rogue"]; ok {
		result["fortipresence_rogue"] = flattenWirelessControllerWtpProfileLbsFortipresenceRogue(i["fortipresence-rogue"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_unassoc"
	if _, ok := i["fortipresence-unassoc"]; ok {
		result["fortipresence_unassoc"] = flattenWirelessControllerWtpProfileLbsFortipresenceUnassoc(i["fortipresence-unassoc"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fortipresence_ble"
	if _, ok := i["fortipresence-ble"]; ok {
		result["fortipresence_ble"] = flattenWirelessControllerWtpProfileLbsFortipresenceBle(i["fortipresence-ble"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "station_locate"
	if _, ok := i["station-locate"]; ok {
		result["station_locate"] = flattenWirelessControllerWtpProfileLbsStationLocate(i["station-locate"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileLbsEkahauBlinkMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsEkahauTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsErcServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsErcServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutApMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMmuReport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMuFactor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsAeroscoutMuTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresence(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServerAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceServerFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresencePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceProject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceFrequency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceRogue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceUnassoc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsFortipresenceBle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileLbsStationLocate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileExtInfoEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileIndoorOutdoorDeployment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongle(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "compliance_level"
	if _, ok := i["compliance-level"]; ok {
		result["compliance_level"] = flattenWirelessControllerWtpProfileEslSesDongleComplianceLevel(i["compliance-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scd_enable"
	if _, ok := i["scd-enable"]; ok {
		result["scd_enable"] = flattenWirelessControllerWtpProfileEslSesDongleScdEnable(i["scd-enable"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "esl_channel"
	if _, ok := i["esl-channel"]; ok {
		result["esl_channel"] = flattenWirelessControllerWtpProfileEslSesDongleEslChannel(i["esl-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "output_power"
	if _, ok := i["output-power"]; ok {
		result["output_power"] = flattenWirelessControllerWtpProfileEslSesDongleOutputPower(i["output-power"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "apc_addr_type"
	if _, ok := i["apc-addr-type"]; ok {
		result["apc_addr_type"] = flattenWirelessControllerWtpProfileEslSesDongleApcAddrType(i["apc-addr-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "apc_fqdn"
	if _, ok := i["apc-fqdn"]; ok {
		result["apc_fqdn"] = flattenWirelessControllerWtpProfileEslSesDongleApcFqdn(i["apc-fqdn"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "apc_ip"
	if _, ok := i["apc-ip"]; ok {
		result["apc_ip"] = flattenWirelessControllerWtpProfileEslSesDongleApcIp(i["apc-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "apc_port"
	if _, ok := i["apc-port"]; ok {
		result["apc_port"] = flattenWirelessControllerWtpProfileEslSesDongleApcPort(i["apc-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "coex_level"
	if _, ok := i["coex-level"]; ok {
		result["coex_level"] = flattenWirelessControllerWtpProfileEslSesDongleCoexLevel(i["coex-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tls_cert_verification"
	if _, ok := i["tls-cert-verification"]; ok {
		result["tls_cert_verification"] = flattenWirelessControllerWtpProfileEslSesDongleTlsCertVerification(i["tls-cert-verification"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "tls_fqdn_verification"
	if _, ok := i["tls-fqdn-verification"]; ok {
		result["tls_fqdn_verification"] = flattenWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification(i["tls-fqdn-verification"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpProfileEslSesDongleComplianceLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleScdEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleEslChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleOutputPower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleApcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleCoexLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleTlsCertVerification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileConsoleLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileWanPortAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileWanPortAuthUsrname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileWanPortAuthPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileWanPortAuthMethods(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpProfileUnii45GhzBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerWtpProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerWtpProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("platform", flattenWirelessControllerWtpProfilePlatform(o["platform"], d, "platform", sv)); err != nil {
			if !fortiAPIPatch(o["platform"]) {
				return fmt.Errorf("Error reading platform: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("platform"); ok {
			if err = d.Set("platform", flattenWirelessControllerWtpProfilePlatform(o["platform"], d, "platform", sv)); err != nil {
				if !fortiAPIPatch(o["platform"]) {
					return fmt.Errorf("Error reading platform: %v", err)
				}
			}
		}
	}

	if err = d.Set("control_message_offload", flattenWirelessControllerWtpProfileControlMessageOffload(o["control-message-offload"], d, "control_message_offload", sv)); err != nil {
		if !fortiAPIPatch(o["control-message-offload"]) {
			return fmt.Errorf("Error reading control_message_offload: %v", err)
		}
	}

	if err = d.Set("apcfg_profile", flattenWirelessControllerWtpProfileApcfgProfile(o["apcfg-profile"], d, "apcfg_profile", sv)); err != nil {
		if !fortiAPIPatch(o["apcfg-profile"]) {
			return fmt.Errorf("Error reading apcfg_profile: %v", err)
		}
	}

	if err = d.Set("ble_profile", flattenWirelessControllerWtpProfileBleProfile(o["ble-profile"], d, "ble_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ble-profile"]) {
			return fmt.Errorf("Error reading ble_profile: %v", err)
		}
	}

	if err = d.Set("syslog_profile", flattenWirelessControllerWtpProfileSyslogProfile(o["syslog-profile"], d, "syslog_profile", sv)); err != nil {
		if !fortiAPIPatch(o["syslog-profile"]) {
			return fmt.Errorf("Error reading syslog_profile: %v", err)
		}
	}

	if err = d.Set("wan_port_mode", flattenWirelessControllerWtpProfileWanPortMode(o["wan-port-mode"], d, "wan_port_mode", sv)); err != nil {
		if !fortiAPIPatch(o["wan-port-mode"]) {
			return fmt.Errorf("Error reading wan_port_mode: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("lan", flattenWirelessControllerWtpProfileLan(o["lan"], d, "lan", sv)); err != nil {
			if !fortiAPIPatch(o["lan"]) {
				return fmt.Errorf("Error reading lan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan"); ok {
			if err = d.Set("lan", flattenWirelessControllerWtpProfileLan(o["lan"], d, "lan", sv)); err != nil {
				if !fortiAPIPatch(o["lan"]) {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			}
		}
	}

	if err = d.Set("energy_efficient_ethernet", flattenWirelessControllerWtpProfileEnergyEfficientEthernet(o["energy-efficient-ethernet"], d, "energy_efficient_ethernet", sv)); err != nil {
		if !fortiAPIPatch(o["energy-efficient-ethernet"]) {
			return fmt.Errorf("Error reading energy_efficient_ethernet: %v", err)
		}
	}

	if err = d.Set("led_state", flattenWirelessControllerWtpProfileLedState(o["led-state"], d, "led_state", sv)); err != nil {
		if !fortiAPIPatch(o["led-state"]) {
			return fmt.Errorf("Error reading led_state: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("led_schedules", flattenWirelessControllerWtpProfileLedSchedules(o["led-schedules"], d, "led_schedules", sv)); err != nil {
			if !fortiAPIPatch(o["led-schedules"]) {
				return fmt.Errorf("Error reading led_schedules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("led_schedules"); ok {
			if err = d.Set("led_schedules", flattenWirelessControllerWtpProfileLedSchedules(o["led-schedules"], d, "led_schedules", sv)); err != nil {
				if !fortiAPIPatch(o["led-schedules"]) {
					return fmt.Errorf("Error reading led_schedules: %v", err)
				}
			}
		}
	}

	if err = d.Set("dtls_policy", flattenWirelessControllerWtpProfileDtlsPolicy(o["dtls-policy"], d, "dtls_policy", sv)); err != nil {
		if !fortiAPIPatch(o["dtls-policy"]) {
			return fmt.Errorf("Error reading dtls_policy: %v", err)
		}
	}

	if err = d.Set("dtls_in_kernel", flattenWirelessControllerWtpProfileDtlsInKernel(o["dtls-in-kernel"], d, "dtls_in_kernel", sv)); err != nil {
		if !fortiAPIPatch(o["dtls-in-kernel"]) {
			return fmt.Errorf("Error reading dtls_in_kernel: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerWtpProfileMaxClients(o["max-clients"], d, "max_clients", sv)); err != nil {
		if !fortiAPIPatch(o["max-clients"]) {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("handoff_rssi", flattenWirelessControllerWtpProfileHandoffRssi(o["handoff-rssi"], d, "handoff_rssi", sv)); err != nil {
		if !fortiAPIPatch(o["handoff-rssi"]) {
			return fmt.Errorf("Error reading handoff_rssi: %v", err)
		}
	}

	if err = d.Set("handoff_sta_thresh", flattenWirelessControllerWtpProfileHandoffStaThresh(o["handoff-sta-thresh"], d, "handoff_sta_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["handoff-sta-thresh"]) {
			return fmt.Errorf("Error reading handoff_sta_thresh: %v", err)
		}
	}

	if err = d.Set("handoff_roaming", flattenWirelessControllerWtpProfileHandoffRoaming(o["handoff-roaming"], d, "handoff_roaming", sv)); err != nil {
		if !fortiAPIPatch(o["handoff-roaming"]) {
			return fmt.Errorf("Error reading handoff_roaming: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("deny_mac_list", flattenWirelessControllerWtpProfileDenyMacList(o["deny-mac-list"], d, "deny_mac_list", sv)); err != nil {
			if !fortiAPIPatch(o["deny-mac-list"]) {
				return fmt.Errorf("Error reading deny_mac_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("deny_mac_list"); ok {
			if err = d.Set("deny_mac_list", flattenWirelessControllerWtpProfileDenyMacList(o["deny-mac-list"], d, "deny_mac_list", sv)); err != nil {
				if !fortiAPIPatch(o["deny-mac-list"]) {
					return fmt.Errorf("Error reading deny_mac_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("ap_country", flattenWirelessControllerWtpProfileApCountry(o["ap-country"], d, "ap_country", sv)); err != nil {
		if !fortiAPIPatch(o["ap-country"]) {
			return fmt.Errorf("Error reading ap_country: %v", err)
		}
	}

	if err = d.Set("ip_fragment_preventing", flattenWirelessControllerWtpProfileIpFragmentPreventing(o["ip-fragment-preventing"], d, "ip_fragment_preventing", sv)); err != nil {
		if !fortiAPIPatch(o["ip-fragment-preventing"]) {
			return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
		}
	}

	if err = d.Set("tun_mtu_uplink", flattenWirelessControllerWtpProfileTunMtuUplink(o["tun-mtu-uplink"], d, "tun_mtu_uplink", sv)); err != nil {
		if !fortiAPIPatch(o["tun-mtu-uplink"]) {
			return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
		}
	}

	if err = d.Set("tun_mtu_downlink", flattenWirelessControllerWtpProfileTunMtuDownlink(o["tun-mtu-downlink"], d, "tun_mtu_downlink", sv)); err != nil {
		if !fortiAPIPatch(o["tun-mtu-downlink"]) {
			return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_path", flattenWirelessControllerWtpProfileSplitTunnelingAclPath(o["split-tunneling-acl-path"], d, "split_tunneling_acl_path", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-path"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_local_ap_subnet", flattenWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(o["split-tunneling-acl-local-ap-subnet"], d, "split_tunneling_acl_local_ap_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-local-ap-subnet"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpProfileSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl", sv)); err != nil {
			if !fortiAPIPatch(o["split-tunneling-acl"]) {
				return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_acl"); ok {
			if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpProfileSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl", sv)); err != nil {
				if !fortiAPIPatch(o["split-tunneling-acl"]) {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			}
		}
	}

	if err = d.Set("allowaccess", flattenWirelessControllerWtpProfileAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("login_passwd_change", flattenWirelessControllerWtpProfileLoginPasswdChange(o["login-passwd-change"], d, "login_passwd_change", sv)); err != nil {
		if !fortiAPIPatch(o["login-passwd-change"]) {
			return fmt.Errorf("Error reading login_passwd_change: %v", err)
		}
	}

	if err = d.Set("lldp", flattenWirelessControllerWtpProfileLldp(o["lldp"], d, "lldp", sv)); err != nil {
		if !fortiAPIPatch(o["lldp"]) {
			return fmt.Errorf("Error reading lldp: %v", err)
		}
	}

	if err = d.Set("poe_mode", flattenWirelessControllerWtpProfilePoeMode(o["poe-mode"], d, "poe_mode", sv)); err != nil {
		if !fortiAPIPatch(o["poe-mode"]) {
			return fmt.Errorf("Error reading poe_mode: %v", err)
		}
	}

	if err = d.Set("frequency_handoff", flattenWirelessControllerWtpProfileFrequencyHandoff(o["frequency-handoff"], d, "frequency_handoff", sv)); err != nil {
		if !fortiAPIPatch(o["frequency-handoff"]) {
			return fmt.Errorf("Error reading frequency_handoff: %v", err)
		}
	}

	if err = d.Set("ap_handoff", flattenWirelessControllerWtpProfileApHandoff(o["ap-handoff"], d, "ap_handoff", sv)); err != nil {
		if !fortiAPIPatch(o["ap-handoff"]) {
			return fmt.Errorf("Error reading ap_handoff: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("radio_1", flattenWirelessControllerWtpProfileRadio1(o["radio-1"], d, "radio_1", sv)); err != nil {
			if !fortiAPIPatch(o["radio-1"]) {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenWirelessControllerWtpProfileRadio1(o["radio-1"], d, "radio_1", sv)); err != nil {
				if !fortiAPIPatch(o["radio-1"]) {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("radio_2", flattenWirelessControllerWtpProfileRadio2(o["radio-2"], d, "radio_2", sv)); err != nil {
			if !fortiAPIPatch(o["radio-2"]) {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenWirelessControllerWtpProfileRadio2(o["radio-2"], d, "radio_2", sv)); err != nil {
				if !fortiAPIPatch(o["radio-2"]) {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("radio_3", flattenWirelessControllerWtpProfileRadio3(o["radio-3"], d, "radio_3", sv)); err != nil {
			if !fortiAPIPatch(o["radio-3"]) {
				return fmt.Errorf("Error reading radio_3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_3"); ok {
			if err = d.Set("radio_3", flattenWirelessControllerWtpProfileRadio3(o["radio-3"], d, "radio_3", sv)); err != nil {
				if !fortiAPIPatch(o["radio-3"]) {
					return fmt.Errorf("Error reading radio_3: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("radio_4", flattenWirelessControllerWtpProfileRadio4(o["radio-4"], d, "radio_4", sv)); err != nil {
			if !fortiAPIPatch(o["radio-4"]) {
				return fmt.Errorf("Error reading radio_4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_4"); ok {
			if err = d.Set("radio_4", flattenWirelessControllerWtpProfileRadio4(o["radio-4"], d, "radio_4", sv)); err != nil {
				if !fortiAPIPatch(o["radio-4"]) {
					return fmt.Errorf("Error reading radio_4: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("lbs", flattenWirelessControllerWtpProfileLbs(o["lbs"], d, "lbs", sv)); err != nil {
			if !fortiAPIPatch(o["lbs"]) {
				return fmt.Errorf("Error reading lbs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lbs"); ok {
			if err = d.Set("lbs", flattenWirelessControllerWtpProfileLbs(o["lbs"], d, "lbs", sv)); err != nil {
				if !fortiAPIPatch(o["lbs"]) {
					return fmt.Errorf("Error reading lbs: %v", err)
				}
			}
		}
	}

	if err = d.Set("ext_info_enable", flattenWirelessControllerWtpProfileExtInfoEnable(o["ext-info-enable"], d, "ext_info_enable", sv)); err != nil {
		if !fortiAPIPatch(o["ext-info-enable"]) {
			return fmt.Errorf("Error reading ext_info_enable: %v", err)
		}
	}

	if err = d.Set("indoor_outdoor_deployment", flattenWirelessControllerWtpProfileIndoorOutdoorDeployment(o["indoor-outdoor-deployment"], d, "indoor_outdoor_deployment", sv)); err != nil {
		if !fortiAPIPatch(o["indoor-outdoor-deployment"]) {
			return fmt.Errorf("Error reading indoor_outdoor_deployment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("esl_ses_dongle", flattenWirelessControllerWtpProfileEslSesDongle(o["esl-ses-dongle"], d, "esl_ses_dongle", sv)); err != nil {
			if !fortiAPIPatch(o["esl-ses-dongle"]) {
				return fmt.Errorf("Error reading esl_ses_dongle: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("esl_ses_dongle"); ok {
			if err = d.Set("esl_ses_dongle", flattenWirelessControllerWtpProfileEslSesDongle(o["esl-ses-dongle"], d, "esl_ses_dongle", sv)); err != nil {
				if !fortiAPIPatch(o["esl-ses-dongle"]) {
					return fmt.Errorf("Error reading esl_ses_dongle: %v", err)
				}
			}
		}
	}

	if err = d.Set("console_login", flattenWirelessControllerWtpProfileConsoleLogin(o["console-login"], d, "console_login", sv)); err != nil {
		if !fortiAPIPatch(o["console-login"]) {
			return fmt.Errorf("Error reading console_login: %v", err)
		}
	}

	if err = d.Set("wan_port_auth", flattenWirelessControllerWtpProfileWanPortAuth(o["wan-port-auth"], d, "wan_port_auth", sv)); err != nil {
		if !fortiAPIPatch(o["wan-port-auth"]) {
			return fmt.Errorf("Error reading wan_port_auth: %v", err)
		}
	}

	if err = d.Set("wan_port_auth_usrname", flattenWirelessControllerWtpProfileWanPortAuthUsrname(o["wan-port-auth-usrname"], d, "wan_port_auth_usrname", sv)); err != nil {
		if !fortiAPIPatch(o["wan-port-auth-usrname"]) {
			return fmt.Errorf("Error reading wan_port_auth_usrname: %v", err)
		}
	}

	if err = d.Set("wan_port_auth_password", flattenWirelessControllerWtpProfileWanPortAuthPassword(o["wan-port-auth-password"], d, "wan_port_auth_password", sv)); err != nil {
		if !fortiAPIPatch(o["wan-port-auth-password"]) {
			return fmt.Errorf("Error reading wan_port_auth_password: %v", err)
		}
	}

	if err = d.Set("wan_port_auth_methods", flattenWirelessControllerWtpProfileWanPortAuthMethods(o["wan-port-auth-methods"], d, "wan_port_auth_methods", sv)); err != nil {
		if !fortiAPIPatch(o["wan-port-auth-methods"]) {
			return fmt.Errorf("Error reading wan_port_auth_methods: %v", err)
		}
	}

	if err = d.Set("unii_4_5ghz_band", flattenWirelessControllerWtpProfileUnii45GhzBand(o["unii-4-5ghz-band"], d, "unii_4_5ghz_band", sv)); err != nil {
		if !fortiAPIPatch(o["unii-4-5ghz-band"]) {
			return fmt.Errorf("Error reading unii_4_5ghz_band: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerWtpProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatform(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok {
		result["type"], _ = expandWirelessControllerWtpProfilePlatformType(d, i["type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfilePlatformMode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ddscan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ddscan"], _ = expandWirelessControllerWtpProfilePlatformDdscan(d, i["ddscan"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfilePlatformType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatformMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePlatformDdscan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileControlMessageOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileApcfgProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileBleProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSyslogProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileWanPortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-mode"], _ = expandWirelessControllerWtpProfileLanPortMode(d, i["port_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-ssid"], _ = expandWirelessControllerWtpProfileLanPortSsid(d, i["port_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port1_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-mode"], _ = expandWirelessControllerWtpProfileLanPort1Mode(d, i["port1_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-ssid"], _ = expandWirelessControllerWtpProfileLanPort1Ssid(d, i["port1_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port2_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-mode"], _ = expandWirelessControllerWtpProfileLanPort2Mode(d, i["port2_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-ssid"], _ = expandWirelessControllerWtpProfileLanPort2Ssid(d, i["port2_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port3_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-mode"], _ = expandWirelessControllerWtpProfileLanPort3Mode(d, i["port3_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-ssid"], _ = expandWirelessControllerWtpProfileLanPort3Ssid(d, i["port3_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port4_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-mode"], _ = expandWirelessControllerWtpProfileLanPort4Mode(d, i["port4_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-ssid"], _ = expandWirelessControllerWtpProfileLanPort4Ssid(d, i["port4_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port5_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-mode"], _ = expandWirelessControllerWtpProfileLanPort5Mode(d, i["port5_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-ssid"], _ = expandWirelessControllerWtpProfileLanPort5Ssid(d, i["port5_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port6_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-mode"], _ = expandWirelessControllerWtpProfileLanPort6Mode(d, i["port6_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-ssid"], _ = expandWirelessControllerWtpProfileLanPort6Ssid(d, i["port6_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port7_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-mode"], _ = expandWirelessControllerWtpProfileLanPort7Mode(d, i["port7_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-ssid"], _ = expandWirelessControllerWtpProfileLanPort7Ssid(d, i["port7_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port8_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-mode"], _ = expandWirelessControllerWtpProfileLanPort8Mode(d, i["port8_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-ssid"], _ = expandWirelessControllerWtpProfileLanPort8Ssid(d, i["port8_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-esl-mode"], _ = expandWirelessControllerWtpProfileLanPortEslMode(d, i["port_esl_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-esl-ssid"], _ = expandWirelessControllerWtpProfileLanPortEslSsid(d, i["port_esl_ssid"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileLanPortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPortSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort1Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort1Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort2Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort2Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort3Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort3Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort4Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort4Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort5Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort5Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort6Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort6Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort7Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort7Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort8Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPort8Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPortEslMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLanPortEslSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEnergyEfficientEthernet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLedState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLedSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileLedSchedulesName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileLedSchedulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDtlsPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDtlsInKernel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileMaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileHandoffRssi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileHandoffStaThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileHandoffRoaming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDenyMacList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerWtpProfileDenyMacListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandWirelessControllerWtpProfileDenyMacListMac(d, i["mac"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileDenyMacListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileDenyMacListMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileApCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileIpFragmentPreventing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileTunMtuUplink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileTunMtuDownlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerWtpProfileSplitTunnelingAclId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dest-ip"], _ = expandWirelessControllerWtpProfileSplitTunnelingAclDestIp(d, i["dest_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileSplitTunnelingAclDestIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLoginPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLldp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfilePoeMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileFrequencyHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileApHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandWirelessControllerWtpProfileRadio1RadioId(d, i["radio_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfileRadio1Mode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpProfileRadio1Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandWirelessControllerWtpProfileRadio1Band5GType(d, i["band_5g_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandWirelessControllerWtpProfileRadio1Drma(d, i["drma"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandWirelessControllerWtpProfileRadio1DrmaSensitivity(d, i["drma_sensitivity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandWirelessControllerWtpProfileRadio1AirtimeFairness(d, i["airtime_fairness"], pre_append, sv)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandWirelessControllerWtpProfileRadio1ProtectionMode(d, i["protection_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandWirelessControllerWtpProfileRadio1PowersaveOptimize(d, i["powersave_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandWirelessControllerWtpProfileRadio1TransmitOptimize(d, i["transmit_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandWirelessControllerWtpProfileRadio1Amsdu(d, i["amsdu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandWirelessControllerWtpProfileRadio1Coexistence(d, i["coexistence"], pre_append, sv)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandWirelessControllerWtpProfileRadio1ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandWirelessControllerWtpProfileRadio1BssColor(d, i["bss_color"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color-mode"], _ = expandWirelessControllerWtpProfileRadio1BssColorMode(d, i["bss_color_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandWirelessControllerWtpProfileRadio1ShortGuardInterval(d, i["short_guard_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandWirelessControllerWtpProfileRadio1ChannelBonding(d, i["channel_bonding"], pre_append, sv)
	}
	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := d.GetOk(pre_append); ok {
		result["optional-antenna"], _ = expandWirelessControllerWtpProfileRadio1OptionalAntenna(d, i["optional_antenna"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandWirelessControllerWtpProfileRadio1AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandWirelessControllerWtpProfileRadio1PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpProfileRadio1PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandWirelessControllerWtpProfileRadio1PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandWirelessControllerWtpProfileRadio1Dtim(d, i["dtim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandWirelessControllerWtpProfileRadio1BeaconInterval(d, i["beacon_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok {
		result["80211d"], _ = expandWirelessControllerWtpProfileRadio180211D(d, i["n80211d"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandWirelessControllerWtpProfileRadio1RtsThreshold(d, i["rts_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandWirelessControllerWtpProfileRadio1FragThreshold(d, i["frag_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferChan(d, i["ap_sniffer_chan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandWirelessControllerWtpProfileRadio1ApSnifferData(d, i["ap_sniffer_data"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandWirelessControllerWtpProfileRadio1SamSsid(d, i["sam_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandWirelessControllerWtpProfileRadio1SamBssid(d, i["sam_bssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandWirelessControllerWtpProfileRadio1SamSecurityType(d, i["sam_security_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandWirelessControllerWtpProfileRadio1SamCaptivePortal(d, i["sam_captive_portal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-username"], _ = expandWirelessControllerWtpProfileRadio1SamCwpUsername(d, i["sam_cwp_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-password"], _ = expandWirelessControllerWtpProfileRadio1SamCwpPassword(d, i["sam_cwp_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-test-url"], _ = expandWirelessControllerWtpProfileRadio1SamCwpTestUrl(d, i["sam_cwp_test_url"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-match-string"], _ = expandWirelessControllerWtpProfileRadio1SamCwpMatchString(d, i["sam_cwp_match_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-success-string"], _ = expandWirelessControllerWtpProfileRadio1SamCwpSuccessString(d, i["sam_cwp_success_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-failure-string"], _ = expandWirelessControllerWtpProfileRadio1SamCwpFailureString(d, i["sam_cwp_failure_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandWirelessControllerWtpProfileRadio1SamUsername(d, i["sam_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandWirelessControllerWtpProfileRadio1SamPassword(d, i["sam_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandWirelessControllerWtpProfileRadio1SamTest(d, i["sam_test"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-type"], _ = expandWirelessControllerWtpProfileRadio1SamServerType(d, i["sam_server_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-ip"], _ = expandWirelessControllerWtpProfileRadio1SamServerIp(d, i["sam_server_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-fqdn"], _ = expandWirelessControllerWtpProfileRadio1SamServerFqdn(d, i["sam_server_fqdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandWirelessControllerWtpProfileRadio1IperfServerPort(d, i["iperf_server_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandWirelessControllerWtpProfileRadio1IperfProtocol(d, i["iperf_protocol"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandWirelessControllerWtpProfileRadio1SamReportIntv(d, i["sam_report_intv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandWirelessControllerWtpProfileRadio1ChannelUtilization(d, i["channel_utilization"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpProfileRadio1SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandWirelessControllerWtpProfileRadio1WidsProfile(d, i["wids_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandWirelessControllerWtpProfileRadio1Darrp(d, i["darrp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["arrp-profile"], _ = expandWirelessControllerWtpProfileRadio1ArrpProfile(d, i["arrp_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandWirelessControllerWtpProfileRadio1MaxClients(d, i["max_clients"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandWirelessControllerWtpProfileRadio1MaxDistance(d, i["max_distance"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["frequency-handoff"], _ = expandWirelessControllerWtpProfileRadio1FrequencyHandoff(d, i["frequency_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-handoff"], _ = expandWirelessControllerWtpProfileRadio1ApHandoff(d, i["ap_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpProfileRadio1VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpProfileRadio1Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpProfileRadio1Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandWirelessControllerWtpProfileRadio1CallAdmissionControl(d, i["call_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandWirelessControllerWtpProfileRadio1CallCapacity(d, i["call_capacity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandWirelessControllerWtpProfileRadio1BandwidthCapacity(d, i["bandwidth_capacity"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio1RadioId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Band5GType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Drma(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AirtimeFairness(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ProtectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1TransmitOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Amsdu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Coexistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BssColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BssColorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ChannelBonding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1OptionalAntenna(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Dtim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio180211D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1RtsThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1FragThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApSnifferData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamBssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamSecurityType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCwpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCwpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCwpTestUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCwpMatchString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCwpSuccessString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamCwpFailureString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamTest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1IperfServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1IperfProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SamReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ChannelUtilization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1WidsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Darrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ArrpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1MaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1MaxDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1FrequencyHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1ApHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileRadio1VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio1VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chan"], _ = expandWirelessControllerWtpProfileRadio1ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio1ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1CallCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio1BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandWirelessControllerWtpProfileRadio2RadioId(d, i["radio_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfileRadio2Mode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpProfileRadio2Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandWirelessControllerWtpProfileRadio2Band5GType(d, i["band_5g_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandWirelessControllerWtpProfileRadio2Drma(d, i["drma"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandWirelessControllerWtpProfileRadio2DrmaSensitivity(d, i["drma_sensitivity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandWirelessControllerWtpProfileRadio2AirtimeFairness(d, i["airtime_fairness"], pre_append, sv)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandWirelessControllerWtpProfileRadio2ProtectionMode(d, i["protection_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandWirelessControllerWtpProfileRadio2PowersaveOptimize(d, i["powersave_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandWirelessControllerWtpProfileRadio2TransmitOptimize(d, i["transmit_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandWirelessControllerWtpProfileRadio2Amsdu(d, i["amsdu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandWirelessControllerWtpProfileRadio2Coexistence(d, i["coexistence"], pre_append, sv)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandWirelessControllerWtpProfileRadio2ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandWirelessControllerWtpProfileRadio2BssColor(d, i["bss_color"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color-mode"], _ = expandWirelessControllerWtpProfileRadio2BssColorMode(d, i["bss_color_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandWirelessControllerWtpProfileRadio2ShortGuardInterval(d, i["short_guard_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandWirelessControllerWtpProfileRadio2ChannelBonding(d, i["channel_bonding"], pre_append, sv)
	}
	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := d.GetOk(pre_append); ok {
		result["optional-antenna"], _ = expandWirelessControllerWtpProfileRadio2OptionalAntenna(d, i["optional_antenna"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandWirelessControllerWtpProfileRadio2AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandWirelessControllerWtpProfileRadio2PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpProfileRadio2PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandWirelessControllerWtpProfileRadio2PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandWirelessControllerWtpProfileRadio2Dtim(d, i["dtim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandWirelessControllerWtpProfileRadio2BeaconInterval(d, i["beacon_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok {
		result["80211d"], _ = expandWirelessControllerWtpProfileRadio280211D(d, i["n80211d"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandWirelessControllerWtpProfileRadio2RtsThreshold(d, i["rts_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandWirelessControllerWtpProfileRadio2FragThreshold(d, i["frag_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferChan(d, i["ap_sniffer_chan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandWirelessControllerWtpProfileRadio2ApSnifferData(d, i["ap_sniffer_data"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandWirelessControllerWtpProfileRadio2SamSsid(d, i["sam_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandWirelessControllerWtpProfileRadio2SamBssid(d, i["sam_bssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandWirelessControllerWtpProfileRadio2SamSecurityType(d, i["sam_security_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandWirelessControllerWtpProfileRadio2SamCaptivePortal(d, i["sam_captive_portal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-username"], _ = expandWirelessControllerWtpProfileRadio2SamCwpUsername(d, i["sam_cwp_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-password"], _ = expandWirelessControllerWtpProfileRadio2SamCwpPassword(d, i["sam_cwp_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-test-url"], _ = expandWirelessControllerWtpProfileRadio2SamCwpTestUrl(d, i["sam_cwp_test_url"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-match-string"], _ = expandWirelessControllerWtpProfileRadio2SamCwpMatchString(d, i["sam_cwp_match_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-success-string"], _ = expandWirelessControllerWtpProfileRadio2SamCwpSuccessString(d, i["sam_cwp_success_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-failure-string"], _ = expandWirelessControllerWtpProfileRadio2SamCwpFailureString(d, i["sam_cwp_failure_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandWirelessControllerWtpProfileRadio2SamUsername(d, i["sam_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandWirelessControllerWtpProfileRadio2SamPassword(d, i["sam_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandWirelessControllerWtpProfileRadio2SamTest(d, i["sam_test"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-type"], _ = expandWirelessControllerWtpProfileRadio2SamServerType(d, i["sam_server_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-ip"], _ = expandWirelessControllerWtpProfileRadio2SamServerIp(d, i["sam_server_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-fqdn"], _ = expandWirelessControllerWtpProfileRadio2SamServerFqdn(d, i["sam_server_fqdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandWirelessControllerWtpProfileRadio2IperfServerPort(d, i["iperf_server_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandWirelessControllerWtpProfileRadio2IperfProtocol(d, i["iperf_protocol"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandWirelessControllerWtpProfileRadio2SamReportIntv(d, i["sam_report_intv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandWirelessControllerWtpProfileRadio2ChannelUtilization(d, i["channel_utilization"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpProfileRadio2SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandWirelessControllerWtpProfileRadio2WidsProfile(d, i["wids_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandWirelessControllerWtpProfileRadio2Darrp(d, i["darrp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["arrp-profile"], _ = expandWirelessControllerWtpProfileRadio2ArrpProfile(d, i["arrp_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandWirelessControllerWtpProfileRadio2MaxClients(d, i["max_clients"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandWirelessControllerWtpProfileRadio2MaxDistance(d, i["max_distance"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["frequency-handoff"], _ = expandWirelessControllerWtpProfileRadio2FrequencyHandoff(d, i["frequency_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-handoff"], _ = expandWirelessControllerWtpProfileRadio2ApHandoff(d, i["ap_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpProfileRadio2VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpProfileRadio2Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpProfileRadio2Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandWirelessControllerWtpProfileRadio2CallAdmissionControl(d, i["call_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandWirelessControllerWtpProfileRadio2CallCapacity(d, i["call_capacity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandWirelessControllerWtpProfileRadio2BandwidthCapacity(d, i["bandwidth_capacity"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio2RadioId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Band5GType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Drma(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AirtimeFairness(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ProtectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2TransmitOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Amsdu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Coexistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BssColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BssColorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelBonding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2OptionalAntenna(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Dtim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio280211D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2RtsThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2FragThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApSnifferData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamBssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamSecurityType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpTestUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpMatchString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpSuccessString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamCwpFailureString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamTest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2IperfServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2IperfProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SamReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelUtilization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2WidsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Darrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ArrpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2MaxDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2FrequencyHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2ApHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileRadio2VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio2VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chan"], _ = expandWirelessControllerWtpProfileRadio2ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio2ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2CallCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio2BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfileRadio3Mode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpProfileRadio3Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandWirelessControllerWtpProfileRadio3Band5GType(d, i["band_5g_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandWirelessControllerWtpProfileRadio3Drma(d, i["drma"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandWirelessControllerWtpProfileRadio3DrmaSensitivity(d, i["drma_sensitivity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandWirelessControllerWtpProfileRadio3AirtimeFairness(d, i["airtime_fairness"], pre_append, sv)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandWirelessControllerWtpProfileRadio3ProtectionMode(d, i["protection_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandWirelessControllerWtpProfileRadio3PowersaveOptimize(d, i["powersave_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandWirelessControllerWtpProfileRadio3TransmitOptimize(d, i["transmit_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandWirelessControllerWtpProfileRadio3Amsdu(d, i["amsdu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandWirelessControllerWtpProfileRadio3Coexistence(d, i["coexistence"], pre_append, sv)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandWirelessControllerWtpProfileRadio3ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandWirelessControllerWtpProfileRadio3BssColor(d, i["bss_color"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color-mode"], _ = expandWirelessControllerWtpProfileRadio3BssColorMode(d, i["bss_color_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandWirelessControllerWtpProfileRadio3ShortGuardInterval(d, i["short_guard_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandWirelessControllerWtpProfileRadio3ChannelBonding(d, i["channel_bonding"], pre_append, sv)
	}
	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := d.GetOk(pre_append); ok {
		result["optional-antenna"], _ = expandWirelessControllerWtpProfileRadio3OptionalAntenna(d, i["optional_antenna"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpProfileRadio3AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpProfileRadio3AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpProfileRadio3AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandWirelessControllerWtpProfileRadio3AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandWirelessControllerWtpProfileRadio3PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpProfileRadio3PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandWirelessControllerWtpProfileRadio3PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandWirelessControllerWtpProfileRadio3Dtim(d, i["dtim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandWirelessControllerWtpProfileRadio3BeaconInterval(d, i["beacon_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok {
		result["80211d"], _ = expandWirelessControllerWtpProfileRadio380211D(d, i["n80211d"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandWirelessControllerWtpProfileRadio3RtsThreshold(d, i["rts_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandWirelessControllerWtpProfileRadio3FragThreshold(d, i["frag_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferChan(d, i["ap_sniffer_chan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandWirelessControllerWtpProfileRadio3ApSnifferData(d, i["ap_sniffer_data"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandWirelessControllerWtpProfileRadio3SamSsid(d, i["sam_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandWirelessControllerWtpProfileRadio3SamBssid(d, i["sam_bssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandWirelessControllerWtpProfileRadio3SamSecurityType(d, i["sam_security_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandWirelessControllerWtpProfileRadio3SamCaptivePortal(d, i["sam_captive_portal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-username"], _ = expandWirelessControllerWtpProfileRadio3SamCwpUsername(d, i["sam_cwp_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-password"], _ = expandWirelessControllerWtpProfileRadio3SamCwpPassword(d, i["sam_cwp_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-test-url"], _ = expandWirelessControllerWtpProfileRadio3SamCwpTestUrl(d, i["sam_cwp_test_url"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-match-string"], _ = expandWirelessControllerWtpProfileRadio3SamCwpMatchString(d, i["sam_cwp_match_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-success-string"], _ = expandWirelessControllerWtpProfileRadio3SamCwpSuccessString(d, i["sam_cwp_success_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-failure-string"], _ = expandWirelessControllerWtpProfileRadio3SamCwpFailureString(d, i["sam_cwp_failure_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandWirelessControllerWtpProfileRadio3SamUsername(d, i["sam_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandWirelessControllerWtpProfileRadio3SamPassword(d, i["sam_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandWirelessControllerWtpProfileRadio3SamTest(d, i["sam_test"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-type"], _ = expandWirelessControllerWtpProfileRadio3SamServerType(d, i["sam_server_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-ip"], _ = expandWirelessControllerWtpProfileRadio3SamServerIp(d, i["sam_server_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-fqdn"], _ = expandWirelessControllerWtpProfileRadio3SamServerFqdn(d, i["sam_server_fqdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandWirelessControllerWtpProfileRadio3IperfServerPort(d, i["iperf_server_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandWirelessControllerWtpProfileRadio3IperfProtocol(d, i["iperf_protocol"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandWirelessControllerWtpProfileRadio3SamReportIntv(d, i["sam_report_intv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandWirelessControllerWtpProfileRadio3ChannelUtilization(d, i["channel_utilization"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpProfileRadio3SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandWirelessControllerWtpProfileRadio3WidsProfile(d, i["wids_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandWirelessControllerWtpProfileRadio3Darrp(d, i["darrp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["arrp-profile"], _ = expandWirelessControllerWtpProfileRadio3ArrpProfile(d, i["arrp_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandWirelessControllerWtpProfileRadio3MaxClients(d, i["max_clients"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandWirelessControllerWtpProfileRadio3MaxDistance(d, i["max_distance"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["frequency-handoff"], _ = expandWirelessControllerWtpProfileRadio3FrequencyHandoff(d, i["frequency_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-handoff"], _ = expandWirelessControllerWtpProfileRadio3ApHandoff(d, i["ap_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpProfileRadio3VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpProfileRadio3Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpProfileRadio3Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandWirelessControllerWtpProfileRadio3CallAdmissionControl(d, i["call_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandWirelessControllerWtpProfileRadio3CallCapacity(d, i["call_capacity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandWirelessControllerWtpProfileRadio3BandwidthCapacity(d, i["bandwidth_capacity"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio3Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Band5GType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Drma(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3AirtimeFairness(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ProtectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3TransmitOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Amsdu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Coexistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3BssColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3BssColorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ChannelBonding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3OptionalAntenna(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Dtim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3BeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio380211D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3RtsThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3FragThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApSnifferData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamBssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamSecurityType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCwpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCwpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCwpTestUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCwpMatchString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCwpSuccessString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamCwpFailureString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamTest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3IperfServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3IperfProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SamReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ChannelUtilization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3WidsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Darrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ArrpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3MaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3MaxDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3FrequencyHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3ApHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileRadio3VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio3VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chan"], _ = expandWirelessControllerWtpProfileRadio3ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio3ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3CallCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio3BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandWirelessControllerWtpProfileRadio4Mode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpProfileRadio4Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band_5g_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["band-5g-type"], _ = expandWirelessControllerWtpProfileRadio4Band5GType(d, i["band_5g_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma"], _ = expandWirelessControllerWtpProfileRadio4Drma(d, i["drma"], pre_append, sv)
	}
	pre_append = pre + ".0." + "drma_sensitivity"
	if _, ok := d.GetOk(pre_append); ok {
		result["drma-sensitivity"], _ = expandWirelessControllerWtpProfileRadio4DrmaSensitivity(d, i["drma_sensitivity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "airtime_fairness"
	if _, ok := d.GetOk(pre_append); ok {
		result["airtime-fairness"], _ = expandWirelessControllerWtpProfileRadio4AirtimeFairness(d, i["airtime_fairness"], pre_append, sv)
	}
	pre_append = pre + ".0." + "protection_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["protection-mode"], _ = expandWirelessControllerWtpProfileRadio4ProtectionMode(d, i["protection_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "powersave_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["powersave-optimize"], _ = expandWirelessControllerWtpProfileRadio4PowersaveOptimize(d, i["powersave_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "transmit_optimize"
	if _, ok := d.GetOk(pre_append); ok {
		result["transmit-optimize"], _ = expandWirelessControllerWtpProfileRadio4TransmitOptimize(d, i["transmit_optimize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "amsdu"
	if _, ok := d.GetOk(pre_append); ok {
		result["amsdu"], _ = expandWirelessControllerWtpProfileRadio4Amsdu(d, i["amsdu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "coexistence"
	if _, ok := d.GetOk(pre_append); ok {
		result["coexistence"], _ = expandWirelessControllerWtpProfileRadio4Coexistence(d, i["coexistence"], pre_append, sv)
	}
	pre_append = pre + ".0." + "zero_wait_dfs"
	if _, ok := d.GetOk(pre_append); ok {
		result["zero-wait-dfs"], _ = expandWirelessControllerWtpProfileRadio4ZeroWaitDfs(d, i["zero_wait_dfs"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandWirelessControllerWtpProfileRadio4BssColor(d, i["bss_color"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color-mode"], _ = expandWirelessControllerWtpProfileRadio4BssColorMode(d, i["bss_color_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "short_guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["short-guard-interval"], _ = expandWirelessControllerWtpProfileRadio4ShortGuardInterval(d, i["short_guard_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_bonding"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-bonding"], _ = expandWirelessControllerWtpProfileRadio4ChannelBonding(d, i["channel_bonding"], pre_append, sv)
	}
	pre_append = pre + ".0." + "optional_antenna"
	if _, ok := d.GetOk(pre_append); ok {
		result["optional-antenna"], _ = expandWirelessControllerWtpProfileRadio4OptionalAntenna(d, i["optional_antenna"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpProfileRadio4AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpProfileRadio4AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpProfileRadio4AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-target"], _ = expandWirelessControllerWtpProfileRadio4AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-mode"], _ = expandWirelessControllerWtpProfileRadio4PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpProfileRadio4PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-value"], _ = expandWirelessControllerWtpProfileRadio4PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dtim"
	if _, ok := d.GetOk(pre_append); ok {
		result["dtim"], _ = expandWirelessControllerWtpProfileRadio4Dtim(d, i["dtim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandWirelessControllerWtpProfileRadio4BeaconInterval(d, i["beacon_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok {
		result["80211d"], _ = expandWirelessControllerWtpProfileRadio480211D(d, i["n80211d"], pre_append, sv)
	}
	pre_append = pre + ".0." + "rts_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["rts-threshold"], _ = expandWirelessControllerWtpProfileRadio4RtsThreshold(d, i["rts_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frag_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["frag-threshold"], _ = expandWirelessControllerWtpProfileRadio4FragThreshold(d, i["frag_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_bufsize"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-bufsize"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferBufsize(d, i["ap_sniffer_bufsize"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_chan"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-chan"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferChan(d, i["ap_sniffer_chan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-addr"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferAddr(d, i["ap_sniffer_addr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_beacon"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-beacon"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(d, i["ap_sniffer_mgmt_beacon"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_probe"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-probe"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(d, i["ap_sniffer_mgmt_probe"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_mgmt_other"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-mgmt-other"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(d, i["ap_sniffer_mgmt_other"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_ctl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-ctl"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferCtl(d, i["ap_sniffer_ctl"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_sniffer_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-sniffer-data"], _ = expandWirelessControllerWtpProfileRadio4ApSnifferData(d, i["ap_sniffer_data"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-ssid"], _ = expandWirelessControllerWtpProfileRadio4SamSsid(d, i["sam_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_bssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-bssid"], _ = expandWirelessControllerWtpProfileRadio4SamBssid(d, i["sam_bssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_security_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-security-type"], _ = expandWirelessControllerWtpProfileRadio4SamSecurityType(d, i["sam_security_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_captive_portal"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-captive-portal"], _ = expandWirelessControllerWtpProfileRadio4SamCaptivePortal(d, i["sam_captive_portal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-username"], _ = expandWirelessControllerWtpProfileRadio4SamCwpUsername(d, i["sam_cwp_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-password"], _ = expandWirelessControllerWtpProfileRadio4SamCwpPassword(d, i["sam_cwp_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_test_url"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-test-url"], _ = expandWirelessControllerWtpProfileRadio4SamCwpTestUrl(d, i["sam_cwp_test_url"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_match_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-match-string"], _ = expandWirelessControllerWtpProfileRadio4SamCwpMatchString(d, i["sam_cwp_match_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_success_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-success-string"], _ = expandWirelessControllerWtpProfileRadio4SamCwpSuccessString(d, i["sam_cwp_success_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_cwp_failure_string"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-cwp-failure-string"], _ = expandWirelessControllerWtpProfileRadio4SamCwpFailureString(d, i["sam_cwp_failure_string"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-username"], _ = expandWirelessControllerWtpProfileRadio4SamUsername(d, i["sam_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-password"], _ = expandWirelessControllerWtpProfileRadio4SamPassword(d, i["sam_password"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_test"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-test"], _ = expandWirelessControllerWtpProfileRadio4SamTest(d, i["sam_test"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-type"], _ = expandWirelessControllerWtpProfileRadio4SamServerType(d, i["sam_server_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-ip"], _ = expandWirelessControllerWtpProfileRadio4SamServerIp(d, i["sam_server_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_server_fqdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-server-fqdn"], _ = expandWirelessControllerWtpProfileRadio4SamServerFqdn(d, i["sam_server_fqdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-server-port"], _ = expandWirelessControllerWtpProfileRadio4IperfServerPort(d, i["iperf_server_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "iperf_protocol"
	if _, ok := d.GetOk(pre_append); ok {
		result["iperf-protocol"], _ = expandWirelessControllerWtpProfileRadio4IperfProtocol(d, i["iperf_protocol"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sam_report_intv"
	if _, ok := d.GetOk(pre_append); ok {
		result["sam-report-intv"], _ = expandWirelessControllerWtpProfileRadio4SamReportIntv(d, i["sam_report_intv"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel_utilization"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel-utilization"], _ = expandWirelessControllerWtpProfileRadio4ChannelUtilization(d, i["channel_utilization"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpProfileRadio4SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "wids_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["wids-profile"], _ = expandWirelessControllerWtpProfileRadio4WidsProfile(d, i["wids_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "darrp"
	if _, ok := d.GetOk(pre_append); ok {
		result["darrp"], _ = expandWirelessControllerWtpProfileRadio4Darrp(d, i["darrp"], pre_append, sv)
	}
	pre_append = pre + ".0." + "arrp_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["arrp-profile"], _ = expandWirelessControllerWtpProfileRadio4ArrpProfile(d, i["arrp_profile"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandWirelessControllerWtpProfileRadio4MaxClients(d, i["max_clients"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_distance"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-distance"], _ = expandWirelessControllerWtpProfileRadio4MaxDistance(d, i["max_distance"], pre_append, sv)
	}
	pre_append = pre + ".0." + "frequency_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["frequency-handoff"], _ = expandWirelessControllerWtpProfileRadio4FrequencyHandoff(d, i["frequency_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ap_handoff"
	if _, ok := d.GetOk(pre_append); ok {
		result["ap-handoff"], _ = expandWirelessControllerWtpProfileRadio4ApHandoff(d, i["ap_handoff"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpProfileRadio4VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpProfileRadio4Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpProfileRadio4Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "call_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-admission-control"], _ = expandWirelessControllerWtpProfileRadio4CallAdmissionControl(d, i["call_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "call_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["call-capacity"], _ = expandWirelessControllerWtpProfileRadio4CallCapacity(d, i["call_capacity"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_admission_control"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-admission-control"], _ = expandWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(d, i["bandwidth_admission_control"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth_capacity"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth-capacity"], _ = expandWirelessControllerWtpProfileRadio4BandwidthCapacity(d, i["bandwidth_capacity"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio4Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Band5GType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Drma(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4DrmaSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4AirtimeFairness(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ProtectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4PowersaveOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4TransmitOptimize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Amsdu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Coexistence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ZeroWaitDfs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4BssColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4BssColorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ShortGuardInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ChannelBonding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4OptionalAntenna(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Dtim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4BeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio480211D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4RtsThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4FragThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferBufsize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferMgmtBeacon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferMgmtProbe(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferMgmtOther(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferCtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApSnifferData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamBssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamSecurityType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCwpUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCwpPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCwpTestUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCwpMatchString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCwpSuccessString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamCwpFailureString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamTest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4IperfServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4IperfProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SamReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ChannelUtilization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4WidsProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Darrp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ArrpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4MaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4MaxDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4FrequencyHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4ApHandoff(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpProfileRadio4VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio4VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chan"], _ = expandWirelessControllerWtpProfileRadio4ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpProfileRadio4ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4CallAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4CallCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4BandwidthAdmissionControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileRadio4BandwidthCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ekahau_blink_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ekahau-blink-mode"], _ = expandWirelessControllerWtpProfileLbsEkahauBlinkMode(d, i["ekahau_blink_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ekahau_tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ekahau-tag"], _ = expandWirelessControllerWtpProfileLbsEkahauTag(d, i["ekahau_tag"], pre_append, sv)
	}
	pre_append = pre + ".0." + "erc_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["erc-server-ip"], _ = expandWirelessControllerWtpProfileLbsErcServerIp(d, i["erc_server_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "erc_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["erc-server-port"], _ = expandWirelessControllerWtpProfileLbsErcServerPort(d, i["erc_server_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout"], _ = expandWirelessControllerWtpProfileLbsAeroscout(d, i["aeroscout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_server_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-server-ip"], _ = expandWirelessControllerWtpProfileLbsAeroscoutServerIp(d, i["aeroscout_server_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_server_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-server-port"], _ = expandWirelessControllerWtpProfileLbsAeroscoutServerPort(d, i["aeroscout_server_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_mu"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMu(d, i["aeroscout_mu"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_ap_mac"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-ap-mac"], _ = expandWirelessControllerWtpProfileLbsAeroscoutApMac(d, i["aeroscout_ap_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_mmu_report"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mmu-report"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMmuReport(d, i["aeroscout_mmu_report"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_mu_factor"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu-factor"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMuFactor(d, i["aeroscout_mu_factor"], pre_append, sv)
	}
	pre_append = pre + ".0." + "aeroscout_mu_timeout"
	if _, ok := d.GetOk(pre_append); ok {
		result["aeroscout-mu-timeout"], _ = expandWirelessControllerWtpProfileLbsAeroscoutMuTimeout(d, i["aeroscout_mu_timeout"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence"], _ = expandWirelessControllerWtpProfileLbsFortipresence(d, i["fortipresence"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_server_addr_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-server-addr-type"], _ = expandWirelessControllerWtpProfileLbsFortipresenceServerAddrType(d, i["fortipresence_server_addr_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-server"], _ = expandWirelessControllerWtpProfileLbsFortipresenceServer(d, i["fortipresence_server"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_server_fqdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-server-fqdn"], _ = expandWirelessControllerWtpProfileLbsFortipresenceServerFqdn(d, i["fortipresence_server_fqdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-port"], _ = expandWirelessControllerWtpProfileLbsFortipresencePort(d, i["fortipresence_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_secret"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-secret"], _ = expandWirelessControllerWtpProfileLbsFortipresenceSecret(d, i["fortipresence_secret"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_project"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-project"], _ = expandWirelessControllerWtpProfileLbsFortipresenceProject(d, i["fortipresence_project"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_frequency"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-frequency"], _ = expandWirelessControllerWtpProfileLbsFortipresenceFrequency(d, i["fortipresence_frequency"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_rogue"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-rogue"], _ = expandWirelessControllerWtpProfileLbsFortipresenceRogue(d, i["fortipresence_rogue"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_unassoc"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-unassoc"], _ = expandWirelessControllerWtpProfileLbsFortipresenceUnassoc(d, i["fortipresence_unassoc"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fortipresence_ble"
	if _, ok := d.GetOk(pre_append); ok {
		result["fortipresence-ble"], _ = expandWirelessControllerWtpProfileLbsFortipresenceBle(d, i["fortipresence_ble"], pre_append, sv)
	}
	pre_append = pre + ".0." + "station_locate"
	if _, ok := d.GetOk(pre_append); ok {
		result["station-locate"], _ = expandWirelessControllerWtpProfileLbsStationLocate(d, i["station_locate"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileLbsEkahauBlinkMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsEkahauTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsErcServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsErcServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutApMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMmuReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMuFactor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsAeroscoutMuTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresence(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServerAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceServerFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresencePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceProject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceFrequency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceRogue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceUnassoc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsFortipresenceBle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileLbsStationLocate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileExtInfoEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileIndoorOutdoorDeployment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "compliance_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["compliance-level"], _ = expandWirelessControllerWtpProfileEslSesDongleComplianceLevel(d, i["compliance_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scd_enable"
	if _, ok := d.GetOk(pre_append); ok {
		result["scd-enable"], _ = expandWirelessControllerWtpProfileEslSesDongleScdEnable(d, i["scd_enable"], pre_append, sv)
	}
	pre_append = pre + ".0." + "esl_channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["esl-channel"], _ = expandWirelessControllerWtpProfileEslSesDongleEslChannel(d, i["esl_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "output_power"
	if _, ok := d.GetOk(pre_append); ok {
		result["output-power"], _ = expandWirelessControllerWtpProfileEslSesDongleOutputPower(d, i["output_power"], pre_append, sv)
	}
	pre_append = pre + ".0." + "apc_addr_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["apc-addr-type"], _ = expandWirelessControllerWtpProfileEslSesDongleApcAddrType(d, i["apc_addr_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "apc_fqdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["apc-fqdn"], _ = expandWirelessControllerWtpProfileEslSesDongleApcFqdn(d, i["apc_fqdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "apc_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["apc-ip"], _ = expandWirelessControllerWtpProfileEslSesDongleApcIp(d, i["apc_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "apc_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["apc-port"], _ = expandWirelessControllerWtpProfileEslSesDongleApcPort(d, i["apc_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "coex_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["coex-level"], _ = expandWirelessControllerWtpProfileEslSesDongleCoexLevel(d, i["coex_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tls_cert_verification"
	if _, ok := d.GetOk(pre_append); ok {
		result["tls-cert-verification"], _ = expandWirelessControllerWtpProfileEslSesDongleTlsCertVerification(d, i["tls_cert_verification"], pre_append, sv)
	}
	pre_append = pre + ".0." + "tls_fqdn_verification"
	if _, ok := d.GetOk(pre_append); ok {
		result["tls-fqdn-verification"], _ = expandWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification(d, i["tls_fqdn_verification"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpProfileEslSesDongleComplianceLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleScdEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleEslChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleOutputPower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleApcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleCoexLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleTlsCertVerification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileEslSesDongleTlsFqdnVerification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileConsoleLogin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileWanPortAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileWanPortAuthUsrname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileWanPortAuthPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileWanPortAuthMethods(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpProfileUnii45GhzBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWtpProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerWtpProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("platform"); ok {
		t, err := expandWirelessControllerWtpProfilePlatform(d, v, "platform", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform"] = t
		}
	}

	if v, ok := d.GetOk("control_message_offload"); ok {
		t, err := expandWirelessControllerWtpProfileControlMessageOffload(d, v, "control_message_offload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-message-offload"] = t
		}
	}

	if v, ok := d.GetOk("apcfg_profile"); ok {
		t, err := expandWirelessControllerWtpProfileApcfgProfile(d, v, "apcfg_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apcfg-profile"] = t
		}
	}

	if v, ok := d.GetOk("ble_profile"); ok {
		t, err := expandWirelessControllerWtpProfileBleProfile(d, v, "ble_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-profile"] = t
		}
	}

	if v, ok := d.GetOk("syslog_profile"); ok {
		t, err := expandWirelessControllerWtpProfileSyslogProfile(d, v, "syslog_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syslog-profile"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_mode"); ok {
		t, err := expandWirelessControllerWtpProfileWanPortMode(d, v, "wan_port_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("lan"); ok {
		t, err := expandWirelessControllerWtpProfileLan(d, v, "lan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan"] = t
		}
	}

	if v, ok := d.GetOk("energy_efficient_ethernet"); ok {
		t, err := expandWirelessControllerWtpProfileEnergyEfficientEthernet(d, v, "energy_efficient_ethernet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["energy-efficient-ethernet"] = t
		}
	}

	if v, ok := d.GetOk("led_state"); ok {
		t, err := expandWirelessControllerWtpProfileLedState(d, v, "led_state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-state"] = t
		}
	}

	if v, ok := d.GetOk("led_schedules"); ok || d.HasChange("led_schedules") {
		t, err := expandWirelessControllerWtpProfileLedSchedules(d, v, "led_schedules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-schedules"] = t
		}
	}

	if v, ok := d.GetOk("dtls_policy"); ok {
		t, err := expandWirelessControllerWtpProfileDtlsPolicy(d, v, "dtls_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-policy"] = t
		}
	}

	if v, ok := d.GetOk("dtls_in_kernel"); ok {
		t, err := expandWirelessControllerWtpProfileDtlsInKernel(d, v, "dtls_in_kernel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtls-in-kernel"] = t
		}
	}

	if v, ok := d.GetOkExists("max_clients"); ok {
		t, err := expandWirelessControllerWtpProfileMaxClients(d, v, "max_clients", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("handoff_rssi"); ok {
		t, err := expandWirelessControllerWtpProfileHandoffRssi(d, v, "handoff_rssi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-rssi"] = t
		}
	}

	if v, ok := d.GetOkExists("handoff_sta_thresh"); ok {
		t, err := expandWirelessControllerWtpProfileHandoffStaThresh(d, v, "handoff_sta_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-sta-thresh"] = t
		}
	}

	if v, ok := d.GetOk("handoff_roaming"); ok {
		t, err := expandWirelessControllerWtpProfileHandoffRoaming(d, v, "handoff_roaming", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["handoff-roaming"] = t
		}
	}

	if v, ok := d.GetOk("deny_mac_list"); ok || d.HasChange("deny_mac_list") {
		t, err := expandWirelessControllerWtpProfileDenyMacList(d, v, "deny_mac_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deny-mac-list"] = t
		}
	}

	if v, ok := d.GetOk("ap_country"); ok {
		t, err := expandWirelessControllerWtpProfileApCountry(d, v, "ap_country", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-country"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_preventing"); ok {
		t, err := expandWirelessControllerWtpProfileIpFragmentPreventing(d, v, "ip_fragment_preventing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-preventing"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_uplink"); ok {
		t, err := expandWirelessControllerWtpProfileTunMtuUplink(d, v, "tun_mtu_uplink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-uplink"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_downlink"); ok {
		t, err := expandWirelessControllerWtpProfileTunMtuDownlink(d, v, "tun_mtu_downlink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-downlink"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_path"); ok {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAclPath(d, v, "split_tunneling_acl_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-path"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAclLocalApSubnet(d, v, "split_tunneling_acl_local_ap_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-local-ap-subnet"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl"); ok || d.HasChange("split_tunneling_acl") {
		t, err := expandWirelessControllerWtpProfileSplitTunnelingAcl(d, v, "split_tunneling_acl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandWirelessControllerWtpProfileAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_change"); ok {
		t, err := expandWirelessControllerWtpProfileLoginPasswdChange(d, v, "login_passwd_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok {
		t, err := expandWirelessControllerWtpProfileLoginPasswd(d, v, "login_passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("lldp"); ok {
		t, err := expandWirelessControllerWtpProfileLldp(d, v, "lldp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp"] = t
		}
	}

	if v, ok := d.GetOk("poe_mode"); ok {
		t, err := expandWirelessControllerWtpProfilePoeMode(d, v, "poe_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poe-mode"] = t
		}
	}

	if v, ok := d.GetOk("frequency_handoff"); ok {
		t, err := expandWirelessControllerWtpProfileFrequencyHandoff(d, v, "frequency_handoff", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["frequency-handoff"] = t
		}
	}

	if v, ok := d.GetOk("ap_handoff"); ok {
		t, err := expandWirelessControllerWtpProfileApHandoff(d, v, "ap_handoff", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-handoff"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok {
		t, err := expandWirelessControllerWtpProfileRadio1(d, v, "radio_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok {
		t, err := expandWirelessControllerWtpProfileRadio2(d, v, "radio_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	if v, ok := d.GetOk("radio_3"); ok {
		t, err := expandWirelessControllerWtpProfileRadio3(d, v, "radio_3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-3"] = t
		}
	}

	if v, ok := d.GetOk("radio_4"); ok {
		t, err := expandWirelessControllerWtpProfileRadio4(d, v, "radio_4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-4"] = t
		}
	}

	if v, ok := d.GetOk("lbs"); ok {
		t, err := expandWirelessControllerWtpProfileLbs(d, v, "lbs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lbs"] = t
		}
	}

	if v, ok := d.GetOk("ext_info_enable"); ok {
		t, err := expandWirelessControllerWtpProfileExtInfoEnable(d, v, "ext_info_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-info-enable"] = t
		}
	}

	if v, ok := d.GetOk("indoor_outdoor_deployment"); ok {
		t, err := expandWirelessControllerWtpProfileIndoorOutdoorDeployment(d, v, "indoor_outdoor_deployment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["indoor-outdoor-deployment"] = t
		}
	}

	if v, ok := d.GetOk("esl_ses_dongle"); ok {
		t, err := expandWirelessControllerWtpProfileEslSesDongle(d, v, "esl_ses_dongle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["esl-ses-dongle"] = t
		}
	}

	if v, ok := d.GetOk("console_login"); ok {
		t, err := expandWirelessControllerWtpProfileConsoleLogin(d, v, "console_login", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["console-login"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_auth"); ok {
		t, err := expandWirelessControllerWtpProfileWanPortAuth(d, v, "wan_port_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-auth"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_auth_usrname"); ok {
		t, err := expandWirelessControllerWtpProfileWanPortAuthUsrname(d, v, "wan_port_auth_usrname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-auth-usrname"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_auth_password"); ok {
		t, err := expandWirelessControllerWtpProfileWanPortAuthPassword(d, v, "wan_port_auth_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-auth-password"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_auth_methods"); ok {
		t, err := expandWirelessControllerWtpProfileWanPortAuthMethods(d, v, "wan_port_auth_methods", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-auth-methods"] = t
		}
	}

	if v, ok := d.GetOk("unii_4_5ghz_band"); ok {
		t, err := expandWirelessControllerWtpProfileUnii45GhzBand(d, v, "unii_4_5ghz_band", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unii-4-5ghz-band"] = t
		}
	}

	return &obj, nil
}
