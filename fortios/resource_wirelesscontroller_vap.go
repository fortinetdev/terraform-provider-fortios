// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Virtual Access Points (VAPs).

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerVap() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerVapCreate,
		Read:   resourceWirelessControllerVapRead,
		Update: resourceWirelessControllerVapUpdate,
		Delete: resourceWirelessControllerVapDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Required:     true,
			},
			"pre_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_pre_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_fast_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mesh_backhaul": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"atf_weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_clients_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ssid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"broadcast_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_obsolete_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pmf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pmf_assoc_comeback_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20),
				Optional:     true,
				Computed:     true,
			},
			"pmf_sa_query_retry_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 5),
				Optional:     true,
				Computed:     true,
			},
			"beacon_protection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"okc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mbo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gas_comeback_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 10000),
				Optional:     true,
				Computed:     true,
			},
			"gas_fragmentation_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(512, 4096),
				Optional:     true,
				Computed:     true,
			},
			"mbo_cell_data_conn_pref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n80211k": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n80211v": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voice_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"neighbor_report_dual_band": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_bss_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ft_mobility_domain": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ft_r0_key_lifetime": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ft_over_ds": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"owe_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"owe_transition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owe_transition_ssid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
			},
			"additional_akms": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eapol_key_retries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tkip_counter_measure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_web": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"external_web_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_logout": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_calling_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_called_station_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"called_station_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_auth_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"radius_mac_auth_block_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 864000),
				Optional:     true,
			},
			"radius_mac_mpsk_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_mac_mpsk_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 864000),
				Optional:     true,
				Computed:     true,
			},
			"radius_mac_auth_usergroups": &schema.Schema{
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
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keyindex": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4),
				Optional:     true,
				Computed:     true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"passphrase": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"sae_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"sae_h2e_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_hnp_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_pk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sae_private_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 359),
				Optional:     true,
			},
			"akm24_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"nas_filter_rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name_stripping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
			},
			"local_standalone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_lease_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"local_standalone_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_standalone_dns_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_lan_partition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_bridging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usergroup": &schema.Schema{
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
			"portal_message_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"portal_message_overrides": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"auth_reject_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"auth_login_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"auth_login_failed_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"selected_usergroups": &schema.Schema{
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
			"security_exempt_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"security_redirect_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"auth_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"auth_portal_addr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"intra_vap_privacy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ldpc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"high_efficiency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_wake_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_macauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port_macauth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 65535),
				Optional:     true,
				Computed:     true,
			},
			"port_macauth_reauth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 65535),
				Optional:     true,
				Computed:     true,
			},
			"bss_color_partial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"mpsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mpsk_concurrent_clients": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"mpsk_key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"passphrase": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"concurrent_clients": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"mpsk_schedules": &schema.Schema{
							Type:     schema.TypeSet,
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
					},
				},
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4094),
				Optional:     true,
			},
			"vlan_auto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_fw_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal_radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"captive_portal_radius_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"captive_portal_macauth_radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"captive_portal_macauth_radius_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"captive_portal_ac_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"captive_portal_auth_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 864000),
				Optional:     true,
			},
			"captive_portal_session_timeout_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 864000),
				Optional:     true,
			},
			"alias": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 25),
				Optional:     true,
				Computed:     true,
			},
			"multicast_rate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_enhance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_address_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_suppression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_rules": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"me_disable_thresh": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 256),
				Optional:     true,
				Computed:     true,
			},
			"mu_mimo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_resp_suppression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"probe_resp_threshold": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"radio_sensitivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radio_5g_threshold": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"radio_2g_threshold": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"vlan_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"vlan_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),
							Optional:     true,
						},
					},
				},
			},
			"vlan_pooling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_pool": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),
							Optional:     true,
							Computed:     true,
						},
						"wtp_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
					},
				},
			},
			"dhcp_option43_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_circuit_id_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_remote_id_insertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptk_rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ptk_rekey_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(600, 864000),
				Optional:     true,
				Computed:     true,
			},
			"gtk_rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtk_rekey_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(600, 864000),
				Optional:     true,
				Computed:     true,
			},
			"eap_reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_reauth_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1800, 864000),
				Optional:     true,
				Computed:     true,
			},
			"roaming_acct_interim_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"qos_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"hotspot20_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"access_control_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"primary_wag_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"secondary_wag_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"tunnel_echo_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"tunnel_fallback_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"rates_11a": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11bg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11n_ss12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11n_ss34": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ac_mcs_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"rates_11ax_mcs_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"rates_11be_mcs_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"rates_11be_mcs_map_160": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"rates_11be_mcs_map_320": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"rates_11ac_ss12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ac_ss34": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ax_ss12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rates_11ax_ss34": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"utm_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"antivirus_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"address_group_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter_policy_other": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_filter_list": &schema.Schema{
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
						"mac_filter_policy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sticky_client_remove": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sticky_client_threshold_5g": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"sticky_client_threshold_2g": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"sticky_client_threshold_6g": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"bstm_rssi_disassoc_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2000),
				Optional:     true,
				Computed:     true,
			},
			"bstm_load_balancing_disassoc_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"bstm_disassociation_imminent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"beacon_advertising": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"osen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_detection_engine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_dscp_marking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_report_intv": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 864000),
				Optional:     true,
				Computed:     true,
			},
			"l3_roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l3_roaming_mode": &schema.Schema{
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

func resourceWirelessControllerVapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerVap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVap resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerVap(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerVap")
	}

	return resourceWirelessControllerVapRead(d, m)
}

func resourceWirelessControllerVapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerVap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVap resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerVap(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerVap")
	}

	return resourceWirelessControllerVapRead(d, m)
}

func resourceWirelessControllerVapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerVap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerVap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerVapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerVap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVap resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPreAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalPreAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapFastRoaming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalFastRoaming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMeshBackhaul(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAtfWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapMaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapMaxClientsAp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapBroadcastSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityObsoleteOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPmf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPmfAssocComebackTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapPmfSaQueryRetryTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapBeaconProtection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapOkc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMbo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapGasComebackDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapGasFragmentationLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapMboCellDataConnPref(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVap80211K(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVap80211V(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapVoiceEnterprise(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapNeighborReportDualBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapFastBssTransition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapFtMobilityDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapFtR0KeyLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapFtOverDs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSaeGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapOweGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapOweTransition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapOweTransitionSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAdditionalAkms(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapEapolKeyRetries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapTkipCounterMeasure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalWeb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalWebFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalLogout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacCallingStationDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacCalledStationDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacCase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCalledStationIdType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacAuthBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuthServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuthBlockInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapRadiusMacMpskAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacMpskTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapRadiusMacAuthUsergroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapRadiusMacAuthUsergroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapRadiusMacAuthUsergroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapEncrypt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapKeyindex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapSaeH2EOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSaeHnpOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSaePk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSaePrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAkm24Only(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapNasFilterRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDomainNameStripping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapLocalStandalone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandaloneNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenWirelessControllerVapDhcpLeaseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapLocalStandaloneDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandaloneDnsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalLanPartition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalBridging(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalLan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapUsergroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapUsergroupName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapUsergroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverrides(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := i["auth-disclaimer-page"]; ok {
		result["auth_disclaimer_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(i["auth-disclaimer-page"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := i["auth-reject-page"]; ok {
		result["auth_reject_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage(i["auth-reject-page"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := i["auth-login-page"]; ok {
		result["auth_login_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage(i["auth-login-page"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := i["auth-login-failed-page"]; ok {
		result["auth_login_failed_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(i["auth-login-failed-page"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapSelectedUsergroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapSelectedUsergroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityExemptList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAuthCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAuthPortalAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapIntraVapPrivacy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("schedule"), "name")
}

func flattenWirelessControllerVapLdpc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapHighEfficiency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapTargetWakeTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortMacauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPortMacauthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapPortMacauthReauthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapBssColorPartial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMpsk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapMpskKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if cur_v, ok := i["key-name"]; ok {
			tmp["key_name"] = flattenWirelessControllerVapMpskKeyKeyName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := i["passphrase"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["passphrase"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if cur_v, ok := i["concurrent-clients"]; ok {
			tmp["concurrent_clients"] = flattenWirelessControllerVapMpskKeyConcurrentClients(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if cur_v, ok := i["comment"]; ok {
			tmp["comment"] = flattenWirelessControllerVapMpskKeyComment(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if cur_v, ok := i["mpsk-schedules"]; ok {
			tmp["mpsk_schedules"] = flattenWirelessControllerVapMpskKeyMpskSchedules(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "key_name", d)
	return result
}

func flattenWirelessControllerVapMpskKeyKeyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyMpskSchedules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapMpskKeyMpskSchedulesName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapMpskKeyMpskSchedulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSplitTunneling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapNac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapNacProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapVlanAuto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalFwAccounting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalMacauthRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalAcName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalAuthTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapCaptivePortalSessionTimeoutInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapAlias(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMulticastRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMulticastEnhance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapIgmpSnooping(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpAddressEnforcement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapBroadcastSuppression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapIpv6Rules(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMeDisableThresh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapMuMimo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapProbeRespSuppression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapProbeRespThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadioSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadio5GThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRadio2GThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapVlanNameName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if cur_v, ok := i["vlan-id"]; ok {
			tmp["vlan_id"] = flattenWirelessControllerVapVlanNameVlanId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapVlanNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanNameVlanId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapVlanPooling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanPool(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerVapVlanPoolId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_group"
		if cur_v, ok := i["wtp-group"]; ok {
			tmp["wtp_group"] = flattenWirelessControllerVapVlanPoolWtpGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerVapVlanPoolId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapVlanPoolWtpGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption43Insertion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82Insertion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82CircuitIdInsertion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82RemoteIdInsertion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPtkRekey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapGtkRekey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapGtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapEapReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapEapReauthIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapRoamingAcctInterimUpdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapQosProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapHotspot20Profile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAccessControlList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapPrimaryWagProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapSecondaryWagProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapTunnelEchoInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapTunnelFallbackInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapRates11A(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11Bg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11NSs12(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11NSs34(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AcMcsMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AxMcsMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11BeMcsMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11BeMcsMap160(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11BeMcsMap320(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AcSs12(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AcSs34(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AxSs12(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AxSs34(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapUtmProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapUtmStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapUtmLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapIpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAntivirusProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapWebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAddressGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapAddressGroupPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterPolicyOther(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerVapMacFilterListId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if cur_v, ok := i["mac"]; ok {
			tmp["mac"] = flattenWirelessControllerVapMacFilterListMac(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if cur_v, ok := i["mac-filter-policy"]; ok {
			tmp["mac_filter_policy"] = flattenWirelessControllerVapMacFilterListMacFilterPolicy(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerVapMacFilterListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapMacFilterListMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterListMacFilterPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientRemove(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientThreshold5G(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientThreshold2G(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapStickyClientThreshold6G(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapBstmRssiDisassocTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapBstmLoadBalancingDisassocTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapBstmDisassociationImminent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapBeaconAdvertising(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapOsen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapApplicationDetectionEngine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapApplicationDscpMarking(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapApplicationReportIntv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerVapL3Roaming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerVapL3RoamingMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerVap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerVapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pre_auth", flattenWirelessControllerVapPreAuth(o["pre-auth"], d, "pre_auth", sv)); err != nil {
		if !fortiAPIPatch(o["pre-auth"]) {
			return fmt.Errorf("Error reading pre_auth: %v", err)
		}
	}

	if err = d.Set("external_pre_auth", flattenWirelessControllerVapExternalPreAuth(o["external-pre-auth"], d, "external_pre_auth", sv)); err != nil {
		if !fortiAPIPatch(o["external-pre-auth"]) {
			return fmt.Errorf("Error reading external_pre_auth: %v", err)
		}
	}

	if err = d.Set("fast_roaming", flattenWirelessControllerVapFastRoaming(o["fast-roaming"], d, "fast_roaming", sv)); err != nil {
		if !fortiAPIPatch(o["fast-roaming"]) {
			return fmt.Errorf("Error reading fast_roaming: %v", err)
		}
	}

	if err = d.Set("external_fast_roaming", flattenWirelessControllerVapExternalFastRoaming(o["external-fast-roaming"], d, "external_fast_roaming", sv)); err != nil {
		if !fortiAPIPatch(o["external-fast-roaming"]) {
			return fmt.Errorf("Error reading external_fast_roaming: %v", err)
		}
	}

	if err = d.Set("mesh_backhaul", flattenWirelessControllerVapMeshBackhaul(o["mesh-backhaul"], d, "mesh_backhaul", sv)); err != nil {
		if !fortiAPIPatch(o["mesh-backhaul"]) {
			return fmt.Errorf("Error reading mesh_backhaul: %v", err)
		}
	}

	if err = d.Set("atf_weight", flattenWirelessControllerVapAtfWeight(o["atf-weight"], d, "atf_weight", sv)); err != nil {
		if !fortiAPIPatch(o["atf-weight"]) {
			return fmt.Errorf("Error reading atf_weight: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerVapMaxClients(o["max-clients"], d, "max_clients", sv)); err != nil {
		if !fortiAPIPatch(o["max-clients"]) {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_clients_ap", flattenWirelessControllerVapMaxClientsAp(o["max-clients-ap"], d, "max_clients_ap", sv)); err != nil {
		if !fortiAPIPatch(o["max-clients-ap"]) {
			return fmt.Errorf("Error reading max_clients_ap: %v", err)
		}
	}

	if err = d.Set("ssid", flattenWirelessControllerVapSsid(o["ssid"], d, "ssid", sv)); err != nil {
		if !fortiAPIPatch(o["ssid"]) {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenWirelessControllerVapBroadcastSsid(o["broadcast-ssid"], d, "broadcast_ssid", sv)); err != nil {
		if !fortiAPIPatch(o["broadcast-ssid"]) {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("security_obsolete_option", flattenWirelessControllerVapSecurityObsoleteOption(o["security-obsolete-option"], d, "security_obsolete_option", sv)); err != nil {
		if !fortiAPIPatch(o["security-obsolete-option"]) {
			return fmt.Errorf("Error reading security_obsolete_option: %v", err)
		}
	}

	if err = d.Set("security", flattenWirelessControllerVapSecurity(o["security"], d, "security", sv)); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("pmf", flattenWirelessControllerVapPmf(o["pmf"], d, "pmf", sv)); err != nil {
		if !fortiAPIPatch(o["pmf"]) {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("pmf_assoc_comeback_timeout", flattenWirelessControllerVapPmfAssocComebackTimeout(o["pmf-assoc-comeback-timeout"], d, "pmf_assoc_comeback_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["pmf-assoc-comeback-timeout"]) {
			return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
		}
	}

	if err = d.Set("pmf_sa_query_retry_timeout", flattenWirelessControllerVapPmfSaQueryRetryTimeout(o["pmf-sa-query-retry-timeout"], d, "pmf_sa_query_retry_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["pmf-sa-query-retry-timeout"]) {
			return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
		}
	}

	if err = d.Set("beacon_protection", flattenWirelessControllerVapBeaconProtection(o["beacon-protection"], d, "beacon_protection", sv)); err != nil {
		if !fortiAPIPatch(o["beacon-protection"]) {
			return fmt.Errorf("Error reading beacon_protection: %v", err)
		}
	}

	if err = d.Set("okc", flattenWirelessControllerVapOkc(o["okc"], d, "okc", sv)); err != nil {
		if !fortiAPIPatch(o["okc"]) {
			return fmt.Errorf("Error reading okc: %v", err)
		}
	}

	if err = d.Set("mbo", flattenWirelessControllerVapMbo(o["mbo"], d, "mbo", sv)); err != nil {
		if !fortiAPIPatch(o["mbo"]) {
			return fmt.Errorf("Error reading mbo: %v", err)
		}
	}

	if err = d.Set("gas_comeback_delay", flattenWirelessControllerVapGasComebackDelay(o["gas-comeback-delay"], d, "gas_comeback_delay", sv)); err != nil {
		if !fortiAPIPatch(o["gas-comeback-delay"]) {
			return fmt.Errorf("Error reading gas_comeback_delay: %v", err)
		}
	}

	if err = d.Set("gas_fragmentation_limit", flattenWirelessControllerVapGasFragmentationLimit(o["gas-fragmentation-limit"], d, "gas_fragmentation_limit", sv)); err != nil {
		if !fortiAPIPatch(o["gas-fragmentation-limit"]) {
			return fmt.Errorf("Error reading gas_fragmentation_limit: %v", err)
		}
	}

	if err = d.Set("mbo_cell_data_conn_pref", flattenWirelessControllerVapMboCellDataConnPref(o["mbo-cell-data-conn-pref"], d, "mbo_cell_data_conn_pref", sv)); err != nil {
		if !fortiAPIPatch(o["mbo-cell-data-conn-pref"]) {
			return fmt.Errorf("Error reading mbo_cell_data_conn_pref: %v", err)
		}
	}

	if err = d.Set("n80211k", flattenWirelessControllerVap80211K(o["80211k"], d, "n80211k", sv)); err != nil {
		if !fortiAPIPatch(o["80211k"]) {
			return fmt.Errorf("Error reading n80211k: %v", err)
		}
	}

	if err = d.Set("n80211v", flattenWirelessControllerVap80211V(o["80211v"], d, "n80211v", sv)); err != nil {
		if !fortiAPIPatch(o["80211v"]) {
			return fmt.Errorf("Error reading n80211v: %v", err)
		}
	}

	if err = d.Set("voice_enterprise", flattenWirelessControllerVapVoiceEnterprise(o["voice-enterprise"], d, "voice_enterprise", sv)); err != nil {
		if !fortiAPIPatch(o["voice-enterprise"]) {
			return fmt.Errorf("Error reading voice_enterprise: %v", err)
		}
	}

	if err = d.Set("neighbor_report_dual_band", flattenWirelessControllerVapNeighborReportDualBand(o["neighbor-report-dual-band"], d, "neighbor_report_dual_band", sv)); err != nil {
		if !fortiAPIPatch(o["neighbor-report-dual-band"]) {
			return fmt.Errorf("Error reading neighbor_report_dual_band: %v", err)
		}
	}

	if err = d.Set("fast_bss_transition", flattenWirelessControllerVapFastBssTransition(o["fast-bss-transition"], d, "fast_bss_transition", sv)); err != nil {
		if !fortiAPIPatch(o["fast-bss-transition"]) {
			return fmt.Errorf("Error reading fast_bss_transition: %v", err)
		}
	}

	if err = d.Set("ft_mobility_domain", flattenWirelessControllerVapFtMobilityDomain(o["ft-mobility-domain"], d, "ft_mobility_domain", sv)); err != nil {
		if !fortiAPIPatch(o["ft-mobility-domain"]) {
			return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
		}
	}

	if err = d.Set("ft_r0_key_lifetime", flattenWirelessControllerVapFtR0KeyLifetime(o["ft-r0-key-lifetime"], d, "ft_r0_key_lifetime", sv)); err != nil {
		if !fortiAPIPatch(o["ft-r0-key-lifetime"]) {
			return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
		}
	}

	if err = d.Set("ft_over_ds", flattenWirelessControllerVapFtOverDs(o["ft-over-ds"], d, "ft_over_ds", sv)); err != nil {
		if !fortiAPIPatch(o["ft-over-ds"]) {
			return fmt.Errorf("Error reading ft_over_ds: %v", err)
		}
	}

	if err = d.Set("sae_groups", flattenWirelessControllerVapSaeGroups(o["sae-groups"], d, "sae_groups", sv)); err != nil {
		if !fortiAPIPatch(o["sae-groups"]) {
			return fmt.Errorf("Error reading sae_groups: %v", err)
		}
	}

	if err = d.Set("owe_groups", flattenWirelessControllerVapOweGroups(o["owe-groups"], d, "owe_groups", sv)); err != nil {
		if !fortiAPIPatch(o["owe-groups"]) {
			return fmt.Errorf("Error reading owe_groups: %v", err)
		}
	}

	if err = d.Set("owe_transition", flattenWirelessControllerVapOweTransition(o["owe-transition"], d, "owe_transition", sv)); err != nil {
		if !fortiAPIPatch(o["owe-transition"]) {
			return fmt.Errorf("Error reading owe_transition: %v", err)
		}
	}

	if err = d.Set("owe_transition_ssid", flattenWirelessControllerVapOweTransitionSsid(o["owe-transition-ssid"], d, "owe_transition_ssid", sv)); err != nil {
		if !fortiAPIPatch(o["owe-transition-ssid"]) {
			return fmt.Errorf("Error reading owe_transition_ssid: %v", err)
		}
	}

	if err = d.Set("additional_akms", flattenWirelessControllerVapAdditionalAkms(o["additional-akms"], d, "additional_akms", sv)); err != nil {
		if !fortiAPIPatch(o["additional-akms"]) {
			return fmt.Errorf("Error reading additional_akms: %v", err)
		}
	}

	if err = d.Set("eapol_key_retries", flattenWirelessControllerVapEapolKeyRetries(o["eapol-key-retries"], d, "eapol_key_retries", sv)); err != nil {
		if !fortiAPIPatch(o["eapol-key-retries"]) {
			return fmt.Errorf("Error reading eapol_key_retries: %v", err)
		}
	}

	if err = d.Set("tkip_counter_measure", flattenWirelessControllerVapTkipCounterMeasure(o["tkip-counter-measure"], d, "tkip_counter_measure", sv)); err != nil {
		if !fortiAPIPatch(o["tkip-counter-measure"]) {
			return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
		}
	}

	if err = d.Set("external_web", flattenWirelessControllerVapExternalWeb(o["external-web"], d, "external_web", sv)); err != nil {
		if !fortiAPIPatch(o["external-web"]) {
			return fmt.Errorf("Error reading external_web: %v", err)
		}
	}

	if err = d.Set("external_web_format", flattenWirelessControllerVapExternalWebFormat(o["external-web-format"], d, "external_web_format", sv)); err != nil {
		if !fortiAPIPatch(o["external-web-format"]) {
			return fmt.Errorf("Error reading external_web_format: %v", err)
		}
	}

	if err = d.Set("external_logout", flattenWirelessControllerVapExternalLogout(o["external-logout"], d, "external_logout", sv)); err != nil {
		if !fortiAPIPatch(o["external-logout"]) {
			return fmt.Errorf("Error reading external_logout: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenWirelessControllerVapMacUsernameDelimiter(o["mac-username-delimiter"], d, "mac_username_delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-username-delimiter"]) {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenWirelessControllerVapMacPasswordDelimiter(o["mac-password-delimiter"], d, "mac_password_delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-password-delimiter"]) {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_calling_station_delimiter", flattenWirelessControllerVapMacCallingStationDelimiter(o["mac-calling-station-delimiter"], d, "mac_calling_station_delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-calling-station-delimiter"]) {
			return fmt.Errorf("Error reading mac_calling_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_called_station_delimiter", flattenWirelessControllerVapMacCalledStationDelimiter(o["mac-called-station-delimiter"], d, "mac_called_station_delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-called-station-delimiter"]) {
			return fmt.Errorf("Error reading mac_called_station_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenWirelessControllerVapMacCase(o["mac-case"], d, "mac_case", sv)); err != nil {
		if !fortiAPIPatch(o["mac-case"]) {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("called_station_id_type", flattenWirelessControllerVapCalledStationIdType(o["called-station-id-type"], d, "called_station_id_type", sv)); err != nil {
		if !fortiAPIPatch(o["called-station-id-type"]) {
			return fmt.Errorf("Error reading called_station_id_type: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenWirelessControllerVapMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass", sv)); err != nil {
		if !fortiAPIPatch(o["mac-auth-bypass"]) {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth", flattenWirelessControllerVapRadiusMacAuth(o["radius-mac-auth"], d, "radius_mac_auth", sv)); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth"]) {
			return fmt.Errorf("Error reading radius_mac_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_server", flattenWirelessControllerVapRadiusMacAuthServer(o["radius-mac-auth-server"], d, "radius_mac_auth_server", sv)); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth-server"]) {
			return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_block_interval", flattenWirelessControllerVapRadiusMacAuthBlockInterval(o["radius-mac-auth-block-interval"], d, "radius_mac_auth_block_interval", sv)); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth-block-interval"]) {
			return fmt.Errorf("Error reading radius_mac_auth_block_interval: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_auth", flattenWirelessControllerVapRadiusMacMpskAuth(o["radius-mac-mpsk-auth"], d, "radius_mac_mpsk_auth", sv)); err != nil {
		if !fortiAPIPatch(o["radius-mac-mpsk-auth"]) {
			return fmt.Errorf("Error reading radius_mac_mpsk_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_mpsk_timeout", flattenWirelessControllerVapRadiusMacMpskTimeout(o["radius-mac-mpsk-timeout"], d, "radius_mac_mpsk_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["radius-mac-mpsk-timeout"]) {
			return fmt.Errorf("Error reading radius_mac_mpsk_timeout: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("radius_mac_auth_usergroups", flattenWirelessControllerVapRadiusMacAuthUsergroups(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups", sv)); err != nil {
			if !fortiAPIPatch(o["radius-mac-auth-usergroups"]) {
				return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radius_mac_auth_usergroups"); ok {
			if err = d.Set("radius_mac_auth_usergroups", flattenWirelessControllerVapRadiusMacAuthUsergroups(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups", sv)); err != nil {
				if !fortiAPIPatch(o["radius-mac-auth-usergroups"]) {
					return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth", flattenWirelessControllerVapAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenWirelessControllerVapEncrypt(o["encrypt"], d, "encrypt", sv)); err != nil {
		if !fortiAPIPatch(o["encrypt"]) {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("keyindex", flattenWirelessControllerVapKeyindex(o["keyindex"], d, "keyindex", sv)); err != nil {
		if !fortiAPIPatch(o["keyindex"]) {
			return fmt.Errorf("Error reading keyindex: %v", err)
		}
	}

	if err = d.Set("sae_h2e_only", flattenWirelessControllerVapSaeH2EOnly(o["sae-h2e-only"], d, "sae_h2e_only", sv)); err != nil {
		if !fortiAPIPatch(o["sae-h2e-only"]) {
			return fmt.Errorf("Error reading sae_h2e_only: %v", err)
		}
	}

	if err = d.Set("sae_hnp_only", flattenWirelessControllerVapSaeHnpOnly(o["sae-hnp-only"], d, "sae_hnp_only", sv)); err != nil {
		if !fortiAPIPatch(o["sae-hnp-only"]) {
			return fmt.Errorf("Error reading sae_hnp_only: %v", err)
		}
	}

	if err = d.Set("sae_pk", flattenWirelessControllerVapSaePk(o["sae-pk"], d, "sae_pk", sv)); err != nil {
		if !fortiAPIPatch(o["sae-pk"]) {
			return fmt.Errorf("Error reading sae_pk: %v", err)
		}
	}

	if err = d.Set("sae_private_key", flattenWirelessControllerVapSaePrivateKey(o["sae-private-key"], d, "sae_private_key", sv)); err != nil {
		if !fortiAPIPatch(o["sae-private-key"]) {
			return fmt.Errorf("Error reading sae_private_key: %v", err)
		}
	}

	if err = d.Set("akm24_only", flattenWirelessControllerVapAkm24Only(o["akm24-only"], d, "akm24_only", sv)); err != nil {
		if !fortiAPIPatch(o["akm24-only"]) {
			return fmt.Errorf("Error reading akm24_only: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenWirelessControllerVapRadiusServer(o["radius-server"], d, "radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["radius-server"]) {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("nas_filter_rule", flattenWirelessControllerVapNasFilterRule(o["nas-filter-rule"], d, "nas_filter_rule", sv)); err != nil {
		if !fortiAPIPatch(o["nas-filter-rule"]) {
			return fmt.Errorf("Error reading nas_filter_rule: %v", err)
		}
	}

	if err = d.Set("domain_name_stripping", flattenWirelessControllerVapDomainNameStripping(o["domain-name-stripping"], d, "domain_name_stripping", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name-stripping"]) {
			return fmt.Errorf("Error reading domain_name_stripping: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenWirelessControllerVapAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval", sv)); err != nil {
		if !fortiAPIPatch(o["acct-interim-interval"]) {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("local_standalone", flattenWirelessControllerVapLocalStandalone(o["local-standalone"], d, "local_standalone", sv)); err != nil {
		if !fortiAPIPatch(o["local-standalone"]) {
			return fmt.Errorf("Error reading local_standalone: %v", err)
		}
	}

	if err = d.Set("local_standalone_nat", flattenWirelessControllerVapLocalStandaloneNat(o["local-standalone-nat"], d, "local_standalone_nat", sv)); err != nil {
		if !fortiAPIPatch(o["local-standalone-nat"]) {
			return fmt.Errorf("Error reading local_standalone_nat: %v", err)
		}
	}

	if err = d.Set("ip", flattenWirelessControllerVapIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_time", flattenWirelessControllerVapDhcpLeaseTime(o["dhcp-lease-time"], d, "dhcp_lease_time", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-lease-time"]) {
			return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns", flattenWirelessControllerVapLocalStandaloneDns(o["local-standalone-dns"], d, "local_standalone_dns", sv)); err != nil {
		if !fortiAPIPatch(o["local-standalone-dns"]) {
			return fmt.Errorf("Error reading local_standalone_dns: %v", err)
		}
	}

	if err = d.Set("local_standalone_dns_ip", flattenWirelessControllerVapLocalStandaloneDnsIp(o["local-standalone-dns-ip"], d, "local_standalone_dns_ip", sv)); err != nil {
		if !fortiAPIPatch(o["local-standalone-dns-ip"]) {
			return fmt.Errorf("Error reading local_standalone_dns_ip: %v", err)
		}
	}

	if err = d.Set("local_lan_partition", flattenWirelessControllerVapLocalLanPartition(o["local-lan-partition"], d, "local_lan_partition", sv)); err != nil {
		if !fortiAPIPatch(o["local-lan-partition"]) {
			return fmt.Errorf("Error reading local_lan_partition: %v", err)
		}
	}

	if err = d.Set("local_bridging", flattenWirelessControllerVapLocalBridging(o["local-bridging"], d, "local_bridging", sv)); err != nil {
		if !fortiAPIPatch(o["local-bridging"]) {
			return fmt.Errorf("Error reading local_bridging: %v", err)
		}
	}

	if err = d.Set("local_lan", flattenWirelessControllerVapLocalLan(o["local-lan"], d, "local_lan", sv)); err != nil {
		if !fortiAPIPatch(o["local-lan"]) {
			return fmt.Errorf("Error reading local_lan: %v", err)
		}
	}

	if err = d.Set("local_authentication", flattenWirelessControllerVapLocalAuthentication(o["local-authentication"], d, "local_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["local-authentication"]) {
			return fmt.Errorf("Error reading local_authentication: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("usergroup", flattenWirelessControllerVapUsergroup(o["usergroup"], d, "usergroup", sv)); err != nil {
			if !fortiAPIPatch(o["usergroup"]) {
				return fmt.Errorf("Error reading usergroup: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("usergroup"); ok {
			if err = d.Set("usergroup", flattenWirelessControllerVapUsergroup(o["usergroup"], d, "usergroup", sv)); err != nil {
				if !fortiAPIPatch(o["usergroup"]) {
					return fmt.Errorf("Error reading usergroup: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_message_override_group", flattenWirelessControllerVapPortalMessageOverrideGroup(o["portal-message-override-group"], d, "portal_message_override_group", sv)); err != nil {
		if !fortiAPIPatch(o["portal-message-override-group"]) {
			return fmt.Errorf("Error reading portal_message_override_group: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides", sv)); err != nil {
			if !fortiAPIPatch(o["portal-message-overrides"]) {
				return fmt.Errorf("Error reading portal_message_overrides: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("portal_message_overrides"); ok {
			if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides", sv)); err != nil {
				if !fortiAPIPatch(o["portal-message-overrides"]) {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_type", flattenWirelessControllerVapPortalType(o["portal-type"], d, "portal_type", sv)); err != nil {
		if !fortiAPIPatch(o["portal-type"]) {
			return fmt.Errorf("Error reading portal_type: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("selected_usergroups", flattenWirelessControllerVapSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups", sv)); err != nil {
			if !fortiAPIPatch(o["selected-usergroups"]) {
				return fmt.Errorf("Error reading selected_usergroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("selected_usergroups"); ok {
			if err = d.Set("selected_usergroups", flattenWirelessControllerVapSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups", sv)); err != nil {
				if !fortiAPIPatch(o["selected-usergroups"]) {
					return fmt.Errorf("Error reading selected_usergroups: %v", err)
				}
			}
		}
	}

	if err = d.Set("security_exempt_list", flattenWirelessControllerVapSecurityExemptList(o["security-exempt-list"], d, "security_exempt_list", sv)); err != nil {
		if !fortiAPIPatch(o["security-exempt-list"]) {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenWirelessControllerVapSecurityRedirectUrl(o["security-redirect-url"], d, "security_redirect_url", sv)); err != nil {
		if !fortiAPIPatch(o["security-redirect-url"]) {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("auth_cert", flattenWirelessControllerVapAuthCert(o["auth-cert"], d, "auth_cert", sv)); err != nil {
		if !fortiAPIPatch(o["auth-cert"]) {
			return fmt.Errorf("Error reading auth_cert: %v", err)
		}
	}

	if err = d.Set("auth_portal_addr", flattenWirelessControllerVapAuthPortalAddr(o["auth-portal-addr"], d, "auth_portal_addr", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal-addr"]) {
			return fmt.Errorf("Error reading auth_portal_addr: %v", err)
		}
	}

	if err = d.Set("intra_vap_privacy", flattenWirelessControllerVapIntraVapPrivacy(o["intra-vap-privacy"], d, "intra_vap_privacy", sv)); err != nil {
		if !fortiAPIPatch(o["intra-vap-privacy"]) {
			return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
		}
	}

	if err = d.Set("schedule", flattenWirelessControllerVapSchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("ldpc", flattenWirelessControllerVapLdpc(o["ldpc"], d, "ldpc", sv)); err != nil {
		if !fortiAPIPatch(o["ldpc"]) {
			return fmt.Errorf("Error reading ldpc: %v", err)
		}
	}

	if err = d.Set("high_efficiency", flattenWirelessControllerVapHighEfficiency(o["high-efficiency"], d, "high_efficiency", sv)); err != nil {
		if !fortiAPIPatch(o["high-efficiency"]) {
			return fmt.Errorf("Error reading high_efficiency: %v", err)
		}
	}

	if err = d.Set("target_wake_time", flattenWirelessControllerVapTargetWakeTime(o["target-wake-time"], d, "target_wake_time", sv)); err != nil {
		if !fortiAPIPatch(o["target-wake-time"]) {
			return fmt.Errorf("Error reading target_wake_time: %v", err)
		}
	}

	if err = d.Set("port_macauth", flattenWirelessControllerVapPortMacauth(o["port-macauth"], d, "port_macauth", sv)); err != nil {
		if !fortiAPIPatch(o["port-macauth"]) {
			return fmt.Errorf("Error reading port_macauth: %v", err)
		}
	}

	if err = d.Set("port_macauth_timeout", flattenWirelessControllerVapPortMacauthTimeout(o["port-macauth-timeout"], d, "port_macauth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["port-macauth-timeout"]) {
			return fmt.Errorf("Error reading port_macauth_timeout: %v", err)
		}
	}

	if err = d.Set("port_macauth_reauth_timeout", flattenWirelessControllerVapPortMacauthReauthTimeout(o["port-macauth-reauth-timeout"], d, "port_macauth_reauth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["port-macauth-reauth-timeout"]) {
			return fmt.Errorf("Error reading port_macauth_reauth_timeout: %v", err)
		}
	}

	if err = d.Set("bss_color_partial", flattenWirelessControllerVapBssColorPartial(o["bss-color-partial"], d, "bss_color_partial", sv)); err != nil {
		if !fortiAPIPatch(o["bss-color-partial"]) {
			return fmt.Errorf("Error reading bss_color_partial: %v", err)
		}
	}

	if err = d.Set("mpsk_profile", flattenWirelessControllerVapMpskProfile(o["mpsk-profile"], d, "mpsk_profile", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk-profile"]) {
			return fmt.Errorf("Error reading mpsk_profile: %v", err)
		}
	}

	if err = d.Set("mpsk", flattenWirelessControllerVapMpsk(o["mpsk"], d, "mpsk", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk"]) {
			return fmt.Errorf("Error reading mpsk: %v", err)
		}
	}

	if err = d.Set("mpsk_concurrent_clients", flattenWirelessControllerVapMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients", sv)); err != nil {
		if !fortiAPIPatch(o["mpsk-concurrent-clients"]) {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key", sv)); err != nil {
			if !fortiAPIPatch(o["mpsk-key"]) {
				return fmt.Errorf("Error reading mpsk_key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_key"); ok {
			if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key", sv)); err != nil {
				if !fortiAPIPatch(o["mpsk-key"]) {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			}
		}
	}

	if err = d.Set("split_tunneling", flattenWirelessControllerVapSplitTunneling(o["split-tunneling"], d, "split_tunneling", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling"]) {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("nac", flattenWirelessControllerVapNac(o["nac"], d, "nac", sv)); err != nil {
		if !fortiAPIPatch(o["nac"]) {
			return fmt.Errorf("Error reading nac: %v", err)
		}
	}

	if err = d.Set("nac_profile", flattenWirelessControllerVapNacProfile(o["nac-profile"], d, "nac_profile", sv)); err != nil {
		if !fortiAPIPatch(o["nac-profile"]) {
			return fmt.Errorf("Error reading nac_profile: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenWirelessControllerVapVlanid(o["vlanid"], d, "vlanid", sv)); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("vlan_auto", flattenWirelessControllerVapVlanAuto(o["vlan-auto"], d, "vlan_auto", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-auto"]) {
			return fmt.Errorf("Error reading vlan_auto: %v", err)
		}
	}

	if err = d.Set("dynamic_vlan", flattenWirelessControllerVapDynamicVlan(o["dynamic-vlan"], d, "dynamic_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-vlan"]) {
			return fmt.Errorf("Error reading dynamic_vlan: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenWirelessControllerVapCaptivePortal(o["captive-portal"], d, "captive_portal", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal"]) {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("captive_portal_fw_accounting", flattenWirelessControllerVapCaptivePortalFwAccounting(o["captive-portal-fw-accounting"], d, "captive_portal_fw_accounting", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-fw-accounting"]) {
			return fmt.Errorf("Error reading captive_portal_fw_accounting: %v", err)
		}
	}

	if err = d.Set("captive_portal_radius_server", flattenWirelessControllerVapCaptivePortalRadiusServer(o["captive-portal-radius-server"], d, "captive_portal_radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-radius-server"]) {
			return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_macauth_radius_server", flattenWirelessControllerVapCaptivePortalMacauthRadiusServer(o["captive-portal-macauth-radius-server"], d, "captive_portal_macauth_radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-macauth-radius-server"]) {
			return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_ac_name", flattenWirelessControllerVapCaptivePortalAcName(o["captive-portal-ac-name"], d, "captive_portal_ac_name", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-ac-name"]) {
			return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
		}
	}

	if err = d.Set("captive_portal_auth_timeout", flattenWirelessControllerVapCaptivePortalAuthTimeout(o["captive-portal-auth-timeout"], d, "captive_portal_auth_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-auth-timeout"]) {
			return fmt.Errorf("Error reading captive_portal_auth_timeout: %v", err)
		}
	}

	if err = d.Set("captive_portal_session_timeout_interval", flattenWirelessControllerVapCaptivePortalSessionTimeoutInterval(o["captive-portal-session-timeout-interval"], d, "captive_portal_session_timeout_interval", sv)); err != nil {
		if !fortiAPIPatch(o["captive-portal-session-timeout-interval"]) {
			return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
		}
	}

	if err = d.Set("alias", flattenWirelessControllerVapAlias(o["alias"], d, "alias", sv)); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("multicast_rate", flattenWirelessControllerVapMulticastRate(o["multicast-rate"], d, "multicast_rate", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-rate"]) {
			return fmt.Errorf("Error reading multicast_rate: %v", err)
		}
	}

	if err = d.Set("multicast_enhance", flattenWirelessControllerVapMulticastEnhance(o["multicast-enhance"], d, "multicast_enhance", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-enhance"]) {
			return fmt.Errorf("Error reading multicast_enhance: %v", err)
		}
	}

	if err = d.Set("igmp_snooping", flattenWirelessControllerVapIgmpSnooping(o["igmp-snooping"], d, "igmp_snooping", sv)); err != nil {
		if !fortiAPIPatch(o["igmp-snooping"]) {
			return fmt.Errorf("Error reading igmp_snooping: %v", err)
		}
	}

	if err = d.Set("dhcp_address_enforcement", flattenWirelessControllerVapDhcpAddressEnforcement(o["dhcp-address-enforcement"], d, "dhcp_address_enforcement", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-address-enforcement"]) {
			return fmt.Errorf("Error reading dhcp_address_enforcement: %v", err)
		}
	}

	if err = d.Set("broadcast_suppression", flattenWirelessControllerVapBroadcastSuppression(o["broadcast-suppression"], d, "broadcast_suppression", sv)); err != nil {
		if !fortiAPIPatch(o["broadcast-suppression"]) {
			return fmt.Errorf("Error reading broadcast_suppression: %v", err)
		}
	}

	if err = d.Set("ipv6_rules", flattenWirelessControllerVapIpv6Rules(o["ipv6-rules"], d, "ipv6_rules", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-rules"]) {
			return fmt.Errorf("Error reading ipv6_rules: %v", err)
		}
	}

	if err = d.Set("me_disable_thresh", flattenWirelessControllerVapMeDisableThresh(o["me-disable-thresh"], d, "me_disable_thresh", sv)); err != nil {
		if !fortiAPIPatch(o["me-disable-thresh"]) {
			return fmt.Errorf("Error reading me_disable_thresh: %v", err)
		}
	}

	if err = d.Set("mu_mimo", flattenWirelessControllerVapMuMimo(o["mu-mimo"], d, "mu_mimo", sv)); err != nil {
		if !fortiAPIPatch(o["mu-mimo"]) {
			return fmt.Errorf("Error reading mu_mimo: %v", err)
		}
	}

	if err = d.Set("probe_resp_suppression", flattenWirelessControllerVapProbeRespSuppression(o["probe-resp-suppression"], d, "probe_resp_suppression", sv)); err != nil {
		if !fortiAPIPatch(o["probe-resp-suppression"]) {
			return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
		}
	}

	if err = d.Set("probe_resp_threshold", flattenWirelessControllerVapProbeRespThreshold(o["probe-resp-threshold"], d, "probe_resp_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["probe-resp-threshold"]) {
			return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
		}
	}

	if err = d.Set("radio_sensitivity", flattenWirelessControllerVapRadioSensitivity(o["radio-sensitivity"], d, "radio_sensitivity", sv)); err != nil {
		if !fortiAPIPatch(o["radio-sensitivity"]) {
			return fmt.Errorf("Error reading radio_sensitivity: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenWirelessControllerVapQuarantine(o["quarantine"], d, "quarantine", sv)); err != nil {
		if !fortiAPIPatch(o["quarantine"]) {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("radio_5g_threshold", flattenWirelessControllerVapRadio5GThreshold(o["radio-5g-threshold"], d, "radio_5g_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["radio-5g-threshold"]) {
			return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_2g_threshold", flattenWirelessControllerVapRadio2GThreshold(o["radio-2g-threshold"], d, "radio_2g_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["radio-2g-threshold"]) {
			return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vlan_name", flattenWirelessControllerVapVlanName(o["vlan-name"], d, "vlan_name", sv)); err != nil {
			if !fortiAPIPatch(o["vlan-name"]) {
				return fmt.Errorf("Error reading vlan_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_name"); ok {
			if err = d.Set("vlan_name", flattenWirelessControllerVapVlanName(o["vlan-name"], d, "vlan_name", sv)); err != nil {
				if !fortiAPIPatch(o["vlan-name"]) {
					return fmt.Errorf("Error reading vlan_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("vlan_pooling", flattenWirelessControllerVapVlanPooling(o["vlan-pooling"], d, "vlan_pooling", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-pooling"]) {
			return fmt.Errorf("Error reading vlan_pooling: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool", sv)); err != nil {
			if !fortiAPIPatch(o["vlan-pool"]) {
				return fmt.Errorf("Error reading vlan_pool: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_pool"); ok {
			if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool", sv)); err != nil {
				if !fortiAPIPatch(o["vlan-pool"]) {
					return fmt.Errorf("Error reading vlan_pool: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp_option43_insertion", flattenWirelessControllerVapDhcpOption43Insertion(o["dhcp-option43-insertion"], d, "dhcp_option43_insertion", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option43-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option43_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_insertion", flattenWirelessControllerVapDhcpOption82Insertion(o["dhcp-option82-insertion"], d, "dhcp_option82_insertion", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id_insertion", flattenWirelessControllerVapDhcpOption82CircuitIdInsertion(o["dhcp-option82-circuit-id-insertion"], d, "dhcp_option82_circuit_id_insertion", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-circuit-id-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id_insertion", flattenWirelessControllerVapDhcpOption82RemoteIdInsertion(o["dhcp-option82-remote-id-insertion"], d, "dhcp_option82_remote_id_insertion", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-remote-id-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
		}
	}

	if err = d.Set("ptk_rekey", flattenWirelessControllerVapPtkRekey(o["ptk-rekey"], d, "ptk_rekey", sv)); err != nil {
		if !fortiAPIPatch(o["ptk-rekey"]) {
			return fmt.Errorf("Error reading ptk_rekey: %v", err)
		}
	}

	if err = d.Set("ptk_rekey_intv", flattenWirelessControllerVapPtkRekeyIntv(o["ptk-rekey-intv"], d, "ptk_rekey_intv", sv)); err != nil {
		if !fortiAPIPatch(o["ptk-rekey-intv"]) {
			return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("gtk_rekey", flattenWirelessControllerVapGtkRekey(o["gtk-rekey"], d, "gtk_rekey", sv)); err != nil {
		if !fortiAPIPatch(o["gtk-rekey"]) {
			return fmt.Errorf("Error reading gtk_rekey: %v", err)
		}
	}

	if err = d.Set("gtk_rekey_intv", flattenWirelessControllerVapGtkRekeyIntv(o["gtk-rekey-intv"], d, "gtk_rekey_intv", sv)); err != nil {
		if !fortiAPIPatch(o["gtk-rekey-intv"]) {
			return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("eap_reauth", flattenWirelessControllerVapEapReauth(o["eap-reauth"], d, "eap_reauth", sv)); err != nil {
		if !fortiAPIPatch(o["eap-reauth"]) {
			return fmt.Errorf("Error reading eap_reauth: %v", err)
		}
	}

	if err = d.Set("eap_reauth_intv", flattenWirelessControllerVapEapReauthIntv(o["eap-reauth-intv"], d, "eap_reauth_intv", sv)); err != nil {
		if !fortiAPIPatch(o["eap-reauth-intv"]) {
			return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
		}
	}

	if err = d.Set("roaming_acct_interim_update", flattenWirelessControllerVapRoamingAcctInterimUpdate(o["roaming-acct-interim-update"], d, "roaming_acct_interim_update", sv)); err != nil {
		if !fortiAPIPatch(o["roaming-acct-interim-update"]) {
			return fmt.Errorf("Error reading roaming_acct_interim_update: %v", err)
		}
	}

	if err = d.Set("qos_profile", flattenWirelessControllerVapQosProfile(o["qos-profile"], d, "qos_profile", sv)); err != nil {
		if !fortiAPIPatch(o["qos-profile"]) {
			return fmt.Errorf("Error reading qos_profile: %v", err)
		}
	}

	if err = d.Set("hotspot20_profile", flattenWirelessControllerVapHotspot20Profile(o["hotspot20-profile"], d, "hotspot20_profile", sv)); err != nil {
		if !fortiAPIPatch(o["hotspot20-profile"]) {
			return fmt.Errorf("Error reading hotspot20_profile: %v", err)
		}
	}

	if err = d.Set("access_control_list", flattenWirelessControllerVapAccessControlList(o["access-control-list"], d, "access_control_list", sv)); err != nil {
		if !fortiAPIPatch(o["access-control-list"]) {
			return fmt.Errorf("Error reading access_control_list: %v", err)
		}
	}

	if err = d.Set("primary_wag_profile", flattenWirelessControllerVapPrimaryWagProfile(o["primary-wag-profile"], d, "primary_wag_profile", sv)); err != nil {
		if !fortiAPIPatch(o["primary-wag-profile"]) {
			return fmt.Errorf("Error reading primary_wag_profile: %v", err)
		}
	}

	if err = d.Set("secondary_wag_profile", flattenWirelessControllerVapSecondaryWagProfile(o["secondary-wag-profile"], d, "secondary_wag_profile", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-wag-profile"]) {
			return fmt.Errorf("Error reading secondary_wag_profile: %v", err)
		}
	}

	if err = d.Set("tunnel_echo_interval", flattenWirelessControllerVapTunnelEchoInterval(o["tunnel-echo-interval"], d, "tunnel_echo_interval", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-echo-interval"]) {
			return fmt.Errorf("Error reading tunnel_echo_interval: %v", err)
		}
	}

	if err = d.Set("tunnel_fallback_interval", flattenWirelessControllerVapTunnelFallbackInterval(o["tunnel-fallback-interval"], d, "tunnel_fallback_interval", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-fallback-interval"]) {
			return fmt.Errorf("Error reading tunnel_fallback_interval: %v", err)
		}
	}

	if err = d.Set("rates_11a", flattenWirelessControllerVapRates11A(o["rates-11a"], d, "rates_11a", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11a"]) {
			return fmt.Errorf("Error reading rates_11a: %v", err)
		}
	}

	if err = d.Set("rates_11bg", flattenWirelessControllerVapRates11Bg(o["rates-11bg"], d, "rates_11bg", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11bg"]) {
			return fmt.Errorf("Error reading rates_11bg: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss12", flattenWirelessControllerVapRates11NSs12(o["rates-11n-ss12"], d, "rates_11n_ss12", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11n-ss12"]) {
			return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss34", flattenWirelessControllerVapRates11NSs34(o["rates-11n-ss34"], d, "rates_11n_ss34", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11n-ss34"]) {
			return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11ac_mcs_map", flattenWirelessControllerVapRates11AcMcsMap(o["rates-11ac-mcs-map"], d, "rates_11ac_mcs_map", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11ac-mcs-map"]) {
			return fmt.Errorf("Error reading rates_11ac_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11ax_mcs_map", flattenWirelessControllerVapRates11AxMcsMap(o["rates-11ax-mcs-map"], d, "rates_11ax_mcs_map", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11ax-mcs-map"]) {
			return fmt.Errorf("Error reading rates_11ax_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11be_mcs_map", flattenWirelessControllerVapRates11BeMcsMap(o["rates-11be-mcs-map"], d, "rates_11be_mcs_map", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11be-mcs-map"]) {
			return fmt.Errorf("Error reading rates_11be_mcs_map: %v", err)
		}
	}

	if err = d.Set("rates_11be_mcs_map_160", flattenWirelessControllerVapRates11BeMcsMap160(o["rates-11be-mcs-map-160"], d, "rates_11be_mcs_map_160", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11be-mcs-map-160"]) {
			return fmt.Errorf("Error reading rates_11be_mcs_map_160: %v", err)
		}
	}

	if err = d.Set("rates_11be_mcs_map_320", flattenWirelessControllerVapRates11BeMcsMap320(o["rates-11be-mcs-map-320"], d, "rates_11be_mcs_map_320", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11be-mcs-map-320"]) {
			return fmt.Errorf("Error reading rates_11be_mcs_map_320: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss12", flattenWirelessControllerVapRates11AcSs12(o["rates-11ac-ss12"], d, "rates_11ac_ss12", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11ac-ss12"]) {
			return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss34", flattenWirelessControllerVapRates11AcSs34(o["rates-11ac-ss34"], d, "rates_11ac_ss34", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11ac-ss34"]) {
			return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss12", flattenWirelessControllerVapRates11AxSs12(o["rates-11ax-ss12"], d, "rates_11ax_ss12", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11ax-ss12"]) {
			return fmt.Errorf("Error reading rates_11ax_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ax_ss34", flattenWirelessControllerVapRates11AxSs34(o["rates-11ax-ss34"], d, "rates_11ax_ss34", sv)); err != nil {
		if !fortiAPIPatch(o["rates-11ax-ss34"]) {
			return fmt.Errorf("Error reading rates_11ax_ss34: %v", err)
		}
	}

	if err = d.Set("utm_profile", flattenWirelessControllerVapUtmProfile(o["utm-profile"], d, "utm_profile", sv)); err != nil {
		if !fortiAPIPatch(o["utm-profile"]) {
			return fmt.Errorf("Error reading utm_profile: %v", err)
		}
	}

	if err = d.Set("utm_status", flattenWirelessControllerVapUtmStatus(o["utm-status"], d, "utm_status", sv)); err != nil {
		if !fortiAPIPatch(o["utm-status"]) {
			return fmt.Errorf("Error reading utm_status: %v", err)
		}
	}

	if err = d.Set("utm_log", flattenWirelessControllerVapUtmLog(o["utm-log"], d, "utm_log", sv)); err != nil {
		if !fortiAPIPatch(o["utm-log"]) {
			return fmt.Errorf("Error reading utm_log: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenWirelessControllerVapIpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenWirelessControllerVapApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("antivirus_profile", flattenWirelessControllerVapAntivirusProfile(o["antivirus-profile"], d, "antivirus_profile", sv)); err != nil {
		if !fortiAPIPatch(o["antivirus-profile"]) {
			return fmt.Errorf("Error reading antivirus_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenWirelessControllerVapWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenWirelessControllerVapScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections", sv)); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("address_group", flattenWirelessControllerVapAddressGroup(o["address-group"], d, "address_group", sv)); err != nil {
		if !fortiAPIPatch(o["address-group"]) {
			return fmt.Errorf("Error reading address_group: %v", err)
		}
	}

	if err = d.Set("address_group_policy", flattenWirelessControllerVapAddressGroupPolicy(o["address-group-policy"], d, "address_group_policy", sv)); err != nil {
		if !fortiAPIPatch(o["address-group-policy"]) {
			return fmt.Errorf("Error reading address_group_policy: %v", err)
		}
	}

	if err = d.Set("mac_filter", flattenWirelessControllerVapMacFilter(o["mac-filter"], d, "mac_filter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-filter"]) {
			return fmt.Errorf("Error reading mac_filter: %v", err)
		}
	}

	if err = d.Set("mac_filter_policy_other", flattenWirelessControllerVapMacFilterPolicyOther(o["mac-filter-policy-other"], d, "mac_filter_policy_other", sv)); err != nil {
		if !fortiAPIPatch(o["mac-filter-policy-other"]) {
			return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list", sv)); err != nil {
			if !fortiAPIPatch(o["mac-filter-list"]) {
				return fmt.Errorf("Error reading mac_filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_filter_list"); ok {
			if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list", sv)); err != nil {
				if !fortiAPIPatch(o["mac-filter-list"]) {
					return fmt.Errorf("Error reading mac_filter_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("sticky_client_remove", flattenWirelessControllerVapStickyClientRemove(o["sticky-client-remove"], d, "sticky_client_remove", sv)); err != nil {
		if !fortiAPIPatch(o["sticky-client-remove"]) {
			return fmt.Errorf("Error reading sticky_client_remove: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_5g", flattenWirelessControllerVapStickyClientThreshold5G(o["sticky-client-threshold-5g"], d, "sticky_client_threshold_5g", sv)); err != nil {
		if !fortiAPIPatch(o["sticky-client-threshold-5g"]) {
			return fmt.Errorf("Error reading sticky_client_threshold_5g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_2g", flattenWirelessControllerVapStickyClientThreshold2G(o["sticky-client-threshold-2g"], d, "sticky_client_threshold_2g", sv)); err != nil {
		if !fortiAPIPatch(o["sticky-client-threshold-2g"]) {
			return fmt.Errorf("Error reading sticky_client_threshold_2g: %v", err)
		}
	}

	if err = d.Set("sticky_client_threshold_6g", flattenWirelessControllerVapStickyClientThreshold6G(o["sticky-client-threshold-6g"], d, "sticky_client_threshold_6g", sv)); err != nil {
		if !fortiAPIPatch(o["sticky-client-threshold-6g"]) {
			return fmt.Errorf("Error reading sticky_client_threshold_6g: %v", err)
		}
	}

	if err = d.Set("bstm_rssi_disassoc_timer", flattenWirelessControllerVapBstmRssiDisassocTimer(o["bstm-rssi-disassoc-timer"], d, "bstm_rssi_disassoc_timer", sv)); err != nil {
		if !fortiAPIPatch(o["bstm-rssi-disassoc-timer"]) {
			return fmt.Errorf("Error reading bstm_rssi_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("bstm_load_balancing_disassoc_timer", flattenWirelessControllerVapBstmLoadBalancingDisassocTimer(o["bstm-load-balancing-disassoc-timer"], d, "bstm_load_balancing_disassoc_timer", sv)); err != nil {
		if !fortiAPIPatch(o["bstm-load-balancing-disassoc-timer"]) {
			return fmt.Errorf("Error reading bstm_load_balancing_disassoc_timer: %v", err)
		}
	}

	if err = d.Set("bstm_disassociation_imminent", flattenWirelessControllerVapBstmDisassociationImminent(o["bstm-disassociation-imminent"], d, "bstm_disassociation_imminent", sv)); err != nil {
		if !fortiAPIPatch(o["bstm-disassociation-imminent"]) {
			return fmt.Errorf("Error reading bstm_disassociation_imminent: %v", err)
		}
	}

	if err = d.Set("beacon_advertising", flattenWirelessControllerVapBeaconAdvertising(o["beacon-advertising"], d, "beacon_advertising", sv)); err != nil {
		if !fortiAPIPatch(o["beacon-advertising"]) {
			return fmt.Errorf("Error reading beacon_advertising: %v", err)
		}
	}

	if err = d.Set("osen", flattenWirelessControllerVapOsen(o["osen"], d, "osen", sv)); err != nil {
		if !fortiAPIPatch(o["osen"]) {
			return fmt.Errorf("Error reading osen: %v", err)
		}
	}

	if err = d.Set("application_detection_engine", flattenWirelessControllerVapApplicationDetectionEngine(o["application-detection-engine"], d, "application_detection_engine", sv)); err != nil {
		if !fortiAPIPatch(o["application-detection-engine"]) {
			return fmt.Errorf("Error reading application_detection_engine: %v", err)
		}
	}

	if err = d.Set("application_dscp_marking", flattenWirelessControllerVapApplicationDscpMarking(o["application-dscp-marking"], d, "application_dscp_marking", sv)); err != nil {
		if !fortiAPIPatch(o["application-dscp-marking"]) {
			return fmt.Errorf("Error reading application_dscp_marking: %v", err)
		}
	}

	if err = d.Set("application_report_intv", flattenWirelessControllerVapApplicationReportIntv(o["application-report-intv"], d, "application_report_intv", sv)); err != nil {
		if !fortiAPIPatch(o["application-report-intv"]) {
			return fmt.Errorf("Error reading application_report_intv: %v", err)
		}
	}

	if err = d.Set("l3_roaming", flattenWirelessControllerVapL3Roaming(o["l3-roaming"], d, "l3_roaming", sv)); err != nil {
		if !fortiAPIPatch(o["l3-roaming"]) {
			return fmt.Errorf("Error reading l3_roaming: %v", err)
		}
	}

	if err = d.Set("l3_roaming_mode", flattenWirelessControllerVapL3RoamingMode(o["l3-roaming-mode"], d, "l3_roaming_mode", sv)); err != nil {
		if !fortiAPIPatch(o["l3-roaming-mode"]) {
			return fmt.Errorf("Error reading l3_roaming_mode: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerVapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerVapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPreAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalPreAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFastRoaming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalFastRoaming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMeshBackhaul(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAtfWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMaxClientsAp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBroadcastSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityObsoleteOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmfAssocComebackTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmfSaQueryRetryTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBeaconProtection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOkc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMbo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGasComebackDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGasFragmentationLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMboCellDataConnPref(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVap80211K(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVap80211V(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVoiceEnterprise(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNeighborReportDualBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFastBssTransition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtMobilityDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtR0KeyLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtOverDs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaeGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOweGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOweTransition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOweTransitionSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAdditionalAkms(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapolKeyRetries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTkipCounterMeasure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalWeb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalWebFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalLogout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacCallingStationDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacCalledStationDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacCase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCalledStationIdType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacAuthBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthBlockInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacMpskAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacMpskTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWirelessControllerVapRadiusMacAuthUsergroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapRadiusMacAuthUsergroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEncrypt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapKeyindex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPassphrase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaeH2EOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaeHnpOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaePk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSaePrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAkm24Only(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNasFilterRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDomainNameStripping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandalone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandaloneNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpLeaseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandaloneDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandaloneDnsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalLanPartition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalBridging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalLan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUsergroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWirelessControllerVapUsergroupName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapUsergroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-disclaimer-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d, i["auth_disclaimer_page"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-reject-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthRejectPage(d, i["auth_reject_page"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-login-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthLoginPage(d, i["auth_login_page"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-login-failed-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d, i["auth_login_failed_page"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthRejectPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWirelessControllerVapSelectedUsergroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapSelectedUsergroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityExemptList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAuthCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAuthPortalAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntraVapPrivacy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"6.2.4"},
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

func expandWirelessControllerVapLdpc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapHighEfficiency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTargetWakeTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortMacauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortMacauthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortMacauthReauthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBssColorPartial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpsk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-name"], _ = expandWirelessControllerVapMpskKeyKeyName(d, i["key_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passphrase"], _ = expandWirelessControllerVapMpskKeyPassphrase(d, i["passphrase"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["passphrase"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["concurrent-clients"], _ = expandWirelessControllerVapMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["concurrent-clients"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWirelessControllerVapMpskKeyComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mpsk_schedules"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mpsk-schedules"], _ = expandWirelessControllerVapMpskKeyMpskSchedules(d, i["mpsk_schedules"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mpsk-schedules"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMpskKeyKeyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyMpskSchedules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWirelessControllerVapMpskKeyMpskSchedulesName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMpskKeyMpskSchedulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSplitTunneling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapNacProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanAuto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalFwAccounting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalRadiusSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalMacauthRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalMacauthRadiusSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalAcName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalAuthTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalSessionTimeoutInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAlias(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMulticastRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMulticastEnhance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIgmpSnooping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpAddressEnforcement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBroadcastSuppression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIpv6Rules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMeDisableThresh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMuMimo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapProbeRespSuppression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapProbeRespThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadioSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadio5GThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadio2GThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerVapVlanNameName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-id"], _ = expandWirelessControllerVapVlanNameVlanId(d, i["vlan_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vlan-id"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapVlanNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanNameVlanId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPooling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerVapVlanPoolId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wtp-group"], _ = expandWirelessControllerVapVlanPoolWtpGroup(d, i["wtp_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["wtp-group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapVlanPoolId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPoolWtpGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption43Insertion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82Insertion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82CircuitIdInsertion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82RemoteIdInsertion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPtkRekey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGtkRekey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapReauthIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRoamingAcctInterimUpdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapQosProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapHotspot20Profile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAccessControlList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPrimaryWagProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecondaryWagProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTunnelEchoInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTunnelFallbackInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11A(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11Bg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11NSs12(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11NSs34(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AcMcsMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AxMcsMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11BeMcsMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11BeMcsMap160(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11BeMcsMap320(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AcSs12(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AcSs34(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AxSs12(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AxSs34(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUtmProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUtmStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUtmLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAntivirusProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapWebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAddressGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAddressGroupPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterPolicyOther(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerVapMacFilterListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandWirelessControllerVapMacFilterListMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac-filter-policy"], _ = expandWirelessControllerVapMacFilterListMacFilterPolicy(d, i["mac_filter_policy"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMacFilterListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterListMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterListMacFilterPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientRemove(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientThreshold5G(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientThreshold2G(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapStickyClientThreshold6G(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBstmRssiDisassocTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBstmLoadBalancingDisassocTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBstmDisassociationImminent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBeaconAdvertising(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOsen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapApplicationDetectionEngine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapApplicationDscpMarking(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapApplicationReportIntv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapL3Roaming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapL3RoamingMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerVapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pre_auth"); ok {
		t, err := expandWirelessControllerVapPreAuth(d, v, "pre_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pre-auth"] = t
		}
	}

	if v, ok := d.GetOk("external_pre_auth"); ok {
		t, err := expandWirelessControllerVapExternalPreAuth(d, v, "external_pre_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-pre-auth"] = t
		}
	}

	if v, ok := d.GetOk("fast_roaming"); ok {
		t, err := expandWirelessControllerVapFastRoaming(d, v, "fast_roaming", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("external_fast_roaming"); ok {
		t, err := expandWirelessControllerVapExternalFastRoaming(d, v, "external_fast_roaming", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("mesh_backhaul"); ok {
		t, err := expandWirelessControllerVapMeshBackhaul(d, v, "mesh_backhaul", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-backhaul"] = t
		}
	}

	if v, ok := d.GetOkExists("atf_weight"); ok {
		t, err := expandWirelessControllerVapAtfWeight(d, v, "atf_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atf-weight"] = t
		}
	}

	if v, ok := d.GetOkExists("max_clients"); ok {
		t, err := expandWirelessControllerVapMaxClients(d, v, "max_clients", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	} else if d.HasChange("max_clients") {
		obj["max-clients"] = nil
	}

	if v, ok := d.GetOkExists("max_clients_ap"); ok {
		t, err := expandWirelessControllerVapMaxClientsAp(d, v, "max_clients_ap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients-ap"] = t
		}
	} else if d.HasChange("max_clients_ap") {
		obj["max-clients-ap"] = nil
	}

	if v, ok := d.GetOk("ssid"); ok {
		t, err := expandWirelessControllerVapSsid(d, v, "ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok {
		t, err := expandWirelessControllerVapBroadcastSsid(d, v, "broadcast_ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("security_obsolete_option"); ok {
		t, err := expandWirelessControllerVapSecurityObsoleteOption(d, v, "security_obsolete_option", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-obsolete-option"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok {
		t, err := expandWirelessControllerVapSecurity(d, v, "security", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok {
		t, err := expandWirelessControllerVapPmf(d, v, "pmf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("pmf_assoc_comeback_timeout"); ok {
		t, err := expandWirelessControllerVapPmfAssocComebackTimeout(d, v, "pmf_assoc_comeback_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-assoc-comeback-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pmf_sa_query_retry_timeout"); ok {
		t, err := expandWirelessControllerVapPmfSaQueryRetryTimeout(d, v, "pmf_sa_query_retry_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-sa-query-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("beacon_protection"); ok {
		t, err := expandWirelessControllerVapBeaconProtection(d, v, "beacon_protection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-protection"] = t
		}
	}

	if v, ok := d.GetOk("okc"); ok {
		t, err := expandWirelessControllerVapOkc(d, v, "okc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["okc"] = t
		}
	}

	if v, ok := d.GetOk("mbo"); ok {
		t, err := expandWirelessControllerVapMbo(d, v, "mbo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo"] = t
		}
	}

	if v, ok := d.GetOk("gas_comeback_delay"); ok {
		t, err := expandWirelessControllerVapGasComebackDelay(d, v, "gas_comeback_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-comeback-delay"] = t
		}
	}

	if v, ok := d.GetOk("gas_fragmentation_limit"); ok {
		t, err := expandWirelessControllerVapGasFragmentationLimit(d, v, "gas_fragmentation_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gas-fragmentation-limit"] = t
		}
	}

	if v, ok := d.GetOk("mbo_cell_data_conn_pref"); ok {
		t, err := expandWirelessControllerVapMboCellDataConnPref(d, v, "mbo_cell_data_conn_pref", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mbo-cell-data-conn-pref"] = t
		}
	}

	if v, ok := d.GetOk("n80211k"); ok {
		t, err := expandWirelessControllerVap80211K(d, v, "n80211k", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211k"] = t
		}
	}

	if v, ok := d.GetOk("n80211v"); ok {
		t, err := expandWirelessControllerVap80211V(d, v, "n80211v", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["80211v"] = t
		}
	}

	if v, ok := d.GetOk("voice_enterprise"); ok {
		t, err := expandWirelessControllerVapVoiceEnterprise(d, v, "voice_enterprise", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-enterprise"] = t
		}
	}

	if v, ok := d.GetOk("neighbor_report_dual_band"); ok {
		t, err := expandWirelessControllerVapNeighborReportDualBand(d, v, "neighbor_report_dual_band", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor-report-dual-band"] = t
		}
	}

	if v, ok := d.GetOk("fast_bss_transition"); ok {
		t, err := expandWirelessControllerVapFastBssTransition(d, v, "fast_bss_transition", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("ft_mobility_domain"); ok {
		t, err := expandWirelessControllerVapFtMobilityDomain(d, v, "ft_mobility_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-mobility-domain"] = t
		}
	}

	if v, ok := d.GetOk("ft_r0_key_lifetime"); ok {
		t, err := expandWirelessControllerVapFtR0KeyLifetime(d, v, "ft_r0_key_lifetime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-r0-key-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("ft_over_ds"); ok {
		t, err := expandWirelessControllerVapFtOverDs(d, v, "ft_over_ds", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-over-ds"] = t
		}
	}

	if v, ok := d.GetOk("sae_groups"); ok {
		t, err := expandWirelessControllerVapSaeGroups(d, v, "sae_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-groups"] = t
		}
	} else if d.HasChange("sae_groups") {
		obj["sae-groups"] = nil
	}

	if v, ok := d.GetOk("owe_groups"); ok {
		t, err := expandWirelessControllerVapOweGroups(d, v, "owe_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-groups"] = t
		}
	} else if d.HasChange("owe_groups") {
		obj["owe-groups"] = nil
	}

	if v, ok := d.GetOk("owe_transition"); ok {
		t, err := expandWirelessControllerVapOweTransition(d, v, "owe_transition", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition"] = t
		}
	}

	if v, ok := d.GetOk("owe_transition_ssid"); ok {
		t, err := expandWirelessControllerVapOweTransitionSsid(d, v, "owe_transition_ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owe-transition-ssid"] = t
		}
	} else if d.HasChange("owe_transition_ssid") {
		obj["owe-transition-ssid"] = nil
	}

	if v, ok := d.GetOk("additional_akms"); ok {
		t, err := expandWirelessControllerVapAdditionalAkms(d, v, "additional_akms", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["additional-akms"] = t
		}
	} else if d.HasChange("additional_akms") {
		obj["additional-akms"] = nil
	}

	if v, ok := d.GetOk("eapol_key_retries"); ok {
		t, err := expandWirelessControllerVapEapolKeyRetries(d, v, "eapol_key_retries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-retries"] = t
		}
	}

	if v, ok := d.GetOk("tkip_counter_measure"); ok {
		t, err := expandWirelessControllerVapTkipCounterMeasure(d, v, "tkip_counter_measure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tkip-counter-measure"] = t
		}
	}

	if v, ok := d.GetOk("external_web"); ok {
		t, err := expandWirelessControllerVapExternalWeb(d, v, "external_web", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web"] = t
		}
	} else if d.HasChange("external_web") {
		obj["external-web"] = nil
	}

	if v, ok := d.GetOk("external_web_format"); ok {
		t, err := expandWirelessControllerVapExternalWebFormat(d, v, "external_web_format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web-format"] = t
		}
	}

	if v, ok := d.GetOk("external_logout"); ok {
		t, err := expandWirelessControllerVapExternalLogout(d, v, "external_logout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-logout"] = t
		}
	} else if d.HasChange("external_logout") {
		obj["external-logout"] = nil
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok {
		t, err := expandWirelessControllerVapMacUsernameDelimiter(d, v, "mac_username_delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok {
		t, err := expandWirelessControllerVapMacPasswordDelimiter(d, v, "mac_password_delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_calling_station_delimiter"); ok {
		t, err := expandWirelessControllerVapMacCallingStationDelimiter(d, v, "mac_calling_station_delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-calling-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_called_station_delimiter"); ok {
		t, err := expandWirelessControllerVapMacCalledStationDelimiter(d, v, "mac_called_station_delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-called-station-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok {
		t, err := expandWirelessControllerVapMacCase(d, v, "mac_case", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("called_station_id_type"); ok {
		t, err := expandWirelessControllerVapCalledStationIdType(d, v, "called_station_id_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["called-station-id-type"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok {
		t, err := expandWirelessControllerVapMacAuthBypass(d, v, "mac_auth_bypass", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth"); ok {
		t, err := expandWirelessControllerVapRadiusMacAuth(d, v, "radius_mac_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_server"); ok {
		t, err := expandWirelessControllerVapRadiusMacAuthServer(d, v, "radius_mac_auth_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-server"] = t
		}
	} else if d.HasChange("radius_mac_auth_server") {
		obj["radius-mac-auth-server"] = nil
	}

	if v, ok := d.GetOk("radius_mac_auth_block_interval"); ok {
		t, err := expandWirelessControllerVapRadiusMacAuthBlockInterval(d, v, "radius_mac_auth_block_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-block-interval"] = t
		}
	} else if d.HasChange("radius_mac_auth_block_interval") {
		obj["radius-mac-auth-block-interval"] = nil
	}

	if v, ok := d.GetOk("radius_mac_mpsk_auth"); ok {
		t, err := expandWirelessControllerVapRadiusMacMpskAuth(d, v, "radius_mac_mpsk_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_mpsk_timeout"); ok {
		t, err := expandWirelessControllerVapRadiusMacMpskTimeout(d, v, "radius_mac_mpsk_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-mpsk-timeout"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_usergroups"); ok || d.HasChange("radius_mac_auth_usergroups") {
		t, err := expandWirelessControllerVapRadiusMacAuthUsergroups(d, v, "radius_mac_auth_usergroups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		t, err := expandWirelessControllerVapAuth(d, v, "auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("encrypt"); ok {
		t, err := expandWirelessControllerVapEncrypt(d, v, "encrypt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt"] = t
		}
	}

	if v, ok := d.GetOk("keyindex"); ok {
		t, err := expandWirelessControllerVapKeyindex(d, v, "keyindex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyindex"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandWirelessControllerVapKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	} else if d.HasChange("key") {
		obj["key"] = nil
	}

	if v, ok := d.GetOk("passphrase"); ok {
		t, err := expandWirelessControllerVapPassphrase(d, v, "passphrase", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	} else if d.HasChange("passphrase") {
		obj["passphrase"] = nil
	}

	if v, ok := d.GetOk("sae_password"); ok {
		t, err := expandWirelessControllerVapSaePassword(d, v, "sae_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	} else if d.HasChange("sae_password") {
		obj["sae-password"] = nil
	}

	if v, ok := d.GetOk("sae_h2e_only"); ok {
		t, err := expandWirelessControllerVapSaeH2EOnly(d, v, "sae_h2e_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-h2e-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_hnp_only"); ok {
		t, err := expandWirelessControllerVapSaeHnpOnly(d, v, "sae_hnp_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-hnp-only"] = t
		}
	}

	if v, ok := d.GetOk("sae_pk"); ok {
		t, err := expandWirelessControllerVapSaePk(d, v, "sae_pk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-pk"] = t
		}
	}

	if v, ok := d.GetOk("sae_private_key"); ok {
		t, err := expandWirelessControllerVapSaePrivateKey(d, v, "sae_private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-private-key"] = t
		}
	} else if d.HasChange("sae_private_key") {
		obj["sae-private-key"] = nil
	}

	if v, ok := d.GetOk("akm24_only"); ok {
		t, err := expandWirelessControllerVapAkm24Only(d, v, "akm24_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["akm24-only"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok {
		t, err := expandWirelessControllerVapRadiusServer(d, v, "radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	} else if d.HasChange("radius_server") {
		obj["radius-server"] = nil
	}

	if v, ok := d.GetOk("nas_filter_rule"); ok {
		t, err := expandWirelessControllerVapNasFilterRule(d, v, "nas_filter_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-filter-rule"] = t
		}
	}

	if v, ok := d.GetOk("domain_name_stripping"); ok {
		t, err := expandWirelessControllerVapDomainNameStripping(d, v, "domain_name_stripping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name-stripping"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok {
		t, err := expandWirelessControllerVapAcctInterimInterval(d, v, "acct_interim_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	} else if d.HasChange("acct_interim_interval") {
		obj["acct-interim-interval"] = nil
	}

	if v, ok := d.GetOk("local_standalone"); ok {
		t, err := expandWirelessControllerVapLocalStandalone(d, v, "local_standalone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_nat"); ok {
		t, err := expandWirelessControllerVapLocalStandaloneNat(d, v, "local_standalone_nat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-nat"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWirelessControllerVapIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_lease_time"); ok {
		t, err := expandWirelessControllerVapDhcpLeaseTime(d, v, "dhcp_lease_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-lease-time"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns"); ok {
		t, err := expandWirelessControllerVapLocalStandaloneDns(d, v, "local_standalone_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_dns_ip"); ok {
		t, err := expandWirelessControllerVapLocalStandaloneDnsIp(d, v, "local_standalone_dns_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-dns-ip"] = t
		}
	} else if d.HasChange("local_standalone_dns_ip") {
		obj["local-standalone-dns-ip"] = nil
	}

	if v, ok := d.GetOk("local_lan_partition"); ok {
		t, err := expandWirelessControllerVapLocalLanPartition(d, v, "local_lan_partition", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan-partition"] = t
		}
	}

	if v, ok := d.GetOk("local_bridging"); ok {
		t, err := expandWirelessControllerVapLocalBridging(d, v, "local_bridging", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-bridging"] = t
		}
	}

	if v, ok := d.GetOk("local_lan"); ok {
		t, err := expandWirelessControllerVapLocalLan(d, v, "local_lan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan"] = t
		}
	}

	if v, ok := d.GetOk("local_authentication"); ok {
		t, err := expandWirelessControllerVapLocalAuthentication(d, v, "local_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-authentication"] = t
		}
	}

	if v, ok := d.GetOk("usergroup"); ok || d.HasChange("usergroup") {
		t, err := expandWirelessControllerVapUsergroup(d, v, "usergroup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroup"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_override_group"); ok {
		t, err := expandWirelessControllerVapPortalMessageOverrideGroup(d, v, "portal_message_override_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-override-group"] = t
		}
	} else if d.HasChange("portal_message_override_group") {
		obj["portal-message-override-group"] = nil
	}

	if v, ok := d.GetOk("portal_message_overrides"); ok {
		t, err := expandWirelessControllerVapPortalMessageOverrides(d, v, "portal_message_overrides", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-overrides"] = t
		}
	}

	if v, ok := d.GetOk("portal_type"); ok {
		t, err := expandWirelessControllerVapPortalType(d, v, "portal_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-type"] = t
		}
	}

	if v, ok := d.GetOk("selected_usergroups"); ok || d.HasChange("selected_usergroups") {
		t, err := expandWirelessControllerVapSelectedUsergroups(d, v, "selected_usergroups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok {
		t, err := expandWirelessControllerVapSecurityExemptList(d, v, "security_exempt_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	} else if d.HasChange("security_exempt_list") {
		obj["security-exempt-list"] = nil
	}

	if v, ok := d.GetOk("security_redirect_url"); ok {
		t, err := expandWirelessControllerVapSecurityRedirectUrl(d, v, "security_redirect_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	} else if d.HasChange("security_redirect_url") {
		obj["security-redirect-url"] = nil
	}

	if v, ok := d.GetOk("auth_cert"); ok {
		t, err := expandWirelessControllerVapAuthCert(d, v, "auth_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-cert"] = t
		}
	} else if d.HasChange("auth_cert") {
		obj["auth-cert"] = nil
	}

	if v, ok := d.GetOk("auth_portal_addr"); ok {
		t, err := expandWirelessControllerVapAuthPortalAddr(d, v, "auth_portal_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal-addr"] = t
		}
	} else if d.HasChange("auth_portal_addr") {
		obj["auth-portal-addr"] = nil
	}

	if v, ok := d.GetOk("intra_vap_privacy"); ok {
		t, err := expandWirelessControllerVapIntraVapPrivacy(d, v, "intra_vap_privacy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-vap-privacy"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandWirelessControllerVapSchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	} else if d.HasChange("schedule") {
		obj["schedule"] = nil
	}

	if v, ok := d.GetOk("ldpc"); ok {
		t, err := expandWirelessControllerVapLdpc(d, v, "ldpc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldpc"] = t
		}
	}

	if v, ok := d.GetOk("high_efficiency"); ok {
		t, err := expandWirelessControllerVapHighEfficiency(d, v, "high_efficiency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high-efficiency"] = t
		}
	}

	if v, ok := d.GetOk("target_wake_time"); ok {
		t, err := expandWirelessControllerVapTargetWakeTime(d, v, "target_wake_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth"); ok {
		t, err := expandWirelessControllerVapPortMacauth(d, v, "port_macauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_timeout"); ok {
		t, err := expandWirelessControllerVapPortMacauthTimeout(d, v, "port_macauth_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("port_macauth_reauth_timeout"); ok {
		t, err := expandWirelessControllerVapPortMacauthReauthTimeout(d, v, "port_macauth_reauth_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port-macauth-reauth-timeout"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_partial"); ok {
		t, err := expandWirelessControllerVapBssColorPartial(d, v, "bss_color_partial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-partial"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_profile"); ok {
		t, err := expandWirelessControllerVapMpskProfile(d, v, "mpsk_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-profile"] = t
		}
	} else if d.HasChange("mpsk_profile") {
		obj["mpsk-profile"] = nil
	}

	if v, ok := d.GetOk("mpsk"); ok {
		t, err := expandWirelessControllerVapMpsk(d, v, "mpsk", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk"] = t
		}
	} else if d.HasChange("mpsk") {
		obj["mpsk"] = nil
	}

	if v, ok := d.GetOkExists("mpsk_concurrent_clients"); ok {
		t, err := expandWirelessControllerVapMpskConcurrentClients(d, v, "mpsk_concurrent_clients", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	} else if d.HasChange("mpsk_concurrent_clients") {
		obj["mpsk-concurrent-clients"] = nil
	}

	if v, ok := d.GetOk("mpsk_key"); ok || d.HasChange("mpsk_key") {
		t, err := expandWirelessControllerVapMpskKey(d, v, "mpsk_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-key"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok {
		t, err := expandWirelessControllerVapSplitTunneling(d, v, "split_tunneling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("nac"); ok {
		t, err := expandWirelessControllerVapNac(d, v, "nac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac"] = t
		}
	}

	if v, ok := d.GetOk("nac_profile"); ok {
		t, err := expandWirelessControllerVapNacProfile(d, v, "nac_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-profile"] = t
		}
	} else if d.HasChange("nac_profile") {
		obj["nac-profile"] = nil
	}

	if v, ok := d.GetOkExists("vlanid"); ok {
		t, err := expandWirelessControllerVapVlanid(d, v, "vlanid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	} else if d.HasChange("vlanid") {
		obj["vlanid"] = nil
	}

	if v, ok := d.GetOk("vlan_auto"); ok {
		t, err := expandWirelessControllerVapVlanAuto(d, v, "vlan_auto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-auto"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_vlan"); ok {
		t, err := expandWirelessControllerVapDynamicVlan(d, v, "dynamic_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-vlan"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok {
		t, err := expandWirelessControllerVapCaptivePortal(d, v, "captive_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_fw_accounting"); ok {
		t, err := expandWirelessControllerVapCaptivePortalFwAccounting(d, v, "captive_portal_fw_accounting", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-fw-accounting"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_server"); ok {
		t, err := expandWirelessControllerVapCaptivePortalRadiusServer(d, v, "captive_portal_radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-server"] = t
		}
	} else if d.HasChange("captive_portal_radius_server") {
		obj["captive-portal-radius-server"] = nil
	}

	if v, ok := d.GetOk("captive_portal_radius_secret"); ok {
		t, err := expandWirelessControllerVapCaptivePortalRadiusSecret(d, v, "captive_portal_radius_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-secret"] = t
		}
	} else if d.HasChange("captive_portal_radius_secret") {
		obj["captive-portal-radius-secret"] = nil
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_server"); ok {
		t, err := expandWirelessControllerVapCaptivePortalMacauthRadiusServer(d, v, "captive_portal_macauth_radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-server"] = t
		}
	} else if d.HasChange("captive_portal_macauth_radius_server") {
		obj["captive-portal-macauth-radius-server"] = nil
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_secret"); ok {
		t, err := expandWirelessControllerVapCaptivePortalMacauthRadiusSecret(d, v, "captive_portal_macauth_radius_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-secret"] = t
		}
	} else if d.HasChange("captive_portal_macauth_radius_secret") {
		obj["captive-portal-macauth-radius-secret"] = nil
	}

	if v, ok := d.GetOk("captive_portal_ac_name"); ok {
		t, err := expandWirelessControllerVapCaptivePortalAcName(d, v, "captive_portal_ac_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ac-name"] = t
		}
	} else if d.HasChange("captive_portal_ac_name") {
		obj["captive-portal-ac-name"] = nil
	}

	if v, ok := d.GetOkExists("captive_portal_auth_timeout"); ok {
		t, err := expandWirelessControllerVapCaptivePortalAuthTimeout(d, v, "captive_portal_auth_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-auth-timeout"] = t
		}
	} else if d.HasChange("captive_portal_auth_timeout") {
		obj["captive-portal-auth-timeout"] = nil
	}

	if v, ok := d.GetOkExists("captive_portal_session_timeout_interval"); ok {
		t, err := expandWirelessControllerVapCaptivePortalSessionTimeoutInterval(d, v, "captive_portal_session_timeout_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-session-timeout-interval"] = t
		}
	} else if d.HasChange("captive_portal_session_timeout_interval") {
		obj["captive-portal-session-timeout-interval"] = nil
	}

	if v, ok := d.GetOk("alias"); ok {
		t, err := expandWirelessControllerVapAlias(d, v, "alias", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("multicast_rate"); ok {
		t, err := expandWirelessControllerVapMulticastRate(d, v, "multicast_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-rate"] = t
		}
	}

	if v, ok := d.GetOk("multicast_enhance"); ok {
		t, err := expandWirelessControllerVapMulticastEnhance(d, v, "multicast_enhance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-enhance"] = t
		}
	}

	if v, ok := d.GetOk("igmp_snooping"); ok {
		t, err := expandWirelessControllerVapIgmpSnooping(d, v, "igmp_snooping", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_address_enforcement"); ok {
		t, err := expandWirelessControllerVapDhcpAddressEnforcement(d, v, "dhcp_address_enforcement", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-address-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_suppression"); ok {
		t, err := expandWirelessControllerVapBroadcastSuppression(d, v, "broadcast_suppression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-suppression"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_rules"); ok {
		t, err := expandWirelessControllerVapIpv6Rules(d, v, "ipv6_rules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-rules"] = t
		}
	}

	if v, ok := d.GetOk("me_disable_thresh"); ok {
		t, err := expandWirelessControllerVapMeDisableThresh(d, v, "me_disable_thresh", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["me-disable-thresh"] = t
		}
	}

	if v, ok := d.GetOk("mu_mimo"); ok {
		t, err := expandWirelessControllerVapMuMimo(d, v, "mu_mimo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mu-mimo"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_suppression"); ok {
		t, err := expandWirelessControllerVapProbeRespSuppression(d, v, "probe_resp_suppression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_threshold"); ok {
		t, err := expandWirelessControllerVapProbeRespThreshold(d, v, "probe_resp_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_sensitivity"); ok {
		t, err := expandWirelessControllerVapRadioSensitivity(d, v, "radio_sensitivity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok {
		t, err := expandWirelessControllerVapQuarantine(d, v, "quarantine", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("radio_5g_threshold"); ok {
		t, err := expandWirelessControllerVapRadio5GThreshold(d, v, "radio_5g_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-5g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_2g_threshold"); ok {
		t, err := expandWirelessControllerVapRadio2GThreshold(d, v, "radio_2g_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("vlan_name"); ok || d.HasChange("vlan_name") {
		t, err := expandWirelessControllerVapVlanName(d, v, "vlan_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-name"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pooling"); ok {
		t, err := expandWirelessControllerVapVlanPooling(d, v, "vlan_pooling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pooling"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pool"); ok || d.HasChange("vlan_pool") {
		t, err := expandWirelessControllerVapVlanPool(d, v, "vlan_pool", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pool"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option43_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption43Insertion(d, v, "dhcp_option43_insertion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option43-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption82Insertion(d, v, "dhcp_option82_insertion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption82CircuitIdInsertion(d, v, "dhcp_option82_circuit_id_insertion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-circuit-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption82RemoteIdInsertion(d, v, "dhcp_option82_remote_id_insertion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-remote-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey"); ok {
		t, err := expandWirelessControllerVapPtkRekey(d, v, "ptk_rekey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey_intv"); ok {
		t, err := expandWirelessControllerVapPtkRekeyIntv(d, v, "ptk_rekey_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey"); ok {
		t, err := expandWirelessControllerVapGtkRekey(d, v, "gtk_rekey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey_intv"); ok {
		t, err := expandWirelessControllerVapGtkRekeyIntv(d, v, "gtk_rekey_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth"); ok {
		t, err := expandWirelessControllerVapEapReauth(d, v, "eap_reauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth_intv"); ok {
		t, err := expandWirelessControllerVapEapReauthIntv(d, v, "eap_reauth_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth-intv"] = t
		}
	}

	if v, ok := d.GetOk("roaming_acct_interim_update"); ok {
		t, err := expandWirelessControllerVapRoamingAcctInterimUpdate(d, v, "roaming_acct_interim_update", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-acct-interim-update"] = t
		}
	}

	if v, ok := d.GetOk("qos_profile"); ok {
		t, err := expandWirelessControllerVapQosProfile(d, v, "qos_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-profile"] = t
		}
	} else if d.HasChange("qos_profile") {
		obj["qos-profile"] = nil
	}

	if v, ok := d.GetOk("hotspot20_profile"); ok {
		t, err := expandWirelessControllerVapHotspot20Profile(d, v, "hotspot20_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspot20-profile"] = t
		}
	} else if d.HasChange("hotspot20_profile") {
		obj["hotspot20-profile"] = nil
	}

	if v, ok := d.GetOk("access_control_list"); ok {
		t, err := expandWirelessControllerVapAccessControlList(d, v, "access_control_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-control-list"] = t
		}
	} else if d.HasChange("access_control_list") {
		obj["access-control-list"] = nil
	}

	if v, ok := d.GetOk("primary_wag_profile"); ok {
		t, err := expandWirelessControllerVapPrimaryWagProfile(d, v, "primary_wag_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-wag-profile"] = t
		}
	} else if d.HasChange("primary_wag_profile") {
		obj["primary-wag-profile"] = nil
	}

	if v, ok := d.GetOk("secondary_wag_profile"); ok {
		t, err := expandWirelessControllerVapSecondaryWagProfile(d, v, "secondary_wag_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-wag-profile"] = t
		}
	} else if d.HasChange("secondary_wag_profile") {
		obj["secondary-wag-profile"] = nil
	}

	if v, ok := d.GetOk("tunnel_echo_interval"); ok {
		t, err := expandWirelessControllerVapTunnelEchoInterval(d, v, "tunnel_echo_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-echo-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("tunnel_fallback_interval"); ok {
		t, err := expandWirelessControllerVapTunnelFallbackInterval(d, v, "tunnel_fallback_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-fallback-interval"] = t
		}
	}

	if v, ok := d.GetOk("rates_11a"); ok {
		t, err := expandWirelessControllerVapRates11A(d, v, "rates_11a", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11a"] = t
		}
	} else if d.HasChange("rates_11a") {
		obj["rates-11a"] = nil
	}

	if v, ok := d.GetOk("rates_11bg"); ok {
		t, err := expandWirelessControllerVapRates11Bg(d, v, "rates_11bg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11bg"] = t
		}
	} else if d.HasChange("rates_11bg") {
		obj["rates-11bg"] = nil
	}

	if v, ok := d.GetOk("rates_11n_ss12"); ok {
		t, err := expandWirelessControllerVapRates11NSs12(d, v, "rates_11n_ss12", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss12"] = t
		}
	} else if d.HasChange("rates_11n_ss12") {
		obj["rates-11n-ss12"] = nil
	}

	if v, ok := d.GetOk("rates_11n_ss34"); ok {
		t, err := expandWirelessControllerVapRates11NSs34(d, v, "rates_11n_ss34", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss34"] = t
		}
	} else if d.HasChange("rates_11n_ss34") {
		obj["rates-11n-ss34"] = nil
	}

	if v, ok := d.GetOk("rates_11ac_mcs_map"); ok {
		t, err := expandWirelessControllerVapRates11AcMcsMap(d, v, "rates_11ac_mcs_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-mcs-map"] = t
		}
	} else if d.HasChange("rates_11ac_mcs_map") {
		obj["rates-11ac-mcs-map"] = nil
	}

	if v, ok := d.GetOk("rates_11ax_mcs_map"); ok {
		t, err := expandWirelessControllerVapRates11AxMcsMap(d, v, "rates_11ax_mcs_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-mcs-map"] = t
		}
	} else if d.HasChange("rates_11ax_mcs_map") {
		obj["rates-11ax-mcs-map"] = nil
	}

	if v, ok := d.GetOk("rates_11be_mcs_map"); ok {
		t, err := expandWirelessControllerVapRates11BeMcsMap(d, v, "rates_11be_mcs_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11be-mcs-map"] = t
		}
	} else if d.HasChange("rates_11be_mcs_map") {
		obj["rates-11be-mcs-map"] = nil
	}

	if v, ok := d.GetOk("rates_11be_mcs_map_160"); ok {
		t, err := expandWirelessControllerVapRates11BeMcsMap160(d, v, "rates_11be_mcs_map_160", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11be-mcs-map-160"] = t
		}
	} else if d.HasChange("rates_11be_mcs_map_160") {
		obj["rates-11be-mcs-map-160"] = nil
	}

	if v, ok := d.GetOk("rates_11be_mcs_map_320"); ok {
		t, err := expandWirelessControllerVapRates11BeMcsMap320(d, v, "rates_11be_mcs_map_320", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11be-mcs-map-320"] = t
		}
	} else if d.HasChange("rates_11be_mcs_map_320") {
		obj["rates-11be-mcs-map-320"] = nil
	}

	if v, ok := d.GetOk("rates_11ac_ss12"); ok {
		t, err := expandWirelessControllerVapRates11AcSs12(d, v, "rates_11ac_ss12", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss12"] = t
		}
	} else if d.HasChange("rates_11ac_ss12") {
		obj["rates-11ac-ss12"] = nil
	}

	if v, ok := d.GetOk("rates_11ac_ss34"); ok {
		t, err := expandWirelessControllerVapRates11AcSs34(d, v, "rates_11ac_ss34", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss34"] = t
		}
	} else if d.HasChange("rates_11ac_ss34") {
		obj["rates-11ac-ss34"] = nil
	}

	if v, ok := d.GetOk("rates_11ax_ss12"); ok {
		t, err := expandWirelessControllerVapRates11AxSs12(d, v, "rates_11ax_ss12", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss12"] = t
		}
	} else if d.HasChange("rates_11ax_ss12") {
		obj["rates-11ax-ss12"] = nil
	}

	if v, ok := d.GetOk("rates_11ax_ss34"); ok {
		t, err := expandWirelessControllerVapRates11AxSs34(d, v, "rates_11ax_ss34", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ax-ss34"] = t
		}
	} else if d.HasChange("rates_11ax_ss34") {
		obj["rates-11ax-ss34"] = nil
	}

	if v, ok := d.GetOk("utm_profile"); ok {
		t, err := expandWirelessControllerVapUtmProfile(d, v, "utm_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-profile"] = t
		}
	} else if d.HasChange("utm_profile") {
		obj["utm-profile"] = nil
	}

	if v, ok := d.GetOk("utm_status"); ok {
		t, err := expandWirelessControllerVapUtmStatus(d, v, "utm_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-status"] = t
		}
	}

	if v, ok := d.GetOk("utm_log"); ok {
		t, err := expandWirelessControllerVapUtmLog(d, v, "utm_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-log"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandWirelessControllerVapIpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	} else if d.HasChange("ips_sensor") {
		obj["ips-sensor"] = nil
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandWirelessControllerVapApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	} else if d.HasChange("application_list") {
		obj["application-list"] = nil
	}

	if v, ok := d.GetOk("antivirus_profile"); ok {
		t, err := expandWirelessControllerVapAntivirusProfile(d, v, "antivirus_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antivirus-profile"] = t
		}
	} else if d.HasChange("antivirus_profile") {
		obj["antivirus-profile"] = nil
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandWirelessControllerVapWebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	} else if d.HasChange("webfilter_profile") {
		obj["webfilter-profile"] = nil
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandWirelessControllerVapScanBotnetConnections(d, v, "scan_botnet_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("address_group"); ok {
		t, err := expandWirelessControllerVapAddressGroup(d, v, "address_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group"] = t
		}
	} else if d.HasChange("address_group") {
		obj["address-group"] = nil
	}

	if v, ok := d.GetOk("address_group_policy"); ok {
		t, err := expandWirelessControllerVapAddressGroupPolicy(d, v, "address_group_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-group-policy"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter"); ok {
		t, err := expandWirelessControllerVapMacFilter(d, v, "mac_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_policy_other"); ok {
		t, err := expandWirelessControllerVapMacFilterPolicyOther(d, v, "mac_filter_policy_other", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-policy-other"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_list"); ok || d.HasChange("mac_filter_list") {
		t, err := expandWirelessControllerVapMacFilterList(d, v, "mac_filter_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-list"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_remove"); ok {
		t, err := expandWirelessControllerVapStickyClientRemove(d, v, "sticky_client_remove", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-remove"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_5g"); ok {
		t, err := expandWirelessControllerVapStickyClientThreshold5G(d, v, "sticky_client_threshold_5g", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-5g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_2g"); ok {
		t, err := expandWirelessControllerVapStickyClientThreshold2G(d, v, "sticky_client_threshold_2g", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-2g"] = t
		}
	}

	if v, ok := d.GetOk("sticky_client_threshold_6g"); ok {
		t, err := expandWirelessControllerVapStickyClientThreshold6G(d, v, "sticky_client_threshold_6g", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sticky-client-threshold-6g"] = t
		}
	}

	if v, ok := d.GetOk("bstm_rssi_disassoc_timer"); ok {
		t, err := expandWirelessControllerVapBstmRssiDisassocTimer(d, v, "bstm_rssi_disassoc_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-rssi-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("bstm_load_balancing_disassoc_timer"); ok {
		t, err := expandWirelessControllerVapBstmLoadBalancingDisassocTimer(d, v, "bstm_load_balancing_disassoc_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-load-balancing-disassoc-timer"] = t
		}
	}

	if v, ok := d.GetOk("bstm_disassociation_imminent"); ok {
		t, err := expandWirelessControllerVapBstmDisassociationImminent(d, v, "bstm_disassociation_imminent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bstm-disassociation-imminent"] = t
		}
	}

	if v, ok := d.GetOk("beacon_advertising"); ok {
		t, err := expandWirelessControllerVapBeaconAdvertising(d, v, "beacon_advertising", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-advertising"] = t
		}
	} else if d.HasChange("beacon_advertising") {
		obj["beacon-advertising"] = nil
	}

	if v, ok := d.GetOk("osen"); ok {
		t, err := expandWirelessControllerVapOsen(d, v, "osen", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["osen"] = t
		}
	}

	if v, ok := d.GetOk("application_detection_engine"); ok {
		t, err := expandWirelessControllerVapApplicationDetectionEngine(d, v, "application_detection_engine", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-detection-engine"] = t
		}
	}

	if v, ok := d.GetOk("application_dscp_marking"); ok {
		t, err := expandWirelessControllerVapApplicationDscpMarking(d, v, "application_dscp_marking", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-dscp-marking"] = t
		}
	}

	if v, ok := d.GetOk("application_report_intv"); ok {
		t, err := expandWirelessControllerVapApplicationReportIntv(d, v, "application_report_intv", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-report-intv"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming"); ok {
		t, err := expandWirelessControllerVapL3Roaming(d, v, "l3_roaming", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming"] = t
		}
	}

	if v, ok := d.GetOk("l3_roaming_mode"); ok {
		t, err := expandWirelessControllerVapL3RoamingMode(d, v, "l3_roaming_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l3-roaming-mode"] = t
		}
	}

	return &obj, nil
}
