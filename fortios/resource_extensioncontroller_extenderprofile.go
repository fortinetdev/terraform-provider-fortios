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

func resourceExtensionControllerExtenderProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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
	return v
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

func flattenExtensionControllerExtenderProfileLoginPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileBandwidthLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenExtensionControllerExtenderProfileCellularDataplanName(i["name"], d, pre_append, sv)
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
	return v
}

func flattenExtensionControllerExtenderProfileCellularControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := i["phone-number"]; ok {
			tmp["phone_number"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(i["phone-number"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			tmp["alert"] = flattenExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(i["alert"], d, pre_append, sv)
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
		result["sim1_pin_code"] = flattenExtensionControllerExtenderProfileCellularModem1Sim1PinCode(i["sim1-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenExtensionControllerExtenderProfileCellularModem1Sim2PinCode(i["sim2-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem1PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem1AutoSwitch(i["auto-switch"], d, pre_append, sv)
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
	return v
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

func flattenExtensionControllerExtenderProfileCellularModem1Sim1PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1Sim2PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
		result["sim1_pin_code"] = flattenExtensionControllerExtenderProfileCellularModem2Sim1PinCode(i["sim1-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenExtensionControllerExtenderProfileCellularModem2Sim2PinCode(i["sim2-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtensionControllerExtenderProfileCellularModem2PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtensionControllerExtenderProfileCellularModem2AutoSwitch(i["auto-switch"], d, pre_append, sv)
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
	return v
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

func flattenExtensionControllerExtenderProfileCellularModem2Sim1PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2Sim2PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
	return v
}

func flattenExtensionControllerExtenderProfileCellularModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			tmp["port"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role"
		if _, ok := i["role"]; ok {
			tmp["role"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulRole(i["role"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {
			tmp["weight"] = flattenExtensionControllerExtenderProfileLanExtensionBackhaulWeight(i["weight"], d, pre_append, sv)
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

	if err = d.Set("login_password", flattenExtensionControllerExtenderProfileLoginPassword(o["login-password"], d, "login_password", sv)); err != nil {
		if !fortiAPIPatch(o["login-password"]) {
			return fmt.Errorf("Error reading login_password: %v", err)
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
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "phone_number"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["phone-number"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverPhoneNumber(d, i["phone_number"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["alert"], _ = expandExtensionControllerExtenderProfileCellularSmsNotificationReceiverAlert(d, i["alert"], pre_append, sv)
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

	return &obj, nil
}
