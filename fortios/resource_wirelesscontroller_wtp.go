// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerWtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpCreate,
		Read:   resourceWirelessControllerWtpRead,
		Update: resourceWirelessControllerWtpUpdate,
		Delete: resourceWirelessControllerWtpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"wtp_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"region_x": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"region_y": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"wtp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"wtp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bonjour_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"override_led_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"led_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_wan_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wan_port_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_ip_fragment": &schema.Schema{
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
			"override_split_tunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"override_lan": &schema.Schema{
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
			"override_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_login_passwd_change": &schema.Schema{
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
						"override_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_txpower": &schema.Schema{
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
						"override_vaps": &schema.Schema{
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
						"override_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"override_band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"spectrum_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"override_txpower": &schema.Schema{
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
						"override_vaps": &schema.Schema{
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
						"override_channel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
					},
				},
			},
			"image_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mesh_bridge_enable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"coordinate_latitude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"coordinate_longitude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerWtpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerWtp(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtp resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWtp(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWtp")
	}

	return resourceWirelessControllerWtpRead(d, m)
}

func resourceWirelessControllerWtpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerWtp(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtp resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWtp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWtp")
	}

	return resourceWirelessControllerWtpRead(d, m)
}

func resourceWirelessControllerWtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerWtp(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerWtp(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtp resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpWtpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpAdmin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegionX(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegionY(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpBonjourProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLedState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLedState(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideWanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpWanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideIpFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpIpFragmentPreventing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpTunMtuUplink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpTunMtuDownlink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideSplitTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclLocalApSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAcl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerWtpSplitTunnelingAclId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := i["dest-ip"]; ok {
			tmp["dest_ip"] = flattenWirelessControllerWtpSplitTunnelingAclDestIp(i["dest-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpSplitTunnelingAclId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclDestIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenWirelessControllerWtpOverrideLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLan(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := i["port-mode"]; ok {
		result["port_mode"] = flattenWirelessControllerWtpLanPortMode(i["port-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port_ssid"
	if _, ok := i["port-ssid"]; ok {
		result["port_ssid"] = flattenWirelessControllerWtpLanPortSsid(i["port-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_mode"
	if _, ok := i["port1-mode"]; ok {
		result["port1_mode"] = flattenWirelessControllerWtpLanPort1Mode(i["port1-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := i["port1-ssid"]; ok {
		result["port1_ssid"] = flattenWirelessControllerWtpLanPort1Ssid(i["port1-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_mode"
	if _, ok := i["port2-mode"]; ok {
		result["port2_mode"] = flattenWirelessControllerWtpLanPort2Mode(i["port2-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := i["port2-ssid"]; ok {
		result["port2_ssid"] = flattenWirelessControllerWtpLanPort2Ssid(i["port2-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_mode"
	if _, ok := i["port3-mode"]; ok {
		result["port3_mode"] = flattenWirelessControllerWtpLanPort3Mode(i["port3-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := i["port3-ssid"]; ok {
		result["port3_ssid"] = flattenWirelessControllerWtpLanPort3Ssid(i["port3-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_mode"
	if _, ok := i["port4-mode"]; ok {
		result["port4_mode"] = flattenWirelessControllerWtpLanPort4Mode(i["port4-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := i["port4-ssid"]; ok {
		result["port4_ssid"] = flattenWirelessControllerWtpLanPort4Ssid(i["port4-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_mode"
	if _, ok := i["port5-mode"]; ok {
		result["port5_mode"] = flattenWirelessControllerWtpLanPort5Mode(i["port5-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := i["port5-ssid"]; ok {
		result["port5_ssid"] = flattenWirelessControllerWtpLanPort5Ssid(i["port5-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_mode"
	if _, ok := i["port6-mode"]; ok {
		result["port6_mode"] = flattenWirelessControllerWtpLanPort6Mode(i["port6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := i["port6-ssid"]; ok {
		result["port6_ssid"] = flattenWirelessControllerWtpLanPort6Ssid(i["port6-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_mode"
	if _, ok := i["port7-mode"]; ok {
		result["port7_mode"] = flattenWirelessControllerWtpLanPort7Mode(i["port7-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := i["port7-ssid"]; ok {
		result["port7_ssid"] = flattenWirelessControllerWtpLanPort7Ssid(i["port7-ssid"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_mode"
	if _, ok := i["port8-mode"]; ok {
		result["port8_mode"] = flattenWirelessControllerWtpLanPort8Mode(i["port8-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := i["port8-ssid"]; ok {
		result["port8_ssid"] = flattenWirelessControllerWtpLanPort8Ssid(i["port8-ssid"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpLanPortMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort1Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort2Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort2Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort3Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort3Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort4Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort4Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort5Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort5Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort6Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort7Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort7Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort8Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort8Ssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpLoginPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpRadio1RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {
		result["override_band"] = flattenWirelessControllerWtpRadio1OverrideBand(i["override-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpRadio1Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {
		result["override_analysis"] = flattenWirelessControllerWtpRadio1OverrideAnalysis(i["override-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio1SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {
		result["override_txpower"] = flattenWirelessControllerWtpRadio1OverrideTxpower(i["override-txpower"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpRadio1AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpRadio1AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpRadio1AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpRadio1PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {
		result["override_vaps"] = flattenWirelessControllerWtpRadio1OverrideVaps(i["override-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpRadio1VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpRadio1Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {
		result["override_channel"] = flattenWirelessControllerWtpRadio1OverrideChannel(i["override-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpRadio1Channel(i["channel"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio1RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerWtpRadio1VapsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpRadio1VapsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Channel(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["chan"] = flattenWirelessControllerWtpRadio1ChannelChan(i["chan"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpRadio1ChannelChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {
		result["radio_id"] = flattenWirelessControllerWtpRadio2RadioId(i["radio-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {
		result["override_band"] = flattenWirelessControllerWtpRadio2OverrideBand(i["override-band"], d, pre_append)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenWirelessControllerWtpRadio2Band(i["band"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {
		result["override_analysis"] = flattenWirelessControllerWtpRadio2OverrideAnalysis(i["override-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {
		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio2SpectrumAnalysis(i["spectrum-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {
		result["override_txpower"] = flattenWirelessControllerWtpRadio2OverrideTxpower(i["override-txpower"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {
		result["auto_power_level"] = flattenWirelessControllerWtpRadio2AutoPowerLevel(i["auto-power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {
		result["auto_power_high"] = flattenWirelessControllerWtpRadio2AutoPowerHigh(i["auto-power-high"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {
		result["auto_power_low"] = flattenWirelessControllerWtpRadio2AutoPowerLow(i["auto-power-low"], d, pre_append)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenWirelessControllerWtpRadio2PowerLevel(i["power-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {
		result["override_vaps"] = flattenWirelessControllerWtpRadio2OverrideVaps(i["override-vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {
		result["vap_all"] = flattenWirelessControllerWtpRadio2VapAll(i["vap-all"], d, pre_append)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {
		result["vaps"] = flattenWirelessControllerWtpRadio2Vaps(i["vaps"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {
		result["override_channel"] = flattenWirelessControllerWtpRadio2OverrideChannel(i["override-channel"], d, pre_append)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenWirelessControllerWtpRadio2Channel(i["channel"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio2RadioId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideBand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Band(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerLow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideVaps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2VapAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vaps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenWirelessControllerWtpRadio2VapsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpRadio2VapsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideChannel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Channel(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["chan"] = flattenWirelessControllerWtpRadio2ChannelChan(i["chan"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerWtpRadio2ChannelChan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpImageDownload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpMeshBridgeEnable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpCoordinateLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpCoordinateLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("wtp_id", flattenWirelessControllerWtpWtpId(o["wtp-id"], d, "wtp_id")); err != nil {
		if !fortiAPIPatch(o["wtp-id"]) {
			return fmt.Errorf("Error reading wtp_id: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerWtpIndex(o["index"], d, "index")); err != nil {
		if !fortiAPIPatch(o["index"]) {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("admin", flattenWirelessControllerWtpAdmin(o["admin"], d, "admin")); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerWtpName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("location", flattenWirelessControllerWtpLocation(o["location"], d, "location")); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("region", flattenWirelessControllerWtpRegion(o["region"], d, "region")); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("region_x", flattenWirelessControllerWtpRegionX(o["region-x"], d, "region_x")); err != nil {
		if !fortiAPIPatch(o["region-x"]) {
			return fmt.Errorf("Error reading region_x: %v", err)
		}
	}

	if err = d.Set("region_y", flattenWirelessControllerWtpRegionY(o["region-y"], d, "region_y")); err != nil {
		if !fortiAPIPatch(o["region-y"]) {
			return fmt.Errorf("Error reading region_y: %v", err)
		}
	}

	if err = d.Set("wtp_profile", flattenWirelessControllerWtpWtpProfile(o["wtp-profile"], d, "wtp_profile")); err != nil {
		if !fortiAPIPatch(o["wtp-profile"]) {
			return fmt.Errorf("Error reading wtp_profile: %v", err)
		}
	}

	if err = d.Set("wtp_mode", flattenWirelessControllerWtpWtpMode(o["wtp-mode"], d, "wtp_mode")); err != nil {
		if !fortiAPIPatch(o["wtp-mode"]) {
			return fmt.Errorf("Error reading wtp_mode: %v", err)
		}
	}

	if err = d.Set("bonjour_profile", flattenWirelessControllerWtpBonjourProfile(o["bonjour-profile"], d, "bonjour_profile")); err != nil {
		if !fortiAPIPatch(o["bonjour-profile"]) {
			return fmt.Errorf("Error reading bonjour_profile: %v", err)
		}
	}

	if err = d.Set("override_led_state", flattenWirelessControllerWtpOverrideLedState(o["override-led-state"], d, "override_led_state")); err != nil {
		if !fortiAPIPatch(o["override-led-state"]) {
			return fmt.Errorf("Error reading override_led_state: %v", err)
		}
	}

	if err = d.Set("led_state", flattenWirelessControllerWtpLedState(o["led-state"], d, "led_state")); err != nil {
		if !fortiAPIPatch(o["led-state"]) {
			return fmt.Errorf("Error reading led_state: %v", err)
		}
	}

	if err = d.Set("override_wan_port_mode", flattenWirelessControllerWtpOverrideWanPortMode(o["override-wan-port-mode"], d, "override_wan_port_mode")); err != nil {
		if !fortiAPIPatch(o["override-wan-port-mode"]) {
			return fmt.Errorf("Error reading override_wan_port_mode: %v", err)
		}
	}

	if err = d.Set("wan_port_mode", flattenWirelessControllerWtpWanPortMode(o["wan-port-mode"], d, "wan_port_mode")); err != nil {
		if !fortiAPIPatch(o["wan-port-mode"]) {
			return fmt.Errorf("Error reading wan_port_mode: %v", err)
		}
	}

	if err = d.Set("override_ip_fragment", flattenWirelessControllerWtpOverrideIpFragment(o["override-ip-fragment"], d, "override_ip_fragment")); err != nil {
		if !fortiAPIPatch(o["override-ip-fragment"]) {
			return fmt.Errorf("Error reading override_ip_fragment: %v", err)
		}
	}

	if err = d.Set("ip_fragment_preventing", flattenWirelessControllerWtpIpFragmentPreventing(o["ip-fragment-preventing"], d, "ip_fragment_preventing")); err != nil {
		if !fortiAPIPatch(o["ip-fragment-preventing"]) {
			return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
		}
	}

	if err = d.Set("tun_mtu_uplink", flattenWirelessControllerWtpTunMtuUplink(o["tun-mtu-uplink"], d, "tun_mtu_uplink")); err != nil {
		if !fortiAPIPatch(o["tun-mtu-uplink"]) {
			return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
		}
	}

	if err = d.Set("tun_mtu_downlink", flattenWirelessControllerWtpTunMtuDownlink(o["tun-mtu-downlink"], d, "tun_mtu_downlink")); err != nil {
		if !fortiAPIPatch(o["tun-mtu-downlink"]) {
			return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
		}
	}

	if err = d.Set("override_split_tunnel", flattenWirelessControllerWtpOverrideSplitTunnel(o["override-split-tunnel"], d, "override_split_tunnel")); err != nil {
		if !fortiAPIPatch(o["override-split-tunnel"]) {
			return fmt.Errorf("Error reading override_split_tunnel: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_path", flattenWirelessControllerWtpSplitTunnelingAclPath(o["split-tunneling-acl-path"], d, "split_tunneling_acl_path")); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-path"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_local_ap_subnet", flattenWirelessControllerWtpSplitTunnelingAclLocalApSubnet(o["split-tunneling-acl-local-ap-subnet"], d, "split_tunneling_acl_local_ap_subnet")); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-local-ap-subnet"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
			if !fortiAPIPatch(o["split-tunneling-acl"]) {
				return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_acl"); ok {
			if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl")); err != nil {
				if !fortiAPIPatch(o["split-tunneling-acl"]) {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_lan", flattenWirelessControllerWtpOverrideLan(o["override-lan"], d, "override_lan")); err != nil {
		if !fortiAPIPatch(o["override-lan"]) {
			return fmt.Errorf("Error reading override_lan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan", flattenWirelessControllerWtpLan(o["lan"], d, "lan")); err != nil {
			if !fortiAPIPatch(o["lan"]) {
				return fmt.Errorf("Error reading lan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan"); ok {
			if err = d.Set("lan", flattenWirelessControllerWtpLan(o["lan"], d, "lan")); err != nil {
				if !fortiAPIPatch(o["lan"]) {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_allowaccess", flattenWirelessControllerWtpOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess")); err != nil {
		if !fortiAPIPatch(o["override-allowaccess"]) {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenWirelessControllerWtpAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("override_login_passwd_change", flattenWirelessControllerWtpOverrideLoginPasswdChange(o["override-login-passwd-change"], d, "override_login_passwd_change")); err != nil {
		if !fortiAPIPatch(o["override-login-passwd-change"]) {
			return fmt.Errorf("Error reading override_login_passwd_change: %v", err)
		}
	}

	if err = d.Set("login_passwd_change", flattenWirelessControllerWtpLoginPasswdChange(o["login-passwd-change"], d, "login_passwd_change")); err != nil {
		if !fortiAPIPatch(o["login-passwd-change"]) {
			return fmt.Errorf("Error reading login_passwd_change: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(o["radio-1"], d, "radio_1")); err != nil {
			if !fortiAPIPatch(o["radio-1"]) {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(o["radio-1"], d, "radio_1")); err != nil {
				if !fortiAPIPatch(o["radio-1"]) {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(o["radio-2"], d, "radio_2")); err != nil {
			if !fortiAPIPatch(o["radio-2"]) {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(o["radio-2"], d, "radio_2")); err != nil {
				if !fortiAPIPatch(o["radio-2"]) {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	if err = d.Set("image_download", flattenWirelessControllerWtpImageDownload(o["image-download"], d, "image_download")); err != nil {
		if !fortiAPIPatch(o["image-download"]) {
			return fmt.Errorf("Error reading image_download: %v", err)
		}
	}

	if err = d.Set("mesh_bridge_enable", flattenWirelessControllerWtpMeshBridgeEnable(o["mesh-bridge-enable"], d, "mesh_bridge_enable")); err != nil {
		if !fortiAPIPatch(o["mesh-bridge-enable"]) {
			return fmt.Errorf("Error reading mesh_bridge_enable: %v", err)
		}
	}

	if err = d.Set("coordinate_latitude", flattenWirelessControllerWtpCoordinateLatitude(o["coordinate-latitude"], d, "coordinate_latitude")); err != nil {
		if !fortiAPIPatch(o["coordinate-latitude"]) {
			return fmt.Errorf("Error reading coordinate_latitude: %v", err)
		}
	}

	if err = d.Set("coordinate_longitude", flattenWirelessControllerWtpCoordinateLongitude(o["coordinate-longitude"], d, "coordinate_longitude")); err != nil {
		if !fortiAPIPatch(o["coordinate-longitude"]) {
			return fmt.Errorf("Error reading coordinate_longitude: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpWtpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpAdmin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLocation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegionX(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegionY(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpBonjourProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLedState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLedState(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideWanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideIpFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpIpFragmentPreventing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpTunMtuUplink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpTunMtuDownlink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideSplitTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclLocalApSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerWtpSplitTunnelingAclId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dest-ip"], _ = expandWirelessControllerWtpSplitTunnelingAclDestIp(d, i["dest_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpSplitTunnelingAclId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclDestIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-mode"], _ = expandWirelessControllerWtpLanPortMode(d, i["port_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port-ssid"], _ = expandWirelessControllerWtpLanPortSsid(d, i["port_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port1_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-mode"], _ = expandWirelessControllerWtpLanPort1Mode(d, i["port1_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port1-ssid"], _ = expandWirelessControllerWtpLanPort1Ssid(d, i["port1_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port2_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-mode"], _ = expandWirelessControllerWtpLanPort2Mode(d, i["port2_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port2-ssid"], _ = expandWirelessControllerWtpLanPort2Ssid(d, i["port2_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port3_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-mode"], _ = expandWirelessControllerWtpLanPort3Mode(d, i["port3_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port3-ssid"], _ = expandWirelessControllerWtpLanPort3Ssid(d, i["port3_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port4_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-mode"], _ = expandWirelessControllerWtpLanPort4Mode(d, i["port4_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port4-ssid"], _ = expandWirelessControllerWtpLanPort4Ssid(d, i["port4_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port5_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-mode"], _ = expandWirelessControllerWtpLanPort5Mode(d, i["port5_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port5-ssid"], _ = expandWirelessControllerWtpLanPort5Ssid(d, i["port5_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port6_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-mode"], _ = expandWirelessControllerWtpLanPort6Mode(d, i["port6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port6-ssid"], _ = expandWirelessControllerWtpLanPort6Ssid(d, i["port6_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port7_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-mode"], _ = expandWirelessControllerWtpLanPort7Mode(d, i["port7_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port7-ssid"], _ = expandWirelessControllerWtpLanPort7Ssid(d, i["port7_ssid"], pre_append)
	}
	pre_append = pre + ".0." + "port8_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-mode"], _ = expandWirelessControllerWtpLanPort8Mode(d, i["port8_mode"], pre_append)
	}
	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := d.GetOk(pre_append); ok {
		result["port8-ssid"], _ = expandWirelessControllerWtpLanPort8Ssid(d, i["port8_ssid"], pre_append)
	}

	return result, nil
}

func expandWirelessControllerWtpLanPortMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort1Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort2Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort2Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort3Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort3Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort4Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort4Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort5Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort5Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort6Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort7Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort7Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort8Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort8Ssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLoginPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandWirelessControllerWtpRadio1RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-band"], _ = expandWirelessControllerWtpRadio1OverrideBand(d, i["override_band"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpRadio1Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-analysis"], _ = expandWirelessControllerWtpRadio1OverrideAnalysis(d, i["override_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio1SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-txpower"], _ = expandWirelessControllerWtpRadio1OverrideTxpower(d, i["override_txpower"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpRadio1AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpRadio1AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpRadio1AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpRadio1PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-vaps"], _ = expandWirelessControllerWtpRadio1OverrideVaps(d, i["override_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpRadio1VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpRadio1Vaps(d, i["vaps"], pre_append)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-channel"], _ = expandWirelessControllerWtpRadio1OverrideChannel(d, i["override_channel"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpRadio1Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpRadio1VapsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1VapsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["chan"], _ = expandWirelessControllerWtpRadio1ChannelChan(d, i["chan"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1ChannelChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-id"], _ = expandWirelessControllerWtpRadio2RadioId(d, i["radio_id"], pre_append)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-band"], _ = expandWirelessControllerWtpRadio2OverrideBand(d, i["override_band"], pre_append)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandWirelessControllerWtpRadio2Band(d, i["band"], pre_append)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-analysis"], _ = expandWirelessControllerWtpRadio2OverrideAnalysis(d, i["override_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio2SpectrumAnalysis(d, i["spectrum_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-txpower"], _ = expandWirelessControllerWtpRadio2OverrideTxpower(d, i["override_txpower"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-level"], _ = expandWirelessControllerWtpRadio2AutoPowerLevel(d, i["auto_power_level"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-high"], _ = expandWirelessControllerWtpRadio2AutoPowerHigh(d, i["auto_power_high"], pre_append)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-power-low"], _ = expandWirelessControllerWtpRadio2AutoPowerLow(d, i["auto_power_low"], pre_append)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandWirelessControllerWtpRadio2PowerLevel(d, i["power_level"], pre_append)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-vaps"], _ = expandWirelessControllerWtpRadio2OverrideVaps(d, i["override_vaps"], pre_append)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {
		result["vap-all"], _ = expandWirelessControllerWtpRadio2VapAll(d, i["vap_all"], pre_append)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["vaps"], _ = expandWirelessControllerWtpRadio2Vaps(d, i["vaps"], pre_append)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-channel"], _ = expandWirelessControllerWtpRadio2OverrideChannel(d, i["override_channel"], pre_append)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandWirelessControllerWtpRadio2Channel(d, i["channel"], pre_append)
	} else {
		result["channel"] = make([]string, 0)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2RadioId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideBand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Band(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerLow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideVaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2VapAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vaps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandWirelessControllerWtpRadio2VapsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2VapsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideChannel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Channel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["chan"], _ = expandWirelessControllerWtpRadio2ChannelChan(d, i["chan"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2ChannelChan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpImageDownload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpMeshBridgeEnable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpCoordinateLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpCoordinateLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("wtp_id"); ok {
		t, err := expandWirelessControllerWtpWtpId(d, v, "wtp_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-id"] = t
		}
	}

	if v, ok := d.GetOkExists("index"); ok {
		t, err := expandWirelessControllerWtpIndex(d, v, "index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok {
		t, err := expandWirelessControllerWtpAdmin(d, v, "admin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWtpName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {
		t, err := expandWirelessControllerWtpLocation(d, v, "location")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok {
		t, err := expandWirelessControllerWtpRegion(d, v, "region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("region_x"); ok {
		t, err := expandWirelessControllerWtpRegionX(d, v, "region_x")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-x"] = t
		}
	}

	if v, ok := d.GetOk("region_y"); ok {
		t, err := expandWirelessControllerWtpRegionY(d, v, "region_y")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-y"] = t
		}
	}

	if v, ok := d.GetOk("wtp_profile"); ok {
		t, err := expandWirelessControllerWtpWtpProfile(d, v, "wtp_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-profile"] = t
		}
	}

	if v, ok := d.GetOk("wtp_mode"); ok {
		t, err := expandWirelessControllerWtpWtpMode(d, v, "wtp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-mode"] = t
		}
	}

	if v, ok := d.GetOk("bonjour_profile"); ok {
		t, err := expandWirelessControllerWtpBonjourProfile(d, v, "bonjour_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bonjour-profile"] = t
		}
	}

	if v, ok := d.GetOk("override_led_state"); ok {
		t, err := expandWirelessControllerWtpOverrideLedState(d, v, "override_led_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-led-state"] = t
		}
	}

	if v, ok := d.GetOk("led_state"); ok {
		t, err := expandWirelessControllerWtpLedState(d, v, "led_state")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-state"] = t
		}
	}

	if v, ok := d.GetOk("override_wan_port_mode"); ok {
		t, err := expandWirelessControllerWtpOverrideWanPortMode(d, v, "override_wan_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_mode"); ok {
		t, err := expandWirelessControllerWtpWanPortMode(d, v, "wan_port_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("override_ip_fragment"); ok {
		t, err := expandWirelessControllerWtpOverrideIpFragment(d, v, "override_ip_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-ip-fragment"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_preventing"); ok {
		t, err := expandWirelessControllerWtpIpFragmentPreventing(d, v, "ip_fragment_preventing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-preventing"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_uplink"); ok {
		t, err := expandWirelessControllerWtpTunMtuUplink(d, v, "tun_mtu_uplink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-uplink"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_downlink"); ok {
		t, err := expandWirelessControllerWtpTunMtuDownlink(d, v, "tun_mtu_downlink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-downlink"] = t
		}
	}

	if v, ok := d.GetOk("override_split_tunnel"); ok {
		t, err := expandWirelessControllerWtpOverrideSplitTunnel(d, v, "override_split_tunnel")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-split-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_path"); ok {
		t, err := expandWirelessControllerWtpSplitTunnelingAclPath(d, v, "split_tunneling_acl_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-path"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok {
		t, err := expandWirelessControllerWtpSplitTunnelingAclLocalApSubnet(d, v, "split_tunneling_acl_local_ap_subnet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-local-ap-subnet"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl"); ok {
		t, err := expandWirelessControllerWtpSplitTunnelingAcl(d, v, "split_tunneling_acl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl"] = t
		}
	}

	if v, ok := d.GetOk("override_lan"); ok {
		t, err := expandWirelessControllerWtpOverrideLan(d, v, "override_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-lan"] = t
		}
	}

	if v, ok := d.GetOk("lan"); ok {
		t, err := expandWirelessControllerWtpLan(d, v, "lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok {
		t, err := expandWirelessControllerWtpOverrideAllowaccess(d, v, "override_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandWirelessControllerWtpAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_login_passwd_change"); ok {
		t, err := expandWirelessControllerWtpOverrideLoginPasswdChange(d, v, "override_login_passwd_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_change"); ok {
		t, err := expandWirelessControllerWtpLoginPasswdChange(d, v, "login_passwd_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok {
		t, err := expandWirelessControllerWtpLoginPasswd(d, v, "login_passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok {
		t, err := expandWirelessControllerWtpRadio1(d, v, "radio_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok {
		t, err := expandWirelessControllerWtpRadio2(d, v, "radio_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	if v, ok := d.GetOk("image_download"); ok {
		t, err := expandWirelessControllerWtpImageDownload(d, v, "image_download")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-download"] = t
		}
	}

	if v, ok := d.GetOk("mesh_bridge_enable"); ok {
		t, err := expandWirelessControllerWtpMeshBridgeEnable(d, v, "mesh_bridge_enable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-bridge-enable"] = t
		}
	}

	if v, ok := d.GetOk("coordinate_latitude"); ok {
		t, err := expandWirelessControllerWtpCoordinateLatitude(d, v, "coordinate_latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinate-latitude"] = t
		}
	}

	if v, ok := d.GetOk("coordinate_longitude"); ok {
		t, err := expandWirelessControllerWtpCoordinateLongitude(d, v, "coordinate_longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinate-longitude"] = t
		}
	}

	return &obj, nil
}
