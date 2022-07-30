// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
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
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
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
			"firmware_provision": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"apcfg_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
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
				ValidateFunc: validation.StringLenBetween(0, 128),
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
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"radio_3": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"radio_4": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"drma_manual_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerWtpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerWtp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtp resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWtp(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerWtp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtp resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWtp(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerWtp(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerWtp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtp resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpWtpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegionX(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRegionY(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpFirmwareProvision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpWtpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpApcfgProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpBonjourProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLedState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLedState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideWanPortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpWanPortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideIpFragment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpIpFragmentPreventing(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpTunMtuUplink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpTunMtuDownlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideSplitTunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclLocalApSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAcl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenWirelessControllerWtpSplitTunnelingAclId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := i["dest-ip"]; ok {

			tmp["dest_ip"] = flattenWirelessControllerWtpSplitTunnelingAclDestIp(i["dest-ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerWtpSplitTunnelingAclId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpSplitTunnelingAclDestIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenWirelessControllerWtpOverrideLan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLan(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := i["port-mode"]; ok {

		result["port_mode"] = flattenWirelessControllerWtpLanPortMode(i["port-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_ssid"
	if _, ok := i["port-ssid"]; ok {

		result["port_ssid"] = flattenWirelessControllerWtpLanPortSsid(i["port-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port1_mode"
	if _, ok := i["port1-mode"]; ok {

		result["port1_mode"] = flattenWirelessControllerWtpLanPort1Mode(i["port1-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := i["port1-ssid"]; ok {

		result["port1_ssid"] = flattenWirelessControllerWtpLanPort1Ssid(i["port1-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port2_mode"
	if _, ok := i["port2-mode"]; ok {

		result["port2_mode"] = flattenWirelessControllerWtpLanPort2Mode(i["port2-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := i["port2-ssid"]; ok {

		result["port2_ssid"] = flattenWirelessControllerWtpLanPort2Ssid(i["port2-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port3_mode"
	if _, ok := i["port3-mode"]; ok {

		result["port3_mode"] = flattenWirelessControllerWtpLanPort3Mode(i["port3-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := i["port3-ssid"]; ok {

		result["port3_ssid"] = flattenWirelessControllerWtpLanPort3Ssid(i["port3-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port4_mode"
	if _, ok := i["port4-mode"]; ok {

		result["port4_mode"] = flattenWirelessControllerWtpLanPort4Mode(i["port4-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := i["port4-ssid"]; ok {

		result["port4_ssid"] = flattenWirelessControllerWtpLanPort4Ssid(i["port4-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port5_mode"
	if _, ok := i["port5-mode"]; ok {

		result["port5_mode"] = flattenWirelessControllerWtpLanPort5Mode(i["port5-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := i["port5-ssid"]; ok {

		result["port5_ssid"] = flattenWirelessControllerWtpLanPort5Ssid(i["port5-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port6_mode"
	if _, ok := i["port6-mode"]; ok {

		result["port6_mode"] = flattenWirelessControllerWtpLanPort6Mode(i["port6-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := i["port6-ssid"]; ok {

		result["port6_ssid"] = flattenWirelessControllerWtpLanPort6Ssid(i["port6-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port7_mode"
	if _, ok := i["port7-mode"]; ok {

		result["port7_mode"] = flattenWirelessControllerWtpLanPort7Mode(i["port7-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := i["port7-ssid"]; ok {

		result["port7_ssid"] = flattenWirelessControllerWtpLanPort7Ssid(i["port7-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port8_mode"
	if _, ok := i["port8-mode"]; ok {

		result["port8_mode"] = flattenWirelessControllerWtpLanPort8Mode(i["port8-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := i["port8-ssid"]; ok {

		result["port8_ssid"] = flattenWirelessControllerWtpLanPort8Ssid(i["port8-ssid"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := i["port-esl-mode"]; ok {

		result["port_esl_mode"] = flattenWirelessControllerWtpLanPortEslMode(i["port-esl-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := i["port-esl-ssid"]; ok {

		result["port_esl_ssid"] = flattenWirelessControllerWtpLanPortEslSsid(i["port-esl-ssid"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpLanPortMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort1Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort1Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort2Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort2Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort3Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort3Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort4Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort4Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort5Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort5Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort6Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort6Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort7Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort7Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort8Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPort8Ssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortEslMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLanPortEslSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpOverrideLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLoginPasswdChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpLoginPasswd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {

		result["radio_id"] = flattenWirelessControllerWtpRadio1RadioId(i["radio-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {

		result["override_band"] = flattenWirelessControllerWtpRadio1OverrideBand(i["override-band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {

		result["band"] = flattenWirelessControllerWtpRadio1Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {

		result["override_analysis"] = flattenWirelessControllerWtpRadio1OverrideAnalysis(i["override-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {

		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio1SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {

		result["override_txpower"] = flattenWirelessControllerWtpRadio1OverrideTxpower(i["override-txpower"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {

		result["auto_power_level"] = flattenWirelessControllerWtpRadio1AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {

		result["auto_power_high"] = flattenWirelessControllerWtpRadio1AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {

		result["auto_power_low"] = flattenWirelessControllerWtpRadio1AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {

		result["auto_power_target"] = flattenWirelessControllerWtpRadio1AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {

		result["power_mode"] = flattenWirelessControllerWtpRadio1PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {

		result["power_level"] = flattenWirelessControllerWtpRadio1PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {

		result["power_value"] = flattenWirelessControllerWtpRadio1PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {

		result["override_vaps"] = flattenWirelessControllerWtpRadio1OverrideVaps(i["override-vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {

		result["vap_all"] = flattenWirelessControllerWtpRadio1VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {

		result["vaps"] = flattenWirelessControllerWtpRadio1Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {

		result["override_channel"] = flattenWirelessControllerWtpRadio1OverrideChannel(i["override-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {

		result["channel"] = flattenWirelessControllerWtpRadio1Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {

		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio1DrmaManualMode(i["drma-manual-mode"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio1RadioId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideTxpower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideVaps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerWtpRadio1VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpRadio1VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1OverrideChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["chan"] = flattenWirelessControllerWtpRadio1ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpRadio1ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio1DrmaManualMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := i["radio-id"]; ok {

		result["radio_id"] = flattenWirelessControllerWtpRadio2RadioId(i["radio-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {

		result["override_band"] = flattenWirelessControllerWtpRadio2OverrideBand(i["override-band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {

		result["band"] = flattenWirelessControllerWtpRadio2Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {

		result["override_analysis"] = flattenWirelessControllerWtpRadio2OverrideAnalysis(i["override-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {

		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio2SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {

		result["override_txpower"] = flattenWirelessControllerWtpRadio2OverrideTxpower(i["override-txpower"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {

		result["auto_power_level"] = flattenWirelessControllerWtpRadio2AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {

		result["auto_power_high"] = flattenWirelessControllerWtpRadio2AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {

		result["auto_power_low"] = flattenWirelessControllerWtpRadio2AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {

		result["auto_power_target"] = flattenWirelessControllerWtpRadio2AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {

		result["power_mode"] = flattenWirelessControllerWtpRadio2PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {

		result["power_level"] = flattenWirelessControllerWtpRadio2PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {

		result["power_value"] = flattenWirelessControllerWtpRadio2PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {

		result["override_vaps"] = flattenWirelessControllerWtpRadio2OverrideVaps(i["override-vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {

		result["vap_all"] = flattenWirelessControllerWtpRadio2VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {

		result["vaps"] = flattenWirelessControllerWtpRadio2Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {

		result["override_channel"] = flattenWirelessControllerWtpRadio2OverrideChannel(i["override-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {

		result["channel"] = flattenWirelessControllerWtpRadio2Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {

		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio2DrmaManualMode(i["drma-manual-mode"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio2RadioId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideTxpower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideVaps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerWtpRadio2VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpRadio2VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2OverrideChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["chan"] = flattenWirelessControllerWtpRadio2ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpRadio2ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio2DrmaManualMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {

		result["override_band"] = flattenWirelessControllerWtpRadio3OverrideBand(i["override-band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {

		result["band"] = flattenWirelessControllerWtpRadio3Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {

		result["override_analysis"] = flattenWirelessControllerWtpRadio3OverrideAnalysis(i["override-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {

		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio3SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {

		result["override_txpower"] = flattenWirelessControllerWtpRadio3OverrideTxpower(i["override-txpower"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {

		result["auto_power_level"] = flattenWirelessControllerWtpRadio3AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {

		result["auto_power_high"] = flattenWirelessControllerWtpRadio3AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {

		result["auto_power_low"] = flattenWirelessControllerWtpRadio3AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {

		result["auto_power_target"] = flattenWirelessControllerWtpRadio3AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {

		result["power_mode"] = flattenWirelessControllerWtpRadio3PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {

		result["power_level"] = flattenWirelessControllerWtpRadio3PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {

		result["power_value"] = flattenWirelessControllerWtpRadio3PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {

		result["override_vaps"] = flattenWirelessControllerWtpRadio3OverrideVaps(i["override-vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {

		result["vap_all"] = flattenWirelessControllerWtpRadio3VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {

		result["vaps"] = flattenWirelessControllerWtpRadio3Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {

		result["override_channel"] = flattenWirelessControllerWtpRadio3OverrideChannel(i["override-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {

		result["channel"] = flattenWirelessControllerWtpRadio3Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {

		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio3DrmaManualMode(i["drma-manual-mode"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio3OverrideBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideTxpower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideVaps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerWtpRadio3VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpRadio3VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3OverrideChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["chan"] = flattenWirelessControllerWtpRadio3ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpRadio3ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio3DrmaManualMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "override_band"
	if _, ok := i["override-band"]; ok {

		result["override_band"] = flattenWirelessControllerWtpRadio4OverrideBand(i["override-band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {

		result["band"] = flattenWirelessControllerWtpRadio4Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_analysis"
	if _, ok := i["override-analysis"]; ok {

		result["override_analysis"] = flattenWirelessControllerWtpRadio4OverrideAnalysis(i["override-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := i["spectrum-analysis"]; ok {

		result["spectrum_analysis"] = flattenWirelessControllerWtpRadio4SpectrumAnalysis(i["spectrum-analysis"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_txpower"
	if _, ok := i["override-txpower"]; ok {

		result["override_txpower"] = flattenWirelessControllerWtpRadio4OverrideTxpower(i["override-txpower"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := i["auto-power-level"]; ok {

		result["auto_power_level"] = flattenWirelessControllerWtpRadio4AutoPowerLevel(i["auto-power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := i["auto-power-high"]; ok {

		result["auto_power_high"] = flattenWirelessControllerWtpRadio4AutoPowerHigh(i["auto-power-high"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := i["auto-power-low"]; ok {

		result["auto_power_low"] = flattenWirelessControllerWtpRadio4AutoPowerLow(i["auto-power-low"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := i["auto-power-target"]; ok {

		result["auto_power_target"] = flattenWirelessControllerWtpRadio4AutoPowerTarget(i["auto-power-target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_mode"
	if _, ok := i["power-mode"]; ok {

		result["power_mode"] = flattenWirelessControllerWtpRadio4PowerMode(i["power-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {

		result["power_level"] = flattenWirelessControllerWtpRadio4PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_value"
	if _, ok := i["power-value"]; ok {

		result["power_value"] = flattenWirelessControllerWtpRadio4PowerValue(i["power-value"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_vaps"
	if _, ok := i["override-vaps"]; ok {

		result["override_vaps"] = flattenWirelessControllerWtpRadio4OverrideVaps(i["override-vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vap_all"
	if _, ok := i["vap-all"]; ok {

		result["vap_all"] = flattenWirelessControllerWtpRadio4VapAll(i["vap-all"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vaps"
	if _, ok := i["vaps"]; ok {

		result["vaps"] = flattenWirelessControllerWtpRadio4Vaps(i["vaps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_channel"
	if _, ok := i["override-channel"]; ok {

		result["override_channel"] = flattenWirelessControllerWtpRadio4OverrideChannel(i["override-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {

		result["channel"] = flattenWirelessControllerWtpRadio4Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := i["drma-manual-mode"]; ok {

		result["drma_manual_mode"] = flattenWirelessControllerWtpRadio4DrmaManualMode(i["drma-manual-mode"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWirelessControllerWtpRadio4OverrideBand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4SpectrumAnalysis(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideTxpower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4AutoPowerTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4PowerValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideVaps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4VapAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Vaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerWtpRadio4VapsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWirelessControllerWtpRadio4VapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4OverrideChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4Channel(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["chan"] = flattenWirelessControllerWtpRadio4ChannelChan(i["chan"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "chan", d)
	return result
}

func flattenWirelessControllerWtpRadio4ChannelChan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpRadio4DrmaManualMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpImageDownload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpMeshBridgeEnable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpCoordinateLatitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpCoordinateLongitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("wtp_id", flattenWirelessControllerWtpWtpId(o["wtp-id"], d, "wtp_id", sv)); err != nil {
		if !fortiAPIPatch(o["wtp-id"]) {
			return fmt.Errorf("Error reading wtp_id: %v", err)
		}
	}

	if err = d.Set("index", flattenWirelessControllerWtpIndex(o["index"], d, "index", sv)); err != nil {
		if !fortiAPIPatch(o["index"]) {
			return fmt.Errorf("Error reading index: %v", err)
		}
	}

	if err = d.Set("uuid", flattenWirelessControllerWtpUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("admin", flattenWirelessControllerWtpAdmin(o["admin"], d, "admin", sv)); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("name", flattenWirelessControllerWtpName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("location", flattenWirelessControllerWtpLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("region", flattenWirelessControllerWtpRegion(o["region"], d, "region", sv)); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("region_x", flattenWirelessControllerWtpRegionX(o["region-x"], d, "region_x", sv)); err != nil {
		if !fortiAPIPatch(o["region-x"]) {
			return fmt.Errorf("Error reading region_x: %v", err)
		}
	}

	if err = d.Set("region_y", flattenWirelessControllerWtpRegionY(o["region-y"], d, "region_y", sv)); err != nil {
		if !fortiAPIPatch(o["region-y"]) {
			return fmt.Errorf("Error reading region_y: %v", err)
		}
	}

	if err = d.Set("firmware_provision", flattenWirelessControllerWtpFirmwareProvision(o["firmware-provision"], d, "firmware_provision", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision"]) {
			return fmt.Errorf("Error reading firmware_provision: %v", err)
		}
	}

	if err = d.Set("firmware_provision_latest", flattenWirelessControllerWtpFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision-latest"]) {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	if err = d.Set("wtp_profile", flattenWirelessControllerWtpWtpProfile(o["wtp-profile"], d, "wtp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["wtp-profile"]) {
			return fmt.Errorf("Error reading wtp_profile: %v", err)
		}
	}

	if err = d.Set("wtp_mode", flattenWirelessControllerWtpWtpMode(o["wtp-mode"], d, "wtp_mode", sv)); err != nil {
		if !fortiAPIPatch(o["wtp-mode"]) {
			return fmt.Errorf("Error reading wtp_mode: %v", err)
		}
	}

	if err = d.Set("apcfg_profile", flattenWirelessControllerWtpApcfgProfile(o["apcfg-profile"], d, "apcfg_profile", sv)); err != nil {
		if !fortiAPIPatch(o["apcfg-profile"]) {
			return fmt.Errorf("Error reading apcfg_profile: %v", err)
		}
	}

	if err = d.Set("bonjour_profile", flattenWirelessControllerWtpBonjourProfile(o["bonjour-profile"], d, "bonjour_profile", sv)); err != nil {
		if !fortiAPIPatch(o["bonjour-profile"]) {
			return fmt.Errorf("Error reading bonjour_profile: %v", err)
		}
	}

	if err = d.Set("override_led_state", flattenWirelessControllerWtpOverrideLedState(o["override-led-state"], d, "override_led_state", sv)); err != nil {
		if !fortiAPIPatch(o["override-led-state"]) {
			return fmt.Errorf("Error reading override_led_state: %v", err)
		}
	}

	if err = d.Set("led_state", flattenWirelessControllerWtpLedState(o["led-state"], d, "led_state", sv)); err != nil {
		if !fortiAPIPatch(o["led-state"]) {
			return fmt.Errorf("Error reading led_state: %v", err)
		}
	}

	if err = d.Set("override_wan_port_mode", flattenWirelessControllerWtpOverrideWanPortMode(o["override-wan-port-mode"], d, "override_wan_port_mode", sv)); err != nil {
		if !fortiAPIPatch(o["override-wan-port-mode"]) {
			return fmt.Errorf("Error reading override_wan_port_mode: %v", err)
		}
	}

	if err = d.Set("wan_port_mode", flattenWirelessControllerWtpWanPortMode(o["wan-port-mode"], d, "wan_port_mode", sv)); err != nil {
		if !fortiAPIPatch(o["wan-port-mode"]) {
			return fmt.Errorf("Error reading wan_port_mode: %v", err)
		}
	}

	if err = d.Set("override_ip_fragment", flattenWirelessControllerWtpOverrideIpFragment(o["override-ip-fragment"], d, "override_ip_fragment", sv)); err != nil {
		if !fortiAPIPatch(o["override-ip-fragment"]) {
			return fmt.Errorf("Error reading override_ip_fragment: %v", err)
		}
	}

	if err = d.Set("ip_fragment_preventing", flattenWirelessControllerWtpIpFragmentPreventing(o["ip-fragment-preventing"], d, "ip_fragment_preventing", sv)); err != nil {
		if !fortiAPIPatch(o["ip-fragment-preventing"]) {
			return fmt.Errorf("Error reading ip_fragment_preventing: %v", err)
		}
	}

	if err = d.Set("tun_mtu_uplink", flattenWirelessControllerWtpTunMtuUplink(o["tun-mtu-uplink"], d, "tun_mtu_uplink", sv)); err != nil {
		if !fortiAPIPatch(o["tun-mtu-uplink"]) {
			return fmt.Errorf("Error reading tun_mtu_uplink: %v", err)
		}
	}

	if err = d.Set("tun_mtu_downlink", flattenWirelessControllerWtpTunMtuDownlink(o["tun-mtu-downlink"], d, "tun_mtu_downlink", sv)); err != nil {
		if !fortiAPIPatch(o["tun-mtu-downlink"]) {
			return fmt.Errorf("Error reading tun_mtu_downlink: %v", err)
		}
	}

	if err = d.Set("override_split_tunnel", flattenWirelessControllerWtpOverrideSplitTunnel(o["override-split-tunnel"], d, "override_split_tunnel", sv)); err != nil {
		if !fortiAPIPatch(o["override-split-tunnel"]) {
			return fmt.Errorf("Error reading override_split_tunnel: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_path", flattenWirelessControllerWtpSplitTunnelingAclPath(o["split-tunneling-acl-path"], d, "split_tunneling_acl_path", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-path"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_path: %v", err)
		}
	}

	if err = d.Set("split_tunneling_acl_local_ap_subnet", flattenWirelessControllerWtpSplitTunnelingAclLocalApSubnet(o["split-tunneling-acl-local-ap-subnet"], d, "split_tunneling_acl_local_ap_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling-acl-local-ap-subnet"]) {
			return fmt.Errorf("Error reading split_tunneling_acl_local_ap_subnet: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl", sv)); err != nil {
			if !fortiAPIPatch(o["split-tunneling-acl"]) {
				return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_acl"); ok {
			if err = d.Set("split_tunneling_acl", flattenWirelessControllerWtpSplitTunnelingAcl(o["split-tunneling-acl"], d, "split_tunneling_acl", sv)); err != nil {
				if !fortiAPIPatch(o["split-tunneling-acl"]) {
					return fmt.Errorf("Error reading split_tunneling_acl: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_lan", flattenWirelessControllerWtpOverrideLan(o["override-lan"], d, "override_lan", sv)); err != nil {
		if !fortiAPIPatch(o["override-lan"]) {
			return fmt.Errorf("Error reading override_lan: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan", flattenWirelessControllerWtpLan(o["lan"], d, "lan", sv)); err != nil {
			if !fortiAPIPatch(o["lan"]) {
				return fmt.Errorf("Error reading lan: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan"); ok {
			if err = d.Set("lan", flattenWirelessControllerWtpLan(o["lan"], d, "lan", sv)); err != nil {
				if !fortiAPIPatch(o["lan"]) {
					return fmt.Errorf("Error reading lan: %v", err)
				}
			}
		}
	}

	if err = d.Set("override_allowaccess", flattenWirelessControllerWtpOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["override-allowaccess"]) {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenWirelessControllerWtpAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("override_login_passwd_change", flattenWirelessControllerWtpOverrideLoginPasswdChange(o["override-login-passwd-change"], d, "override_login_passwd_change", sv)); err != nil {
		if !fortiAPIPatch(o["override-login-passwd-change"]) {
			return fmt.Errorf("Error reading override_login_passwd_change: %v", err)
		}
	}

	if err = d.Set("login_passwd_change", flattenWirelessControllerWtpLoginPasswdChange(o["login-passwd-change"], d, "login_passwd_change", sv)); err != nil {
		if !fortiAPIPatch(o["login-passwd-change"]) {
			return fmt.Errorf("Error reading login_passwd_change: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(o["radio-1"], d, "radio_1", sv)); err != nil {
			if !fortiAPIPatch(o["radio-1"]) {
				return fmt.Errorf("Error reading radio_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_1"); ok {
			if err = d.Set("radio_1", flattenWirelessControllerWtpRadio1(o["radio-1"], d, "radio_1", sv)); err != nil {
				if !fortiAPIPatch(o["radio-1"]) {
					return fmt.Errorf("Error reading radio_1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(o["radio-2"], d, "radio_2", sv)); err != nil {
			if !fortiAPIPatch(o["radio-2"]) {
				return fmt.Errorf("Error reading radio_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_2"); ok {
			if err = d.Set("radio_2", flattenWirelessControllerWtpRadio2(o["radio-2"], d, "radio_2", sv)); err != nil {
				if !fortiAPIPatch(o["radio-2"]) {
					return fmt.Errorf("Error reading radio_2: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_3", flattenWirelessControllerWtpRadio3(o["radio-3"], d, "radio_3", sv)); err != nil {
			if !fortiAPIPatch(o["radio-3"]) {
				return fmt.Errorf("Error reading radio_3: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_3"); ok {
			if err = d.Set("radio_3", flattenWirelessControllerWtpRadio3(o["radio-3"], d, "radio_3", sv)); err != nil {
				if !fortiAPIPatch(o["radio-3"]) {
					return fmt.Errorf("Error reading radio_3: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("radio_4", flattenWirelessControllerWtpRadio4(o["radio-4"], d, "radio_4", sv)); err != nil {
			if !fortiAPIPatch(o["radio-4"]) {
				return fmt.Errorf("Error reading radio_4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("radio_4"); ok {
			if err = d.Set("radio_4", flattenWirelessControllerWtpRadio4(o["radio-4"], d, "radio_4", sv)); err != nil {
				if !fortiAPIPatch(o["radio-4"]) {
					return fmt.Errorf("Error reading radio_4: %v", err)
				}
			}
		}
	}

	if err = d.Set("image_download", flattenWirelessControllerWtpImageDownload(o["image-download"], d, "image_download", sv)); err != nil {
		if !fortiAPIPatch(o["image-download"]) {
			return fmt.Errorf("Error reading image_download: %v", err)
		}
	}

	if err = d.Set("mesh_bridge_enable", flattenWirelessControllerWtpMeshBridgeEnable(o["mesh-bridge-enable"], d, "mesh_bridge_enable", sv)); err != nil {
		if !fortiAPIPatch(o["mesh-bridge-enable"]) {
			return fmt.Errorf("Error reading mesh_bridge_enable: %v", err)
		}
	}

	if err = d.Set("coordinate_latitude", flattenWirelessControllerWtpCoordinateLatitude(o["coordinate-latitude"], d, "coordinate_latitude", sv)); err != nil {
		if !fortiAPIPatch(o["coordinate-latitude"]) {
			return fmt.Errorf("Error reading coordinate_latitude: %v", err)
		}
	}

	if err = d.Set("coordinate_longitude", flattenWirelessControllerWtpCoordinateLongitude(o["coordinate-longitude"], d, "coordinate_longitude", sv)); err != nil {
		if !fortiAPIPatch(o["coordinate-longitude"]) {
			return fmt.Errorf("Error reading coordinate_longitude: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerWtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerWtpWtpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegionX(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRegionY(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpFirmwareProvision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWtpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpApcfgProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpBonjourProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLedState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLedState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideWanPortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpWanPortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideIpFragment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpIpFragmentPreventing(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpTunMtuUplink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpTunMtuDownlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideSplitTunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclLocalApSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAcl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandWirelessControllerWtpSplitTunnelingAclId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dest_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["dest-ip"], _ = expandWirelessControllerWtpSplitTunnelingAclDestIp(d, i["dest_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpSplitTunnelingAclId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpSplitTunnelingAclDestIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "port_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port-mode"], _ = expandWirelessControllerWtpLanPortMode(d, i["port_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port-ssid"], _ = expandWirelessControllerWtpLanPortSsid(d, i["port_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port1_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port1-mode"], _ = expandWirelessControllerWtpLanPort1Mode(d, i["port1_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port1_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port1-ssid"], _ = expandWirelessControllerWtpLanPort1Ssid(d, i["port1_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port2_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port2-mode"], _ = expandWirelessControllerWtpLanPort2Mode(d, i["port2_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port2_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port2-ssid"], _ = expandWirelessControllerWtpLanPort2Ssid(d, i["port2_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port3_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port3-mode"], _ = expandWirelessControllerWtpLanPort3Mode(d, i["port3_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port3_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port3-ssid"], _ = expandWirelessControllerWtpLanPort3Ssid(d, i["port3_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port4_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port4-mode"], _ = expandWirelessControllerWtpLanPort4Mode(d, i["port4_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port4_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port4-ssid"], _ = expandWirelessControllerWtpLanPort4Ssid(d, i["port4_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port5_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port5-mode"], _ = expandWirelessControllerWtpLanPort5Mode(d, i["port5_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port5_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port5-ssid"], _ = expandWirelessControllerWtpLanPort5Ssid(d, i["port5_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port6_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port6-mode"], _ = expandWirelessControllerWtpLanPort6Mode(d, i["port6_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port6_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port6-ssid"], _ = expandWirelessControllerWtpLanPort6Ssid(d, i["port6_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port7_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port7-mode"], _ = expandWirelessControllerWtpLanPort7Mode(d, i["port7_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port7_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port7-ssid"], _ = expandWirelessControllerWtpLanPort7Ssid(d, i["port7_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port8_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port8-mode"], _ = expandWirelessControllerWtpLanPort8Mode(d, i["port8_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port8_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port8-ssid"], _ = expandWirelessControllerWtpLanPort8Ssid(d, i["port8_ssid"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_esl_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["port-esl-mode"], _ = expandWirelessControllerWtpLanPortEslMode(d, i["port_esl_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "port_esl_ssid"
	if _, ok := d.GetOk(pre_append); ok {

		result["port-esl-ssid"], _ = expandWirelessControllerWtpLanPortEslSsid(d, i["port_esl_ssid"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpLanPortMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort1Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort1Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort2Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort2Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort3Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort3Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort4Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort4Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort5Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort5Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort6Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort6Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort7Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort7Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort8Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPort8Ssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortEslMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLanPortEslSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpOverrideLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLoginPasswdChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpLoginPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["radio-id"], _ = expandWirelessControllerWtpRadio1RadioId(d, i["radio_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-band"], _ = expandWirelessControllerWtpRadio1OverrideBand(d, i["override_band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {

		result["band"], _ = expandWirelessControllerWtpRadio1Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-analysis"], _ = expandWirelessControllerWtpRadio1OverrideAnalysis(d, i["override_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio1SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-txpower"], _ = expandWirelessControllerWtpRadio1OverrideTxpower(d, i["override_txpower"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-level"], _ = expandWirelessControllerWtpRadio1AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-high"], _ = expandWirelessControllerWtpRadio1AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-low"], _ = expandWirelessControllerWtpRadio1AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-target"], _ = expandWirelessControllerWtpRadio1AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-mode"], _ = expandWirelessControllerWtpRadio1PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-level"], _ = expandWirelessControllerWtpRadio1PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-value"], _ = expandWirelessControllerWtpRadio1PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-vaps"], _ = expandWirelessControllerWtpRadio1OverrideVaps(d, i["override_vaps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {

		result["vap-all"], _ = expandWirelessControllerWtpRadio1VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["vaps"], _ = expandWirelessControllerWtpRadio1Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-channel"], _ = expandWirelessControllerWtpRadio1OverrideChannel(d, i["override_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["channel"], _ = expandWirelessControllerWtpRadio1Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio1DrmaManualMode(d, i["drma_manual_mode"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1RadioId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideTxpower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideVaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerWtpRadio1VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1OverrideChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["chan"], _ = expandWirelessControllerWtpRadio1ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio1ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio1DrmaManualMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "radio_id"
	if _, ok := d.GetOk(pre_append); ok {

		result["radio-id"], _ = expandWirelessControllerWtpRadio2RadioId(d, i["radio_id"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-band"], _ = expandWirelessControllerWtpRadio2OverrideBand(d, i["override_band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {

		result["band"], _ = expandWirelessControllerWtpRadio2Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-analysis"], _ = expandWirelessControllerWtpRadio2OverrideAnalysis(d, i["override_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio2SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-txpower"], _ = expandWirelessControllerWtpRadio2OverrideTxpower(d, i["override_txpower"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-level"], _ = expandWirelessControllerWtpRadio2AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-high"], _ = expandWirelessControllerWtpRadio2AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-low"], _ = expandWirelessControllerWtpRadio2AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-target"], _ = expandWirelessControllerWtpRadio2AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-mode"], _ = expandWirelessControllerWtpRadio2PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-level"], _ = expandWirelessControllerWtpRadio2PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-value"], _ = expandWirelessControllerWtpRadio2PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-vaps"], _ = expandWirelessControllerWtpRadio2OverrideVaps(d, i["override_vaps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {

		result["vap-all"], _ = expandWirelessControllerWtpRadio2VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["vaps"], _ = expandWirelessControllerWtpRadio2Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-channel"], _ = expandWirelessControllerWtpRadio2OverrideChannel(d, i["override_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["channel"], _ = expandWirelessControllerWtpRadio2Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio2DrmaManualMode(d, i["drma_manual_mode"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2RadioId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideTxpower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideVaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerWtpRadio2VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2OverrideChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["chan"], _ = expandWirelessControllerWtpRadio2ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio2ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio2DrmaManualMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-band"], _ = expandWirelessControllerWtpRadio3OverrideBand(d, i["override_band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {

		result["band"], _ = expandWirelessControllerWtpRadio3Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-analysis"], _ = expandWirelessControllerWtpRadio3OverrideAnalysis(d, i["override_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio3SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-txpower"], _ = expandWirelessControllerWtpRadio3OverrideTxpower(d, i["override_txpower"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-level"], _ = expandWirelessControllerWtpRadio3AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-high"], _ = expandWirelessControllerWtpRadio3AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-low"], _ = expandWirelessControllerWtpRadio3AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-target"], _ = expandWirelessControllerWtpRadio3AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-mode"], _ = expandWirelessControllerWtpRadio3PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-level"], _ = expandWirelessControllerWtpRadio3PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-value"], _ = expandWirelessControllerWtpRadio3PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-vaps"], _ = expandWirelessControllerWtpRadio3OverrideVaps(d, i["override_vaps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {

		result["vap-all"], _ = expandWirelessControllerWtpRadio3VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["vaps"], _ = expandWirelessControllerWtpRadio3Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-channel"], _ = expandWirelessControllerWtpRadio3OverrideChannel(d, i["override_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["channel"], _ = expandWirelessControllerWtpRadio3Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio3DrmaManualMode(d, i["drma_manual_mode"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio3OverrideBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideTxpower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideVaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerWtpRadio3VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio3VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3OverrideChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["chan"], _ = expandWirelessControllerWtpRadio3ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio3ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio3DrmaManualMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "override_band"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-band"], _ = expandWirelessControllerWtpRadio4OverrideBand(d, i["override_band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {

		result["band"], _ = expandWirelessControllerWtpRadio4Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-analysis"], _ = expandWirelessControllerWtpRadio4OverrideAnalysis(d, i["override_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "spectrum_analysis"
	if _, ok := d.GetOk(pre_append); ok {

		result["spectrum-analysis"], _ = expandWirelessControllerWtpRadio4SpectrumAnalysis(d, i["spectrum_analysis"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_txpower"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-txpower"], _ = expandWirelessControllerWtpRadio4OverrideTxpower(d, i["override_txpower"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-level"], _ = expandWirelessControllerWtpRadio4AutoPowerLevel(d, i["auto_power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_high"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-high"], _ = expandWirelessControllerWtpRadio4AutoPowerHigh(d, i["auto_power_high"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_low"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-low"], _ = expandWirelessControllerWtpRadio4AutoPowerLow(d, i["auto_power_low"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_power_target"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-power-target"], _ = expandWirelessControllerWtpRadio4AutoPowerTarget(d, i["auto_power_target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-mode"], _ = expandWirelessControllerWtpRadio4PowerMode(d, i["power_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-level"], _ = expandWirelessControllerWtpRadio4PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_value"
	if _, ok := d.GetOk(pre_append); ok {

		result["power-value"], _ = expandWirelessControllerWtpRadio4PowerValue(d, i["power_value"], pre_append, sv)
	}
	pre_append = pre + ".0." + "override_vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-vaps"], _ = expandWirelessControllerWtpRadio4OverrideVaps(d, i["override_vaps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vap_all"
	if _, ok := d.GetOk(pre_append); ok {

		result["vap-all"], _ = expandWirelessControllerWtpRadio4VapAll(d, i["vap_all"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vaps"
	if _, ok := d.GetOk(pre_append); ok {

		result["vaps"], _ = expandWirelessControllerWtpRadio4Vaps(d, i["vaps"], pre_append, sv)
	} else {
		result["vaps"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "override_channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["override-channel"], _ = expandWirelessControllerWtpRadio4OverrideChannel(d, i["override_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {

		result["channel"], _ = expandWirelessControllerWtpRadio4Channel(d, i["channel"], pre_append, sv)
	} else {
		result["channel"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "drma_manual_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["drma-manual-mode"], _ = expandWirelessControllerWtpRadio4DrmaManualMode(d, i["drma_manual_mode"], pre_append, sv)
	}

	return result, nil
}

func expandWirelessControllerWtpRadio4OverrideBand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4SpectrumAnalysis(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideTxpower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4AutoPowerTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4PowerValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideVaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4VapAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Vaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerWtpRadio4VapsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio4VapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4OverrideChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["chan"], _ = expandWirelessControllerWtpRadio4ChannelChan(d, i["chan"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpRadio4ChannelChan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpRadio4DrmaManualMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpImageDownload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpMeshBridgeEnable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpCoordinateLatitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpCoordinateLongitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("wtp_id"); ok {

		t, err := expandWirelessControllerWtpWtpId(d, v, "wtp_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-id"] = t
		}
	}

	if v, ok := d.GetOkExists("index"); ok {

		t, err := expandWirelessControllerWtpIndex(d, v, "index", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["index"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {

		t, err := expandWirelessControllerWtpUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok {

		t, err := expandWirelessControllerWtpAdmin(d, v, "admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerWtpName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("location"); ok {

		t, err := expandWirelessControllerWtpLocation(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}

	if v, ok := d.GetOk("region"); ok {

		t, err := expandWirelessControllerWtpRegion(d, v, "region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	}

	if v, ok := d.GetOk("region_x"); ok {

		t, err := expandWirelessControllerWtpRegionX(d, v, "region_x", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-x"] = t
		}
	}

	if v, ok := d.GetOk("region_y"); ok {

		t, err := expandWirelessControllerWtpRegionY(d, v, "region_y", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-y"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision"); ok {

		t, err := expandWirelessControllerWtpFirmwareProvision(d, v, "firmware_provision", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok {

		t, err := expandWirelessControllerWtpFirmwareProvisionLatest(d, v, "firmware_provision_latest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	if v, ok := d.GetOk("wtp_profile"); ok {

		t, err := expandWirelessControllerWtpWtpProfile(d, v, "wtp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-profile"] = t
		}
	}

	if v, ok := d.GetOk("wtp_mode"); ok {

		t, err := expandWirelessControllerWtpWtpMode(d, v, "wtp_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtp-mode"] = t
		}
	}

	if v, ok := d.GetOk("apcfg_profile"); ok {

		t, err := expandWirelessControllerWtpApcfgProfile(d, v, "apcfg_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["apcfg-profile"] = t
		}
	}

	if v, ok := d.GetOk("bonjour_profile"); ok {

		t, err := expandWirelessControllerWtpBonjourProfile(d, v, "bonjour_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bonjour-profile"] = t
		}
	}

	if v, ok := d.GetOk("override_led_state"); ok {

		t, err := expandWirelessControllerWtpOverrideLedState(d, v, "override_led_state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-led-state"] = t
		}
	}

	if v, ok := d.GetOk("led_state"); ok {

		t, err := expandWirelessControllerWtpLedState(d, v, "led_state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["led-state"] = t
		}
	}

	if v, ok := d.GetOk("override_wan_port_mode"); ok {

		t, err := expandWirelessControllerWtpOverrideWanPortMode(d, v, "override_wan_port_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("wan_port_mode"); ok {

		t, err := expandWirelessControllerWtpWanPortMode(d, v, "wan_port_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-port-mode"] = t
		}
	}

	if v, ok := d.GetOk("override_ip_fragment"); ok {

		t, err := expandWirelessControllerWtpOverrideIpFragment(d, v, "override_ip_fragment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-ip-fragment"] = t
		}
	}

	if v, ok := d.GetOk("ip_fragment_preventing"); ok {

		t, err := expandWirelessControllerWtpIpFragmentPreventing(d, v, "ip_fragment_preventing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-fragment-preventing"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_uplink"); ok {

		t, err := expandWirelessControllerWtpTunMtuUplink(d, v, "tun_mtu_uplink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-uplink"] = t
		}
	}

	if v, ok := d.GetOk("tun_mtu_downlink"); ok {

		t, err := expandWirelessControllerWtpTunMtuDownlink(d, v, "tun_mtu_downlink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tun-mtu-downlink"] = t
		}
	}

	if v, ok := d.GetOk("override_split_tunnel"); ok {

		t, err := expandWirelessControllerWtpOverrideSplitTunnel(d, v, "override_split_tunnel", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-split-tunnel"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_path"); ok {

		t, err := expandWirelessControllerWtpSplitTunnelingAclPath(d, v, "split_tunneling_acl_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-path"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl_local_ap_subnet"); ok {

		t, err := expandWirelessControllerWtpSplitTunnelingAclLocalApSubnet(d, v, "split_tunneling_acl_local_ap_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl-local-ap-subnet"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_acl"); ok || d.HasChange("split_tunneling_acl") {

		t, err := expandWirelessControllerWtpSplitTunnelingAcl(d, v, "split_tunneling_acl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-acl"] = t
		}
	}

	if v, ok := d.GetOk("override_lan"); ok {

		t, err := expandWirelessControllerWtpOverrideLan(d, v, "override_lan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-lan"] = t
		}
	}

	if v, ok := d.GetOk("lan"); ok {

		t, err := expandWirelessControllerWtpLan(d, v, "lan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok {

		t, err := expandWirelessControllerWtpOverrideAllowaccess(d, v, "override_allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {

		t, err := expandWirelessControllerWtpAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_login_passwd_change"); ok {

		t, err := expandWirelessControllerWtpOverrideLoginPasswdChange(d, v, "override_login_passwd_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_change"); ok {

		t, err := expandWirelessControllerWtpLoginPasswdChange(d, v, "login_passwd_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-change"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok {

		t, err := expandWirelessControllerWtpLoginPasswd(d, v, "login_passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	if v, ok := d.GetOk("radio_1"); ok {

		t, err := expandWirelessControllerWtpRadio1(d, v, "radio_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-1"] = t
		}
	}

	if v, ok := d.GetOk("radio_2"); ok {

		t, err := expandWirelessControllerWtpRadio2(d, v, "radio_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-2"] = t
		}
	}

	if v, ok := d.GetOk("radio_3"); ok {

		t, err := expandWirelessControllerWtpRadio3(d, v, "radio_3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-3"] = t
		}
	}

	if v, ok := d.GetOk("radio_4"); ok {

		t, err := expandWirelessControllerWtpRadio4(d, v, "radio_4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radio-4"] = t
		}
	}

	if v, ok := d.GetOk("image_download"); ok {

		t, err := expandWirelessControllerWtpImageDownload(d, v, "image_download", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-download"] = t
		}
	}

	if v, ok := d.GetOk("mesh_bridge_enable"); ok {

		t, err := expandWirelessControllerWtpMeshBridgeEnable(d, v, "mesh_bridge_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-bridge-enable"] = t
		}
	}

	if v, ok := d.GetOk("coordinate_latitude"); ok {

		t, err := expandWirelessControllerWtpCoordinateLatitude(d, v, "coordinate_latitude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinate-latitude"] = t
		}
	}

	if v, ok := d.GetOk("coordinate_longitude"); ok {

		t, err := expandWirelessControllerWtpCoordinateLongitude(d, v, "coordinate_longitude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinate-longitude"] = t
		}
	}

	return &obj, nil
}
