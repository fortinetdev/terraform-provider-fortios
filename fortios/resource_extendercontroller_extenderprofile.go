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

func resourceExtenderControllerExtenderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerExtenderProfileCreate,
		Read:   resourceExtenderControllerExtenderProfileRead,
		Update: resourceExtenderControllerExtenderProfileUpdate,
		Delete: resourceExtenderControllerExtenderProfileDelete,

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
				Computed: true,
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
										Computed:     true,
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
													Computed:     true,
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
													Computed:     true,
												},
												"alert": &schema.Schema{
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
										Computed:     true,
									},
									"conn_status": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
									},
									"sim2_pin_code": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 27),
										Optional:     true,
									},
									"preferred_carrier": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
										Computed:     true,
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
													Computed: true,
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
										Computed:     true,
									},
									"conn_status": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
									},
									"sim2_pin_code": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 27),
										Optional:     true,
									},
									"preferred_carrier": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
										Computed:     true,
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
													Computed: true,
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
							Computed:     true,
						},
						"backhaul_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"backhaul_ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
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
										Computed:     true,
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

func resourceExtenderControllerExtenderProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtenderControllerExtenderProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtenderProfile resource while getting object: %v", err)
	}

	o, err := c.CreateExtenderControllerExtenderProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtenderProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtenderProfile")
	}

	return resourceExtenderControllerExtenderProfileRead(d, m)
}

func resourceExtenderControllerExtenderProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtenderControllerExtenderProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtenderProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateExtenderControllerExtenderProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtenderProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtenderProfile")
	}

	return resourceExtenderControllerExtenderProfileRead(d, m)
}

func resourceExtenderControllerExtenderProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtenderControllerExtenderProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerExtenderProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtenderControllerExtenderProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtenderProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerExtenderProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtenderProfile resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerExtenderProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileExtension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLoginPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileBandwidthLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellular(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtenderProfileCellularDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "controller_report"
	if _, ok := i["controller-report"]; ok {
		result["controller_report"] = flattenExtenderControllerExtenderProfileCellularControllerReport(i["controller-report"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sms_notification"
	if _, ok := i["sms-notification"]; ok {
		result["sms_notification"] = flattenExtenderControllerExtenderProfileCellularSmsNotification(i["sms-notification"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "modem1"
	if _, ok := i["modem1"]; ok {
		result["modem1"] = flattenExtenderControllerExtenderProfileCellularModem1(i["modem1"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "modem2"
	if _, ok := i["modem2"]; ok {
		result["modem2"] = flattenExtenderControllerExtenderProfileCellularModem2(i["modem2"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtenderControllerExtenderProfileCellularDataplanName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtenderControllerExtenderProfileCellularDataplanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularControllerReport(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderProfileCellularControllerReportStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenExtenderControllerExtenderProfileCellularControllerReportInterval(i["interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(i["signal-threshold"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularControllerReportStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularControllerReportInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotification(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "alert"
	if _, ok := i["alert"]; ok {
		result["alert"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlert(i["alert"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "receiver"
	if _, ok := i["receiver"]; ok {
		result["receiver"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiver(i["receiver"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlert(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := i["system-reboot"]; ok {
		result["system_reboot"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(i["system-reboot"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := i["data-exhausted"]; ok {
		result["data_exhausted"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(i["data-exhausted"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := i["session-disconnect"]; ok {
		result["session_disconnect"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(i["session-disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := i["low-signal-strength"]; ok {
		result["low_signal_strength"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(i["low-signal-strength"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := i["os-image-fallback"]; ok {
		result["os_image_fallback"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(i["os-image-fallback"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "mode_switch"
	if _, ok := i["mode-switch"]; ok {
		result["mode_switch"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(i["mode-switch"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := i["fgt-backup-mode-switch"]; ok {
		result["fgt_backup_mode_switch"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(i["fgt-backup-mode-switch"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiver(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if cur_v, ok := i["phone-number"]; ok {
			tmp["phone_number"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if cur_v, ok := i["alert"]; ok {
			tmp["alert"] = flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtenderControllerExtenderProfileCellularModem1RedundantMode(i["redundant-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtenderControllerExtenderProfileCellularModem1RedundantIntf(i["redundant-intf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtenderControllerExtenderProfileCellularModem1ConnStatus(i["conn-status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtenderControllerExtenderProfileCellularModem1DefaultSim(i["default-sim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtenderControllerExtenderProfileCellularModem1Gps(i["gps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtenderControllerExtenderProfileCellularModem1Sim1Pin(i["sim1-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtenderControllerExtenderProfileCellularModem1Sim2Pin(i["sim2-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenExtenderControllerExtenderProfileCellularModem1Sim1PinCode(i["sim1-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenExtenderControllerExtenderProfileCellularModem1Sim2PinCode(i["sim2-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtenderControllerExtenderProfileCellularModem1PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitch(i["auto-switch"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1Gps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1Sim1PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1Sim2PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(i["signal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtenderControllerExtenderProfileCellularModem2RedundantMode(i["redundant-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtenderControllerExtenderProfileCellularModem2RedundantIntf(i["redundant-intf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtenderControllerExtenderProfileCellularModem2ConnStatus(i["conn-status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtenderControllerExtenderProfileCellularModem2DefaultSim(i["default-sim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtenderControllerExtenderProfileCellularModem2Gps(i["gps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtenderControllerExtenderProfileCellularModem2Sim1Pin(i["sim1-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtenderControllerExtenderProfileCellularModem2Sim2Pin(i["sim2-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenExtenderControllerExtenderProfileCellularModem2Sim1PinCode(i["sim1-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenExtenderControllerExtenderProfileCellularModem2Sim2PinCode(i["sim2-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtenderControllerExtenderProfileCellularModem2PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitch(i["auto-switch"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2Gps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2Sim1PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2Sim2PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(i["signal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := i["link-loadbalance"]; ok {
		result["link_loadbalance"] = flattenExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(i["link-loadbalance"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {
		result["ipsec_tunnel"] = flattenExtenderControllerExtenderProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {
		result["backhaul_interface"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {
		result["backhaul_ip"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul"
	if _, ok := i["backhaul"]; ok {
		result["backhaul"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaul(i["backhaul"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaul(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaulName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaulPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if cur_v, ok := i["role"]; ok {
			tmp["role"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaulRole(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if cur_v, ok := i["weight"]; ok {
			tmp["weight"] = flattenExtenderControllerExtenderProfileLanExtensionBackhaulWeight(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaulName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaulPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaulRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfileLanExtensionBackhaulWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtenderControllerExtenderProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenExtenderControllerExtenderProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtenderControllerExtenderProfileId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("model", flattenExtenderControllerExtenderProfileModel(o["model"], d, "model", sv)); err != nil {
		if !fortiAPIPatch(o["model"]) {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("extension", flattenExtenderControllerExtenderProfileExtension(o["extension"], d, "extension", sv)); err != nil {
		if !fortiAPIPatch(o["extension"]) {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtenderControllerExtenderProfileAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenExtenderControllerExtenderProfileLoginPasswordChange(o["login-password-change"], d, "login_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["login-password-change"]) {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("login_password", flattenExtenderControllerExtenderProfileLoginPassword(o["login-password"], d, "login_password", sv)); err != nil {
		if !fortiAPIPatch(o["login-password"]) {
			return fmt.Errorf("Error reading login_password: %v", err)
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtenderControllerExtenderProfileEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-bandwidth"]) {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtenderControllerExtenderProfileBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-limit"]) {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("cellular", flattenExtenderControllerExtenderProfileCellular(o["cellular"], d, "cellular", sv)); err != nil {
			if !fortiAPIPatch(o["cellular"]) {
				return fmt.Errorf("Error reading cellular: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cellular"); ok {
			if err = d.Set("cellular", flattenExtenderControllerExtenderProfileCellular(o["cellular"], d, "cellular", sv)); err != nil {
				if !fortiAPIPatch(o["cellular"]) {
					return fmt.Errorf("Error reading cellular: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("lan_extension", flattenExtenderControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension", sv)); err != nil {
			if !fortiAPIPatch(o["lan-extension"]) {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenExtenderControllerExtenderProfileLanExtension(o["lan-extension"], d, "lan_extension", sv)); err != nil {
				if !fortiAPIPatch(o["lan-extension"]) {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtenderControllerExtenderProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtenderControllerExtenderProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLoginPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileBandwidthLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellular(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtenderControllerExtenderProfileCellularDataplan(d, i["dataplan"], pre_append, sv)
	} else {
		result["dataplan"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "controller_report"
	if _, ok := d.GetOk(pre_append); ok {
		result["controller-report"], _ = expandExtenderControllerExtenderProfileCellularControllerReport(d, i["controller_report"], pre_append, sv)
	} else {
		result["controller-report"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sms_notification"
	if _, ok := d.GetOk(pre_append); ok {
		result["sms-notification"], _ = expandExtenderControllerExtenderProfileCellularSmsNotification(d, i["sms_notification"], pre_append, sv)
	} else {
		result["sms-notification"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "modem1"
	if _, ok := d.GetOk(pre_append); ok {
		result["modem1"], _ = expandExtenderControllerExtenderProfileCellularModem1(d, i["modem1"], pre_append, sv)
	} else {
		result["modem1"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "modem2"
	if _, ok := d.GetOk(pre_append); ok {
		result["modem2"], _ = expandExtenderControllerExtenderProfileCellularModem2(d, i["modem2"], pre_append, sv)
	} else {
		result["modem2"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandExtenderControllerExtenderProfileCellularDataplanName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularDataplanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularControllerReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtenderControllerExtenderProfileCellularControllerReportStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["interval"], _ = expandExtenderControllerExtenderProfileCellularControllerReportInterval(d, i["interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal-threshold"], _ = expandExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(d, i["signal_threshold"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularControllerReportStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularControllerReportInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "alert"
	if _, ok := d.GetOk(pre_append); ok {
		result["alert"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlert(d, i["alert"], pre_append, sv)
	} else {
		result["alert"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "receiver"
	if _, ok := d.GetOk(pre_append); ok {
		result["receiver"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d, i["receiver"], pre_append, sv)
	} else {
		result["receiver"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "system_reboot"
	if _, ok := d.GetOk(pre_append); ok {
		result["system-reboot"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d, i["system_reboot"], pre_append, sv)
	}
	pre_append = pre + ".0." + "data_exhausted"
	if _, ok := d.GetOk(pre_append); ok {
		result["data-exhausted"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d, i["data_exhausted"], pre_append, sv)
	}
	pre_append = pre + ".0." + "session_disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["session-disconnect"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d, i["session_disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "low_signal_strength"
	if _, ok := d.GetOk(pre_append); ok {
		result["low-signal-strength"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d, i["low_signal_strength"], pre_append, sv)
	}
	pre_append = pre + ".0." + "os_image_fallback"
	if _, ok := d.GetOk(pre_append); ok {
		result["os-image-fallback"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d, i["os_image_fallback"], pre_append, sv)
	}
	pre_append = pre + ".0." + "mode_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["mode-switch"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d, i["mode_switch"], pre_append, sv)
	}
	pre_append = pre + ".0." + "fgt_backup_mode_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["fgt-backup-mode-switch"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d, i["fgt_backup_mode_switch"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertSystemReboot(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertDataExhausted(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertSessionDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertLowSignalStrength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertOsImageFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertModeSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationAlertFgtBackupModeSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationReceiver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["phone-number"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d, i["phone_number"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["alert"], _ = expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(d, i["alert"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularSmsNotificationReceiverAlert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-mode"], _ = expandExtenderControllerExtenderProfileCellularModem1RedundantMode(d, i["redundant_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-intf"], _ = expandExtenderControllerExtenderProfileCellularModem1RedundantIntf(d, i["redundant_intf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {
		result["conn-status"], _ = expandExtenderControllerExtenderProfileCellularModem1ConnStatus(d, i["conn_status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-sim"], _ = expandExtenderControllerExtenderProfileCellularModem1DefaultSim(d, i["default_sim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {
		result["gps"], _ = expandExtenderControllerExtenderProfileCellularModem1Gps(d, i["gps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin"], _ = expandExtenderControllerExtenderProfileCellularModem1Sim1Pin(d, i["sim1_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin"], _ = expandExtenderControllerExtenderProfileCellularModem1Sim2Pin(d, i["sim2_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin-code"], _ = expandExtenderControllerExtenderProfileCellularModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin-code"], _ = expandExtenderControllerExtenderProfileCellularModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {
		result["preferred-carrier"], _ = expandExtenderControllerExtenderProfileCellularModem1PreferredCarrier(d, i["preferred_carrier"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-switch"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitch(d, i["auto_switch"], pre_append, sv)
	} else {
		result["auto-switch"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1Gps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-threshold"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-period"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(d, i["signal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(d, i["dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-time"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-timer"], _ = expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-mode"], _ = expandExtenderControllerExtenderProfileCellularModem2RedundantMode(d, i["redundant_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-intf"], _ = expandExtenderControllerExtenderProfileCellularModem2RedundantIntf(d, i["redundant_intf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {
		result["conn-status"], _ = expandExtenderControllerExtenderProfileCellularModem2ConnStatus(d, i["conn_status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-sim"], _ = expandExtenderControllerExtenderProfileCellularModem2DefaultSim(d, i["default_sim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {
		result["gps"], _ = expandExtenderControllerExtenderProfileCellularModem2Gps(d, i["gps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin"], _ = expandExtenderControllerExtenderProfileCellularModem2Sim1Pin(d, i["sim1_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin"], _ = expandExtenderControllerExtenderProfileCellularModem2Sim2Pin(d, i["sim2_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin-code"], _ = expandExtenderControllerExtenderProfileCellularModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin-code"], _ = expandExtenderControllerExtenderProfileCellularModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {
		result["preferred-carrier"], _ = expandExtenderControllerExtenderProfileCellularModem2PreferredCarrier(d, i["preferred_carrier"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-switch"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitch(d, i["auto_switch"], pre_append, sv)
	} else {
		result["auto-switch"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2Gps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-threshold"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-period"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(d, i["signal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(d, i["dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-time"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-timer"], _ = expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileCellularModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "link_loadbalance"
	if _, ok := d.GetOk(pre_append); ok {
		result["link-loadbalance"], _ = expandExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(d, i["link_loadbalance"], pre_append, sv)
	}
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["ipsec-tunnel"], _ = expandExtenderControllerExtenderProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok {
		result["backhaul-interface"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["backhaul-ip"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul"
	if _, ok := d.GetOk(pre_append); ok {
		result["backhaul"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaul(d, i["backhaul"], pre_append, sv)
	} else {
		result["backhaul"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileLanExtensionLinkLoadbalance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaul(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaulName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaulPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaulRole(d, i["role"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["weight"], _ = expandExtenderControllerExtenderProfileLanExtensionBackhaulWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaulName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaulPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaulRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfileLanExtensionBackhaulWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtenderControllerExtenderProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtenderControllerExtenderProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandExtenderControllerExtenderProfileId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("model"); ok {
		t, err := expandExtenderControllerExtenderProfileModel(d, v, "model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	}

	if v, ok := d.GetOk("extension"); ok {
		t, err := expandExtenderControllerExtenderProfileExtension(d, v, "extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandExtenderControllerExtenderProfileAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok {
		t, err := expandExtenderControllerExtenderProfileLoginPasswordChange(d, v, "login_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok {
		t, err := expandExtenderControllerExtenderProfileLoginPassword(d, v, "login_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok {
		t, err := expandExtenderControllerExtenderProfileEnforceBandwidth(d, v, "enforce_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok {
		t, err := expandExtenderControllerExtenderProfileBandwidthLimit(d, v, "bandwidth_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("cellular"); ok {
		t, err := expandExtenderControllerExtenderProfileCellular(d, v, "cellular", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cellular"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok {
		t, err := expandExtenderControllerExtenderProfileLanExtension(d, v, "lan_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	return &obj, nil
}
