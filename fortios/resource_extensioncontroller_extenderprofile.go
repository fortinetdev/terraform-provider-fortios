// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: FortiExtender extender profile configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderProfileCreate,
		Read:   resourceExtensionControllerExtenderProfileRead,
		Update: resourceExtensionControllerExtenderProfileUpdate,
		Delete: resourceExtensionControllerExtenderProfileDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 102400000),
				Optional:     true,
				Computed:     true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"cellular": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataplan": &schema.Schema{
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
						"controller_report": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interval": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"signal_threshold": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(10, 50),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"sms_notification": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"alert": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"system_reboot": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"data_exhausted": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"session_disconnect": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"low_signal_strength": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"os_image_fallback": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"mode_switch": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"fgt_backup_mode_switch": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
									"receiver": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),
													Optional:     true,
												},
												"status": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"phone_number": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),
													Optional:     true,
												},
												"alert": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"modem1": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redundant_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"redundant_intf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
									"conn_status": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"default_sim": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gps": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim2_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin_code": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 27),
										Optional:     true,
										Sensitive:    true,
									},
									"sim2_pin_code": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 27),
										Optional:     true,
										Sensitive:    true,
									},
									"preferred_carrier": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"auto_switch": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"disconnect_threshold": &schema.Schema{
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 100),
													Optional:     true,
													Computed:     true,
												},
												"disconnect_period": &schema.Schema{
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(600, 18000),
													Optional:     true,
													Computed:     true,
												},
												"signal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dataplan": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"switch_back": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"switch_back_time": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),
													Optional:     true,
													Computed:     true,
												},
												"switch_back_timer": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"multiple_pdn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pdn1_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pdn2_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pdn3_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pdn4_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
								},
							},
						},
						"modem2": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redundant_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"redundant_intf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
									"conn_status": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"default_sim": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"gps": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim2_pin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sim1_pin_code": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 27),
										Optional:     true,
										Sensitive:    true,
									},
									"sim2_pin_code": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 27),
										Optional:     true,
										Sensitive:    true,
									},
									"preferred_carrier": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"auto_switch": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"disconnect": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"disconnect_threshold": &schema.Schema{
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(1, 100),
													Optional:     true,
													Computed:     true,
												},
												"disconnect_period": &schema.Schema{
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntBetween(600, 18000),
													Optional:     true,
													Computed:     true,
												},
												"signal": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dataplan": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"switch_back": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												"switch_back_time": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),
													Optional:     true,
													Computed:     true,
												},
												"switch_back_timer": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"multiple_pdn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pdn1_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pdn2_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pdn3_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pdn4_dataplan": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"lan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"link_loadbalance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipsec_tunnel": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"backhaul_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"backhaul_ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"backhaul": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"role": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"weight": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 256),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"downlinks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"vap": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"pvid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4089),
										Optional:     true,
									},
								},
							},
						},
						"traffic_split_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"vsdb": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"address": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"service": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"wifi": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": &schema.Schema{
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
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"operating_standard": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"guard_interval": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"channel": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"bandwidth": &schema.Schema{
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
									"beacon_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(100, 3500),
										Optional:     true,
										Computed:     true,
									},
									"n80211d": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_clients": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 512),
										Optional:     true,
									},
									"extension_channel": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bss_color_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bss_color": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 63),
										Optional:     true,
									},
									"lan_ext_vap": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"local_vaps": &schema.Schema{
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
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"operating_standard": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"guard_interval": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"channel": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"bandwidth": &schema.Schema{
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
									"beacon_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(100, 3500),
										Optional:     true,
										Computed:     true,
									},
									"n80211d": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_clients": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 512),
										Optional:     true,
									},
									"extension_channel": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bss_color_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bss_color": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 63),
										Optional:     true,
									},
									"lan_ext_vap": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
									"local_vaps": &schema.Schema{
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
								},
							},
						},
					},
				},
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceExtensionControllerExtenderProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfile resource while getting object: %v", err)
	}

	o, err := c.CreateExtensionControllerExtenderProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerExtenderProfile")
	}

	return resourceExtensionControllerExtenderProfileRead(d, m)
}

func resourceExtensionControllerExtenderProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateExtensionControllerExtenderProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerExtenderProfile")
	}

	return resourceExtensionControllerExtenderProfileRead(d, m)
}

func resourceExtensionControllerExtenderProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtensionControllerExtenderProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderProfile resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileExtension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileBandwidthLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellular(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "controller_report"
	if _, ok := i["controller-report"]; ok {
		result["controller_report"] = flattenExtensionControllerExtenderProfileCellularControllerReport(i["controller-report"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sms_notification"
	if _, ok := i["sms-notification"]; ok {
		result["sms_notification"] = flattenExtensionControllerExtenderProfileCellularSmsNotification(i["sms-notification"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "modem1"
	if _, ok := i["modem1"]; ok {
		result["modem1"] = flattenExtensionControllerExtenderProfileCellularModem1(i["modem1"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "modem2"
	if _, ok := i["modem2"]; ok {
		result["modem2"] = flattenExtensionControllerExtenderProfileCellularModem2(i["modem2"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileCellularDataplanName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileCellularDataplanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReport(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileCellularControllerReportStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenExtensionControllerExtenderProfileCellularControllerReportInterval(i["interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(i["signal-threshold"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularControllerReportStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularSmsNotification(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert(i["alert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver(i["receiver"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlert(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(i["system-reboot"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(i["data-exhausted"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(i["session-disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(i["low-signal-strength"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(i["os-image-fallback"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(i["mode-switch"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(i["fgt-backup-mode-switch"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiver(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if cur_v, ok := i["phone-number"]; ok {
			tmp["phone_number"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if cur_v, ok := i["alert"]; ok {
			tmp["alert"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtensionControllerExtenderProfileCellularModem1RedundantMode(i["redundant-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtensionControllerExtenderProfileCellularModem1RedundantIntf(i["redundant-intf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtensionControllerExtenderProfileCellularModem1ConnStatus(i["conn-status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtensionControllerExtenderProfileCellularModem1DefaultSim(i["default-sim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtensionControllerExtenderProfileCellularModem1Gps(i["gps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtensionControllerExtenderProfileCellularModem1Sim1Pin(i["sim1-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtensionControllerExtenderProfileCellularModem1Sim2Pin(i["sim2-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim1_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim2_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch(i["auto-switch"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := i["multiple-PDN"]; ok {
		result["multiple_pdn"] = flattenExtensionControllerExtenderProfileCellularModem1MultiplePdn(i["multiple-PDN"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := i["pdn1-dataplan"]; ok {
		result["pdn1_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(i["pdn1-dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := i["pdn2-dataplan"]; ok {
		result["pdn2_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(i["pdn2-dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := i["pdn3-dataplan"]; ok {
		result["pdn3_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(i["pdn3-dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := i["pdn4-dataplan"]; ok {
		result["pdn4_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(i["pdn4-dataplan"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Gps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(i["signal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem1MultiplePdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtensionControllerExtenderProfileCellularModem2RedundantMode(i["redundant-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtensionControllerExtenderProfileCellularModem2RedundantIntf(i["redundant-intf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtensionControllerExtenderProfileCellularModem2ConnStatus(i["conn-status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtensionControllerExtenderProfileCellularModem2DefaultSim(i["default-sim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtensionControllerExtenderProfileCellularModem2Gps(i["gps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtensionControllerExtenderProfileCellularModem2Sim1Pin(i["sim1-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtensionControllerExtenderProfileCellularModem2Sim2Pin(i["sim2-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim1_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim2_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch(i["auto-switch"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := i["multiple-PDN"]; ok {
		result["multiple_pdn"] = flattenExtensionControllerExtenderProfileCellularModem2MultiplePdn(i["multiple-PDN"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := i["pdn1-dataplan"]; ok {
		result["pdn1_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(i["pdn1-dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := i["pdn2-dataplan"]; ok {
		result["pdn2_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(i["pdn2-dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := i["pdn3-dataplan"]; ok {
		result["pdn3_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(i["pdn3-dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := i["pdn4-dataplan"]; ok {
		result["pdn4_dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(i["pdn4-dataplan"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Gps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(i["signal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileCellularModem2MultiplePdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := i["link-loadbalance"]; ok {
		result["link_loadbalance"] = flattenExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(i["link-loadbalance"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {
		result["ipsec_tunnel"] = flattenExtensionControllerExtenderProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {
		result["backhaul_interface"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {
		result["backhaul_ip"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul"
	if _, ok := i["backhaul"]; ok {
		result["backhaul"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaul(i["backhaul"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "downlinks"
	if _, ok := i["downlinks"]; ok {
		result["downlinks"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinks(i["downlinks"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "traffic_split_services"
	if _, ok := i["traffic-split-services"]; ok {
		result["traffic_split_services"] = flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(i["traffic-split-services"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaul(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if cur_v, ok := i["role"]; ok {
			tmp["role"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if cur_v, ok := i["weight"]; ok {
			tmp["weight"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinks(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinksName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinksType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinksPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vap"
		if cur_v, ok := i["vap"]; ok {
			tmp["vap"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinksVap(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pvid"
		if cur_v, ok := i["pvid"]; ok {
			tmp["pvid"] = flattenExtensionControllerExtenderProfileLanExtensionDownlinksPvid(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksVap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionDownlinksPvid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vsdb"
		if cur_v, ok := i["vsdb"]; ok {
			tmp["vsdb"] = flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if cur_v, ok := i["address"]; ok {
			tmp["address"] = flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if cur_v, ok := i["service"]; ok {
			tmp["service"] = flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifi(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "country"
	if _, ok := i["country"]; ok {
		result["country"] = flattenExtensionControllerExtenderProfileWifiCountry(i["country"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "radio_1"
	if _, ok := i["radio-1"]; ok {
		result["radio_1"] = flattenExtensionControllerExtenderProfileWifiRadio1(i["radio-1"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "radio_2"
	if _, ok := i["radio-2"]; ok {
		result["radio_2"] = flattenExtensionControllerExtenderProfileWifiRadio2(i["radio-2"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileWifiCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenExtensionControllerExtenderProfileWifiRadio1Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenExtensionControllerExtenderProfileWifiRadio1Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileWifiRadio1Status(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "operating_standard"
	if _, ok := i["operating-standard"]; ok {
		result["operating_standard"] = flattenExtensionControllerExtenderProfileWifiRadio1OperatingStandard(i["operating-standard"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "guard_interval"
	if _, ok := i["guard-interval"]; ok {
		result["guard_interval"] = flattenExtensionControllerExtenderProfileWifiRadio1GuardInterval(i["guard-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenExtensionControllerExtenderProfileWifiRadio1Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth"
	if _, ok := i["bandwidth"]; ok {
		result["bandwidth"] = flattenExtensionControllerExtenderProfileWifiRadio1Bandwidth(i["bandwidth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenExtensionControllerExtenderProfileWifiRadio1PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenExtensionControllerExtenderProfileWifiRadio1BeaconInterval(i["beacon-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenExtensionControllerExtenderProfileWifiRadio180211D(i["80211d"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenExtensionControllerExtenderProfileWifiRadio1MaxClients(i["max-clients"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "extension_channel"
	if _, ok := i["extension-channel"]; ok {
		result["extension_channel"] = flattenExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(i["extension-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenExtensionControllerExtenderProfileWifiRadio1BssColorMode(i["bss-color-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenExtensionControllerExtenderProfileWifiRadio1BssColor(i["bss-color"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := i["lan-ext-vap"]; ok {
		result["lan_ext_vap"] = flattenExtensionControllerExtenderProfileWifiRadio1LanExtVap(i["lan-ext-vap"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "local_vaps"
	if _, ok := i["local-vaps"]; ok {
		result["local_vaps"] = flattenExtensionControllerExtenderProfileWifiRadio1LocalVaps(i["local-vaps"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileWifiRadio1Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1OperatingStandard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1GuardInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Channel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1Bandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1BeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio180211D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1MaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BssColorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1BssColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio1LanExtVap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio1LocalVaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileWifiRadio1LocalVapsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileWifiRadio1LocalVapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenExtensionControllerExtenderProfileWifiRadio2Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "band"
	if _, ok := i["band"]; ok {
		result["band"] = flattenExtensionControllerExtenderProfileWifiRadio2Band(i["band"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtensionControllerExtenderProfileWifiRadio2Status(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "operating_standard"
	if _, ok := i["operating-standard"]; ok {
		result["operating_standard"] = flattenExtensionControllerExtenderProfileWifiRadio2OperatingStandard(i["operating-standard"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "guard_interval"
	if _, ok := i["guard-interval"]; ok {
		result["guard_interval"] = flattenExtensionControllerExtenderProfileWifiRadio2GuardInterval(i["guard-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "channel"
	if _, ok := i["channel"]; ok {
		result["channel"] = flattenExtensionControllerExtenderProfileWifiRadio2Channel(i["channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bandwidth"
	if _, ok := i["bandwidth"]; ok {
		result["bandwidth"] = flattenExtensionControllerExtenderProfileWifiRadio2Bandwidth(i["bandwidth"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "power_level"
	if _, ok := i["power-level"]; ok {
		result["power_level"] = flattenExtensionControllerExtenderProfileWifiRadio2PowerLevel(i["power-level"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := i["beacon-interval"]; ok {
		result["beacon_interval"] = flattenExtensionControllerExtenderProfileWifiRadio2BeaconInterval(i["beacon-interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "n80211d"
	if _, ok := i["80211d"]; ok {
		result["n80211d"] = flattenExtensionControllerExtenderProfileWifiRadio280211D(i["80211d"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_clients"
	if _, ok := i["max-clients"]; ok {
		result["max_clients"] = flattenExtensionControllerExtenderProfileWifiRadio2MaxClients(i["max-clients"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "extension_channel"
	if _, ok := i["extension-channel"]; ok {
		result["extension_channel"] = flattenExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(i["extension-channel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := i["bss-color-mode"]; ok {
		result["bss_color_mode"] = flattenExtensionControllerExtenderProfileWifiRadio2BssColorMode(i["bss-color-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bss_color"
	if _, ok := i["bss-color"]; ok {
		result["bss_color"] = flattenExtensionControllerExtenderProfileWifiRadio2BssColor(i["bss-color"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := i["lan-ext-vap"]; ok {
		result["lan_ext_vap"] = flattenExtensionControllerExtenderProfileWifiRadio2LanExtVap(i["lan-ext-vap"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "local_vaps"
	if _, ok := i["local-vaps"]; ok {
		result["local_vaps"] = flattenExtensionControllerExtenderProfileWifiRadio2LocalVaps(i["local-vaps"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderProfileWifiRadio2Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Band(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2OperatingStandard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2GuardInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Channel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2Bandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2PowerLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio2BeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio280211D(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2MaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2BssColorMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2BssColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenExtensionControllerExtenderProfileWifiRadio2LanExtVap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileWifiRadio2LocalVaps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtensionControllerExtenderProfileWifiRadio2LocalVapsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtensionControllerExtenderProfileWifiRadio2LocalVapsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenExtensionControllerExtenderProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerExtenderProfileId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("model", flattenExtensionControllerExtenderProfileModel(o["model"], d, "model", sv)); err != nil {
		if !fortiAPIPatch(o["model"]) {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("extension", flattenExtensionControllerExtenderProfileExtension(o["extension"], d, "extension", sv)); err != nil {
		if !fortiAPIPatch(o["extension"]) {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtensionControllerExtenderProfileAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenExtensionControllerExtenderProfileLoginPasswordChange(o["login-password-change"], d, "login_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["login-password-change"]) {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtensionControllerExtenderProfileEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-bandwidth"]) {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtensionControllerExtenderProfileBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-limit"]) {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("cellular", flattenExtensionControllerExtenderProfileCellular(o["cellular"], d, "cellular", sv)); err != nil {
			if !fortiAPIPatch(o["cellular"]) {
				return fmt.Errorf("Error reading cellular: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cellular"); ok {
			if err = d.Set("cellular", flattenExtensionControllerExtenderProfileCellular(o["cellular"], d, "cellular", sv)); err != nil {
				if !fortiAPIPatch(o["cellular"]) {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("lan_extension", flattenExtensionControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension", sv)); err != nil {
			if !fortiAPIPatch(o["lan-extension"]) {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenExtensionControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension", sv)); err != nil {
				if !fortiAPIPatch(o["lan-extension"]) {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("wifi", flattenExtensionControllerExtenderProfileWifi(o["wifi"], d, "wifi", sv)); err != nil {
			if !fortiAPIPatch(o["wifi"]) {
				return fmt.Errorf("Error reading wifi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wifi"); ok {
			if err = d.Set("wifi", flattenExtensionControllerExtenderProfileWifi(o["wifi"], d, "wifi", sv)); err != nil {
				if !fortiAPIPatch(o["wifi"]) {
					return fmt.Errorf("Error reading wifi: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtensionControllerExtenderProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtensionControllerExtenderProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLoginPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileBandwidthLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellular(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularDataplan(d, i["dataplan"], pre_append, sv)
	} else {
		result["dataplan"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "controller_report"
	if _, ok := d.GetOk(pre_append); ok {
		result["controller-report"], _ = expandExtensionControllerExtenderProfileCellularControllerReport(d, i["controller_report"], pre_append, sv)
	} else {
		result["controller-report"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sms_notification"
	if _, ok := d.GetOk(pre_append); ok {
		result["sms-notification"], _ = expandExtensionControllerExtenderProfileCellularSmsNotification(d, i["sms_notification"], pre_append, sv)
	} else {
		result["sms-notification"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "modem1"
	if _, ok := d.GetOk(pre_append); ok {
		result["modem1"], _ = expandExtensionControllerExtenderProfileCellularModem1(d, i["modem1"], pre_append, sv)
	} else {
		result["modem1"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "modem2"
	if _, ok := d.GetOk(pre_append); ok {
		result["modem2"], _ = expandExtensionControllerExtenderProfileCellularModem2(d, i["modem2"], pre_append, sv)
	} else {
		result["modem2"] = make([]string, 0)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandExtensionControllerExtenderProfileCellularDataplanName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularDataplanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtensionControllerExtenderProfileCellularControllerReportStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["interval"], _ = expandExtensionControllerExtenderProfileCellularControllerReportInterval(d, i["interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal-threshold"], _ = expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(d, i["signal_threshold"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok {
		result["alert"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlert(d, i["alert"], pre_append, sv)
	} else {
		result["alert"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok {
		result["receiver"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d, i["receiver"], pre_append, sv)
	} else {
		result["receiver"] = make([]string, 0)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok {
		result["system-reboot"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d, i["system_reboot"], pre_append, sv)
	}
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok {
		result["data-exhausted"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d, i["data_exhausted"], pre_append, sv)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["session-disconnect"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d, i["session_disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok {
		result["low-signal-strength"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d, i["low_signal_strength"], pre_append, sv)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok {
		result["os-image-fallback"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d, i["os_image_fallback"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d, i["mode_switch"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["fgt-backup-mode-switch"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d, i["fgt_backup_mode_switch"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["phone-number"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d, i["phone_number"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["phone-number"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["alert"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d, i["alert"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["alert"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-mode"], _ = expandExtensionControllerExtenderProfileCellularModem1RedundantMode(d, i["redundant_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-intf"], _ = expandExtensionControllerExtenderProfileCellularModem1RedundantIntf(d, i["redundant_intf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {
		result["conn-status"], _ = expandExtensionControllerExtenderProfileCellularModem1ConnStatus(d, i["conn_status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-sim"], _ = expandExtensionControllerExtenderProfileCellularModem1DefaultSim(d, i["default_sim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {
		result["gps"], _ = expandExtensionControllerExtenderProfileCellularModem1Gps(d, i["gps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim1Pin(d, i["sim1_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim2Pin(d, i["sim2_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {
		result["preferred-carrier"], _ = expandExtensionControllerExtenderProfileCellularModem1PreferredCarrier(d, i["preferred_carrier"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-switch"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitch(d, i["auto_switch"], pre_append, sv)
	} else {
		result["auto-switch"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["multiple-PDN"], _ = expandExtensionControllerExtenderProfileCellularModem1MultiplePdn(d, i["multiple_pdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn1-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(d, i["pdn1_dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn2-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(d, i["pdn2_dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn3-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(d, i["pdn3_dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn4-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(d, i["pdn4_dataplan"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Gps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-threshold"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-period"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(d, i["signal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(d, i["dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-time"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-timer"], _ = expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1MultiplePdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn1Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn2Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn3Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem1Pdn4Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-mode"], _ = expandExtensionControllerExtenderProfileCellularModem2RedundantMode(d, i["redundant_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-intf"], _ = expandExtensionControllerExtenderProfileCellularModem2RedundantIntf(d, i["redundant_intf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {
		result["conn-status"], _ = expandExtensionControllerExtenderProfileCellularModem2ConnStatus(d, i["conn_status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-sim"], _ = expandExtensionControllerExtenderProfileCellularModem2DefaultSim(d, i["default_sim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {
		result["gps"], _ = expandExtensionControllerExtenderProfileCellularModem2Gps(d, i["gps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim1Pin(d, i["sim1_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim2Pin(d, i["sim2_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin-code"], _ = expandExtensionControllerExtenderProfileCellularModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {
		result["preferred-carrier"], _ = expandExtensionControllerExtenderProfileCellularModem2PreferredCarrier(d, i["preferred_carrier"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-switch"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitch(d, i["auto_switch"], pre_append, sv)
	} else {
		result["auto-switch"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "multiple_pdn"
	if _, ok := d.GetOk(pre_append); ok {
		result["multiple-PDN"], _ = expandExtensionControllerExtenderProfileCellularModem2MultiplePdn(d, i["multiple_pdn"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn1_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn1-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(d, i["pdn1_dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn2_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn2-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(d, i["pdn2_dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn3_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn3-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(d, i["pdn3_dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "pdn4_dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["pdn4-dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(d, i["pdn4_dataplan"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Gps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-threshold"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-period"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(d, i["signal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(d, i["dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-time"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-timer"], _ = expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2MultiplePdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn1Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn2Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn3Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileCellularModem2Pdn4Dataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := d.GetOk(pre_append); ok {
		result["link-loadbalance"], _ = expandExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(d, i["link_loadbalance"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipsec-tunnel"], _ = expandExtensionControllerExtenderProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok {
		result["backhaul-interface"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["backhaul-ip"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul"
	if _, ok := d.GetOk(pre_append); ok {
		result["backhaul"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaul(d, i["backhaul"], pre_append, sv)
	} else {
		result["backhaul"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "downlinks"
	if _, ok := d.GetOk(pre_append); ok {
		result["downlinks"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinks(d, i["downlinks"], pre_append, sv)
	} else {
		result["downlinks"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "traffic_split_services"
	if _, ok := d.GetOk(pre_append); ok {
		result["traffic-split-services"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d, i["traffic_split_services"], pre_append, sv)
	} else {
		result["traffic-split-services"] = make([]string, 0)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionLinkLoadbalance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulRole(d, i["role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandExtensionControllerExtenderProfileLanExtensionBackhaulWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionBackhaulWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vap"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksVap(d, i["vap"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vap"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pvid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pvid"], _ = expandExtensionControllerExtenderProfileLanExtensionDownlinksPvid(d, i["pvid"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["pvid"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksVap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionDownlinksPvid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vsdb"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vsdb"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(d, i["vsdb"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(d, i["address"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["address"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["service"], _ = expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(d, i["service"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["service"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesVsdb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileLanExtensionTrafficSplitServicesService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "country"
	if _, ok := d.GetOk(pre_append); ok {
		result["country"], _ = expandExtensionControllerExtenderProfileWifiCountry(d, i["country"], pre_append, sv)
	}
	pre_append = pre + ".0." + "radio_1"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-1"], _ = expandExtensionControllerExtenderProfileWifiRadio1(d, i["radio_1"], pre_append, sv)
	} else {
		result["radio-1"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "radio_2"
	if _, ok := d.GetOk(pre_append); ok {
		result["radio-2"], _ = expandExtensionControllerExtenderProfileWifiRadio2(d, i["radio_2"], pre_append, sv)
	} else {
		result["radio-2"] = make([]string, 0)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandExtensionControllerExtenderProfileWifiRadio1Mode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandExtensionControllerExtenderProfileWifiRadio1Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtensionControllerExtenderProfileWifiRadio1Status(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "operating_standard"
	if _, ok := d.GetOk(pre_append); ok {
		result["operating-standard"], _ = expandExtensionControllerExtenderProfileWifiRadio1OperatingStandard(d, i["operating_standard"], pre_append, sv)
	}
	pre_append = pre + ".0." + "guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["guard-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio1GuardInterval(d, i["guard_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandExtensionControllerExtenderProfileWifiRadio1Channel(d, i["channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth"], _ = expandExtensionControllerExtenderProfileWifiRadio1Bandwidth(d, i["bandwidth"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandExtensionControllerExtenderProfileWifiRadio1PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio1BeaconInterval(d, i["beacon_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok {
		result["80211d"], _ = expandExtensionControllerExtenderProfileWifiRadio180211D(d, i["n80211d"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandExtensionControllerExtenderProfileWifiRadio1MaxClients(d, i["max_clients"], pre_append, sv)
	}
	pre_append = pre + ".0." + "extension_channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["extension-channel"], _ = expandExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(d, i["extension_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color-mode"], _ = expandExtensionControllerExtenderProfileWifiRadio1BssColorMode(d, i["bss_color_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandExtensionControllerExtenderProfileWifiRadio1BssColor(d, i["bss_color"], pre_append, sv)
	}
	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := d.GetOk(pre_append); ok {
		result["lan-ext-vap"], _ = expandExtensionControllerExtenderProfileWifiRadio1LanExtVap(d, i["lan_ext_vap"], pre_append, sv)
	}
	pre_append = pre + ".0." + "local_vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-vaps"], _ = expandExtensionControllerExtenderProfileWifiRadio1LocalVaps(d, i["local_vaps"], pre_append, sv)
	} else {
		result["local-vaps"] = make([]string, 0)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1OperatingStandard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1GuardInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1Bandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio180211D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1MaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1ExtensionChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BssColorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1BssColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LanExtVap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LocalVaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandExtensionControllerExtenderProfileWifiRadio1LocalVapsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiRadio1LocalVapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode"], _ = expandExtensionControllerExtenderProfileWifiRadio2Mode(d, i["mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "band"
	if _, ok := d.GetOk(pre_append); ok {
		result["band"], _ = expandExtensionControllerExtenderProfileWifiRadio2Band(d, i["band"], pre_append, sv)
	}
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtensionControllerExtenderProfileWifiRadio2Status(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "operating_standard"
	if _, ok := d.GetOk(pre_append); ok {
		result["operating-standard"], _ = expandExtensionControllerExtenderProfileWifiRadio2OperatingStandard(d, i["operating_standard"], pre_append, sv)
	}
	pre_append = pre + ".0." + "guard_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["guard-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio2GuardInterval(d, i["guard_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["channel"], _ = expandExtensionControllerExtenderProfileWifiRadio2Channel(d, i["channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bandwidth"
	if _, ok := d.GetOk(pre_append); ok {
		result["bandwidth"], _ = expandExtensionControllerExtenderProfileWifiRadio2Bandwidth(d, i["bandwidth"], pre_append, sv)
	}
	pre_append = pre + ".0." + "power_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["power-level"], _ = expandExtensionControllerExtenderProfileWifiRadio2PowerLevel(d, i["power_level"], pre_append, sv)
	}
	pre_append = pre + ".0." + "beacon_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["beacon-interval"], _ = expandExtensionControllerExtenderProfileWifiRadio2BeaconInterval(d, i["beacon_interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "n80211d"
	if _, ok := d.GetOk(pre_append); ok {
		result["80211d"], _ = expandExtensionControllerExtenderProfileWifiRadio280211D(d, i["n80211d"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_clients"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-clients"], _ = expandExtensionControllerExtenderProfileWifiRadio2MaxClients(d, i["max_clients"], pre_append, sv)
	}
	pre_append = pre + ".0." + "extension_channel"
	if _, ok := d.GetOk(pre_append); ok {
		result["extension-channel"], _ = expandExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(d, i["extension_channel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color-mode"], _ = expandExtensionControllerExtenderProfileWifiRadio2BssColorMode(d, i["bss_color_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bss_color"
	if _, ok := d.GetOk(pre_append); ok {
		result["bss-color"], _ = expandExtensionControllerExtenderProfileWifiRadio2BssColor(d, i["bss_color"], pre_append, sv)
	}
	pre_append = pre + ".0." + "lan_ext_vap"
	if _, ok := d.GetOk(pre_append); ok {
		result["lan-ext-vap"], _ = expandExtensionControllerExtenderProfileWifiRadio2LanExtVap(d, i["lan_ext_vap"], pre_append, sv)
	}
	pre_append = pre + ".0." + "local_vaps"
	if _, ok := d.GetOk(pre_append); ok {
		result["local-vaps"], _ = expandExtensionControllerExtenderProfileWifiRadio2LocalVaps(d, i["local_vaps"], pre_append, sv)
	} else {
		result["local-vaps"] = make([]string, 0)
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Band(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2OperatingStandard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2GuardInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Channel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2Bandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2PowerLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2BeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio280211D(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2MaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2ExtensionChannel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2BssColorMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2BssColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2LanExtVap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2LocalVaps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandExtensionControllerExtenderProfileWifiRadio2LocalVapsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtensionControllerExtenderProfileWifiRadio2LocalVapsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtensionControllerExtenderProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandExtensionControllerExtenderProfileId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("model"); ok {
		t, err := expandExtensionControllerExtenderProfileModel(d, v, "model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	}

	if v, ok := d.GetOk("extension"); ok {
		t, err := expandExtensionControllerExtenderProfileExtension(d, v, "extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandExtensionControllerExtenderProfileAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	} else if d.HasChange("allowaccess") {
		obj["allowaccess"] = nil
	}

	if v, ok := d.GetOk("login_password_change"); ok {
		t, err := expandExtensionControllerExtenderProfileLoginPasswordChange(d, v, "login_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok {
		t, err := expandExtensionControllerExtenderProfileLoginPassword(d, v, "login_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	} else if d.HasChange("login_password") {
		obj["login-password"] = nil
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok {
		t, err := expandExtensionControllerExtenderProfileEnforceBandwidth(d, v, "enforce_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok {
		t, err := expandExtensionControllerExtenderProfileBandwidthLimit(d, v, "bandwidth_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("cellular"); ok {
		t, err := expandExtensionControllerExtenderProfileCellular(d, v, "cellular", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cellular"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok {
		t, err := expandExtensionControllerExtenderProfileLanExtension(d, v, "lan_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	if v, ok := d.GetOk("wifi"); ok {
		t, err := expandExtensionControllerExtenderProfileWifi(d, v, "wifi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi"] = t
		}
	}

	return &obj, nil
}
