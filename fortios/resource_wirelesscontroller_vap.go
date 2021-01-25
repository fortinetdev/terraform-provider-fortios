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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Required:     true,
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
				Computed: true,
			},
			"max_clients_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"okc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voice_enterprise": &schema.Schema{
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
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"external_logout": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
			},
			"radius_mac_auth_usergroups": &schema.Schema{
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
			"radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
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
			"portal_message_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"portal_message_overrides": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_reject_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_login_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_login_failed_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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
			"security_exempt_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"security_redirect_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"intra_vap_privacy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ldpc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mpsk_concurrent_clients": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
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
							Computed:     true,
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
							Computed:     true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4094),
				Optional:     true,
				Computed:     true,
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
			"captive_portal_radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
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
				Computed:     true,
			},
			"captive_portal_session_timeout_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 864000),
				Optional:     true,
				Computed:     true,
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
			"broadcast_suppression": &schema.Schema{
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
							Computed:     true,
						},
					},
				},
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
				ValidateFunc: validation.IntBetween(1800, 864000),
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
				ValidateFunc: validation.IntBetween(1800, 864000),
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
			"qos_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hotspot20_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"rates_11a": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rates_11bg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rates_11n_ss12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rates_11n_ss34": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rates_11ac_ss12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rates_11ac_ss34": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"utm_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
			"dynamic_sort_subtable": &schema.Schema{
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

	obj, err := getObjectWirelessControllerVap(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerVap resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerVap(obj)

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

	obj, err := getObjectWirelessControllerVap(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerVap resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerVap(obj, mkey)
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

	err := c.DeleteWirelessControllerVap(mkey)
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

	o, err := c.ReadWirelessControllerVap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerVap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerVap resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerVapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalFastRoaming(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMeshBackhaul(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAtfWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMaxClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMaxClientsAp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBroadcastSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityObsoleteOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPmf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPmfAssocComebackTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPmfSaQueryRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapOkc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVoiceEnterprise(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFastBssTransition(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFtMobilityDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFtR0KeyLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapFtOverDs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEapolKeyRetries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapTkipCounterMeasure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuthServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusMacAuthUsergroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapRadiusMacAuthUsergroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapRadiusMacAuthUsergroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEncrypt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapKeyindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandalone(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalStandaloneNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpLeaseTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalBridging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLocalAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapUsergroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapUsergroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapUsergroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverrides(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := i["auth-disclaimer-page"]; ok {
		result["auth_disclaimer_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(i["auth-disclaimer-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := i["auth-reject-page"]; ok {
		result["auth_reject_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage(i["auth-reject-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := i["auth-login-page"]; ok {
		result["auth_login_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage(i["auth-login-page"], d, pre_append)
	}

	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := i["auth-login-failed-page"]; ok {
		result["auth_login_failed_page"] = flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(i["auth-login-failed-page"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthRejectPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPortalType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerVapSelectedUsergroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerVapSelectedUsergroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapIntraVapPrivacy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapLdpc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := i["key-name"]; ok {
			tmp["key_name"] = flattenWirelessControllerVapMpskKeyKeyName(i["key-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := i["passphrase"]; ok {
			tmp["passphrase"] = flattenWirelessControllerVapMpskKeyPassphrase(i["passphrase"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["passphrase"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := i["concurrent-clients"]; ok {
			tmp["concurrent_clients"] = flattenWirelessControllerVapMpskKeyConcurrentClients(i["concurrent-clients"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			tmp["comment"] = flattenWirelessControllerVapMpskKeyComment(i["comment"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "key_name", d)
	return result
}

func flattenWirelessControllerVapMpskKeyKeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyPassphrase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyConcurrentClients(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMpskKeyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapSplitTunneling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanAuto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDynamicVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalMacauthRadiusServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalMacauthRadiusSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapCaptivePortalSessionTimeoutInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMulticastRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMulticastEnhance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapBroadcastSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMeDisableThresh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapProbeRespSuppression(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapProbeRespThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadioSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadio5GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRadio2GThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanPooling(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanPool(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerVapVlanPoolId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_group"
		if _, ok := i["wtp-group"]; ok {
			tmp["wtp_group"] = flattenWirelessControllerVapVlanPoolWtpGroup(i["wtp-group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerVapVlanPoolId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapVlanPoolWtpGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82Insertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82CircuitIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapDhcpOption82RemoteIdInsertion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapPtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGtkRekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapGtkRekeyIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEapReauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapEapReauthIntv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapQosProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapHotspot20Profile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11A(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11Bg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11NSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11NSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AcSs12(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapRates11AcSs34(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapUtmProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterPolicyOther(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerVapMacFilterListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {
			tmp["mac"] = flattenWirelessControllerVapMacFilterListMac(i["mac"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := i["mac-filter-policy"]; ok {
			tmp["mac_filter_policy"] = flattenWirelessControllerVapMacFilterListMacFilterPolicy(i["mac-filter-policy"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerVapMacFilterListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterListMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerVapMacFilterListMacFilterPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerVap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerVapName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fast_roaming", flattenWirelessControllerVapFastRoaming(o["fast-roaming"], d, "fast_roaming")); err != nil {
		if !fortiAPIPatch(o["fast-roaming"]) {
			return fmt.Errorf("Error reading fast_roaming: %v", err)
		}
	}

	if err = d.Set("external_fast_roaming", flattenWirelessControllerVapExternalFastRoaming(o["external-fast-roaming"], d, "external_fast_roaming")); err != nil {
		if !fortiAPIPatch(o["external-fast-roaming"]) {
			return fmt.Errorf("Error reading external_fast_roaming: %v", err)
		}
	}

	if err = d.Set("mesh_backhaul", flattenWirelessControllerVapMeshBackhaul(o["mesh-backhaul"], d, "mesh_backhaul")); err != nil {
		if !fortiAPIPatch(o["mesh-backhaul"]) {
			return fmt.Errorf("Error reading mesh_backhaul: %v", err)
		}
	}

	if err = d.Set("atf_weight", flattenWirelessControllerVapAtfWeight(o["atf-weight"], d, "atf_weight")); err != nil {
		if !fortiAPIPatch(o["atf-weight"]) {
			return fmt.Errorf("Error reading atf_weight: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerVapMaxClients(o["max-clients"], d, "max_clients")); err != nil {
		if !fortiAPIPatch(o["max-clients"]) {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("max_clients_ap", flattenWirelessControllerVapMaxClientsAp(o["max-clients-ap"], d, "max_clients_ap")); err != nil {
		if !fortiAPIPatch(o["max-clients-ap"]) {
			return fmt.Errorf("Error reading max_clients_ap: %v", err)
		}
	}

	if err = d.Set("ssid", flattenWirelessControllerVapSsid(o["ssid"], d, "ssid")); err != nil {
		if !fortiAPIPatch(o["ssid"]) {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenWirelessControllerVapBroadcastSsid(o["broadcast-ssid"], d, "broadcast_ssid")); err != nil {
		if !fortiAPIPatch(o["broadcast-ssid"]) {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("security_obsolete_option", flattenWirelessControllerVapSecurityObsoleteOption(o["security-obsolete-option"], d, "security_obsolete_option")); err != nil {
		if !fortiAPIPatch(o["security-obsolete-option"]) {
			return fmt.Errorf("Error reading security_obsolete_option: %v", err)
		}
	}

	if err = d.Set("security", flattenWirelessControllerVapSecurity(o["security"], d, "security")); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("pmf", flattenWirelessControllerVapPmf(o["pmf"], d, "pmf")); err != nil {
		if !fortiAPIPatch(o["pmf"]) {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("pmf_assoc_comeback_timeout", flattenWirelessControllerVapPmfAssocComebackTimeout(o["pmf-assoc-comeback-timeout"], d, "pmf_assoc_comeback_timeout")); err != nil {
		if !fortiAPIPatch(o["pmf-assoc-comeback-timeout"]) {
			return fmt.Errorf("Error reading pmf_assoc_comeback_timeout: %v", err)
		}
	}

	if err = d.Set("pmf_sa_query_retry_timeout", flattenWirelessControllerVapPmfSaQueryRetryTimeout(o["pmf-sa-query-retry-timeout"], d, "pmf_sa_query_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["pmf-sa-query-retry-timeout"]) {
			return fmt.Errorf("Error reading pmf_sa_query_retry_timeout: %v", err)
		}
	}

	if err = d.Set("okc", flattenWirelessControllerVapOkc(o["okc"], d, "okc")); err != nil {
		if !fortiAPIPatch(o["okc"]) {
			return fmt.Errorf("Error reading okc: %v", err)
		}
	}

	if err = d.Set("voice_enterprise", flattenWirelessControllerVapVoiceEnterprise(o["voice-enterprise"], d, "voice_enterprise")); err != nil {
		if !fortiAPIPatch(o["voice-enterprise"]) {
			return fmt.Errorf("Error reading voice_enterprise: %v", err)
		}
	}

	if err = d.Set("fast_bss_transition", flattenWirelessControllerVapFastBssTransition(o["fast-bss-transition"], d, "fast_bss_transition")); err != nil {
		if !fortiAPIPatch(o["fast-bss-transition"]) {
			return fmt.Errorf("Error reading fast_bss_transition: %v", err)
		}
	}

	if err = d.Set("ft_mobility_domain", flattenWirelessControllerVapFtMobilityDomain(o["ft-mobility-domain"], d, "ft_mobility_domain")); err != nil {
		if !fortiAPIPatch(o["ft-mobility-domain"]) {
			return fmt.Errorf("Error reading ft_mobility_domain: %v", err)
		}
	}

	if err = d.Set("ft_r0_key_lifetime", flattenWirelessControllerVapFtR0KeyLifetime(o["ft-r0-key-lifetime"], d, "ft_r0_key_lifetime")); err != nil {
		if !fortiAPIPatch(o["ft-r0-key-lifetime"]) {
			return fmt.Errorf("Error reading ft_r0_key_lifetime: %v", err)
		}
	}

	if err = d.Set("ft_over_ds", flattenWirelessControllerVapFtOverDs(o["ft-over-ds"], d, "ft_over_ds")); err != nil {
		if !fortiAPIPatch(o["ft-over-ds"]) {
			return fmt.Errorf("Error reading ft_over_ds: %v", err)
		}
	}

	if err = d.Set("eapol_key_retries", flattenWirelessControllerVapEapolKeyRetries(o["eapol-key-retries"], d, "eapol_key_retries")); err != nil {
		if !fortiAPIPatch(o["eapol-key-retries"]) {
			return fmt.Errorf("Error reading eapol_key_retries: %v", err)
		}
	}

	if err = d.Set("tkip_counter_measure", flattenWirelessControllerVapTkipCounterMeasure(o["tkip-counter-measure"], d, "tkip_counter_measure")); err != nil {
		if !fortiAPIPatch(o["tkip-counter-measure"]) {
			return fmt.Errorf("Error reading tkip_counter_measure: %v", err)
		}
	}

	if err = d.Set("external_web", flattenWirelessControllerVapExternalWeb(o["external-web"], d, "external_web")); err != nil {
		if !fortiAPIPatch(o["external-web"]) {
			return fmt.Errorf("Error reading external_web: %v", err)
		}
	}

	if err = d.Set("external_logout", flattenWirelessControllerVapExternalLogout(o["external-logout"], d, "external_logout")); err != nil {
		if !fortiAPIPatch(o["external-logout"]) {
			return fmt.Errorf("Error reading external_logout: %v", err)
		}
	}

	if err = d.Set("mac_auth_bypass", flattenWirelessControllerVapMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if !fortiAPIPatch(o["mac-auth-bypass"]) {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth", flattenWirelessControllerVapRadiusMacAuth(o["radius-mac-auth"], d, "radius_mac_auth")); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth"]) {
			return fmt.Errorf("Error reading radius_mac_auth: %v", err)
		}
	}

	if err = d.Set("radius_mac_auth_server", flattenWirelessControllerVapRadiusMacAuthServer(o["radius-mac-auth-server"], d, "radius_mac_auth_server")); err != nil {
		if !fortiAPIPatch(o["radius-mac-auth-server"]) {
			return fmt.Errorf("Error reading radius_mac_auth_server: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radius_mac_auth_usergroups", flattenWirelessControllerVapRadiusMacAuthUsergroups(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups")); err != nil {
			if !fortiAPIPatch(o["radius-mac-auth-usergroups"]) {
				return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radius_mac_auth_usergroups"); ok {
			if err = d.Set("radius_mac_auth_usergroups", flattenWirelessControllerVapRadiusMacAuthUsergroups(o["radius-mac-auth-usergroups"], d, "radius_mac_auth_usergroups")); err != nil {
				if !fortiAPIPatch(o["radius-mac-auth-usergroups"]) {
					return fmt.Errorf("Error reading radius_mac_auth_usergroups: %v", err)
				}
			}
		}
	}

	if err = d.Set("auth", flattenWirelessControllerVapAuth(o["auth"], d, "auth")); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("encrypt", flattenWirelessControllerVapEncrypt(o["encrypt"], d, "encrypt")); err != nil {
		if !fortiAPIPatch(o["encrypt"]) {
			return fmt.Errorf("Error reading encrypt: %v", err)
		}
	}

	if err = d.Set("keyindex", flattenWirelessControllerVapKeyindex(o["keyindex"], d, "keyindex")); err != nil {
		if !fortiAPIPatch(o["keyindex"]) {
			return fmt.Errorf("Error reading keyindex: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenWirelessControllerVapRadiusServer(o["radius-server"], d, "radius_server")); err != nil {
		if !fortiAPIPatch(o["radius-server"]) {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenWirelessControllerVapAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval")); err != nil {
		if !fortiAPIPatch(o["acct-interim-interval"]) {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("local_standalone", flattenWirelessControllerVapLocalStandalone(o["local-standalone"], d, "local_standalone")); err != nil {
		if !fortiAPIPatch(o["local-standalone"]) {
			return fmt.Errorf("Error reading local_standalone: %v", err)
		}
	}

	if err = d.Set("local_standalone_nat", flattenWirelessControllerVapLocalStandaloneNat(o["local-standalone-nat"], d, "local_standalone_nat")); err != nil {
		if !fortiAPIPatch(o["local-standalone-nat"]) {
			return fmt.Errorf("Error reading local_standalone_nat: %v", err)
		}
	}

	if err = d.Set("ip", flattenWirelessControllerVapIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("dhcp_lease_time", flattenWirelessControllerVapDhcpLeaseTime(o["dhcp-lease-time"], d, "dhcp_lease_time")); err != nil {
		if !fortiAPIPatch(o["dhcp-lease-time"]) {
			return fmt.Errorf("Error reading dhcp_lease_time: %v", err)
		}
	}

	if err = d.Set("local_bridging", flattenWirelessControllerVapLocalBridging(o["local-bridging"], d, "local_bridging")); err != nil {
		if !fortiAPIPatch(o["local-bridging"]) {
			return fmt.Errorf("Error reading local_bridging: %v", err)
		}
	}

	if err = d.Set("local_lan", flattenWirelessControllerVapLocalLan(o["local-lan"], d, "local_lan")); err != nil {
		if !fortiAPIPatch(o["local-lan"]) {
			return fmt.Errorf("Error reading local_lan: %v", err)
		}
	}

	if err = d.Set("local_authentication", flattenWirelessControllerVapLocalAuthentication(o["local-authentication"], d, "local_authentication")); err != nil {
		if !fortiAPIPatch(o["local-authentication"]) {
			return fmt.Errorf("Error reading local_authentication: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("usergroup", flattenWirelessControllerVapUsergroup(o["usergroup"], d, "usergroup")); err != nil {
			if !fortiAPIPatch(o["usergroup"]) {
				return fmt.Errorf("Error reading usergroup: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("usergroup"); ok {
			if err = d.Set("usergroup", flattenWirelessControllerVapUsergroup(o["usergroup"], d, "usergroup")); err != nil {
				if !fortiAPIPatch(o["usergroup"]) {
					return fmt.Errorf("Error reading usergroup: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_message_override_group", flattenWirelessControllerVapPortalMessageOverrideGroup(o["portal-message-override-group"], d, "portal_message_override_group")); err != nil {
		if !fortiAPIPatch(o["portal-message-override-group"]) {
			return fmt.Errorf("Error reading portal_message_override_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides")); err != nil {
			if !fortiAPIPatch(o["portal-message-overrides"]) {
				return fmt.Errorf("Error reading portal_message_overrides: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("portal_message_overrides"); ok {
			if err = d.Set("portal_message_overrides", flattenWirelessControllerVapPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides")); err != nil {
				if !fortiAPIPatch(o["portal-message-overrides"]) {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_type", flattenWirelessControllerVapPortalType(o["portal-type"], d, "portal_type")); err != nil {
		if !fortiAPIPatch(o["portal-type"]) {
			return fmt.Errorf("Error reading portal_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("selected_usergroups", flattenWirelessControllerVapSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups")); err != nil {
			if !fortiAPIPatch(o["selected-usergroups"]) {
				return fmt.Errorf("Error reading selected_usergroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("selected_usergroups"); ok {
			if err = d.Set("selected_usergroups", flattenWirelessControllerVapSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups")); err != nil {
				if !fortiAPIPatch(o["selected-usergroups"]) {
					return fmt.Errorf("Error reading selected_usergroups: %v", err)
				}
			}
		}
	}

	if err = d.Set("security_exempt_list", flattenWirelessControllerVapSecurityExemptList(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if !fortiAPIPatch(o["security-exempt-list"]) {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenWirelessControllerVapSecurityRedirectUrl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if !fortiAPIPatch(o["security-redirect-url"]) {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("intra_vap_privacy", flattenWirelessControllerVapIntraVapPrivacy(o["intra-vap-privacy"], d, "intra_vap_privacy")); err != nil {
		if !fortiAPIPatch(o["intra-vap-privacy"]) {
			return fmt.Errorf("Error reading intra_vap_privacy: %v", err)
		}
	}

	if err = d.Set("schedule", flattenWirelessControllerVapSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("ldpc", flattenWirelessControllerVapLdpc(o["ldpc"], d, "ldpc")); err != nil {
		if !fortiAPIPatch(o["ldpc"]) {
			return fmt.Errorf("Error reading ldpc: %v", err)
		}
	}

	if err = d.Set("mpsk", flattenWirelessControllerVapMpsk(o["mpsk"], d, "mpsk")); err != nil {
		if !fortiAPIPatch(o["mpsk"]) {
			return fmt.Errorf("Error reading mpsk: %v", err)
		}
	}

	if err = d.Set("mpsk_concurrent_clients", flattenWirelessControllerVapMpskConcurrentClients(o["mpsk-concurrent-clients"], d, "mpsk_concurrent_clients")); err != nil {
		if !fortiAPIPatch(o["mpsk-concurrent-clients"]) {
			return fmt.Errorf("Error reading mpsk_concurrent_clients: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key")); err != nil {
			if !fortiAPIPatch(o["mpsk-key"]) {
				return fmt.Errorf("Error reading mpsk_key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mpsk_key"); ok {
			if err = d.Set("mpsk_key", flattenWirelessControllerVapMpskKey(o["mpsk-key"], d, "mpsk_key")); err != nil {
				if !fortiAPIPatch(o["mpsk-key"]) {
					return fmt.Errorf("Error reading mpsk_key: %v", err)
				}
			}
		}
	}

	if err = d.Set("split_tunneling", flattenWirelessControllerVapSplitTunneling(o["split-tunneling"], d, "split_tunneling")); err != nil {
		if !fortiAPIPatch(o["split-tunneling"]) {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenWirelessControllerVapVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("vlan_auto", flattenWirelessControllerVapVlanAuto(o["vlan-auto"], d, "vlan_auto")); err != nil {
		if !fortiAPIPatch(o["vlan-auto"]) {
			return fmt.Errorf("Error reading vlan_auto: %v", err)
		}
	}

	if err = d.Set("dynamic_vlan", flattenWirelessControllerVapDynamicVlan(o["dynamic-vlan"], d, "dynamic_vlan")); err != nil {
		if !fortiAPIPatch(o["dynamic-vlan"]) {
			return fmt.Errorf("Error reading dynamic_vlan: %v", err)
		}
	}

	if err = d.Set("captive_portal_radius_server", flattenWirelessControllerVapCaptivePortalRadiusServer(o["captive-portal-radius-server"], d, "captive_portal_radius_server")); err != nil {
		if !fortiAPIPatch(o["captive-portal-radius-server"]) {
			return fmt.Errorf("Error reading captive_portal_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_macauth_radius_server", flattenWirelessControllerVapCaptivePortalMacauthRadiusServer(o["captive-portal-macauth-radius-server"], d, "captive_portal_macauth_radius_server")); err != nil {
		if !fortiAPIPatch(o["captive-portal-macauth-radius-server"]) {
			return fmt.Errorf("Error reading captive_portal_macauth_radius_server: %v", err)
		}
	}

	if err = d.Set("captive_portal_ac_name", flattenWirelessControllerVapCaptivePortalAcName(o["captive-portal-ac-name"], d, "captive_portal_ac_name")); err != nil {
		if !fortiAPIPatch(o["captive-portal-ac-name"]) {
			return fmt.Errorf("Error reading captive_portal_ac_name: %v", err)
		}
	}

	if err = d.Set("captive_portal_session_timeout_interval", flattenWirelessControllerVapCaptivePortalSessionTimeoutInterval(o["captive-portal-session-timeout-interval"], d, "captive_portal_session_timeout_interval")); err != nil {
		if !fortiAPIPatch(o["captive-portal-session-timeout-interval"]) {
			return fmt.Errorf("Error reading captive_portal_session_timeout_interval: %v", err)
		}
	}

	if err = d.Set("alias", flattenWirelessControllerVapAlias(o["alias"], d, "alias")); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("multicast_rate", flattenWirelessControllerVapMulticastRate(o["multicast-rate"], d, "multicast_rate")); err != nil {
		if !fortiAPIPatch(o["multicast-rate"]) {
			return fmt.Errorf("Error reading multicast_rate: %v", err)
		}
	}

	if err = d.Set("multicast_enhance", flattenWirelessControllerVapMulticastEnhance(o["multicast-enhance"], d, "multicast_enhance")); err != nil {
		if !fortiAPIPatch(o["multicast-enhance"]) {
			return fmt.Errorf("Error reading multicast_enhance: %v", err)
		}
	}

	if err = d.Set("broadcast_suppression", flattenWirelessControllerVapBroadcastSuppression(o["broadcast-suppression"], d, "broadcast_suppression")); err != nil {
		if !fortiAPIPatch(o["broadcast-suppression"]) {
			return fmt.Errorf("Error reading broadcast_suppression: %v", err)
		}
	}

	if err = d.Set("me_disable_thresh", flattenWirelessControllerVapMeDisableThresh(o["me-disable-thresh"], d, "me_disable_thresh")); err != nil {
		if !fortiAPIPatch(o["me-disable-thresh"]) {
			return fmt.Errorf("Error reading me_disable_thresh: %v", err)
		}
	}

	if err = d.Set("probe_resp_suppression", flattenWirelessControllerVapProbeRespSuppression(o["probe-resp-suppression"], d, "probe_resp_suppression")); err != nil {
		if !fortiAPIPatch(o["probe-resp-suppression"]) {
			return fmt.Errorf("Error reading probe_resp_suppression: %v", err)
		}
	}

	if err = d.Set("probe_resp_threshold", flattenWirelessControllerVapProbeRespThreshold(o["probe-resp-threshold"], d, "probe_resp_threshold")); err != nil {
		if !fortiAPIPatch(o["probe-resp-threshold"]) {
			return fmt.Errorf("Error reading probe_resp_threshold: %v", err)
		}
	}

	if err = d.Set("radio_sensitivity", flattenWirelessControllerVapRadioSensitivity(o["radio-sensitivity"], d, "radio_sensitivity")); err != nil {
		if !fortiAPIPatch(o["radio-sensitivity"]) {
			return fmt.Errorf("Error reading radio_sensitivity: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenWirelessControllerVapQuarantine(o["quarantine"], d, "quarantine")); err != nil {
		if !fortiAPIPatch(o["quarantine"]) {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("radio_5g_threshold", flattenWirelessControllerVapRadio5GThreshold(o["radio-5g-threshold"], d, "radio_5g_threshold")); err != nil {
		if !fortiAPIPatch(o["radio-5g-threshold"]) {
			return fmt.Errorf("Error reading radio_5g_threshold: %v", err)
		}
	}

	if err = d.Set("radio_2g_threshold", flattenWirelessControllerVapRadio2GThreshold(o["radio-2g-threshold"], d, "radio_2g_threshold")); err != nil {
		if !fortiAPIPatch(o["radio-2g-threshold"]) {
			return fmt.Errorf("Error reading radio_2g_threshold: %v", err)
		}
	}

	if err = d.Set("vlan_pooling", flattenWirelessControllerVapVlanPooling(o["vlan-pooling"], d, "vlan_pooling")); err != nil {
		if !fortiAPIPatch(o["vlan-pooling"]) {
			return fmt.Errorf("Error reading vlan_pooling: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool")); err != nil {
			if !fortiAPIPatch(o["vlan-pool"]) {
				return fmt.Errorf("Error reading vlan_pool: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vlan_pool"); ok {
			if err = d.Set("vlan_pool", flattenWirelessControllerVapVlanPool(o["vlan-pool"], d, "vlan_pool")); err != nil {
				if !fortiAPIPatch(o["vlan-pool"]) {
					return fmt.Errorf("Error reading vlan_pool: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp_option82_insertion", flattenWirelessControllerVapDhcpOption82Insertion(o["dhcp-option82-insertion"], d, "dhcp_option82_insertion")); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option82_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id_insertion", flattenWirelessControllerVapDhcpOption82CircuitIdInsertion(o["dhcp-option82-circuit-id-insertion"], d, "dhcp_option82_circuit_id_insertion")); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-circuit-id-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id_insertion: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id_insertion", flattenWirelessControllerVapDhcpOption82RemoteIdInsertion(o["dhcp-option82-remote-id-insertion"], d, "dhcp_option82_remote_id_insertion")); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-remote-id-insertion"]) {
			return fmt.Errorf("Error reading dhcp_option82_remote_id_insertion: %v", err)
		}
	}

	if err = d.Set("ptk_rekey", flattenWirelessControllerVapPtkRekey(o["ptk-rekey"], d, "ptk_rekey")); err != nil {
		if !fortiAPIPatch(o["ptk-rekey"]) {
			return fmt.Errorf("Error reading ptk_rekey: %v", err)
		}
	}

	if err = d.Set("ptk_rekey_intv", flattenWirelessControllerVapPtkRekeyIntv(o["ptk-rekey-intv"], d, "ptk_rekey_intv")); err != nil {
		if !fortiAPIPatch(o["ptk-rekey-intv"]) {
			return fmt.Errorf("Error reading ptk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("gtk_rekey", flattenWirelessControllerVapGtkRekey(o["gtk-rekey"], d, "gtk_rekey")); err != nil {
		if !fortiAPIPatch(o["gtk-rekey"]) {
			return fmt.Errorf("Error reading gtk_rekey: %v", err)
		}
	}

	if err = d.Set("gtk_rekey_intv", flattenWirelessControllerVapGtkRekeyIntv(o["gtk-rekey-intv"], d, "gtk_rekey_intv")); err != nil {
		if !fortiAPIPatch(o["gtk-rekey-intv"]) {
			return fmt.Errorf("Error reading gtk_rekey_intv: %v", err)
		}
	}

	if err = d.Set("eap_reauth", flattenWirelessControllerVapEapReauth(o["eap-reauth"], d, "eap_reauth")); err != nil {
		if !fortiAPIPatch(o["eap-reauth"]) {
			return fmt.Errorf("Error reading eap_reauth: %v", err)
		}
	}

	if err = d.Set("eap_reauth_intv", flattenWirelessControllerVapEapReauthIntv(o["eap-reauth-intv"], d, "eap_reauth_intv")); err != nil {
		if !fortiAPIPatch(o["eap-reauth-intv"]) {
			return fmt.Errorf("Error reading eap_reauth_intv: %v", err)
		}
	}

	if err = d.Set("qos_profile", flattenWirelessControllerVapQosProfile(o["qos-profile"], d, "qos_profile")); err != nil {
		if !fortiAPIPatch(o["qos-profile"]) {
			return fmt.Errorf("Error reading qos_profile: %v", err)
		}
	}

	if err = d.Set("hotspot20_profile", flattenWirelessControllerVapHotspot20Profile(o["hotspot20-profile"], d, "hotspot20_profile")); err != nil {
		if !fortiAPIPatch(o["hotspot20-profile"]) {
			return fmt.Errorf("Error reading hotspot20_profile: %v", err)
		}
	}

	if err = d.Set("rates_11a", flattenWirelessControllerVapRates11A(o["rates-11a"], d, "rates_11a")); err != nil {
		if !fortiAPIPatch(o["rates-11a"]) {
			return fmt.Errorf("Error reading rates_11a: %v", err)
		}
	}

	if err = d.Set("rates_11bg", flattenWirelessControllerVapRates11Bg(o["rates-11bg"], d, "rates_11bg")); err != nil {
		if !fortiAPIPatch(o["rates-11bg"]) {
			return fmt.Errorf("Error reading rates_11bg: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss12", flattenWirelessControllerVapRates11NSs12(o["rates-11n-ss12"], d, "rates_11n_ss12")); err != nil {
		if !fortiAPIPatch(o["rates-11n-ss12"]) {
			return fmt.Errorf("Error reading rates_11n_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11n_ss34", flattenWirelessControllerVapRates11NSs34(o["rates-11n-ss34"], d, "rates_11n_ss34")); err != nil {
		if !fortiAPIPatch(o["rates-11n-ss34"]) {
			return fmt.Errorf("Error reading rates_11n_ss34: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss12", flattenWirelessControllerVapRates11AcSs12(o["rates-11ac-ss12"], d, "rates_11ac_ss12")); err != nil {
		if !fortiAPIPatch(o["rates-11ac-ss12"]) {
			return fmt.Errorf("Error reading rates_11ac_ss12: %v", err)
		}
	}

	if err = d.Set("rates_11ac_ss34", flattenWirelessControllerVapRates11AcSs34(o["rates-11ac-ss34"], d, "rates_11ac_ss34")); err != nil {
		if !fortiAPIPatch(o["rates-11ac-ss34"]) {
			return fmt.Errorf("Error reading rates_11ac_ss34: %v", err)
		}
	}

	if err = d.Set("utm_profile", flattenWirelessControllerVapUtmProfile(o["utm-profile"], d, "utm_profile")); err != nil {
		if !fortiAPIPatch(o["utm-profile"]) {
			return fmt.Errorf("Error reading utm_profile: %v", err)
		}
	}

	if err = d.Set("mac_filter", flattenWirelessControllerVapMacFilter(o["mac-filter"], d, "mac_filter")); err != nil {
		if !fortiAPIPatch(o["mac-filter"]) {
			return fmt.Errorf("Error reading mac_filter: %v", err)
		}
	}

	if err = d.Set("mac_filter_policy_other", flattenWirelessControllerVapMacFilterPolicyOther(o["mac-filter-policy-other"], d, "mac_filter_policy_other")); err != nil {
		if !fortiAPIPatch(o["mac-filter-policy-other"]) {
			return fmt.Errorf("Error reading mac_filter_policy_other: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list")); err != nil {
			if !fortiAPIPatch(o["mac-filter-list"]) {
				return fmt.Errorf("Error reading mac_filter_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_filter_list"); ok {
			if err = d.Set("mac_filter_list", flattenWirelessControllerVapMacFilterList(o["mac-filter-list"], d, "mac_filter_list")); err != nil {
				if !fortiAPIPatch(o["mac-filter-list"]) {
					return fmt.Errorf("Error reading mac_filter_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerVapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerVapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalFastRoaming(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMeshBackhaul(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAtfWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMaxClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMaxClientsAp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBroadcastSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityObsoleteOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmfAssocComebackTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPmfSaQueryRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapOkc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVoiceEnterprise(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFastBssTransition(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtMobilityDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtR0KeyLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapFtOverDs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapolKeyRetries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapTkipCounterMeasure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusMacAuthUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerVapRadiusMacAuthUsergroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapRadiusMacAuthUsergroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEncrypt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapKeyindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandalone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalStandaloneNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpLeaseTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalBridging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLocalAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUsergroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerVapUsergroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapUsergroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverrides(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-disclaimer-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d, i["auth_disclaimer_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-reject-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthRejectPage(d, i["auth_reject_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-login-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthLoginPage(d, i["auth_login_page"], pre_append)
	}
	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := d.GetOk(pre_append); ok {
		result["auth-login-failed-page"], _ = expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d, i["auth_login_failed_page"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthDisclaimerPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthRejectPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalMessageOverridesAuthLoginFailedPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPortalType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerVapSelectedUsergroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapSelectedUsergroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapIntraVapPrivacy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapLdpc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-name"], _ = expandWirelessControllerVapMpskKeyKeyName(d, i["key_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passphrase"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passphrase"], _ = expandWirelessControllerVapMpskKeyPassphrase(d, i["passphrase"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "concurrent_clients"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["concurrent-clients"], _ = expandWirelessControllerVapMpskKeyConcurrentClients(d, i["concurrent_clients"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWirelessControllerVapMpskKeyComment(d, i["comment"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMpskKeyKeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyPassphrase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyConcurrentClients(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMpskKeyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapSplitTunneling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanAuto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDynamicVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalMacauthRadiusServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalMacauthRadiusSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapCaptivePortalSessionTimeoutInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMulticastRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMulticastEnhance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapBroadcastSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMeDisableThresh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapProbeRespSuppression(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapProbeRespThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadioSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadio5GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRadio2GThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPooling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerVapVlanPoolId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wtp-group"], _ = expandWirelessControllerVapVlanPoolWtpGroup(d, i["wtp_group"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapVlanPoolId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapVlanPoolWtpGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82Insertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82CircuitIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapDhcpOption82RemoteIdInsertion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapPtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGtkRekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapGtkRekeyIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapReauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapEapReauthIntv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapQosProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapHotspot20Profile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11A(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11Bg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11NSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11NSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AcSs12(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapRates11AcSs34(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapUtmProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterPolicyOther(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerVapMacFilterListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandWirelessControllerVapMacFilterListMac(d, i["mac"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_filter_policy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac-filter-policy"], _ = expandWirelessControllerVapMacFilterListMacFilterPolicy(d, i["mac_filter_policy"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerVapMacFilterListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterListMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerVapMacFilterListMacFilterPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerVap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerVapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fast_roaming"); ok {
		t, err := expandWirelessControllerVapFastRoaming(d, v, "fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("external_fast_roaming"); ok {
		t, err := expandWirelessControllerVapExternalFastRoaming(d, v, "external_fast_roaming")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-fast-roaming"] = t
		}
	}

	if v, ok := d.GetOk("mesh_backhaul"); ok {
		t, err := expandWirelessControllerVapMeshBackhaul(d, v, "mesh_backhaul")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-backhaul"] = t
		}
	}

	if v, ok := d.GetOkExists("atf_weight"); ok {
		t, err := expandWirelessControllerVapAtfWeight(d, v, "atf_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["atf-weight"] = t
		}
	}

	if v, ok := d.GetOkExists("max_clients"); ok {
		t, err := expandWirelessControllerVapMaxClients(d, v, "max_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOkExists("max_clients_ap"); ok {
		t, err := expandWirelessControllerVapMaxClientsAp(d, v, "max_clients_ap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients-ap"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok {
		t, err := expandWirelessControllerVapSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok {
		t, err := expandWirelessControllerVapBroadcastSsid(d, v, "broadcast_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("security_obsolete_option"); ok {
		t, err := expandWirelessControllerVapSecurityObsoleteOption(d, v, "security_obsolete_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-obsolete-option"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok {
		t, err := expandWirelessControllerVapSecurity(d, v, "security")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok {
		t, err := expandWirelessControllerVapPmf(d, v, "pmf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("pmf_assoc_comeback_timeout"); ok {
		t, err := expandWirelessControllerVapPmfAssocComebackTimeout(d, v, "pmf_assoc_comeback_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-assoc-comeback-timeout"] = t
		}
	}

	if v, ok := d.GetOk("pmf_sa_query_retry_timeout"); ok {
		t, err := expandWirelessControllerVapPmfSaQueryRetryTimeout(d, v, "pmf_sa_query_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf-sa-query-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("okc"); ok {
		t, err := expandWirelessControllerVapOkc(d, v, "okc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["okc"] = t
		}
	}

	if v, ok := d.GetOk("voice_enterprise"); ok {
		t, err := expandWirelessControllerVapVoiceEnterprise(d, v, "voice_enterprise")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-enterprise"] = t
		}
	}

	if v, ok := d.GetOk("fast_bss_transition"); ok {
		t, err := expandWirelessControllerVapFastBssTransition(d, v, "fast_bss_transition")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-bss-transition"] = t
		}
	}

	if v, ok := d.GetOk("ft_mobility_domain"); ok {
		t, err := expandWirelessControllerVapFtMobilityDomain(d, v, "ft_mobility_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-mobility-domain"] = t
		}
	}

	if v, ok := d.GetOk("ft_r0_key_lifetime"); ok {
		t, err := expandWirelessControllerVapFtR0KeyLifetime(d, v, "ft_r0_key_lifetime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-r0-key-lifetime"] = t
		}
	}

	if v, ok := d.GetOk("ft_over_ds"); ok {
		t, err := expandWirelessControllerVapFtOverDs(d, v, "ft_over_ds")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ft-over-ds"] = t
		}
	}

	if v, ok := d.GetOk("eapol_key_retries"); ok {
		t, err := expandWirelessControllerVapEapolKeyRetries(d, v, "eapol_key_retries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eapol-key-retries"] = t
		}
	}

	if v, ok := d.GetOk("tkip_counter_measure"); ok {
		t, err := expandWirelessControllerVapTkipCounterMeasure(d, v, "tkip_counter_measure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tkip-counter-measure"] = t
		}
	}

	if v, ok := d.GetOk("external_web"); ok {
		t, err := expandWirelessControllerVapExternalWeb(d, v, "external_web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-web"] = t
		}
	}

	if v, ok := d.GetOk("external_logout"); ok {
		t, err := expandWirelessControllerVapExternalLogout(d, v, "external_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-logout"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok {
		t, err := expandWirelessControllerVapMacAuthBypass(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth"); ok {
		t, err := expandWirelessControllerVapRadiusMacAuth(d, v, "radius_mac_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_server"); ok {
		t, err := expandWirelessControllerVapRadiusMacAuthServer(d, v, "radius_mac_auth_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-server"] = t
		}
	}

	if v, ok := d.GetOk("radius_mac_auth_usergroups"); ok {
		t, err := expandWirelessControllerVapRadiusMacAuthUsergroups(d, v, "radius_mac_auth_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-mac-auth-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		t, err := expandWirelessControllerVapAuth(d, v, "auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("encrypt"); ok {
		t, err := expandWirelessControllerVapEncrypt(d, v, "encrypt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encrypt"] = t
		}
	}

	if v, ok := d.GetOk("keyindex"); ok {
		t, err := expandWirelessControllerVapKeyindex(d, v, "keyindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyindex"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandWirelessControllerVapKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok {
		t, err := expandWirelessControllerVapPassphrase(d, v, "passphrase")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok {
		t, err := expandWirelessControllerVapRadiusServer(d, v, "radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok {
		t, err := expandWirelessControllerVapAcctInterimInterval(d, v, "acct_interim_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone"); ok {
		t, err := expandWirelessControllerVapLocalStandalone(d, v, "local_standalone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone"] = t
		}
	}

	if v, ok := d.GetOk("local_standalone_nat"); ok {
		t, err := expandWirelessControllerVapLocalStandaloneNat(d, v, "local_standalone_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-standalone-nat"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWirelessControllerVapIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_lease_time"); ok {
		t, err := expandWirelessControllerVapDhcpLeaseTime(d, v, "dhcp_lease_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-lease-time"] = t
		}
	}

	if v, ok := d.GetOk("local_bridging"); ok {
		t, err := expandWirelessControllerVapLocalBridging(d, v, "local_bridging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-bridging"] = t
		}
	}

	if v, ok := d.GetOk("local_lan"); ok {
		t, err := expandWirelessControllerVapLocalLan(d, v, "local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-lan"] = t
		}
	}

	if v, ok := d.GetOk("local_authentication"); ok {
		t, err := expandWirelessControllerVapLocalAuthentication(d, v, "local_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-authentication"] = t
		}
	}

	if v, ok := d.GetOk("usergroup"); ok {
		t, err := expandWirelessControllerVapUsergroup(d, v, "usergroup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroup"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_override_group"); ok {
		t, err := expandWirelessControllerVapPortalMessageOverrideGroup(d, v, "portal_message_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-override-group"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_overrides"); ok {
		t, err := expandWirelessControllerVapPortalMessageOverrides(d, v, "portal_message_overrides")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-overrides"] = t
		}
	}

	if v, ok := d.GetOk("portal_type"); ok {
		t, err := expandWirelessControllerVapPortalType(d, v, "portal_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-type"] = t
		}
	}

	if v, ok := d.GetOk("selected_usergroups"); ok {
		t, err := expandWirelessControllerVapSelectedUsergroups(d, v, "selected_usergroups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-usergroups"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok {
		t, err := expandWirelessControllerVapSecurityExemptList(d, v, "security_exempt_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	}

	if v, ok := d.GetOk("security_redirect_url"); ok {
		t, err := expandWirelessControllerVapSecurityRedirectUrl(d, v, "security_redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("intra_vap_privacy"); ok {
		t, err := expandWirelessControllerVapIntraVapPrivacy(d, v, "intra_vap_privacy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intra-vap-privacy"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandWirelessControllerVapSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("ldpc"); ok {
		t, err := expandWirelessControllerVapLdpc(d, v, "ldpc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldpc"] = t
		}
	}

	if v, ok := d.GetOk("mpsk"); ok {
		t, err := expandWirelessControllerVapMpsk(d, v, "mpsk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk"] = t
		}
	}

	if v, ok := d.GetOkExists("mpsk_concurrent_clients"); ok {
		t, err := expandWirelessControllerVapMpskConcurrentClients(d, v, "mpsk_concurrent_clients")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-concurrent-clients"] = t
		}
	}

	if v, ok := d.GetOk("mpsk_key"); ok {
		t, err := expandWirelessControllerVapMpskKey(d, v, "mpsk_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mpsk-key"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok {
		t, err := expandWirelessControllerVapSplitTunneling(d, v, "split_tunneling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOkExists("vlanid"); ok {
		t, err := expandWirelessControllerVapVlanid(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("vlan_auto"); ok {
		t, err := expandWirelessControllerVapVlanAuto(d, v, "vlan_auto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-auto"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_vlan"); ok {
		t, err := expandWirelessControllerVapDynamicVlan(d, v, "dynamic_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-vlan"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_server"); ok {
		t, err := expandWirelessControllerVapCaptivePortalRadiusServer(d, v, "captive_portal_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_radius_secret"); ok {
		t, err := expandWirelessControllerVapCaptivePortalRadiusSecret(d, v, "captive_portal_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_server"); ok {
		t, err := expandWirelessControllerVapCaptivePortalMacauthRadiusServer(d, v, "captive_portal_macauth_radius_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-server"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_macauth_radius_secret"); ok {
		t, err := expandWirelessControllerVapCaptivePortalMacauthRadiusSecret(d, v, "captive_portal_macauth_radius_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-macauth-radius-secret"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal_ac_name"); ok {
		t, err := expandWirelessControllerVapCaptivePortalAcName(d, v, "captive_portal_ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-ac-name"] = t
		}
	}

	if v, ok := d.GetOkExists("captive_portal_session_timeout_interval"); ok {
		t, err := expandWirelessControllerVapCaptivePortalSessionTimeoutInterval(d, v, "captive_portal_session_timeout_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal-session-timeout-interval"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok {
		t, err := expandWirelessControllerVapAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("multicast_rate"); ok {
		t, err := expandWirelessControllerVapMulticastRate(d, v, "multicast_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-rate"] = t
		}
	}

	if v, ok := d.GetOk("multicast_enhance"); ok {
		t, err := expandWirelessControllerVapMulticastEnhance(d, v, "multicast_enhance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-enhance"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_suppression"); ok {
		t, err := expandWirelessControllerVapBroadcastSuppression(d, v, "broadcast_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-suppression"] = t
		}
	}

	if v, ok := d.GetOk("me_disable_thresh"); ok {
		t, err := expandWirelessControllerVapMeDisableThresh(d, v, "me_disable_thresh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["me-disable-thresh"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_suppression"); ok {
		t, err := expandWirelessControllerVapProbeRespSuppression(d, v, "probe_resp_suppression")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-suppression"] = t
		}
	}

	if v, ok := d.GetOk("probe_resp_threshold"); ok {
		t, err := expandWirelessControllerVapProbeRespThreshold(d, v, "probe_resp_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["probe-resp-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_sensitivity"); ok {
		t, err := expandWirelessControllerVapRadioSensitivity(d, v, "radio_sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok {
		t, err := expandWirelessControllerVapQuarantine(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("radio_5g_threshold"); ok {
		t, err := expandWirelessControllerVapRadio5GThreshold(d, v, "radio_5g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-5g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("radio_2g_threshold"); ok {
		t, err := expandWirelessControllerVapRadio2GThreshold(d, v, "radio_2g_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2g-threshold"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pooling"); ok {
		t, err := expandWirelessControllerVapVlanPooling(d, v, "vlan_pooling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pooling"] = t
		}
	}

	if v, ok := d.GetOk("vlan_pool"); ok {
		t, err := expandWirelessControllerVapVlanPool(d, v, "vlan_pool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-pool"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption82Insertion(d, v, "dhcp_option82_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption82CircuitIdInsertion(d, v, "dhcp_option82_circuit_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-circuit-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id_insertion"); ok {
		t, err := expandWirelessControllerVapDhcpOption82RemoteIdInsertion(d, v, "dhcp_option82_remote_id_insertion")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-option82-remote-id-insertion"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey"); ok {
		t, err := expandWirelessControllerVapPtkRekey(d, v, "ptk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("ptk_rekey_intv"); ok {
		t, err := expandWirelessControllerVapPtkRekeyIntv(d, v, "ptk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ptk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey"); ok {
		t, err := expandWirelessControllerVapGtkRekey(d, v, "gtk_rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey"] = t
		}
	}

	if v, ok := d.GetOk("gtk_rekey_intv"); ok {
		t, err := expandWirelessControllerVapGtkRekeyIntv(d, v, "gtk_rekey_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtk-rekey-intv"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth"); ok {
		t, err := expandWirelessControllerVapEapReauth(d, v, "eap_reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth"] = t
		}
	}

	if v, ok := d.GetOk("eap_reauth_intv"); ok {
		t, err := expandWirelessControllerVapEapReauthIntv(d, v, "eap_reauth_intv")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-reauth-intv"] = t
		}
	}

	if v, ok := d.GetOk("qos_profile"); ok {
		t, err := expandWirelessControllerVapQosProfile(d, v, "qos_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-profile"] = t
		}
	}

	if v, ok := d.GetOk("hotspot20_profile"); ok {
		t, err := expandWirelessControllerVapHotspot20Profile(d, v, "hotspot20_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hotspot20-profile"] = t
		}
	}

	if v, ok := d.GetOk("rates_11a"); ok {
		t, err := expandWirelessControllerVapRates11A(d, v, "rates_11a")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11a"] = t
		}
	}

	if v, ok := d.GetOk("rates_11bg"); ok {
		t, err := expandWirelessControllerVapRates11Bg(d, v, "rates_11bg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11bg"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss12"); ok {
		t, err := expandWirelessControllerVapRates11NSs12(d, v, "rates_11n_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11n_ss34"); ok {
		t, err := expandWirelessControllerVapRates11NSs34(d, v, "rates_11n_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11n-ss34"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss12"); ok {
		t, err := expandWirelessControllerVapRates11AcSs12(d, v, "rates_11ac_ss12")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss12"] = t
		}
	}

	if v, ok := d.GetOk("rates_11ac_ss34"); ok {
		t, err := expandWirelessControllerVapRates11AcSs34(d, v, "rates_11ac_ss34")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rates-11ac-ss34"] = t
		}
	}

	if v, ok := d.GetOk("utm_profile"); ok {
		t, err := expandWirelessControllerVapUtmProfile(d, v, "utm_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-profile"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter"); ok {
		t, err := expandWirelessControllerVapMacFilter(d, v, "mac_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_policy_other"); ok {
		t, err := expandWirelessControllerVapMacFilterPolicyOther(d, v, "mac_filter_policy_other")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-policy-other"] = t
		}
	}

	if v, ok := d.GetOk("mac_filter_list"); ok {
		t, err := expandWirelessControllerVapMacFilterList(d, v, "mac_filter_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-filter-list"] = t
		}
	}

	return &obj, nil
}
